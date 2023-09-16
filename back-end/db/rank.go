package db

import "gorm.io/gorm"

type Rank struct {
	gorm.Model `json:"-"`
	ID         uint   `gorm:"primaryKey" json:"id"`
	Title      string `gorm:"unique" json:"title"`
	Point      int    `json:"-"`
}

func (db *DB) InsertRank(title string, point int) (int, error) {
	rank := Rank{
		Title: title,
		Point: point,
	}

	result := db.Create(&rank)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(rank.ID), nil
}

func (db *DB) GetRankById(id int) (*Rank, error) {
	var rank Rank
	result := db.Where("id = ?", id).First(&rank)
	if result.Error != nil {
		return nil, result.Error
	}

	return &rank, nil
}

func (db *DB) UpdateRank(id int, title string, point int) error {
	result := db.Model(&Rank{}).Where("id = ?", id).Updates(Rank{Title: title, Point: point})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) DeleteRank(id int) error {
	result := db.Delete(&Rank{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetAllRanks() ([]*Rank, error) {
	var ranks []*Rank
	result := db.Find(&ranks)
	if result.Error != nil {
		return nil, result.Error
	}

	return ranks, nil
}

func seedRanks(db *DB) error {
	ranks := []Rank{
		{
			Title: "First",
			Point: 11,
		},
		{
			Title: "Second",
			Point: 9,
		},
		{
			Title: "Third",
			Point: 8,
		},
		{
			Title: "Fourth",
			Point: 7,
		},
		{
			Title: "Fifth",
			Point: 6,
		},
		{
			Title: "Sixth",
			Point: 5,
		},
		{
			Title: "Seventh",
			Point: 4,
		},
		{
			Title: "Eighth",
			Point: 3,
		},
		{
			Title: "Ninth",
			Point: 3,
		},
		{
			Title: "Tenth",
			Point: 3,
		},
	}

	// Delete all ranks
	if result := db.Exec("DELETE FROM ranks"); result.Error != nil {
		return result.Error
	}

	// `result.RowsAffected` will give you the number of rows affected by the query

	// Insert ranks
	for _, rank := range ranks {
		result := db.Create(&rank)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
