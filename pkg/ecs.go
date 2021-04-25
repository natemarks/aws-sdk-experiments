package pkg

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

// ECSGetClusters retrieves a list of cluster arns as a slice of strings
func ECSGetClusters() ([]string, error) {

	// get the aws sdk client config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ecs.NewFromConfig(cfg)

	input := &ecs.ListClustersInput{}

	result, err := client.ListClusters(context.TODO(), input)
	if result == nil {

		return []string{}, err
	}
	return result.ClusterArns, err

}
