# Dev Notes

Somes notes I keep for the development of the app

## TODO

- Two existing exercises written properly, not with lorem ipsum
- Exercise state updated by pressing enter (Deployed/Not Deployed)
  - Presisted between sessions
- Help updated based on deployment state of exercise

## Some variables and things

```sh
APP_NAME=iep # This will change once I come up with a better name
```

## Exercise migrations

Right now the migrations are run at startup, which won't work when it's
downloaded from a package managed. The database needs to be migrated ahead of
time and packaged with the application.

- Decouple the migrations from startup. Turn it into a make rule instead.
- Document migration mechanism
- Get rid of the migration registry and figure out how to list the migrations
  dynamically

## XDG Variable Reference

Just storing these so I don't have to keep looking them up

```sh
XDG_DATA_HOME = $HOME/.local/share
XDG_STATE_HOME = $HOME/.local/state
XDG_CONFIG_HOME = $HOME/.config
```

## Files

### exercises.db

```sh
${XDG_DATA_HOME}/${APP_NAME}/exercises.db
```
