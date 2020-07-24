package config

import (
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"github.com/zngue/go_tool/src/fun/file"
	"github.com/zngue/go_tool/src/log"
	"github.com/zngue/go_tool/src/sign_chan"
	"io/ioutil"
)
const ConfigJson = ""
const ConfigYaml = "config.yaml"
func JsonToStruck() *Config  {
	var config Config
	data, err := ioutil.ReadFile(ConfigJson)
	if err != nil {
		return nil
	}
	jerr:=json.Unmarshal(data,&config)
	if jerr!=nil {
		log.LogInfo(jerr)
		return nil
	}
	return  &config
}
func YamlToStruck()(configinfo *Config) {
	var config Config
	if !file.FileExist(ConfigYaml){
		sign_chan.SignLog(errors.New("config.yaml is not Exist "))
		return
	}
	v := viper.New()
	v.SetConfigFile(ConfigYaml)
	err := v.ReadInConfig()
	if err != nil {
		sign_chan.SignLog(err)
		return
	}
	v.WatchConfig()
	if err := v.Unmarshal(&config); err != nil {
		sign_chan.SignLog(err)
		return
	}
	configinfo = &config
	return
}