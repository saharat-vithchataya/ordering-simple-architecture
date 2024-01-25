package customer

import (
	"context"
	"time"

	domain "github.com/saharat-vithchataya/ordering/domain/customer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customerRepositoryMongoDB struct {
	// mongodb driver
	db         *mongo.Database
	collection string
}

func NewCustomerRepositoryMongoDB(db *mongo.Database, collection string) domain.CustomerRepository {
	return customerRepositoryMongoDB{db: db, collection: collection}
}

func (repo customerRepositoryMongoDB) NextIdentity() string {
	return primitive.NewObjectID().Hex()
}

func (repo customerRepositoryMongoDB) FromID(customerID string) (*domain.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var customer domain.Customer

	objectID, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
	}

	if err = repo.db.Collection(repo.collection).FindOne(ctx, filter).Decode(&customer); err != nil {
		return nil, err
	}

	return &customer, nil
}

func (repo customerRepositoryMongoDB) Save(entity *domain.Customer) error {
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
			"name":    entity.Name,
			"phone":   entity.Phone,
			"address": entity.Address,
		},
	}

	if _, err = repo.db.Collection(repo.collection).UpdateOne(ctx, filter, update, options.Update().SetUpsert(true)); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrCustomerNotFound
		}
		return err
	}

	return nil
}
