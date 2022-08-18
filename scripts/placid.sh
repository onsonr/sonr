#!/bin/bash

VERSION=$1
URL=$2
OUT_PATH=$3
JSON_FMT='{ "layers": { "tag": { "text": "version - %s" }, "changelog": { "text": "Full Release notes: %s" } } }'
JSON=$(printf "$JSON_FMT" "$VERSION" "$URL")

echo "🔷 Placid Image Request:\n\t${JSON}"
echo "Sending Placid Image Request..."
mkdir -p tmp
STAGE_RSP=$(curl -X POST -H "Authorization: ${PLACID_APP_BEARER}" -H "Content-Type: application/json" -d "${JSON}" https://api.placid.app/api/rest/5ijokwc0v )
IMAGE_ID=$(echo $STAGE_RSP | jq -r '.id')
echo "Queued with image id: $IMAGE_ID"

echo "Waiting for image to be ready before downloading..."
sleep 5
GET_IMAGE_RSP=$(curl -X GET -H "Authorization: ${PLACID_APP_BEARER}" https://api.placid.app/api/rest/images/${IMAGE_ID})
echo $GET_IMAGE_RSP
IMAGE_URL=$(echo $GET_IMAGE_RSP | jq -r '.image_url')
curl $IMAGE_URL -o $OUT_PATH
echo "✅ Downloaded image to $OUT_PATH"
