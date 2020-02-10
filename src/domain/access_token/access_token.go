package access_token

import (
	"fmt"
	"github.com/narinjtp/bookstore_users-api/utils/errors"
	"strings"
	"time"
)

const(
	expirationTime = 24
	grantTypePassword = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope string `json:"scope"`

	//Used for password grant type
	Username string `json:"username`
	Password string `json:"password"`

	//Used fot client_credentials grant type
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId int64 `json:"user_id"`
	ClientId int64 `json:"client_id"`
	Expires int64 `json:"expires"`
}
//type Repository interface {
//	GetById(string) (*AccessToken, *errors.RestErr)
//}
func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}
	 //if at.GrantType != grantTypePassword || at.GrantType != grantTypeClientCredentials{
		//return errors.NewBadRequestError("invalid grant_type parameter")
	 //}
	 //TODO: Validate parameters for each grant_type
	 return nil
}
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == ""{
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0{
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0{
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires:time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	fmt.Println(expirationTime)
	return expirationTime.Before(now)
}