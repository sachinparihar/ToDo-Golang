package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Todo string             `json:"todo" bson:"todo"`
	Done bool               `json:"done" bson:"done"`
}

func CreateTodo(todo string) error {
	_, err := db.Collection("todos").InsertOne(context.TODO(), bson.M{"todo": todo, "done": false})
	return err
}

func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	cursor, err := db.Collection("todos").Find(context.TODO(), bson.D{})
	if err != nil {
		return todos, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var todo Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodoByID(id primitive.ObjectID) (Todo, error) {
	var todo Todo
	filter := bson.M{"_id": id}
	err := db.Collection("todos").FindOne(context.TODO(), filter).Decode(&todo)
	return todo, err
}

func MarkTodoDone(id primitive.ObjectID, done bool) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"done": done}}
	_, err := db.Collection("todos").UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteTodoByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := db.Collection("todos").DeleteOne(context.TODO(), filter)
	return err
}
