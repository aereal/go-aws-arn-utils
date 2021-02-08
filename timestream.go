package awsarnutils

import (
	"strings"

	arnv2 "github.com/aws/aws-sdk-go-v2/aws/arn"
)

// ParseTimestreamTableARN extracts Timestream database name and table name from tableArn
func ParseTimestreamTableARN(tableArn string) (databaseName string, tableName string, err error) {
	parsed, err := arnv2.Parse(tableArn)
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
