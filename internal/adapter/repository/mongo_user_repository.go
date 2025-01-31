package repository

import (
	"context"
	"library-app/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *mongoUserRepository {
	return &mongoUserRepository{collection: collection}
}

// Implement UserRepository methods
func (r *mongoUserRepository) Create(user *domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *mongoUserRepository) GetByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mongoUserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, &user)
	}
	return users, nil
}

func (r *mongoUserRepository) Update(id string, user *domain.User) error {
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": id}, user)
	return err
}

func (r *mongoUserRepository) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *mongoUserRepository) CreateUser(user *domain.User) error {
	return r.Create(user)
}

func (r *mongoUserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
