package awsarnutils

import "testing"

func TestParseECSServiceARN(t *testing.T) {
	cases := []struct {
		name            string
		serviceArn      string
		wantClusterName string
		wantServiceName string
		wantErr         bool
	}{
		{"ok", "arn:aws:ecs:us-east-1:1234567890:service/cluster1/service1", "cluster1", "service1", false},
		{"invalid ARN", "http://example.com/", "", "", true},
		{"other service", "arn:aws:dynamodb:us-east-1:1234567890:table/table1", "", "", true},
		{"ill-formed resource", "arn:aws:ecs:us-east-1:1234567890:service/cluster1", "", "", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotClusterName, gotServiceName, err := ParseECSServiceARN(c.serviceArn)
			if (err != nil) != c.wantErr {
				t.Errorf("error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if gotClusterName != c.wantClusterName {
				t.Errorf("gotClusterName = %v, want %v", gotClusterName, c.wantClusterName)
			}
			if gotServiceName != c.wantServiceName {
				t.Errorf("gotServiceName = %v, want %v", gotServiceName, c.wantServiceName)
			}
		})
	}
}
