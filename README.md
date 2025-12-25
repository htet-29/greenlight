## Creating migration file
```sh
$ migrate create -seq -ext=.sql -dir=./migrations create_movies_table
```

```sh
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
```

```sh
$ migrate -path=./migrations -database=$EXAMPLE_DSN goto 1
```