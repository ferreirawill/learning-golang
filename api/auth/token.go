package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"gitlab.com/Ferreira.will/ormapi/models"
)

func CreateToken(userUuid string)(*models.TokenDetails, error){
	jwtToken := &models.TokenDetails{}
	var err error
	jwtToken.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	u,_ :=  uuid.NewV4()
	jwtToken.AccessUuid = u.String()

	jwtToken.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	u,_ = uuid.NewV4()
	jwtToken.RefreshUuid = u.String()




	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["access_id"] = jwtToken.AccessUuid
	atClaims["user_id"] = userUuid
	atClaims["exp"] = jwtToken.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256,atClaims)

	jwtToken.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))


	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = jwtToken.RefreshUuid
	rtClaims["user_id"] = userUuid
	rtClaims["exp"] = jwtToken.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256,rtClaims)

	jwtToken.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))


	if err != nil {
		return nil, err
	}
	return jwtToken,nil
}