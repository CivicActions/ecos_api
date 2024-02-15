package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//https://marketplace.api.healthcare.gov/api/v1/counties/by/zip/78778?apikey=d687412e7b53146b2631dc01974ad0a4

func GetCounties(ZipCode int) (*GetCountiesResponse, error) {
	var urlPrefix string = "https://marketplace.api.healthcare.gov/api/v1/counties/by/zip/"
	var urlSuffix string = "?apikey=d687412e7b53146b2631dc01974ad0a4"
	var url string = urlPrefix + strconv.Itoa(ZipCode) + urlSuffix
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
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error reading response")
	}
	//translate body
	//verify everything is there
	//fmt.Println(string(body[:]))
	//take response from getCounties
	//translate to string it that can respond with
	var request GetCountiesRequest
	if err := json.Unmarshal([]byte(body), &request); err != nil {
		fmt.Println("Error unmarshalling json to struct for GinGetCounties()")
	}
	//fmt.Printf("%+v", request)
	var intZipCode = 0
	// check root zipcode
	if request.ZipCode != "" {
		intZipCode, err = strconv.Atoi(request.ZipCode)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid zipcode format")
			return
		}
		//check campaign zipcode
	} else {
		if request.Campaign != nil && request.Campaign.ZipCode != "" {
			intZipCode, err = strconv.Atoi(request.Campaign.ZipCode)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid zipcode format")
				return
			}
		} else {
			c.String(http.StatusBadRequest, "Missing zipcode")
			return
		}
	}

	resp, err := GetCounties(intZipCode)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error retrieving counties")
		return
	}
	//to send response back to caller
	c.IndentedJSON(http.StatusOK, resp)
	return
}
