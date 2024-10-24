package services

import (
	"context"
	"hss/internal/models"
	"hss/internal/repositories"
)

type CompanyService struct {
	CompanyRepository repositories.CompanyRepository
}

func (s *CompanyService) InsertCompany(ctx context.Context, company *models.Company) error {
	err := s.CompanyRepository.InsertCompany(context.Background(), company)
	if err != nil {
		return err
	}
	return nil
}
