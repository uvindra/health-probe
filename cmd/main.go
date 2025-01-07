package main

func main() {
	setup()

	for _, service := range serviceLookup {
		service.Start()
	}

}
