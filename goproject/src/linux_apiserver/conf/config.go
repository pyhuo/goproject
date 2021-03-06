package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	name string
}

//初始化config
func Init(cfg string) error {
	config := Config{name: cfg}
	if err := config.InitConfit(); err != nil {
		return err
	}

	//初始化log信息
	initLog()
	config.watchConfig()
	return nil
}

// 通过指定配置文件可以很方便地连接不同的环境（开发环境、测试环境）并加载不同的配置，方便开发和测试。
func (c *Config) InitConfit() error {
	name := c.name
	viper.SetConfigType("yaml") //配置类型
	if name == "" {
		viper.AddConfigPath("./conf")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigFile(name)
	}
	viper.AutomaticEnv()            //读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") //// 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		log.Infof("错误是：", err)
		return err
	}

	return nil
}

//初始化log信息
func initLog() {
	//读取配置文件信息
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}
// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}