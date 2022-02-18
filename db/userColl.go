package db

import (
	"context"
	"log"
	"wetime-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

type UserColl struct {
	name string
}

var User = UserColl{name: "user"}

func (c *UserColl) SaveUser(user models.User) (err error) {
	collection := DB.Collection(c.name)

	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Printf("user insert error: %s", err)
		return err
	}
	log.Printf("Inserted a single document: %s", insertResult.InsertedID)
	return nil
}

func (c *UserColl) FindUser(user models.User) (res *models.User, err error) {
	collection := DB.Collection(c.name)
	filter := bson.M{"userID": user.UserID}

	var result *models.User
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Printf("user find error: %s", err)
		return nil, err
	}
	log.Printf("Find a single document: %s", result.UserID)
	return result, nil
}

func (c *UserColl) FindUsers(userIDs []string) (res []*models.User, err error) {
	collection := DB.Collection(c.name)
	filter := bson.M{"userID": bson.M{"$in": userIDs}}

	var results []*models.User
	users, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("users find error: %s", err)
		return nil, err
	}
	err = users.All(context.TODO(), &results)
	if err != nil {
		log.Printf("users find error: %s", err)
		return nil, err
	}
	//log.Printf("Find a single document: %s", result.UserID)
	return results, nil
}
