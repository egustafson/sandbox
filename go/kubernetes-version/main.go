package main

func main() {

	apiCtrl, err := InitKubeAPIController()
	if err != nil {
		panic(err)
	}

	webServer, err := InitWebServer(apiCtrl)
	if err != nil {
		panic(err)
	}

	if err := webServer.Start(); err != nil {
		panic(err)
	}
}
