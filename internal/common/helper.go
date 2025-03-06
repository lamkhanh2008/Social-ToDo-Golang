package core

import (
	"time"
)

type ItemAPICaller interface {
	GetServiceURL() string
}

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
}

func GetRequesterID(requester Requester) int {
	uid, _ := UIDFromString(requester.GetUID())
	return int(uid.GetLocalID())
}
