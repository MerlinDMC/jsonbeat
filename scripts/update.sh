#!/bin/bash

BEAT_NAME=${BEAT_NAME:-jsonbeat}
BEAT_INDEX_PREFIX=${BEAT_INDEX_PREFIX:-${BEAT_NAME}}


cat _meta/beat.yml | sed -e "s/beatname/${BEAT_NAME}/g;s/beat-index-prefix/${BEAT_INDEX_PREFIX}/g" > ${BEAT_NAME}.yml

if [ -e _meta/beat.reference.yml ]; then
    cat _meta/beat.reference.yml | sed -e "s/beatname/${BEAT_NAME}/g;s/beat-index-prefix/${BEAT_INDEX_PREFIX}/g" > ${BEAT_NAME}.reference.yml
else
    cat _meta/beat.yml | sed -e "s/beatname/${BEAT_NAME}/g;s/beat-index-prefix/${BEAT_INDEX_PREFIX}/g" > ${BEAT_NAME}.reference.yml
fi

chmod 0640 ${BEAT_NAME}.yml ${BEAT_NAME}.reference.yml


go get github.com/elastic/beats/v7/libbeat/scripts/cmd/global_fields
global_fields -es_beats_path vendor/github.com/elastic/beats/v7 -out fields.yml
go run dev-tools/cmd/asset/asset.go -pkg include -in fields.yml -out include/fields.go ${BEAT_NAME} fields.yml
