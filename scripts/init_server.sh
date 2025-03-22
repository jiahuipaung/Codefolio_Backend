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

# 检查是否为 root 用户
check_root() {
    if [ "$EUID" -ne 0 ]; then
        print_error "Please run as root"
        exit 1
    fi
}

# 更新系统
update_system() {
    print_info "Updating system..."
    apt-get update
    apt-get upgrade -y
}

# 安装必要的软件包
install_packages() {
    print_info "Installing required packages..."
    apt-get install -y \
        apt-transport-https \
        ca-certificates \
        curl \
        software-properties-common \
        git
}

# 安装 Docker
install_docker() {
    print_info "Installing Docker..."
    
    # 添加 Docker 的官方 GPG 密钥
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
    
    # 设置稳定版仓库
    echo \
        "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
    
    # 安装 Docker
    apt-get update
    apt-get install -y docker-ce docker-ce-cli containerd.io
    
    # 启动 Docker
    systemctl start docker
    systemctl enable docker
}

# 安装 Docker Compose
install_docker_compose() {
    print_info "Installing Docker Compose..."
    curl -L "https://github.com/docker/compose/releases/download/v2.24.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
}

# 配置防火墙
configure_firewall() {
    print_info "Configuring firewall..."
    ufw allow 22/tcp
    ufw allow 80/tcp
    ufw allow 443/tcp
    ufw allow 8080/tcp
    ufw allow 3306/tcp
    ufw --force enable
}

# 创建应用目录
create_app_directory() {
    print_info "Creating application directory..."
    mkdir -p /opt/codefolio
    chown -R $SUDO_USER:$SUDO_USER /opt/codefolio
}

# 主函数
main() {
    print_info "Starting server initialization..."
    
    # 检查是否为 root 用户
    check_root
    
    # 更新系统
    update_system
    
    # 安装必要的软件包
    install_packages
    
    # 安装 Docker
    install_docker
    
    # 安装 Docker Compose
    install_docker_compose
    
    # 配置防火墙
    configure_firewall
    
    # 创建应用目录
    create_app_directory
    
    print_info "Server initialization completed successfully!"
    print_info "Please clone your repository to /opt/codefolio and run deploy.sh"
}

# 执行主函数
main 