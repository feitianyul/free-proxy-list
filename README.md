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

最后更新：2026-04-01 21:48:47 UTC（2026-04-02 05:48:47 UTC+8）

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
| 43.99.54.236:5555 | ✓ 797ms | ✓ 1107ms | ✓ 904ms | ✓ 992ms | ✓ 775ms | http |
| 147.161.210.140:8800 | ✓ 996ms | ✓ 1800ms | ✓ 1046ms | ✓ 1346ms | ✓ 1231ms | http |
| 1.231.81.166:3128 | ✓ 934ms | ✓ 1380ms | ✓ 1629ms | ✓ 1208ms | ✓ 1018ms | http |
| 203.80.138.81:50000 | ✓ 972ms | ✓ 1328ms | ✓ 1224ms | ✓ 1109ms | ✓ 975ms | http |
| 95.213.217.168:52004 | ✓ 1502ms | 否 | ✓ 1277ms | 否 | ✓ 1609ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1514ms | ✓ 1535ms | ✓ 1422ms | ✓ 1117ms | http |
| 42.96.16.158:1311 | ✓ 1120ms | 否 | ✓ 1841ms | ✓ 1477ms | ✓ 1263ms | http |
| 167.103.115.102:8800 | ✓ 1714ms | ✓ 1793ms | ✓ 1308ms | ✓ 1800ms | ✓ 1727ms | http |
| 167.103.34.108:8800 | ✓ 1772ms | 否 | 否 | ✓ 1735ms | ✓ 1643ms | http |
| 35.225.22.61:80 | 否 | ✓ 1701ms | 否 | ✓ 1248ms | ✓ 1002ms | http |
| 34.96.238.40:8080 | ✓ 1019ms | ✓ 1203ms | ✓ 1385ms | ✓ 1133ms | ✓ 1170ms | http |
| 45.12.151.226:2829 | ✓ 696ms | ✓ 1589ms | ✓ 1293ms | 否 | ✓ 1950ms | http |
| 167.103.144.127:8800 | ✓ 1359ms | ✓ 1949ms | ✓ 1550ms | 否 | ✓ 1650ms | http |
| 208.87.243.199:7878 | ✓ 452ms | ✓ 993ms | ✓ 1017ms | ✓ 1061ms | ✓ 835ms | http |
| 165.232.146.249:3128 | ✓ 490ms | ✓ 1314ms | ✓ 1124ms | ✓ 901ms | ✓ 748ms | http |
| 178.128.243.121:3128 | ✓ 1004ms | 否 | ✓ 520ms | ✓ 1292ms | ✓ 1004ms | http |
| 133.242.138.34:8100 | ✓ 918ms | ✓ 1246ms | ✓ 941ms | ✓ 1466ms | ✓ 1866ms | http |
| 167.103.31.122:8800 | ✓ 1711ms | 否 | ✓ 1284ms | ✓ 1969ms | ✓ 1507ms | http |
| 5.104.87.17:8051 | ✓ 1144ms | 否 | ✓ 1117ms | ✓ 1028ms | ✓ 817ms | http |
| 147.161.239.240:8800 | ✓ 1114ms | ✓ 1645ms | ✓ 1296ms | ✓ 1250ms | ✓ 1071ms | http |
| 38.145.218.212:8448 | ✓ 725ms | ✓ 1516ms | ✓ 1347ms | ✓ 1027ms | ✓ 747ms | http |
| 38.145.218.232:8448 | ✓ 729ms | ✓ 1487ms | ✓ 1356ms | ✓ 1029ms | ✓ 742ms | http |
| 101.43.127.100:8877 | ✓ 1034ms | ✓ 1277ms | ✓ 1382ms | ✓ 1152ms | ✓ 1020ms | http |
| 38.34.179.65:8450 | ✓ 962ms | ✓ 1894ms | ✓ 352ms | ✓ 1441ms | 否 | http |
| 160.250.4.245:1 | ✓ 1485ms | 否 | ✓ 1688ms | ✓ 1630ms | ✓ 1219ms | http |
| 92.119.127.212:6005 | ✓ 1126ms | ✓ 1950ms | ✓ 1245ms | ✓ 1779ms | 否 | http |
| 45.136.131.36:8450 | ✓ 1823ms | 否 | 否 | ✓ 1410ms | ✓ 703ms | http |
| 106.10.55.212:1121 | ✓ 799ms | ✓ 1408ms | 否 | ✓ 1381ms | 否 | http |
| 177.234.217.88:999 | ✓ 1501ms | ✓ 1810ms | ✓ 1672ms | ✓ 1939ms | ✓ 1530ms | http |
| 45.140.147.82:1081 | ✓ 1137ms | 否 | ✓ 1210ms | 否 | ✓ 1129ms | http |
| 167.160.184.231:6005 | ✓ 306ms | ✓ 1173ms | ✓ 1057ms | ✓ 1164ms | ✓ 970ms | http |
| 38.34.183.13:8449 | ✓ 1234ms | ✓ 1236ms | ✓ 979ms | ✓ 974ms | ✓ 754ms | http |
| 209.126.84.232:8888 | ✓ 801ms | ✓ 1815ms | 否 | ✓ 1644ms | ✓ 1320ms | http |
| 38.34.179.61:8445 | ✓ 1735ms | ✓ 1708ms | ✓ 391ms | ✓ 1209ms | ✓ 1949ms | http |
| 212.58.132.5:8888 | ✓ 1502ms | 否 | ✓ 1511ms | ✓ 1615ms | ✓ 1266ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1601ms | ✓ 1038ms | 否 | ✓ 1881ms | http |
| 167.160.191.204:6005 | ✓ 469ms | ✓ 1828ms | ✓ 1130ms | ✓ 1144ms | 否 | http |
| 38.34.179.86:8452 | ✓ 742ms | ✓ 1368ms | ✓ 746ms | 否 | ✓ 1220ms | http |
| 34.101.184.164:3128 | ✓ 1616ms | 否 | ✓ 1055ms | ✓ 1303ms | ✓ 1020ms | http |
| 160.250.5.22:1 | ✓ 1425ms | 否 | ✓ 1331ms | ✓ 1521ms | ✓ 1045ms | http |
| 120.92.212.16:7890 | ✓ 1105ms | ✓ 1367ms | ✓ 1108ms | 否 | ✓ 1957ms | http |
| 5.102.109.41:999 | ✓ 688ms | ✓ 1501ms | 否 | ✓ 1904ms | ✓ 1997ms | http |
| 121.230.9.161:1080 | 否 | 否 | ✓ 1949ms | ✓ 1944ms | ✓ 1284ms | http |
| 38.145.208.207:8445 | ✓ 1470ms | ✓ 1863ms | ✓ 800ms | ✓ 1522ms | ✓ 1681ms | http |
| 217.217.249.160:8080 | ✓ 1279ms | 否 | ✓ 1349ms | 否 | ✓ 1558ms | http |
| 8.219.97.248:80 | ✓ 1824ms | 否 | ✓ 1074ms | ✓ 1652ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1815ms | ✓ 1650ms | ✓ 1373ms | ✓ 1410ms | ✓ 1452ms | http |
| 38.180.2.107:3128 | ✓ 1802ms | ✓ 1674ms | 否 | 否 | ✓ 1904ms | http |
| 91.233.223.147:3128 | ✓ 1654ms | 否 | ✓ 1105ms | 否 | ✓ 1960ms | http |
| 192.203.0.210:999 | ✓ 1389ms | ✓ 1637ms | ✓ 1294ms | ✓ 1879ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1070ms | 否 | ✓ 1209ms | ✓ 1526ms | 否 | http |
| 158.160.215.167:8125 | ✓ 1561ms | ✓ 1932ms | ✓ 1854ms | 否 | 否 | http |
| 150.107.140.238:3128 | ✓ 1623ms | 否 | 否 | ✓ 1408ms | ✓ 1194ms | http |
| 190.52.108.184:999 | ✓ 1910ms | ✓ 1594ms | ✓ 1625ms | ✓ 1469ms | ✓ 1261ms | http |
| 171.234.52.12:5102 | ✓ 1718ms | 否 | 否 | ✓ 1787ms | ✓ 1511ms | http |
| 116.80.65.79:3172 | ✓ 1606ms | 否 | ✓ 1651ms | ✓ 1946ms | 否 | http |
| 186.116.148.52:8080 | ✓ 1163ms | ✓ 1886ms | ✓ 1891ms | 否 | 否 | http |
| 121.230.8.111:1080 | ✓ 1383ms | ✓ 1507ms | ✓ 1949ms | ✓ 1509ms | ✓ 1287ms | http |
| 38.34.179.202:8449 | ✓ 977ms | ✓ 1278ms | ✓ 850ms | ✓ 1291ms | ✓ 1077ms | http |
| 38.145.203.135:8444 | 否 | ✓ 825ms | ✓ 571ms | 否 | ✓ 1382ms | http |
| 38.34.179.16:8451 | 否 | ✓ 996ms | ✓ 986ms | 否 | ✓ 779ms | http |
| 38.145.208.220:8448 | ✓ 1595ms | ✓ 853ms | 否 | ✓ 1638ms | ✓ 841ms | http |
| 203.150.113.51:8080 | ✓ 1591ms | 否 | 否 | ✓ 1793ms | ✓ 1684ms | http |
| 38.145.220.102:8453 | ✓ 1104ms | 否 | ✓ 501ms | ✓ 911ms | ✓ 1057ms | http |
| 37.187.109.70:10111 | ✓ 1274ms | ✓ 1472ms | ✓ 962ms | 否 | ✓ 1977ms | http |
| 106.117.208.101:7890 | ✓ 1476ms | ✓ 1306ms | ✓ 1184ms | ✓ 1673ms | ✓ 1297ms | http |
| 45.149.92.147:5001 | ✓ 772ms | 否 | ✓ 789ms | ✓ 969ms | ✓ 742ms | http |
| 159.223.71.162:443 | ✓ 879ms | 否 | ✓ 1090ms | ✓ 1396ms | 否 | http |
| 178.156.224.42:3128 | ✓ 894ms | 否 | ✓ 1577ms | 否 | ✓ 1952ms | http |
| 45.129.141.143:3128 | ✓ 755ms | ✓ 1898ms | ✓ 1833ms | ✓ 1738ms | ✓ 1590ms | http |
| 190.2.213.169:999 | ✓ 1678ms | ✓ 1598ms | ✓ 1601ms | 否 | 否 | http |
| 47.105.98.23:3128 | ✓ 1224ms | ✓ 1308ms | ✓ 1054ms | ✓ 1363ms | ✓ 1105ms | http |
| 103.189.116.20:8080 | ✓ 1846ms | 否 | ✓ 1686ms | ✓ 1778ms | 否 | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1953ms | ✓ 1472ms | ✓ 1551ms | http |
| 31.192.106.135:8010 | ✓ 1304ms | 否 | ✓ 1272ms | 否 | ✓ 1970ms | http |
| 45.174.253.64:999 | ✓ 732ms | 否 | ✓ 993ms | ✓ 1214ms | 否 | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1248ms | ✓ 1511ms | ✓ 1300ms | http |
| 38.145.208.246:8450 | ✓ 270ms | ✓ 893ms | ✓ 418ms | ✓ 923ms | ✓ 692ms | http |
| 45.136.130.168:8448 | ✓ 275ms | ✓ 804ms | ✓ 803ms | ✓ 898ms | ✓ 758ms | http |
| 45.136.130.173:8448 | ✓ 285ms | ✓ 808ms | ✓ 789ms | ✓ 926ms | ✓ 792ms | http |
| 45.136.130.170:8448 | ✓ 316ms | ✓ 809ms | ✓ 756ms | ✓ 908ms | ✓ 742ms | http |
| 45.136.131.30:8451 | ✓ 256ms | ✓ 837ms | ✓ 975ms | ✓ 1079ms | ✓ 863ms | http |
| 121.230.9.248:1080 | ✓ 1215ms | ✓ 1552ms | ✓ 1031ms | ✓ 1559ms | ✓ 1114ms | http |
| 128.199.121.61:9090 | ✓ 1804ms | 否 | ✓ 1780ms | 否 | ✓ 1247ms | http |
| 38.34.179.203:8450 | 否 | ✓ 1227ms | ✓ 781ms | 否 | ✓ 1010ms | http |
| 121.230.8.97:1080 | ✓ 1280ms | ✓ 1734ms | ✓ 1437ms | ✓ 1742ms | ✓ 1390ms | http |
| 198.59.68.130:3128 | ✓ 547ms | ✓ 1160ms | ✓ 1343ms | ✓ 1523ms | ✓ 1272ms | http |
| 103.87.171.180:32650 | ✓ 1934ms | 否 | 否 | ✓ 1961ms | ✓ 1939ms | http |
| 38.145.220.9:8448 | 否 | ✓ 1255ms | ✓ 1702ms | 否 | ✓ 805ms | http |
| 38.34.183.47:8452 | 否 | ✓ 1486ms | ✓ 319ms | ✓ 1285ms | ✓ 1557ms | http |
| 38.145.208.242:8451 | 否 | 否 | ✓ 1129ms | ✓ 1708ms | ✓ 1676ms | http |
| 166.0.192.117:8888 | ✓ 279ms | ✓ 872ms | ✓ 272ms | ✓ 948ms | ✓ 697ms | http |

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
