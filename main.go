package main

import "school_app/config"

func main() {

	app := &config.App{}
	app.Connect(config.GetConfig())
	app.Run(":9000")
}
