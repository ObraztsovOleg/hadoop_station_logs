build:
	cd ./error_map_reduce/mapper && ${MAKE} build && cd ..
	cd ./error_map_reduce/reducer && ${MAKE} build && cd ..
	cd ./h_map_reduce/mapper && ${MAKE} build && cd ..
	cd ./h_map_reduce/reducer && ${MAKE} build && cd ..
	cd ./general_reduce/reducer && ${MAKE} build && cd ..

clean:
	cd ./error_map_reduce/mapper && ${MAKE} clean && cd ..
	cd ./error_map_reduce/reducer && ${MAKE} clean && cd ..
	cd ./h_map_reduce/mapper && ${MAKE} clean && cd ..
	cd ./h_map_reduce/reducer && ${MAKE} clean && cd ..
	cd ./general_reduce/reducer && ${MAKE} clean && cd ..
	rm -rf ./general_reduce/bin
	rm -rf ./h_map_reduce/bin
	rm -rf ./error_map_reduce/bin

test: $(eval SHELL:=/bin/bash)
	./error_map_reduce/bin/error_mapper < /home/obrol/Downloads/h31/station_log/stationlog.h31.csv | ./error_map_reduce/bin/error_reducer >> log.log && \
	./h_map_reduce/bin/h_mapper < /home/obrol/Downloads/h31/user_log/userlog.h31.u1.csv | ./h_map_reduce/bin/h_reducer >> log.log && \
	./general_reduce/bin/general_reducer < log.log