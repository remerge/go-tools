PROJECT := go-tools
PACKAGE := github.com/remerge/$(PROJECT)
REVIVELINTER_EXCLUDES = $(foreach p,$(wildcard **/*_fsm.go),-exclude $(p)) 

include mkf/Makefile.common
