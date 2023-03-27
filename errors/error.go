package errors

type AppError struct {
	// ErrCode -> Response and Logs
	// Message -> Response Message
	// Err -> Log Message
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}

func (appErr *AppError) Unwrap() error {
	return appErr.Err
}
