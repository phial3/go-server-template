package setting

import "time"

//ServerSettingS 服务端配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

//AppSettingS 服务器软件配置
type AppSettingS struct {
	DefaultPageSize         int
	MaxPageSize             int
	NormalUserDeviceUplimit int
	VipUserDeviceUplimit    int
	LogSavePath             string
	LogFileName             string
	LogFileExt              string
}

//DatabaseSettingS 数据库配置信息结构体
type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

//JWTSettingS jwt配置信息
type JWTSettingS struct {
	AppSecret string
	Issuer    string
	Expire    time.Duration
}

//EmailSettingS 告警邮件配置信息
type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

//ReadSection 读取配置信息
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
