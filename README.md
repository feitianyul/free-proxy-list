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

最后更新：2026-05-21 09:23:13 UTC（2026-05-21 17:23:13 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.130.126.146:6688 | ✓ 1618ms | 否 | ✓ 792ms | ✓ 1325ms | 否 | http |
| 174.137.134.182:2999 | ✓ 1977ms | ✓ 1092ms | ✓ 858ms | ✓ 1263ms | ✓ 923ms | http |
| 192.99.8.15:8850 | ✓ 1245ms | 否 | ✓ 946ms | ✓ 1009ms | ✓ 1016ms | http |
| 81.30.156.115:8080 | ✓ 946ms | ✓ 1291ms | ✓ 1779ms | ✓ 1766ms | ✓ 1505ms | http |
| 89.58.50.94:11140 | ✓ 951ms | 否 | ✓ 1366ms | ✓ 1906ms | ✓ 1499ms | http |
| 167.86.95.198:3128 | 否 | ✓ 1887ms | ✓ 906ms | 否 | ✓ 1474ms | http |
| 65.109.190.168:8080 | ✓ 976ms | 否 | ✓ 1542ms | 否 | ✓ 1672ms | http |
| 188.253.125.38:28798 | ✓ 1218ms | 否 | ✓ 1471ms | ✓ 1328ms | ✓ 1079ms | http |
| 159.65.5.53:8080 | ✓ 1122ms | 否 | ✓ 1570ms | ✓ 1402ms | ✓ 1107ms | http |
| 176.111.37.5:39811 | 否 | ✓ 1558ms | ✓ 1563ms | 否 | ✓ 1938ms | http |
| 1.231.81.166:3128 | ✓ 1175ms | 否 | ✓ 1672ms | ✓ 1374ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1826ms | 否 | ✓ 1236ms | ✓ 1605ms | ✓ 1340ms | http |
| 116.171.106.15:3443 | 否 | ✓ 1725ms | ✓ 1591ms | ✓ 1853ms | 否 | http |
| 82.22.175.77:8080 | ✓ 1483ms | ✓ 1449ms | ✓ 1232ms | ✓ 1841ms | ✓ 1530ms | http |
| 138.2.92.70:8100 | ✓ 1123ms | 否 | ✓ 1810ms | ✓ 1457ms | ✓ 1293ms | http |
| 138.2.90.82:8100 | ✓ 1231ms | 否 | ✓ 1750ms | ✓ 1563ms | ✓ 1308ms | http |
| 138.2.78.251:8100 | ✓ 1130ms | 否 | ✓ 1759ms | ✓ 1444ms | ✓ 1356ms | http |
| 168.138.171.204:8100 | ✓ 1392ms | 否 | ✓ 1620ms | ✓ 1600ms | ✓ 1577ms | http |
| 103.163.132.178:3128 | ✓ 455ms | ✓ 1473ms | ✓ 414ms | ✓ 1466ms | ✓ 1123ms | http |
| 176.111.37.216:39811 | ✓ 457ms | ✓ 1501ms | ✓ 407ms | ✓ 1555ms | ✓ 1187ms | http |
| 45.117.163.134:3128 | ✓ 999ms | 否 | ✓ 1029ms | ✓ 1362ms | ✓ 1066ms | http |
| 185.200.188.234:10001 | ✓ 985ms | 否 | ✓ 1148ms | 否 | ✓ 1555ms | http |
| 202.28.194.139:31280 | ✓ 1784ms | 否 | ✓ 1745ms | 否 | ✓ 1954ms | http |
| 149.126.168.216:3128 | ✓ 1798ms | ✓ 1734ms | ✓ 627ms | 否 | 否 | http |
| 144.124.227.88:3128 | ✓ 654ms | 否 | ✓ 1101ms | ✓ 1992ms | ✓ 1728ms | http |
| 181.119.97.24:999 | ✓ 1650ms | ✓ 1876ms | ✓ 1239ms | ✓ 1783ms | 否 | http |
| 167.61.202.55:3128 | ✓ 865ms | 否 | ✓ 717ms | 否 | ✓ 1585ms | http |
| 45.8.229.228:8080 | ✓ 1092ms | ✓ 1689ms | ✓ 1765ms | 否 | 否 | http |
| 8.212.167.186:8080 | ✓ 1607ms | 否 | ✓ 1193ms | 否 | ✓ 1357ms | http |
| 74.208.192.81:3129 | ✓ 1008ms | ✓ 840ms | ✓ 1208ms | ✓ 1467ms | ✓ 1503ms | http |
| 152.67.191.232:6800 | ✓ 1135ms | 否 | ✓ 1111ms | ✓ 1391ms | 否 | http |
| 207.254.71.62:8088 | ✓ 570ms | 否 | ✓ 1652ms | 否 | ✓ 1598ms | http |
| 210.223.44.230:3128 | ✓ 1182ms | ✓ 1147ms | ✓ 872ms | 否 | 否 | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 930ms | ✓ 1356ms | ✓ 1063ms | http |
| 128.199.116.219:9090 | ✓ 1540ms | 否 | 否 | ✓ 1404ms | ✓ 1050ms | http |
| 46.30.46.133:3128 | ✓ 671ms | ✓ 1153ms | 否 | ✓ 1957ms | ✓ 1667ms | http |
| 128.199.113.85:9090 | ✓ 1532ms | 否 | ✓ 946ms | ✓ 1372ms | ✓ 1048ms | http |
| 128.199.121.61:9090 | ✓ 1530ms | 否 | ✓ 961ms | ✓ 1316ms | ✓ 1075ms | http |
| 148.230.4.241:999 | ✓ 685ms | ✓ 1753ms | ✓ 760ms | ✓ 1579ms | ✓ 1323ms | http |
| 122.2.48.121:8080 | ✓ 1539ms | 否 | ✓ 1521ms | ✓ 1698ms | ✓ 1589ms | http |
| 34.87.80.221:30000 | ✓ 929ms | 否 | ✓ 1418ms | ✓ 1403ms | ✓ 1045ms | http |
| 84.47.150.125:1080 | ✓ 953ms | 否 | ✓ 1378ms | 否 | ✓ 1919ms | http |
| 23.230.14.90:8080 | ✓ 1402ms | 否 | ✓ 1488ms | ✓ 1988ms | ✓ 1514ms | http |
| 5.252.33.13:2025 | ✓ 1525ms | 否 | ✓ 1835ms | ✓ 1997ms | ✓ 1753ms | http |
| 5.55.156.85:8080 | 否 | 否 | ✓ 1814ms | ✓ 1920ms | ✓ 1804ms | http |
| 144.124.227.90:21074 | ✓ 1785ms | 否 | ✓ 441ms | ✓ 1706ms | 否 | http |
| 83.171.227.44:3128 | 否 | ✓ 1493ms | 否 | ✓ 1777ms | ✓ 1216ms | http |
| 20.164.75.153:8080 | ✓ 1605ms | 否 | ✓ 1066ms | 否 | ✓ 1986ms | http |
| 108.75.132.171:8080 | 否 | 否 | ✓ 945ms | ✓ 1710ms | ✓ 1838ms | http |
| 128.199.114.189:9090 | ✓ 1045ms | 否 | ✓ 1670ms | ✓ 1361ms | ✓ 1034ms | http |
| 85.192.29.60:3128 | ✓ 861ms | 否 | ✓ 895ms | ✓ 1465ms | ✓ 1265ms | http |
| 170.106.136.181:31002 | ✓ 1963ms | 否 | ✓ 690ms | ✓ 1749ms | ✓ 1283ms | http |
| 147.45.186.28:3128 | ✓ 1235ms | 否 | 否 | ✓ 1386ms | ✓ 1381ms | http |
| 107.150.97.83:3128 | ✓ 570ms | 否 | ✓ 675ms | ✓ 1237ms | ✓ 837ms | http |
| 139.59.105.64:8080 | ✓ 1702ms | 否 | ✓ 1383ms | ✓ 1828ms | ✓ 1581ms | http |
| 147.45.78.89:1080 | ✓ 1059ms | 否 | ✓ 1093ms | ✓ 1594ms | ✓ 1518ms | http |
| 8.219.97.248:80 | ✓ 1708ms | 否 | ✓ 1362ms | 否 | ✓ 1736ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1873ms | ✓ 1531ms | ✓ 1070ms | http |
| 146.190.80.158:9090 | 否 | 否 | ✓ 1408ms | ✓ 1339ms | ✓ 1787ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1482ms | 否 | ✓ 1821ms | ✓ 1289ms | http |
| 34.101.184.164:3128 | ✓ 1747ms | 否 | ✓ 940ms | ✓ 1445ms | ✓ 1131ms | http |
| 64.188.77.26:3128 | ✓ 560ms | ✓ 1672ms | ✓ 594ms | 否 | ✓ 1565ms | http |
| 103.82.23.118:5216 | ✓ 1869ms | 否 | ✓ 1629ms | ✓ 1999ms | ✓ 1781ms | http |
| 120.92.212.16:8890 | ✓ 1586ms | ✓ 1745ms | 否 | 否 | ✓ 1240ms | http |
| 174.138.161.174:33292 | 否 | 否 | ✓ 947ms | ✓ 1296ms | ✓ 1082ms | http |
| 111.79.111.126:3128 | ✓ 1441ms | ✓ 1733ms | ✓ 1884ms | 否 | 否 | http |
| 129.80.217.21:444 | ✓ 1935ms | ✓ 871ms | ✓ 792ms | ✓ 1096ms | ✓ 701ms | http |
| 129.80.238.83:444 | 否 | ✓ 880ms | ✓ 720ms | ✓ 1135ms | ✓ 691ms | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 951ms | ✓ 1304ms | ✓ 1036ms | http |
| 31.172.78.12:3128 | ✓ 1231ms | 否 | ✓ 1965ms | ✓ 1703ms | 否 | http |
| 43.162.123.66:24051 | ✓ 1499ms | ✓ 1063ms | ✓ 473ms | ✓ 927ms | ✓ 690ms | http |
| 121.130.177.28:8888 | ✓ 1339ms | 否 | ✓ 1815ms | ✓ 1612ms | 否 | http |
| 144.31.73.173:3128 | ✓ 910ms | 否 | ✓ 1130ms | 否 | ✓ 1241ms | http |
| 104.248.151.93:9090 | ✓ 981ms | 否 | ✓ 939ms | ✓ 1260ms | ✓ 1075ms | http |
| 159.223.41.216:9090 | ✓ 981ms | 否 | ✓ 943ms | ✓ 1407ms | ✓ 1086ms | http |
| 106.10.55.212:1121 | ✓ 1472ms | ✓ 1353ms | ✓ 1284ms | 否 | ✓ 1183ms | http |
| 3.101.133.120:80 | ✓ 588ms | 否 | ✓ 1217ms | ✓ 1494ms | ✓ 1366ms | http |
| 1.20.207.244:8080 | ✓ 1686ms | 否 | ✓ 1934ms | ✓ 1717ms | ✓ 1757ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1186ms | ✓ 1531ms | ✓ 1379ms | http |
| 223.16.170.103:80 | ✓ 1433ms | 否 | ✓ 1339ms | ✓ 1779ms | ✓ 1368ms | http |
| 34.71.229.255:3128 | ✓ 1516ms | ✓ 1484ms | ✓ 800ms | ✓ 1164ms | ✓ 1226ms | http |
| 208.67.248.60:3128 | ✓ 1538ms | ✓ 1774ms | ✓ 970ms | ✓ 1797ms | ✓ 1718ms | http |
| 121.230.8.133:1080 | ✓ 1304ms | ✓ 1656ms | ✓ 1760ms | 否 | ✓ 1209ms | http |
| 61.52.131.172:8443 | ✓ 859ms | ✓ 1235ms | ✓ 1030ms | ✓ 1260ms | ✓ 1004ms | http |
| 185.41.152.110:3128 | ✓ 955ms | ✓ 1428ms | ✓ 1584ms | ✓ 1951ms | 否 | http |
| 116.171.106.26:3443 | ✓ 1563ms | ✓ 1630ms | ✓ 1581ms | 否 | 否 | http |
| 190.12.150.244:999 | ✓ 1166ms | ✓ 1812ms | ✓ 1490ms | 否 | 否 | http |
| 114.214.165.78:10810 | ✓ 1306ms | 否 | ✓ 1558ms | ✓ 1545ms | ✓ 1249ms | http |
| 121.230.8.136:1080 | 否 | ✓ 1758ms | ✓ 1603ms | 否 | ✓ 1270ms | http |
| 103.189.197.43:7778 | ✓ 1318ms | 否 | ✓ 1927ms | ✓ 1838ms | 否 | http |

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
