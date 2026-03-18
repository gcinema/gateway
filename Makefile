include .env
export

export PROJECT_ROOT := $(abspath ..)
export LOGGER_FOLDER := ${PROJECT_ROOT}/gateway/out/
export CONFIG_PATH := ${PROJECT_ROOT}/gateway/config/local.yaml

run:
	@go run ./cmd


