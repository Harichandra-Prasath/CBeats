run: build
	@./bin/CBeats $(ARGS)

build:
	@go build -o ./bin/CBeats
