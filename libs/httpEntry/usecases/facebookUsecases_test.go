package usecases

import (
	"default-repo/helpers"
	FBModel "default-repo/libs/httpEntry/models"
	"fmt"
	"testing"
)

func Test_facebookEntry_Init(t *testing.T) {
	helpers.InitConfig()
	f := NewFacebookEntry()
	f.Init()
	tests := []struct {
		name    string
		req     FBModel.CreateCampaignReq
		res     FBModel.CreateCampaignRes
		wantErr bool
	}{
		{
			name: "test success create campaign",
			req: FBModel.CreateCampaignReq{
				Name:                "test create campaign facebook",
				Objective:           "OUTCOME_APP_PROMOTION",
				Status:              "",
				SpecialAdCategories: []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := f.CreateCampaign(&tt.req)
			if res != nil {
				fmt.Println("SUCCESS")
			}
		})
	}
}
