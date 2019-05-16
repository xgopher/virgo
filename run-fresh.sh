#!/bin/bash

#
# fresh 工具自动安装 && 执行
#
if  ! command -v fresh > /dev/null; then
    echo "fresh command has not found"
    echo "Begin install fresh ..."
    # 切到应用外部去安装 fresh
    cd $GOPATH
    go get -v github.com/pilu/fresh
    echo "Done ..."
    cd -
fi

fresh -c runner.conf