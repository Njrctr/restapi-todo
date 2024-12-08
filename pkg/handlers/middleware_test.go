package handler

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/Njrctr/restapi-todo/pkg/service"
	mock_service "github.com/Njrctr/restapi-todo/pkg/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_userIdentify(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAutorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAutorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `1`,
		},
		{
			name:                 "No Header",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(s *mock_service.MockAutorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"message\":\"Empty auth header\"}\n",
		},
		{
			name:                 "Invalid Bearer",
			headerName:           "Authorization",
			headerValue:          "Bearrer token",
			token:                "token",
			mockBehavior:         func(s *mock_service.MockAutorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"message\":\"Invalid auth header\"}\n",
		},
		{
			name:                 "Invalid Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			mockBehavior:         func(s *mock_service.MockAutorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"message\":\"Token is empty\"}\n",
		},
		{
			name:        "Service Failure",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAutorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"message\":\"invalid token\"}\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mock_service.NewMockAutorization(controller)
			testCase.mockBehavior(repo, testCase.token)

			services := &service.Service{Autorization: repo}
			handler := NewHandler(services)

			//Test Server
			router := gin.New()
			router.GET("/protected", handler.userIdentify, func(ctx *gin.Context) {
				id, _ := ctx.Get(userCtx)
				ctx.String(200, fmt.Sprintf("%d", id.(int)))
			})

			// Test Request
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			// Make Request
			router.ServeHTTP(rec, req)

			// Assert
			assert.Equal(t, rec.Code, testCase.expectedStatusCode)
			assert.Equal(t, rec.Body.String(), testCase.expectedResponseBody)
		})
	}
}

func TestGetUserId(t *testing.T) {
	var getContext = func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Set(userCtx, id)
		return ctx
	}

	testTable := []struct {
		name       string
		ctx        *gin.Context
		id         int
		shouldFail bool
	}{
		{
			name: "Ok",
			ctx:  getContext(1),
			id:   1,
		},
		{
			name:       "Empty",
			ctx:        &gin.Context{},
			shouldFail: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			id, err := getUserId(test.ctx)
			if test.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, id, test.id)
		})
	}
}
