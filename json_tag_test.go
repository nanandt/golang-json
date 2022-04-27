package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTagEncode(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "HP Second",
		ImageURL: "https://example.com/image.png",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"Id":"P001","Name":"HP Second","Image_URL":"https://example.com/image.png"}`
	jsonByte := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonByte, product)
	fmt.Println(product)
}
