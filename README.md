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

最后更新：2026-02-22 22:19:00 UTC（2026-02-23 06:19:00 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 103.84.95.54:7890 | ✓ 1293ms | 否 | ✓ 986ms | ✓ 1042ms | 否 | http |
| 205.209.118.30:3138 | ✓ 275ms | ✓ 1806ms | 否 | ✓ 1077ms | ✓ 1078ms | http |
| 62.113.119.14:8080 | ✓ 1448ms | 否 | 否 | ✓ 1487ms | ✓ 1070ms | http |
| 168.235.110.63:3128 | ✓ 335ms | 否 | ✓ 831ms | 否 | ✓ 975ms | http |
| 43.130.6.42:80 | ✓ 697ms | 否 | ✓ 1154ms | ✓ 1694ms | 否 | http |
| 18.229.170.122:3128 | ✓ 1341ms | 否 | ✓ 587ms | 否 | ✓ 1504ms | http |
| 72.56.59.56:63127 | ✓ 1436ms | 否 | ✓ 1483ms | 否 | ✓ 1809ms | http |
| 72.56.59.62:63133 | ✓ 1450ms | 否 | ✓ 1514ms | 否 | ✓ 1813ms | http |
| 211.230.49.122:3128 | ✓ 879ms | ✓ 1933ms | 否 | ✓ 1584ms | ✓ 1387ms | http |
| 104.238.30.39:59741 | ✓ 1569ms | 否 | ✓ 1679ms | 否 | ✓ 1839ms | http |
| 104.238.30.68:63744 | ✓ 1572ms | 否 | ✓ 1643ms | 否 | ✓ 1871ms | http |
| 104.238.30.63:63744 | ✓ 1706ms | 否 | ✓ 1715ms | 否 | ✓ 1835ms | http |
| 104.238.30.58:63744 | ✓ 1580ms | 否 | ✓ 1647ms | 否 | ✓ 1839ms | http |
| 104.238.30.45:59741 | ✓ 1575ms | 否 | ✓ 1651ms | 否 | ✓ 1903ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1626ms | ✓ 1322ms | ✓ 1537ms | 否 | http |
| 51.81.46.174:3128 | ✓ 670ms | 否 | ✓ 1051ms | ✓ 883ms | ✓ 666ms | http |
| 138.124.53.25:7443 | ✓ 908ms | 否 | ✓ 1501ms | ✓ 1487ms | ✓ 1353ms | http |
| 188.166.208.168:9876 | ✓ 1054ms | 否 | ✓ 1199ms | ✓ 1295ms | ✓ 1027ms | http |
| 217.216.109.116:8080 | ✓ 1061ms | 否 | ✓ 1234ms | ✓ 1467ms | ✓ 1093ms | http |
| 202.152.44.19:8081 | ✓ 1504ms | 否 | ✓ 1366ms | ✓ 1423ms | ✓ 1149ms | http |
| 202.152.44.18:8081 | ✓ 1490ms | 否 | ✓ 1405ms | ✓ 1556ms | ✓ 1246ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1441ms | ✓ 1217ms | 否 | ✓ 1126ms | http |
| 104.238.30.38:59741 | ✓ 1532ms | 否 | ✓ 1615ms | 否 | ✓ 1871ms | http |
| 186.148.180.46:999 | ✓ 1605ms | ✓ 1798ms | ✓ 1315ms | ✓ 1614ms | ✓ 1224ms | http |
| 190.242.157.215:8080 | ✓ 1231ms | ✓ 1573ms | ✓ 1122ms | 否 | 否 | http |
| 104.238.30.86:63900 | ✓ 1560ms | 否 | ✓ 1611ms | 否 | ✓ 1903ms | http |
| 72.56.50.17:59787 | ✓ 1357ms | 否 | ✓ 1455ms | 否 | ✓ 1805ms | http |
| 72.56.59.23:61937 | ✓ 1427ms | 否 | ✓ 1589ms | 否 | ✓ 1770ms | http |
| 104.238.30.37:59741 | ✓ 1540ms | 否 | ✓ 1643ms | 否 | ✓ 1907ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1680ms | ✓ 1370ms | ✓ 1473ms | ✓ 1458ms | http |
| 104.238.30.91:63900 | ✓ 1560ms | 否 | ✓ 1614ms | 否 | ✓ 1870ms | http |
| 101.43.255.96:80 | ✓ 1279ms | ✓ 1546ms | ✓ 1197ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1300ms | ✓ 1582ms | ✓ 1152ms | 否 | 否 | http |
| 124.16.93.70:7890 | ✓ 1200ms | ✓ 1370ms | ✓ 1068ms | ✓ 1398ms | 否 | http |
| 8.219.97.248:80 | ✓ 1414ms | 否 | ✓ 1396ms | 否 | ✓ 1512ms | http |
| 5.101.0.233:3128 | ✓ 1032ms | ✓ 1775ms | ✓ 1592ms | 否 | ✓ 1769ms | http |
| 36.136.27.2:4999 | ✓ 1327ms | ✓ 1518ms | ✓ 1425ms | ✓ 1580ms | ✓ 1294ms | http |
| 120.46.152.136:3128 | ✓ 1233ms | ✓ 1563ms | ✓ 1581ms | ✓ 1586ms | ✓ 1290ms | http |
| 85.234.69.102:3128 | ✓ 1141ms | 否 | ✓ 1767ms | 否 | ✓ 1808ms | http |
| 59.127.212.110:4431 | ✓ 1850ms | ✓ 1587ms | ✓ 1725ms | ✓ 1516ms | ✓ 1124ms | http |
| 125.128.12.94:3128 | ✓ 1681ms | 否 | ✓ 1771ms | 否 | ✓ 1047ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1541ms | ✓ 1495ms | 否 | ✓ 1077ms | http |
| 217.216.109.116:80 | ✓ 1055ms | 否 | ✓ 1019ms | ✓ 1433ms | ✓ 1149ms | http |
| 47.101.159.19:8899 | ✓ 1022ms | ✓ 1287ms | ✓ 1036ms | ✓ 1299ms | ✓ 1119ms | http |
| 213.220.3.234:20573 | ✓ 995ms | ✓ 1343ms | ✓ 1471ms | ✓ 1816ms | ✓ 1504ms | http |
| 100.52.6.187:80 | ✓ 177ms | ✓ 1342ms | ✓ 764ms | ✓ 1557ms | ✓ 1110ms | http |
| 172.86.92.68:31337 | ✓ 1043ms | 否 | ✓ 1822ms | ✓ 1940ms | ✓ 1464ms | http |
| 137.220.150.22:6005 | ✓ 894ms | 否 | ✓ 1165ms | ✓ 1320ms | ✓ 1030ms | http |
| 132.145.93.138:1080 | ✓ 1321ms | 否 | ✓ 1767ms | ✓ 1888ms | ✓ 1559ms | http |
| 104.238.30.40:59741 | ✓ 1572ms | 否 | ✓ 1677ms | 否 | ✓ 1903ms | http |
| 81.177.48.54:2080 | 否 | 否 | ✓ 1895ms | ✓ 1668ms | ✓ 1360ms | http |
| 85.208.108.43:2094 | ✓ 775ms | 否 | ✓ 789ms | ✓ 1091ms | ✓ 917ms | http |
| 178.253.22.108:65431 | ✓ 561ms | ✓ 1811ms | ✓ 1133ms | ✓ 1642ms | ✓ 1055ms | http |
| 43.228.85.73:8888 | 否 | 否 | ✓ 1017ms | ✓ 1563ms | ✓ 1227ms | http |
| 14.56.177.162:3128 | ✓ 964ms | 否 | ✓ 1332ms | 否 | ✓ 1583ms | http |
| 72.56.59.17:61931 | ✓ 1929ms | 否 | ✓ 1619ms | 否 | ✓ 1867ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1971ms | 否 | ✓ 1022ms | ✓ 754ms | http |
| 78.13.231.158:3128 | ✓ 249ms | ✓ 1070ms | ✓ 308ms | ✓ 1780ms | ✓ 882ms | http |
| 36.147.78.166:80 | ✓ 1944ms | ✓ 1853ms | 否 | 否 | ✓ 1913ms | http |
| 160.238.65.5:3128 | ✓ 632ms | ✓ 1831ms | ✓ 1437ms | 否 | 否 | http |
| 160.238.65.2:3128 | ✓ 632ms | ✓ 1875ms | ✓ 1394ms | 否 | 否 | http |
| 160.238.65.3:3128 | ✓ 632ms | ✓ 1861ms | ✓ 1407ms | 否 | 否 | http |
| 160.238.65.8:3128 | ✓ 405ms | ✓ 1110ms | ✓ 491ms | ✓ 1269ms | ✓ 1019ms | http |
| 160.238.65.6:3128 | ✓ 403ms | ✓ 1137ms | ✓ 472ms | ✓ 1266ms | ✓ 1015ms | http |
| 103.236.64.247:8888 | ✓ 1441ms | 否 | ✓ 1330ms | ✓ 1487ms | 否 | http |
| 121.230.8.138:1080 | ✓ 1221ms | ✓ 1384ms | ✓ 1370ms | 否 | 否 | http |
| 91.238.104.171:2023 | ✓ 589ms | ✓ 1764ms | ✓ 1232ms | 否 | 否 | http |
| 35.72.254.71:3128 | 否 | ✓ 1387ms | ✓ 769ms | ✓ 1356ms | ✓ 1222ms | http |
| 18.215.106.202:80 | ✓ 315ms | 否 | ✓ 1724ms | ✓ 1494ms | ✓ 1215ms | http |
| 45.151.182.9:3128 | ✓ 1213ms | ✓ 1936ms | ✓ 1859ms | ✓ 1855ms | 否 | http |
| 45.174.77.224:999 | ✓ 810ms | ✓ 1422ms | ✓ 1106ms | ✓ 1327ms | ✓ 1166ms | http |
| 103.3.246.71:3128 | ✓ 1810ms | 否 | ✓ 1080ms | ✓ 1353ms | ✓ 1095ms | http |
| 43.155.156.147:3128 | 否 | ✓ 1558ms | ✓ 1822ms | ✓ 1722ms | ✓ 1668ms | http |
| 104.238.30.50:59741 | ✓ 1583ms | 否 | ✓ 1679ms | 否 | ✓ 1903ms | http |
| 217.217.254.94:8080 | ✓ 1044ms | 否 | ✓ 1664ms | ✓ 1385ms | 否 | http |
| 195.123.209.48:3128 | ✓ 967ms | ✓ 1771ms | ✓ 1764ms | 否 | ✓ 1617ms | http |
| 103.82.23.118:5182 | 否 | ✓ 1905ms | ✓ 1421ms | ✓ 1908ms | ✓ 1932ms | http |
| 135.125.97.184:43205 | ✓ 1092ms | ✓ 1885ms | ✓ 1776ms | 否 | ✓ 1985ms | http |
| 14.56.177.172:3128 | ✓ 1651ms | ✓ 1759ms | ✓ 1547ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 634ms | 否 | ✓ 861ms | ✓ 1077ms | 否 | http |
| 160.16.204.90:3128 | ✓ 799ms | ✓ 1092ms | 否 | ✓ 1166ms | ✓ 1115ms | http |
| 123.57.2.231:2020 | ✓ 1118ms | ✓ 1389ms | ✓ 1159ms | ✓ 1593ms | ✓ 1177ms | http |
| 91.238.105.64:2024 | ✓ 722ms | ✓ 1792ms | ✓ 1599ms | 否 | 否 | http |
| 34.96.238.40:8080 | ✓ 1397ms | 否 | ✓ 1299ms | ✓ 1749ms | 否 | http |
| 14.56.118.4:3128 | ✓ 1732ms | 否 | ✓ 1241ms | 否 | ✓ 1420ms | http |
| 45.140.147.82:1081 | ✓ 1020ms | ✓ 1968ms | ✓ 1110ms | 否 | ✓ 1803ms | http |
| 147.45.159.213:48206 | ✓ 1642ms | 否 | ✓ 954ms | 否 | ✓ 1211ms | http |
| 14.56.177.108:3128 | 否 | 否 | ✓ 965ms | ✓ 1875ms | ✓ 1141ms | http |
| 121.230.8.34:1080 | 否 | ✓ 1898ms | ✓ 1269ms | ✓ 1772ms | ✓ 1436ms | http |
| 113.45.250.180:443 | ✓ 1140ms | ✓ 1432ms | ✓ 1175ms | 否 | ✓ 1105ms | http |
| 61.109.216.213:8080 | ✓ 1836ms | 否 | ✓ 1429ms | ✓ 1800ms | ✓ 1739ms | http |
| 173.224.122.240:3128 | ✓ 1775ms | 否 | ✓ 1663ms | ✓ 1162ms | ✓ 1163ms | http |

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
