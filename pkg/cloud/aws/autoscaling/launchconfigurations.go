package autoscaling

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func (s *Service) DeleteLaunchConfigurations(launchConfigList []string) error {

	for _, launchConfigName := range launchConfigList {

		fmt.Printf("Deleting launch configuration: %s", launchConfigName)

		input := &autoscaling.DeleteLaunchConfigurationInput{
			LaunchConfigurationName: &launchConfigName,
		}
		_, err := s.AutoScaling.DeleteLaunchConfiguration(input)
		if err != nil {
			return err
		}
	}

	return nil
}
