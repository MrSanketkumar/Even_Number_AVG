USERNAME ?= your_username
PASSWORD ?= your_password
# CLUSTERUSERNAME ?= your_username
# CLUSTERPASSWORD ?= your_password 
PROJECT_NAME ?= your_project_name
APPLICATION_NAME ?= your_application_name
API_URL = https://api.cluster01.lab.psi.pnq2.redhat.com:6443

help:
	@echo "Available commands:"
	@echo "  make man          - Manual Page for evennumberaverage application  "
	@echo "  make build        - To build the evennumberaverage application"
	@echo "  make run          - Run the main.go"
	@echo "  make test         - Run test cases"
	@echo "  make podman-build - Build the Podman image"
	@echo "  make podman-run   - Run the Podman container"
	@echo "  make podman-push  - Push the Podman image to quay.io"
	@echo "  make clean        - Clean up the build files"
	@echo "  make all          - Used to test and run the application"
	@echo "  make podman-all   - Used to build,push and run the container"

.PHONY: all clean 

all: clean test run 

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

podman-run: podman-build
	podman run -p 8080:8080 quay.io/sanket/evennumberaverage

podman-push:
	podman login quay.io -u sanket -p ${{secrets.PASSWORD_QUAY}}
	podman push quay.io/sanket/evennumberaverage
	@echo "The image is pushed to quay.io"

clean:
	rm -f main
	@echo "All files have been cleaned successfully"

podman-all: podman-build podman-push podman-run

# oc-deploy:
# ifdef CLUSTERUSERNAME
#     oc login https://api.cluster01.lab.psi.pnq2.redhat.com:6443 -u $(CLUSTERUSERNAME) -p $(CLUSTERPASSWORD) --insecure-skip-tls-verify
# else
#     @echo "Error: CLUSTERUSERNAME is not defined. Please set CLUSTERUSERNAME variable before running this command."
#     @exit 1
# endif

#     # oc new-project ${PROJECT_NAME}

#     oc new-app . --name ${APPLICATION_NAME} --strategy=docker

#     # oc waiti --for=condition=complete --timeout=600s  build/${APPLICATION_NAME}

#     sleep 5

#     oc start-build ${APPLICATION_NAME} --from-dir=./ --follow=true --wait=true

#     oc expose svc ${APPLICATION_NAME}

#     oc get routes.route.openshift.io 
# oc-delete:
# 	@echo "Switching to project: ${APPLICATION_NAME}"
# 	oc project ${APPLICATION_NAME}
# 	@echo "Deleting all resources in project ${APPLICATION_NAME}..."
# 	oc delete all --all
# 	@echo "All resources deleted successfully."

man:
	@echo "Manual for evennumberaverage application:"
	@echo 
	@echo "  This Makefile will provide a set of commands to build, test, and run the evennumberaverage application."
	@echo 
	@echo "  To set your username and password for quay, use the following command:"
	@echo "     'make podman-push USERNAME=your_username PASSWORD=your_password'"
	@echo 
	@echo "  Example:"
	@echo "     'make podman-push USERNAME=rh-ee-sanket PASSWORD=123'"
	@echo 
	@echo "  To see all available commands, use:"
	@echo "      'make help'"
	@echo 


s2i-build:
	wget https://github.com/openshift/source-to-image/releases/download/v1.1.14/source-to-image-v1.1.14-874754de-linux-386.tar.gz
	tar -xvf source-to-image-v1.1.14-874754de-linux-386.tar.gz
	cp s2i /usr/local/bin
	@sudo s2i build . quay.io/sanket/my-builder-image quay.io/sanket/my-even-application
	@echo " s2i builer image is created successfully"

s2i-push: s2i-build
	@echo " quay login successful"
	@sudo docker push quay.io/sanket/my-even-application
	@echo " my-even-application image is pushed to quay.io"

changelog:
	@PREVIOUS_TAG=$$(git describe --abbrev=0 --tags $$(git rev-list --tags --skip=1 --max-count=1)); \
	git log --oneline --decorate --first-parent $$PREVIOUS_TAG..HEAD; \
	
cleanlog:
	rm -f changelog