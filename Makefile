PROJECT := go-tools
PACKAGE := github.com/remerge/$(PROJECT)

GOMETALINTER_OPTS=--vendor --vendored-linters --errors -D gotype

include Makefile.common
