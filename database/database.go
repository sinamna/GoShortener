package database
import (
	"fmt"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"ChizShortener/graph/model"
	"crypto/sha256"
)


type Database struct{
	client *mongo.Client
}
func (db *Database)getChizCollection()*mongo.Collection{
	return db.client.Database("ChizDatabase").Collection("links")
}
func ConnectDb() *Database{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to Database")
	return &Database{
		client: client,
	}
}

func (db *Database) SaveLink(ctx context.Context,newLink *model.NewLink) *model.Link{
	shortLink,err := generateHash(ctx,newLink.LongLink,db)
	if err!=nil{
		log.Fatal(err)
	}
	linkToSave :=&model.Link{
		ShortLink:shortLink,
		LongLink:newLink.LongLink,
	}
	linkCollection:=db.getChizCollection()
	_, err = linkCollection.InsertOne(ctx, linkToSave)
	if err!=nil{
		log.Fatal(err)
	}
	return linkToSave
}
func (db *Database) GetAllLinks(ctx context.Context) []*model.Link{
	linkCollection:= db.getChizCollection()
	cursor, err := linkCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var links []*model.Link
	for cursor.Next(ctx) {
		var link model.Link
		if err = cursor.Decode(&link); err != nil {
			log.Fatal(err)
		}
		links=append(links,&link)
	}
	return links
}
func generateHash(ctx context.Context,longLink string,db *Database) (string,error){
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(longLink)))[:6]
	linkCollection := db.getChizCollection()
	var savedLink bson.M
	if err := linkCollection.FindOne(ctx,bson.M{"shortLink":hash}).Decode(&savedLink); err!=nil{
		return hash,nil
	}
	return generateHash(ctx,longLink+"1",db)
}
func(db *Database) GetLink(ctx context.Context,identifier string,wantedLink string) (*model.Link,error){
	linkCollection := db.getChizCollection()
	var link model.Link
	var all []model.Link
	cursor, err := linkCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &all); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v",all)
	if err := linkCollection.FindOne(ctx,bson.M{identifier:wantedLink}).Decode(&link); err!=nil{
		return nil,err
	}
	return &link,nil
}
