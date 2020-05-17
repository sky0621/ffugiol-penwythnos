#!/usr/bin/env bash

set -euox pipefail

SCRIPT_FILE=$(basename $0)
echo ${SCRIPT_FILE}

SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}

cd ${SCRIPT_DIR} && cd ../src/graph/model/

dataloaden ViewingHistoryLoader string []*github.com/sky0621/fs-mng-backend/graph/model.ViewingHistory
