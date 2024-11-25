package tests

// import (
// 	FBModel "default-repo/libs/httpEntry/models"
// 	"default-repo/libs/httpEntry/usecases"
// 	"net/http"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/jarcoal/httpmock"
// )

// func TestCreateCustomAudience(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockViper := mocks.NewMockViper(mockCtrl)
// 	mockViper.EXPECT().GetString("facebook.accessToken").Return("mockAccessToken")
// 	mockViper.EXPECT().GetString("facebook.actID").Return("mockActID")

// 	facebookEntry := usecases.NewFacebookEntry()

// 	req := &FBModel.CreateCustomAudienceReq{
// 		Name:        "Test Audience",
// 		Description: "Test Description",
// 	}

// 	httpClient := &http.Client{}
// 	httpmock.ActivateNonDefault(httpClient)
// 	defer httpmock.DeactivateAndReset()

// 	httpmock.RegisterResponder("POST", "https://graph.facebook.com/v21.0/act_mockActID/customaudiences",
// 		func(req *http.Request) (*http.Response, error) {
// 			res := httpmock.NewStringResponse(200, `{"id": "12345"}`)
// 			return res, nil
// 		},
// 	)

// 	res, err := facebookEntry.CreateCustomAudience(req)
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if res.ID != "12345" {
// 		t.Errorf("expected ID to be '12345', got %v", res.ID)
// 	}
// }
