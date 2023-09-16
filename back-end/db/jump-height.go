package db

import "gorm.io/gorm"

type JumpHeight struct {
	gorm.Model `json:"-"`
	ID         uint   `gorm:"primaryKey" json:"id"`
	MinHeight  int    `json:"-"`
	MaxHeight  int    `json:"-"`
	Multiplier int    `json:"-"`
	Title      string `json:"title"`
}

func (db *DB) InsertJumpHeight(minHeight, maxHeight, multiplier int, title string) (int, error) {
	jumpHeight := JumpHeight{
		MinHeight:  minHeight,
		MaxHeight:  maxHeight,
		Multiplier: multiplier,
		Title:      title,
	}

	result := db.Create(&jumpHeight)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(jumpHeight.ID), nil
}

func (db *DB) GetJumpHeightById(id int) (*JumpHeight, error) {
	var jumpHeight JumpHeight
	result := db.Where("id = ?", id).First(&jumpHeight)
	if result.Error != nil {
		return nil, result.Error
	}

	return &jumpHeight, nil
}

func (db *DB) UpdateJumpHeight(id, minHeight, maxHeight, multiplier int, title string) error {
	result := db.Model(&JumpHeight{}).Where("id = ?", id).Updates(JumpHeight{MinHeight: minHeight, MaxHeight: maxHeight, Multiplier: multiplier, Title: title})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) DeleteJumpHeight(id int) error {
	result := db.Delete(&JumpHeight{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetAllJumpHeights() ([]*JumpHeight, error) {
	var jumpHeights []*JumpHeight
	result := db.Find(&jumpHeights)
	if result.Error != nil {
		return nil, result.Error
	}

	return jumpHeights, nil
}

func seedJumpHeights(db *DB) error {
	jumpHeights := []JumpHeight{
		{
			MinHeight:  0,
			MaxHeight:  95,
			Multiplier: 1,
			Title:      "Upto and with 95cm",
		},
		{
			MinHeight:  100,
			MaxHeight:  105,
			Multiplier: 5,
			Title:      "100/105cm",
		},
		{
			MinHeight:  110,
			MaxHeight:  115,
			Multiplier: 10,
			Title:      "110/115cm",
		},
		{
			MinHeight:  120,
			MaxHeight:  125,
			Multiplier: 18,
			Title:      "120/125cm",
		},
		{
			MinHeight:  130,
			MaxHeight:  135,
			Multiplier: 40,
			Title:      "130/135cm",
		},
		{
			MinHeight:  140,
			MaxHeight:  145,
			Multiplier: 85,
			Title:      "140/145cm",
		},
		{
			MinHeight:  150,
			MaxHeight:  155,
			Multiplier: 120,
			Title:      "150/155cm",
		},
		{
			MinHeight:  160,
			MaxHeight:  165,
			Multiplier: 170,
			Title:      "From 160cm",
		},
	}

	// Delete old heights
	if result := db.Exec("DELETE FROM jump_heights"); result.Error != nil {
		return result.Error
	}

	for _, jumpHeight := range jumpHeights {
		result := db.Create(&jumpHeight)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
