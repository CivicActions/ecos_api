package main

type County struct {
	ZipCode string `json:"zipcode,omitempty"`
	Name    string `json:"name,omitempty"`
	Fips    string `json:"fips,omitempty"`
	State   string `json:"state,omitempty"`
}

type GetCountiesResponse struct {
	Counties *[]County `json:"counties,omitempty"`
}

type Place struct {
	CountyFips int32
	State      string
	ZipCode    string
}

type IntRange struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

type Campaign struct {
	CampaignID  string    `json:"campaign_id,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	IncomeRange *IntRange `json:"income_range,omitempty"`
	AgeRange    *IntRange `json:"age_range,omitempty"`
	ZipCode     string    `json:"zipcode,omitempty"`
	IsParent    bool      `json:"is_parent,omitempty"`
}

// may need to rename this struct to specific endpoint name like county request
type GetCountiesRequest struct {
	Campaign    *Campaign `json:"campaign,omitempty"`
	ZipCode     string    `json:"zipcode,omitempty"`
	Income      float64   `json:"income,omitempty"`
	UsesTobacco bool      `json:"uses_tobacco,omitempty"`
}

func NewCampaign() *Campaign {
	var IncomeRange IntRange
	var AgeRange IntRange
	var Campaign Campaign

	Campaign.IncomeRange = &IncomeRange
	Campaign.AgeRange = &AgeRange

	return &Campaign
}

func NewRequest() *GetCountiesRequest {
	var Request GetCountiesRequest

	Request.Campaign = NewCampaign()

	return &Request
}

type People struct {
	AptcEligible bool
	Age          int32
	HasMec       bool
	IsPregnant   bool
	IsParent     bool
	UsesTobacco  bool
	Gender       string
}

type Household struct {
	Income               float64
	People               People
	HasMarriedCouple     bool
	UnemploymentReceived string
}
