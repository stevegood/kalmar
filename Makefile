BUILD_PATH=./build
BIN_PATH=${BUILD_PATH}/bin
KALMAR_BIN=${BIN_PATH}/kalmar

clean:
	rm -rf ${BUILD_PATH}

build: clean
	mkdir -p ${BIN_PATH}
	go build -o ${BIN_PATH} .

cmd-serve:
	${KALMAR_BIN} serve

cmd-test:
	${KALMAR_BIN} test

cmd-setup:
	${KALMAR_BIN} setup

test:
	go test ./...

test-verbose:
	go test -v ./...
