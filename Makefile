.PHONY: all

all: build run clean
# Создать бинарник
build:
	go build -o bin/main main.go

# Запустить бинарник
run:
	./bin/main

# Удалить бинарник
clean:
	rm -rf  bin
