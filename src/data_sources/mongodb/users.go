package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/gretchelg/Go_BudgetApp/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// dbUser defines a user as specified in the DB
// the `bson` tag maps to the field names in the db.
type dbUser struct {
	Email       string `bson:"email"`
	AccessToken string `bson:"access_token"`
	FirstName   string `bson:"first_name"`
	LastName    string `bson:"last_name"`
	// TODO add more
}

// GetAllUsers returns all Users
func (d *DBClient) GetAllUsers(ctx context.Context) ([]models.User, error) {
	// create context used to enforce timeouts
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// get all users
	cursor, err := d.usersCollection.Find(ctxWithTimeout, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctxWithTimeout)

	// parse the db call response.
	// for each row, convert the data from dbModel to appModel and append to the list of results.
	var results []models.User
	for cursor.Next(ctxWithTimeout) {

		// parse the db row date into the dbModel
		var aDbUser dbUser
		err = cursor.Decode(&aDbUser)
		if err != nil {
			return nil, err
		}

		// convert dbModel to our appModel
		aUser := convertUser(aDbUser)

		// append to the list of results
		results = append(results, aUser)
	}

	// final check for any errors reported
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	// respond
	return results, nil
}

// GetUserByEmail returns one user specified by the given email
func (d *DBClient) GetUserByEmail(ctx0 context.Context, email string) (*models.User, error) {
	// create context used to enforce timeouts
	ctxWithTimeout, cancel := context.WithTimeout(ctx0, timeout)
	defer cancel()

	// create filter that matches the given ID
	filter := bson.D{
		{
			Key:   "email",
			Value: email,
		},
	}

	// do find
	var aUser dbUser
	err := d.usersCollection.FindOne(ctxWithTimeout, filter).Decode(&aUser)
	if err == mongo.ErrNoDocuments {
		// if no matching docs found, return sentinel error "models.ErrorNotFound" that callers can inspect in order to
		// handle in a custom way, such as returning 404-NotFound rather than a generic 500-InternalServerError
		return nil, fmt.Errorf("DB.GetUserByEmail: %w", models.ErrorNotFound)
	}

	if err != nil {
		return nil, err
	}

	// respond
	result := convertUser(aUser)
	return &result, nil

}

// convertTransaction converts from the internal db model to the application-wide data model.
func convertUser(dbModel dbUser) models.User {
	// response
	return models.User{
		Email:       dbModel.Email,
		AccessToken: dbModel.AccessToken,
		FirstName:   dbModel.FirstName,
		LastName:    dbModel.LastName,
	}
}
