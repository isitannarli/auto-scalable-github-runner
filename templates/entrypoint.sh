#!/bin/sh

./config.sh -name $($RUNNER_NAME) \
            --token ${RUNNER_TOKEN} \
            --url https://github.com/${GITHUB_OWNER}/${GITHUB_REPOSITORY} \
            --work ${RUNNER_WORKDIR} \
            --unattended \
            --replace

remove() {
    ./config.sh remove --unattended --token "${RUNNER_TOKEN}"
}

trap 'remove; exit 130' INT
trap 'remove; exit 143' TERM

./run.sh "$*" &

wait $!