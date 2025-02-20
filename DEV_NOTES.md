# Dev Notes

Somes notes I keep for the development of the app

## TODO

### Final push for MVP

#### Bugs

- Some exercises in list being skipped when navigating up and down
- UI elements fallen off right of screen

#### Features

- Write a single exercise properly about installing and configuring nginx
  - Only testing nginx, not AWS. Do all AWS setup during provisioning
- Deploy an EC2 instance to AWS
  - Wrap around Terraform for this
  - Show a progress bar
  - Print terraform output to output console
- Run tests against deployed EC2 instance
  - Start all tests with failed styling in the ecercise description
  - Mark test as passed or failed by updating style in exercise description
  - Persist the state of each test in a user database of some description
    - Probably state stored in the exercises db
- Help updated based on deployment state of exercise
  - When undeployed, show how to deploy
  - When deployed, show how to undeploy

### Walking skeleton post-MVP

- Package for Debian
- Package for Ubuntu
- Package for Fedora
- Package for macOS (both Homebrew and a dmg installer file)
- Package for Windows (both chocolately and an exe installer file)
- Do all of these automatically in a pipeline
- Documentation in README
  - Installation instructions
  - Gifs showing basic features
    - Navigating the list and deploying an exercise
    - Running tests against an exercise
    - Tearing down exercse infrastructure

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
