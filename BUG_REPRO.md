# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

按倒序查询事件后再输出 JSON，得到的结果又变回时间升序，和查询请求的顺序相反；文本输出顺序是对的，只有 JSON 输出被改掉。请修复事件 JSON 输出，让它保持调用方给定的顺序，同时不改变事件内容和其他报告格式，并保证全量测试通过。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/gogo-48
- 仓库地址：https://github.com/zhanglei10281852-gif/gogo-48.git
- parent SHA：8463700aafc5445ed768fe741ffc91e772303732

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/gogo-48.git bug-repro
cd bug-repro
git checkout --detach 8463700aafc5445ed768fe741ffc91e772303732
go test ./internal/report -run "^TestEventJSONPreservesCallerOrdering$" -count=1 -v
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/report -run "^TestEventJSONPreservesCallerOrdering$" -count=1 -v
=== RUN   TestEventJSONPreservesCallerOrdering
    ordering_regression_test.go:24: JSON output changed requested descending order: [3fa6a1c8db426f44db5bacc9 1f988077cc2b48829c091a17]
--- FAIL: TestEventJSONPreservesCallerOrdering (0.00s)
FAIL
FAIL	LogPilot/internal/report	0.002s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/report -run "^TestEventJSONPreservesCallerOrdering$" -count=1 -v
=== RUN   TestEventJSONPreservesCallerOrdering
    ordering_regression_test.go:24: JSON output changed requested descending order: [3fa6a1c8db426f44db5bacc9 1f988077cc2b48829c091a17]
--- FAIL: TestEventJSONPreservesCallerOrdering (0.03s)
FAIL
FAIL	LogPilot/internal/report	0.135s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

JSON 输出严格保持调用方传入顺序（含倒序）；事件字段内容与其他报告输出不回归；双架构定向、全量、build/vet 通过。
