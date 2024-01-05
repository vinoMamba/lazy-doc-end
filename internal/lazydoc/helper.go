package lazydoc

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/log"
	"github.com/vinoMamba/lazydoc/pkg/db"
)

const (
	defDir      = ".lazydoc"
	defFileName = "lazydoc.yaml"
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(filepath.Join(home, defDir))
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(defFileName)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("LAZYDOC")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err == nil {
		log.Infow("Using config file", "file", viper.ConfigFileUsed())
	}
}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

func initStore() error {
	mysqlOpts := &db.MySqlOpts{
		Host:     viper.GetString("mysql.host"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
		LogLevel: viper.GetInt("mysql.log-level"),
	}
	d, err := db.NewMySql(mysqlOpts)

	if err != nil {
		return err
	}

	store.NewStore(d)

	log.Infow("Store initialized", "store", "mysql")
	return nil
}
