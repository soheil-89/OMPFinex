package environment

import (
	"github.com/spf13/viper"
)

func Load() *viper.Viper {
	viper.SetConfigName("Apis")        // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/auth/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.auth") // call multiple times to add many search paths
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(err)
	}
	return viper.GetViper()
}
