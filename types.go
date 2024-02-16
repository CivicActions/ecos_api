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
	CountyFips string `json:"countyfips,omitempty"`
	State      string `json:"state,omitempty"`
	ZipCode    string `json:"zipcode,omitempty"`
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

type CampaignRequest struct {
	CampaignID string     `json:"campaign_id,omitempty"`
	Place      *Place     `json:"place,omitempty"`
	Market     string     `json:"market,omitempty"`
	Household  *Household `json:"household,omitempty"`
}

// may need to rename this struct to specific endpoint name like county request
type GetCountiesRequest struct {
	Campaign    *Campaign `json:"campaign,omitempty"`
	ZipCode     string    `json:"zipcode,omitempty"`
	Income      float64   `json:"income,omitempty"`
	UsesTobacco bool      `json:"uses_tobacco"`
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
	AptcEligible bool   `json:"aptc_elibigle"`
	Age          int32  `json:"age,omitempty"`
	HasMec       bool   `json:"has_mec"`
	IsPregnant   bool   `json:"is_pregnant"`
	IsParent     bool   `json:"is_parent"`
	UsesTobacco  bool   `json:"uses_tobacco"`
	Gender       string `json:"gender,omitempty"`
}

type Household struct {
	Income               float64   `json:"income,omitempty"`
	People               *[]People `json:"people,omitempty"`
	HasMarriedCouple     bool      `json:"has_married_couple`
	UnemploymentReceived string    `json:"unemployment_received,omitempty"`
}

type HouseholdsEligibilityEstimatesRequest struct {
	CampaignId string     `json:"campaignId,omitempty"`
	Place      *Place     `json:"place,omitempty"`
	Market     string     `json:"market,omitempty"`
	Household  *Household `json:"household,omitempty"`
}

type Estimates struct {
	Aptc               int    `json:"aptc,omitempty"`
	Csr                string `json:"csr,omitempty"`
	Hardship_exemption bool   `json:"hardship_exemption"`
	Is_medicaid_chip   bool   `json:"is_medicaid_chip"`
	In_coverage_gap    bool   `json:"in_coverage_gap"`
}

type HouseholdsEligibilityEstimatesRespond struct {
	Estimates *[]Estimates `json:"estimates,omitempty"`
}

var Campaigns = []Campaign{
	{CampaignID: "campaign_1", Gender: "Female", IncomeRange: &IntRange{Min: 30000, Max: 40000}, AgeRange: &IntRange{Min: 35, Max: 45}, ZipCode: "73301", IsParent: true},
	{CampaignID: "campaign_2", Gender: "Female", IncomeRange: &IntRange{Min: 30000, Max: 40000}, AgeRange: &IntRange{Min: 40, Max: 50}, ZipCode: "73301", IsParent: true},
}
