package domain

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct {
	collection *mongo.Collection
}

func NewMongoRepo(uri string) (Repository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	collection := client.Database("auth").Collection("users")
	return &repo{collection: collection}, nil
}

func (r *repo) CreateUser(u User) error {
	_, err := r.collection.InsertOne(context.Background(), u)
	return err
}

func (r *repo) UserByLogin(username, password string) (User, error) {
	u := new(User)
	filter := bson.M{"username": username, "password": password}
	err := r.collection.FindOne(context.Background(), filter).Decode(&u)

	return *u, err
}

func (r *repo) UserById(id primitive.ObjectID) (User, error) {
	u := new(User)
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(context.Background(), filter).Decode(&u)

	return *u, err
}

func (r *repo) UpdateUser(id primitive.ObjectID, u UserUpdate) error {
	filter := bson.M{"_id": id}
	update := bson.M{}

	if u.Username != nil {
		update["username"] = *u.Username
	}
	if u.Password != nil {
		update["password"] = *u.Password
	}
	if u.Role != nil {
		update["role"] = *u.Role
	}
	if u.Phone != nil {
		update["phone"] = *u.Phone
	}
	if u.Address != nil {
		update["address"] = *u.Address
	}
	if u.FullName != nil {
		update["full_name"] = *u.FullName
	}

	fmt.Println(update)
	_, err := r.collection.UpdateOne(context.Background(), filter, bson.M{"$set": update})
	//fmt.Println(res.ModifiedCount)

	return err
}

func (r *repo) DeleteUser(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)

	return err
}

func (r *repo) BecomeVolunteer(userID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"role": RoleVolunteer}}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (r *repo) CreateOrganization(org Organization) error {
	_, err := r.collection.InsertOne(context.Background(), org)
	return err
}

func (r *repo) JoinOrganization(userID, orgID primitive.ObjectID) error {
	filter := bson.M{"_id": orgID}
	update := bson.M{"$push": bson.M{"members": userID}}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (r *repo) LeaveOrganization(userID, orgID primitive.ObjectID) error {
	filter := bson.M{"_id": orgID}
	update := bson.M{"$pull": bson.M{"members": userID}}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)

	return err

}

func (r *repo) GetOrganizationById(id primitive.ObjectID) (Organization, error) {
	org := new(Organization)
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(context.Background(), filter).Decode(&org)

	return *org, err
}

func (r *repo) GetAllUsers() ([]User, error) {
	var users []User
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &users)
	return users, err
}

func (r *repo) GetAllOrganizations() ([]Organization, error) {
	var orgs []Organization
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &orgs)
	return orgs, err
}
