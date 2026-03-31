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

最后更新：2026-03-31 16:00:47 UTC（2026-04-01 00:00:47 UTC+8）

**代理总数：74**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1540ms | 否 | ✓ 1030ms | ✓ 1220ms | ✓ 1206ms | http |
| 45.136.198.40:3128 | ✓ 1096ms | 否 | ✓ 1611ms | ✓ 1829ms | ✓ 1623ms | http |
| 95.213.217.168:52004 | ✓ 1110ms | ✓ 1501ms | 否 | 否 | ✓ 1565ms | http |
| 1.231.81.166:3128 | ✓ 1567ms | 否 | ✓ 1618ms | ✓ 1244ms | ✓ 1130ms | http |
| 167.103.115.102:8800 | ✓ 1355ms | 否 | ✓ 1822ms | ✓ 1501ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1264ms | 否 | ✓ 1276ms | ✓ 1376ms | ✓ 1521ms | http |
| 113.160.132.26:8080 | ✓ 1083ms | 否 | 否 | ✓ 1523ms | ✓ 1176ms | http |
| 167.103.144.127:8800 | ✓ 1472ms | 否 | ✓ 1595ms | ✓ 1836ms | ✓ 1823ms | http |
| 45.167.124.52:8080 | ✓ 1100ms | ✓ 1656ms | ✓ 515ms | ✓ 1550ms | 否 | http |
| 35.225.22.61:80 | ✓ 398ms | 否 | ✓ 945ms | ✓ 1559ms | ✓ 940ms | http |
| 208.87.243.199:7878 | ✓ 559ms | ✓ 1154ms | 否 | ✓ 1229ms | ✓ 719ms | http |
| 167.103.31.122:8800 | ✓ 1495ms | 否 | ✓ 1240ms | ✓ 1958ms | ✓ 1475ms | http |
| 195.19.217.200:3128 | ✓ 1451ms | 否 | ✓ 1437ms | 否 | ✓ 1938ms | http |
| 147.161.239.240:8800 | ✓ 1273ms | 否 | ✓ 1073ms | ✓ 1550ms | ✓ 1444ms | http |
| 133.242.138.34:8100 | ✓ 1304ms | 否 | 否 | ✓ 1859ms | ✓ 1450ms | http |
| 101.43.127.100:8877 | ✓ 1084ms | ✓ 1284ms | 否 | ✓ 1417ms | ✓ 1824ms | http |
| 120.92.212.16:8890 | ✓ 1709ms | 否 | ✓ 1167ms | ✓ 1552ms | 否 | http |
| 177.234.217.88:999 | ✓ 1334ms | 否 | ✓ 1666ms | 否 | ✓ 1854ms | http |
| 31.192.106.135:8010 | ✓ 878ms | ✓ 1896ms | 否 | 否 | ✓ 1712ms | http |
| 180.250.219.58:53281 | ✓ 1788ms | 否 | ✓ 1701ms | 否 | ✓ 1780ms | http |
| 101.47.73.135:3128 | ✓ 1747ms | 否 | ✓ 1121ms | ✓ 1479ms | 否 | http |
| 209.126.84.232:8888 | ✓ 537ms | ✓ 1622ms | ✓ 440ms | ✓ 1795ms | ✓ 1298ms | http |
| 120.92.212.16:7890 | ✓ 1460ms | 否 | ✓ 1415ms | ✓ 1683ms | 否 | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1854ms | ✓ 1805ms | ✓ 1211ms | http |
| 38.145.208.234:8453 | ✓ 1028ms | 否 | ✓ 1401ms | ✓ 1386ms | 否 | http |
| 103.113.70.189:1081 | ✓ 185ms | 否 | ✓ 1143ms | ✓ 936ms | ✓ 712ms | http |
| 43.99.54.236:5555 | ✓ 971ms | ✓ 1415ms | ✓ 925ms | ✓ 1063ms | ✓ 854ms | http |
| 103.180.123.229:8181 | 否 | 否 | ✓ 1577ms | ✓ 1750ms | ✓ 1676ms | http |
| 38.145.218.87:8445 | 否 | ✓ 1589ms | 否 | ✓ 1279ms | ✓ 966ms | http |
| 38.34.179.13:8445 | ✓ 480ms | 否 | ✓ 903ms | ✓ 913ms | ✓ 1264ms | http |
| 38.145.208.229:8446 | ✓ 763ms | ✓ 956ms | ✓ 1851ms | ✓ 1922ms | 否 | http |
| 185.76.241.128:10001 | ✓ 1220ms | 否 | ✓ 1531ms | 否 | ✓ 1934ms | http |
| 86.53.183.16:1080 | ✓ 759ms | ✓ 1545ms | ✓ 1505ms | 否 | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1113ms | ✓ 1515ms | ✓ 1126ms | http |
| 103.82.93.100:3128 | ✓ 1956ms | 否 | ✓ 1506ms | ✓ 1543ms | ✓ 1200ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1932ms | ✓ 1491ms | ✓ 1501ms | ✓ 1858ms | http |
| 103.184.99.194:8080 | 否 | 否 | ✓ 1607ms | ✓ 1915ms | ✓ 1787ms | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 762ms | ✓ 1510ms | ✓ 1720ms | http |
| 45.136.131.66:8446 | 否 | ✓ 856ms | ✓ 480ms | ✓ 1376ms | ✓ 1693ms | http |
| 38.145.208.235:8445 | ✓ 495ms | ✓ 1065ms | ✓ 284ms | ✓ 1117ms | ✓ 1907ms | http |
| 38.34.183.130:8451 | ✓ 543ms | ✓ 1878ms | ✓ 343ms | ✓ 1284ms | ✓ 1085ms | http |
| 38.145.220.40:8449 | ✓ 568ms | ✓ 1319ms | ✓ 730ms | ✓ 1053ms | ✓ 817ms | http |
| 38.145.220.43:8449 | ✓ 571ms | ✓ 1852ms | ✓ 282ms | ✓ 943ms | ✓ 831ms | http |
| 38.34.179.66:8446 | ✓ 597ms | ✓ 1936ms | ✓ 916ms | ✓ 1100ms | ✓ 895ms | http |
| 38.34.179.100:8452 | ✓ 432ms | ✓ 981ms | ✓ 1392ms | ✓ 1960ms | ✓ 780ms | http |
| 45.136.130.246:8446 | ✓ 462ms | ✓ 1951ms | 否 | ✓ 1033ms | ✓ 936ms | http |
| 38.145.208.224:8446 | ✓ 551ms | ✓ 1549ms | ✓ 1319ms | ✓ 1845ms | ✓ 921ms | http |
| 38.145.218.134:8445 | ✓ 587ms | 否 | ✓ 1714ms | ✓ 992ms | ✓ 971ms | http |
| 38.34.183.8:8448 | ✓ 609ms | ✓ 1091ms | 否 | ✓ 1624ms | ✓ 1534ms | http |
| 38.145.218.232:8448 | ✓ 759ms | ✓ 961ms | ✓ 1419ms | 否 | ✓ 943ms | http |
| 38.145.218.212:8448 | ✓ 762ms | ✓ 1603ms | ✓ 894ms | 否 | ✓ 1094ms | http |
| 38.34.179.19:8452 | ✓ 1499ms | ✓ 1121ms | ✓ 1625ms | 否 | ✓ 1756ms | http |
| 38.34.183.16:8445 | ✓ 1301ms | ✓ 1847ms | 否 | ✓ 1775ms | ✓ 1153ms | http |
| 45.136.131.37:8452 | ✓ 1772ms | ✓ 1888ms | ✓ 1165ms | 否 | ✓ 927ms | http |
| 45.136.131.43:8452 | ✓ 1779ms | 否 | ✓ 1357ms | 否 | ✓ 743ms | http |
| 217.217.249.160:80 | ✓ 1981ms | 否 | ✓ 1132ms | 否 | ✓ 1969ms | http |
| 38.145.203.45:8452 | ✓ 868ms | ✓ 1872ms | 否 | ✓ 931ms | ✓ 1403ms | http |
| 38.145.218.161:8444 | ✓ 590ms | ✓ 1186ms | ✓ 1366ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1137ms | 否 | 否 | ✓ 1606ms | ✓ 1600ms | http |
| 107.178.115.140:3128 | ✓ 534ms | 否 | ✓ 1070ms | 否 | ✓ 765ms | http |
| 104.248.81.109:3128 | ✓ 916ms | ✓ 1553ms | 否 | ✓ 1356ms | ✓ 1291ms | http |
| 205.164.46.6:3128 | ✓ 187ms | ✓ 951ms | ✓ 745ms | ✓ 838ms | ✓ 937ms | http |
| 65.108.203.37:28080 | ✓ 1414ms | 否 | ✓ 1766ms | 否 | ✓ 1759ms | http |
| 103.39.51.190:8080 | ✓ 1991ms | 否 | 否 | ✓ 1565ms | ✓ 1598ms | http |
| 217.217.249.160:8080 | ✓ 1711ms | 否 | ✓ 863ms | 否 | ✓ 1754ms | http |
| 192.71.213.85:9091 | ✓ 780ms | 否 | ✓ 619ms | ✓ 1907ms | 否 | http |
| 192.71.213.85:5678 | ✓ 1526ms | 否 | ✓ 1505ms | ✓ 1992ms | 否 | http |
| 205.164.46.6:3129 | ✓ 210ms | ✓ 924ms | ✓ 1144ms | ✓ 846ms | ✓ 953ms | http |
| 91.233.223.147:3128 | ✓ 1007ms | 否 | ✓ 1979ms | 否 | ✓ 1860ms | http |
| 61.52.131.172:8443 | ✓ 1118ms | ✓ 1454ms | ✓ 1134ms | ✓ 1453ms | ✓ 1120ms | http |
| 164.90.151.28:3128 | ✓ 515ms | ✓ 1105ms | ✓ 984ms | ✓ 996ms | ✓ 758ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 952ms | ✓ 1612ms | ✓ 1120ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1575ms | ✓ 1253ms | ✓ 1560ms | ✓ 1237ms | http |
| 150.241.71.15:1080 | 否 | ✓ 1848ms | ✓ 459ms | 否 | ✓ 1322ms | http |

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
