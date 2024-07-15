package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCargo(cargo Cargo) error {
	return s.repo.CreateCargo(cargo)
}

func (s *service) GetCargoByID(id string) (Cargo, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Cargo{}, err
	}
	return s.repo.GetCargoByID(objId)
}

func (s *service) UpdateCargoStatus(id, status string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.repo.UpdateCargoStatus(objId, status)
}

func (s *service) AssignVolunteerToCargo(cargoID, volunteerID string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoID)
	if err != nil {
		return err
	}
	volunteerObjId, err := primitive.ObjectIDFromHex(volunteerID)
	if err != nil {
		return err
	}
	return s.repo.AssignVolunteerToCargo(cargoObjId, volunteerObjId)
}

func (s *service) AddDeliverySteps(cargoID string, steps []DeliveryStep) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoID)
	if err != nil {
		return err
	}
	return s.repo.AddDeliverySteps(cargoObjId, steps)
}

func (s *service) CompleteDeliveryStep(cargoID, stepID, photo string, rfidScanned bool, lossDescription string) error {
	cargoObjId, err := primitive.ObjectIDFromHex(cargoID)
	if err != nil {
		return err
	}
	stepObjId, err := primitive.ObjectIDFromHex(stepID)
	if err != nil {
		return err
	}
	return s.repo.CompleteDeliveryStep(CompletedStep{
		StepID:  stepObjId,
		CargoID: cargoObjId,
		//Photo:       photo,
		RFIDScanned: rfidScanned,
		LossDesc:    lossDescription,
	})
}

func (s *service) GetCargoByRFID(rfid string) (Cargo, error) {
	return s.repo.GetCargoByRFID(rfid)
}
