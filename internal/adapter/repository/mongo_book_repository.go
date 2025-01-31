package repository

import (
	"context"
	"library-app/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoBookRepository struct {
	collection *mongo.Collection
}

func NewMongoBookRepository(collection *mongo.Collection) *mongoBookRepository {
	return &mongoBookRepository{collection: collection}
}

func (r *mongoBookRepository) Create(book *domain.Book) error {
	_, err := r.collection.InsertOne(context.Background(), book)
	return err
}

func (r *mongoBookRepository) GetByID(id string) (*domain.Book, error) {
	var book domain.Book
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *mongoBookRepository) GetAll() ([]*domain.Book, error) {
	var books []*domain.Book
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var book domain.Book
		cursor.Decode(&book)
		books = append(books, &book)
	}
	return books, nil
}

func (r *mongoBookRepository) Update(id string, book *domain.Book) error {
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": id}, book)
	return err
}

func (r *mongoBookRepository) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
