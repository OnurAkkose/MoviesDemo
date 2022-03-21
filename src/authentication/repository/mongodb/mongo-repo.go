package mongo

import (
	"authentication/config"
	"authentication/entity"
	"authentication/repository"
	"context"
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	client *mongo.Client
}

var database string

func newMongoClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	var configMongoDB config.MongoDBConfig
	if err := cleanenv.ReadEnv(&configMongoDB); err != nil {
		return nil, err
	}
	mongoURL := fmt.Sprintf("mongodb://%s:%d", configMongoDB.Host, configMongoDB.Port)
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		AuthMechanism: configMongoDB.AuthMechanism,
		Username:      configMongoDB.User,
		Password:      configMongoDB.Password,
	})
	database = configMongoDB.Database
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return client, err
}
func NewMongoRepository() (repository.UserRepository, error) {
	repo := &mongoRepository{}
	client, err := newMongoClient()
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (repo *mongoRepository) Find(username string) (*entity.User, error) {
	collection := repo.client.Database(database).Collection("users")
	filter := bson.M{"username": username}
	user := &entity.User{}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *mongoRepository) GetAll() ([]entity.User, error) {

	collection := repo.client.Database(database).Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		defer cursor.Close(context.TODO())
		return nil, err
	}
	var users []entity.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil

}
func (repo *mongoRepository) Save(user *entity.User) (*entity.User, error) {
	collection := repo.client.Database(database).Collection("users")
	user.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return user, err
}
