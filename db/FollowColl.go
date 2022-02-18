package db

import (
	"context"
	"log"
	"wetime-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

type FollowColl struct {
	name string
}

var Follow = FollowColl{name: "follow"}

func (c *FollowColl) SaveFollow(follow models.Follow) (err error) {
	collection := DB.Collection(c.name)

	insertResult, err := collection.InsertOne(context.TODO(), follow)

	if err != nil {
		log.Printf("follow insert error: %s", err)
		return err
	}
	log.Printf("Inserted a single document: %s", insertResult.InsertedID)
	return nil
}

func (c *FollowColl) FindFollow(follow models.Follow) (res []*models.Follow, err error) {
	collection := DB.Collection(c.name)
	filter := bson.M{"userID": follow.UserID}

	results := []*models.Follow{}
	follows, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("follows find error: %s", err)
		return nil, err
	}
	err = follows.All(context.TODO(), &results)
	if err != nil {
		log.Printf("follows find error: %s", err)
		return nil, err
	}

	//log.Printf("Find documents: %s", result.UserID)
	return results, nil
}
