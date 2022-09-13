package errno

var (
	// OK represents a successful request.
	OK = &Errno{Code: 0, Message: "OK"}

	// InternalServerError represents all unknown server-side errors.
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
)
