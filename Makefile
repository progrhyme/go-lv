.PHONY: test version

VERSION := $(shell cat version.txt)

test:
	go test -v ./...

version:
	git commit -m $(VERSION)
	git tag -a v$(VERSION) -m $(VERSION)
	git push origin v$(VERSION)
	git push origin master
