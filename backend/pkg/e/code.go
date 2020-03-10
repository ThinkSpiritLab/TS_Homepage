package e

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_NOT_LOGIN = 10000
	ERROR_AUTH_STID_PSW = 10001
	ERROR_AUTH_CHECK_TOKEN_FAIL = 10002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10003
	ERROR_AUTH_OTHER_ERROR = 10004

	ERROR_NOT_EXIST_USER = 10010
	ERROR_HAS_EXIST_USER = 10011

	ERROR_UPLOAD_IMAGE_TOO_LARGE = 10020
	ERROR_UPLOAD_IMAGE_WRONG_TYPE = 10021
	ERROR_UPLOAD_IMAGE_FAIL = 10022
)