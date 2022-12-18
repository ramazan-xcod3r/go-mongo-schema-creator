package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var test *bool
func main() {
	// Parse the command-line flags
	test = flag.Bool("test", false, "Test mode")
	flag.Parse()
	fmt.Println(*test)
	// Set up the connection to the MongoDB server
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, errs := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if errs != nil {
		return
	}
	defer cancel()
	
	args := flag.Args()
	fmt.Println(args)
	if len(args) > 0 {
		filename := args[0]
		err := AddDataFromJSONFile(client, filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func AddDataFromJSONFile(client *mongo.Client, filename string) error {
	// Read the JSON file
	data, err := ReadJSONFile(filename)
	if err != nil {
		return err
	}
	// Create a BSON document based on the data
	doc := bson.M{}
	for k, v := range data {
		if v == nil {
			doc[k] = "null"
		} else {
			switch reflect.TypeOf(v).Kind() {
			case reflect.String:
				doc[k] = "string"
			case reflect.Float64:
				doc[k] = "number"
			case reflect.Map:
				doc[k] = "object"
			case reflect.Slice:
				doc[k] = "array"
			case reflect.Bool:
				doc[k] = "boolean"
			}
		}
	}
	//doc = append(doc, bson.E{Key: k, Value: fmt.Sprintf(":%T",v)})

	jsonStr, _ := bson.MarshalExtJSONIndent(doc, true, true, "", "	")
	// Output the BSON document to the console
	fmt.Println(string(jsonStr))

	if *test {
		// Insert the data into the database
		err = InsertData(client, data)
		if err != nil {
			return err
		}
		fmt.Println("Successfully added user to the database.")
	}
	return nil
}
func InsertData(client *mongo.Client, data interface{}) error {
	// Create the "test" database and "data" collection if they don't already exist
	db := client.Database("test")
	usersColl := db.Collection("data")

	// Insert the data into the collection
	_, err := usersColl.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

// ReadJSONFile reads a JSON file and returns the data as an interface
func ReadJSONFile(filename string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
