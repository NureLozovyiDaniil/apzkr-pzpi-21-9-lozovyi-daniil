package domain

type Service interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (Identity, error)

	CreateUser(user User) error
	UserById(id string) (User, error)
	UpdateUser(user UserUpdate) error
	DeleteUser(id string) error
	GetAllUsers() ([]User, error)

	BecomeVolunteer(userID string) error

	CreateOrganization(userId string, org Organization) error
	JoinOrganization(userID, orgID string) error
	LeaveOrganization(userID, orgID string) error
	GetOrganizationById(id string) (Organization, error)
	GetAllOrganizations() ([]Organization, error)
}
