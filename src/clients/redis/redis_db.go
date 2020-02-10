package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/narinjtp/bookstore_users-api/utils/errors"
)
var(
	redisDb *redis.Pool
)
func GetConnection() (redis.Conn, *errors.RestErr){
	//var conn redis.Conn
	var pool *redis.Pool
	pool, err := newPool()
	if err != nil {
		return nil, err
	}
	return pool.Get(), nil
	// Send PING command to Redis
	//pong, err := c.Do("PING")
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//
	//// PING command returns a Redis "Simple String"
	//// Use redis.String to convert the interface type to string
	//s, err := redis.String(pong, err)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//
	//fmt.Printf("Connection redis success; PING Response = %s\n", s)
	// Output: PONG
}

func newPool() (*redis.Pool, *errors.RestErr){
	//var pool redis.Pool
	//
	//pool.MaxIdle = 80
	//pool.MaxActive = 12000
	//pool.Dial
	var poolErr *errors.RestErr
	pool := redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				fmt.Println(err.Error())
				poolErr = errors.NewInternalServerError(err.Error())
				//panic(err.Error())
				//return nil, errors.NewInternalServerError(err.Error)
			}
			return c, err
		},
	}
	if poolErr != nil {
		return nil, poolErr
	}
	return &pool, nil
}