package models

// facebook create camppaign request and response
type (
	CreateCampaignReq struct {
		Name                string
		Objective           string
		Status              string
		SpecialAdCategories []string
	}
	CreateCampaignRes struct {
		ID string `json:"id"`
	}
)

// facebook get campaign request and response
type (
	GetCampaignReq struct {
		ID string
	}
	GetCampaignRes struct {
		ID                  string   `json:"id"`
		AccountID           string   `json:"account_id"`
		Objective           string   `json:"objective"`
		Name                string   `json:"name"`
		ConfiguredStatus    string   `json:"configured_status"`
		EffectiveStatus     string   `json:"effective_status"`
		BuyingType          string   `json:"buying_type"`
		CreatedTime         string   `json:"created_time"`
		UpdatedTime         string   `json:"updated_time"`
		SpendCap            string   `json:"spend_cap"`
		CanUseSpendCap      bool     `json:"can_use_spend_cap"`
		IssuesInfo          []string `json:"issues_info"`
		SpecialAdCategories []string `json:"special_ad_categories"`
	}
)

type (
	DefineTargetReq struct {
		Name             string
		OptimizationGoal string
		BillingEvent     string
		BidAmount        int
		DailyBudget      int
		CampaignID       string
		Targeting        struct {
			GeoLocations struct {
				Countries []string
			}
			FacebookPositions []string
		}
		Status         string
		PromotedObject struct {
			PageID string
		}
	}

	DefineTargetRes struct {
	}
)

// facebook create custom audience
type (
	CreateCustomAudienceReq struct {
		Name               string
		Subtype            string
		Description        string
		CustomerFileSource string
	}
	CreateCustomAudienceRes struct {
		ID string `json:"id"`
	}
)

// facebook get custom audience
type (
	GetCustomAudienceReq struct {
		ID string
	}

	GetCustomAudienceRes struct {
		ID                 string `json:"id"`
		AccountID          string `json:"account_id"`
		CustomerFileSource string `json:"customer_file_source"`
		DataSource         struct {
			Type           string `json:"type"`
			SubType        string `json:"sub_type"`
			CreationParams string `json:"creation_params"`
		} `json:"data_source"`
		Description string `json:"description"`
	}
)
