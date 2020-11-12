package auth

import (
	"fmt"

	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"gitlab.com/Ferreira.will/ormapi/models"
)
	
type AccessDetails struct {
    AccessUuid string
    UserUuid   string
}


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
	atClaims["access_uuid"] = jwtToken.AccessUuid
	atClaims["user_uuid"] = userUuid
	atClaims["exp"] = jwtToken.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256,atClaims)

	jwtToken.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))


	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = jwtToken.RefreshUuid
	rtClaims["user_uuid"] = userUuid
	rtClaims["exp"] = jwtToken.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256,rtClaims)

	jwtToken.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))


	if err != nil {
		return nil, err
	}
	return jwtToken,nil
}

func ExtractToken(r *http.Request) string{
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func VerifyToken(r * http.Request)(*jwt.Token,error){
	tokenString := ExtractToken(r)
	jwtToken ,err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",t.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")),nil
	})

	if err != nil{
		return nil, err
	}

	return jwtToken, nil
}

func TokenValid(r *http.Request)error{
 jwtToken, err := VerifyToken(r)
 if err != nil{
	 return err
 }

 if _,ok := jwtToken.Claims.(jwt.Claims); !ok && !jwtToken.Valid {
	 return err
 }

 return nil

}

func ExtractTokenMetadata(r *http.Request)(*AccessDetails,error){
	jwtToken, err := VerifyToken(r)

	if err !=nil {
		return nil ,err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if ok && jwtToken.Valid {
		accessUuid, ok := claims["access_uuid"].(string)

		if !ok {
			return nil, err
		}

		userUuid := claims["user_uuid"].(string)

		return &AccessDetails{
			AccessUuid: accessUuid,
			UserUuid: userUuid,
		},nil
	}



	return nil, err

}

