# Makefile

sqlc-init-default:
	sqlc init

sqlc-generate-default:
	sqlc generate

get-mocker:
	go install github.com/golang/mock/mockgen@v1.6.0

mock-user:
	scripts/mocks/mock-user.sh

mock-task:
	scripts/mocks/mock-task.sh

mock-all:
	@make \
	get-mocker \
	mock-user \
	mock-task;

test-all:
	go test ./...
