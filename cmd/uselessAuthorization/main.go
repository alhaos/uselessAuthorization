package main

import (
	"github.com/alhaos/uselessAuthorization/cmd/internal/autorizaton"
	"github.com/alhaos/uselessAuthorization/cmd/internal/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	// init auth
	auth, err := autorizaton.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// init controllers
	ec := controllers.New(auth)

	// init router
	router := gin.Default()

	// register routes
	ec.RegisterRoutes(router)

	// set templates
	err = controllers.SetTemplates(router)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// web server start
	err = router.Run(":80")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
