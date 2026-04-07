all: funlang

funlang: cmd/*.go internal/*/*.go
	go build -o funlang cmd/main.go

.PHONY: clean

clean:
	$(RM) funlang

.PHONY: run

run: funlang
	./funlang