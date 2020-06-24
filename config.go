package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	name        string
	viper       *viper.Viper
	conf2Value  map[string]reflect.Value
	watchPrefix string
}

// Can initialize multiple configuration file objects, or a configuration file bound multiple struct
func Init(cfg, watchPrefix string, watchObj interface{}) error {
	c := config{
		name:        cfg,
		viper:       viper.New(),
		conf2Value:  make(map[string]reflect.Value),
		watchPrefix: watchPrefix,
	}

	// Initialize the configuration file
	if err := c.initConfig(); err != nil {
		return err
	}

	// Monitor configuration file changes and hot load programs
	c.watchConfig()
	// Get configuration object reflection collection
	c.newWithWatcher(watchObj, watchPrefix)
	// Assign original values ​​from the configuration file at startup
	for con, val := range c.conf2Value {
		c.changeValue(con, val)
	}
	return nil
}

func (c *config) initConfig() error {
	if !fileExists(c.name) {
		log.Fatalln("config file not found.")
	}

	c.viper.SetConfigFile(c.name)
	// Obtain the configuration file type from the configuration file suffix. Unsupported types are reported by viper.
	c.viper.SetConfigType(strings.TrimLeft(filepath.Ext(c.name), "."))

	// Viper parses the configuration file
	if err := c.viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// Monitor configuration file changes and hot load programs
func (c *config) watchConfig() {
	c.viper.WatchConfig()
	c.viper.OnConfigChange(func(e fsnotify.Event) {
		for con, val := range c.conf2Value {
			c.changeValue(con, val)
		}
	})
}

func (c *config) newWithWatcher(i interface{}, prefix string) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		log.Fatalf("`NewWithFile` need struct type! %v", v.Kind())
	}
	for i := 0; i < t.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)

		// tag name use by config file type
		yml := ft.Tag.Get(strings.TrimLeft(filepath.Ext(c.name), "."))
		key := fmt.Sprintf("%s.%s", prefix, yml)
		if prefix == "" {
			key = yml
		}

		if fv.Kind() != reflect.Struct {
			c.conf2Value[key] = fv
		} else {
			c.newWithWatcher(fv.Addr().Interface(), key)
		}
	}
}

func (c *config) changeValue(field string, val interface{}) {
	var refVal reflect.Value
	var ok bool
	if refVal, ok = c.conf2Value[field]; !ok {
		return
	}
	refVal.Set(reflect.ValueOf(c.formatAtom(refVal, field)))
}

func (c *config) formatAtom(v reflect.Value, field string) interface{} {
	switch v.Kind() {
	case reflect.Invalid:
		log.Fatalf("load config failed: invalid config type.")
	case reflect.Int:
		return c.viper.GetInt(field)
	case reflect.Int8:
		return c.viper.GetInt(field)
	case reflect.Int16:
		return c.viper.GetInt(field)
	case reflect.Int32:
		return c.viper.GetInt32(field)
	case reflect.Int64:
		return c.viper.GetInt64(field)
	case reflect.Uint:
		return c.viper.GetUint(field)
	case reflect.Uint8:
		return c.viper.GetUint(field)
	case reflect.Uint16:
		return c.viper.GetUint(field)
	case reflect.Uint32:
		return c.viper.GetUint32(field)
	case reflect.Uint64:
		return c.viper.GetUint64(field)
	case reflect.Bool:
		return c.viper.GetBool(field)
	case reflect.String:
		return c.viper.GetString(field)
	case reflect.Slice:
		switch v.Type().String() {
		case "[]string":
			return c.viper.GetStringSlice(field)
		case "[]int":
			return c.viper.GetIntSlice(field)
		default:
			log.Fatalf("unsupported type :%s", v.Type().String())
		}
	case reflect.Map:
		switch v.Type().String() {
		case "map[string]interface {}":
			return c.viper.GetStringMap(field)
		case "map[string]string":
			return c.viper.GetStringMapString(field)
		case "map[string][]string":
			return c.viper.GetStringMapStringSlice(field)
		default:
			log.Fatalf("unsupported type :%s", v.Type().String())
		}
	default: // reflect.Array, reflect.Struct, reflect.Interface
		log.Fatalf("unsupported type :%s", v.Type().String())
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
