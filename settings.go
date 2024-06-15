package weaver

type settings struct {
	// If true, each incoming request will be logged in terminal.
	// By default - true.
	LogIncomingRequests bool
	// Access-Control-Allow-Origin header value.
	// By default - "*".
	AccessControlAllowOrigin string
}

var Settings *settings = &settings{
	LogIncomingRequests:      true,
	AccessControlAllowOrigin: "*",
}
