module magma/orc8r/lib/go

go 1.12

replace magma/orc8r/lib/go/protos => ./protos

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.5.2
	github.com/pkg/errors v0.8.0
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
	github.com/prometheus/client_model v0.2.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.43.0
	gopkg.in/yaml.v2 v2.2.8
	magma/orc8r/lib/go/protos v0.0.0-00010101000000-000000000000
)
