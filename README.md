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

最后更新：2026-03-09 03:30:13 UTC（2026-03-09 11:30:13 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1581ms | ✓ 1370ms | ✓ 1432ms | ✓ 1025ms | ✓ 820ms | http |
| 89.185.85.138:1080 | ✓ 496ms | 否 | 否 | ✓ 1601ms | ✓ 1865ms | http |
| 1.225.116.115:1080 | ✓ 987ms | ✓ 1588ms | ✓ 1255ms | ✓ 1683ms | ✓ 1101ms | http |
| 120.92.212.16:8890 | ✓ 1037ms | ✓ 1331ms | ✓ 1235ms | ✓ 1299ms | ✓ 1028ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1442ms | ✓ 1605ms | ✓ 1292ms | http |
| 121.128.121.54:3128 | ✓ 1554ms | ✓ 1266ms | ✓ 983ms | ✓ 1145ms | ✓ 892ms | http |
| 152.42.213.210:8080 | ✓ 818ms | 否 | ✓ 1564ms | ✓ 1418ms | ✓ 900ms | http |
| 202.155.12.161:443 | ✓ 1949ms | 否 | 否 | ✓ 1784ms | ✓ 1366ms | http |
| 162.248.165.72:1080 | ✓ 1510ms | 否 | ✓ 1984ms | 否 | ✓ 1899ms | http |
| 14.225.222.164:7890 | 否 | 否 | ✓ 1607ms | ✓ 1696ms | ✓ 1407ms | http |
| 115.231.181.40:8128 | ✓ 1023ms | ✓ 1584ms | ✓ 1037ms | 否 | ✓ 960ms | http |
| 165.227.5.10:8888 | ✓ 566ms | 否 | 否 | ✓ 1140ms | ✓ 789ms | http |
| 35.225.22.61:80 | ✓ 332ms | ✓ 1333ms | 否 | ✓ 1030ms | 否 | http |
| 194.213.18.200:443 | ✓ 1627ms | ✓ 1935ms | 否 | 否 | ✓ 1955ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1069ms | ✓ 1202ms | ✓ 1848ms | http |
| 14.56.107.244:3128 | ✓ 1576ms | ✓ 1287ms | ✓ 725ms | ✓ 1088ms | ✓ 896ms | http |
| 61.72.221.194:3128 | ✓ 1993ms | ✓ 1524ms | ✓ 670ms | ✓ 1097ms | ✓ 1122ms | http |
| 61.72.110.54:3128 | 否 | ✓ 1567ms | ✓ 1303ms | ✓ 1280ms | ✓ 921ms | http |
| 103.84.95.54:7890 | ✓ 1910ms | 否 | ✓ 770ms | ✓ 1665ms | ✓ 744ms | http |
| 125.128.12.14:3128 | ✓ 703ms | ✓ 1510ms | ✓ 890ms | ✓ 1075ms | ✓ 860ms | http |
| 61.72.221.94:3128 | ✓ 854ms | ✓ 1491ms | ✓ 1279ms | ✓ 1170ms | ✓ 1072ms | http |
| 178.236.245.17:3128 | ✓ 848ms | 否 | ✓ 1053ms | 否 | ✓ 1713ms | http |
| 81.70.169.194:80 | ✓ 1024ms | ✓ 1347ms | ✓ 1142ms | ✓ 1395ms | ✓ 1212ms | http |
| 101.43.255.96:80 | ✓ 1047ms | ✓ 1429ms | ✓ 1040ms | ✓ 1607ms | ✓ 1320ms | http |
| 178.236.245.59:3128 | ✓ 1274ms | 否 | ✓ 848ms | ✓ 1842ms | ✓ 1773ms | http |
| 125.128.12.144:3128 | ✓ 1541ms | ✓ 1533ms | ✓ 1528ms | ✓ 1678ms | ✓ 1318ms | http |
| 168.235.110.63:3128 | ✓ 1199ms | ✓ 1148ms | ✓ 1052ms | ✓ 1050ms | ✓ 979ms | http |
| 159.89.31.62:8080 | ✓ 969ms | ✓ 1707ms | ✓ 1485ms | ✓ 1679ms | ✓ 1407ms | http |
| 152.42.213.210:80 | ✓ 1473ms | 否 | ✓ 1197ms | ✓ 1224ms | ✓ 972ms | http |
| 103.215.36.88:18465 | ✓ 1004ms | ✓ 1292ms | ✓ 1057ms | ✓ 1472ms | ✓ 1026ms | http |
| 47.95.231.180:8084 | ✓ 1031ms | ✓ 1253ms | ✓ 940ms | ✓ 1276ms | ✓ 995ms | http |
| 160.250.4.13:1 | ✓ 1520ms | 否 | ✓ 1476ms | ✓ 1401ms | ✓ 1151ms | http |
| 194.87.43.49:8888 | ✓ 1575ms | 否 | ✓ 1576ms | 否 | ✓ 1982ms | http |
| 222.109.119.178:3128 | ✓ 775ms | ✓ 1221ms | ✓ 830ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 875ms | 否 | ✓ 1341ms | ✓ 1306ms | ✓ 1036ms | http |
| 5.101.0.233:3128 | ✓ 755ms | 否 | ✓ 1681ms | 否 | ✓ 1276ms | http |
| 205.209.118.30:3138 | ✓ 1136ms | 否 | ✓ 439ms | ✓ 1723ms | ✓ 1241ms | http |
| 14.56.177.44:3128 | ✓ 1178ms | ✓ 967ms | ✓ 1127ms | ✓ 1692ms | ✓ 969ms | http |
| 116.80.82.227:3172 | ✓ 1646ms | 否 | ✓ 1614ms | ✓ 1955ms | 否 | http |
| 138.124.53.25:7443 | ✓ 971ms | 否 | ✓ 1532ms | 否 | ✓ 1299ms | http |
| 121.230.8.158:1080 | ✓ 1745ms | ✓ 1886ms | ✓ 1445ms | 否 | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1329ms | ✓ 1495ms | ✓ 1003ms | http |
| 190.9.109.207:999 | ✓ 944ms | ✓ 1527ms | ✓ 1227ms | ✓ 1631ms | ✓ 1257ms | http |
| 190.9.109.198:999 | ✓ 877ms | ✓ 1555ms | ✓ 1152ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 992ms | ✓ 1873ms | ✓ 1148ms | ✓ 1521ms | ✓ 1117ms | http |
| 116.58.162.45:3128 | ✓ 1737ms | ✓ 1764ms | ✓ 1560ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 708ms | 否 | ✓ 1008ms | ✓ 1646ms | ✓ 1194ms | http |
| 39.104.201.40:7890 | ✓ 1840ms | ✓ 1318ms | ✓ 1005ms | ✓ 1315ms | ✓ 1868ms | http |
| 103.215.36.88:17565 | ✓ 1013ms | ✓ 1631ms | ✓ 1439ms | ✓ 1412ms | ✓ 1097ms | http |
| 187.216.141.46:3128 | ✓ 1091ms | 否 | ✓ 1323ms | ✓ 1390ms | 否 | http |
| 121.230.8.97:1080 | 否 | ✓ 1815ms | ✓ 1366ms | 否 | ✓ 1292ms | http |
| 121.230.9.241:1080 | ✓ 1289ms | ✓ 1487ms | ✓ 1644ms | 否 | ✓ 1480ms | http |
| 45.123.142.82:1111 | ✓ 1515ms | 否 | 否 | ✓ 1824ms | ✓ 1671ms | http |
| 46.39.105.157:8080 | 否 | ✓ 1846ms | ✓ 1736ms | ✓ 1429ms | 否 | http |
| 45.88.0.117:3128 | ✓ 995ms | 否 | ✓ 1037ms | ✓ 1947ms | ✓ 1414ms | http |
| 45.88.0.115:3128 | ✓ 1005ms | 否 | ✓ 1035ms | ✓ 1940ms | ✓ 1418ms | http |
| 88.80.150.82:8080 | ✓ 1597ms | ✓ 1939ms | 否 | 否 | ✓ 1944ms | https |
| 103.183.10.169:3125 | ✓ 1896ms | 否 | 否 | ✓ 1549ms | ✓ 1974ms | http |
| 45.88.0.98:3128 | ✓ 905ms | ✓ 1498ms | ✓ 1898ms | ✓ 1980ms | ✓ 1592ms | http |
| 45.88.0.111:3128 | ✓ 906ms | ✓ 1697ms | ✓ 1699ms | ✓ 1963ms | ✓ 1609ms | http |
| 213.220.62.62:3128 | ✓ 906ms | ✓ 1527ms | ✓ 1868ms | ✓ 1981ms | ✓ 1587ms | http |
| 45.88.0.113:3128 | ✓ 905ms | ✓ 1526ms | ✓ 1873ms | ✓ 1974ms | ✓ 1597ms | http |
| 45.88.0.116:3128 | ✓ 901ms | ✓ 1564ms | ✓ 1836ms | ✓ 1967ms | ✓ 1601ms | http |
| 45.88.0.114:3128 | ✓ 911ms | 否 | ✓ 1396ms | ✓ 1989ms | ✓ 1581ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1118ms | 否 | ✓ 1199ms | ✓ 895ms | http |
| 94.176.3.43:7443 | ✓ 1510ms | 否 | ✓ 1891ms | ✓ 1694ms | 否 | http |
| 116.80.82.223:3172 | ✓ 1719ms | 否 | ✓ 1567ms | 否 | ✓ 1764ms | http |
| 107.172.125.217:3128 | ✓ 888ms | 否 | ✓ 686ms | ✓ 819ms | ✓ 772ms | http |
| 113.177.131.2:3128 | ✓ 1574ms | 否 | ✓ 1071ms | ✓ 1171ms | ✓ 1149ms | http |
| 116.80.82.218:3172 | 否 | 否 | ✓ 1625ms | ✓ 1961ms | ✓ 1720ms | http |
| 20.210.39.153:8561 | ✓ 553ms | ✓ 1155ms | ✓ 811ms | ✓ 1095ms | ✓ 1012ms | http |
| 20.78.26.206:8561 | ✓ 548ms | ✓ 1477ms | ✓ 578ms | ✓ 1176ms | ✓ 954ms | http |
| 20.78.118.91:8561 | ✓ 554ms | ✓ 1928ms | ✓ 788ms | ✓ 1078ms | ✓ 865ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1496ms | ✓ 1714ms | ✓ 1520ms | http |
| 116.80.82.232:3172 | ✓ 1602ms | 否 | 否 | ✓ 1945ms | ✓ 1748ms | http |
| 172.212.68.37:3128 | ✓ 966ms | 否 | ✓ 545ms | ✓ 1653ms | 否 | http |
| 45.136.198.40:3128 | ✓ 799ms | 否 | 否 | ✓ 1930ms | ✓ 1775ms | http |
| 45.88.0.99:3128 | ✓ 884ms | ✓ 1576ms | ✓ 1038ms | 否 | 否 | http |
| 160.250.4.245:1 | ✓ 1795ms | 否 | ✓ 1481ms | ✓ 1390ms | ✓ 1037ms | http |
| 103.215.36.88:15842 | ✓ 1492ms | ✓ 1594ms | ✓ 1902ms | ✓ 1413ms | 否 | http |
| 177.93.33.55:999 | ✓ 1116ms | 否 | ✓ 1541ms | ✓ 1788ms | 否 | http |
| 20.120.225.109:3128 | ✓ 663ms | 否 | ✓ 1293ms | ✓ 1393ms | ✓ 817ms | http |
| 159.223.42.219:3128 | ✓ 817ms | 否 | ✓ 1396ms | ✓ 1761ms | ✓ 1020ms | http |
| 103.236.89.228:7890 | ✓ 1085ms | 否 | ✓ 1026ms | ✓ 1385ms | ✓ 1115ms | http |
| 138.124.93.82:1080 | ✓ 836ms | 否 | ✓ 1753ms | 否 | ✓ 1795ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1149ms | 否 | ✓ 1146ms | ✓ 870ms | http |
| 91.233.223.147:3128 | ✓ 1197ms | 否 | ✓ 1311ms | 否 | ✓ 1629ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1856ms | ✓ 1655ms | ✓ 1650ms | http |
| 180.130.80.196:9003 | 否 | ✓ 1555ms | ✓ 1429ms | ✓ 1524ms | ✓ 1162ms | http |
| 103.39.51.190:8080 | ✓ 1835ms | 否 | 否 | ✓ 1615ms | ✓ 1556ms | http |

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
