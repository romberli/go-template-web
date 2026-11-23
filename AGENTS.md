# 四条核心原则

## Think Before Coding
- 明确陈述假设, 不确定就问, 达成一致最重要
- 多种解读时呈现选项, 不要默默选一个
- 有更简单的方案就说, 该 push back 就 push back

## Simplicity First
- 最少代码解决问题, 不做投机性抽象
- 不加没要求的功能、灵活性、配置项
- 不为不可能的场景写错误处理
- 200 行能用 50 行解决？重写

## Surgical Changes
- 不"改进"相邻代码/注释/格式
- 匹配现有风格, 即使你会 differently
- 只清理你的改动造成的孤儿代码
- 发现无关死代码？提一嘴, 别删

## Goal-Driven Execution
- 把任务转成可验证目标：
  - "加验证" → "写无效输入测试, 然后让它过"
  - "修 bug" → "写复现测试, 然后让它过"
- 多步任务列计划：1. [步骤] → verify: [检查]

## 配置访问
- 只通过 `viper.GetXxx(config.XxxKey)` 读取配置, 禁止硬编码配置值
- key 常量必须定义在 `config/` 包中

## 依赖包规范
- 错误处理统一使用 `github.com/pingcap/errors`, 禁止使用标准库 `errors` 或其他错误包
- 日志输出统一使用 `github.com/romberli/log`, 禁止使用标准库 `log` 或其他日志包

# 项目约定

- 后端数据库为gtw, 可通过mcp访问, mysql-gtw
- 项目分为api, service, controller, repository层, controller与repository层的可导出代码(大写字母开头的), 必须有单元测试, 所有api必须在test/goland下有http测试
- Redis可通过mcp访问, redis-gls
- 所有的表都必须包含如下3个字段:
  ```sql
  del_flag         tinyint     DEFAULT 0                    NOT NULL COMMENT '删除标记: 0-未删除, 1-已删除',
  create_time      datetime(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL COMMENT '创建时间',
  last_update_time datetime(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间'
  ```
- 开发api时, 如果没有明确说明, 一律使用POST接口
- 禁止在if语句中内联短变量声明, 必须先声明再判断, 例如:
  ```go
  // 正确
  _, ok := dbSecrets[clientID]
  if !ok {
      ...
  }
  // 错误
  if _, ok := dbSecrets[clientID]; !ok {
      ...
  }
  ```
- 类型一律使用any, 禁止使用interface{}
- 配置文件: 不以-default.yaml结尾的配置文件不纳入git版本管理, 仅保留在本地(已在.gitignore中排除); 例外: 以_default.yaml, -default.yaml结尾的默认配置模板、docs/swagger.yaml(swagger生成的API文档)纳入版本管理

# 禁止事项
- 禁止在 Handler 中直接访问数据库
- 禁止在 Repository 中包含业务逻辑
- 如果返回的错误类型不是 `github.com/pingcap/errors`（通常是调用依赖包函数返回的错误）, 则必须用 `errors.Trace(err)` 包裹后返回
- 禁止未经明确同意直接提交代码, 仅生成commit message后提醒用户来提交即可
- 用户提问题的时候, 只要回答问题然后询问是否要进行相应的修改就可以了, 禁止一提问就未经同意直接修改
