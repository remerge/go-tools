PROJECT := go-tools
PACKAGE := github.com/remerge/$(PROJECT)
REVIVELINTER_EXCLUDES = $(foreach p,$(wildcard **/*_fsm.go),-exclude $(p)) 
# state machines generate some unreachable code
VET_FLAGS = -unreachable=false
include mkf/Makefile.common
