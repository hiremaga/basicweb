all: build

build:
	docker build -t hiremaga/basicweb --rm=true .

run:
	docker run -d -P hiremaga/basicweb

push:
	docker push hiremaga/basicweb
