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

最后更新：2026-04-28 00:51:40 UTC（2026-04-28 08:51:40 UTC+8）

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
| 218.108.131.186:17890 | ✓ 873ms | ✓ 1136ms | ✓ 934ms | ✓ 1173ms | 否 | http |
| 43.153.149.53:1080 | ✓ 1719ms | 否 | 否 | ✓ 1524ms | ✓ 1455ms | http |
| 113.160.132.26:8080 | ✓ 1690ms | 否 | ✓ 1550ms | ✓ 1956ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1283ms | ✓ 1456ms | 否 | ✓ 1293ms | 否 | http |
| 77.110.119.136:3128 | ✓ 1519ms | 否 | ✓ 735ms | ✓ 1728ms | ✓ 1439ms | http |
| 62.113.119.14:8080 | ✓ 1226ms | ✓ 1723ms | ✓ 847ms | 否 | ✓ 1246ms | http |
| 46.101.95.183:8888 | ✓ 1297ms | 否 | ✓ 896ms | ✓ 1747ms | 否 | http |
| 8.211.166.184:8081 | ✓ 1480ms | ✓ 1155ms | ✓ 716ms | ✓ 931ms | 否 | http |
| 120.92.212.16:7890 | ✓ 962ms | ✓ 1743ms | ✓ 1552ms | 否 | 否 | http |
| 45.129.141.143:3128 | ✓ 1055ms | ✓ 1949ms | 否 | 否 | ✓ 1817ms | http |
| 103.157.200.126:3128 | ✓ 1686ms | 否 | ✓ 1682ms | ✓ 1919ms | 否 | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1462ms | ✓ 1748ms | ✓ 978ms | http |
| 47.110.42.192:9003 | ✓ 1760ms | ✓ 1765ms | ✓ 1347ms | ✓ 1717ms | ✓ 1701ms | http |
| 47.85.51.197:1080 | ✓ 180ms | ✓ 1643ms | ✓ 1527ms | 否 | 否 | http |
| 202.129.206.239:3128 | ✓ 1595ms | 否 | ✓ 1547ms | ✓ 1659ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1546ms | 否 | ✓ 1155ms | ✓ 1635ms | ✓ 971ms | http |
| 152.32.132.190:7890 | ✓ 1854ms | 否 | 否 | ✓ 1302ms | ✓ 1611ms | http |
| 210.223.44.230:3128 | ✓ 1620ms | ✓ 1559ms | ✓ 1887ms | 否 | 否 | http |
| 159.223.225.118:8888 | ✓ 851ms | 否 | 否 | ✓ 1789ms | ✓ 1385ms | http |
| 120.92.108.86:7890 | ✓ 1464ms | 否 | ✓ 1916ms | ✓ 1925ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1384ms | 否 | ✓ 1626ms | ✓ 1647ms | ✓ 1335ms | http |
| 86.104.72.220:1081 | ✓ 1082ms | 否 | ✓ 1684ms | 否 | ✓ 1475ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1337ms | ✓ 1011ms | 否 | ✓ 1071ms | http |
| 121.230.9.161:1080 | ✓ 1176ms | ✓ 1269ms | ✓ 1146ms | ✓ 1250ms | ✓ 1053ms | http |
| 190.248.145.114:999 | ✓ 1404ms | ✓ 1966ms | ✓ 1655ms | ✓ 1821ms | 否 | http |
| 143.198.223.214:1084 | ✓ 928ms | 否 | ✓ 1151ms | ✓ 1083ms | ✓ 865ms | http |
| 34.101.184.164:3128 | ✓ 1704ms | 否 | ✓ 1419ms | 否 | ✓ 1752ms | http |
| 124.156.179.148:3128 | ✓ 662ms | ✓ 930ms | ✓ 650ms | ✓ 844ms | ✓ 679ms | http |
| 168.110.52.228:3128 | ✓ 1356ms | ✓ 989ms | ✓ 698ms | ✓ 808ms | 否 | http |
| 146.56.164.121:3128 | ✓ 1412ms | 否 | ✓ 1087ms | ✓ 1221ms | 否 | http |
| 8.209.238.110:47701 | ✓ 509ms | 否 | ✓ 719ms | ✓ 875ms | ✓ 803ms | http |
| 223.205.72.13:8080 | ✓ 1511ms | 否 | ✓ 1708ms | ✓ 1586ms | 否 | http |
| 103.172.121.51:8080 | ✓ 1923ms | 否 | ✓ 1684ms | ✓ 1496ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1585ms | ✓ 1745ms | ✓ 1001ms | 否 | 否 | http |
| 183.232.248.73:7890 | ✓ 1311ms | ✓ 1193ms | ✓ 1001ms | ✓ 1484ms | 否 | http |
| 106.44.155.90:2222 | ✓ 1049ms | ✓ 1310ms | ✓ 1093ms | ✓ 1401ms | 否 | http |

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
