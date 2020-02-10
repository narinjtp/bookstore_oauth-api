package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@mail.com","password":"password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("test@mail.com","password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Messsage)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@mail.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status":"404", "error":"not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("test@mail.com","password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Messsage)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@mail.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status":"404", "error":"not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("test@mail.com","password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Messsage)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@mail.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"id": "1","first_name": "narin","last_name": "ja","email": "narin.jtp@gmail.com"}`,
	})


	repository := usersRepository{}

	user, err := repository.LoginUser("test@mail.com","password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal users login response", err.Messsage)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@mail.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1,"first_name": "narin","last_name": "ja","email": "narin.jtp@gmail.com"}`,
	})


	repository := usersRepository{}

	user, err := repository.LoginUser("test@mail.com","password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "narin", user.FirstName)
	assert.EqualValues(t, "ja", user.LastName)
	assert.EqualValues(t, "narin.jtp@gmail.com", user.Email)
}