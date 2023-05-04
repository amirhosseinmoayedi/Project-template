# Template for starting project

- dockerfile and docker-compose
- makfile
- live reloading

## used packages

- viper
- cobra
- logrus
- echo

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
