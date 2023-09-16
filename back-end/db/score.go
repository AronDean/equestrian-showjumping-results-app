package db

import (
	"time"

	"gorm.io/gorm"
)

type Score struct {
	gorm.Model `json:"-"`

	ID uint `gorm:"primaryKey" json:"id"`

	Horse   Horse `gorm:"foreignKey:HorseID;constraint:OnDelete:Set Null" json:"horse"`
	HorseID int   `gorm:"index" json:"-"`

	JumpHeight   JumpHeight `gorm:"foreignKey:JumpHeightID;constraint:OnDelete:Set Null" json:"jump_height"`
	JumpHeightID int        `gorm:"index" json:"-"`

	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:Set Null" json:"-"`
	UserID int  `gorm:"index" json:"-"`

	Rank   Rank `gorm:"foreignKey:RankID;constraint:OnDelete:Set Null" json:"rank"`
	RankID int  `gorm:"index" json:"-"`

	PointsEarned int       `json:"points_earned"`
	CreatedAt    time.Time `json:"created_at"`
}

func (db *DB) InsertScore(score *Score) (int, error) {
	result := db.Create(&score)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(score.ID), nil
}

func (db *DB) GetScoreById(id int) (*Score, error) {
	var score Score
	result := db.Where("id = ?", id).
		Preload("Horse").
		First(&score)
	if result.Error != nil {
		return nil, result.Error
	}

	return &score, nil
}

func (db *DB) UpdateScore(id, horseID, jumpHeightID, userID, rankID, pointsEarned int) error {
	result := db.Model(&Score{}).Where("id = ?", id).Updates(Score{HorseID: horseID, JumpHeightID: jumpHeightID, UserID: userID, RankID: rankID, PointsEarned: pointsEarned})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) DeleteScore(id int) error {
	result := db.Delete(&Score{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetAllScores(page int, pageSize int, horseID int, jumpHeightID int, userID int, rankID int) (map[string]interface{}, error) {
	var scores []Score
	var count int64

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	var query *gorm.DB

	if userID != 0 {
		query = db.Model(Score{}).Where("user_id = ?", userID)
	}

	if horseID != 0 {
		query = query.Where("horse_id = ?", horseID)
	}
	if jumpHeightID != 0 {
		query = query.Where("jump_height_id = ?", jumpHeightID)
	}

	if rankID != 0 {
		query = query.Where("rank_id = ?", rankID)
	}

	err := query.Count(&count).Error
	if err != nil {
		return nil, err
	}

	query = query.Offset(offset).Limit(pageSize)
	query = query.Order("created_at desc")

	err = query.Preload("Horse").
		Preload("User").
		Preload("JumpHeight").
		Preload("Rank").
		Find(&scores).Error
	if err != nil {
		return nil, err
	}

	scoreData := map[string]interface{}{
		"count":  count,
		"scores": scores,
	}

	return scoreData, nil
}

func (db *DB) GetUsersByPointsEarned(startTime time.Time, endTime time.Time) ([]UserPoints, error) {
	var userPoints []UserPoints

	err := db.Table("users").
		Select("users.name, users.email, COALESCE(SUM(scores.points_earned), 0) as points_earned, COALESCE(COUNT(scores.id), 0) as game_count").
		Joins("LEFT JOIN scores ON users.id = scores.user_id AND scores.created_at BETWEEN ? AND ?", startTime, endTime).
		Group("users.id").Order("points_earned desc").Scan(&userPoints).Error

	if err != nil {
		return nil, err
	}
	return userPoints, nil
}

type UserPoints struct {
	GameCount    int    `json:"game_count"`
	PointsEarned int    `json:"points_earned"`
	Name         string `json:"name"`
	Email        string `json:"email"`
}
