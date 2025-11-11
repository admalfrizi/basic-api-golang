package main

import (
	"api-golang/modules/station"
	"github.com/gin-gonic/gin"
);

func main() {
    Initiate()
}

func Initiate() {
	var (
		router = gin.Default();
 		api = router.Group("v1/api");
	) 

	station.Initiate(api)

	router.Run(":8080");
}