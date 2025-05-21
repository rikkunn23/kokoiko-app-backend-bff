TAGS=$1
BASEDIR=/go/src/github.com/rikkunn23/kokoiko-app-backend-bff
BASE_INFILE=${BASEDIR}/kokoiko-app-backend-api-doc/bff/01_IF/openapi.yaml
INFILE=${BASEDIR}/kokoiko-app-backend-api-doc/bff/01_IF/openapi_merge.yaml

docker compose exec -T app swagger-cli bundle -o ${INFILE} -t yaml ${BASE_INFILE}

for tag in $TAGS
do
  mkdir -p gen/api/${tag}
  rm -f gen/api/${tag}/api.go
  OUTFILE=$(pwd)/gen/api/${tag}/api.go
  # テンプレート指定をする場合
  # docker compose exec -T app oapi-codegen -package ${tag} -include-tags ${tag} -generate types,chi-server -templates ./tools/swagger/templates ${INFILE} > ${OUTFILE}
  docker compose exec -T app oapi-codegen -package ${tag} -include-tags ${tag} -generate types,chi-server ${INFILE} > ${OUTFILE}
  echo "finished gen ${tag}"
done

docker compose exec -T app rm -f ${INFILE}
