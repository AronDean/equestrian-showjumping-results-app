package db

import "gorm.io/gorm"

type User struct {
	gorm.Model     `json:"-"`
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
}

func (db *DB) InsertUser(email, name, hashedPassword string) (*User, error) {
	user := User{
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (db *DB) GetUserById(id int) (*User, error) {
	var user User
	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (db *DB) UpdateUserHashedPassword(id int, hashedPassword string) error {
	result := db.Model(&User{}).Where("id = ?", id).Update("hashed_password", hashedPassword)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) DeleteUser(id int) error {
	result := db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetAllUsers() ([]*User, error) {
	var users []*User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
