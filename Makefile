package := ufolio
out_dir := ./.out

ifneq (,$(wildcard .env))
    include .env
endif

export GITHUB_TOKEN
export TURNSTILE_SECRET_KEY=1x0000000000000000000000000000000AA
# export X_EDGE_SIGNATURE
export DEV=1
export TURNSTILE_SITE_KEY=1x00000000000000000000AA

build: clean
	go build -o $(out_dir)/$(package)
run: build
	 $(out_dir)/$(package)
clean:
	rm -f $(out_dir)/$(package)
docker-build:
	docker build -t $(package) .
docker-run: docker-build
	docker run --env-file .env -it --rm -p 8080:8080 $(package)
test:
	go test ./...
