package db

import "gorm.io/gorm"

type Horse struct {
	gorm.Model `json:"-"`
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"unique" json:"name" validate:"required,min=2,max=100"`
	Color      string `json:"color" validate:"required,min=2,max=100"`
	Age        int    `json:"age" validate:"required,min=1,max=100"`

	Owner   User `gorm:"foreignKey:OwnerID;constraint:OnDelete:Set Null" json:"-"`
	OwnerID int  `gorm:"index" json:"-"`
}

func (db *DB) InsertHorse(horse *Horse) (int, error) {

	result := db.Create(horse)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(horse.ID), nil
}

func (db *DB) GetHorseById(id int, ownerID int) (*Horse, error) {
	var horse Horse
	result := db.Where("id = ?", id).
		Where("owner_id = ?", ownerID).
		First(&horse)

	if result.Error != nil {
		return nil, result.Error
	}

	return &horse, nil
}

func (db *DB) UpdateHorse(id int, name, color string, age,ownerID int) error {
	result := db.Model(&Horse{}).
		Where("id = ?", id).
		Where("owner_id = ?", ownerID).
		Updates(Horse{Name: name, Color: color, Age: age})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) DeleteHorse(id, ownerID int) error {
	result := db.Delete(&Horse{}, id, ownerID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetAllHorses(ownerID int) ([]*Horse, error) {
	var horses []*Horse
	result := db.Where("owner_id = ?", ownerID).
		// Preload("Owner").
		Find(&horses)
	if result.Error != nil {
		return nil, result.Error
	}

	return horses, nil
}
