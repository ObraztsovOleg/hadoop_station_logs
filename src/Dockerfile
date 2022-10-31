FROM bde2020/hadoop-base:2.0.0-hadoop3.3.3-java8

WORKDIR /

RUN mkdir -p map_reduce

COPY ./error_map_reduce/bin/* /map_reduce/
COPY ./general_reduce/bin/* /map_reduce/
COPY ./h_map_reduce/bin/* /map_reduce/

RUN \
    hadoop fs -mkdir -p input && \
    hdfs dfs -put ./error_output/ERR_h31/* input && \
    hdfs dfs -put ./output/h31/* input && \
    hadoop jar /opt/hadoop-3.3.3/share/hadoop/tools/lib/hadoop-streaming-3.3.3.jar \
    -files /map_reduce/general_reducer \
    -reducer "./general_reducer" \
    -input /input -output /output/h31
RUN \
    hadoop fs -mkdir -p input && \
    hdfs dfs -put ./error_output/ERR_h31/* input && \
    hdfs dfs -put ./output/h31/* input && \
    hadoop jar /opt/hadoop-3.3.3/share/hadoop/tools/lib/hadoop-streaming-3.3.3.jar \
    -files /map_reduce/general_reducer \
    -reducer "./general_reducer" \
    -input /input -output /output/h31
RUN \
    hadoop fs -mkdir -p input && \
    hdfs dfs -put ./error_output/ERR_h31/* input && \
    hdfs dfs -put ./output/h31/* input && \
    hadoop jar /opt/hadoop-3.3.3/share/hadoop/tools/lib/hadoop-streaming-3.3.3.jar \
    -files /map_reduce/general_reducer \
    -reducer "./general_reducer" \
    -input /input -output /output/h31
