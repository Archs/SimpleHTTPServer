package main

import (
	"flag"
	"github.com/Archs/web"
)

var (
	port = flag.String("p", "8000", "default http server port")
)

func main() {
	flag.Parse()
	web.Run(":" + *port)
}
