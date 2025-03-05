#!/bin/bash

# 设置工作目录
cd /home/ec2-user/data1/boomstage/admin || exit

# 拉取最新的 dev 分支
echo "拉取最新的 dev 分支..."
git checkout dev
git pull origin dev

# 构建应用
echo "构建应用..."
go build -o boomstatge .

# 检查是否已有应用在运行，如果运行中则停止它
if pgrep -f "boomstatge" > /dev/null
then
    echo "停止正在运行的 boomstatge..."
    pkill -f "boomstatge"
fi

# 启动新构建的应用
echo "启动新构建的 boomstatge..."
nohup ./boomstatge > ./logs/boomstatge.log 2>&1 &

echo "操作完成，boomstatge 应用已重启。"
