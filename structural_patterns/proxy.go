package structural_patterns

import "fmt"

type iApplication interface {
	handleRequest(url, method string) (int, string)
}

type application struct {}

func (app *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/user/create" && method == "POST" {
		return 201, "User created"
	}

	return 404, "Not found"
}

type appProxy struct {
	app iApplication
	maxAllowedRequest int
	rateLimit map[string]int
}

func newAppProxy() iApplication {
	return &appProxy{
		app: &application{},
		maxAllowedRequest: 2,
		rateLimit: make(map[string]int),
	}
}

func (appPrx *appProxy) handleRequest(url, method string) (int, string){
	appPrx.rateLimit[url]++

	if appPrx.rateLimit[url] > appPrx.maxAllowedRequest {
		return 403, "Not allowed"
	}

	return appPrx.app.handleRequest(url, method)
}

func ProxyPattern() {
	appStatusURL := "/app/status"
	createUserURL := "/user/create"
	
	server := newAppProxy()

	httpCode, body := server.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = server.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = server.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = server.handleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = server.handleRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

}