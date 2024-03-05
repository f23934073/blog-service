package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "導入參數錯誤")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "驗證失敗，找不到對應的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "驗證失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(10000005, "驗證失敗，Token逾時")
	UnauthorizedTokenGenerate = NewError(10000006, "驗證失敗，Token生成失敗")
	TooManyRequests           = NewError(10000007, "請求過多")
)

// 2001xxx 標籤相關錯誤
var (
	ErrorGetTagListFail = NewError(20010001, "獲取標籤列表失敗")
	ErrorCreateTagFail  = NewError(20010002, "創建標籤失敗")
	ErrorUpdateTagFail  = NewError(20010003, "更新標籤失敗")
	ErrorDeleteTagFail  = NewError(20010004, "刪除標籤失敗")
	ErrorCountTagFail   = NewError(20010005, "統計標籤失敗")
)

// 2003xxx 上傳文件相關錯誤
var (
	ErrorUploadFileFail = NewError(20030001, "上傳文件失敗")
)
