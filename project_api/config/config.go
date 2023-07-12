package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"pomfret.cn/project_common/logs"
)

var C = InitConfig()

type Config struct {
	viper      *viper.Viper
	SC         *ServerConfig
	GC         *GrpcConfig
	EtcdConfig *EtcdConfig
}

type ServerConfig struct {
	Name string
	Addr string
}
type EtcdConfig struct {
	Addrs []string
}
type GrpcConfig struct {
	Name string
	Addr string
}

func InitConfig() *Config {
	conf := &Config{
		viper: viper.New(),
	}
	workDir, _ := os.Getwd()
	// SetConfigName 设置配置文件的名称。
	conf.viper.SetConfigName("config")
	// SetConfigType设置由远程源返回的配置类型。
	conf.viper.SetConfigType("yaml")
	//AddConfigPath为Viper添加了一个搜索配置文件的路径。
	conf.viper.AddConfigPath("/etc/ms_project/user")
	conf.viper.AddConfigPath(workDir + "/config")

	// ReadInConfig将发现并从磁盘加载配置文件
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.ReadEtcdConfig()
	return conf
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.DebugFileName"),
		InfoFileName:  c.viper.GetString("zap.InfoFileName"),
		WarnFileName:  c.viper.GetString("zap.WarnFileName"),
		MaxSize:       c.viper.GetInt("zap.MaxSize"),
		MaxAge:        c.viper.GetInt("zap.MaxAge"),
		MaxBackups:    c.viper.GetInt("zap.MaxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}
func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}

	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addrs = addrs
	c.EtcdConfig = ec

}
