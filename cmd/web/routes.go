package main

import "github.com/gin-gonic/gin"

func (app *application) RunREST() {
	r := gin.Default()
	v1 := r.Group("/v1")
	h := NewRESTHandler(app.vault)

	{
		v1.GET("/get/:token", h.Get)
		v1.POST("/put", h.Put)
	}

	r.Run(app.port)
}
