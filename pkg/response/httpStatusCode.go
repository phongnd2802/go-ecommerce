package response

const (
	CodeSuccess = 20000
	CodeCreated = 20001
	CodeUpdated = 20004

	CodePublishProductSuccess   = 20011
	CodeUnPublishProductSuccess = 20012
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
	ErrCodeFailedUpdateDB           = 50003
	ErrCodeInvalidProductType       = 40023
	ErrCodeNotFoundProduct          = 40024
)

var msg = map[int]string{
	CodeSuccess:                     "Success",
	CodeCreated:                     "Created",
	CodeUpdated:                     "Updated",
	CodePublishProductSuccess:       "Publish Product Success",
	CodeUnPublishProductSuccess:     "UnPublish Product Success",
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
	ErrCodeNotFoundProduct:          "Not Found Product For Shop",
	ErrCodeFailedUpdateDB:           "Update Resource Failed",
}
