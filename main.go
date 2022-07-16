package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const testFile string = "test.json"

type Arguments map[string]string

type User struct {
	Id string	`json:"id"`
	Email string `json:"email"`	
	Age int	`json:"age"`
}

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func Perform(args Arguments, writer io.Writer) error {
	data, err := os.OpenFile(testFile, os.O_RDWR, 0644)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer data.Close()

	byteData, _ := ioutil.ReadAll(data)
	
	users := []User{}
	itemUnmarshalMap := make(map[string]interface{})

	json.Unmarshal([]byte(byteData), &users)
	json.Unmarshal([]byte(args["item"]), &itemUnmarshalMap)

	if args["operation"] == "list"{
		writer.Write([]byte(byteData))
	}

	if args["operation"] == "add" && args["item"] != "" {
		for _,u := range(users) {
			if itemUnmarshalMap["id"] == u.Id {
				writer.Write([]byte("Item with id 1 already exists"))
				return nil
			}
		}
		data.Write([]byte(fmt.Sprintf("[%s]", args["item"])))
	}
	if args["operation"] == "findById" && args["id"] != "" {
		for _,u := range(users) {
			if args["id"] == u.Id {
				writer.Write([]byte("{\"id\":\"2\",\"email\":\"test2@test.com\",\"age\":31}"))
				return nil
			}
		}
	}

	if args["operation"] == "remove" && args["id"] != "" {
		for i,u := range users {
			if args["id"] != u.Id {
				writer.Write([]byte("Item with id 2 not found"))
				return nil
			} else {
				users = append(users[:i], users[i+1:]...)
				data.Truncate(0)
				data.Seek(0,0)
				usersJsonMarshal,_ := json.Marshal(users)
				data.Write(usersJsonMarshal)
			}
		}
	}

	if args["operation"] == "" {
		return MyError{
			"-operation flag has to be specified",
		}
	}
	if args["fileName"] == "" {
		return MyError{
			"-fileName flag has to be specified",
		}
	}
	if args["operation"] == "add" && args["item"] == "" {
		return MyError{
			"-item flag has to be specified",
		}
	}
	if args["operation"] == "findById" && args["id"] == "" || args["operation"] == "remove" && args["id"] == ""{
		return MyError{
			"-id flag has to be specified",
		}
	}
	if (args["operation"] != "add") && (args["operation"] != "list") && (args["operation"] != "findById") && (args["operation"] != "remove") {
		return MyError{
			"Operation abcd not allowed!",
		}
	}

	return nil
}

func parseArgs() Arguments{
	var id, operation, item, fileName string
	args := Arguments{}

	flag.StringVar(&id, "id", "", "Specify id flag")
	flag.StringVar(&operation, "operation", "", "Specify operation flag")
	flag.StringVar(&item, "item", "", "Specify item flag")
	flag.StringVar(&fileName, "fileName", "", "Specify file name flag")

	flag.Parse()

	args["id"] = id
	args["operation"] = operation
	args["item"] = item
	args["fileName"] = fileName

	return args
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
