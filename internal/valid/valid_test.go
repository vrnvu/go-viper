package valid

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/vrnvu/go-viper/internal/models"
)

var configPath = "../../configs/"

func TestFull(t *testing.T) {
	v := NewViper(configPath, "full", "yaml")
	ReadInConfig(v)

	var c models.ValidConfigViper
	err := Unmarshall(v, &c)

	assert.NoError(t, err)
	assert.Equal(t, "service-name", c.Service.Name)
	assert.Equal(t, true, c.Service.Enabled)
	assert.Equal(t, 10, c.Service.Number)
	assert.Equal(t, []string{"a", "b", "c"}, c.Service.List)
}

func TestPartialIsErr(t *testing.T) {
	v := NewViper(configPath, "partial", "yaml")
	ReadInConfig(v)

	var c models.ValidConfigViper
	err := Unmarshall(v, &c)

	expectedFieldErrs := map[string]string{
		"ValidConfigViper.Service.Name":    "required",
		"ValidConfigViper.Service.Enabled": "required",
		"ValidConfigViper.Service.List":    "required",
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			wantTag, ok := expectedFieldErrs[e.Namespace()]
			assert.True(t, ok)
			assert.Equal(t, wantTag, e.Tag())
		}
	}
}
