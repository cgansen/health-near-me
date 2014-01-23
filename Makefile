ESHOST=http://localhost:9200
LINUX_AMD64_FLAGS=GOARCH="amd64" GOBIN="" GOCHAR="6" GOEXE="" GOHOSTARCH="amd64" GOHOSTOS="darwin" GOOS="linux" GORACE="" CC="gcc" GOGCCFLAGS="-g -O2 -fPIC -m64" CGO_ENABLED="0"

install:
	go install -v ./...

fmt:
	go fmt ./...

install-linux:
	$(LINUX_AMD64_FLAGS) go install -v ./...

deploy: install-linux
	rsync -avz ~/go/bin/linux_amd64/{api,loader} Makefile create_index.sh data tmpl ec2-user@api.healthnear.me:/var/www/healthnearme-api
	
deleteIndex:
	curl -XDELETE $(ESHOST)/health-near-me
	
createIndex:
	curl -XPOST $(ESHOST)/health-near-me
	
createMapping:
	curl -XPOST $(ESHOST)/health-near-me/health-provider/_mapping -d '{"health-provider":{"properties":{"address":{"type":"string"},"city":{"type":"string"},"hours_of_operation":{"type":"string"},"location":{"properties":{"human_address":{"type":"string"},"lat_lon":{"type":"geo_point"},"latitude":{"type":"string"},"longitude":{"type":"string"}}},"name":{"type":"string"},"phone":{"type":"string"},"provider_type":{"type":"long"},"state":{"type":"string"},"zip_code":{"type":"string"}}}}'
	
reset: deleteIndex createIndex createMapping
	@echo "\n\ndeleted and recreated index"
	
load:
	go run loader/main.go	
	
search:
	curl -XGET 'http://localhost:9200/health-near-me/health-provider/_search' -d '{ "query": { "filtered" : { "query": { "match_all": {} }, "filter" : { "geo_distance" : { "distance" : "1mi", "location.lat_lon" : { "lat" : 41.866592082671566, "lon" : -87.69819577390969 } } } } } }'

es:
	elasticsearch -f -D es.config=/usr/local/opt/elasticsearch/config/elasticsearch.yml