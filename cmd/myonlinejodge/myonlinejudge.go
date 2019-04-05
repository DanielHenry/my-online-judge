package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := routes.setup();
	// Listen and Server in 0.0.0.0:8090
	r.Run(":8090");
}
