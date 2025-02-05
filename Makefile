build:
	@echo "Building dsh module..."
	cd modules/dsh && go build -o ./cmd/dsh ./cmd/dsh

	@echo "Building Duckshell..."
	cargo build

	@echo "Build complete!"

clean:
	@echo "Cleaning dsh module..."
	cd modules/dsh && go clean ./cmd/dsh/dsh

	@echo "Cleaning Duckshell..."
	cargo clean

	@echo "Removing compiled files..."
	rm -rf target

	@echo "Clean complete!"

run:
	@echo "Running Duckshell..."
	./target/debug/duckshell

.PHONY: build clean run