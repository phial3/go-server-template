package setting

import "github.com/spf13/viper"

//Setting 配置管理结构体
type Setting struct {
	vp *viper.Viper
}

//NewSetting 实例化配置管理结构体
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
