build:
	cd ./error_map_reduce/mapper && ${MAKE} build && cd ..
	cd ./error_map_reduce/reducer && ${MAKE} build && cd ..
	cd ./h_map_reduce/mapper && ${MAKE} build && cd ..
	cd ./h_map_reduce/reducer && ${MAKE} build && cd ..
	cd ./general_reduce/reducer && ${MAKE} build && cd ..
	cd ./general_reduce/mapper && ${MAKE} build && cd ..

clean:
	cd ./error_map_reduce/mapper && ${MAKE} clean && cd ..
	cd ./error_map_reduce/reducer && ${MAKE} clean && cd ..
	cd ./h_map_reduce/mapper && ${MAKE} clean && cd ..
	cd ./h_map_reduce/reducer && ${MAKE} clean && cd ..
	cd ./general_reduce/reducer && ${MAKE} clean && cd ..
	cd ./general_reduce/mapper && ${MAKE} clean && cd ..
	rm -rf ./general_reduce/bin
	rm -rf ./h_map_reduce/bin
	rm -rf ./error_map_reduce/bin

test: $(eval SHELL:=/bin/bash)
#	./error_map_reduce/bin/error_mapper < ../../data/h31/station_logs/stationlog.h31.csv | ./error_map_reduce/bin/error_reducer
	./h_map_reduce/bin/h_mapper < <(cat ../../data/h31/user_logs/userlog.h31.u1.csv ../../data/h31/user_logs/userlog.h31.u2.csv) | ./h_map_reduce/bin/h_reducer
#	./general_reduce/bin/general_reducer < log.log

generate_bin:
	mkdir ../../map_reduce
	mv ./error_map_reduce/bin/* ../../map_reduce/
	mv ./h_map_reduce/bin/* ../../map_reduce/
	mv ./general_reduce/bin/* ../../map_reduce/
