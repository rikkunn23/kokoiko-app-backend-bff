.PHONY: build
# 機能カテゴリが増えたら半角空白区切りでTAGSに追加する。
TAGS := "job master user entry work auth identification"
EXEC_APP := docker compose exec app
EXEC_APP_TEST := docker compose exec app-test
EXEC_T_APP := docker compose exec -T app
EXEC_T_APP_TEST := docker compose exec -T app-test
PKGS = $(shell arc= $(EXEC_T_APP) go list ./... | grep -v vendor )
INTERFACES = $(shell arc= $(EXEC_T_APP) find -name "interface.go" )
GO_VER = 1.23

init-env:
  @cp -p ./deploy/env/.env.local ./deploy/env/.env.private.local

init-doc:
  @make init-docb br=main

init-docb:
  rm -rf ./baitorufree-free-app-backend-api-doc
  git clone -b $(br) https://github.com/rikkunn23/kokoiko-app-backend-api-doc.git
  chmod -R 755 ./kokoiko-app-backend-api-doc

up:
  @docker compose up -d
# CIでのみ使用

up-test:
  @make create-network
  @docker compose up -d app-test postgres-test localstack
  @make wait-test-postgres

c-up:
  @rm -rf ./.localstack
  @docker compose up -d --build

down:
  @docker compose down

ps:
  @docker compose ps

tidy:
  @$(EXEC_T_APP) go mod tidy -go=${GO_VER}

pretest:
  @$(EXEC_T_APP_TEST) golangci-lint run --timeout 10m
  @$(EXEC_T_APP_TEST) sh ./build/format-check.sh

test:
  @$(EXEC_T_APP_TEST) go test ./... -v -count=1 -coverprofile=cover.out ./...
  @$(EXEC_T_APP_TEST) go tool cover -html=cover.out -o cover.html

gen-api:
  ./tools/swagger/codegen.sh ${TAGS}

gen-ms-user:
  ./tools/swagger/codegen_ms.sh "user"

gen-ms-job:
  ./tools/swagger/codegen_ms.sh "job"

gen-api-all:
  @make gen-api
  @make gen-ms-user
  @make gen-ms-job

gen-mock:
  @$(foreach src,$(INTERFACES), docker compose exec app bash tools/mockgen/codegen.sh ${src} || exit;)

enter-app:
  @$(EXEC_APP) bash

e2e:
  @$(EXEC_APP_TEST) runn run tools/runn/e2e.yaml

wait-test-postgres:
  @sh ./build/postgres/wait_for_postgres.sh -n postgres-test
# 初回起動時などネットワークが存在しない場合は作成する

create-network:
  @docker network create external
