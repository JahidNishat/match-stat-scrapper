package models

import (
	"gorm.io/gorm"
)

type PlayersData struct {
	gorm.Model
	PlayerName string `json:"playerName"`
	Source     string `json:"source"`
	Url        string `json:"url" gorm:"uniqueIndex"`
}

type ScrappedData struct {
	gorm.Model
	Hash string `json:"scrappedHash" gorm:"column:hash;size:255;index"`
	Data string `json:"data" gorm:"column:data"`
}

type Subscriber struct {
	gorm.Model
	ChatID   int64 `json:"chatID" gorm:"uniqueIndex"`
	Approved bool  `json:"approved"`
}
