package e

var MsgFlags = map[Code]string{
	SUCCESS:                    "成功",
	ERROR:                      "失败",
	InvalidParams:              "请求参数错误",
	ErrorUploadSaveFileFail:    "保存文件失败",
	ErrorUploadCheckFileFail:   "检查文件失败",
	ErrorUploadCheckFileFormat: "校验文件错误，文件格式错误",
}

func GetMsg(code Code) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
