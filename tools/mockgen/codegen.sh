#!/bin/bash

# mockを作成したいInterfaceが定義されているファイルを指定すると、同じディレクトリにmock.goができる

SRC=$1
DIR=dirname ${SRC}
PKG=echo ${DIR} | sed  "s/\//\ /g" | awk '{print $NF}'
GEN_FILE=${DIR}/mock.go

mockgen -source=${SRC} -destination ${GEN_FILE} -package ${PKG}
echo "create ${GEN_FILE}"
