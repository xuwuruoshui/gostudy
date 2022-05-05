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

// 断开连接
func close(){
	err := Client.Connect(context.TODO())
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println("Connection is MongoDB closed.")
}

// 连接池模式
//func ConnectToDB(uri,name string,timeout time.Duration,num uint64)(*mongo.Database,error){
//	// 设置超时
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	defer cancel()
//	o := options.Client().ApplyURI(uri)
//	o.SetMaxPoolSize(num)
//
//	client, err := mongo.Connect(ctx, o)
//	if err!=nil{
//		return nil,err
//	}
//	
//	return client.Database(name),nil
//}

type Student struct {
	Name string
	Age int
}

func main() {
	defer close()
	//InsertDocument()
	//InsertManyDocument()
	updateDocuemnt()
	
}

// 新增
func InsertDocument(){
	collection := Client.Database("test").Collection("student")

	s1 := Student{"小红",12}
	
	insertResult, err := collection.InsertOne(context.TODO(), s1)
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println("回显id",insertResult.InsertedID)
}

// 插入多条文档
func InsertManyDocument(){
	collection := Client.Database("test").Collection("student")
	
	s2 := Student{"小兰", 10}
	s3 := Student{"小黄", 11}

	studnets := []interface{}{s2, s3}
	insertResult, err := collection.InsertMany(context.TODO(), studnets)
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println("回显ids",insertResult.InsertedIDs)
}

// 更新文档
func updateDocuemnt(){
	collection := Client.Database("test").Collection("student")
	// 筛选
	filter := bson.D{{"name","小兰"}}
	update := bson.D{
		{
			"$inc",bson.D{
				{"age",1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err!=nil{
		log.Fatalln(err)
	}
	
	fmt.Println("更新后返回",updateResult.ModifiedCount,updateResult.MatchedCount)
}