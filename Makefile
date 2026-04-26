APP_NAME := kbot
IMAGE_TAG := quay.io/viktorshlapak/kbot:v1.0.0

.PHONY: linux arm macos windows image clean

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(APP_NAME)-linux-amd64 .

arm:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bin/$(APP_NAME)-linux-arm64 .

macos:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o bin/$(APP_NAME)-darwin-arm64 .

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(APP_NAME)-windows-amd64.exe .

image:
	docker build -t $(IMAGE_TAG) .

clean:
	rm -rf bin
	docker rmi $(IMAGE_TAG)
