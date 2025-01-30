package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NXRts/music-catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name             string
		mockFn           func()
		expectStatusCode int
	}{
		{
			name: "Success User Signed Up",
			mockFn: func() {
				mockSvc.EXPECT().SingUp(memberships.SingUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(nil)
			},
			expectStatusCode: 201,
		},
		{
			name: "Failed User Signed Up",
			mockFn: func() {
				mockSvc.EXPECT().SingUp(memberships.SingUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(errors.New("failed to sign up"))
			},
			expectStatusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoutes()
			w := httptest.NewRecorder()

			endpoint := `/memberships/sign_up`
			jsonReq := memberships.SingUpRequest{
				Email:    "test@gmail.com",
				Username: "testusername",
				Password: "password",
			}

			val, err := json.Marshal(jsonReq)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			httpReq, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(w, httpReq)

			assert.Equal(t, tt.expectStatusCode, w.Code)
		})
	}
}
