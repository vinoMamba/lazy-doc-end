package config

import (
	"os"

	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/vinoMamba/lazy-doc-end/pkg/helpers"
)

var vp *viper.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {
	vp = viper.New()

	vp.SetConfigType("env")

	vp.AddConfigPath(".")

	vp.SetEnvPrefix("lazy_doc_env")

	vp.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(envSuffix string) {
	loadEnv(envSuffix)

	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		vp.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := envPath + envSuffix
		if _, err := os.Stat(filePath); err == nil {
			envPath = filePath
		}
	}

	vp.SetConfigName(envPath)
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}

	vp.WatchConfig()
}

func Add(name string, configFN ConfigFunc) {
	ConfigFuncs[name] = configFN
}

func Env(key string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(key, defaultValue[0])
	}
	return internalGet(key)
}

func internalGet(key string, defaultValue ...interface{}) interface{} {
	if !vp.IsSet(key) || helpers.Empty(vp.Get(key)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return vp.Get(key)
}

func Get(key string, defaultValue ...interface{}) string {
	return GetString(key, defaultValue...)
}

func GetString(key string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(key, defaultValue...))
}
