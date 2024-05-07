package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const Version = "1.0.0"
const Priority = 1

type MyConfig struct {
	Key string `json:"key"`
}

func New() interface{} {
	return &MyConfig{}
}

func main() {
	server.StartServer(New, Version, Priority)
}

func (conf MyConfig) Access(kong *pdk.PDK) {
	return
}
