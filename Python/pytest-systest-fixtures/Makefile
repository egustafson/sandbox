# Makefile for example
## -----------------------------------------------
## Vars

SVC_IMAGE = example-image
SVC_VER   = 0.1

## -----------------------------------------------

.PHONY: all test
all:
	echo "do nutin'"

test:
	pytest


## --  System Test  ------------------------------

.PHONY: kv-up kv-down systest
kv-up:
	docker-compose up -d

kv-down:
	docker-compose down

systest:
	pytest --systest
