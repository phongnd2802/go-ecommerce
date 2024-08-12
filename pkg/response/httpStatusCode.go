package response

const (
	CodeSuccess = 20000
	CodeCreated = 20001
)

const (
	ErrCodeBadRequest               = 40001
	ErrCodeInternalServer           = 50000
	ErrCodeShopExist                = 40011
	ErrCodeEmailOrPasswordIncorrect = 40010
	ErrCodeForbidden                = 40003
	ErrCodeNotFound                 = 40004
	ErrCodeFailedVerifyJWT          = 40020
	ErrCodeShopNotExist             = 40021
	ErrCodeRefreshTokenUsed         = 40022
	ErrCodeFailedInsertDB           = 50001
	ErrCodeFailedQueryDB            = 50002
	ErrCodeInvalidProductType       = 40023
)

var msg = map[int]string{
	CodeSuccess: "Success",
	CodeCreated: "Created",

	ErrCodeBadRequest:               "Bad Request",
	ErrCodeInternalServer:           "Internal Server Error",
	ErrCodeShopExist:                "Email already exists",
	ErrCodeEmailOrPasswordIncorrect: "Email or Password Incorrect",
	ErrCodeForbidden:                "Forbidden Error",
	ErrCodeNotFound:                 "Not Found",
	ErrCodeFailedVerifyJWT:          "Failed to verify JWT token",
	ErrCodeShopNotExist:             "Shop not regitered",
	ErrCodeRefreshTokenUsed:         "Please Login again",
	ErrCodeFailedInsertDB:           "Insert Record DB Error",
	ErrCodeInvalidProductType:       "Invalid Product Type",
	ErrCodeFailedQueryDB:            "Query Record DB Error",
}
