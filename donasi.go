package donasi

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"os"
	

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDonasi(jenis string, tanggal string, lokasi string, deskripsi string, biodata Donasi) (insertedID interface{}) {
	tanggalTime, _ := time.Parse("2006-01-02", tanggal)
	bencana := BencanaAlam{
		Jenis:     jenis,
		Tanggal:   tanggalTime,
		Lokasi:    lokasi,
		Deskripsi: deskripsi,
		Datetime:  [12]byte{},
	}
	return InsertOneDoc("Donasi", "donasi", bencana)
}

func GetInsertBencanaAlam(nim string) (m Donasi) {
	col := MongoConnect("Donasi").Collection("donasi")
	filter := bson.M{"nim": nim}
	err := col.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		fmt.Printf("GetBencanaAlam: %v\n", err)
	}
	return m
}



func GetAllDonasi() (data []Donasi) {
	karyawan := MongoConnect("Donasi").Collection("donasi")
	filter := bson.M{}
	cursor, err := karyawan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllDonasi: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

