package memberships

import "github.com/NXRts/music-catalog/internal/models/memberships"

func (r *repository) CreateUser(models memberships.User) error {
	return r.db.Create(&models).Error
}

func (r *repository) GetUser(email, username string, id uint) (memberships.User, error) {
	user := memberships.User{}
	res := r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
