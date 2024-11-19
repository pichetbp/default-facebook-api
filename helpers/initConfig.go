package helpers

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	err := viper.ReadInConfig()
	if err != nil {

	}
}

func GetConnectionString() string {
	mssqlHost := viper.Get("db.host")
	mssqlUsername := viper.Get("db.username")
	mssqlPassword := viper.Get("db.password")
	mssqlDatabase := viper.Get("db.database")
	return fmt.Sprintf("sqlserver://%v:%v@%v?database=%s", mssqlUsername, mssqlPassword, mssqlHost, mssqlDatabase)
}

func GetDriverString() string {
	return fmt.Sprint(viper.Get("db.driver"))
}
