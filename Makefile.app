clean::
	-rm -rf .build/*


# Buidling binaries

CI_COMMIT := $(CIRCLE_SHA1)
CI_REPO := $(CIRCLE_PROJECT_REPONAME)
CI_NUM := $(CIRCLE_BUILD_NUM)

# backwards compatibility with Travis CI
ifndef CI_COMMIT
override CI_COMMIT := $(TRAVIS_COMMIT)
endif
ifndef CI_REPO
override CI_REPO := $(TRAVIS_REPO_SLUG)
endif
ifndef CI_NUM
override CI_NUM := $(TRAVIS_JOB_NUMBER)
endif

# local development fallbacks
ifndef CI_COMMIT
override CI_COMMIT := dev.$(shell whoami).$(shell git rev-parse --short HEAD)
endif
ifndef CI_REPO
override CI_REPO := $(PACKAGE)
endif
ifndef CI_NUM
override CI_NUM := $(shell whoami)
endif

LDIMPORT=github.com/remerge/go-service
LDFLAGS=-X $(LDIMPORT).CodeVersion=$(CI_COMMIT) -X $(LDIMPORT).CodeBuild=$(CI_REPO)\#$(CI_NUM)@$(shell date -u +%FT%TZ)
MAIN ?= main

# NOTE: logically pull out binaries from source tree
.build/$(PROJECT).%: $(GO_SOURCES) go.mod
	mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(basename $*) GOARCH=$(patsubst .%,%,$(suffix $*)) go build -o $@ -ldflags "$(LDFLAGS)" $(PACKAGE)/$(MAIN)

dist: .build/$(PROJECT).linux.amd64	## linux amd64 binary

.BIN_LOCAL = $(PROJECT).$(shell go env GOOS).$(shell go env GOARCH)

local: .build/$(.BIN_LOCAL)

# Deployment

release:
	git push origin master:production

deploy:
	/bin/bash -c 'cd ../chef.new && knife ssh roles:$(PROJECT) sudo chef-client'
