package := ufolio
out_dir := ./.out

ifneq (,$(wildcard .env))
    include .env
endif

export GITHUB_TOKEN

build: clean
	go build -o $(out_dir)/$(package)
run: build
	 $(out_dir)/$(package)
clean:
	rm -f $(out_dir)/$(package)
docker-build:
	docker build -t $(package) .
docker-run: docker-build
	docker run -it --rm -p 8080:8080 $(package)
