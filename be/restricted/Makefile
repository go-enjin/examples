#!/usr/bin/make -f

# Copyright (c) 2022  The Go-Enjin Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#: uncomment to echo instead of execute
#CMD=echo
SHELL = /bin/bash

APP_NAME    ?= be-restricted
APP_SUMMARY ?= Restricted Enjin
ENV_PREFIX  ?= RST

LISTEN        ?=
PORT          ?= 3334
PREFIX        ?= ${USER}
DEBUG         ?= false
STRICT        ?= false
DOMAIN        ?=
DENY_DURATION ?= 86400

BUILD_TAGS = basic_auth,locals
EXTRA_PKGS =

GOLANG ?= 1.17.7
NODEJS ?=

RELEASE_BUILD = false

GO_ENJIN_PKG = github.com/go-enjin/be

.PHONY: help clean build dev release run enjenv golang

help:
	@echo "${APP_NAME} - ${APP_SUMMARY}"
	@echo
	@echo "This is a Go-Enjin project which depends upon enjenv to"
	@echo "manage the Golang and Nodejs SDKs and build this Enjin."
	@echo
	@echo "usage: make <clean|dist-clean|build|dev|release|run>"
	@echo
	@echo "cleanup targets:"
	@echo "  clean       remove built binary"
	@echo "  dist-clean  clean and purge local .bin ${ENJENV_DIR}"
	@echo
	@echo "compile targets:"
	@echo "  build       build a debug-able version of ${APP_NAME}"
	@echo "  release     build a release version of ${APP_NAME}"
	@echo
	@echo "runtime targets:"
	@echo "  dev         set for DEBUG mode and run ./${APP_NAME}"
	@echo "  run         execute ./${APP_NAME}"
	@echo
	@echo "dependency targets:"
	@echo "  enjenv      install enjenv if not already present"
	@echo "  golang      enjenv init golang if not already present"
	@echo "  nodejs      enjenv init nodejs if not already present"
	@echo
	@echo "helper targets:"
	@echo "  help        this screen of output"
	@echo "  tidy        run go mod tidy -go=1.16 && -go=1.17"
	@echo "  local       add go mod -replace for ${GO_ENJIN_PKG}"
	@echo "  unlocal     go mod -dropreplace for ${GO_ENJIN_PKG}"
	@echo "  be-update   go clean and get -u ${GO_ENJIN_PKG}"

define _be_local_path =
$(shell \
	if [ "${BE_LOCAL_PATH}" != "" -a -d "${BE_LOCAL_PATH}" ]; then \
		echo "${BE_LOCAL_PATH}"; \
	elif [ "${GOPATH}" != "" ]; then \
		if [ -d "${GOPATH}/src/${GO_ENJIN_PKG}" ]; then \
			echo "${GOPATH}/src/${GO_ENJIN_PKG}"; \
		fi; \
	fi)
endef

define _be_version =
$(shell ${ENJENV_EXE} git-tag --untagged v0.0.0)
endef

define _be_release =
$(shell ${ENJENV_EXE} rel-ver)
endef

define _enjenv_bin =
$(shell \
	if [ "${ENJENV_BIN}" != "" -a -x "${ENJENV_BIN}" ]; then \
		echo "${ENJENV_BIN}"; \
	else \
		if [ "$(1)" != "" ]; then \
			echo "$(1)"; \
		elif [ -d .bin -a -x .bin/enjenv ]; then \
			echo "${PWD}/.bin/enjenv"; \
		fi; \
	fi)
endef

define _enjenv_path =
$(shell \
	if [ -x "${ENJENV_EXE}" ]; then \
		${ENJENV_EXE}; \
	elif [ -d "./${ENJENV_DIR}" ]; then \
		echo "${PWD}/${ENJENV_DIR}"; \
	fi)
endef

define _build_tags =
$(shell if [ "${BUILD_TAGS}" != "" ]; then echo "-tags ${BUILD_TAGS}"; fi)
endef

define _build_label =
$(shell \
	if [ "${RELEASE_BUILD}" == "true" ]; then \
		echo "Building release"; \
	else \
		echo "Building"; \
	fi)
endef

define _build_args =
$(shell \
	if [ "${RELEASE_BUILD}" == "true" ]; then \
		echo "--optimize"; \
	fi)
endef

define _env_build_vars =
BE_APP_NAME="${APP_NAME}" \
BE_SUMMARY="${APP_SUMMARY}" \
BE_ENV_PREFIX="${ENV_PREFIX}" \
BE_VERSION="${VERSION}" \
BE_RELEASE="${RELEASE}"
endef

define _env_run_vars =
${ENV_PREFIX}_RESTRICT_WITH_BASIC_ENV="true" \
${ENV_PREFIX}_RESTRICT_ALL_DATA="true" \
${ENV_PREFIX}_RESTRICT_BASIC_USER_ZEN='$$1$$dlPL2MqE$$oQmn16q49SqdmhenQuNgs1' \
${ENV_PREFIX}_RESTRICT_BASIC_USER_ONE_MORE='$$1$$dlPL2MqE$$oQmn16q49SqdmhenQuNgs1' \
${ENV_PREFIX}_RESTRICT_BASIC_GROUP_ONLY_ONES='zen,one-more' \
${ENV_PREFIX}_RESTRICT_BASIC_GROUP_ONES_TWOS='zen' \
${ENV_PREFIX}_DEBUG="${DEBUG}" \
${ENV_PREFIX}_LISTEN="${LISTEN}" \
${ENV_PREFIX}_PORT="${PORT}" \
${ENV_PREFIX}_PREFIX="${PREFIX}" \
${ENV_PREFIX}_DEBUG="${DEBUG}" \
${ENV_PREFIX}_STRICT="${STRICT}" \
${ENV_PREFIX}_DOMAIN="${DOMAIN}" \
${ENV_PREFIX}_DENY_DURATION="${DENY_DURATION}"
endef

BE_PATH ?= $(call _be_local_path)

ENJENV_EXE ?= $(call _enjenv_bin,$(shell which enjenv))
ENJENV_URL ?= https://github.com/go-enjin/enjenv-heroku-buildpack/raw/trunk/bin/enjenv
ENJENV_PKG ?= github.com/go-enjin/enjenv/cmd/enjenv@latest
ENJENV_DIR_NAME ?= .enjenv
ENJENV_DIR ?= ${ENJENV_DIR_NAME}
ENJENV_PATH ?= $(call _enjenv_path)

VERSION ?= $(call _be_version)
RELEASE ?= $(call _be_release)

enjenv: gobin=$(shell which go)
enjenv:
	@if [ "${ENJENV_EXE}" = "" ]; then \
		if [ "${gobin}" = "" ]; then \
			echo "# downloading enjenv..."; \
			wget -q -c ${ENJENV_URL}; \
			chmod +x ./enjenv; \
			echo "# installing enjenv..."; \
			if [ "${GOPATH}" = "" ]; then \
				mkdir -v .bin; \
				mv -v ./enjenv ./.bin/; \
			else \
				mv -v ./enjenv ${GOPATH}/bin/enjenv; \
			fi; \
		else \
			echo "# go install enjenv..."; \
			go install ${ENJENV_PKG}; \
		fi; \
	fi

golang: enjenv
	@if [ "${ENJENV_PATH}" != "" ]; then \
		if [ ! -d "${ENJENV_PATH}/golang" ]; then \
			if [ "${GOLANG}" != "" ]; then \
				${CMD} ${ENJENV_EXE} golang init --golang "${GOLANG}"; \
				${CMD} ${ENJENV_EXE} write-scripts; \
			else \
				${CMD} ${ENJENV_EXE} golang init; \
				${CMD} ${ENJENV_EXE} write-scripts; \
			fi; \
		fi; \
		if [ "${NODEJS}" != "" ]; then \
			if [ ! -d "${ENJENV_PATH}/nodejs" ]; then \
				${CMD} ${ENJENV_EXE} nodejs init --nodejs "${NODEJS}"; \
				${CMD} ${ENJENV_EXE} write-scripts; \
			fi; \
		fi; \
	else \
		echo "# missing enjenv path"; \
		false; \
	fi

tidy: golang
	@echo "# go mod tidy -go=1.16 && go mod tidy -go=1.17"
	@source "${ENJENV_PATH}/activate" \
		&& ${CMD} go mod tidy -go=1.16 \
		&& ${CMD} go mod tidy -go=1.17

local: enjenv
	@if [ "${BE_PATH}" = "" ]; then \
		echo "missing BE_PATH"; \
		false; \
	fi
	@echo "# localizing ${GO_ENJIN_PKG}"
	@${CMD} ${ENJENV_EXE} go-local "${BE_PATH}"

unlocal: enjenv
	@echo "# restoring ${GO_ENJIN_PKG}"
	@${CMD} ${ENJENV_EXE} go-unlocal

be-update: golang
	@echo "# go get -u ${GO_ENJIN_PKG} ${EXTRA_PKGS}"
	@source "${ENJENV_PATH}/activate" \
		&& ${CMD} GOPROXY=direct go get -u \
			$(call _build_tags) \
			${GO_ENJIN_PKG} \
			${EXTRA_PKGS}

clean:
	@if [ -f "${APP_NAME}" ]; then rm -fv "${APP_NAME}"; fi

dist-clean: clean
	@if [ -d "./${ENJENV_DIR}" ]; then \
		ENJENV_PATH="${PWD}/${ENJENV_DIR}" ${ENJENV_EXE} clean --force; \
	fi
	@if [ -d .bin ]; then \
		rm -rfv .bin; \
	fi

build: golang
	@echo "$(call _build_label): ${VERSION}, ${RELEASE}"
	@${CMD} \
			$(call _env_build_vars) \
			${ENJENV_EXE} golang build \
				$(call _build_opts) \
				-- -v $(call _build_tags)

release: RELEASE_BUILD="true"
release: build

run:
	@if [ ! -x "${APP_NAME}" ]; then \
		echo "# ${APP_NAME} not found or not executable"; \
	fi
	@echo "# running ${APP_NAME}"
	@${CMD} $(call _env_run_vars) ./${APP_NAME}

dev: DEBUG=true
dev: run
