# GOrder v2

> ⚠️ **项目状态**: 此项目正在积极开发中，部分功能尚未完善。欢迎贡献代码和反馈！

基于 Go 语言构建的微服务订单管理系统，采用领域驱动设计（DDD）和六边形架构模式，实现了订单处理、支付管理、库存管理等核心业务功能。

## ✨ 特性

- 🏗️ **微服务架构**: 基于领域驱动设计的松耦合服务
- 🔄 **分布式追踪**: 集成 Jaeger 实现完整的链路追踪
- 📨 **事件驱动**: 基于 RabbitMQ 的异步消息通信
- 🔍 **服务发现**: 使用 Consul 进行动态服务注册与发现
- 💳 **支付集成**: 集成 Stripe 支付网关
- 🛡️ **防腐层**: 抽象外部服务集成，提高系统稳定性
- 📋 **代码规范**: 严格的 Go 代码规范和质量检查

## 🏛️ 架构概览

### 微服务组件

| 服务 | 描述 | 协议 | 端口 |
|------|------|------|------|
| **Order Service** | 订单管理核心服务 | HTTP + gRPC | 8282 (HTTP), 5002 (gRPC) |
| **Payment Service** | 支付处理服务 | HTTP | 8284 |
| **Stock Service** | 库存管理服务 | gRPC | 5003 |
| **Kitchen Service** | 厨房服务（开发中） | - | - |

### 基础设施组件

- **Consul** (8500): 服务发现与配置管理
- **RabbitMQ** (5672, 15672): 消息队列
- **Jaeger** (16686): 分布式追踪
- **Stripe**: 支付网关集成

### 架构模式

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Order Service │    │ Payment Service │    │  Stock Service  │
│                 │    │                 │    │                 │
│ ┌─────────────┐ │    │ ┌─────────────┐ │    │ ┌─────────────┐ │
│ │   Domain    │ │    │ │   Domain    │ │    │ │   Domain    │ │
│ │  Business   │ │    │ │  Business   │ │    │ │  Business   │ │
│ │    Logic    │ │    │ │    Logic    │ │    │ │    Logic    │ │
│ └─────────────┘ │    │ └─────────────┘ │    │ └─────────────┘ │
│ ┌─────────────┐ │    │ ┌─────────────┐ │    │ ┌─────────────┐ │
│ │   Adapters  │ │    │ │   Adapters  │ │    │ │   Adapters  │ │
│ │    Ports    │ │    │ │    Ports    │ │    │ │    Ports    │ │
│ └─────────────┘ │    │ └─────────────┘ │    │ └─────────────┘ │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                        │                        │
         └────────────────────────┼────────────────────────┘
                                  │
                    ┌─────────────────────────┐
                    │     Common Module       │
                    │   (Anti-corruption      │
                    │       Layer)            │
                    │                         │
                    │ • Service Discovery     │
                    │ • Message Broker        │
                    │ • Distributed Tracing   │
                    │ • Configuration         │
                    │ • Logging & Metrics     │
                    └─────────────────────────┘
```

## 🚀 快速开始

### 前置要求

- Go 1.21+
- Docker & Docker Compose
- Make
- Git

### 1. 克隆项目

```bash
git clone https://github.com/jialechen7/gorder-v2.git
cd gorder-v2
```

### 2. 启动基础设施

```bash
# 启动 Consul、RabbitMQ 和 Jaeger
docker-compose up -d

# 验证服务状态
docker-compose ps
```

### 3. 配置环境变量

```bash
# 设置 Stripe 相关环境变量（用于支付功能）
export STRIPE_KEY="your_stripe_secret_key"
export ENDPOINT_STRIPE_SECRET="your_webhook_endpoint_secret"
```

### 4. 生成代码

```bash
# 生成 protobuf 和 OpenAPI 代码
make gen
```

### 5. 启动服务

```bash
# 启动 Stock Service
cd internal/stock
go run main.go

# 启动 Order Service
cd internal/order  
go run main.go

# 启动 Payment Service
cd internal/payment
go run main.go
```

### 6. 验证部署

访问以下地址验证服务状态：

- **Consul UI**: http://localhost:8500
- **RabbitMQ Management**: http://localhost:15672 (guest/guest)
- **Jaeger UI**: http://localhost:16686
- **Order Service**: http://localhost:8282/api

## 🛠️ 开发指南

### 代码生成

```bash
make gen          # 生成 protobuf 和 OpenAPI 代码
make genproto     # 仅生成 protobuf 代码
make genopenapi   # 仅生成 OpenAPI 代码
```

### 代码质量

```bash
make lint         # 运行代码检查 (golangci-lint + go-cleanarch)
make fmt          # 格式化代码 (goimports)
```

### Stripe 开发

```bash
# 启动 Stripe webhook 监听器（开发环境）
make listen
```

### 目录结构

```
├── api/                    # API 协议定义
│   ├── openapi/           # OpenAPI 规范
│   └── orderpb/           # protobuf 定义
├── internal/              # 内部模块
│   ├── common/           # 共享组件（防腐层）
│   │   ├── broker/       # 消息队列
│   │   ├── discovery/    # 服务发现
│   │   ├── tracing/      # 分布式追踪
│   │   └── config/       # 配置管理
│   ├── order/            # 订单服务
│   ├── payment/          # 支付服务
│   ├── stock/            # 库存服务
│   └── kitchen/          # 厨房服务
├── scripts/              # 构建和工具脚本
└── public/               # 静态资源
```

### 编码规范

项目遵循严格的 Go 编码规范：

- 使用 `golangci-lint` 进行代码质量检查
- 遵循 Go 官方命名约定
- 实施 clean architecture 检查
- 强制错误处理和上下文传递规范

## 📋 项目状态

### ✅ 已完成

- [x] 基础微服务架构搭建
- [x] 服务发现与配置管理
- [x] 分布式追踪集成
- [x] 消息队列通信机制
- [x] Order Service 核心功能
- [x] Stock Service 基本功能
- [x] Payment Service 和 Stripe 集成
- [x] 代码质量工具链

### 🚧 开发中

- [ ] Kitchen Service 完整实现
- [ ] 单元测试和集成测试
- [ ] API 文档完善
- [ ] 容器化部署方案
- [ ] 监控和告警系统
- [ ] 性能优化

### 📈 计划中

- [ ] 用户认证和授权
- [ ] 数据持久化层
- [ ] 缓存策略
- [ ] API 网关集成
- [ ] CI/CD 流水线
- [ ] 生产环境配置

## 🤝 贡献指南

欢迎任何形式的贡献！请遵循以下步骤：

1. Fork 本仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### 代码贡献规范

- 提交前运行 `make lint` 确保代码质量
- 为新功能添加适当的测试
- 更新相关文档
- 遵循现有的代码风格和架构模式

## 📄 许可证

本项目采用 [MIT 许可证](LICENSE)。

## 📞 联系方式

- 作者: jialechen7
- 项目地址: https://github.com/jialechen7/gorder-v2
- 问题反馈: [GitHub Issues](https://github.com/jialechen7/gorder-v2/issues)

---

> 💡 **提示**: 如果您在使用过程中遇到问题，请查看 [CLAUDE.md](CLAUDE.md) 获取更多技术细节，或在 Issues 中寻求帮助。