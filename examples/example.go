package main

import "djgo"

type HomeController struct {
	djgo.Controller
}

func main() {
	djgo.Router("/", &HomeController{})
	djgo.Run()
}
