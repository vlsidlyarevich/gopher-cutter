
install: go-get go-build

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@pwd
	go get -d ./...

go-build:
	@echo "  >  Building the executable..."
	@pwd
	go build ./cmd/gopher-cutter/main.go