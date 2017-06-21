package daos

import (
	"log"

	"github.com/mammenj/goboot/models"
)

// StaticUserImpl static implementation of API
type StaticUserImpl struct {
}

// Create static implementation of API
func (dao StaticUserImpl) Create(u *models.User) error {
	log.Println("Created User ....", u.Name)
	return nil
}

// GetAll static implementation of API
func (dao StaticUserImpl) GetAll() ([]models.User, error) {
	users := make([]models.User, 0)
	log.Println("GetAll User ....")
	return users, nil
}

// Delete static implementation of API
func (dao StaticUserImpl) Delete(id int) error {
	log.Printf("... Deleted User ....%d", id)
	return nil
}

// Get static implementation of API
func (dao StaticUserImpl) Get(id int) (models.User, error) {
	log.Printf("Getting User ....%d", id)
	u := models.User{Name: "john", Age: 30, Gender: "Male"}
	return u, nil
}

// Update static implementation of API
func (dao StaticUserImpl) Update(u *models.User) error {
	log.Println("Updated User ....", u.Name)
	return nil
}
