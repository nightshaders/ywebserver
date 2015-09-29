

run: install	
	ywebserver web --serve-embedded-assets

embed: .FORCE
	rm -rf embedded/* .files/_site
	mkdir -p embedded .files/_site
	cp ../webground/dest/_site/* .files/_site
	go-bindata -nocompress -o embedded/assets.go -pkg embedded -prefix .files/_site  .files/_site

fmt: .FORCE
	gofmt -w */

install: .FORCE
	go install ywebserver

.FORCE:
