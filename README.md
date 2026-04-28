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

最后更新：2026-04-28 21:59:27 UTC（2026-04-29 05:59:27 UTC+8）

**代理总数：34**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 34 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 34 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 15.204.230.28:3128 | ✓ 773ms | ✓ 1613ms | 否 | ✓ 882ms | ✓ 678ms | http |
| 45.167.124.71:999 | ✓ 1415ms | ✓ 1843ms | ✓ 1954ms | ✓ 1573ms | ✓ 1284ms | http |
| 8.211.166.184:8081 | ✓ 1878ms | 否 | ✓ 959ms | ✓ 1118ms | ✓ 1673ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1564ms | ✓ 1119ms | ✓ 1673ms | 否 | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 896ms | ✓ 1535ms | ✓ 851ms | http |
| 218.108.131.186:17890 | ✓ 1034ms | ✓ 1616ms | ✓ 1886ms | 否 | ✓ 1026ms | http |
| 86.104.72.220:1082 | ✓ 1302ms | ✓ 1521ms | 否 | ✓ 1420ms | ✓ 969ms | http |
| 86.104.72.220:1081 | 否 | 否 | ✓ 745ms | ✓ 1547ms | ✓ 913ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1332ms | ✓ 1186ms | ✓ 1359ms | 否 | http |
| 46.101.95.183:8888 | 否 | 否 | ✓ 1676ms | ✓ 1717ms | ✓ 1282ms | http |
| 212.58.132.5:8888 | ✓ 1467ms | 否 | ✓ 1158ms | 否 | ✓ 1588ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 719ms | ✓ 951ms | ✓ 1690ms | http |
| 86.104.72.219:1081 | ✓ 631ms | ✓ 1786ms | ✓ 1709ms | ✓ 1538ms | ✓ 1017ms | http |
| 45.76.207.177:40000 | ✓ 1941ms | 否 | ✓ 1216ms | 否 | ✓ 1344ms | http |
| 82.114.228.67:1080 | ✓ 1259ms | 否 | ✓ 1572ms | ✓ 1768ms | 否 | http |
| 94.158.219.111:3128 | ✓ 998ms | ✓ 1812ms | ✓ 1321ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1149ms | ✓ 1459ms | ✓ 1232ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1960ms | ✓ 1419ms | ✓ 1412ms | 否 | 否 | http |
| 8.154.21.175:3128 | ✓ 1040ms | ✓ 1274ms | ✓ 1596ms | ✓ 1336ms | ✓ 1424ms | http |
| 154.64.232.35:8080 | ✓ 893ms | ✓ 1008ms | ✓ 1146ms | ✓ 1163ms | ✓ 912ms | http |
| 172.236.145.31:7890 | ✓ 1114ms | 否 | 否 | ✓ 1316ms | ✓ 1429ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1534ms | ✓ 1859ms | 否 | ✓ 1345ms | http |
| 112.209.22.22:8083 | ✓ 1612ms | 否 | ✓ 1675ms | ✓ 1679ms | ✓ 1654ms | http |
| 80.92.204.47:1081 | 否 | ✓ 1429ms | ✓ 1160ms | 否 | ✓ 1400ms | http |
| 129.226.81.110:7890 | ✓ 1976ms | ✓ 1858ms | 否 | ✓ 1841ms | ✓ 1721ms | http |
| 64.188.67.154:1080 | ✓ 1032ms | ✓ 1636ms | 否 | 否 | ✓ 1369ms | http |
| 86.104.74.110:1082 | ✓ 1337ms | ✓ 1484ms | ✓ 873ms | ✓ 1769ms | ✓ 1777ms | http |
| 86.104.74.110:1081 | ✓ 1343ms | ✓ 1703ms | ✓ 758ms | ✓ 1691ms | ✓ 1765ms | http |
| 130.61.174.200:1080 | ✓ 1340ms | ✓ 1844ms | 否 | 否 | ✓ 1453ms | http |
| 146.19.56.212:40002 | ✓ 314ms | ✓ 1478ms | ✓ 345ms | ✓ 1539ms | ✓ 1216ms | http |
| 209.126.10.139:3128 | 否 | ✓ 1109ms | ✓ 1466ms | ✓ 1219ms | ✓ 1022ms | http |
| 61.52.131.172:8443 | ✓ 1041ms | ✓ 1379ms | ✓ 1159ms | ✓ 1421ms | ✓ 1186ms | http |
| 34.71.229.255:3128 | ✓ 1405ms | 否 | ✓ 1510ms | 否 | ✓ 1976ms | http |
| 34.96.238.40:8080 | ✓ 1610ms | 否 | ✓ 1311ms | 否 | ✓ 1465ms | http |

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
