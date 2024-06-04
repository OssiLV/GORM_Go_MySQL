package main

import (
	"name/configs"
)

func main() {
	config := configs.New()
	configs.InitDb(*config, false)

	
}
