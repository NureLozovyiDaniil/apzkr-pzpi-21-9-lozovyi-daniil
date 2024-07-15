package domain

type Service interface {
	IsOwner(userID, reqID string) bool

	CreateHelpRequest(request HelpRequest) error
	DeleteRequest(id string) error
	GetRequests() ([]HelpRequest, error)
	GetRequestByID(id string) (HelpRequest, error)
	UpdateRequestStatus(id, status string) error
	AssignRequestToCargo(requestID, cargoID string) error
	GetRequestsByCargo(cargoID string) ([]HelpRequest, error)
}
