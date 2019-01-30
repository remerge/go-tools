# go-makefile

Common Makefile for remerge Go projects. Various targets are provided:

*TBD: update while working on: https://trello.com/c/vROA2Va8/115-improve-and-share-makefilecommon-revivetoml-between-go-services*

## Setup

### If the makefile is not yet part of the repository

Add this repository as a subtree to the go project repository.

```
git remote add makefile https://github.com/remerge/go-makefile.git
git subtree add --squash --prefix mkf/ makefile master
```
Afterwards `mkf/Makefile.common` can be included in the parent project.

### If the makefile is already available

`make update-makefile`

## Example

```
PROJECT := go-tools
PACKAGE := github.com/remerge/$(PROJECT)
REVIVELINTER_EXCLUDES = $(foreach p,$(wildcard **/*_fsm.go),-exclude $(p))

include mkf/Makefile.common
```

## Modules

### Vendor

From Go 1.11 modules used instead vendor. To maintain vendor directory use

```
GO111MODULES=on go mod vendor
```

## Targets

### lint

This target invokes Go format check, vet and revive linter. Revive will be
installed automatically if it not present on host. To override revive config 
put `revive.toml` in root of build tree. Use `REVIVELINTER_EXCLUDES` to add 
excludes.

```
REVIVELINTER_EXCLUDES = $(foreach p,$(wildcard **/*_fsm.go),-exclude $(p))
```

## Isolated test deployment (diversion)

Sometimes run application on isolated instance is the best way to test it. *This is
not safe regular deployment!* You must warn your colleagues.

### Requirements

1. SSH access and sudo rights on target machine
1. Minimal knowledge about SystemD
1. Go with crossplatform build on developer machine

### Setup and teardown

For each operation you need to define the `DIVERT_SSH` environment variable.

```shell
$ DIVERT_SSH=user@app.machine make ...
$ make ... DIVERT_SSH=user@app.machine
```

Alternatively `DIVERT_SSH` can be defined globally:

```shell
$ export DIVERT_SSH=user@app.machine
```

> Diversion status may be checked by `.CHECK-divert-on` and
  `.CHECK-divert-on` targets.

Now you can run the `divert-setup` target. This target will prepare the basic setup and
stop `chef-client`.

To revert diversion environment and start `chef-client` use `divert-teardown` target.

### Journal

Use `divert-journal` to follow application log. This target is not depends on
diverted environment.

### Deploy

Use `divert-do` target to deploy dev version.