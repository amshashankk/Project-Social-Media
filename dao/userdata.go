package dao

import (
	"context"
	"time"

	"github.com/amshashankk/GoAuth/database"
	"github.com/amshashankk/GoAuth/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserData(User *models.User) (*mongo.InsertOneResult, error) {

	db := database.Get()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	return db.Collection("User").InsertOne(ctx, User)

}
