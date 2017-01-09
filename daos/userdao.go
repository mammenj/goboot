package daos

import "github.com/mammenj/goboot/models"

type UserDao interface {
	Get(i int) (models.User, error)
	GetAll() ([]models.User, error)
	Create(u *models.User) error
	Delete(i int) error
	Update(u *models.User) error
}
