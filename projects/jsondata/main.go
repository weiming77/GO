/*
Function/Type	Work with
Unmarshal		[]byte -> Go
Marshal			Go -> []byte
Decoder			io.Reader -> Go
Encoder			Go -> io.Writer

Go Type 				JSON Type
bool					Boolean
int,float64				number
string					string
nil						null
[]interface{}			array
map[string]interface{}	object
time.Time				string
[]bytes					string(base64 encoded)
*/
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const jsonData = `{"intValue":123456,
"boolValue":false,
"stringValue":"Hell No!",
"dateValue":"2022-03-02T09:10:00Z",
"objectValue":{"arrayValue":[1,2,3,4,5]},
"nullStringValue":null,
"nullIntValue":null}`

type TMyJSON struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"dateValue"`
	ObjectValue     *TMyOBJ   `json:"myOBJ"`
	NullStringValue *string   `json:"nullStringValue"`
	NullIntValue    *int      `json:"nullIntValue`
}

type TMyOBJ struct {
	ArrayValue []int `json:"arrayValue"`
}

func genJSONbyMap() {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello world",
		"dateValue":   time.Date(2023, 5, 1, 17, 14, 45, 200, time.UTC),
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 5, 6},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
}

func parseJSONbyMap(JSONData string) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(JSONData), &data)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("JSON map: %v\n", data)
}

func genJSONbyStruct() {
	otherInt := 4321
	otherData := &TMyJSON{
		IntValue:    1234,
		BoolValue:   false,
		StringValue: "Hello Hell",
		DateValue:   time.Date(2023, 5, 1, 9, 30, 45, 600, time.UTC),
		ObjectValue: &TMyOBJ{
			ArrayValue: []int{1, 1, 2, 2, 2, 3, 34, 4},
		},
		NullStringValue: nil,
		NullIntValue:    &otherInt,
	}

	jsonOther, er := json.Marshal(otherData)
	if er != nil {
		fmt.Printf("could not marshal json: %s\n", er)
		return
	}

	fmt.Printf("json Other: %s\n", jsonOther)
}

func parseJSONbyStruct(JSONData string) {
	var data *TMyJSON
	err := json.Unmarshal([]byte(JSONData), &data)
	if err != nil {
		fmt.Printf("Coild not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json struct: %#v\n", data)
	fmt.Printf("dateValue: %#v\n", data.DateValue)
	fmt.Printf("objectValue: %#v\n", data.ObjectValue)
}

func main() {
	MainFunc()
	genJSONbyMap()
	genJSONbyStruct()
	parseJSONbyMap(jsonData)
	parseJSONbyStruct(jsonData)
}
