FROM bde2020/hadoop-base:2.0.0-hadoop3.3.3-java8

WORKDIR /

RUN mkdir -p map_reduce

COPY ./error_map_reduce/bin/* /map_reduce/
COPY ./general_reduce/bin/* /map_reduce/
COPY ./h_map_reduce/bin/* /map_reduce/

ADD run.sh /run.sh
RUN chmod a+x /run.sh

CMD ["/run.sh"]
