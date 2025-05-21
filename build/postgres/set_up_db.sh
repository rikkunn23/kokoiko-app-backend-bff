#!/bin/bash

set +e
LOG_FILE=/dev/null
echo "CREATING DATABASE..."
PGDATA=var/lib/postgresql/data
# posの初期設定と起動
su postgres -c "initdb -D /var/lib/postgresql/data"
echo host all all 0.0.0.0/0 trust>>/var/lib/postgresql/data/pg_hba.conf
su postgres -c "/usr/lib/postgresql/14/bin/pg_ctl start -D ${PGDATA} -s  -w -t 300"
su postgres -c "/usr/lib/postgresql/14/bin/pg_ctl status -D ${PGDATA}"
psql -h localhost -p 5432 -U postgres -d postgres -c "CREATE ROLE kokoiko_admin with superuser login PASSWORD 'password123';"
psql -h localhost -p 5432 -U postgres -d postgres -c "CREATE DATABASE kokoiko OWNER kokoiko_admin;"
# DDLの実行
dirs=(
    "ddl/init/roles"
    "ddl/init/schemas"
    "ddl/create/master"
    "seed/master"
)
for dir in ${dirs[@]}; do
    echo $dir"を実行";
    cd /postgres/$dir
    for file in `\find . -type f -name "*.sql"`; do
        echo $dir$file"を実行";
        psql -h localhost -p 5432 -U kokoiko_admin -d kokoiko -f "$file"
    done
done
echo "DATABASE IS READY!!";
set -e
/usr/bin/tail -f ${LOG_FILE}
