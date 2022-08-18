#!/bin/bash

VERSION=$1
URL=$1
JSON_FMT='{ "layers": { "tag": { "text": "version - %s" }, "changelog": { "text": "Full Release notes: %s" } } }'
JSON=$(printf "$JSON_FMT" "$VERSION" "$URL")

function gen_image (){
    STAGE_RSP=$(curl -X POST -H "Authorization: ${PLACID_APP_BEARER}" -H "Content-Type: application/json" -d "${JSON}" https://api.placid.app/api/rest/5ijokwc0v )
    IMAGE_ID=$(echo $STAGE_RSP | jq -r '.id')
    echo "Queued with image id: $IMAGE_ID"
    return 0
}

function download_image (){
    sleep 5
    GET_IMAGE_RSP=$(curl -X GET -H "Authorization: ${PLACID_APP_BEARER}" https://api.placid.app/api/rest/images/${IMAGE_ID})
    echo $GET_IMAGE_RSP
    IMAGE_URL=$(echo $GET_IMAGE_RSP | jq -r '.image_url')
    curl $IMAGE_URL -o tmp/cover.png
    echo "✅ Downloaded image to tmp/cover.png"
    return 0
}

echo "🔷 Placid Image Request:"
echo "\t${JSON}"
mkdir -p tmp
gen_image
download_image
