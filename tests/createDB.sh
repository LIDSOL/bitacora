#!/usr/bin/env bash

podman run --detach \
    --replace \
    --name dbTest \
    -p 3306:3306 \
    -v `pwd`/datadb:/var/lib/mysql:Z \
    --env MYSQL_DATABASE=bitacoraDB \
    --env MARIADB_USER=bitacoraU \
    --env MARIADB_PASSWORD=test-passwd \
    --env MARIADB_ROOT_PASSWORD=root-passwd \
    docker.io/library/mariadb

echo 'mysql -h 127.0.0.1 -u bitacoraU -p bitacoraDB'
