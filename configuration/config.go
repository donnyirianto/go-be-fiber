package configuration

import (
	"strconv"

	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/spf13/viper"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

type configImpl struct{}

func (config *configImpl) GetString(key string) string {
	return viper.GetString(key)
}

func (config *configImpl) GetInt(key string) int {
	value, err := strconv.Atoi(viper.GetString(key))
	exception.PanicLogging(err)
	return value
}

func (config *configImpl) GetBool(key string) bool {
	value, err := strconv.ParseBool(viper.GetString(key))
	exception.PanicLogging(err)
	return value
}

func New(filenames ...string) Config {
	// Set the default configuration file name(s)
	if len(filenames) > 0 {
		for _, filename := range filenames {
			viper.SetConfigFile(filename)
			if err := viper.ReadInConfig(); err != nil {
				// Handle the error (e.g., log it) and continue
				exception.PanicLogging(err)
			}
		}
	} else {
		viper.SetConfigFile(".env") // Default to .env if no filename is provided
		if err := viper.ReadInConfig(); err != nil {
			// Handle the error (e.g., log it) and continue
			exception.PanicLogging(err)
		}
	}

	// Allow environment variables to override configuration values
	viper.AutomaticEnv()

	return &configImpl{}
}
