package valid

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/vrnvu/go-viper/internal/models"
)

var validate *validator.Validate = validator.New()

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

func unmarshall(v *viper.Viper, c *models.ValidConfigViper) error {
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}
	if err := validate.Struct(c); err != nil {
		c = nil
		return err
	}
	return nil
}
