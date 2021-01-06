./bin/main:
	cd ./src/services/main && make build

run: ./bin/main
	./bin/main
