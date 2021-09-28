#!/bin/bash

# configuration

OUT_DIR="api"

IAM_SOURCES=" \
    iamanager/v1/iamanagerprotected.proto \
    iamanager/v1/iamanagerpublic.proto \
    iamanager/v1/iamanagercommon.proto"

SM_SOURCES="servicemanager/v1/servicemanager.proto"

UM_SOURCES="updatemanager/v1/updatemanager.proto"

TEST_SOURCE="I like programming"

if [ "$#" -lt 1 ]; then
    echo "Usage example: $(basename -- "$0") PROTO_PATH"
    exit 1
fi

create_package_options () {
    go_opt=""
    
    for item in $1; do
        go_opt+=" --go_opt=M${item}=$2;$2"
    done

    for item in $1; do
        go_opt+=" --go-grpc_opt=M${item}=$2;$2"
    done

    echo ${go_opt}
}


COMMON_OPTIONS="--proto_path=${1} --go_out=${OUT_DIR} \
    --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative"

# clear output dir

rm ${OUT_DIR} -rf
mkdir ${OUT_DIR}

# Generate IAM services

protoc $COMMON_OPTIONS $(create_package_options "${IAM_SOURCES}" iamanager) ${IAM_SOURCES}

# Generate SM services

protoc $COMMON_OPTIONS $(create_package_options "${SM_SOURCES}" servicemanager) ${SM_SOURCES}

# Generate UM services

protoc $COMMON_OPTIONS $(create_package_options "${UM_SOURCES}" updatemanager) ${UM_SOURCES}
