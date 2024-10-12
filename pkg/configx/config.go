package configx

import (
	"commentService/pkg/util"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func Load(path string, v any) error {
	fi := buildFieldInfo(reflect.TypeOf(v))
	for k, v := range fi.defaultValues {
		viper.SetDefault(k, v)
	}
	viper.SetEnvPrefix("comment")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.AutomaticEnv()
	return viper.Unmarshal(v, func(c *mapstructure.DecoderConfig) {
		c.TagName = "config"
		c.Squash = true
	})
}

func MustLoad(path string, v any) {
	util.PanicError(Load(path, v))
}
