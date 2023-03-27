package errors

type ErrCode string

const (
	Unknown  ErrCode = "U000"
	NotFound ErrCode = "U001"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"
	DeleteDataFailed ErrCode = "S006"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{ErrCode: code, Message: message, Err: err}
}
