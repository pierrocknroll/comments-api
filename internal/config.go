package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func checkConfigParameters() error {
	commentsDatabase := viper.GetString("COMMENTS_DATABASE_ADDRESSS")
	if commentsDatabase == "" {
		return fmt.Errorf("COMMENTS_DATABASE_ADDRESSS not found in config file")
	}

	accessKey := viper.GetString("ACCESS_KEY")
	if accessKey == "" {
		return fmt.Errorf("ACCESS_KEY parameter not found in config file")
	}

	return nil
}

func initializeConfig(configFilename string) error {

	viper.SetDefault("LOG_FILENAME", "")
	viper.SetDefault("LOG_LEVEL", "INFO")
	viper.SetDefault("PORT_NUMBER", 8080)

	// Allows the environment variables to overwrite the parameters in the config file
	viper.AutomaticEnv()
	viper.SetConfigFile(configFilename)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	// If the config file doesn't exist, it may be because environment variables are already defined.
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			_, ok = err.(*os.PathError)
			if !ok {
				return err
			}
		}
	}

	err = checkConfigParameters()
	return err
}
