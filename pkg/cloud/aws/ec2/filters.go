package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	ClusterIdKey = "tag:clusterid"
)

// Returns an EC2 filter based on cluster Id
func (s *Service) filterByClusterId(clusterName string) *ec2.Filter {
	return &ec2.Filter{
		Name:   aws.String(ClusterIdKey),
		Values: aws.StringSlice([]string{clusterName}),
	}
}
