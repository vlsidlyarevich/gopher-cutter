
install: go-get



go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@pwd
	go get -d ./...