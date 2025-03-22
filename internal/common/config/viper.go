package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	Storage StorageConfig `mapstructure:"storage"`
}

type ServerConfig struct {
	Port      string `mapstructure:"port"`
	JWTSecret string `mapstructure:"jwt_secret"`
}

type StorageConfig struct {
	Type     string         `mapstructure:"type"`
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

var GlobalConfig Config

func Init() error {
	// 设置配置文件信息
	viper.SetConfigName("global")
	viper.SetConfigType("yaml")

	// 添加配置文件搜索路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("./internal/common/config")
	viper.AddConfigPath("../internal/common/config")
	viper.AddConfigPath("../../internal/common/config")

	// 设置环境变量前缀
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// 将配置解析到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

// GetConfig 返回全局配置
func GetConfig() *Config {
	return &GlobalConfig
}

// GetDatabaseDSN 返回数据库连接字符串
func (c *DatabaseConfig) GetDatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}
