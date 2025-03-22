package consts

const (
	ErrnoSuccess      = 0
	ErrorUnknownError = 1
)

var ErrMsg = map[int]string{
	ErrnoSuccess:      "success",
	ErrorUnknownError: "unknown error",
}
