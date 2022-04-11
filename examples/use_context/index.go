package use_context

/**
 * open browser try : http://localhost:8080/search?q=${keyword}&timeout=${time}
 *
 * exp: http://localhost:8080/search?q=golang&timeout=0.2s
 */
func UseContext() {
	go UseMockServer()
	UseSearchGoogle()
}
