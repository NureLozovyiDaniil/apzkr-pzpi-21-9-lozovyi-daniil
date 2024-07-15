package domain

type Service interface {
	CreateCargo(cargo Cargo) error
	GetCargoByID(id string) (Cargo, error)
	UpdateCargoStatus(id, status string) error
	AssignVolunteerToCargo(cargoID, volunteerID string) error
	AddDeliverySteps(cargoID string, steps []DeliveryStep) error
	CompleteDeliveryStep(cargoID, stepID, photo string, rfidScanned bool, lossDescription string) error
	GetCargoByRFID(rfid string) (Cargo, error)
}
