#!/usr/bin/env bash

# -e : 各コマンド正常終了以外の時にスクリプト終了
# -u : 未定義変数参照時にスクリプト終了
# -o pipefail : pipe使用時の各コマンド正常終了以外の時にスクリプト終了
# -x : デバッグ情報出力
set -euox pipefail

# スクリプトファイル名の出力
SCRIPT_FILE=$(basename $0)
echo ${SCRIPT_FILE}

# スクリプトファイル格納ディレクトリの出力
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}

cd ${SCRIPT_DIR} && cd ../src/graph/model/

dataloaden ViewingHistoryLoader int []*github.com/sky0621/fs-mng-backend/graph/model.ViewingHistory
