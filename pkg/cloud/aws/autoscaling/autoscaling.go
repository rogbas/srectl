package autoscaling

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/pkg/errors"
)

func (s *Service) GetAutoScallingGroupsByCluster(clusterName string) ([]*autoscaling.Group, error) {
	asGroups := []*autoscaling.Group{}

	// input := &autoscaling.DescribeAutoScalingGroupsInput{}

	out, err := s.AutoScaling.DescribeAutoScalingGroups(nil)
	switch {
	// case IsNotFound(err):
	// 	return nil, nil
	case err != nil:
		return nil, errors.Wrap(err, "failed to describe instances by tags")
	}

	for _, asg := range out.AutoScalingGroups {
		for _, tag := range asg.Tags {
			if (*tag.Key == "clusterid") && (*tag.Value == clusterName) {
				asGroups = append(asGroups, asg)
				break
			}
		}
	}
	// fmt.Printf("ASGs: %s", out)
	return asGroups, nil
}

func (s *Service) DeleteAutoScalingGroupsByCluster(clusterName string) ([]string, error) {
	var launchConfigList []string

	asgs, err := s.GetAutoScallingGroupsByCluster(clusterName)
	if err != nil {
		return nil, err
	}

	launchConfigList, err = s.DeleteAutoScalingGroupsAndWait(asgs)
	if err != nil {
		return nil, err
	}

	return launchConfigList, nil
}

func (s *Service) DeleteAutoScalingGroupsAndWait(asgs []*autoscaling.Group) ([]string, error) {
	forceDelete := true
	var asgsNames []string
	var launchConfigList []string

	for _, asg := range asgs {

		// Keep list of deleted ASGs for the waiting step
		asgsNames = append(asgsNames, *asg.AutoScalingGroupName)
		launchConfigList = append(launchConfigList, *asg.LaunchConfigurationName)

		fmt.Printf("Deleting autoscaling group: %s\n", *asg.AutoScalingGroupName)
		input := &autoscaling.DeleteAutoScalingGroupInput{
			AutoScalingGroupName: asg.AutoScalingGroupName,
			ForceDelete:          &forceDelete,
		}

		_, err := s.AutoScaling.DeleteAutoScalingGroup(input)
		if err != nil {
			return nil, err
		}
	}

	fmt.Printf("Waiting for deletion to be completed on: %s\n", asgsNames)

	inputWait := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: aws.StringSlice(asgsNames),
	}

	err := s.AutoScaling.WaitUntilGroupNotExists(inputWait)
	if err != nil {
		return nil, err
	}

	return launchConfigList, nil
}
