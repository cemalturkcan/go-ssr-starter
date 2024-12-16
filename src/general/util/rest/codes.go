package rest

const (
	ServerInternalError = "500"
	NotFound            = "404"
)

var ErrorCode = map[string]string{
	NotFound:            "Not Found",
	ServerInternalError: "Server Internal Error",
}

func Error(error error) (string, string) {
	err := ErrorCode[error.Error()]
	if err == "" {
		return ServerInternalError, ErrorCode[ServerInternalError]
	}
	return error.Error(), err
}
