package main

import (
	"router/router"
)

func main() {
	r := *router.SetupRouter()
	r.Run(":8080")

}
