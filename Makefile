USERNAME ?= your_username
PASSWORD ?= your_password

build:
	go build -o main ./main.go
	@echo "evennumberaverage application is built successfully"

run: build
	./main

test:
	go test -v ./...

podman-build:
	podman build -t quay.io/sanket/evennumberaverage -f Dockerfile
	@echo "evennumberaverage application image is created successfully"

podman-run:
	podman run -p 8080:8080 quay.io/sanket/evennumberaverage

podman-push:
	@podman login quay.io --username $(USERNAME) --password $(PASSWORD)
	podman push quay.io/sanket/evennumberaverage
	@echo "The image is pushed to quay.io"

clean:
	rm -f main
	@echo "All files have been cleaned successfully"

help:
	@echo "Available commands:"
	@echo "  make build        - To build the evennumberaverage application"
	@echo "  make run          - Run the main.go"
	@echo "  make test         - Run test cases"
	@echo "  make podman-build - Build the Podman image"
	@echo "  make podman-run   - Run the Podman container"
	@echo "  make podman-push  - Push the Podman image to quay.io"
	@echo "  make clean        - Clean up the build files"
	@echo "  make man          - Manual Page for evennumberaverage application  "

# man:
# 	@echo "Manual for evennumberaverage application:"
# 	@echo 
# 	@echo "  This Makefile will provide a set of commands to build, test, and run the evennumberaverage application."
# 	@echo 
# 	@echo "  To set your username and password for quay, use the following command:"
# 	@echo "     'make podman-push USERNAME=your_username PASSWORD=your_password'"
# 	@echo 
# 	@echo "  Example:"
# 	@echo "     'make podman-push USERNAME=rh-ee-sanket PASSWORD=123'"
# 	@echo 
# 	@echo "  To see all available commands, use:"
# 	@echo "      'make help'"
# 	@echo 