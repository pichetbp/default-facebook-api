package usecases

import (
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
	Init()
	CreateCampaign(req *FBModel.CreateCampaignReq) (*FBModel.CreateCampaignRes, error)
	CreateCustomAudience(req *FBModel.CreateCustomAudienceReq) (*FBModel.CreateCustomAudienceRes, error)
}

type facebookEntry struct {
	accessToken string
	actID       string
}

func NewFacebookEntry() FacebookEntryInterface {
	FacebookEntry = &facebookEntry{}
	FacebookEntry.Init()
	return FacebookEntry
}

func (f *facebookEntry) Init() {
	f.accessToken = viper.GetString("facebook.accessToken")
	f.actID = viper.GetString("facebook.actID")
}

// create campaign : https://developers.facebook.com/docs/marketing-apis/get-started/#campaign
func (f *facebookEntry) CreateCampaign(req *FBModel.CreateCampaignReq) (*FBModel.CreateCampaignRes, error) {

	url := "https://graph.facebook.com/v21.0/" + f.actID + "/campaigns"
	payload := map[string]string{
		"name":                  req.Name,
		"objective":             req.Objective,
		"status":                req.Status,
		"special_ad_categories": helpers.ResponseStringFromList(req.SpecialAdCategories),
	}

	return doHttp[FBModel.CreateCampaignRes](url, "POST", payload, f)
}

func (f *facebookEntry) GetCampaigns(req *FBModel.GetCampaignReq) (*FBModel.GetCampaignRes, error) {
	url := "https://graph.facebook.com/v21.0/"
	payload := map[string]string{
		"fields": helpers.GetFieldMap(FBModel.GetCampaignRes{}),
		"id":     req.ID,
	}
	return doHttp[FBModel.GetCampaignRes](url, "GET", payload, f)
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

// create custom audience : https://developers.facebook.com/docs/marketing-api/audiences/guides/custom-audiences
func (f *facebookEntry) CreateCustomAudience(req *FBModel.CreateCustomAudienceReq) (*FBModel.CreateCustomAudienceRes, error) {
	url := "https://graph.facebook.com/v21.0/" + f.actID + "/customaudiences"
	payload := map[string]string{
		"name":                 req.Name,
		"subtype":              "CUSTOM",
		"description":          req.Description,
		"customer_file_source": "USER_PROVIDED_ONLY",
	}
	return doHttp[FBModel.CreateCustomAudienceRes](url, "POST", payload, f)
}

func (f *facebookEntry) GetCustomAudience(req *FBModel.GetCustomAudienceReq) (*FBModel.GetCustomAudienceRes, error) {
	url := "https://graph.facebook.com/v21.0/" + req.ID
	payload := map[string]string{
		"fields": helpers.GetFieldMap(FBModel.GetCustomAudienceRes{}),
	}
	return doHttp[FBModel.GetCustomAudienceRes](url, "GET", payload, f)
}

// audiance targeting : https://developers.facebook.com/docs/marketing-api/audiences/overview
func (f *facebookEntry) AudienceTargeting() {
	// audience targeting

}

// audience rule : https://developers.facebook.com/docs/marketing-api/audiences/guides/audience-rules

// convert this to function call do smt
func doHttp[resT any](url, methods string, req map[string]string, f *facebookEntry) (*resT, error) {

	client := &http.Client{}
	request, err := http.NewRequest(methods, url, nil)
	if err != nil {
		return nil, err
	}

	reqQuery := request.URL.Query()
	for key, value := range req {
		reqQuery.Add(key, value)
	}

	request.URL.RawQuery = reqQuery.Encode()

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+f.accessToken)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("failed to get custom audience: %s", string(bodyBytes))
	}

	var res resT
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
