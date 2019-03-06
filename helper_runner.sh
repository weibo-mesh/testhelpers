#!/usr/bin/env sh

### BEGIN ###
# Author: idevz
# Since: 14:18:26 2019/03/06
# Description:       runner for docker container
# helper_runner          ./helper_runner.sh
#
# Environment variables that control this script:
#
### END ###

set -ex

/mtserver &
/tmesh -c ./mesh-confs -pool testhelper
