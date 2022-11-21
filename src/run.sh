#!/bin/bash

$HADOOP_HOME/bin/hadoop jar /opt/hadoop-3.3.3/share/hadoop/tools/lib/hadoop-streaming-3.3.3.jar \
    -files /map_reduce/${REDUCER},/map_reduce/${MAPPER} \
    -reducer "./${REDUCER}" \
    -mapper "./${MAPPER}" \
    -input ${INPUT} -output /output/${OUTPUT}

