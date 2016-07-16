package parser

import "github.com/spf13/viper"

// GetString return the key from the config as a string. Optional default value as second argument.
func GetString(key string, params ...string) (value string) {
	value = viper.GetString(key)
	if value == "" {
		value = params[0]
	}
	return
}
