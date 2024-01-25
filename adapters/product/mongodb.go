package product

import (
	"context"
	"time"

	domain "github.com/saharat-vithchataya/ordering/domain/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepositoryMongoDB struct {
	// mongodb driver
	db         *mongo.Database
	collection string
}

func NewProductRepositoryMongoDB(db *mongo.Database, collection string) domain.ProductRepository {
	return productRepositoryMongoDB{db: db, collection: collection}
}

func (repo productRepositoryMongoDB) NextIdentity() string {
	return primitive.NewObjectID().Hex()
}

func (repo productRepositoryMongoDB) FromID(productID string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var product domain.Product

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
	}

	if err = repo.db.Collection(repo.collection).FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo productRepositoryMongoDB) Save(entity *domain.Product) error {
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
			"name":        entity.Name,
			"description": entity.Description,
			"price":       entity.Price,
			"quantity":    entity.Quantity,
		},
	}

	if _, err = repo.db.Collection(repo.collection).UpdateOne(ctx, filter, update, options.Update().SetUpsert(true)); err != nil {
		return err
	}

	return nil
}
