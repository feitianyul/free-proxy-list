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

最后更新：2026-04-02 21:34:06 UTC（2026-04-03 05:34:06 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 565ms | ✓ 1124ms | 否 | 否 | ✓ 1087ms | http |
| 208.87.243.199:7878 | ✓ 1019ms | ✓ 873ms | ✓ 1171ms | ✓ 1076ms | ✓ 712ms | http |
| 43.99.54.236:5555 | ✓ 897ms | ✓ 1158ms | ✓ 890ms | ✓ 1068ms | ✓ 861ms | http |
| 203.80.138.81:50000 | ✓ 1172ms | ✓ 1392ms | ✓ 1115ms | ✓ 1142ms | ✓ 1101ms | http |
| 95.213.217.168:52004 | ✓ 585ms | ✓ 1556ms | 否 | ✓ 1941ms | ✓ 1622ms | http |
| 1.231.81.166:3128 | ✓ 1668ms | ✓ 1204ms | ✓ 1958ms | ✓ 1214ms | ✓ 981ms | http |
| 147.161.210.140:8800 | ✓ 1635ms | 否 | ✓ 1111ms | ✓ 1331ms | ✓ 1335ms | http |
| 159.223.71.162:8080 | ✓ 1695ms | 否 | ✓ 1695ms | ✓ 1307ms | ✓ 1075ms | http |
| 212.58.132.5:8888 | ✓ 1692ms | ✓ 1911ms | ✓ 1599ms | ✓ 1465ms | ✓ 1194ms | http |
| 167.103.115.102:8800 | ✓ 1702ms | ✓ 1944ms | ✓ 1179ms | 否 | ✓ 1418ms | http |
| 113.160.132.26:8080 | ✓ 1961ms | ✓ 1480ms | ✓ 1527ms | ✓ 1514ms | ✓ 1725ms | http |
| 167.103.34.108:8800 | ✓ 1725ms | ✓ 1954ms | ✓ 1596ms | 否 | ✓ 1484ms | http |
| 31.192.106.135:8005 | ✓ 1798ms | 否 | ✓ 1782ms | 否 | ✓ 1921ms | http |
| 115.231.181.40:8128 | ✓ 1150ms | ✓ 1440ms | ✓ 1398ms | ✓ 1458ms | ✓ 1649ms | http |
| 120.92.212.16:8890 | ✓ 1109ms | ✓ 1424ms | 否 | 否 | ✓ 1528ms | http |
| 59.46.216.131:30001 | ✓ 1209ms | 否 | 否 | ✓ 1638ms | ✓ 1219ms | http |
| 45.167.124.52:8080 | ✓ 510ms | ✓ 1615ms | ✓ 508ms | 否 | 否 | http |
| 45.167.125.21:999 | ✓ 773ms | ✓ 1791ms | ✓ 1456ms | ✓ 1715ms | ✓ 1655ms | http |
| 167.103.31.122:8800 | ✓ 1281ms | 否 | ✓ 1190ms | ✓ 1485ms | ✓ 1372ms | http |
| 133.242.138.34:8100 | ✓ 1587ms | ✓ 1502ms | ✓ 812ms | ✓ 1765ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1406ms | 否 | ✓ 1291ms | ✓ 1476ms | ✓ 1344ms | http |
| 120.92.212.16:7890 | ✓ 1149ms | 否 | ✓ 1538ms | ✓ 1756ms | 否 | http |
| 147.161.239.240:8800 | ✓ 687ms | ✓ 1477ms | ✓ 1188ms | ✓ 1651ms | ✓ 1498ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 1017ms | ✓ 1380ms | ✓ 1008ms | http |
| 101.43.127.100:8877 | ✓ 1036ms | ✓ 1280ms | ✓ 1025ms | ✓ 1322ms | ✓ 1086ms | http |
| 160.250.5.22:1 | ✓ 1763ms | 否 | ✓ 1438ms | 否 | ✓ 1478ms | http |
| 128.199.121.61:9090 | ✓ 1155ms | 否 | ✓ 1310ms | ✓ 1351ms | 否 | http |
| 38.145.208.179:8447 | ✓ 973ms | ✓ 1113ms | 否 | ✓ 1163ms | ✓ 796ms | http |
| 128.199.114.189:9090 | ✓ 967ms | 否 | ✓ 1551ms | ✓ 1355ms | ✓ 1226ms | http |
| 82.114.228.67:1080 | ✓ 593ms | 否 | ✓ 1621ms | ✓ 1486ms | 否 | http |
| 177.234.217.88:999 | ✓ 1338ms | ✓ 1820ms | ✓ 1759ms | ✓ 1851ms | ✓ 1663ms | http |
| 45.136.130.169:8446 | ✓ 591ms | ✓ 1100ms | 否 | ✓ 973ms | ✓ 1003ms | http |
| 146.190.80.158:9090 | ✓ 1601ms | 否 | ✓ 1058ms | ✓ 1337ms | ✓ 1133ms | http |
| 5.102.109.41:999 | ✓ 678ms | ✓ 1600ms | ✓ 331ms | ✓ 1259ms | ✓ 1394ms | http |
| 150.241.71.15:1080 | ✓ 1729ms | 否 | ✓ 1069ms | 否 | ✓ 1353ms | http |
| 195.123.209.48:3128 | ✓ 591ms | ✓ 1447ms | ✓ 1142ms | ✓ 1855ms | ✓ 1485ms | http |
| 45.144.232.5:11741 | ✓ 652ms | 否 | ✓ 1581ms | 否 | ✓ 1616ms | http |
| 34.101.184.164:3128 | ✓ 1599ms | 否 | ✓ 939ms | ✓ 1364ms | ✓ 1091ms | http |
| 92.119.127.212:6005 | ✓ 1651ms | ✓ 1855ms | ✓ 1837ms | ✓ 1756ms | ✓ 1887ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1937ms | ✓ 1534ms | ✓ 1252ms | http |
| 42.96.16.158:1311 | ✓ 1825ms | 否 | ✓ 1027ms | ✓ 1744ms | 否 | http |
| 177.93.33.55:999 | ✓ 1056ms | ✓ 1638ms | ✓ 1444ms | 否 | ✓ 1730ms | http |
| 45.140.147.155:1081 | ✓ 1425ms | ✓ 1352ms | ✓ 520ms | ✓ 1208ms | ✓ 1426ms | http |
| 159.223.71.162:443 | ✓ 1646ms | 否 | ✓ 1312ms | ✓ 1318ms | ✓ 1237ms | http |
| 31.192.106.135:8010 | ✓ 1220ms | 否 | 否 | ✓ 1943ms | ✓ 1599ms | http |
| 45.12.151.226:2829 | ✓ 978ms | 否 | ✓ 1167ms | 否 | ✓ 1856ms | http |
| 89.208.106.138:10808 | ✓ 1269ms | ✓ 1867ms | ✓ 1048ms | ✓ 1746ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1387ms | ✓ 1974ms | ✓ 1311ms | ✓ 1877ms | ✓ 1928ms | http |
| 45.140.147.155:1082 | ✓ 710ms | ✓ 1384ms | ✓ 551ms | ✓ 1771ms | ✓ 1289ms | http |
| 45.186.6.104:3128 | ✓ 1357ms | ✓ 1889ms | ✓ 1808ms | 否 | 否 | http |
| 38.145.220.198:8448 | 否 | ✓ 1413ms | ✓ 1976ms | ✓ 1156ms | ✓ 1595ms | http |
| 38.34.179.21:8446 | 否 | ✓ 861ms | ✓ 394ms | ✓ 1100ms | ✓ 1136ms | http |
| 38.34.179.100:8449 | 否 | ✓ 1087ms | ✓ 313ms | ✓ 914ms | ✓ 879ms | http |
| 38.34.179.202:8449 | 否 | ✓ 940ms | ✓ 765ms | ✓ 1461ms | ✓ 822ms | http |
| 38.34.179.167:8448 | 否 | ✓ 1728ms | ✓ 1227ms | ✓ 990ms | 否 | http |
| 47.105.98.23:3128 | ✓ 1626ms | ✓ 1410ms | ✓ 1179ms | ✓ 1446ms | ✓ 1181ms | http |
| 128.199.113.85:9090 | ✓ 921ms | 否 | ✓ 964ms | ✓ 1361ms | ✓ 1054ms | http |
| 106.117.208.101:7890 | ✓ 1136ms | ✓ 1561ms | ✓ 1199ms | ✓ 1491ms | ✓ 1238ms | http |
| 27.254.99.183:8118 | ✓ 1055ms | 否 | ✓ 1416ms | 否 | ✓ 1230ms | http |
| 86.53.183.16:1080 | ✓ 700ms | 否 | ✓ 1042ms | ✓ 1997ms | ✓ 1734ms | http |
| 38.180.2.107:3128 | ✓ 1129ms | ✓ 1561ms | ✓ 1429ms | 否 | ✓ 1928ms | http |
| 217.217.249.160:8080 | ✓ 1348ms | 否 | ✓ 1275ms | 否 | ✓ 1378ms | http |
| 168.110.52.228:3128 | ✓ 1721ms | 否 | ✓ 667ms | ✓ 1046ms | ✓ 1894ms | http |
| 38.34.179.27:8451 | ✓ 1856ms | ✓ 885ms | ✓ 828ms | 否 | ✓ 708ms | http |
| 34.96.238.40:8080 | ✓ 1319ms | ✓ 1329ms | 否 | 否 | ✓ 1273ms | http |
| 45.136.198.40:3128 | 否 | ✓ 1388ms | ✓ 1931ms | ✓ 1941ms | ✓ 1751ms | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1134ms | ✓ 1394ms | ✓ 1042ms | http |
| 45.136.130.177:8448 | ✓ 506ms | ✓ 1174ms | ✓ 1921ms | ✓ 1117ms | ✓ 1289ms | http |
| 104.248.151.93:9090 | ✓ 1284ms | 否 | ✓ 1246ms | ✓ 1342ms | ✓ 1070ms | http |
| 103.113.70.189:1081 | ✓ 264ms | ✓ 1013ms | ✓ 722ms | ✓ 962ms | ✓ 1800ms | http |
| 223.16.170.103:80 | ✓ 1119ms | ✓ 1951ms | ✓ 1074ms | ✓ 1368ms | ✓ 1659ms | http |
| 223.16.170.103:3128 | ✓ 1392ms | ✓ 1901ms | ✓ 1291ms | ✓ 1332ms | ✓ 1355ms | http |
| 103.39.51.190:8080 | ✓ 1989ms | 否 | 否 | ✓ 1862ms | ✓ 1616ms | http |
| 167.71.196.28:8080 | ✓ 1325ms | 否 | ✓ 1184ms | ✓ 1321ms | 否 | http |
| 38.34.179.51:8449 | ✓ 729ms | ✓ 1285ms | ✓ 1247ms | ✓ 955ms | ✓ 844ms | http |
| 117.86.6.209:1080 | ✓ 1337ms | ✓ 1422ms | ✓ 1749ms | ✓ 1708ms | ✓ 1285ms | http |
| 45.136.130.195:8448 | ✓ 950ms | 否 | ✓ 894ms | 否 | ✓ 1886ms | http |
| 45.136.130.192:8448 | ✓ 948ms | 否 | ✓ 859ms | 否 | ✓ 1925ms | http |
| 190.97.231.0:999 | ✓ 1105ms | ✓ 1628ms | ✓ 1212ms | ✓ 1728ms | 否 | http |
| 185.135.99.14:8080 | 否 | ✓ 1923ms | ✓ 1400ms | ✓ 1886ms | 否 | http |
| 121.230.8.247:1080 | ✓ 1644ms | ✓ 1757ms | ✓ 1431ms | 否 | 否 | http |
| 20.120.225.109:3128 | 否 | ✓ 1504ms | ✓ 1190ms | ✓ 1578ms | ✓ 1028ms | http |
| 72.11.151.159:6005 | 否 | 否 | ✓ 1365ms | ✓ 1455ms | ✓ 1455ms | http |
| 38.145.203.19:8447 | ✓ 1320ms | 否 | ✓ 661ms | ✓ 1544ms | 否 | http |
| 158.160.215.167:8127 | ✓ 852ms | ✓ 1791ms | ✓ 1562ms | 否 | ✓ 1952ms | http |

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
