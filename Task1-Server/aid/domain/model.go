package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	RoleUser = iota + 1
	RoleVolunteer
	RoleOrganization
	RoleAdmin

	AidStatusOpen    = "OPEN"
	AidStatusPending = "IN PROGRESS"
	AidStatusClosed  = "CLOSED"

	ReqColl = "aid_requests"
)

type HelpRequest struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	CargoID     primitive.ObjectID `bson:"cargo_id,omitempty" json:"cargo_id"`
}
