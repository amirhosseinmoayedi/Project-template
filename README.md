# Template for starting project

- dockerfile and docker-compose
- makfile
- live reloading

## used packages

- viper
- cobra
- logrus
- echo
- gorm
- golang-migrate

### add command to cobra

```bash
cobra-cli add command_name
```

### air package is used for live reloading

install:

```bash
go install github.com/cosmtrek/air@latest
```

change the setting in `.air.toml` file <br>
run it :

```bash
air -c .air.toml
```

`clean_on_exit ` option will delete the `tmp` directory on air exit, change it if you don't want to.


### create migration
create migration
```bash
migrate create -ext sql -dir internall/infrastructure/migration/ -seq init_mg
```
`-seq` for sequential, `UP` file for moving db to new state, `DOWN` for reverting <br>
apply migration
```bash
migrate -path internall/infrastructure/migration/ -database "postgresql://username:password@host:port/database_name?sslmode=disable" -verbose up
```
rollback
```bash
migrate -path internall/infrastructure/migration/ -database "postgresql://username:password@host:port/database_name?sslmode=disable" -verbose down
```