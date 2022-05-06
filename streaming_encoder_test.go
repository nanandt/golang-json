package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Muhamad",
		MiddleName: "Fatih",
		LastName:   "Rizky",
		Married:    false,
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}
