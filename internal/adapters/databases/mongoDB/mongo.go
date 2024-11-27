package mongoDB

import (
	"book-app/internal/domain/entities"
	"book-app/internal/domain/repositories"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	coll *mongo.Collection
	ctx  context.Context
}

func NewMongoDB(coll *mongo.Collection, ctx context.Context) repositories.BookRepository {
	return &MongoDB{coll: coll, ctx: ctx}
}

func (md *MongoDB) CreateBook(book *entities.Book) (string, error) {
	res, err := md.coll.InsertOne(md.ctx, book)
	if err != nil {
		return "", fmt.Errorf("failed to insert book: %v", err)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (md *MongoDB) GetBookByID(bookID string) (*entities.Book, error) {
	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %v", err)
	}

	filter := bson.D{
		{"_id", objectID},
	}

	var book entities.Book

	err = md.coll.FindOne(md.ctx, filter).Decode(&book)
	if err != nil {
		return nil, fmt.Errorf("failed to find book: %v", err)
	}
	return &book, nil
}

func (md *MongoDB) GetAllBooks() ([]*entities.Book, error) {
	cursor, err := md.coll.Find(md.ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrive all books: %v", err)
	}
	var books []*entities.Book
	err = cursor.All(md.ctx, &books)
	if err != nil {
		return nil, fmt.Errorf("failed to return all books: %v", err)
	}
	return books, nil
}

func (md *MongoDB) UpdateBook(bookID string, book *entities.Book) error {
	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return fmt.Errorf("failed to get DB book id from provided id: %v", err)
	}
	updateData, err := bson.Marshal(book)
	if err != nil {
		return fmt.Errorf("failed to marshal book: %v", err)
	}
	var updateMap bson.M
	err = bson.Unmarshal(updateData, &updateMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal updateMap: %v", err)
	}
	update := bson.D{
		{"$set", updateMap},
	}
	res, err := md.coll.UpdateByID(md.ctx, id, &update)
	if err != nil {
		return fmt.Errorf("failed to update book: %v", err)
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("failed to find book: %v", bookID)
	}
	return nil

}

func (md *MongoDB) DeleteBook(bookID string) error {
	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return fmt.Errorf("failed to get DB book id from provided id: %v", err)
	}

	filter := bson.D{
		{"_id", id},
	}
	res, err := md.coll.DeleteOne(md.ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete book: %v", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("failed to find book: %v", bookID)
	}

	return nil
}
