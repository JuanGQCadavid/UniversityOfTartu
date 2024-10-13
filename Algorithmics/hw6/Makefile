ex1:
	rm *.png \
	&& rm *.csv \
	&& rm cmd/task1/*.png || true \
	&& rm cmd/task1/*.csv || true \
	&& go run cmd/task1/main.go \
	&& mv *.png cmd/task1 \
	&& mv *.csv cmd/task1
