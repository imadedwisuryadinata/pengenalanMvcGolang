package model

import (
	"fmt"
	"log"
	"strings"

	"firebase.google.com/go/db"
)

//var data []Antrian

type Antrian struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}

//fungsi get antrian
func GetAntrian() (bool, []map[string]interface{}, error) {
	//return true, data, nil

	//menggunakan firebase:
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database: ", err)
		return false, nil, err
	}
	return true, data, nil
}

//fungsi menambahkan antrian
func AddAntrian() (bool, error) {
	_, dataAntrian, _ := GetAntrian()
	var ID string
	var antrianRef *db.Ref
	ref := client.NewRef("antrian")

	if dataAntrian == nil {
		ID = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	} else {
		ID = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}

	antrian := Antrian{
		ID:     ID,
		Status: false,
	}

	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

//fungsi update antrian
func UpdateAntrian(idAntrian string) (bool, error) {
	// for i := range data {
	// 	if data[i].ID == idAntrian {
	// 		data[i].Status = true
	// 		break
	// 	}
	// }
	//menggunakan firebase:
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	antrian := Antrian{
		ID:     idAntrian,
		Status: true,
	}

	if err := childRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

//fungsi delete antrian
func DeleteAntrian(idAntrian string) (bool, error) {
	// for i := range data {
	// 	if data[i].ID == idAntrian {
	// 		data = append(data[:i], data[i+1:]...)
	// 	}
	// }
	//menggunakan firebase:
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])

	if err := childRef.Delete(ctx); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
