package consts

const (
	//10000-19999 通用错误码

	ErrUserExist    = 10001
	ErrUserExistMsg = "用户已存在"

	// ErrUserNotExist 10002 用户不存在
	ErrUserNotExist    = 10002
	ErrUserNotExistMsg = "用户不存在"

	// ErrUserPasswordError 10003 用户密码错误
	ErrUserPasswordError    = 10003
	ErrUserPasswordErrorMsg = "用户密码错误"

	// ErrCloseThisInterface 10007 暂时关闭此接口
	ErrCloseThisInterface    = 10007
	ErrCloseThisInterfaceMsg = "暂时关闭此接口"
)
