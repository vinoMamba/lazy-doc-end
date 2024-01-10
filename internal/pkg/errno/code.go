package errno

var (
	OK                   = &Errno{Status: 200, Code: "", Message: "OK"}
	InternalServerError  = &Errno{Status: 500, Code: "InternalError", Message: "Internal server error."}
	NotFound             = &Errno{Status: 404, Code: "NotFound", Message: "The requested resource was not found."}
	BadRequest           = &Errno{Status: 400, Code: "BadRequest", Message: "Bad request."}
	ErrEmailAlreadyInUse = &Errno{Status: 400, Code: "ErrEmailAlreadyInUse", Message: "The email address is already in use."}
	ErrConfirmPassword   = &Errno{Status: 400, Code: "ErrConfirmPassword", Message: "The password and confirmation password do not match."}
)
