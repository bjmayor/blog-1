deploy: clean build_app archive


clean:
	rm -rf deploy


build_app:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=prod -ldflags "-s -w" -o blog main.go
	cd vue && npm run build

archive:
	mkdir deploy
	cp blog deploy/blog
	cp -r views deploy/
	cp -r static deploy/
	cp -r dist deploy/
	tar zcvf blog.tar.gz deploy
	cp blog.tar.gz ~/Documents/blog
