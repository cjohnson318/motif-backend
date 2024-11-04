package main

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "google.golang.org/grpc"
    "net/http"
    pb "github.com/cjohnson318/motif-backend/pkg/proto"

)

func main() {
    fmt.Println("Hello")
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
        if err != nil {
            panic(err)
        }
        defer conn.Close()

        client := pb.NewGreeterClient(conn)

        ctx := context.Background()
        resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Dork"})
        if err != nil {
            panic(err)
        }

        fmt.Println("Greeting:", resp.Message)
        c.JSON(http.StatusOK, gin.H{"data": resp.Message})    
    })

    r.Run(":50052")
}
