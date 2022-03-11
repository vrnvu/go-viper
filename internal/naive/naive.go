package naive

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/vrnvu/go-viper/internal/models"
)

func NewViper(configPath, configName, configType string) *viper.Viper {
	v := viper.New()
	d, _ := os.Getwd()
	fmt.Println(d)
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	return v
}

func ReadInConfig(v *viper.Viper) {
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Unmarshall(v *viper.Viper, c *models.ConfigViper) {
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}
}
