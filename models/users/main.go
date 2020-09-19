package users

import (
	"api/models"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty" json:"name,omitempty"`
	Email string             `bson:"email,omitempty" json:"email,omitempty"`
}

func New(user *User) (error, string) {
	_, err := json.Marshal(user)

	if err != nil {
		return err, "Elemento enviado nao corresponde a um json"
	}

	mongo := models.GetConnection()
	collection := mongo.Database("getty").Collection("users")

	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return err, "Houve um erro na insercao"
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return err, "Objeto criado, falha no retorno do Id"
	}

	user.Id = oid

	return nil, ""
}

func List(users *[]User) (error, string) {
	mongo := models.GetConnection()
	collection := mongo.Database("getty").Collection("users")

	result, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err, "Erro ao realizar solicitacao para o banco"
	}

	defer result.Close(context.TODO())

	for result.Next(context.TODO()) {

		var user User
		err := result.Decode(&user)
		if err != nil {
			return err, "Erro ao realizar o decode do usuario"
		}

		// add item our array
		*users = append(*users, user)
	}

	return nil, ""
}

func Delete(_id primitive.ObjectID) (error, string) {
	mongo := models.GetConnection()
	collection := mongo.Database("getty").Collection("users")

	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": _id})

	if err != nil {
		return err, "Falha ao deletar usuario"
	}

	if res.DeletedCount < 1 {
		return errors.New("Nenhum item foi deletado"), "Nenhum item foi deletado"
	}

	return nil, ""
}

func Update(_id primitive.ObjectID, user User) (error, string) {
	user.Id = _id
	_, err := json.Marshal(user)

	if err != nil {
		return err, "Elemento enviado nao corresponde a um json"
	}

	mongo := models.GetConnection()
	collection := mongo.Database("getty").Collection("users")

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": user.Id}, bson.D{{Key: "$set", Value: user}})

	if err != nil {
		return err, "Ocorreu um erro ao atualizar"
	}

	return nil, ""
}

func GetOne(_id primitive.ObjectID, user *User) (error, string) {
	user.Id = _id
	_, err := json.Marshal(user)

	if err != nil {
		return err, "Elemento enviado nao corresponde a um json"
	}

	mongo := models.GetConnection()
	collection := mongo.Database("getty").Collection("users")

	if err := collection.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&user); err != nil {
		return nil, "Nada encontrado"
	}

	return nil, ""
}
