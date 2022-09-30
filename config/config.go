package config

import (
	"fmt"
	utils "github.com/davidchou93/wallets/pkg/slice"
	"github.com/spf13/viper"
	"strings"
)

type AppEnv string

const (
	EnvProd AppEnv = "prod"
	EnvTest AppEnv = "test"
	EnvDev  AppEnv = "dev"
)

func GetAppEnv(conf *Config) AppEnv {
	env := conf.AppEnv
	if !utils.SliceContainElement([]AppEnv{EnvDev, EnvTest, EnvProd}, AppEnv(env)) {
		fmt.Println(conf)
		panic("APP_ENV not valid")
	}
	return AppEnv(conf.AppEnv)
}

type Config struct {
	// Router part
	// Env = prod/test/dev
	AppEnv    string   `mapstructure:"app_env"`
	Port      string   `mapstructure:"port"`
	MasterKey string   `mapstructure:"master_key"`
	DB        DBConfig `mapstructure:"db"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"db_name"`
	Password string `mapstructure:"password"`
}

func GetConfig() (*Config, error) {
	vpr := viper.New()
	vpr.SetDefault("APP_ENV", "dev")
	vpr.AutomaticEnv()
	vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vpr.SetConfigType("yaml")
	vpr.SetConfigName(vpr.GetString("APP_ENV"))
	vpr.AddConfigPath("./")           // for ./unit_testing.yml
	vpr.AddConfigPath("./config")     // for go run main.go
	vpr.AddConfigPath("../config")    // for running under bin/
	vpr.AddConfigPath("../../config") // for running unit tests

	if err := vpr.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := vpr.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
