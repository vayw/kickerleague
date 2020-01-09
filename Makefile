all: build done

build:
		@echo "Building..."
			@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
test:
		@echo "Running tests..."
			go test ./...

swagger:	
		@echo "Generating swagger files..."
			swag init -g api/api.go
			
clean:
		@echo "Cleanup..."
			@rm kickerleague
done:
		@echo "Done."
