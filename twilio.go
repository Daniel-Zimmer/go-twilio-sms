package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const urlRoot = "https://api.twilio.com/2010-04-01/Accounts/"

type Service struct {
	accountSid string
	authToken  string
	url        string
}

func NewService(AccountSid, AuthToken string) *Service {
	return &Service{
		accountSid: AccountSid,
		authToken:  AuthToken,
		url: urlRoot + AccountSid + "/Messages.json",
	}
}

func (service Service) SendSms(from, to, text string) error {
	message := url.Values{}
	message.Set("From", from)
	message.Set("To", to)
	message.Set("Body", text)

	req, _ := http.NewRequest("POST", service.url, strings.NewReader(message.Encode()))
	req.SetBasicAuth(service.accountSid, service.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := (&http.Client{}).Do(req)

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return nil
	}

	var body map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&body)

	return fmt.Errorf("%s\n", body["message"])
}
