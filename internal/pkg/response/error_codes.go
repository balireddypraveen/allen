package response

type ErrorCode string

const (
	CodeUnknown          ErrorCode = "CODE_UNKNOWN"
	CodePanic            ErrorCode = "CODE_PANIC"
	BadRequest           ErrorCode = "BAD_REQUEST"
	ErrorWhileProcessing ErrorCode = "ERROR_WHILE_PROCESSING"
	Unauthorized         ErrorCode = "UNAUTHORIZED"
	InvalidInput         ErrorCode = "INVALID_INPUT"
)
