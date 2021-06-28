#!/bin/bash

echo Migrating database
until /project/tools/shmig.sh -t postgresql -H database -l capitalUser -p hARdAYTEReIsUlANTanEOLIB -d capital -m /project/tools/migrations -s migrations up
do
    echo retrying migration
    sleep 3
done
