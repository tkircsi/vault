package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tkircsi/vault/api/rest"
)

func (app *application) RunREST() {
	r := gin.Default()
	v1 := r.Group("/v1")
	h := rest.NewRESTHandler(app.vault)

	{
		v1.GET("/get/:token", h.Get)
		v1.POST("/put", rest.JSONContent(), h.Put)
		v1.GET("/health", h.HealthCheck)
	}

	log.Println("vault REST service started...")
	log.Fatal(r.Run(app.RESTPort))
}
