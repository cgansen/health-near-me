ESHOST=http://localhost:9200
	
deleteIndex:
	curl -XDELETE $(ESHOST)/health-near-me
	
createIndex:
	curl -XPOST $(ESHOST)/health-near-me
	
createMapping:
	curl -XPOST $(ESHOST)/health-near-me/health-provider/_mapping -d '{"health-provider":{"properties":{"address":{"type":"string"},"city":{"type":"string"},"hours_of_operation":{"type":"string"},"location":{"properties":{"human_address":{"type":"string"},"lat_lon":{"type":"geo_point"},"latitude":{"type":"string"},"longitude":{"type":"string"}}},"name":{"type":"string"},"phone":{"type":"string"},"provider_type":{"type":"long"},"state":{"type":"string"},"zip_code":{"type":"string"}}}}'
	
reset: deleteIndex createIndex createMapping
	@echo "\n\ndeleted and recreated index"
	
load:
	go run loader.go	
	
search:
	curl -XGET 'http://localhost:9200/health-near-me/health-provider/_search' -d '{ "query": { "filtered" : { "query": { "match_all": {} }, "filter" : { "geo_distance" : { "distance" : "1mi", "location.lat_lon" : { "lat" : 41.866592082671566, "lon" : -87.69819577390969 } } } } } }'
