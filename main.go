package main


import (
    "context"
    "fmt"
    "io/ioutil"
    "log"

    "cloud.google.com/go/storage"
    "google.golang.org/api/option"
    //     "database/sql"
    // "fmt"

    // _ "github.com/lib/pq"
)

func main() {
    // Build GCS client
    client, err := storage.NewClient(context.TODO(), option.WithoutAuthentication(), option.WithEndpoint("http://gcs:4443/storage/v1/"))
    if err != nil {
        log.Fatalf("failed to create client: %v", err)
    }
    const (
        bucketName = "tables"
        fileKey    = "users.json"
    )

    // Read file
    reader, err := client.Bucket(bucketName).Object(fileKey).NewReader(context.TODO())
    if err != nil {
        fmt.Println(err)
    }
    defer reader.Close()

    data, _ := ioutil.ReadAll(reader)

    fmt.Printf("contents of %s/%s: %s\n", bucketName, fileKey, data)

    // Build PostgreSQL client
    // db, err := sql.Open("postgres", "host=postgresql port=5432 user=postgres sslmode=disable")
    // defer db.Close()

    // if err != nil {
    //     fmt.Println(err)
    // }

    // // INSERT
    // res, err := db.Exec("INSERT INTO users (id, name, age) VALUES (1, 'tom', 18)")
    // if err != nil {
    //     fmt.Println(err)
    // }

    // fmt.Printf("%+v", res)

    // // SELECT
    // rows, err := db.Query("SELECT * FROM users")

    // if err != nil {
    //     fmt.Println(err)
    // }

    // fmt.Printf("%+v", *rows)
}
