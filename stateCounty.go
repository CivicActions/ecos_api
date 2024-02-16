package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

//https://marketplace.api.healthcare.gov/api/v1/counties/by/zip/78778?apikey=d687412e7b53146b2631dc01974ad0a4

func GetCounties(ZipCode string) (*GetCountiesResponse, error) {
	//validate that zipcode is appropriate format using regex
	var urlPrefix string = "https://marketplace.api.healthcare.gov/api/v1/counties/by/zip/"
	var urlSuffix string = "?apikey=d687412e7b53146b2631dc01974ad0a4"
	var url string = urlPrefix + ZipCode + urlSuffix
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println("Error with Get response")
		return nil, fmt.Errorf("Error with Get response\r\n%s", err.Error())
	}
	defer resp.Body.Close() //to set aside call to close body of response so it's closed when func exits
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response")
		return nil, err
	}

	//translate the body from json to struct
	//return resulting struct
	var counties GetCountiesResponse
	if err := json.Unmarshal(body, &counties); err != nil {
		fmt.Println(string(body[:]))
		fmt.Println("Error unmarshalling json to struct for GetCounties()")
		return nil, err
	}

	return &counties, nil
}

func GinGetCounties(c *gin.Context) {
	resp, err := GetCounties(c.Param("zipcode"))
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error retrieving counties")
		return
	}
	//to send response back to caller
	c.IndentedJSON(http.StatusOK, resp)
	return
}
