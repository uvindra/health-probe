package main

import "log"

func main() {
	log.Println("Starting...")
	setup()

	log.Println("Number of services to run: ", len(services))

	for _, service := range services {
		go service.Start()
	}

	controler.Start()
}
