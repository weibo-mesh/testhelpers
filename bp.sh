#!/usr/bin/env bash

### BEGIN ###
# Author: idevz
# Since: 09:18:11 2019/01/10
# Description:       bp for WeiboMesh testhelpers
# bp          ./bp.sh
#
# Environment variables that control this script:
#
### END ###

set -e

images_bp() {
	docker build -t zhoujing/wm-testhelper-mesh -f mesh.Dockerfile .
	docker push zhoujing/wm-testhelper-mesh

	docker build -t zhoujing/wm-testhelper-server -f server.Dockerfile .
	docker push zhoujing/wm-testhelper-server
}

images_bp
