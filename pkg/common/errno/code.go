package errno

var (
	// OK represents a successful request.
	OK = &Errno{Code: 0, Message: "OK"}
	ErrBind = &Errno{Code: 10002, Message: "Error occurred when binding request."}
	ErrUserNotFound = &Errno{Code: 10003, Message: "User not found."}
	ErrPasswordIncorrect = &Errno{Code: 10004, Message: "Password incorrect."}
	ErrToken = &Errno{Code: 10005, Message: "Error Token."}

	// InternalServerError represents all unknown server-side errors.
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
)
