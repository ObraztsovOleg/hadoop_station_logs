FROM bde2020/hadoop-base:2.0.0-hadoop3.3.3-java8

WORKDIR /

RUN mkdir -p map_reduce/input/h31

COPY ./bin/h_* /map_reduce/
COPY ./data/h31/* /map_reduce/input/h31/

RUN \
    hadoop fs -mkdir -p input/h31 && \
    hdfs dfs -put ./map_reduce/input/h31/* input/h31 && \
    hadoop jar /opt/hadoop-3.3.3/share/hadoop/tools/lib/hadoop-streaming-3.3.3.jar \
    -files /map_reduce/h_mapper,/map_reduce/h_reducer \
    -mapper "./h_mapper" \
    -reducer "./h_reducer" \
    -input /input/h31 -output /output/h31
