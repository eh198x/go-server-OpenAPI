package swagger

import (
	"context"
	"fmt"

	//swagger "github.com/ehadjikyriacou/go-server-OpenAPI/g"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_URI = "mongodb://mongoadmin:enigma@localhost:27017"
const DBNAME = "oai"
const COLLECTIONAME = "users"

func InsertDataSimple(data []interface{}) error {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}

	err = insertData(client, DBNAME, COLLECTIONAME, data)
	if err != nil {
		panic(err)
	}
	return nil
}

func insertData(client *mongo.Client, dbName string, collectionName string, data []interface{}) error {
	ctx := context.TODO()

	// Get a handle to the target collection
	collection := client.Database(dbName).Collection(collectionName)

	// Insert the data
	_, err := collection.InsertMany(ctx, data)
	if err != nil {
		return err
	}

	fmt.Println("Data inserted successfully!")
	return nil
}

func ViewInsertedDataSimple() error {
	var err error

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}

	err = viewInsertedData(client, DBNAME, COLLECTIONAME)
	if err != nil {
		panic(err)
	}
	return nil
}
func viewInsertedData(client *mongo.Client, dbName string, collectionName string) error {
	ctx := context.TODO()

	// Get a handle to the target collection
	collection := client.Database(dbName).Collection(collectionName)

	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var err error
	var users []User
	if err = cur.All(ctx, &users); err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println("Full name:", user.FullName, ", Roles:", user.Roles, ", Email:", user.Email)
	}
	return nil
}

/*
func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}

	// Define the data to insert
	data := []interface{}{
		bson.D{{Key: "Email", Value: "e@e.cy"}, {Key: "Password", Value: "hokuto1h"}, {Key: "FullName", Value: "KENS A"}},
		bson.D{{Key: "Email", Value: "a@e.cy"}, {Key: "Password", Value: "hokuto2h"}, {Key: "FullName", Value: "KENS B"}},
		bson.D{{Key: "Email", Value: "b@e.cy"}, {Key: "Password", Value: "hokutoh3"}, {Key: "FullName", Value: "KENS C"}},
	}

	fmt.Println(data)
	fmt.Printf("\n\n")
	// Insert the data
	//err = insertData(client, DBNAME, COLLECTIONAME, data)
	if err != nil {
		panic(err)l {
		anic(err)(data)
	mt.Printf("\n\n")
// Insert the data
	//err = insertData(client, DBNAME, COLLECTIONAME, daa)
	if err != nil
		panic(errl {
	anic(err)"\n\n")
/ Insert the data
//err = insertData(client, DBNAME, COLLECTIONAME, daa)
	if err != nil
		panic(errl {
	anic(err)l {
anic(err)"\n\n")
 Insert the data
//err = insertData(client, DBNAME, COLLECTIONAME, da)
	if err != nil
		panic(err {
	nic(err)


	err = viewInseredData(client, BNAME,OLLECTIONAME)
	if err != nl
		anic(err)
}

/Close the MongoDB connection
	r= client.Diconnect(context.TODO())
	i err != nl {
		anic(err)
}

}
*/
