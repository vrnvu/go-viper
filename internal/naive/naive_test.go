package naive

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vrnvu/go-viper/internal/models"
)

var configPath = "../../configs/"

func TestConfigPath(t *testing.T) {
	v := NewViper(configPath, "full", "yaml")
	ReadInConfig(v)
	assert.Equal(t, "service-name", v.GetString("service.name"))
}

func TestFull(t *testing.T) {
	v := NewViper(configPath, "full", "yaml")
	ReadInConfig(v)

	var c models.ConfigViper
	Unmarshall(v, &c)

	assert.Equal(t, "service-name", c.Service.Name)
	assert.Equal(t, true, c.Service.Enabled)
	assert.Equal(t, 10, c.Service.Number)
	assert.Equal(t, []string{"a", "b", "c"}, c.Service.List)
}

func TestPartialUsesDefaults(t *testing.T) {
	v := NewViper(configPath, "partial", "yaml")
	ReadInConfig(v)

	var c models.ConfigViper
	Unmarshall(v, &c)

	assert.Equal(t, "", c.Service.Name)
	assert.Equal(t, false, c.Service.Enabled)
	assert.Equal(t, 10, c.Service.Number)
	assert.Equal(t, []string(nil), c.Service.List)
}

func TestAppendToNilList(t *testing.T) {
	v := NewViper(configPath, "partial", "yaml")
	ReadInConfig(v)

	var c models.ConfigViper
	Unmarshall(v, &c)

	assert.Equal(t, []string(nil), c.Service.List)
	c.Service.List = append(c.Service.List, "a")
	assert.Equal(t, []string{"a"}, c.Service.List)
}

func TestEmptyConfig(t *testing.T) {
	v := NewViper(configPath, "empty", "yaml")
	ReadInConfig(v)

	var c models.ConfigViper
	Unmarshall(v, &c)

	assert.Equal(t, "", c.Service.Name)
	assert.Equal(t, false, c.Service.Enabled)
	assert.Equal(t, 0, c.Service.Number)
	assert.Equal(t, []string(nil), c.Service.List)
}
