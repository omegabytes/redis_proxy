package main

import (
	"redis_proxy/app"
)

func main() {
	a := app.App{}
	a.Initialize()
	a.Run()
}
