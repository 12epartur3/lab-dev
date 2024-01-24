#!/bin/bash

for arg in $@
do
	echo "create project id ${arg}"
        curl localhost:8080/debug/jobs/create-evaluation-snapshot -H 'Content-Type: application/json' -d '{"project_id":'${arg}'}'
done
