package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func init(){
	// 1、mongodb配置连接
	clientOptions := options.Client().ApplyURI("mongodb://120.78.159.42:27017")

	// 2、连接
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}

	// 3、ping
	err = client.Ping(context.TODO(), nil)
	if err!=nil{
		log.Fatalln(err)
	}

	Client = client
	fmt.Println("Connect to MongoDB!!!")
}


func main()  {
	collection := Client.Database("test").Collection("student")

	i := int64(1)

	ctx := context.TODO()
	cur, _ := collection.Find(ctx, bson.M{
		"classId":"62778e739d5500002d003223",
	}, &options.FindOptions{
		Limit: &i,
	})
	
	results := make([]map[string]interface{},0)
	
	for cur.Next(ctx){
		res := map[string]interface{}{}
		err := cur.Decode(res)
		if err!=nil{
			log.Fatalln(err)
		}
		results = append(results, res)
	}
	
	fmt.Println(results)
}
