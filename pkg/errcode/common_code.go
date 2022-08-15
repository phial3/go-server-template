package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务器内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	Notfound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
	ErrorCreateUserFail       = NewError(20000001, "创建用户失败") /*用户相关*/
	ErrorUpdateUserNameFail   = NewError(20000002, "更新用吗名失败")
	ErrorUpdateUserPwdFail    = NewError(20000003, "更新用户密码失败")
	ErrorGetUserInfoFail      = NewError(20000004, "获取用户信息失败")
	ErrorUserNotFound         = NewError(20000005, "用户不存在")
	ErrorUserExist            = NewError(20000006, "用户已存在")
	ErrorWrongPassword        = NewError(20000007, "用户密码错误")
	ErrorPhoneExist           = NewError(20000008, "手机号已被注册")
	ErrorDeviceCreateFail     = NewError(20000020, "设备创建失败") //设备相关错误
	ErrorDeviceUpperLimit     = NewError(20000021, "用户设备达到上限")
	ErrorDeviceUpdateFail     = NewError(20000022, "更新设备失败")
	ErrorDeviceDelFail        = NewError(20000023, "删除设备失败")
	ErrorDeviceCountFail      = NewError(20000024, "获取设备总数失败")
	ErrorDeviceListFail       = NewError(20000025, "获取设备列表失败")
)
