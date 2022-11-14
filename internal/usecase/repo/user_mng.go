package repo

import (
	"auth_reg/internal/entity"
	"auth_reg/internal/usecase"
	"auth_reg/pkg/mongodb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepo struct {
	mngCollection *mongo.Collection
}

var _ usecase.UserRp = (*UserRepo)(nil)

func NewUserRepo(mng *mongodb.Mongo, collectionName string) *UserRepo {
	return &UserRepo{
		mngCollection: mng.Db.Collection(collectionName),
	}
}

func (u *UserRepo) StoreUser(ctx context.Context, user entity.User) error {
	_, err := u.mngCollection.InsertOne(ctx, bson.D{
		{"email", user.Email},
		{"password", user.Password},
	})
	if err != nil {
		log.Println("Cannot execute user inserting")
		return err
	}
	return nil
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	filter := bson.M{"email": email}

	var usr entity.User
	err := u.mngCollection.FindOne(ctx, filter).Decode(&usr)
	if err != nil {
		log.Println("err in getting user")
		if err == mongo.ErrNoDocuments {
			log.Println("Cannot find user in db")
			return entity.User{}, fmt.Errorf("cannot find user in db")
		}
	}
	return usr, nil
}

func (u *UserRepo) CheckExistenceByEmail(ctx context.Context, email string) (bool, error) {
	filter := bson.M{"email": email}

	count, err := u.mngCollection.CountDocuments(ctx, filter)
	if err != nil {
		log.Println("error in counting users")
		return false, fmt.Errorf("cannot execute query %w", err)
	}
	return count != 0, nil
}
