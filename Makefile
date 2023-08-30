check: testall build386 buildlinuxarm # Perform basic checks and tests

testall: testapi testconfig testdb testdocker testschema ## Run all tests

testapi: ## Test the api package
		go test ./api

testconfig: ## Test the config package
		go test ./config

testdb: ## Test the db package
		go test ./db

testdocker: ## Test the docker package
		go test ./docker

testschema: ## Test the schema package
		go test ./schema

testtypes: ## Test the schema package
		go test ./types

build386: ## Build for linux/386
		GOOS=linux GOARCH=386 go install .

buildlinuxarm: ## Build for linux/arm
		GOOS=linux GOARCH=arm go install .

tidy:	## Run go mod tidy
		go mod tidy