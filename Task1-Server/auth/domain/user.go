package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	RoleUser = iota + 1
	RoleVolunteer
	RoleOrganization
	RoleAdmin
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"password"`
	DateOfBirth primitive.DateTime `bson:"date_of_birth" json:"date_of_birth"`
	FullName    string             `bson:"full_name" json:"full_name"`
	PhotoURL    string             `bson:"photo_url" json:"photo_url"`
	Role        int                `bson:"role" json:"role"`
	Phone       string             `bson:"phone" json:"phone"`
	Address     string             `bson:"address" json:"address"`
}

type UserUpdate struct {
	Id       string  `json:"id"`
	FullName *string `json:"full_name"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Role     *int    `json:"role"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}

type Organization struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Name        string               `bson:"name" json:"name"`
	Description string               `bson:"description" json:"description"`
	UserID      primitive.ObjectID   `bson:"user_id"`
	Members     []primitive.ObjectID `bson:"members"`
}

type Identity struct {
	Id   string `json:"user_id"`
	Role int    `json:"role"`
}
