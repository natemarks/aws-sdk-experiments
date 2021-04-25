package main

import (
	"fmt"
	"github.com/natemarks/aws-sdk-experiments/app/build"
	"github.com/natemarks/aws-sdk-experiments/pkg"
)

func main() {
	println("Version: ", build.Version)
	arns, _ := pkg.ECSGetClusters()
	for _, item := range arns {
		fmt.Println(item)
	}
	fmt.Println("Found", len(arns), "Cluster Arns")
}
