package DB

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var clientoptions *options.ClientOptions

func init() {
 clientoptions = options.Client().ApplyURI("mongodb+srv://karan5599:p3fPlJQI0LZsftj8@cluster0.xnnlgsv.mongodb.net/")
 var err error
 Client,err=mongo.Connect(context.TODO(),clientoptions)
 if err!=nil{
	log.Println(err)
 }
 err= Client.Ping(context.TODO(),nil)
 if err!=nil{
	log.Print(err)
 }
fmt.Println("mongo connected sucessfully")
}