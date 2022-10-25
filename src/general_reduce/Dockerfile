FROM bde2020/hadoop-base:2.0.0-hadoop3.3.3-java8

WORKDIR /

RUN mkdir -p map_reduce

COPY ./bin/* /map_reduce/

RUN \
    hadoop fs -mkdir -p input && \
    hdfs dfs -put ./error_output/ERR_h31/* input && \
    hdfs dfs -put ./output/h31/* input && \
    hadoop jar /opt/hadoop-3.3.3/share/hadoop/tools/lib/hadoop-streaming-3.3.3.jar \
    -files /map_reduce/error_mapper,/map_reduce/error_reducer \
    -mapper "./error_mapper" \
    -reducer "./error_reducer" \
    -input /input/h31 -output /output/h31
