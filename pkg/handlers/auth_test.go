package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	todo "github.com/Njrctr/restapi-todo/models"
	"github.com/Njrctr/restapi-todo/pkg/service"
	mock_service "github.com/Njrctr/restapi-todo/pkg/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAutorization, user todo.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           todo.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"TestName", "username":"testusername", "password":"testpassword"}`,
			inputUser: todo.User{
				Name:     "TestName",
				Username: "testusername",
				Password: "testpassword",
			},
			mockBehavior: func(s *mock_service.MockAutorization, user todo.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "{\"id\":1}\n",
		},
		{
			name:                "Empty Fields",
			inputBody:           `{ "username":"testusername", "password":"testpassword"}`,
			mockBehavior:        func(s *mock_service.MockAutorization, user todo.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: "{\"message\":\"invalid input body\"}\n",
		},
		{
			name:      "Service Failure",
			inputBody: `{"name":"TestName", "username":"testusername", "password":"testpassword"}`,
			inputUser: todo.User{
				Name:     "TestName",
				Username: "testusername",
				Password: "testpassword",
			},
			mockBehavior: func(s *mock_service.MockAutorization, user todo.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: "{\"message\":\"service failure\"}\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			controller := gomock.NewController(t)
			defer controller.Finish()

			auth := mock_service.NewMockAutorization(controller)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Autorization: auth}
			handler := NewHandler(services)

			//Test Server
			router := gin.New()
			router.POST("/sign_up", handler.signUp)

			//Test Request
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign_up", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			router.ServeHTTP(rec, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, rec.Code)
			assert.Equal(t, testCase.expectedRequestBody, rec.Body.String())
		})
	}
}
