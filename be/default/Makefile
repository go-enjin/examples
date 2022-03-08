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

BE_LOCAL_PATH ?= ../../../be

DEBUG ?= false

APP_NAME    ?= be-default
APP_SUMMARY ?= Default Enjin
APP_VERSION ?= $(shell enjenv git-tag --untagged v0.0.1)
REL_VERSION ?= $(shell enjenv rel-ver)
ENV_PREFIX  ?= BE

BUILD_TAGS =

RELEASE = ""

.PHONY: all help clean build dev release run

help:
	@echo "usage: make <help|clean|build|dev|release|run>"

clean:
	@if [ -f "${APP_NAME}" ]; then rm -fv "${APP_NAME}"; fi

build: LABEL=$(shell if [ "${RELEASE}" = "true" ]; then echo "Building release"; else echo "Building"; fi)
build: OPTION=$(shell if [ "${RELEASE}" = "true" ]; then echo "--optimize"; fi)
build:
	@echo "${LABEL}: ${APP_VERSION}, ${REL_VERSION}"
	@${CMD} enjenv golang build \
			--be-app-name "${APP_NAME}" \
			--be-summary "${APP_SUMMARY}" \
			--be-version "${APP_VERSION}" \
			--be-release "${REL_VERSION}" \
			--be-env-prefix "${ENV_PREFIX}" \
			${OPTION} -- -v -tags=${BUILD_TAGS}

release: RELEASE="true"
release: build

run:
	@if [ -x "${APP_NAME}" ]; then \
		echo "# running ${APP_NAME}"; \
		${ENV_PREFIX}_DEBUG=${DEBUG} ./${APP_NAME}; \
	else \
		echo "# ${APP_NAME} not found"; \
	fi

dev: DEBUG=true
dev: run

tidy:
	@go mod tidy -go=1.16 && go mod tidy -go=1.17
