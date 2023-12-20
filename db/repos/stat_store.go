package repos

import (
	"gorm.io/gorm"
	"log"
	"match_statistics_scrapper/models"
)

type StatsStore struct {
	DB *gorm.DB
}

func (s *StatsStore) InsertData(sd *models.ScrappedData) error {
	res := s.DB.Create(sd)
	if res.Error != nil {
		log.Println("Error while creating entry in db", res.Error)
		return res.Error
	}
	return nil
}

func (s *StatsStore) FindHash(hash string) (*models.ScrappedData, error) {
	sd := &models.ScrappedData{}
	res := s.DB.Where("hash = ?", hash).First(sd)
	if res.Error != nil {
		log.Println("Error while getting user in db", res.Error)
		return nil, res.Error
	}
	return sd, nil
}
