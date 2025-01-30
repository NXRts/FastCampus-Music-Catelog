package memberships

import (
	"database/sql"
	"errors"

	"github.com/NXRts/music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SingUp(request memberships.SingUpRequest) error {
	existingUser, err := s.repository.GetUser(request.Email, request.Username, 0)
	if err != nil && err != sql.ErrNoRows {
		log.Error().Err(err).Msg("error get user")
		return err
	}

	if existingUser.ID != 0 {
		return errors.New("email or username exist")
	}

	// if existingUser != nil {
	// return errors.New("email or username exist")
	// }

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("error hash password")
		return err
	}

	model := memberships.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  string(pass),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}

	return s.repository.CreateUser(model)
}
