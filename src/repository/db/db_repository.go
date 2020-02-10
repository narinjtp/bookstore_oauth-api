package db

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	redis2 "github.com/narinjtp/bookstore_oauth-api/src/clients/redis"
	"github.com/narinjtp/bookstore_oauth-api/src/domain/access_token"
	"github.com/narinjtp/bookstore_users-api/utils/errors"
)

func NewDbRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {

}

func (r* dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	const objectPrefix string = "tokenId:"
	c, connErr := redis2.GetConnection()
	if connErr != nil {
		return nil, errors.NewInternalServerError(connErr.Error)
	}
	defer c.Close()
	s, err := redis.String(c.Do("GET", objectPrefix+id))
	if err == redis.ErrNil {
		fmt.Println("session id does not exist")
		//return nil, errors.NewInternalServerError(err.Error())
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	at := access_token.AccessToken{}
	err = json.Unmarshal([]byte(s), &at)

	fmt.Printf("%+v\n", at)

	//return nil
	//redisDb := redis.NewPool()
	return &at, nil
}

func (r* dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	const objectPrefix string = "tokenId:"
	c, connErr := redis2.GetConnection()
	if connErr != nil {
		return errors.NewInternalServerError(connErr.Error)
	}
	defer c.Close()

	// serialize User object to JSON
	json, err := json.Marshal(at)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	// SET object
	_, err = c.Do("SET", objectPrefix+at.AccessToken, json)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r* dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	const objectPrefix string = "tokenId:"
	c, connErr := redis2.GetConnection()
	if connErr != nil {
		return errors.NewInternalServerError(connErr.Error)
	}
	defer c.Close()

	// serialize User object to JSON
	json, err := json.Marshal(at)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	// SET object
	_, err = c.Do("SET", objectPrefix+at.AccessToken, json)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}