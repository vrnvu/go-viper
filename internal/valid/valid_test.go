package valid

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/vrnvu/go-viper/internal/models"
)

const (
	fieldErrMsg = "Key: '%s' Error:Field validation for '%s' failed on the '%s' tag"
)

var configPath = "../../configs/"

func TestFull(t *testing.T) {
	v := newViper(configPath, "full", "yaml")
	readInConfig(v)

	var c models.ValidConfigViper
	err := unmarshall(v, &c)

	assert.NoError(t, err)
	assert.Equal(t, "service-name", c.Service.Name)
	assert.Equal(t, true, c.Service.Enabled)
	assert.Equal(t, 10, c.Service.Number)
	assert.Equal(t, []string{"a", "b", "c"}, c.Service.List)
}

func TestPartialIsErr(t *testing.T) {
	v := newViper(configPath, "partial", "yaml")
	readInConfig(v)

	var c models.ValidConfigViper
	err := unmarshall(v, &c)

	// TODO
	var validationErrors *validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for e := range err {

		}
		var fieldErr *validator.FieldError
		assert.ErrorIs(t, err, *fieldErr)
		assertFieldError(t, err, "ValidConfigViper.Service.Name", "hello", "required")
	}
}

func assertFieldError(t *testing.T, err error, key, field, tag string) {
	assert.ErrorIs(t, err, fmt.Errorf(fieldErrMsg, key, field, tag))
}
