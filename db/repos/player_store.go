package repos

import (
	"gorm.io/gorm"
	"log"
	"match_statistics_scrapper/models"
)

type PlayersStore struct {
	DB *gorm.DB
}

func (s *PlayersStore) InsertData(sd *models.PlayersData) error {
	res := s.DB.Create(sd)
	if res.Error != nil {
		log.Println("Error while creating entry in db", res.Error)
		return res.Error
	}
	return nil
}

func (s *PlayersStore) GetPlayersData() ([]models.PlayersData, error) {
	var sd []models.PlayersData
	res := s.DB.Model(&models.PlayersData{}).Find(&sd)
	if res.Error != nil {
		log.Println("Error while getting players in db", res.Error)
		return nil, res.Error
	}
	return sd, nil
}
