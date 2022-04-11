package use_context

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	//reimplement
	"github.com/hotsnow199/go_util/context"
)

func handleSearch(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	// check timeout query
	timeout, err := time.ParseDuration(r.FormValue("timeout"))
	if err == nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}

	if deadline, ok := ctx.Deadline(); ok {
		now := time.Now()
		fmt.Println("deadline_at:", deadline, "deadline_duration:", deadline.Sub(now))
	}

	defer cancel() // cancel ctx as soon as handleSearch returns, free resources
	//check search query
	query := r.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	// store user ip
	userIP, err := GetIPFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx = NewUserIPContext(ctx, userIP)
	start := time.Now()
	results, err := GOOGLESearch(ctx, query)

	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := resultsTemplate.Execute(w, struct {
		Results          Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Println(err)
		return
	}

}

func UseSearchGoogle() {
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))
