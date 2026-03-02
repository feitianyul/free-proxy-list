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

最后更新：2026-03-02 23:28:11 UTC（2026-03-03 07:28:11 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 91.99.99.83:9000 | ✓ 486ms | 否 | ✓ 1748ms | 否 | ✓ 1952ms | http |
| 205.209.118.30:3138 | 否 | ✓ 1951ms | ✓ 1686ms | ✓ 1183ms | 否 | http |
| 211.171.114.154:3128 | 否 | ✓ 1870ms | ✓ 1731ms | ✓ 1567ms | ✓ 1565ms | http |
| 2.56.178.131:443 | ✓ 1076ms | 否 | ✓ 1457ms | 否 | ✓ 1851ms | http |
| 91.238.104.172:2024 | 否 | ✓ 1847ms | ✓ 1083ms | ✓ 1719ms | ✓ 1341ms | http |
| 35.225.22.61:80 | ✓ 267ms | ✓ 1073ms | ✓ 983ms | ✓ 985ms | ✓ 885ms | http |
| 210.223.44.230:3128 | ✓ 1648ms | 否 | 否 | ✓ 1778ms | ✓ 1587ms | http |
| 103.239.201.50:1 | 否 | 否 | ✓ 1618ms | ✓ 1519ms | ✓ 1534ms | http |
| 142.171.85.32:1080 | ✓ 852ms | ✓ 1359ms | ✓ 1778ms | ✓ 1170ms | ✓ 1178ms | http |
| 61.72.110.94:3128 | ✓ 914ms | 否 | 否 | ✓ 1640ms | ✓ 1696ms | http |
| 160.238.65.6:3128 | ✓ 907ms | ✓ 1844ms | ✓ 763ms | 否 | ✓ 1025ms | http |
| 160.238.65.2:3128 | ✓ 903ms | ✓ 1159ms | ✓ 462ms | ✓ 1315ms | ✓ 1001ms | http |
| 160.238.65.4:3128 | ✓ 907ms | 否 | ✓ 609ms | 否 | ✓ 1010ms | http |
| 81.70.169.194:80 | ✓ 1170ms | ✓ 1316ms | ✓ 1102ms | ✓ 1333ms | ✓ 1137ms | http |
| 101.43.255.96:80 | ✓ 1006ms | ✓ 1438ms | ✓ 1111ms | ✓ 1390ms | ✓ 1064ms | http |
| 14.56.107.244:3128 | ✓ 1132ms | ✓ 1630ms | ✓ 1074ms | ✓ 1199ms | 否 | http |
| 91.238.104.171:2023 | ✓ 708ms | ✓ 1539ms | ✓ 1551ms | 否 | 否 | http |
| 160.238.65.5:3128 | ✓ 907ms | ✓ 1962ms | ✓ 1474ms | ✓ 1304ms | ✓ 1031ms | http |
| 160.238.65.8:3128 | ✓ 1120ms | 否 | ✓ 1471ms | ✓ 1289ms | ✓ 1050ms | http |
| 160.238.65.9:3128 | ✓ 903ms | ✓ 1331ms | ✓ 1278ms | 否 | ✓ 1013ms | http |
| 118.113.247.73:1080 | ✓ 1374ms | ✓ 1730ms | ✓ 1823ms | 否 | ✓ 1314ms | http |
| 61.72.221.194:3128 | 否 | ✓ 1314ms | ✓ 1083ms | 否 | ✓ 1485ms | http |
| 45.88.0.117:3128 | 否 | 否 | ✓ 1541ms | ✓ 1371ms | ✓ 1338ms | http |
| 115.76.5.32:10007 | 否 | 否 | ✓ 1680ms | ✓ 1982ms | ✓ 1732ms | http |
| 115.231.181.40:8128 | ✓ 1046ms | ✓ 1294ms | ✓ 1003ms | 否 | 否 | http |
| 125.128.12.144:3128 | ✓ 1816ms | 否 | ✓ 1467ms | 否 | ✓ 1386ms | http |
| 45.88.0.114:3128 | ✓ 1808ms | 否 | ✓ 462ms | 否 | ✓ 1661ms | http |
| 166.0.192.117:8888 | ✓ 1391ms | 否 | ✓ 1258ms | ✓ 1417ms | ✓ 771ms | http |
| 103.84.95.54:7890 | ✓ 1053ms | 否 | ✓ 982ms | 否 | ✓ 1074ms | http |
| 121.128.121.54:3128 | ✓ 1372ms | ✓ 1285ms | 否 | 否 | ✓ 854ms | http |
| 116.80.63.64:7777 | ✓ 1931ms | 否 | 否 | ✓ 1976ms | ✓ 1792ms | http |
| 51.79.71.106:8080 | ✓ 711ms | 否 | 否 | ✓ 1338ms | ✓ 1176ms | http |
| 121.230.9.113:1080 | ✓ 1335ms | ✓ 1751ms | ✓ 1709ms | ✓ 1856ms | ✓ 1326ms | http |
| 45.88.0.116:3128 | ✓ 547ms | ✓ 1392ms | 否 | 否 | ✓ 1373ms | http |
| 181.78.49.177:999 | ✓ 881ms | ✓ 1796ms | ✓ 765ms | ✓ 1753ms | 否 | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1531ms | ✓ 1857ms | ✓ 1217ms | http |
| 61.72.221.94:3128 | ✓ 1828ms | ✓ 1934ms | ✓ 1310ms | 否 | 否 | http |
| 45.88.0.111:3128 | ✓ 1509ms | 否 | ✓ 998ms | 否 | ✓ 1351ms | http |
| 138.124.53.25:7443 | ✓ 1202ms | ✓ 1983ms | 否 | 否 | ✓ 1547ms | http |
| 120.92.212.16:7890 | ✓ 1129ms | 否 | ✓ 1362ms | 否 | ✓ 1867ms | http |
| 35.234.17.221:8080 | ✓ 939ms | ✓ 1609ms | 否 | ✓ 1437ms | 否 | http |
| 171.234.62.116:10005 | ✓ 1822ms | 否 | ✓ 1629ms | 否 | ✓ 1897ms | http |
| 45.88.0.115:3128 | ✓ 1846ms | 否 | ✓ 1523ms | 否 | ✓ 1951ms | http |
| 58.220.95.11:12421 | ✓ 1030ms | ✓ 1309ms | ✓ 987ms | ✓ 1282ms | ✓ 1057ms | http |
| 158.160.215.167:8127 | ✓ 1059ms | 否 | ✓ 978ms | ✓ 1722ms | ✓ 1267ms | http |
| 120.92.212.16:8890 | ✓ 1091ms | ✓ 1697ms | 否 | 否 | ✓ 1096ms | http |
| 61.72.110.54:3128 | ✓ 939ms | 否 | ✓ 813ms | 否 | ✓ 875ms | http |
| 150.107.140.238:3128 | ✓ 1809ms | 否 | ✓ 1671ms | 否 | ✓ 1003ms | http |
| 121.230.8.49:1080 | ✓ 1466ms | ✓ 1421ms | ✓ 1473ms | ✓ 1497ms | ✓ 1731ms | http |
| 1.225.116.115:1080 | ✓ 1740ms | 否 | ✓ 1478ms | ✓ 1517ms | ✓ 1341ms | http |
| 121.230.9.101:1080 | ✓ 1313ms | ✓ 1873ms | ✓ 1408ms | ✓ 1852ms | ✓ 1612ms | http |
| 39.98.86.246:8118 | 否 | ✓ 1338ms | ✓ 1428ms | 否 | ✓ 1106ms | http |
| 121.230.8.109:1080 | ✓ 1111ms | ✓ 1688ms | ✓ 1358ms | ✓ 1542ms | ✓ 1346ms | http |
| 58.220.95.12:11904 | ✓ 1940ms | ✓ 1195ms | ✓ 985ms | ✓ 1252ms | ✓ 1010ms | http |
| 94.177.131.12:3128 | ✓ 577ms | ✓ 1850ms | ✓ 1462ms | ✓ 1335ms | ✓ 1738ms | http |
| 61.72.221.234:3128 | ✓ 1726ms | 否 | ✓ 1405ms | 否 | ✓ 1965ms | http |
| 121.230.8.136:1080 | ✓ 1272ms | ✓ 1669ms | ✓ 1319ms | ✓ 1637ms | ✓ 1519ms | http |
| 195.123.209.48:3128 | ✓ 705ms | 否 | ✓ 1706ms | 否 | ✓ 1758ms | http |
| 45.125.67.37:8443 | ✓ 1118ms | 否 | 否 | ✓ 1274ms | ✓ 1086ms | http |
| 221.127.195.224:8888 | 否 | 否 | ✓ 1271ms | ✓ 1506ms | ✓ 1361ms | http |
| 209.38.51.97:3128 | ✓ 414ms | 否 | ✓ 1207ms | ✓ 1049ms | 否 | http |
| 91.233.223.147:3128 | ✓ 978ms | 否 | ✓ 948ms | ✓ 1904ms | ✓ 1465ms | http |
| 45.136.198.40:3128 | ✓ 1181ms | ✓ 1794ms | ✓ 1984ms | ✓ 1615ms | ✓ 1542ms | http |
| 45.129.141.143:3128 | ✓ 1169ms | 否 | ✓ 1779ms | ✓ 1889ms | ✓ 1762ms | http |
| 47.110.42.192:9003 | ✓ 1626ms | ✓ 1416ms | ✓ 1565ms | ✓ 1754ms | ✓ 1729ms | http |
| 194.28.224.99:9443 | ✓ 597ms | 否 | ✓ 1231ms | 否 | ✓ 1819ms | http |
| 111.79.111.126:3128 | ✓ 1946ms | ✓ 1831ms | ✓ 1985ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1884ms | 否 | ✓ 1388ms | ✓ 1560ms | 否 | http |
| 171.234.62.116:10002 | ✓ 1635ms | 否 | ✓ 1838ms | 否 | ✓ 1919ms | http |
| 5.75.196.26:40000 | ✓ 474ms | 否 | 否 | ✓ 1842ms | ✓ 1593ms | http |
| 115.76.5.32:10005 | ✓ 1798ms | 否 | ✓ 1814ms | ✓ 1629ms | ✓ 1500ms | http |
| 47.101.149.27:9010 | ✓ 1487ms | ✓ 1409ms | ✓ 1546ms | ✓ 1621ms | 否 | http |
| 123.20.24.166:8118 | ✓ 1473ms | 否 | ✓ 1375ms | 否 | ✓ 1265ms | http |
| 144.31.69.170:1080 | 否 | 否 | ✓ 1617ms | ✓ 1911ms | ✓ 1674ms | http |
| 45.140.147.82:1082 | ✓ 608ms | 否 | ✓ 1197ms | ✓ 1954ms | 否 | http |
| 103.131.19.42:8181 | ✓ 1592ms | 否 | ✓ 1860ms | ✓ 1858ms | ✓ 1665ms | http |
| 74.208.234.198:443 | ✓ 829ms | ✓ 1486ms | 否 | 否 | ✓ 1766ms | http |
| 115.76.5.32:10010 | 否 | 否 | ✓ 1488ms | ✓ 1772ms | ✓ 1493ms | http |

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
