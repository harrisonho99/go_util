package use_context

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//reimplement
	"github.com/hotsnow199/go_util/context"
)

// Results is an ordered list of search results.
type Results []Result

// A Result contains the title and URL of a search result.
type Result struct {
	Title, URL string
}

const API_URL = "http://localhost:3000/search"

//new Version
func GOOGLESearch(ctx context.Context, query string) (Results, error) {

	// prepare search request
	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)

	// If ctx is carrying the user IP address, forward it to the server.
	// Google APIs use the user IP to distinguish server-initiated requests
	// from end-user requests.
	// if userIP, ok := UserIPFromContext(ctx); ok {
	// 	q.Set("userip", userIP.String())
	// }

	req.URL.RawQuery = q.Encode()

	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var data struct {
			ResponseData struct {
				Results []struct {
					TitleNoFormatting string
					URL               string
				}
			}
		}
		fmt.Println("----Error : ", err)
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}

		for _, res := range data.ResponseData.Results {
			results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
		}
		return nil
	})
	return results, err
}
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	c := make(chan error, 1)
	go func() {
		c <- f(do(ctx, req))
	}()
	select {
	//currently only timeout active,
	//cancel is used in this goroutine to clear timeout context when google server response
	case <-ctx.Done():
		return ctx.Err()
	case err := <-c:
		return err
	}
}

// helper func request with custom context
func do(ctx context.Context, req *http.Request) (*http.Response, error) {
	var client *http.Client
	if deadline, ok := ctx.Deadline(); ok {
		client = &http.Client{Timeout: deadline.Sub(time.Now())}
	} else {
		client = &http.Client{}

	}
	return client.Do(req)
}
