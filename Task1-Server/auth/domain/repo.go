package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Repository interface {
	CreateUser(u User) error
	UserByLogin(username, password string) (User, error)
	UserById(id primitive.ObjectID) (User, error)
	UpdateUser(id primitive.ObjectID, u UserUpdate) error
	DeleteUser(id primitive.ObjectID) error
	BecomeVolunteer(userID primitive.ObjectID) error

	CreateOrganization(org Organization) error
	JoinOrganization(userID, orgID primitive.ObjectID) error
	LeaveOrganization(userID, orgID primitive.ObjectID) error
	GetOrganizationById(id primitive.ObjectID) (Organization, error)
	GetAllUsers() ([]User, error)
	GetAllOrganizations() ([]Organization, error)
}
