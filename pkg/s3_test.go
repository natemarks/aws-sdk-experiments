package pkg

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"testing"
)

func TestGetObjects(t *testing.T) {
	tests := []struct {
		name    string
		bucket  string
		want    string
		wantErr bool
	}{
		{"simple", "com.imprivata.371143864265.us-east-1.personal", "sdf", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := config.LoadDefaultConfig(context.TODO())
			if err != nil {
				panic("configuration error, " + err.Error())
			}

			client := s3.NewFromConfig(cfg)

			input := &s3.ListObjectsV2Input{
				Bucket: &tt.bucket,
			}

			got, err := GetObjects(context.TODO(), client, input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || len(got.Contents) == 0 {
				t.Errorf("Got an empty list")
			}
		})
	}
}
