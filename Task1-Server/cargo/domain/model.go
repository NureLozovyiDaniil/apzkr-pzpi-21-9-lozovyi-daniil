package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cargo struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Description    string               `bson:"description"`
	Status         string               `bson:"status"`
	OrganizationID primitive.ObjectID   `bson:"organization_id,omitempty"`
	VolunteerID    primitive.ObjectID   `bson:"volunteer_id,omitempty"`
	RFID           string               `bson:"rfid"`
	Steps          []DeliveryStep       `bson:"steps"`
	Requests       []primitive.ObjectID `bson:"requests"`
}

type DeliveryStep struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	City             string             `bson:"city"`
	Comment          string             `bson:"comment,omitempty"`
	RequirePhoto     bool               `bson:"require_photo"`
	RequireRFIDScan  bool               `bson:"require_rfid_scan"`
	LossDescription  string             `bson:"loss_description,omitempty"`
	CompletionStatus string             `bson:"completion_status"`
	CompletedAt      primitive.DateTime `bson:"completed_at,omitempty"`
}

type CompletedStep struct {
	StepID  primitive.ObjectID `bson:"step_id,omitempty"`
	CargoID primitive.ObjectID `bson:"cargo_id"`
	//Photo       string             `bson:"photo"`
	RFIDScanned bool   `bson:"rfid_scanned"`
	LossDesc    string `bson:"loss_description"`
}
