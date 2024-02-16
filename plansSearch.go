package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinPlansSearch(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading response")
		return
	}

	var resp HouseholdsEligibilityEstimatesRequest
	if err := json.Unmarshal(body, &resp); err != nil {
		fmt.Println(string(body[:]))
		fmt.Println(err)
		fmt.Println("Error unmarshalling json to struct for GinHouseholdsEligibilityEstimates()")
		return
	}

	c.IndentedJSON(http.StatusOK, resp)
}
