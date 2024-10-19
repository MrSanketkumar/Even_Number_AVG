package constant

// Error Messages
const (
	ErrorReadingFileMessage        = "error reading file:"
	ErrorNoNumbersProvidedMessage  = "please provide even numbers"
	ErrorInvalidNumberMessage      = "invalid number provided"
	ErrorServerClosedMessage       = "server is closed"
	ErrorStartingServerMessage     = "error starting server: %s\n"
	ErrorInvalidInputMessage       = "invalid input received"
	ErrorResponseDecodeMessage     = "failed to decode response: %v"
	ErrorMessageFormat             = "Expected error %v, got %v"
	FailedToSendRequestMessage     = "failed to send request: %v"
)

// HTTP Method Errors
const (
	ErrorInvalidHttpMethodMessage = "invalid HTTP method"
)
