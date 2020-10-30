package awsarnutils

import "testing"

func TestParseTimestreamTableARN(t *testing.T) {
	cases := []struct {
		name             string
		tableArn         string
		wantDatabaseName string
		wantTableName    string
		wantErr          bool
	}{
		{"ok", "arn:aws:timestream:us-east-1:1234567890:database/db1/table/table1", "db1", "table1", false},
		{"invalid ARN", "http://example.com/", "", "", true},
		{"other service", "arn:aws:dynamodb:us-east-1:1234567890:table/table1", "", "", true},
		{"ill-formed resource", "arn:aws:timestream:us-east-1:1234567890:database/db1", "", "", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotDatabaseName, gotTableName, err := ParseTimestreamTableARN(c.tableArn)
			if (err != nil) != c.wantErr {
				t.Errorf("error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if gotDatabaseName != c.wantDatabaseName {
				t.Errorf("gotDatabaseName = %v, want %v", gotDatabaseName, c.wantDatabaseName)
			}
			if gotTableName != c.wantTableName {
				t.Errorf("gotTableName = %v, want %v", gotTableName, c.wantTableName)
			}
		})
	}
}
