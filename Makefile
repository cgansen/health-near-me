LINUX_AMD64_FLAGS=GOARCH="amd64" GOBIN="" GOCHAR="6" GOEXE="" GOHOSTARCH="amd64" GOHOSTOS="darwin" GOOS="linux" GORACE="" CC="gcc" GOGCCFLAGS="-g -O2 -fPIC -m64" CGO_ENABLED="0"

install:
	go install -v ./...

fmt:
	go fmt ./...

install-linux:
	$(LINUX_AMD64_FLAGS) go install -v ./...

deploy: install-linux
	rsync -e 'ssh -l ec2-user -i /Users/cgansen/.ssh/cgansen-aws.pem' -avz ~/go/bin/linux_amd64/{api,loader} create_index.sh data tmpl ec2-user@healthnearme:/var/www/healthnearme-api
