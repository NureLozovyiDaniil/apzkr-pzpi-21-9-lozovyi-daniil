package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type repo struct {
	coll *mongo.Collection
}

func NewRepo(uri string) (Repository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	collection := client.Database("cargo").Collection("cargos")
	return &repo{coll: collection}, nil
}

func (r *repo) CreateCargo(cargo Cargo) error {
	_, err := r.coll.InsertOne(context.Background(), cargo)
	return err
}

func (r *repo) GetCargoByID(id primitive.ObjectID) (Cargo, error) {
	filter := bson.M{"_id": id}
	cargo := new(Cargo)
	err := r.coll.FindOne(context.Background(), filter).Decode(&cargo)

	return *cargo, err
}

func (r *repo) UpdateCargoStatus(id primitive.ObjectID, status string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err := r.coll.UpdateOne(context.Background(), filter, update)

	return err
}

func (r *repo) AssignVolunteerToCargo(cargoID, volunteerID primitive.ObjectID) error {
	filter := bson.M{"_id": cargoID}
	update := bson.M{"$set": bson.M{"volunteer": volunteerID}}
	_, err := r.coll.UpdateOne(context.Background(), filter, update)

	return err
}

func (r *repo) AddDeliverySteps(cargoID primitive.ObjectID, steps []DeliveryStep) error {
	filter := bson.M{"_id": cargoID}
	update := bson.M{"$set": bson.M{"steps": steps}}
	_, err := r.coll.UpdateOne(context.Background(), filter, update)

	return err
}

// TODO: Test this method

func (r *repo) CompleteDeliveryStep(cStep CompletedStep) error {
	filter := bson.M{"_id": cStep.CargoID, "steps._id": cStep.StepID}
	update := bson.M{
		"$set": bson.M{
			"steps.$.completion_status": "completed",
			"steps.$.completed_at":      primitive.NewDateTimeFromTime(time.Now()),
			//"steps.$.photo":             cStep.Photo,
			"steps.$.rfid_scanned": cStep.RFIDScanned,
			"steps.$.loss_desc":    cStep.LossDesc,
		}}
	_, err := r.coll.UpdateOne(context.Background(), filter, update)

	return err
}

func (r *repo) GetCargoByRFID(rfid string) (Cargo, error) {
	filter := bson.M{"rfid": rfid}
	cargo := new(Cargo)
	err := r.coll.FindOne(context.Background(), filter).Decode(&cargo)

	return *cargo, err
}
