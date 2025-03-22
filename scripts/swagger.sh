#!/usr/bin/env bash

set -euo pipefail

if ! [[ "$0" =~ scripts/swagger.sh ]]; then
  echo "must be run from repo root"
  exit 255
fi

source ./scripts/lib.sh

# 创建临时目录
TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT

# 安装 swagger-ui-dist
log_callout "Installing swagger-ui-dist..."
npm install swagger-ui-dist --prefix "$TEMP_DIR"

# 创建 HTML 文件
cat > "$TEMP_DIR/index.html" << 'EOF'
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>API Documentation</title>
    <link rel="stylesheet" href="node_modules/swagger-ui-dist/swagger-ui.css" />
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="node_modules/swagger-ui-dist/swagger-ui-bundle.js"></script>
    <script>
        window.onload = () => {
            window.ui = SwaggerUIBundle({
                url: './user.yml',
                dom_id: '#swagger-ui',
            });
        };
    </script>
</body>
</html>
EOF

# 复制 OpenAPI 文件
cp api/openapi/user.yml "$TEMP_DIR/"

# 启动 HTTP 服务器
log_callout "Starting Swagger UI..."
log_callout "Visit http://localhost:8081 to view the API documentation"
cd "$TEMP_DIR" && python3 -m http.server 8081 