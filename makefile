

run: install	
	ywebserver web

embed: .FORCE
	rm -rf embedded/* .files/_site
	mkdir -p embedded .files/_site
	cp ~/Projects/Combo/webjournal/dest/_site/* .files/_site
	go-bindata -nocompress -o embedded/assets.go -pkg embedded -prefix .files/_site  .files/_site

install: .FORCE
	go install ywebserver

.FORCE:
