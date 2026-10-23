# Agent Note: 非平凡改动的理由落在 .agents/notes

Status: implemented

## Problem

仓库已经有 AGENTS.md 说明规矩、docs/subsystems 说明子系统怎么工作、LESSONS.md 说明踩过什么坑，但没有一处放“为什么选 A 不选 B”。这类理由目前只存在于 PR 描述、提交信息或代码旁注里：下一个改同一处的人得重新推断，或者把已经否决的方案再走一遍。具体形态是重试上限改过去又改回来、同一份 provider 迁移策略在多个 PR 里各解释一遍，而 LESSONS.md 面向的是事后复盘，装不下一个还在生效的选型。

## Decision

非平凡改动的决定、取舍和被否掉的备选，写在 `.agents/notes/` 下，路径即分类：`{proposed,implemented,rejected}/{feature,bug-fix,simplification,architecture,process,testing}/yyyy-mm-dd-topic.md`。`archived/` 只收冻结的历史笔记。空目录用到再建，不预铺全部组合。

- **前看**：新立的行为规则、跨模块契约、落盘/网络格式，记成 `implemented`，与代码同一批提交。
- **回看**：为看不见的约束放弃了直觉解法，把对手方案的最强理由和被否原因写进 `## Alternatives considered`。
- **减法**：删能力、收窄 API 面，写清退出条件和迁移边界。
- **翻转**：事实变了就地改老笔记的路径、符号、默认值；决定本身翻转则新开一篇并互链，禁止把 `## Decision` 改写成反面。
- **纯机械改动**：格式化、错别字、无歧义重命名、样式调整、依赖补丁、发布打标、单模块内看 diff 即懂的修复，不写笔记。

笔记用简体中文，骨架固定为 `## Problem` / `## Decision` / `## Alternatives considered` / `## Consequences`，`implemented` 用现在时，不出现提案式标题。校验入口是 `make verify-notes`：树结构、头块与必需节、归档封印三项串跑，Node 22+ 直接执行 Skill 里的 TypeScript 脚本（`--experimental-strip-types`），不需要 `npm install`。

AGENTS.md 增加“Agent notes”一节指向这里；`.github/workflows/test.yml` 的 `notes` job 在 Linux 上跑同一命令。归档封印的 append-only 对照用 `AGENT_NOTE_ARCHIVE_BASE_REF` 指定变更前的 commit（PR 用 `github.event.pull_request.base.sha`，push 用 `github.event.before`），因为脚本默认的 `HEAD` 包含本次改动，等于没查。

## Alternatives considered

- **`docs/adr/` 风格的架构决策记录** — 业界标准，评审流程成熟。但 ADR 只有“已接受”一种终态，装不下 proposed/rejected 和按 class 的切片；再建一套 docs 目录也会和 docs/subsystems 的设计说明重叠。选 lifecycle × class 的四层结构，让检索可以先切意图再搜关键词。
- **单文件 `NOTES.md` 或继续写进 `LESSONS.md`** — 维护成本最低，一眼能看完。但单文件在并行 PR 里必然冲突，也没有机械校验能约束一条记录的完整性；LESSONS.md 的语义是事后教训，把前瞻决策混进去会稀释两者。
- **只靠 PR 描述和提交信息** — 零额外文件，写的时候最省事。但 code review 之后就没有人从源码走到那段理由，搜索也依赖平台的索引；决定和代码的对应关系会断。
- **写笔记但不接门禁** — 摩擦最小，避免为了格式返工。代价是头块、必需节和归档封印会随第一批赶工的笔记一起腐化；脚本只检查结构，本来就不检查理由是否成立，这点摩擦由写作者承担。

## Consequences

- **收益**：选型和被否方案能从仓库内搜到，并与代码同批复审；`make verify-notes` 让格式与归档冻结可机械检查，不依赖记忆。
- **代价与已知上限**：非平凡改动多出一道书写成本；脚本只保证结构，不保证理由诚实或完整，语义仍靠作者和 reviewer。CI 多一个 Node 22 的 job，`notes` 失败会挡住合并。归档后笔记不可改，发现事实错误只能新开一篇互链。
- **重访信号**：如果连续多次 PR 的笔记没人读、或 `notes` job 长期只做格式修补，说明分类或门禁过重，应缩减检查项而不是继续加规则。

## Verification

- `make verify-notes` 在当前树通过（1 篇笔记，0 条归档封印）。
- 脚本位于 `.agents/skills/write-notes-like-deepseek/scripts/`，`verify-agent-note-tree.ts` 校验路径与相对链接，`verify-agent-note-format.ts` 校验头块与必需节，`verify-archived-agent-notes.ts` 校验 `archived/manifest.json` 的 sha256 封印。
- CI 入口在 `.github/workflows/test.yml` 的 `notes` job；本地入口在 `Makefile` 的 `verify-notes` target。
