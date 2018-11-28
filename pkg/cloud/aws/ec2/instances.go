package ec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

func (s *Service) GetInstanceByCluster(clusterName string) ([]*ec2.Instance, error) {
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
		return nil, errors.Wrap(err, "failed to describe instances by tags")
	}

	for _, res := range out.Reservations {
		for _, inst := range res.Instances {

			instances = append(instances, inst)
		}
	}

	// fmt.Printf("Instances: %s", instances)
	return instances, nil
}
