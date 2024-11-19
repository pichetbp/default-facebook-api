package usecases

import (
	"bytes"
	"default-repo/helpers"
	FBModel "default-repo/libs/httpEntry/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

var FacebookEntry FacebookEntryInterface

type FacebookEntryInterface interface {
}

type facebookEntry struct {
	accessToken string
	actID       string
}

func (f *facebookEntry) Init() {
	f.accessToken = viper.GetString("faceook.accessToken")
	f.actID = viper.GetString("facebook.actID")
}

func NewFacebookEntry() FacebookEntryInterface {
	FacebookEntry = &facebookEntry{}
	return FacebookEntry
}

// create campaign : https://developers.facebook.com/docs/marketing-apis/get-started/#campaign
func (f *facebookEntry) CreateCampaign(req *FBModel.CreateCampaignReq) (*FBModel.CreateCampaignRes, error) {

	url := "https://graph.facebook.com/v21.0/act_" + f.actID + "/campaigns"
	payload := map[string]interface{}{
		"name":                req.Name,
		"objective":           req.Objective,
		"status":              req.Status,
		"special_ad_category": req.SpecialAdCategories,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+f.accessToken)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("failed to create campaign: %s", string(bodyBytes))
	}

	var res FBModel.CreateCampaignRes
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (f *facebookEntry) GetCampaigns(req *FBModel.GetCampaignReq) (*FBModel.GetCampaignRes, error) {
	url := "https://graph.facebook.com/v21.0/"
	payload := map[string]any{
		"fields": helpers.GetFieldMap(FBModel.GetCampaignRes{}),
		"id":     req.ID,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+f.accessToken)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("failed to get campaign: %s", string(bodyBytes))
	}

	var res FBModel.GetCampaignRes
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// define target
func (f *facebookEntry) DefineTarget() {
	// define target
}

// create ad set ==> using the flexible ad format : https://developers.facebook.com/docs/marketing-api/flexible-ad-format
func (f *facebookEntry) CreateAdSet() {
	// create ad set
}

// provide ad create
func (f *facebookEntry) CreateAd() {
	// create ad
}

// schedule delivery
func (f *facebookEntry) ScheduleDelivery() {
	// schedule delivery
}

// audiance targeting : https://developers.facebook.com/docs/marketing-api/audiences/overview
func (f *facebookEntry) AudienceTargeting() {
	// audience targeting
}

// audience rule : https://developers.facebook.com/docs/marketing-api/audiences/guides/audience-rules
