export MAIN_BRANCH ?= main

.DEFAULT_GOAL := help
.PHONY: test build clean test/coverage release/prepare release/tag .check_bump_type .check_git_clean help

GIT_BRANCH := $(shell git symbolic-ref --short HEAD)
WORKTREE_CLEAN := $(shell git status --porcelain 1>/dev/null 2>&1; echo $$?)
SCRIPTS_DIR := $(CURDIR)/scripts

curVersion := $$(sed -n -E 's/^const SDKVersion = "v?(.+)"$/\1/p' connect/version.go)

test:	## Run test suite
	go test -v ./...

test/coverage:	## Run test suite with coverage report
	go test -v ./... -cover

build:	## Build project
	go build ./...


## Release functions =====================

release/prepare: .check_git_clean	## Bumps version and creates release branch (call with 'release/prepare version=<new_version_number>')

	@test $(version) || (echo "[ERROR] version argument not set."; exit 1)
	@git fetch --quiet origin $(MAIN_BRANCH)

	@sed -i.tmp -E 's/^(const SDKVersion) = "v?(.+)"$$/\1 = "$(version)"/' connect/version.go
	@NEW_VERSION=$(version) $(SCRIPTS_DIR)/prepare-release.sh

	@rm -f connect/version.go.tmp


release/tag: .check_git_clean	## Creates git tag using version from package.json
	@git pull --ff-only
	@echo "Applying tag 'v$(curVersion)' to HEAD..."
	@git tag --sign "v$(curVersion)" -m "Release v$(curVersion)"
	@echo "[OK] Success!"
	@echo "Remember to call 'git push --tags' to persist the tag."

## Helper functions =====================

.check_git_clean:
ifneq ($(GIT_BRANCH), $(MAIN_BRANCH))
	@echo "[ERROR] Please checkout default branch '$(MAIN_BRANCH)' and re-run this command."; exit 1;
endif
ifneq ($(WORKTREE_CLEAN), 0)
	@echo "[ERROR] Uncommitted changes found in worktree. Address them and try again."; exit 1;
endif
