run-ex1:
	chmod +x repositories/base_python_code.py \
	&& go run cmd/ex1/main.go

run-ex2:
	chmod +x repositories/base_python_code.py \
	&& go run cmd/ex2/main.go

test-stacks:
	go test -v ./internal/stacks

test-queues:
	go test -v ./internal/queues

test-all: test-stacks test-queues