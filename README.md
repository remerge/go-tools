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
