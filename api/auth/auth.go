package auth

import (

	"time"

	"github.com/go-redis/redis/v7"
	"gitlab.com/Ferreira.will/ormapi/models"
)

func CreateAuth(userUuid string, jwtToken *models.TokenDetails,redisClient *redis.Client) error{
	at := time.Unix(jwtToken.AtExpires,0)
	rt := time.Unix(jwtToken.RtExpires,0)

	now := time.Now()

	errAccess := redisClient.Set(jwtToken.AccessUuid,userUuid,at.Sub(now)).Err()

	if errAccess != nil {
		return errAccess
	}
	errRefresh := redisClient.Set(jwtToken.RefreshUuid,userUuid,rt.Sub(now)).Err()

	if errRefresh != nil {
		return errRefresh
	}

	return nil

}