**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-03 12:39:00 UTC（2026-03-03 20:39:00 UTC+8）

**代理总数：36**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 36 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 36 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 225ms | 否 | ✓ 122ms | ✓ 1420ms | ✓ 1282ms | http |
| 45.95.0.142:443 | ✓ 438ms | ✓ 1837ms | ✓ 1358ms | ✓ 1888ms | ✓ 1721ms | http |
| 205.209.118.30:3138 | ✓ 1308ms | 否 | ✓ 706ms | 否 | ✓ 1935ms | http |
| 166.0.192.117:8888 | ✓ 958ms | 否 | ✓ 1669ms | 否 | ✓ 1687ms | http |
| 81.70.169.194:80 | ✓ 1123ms | ✓ 1467ms | ✓ 1728ms | ✓ 1357ms | ✓ 1094ms | http |
| 115.231.181.40:8128 | ✓ 1015ms | 否 | 否 | ✓ 1862ms | ✓ 1358ms | http |
| 101.43.255.96:80 | ✓ 1027ms | ✓ 1520ms | ✓ 1106ms | ✓ 1334ms | ✓ 1075ms | http |
| 120.92.212.16:7890 | ✓ 1158ms | ✓ 1572ms | ✓ 1582ms | 否 | ✓ 1529ms | http |
| 195.123.209.48:3128 | ✓ 853ms | 否 | ✓ 1801ms | 否 | ✓ 1499ms | http |
| 35.234.17.221:8080 | ✓ 1506ms | 否 | ✓ 1461ms | 否 | ✓ 956ms | http |
| 95.85.252.153:21064 | ✓ 521ms | ✓ 1547ms | ✓ 1617ms | 否 | 否 | http |
| 125.128.12.144:3128 | 否 | ✓ 1709ms | ✓ 1842ms | ✓ 1135ms | ✓ 876ms | http |
| 121.230.9.205:1080 | ✓ 1084ms | 否 | ✓ 1080ms | ✓ 1915ms | 否 | http |
| 121.128.121.54:3128 | 否 | ✓ 1015ms | ✓ 1781ms | ✓ 1130ms | ✓ 1957ms | http |
| 91.193.240.157:9877 | ✓ 1065ms | 否 | ✓ 1622ms | 否 | ✓ 1732ms | http |
| 120.92.212.16:8890 | ✓ 1945ms | 否 | ✓ 1054ms | ✓ 1619ms | ✓ 1098ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1131ms | ✓ 1149ms | ✓ 907ms | http |
| 192.166.82.55:1080 | ✓ 1185ms | 否 | 否 | ✓ 1935ms | ✓ 1662ms | http |
| 132.226.235.199:1080 | ✓ 1279ms | 否 | ✓ 1488ms | ✓ 1237ms | ✓ 1069ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 983ms | ✓ 1208ms | ✓ 769ms | http |
| 138.124.53.25:7443 | ✓ 639ms | 否 | ✓ 1873ms | 否 | ✓ 1686ms | http |
| 61.72.110.94:3128 | 否 | 否 | ✓ 1779ms | ✓ 1173ms | ✓ 1916ms | http |
| 194.87.250.190:1080 | 否 | 否 | ✓ 561ms | ✓ 1779ms | ✓ 1076ms | http |
| 150.107.140.238:3128 | ✓ 1821ms | 否 | ✓ 904ms | 否 | ✓ 1085ms | http |
| 91.233.223.147:3128 | ✓ 881ms | 否 | ✓ 908ms | ✓ 1957ms | ✓ 1538ms | http |
| 107.173.111.110:7890 | ✓ 990ms | ✓ 1543ms | ✓ 790ms | ✓ 909ms | ✓ 713ms | http |
| 45.136.198.40:3128 | ✓ 755ms | 否 | ✓ 1370ms | 否 | ✓ 1841ms | http |
| 103.215.36.88:16777 | ✓ 1922ms | 否 | ✓ 1294ms | ✓ 1741ms | ✓ 1790ms | http |
| 14.56.107.244:3128 | ✓ 1728ms | 否 | ✓ 1183ms | 否 | ✓ 1894ms | http |
| 121.230.8.111:1080 | ✓ 1258ms | 否 | 否 | ✓ 1649ms | ✓ 1122ms | http |
| 125.128.12.14:3128 | ✓ 1811ms | ✓ 1231ms | ✓ 750ms | 否 | 否 | http |
| 192.71.213.85:9812 | ✓ 475ms | 否 | ✓ 464ms | ✓ 1506ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1912ms | 否 | ✓ 1488ms | 否 | ✓ 1423ms | http |
| 210.223.44.230:3128 | ✓ 907ms | ✓ 1612ms | ✓ 755ms | ✓ 1029ms | ✓ 832ms | http |
| 103.39.51.190:8080 | ✓ 1886ms | 否 | 否 | ✓ 1502ms | ✓ 1449ms | http |
| 103.82.23.118:5207 | ✓ 1812ms | 否 | ✓ 1458ms | ✓ 1891ms | ✓ 1349ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
