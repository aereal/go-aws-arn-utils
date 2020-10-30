![CI][ci-status]
[![PkgGoDev][pkg-go-dev-badge]][pkg-go-dev]

# go-aws-arn-utils

Utilities to extract more detailed components from AWS ARNs.

## Install

```sh
go get github.com/aereal/go-aws-arn-utils
```

## Usage

### Timestream

```go
import (
  "github.com/aereal/go-aws-arn-utils"
)

dbName, tableName, err := awsarnutils.ParseTimestreamTableARN("arn:aws:timestream:us-east-1:1234567890:database/db1/table/table1")
// dbName = "db1", tableName = "table1"
```

## License

See LICENSE file.

[pkg-go-dev]: https://pkg.go.dev/github.com/aereal/go-aws-arn-utils
[pkg-go-dev-badge]: https://pkg.go.dev/badge/aereal/go-aws-arn-utils
[ci-status]: https://github.com/aereal/go-aws-arn-utils/workflows/CI/badge.svg?branch=main
