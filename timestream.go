package awsarnutils

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
)

var (
	// ErrInvalidService indicates the ARN is not of Timestream
	ErrInvalidService = errors.New("invalid ARN: service is not timestream")

	// ErrIllFormedResource indicates resource part of the ARN is not a table
	ErrIllFormedResource = errors.New("invalid ARN: resource part is ill-formed")
)

// ParseTimestreamTableARN extracts Timestream database name and table name from tableArn
func ParseTimestreamTableARN(tableArn string) (databaseName string, tableName string, err error) {
	parsed, err := arn.Parse(tableArn)
	if err != nil {
		return
	}
	if parsed.Service != "timestream" {
		err = ErrInvalidService
		return
	}
	parts := strings.Split(parsed.Resource, "/")
	if len(parts) != 4 {
		err = ErrIllFormedResource
		return
	}
	databaseName, tableName = parts[1], parts[3]
	return
}
