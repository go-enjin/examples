#!/usr/bin/make -f

SHELL = /bin/bash

SUB_DIRS = $(shell find be cmd -type d)

define _run_make =
	for dir in ${SUB_DIRS}; do \
		if [ -f "$${dir}/Makefile" ]; then \
			if egrep -q "^$(1)\:" "$${dir}/Makefile"; then \
				echo "# making $(1) in $${dir}"; \
				cd $${dir} > /dev/null; \
				make $(1); \
				cd - > /dev/null; \
			fi; \
		fi; \
	done
endef

clean:
	@$(call _run_make,clean)

dist-clean:
	@$(call _run_make,dist-clean)

build:
	@$(call _run_make,build)

tidy:
	@$(call _run_make,tidy)

local:
	@$(call _run_make,local)

unlocal:
	@$(call _run_make,unlocal)

be-update:
	@$(call _run_make,be-update)
