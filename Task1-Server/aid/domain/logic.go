package domain

import (
	"context"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const AidDBName = "aid"

type service struct {
	client *mongo.Client
	db     *mongo.Database
	logger log.Logger
}

func NewSvc(dbURI string) Service {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURI))
	if err != nil {
		log.With(log.NewNopLogger()).Log("msg", "failed to connect to MongoDB", "err", err)
		panic(err)
	}

	return &service{
		client: client,
		db:     client.Database(AidDBName),
	}
}

func (s *service) IsOwner(userID, reqID string) bool {
	userObjID, _ := primitive.ObjectIDFromHex(userID)
	obj := s.db.Collection(ReqColl).FindOne(context.Background(), bson.M{"_id": reqID})

	var req HelpRequest
	if err := obj.Decode(&req); err != nil {
		return false
	}

	return req.UserID == userObjID
}

func (s *service) CreateHelpRequest(req HelpRequest) error {
	req.Status = AidStatusOpen
	req.CargoID = primitive.NilObjectID
	_, err := s.db.Collection(ReqColl).InsertOne(context.Background(), req)

	return err
}

func (s *service) DeleteRequest(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.db.Collection(ReqColl).DeleteOne(context.Background(), bson.M{"_id": objId})
	return err
}

func (s *service) GetRequests() ([]HelpRequest, error) {
	cursor, err := s.db.Collection(ReqColl).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var requests []HelpRequest
	if err := cursor.All(context.Background(), &requests); err != nil {
		return nil, err
	}

	return requests, nil
}

func (s *service) GetRequestByID(id string) (HelpRequest, error) {
	var req HelpRequest
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return HelpRequest{}, err
	}

	obj := s.db.Collection(ReqColl).FindOne(context.Background(), bson.M{"_id": objId})
	if err := obj.Decode(&req); err != nil {
		return HelpRequest{}, err
	}

	return req, nil
}

func (s *service) UpdateRequestStatus(id, status string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err = s.db.Collection(ReqColl).UpdateOne(context.Background(), filter, update)

	return err
}

func (s *service) AssignRequestToCargo(reqID, cargoID string) error {
	reqObjID, err := primitive.ObjectIDFromHex(reqID)
	if err != nil {
		return err
	}

	cargoObjID, err := primitive.ObjectIDFromHex(cargoID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": reqObjID}
	update := bson.M{"$set": bson.M{"cargo_id": cargoObjID}}
	_, err = s.db.Collection(ReqColl).UpdateOne(context.Background(), filter, update)

	return err
}

func (s *service) GetRequestsByCargo(cargoID string) ([]HelpRequest, error) {
	cargoObjID, err := primitive.ObjectIDFromHex(cargoID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"cargo_id": cargoObjID}
	cursor, err := s.db.Collection(ReqColl).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var requests []HelpRequest
	if err := cursor.All(context.Background(), &requests); err != nil {
		return nil, err
	}

	return requests, nil
}
