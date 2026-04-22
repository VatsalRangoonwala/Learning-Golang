package repository

import (
	"context"
	"project/db"
	"project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertOrder(ctx context.Context, order models.Order) error {
	collection := db.Client.Database("shop").Collection("orders")
	_, err := collection.InsertOne(ctx, order)
	return err
}

func GetAllOrders() ([]models.Order, error) {
	collection := db.Client.Database("shop").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var orders []models.Order

	for cursor.Next(ctx) {
		var order models.Order
		err := cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func GetOrderByID(id string) (models.Order, error) {
	collection := db.Client.Database("shop").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var order models.Order

	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}