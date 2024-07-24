package response

type ErrorMessages string

const (
	EmSomethingWentWrong ErrorMessages = "Something went wrong, Please try again later"
	MalformedRequest     ErrorMessages = "Bad Input request"
	EmValidationError    ErrorMessages = "Error while validating the request"
)
