#!/usr/bin/env bash

set -euo pipefail

if ! [[ "$0" =~ scripts/start.sh ]]; then
  echo "must be run from repo root"
  exit 255
fi

source ./scripts/lib.sh

# 设置环境变量
export JWT_SECRET="your-secret-key-here"  # 在生产环境中应该使用更安全的方式存储
export PORT="8080"

# 运行服务
log_callout "Starting user service..."
go run ./internal/user/main.go