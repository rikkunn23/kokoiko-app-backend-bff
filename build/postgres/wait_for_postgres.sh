#!/bin/bash

usage() {
  echo "Usage: $0 -n container_name" 1>&2
  exit 1
}
while getopts n: OPT
do
  case ${OPT} in
    n) CONTAINER_NAME=${OPTARG}
      ;;
    \?) usage
      ;;
  esac
done
if [ ${CONTAINER_NAME} = "" ]; then
  usage
fi
echo ${CONTAINER_NAME}
# コンテナの存在確認
if [ $(docker ps -f name=${CONTAINER_NAME} -q) = "" ]; then
  echo "given a invalid container_name"
  exit 1
fi
# posが使用可能になり初期化SQLを流し終えると、'DATA BASE IS READY!!'というログを吐くので、
# 30秒ごとにログをチェックする
echo '==================================='
echo '# Start waiting for PostgreSQL Database'
echo '==================================='
for i in `seq 1 40`; do
  # 「2>/dev/null」で標準ログのみ抽出する
  if [ $(docker logs ${CONTAINER_NAME} 2>/dev/null | grep -c 'DATABASE IS READY!!') -gt 0 ]; then
    echo "PostgreSQL DATABASE IS READY TO USE!"
    exit 0
  else
    echo "WAITING FOR PostgreSQL DATABASE..."
    sleep 30
  fi
done
exit 1
