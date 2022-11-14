$HADOOP_HOME/bin/hadoop jar /opt/hadoop-3.3.3/share/hadoop/tools/lib/hadoop-streaming-3.3.3.jar \
    -files /map_reduce/h_reducer,/map_reduce/h_mapper \
    -reducer "./h_reducer" \
    -mapper "./h_mapper" \
    -input /h31/user_logs -output /output/h31_log/

