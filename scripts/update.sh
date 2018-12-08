#!/bin/bash

BEAT_NAME=${BEAT_NAME:-jsonbeat}
BEAT_INDEX_PREFIX=${BEAT_INDEX_PREFIX:-${BEAT_NAME}}


cat _meta/beat.yml vendor/github.com/elastic/beats/libbeat/_meta/config.yml | sed -e "s/beatname/${BEAT_NAME}/g;s/beat-index-prefix/${BEAT_INDEX_PREFIX}/g" > ${BEAT_NAME}.yml

if [ -e _meta/beat.reference.yml ]; then
    cat _meta/beat.reference.yml vendor/github.com/elastic/beats/libbeat/_meta/config.reference.yml | sed -e "s/beatname/${BEAT_NAME}/g;s/beat-index-prefix/${BEAT_INDEX_PREFIX}/g" > ${BEAT_NAME}.reference.yml
else
    cat _meta/beat.yml vendor/github.com/elastic/beats/libbeat/_meta/config.reference.yml | sed -e "s/beatname/${BEAT_NAME}/g;s/beat-index-prefix/${BEAT_INDEX_PREFIX}/g" > ${BEAT_NAME}.reference.yml
fi

chmod 0640 ${BEAT_NAME}.yml ${BEAT_NAME}.reference.yml


go run vendor/github.com/elastic/beats/libbeat/scripts/cmd/global_fields/main.go -es_beats_path vendor/github.com/elastic/beats -out fields.yml
