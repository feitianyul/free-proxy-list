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

最后更新：2026-02-28 06:32:29 UTC（2026-02-28 14:32:29 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 1158ms | ✓ 1945ms | ✓ 1365ms | ✓ 1708ms | ✓ 1540ms | http |
| 205.209.118.30:3138 | ✓ 1169ms | 否 | ✓ 1304ms | ✓ 1580ms | ✓ 1133ms | http |
| 120.92.212.16:7890 | ✓ 1137ms | ✓ 1188ms | ✓ 968ms | ✓ 1195ms | ✓ 928ms | http |
| 147.45.216.148:1080 | ✓ 792ms | 否 | ✓ 1628ms | 否 | ✓ 1536ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1222ms | ✓ 1486ms | ✓ 1336ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1332ms | 否 | ✓ 1617ms | ✓ 1112ms | http |
| 14.56.107.244:3128 | ✓ 609ms | ✓ 1283ms | 否 | ✓ 1248ms | ✓ 1203ms | http |
| 210.223.44.230:3128 | ✓ 1115ms | 否 | ✓ 1811ms | 否 | ✓ 1656ms | http |
| 43.161.214.161:1081 | ✓ 1671ms | ✓ 1296ms | ✓ 850ms | 否 | 否 | http |
| 168.235.110.63:3128 | ✓ 520ms | 否 | ✓ 768ms | ✓ 1367ms | ✓ 1025ms | http |
| 36.147.78.166:80 | ✓ 1691ms | 否 | ✓ 1656ms | ✓ 1931ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1762ms | ✓ 1177ms | 否 | 否 | ✓ 1242ms | http |
| 101.43.255.96:80 | ✓ 1862ms | ✓ 1324ms | ✓ 1052ms | ✓ 1198ms | ✓ 1948ms | http |
| 81.70.169.194:80 | ✓ 1044ms | ✓ 1985ms | 否 | ✓ 1256ms | ✓ 1065ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 585ms | ✓ 1141ms | ✓ 1116ms | http |
| 45.88.0.98:3128 | ✓ 749ms | 否 | ✓ 1525ms | 否 | ✓ 1954ms | http |
| 45.88.0.114:3128 | ✓ 789ms | ✓ 1530ms | ✓ 1977ms | 否 | 否 | http |
| 61.72.110.54:3128 | 否 | ✓ 1706ms | ✓ 1912ms | ✓ 1453ms | 否 | http |
| 100.50.41.9:80 | ✓ 1291ms | 否 | ✓ 1940ms | ✓ 1675ms | ✓ 1564ms | http |
| 52.188.28.218:3128 | ✓ 315ms | 否 | ✓ 313ms | ✓ 1122ms | 否 | http |
| 103.104.99.89:80 | ✓ 1780ms | 否 | ✓ 1552ms | ✓ 1201ms | ✓ 1385ms | http |
| 61.72.110.24:3128 | 否 | ✓ 1999ms | ✓ 1828ms | 否 | ✓ 837ms | http |
| 167.103.6.46:11197 | ✓ 1719ms | 否 | ✓ 1156ms | ✓ 1705ms | 否 | http |
| 165.225.120.17:11995 | ✓ 1267ms | 否 | ✓ 1345ms | 否 | ✓ 1140ms | http |
| 165.225.120.17:10919 | ✓ 1266ms | 否 | ✓ 1345ms | 否 | ✓ 1156ms | http |
| 165.225.120.17:11070 | ✓ 1266ms | 否 | ✓ 1344ms | ✓ 1845ms | ✓ 1326ms | http |
| 165.225.120.17:12215 | ✓ 1266ms | 否 | ✓ 1400ms | 否 | ✓ 1124ms | http |
| 165.225.120.17:10906 | ✓ 1268ms | 否 | ✓ 1344ms | 否 | ✓ 1189ms | http |
| 165.225.120.17:11745 | ✓ 1267ms | 否 | ✓ 1345ms | 否 | ✓ 1189ms | http |
| 165.225.120.17:11702 | ✓ 1266ms | 否 | ✓ 1345ms | ✓ 1941ms | ✓ 1420ms | http |
| 115.231.181.40:8128 | ✓ 1093ms | ✓ 1083ms | ✓ 872ms | ✓ 1787ms | ✓ 1076ms | http |
| 138.124.53.25:7443 | ✓ 1091ms | 否 | ✓ 1428ms | 否 | ✓ 1715ms | http |
| 45.88.0.99:3128 | ✓ 1736ms | 否 | 否 | ✓ 1530ms | ✓ 1683ms | http |
| 120.202.127.234:10808 | ✓ 852ms | ✓ 1080ms | ✓ 869ms | ✓ 1098ms | ✓ 1469ms | http |
| 39.104.201.40:7890 | ✓ 868ms | ✓ 1151ms | ✓ 906ms | ✓ 1205ms | ✓ 908ms | http |
| 121.237.181.137:8888 | ✓ 1204ms | 否 | 否 | ✓ 1186ms | ✓ 882ms | http |
| 36.147.78.166:443 | 否 | ✓ 1518ms | ✓ 1608ms | ✓ 1594ms | ✓ 1561ms | http |
| 103.104.99.29:80 | ✓ 1487ms | 否 | ✓ 1491ms | ✓ 1568ms | ✓ 1258ms | http |
| 103.26.116.138:8080 | ✓ 1472ms | 否 | 否 | ✓ 1777ms | ✓ 1416ms | http |
| 45.140.147.155:1081 | ✓ 1908ms | ✓ 1837ms | ✓ 1474ms | 否 | 否 | http |
| 103.84.95.54:7890 | ✓ 902ms | 否 | ✓ 669ms | ✓ 849ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1398ms | 否 | 否 | ✓ 1718ms | ✓ 1225ms | http |
| 100.49.169.232:80 | 否 | ✓ 1960ms | ✓ 1119ms | ✓ 1648ms | ✓ 1480ms | http |
| 195.123.209.48:3128 | ✓ 894ms | 否 | ✓ 1776ms | ✓ 1762ms | 否 | http |
| 81.177.48.54:2080 | ✓ 1901ms | 否 | ✓ 1773ms | 否 | ✓ 1536ms | http |
| 103.236.64.247:8888 | ✓ 1305ms | ✓ 1981ms | 否 | 否 | ✓ 1632ms | http |
| 35.234.17.221:8080 | ✓ 853ms | ✓ 1236ms | 否 | 否 | ✓ 1429ms | http |
| 45.136.198.40:3128 | ✓ 1301ms | 否 | ✓ 1767ms | 否 | ✓ 1850ms | http |
| 136.226.254.24:9400 | 否 | 否 | ✓ 1526ms | ✓ 1836ms | ✓ 1542ms | http |
| 136.226.254.24:10517 | ✓ 871ms | 否 | ✓ 887ms | ✓ 1452ms | ✓ 1140ms | http |
| 193.181.35.5:8118 | ✓ 1011ms | 否 | ✓ 1141ms | 否 | ✓ 1996ms | http |
| 171.234.131.41:6620 | ✓ 1903ms | 否 | ✓ 1562ms | ✓ 1514ms | ✓ 1645ms | http |
| 103.39.51.190:8080 | ✓ 1632ms | 否 | ✓ 1162ms | 否 | ✓ 1876ms | http |
| 57.128.188.167:9196 | ✓ 1951ms | 否 | ✓ 1870ms | 否 | ✓ 1823ms | http |
| 85.208.108.43:2094 | ✓ 669ms | 否 | ✓ 1176ms | ✓ 1086ms | ✓ 892ms | http |
| 90.84.188.97:8000 | ✓ 1230ms | ✓ 1829ms | 否 | 否 | ✓ 1502ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1191ms | 否 | ✓ 1501ms | ✓ 1289ms | http |
| 94.177.131.12:3128 | ✓ 924ms | 否 | ✓ 1142ms | ✓ 909ms | ✓ 998ms | http |
| 103.131.19.42:8181 | ✓ 1349ms | 否 | ✓ 1698ms | ✓ 1594ms | ✓ 1382ms | http |
| 8.219.97.248:80 | ✓ 1414ms | 否 | ✓ 1751ms | ✓ 1498ms | 否 | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1193ms | ✓ 1796ms | ✓ 1261ms | http |
| 61.72.110.94:3128 | ✓ 1122ms | 否 | ✓ 717ms | ✓ 1030ms | 否 | http |

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
