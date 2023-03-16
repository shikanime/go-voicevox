.PHONY: all
all: gen format

.PHONY: format
format:
	go fmt ./...

.PHONY: gen
gen:
	openapi-generator-cli generate \
		-i openapi.json \
		-g go \
		--git-host github.com \
		--git-repo-id go-voicevox \
		--git-user-id infinity-blackhole \
		--package-name voicevox \
		-c config.yaml