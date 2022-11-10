package main

import "github.com/Eliwelton-The-Espada/go-expert/apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
