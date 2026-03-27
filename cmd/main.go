package main

import (
	"gin_mall/conf"
	"gin_mall/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
