package naive

import (
	"github.com/spf13/viper"
	"github.com/vrnvu/go-viper/internal/models"
)

func newViper(configPath, configName, configType string) *viper.Viper {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	return v
}

func readInConfig(v *viper.Viper) {
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
}

func unmarshall(v *viper.Viper, c *models.ConfigViper) {
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}
}
