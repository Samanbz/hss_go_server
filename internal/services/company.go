package services

import (
	"context"
	"hss/internal/models"
	"hss/internal/repositories"
)

type CompanyService struct {
	companyRepository *repositories.CompanyRepository
}

func NewCompanyService(companyRepository *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{companyRepository: companyRepository}
}

func (s *CompanyService) InsertCompany(ctx context.Context, company *models.Company) error {
	err := s.companyRepository.InsertCompany(context.Background(), company)
	if err != nil {
		return err
	}
	return nil
}

func (s *CompanyService) GetAllCompanies(ctx context.Context) ([]models.Company, error) {
	companies, err := s.companyRepository.GetAllCompanies(context.Background())
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (s *CompanyService) GetCompanyByID(ctx context.Context, id int) (*models.Company, error) {
	company, err := s.companyRepository.GetCompanyByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return company, nil
}
