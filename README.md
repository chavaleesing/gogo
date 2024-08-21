# gogo

For run app
- use launch.json OR `go run app/main.go`

For run temp file
`go run temp.go`

## Removing dependency
1. Remove all the imports and code related to the package
2. Run `go mod tidy -v command` to remove any unused package