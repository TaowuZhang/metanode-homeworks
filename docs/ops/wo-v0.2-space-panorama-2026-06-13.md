# 沃壤 v0.2 · 全景合体图（Notion + GitHub 同框，2026-06-13）

> 本图是 v0.2 快照的**全景伴随图**，与主快照里的「v0.2 全空间运行流图」同尺度、互为参照：把 Notion 活现场与 GitHub 沉地层放进同一张照片，看清两边怎么经由栈桥 / 泊坞 / 借港出入场。
>
> - 镜头尺标：全景 ＝ Notion + GitHub + 一切（本图所在层级）
> - 来源：2026-06-13 Notion 对话回填。对话中的原图已滚出可见窗口，本图按已验证拓扑（星盘 / 腐海 / 沃壤）与 GitHub 地层重建。

## 全景流图

```mermaid
flowchart TB
    subgraph NOTION["🌊 Notion —— 活现场"]
        subgraph LAW["公约数 · 立法层"]
            L1["出入场底谱 · 物理法则界碑<br/>L1 物理学 · 隐喻受力面 · 摩擦沉积床"]
        end
        subgraph FUHAI["腐海 · 生态现场"]
            F1["七判 · 涌"]
            F2["AI共台讨论场 · 深水解理台"]
            F3["跨境资金锚点 v0"]
        end
        subgraph COCKPIT["接续入口"]
            C1["驾驶舱 · 风险审计待办总账 · 外显事实源面板"]
        end
        subgraph SOIL["沃壤 · 地层指针"]
            S1["薄指针 → 指向 GitHub 地层"]
        end
    end

    subgraph BRIDGE["⚓ 栈桥 / 泊坞 / 借港 —— 出入场"]
        BR1["泊坞：内运（Notion→GitHub 沉积）"]
        BR2["借港：外部 MCP / 平台接口"]
        BR3["栈桥：对外折射面（薄，等真实桥桩）"]
    end

    subgraph GITHUB["🪨 GitHub（worang）—— 沉地层"]
        GH1["入口页 + 成为 / 领域 / 资源"]
        GH2["wo CLI（Rust）+ docs/ops 协议快照"]
        GH3["workflows / scripts / workers · 自动化"]
        GH4["Issues #190–#197 硬锚点 · PR 候选沉积"]
    end

    OUT["🌍 真实的人与领域<br/>（求职 / 作品 / 现实摩擦）"]

    LAW --> FUHAI --> COCKPIT
    FUHAI -->|"对话长厚"| BR1
    BR1 ==>|"沉积"| GITHUB
    GITHUB ==>|"硬事实回供"| SOIL
    SOIL --> COCKPIT
    BR2 -.->|"多源来物"| GITHUB
    GITHUB ==>|"成熟出土"| BR3
    BR3 -.->|"目前很薄"| OUT
    OUT -.->|"现实摩擦回流"| FUHAI

    classDef sediment fill:#efe3d5,stroke:#8a6d3b
    classDef gauge fill:#dbe7f3,stroke:#3672c2
    classDef pending fill:#ffe3c2,stroke:#e08a1e,stroke-dasharray:5 3
    class GH1,GH2,GH3,GH4 sediment
    class C1 gauge
    class BR3 pending
```

> 图例（线义统一 · 全图谱通用 ＋ 色温双通道，依出入场底谱）：
> - 实线 `-->`：主流向 / 内部沉积链（东西真往这走）。
> - 粗实线 `==>`：跨层主沉积链（沉积 · 硬事实回供 · 成熟出土）。
> - 虚线 `-.->`：支撑 / 反向拉动 / 漏出 / 仍薄未接线。
> - 实线框节点：已建（在转）。
> - 虚线框节点：待建 / 半轮 / 很薄（设计在、未闭合，如栈桥 BR3）。
> - 💤：休眠段（设计完整、暂未启用、即醒，依「休眠不摘牌」仍画）。
> - ⚠：漏点（链路断在这就变负复利）。
> - 色温：🟤 棕＝GitHub 沉地层（旧事实 / 已沉积）；🔵 蓝＝判断 / 仪表（驾驶舱 · 外显事实源）；🟠 橙＝待建 / 注意（栈桥对外折射面薄，等真实桥桩）；⚪ 无色＝Notion 活现场 / 中性。颜色与文字双通道，同一张图同色单义。

## 一句话读法

里头（公约数立法 → 腐海现场 → 沃壤地层 → GitHub 沉积）已经是闭环、互证、稳的；整张图唯一薄的一段是 **栈桥（对外折射面）**——东西沉得很厚，但还没真正流出去被真实的人用到。全景看下来，下一程不在纵深，在边缘：朝外走。
