package services

import (
	"context"
	"hss/internal/api/repositories"
	"hss/internal/models"
)

type AddressService struct {
	AddressRepository repositories.AddressRepository
}

func NewAddressService(addressRepository repositories.AddressRepository) AddressService {
	return AddressService{AddressRepository: addressRepository}
}

func (s *AddressService) InsertAddress(ctx context.Context, address *models.Address) error {
	err := s.AddressRepository.InsertAddress(context.Background(), address)
	if err != nil {
		return err
	}
	return nil
}
