#!/usr/bin/env bash

### BEGIN ###
# Author: idevz
# Since: 09:18:11 2019/01/10
# Description:       run for WeiboMesh testhelpers
# run          ./run.sh
#
# Environment variables that control this script:
#
### END ###

set -ex
BASE_DIR=${BASE_DIR:-"$(readlink -f "$(dirname "$0")")"}

# the default version same to motan-go version
VERSION=${V:-"1.0.0"}
ORI_HELP_IMG=${OHIMG:-"zhoujing/wm-testhelper"}

# the Golang runner image should using the version as Golang it Self
GOLANG_IMG=${GIMG:-"zhoujing/idevz-runx-golang:1.12"}
TEST_HELP_IMG=${TIMG:-"${ORI_HELP_IMG}:${VERSION}"}
MTSERVER_EXEC=${MTS:-"mtserver"}
TMESH_EXEC=${TMSH:-"tmesh"}

MTCONTAINER_NAME=${MTC_NAME:-"wm-testhelper"}

build_exec() {
	local cmd=${1}
	sudo docker run -e CGO_ENABLED=0 \
		-e GOOS=linux -e GOARCH=amd64 \
		-v $GIT/go:/go \
		-v $MCODE/z/idevz-go:/z/idevz-go \
		-v $GIT/weibo-mesh/testhelpers:/runX/testhelpers \
		-w /runX/testhelpers/cmd/${cmd} \
		"${GOLANG_IMG}" go build -a -installsuffix cgo -o "/runX/testhelpers/${cmd}"
}

images_bp() {
	local mt_exec="${BASE_DIR}/${MTSERVER_EXEC}"
	local tmesh_exec="${BASE_DIR}/${TMESH_EXEC}"
	[ -f "${mt_exec}" ] && sudo rm "${mt_exec}"
	[ -f "${tmesh_exec}" ] && sudo rm "${tmesh_exec}"
	build_exec ${MTSERVER_EXEC}
	build_exec ${TMESH_EXEC}

	sudo docker build -t "${ORI_HELP_IMG}" .
	sudo docker tag "${ORI_HELP_IMG}" "${TEST_HELP_IMG}"
	sudo docker push "${ORI_HELP_IMG}"
	sudo docker push "${TEST_HELP_IMG}"
}

show_help() {
	echo "
    ./run.sh                show this help
    ./run.sh -h             show this help


    ./run.sh bp             build and push docker images 
    ./run.sh upd            start the test server and mesh 
    ./run.sh downd          stop and clean the server and mesh containers
    "
}

case "${1}" in
bp)
	images_bp
	;;
mc)
	sudo docker run --network host -d --rm --name mc -v memcached
	;;
upd)
	sudo docker run --network host -d --rm --name "${MTCONTAINER_NAME}" \
		-v $(pwd)/snapshot:/snapshot "${ORI_HELP_IMG}"
	;;
downd)
	sudo docker stop "${MTCONTAINER_NAME}"
	;;
up)
	./mtserver -c ./mesh.yaml &
	;;
down)
	pkill mtserver
	rm -rf ./agent_runtime ./mesh-logs ./serverlogs ./agent.pid
	;;
*)
	show_help
	;;
esac
