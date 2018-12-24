package main

import (
	"net/http"

    "github.com/DanielHenry/my-online-judge/routes"
    "github.com/DanielHenry/my-online-judge/app/Http/Controllers"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	r := routes.setup();
	// Listen and Server in 0.0.0.0:8090
	r.Run(":8090");
}
