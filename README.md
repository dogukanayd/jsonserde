# jsonserde

[![Coverage Status](https://coveralls.io/repos/github/dogukanayd/jsonserde/badge.svg?branch=main)](https://coveralls.io/github/dogukanayd/jsonserde?branch=main)

I developed this very easy package to use with my projects which I need to send a log to the AWS Athena.
As you know you can't store a log in S3 as JSON Array, you need a different format to create a table on
Athena using the data in the S3 bucket.

## Get the package
```sh
go get github.com/dogukanayd/jsonserde
```

## Usage

```go
	exampleData := []byte(`[{"name":"John","age":30},{"name":"Mike","age":25}]`)
	output, _ := jsonserde.Convert(exampleData)
    // output: {"age":30,"name":"John"}{"age":25,"name":"Mike"}
```

now you can put the result to the S3 and create a table on Athena
