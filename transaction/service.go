package transaction

import (
	"errors"
	"golang-crowdfunding-backend/campaign"
)

type Service interface {
	GetTransactionsByCampaignID(input GetTransactionDetailInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetTransactionDetailInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Unauthorized access")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	// campaign, err := s.campaignRepository.FindByID(input.ID)
	// if err != nil {
	// 	return []Transaction{}, err
	// }

	// if campaign.UserID != input.User.ID {
	// 	return []Transaction{}, errors.New("Unauthorized access")
	// }

	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.UserID = input.User.ID
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.Code = ""

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil

}
