package main

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	customerMongoDBRepository "github.com/saharat-vithchataya/ordering/adapters/customer"
	orderMongoDBRepository "github.com/saharat-vithchataya/ordering/adapters/order"
	productMongoDBRepository "github.com/saharat-vithchataya/ordering/adapters/product"
	"github.com/saharat-vithchataya/ordering/handlers"
	"github.com/saharat-vithchataya/ordering/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db := initDB()
	app := fiber.New()

	orderRepo := orderMongoDBRepository.NewOrderRepositoryMongoDB(db, "orders")
	orderSrv := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderSrv)

	customerRepo := customerMongoDBRepository.NewCustomerRepositoryMongoDB(db, "customers")
	customerSrv := services.NewCustomerService(customerRepo)
	customerHandler := handlers.NewCustomerHandler(customerSrv)

	productRepo := productMongoDBRepository.NewProductRepositoryMongoDB(db, "products")
	productSrv := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productSrv)

	app.Get("/customer/:customer_id", customerHandler.GetCustomer)
	app.Post("/customer", customerHandler.CreateNewCustomer)

	app.Get("/order/:order_id", orderHandler.GetOrder)
	app.Post("/order/:customer_id/create", orderHandler.CreateNewOrder)
	app.Put("/order/:order_id", orderHandler.UpdateOrderItem)
	app.Put("/order/:order_id/submit", orderHandler.SubmitOrder)

	app.Get("/product/:product_id", productHandler.GetProduct)
	app.Post("/product", productHandler.CreateNewProduct)

	app.Listen(":8000")
}

func initDB() *mongo.Database {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb://root:example@localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	db := client.Database("swiss")

	return db
}
