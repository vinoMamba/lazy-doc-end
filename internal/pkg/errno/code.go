package errno

var (
	OK                  = &Errno{Status: 200, Code: "", Message: "OK"}
	InternalServerError = &Errno{Status: 500, Code: "InternalError", Message: "Internal server error."}
	NotFound            = &Errno{Status: 404, Code: "NotFound", Message: "The requested resource was not found."}
)
