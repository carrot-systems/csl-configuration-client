package main

import (
	configurationClient "github.com/carrot-systems/csl-configuration-client"
	discoveryClient "github.com/carrot-systems/csl-discovery-client"
	env "github.com/carrot-systems/csl-env"
	"log"
)

func main() {
	env.LoadEnv()
	discovery := discoveryClient.NewClient()
	err := discovery.Register("example")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	configuration := configurationClient.NewClient(discovery)
	err = configuration.LoadConfiguration()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	println(env.RequireEnvString("c"))
}
