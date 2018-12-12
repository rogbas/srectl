package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/glog"
)

func (s *Service) TerminateInstance(instanceID string) error {
	fmt.Printf("Attempting to terminate instance with id %q", instanceID)

	input := &ec2.TerminateInstancesInput{
		InstanceIds: aws.StringSlice([]string{instanceID}),
	}

	if _, err := s.EC2.TerminateInstances(input); err != nil {
		return err
	}

	glog.V(2).Infof("Terminated instance with id %q", instanceID)
	return nil
}

func (s *Service) TerminateInstanceAndWait(instanceID string) error {
	if err := s.TerminateInstance(instanceID); err != nil {
		return err
	}

	fmt.Printf("Waiting for EC2 instance with id %q to terminate", instanceID)

	input := &ec2.DescribeInstancesInput{
		InstanceIds: aws.StringSlice([]string{instanceID}),
	}

	if err := s.EC2.WaitUntilInstanceTerminated(input); err != nil {
		return err
	}

	return nil
}

func (s *Service) TerminateInstancesByCluster(clusterName string) error {
	// glog.V(2).Infof("Looking for existing instance for machine %q in cluster %q", machine.Name, cluster.Name)

	instances := []*ec2.Instance{}

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			s.filterByClusterId(clusterName),
		},
	}

	out, err := s.EC2.DescribeInstances(input)
	switch {
	// case IsNotFound(err):
	// 	return nil, nil
	case err != nil:
		return err
	}

	for _, res := range out.Reservations {
		for _, inst := range res.Instances {

			instances = append(instances, inst)
		}
	}

	// fmt.Printf("Instances: %s", instances)
	return nil
}
