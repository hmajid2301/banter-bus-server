package controllers_test

import (
	"fmt"
	"net/http"
	"testing"

	serverModels "banter-bus-server/src/server/models"
	"banter-bus-server/tests/data"

	"github.com/gavv/httpexpect"
)

func (s *Tests) SubTestAddUser(t *testing.T) {
	for _, tc := range data.AddUser {
		testName := fmt.Sprintf("Add New User: %s", tc.TestDescription)
		t.Run(testName, func(t *testing.T) {
			s.httpExpect.POST("/user").
				WithJSON(tc.Payload).
				Expect().
				Status(tc.ExpectedStatus)
		})
	}
}

func (s *Tests) SubTestGetUser(t *testing.T) {
	for _, tc := range data.GetUser {
		testName := fmt.Sprintf("Get User: %s", tc.TestDescription)
		t.Run(testName, func(t *testing.T) {
			getUser(tc.Username, tc.ExpectedStatus, tc.ExpectedUser, s.httpExpect)
		})
	}
}

func (s *Tests) SubTestGetAllUsers(t *testing.T) {
	for _, tc := range data.GetAllUsers {
		testName := fmt.Sprintf("Get All Users: %s", tc.TestDescription)
		t.Run(testName, func(t *testing.T) {
			response := s.httpExpect.GET("/user").WithQueryObject(tc.Filter)
			response.
				Expect().
				Status(http.StatusOK).JSON().Array().Equal(tc.ExpectedUsernames)
		})
	}
}

func getUser(user string, expectedStatus int, expectedResult serverModels.User, httpExpect *httpexpect.Expect) {
	endpoint := fmt.Sprintf("/user/%s", user)
	response := httpExpect.GET(endpoint).
		Expect().
		Status(expectedStatus)

	if expectedStatus == http.StatusOK {
		response.JSON().Object().Equal(expectedResult)
	}
}
