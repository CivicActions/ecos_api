package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/counties/by/zip/:zipcode", GinGetCounties)
	router.POST("/plans/search", GinPlansSearch)
	router.POST("/households/eligibility/estimates", GinHouseholdsEligibilityEstimates)
	//uncomment to run locally
	//router.Run("localhost:8080")
}
