//YOOPY CHRISTIAN

package config

import (
	"challenge/utils"
	"fmt"
)

var (
	DbUser,
	DbPassword,
	DbHost,
	DbPort,
	DbSchema string
)

func Env_conn() (dbEngine, dbSource string) {
	dbEngine = utils.ViperGetEnv("DB_ENGINE", "mysql")
	DbUser = utils.ViperGetEnv("DB_USER", "root")
	DbPassword = utils.ViperGetEnv("DB_PASSWORD", "admin")
	DbHost = utils.ViperGetEnv("DB_HOST", "localhost")
	DbPort = utils.ViperGetEnv("DB_PORT", "3306")
	DbSchema = utils.ViperGetEnv("DB_SCHEMA", "enigmart")

	dbSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbSchema)
	return dbEngine, dbSource
}
