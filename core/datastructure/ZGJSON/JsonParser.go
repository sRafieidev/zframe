package ZGJSON

import (
	"encoding/json"
	"fmt"
)

func TestJsonParse() {
	var JsonData string = "{\n  \"users\": [\n    {\n      \"name\": \"Elliot\",\n      \"type\": \"Reader\",\n      \"age\": 23,\n      \"social\": {\n        \"facebook\": \"https://facebook.com\",\n        \"twitter\": \"https://twitter.com\"\n      }\n    },\n    {\n      \"name\": \"Fraser\",\n      \"type\": \"Author\",\n      \"age\": 17,\n      \"social\": {\n        \"facebook\": \"https://facebook.com\",\n        \"twitter\": \"https://twitter.com\"\n      }\n    }\n  ]\n}"
	TestJsonParsing(JsonData)
}

func TestJsonParsing(inputjsondata string) {

	byteValue := []byte(inputjsondata)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return
	}
	//color.Blue(result["users"])
	fmt.Println(result["users"])

}
