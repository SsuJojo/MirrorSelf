<div align="center">

# MirrorSelf

**一个带有情绪表达的轻量投喂 / 餐食分享网页。**

> 如果结局注定是分别，那么相遇还有意义吗？  
> 别害怕别离——因为她改变的那部分你，已化作镜中的自己。

</div>

## 项目是什么

MirrorSelf 是一个 Vue + Go 的全栈网页原型。当前代码实现的核心场景很简单：用户在页面里输入菜品名称或外卖链接，前端提交后由 Go 服务记录请求，并写入 PocketBase；页面会记住最近一次提交内容，后端还可以触发外部通知。

它不是一个通用社交平台，更像是一个围绕两个人之间“今天吃什么 / 投喂什么”这一小场景做出的完整可运行产品原型。

## 当前能力

- 输入菜品名称或外卖链接并提交
- Vue 页面即时显示最近一次提交内容
- 使用 `localStorage` 保留本机最近一次投喂记录
- 通过 Go API 接收并记录提交内容
- 将记录写入 PocketBase
- 对短时间内的高频重复提交做前端限制
- 后端使用 Zap 同时输出控制台日志与文件日志
- Go 服务同时托管构建后的前端静态文件
- Dockerfile 可用于容器化部署

## 技术栈

| 层级 | 技术 |
| --- | --- |
| Frontend | Vue 3, Vite, Naive UI, Axios |
| Backend | Go, Fiber |
| Data | PocketBase |
| Logging | Zap |
| Deployment | Docker |

## 工作流程

```text
用户输入菜品 / 外卖链接
        ↓
Vue 前端
        ↓ POST /api/meal
Go + Fiber
   ├─ 记录日志
   ├─ 触发通知
   └─ 写入 PocketBase
        ↓
页面展示最近一次提交内容
```

## 目录结构

```text
MirrorSelf/
├── frontend/          # Vue 3 + Vite 前端
│   └── src/
│       ├── components/
│       └── lib/
├── backend/           # Go + Fiber 后端
│   ├── main.go
│   ├── pb/            # PocketBase 集成
│   └── pb_data/       # 本地 PocketBase 数据
├── Dockerfile
└── LICENSE
```

## 本地运行

### 1. 构建前端

```bash
cd frontend
npm install
npm run build
```

### 2. 启动后端

```bash
cd ../backend
go run .
```

默认情况下，Go 服务监听 `:3001`，并托管 `frontend/dist` 中的前端构建产物。

## API

### `POST /api/meal`

提交一条投喂内容：

```json
{
  "meal": "今晚吃寿司"
}
```

成功时返回：

```json
{
  "status": "recorded"
}
```

### `POST /api/msgsomany`

用于记录短时间内高频点击行为。

## 项目状态

**Prototype / 可运行原型。**

核心的“输入 → 提交 → 持久化 → 展示 / 通知”链路已经存在，但当前项目仍保留明显的个人化实现与硬编码配置，不应直接视作生产级服务。若继续开发，优先项应是配置外置、鉴权、测试与部署环境整理。

## License

许可证内容见 [`LICENSE`](./LICENSE)。
