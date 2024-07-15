package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	TrackingColl = "tracking"
	FeedbackColl = "feedback"
)

type Service interface {
	AddTrackingInfo(info TrackingInfo) error
	GetCargoTrackingHistory(cargoID string) ([]TrackingInfo, error)
	SubmitFeedback(feedback Feedback) error
	GetRequestFeedback(requestID string) ([]Feedback, error)
}

type svc struct {
	db *mongo.Database
}

func NewService(uri string) (Service, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	database := client.Database("aid")

	return &svc{db: database}, nil
}

func (s *svc) AddTrackingInfo(info TrackingInfo) error {
	_, err := s.db.Collection(TrackingColl).InsertOne(context.Background(), info)
	return err
}

func (s *svc) GetCargoTrackingHistory(cargoID string) ([]TrackingInfo, error) {
	filter := bson.M{"cargo_id": cargoID}
	cursor, err := s.db.Collection(TrackingColl).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var history []TrackingInfo
	if err := cursor.All(context.Background(), &history); err != nil {
		return nil, err
	}

	return history, nil
}

func (s *svc) SubmitFeedback(feedback Feedback) error {
	_, err := s.db.Collection(FeedbackColl).InsertOne(context.Background(), feedback)
	return err
}

func (s *svc) GetRequestFeedback(requestID string) ([]Feedback, error) {
	filter := bson.M{"request_id": requestID}
	cursor, err := s.db.Collection(FeedbackColl).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var feedbacks []Feedback
	if err := cursor.All(context.Background(), &feedbacks); err != nil {
		return nil, err
	}

	return feedbacks, nil
}
