package repos

import (
	"gorm.io/gorm"
	"log"
	"match_statistics_scrapper/models"
)

type SubscriberStore struct {
	DB *gorm.DB
}

func (s *SubscriberStore) Save(sd *models.Subscriber) error {
	res := s.DB.Save(sd)
	if res.Error != nil {
		log.Println("Error while creating entry in db", res.Error)
		return res.Error
	}
	return nil
}

func (s *SubscriberStore) GetSubscriberData(chatID int64) (*models.Subscriber, error) {
	var sd models.Subscriber
	res := s.DB.Model(&models.Subscriber{}).Where("chat_id = ?", chatID).First(&sd)
	if res.Error != nil {
		log.Println("Error while getting subscriber in db", res.Error)
		return nil, res.Error
	}
	return &sd, nil
}

func (s *SubscriberStore) GetAllSubscribers() ([]models.Subscriber, error) {
	var sd []models.Subscriber
	res := s.DB.Model(&models.Subscriber{}).Where("approved = ? ", true).Find(&sd)
	if res.Error != nil {
		log.Println("Error while getting players in db", res.Error)
		return nil, res.Error
	}
	return sd, nil
}
