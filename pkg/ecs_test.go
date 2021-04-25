package pkg

import (
	"strings"
	"testing"
)

func clusterArnsLookGood(arnList []string) bool {
	return strings.HasPrefix(arnList[0], "arn")
}

func TestECSGetClusters(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"simple", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ECSGetClusters()
			if (err != nil) != tt.wantErr {
				t.Errorf("ECSGetClusters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(clusterArnsLookGood(got)) {
				t.Errorf("First cluster arn doesn't begin with arn")
			}
		})
	}
}
