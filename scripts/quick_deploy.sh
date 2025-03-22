#!/bin/bash

# 设置错误时退出
set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# 打印带颜色的信息
print_info() {
    echo -e "${GREEN}[INFO] $1${NC}"
}

print_warn() {
    echo -e "${YELLOW}[WARN] $1${NC}"
}

print_error() {
    echo -e "${RED}[ERROR] $1${NC}"
}

# 检查必要的命令
check_commands() {
    local commands=("docker" "docker-compose")
    for cmd in "${commands[@]}"; do
        if ! command -v $cmd &> /dev/null; then
            print_error "$cmd is not installed"
            exit 1
        fi
    done
}

# 快速构建并重启服务
quick_rebuild() {
    print_info "Quick rebuilding service..."
    
    # 只重新构建 user-service
    docker-compose build user-service
    
    # 重启 user-service
    docker-compose restart user-service
}

# 检查服务状态
check_services() {
    print_info "Checking services status..."
    docker-compose ps
}

# 主函数
main() {
    print_info "Starting quick deployment process..."
    
    # 检查必要的命令
    check_commands
    
    # 快速重建并重启服务
    quick_rebuild
    
    # 检查服务状态
    check_services
    
    print_info "Quick deployment completed successfully!"
}

# 执行主函数
main 