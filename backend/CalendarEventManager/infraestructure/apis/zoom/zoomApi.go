package zoom

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
)

type ZoomAPI struct {
	ClientId     string
	ClientSecret string
	Url          string
	AccountId    string
}

func NewZoomApi(envs *config.Env) *ZoomAPI {
	return &ZoomAPI{ClientId: envs.ZOOM_CLIENT_ID,
		ClientSecret: envs.ZOOM_CLIENT_SECRET,
		Url:          envs.ZOOM_URL,
		AccountId:    envs.ZOOM_ACCOUNT_ID}
}

func (api *ZoomAPI) GenerateAccessToken() *entities.AccessTokenResponse {
	url := fmt.Sprintf(`%s/oauth/token`, api.Url)
	method := "POST"
	credentialsBasic := fmt.Sprintf(`%s:%s`, api.ClientId, api.ClientSecret)
	encodeCredentials := base64.StdEncoding.EncodeToString([]byte(credentialsBasic))
	payload := strings.NewReader(fmt.Sprintf(`grant_type=account_credentials&account_id=%s`, api.AccountId))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf(`Basic %s`, encodeCredentials))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var accessToken entities.AccessTokenResponse
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &accessToken
}

func (api *ZoomAPI) CreateZoomMeeting(token string, userId string, meeting entities.MeetingResponse) *entities.MeetingResponse {

	url := fmt.Sprintf(`%s/v2/users/%s/meetings`, api.Url, userId)
	method := "POST"
	data, err := json.Marshal(meeting)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	payload := strings.NewReader(string(data))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var meetingResponse entities.MeetingResponse
	err = json.Unmarshal(body, &meetingResponse)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &meetingResponse
}
