package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/natemarks/aws-sdk-experiments/app/build"
	"github.com/natemarks/aws-sdk-experiments/pkg"
)

func main() {
	// Get the bucket name from flags
	println("Version: ", build.Version)
	bucket := flag.String("b", "", "The name of the bucket")
	flag.Parse()

	if *bucket == "" {
		fmt.Println("You must supply the name of a bucket (-b BUCKET)")
		return
	}

	// get the aws sdk client config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)

	// generate the list objects request
	input := &s3.ListObjectsV2Input{
		Bucket: bucket,
	}

	// Execute the request
	resp, err := pkg.GetObjects(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got error retrieving list of objects:")
		fmt.Println(err)
		return
	}

	// Print the list of objects in the bucket
	fmt.Println("Objects in " + *bucket + ":")

	for _, item := range resp.Contents {
		fmt.Println("Name:          ", *item.Key)
		fmt.Println("Last modified: ", *item.LastModified)
		fmt.Println("Size:          ", item.Size)
		fmt.Println("Storage class: ", item.StorageClass)
		fmt.Println("")
	}

	fmt.Println("Found", len(resp.Contents), "items in bucket", *bucket)
	fmt.Println("")
}
