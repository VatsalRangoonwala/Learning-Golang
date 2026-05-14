package repository

import (
	"context"
	"errors"
	"project/db"
	"project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DBName          = "shop"
	OrderCollection = "orders"
)

func InsertOrder(ctx context.Context, order models.Order) error {
	collection := db.Client.Database("shop").Collection("orders")
	_, err := collection.InsertOne(ctx, order)
	return err
}

func GetAllOrders(ctx context.Context) ([]models.Order, error) {
	collection := db.Client.Database("shop").Collection("orders")

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
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(ctx context.Context, id string) (models.Order, error) {
	collection := db.Client.Database(DBName).Collection(OrderCollection)

	var order models.Order

	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&order)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return order, errors.New("order not found")
		}
		return order, err
	}

	return order, nil
}
