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

# 备份数据库
backup_database() {
    if [ -d "mysql-data" ]; then
        print_info "Backing up database..."
        timestamp=$(date +%Y%m%d_%H%M%S)
        tar -czf "mysql-backup_${timestamp}.tar.gz" mysql-data/
    fi
}

# 停止并删除旧容器
cleanup_old_containers() {
    print_info "Cleaning up old containers..."
    docker-compose down || true
}

# 拉取最新代码
pull_latest_code() {
    print_info "Pulling latest code..."
    git pull origin main
}

# 构建新镜像
build_new_image() {
    print_info "Building new image..."
    docker-compose build --no-cache
}

# 启动服务
start_services() {
    print_info "Starting services..."
    docker-compose up -d
}

# 检查服务状态
check_services() {
    print_info "Checking services status..."
    docker-compose ps
}

# 主函数
main() {
    print_info "Starting deployment process..."
    
    # 检查必要的命令
    check_commands
    
    # 备份数据库
    backup_database
    
    # 停止并删除旧容器
    cleanup_old_containers
    
    # 拉取最新代码
    pull_latest_code
    
    # 构建新镜像
    build_new_image
    
    # 启动服务
    start_services
    
    # 检查服务状态
    check_services
    
    print_info "Deployment completed successfully!"
}

# 执行主函数
main 