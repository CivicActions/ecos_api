package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinHouseholdsEligibilityEstimates(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading response")
		return
	}
	var req CampaignRequest
	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Println(string(body[:]))
		fmt.Println(err)
		fmt.Println("Error unmarshalling json to struct for GinHouseholdsEligibilityEstimates()")
		return
	}

	// code for averaging and replacing age range
	// if req.Household != nil {
	// 	if req.Household.People != nil {
	// 		for index, person := range *req.Household.People {
	// 			if person.Age == 0 {
	// 				ageRange := SearchByCampID(req.CampaignID).AgeRange
	// 				fmt.Println(ageRange)
	// 				result := (ageRange.Min + ageRange.Max) / 2
	// 				(*req.Household.People)[index].Age = int32(result)
	// 			}
	// 		}
	// 	}
	// }

	var url string = "https://marketplace.api.healthcare.gov/api/v1/households/eligibility/estimates?apikey=d687412e7b53146b2631dc01974ad0a4&year=2024"
	reqByte, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(reqByte))
	estimateRes, err := http.Post(url, "application/json", bytes.NewBuffer(reqByte))
	if err != nil {
		panic(err)
	}
	defer estimateRes.Body.Close()
	finalBody, err := ioutil.ReadAll(estimateRes.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(finalBody))

	var estimateResponse HouseholdsEligibilityEstimatesRespond
	if err := json.Unmarshal(finalBody, &estimateResponse); err != nil {
		fmt.Println(string(body[:]))
		fmt.Println("Error unmarshalling json to struct for estimateResponse()")
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, estimateResponse)
}

func SearchByCampID(camp_id string) Campaign {
	rtnCampaign := Campaign{}
	for _, campaign := range Campaigns {
		if campaign.CampaignID == camp_id {
			rtnCampaign = campaign
		}
	}
	return rtnCampaign
}
