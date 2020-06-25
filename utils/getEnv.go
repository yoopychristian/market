//YOOPY CHRISTIAN

package utils

import "github.com/spf13/viper"

func ViperGetEnv(key, defaultValue string) string {
	viper.AutomaticEnv()

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
