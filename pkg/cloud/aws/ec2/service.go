package ec2

import (
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type Service struct {
	EC2 ec2iface.EC2API
}

// NewService returns a new service given the ec2 api client.
func NewService(i ec2iface.EC2API) *Service {
	return &Service{
		EC2: i,
	}
}
