package response

const (
	CodeSuccess = 20000
	CodeCreated = 20001
)

const (
	ErrCodeBadRequest     = 40001
	ErrCodeInternalServer = 50000
	ErrCodeShopExist      = 40011
)

var msg = map[int]string{
	CodeSuccess: "Success",
	CodeCreated: "Created",

	ErrCodeBadRequest:     "Bad Request",
	ErrCodeInternalServer: "Internal Server Error",
	ErrCodeShopExist:      "Email already exists",
}
