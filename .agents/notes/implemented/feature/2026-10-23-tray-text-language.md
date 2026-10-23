# Agent Note: 托盘的用量文字随页面语言翻译

Status: implemented

## Problem

托盘菜单从 #301 起就有 zh、zh-TW、ja、de 四份译文（`trayWords`），但托盘真正长期显示的内容——菜单栏图标旁的百分比、每个额度窗口的 tooltip、没有卡片时的 idle tooltip、magpie 自己的今日/7 天合计——仍是英文拼接：`used`、`left`、`pace`、`resets in`、`as of … couldn't be read just now`、`Today`、`tokens`、`requests`。页面本身已按 `settings.Lang` 翻译（`t()` / `i18n.js`），于是中文用户看到的是“菜单是中文、悬浮提示是英文”的混合界面；日期和时长也固定为英文的 `Sep 27 09:05`、`2h 10m`。

## Decision

托盘的原生文字全部走 `traymenu.go` 的 `trayLang` + `trayText`，语言与菜单一致：`settings.Lang`，为 `system` 时按系统语言判定（zh-TW 在台湾、香港、澳门和任何 zh-Hant 下成立）。

- **菜单、卡片 tooltip、idle tooltip、合计 tooltip** 的固定词都在 `trayWords` 里；英文是 key，缺译时回落英文。
- **窗口名**用 `trayWindowName` 走同一张表，因此 `5 hours`、`Weekly`、`Month`、`Credits`、`Session`、`Allowance`、`Limit` 等与 Usage 页一致；表里没有的名字（插件或厂商自定义）原样保留。
- **时长**由 `trayUntil` 按语言的单位表生成：`3d 4h` / `3 天 4 小时` / `3 日 4 時間` / `3 T. 4 Std.`；**日期**由 `trayDateLayout` 给出 `Jan 2 15:04`、`1月2日 15:04` 或 `2. Jan. 15:04`。
- **刷新时机**：菜单由 `onLang` 立即重标；设置保存检测到 `Lang` 变化时同时调用 `onTrayUsage`，唤醒 `watchTrayUsage` 重画卡片，不必等下一次 3 分钟轮询。面板仍是页面的 `t()`，本改动不碰它。
- **双语表是刻意的**：`internal/gui/assets/i18n.js` 服务 webview，托盘在面板尚未加载、窗口关闭甚至 WebKit 不可用时也要立刻画出 tooltip，所以原生侧保留自己的表；`TestTrayWordsCoverEachLanguage` 保证新增的托盘 key 在四种语言里都有值，但不检查与 `i18n.js` 逐字一致。

## Alternatives considered

- **保持现状：只翻译菜单，用量文字留英文** — 维护成本最低，英文字段也看得懂。但托盘最常被看到的是悬浮提示，菜单反而很少打开；混排正是这次要修的东西。
- **把托盘文字也放进 `i18n.js`，由页面翻译后回传** — 单一翻译源，最不容易漂移。但托盘是原生实现的（macOS 的 `NSStatusItem` tooltip、Linux 的 StatusNotifierItem `ToolTip`），启动时页面可能还没加载，窗口关闭时也没有可用的 webview；让托盘显示依赖页面执行会把一条同步路径变成异步且可能失败。
- **构建时从 `i18n.js` 生成 Go 表** — 仍是单一源，能机械对齐。但会引入一个生成步骤和生成物入库/校验问题，而且托盘只需要几十个 key；当前仓库已经接受菜单表的重复（#301），这次沿用同一模式，把风险写进后果。
- **只翻译窗口名，不翻译句子** — 能解决“英文窗口名”的一部分，但 `used`、`resets in` 仍混合；且窗口名的翻译量并不比句子小。

## Consequences

- **收益**：托盘与页面使用同一语言，中文、繁体、日文、德文用户不再看到英文句子；窗口名和时长、日期也随语言变化；语言切换立即生效。
- **代价与已知上限**：`traymenu.go` 与 `i18n.js` 是两份译文，新增托盘词条要同时改两处，可能漂移；窗口名只覆盖页面已翻译的固定集合，插件自定义名仍是原文；`q.Error` 仍是厂商/插件给出的原文（页面用 `quotaError` 做模式翻译，托盘没有照搬那套模式匹配）；`trayWords` 的完整性测试只查 key 存在，不查译文质量。
- **重访信号**：如果托盘词条继续增长、或 `i18n.js` 已有可复用的生成机制，应改为生成/共享表；若要翻译错误文案，另开一篇记录匹配规则，不要在托盘里复制 `quotaError` 的全部正则。

## Verification

- `go test -tags nogui ./internal/gui -run 'TestTray|TestSettingsLang'`：新用例逐语言断言 tooltip、idle tip、合计、时长和日期，并断言语言保存会触发一次重画。
- `go test -tags nogui ./internal/gui`：整包通过。
- `go test -tags nogui ./...`：`internal/gui` 通过；`internal/agent` 与 `internal/davsync` 的失败在未改动的 HEAD worktree 上同样出现（PATH 里缺 CLI、同步时间断言），不属于本改动。
- `GOOS=windows go build ./internal/gui` 与 `GOOS=windows go vet ./internal/gui`：通过，覆盖了只在 `!nogui` 下编译的 `trayusage_tray.go`。
- `go vet -tags nogui ./...`、`gofmt -l` 对改动文件无输出。
- 未运行 `make test-ui`：本改动不涉及 `internal/gui/assets`，页面翻译与布局未变。
- 本机没有 GTK/WebKitGTK 开发包（`pkg-config` 缺失），因此 Linux 的 `!nogui`（GTK）构建无法在这里编译；Windows 交叉构建覆盖了托盘调用点，macOS 的 AppKit 构建仍需在 CI 或桌面上确认。
