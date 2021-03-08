package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
	GetbyUserID(userID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("campaign_id = ?", campaignID).Preload("User").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetbyUserID(userID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("user_id = ?", userID).Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
