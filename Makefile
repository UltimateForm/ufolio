package := ufolio
out_dir := ./.out

build: clean
	go build -o $(out_dir)/$(package)
run: build
	$(out_dir)/$(package)
clean:
	rm -f $(out_dir)/$(package)
docker-build:
	docker build -t $(package) .
docker-run: docker-build
	docker run -it --rm $(package)
