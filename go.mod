module github.com/drnatewhims/connect-sdk-go

go 1.20

require (
	github.com/1Password/connect-sdk-go v1.5.1
	github.com/google/uuid v1.3.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/stretchr/testify v1.8.4
	github.com/uber/jaeger-client-go v2.30.0+incompatible
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/1Password/connect-sdk-go => ./
