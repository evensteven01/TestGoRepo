package main

import (
    "fmt"
    "context"
    "github.com/aws/aws-lambda-go/lambda"
    "log"
)

type MyEvent struct {
    Name string `json: "name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
    log.Print("Handling event " + name.Name)
    return fmt.Sprintf("Hello %s!", name.Name), nil
}

func AddYouSaid(str string) {
    return "You said: " + str
}

func main() {
    fmt.Printf("Hello, world!\n")
    lambda.Start(HandleRequest)
}
