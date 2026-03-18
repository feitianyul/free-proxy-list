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

最后更新：2026-03-18 12:39:33 UTC（2026-03-18 20:39:33 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1096ms | ✓ 1419ms | ✓ 1055ms | ✓ 1175ms | ✓ 1424ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1098ms | ✓ 1487ms | ✓ 1646ms | http |
| 113.160.132.26:8080 | ✓ 1743ms | 否 | ✓ 1237ms | ✓ 1499ms | ✓ 1327ms | http |
| 168.235.110.63:3128 | ✓ 528ms | ✓ 902ms | ✓ 668ms | ✓ 951ms | ✓ 1623ms | http |
| 35.225.22.61:80 | ✓ 940ms | ✓ 1809ms | 否 | ✓ 1266ms | ✓ 1034ms | http |
| 193.23.200.251:10808 | ✓ 1017ms | 否 | ✓ 733ms | ✓ 1989ms | ✓ 1480ms | http |
| 4.216.195.194:3128 | ✓ 725ms | 否 | 否 | ✓ 1422ms | ✓ 1017ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1394ms | ✓ 1491ms | ✓ 1196ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1715ms | 否 | ✓ 1772ms | ✓ 1419ms | http |
| 45.167.124.52:8080 | ✓ 526ms | 否 | ✓ 488ms | ✓ 1534ms | ✓ 1261ms | http |
| 185.191.236.162:3128 | ✓ 1600ms | 否 | ✓ 1009ms | ✓ 1504ms | ✓ 1025ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 1389ms | ✓ 1735ms | ✓ 1406ms | http |
| 162.240.154.26:3128 | ✓ 1705ms | ✓ 1936ms | 否 | ✓ 1518ms | ✓ 809ms | http |
| 147.161.239.240:8800 | ✓ 450ms | ✓ 1608ms | ✓ 789ms | ✓ 1169ms | ✓ 980ms | http |
| 38.145.218.163:8451 | ✓ 1065ms | 否 | ✓ 1737ms | ✓ 936ms | ✓ 1258ms | http |
| 178.236.245.17:3128 | ✓ 733ms | 否 | ✓ 1419ms | ✓ 1856ms | ✓ 1890ms | http |
| 219.117.204.211:7799 | ✓ 1343ms | 否 | ✓ 1790ms | ✓ 1525ms | ✓ 1024ms | http |
| 1.231.81.166:3128 | ✓ 1621ms | ✓ 1581ms | ✓ 1823ms | ✓ 1730ms | ✓ 1434ms | http |
| 101.43.127.100:8877 | ✓ 1158ms | ✓ 1391ms | ✓ 1079ms | ✓ 1704ms | 否 | http |
| 49.156.44.114:8080 | 否 | 否 | ✓ 1617ms | ✓ 1689ms | ✓ 1737ms | http |
| 178.236.245.59:3128 | ✓ 703ms | ✓ 1956ms | ✓ 747ms | ✓ 1615ms | 否 | http |
| 38.145.220.35:8446 | ✓ 1938ms | 否 | ✓ 1636ms | 否 | ✓ 1820ms | http |
| 137.220.150.170:6005 | ✓ 1263ms | 否 | ✓ 947ms | ✓ 1427ms | ✓ 1040ms | http |
| 210.223.44.230:3128 | ✓ 1648ms | 否 | ✓ 1056ms | ✓ 1299ms | ✓ 946ms | http |
| 137.220.150.104:6005 | ✓ 1227ms | 否 | ✓ 1182ms | ✓ 1897ms | 否 | http |
| 194.5.212.40:8080 | ✓ 647ms | 否 | ✓ 1322ms | ✓ 1729ms | ✓ 1251ms | http |
| 85.208.108.43:10808 | ✓ 324ms | 否 | ✓ 365ms | 否 | ✓ 799ms | http |
| 111.79.111.126:3128 | ✓ 1680ms | ✓ 1828ms | 否 | ✓ 1923ms | 否 | http |
| 45.136.131.59:8450 | ✓ 1311ms | 否 | ✓ 1038ms | ✓ 1765ms | 否 | http |
| 110.172.29.131:3128 | ✓ 1716ms | 否 | ✓ 998ms | ✓ 1412ms | ✓ 1115ms | http |
| 38.180.2.107:3128 | ✓ 1803ms | 否 | ✓ 1690ms | 否 | ✓ 1665ms | http |
| 116.80.96.108:3172 | ✓ 1670ms | 否 | ✓ 1701ms | 否 | ✓ 1783ms | http |
| 103.139.138.194:3128 | ✓ 1249ms | 否 | ✓ 1589ms | ✓ 1632ms | ✓ 1374ms | http |
| 115.231.181.40:8128 | ✓ 1116ms | ✓ 1309ms | ✓ 1185ms | 否 | 否 | http |
| 38.145.220.22:8451 | ✓ 943ms | 否 | ✓ 797ms | ✓ 891ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1093ms | 否 | ✓ 1207ms | ✓ 1668ms | ✓ 1624ms | http |
| 165.225.113.220:10884 | ✓ 1647ms | 否 | ✓ 1359ms | ✓ 1861ms | 否 | http |
| 165.225.113.220:11175 | ✓ 1640ms | 否 | ✓ 1329ms | ✓ 1852ms | 否 | http |
| 38.34.179.61:8445 | ✓ 1850ms | ✓ 1063ms | ✓ 864ms | ✓ 1276ms | ✓ 1397ms | http |
| 45.136.130.197:8452 | ✓ 1770ms | ✓ 879ms | ✓ 807ms | 否 | ✓ 1310ms | http |
| 165.225.113.220:10958 | ✓ 1366ms | 否 | ✓ 1110ms | ✓ 1612ms | 否 | http |
| 187.190.58.152:8081 | ✓ 1836ms | ✓ 1788ms | ✓ 1722ms | 否 | 否 | http |
| 150.249.255.91:3128 | ✓ 1231ms | 否 | 否 | ✓ 1004ms | ✓ 813ms | http |
| 133.242.138.34:8100 | ✓ 1587ms | 否 | ✓ 1657ms | ✓ 1833ms | ✓ 1025ms | http |
| 85.198.96.242:3128 | ✓ 1668ms | 否 | 否 | ✓ 1603ms | ✓ 1210ms | http |
| 165.225.113.220:11462 | ✓ 924ms | 否 | ✓ 913ms | ✓ 1266ms | ✓ 1006ms | http |
| 150.5.166.194:1080 | ✓ 972ms | 否 | ✓ 974ms | ✓ 1102ms | ✓ 885ms | http |
| 88.80.150.82:8080 | ✓ 1185ms | 否 | ✓ 1952ms | 否 | ✓ 1198ms | https |
| 59.46.216.131:30001 | ✓ 1948ms | 否 | ✓ 1873ms | 否 | ✓ 1999ms | http |
| 137.220.151.110:6005 | ✓ 1020ms | 否 | ✓ 936ms | ✓ 1297ms | 否 | http |
| 165.227.5.10:8888 | ✓ 1910ms | ✓ 1078ms | ✓ 1118ms | ✓ 1200ms | ✓ 1766ms | http |
| 160.238.65.4:3128 | ✓ 1162ms | 否 | ✓ 1083ms | ✓ 1874ms | ✓ 1379ms | http |
| 160.238.65.6:3128 | ✓ 1161ms | 否 | ✓ 1086ms | ✓ 1880ms | ✓ 1374ms | http |
| 160.238.65.7:3128 | ✓ 1163ms | 否 | ✓ 1082ms | ✓ 1883ms | ✓ 1381ms | http |
| 160.238.65.9:3128 | ✓ 1160ms | 否 | ✓ 1086ms | ✓ 1876ms | ✓ 1382ms | http |
| 160.238.65.3:3128 | ✓ 1162ms | 否 | ✓ 1086ms | ✓ 1880ms | ✓ 1381ms | http |
| 160.238.65.2:3128 | ✓ 1163ms | 否 | ✓ 1084ms | ✓ 1877ms | ✓ 1385ms | http |
| 160.238.65.5:3128 | ✓ 1160ms | 否 | ✓ 1094ms | ✓ 1868ms | ✓ 1383ms | http |
| 180.127.149.247:1080 | 否 | ✓ 1692ms | ✓ 1186ms | ✓ 1498ms | ✓ 1146ms | http |
| 38.145.220.33:8448 | ✓ 366ms | ✓ 1465ms | ✓ 475ms | ✓ 899ms | ✓ 698ms | http |
| 165.225.113.220:11396 | 否 | 否 | ✓ 1334ms | ✓ 1838ms | ✓ 1390ms | http |
| 38.145.203.32:8452 | ✓ 845ms | ✓ 1192ms | ✓ 1071ms | ✓ 1323ms | ✓ 1555ms | http |
| 160.238.65.8:3128 | ✓ 970ms | 否 | 否 | ✓ 1954ms | ✓ 1301ms | http |
| 8.219.97.248:80 | ✓ 1765ms | 否 | ✓ 1814ms | ✓ 1458ms | 否 | http |
| 186.116.148.52:8080 | ✓ 1092ms | ✓ 1529ms | 否 | ✓ 1750ms | 否 | http |
| 217.77.102.18:3128 | 否 | 否 | ✓ 967ms | ✓ 1686ms | ✓ 1374ms | http |
| 3.99.169.21:31177 | ✓ 1087ms | 否 | ✓ 871ms | ✓ 1988ms | ✓ 1604ms | http |
| 3.8.3.11:10031 | ✓ 1065ms | 否 | ✓ 1608ms | 否 | ✓ 1815ms | http |
| 161.35.158.192:3128 | 否 | 否 | ✓ 1403ms | ✓ 1606ms | ✓ 1781ms | http |
| 37.187.109.70:10111 | ✓ 895ms | 否 | ✓ 394ms | ✓ 1429ms | ✓ 1612ms | http |
| 165.225.113.220:11446 | ✓ 1070ms | 否 | ✓ 1146ms | ✓ 1560ms | 否 | http |
| 45.140.147.155:1081 | ✓ 488ms | 否 | ✓ 1142ms | ✓ 1240ms | 否 | http |
| 165.225.113.220:10018 | 否 | 否 | ✓ 1073ms | ✓ 1504ms | ✓ 1282ms | http |
| 185.241.5.57:3128 | ✓ 1517ms | 否 | ✓ 1297ms | 否 | ✓ 1901ms | http |
| 165.225.113.220:10819 | ✓ 1715ms | 否 | 否 | ✓ 1548ms | ✓ 1243ms | http |
| 165.225.113.220:10017 | ✓ 1743ms | 否 | 否 | ✓ 1551ms | ✓ 1254ms | http |

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
