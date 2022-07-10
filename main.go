package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name string
	Age  int
}

// JSON文字列 -> Go構造体
func jsonToStruct() {
	jsonStr := `{"name": "user1", "age":35 }`

	////// json.Unmarshalを使用した場合
	var user User
	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user) // {Name:user1 Age:35}

	////// json.Decoderを使用した場合
	if err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&user); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user) // {Name:user1 Age:35}
}

// Go構造体 -> JSON文字列
func StructToJson() {
	user := User{Name: "user2", Age: 33}

	////// json.Marshalを使用した場合
	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", jsonStr) // {"Name":"user2","Age":33}

	////// json.Encoderを使用した場合
	err = json.NewEncoder(os.Stdout).Encode(user)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func jsonArrayToStruct() {
	jsonStr := `[{"name": "user1", "age":35},{"name": "user2", "age":36}]`
	var users []User
	if err := json.Unmarshal([]byte(jsonStr), &users); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", users) // [{Name:user1 Age:35} {Name:user2 Age:36}]
}

func arrayToJson() {
	users := []User{{Name: "user1", Age: 35}, {Name: "user2", Age: 36}}
	jsonStr, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", jsonStr) // [{"Name":"user1","Age":35},{"Name":"user2","Age":36}]
}

func main() {
	jsonToStruct()
	StructToJson()
	jsonArrayToStruct()
	arrayToJson()
}
