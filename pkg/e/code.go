package e

type Code int

const (
	SUCCESS       Code = 200
	ERROR         Code = 500
	InvalidParams Code = 400
)
const (
	ErrorUploadSaveFileFail Code = iota + 10001
	ErrorUploadCheckFileFail
	ErrorUploadCheckFileFormat
)
