package namapackage

import (
	"fmt"
	"testing"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	
)

func TestInsertDonasi(t *testing.T) {
	jenis := "Bencana Banjir"
	tanggal := "2024-05-16" // Format tanggal sesuaikan dengan kebutuhan
	lokasi := "Rumah"
	deskripsi := "Donasi untuk korban bencana banjir di daerah X"
	biodata := Donasi{
		Nama:         "Hammi",
		Phone_number: "628456456211",
		Jenis_donasi: "Uang",
		Jumlah:       1000000,
	}
	hasil := InsertDonasi(jenis, tanggal, lokasi, deskripsi, biodata)
	fmt.Println(hasil)
}

func GetDonasiFromPhoneNumber(phone_number string) (donasi Donasi) {
	donasiCollection := MongoConnect("Donasi").Collection("donasi")
	filter := bson.M{"phone_number": phone_number}
	err := donasiCollection.FindOne(context.TODO(), filter).Decode(&donasi)
	if err != nil {
		fmt.Printf("GetDonasiFromPhoneNumber: %v\n", err)
	}
	return donasi
}

func InsertBencanaAlam(jenis string, lokasi string, tanggal time.Time, deskripsi string) (insertedID interface{}) {
    bencana := BencanaAlam{
        Jenis:     jenis,
        Lokasi:    lokasi,
        Tanggal:   tanggal,
        Deskripsi: deskripsi,
    }
    insertResult, err := MongoConnect("Donasi").Collection("bencana_alam").InsertOne(context.TODO(), bencana)
    if err != nil {
        fmt.Printf("InsertBencanaAlam: %v\n", err)
    }
    fmt.Println("InsertedID:", insertResult.InsertedID)
    return insertResult.InsertedID
}


func TestGetAllDonasi(t *testing.T) {
	data := GetAllDonasi()
	fmt.Println(data)
}
