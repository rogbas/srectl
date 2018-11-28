package autoscaling

import "github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"

type Service struct {
	AutoScaling autoscalingiface.AutoScalingAPI
}

// NewService returns a new service given the ec2 api client.
func NewService(i autoscalingiface.AutoScalingAPI) *Service {
	return &Service{
		AutoScaling: i,
	}
}
