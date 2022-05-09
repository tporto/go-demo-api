# go-demo-api

#### run postgresql via docker
```sh
$ docker-compose up -d
```

Acesse o psql para criar as tabelas que encontra-se em sql/

```sh
$ docker exec -it devbook-db /bin/sh
$ psql -U postgres;
```