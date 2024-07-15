package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Repository interface {
	CreateCargo(cargo Cargo) error
	GetCargoByID(id primitive.ObjectID) (Cargo, error)
	UpdateCargoStatus(id primitive.ObjectID, status string) error
	AssignVolunteerToCargo(cargoID, volunteerID primitive.ObjectID) error
	AddDeliverySteps(cargoID primitive.ObjectID, steps []DeliveryStep) error
	CompleteDeliveryStep(cStep CompletedStep) error
	GetCargoByRFID(rfid string) (Cargo, error)
}
