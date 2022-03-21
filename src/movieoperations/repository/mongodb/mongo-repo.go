package mongo

import (
	"context"
	"fmt"
	"moviesdemo/config"
	"moviesdemo/entity"
	"moviesdemo/repository"
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
func NewMongoRepository() (repository.MovieRepository, error) {
	repo := &mongoRepository{}
	client, err := newMongoClient()
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (repo *mongoRepository) Find(id string) (*entity.Movie, error) {
	movieId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	collection := repo.client.Database(database).Collection("movies")
	filter := bson.M{"_id": movieId}
	movie := &entity.Movie{}
	err = collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
func (repo *mongoRepository) GetAll() ([]entity.Movie, error) {

	collection := repo.client.Database(database).Collection("movies")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		defer cursor.Close(context.TODO())
		return nil, err
	}
	var movies []entity.Movie
	if err = cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}
	return movies, nil

}
func (repo *mongoRepository) Save(movie *entity.Movie) (*entity.Movie, error) {
	collection := repo.client.Database(database).Collection("movies")
	movie.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return nil, err
	}
	return movie, err
}
