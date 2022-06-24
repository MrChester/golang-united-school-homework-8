package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

//const usersFile string = "users.json"

type Arguments map[string]string

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func Perform(args Arguments, writer io.Writer) error {
	if args["operation"] == "" {
		return MyError{
			"-operation flag has to be specified",
		}
	}
	if (args["operation"] != "add") && (args["operation"] != "list") && (args["operation"] != "findById") && (args["operation"] != "remove") {
		return MyError{
			"Operation abcd not allowed!",
		}
	}
	if args["fileName"] == "" {
		return MyError{
			"-fileName flag has to be specified",
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
