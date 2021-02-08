package awsarnutils

import "errors"

var (
	// ErrInvalidService indicates the ARN is not an expected one
	ErrInvalidService = errors.New("invalid ARN: service is incorrect")

	// ErrIllFormedResource indicates resource part of the ARN is not an expected one
	ErrIllFormedResource = errors.New("invalid ARN: resource part is ill-formed")
)
