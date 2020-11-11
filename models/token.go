package models


type TokenDetails struct {
	AccessToken  string 
	RefreshToken string
	AccessUuid   string `json:"-"`
	RefreshUuid  string `json:"-"`
	AtExpires    int64	`json:"-"`
	RtExpires    int64  `json:"-"`
  }
  