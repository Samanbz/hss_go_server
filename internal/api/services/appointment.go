package services

import (
	"context"
	"hss/internal/api/repositories"
	"hss/internal/models"
)

type AppointmentService struct {
	appointmentRepository repositories.AppointmentRepository
}

func NewAppointmentService(appointmentRepository repositories.AppointmentRepository) AppointmentService {
	return AppointmentService{appointmentRepository: appointmentRepository}
}

func (s *AppointmentService) InsertAppointment(ctx context.Context, appointment *models.Appointment) error {
	err := s.appointmentRepository.InsertAppointment(context.Background(), appointment)
	if err != nil {
		return err
	}
	return nil
}

func (s *AppointmentService) GetAppointmentByID(ctx context.Context, id int) (*models.Appointment, error) {
	appointment, err := s.appointmentRepository.GetAppointmentByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *AppointmentService) GetAppointmentsByCompanyID(ctx context.Context, company_id int) (*[]models.Appointment, error) {
	appointments, err := s.appointmentRepository.GetAppointmentsByCompanyID(context.Background(), company_id)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}
