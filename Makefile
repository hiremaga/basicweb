all: build

build:
	GOOS=linux go build
	docker build -t hiremaga/basicweb --rm=true .

run:
	docker run -d -P hiremaga/basicweb

push:
	docker push hiremaga/basicweb
