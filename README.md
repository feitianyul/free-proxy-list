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

最后更新：2026-05-20 17:46:04 UTC（2026-05-21 01:46:04 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 57.129.144.178:40000 | ✓ 612ms | ✓ 1600ms | ✓ 1268ms | 否 | ✓ 1679ms | http |
| 89.58.50.94:11140 | ✓ 476ms | 否 | ✓ 1579ms | 否 | ✓ 1486ms | http |
| 1.231.81.166:3128 | ✓ 1910ms | ✓ 1402ms | ✓ 1240ms | ✓ 1136ms | ✓ 1014ms | http |
| 192.99.8.15:8850 | ✓ 619ms | 否 | ✓ 1388ms | 否 | ✓ 1364ms | http |
| 185.200.188.234:10001 | ✓ 1064ms | 否 | ✓ 1583ms | 否 | ✓ 1700ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1510ms | ✓ 1043ms | 否 | ✓ 1087ms | http |
| 152.70.91.193:40000 | ✓ 1933ms | 否 | 否 | ✓ 1684ms | ✓ 1960ms | http |
| 138.2.78.251:8100 | ✓ 1775ms | 否 | ✓ 1997ms | ✓ 1440ms | ✓ 1312ms | http |
| 174.137.134.182:2999 | 否 | ✓ 1717ms | ✓ 465ms | ✓ 904ms | ✓ 668ms | http |
| 43.130.126.146:6688 | ✓ 441ms | 否 | ✓ 890ms | ✓ 1156ms | 否 | http |
| 45.117.163.134:3128 | ✓ 1046ms | 否 | ✓ 1010ms | ✓ 1220ms | ✓ 971ms | http |
| 74.208.192.81:3129 | ✓ 1601ms | 否 | 否 | ✓ 1574ms | ✓ 1368ms | http |
| 168.110.52.228:3128 | ✓ 1549ms | 否 | ✓ 629ms | ✓ 961ms | ✓ 757ms | http |
| 5.252.33.13:2025 | ✓ 1402ms | 否 | ✓ 1227ms | 否 | ✓ 1774ms | http |
| 202.28.194.139:31280 | ✓ 1730ms | 否 | ✓ 1757ms | ✓ 1943ms | 否 | http |
| 85.192.29.60:3128 | 否 | 否 | ✓ 1001ms | ✓ 1942ms | ✓ 1566ms | http |
| 190.12.150.244:999 | ✓ 1819ms | ✓ 1605ms | ✓ 878ms | ✓ 1694ms | ✓ 1438ms | http |
| 38.250.126.225:999 | ✓ 1593ms | 否 | ✓ 637ms | 否 | ✓ 1951ms | http |
| 176.111.37.216:39811 | ✓ 570ms | ✓ 1654ms | ✓ 865ms | ✓ 1698ms | ✓ 1694ms | http |
| 176.111.37.5:39811 | 否 | 否 | ✓ 809ms | ✓ 1884ms | ✓ 1443ms | http |
| 129.80.217.21:444 | ✓ 1323ms | ✓ 919ms | ✓ 1197ms | ✓ 1266ms | ✓ 946ms | http |
| 8.154.21.175:3128 | ✓ 952ms | ✓ 1166ms | ✓ 998ms | ✓ 1272ms | ✓ 1033ms | http |
| 128.199.121.61:9090 | ✓ 1854ms | 否 | ✓ 1492ms | ✓ 1285ms | ✓ 974ms | http |
| 128.199.254.13:9090 | ✓ 1285ms | 否 | ✓ 1504ms | ✓ 1591ms | ✓ 1187ms | http |
| 148.230.4.241:999 | ✓ 1369ms | ✓ 1643ms | ✓ 638ms | ✓ 1884ms | ✓ 1547ms | http |
| 129.80.238.83:444 | ✓ 203ms | ✓ 973ms | ✓ 476ms | ✓ 962ms | ✓ 737ms | http |
| 152.42.170.187:9090 | ✓ 1602ms | 否 | ✓ 1391ms | ✓ 1263ms | 否 | http |
| 128.199.114.189:9090 | ✓ 1602ms | 否 | ✓ 1370ms | ✓ 1198ms | ✓ 971ms | http |
| 64.188.77.221:3128 | 否 | ✓ 1655ms | 否 | ✓ 1935ms | ✓ 1650ms | http |
| 170.106.136.181:31002 | 否 | ✓ 1891ms | ✓ 458ms | ✓ 809ms | ✓ 682ms | http |
| 31.172.78.12:3128 | ✓ 930ms | 否 | ✓ 781ms | ✓ 1618ms | ✓ 1538ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1147ms | ✓ 1941ms | 否 | ✓ 1897ms | http |
| 57.128.188.167:9147 | ✓ 1810ms | 否 | ✓ 1828ms | 否 | ✓ 1899ms | http |
| 147.45.186.28:3128 | ✓ 1179ms | ✓ 1782ms | ✓ 869ms | ✓ 1849ms | 否 | http |
| 207.254.71.62:8088 | ✓ 1049ms | ✓ 1601ms | ✓ 1569ms | 否 | ✓ 1746ms | http |
| 212.58.132.5:8888 | ✓ 1634ms | 否 | ✓ 1221ms | 否 | ✓ 1224ms | http |
| 114.214.165.78:10810 | ✓ 1350ms | 否 | ✓ 1489ms | ✓ 1572ms | ✓ 1910ms | http |
| 86.104.72.219:1082 | ✓ 718ms | ✓ 975ms | ✓ 1020ms | ✓ 1035ms | ✓ 1129ms | http |
| 121.230.9.184:1080 | ✓ 1220ms | ✓ 1339ms | ✓ 1192ms | ✓ 1589ms | ✓ 1353ms | http |
| 121.230.8.99:1080 | ✓ 1296ms | ✓ 1603ms | ✓ 1164ms | ✓ 1465ms | ✓ 1382ms | http |
| 121.230.9.161:1080 | ✓ 1349ms | 否 | ✓ 1276ms | ✓ 1710ms | ✓ 1756ms | http |
| 86.104.72.219:1081 | ✓ 718ms | ✓ 1612ms | ✓ 384ms | ✓ 1138ms | 否 | http |
| 38.224.56.215:999 | ✓ 1371ms | ✓ 1587ms | ✓ 1669ms | ✓ 1735ms | ✓ 1608ms | http |
| 190.60.54.221:999 | ✓ 1449ms | ✓ 1903ms | ✓ 1692ms | ✓ 1957ms | ✓ 1810ms | http |
| 64.188.77.26:3128 | ✓ 1540ms | ✓ 1869ms | ✓ 714ms | ✓ 1574ms | 否 | http |
| 104.248.151.93:9090 | ✓ 863ms | 否 | ✓ 889ms | ✓ 1237ms | ✓ 941ms | http |
| 128.199.113.85:9090 | ✓ 1379ms | 否 | ✓ 895ms | ✓ 1217ms | ✓ 1115ms | http |
| 3.101.133.120:80 | ✓ 385ms | ✓ 1423ms | ✓ 425ms | ✓ 1091ms | ✓ 1048ms | http |
| 223.16.170.103:80 | ✓ 1698ms | 否 | ✓ 1456ms | ✓ 1428ms | 否 | http |
| 105.159.134.130:4948 | ✓ 1204ms | ✓ 1868ms | 否 | 否 | ✓ 1987ms | http |
| 38.188.247.12:999 | 否 | 否 | ✓ 432ms | ✓ 1402ms | ✓ 1512ms | http |
| 38.210.201.104:999 | ✓ 388ms | 否 | ✓ 447ms | ✓ 1258ms | ✓ 982ms | http |
| 152.67.191.232:6800 | ✓ 1848ms | 否 | ✓ 1256ms | ✓ 1424ms | ✓ 1145ms | http |
| 146.190.80.158:9090 | ✓ 1602ms | 否 | 否 | ✓ 1624ms | ✓ 996ms | http |
| 180.130.80.196:9003 | ✓ 1647ms | ✓ 1744ms | ✓ 1543ms | ✓ 1576ms | ✓ 1230ms | http |
| 114.214.163.108:6789 | ✓ 1282ms | ✓ 1604ms | ✓ 1308ms | ✓ 1613ms | ✓ 1284ms | http |
| 34.87.80.221:30000 | 否 | ✓ 1497ms | ✓ 1239ms | ✓ 1276ms | 否 | http |
| 158.255.212.55:9005 | ✓ 1148ms | 否 | ✓ 1883ms | ✓ 1860ms | 否 | http |
| 158.255.212.55:9480 | ✓ 1147ms | 否 | ✓ 1878ms | ✓ 1855ms | 否 | http |
| 158.255.212.55:7497 | ✓ 1147ms | 否 | ✓ 1882ms | ✓ 1855ms | 否 | http |
| 158.255.212.55:3256 | ✓ 1151ms | 否 | ✓ 1879ms | ✓ 1861ms | 否 | http |
| 209.97.149.157:80 | ✓ 1515ms | ✓ 1217ms | ✓ 1472ms | ✓ 1457ms | ✓ 1205ms | http |
| 46.30.46.133:3128 | ✓ 1537ms | ✓ 1595ms | ✓ 889ms | ✓ 1688ms | ✓ 1409ms | http |
| 45.173.12.140:1994 | ✓ 1008ms | ✓ 1526ms | ✓ 985ms | 否 | 否 | http |
| 38.19.41.126:999 | 否 | ✓ 1882ms | ✓ 1533ms | ✓ 1739ms | ✓ 1796ms | http |
| 103.158.242.58:83 | ✓ 1789ms | 否 | 否 | ✓ 1982ms | ✓ 1841ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 1064ms | ✓ 1194ms | ✓ 1003ms | http |
| 152.32.132.190:7890 | ✓ 1819ms | ✓ 1343ms | 否 | 否 | ✓ 977ms | http |
| 182.53.202.208:8080 | ✓ 1287ms | 否 | ✓ 1432ms | ✓ 1737ms | 否 | http |
| 47.84.205.9:8100 | 否 | ✓ 1851ms | ✓ 920ms | 否 | ✓ 1137ms | http |
| 116.198.224.150:9000 | 否 | 否 | ✓ 1703ms | ✓ 1900ms | ✓ 1771ms | http |
| 116.63.160.98:8899 | ✓ 1166ms | ✓ 1343ms | ✓ 1081ms | ✓ 1433ms | ✓ 1139ms | http |
| 61.52.131.172:8443 | ✓ 961ms | ✓ 1310ms | ✓ 989ms | 否 | 否 | http |
| 104.247.51.76:3128 | 否 | 否 | ✓ 1051ms | ✓ 1322ms | ✓ 840ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1819ms | ✓ 1307ms | ✓ 1056ms | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1564ms | ✓ 1524ms | ✓ 1556ms | http |
| 45.186.6.104:3128 | ✓ 1683ms | ✓ 1546ms | ✓ 1719ms | 否 | 否 | http |

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
