package errno

var (
	OK                   = &Errno{Status: 200, Code: "200", Message: "OK"}
	InternalServerError  = &Errno{Status: 500, Code: "InternalError", Message: "Internal server error."}
	NotFound             = &Errno{Status: 404, Code: "NotFound", Message: "The requested resource was not found."}
	BadRequest           = &Errno{Status: 400, Code: "BadRequest", Message: "Bad request."}
	ErrEmailAlreadyInUse = &Errno{Status: 400, Code: "ErrEmailAlreadyInUse", Message: "The email address is already in use."}
	ErrConfirmPassword   = &Errno{Status: 400, Code: "ErrConfirmPassword", Message: "The password and confirmation password do not match."}
	ErrPassswordNotMatch = &Errno{Status: 400, Code: "ErrPassswordNotMatch", Message: "The password was not match."}
	ErrUserNotFound      = &Errno{Status: 400, Code: "ErrUserNotFound", Message: "The user was not found."}

	ErrTokenInvalid = &Errno{Status: 401, Code: "ErrTokenInvalid", Message: "The token was invalid."}
)
