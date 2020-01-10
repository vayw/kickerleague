all: build done

build:
		@echo "Building..."
			go build
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
