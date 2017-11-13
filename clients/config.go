package main

import (
	"github.com/choueric/jconfig"
)

const DefContent = `{
	"server": "127.0.0.1:8088"
}
`

type Config struct {
	Server string `json:"server"`
}

func getConfig(file string) (*Config, error) {
	jc := jconfig.New(file, Config{})

	if _, err := jc.Load(DefContent); err != nil {
		return nil, err
	}

	return jc.Data().(*Config), nil
}
