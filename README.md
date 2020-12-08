# aos_common

[![pipeline status](https://gitpct.epam.com/epmd-aepr/aos_common/badges/master/pipeline.svg)](https://gitpct.epam.com/epmd-aepr/aos_common/commits/master) [![coverage report](https://gitpct.epam.com/epmd-aepr/aos_common/badges/master/coverage.svg)](https://gitpct.epam.com/epmd-aepr/aos_common/commits/master)

Contains common aos packages:

* [umprotocol](doc/umprotocol.md) - AOS update manager protocol.

Generate gRPC api:

```bash
cd api
protoc -I iamanager/ iamanager/iamanager.proto --go_out=plugins=grpc:iamanager --go_opt=paths=source_relative

protoc -I updatemanager/ updatemanager/updatemanager.proto --go_out=plugins=grpc:updatemanager --go_opt=paths=source_relative
```
