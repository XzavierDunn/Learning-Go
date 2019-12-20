package main

import(
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sns"
)

func main() {
    ses := session.Must(session.NewSession())
    svc := sns.New(ses)

    params := &sns.PublishInput{
        Message: aws.String("beans"),
        PhoneNumber: aws.String("+19707992457"),
    }
    resp, err := svc.Publish(params)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println(resp)
}
