package order

import (
	"context"
	"time"

	domain "github.com/saharat-vithchataya/ordering/domain/order"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type orderRepositoryMongoDB struct {
	// mongodb driver
	db         *mongo.Database
	collection string
}

func NewOrderRepositoryMongoDB(db *mongo.Database, collection string) domain.OrderRepository {
	return orderRepositoryMongoDB{db: db, collection: collection}
}

func (repo orderRepositoryMongoDB) NextIdentity() string {
	return primitive.NewObjectID().Hex()
}

func (repo orderRepositoryMongoDB) FromID(orderID string) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var order domain.Order

	objectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
	}

	if err = repo.db.Collection(repo.collection).FindOne(ctx, filter).Decode(&order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (repo orderRepositoryMongoDB) Save(entity *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(entity.ID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objectID,
	}
	update := bson.M{
		"$set": bson.M{
			"customer_id": entity.CustomerID,
			"items":       entity.Items,
			"submitted":   entity.Submitted,
		},
	}

	_, err = repo.db.Collection(repo.collection).UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrOrderNotFound
		}
		return err
	}

	return nil
}
