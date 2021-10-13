package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type text struct {
	Title string `json:"heading,omitempty"`
	Words string `json:"content,omitempty"`
}

type book struct {
	abc  text
	isbn int
	text int
}

func main() {
	// t := text{
	// 	Title: "sample text",
	// 	Words: "text content",
	// }

	// fmt.Printf("Text Type:%T\n", t)
	// fmt.Printf("Text content:%+v\n", t)

	// textJson, err := json.Marshal(t)
	// if err != nil {
	// 	fmt.Printf("Error:%v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("Text as json:%+v\n\n", string(textJson))

	t := []byte(`{"heading":"sample text","content":"text content"}`)

	var text text
	err := json.Unmarshal(t, &text)
	if err != nil {
		fmt.Printf("Error:%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Text Type:%T\n", text)
	fmt.Printf("Text content:%+v\n", text)

	// fmt.Println("***********")

	// b := book{
	// 	abc:  t,
	// 	isbn: 1234,
	// 	text: 12,
	// }
	// fmt.Printf("Book Type:%T\n", b)
	// fmt.Printf("Book content:%+v\n", b)
	// fmt.Printf("Book title:%v\n", b.abc.title)
}
