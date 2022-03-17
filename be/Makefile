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

SHELL = /bin/bash

.PHONY: all help clean build local unlocal be-update tidy

help:
	@echo "usage: make <clean|build>"

clean:
	@for p in $(shell ls -1 | grep -v Makefile); do \
		if [ -d "$${p}" ]; then \
			cd "$${p}" > /dev/null; \
			echo "# clean: $${p}"; \
			if [ -f Makefile ]; then \
				make clean || true; \
			fi; \
			cd - > /dev/null; \
		fi; \
	done

build:
	@for p in $(shell ls -1 | grep -v Makefile); do \
		if [ -d "$${p}" ]; then \
			cd "$${p}" > /dev/null; \
			echo "# build: $${p}"; \
			if [ -f Makefile ]; then \
				make build || true; \
			fi; \
			cd - > /dev/null; \
		fi; \
	done

local: export BE_LOCAL_PATH=$(shell realpath ../../be)
local:
	@for p in $(shell ls -1 | grep -v Makefile); do \
		if [ -d "$${p}" ]; then \
			cd "$${p}" > /dev/null; \
			echo "# local: $${p}"; \
			if [ -f go.mod ]; then \
				enjenv go-local || true; \
			fi; \
			cd - > /dev/null; \
		fi; \
	done

unlocal:
	@for p in $(shell ls -1 | grep -v Makefile); do \
		if [ -d "$${p}" ]; then \
			cd "$${p}" > /dev/null; \
			echo "# unlocal: $${p}"; \
			if [ -f go.mod ]; then \
				enjenv go-unlocal || true; \
			fi; \
			cd - > /dev/null; \
		fi; \
	done

be-update:
	@for p in $(shell ls -1 | grep -v Makefile); do \
		if [ -d "$${p}" ]; then \
			cd "$${p}" > /dev/null; \
			echo "# be-update: $${p}"; \
			if [ -f go.mod ]; then \
				GOPROXY=direct go get -u github.com/go-enjin/be || true; \
			fi; \
			cd - > /dev/null; \
		fi; \
	done

tidy:
	@for p in $(shell ls -1 | grep -v Makefile); do \
		if [ -d "$${p}" ]; then \
			cd "$${p}" > /dev/null; \
			echo "# tidy: $${p}"; \
			if [ -f go.mod ]; then \
				(go mod tidy -go=1.16 && go mod tidy -go=1.17) || true; \
			fi; \
			cd - > /dev/null; \
		fi; \
	done
