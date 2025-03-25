![1](https://github.com/user-attachments/assets/86556c4e-53be-4c30-ac46-049adc447d61)

关键设计说明
 (1) 按实体分离路由
- 每个实体（如`user`、`product`）在`cmd/api/routes`目录下有独立的路由注册文件。
- 路由文件通过`RegisterXxxRoutes`函数暴露注册逻辑。

 (2) 依赖集中管理
- 数据库连接、配置加载等全局依赖在`main.go`中初始化。
- 具体实体的Repository和Service在`registerRoutes`中创建，并通过参数传递给路由注册函数。

 (3) 分层清晰
- `main.go`仅负责：
  - 全局初始化（配置、日志、数据库）
  - 路由注册协调
  - 服务启动
- 业务细节（如路由处理、数据库操作）下沉到对应分层模块。

 (4) 扩展性
- 新增实体（如`product`）时只需：
  1. 在`internal`层补充`handler/service/repository`代码。
  2. 创建`cmd/api/routes/product.go`。
  3. 在`registerRoutes`中注入依赖并调用路由注册。

---

 4. 完整调用流程
text
main.go 
→ initDatabase() 
→ registerRoutes() 
  → 创建UserRepository 
  → 创建UserService 
  → 调用routes.RegisterUserRoutes() 
    → 初始化UserHandler 
    → 注册用户相关路由 
→ 启动Fiber服务

---

 5. 优势
1. 高内聚：每个文件/模块只关注单一职责。
2. 易维护：新增实体无需修改`main.go`主逻辑。
3. 可测试：路由模块可通过Mock依赖独立测试。
4. 低耦合：路由层不直接依赖具体数据库实现。

通过这种结构，项目可以轻松扩展多个实体，同时保持代码整洁和可维护性。
