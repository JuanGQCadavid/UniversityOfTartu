ex1: clean-root clean-ex1
	go run cmd/task1/main.go \
	&& mv *.png cmd/task1 \
	&& mv *.csv cmd/task1

clean-ex1:
	rm cmd/task1/*.png || true \
	&& rm cmd/task1/*.csv || true \

clean-root:
	rm *.png || true \
	&& rm *.csv || true


ex3:
	go run cmd/task3/main.go 