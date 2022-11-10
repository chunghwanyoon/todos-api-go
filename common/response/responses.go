package response

type MessageLiteral string

type ResponseData struct {
	ResponseMessage MessageLiteral `json:"message"`
	StatusCode      int            `json:"-"`
}
