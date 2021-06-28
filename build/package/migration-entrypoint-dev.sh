#!/bin/bash

echo Database migration started

until /project/tools/shmig.sh -t postgresql -H database -l ${POSTGRES_USER} -p ${POSTGRES_PASSWORD} -d capital -m /project/tools/migrations -s migrations up
do
    echo retrying migration
    sleep 3
done

echo Database migration finished
