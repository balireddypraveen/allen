package configs

import (
	"fmt"
	"github.com/balireddypraveen/allen/internal/common/configs"
	"github.com/balireddypraveen/allen/internal/common/constants"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func SetupConfig() {

	configName := os.Getenv(configs.VKEYS_HOST_TYPE)

	fmt.Printf("Host Type: %v\n", configName)
	if configName == constants.EmptyString {
		configName = constants.Local
	}
	path := constants.ConfigPath
	viper.SetConfigName(configName)
	viper.AddConfigPath(path)
	viper.SetConfigType(constants.YML)

	fmt.Printf("Config file name: %s\n", configName)
	fmt.Printf("Config file path: %s\n", path)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	// Reading variables using the model
	fmt.Printf("Running in %s environment\n", configName)

	//This pulls out configs from ENV and stores in VIPER
	LoadConfig()
}

// LoadConfig This pulls the config from environment variable and saves it inside viper.
// Make sure viper is initialized with config from yaml file before calling
// only keys starting with $ will be considered ex. $FERNET_KEY
func LoadConfig() {
	// CAUTION:- DO NOT PUT LOGGING STATEMENTS IN THIS FUNCTION TO PRINT ANY SENSITIVE INFO

	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		if strings.HasPrefix(value, "$") {
			// strip the prefix
			value = strings.TrimPrefix(value, "$")
			param := os.Getenv(value)
			viper.Set(key, param)
		} else {
			newValue := viper.Get(key)
			viper.Set(key, newValue)
		}
	}
}
