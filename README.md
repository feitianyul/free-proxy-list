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

最后更新：2026-06-06 12:57:35 UTC（2026-06-06 20:57:35 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 37.49.224.15:3128 | ✓ 1156ms | 否 | ✓ 1291ms | ✓ 1977ms | ✓ 1490ms | http |
| 185.200.188.234:10001 | ✓ 1194ms | 否 | ✓ 734ms | 否 | ✓ 1740ms | http |
| 13.201.56.14:3128 | ✓ 1572ms | ✓ 1592ms | ✓ 1999ms | ✓ 1704ms | ✓ 1346ms | http |
| 113.160.132.26:8080 | ✓ 1556ms | ✓ 1816ms | ✓ 1023ms | ✓ 1353ms | ✓ 1156ms | http |
| 202.28.194.139:31280 | ✓ 1913ms | 否 | 否 | ✓ 1870ms | ✓ 1857ms | http |
| 38.123.220.147:999 | 否 | 否 | ✓ 471ms | ✓ 1458ms | ✓ 1203ms | http |
| 116.104.252.1:2035 | ✓ 1728ms | 否 | ✓ 1438ms | ✓ 1682ms | 否 | http |
| 45.89.106.12:8080 | ✓ 1575ms | 否 | ✓ 1393ms | ✓ 1593ms | ✓ 1561ms | http |
| 85.234.100.149:8080 | ✓ 902ms | 否 | ✓ 1354ms | ✓ 1930ms | ✓ 1243ms | http |
| 85.234.100.149:1080 | ✓ 1762ms | 否 | ✓ 1084ms | ✓ 1570ms | ✓ 1309ms | http |
| 95.3.69.222:8080 | ✓ 1106ms | ✓ 1999ms | ✓ 1129ms | ✓ 1819ms | ✓ 1637ms | http |
| 116.104.252.1:2070 | ✓ 1392ms | 否 | ✓ 1426ms | ✓ 1641ms | ✓ 1396ms | http |
| 216.9.225.157:3128 | ✓ 582ms | 否 | ✓ 883ms | 否 | ✓ 1352ms | http |
| 2.27.50.150:8080 | ✓ 701ms | 否 | ✓ 1699ms | 否 | ✓ 1831ms | http |
| 59.66.37.236:6789 | ✓ 1132ms | ✓ 1445ms | 否 | ✓ 1447ms | 否 | http |
| 45.84.222.25:1080 | ✓ 566ms | 否 | ✓ 1407ms | ✓ 1577ms | ✓ 1504ms | http |
| 176.111.37.216:39811 | ✓ 1285ms | ✓ 1794ms | ✓ 1128ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 1504ms | ✓ 1381ms | ✓ 730ms | ✓ 1000ms | ✓ 770ms | http |
| 43.156.228.168:80 | ✓ 1265ms | ✓ 1798ms | ✓ 1280ms | 否 | 否 | http |
| 180.2.108.38:8080 | ✓ 1481ms | ✓ 1496ms | ✓ 1556ms | 否 | ✓ 723ms | http |
| 43.134.141.85:80 | ✓ 1285ms | ✓ 1898ms | ✓ 1764ms | ✓ 1343ms | ✓ 1416ms | http |
| 192.99.8.15:8850 | ✓ 664ms | 否 | ✓ 704ms | ✓ 1842ms | 否 | http |
| 39.100.88.235:3256 | ✓ 976ms | ✓ 1291ms | ✓ 1038ms | ✓ 1279ms | ✓ 1095ms | http |
| 2.26.3.66:8080 | 否 | ✓ 1971ms | ✓ 1484ms | 否 | ✓ 1316ms | http |
| 121.230.8.55:1080 | ✓ 1396ms | 否 | ✓ 1295ms | 否 | ✓ 1414ms | http |
| 165.227.133.230:8888 | ✓ 709ms | ✓ 1901ms | 否 | ✓ 1557ms | ✓ 1668ms | http |
| 160.238.65.9:3128 | ✓ 1581ms | 否 | ✓ 627ms | 否 | ✓ 1745ms | http |
| 147.45.179.108:1080 | ✓ 1263ms | 否 | ✓ 510ms | 否 | ✓ 1876ms | http |
| 129.153.7.7:60000 | 否 | ✓ 1718ms | ✓ 1217ms | 否 | ✓ 1761ms | http |
| 84.47.150.125:1080 | ✓ 648ms | 否 | ✓ 1908ms | 否 | ✓ 1340ms | http |
| 34.84.162.206:38080 | 否 | ✓ 1863ms | ✓ 1705ms | 否 | ✓ 1529ms | http |
| 170.106.136.181:31002 | ✓ 382ms | ✓ 733ms | ✓ 679ms | ✓ 749ms | ✓ 575ms | http |
| 216.236.30.14:7443 | ✓ 1587ms | 否 | 否 | ✓ 1975ms | ✓ 1576ms | http |
| 43.228.215.32:8080 | ✓ 1865ms | 否 | ✓ 1771ms | ✓ 1750ms | ✓ 1869ms | http |
| 2.26.87.216:1080 | ✓ 1870ms | 否 | ✓ 1498ms | 否 | ✓ 1905ms | http |
| 94.131.118.129:1081 | ✓ 1183ms | ✓ 1230ms | ✓ 1058ms | 否 | 否 | http |
| 160.238.65.5:3128 | ✓ 649ms | 否 | ✓ 1541ms | 否 | ✓ 1647ms | http |
| 185.141.26.131:3128 | ✓ 1358ms | 否 | ✓ 1396ms | 否 | ✓ 1738ms | http |
| 160.238.65.6:3128 | 否 | 否 | ✓ 1973ms | ✓ 1852ms | ✓ 1751ms | http |
| 8.216.132.206:20002 | ✓ 827ms | ✓ 1215ms | ✓ 821ms | ✓ 1159ms | 否 | http |
| 18.180.59.181:80 | ✓ 1526ms | ✓ 994ms | ✓ 1632ms | ✓ 982ms | ✓ 803ms | http |
| 104.154.186.48:80 | 否 | ✓ 1412ms | ✓ 998ms | ✓ 1665ms | ✓ 1369ms | http |
| 2.56.126.204:3128 | ✓ 1100ms | 否 | ✓ 635ms | ✓ 1917ms | ✓ 1208ms | http |
| 80.150.246.98:443 | ✓ 569ms | 否 | ✓ 1966ms | 否 | ✓ 1648ms | http |
| 59.127.212.110:4431 | ✓ 1088ms | ✓ 1688ms | 否 | ✓ 1639ms | ✓ 1347ms | http |
| 207.254.71.62:8088 | ✓ 973ms | 否 | ✓ 1537ms | 否 | ✓ 1848ms | http |
| 147.45.75.124:8080 | ✓ 1248ms | 否 | ✓ 1670ms | 否 | ✓ 1725ms | http |
| 59.66.29.199:6382 | ✓ 1323ms | ✓ 1537ms | 否 | 否 | ✓ 1876ms | http |
| 36.147.78.166:443 | ✓ 1763ms | ✓ 1734ms | 否 | 否 | ✓ 1847ms | http |
| 36.147.78.166:80 | ✓ 1810ms | ✓ 1737ms | ✓ 1811ms | ✓ 1695ms | ✓ 1736ms | http |
| 115.231.181.40:8128 | ✓ 1017ms | ✓ 1210ms | ✓ 962ms | ✓ 1935ms | ✓ 1006ms | http |
| 81.200.154.236:48503 | ✓ 1164ms | ✓ 1487ms | 否 | ✓ 1217ms | 否 | http |
| 206.135.55.224:999 | 否 | ✓ 1201ms | 否 | ✓ 1367ms | ✓ 1100ms | http |
| 162.222.206.167:8080 | ✓ 1842ms | ✓ 1904ms | ✓ 1432ms | 否 | ✓ 1174ms | http |
| 34.87.80.221:30000 | ✓ 1656ms | 否 | ✓ 1651ms | ✓ 1549ms | ✓ 1347ms | http |
| 116.206.141.213:8080 | ✓ 1731ms | 否 | ✓ 1372ms | ✓ 1760ms | ✓ 1443ms | http |
| 8.219.97.248:80 | ✓ 1569ms | 否 | ✓ 1249ms | ✓ 1547ms | 否 | http |
| 152.32.132.190:7890 | ✓ 732ms | ✓ 1002ms | ✓ 1379ms | 否 | 否 | http |
| 187.175.168.26:8080 | ✓ 1164ms | 否 | ✓ 1440ms | ✓ 1867ms | ✓ 1256ms | http |
| 158.255.212.55:3256 | ✓ 1179ms | 否 | ✓ 1586ms | ✓ 1944ms | 否 | http |
| 158.255.212.55:9480 | ✓ 1174ms | 否 | ✓ 1586ms | ✓ 1964ms | 否 | http |
| 158.255.212.55:7497 | ✓ 1173ms | 否 | ✓ 1585ms | ✓ 1975ms | 否 | http |
| 158.255.212.55:7839 | ✓ 1170ms | 否 | ✓ 1580ms | ✓ 1938ms | 否 | http |
| 158.255.212.55:9005 | ✓ 1178ms | 否 | ✓ 1579ms | ✓ 1974ms | 否 | http |
| 174.138.3.101:8443 | ✓ 1092ms | 否 | 否 | ✓ 1717ms | ✓ 1778ms | http |
| 58.187.104.56:2105 | ✓ 1621ms | 否 | ✓ 1684ms | ✓ 1791ms | 否 | http |
| 92.118.112.25:1082 | 否 | 否 | ✓ 1913ms | ✓ 1969ms | ✓ 1553ms | http |
| 8.154.21.175:3128 | ✓ 899ms | ✓ 1156ms | ✓ 953ms | ✓ 1210ms | ✓ 988ms | http |
| 116.254.118.180:80 | 否 | 否 | ✓ 1304ms | ✓ 1388ms | ✓ 1104ms | http |
| 43.161.239.147:8888 | 否 | 否 | ✓ 968ms | ✓ 1121ms | ✓ 1760ms | http |
| 122.233.222.209:2222 | ✓ 1132ms | ✓ 1317ms | 否 | ✓ 1194ms | ✓ 1308ms | http |
| 49.151.189.189:8082 | ✓ 1602ms | 否 | ✓ 1306ms | ✓ 1412ms | ✓ 1382ms | http |
| 191.53.189.10:8080 | ✓ 1356ms | 否 | ✓ 1101ms | 否 | ✓ 1942ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1517ms | 否 | ✓ 1316ms | ✓ 1679ms | http |
| 144.31.73.173:3128 | ✓ 644ms | 否 | ✓ 1600ms | 否 | ✓ 1652ms | http |
| 31.56.205.167:80 | 否 | ✓ 1677ms | ✓ 1749ms | ✓ 1493ms | 否 | http |
| 209.141.46.220:9091 | 否 | ✓ 881ms | ✓ 1552ms | ✓ 1954ms | 否 | http |
| 50.114.102.16:8888 | ✓ 534ms | ✓ 1818ms | ✓ 593ms | ✓ 1484ms | ✓ 1143ms | http |
| 169.40.6.114:3128 | ✓ 856ms | 否 | ✓ 1366ms | ✓ 1948ms | 否 | http |
| 116.104.252.1:2105 | ✓ 1416ms | 否 | ✓ 1392ms | ✓ 1568ms | 否 | http |
| 167.86.74.161:3128 | ✓ 1216ms | 否 | ✓ 1677ms | 否 | ✓ 1451ms | http |
| 61.52.131.172:8443 | ✓ 977ms | ✓ 1198ms | ✓ 1021ms | ✓ 1417ms | ✓ 1035ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1743ms | ✓ 1890ms | ✓ 1884ms | http |

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
