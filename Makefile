DISTFILE := mdp

.PHONY: dist

run:
	@go run main.go

dist:
	# Build for darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -o ./dist/$(DISTFILE)
	cd dist && tar cf $(DISTFILE)-osx.tar.gz $(DISTFILE)
	# Build for linux-amd64
	GOOS=linux GOARCH=amd64 go build -o ./dist/$(DISTFILE)
	cd dist && tar cf $(DISTFILE)-linux-amd64.tar.gz $(DISTFILE)
	# Build for linux-386
	GOOS=linux GOARCH=386 go build -o ./dist/$(DISTFILE)
	cd dist && tar cf $(DISTFILE)-linux-i386.tar.gz $(DISTFILE)
