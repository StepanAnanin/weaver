package weaver

type settings struct {
	// If true, each incoming request will be logged in terminal.
	// By default - true.
	LogIncomingRequests bool
	// If true, then CORS headers won't be apply on preprocessing (request.Preprocessing).
	// By default - false.
	DisableCORS bool
	// Default value of Access-Control-Allow-Origin header.
	// By default - "*".
	DefaultOrigin string
	// If true, then OPTIONS requests won't be canceled on preprocessing (request.Preprocessing).
	// By default - false.
	//
	// # If `DisableCORS` if false, then don't forget to pass OPTIONS method in array of allowed methods.
	// (or u'll get 405 error response on preprocessing)
	PassOptionsRequestsOnPreprocessing bool
}

var Settings *settings = &settings{
	LogIncomingRequests:                true,
	DisableCORS:                        false,
	DefaultOrigin:                      "*",
	PassOptionsRequestsOnPreprocessing: false,
}
