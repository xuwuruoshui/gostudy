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
//	hello, err := mongo.Connect(ctx, o)
//	if err!=nil{
//		return nil,err
//	}
//	
//	return hello.Database(name),nil
//}

type Student struct {
	Name string
	Age int
}

func main() {
	defer close()
	//InsertDocument()
	//InsertManyDocument()
	//updateDocuemnt()
	//SearchDocument()
	//SearchManyDocument()
	//DeleteDocument()
	DeleteManyDocument()
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

// 查找文档
func SearchDocument(){
	collection := Client.Database("test").Collection("student")
	var result Student
	err := collection.FindOne(context.TODO(), bson.D{{"name", "小兰"}}).Decode(&result)
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println("单个document查询:",result)
}

// 查询多个文档
func SearchManyDocument(){
	collection := Client.Database("test").Collection("student")
	findOptions := options.Find()
	findOptions.SetLimit(2)
	
	

	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err!=nil{
		log.Fatalln(err)
	}

	var results []*Student
	for cur.Next(context.TODO()){
		var elemn Student
		err := cur.Decode(&elemn)
		if err!=nil{
			log.Fatalln(err)
		}
		results = append(results,&elemn)
	}

	if err := cur.Err();err!=nil{
		log.Fatalln(err)
	}
	
	// 完成后关闭游标
	cur.Close(context.TODO())
	fmt.Println("多条数据:")
	for _, v := range results {
		fmt.Println(v)
	}
}

// 删除文档
func DeleteDocument(){
	collection := Client.Database("test").Collection("student")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"name", "小黄"}})
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println("删除的数量:",deleteResult.DeletedCount)

	
}

func DeleteManyDocument()  {
	collection := Client.Database("test").Collection("student")
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err!=nil{
		log.Fatalln(err)
	}

	fmt.Println("删除的数量",deleteResult.DeletedCount)
}