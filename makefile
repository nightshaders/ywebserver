

run: embed install
	ywebserver web --serve-embedded-assets

embed: .FORCE
	rm -rf .files/_site
	mkdir -p .files/_site
	cp ../webground/dest/_site/* .files/_site
	go-bindata -nocompress -o embedded.go -pkg main -prefix .files/_site  .files/_site

fmt: .FORCE
	gofmt -w */

install: .FORCE
	go install github.com/nightshaders/ywebserver

.FORCE:
