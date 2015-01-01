#
#
_bootstrap_Makefile := $(wildcard [M]akefile)
_bootstrap_config-status := $(wildcard config.status)

ALL_RECURSIVE_TARGETS=

ifneq ($(_bootstrap_Makefile),)
  include Makefile
else
  ifneq ($(_bt_config-status),)
    $(srcdir)/config.status
    $(MAKE) $(AM_MAKEFLAGS) configure
  endif

srcdir= .

.PHONY: basic_build
basic_build: Makefile
	@$(MAKE) $(AM_MAKEFLAGS)

Makefile: configure
	@$(srcdir)/configure

configure:
	@autoreconf -fi

endif

ALL_RECURSIVE_TARGETS+= $(AM_RECURSIVE_TARGETS)

ifneq ($(word 2, $(MAKECMDGOALS)), )
ifneq ($(filter $(ALL_RECURSIVE_TARGETS), $(MAKECMDGOALS)), )
.NOTPARALLEL:
endif
endif
