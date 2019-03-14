.PHONY: test

fmt:
	@find . -iname "*.go" -not -path "./vendor/**" | xargs gofmt -s -w

imports:
	@goimports -w $$(find . -iname "*.go" -not -path "./vendor/**")

fix: ## Fix code lint issues
	@$(MAKE) fmt
	@$(MAKE) imports

metalinter:
	@gometalinter --vendor --disable-all --enable=gofmt --enable=goimports ./...

test: ## Run project tests using `go test`
	@go test ./...

dep: ## Install project dependencies using `dep`
	@go get github.com/golang/dep/cmd/dep
	@dep ensure
	@find vendor -type l -delete
