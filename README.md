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

最后更新：2026-05-08 08:45:13 UTC（2026-05-08 16:45:13 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | 否 | 否 | ✓ 855ms | ✓ 1168ms | ✓ 907ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 894ms | ✓ 1884ms | ✓ 925ms | http |
| 103.147.152.12:1080 | 否 | ✓ 1977ms | ✓ 1726ms | 否 | ✓ 1680ms | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 1482ms | ✓ 1803ms | ✓ 1468ms | http |
| 38.194.254.134:999 | ✓ 1150ms | ✓ 1662ms | ✓ 1492ms | ✓ 1829ms | ✓ 1503ms | http |
| 45.125.67.37:8443 | ✓ 1114ms | 否 | ✓ 841ms | ✓ 1043ms | ✓ 1084ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1303ms | ✓ 570ms | ✓ 1367ms | 否 | http |
| 77.110.107.80:1080 | ✓ 1059ms | 否 | ✓ 590ms | 否 | ✓ 1694ms | http |
| 77.110.107.80:8080 | ✓ 982ms | 否 | ✓ 591ms | 否 | ✓ 1715ms | http |
| 212.224.88.212:443 | ✓ 1090ms | 否 | ✓ 1036ms | ✓ 1981ms | ✓ 1521ms | http |
| 168.110.52.228:3128 | ✓ 1587ms | 否 | 否 | ✓ 1416ms | ✓ 1041ms | http |
| 185.221.237.57:8443 | ✓ 700ms | ✓ 1707ms | ✓ 1618ms | 否 | ✓ 1830ms | http |
| 120.92.108.86:7890 | ✓ 1601ms | 否 | ✓ 1681ms | 否 | ✓ 1510ms | http |
| 77.110.119.136:3128 | 否 | 否 | ✓ 499ms | ✓ 1339ms | ✓ 1333ms | http |
| 38.188.247.12:999 | ✓ 1257ms | ✓ 1648ms | ✓ 1336ms | ✓ 1794ms | ✓ 1352ms | http |
| 91.242.229.129:8092 | 否 | ✓ 1749ms | ✓ 1843ms | 否 | ✓ 1905ms | http |
| 45.153.231.229:8080 | ✓ 1486ms | ✓ 1890ms | ✓ 1271ms | 否 | 否 | http |
| 103.82.93.219:3128 | 否 | 否 | ✓ 1594ms | ✓ 1920ms | ✓ 1868ms | http |
| 43.133.44.89:8888 | ✓ 1446ms | ✓ 1664ms | ✓ 1319ms | ✓ 1466ms | 否 | http |
| 103.93.93.77:8050 | 否 | 否 | ✓ 1583ms | ✓ 1630ms | ✓ 1536ms | http |
| 190.12.150.244:999 | ✓ 1729ms | 否 | ✓ 987ms | ✓ 1752ms | ✓ 1536ms | http |
| 45.239.48.99:999 | 否 | ✓ 1744ms | ✓ 1438ms | 否 | ✓ 1974ms | http |
| 223.16.170.103:80 | ✓ 1063ms | 否 | ✓ 989ms | 否 | ✓ 1075ms | http |
| 185.125.100.115:40000 | ✓ 1942ms | 否 | ✓ 1205ms | ✓ 1778ms | 否 | http |
| 65.109.125.111:8443 | ✓ 786ms | ✓ 1812ms | ✓ 1028ms | 否 | 否 | http |
| 220.121.143.33:3128 | ✓ 1405ms | ✓ 1178ms | ✓ 1263ms | 否 | 否 | http |
| 43.156.132.113:3128 | ✓ 1389ms | 否 | ✓ 966ms | ✓ 1067ms | ✓ 1194ms | http |
| 45.59.122.132:80 | ✓ 1390ms | 否 | ✓ 1438ms | ✓ 1981ms | ✓ 1312ms | http |
| 148.230.4.241:999 | ✓ 649ms | ✓ 1562ms | ✓ 482ms | ✓ 1582ms | ✓ 1279ms | http |
| 103.158.242.58:83 | ✓ 1744ms | 否 | ✓ 1722ms | ✓ 1872ms | 否 | http |
| 8.154.21.175:3128 | ✓ 1104ms | ✓ 1498ms | ✓ 1678ms | ✓ 1078ms | 否 | http |
| 103.35.190.69:1082 | 否 | 否 | ✓ 1325ms | ✓ 1334ms | ✓ 1865ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1789ms | ✓ 1619ms | ✓ 1281ms | http |
| 121.230.9.54:1080 | 否 | ✓ 1493ms | ✓ 1114ms | ✓ 1422ms | 否 | http |
| 152.32.132.190:7890 | ✓ 1582ms | 否 | 否 | ✓ 1489ms | ✓ 1718ms | http |
| 45.140.147.155:1082 | ✓ 1369ms | 否 | ✓ 1240ms | ✓ 1847ms | 否 | http |
| 147.45.178.211:14658 | ✓ 1325ms | 否 | ✓ 1602ms | 否 | ✓ 1812ms | http |
| 84.47.150.125:1080 | ✓ 1165ms | 否 | ✓ 1897ms | 否 | ✓ 1991ms | http |
| 62.133.60.126:24558 | ✓ 818ms | 否 | ✓ 1334ms | 否 | ✓ 1481ms | http |
| 210.223.44.230:3128 | ✓ 1683ms | ✓ 1353ms | ✓ 897ms | 否 | ✓ 988ms | http |
| 79.137.205.44:40000 | 否 | 否 | ✓ 1073ms | ✓ 1999ms | ✓ 1289ms | http |
| 59.46.216.131:30001 | ✓ 1917ms | ✓ 1442ms | 否 | ✓ 1378ms | 否 | http |
| 8.209.238.110:47701 | ✓ 1324ms | 否 | ✓ 686ms | ✓ 812ms | ✓ 665ms | http |
| 194.59.247.34:10808 | ✓ 681ms | ✓ 1844ms | ✓ 1369ms | 否 | ✓ 1976ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1415ms | ✓ 1286ms | ✓ 1106ms | http |
| 120.92.211.211:7890 | 否 | ✓ 1311ms | ✓ 1893ms | ✓ 1544ms | ✓ 1941ms | http |
| 140.245.66.105:8081 | ✓ 1882ms | 否 | ✓ 1347ms | ✓ 1349ms | ✓ 883ms | http |
| 51.159.156.201:3128 | ✓ 1371ms | 否 | ✓ 1370ms | ✓ 1577ms | ✓ 1186ms | http |
| 1.231.81.166:3128 | ✓ 1579ms | ✓ 1320ms | 否 | ✓ 1148ms | ✓ 1145ms | http |
| 143.198.211.194:8080 | ✓ 1740ms | 否 | ✓ 1116ms | ✓ 1392ms | ✓ 1033ms | http |
| 94.131.118.129:1081 | ✓ 656ms | 否 | ✓ 553ms | ✓ 1722ms | ✓ 993ms | http |
| 45.140.147.155:1081 | ✓ 822ms | 否 | ✓ 1504ms | 否 | ✓ 1303ms | http |
| 95.217.103.18:4567 | ✓ 1828ms | 否 | ✓ 951ms | 否 | ✓ 1742ms | http |
| 103.157.200.126:3128 | ✓ 1775ms | 否 | ✓ 1417ms | 否 | ✓ 1505ms | http |
| 49.65.127.215:3128 | 否 | ✓ 1122ms | ✓ 902ms | ✓ 1156ms | ✓ 858ms | http |
| 138.197.68.35:4857 | 否 | ✓ 1275ms | ✓ 1228ms | ✓ 1737ms | 否 | http |
| 94.131.118.129:1082 | ✓ 1633ms | ✓ 1483ms | ✓ 839ms | ✓ 1932ms | ✓ 1258ms | http |
| 116.171.106.111:3443 | 否 | 否 | ✓ 1485ms | ✓ 1977ms | ✓ 1269ms | http |
| 150.107.140.238:3128 | ✓ 1611ms | 否 | ✓ 941ms | ✓ 1348ms | ✓ 1814ms | http |
| 8.219.97.248:80 | ✓ 1150ms | 否 | ✓ 944ms | 否 | ✓ 1207ms | http |
| 221.122.91.36:11273 | ✓ 928ms | ✓ 1197ms | ✓ 892ms | ✓ 1186ms | ✓ 1010ms | http |
| 3.101.133.120:80 | 否 | ✓ 1434ms | ✓ 692ms | ✓ 1036ms | ✓ 937ms | http |
| 20.78.213.56:80 | ✓ 1425ms | ✓ 1879ms | 否 | ✓ 1137ms | ✓ 998ms | http |
| 185.221.237.57:443 | ✓ 1318ms | 否 | ✓ 1258ms | 否 | ✓ 1780ms | http |
| 207.254.71.62:8088 | ✓ 1465ms | ✓ 1952ms | 否 | ✓ 1980ms | 否 | http |
| 62.113.119.14:8080 | 否 | ✓ 1627ms | ✓ 789ms | ✓ 1615ms | ✓ 1186ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1674ms | ✓ 1024ms | ✓ 1496ms | 否 | http |
| 38.211.245.34:999 | ✓ 1159ms | 否 | ✓ 1130ms | 否 | ✓ 1965ms | http |
| 38.211.245.18:999 | ✓ 1314ms | 否 | ✓ 1122ms | 否 | ✓ 1990ms | http |
| 103.35.190.69:1081 | ✓ 299ms | ✓ 1222ms | ✓ 516ms | 否 | ✓ 881ms | http |
| 110.172.28.217:3128 | ✓ 1599ms | 否 | ✓ 1351ms | ✓ 1221ms | ✓ 961ms | http |
| 35.194.4.51:3128 | ✓ 687ms | 否 | ✓ 1191ms | ✓ 1132ms | ✓ 1214ms | http |
| 213.220.3.38:3128 | ✓ 939ms | ✓ 1898ms | ✓ 1651ms | ✓ 1765ms | ✓ 1487ms | http |
| 137.184.0.30:3128 | ✓ 294ms | ✓ 811ms | ✓ 779ms | ✓ 730ms | ✓ 549ms | http |
| 20.118.221.52:3128 | ✓ 359ms | ✓ 1377ms | ✓ 1057ms | ✓ 938ms | ✓ 704ms | http |
| 154.90.48.209:9090 | 否 | 否 | ✓ 844ms | ✓ 1251ms | ✓ 997ms | http |
| 152.42.177.32:8888 | ✓ 981ms | 否 | ✓ 1614ms | ✓ 1273ms | ✓ 1291ms | http |
| 104.161.23.122:5075 | ✓ 1239ms | 否 | ✓ 1328ms | 否 | ✓ 1939ms | http |
| 159.89.31.62:8080 | ✓ 1663ms | 否 | ✓ 1931ms | 否 | ✓ 1406ms | http |
| 139.135.139.82:8082 | ✓ 1322ms | 否 | ✓ 1452ms | ✓ 1485ms | ✓ 1139ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 833ms | ✓ 1040ms | ✓ 1225ms | http |
| 158.160.215.167:8127 | ✓ 1346ms | 否 | ✓ 1932ms | 否 | ✓ 1802ms | http |
| 120.92.212.16:8890 | ✓ 1836ms | ✓ 1242ms | ✓ 1879ms | 否 | ✓ 1957ms | http |
| 120.92.212.16:7890 | ✓ 907ms | ✓ 1251ms | 否 | ✓ 1495ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1672ms | 否 | ✓ 1761ms | ✓ 1848ms | ✓ 1732ms | http |
| 158.160.215.167:8124 | ✓ 1975ms | 否 | ✓ 1727ms | 否 | ✓ 1902ms | http |

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
