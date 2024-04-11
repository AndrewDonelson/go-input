# Set the application name and version
APP_NAME := goinput
APP_VERSION := 1.0.0

# Set the build flags
BUILD_FLAGS := -ldflags="-s -w -X main.version=$(APP_VERSION)"

all:
	rm -f $(APP_NAME)-*
	@echo "Building $(APP_NAME)"
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(APP_NAME)-linux-amd64