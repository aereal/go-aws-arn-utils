package awsarnutils

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

// ParseECSServiceARN extracts ECS cluster name and service name from serviceARN
func ParseECSServiceARN(serviceARN string) (clusterName, serviceName string, err error) {
	parsed, err := arn.Parse(serviceARN)
	if err != nil {
		return
	}
	if parsed.Service != "ecs" {
		err = ErrInvalidService
		return
	}
	parts := strings.Split(parsed.Resource, "/")
	if len(parts) != 3 {
		err = ErrIllFormedResource
		return
	}
	clusterName, serviceName = parts[1], parts[2]
	return
}
