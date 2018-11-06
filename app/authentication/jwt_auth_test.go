package authentication_test

import (
	"app/authentication"
	"app/service"

	"app/models"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestAuthenticationMiddleware(t *testing.T) {

	userLogin1 := models.UserLogin{
		Email:    "test1@test1.com",
		Password: "asf22qw3",
	}

	signedToken1, err := authentication.SignedJwtToken(&userLogin1)

	if err != nil {
		t.Fatalf("Test failed with %v", err)
	}

	userLoginEmpty := models.UserLogin{
		Email:    "",
		Password: "",
	}

	signedTokenEmptyUser, err := authentication.SignedJwtToken(&userLoginEmpty)

	if err != nil {
		t.Fatalf("Test failed with %v", err)
	}

	tt := []struct {
		name               string
		user               *models.UserLogin
		signedToken        string
		authHeaderSet      bool
		expectedStatusCode int
	}{
		{
			name:               "Good token",
			user:               &userLogin1,
			signedToken:        signedToken1,
			authHeaderSet:      true,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Good token with empty user login",
			user:               &userLoginEmpty,
			signedToken:        signedTokenEmptyUser,
			authHeaderSet:      true,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Good token but Authorization header is required",
			user:               &userLogin1,
			signedToken:        signedToken1,
			authHeaderSet:      false,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:               "Good token with empty user login but Authorization header is required",
			user:               &userLoginEmpty,
			signedToken:        signedTokenEmptyUser,
			authHeaderSet:      false,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:               "Bad token",
			user:               &userLogin1,
			signedToken:        "dummytoken",
			authHeaderSet:      true,
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range tt {
		//test request
		req, err := http.NewRequest("GET", "/authindex", nil)
		if err != nil {
			t.Fatalf("Test '%s' failed with %v", tc.name, err)
		}

		// add the token to the request headers
		if tc.authHeaderSet {
			req.Header.Add("Authorization", "Bearer "+tc.signedToken)
		}
		// record test request
		rr := httptest.NewRecorder()

		router := httprouter.New()
		router.Handle("GET", "/authindex", authentication.AuthenticationMiddleware(service.Index))

		router.ServeHTTP(rr, req)
		if status := rr.Code; status != tc.expectedStatusCode {
			t.Errorf("Test '%s' failed with: handler did not return %d: got %d", tc.name, tc.expectedStatusCode, status)
		}
	}

}
