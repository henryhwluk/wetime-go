package db

import (
	"context"
	"log"
	"wetime-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostColl struct {
	name string
}

var Post = PostColl{name: "post"}

func (c *PostColl) SavePost(post models.Post) (err error) {
	collection := DB.Collection(c.name)

	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Printf("post insert error: %s", err)
		return err
	}
	log.Printf("Inserted a single document: %s", insertResult.InsertedID)
	return nil
}

func (c *PostColl) FindPost(userIDs []string) (res []*models.Post, err error) {
	collection := DB.Collection(c.name)
	filter := bson.M{"userID": bson.M{"$in": userIDs}}
	var results []*models.Post

	findOptions := options.Find()
	// Sort by `postTime` field descending
	findOptions.SetSort(bson.M{"postTime": -1})

	posts, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Printf("post find error: %s", err)
		return nil, err
	}
	err = posts.All(context.TODO(), &results)
	if err != nil {
		log.Printf("post find error: %s", err)
		return nil, err
	}
	return results, nil
}
