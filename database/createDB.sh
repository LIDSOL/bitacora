#!/usr/bin/env bash

podman run --detach \
    --replace \
    --name dbTest \
    -p 3306:3306 \
    -v ./datadb:/var/lib/mysql:Z \
    --env MYSQL_DATABASE=bitacoraDB \
    --env MARIADB_USER=bitacoraU \
    --env MARIADB_PASSWORD=test-passwd \
    --env MARIADB_ROOT_PASSWORD=root-passwd \
    docker.io/library/mariadb

mysql -h 127.0.0.1 -u bitacoraU -p bitacoraDB
