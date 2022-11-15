GO_FILES=`go list ./... | grep -v -E "mock|store|test|fake|cmd"`



lint: ## Lint Golang files
	@golint  ${GO_FILES}


test:
	@go test $(GO_FILES) -coverprofile .cover.txt
	@go tool cover -func .cover.txt

clean: ## Remove previous build
	@rm .cover.txt