package memberships

import (
	"database/sql"
	"testing"

	"github.com/NXRts/music-catalog/internal/configs"
	"github.com/NXRts/music-catalog/internal/models/memberships"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_service_SingUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.SingUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success create user",
			args: args{
				request: memberships.SingUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(memberships.User{}, sql.ErrNoRows)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "error get user",
			args: args{
				request: memberships.SingUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(memberships.User{}, assert.AnError)
			},
		},
		{
			name: "error Create user",
			args: args{
				request: memberships.SingUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(memberships.User{}, sql.ErrNoRows)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.SingUp(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.SingUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
