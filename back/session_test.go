package main

import (
	"GameCaster/main/sqlobjects"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

const SESSION_TEST_PORT = "13456"

type SessionTestSuite struct {
	suite.Suite
	database *gorm.DB
}

func (suite *SessionTestSuite) SetupSuite() {
	// Create test database
	suite.database = OpenDB("test_session.db")
	// Start server in background
	go Serve(mux.NewRouter(), suite.database, SESSION_TEST_PORT)
	time.Sleep(15 * time.Millisecond)
}

func TestSessionSuite(t *testing.T) {
	suite.Run(t, new(SessionTestSuite))
}

func (suite *SessionTestSuite) Post(path string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return http.Post("http://localhost:"+SESSION_TEST_PORT+path, contentType, body)
}

func (suite *SessionTestSuite) Get(path string) (resp *http.Response, err error) {
	return http.Get("http://localhost:" + SESSION_TEST_PORT + path)
}

func (suite *SessionTestSuite) GetWithCookie(path string, cookie *http.Cookie) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:"+SESSION_TEST_PORT+path, nil)
	req.AddCookie(cookie)
	return client.Do(req)
}

func (suite *SessionTestSuite) PostWithCookie(path string, contentType string, body io.Reader, cookie *http.Cookie) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:"+SESSION_TEST_PORT+path, body)
	req.Header.Set("Content-Type", contentType)
	req.AddCookie(cookie)
	return client.Do(req)
}

func (suite *SessionTestSuite) TestUserRoleSelection() {
	// Create a test user
	user, err := sqlobjects.CreateUser("roletest@example.com", "password123", suite.database)
	suite.Nil(err)
	suite.NotNil(user)

	// Create a session token
	token, err := sqlobjects.CreateSessionToken(user, suite.database)
	suite.Nil(err)
	suite.NotEmpty(token)

	// Test setting role to DM
	reqBody := map[string]string{"role": "DM"}
	jsonBody, _ := json.Marshal(reqBody)

	cookie := &http.Cookie{Name: "authToken", Value: token}
	resp, err := suite.PostWithCookie("/api/authenticated/user/role", "application/json", bytes.NewBuffer(jsonBody), cookie)
	suite.Nil(err)
	defer resp.Body.Close()

	suite.Equal(http.StatusOK, resp.StatusCode)

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	suite.Equal("success", response["status"])
	suite.Equal("DM", response["role"])

	// Verify the role was saved in the database
	updatedUser, err := sqlobjects.GetUserByID(int(user.ID), suite.database)
	suite.Nil(err)
	suite.Equal(sqlobjects.RoleDM, updatedUser.Role)
}

func (suite *SessionTestSuite) TestUserRoleSelectionInvalidRole() {
	// Create a test user
	user, err := sqlobjects.CreateUser("invalidroletest@example.com", "password123", suite.database)
	suite.Nil(err)

	// Create a session token
	token, err := sqlobjects.CreateSessionToken(user, suite.database)
	suite.Nil(err)

	// Test setting invalid role
	reqBody := map[string]string{"role": "InvalidRole"}
	jsonBody, _ := json.Marshal(reqBody)

	cookie := &http.Cookie{Name: "authToken", Value: token}
	resp, err := suite.PostWithCookie("/api/authenticated/user/role", "application/json", bytes.NewBuffer(jsonBody), cookie)
	suite.Nil(err)
	defer resp.Body.Close()

	suite.Equal(http.StatusBadRequest, resp.StatusCode)
}

func (suite *SessionTestSuite) TestUserRoleSelectionAllRoles() {
	roles := []sqlobjects.UserRole{
		sqlobjects.RoleDM,
		sqlobjects.RolePlayer,
		sqlobjects.RoleCasting,
	}

	for _, expectedRole := range roles {
		// Create a test user for each role
		user, err := sqlobjects.CreateUser(string(expectedRole)+"test@example.com", "password123", suite.database)
		suite.Nil(err)

		// Create a session token
		token, err := sqlobjects.CreateSessionToken(user, suite.database)
		suite.Nil(err)

		// Test setting the role
		reqBody := map[string]string{"role": string(expectedRole)}
		jsonBody, _ := json.Marshal(reqBody)

		cookie := &http.Cookie{Name: "authToken", Value: token}
		resp, err := suite.PostWithCookie("/api/authenticated/user/role", "application/json", bytes.NewBuffer(jsonBody), cookie)
		suite.Nil(err)
		defer resp.Body.Close()

		suite.Equal(http.StatusOK, resp.StatusCode)

		// Verify the role was saved
		updatedUser, err := sqlobjects.GetUserByID(int(user.ID), suite.database)
		suite.Nil(err)
		suite.Equal(expectedRole, updatedUser.Role)
	}
}

func (suite *SessionTestSuite) TestGetUserInfo() {
	// Create a test user with a role
	user, err := sqlobjects.CreateUser("userinfo@example.com", "password123", suite.database)
	suite.Nil(err)

	// Set the role
	err = sqlobjects.UpdateUserRole(user.ID, sqlobjects.RolePlayer, suite.database)
	suite.Nil(err)

	// Create a session token
	token, err := sqlobjects.CreateSessionToken(user, suite.database)
	suite.Nil(err)

	// Test getting user info
	cookie := &http.Cookie{Name: "authToken", Value: token}
	resp, err := suite.GetWithCookie("/api/authenticated/user", cookie)
	suite.Nil(err)
	defer resp.Body.Close()

	suite.Equal(http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	suite.Equal("userinfo@example.com", response["email"])
	suite.Equal("Player", response["role"])
}

func (suite *SessionTestSuite) TestGetUserInfoUnauthorized() {
	// Test getting user info without authentication
	resp, err := suite.Get("/api/authenticated/user")
	suite.Nil(err)
	defer resp.Body.Close()

	suite.Equal(http.StatusUnauthorized, resp.StatusCode)
}
