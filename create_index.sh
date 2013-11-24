#!/bin/bash

ESHOST="http://localhost:9200"
INDEX_NAME="health-near-me"
INDEX_TYPE="health-provider"

# nuke existing
echo ''
echo 'deleting existing index'
curl -XDELETE "$ESHOST/$INDEX_NAME"

# create
echo ''
echo "creating new index"
curl -XPOST "$ESHOST/$INDEX_NAME"

# create mapping
echo ''
echo "creating mapping"
curl -XPOST "$ESHOST/$INDEX_NAME/$INDEX_TYPE/_mapping" -d '{"health-provider":{"properties":{"address":{"type":"string"},"city":{"type":"string"},"hours_of_operation":{"type":"string"},"location":{"properties":{"human_address":{"type":"string"},"lat_lon":{"type":"geo_point"},"latitude":{"type":"string"},"longitude":{"type":"string"}}},"name":{"type":"string"},"phone":{"type":"string"},"provider_type":{"type":"long"},"state":{"type":"string"},"zip_code":{"type":"string"}}}}'

echo "done"