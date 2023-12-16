# Notes

## Adding gRPC

To add a gRPC service we first need to define a protobuf file, generate Go code
from it, and then implement the service.

1: Define a protobuf file `service.proto` in the service package.

```protobuf
syntax = "proto3";

package service;

service ServiceCatalog {
  rpc ListProducts (ListProductsRequest) returns (ListProductsResponse);
}

message ListProductsRequest {}

message ListProductsResponse {
  repeated Service products = 1;

  message Product {
    string id = 1;
    string name = 2;
    string description = 3;
    int32 version_count = 4;
  }
}
```

2: Generate Go code from the protobuf file. We'll need to install the `protoc`
compiler, and the `protoc-gen-go` and `protoc-gen-go-grpc` plugins first.

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service/service.proto
```

3: Implement the service in `service/grpc.go`.

```golang
package service

import (
  "context"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)

type ServiceCatalogServer struct {
  UnimplementedServiceCatalogServer
}

func (s *ServiceCatalogServer) ListProducts(ctx context.Context, req *ListProductsRequest) (*ListProductsResponse, error) {
  // TODO: Implement this
  return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
```

4: Register the service with a gRPC server in the `main` function.

```golang
import (
  "google.golang.org/grpc"
  "net"
)

func main() {
  lis, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  grpcServer := grpc.NewServer()
  service.RegisterServiceCatalogServer(grpcServer, &service.ServiceCatalogServer{})
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
```
