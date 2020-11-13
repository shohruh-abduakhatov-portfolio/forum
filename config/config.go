package config

import "os"

var (
	conf = map[string]string{}
)

func InitConfig() {
	for k, v := range conf {
		os.Setenv(k, v)
	}
}
