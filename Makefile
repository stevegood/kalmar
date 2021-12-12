BUILD_PATH=./build
BIN_PATH=${BUILD_PATH}/bin
KALMAR_BIN=${BIN_PATH}/kalmar
CONFIG_PATH=./config
GQLGEN_CONFIG_PATH=${CONFIG_PATH}/gqlgen.yml
WEB_PATH=./web

clean:
	rm -rf ${BUILD_PATH}

build: clean
	cd ${WEB_PATH} && yarn build && cd ..
	mkdir -p ${BIN_PATH}
	go build -o ${BIN_PATH} .

cmd-serve:
	${KALMAR_BIN} serve

cmd-test:
	${KALMAR_BIN} test

cmd-setup:
	${KALMAR_BIN} setup

web-setup:
	cd ${WEB_PATH} && yarn

web-dev:
	cd ${WEB_PATH} && yarn start

test:
	go test ./...

test-verbose:
	go test -v ./...

gql-generate:
	go get github.com/99designs/gqlgen/cmd
	go run github.com/99designs/gqlgen generate --verbose --config ${GQLGEN_CONFIG_PATH}
