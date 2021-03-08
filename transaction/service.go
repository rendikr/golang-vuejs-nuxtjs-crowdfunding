package transaction

type Service interface {
	GetTransactionsByCampaignID(input GetTransactionDetailInput) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionsByCampaignID(input GetTransactionDetailInput) ([]Transaction, error) {
	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
