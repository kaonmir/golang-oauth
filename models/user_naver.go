package models

var UserNaverModel = new(UserNaver)

type UserNaver struct {
	SNSID          string `json:"sns_id"`
	SNSType        string `json:"sns_type"`
	SNSName        string `json:"sns_name"`
	SNSProfile     string `json:"sns_profile"`
	SNSConnectDate string `json:"sns_connect_date"`
}
