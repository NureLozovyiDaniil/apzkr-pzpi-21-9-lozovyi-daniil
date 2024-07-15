package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrackingInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CargoID   primitive.ObjectID `bson:"cargo_id"`
	StepID    primitive.ObjectID `bson:"step_id"`
	Location  string             `bson:"location"`
	Timestamp primitive.DateTime `bson:"timestamp"`
	Photo     string             `bson:"photo,omitempty"`
	RFIDScan  bool               `bson:"rfid_scan"`
}

type Feedback struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	RequestID primitive.ObjectID `bson:"request_id"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Rating    int                `bson:"rating"`
	Comment   string             `bson:"comment"`
}

type SubmitFeedback struct {
	ID        string `json:"id" validate:"required"`
	RequestID string `json:"request_id" validate:"required"`
	UserID    string `json:"user_id" validate:"required"`
	Rating    int    `json:"rating" validate:"required"`
	Comment   string `json:"comment" validate:"required"`
}
