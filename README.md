# aos_common

[![pipeline status](https://gitpct.epam.com/epmd-aepr/aos_common/badges/master/pipeline.svg)](https://gitpct.epam.com/epmd-aepr/aos_common/commits/master) [![coverage report](https://gitpct.epam.com/epmd-aepr/aos_common/badges/master/coverage.svg)](https://gitpct.epam.com/epmd-aepr/aos_common/commits/master)

Contains common aos packages:

* [umprotocol](doc/umprotocol.md) - AOS update manager protocol.

Generate gRPC api:

```bash
cd api
protoc --go_out=./iamanager  --go_opt=paths=source_relative --go-grpc_out=./iamanager  --go-grpc_opt=paths=source_relative iamanager/iamanager.proto -I iamanager/ -I ./iamanager/
protoc --go_out=./iamanager --go_opt=paths=source_relative --go-grpc_out=./iamanager --go-grpc_opt=paths=source_relative iamanager/iamanagerpublic.proto -I ./iamanager/
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative updatemanager/updatemanager.proto
```
