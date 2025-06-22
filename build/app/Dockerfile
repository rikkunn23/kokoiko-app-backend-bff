FROM golang:1.23.0-bookworm
ENV TZ=Asia/Tokyo
RUN update-ca-certificates
ENV REPOSITORY github.com/rikkunn23/kokoiko-app-backend-bff
ARG WORKDIR=${GOPATH}/src/${REPOSITORY}
WORKDIR ${WORKDIR}
RUN apt-get update && apt-get install -y \
    jq nodejs npm
# swagger-cli install
RUN npm install -g @apidevtools/swagger-cli
RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/cosmtrek/air@v1.52.1
RUN go install github.com/go-delve/delve/cmd/dlv@v1.23.0
RUN go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0
ENTRYPOINT ["air", "-c", ".air.toml"]
