package boot

import "github.com/spf13/viper"

// InitConfig will init the config context
// path string is the directory of the config
// fileName string is the file name without extension
func InitConfig(path, fileName string) error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	return viper.ReadInConfig()
}

func ParseConfig() {
	var apiKeys = viper.Get("keys").([]interface{})
	var apiKeyMap = make(map[string]struct{})
	for _, v := range apiKeys {
		apiKeyMap[v.(string)] = struct{}{}
	}
	viper.Set("keys", apiKeyMap)
}
