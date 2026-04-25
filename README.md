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

最后更新：2026-04-25 00:32:18 UTC（2026-04-25 08:32:18 UTC+8）

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
| 47.85.51.197:1080 | ✓ 577ms | ✓ 1101ms | ✓ 698ms | ✓ 1283ms | ✓ 1703ms | http |
| 206.206.126.177:2412 | ✓ 1465ms | ✓ 1608ms | ✓ 1338ms | ✓ 1015ms | ✓ 800ms | http |
| 46.101.95.183:8888 | ✓ 1248ms | ✓ 1975ms | ✓ 1620ms | 否 | ✓ 1122ms | http |
| 15.160.132.166:8800 | ✓ 1282ms | 否 | ✓ 1138ms | ✓ 1870ms | ✓ 1550ms | http |
| 8.219.195.129:1080 | ✓ 910ms | ✓ 1684ms | ✓ 927ms | ✓ 1319ms | ✓ 976ms | http |
| 2.27.54.161:1080 | ✓ 1102ms | 否 | ✓ 1431ms | ✓ 1916ms | 否 | http |
| 217.76.245.80:999 | ✓ 1133ms | ✓ 1729ms | ✓ 1327ms | ✓ 1431ms | ✓ 1481ms | http |
| 59.46.216.131:30001 | ✓ 1063ms | ✓ 1359ms | ✓ 1125ms | 否 | 否 | http |
| 16.62.123.236:48789 | ✓ 1213ms | 否 | ✓ 893ms | ✓ 1828ms | ✓ 1724ms | http |
| 15.160.116.45:13815 | ✓ 1462ms | 否 | ✓ 886ms | 否 | ✓ 1701ms | http |
| 15.157.63.22:15433 | ✓ 1411ms | 否 | ✓ 1944ms | 否 | ✓ 1824ms | http |
| 54.229.201.146:48867 | ✓ 1210ms | 否 | ✓ 1807ms | ✓ 1797ms | 否 | http |
| 51.34.28.236:26666 | ✓ 1949ms | 否 | ✓ 869ms | 否 | ✓ 1842ms | http |
| 52.59.218.12:616 | ✓ 1914ms | 否 | ✓ 1639ms | 否 | ✓ 1621ms | http |
| 91.217.81.131:1080 | ✓ 1234ms | ✓ 1992ms | ✓ 1677ms | 否 | 否 | http |
| 52.59.51.29:17422 | ✓ 1921ms | 否 | ✓ 1924ms | ✓ 1633ms | 否 | http |
| 161.35.181.96:999 | ✓ 923ms | ✓ 1152ms | ✓ 956ms | ✓ 1161ms | ✓ 1043ms | http |
| 125.64.244.100:8889 | ✓ 1520ms | ✓ 1736ms | ✓ 1461ms | ✓ 1930ms | ✓ 1576ms | http |
| 43.132.188.134:443 | ✓ 874ms | ✓ 1546ms | 否 | 否 | ✓ 1884ms | http |
| 47.101.159.19:8899 | ✓ 1405ms | ✓ 1069ms | ✓ 900ms | ✓ 1097ms | ✓ 916ms | http |
| 149.62.191.202:3128 | ✓ 1322ms | ✓ 1801ms | ✓ 766ms | 否 | ✓ 1376ms | http |
| 202.47.185.178:8080 | ✓ 1847ms | 否 | ✓ 1342ms | ✓ 1384ms | 否 | http |
| 35.225.22.61:80 | ✓ 827ms | 否 | ✓ 508ms | 否 | ✓ 1004ms | http |
| 218.108.131.186:17890 | ✓ 840ms | ✓ 1048ms | ✓ 1398ms | ✓ 1088ms | ✓ 924ms | http |
| 120.92.108.86:7890 | ✓ 1965ms | 否 | 否 | ✓ 1746ms | ✓ 1617ms | http |
| 113.160.132.26:8080 | ✓ 1467ms | 否 | ✓ 1078ms | ✓ 1226ms | ✓ 1035ms | http |
| 177.93.132.244:3128 | ✓ 1521ms | 否 | ✓ 1568ms | 否 | ✓ 1808ms | http |
| 160.238.65.7:3128 | ✓ 1689ms | 否 | ✓ 1761ms | 否 | ✓ 1281ms | http |
| 36.155.100.217:8080 | ✓ 1342ms | ✓ 1071ms | ✓ 1171ms | ✓ 1234ms | ✓ 1056ms | http |
| 62.113.119.14:8080 | ✓ 1254ms | ✓ 1780ms | ✓ 1098ms | ✓ 1601ms | ✓ 1210ms | http |
| 130.61.174.200:1080 | ✓ 1233ms | ✓ 1944ms | ✓ 1860ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 888ms | ✓ 1321ms | ✓ 1406ms | ✓ 1487ms | ✓ 1527ms | http |
| 38.180.192.119:3128 | ✓ 179ms | ✓ 757ms | ✓ 359ms | ✓ 932ms | ✓ 607ms | http |
| 160.238.65.9:3128 | ✓ 1208ms | 否 | ✓ 1303ms | ✓ 1540ms | 否 | http |
| 121.230.8.220:1080 | ✓ 1190ms | ✓ 1323ms | ✓ 1010ms | ✓ 1341ms | ✓ 1099ms | http |
| 121.138.61.78:8943 | ✓ 1239ms | 否 | ✓ 1285ms | 否 | ✓ 1191ms | http |
| 34.71.229.255:3128 | ✓ 1193ms | ✓ 1784ms | ✓ 1812ms | 否 | ✓ 1770ms | http |
| 160.238.65.4:3128 | ✓ 1637ms | 否 | ✓ 1660ms | ✓ 1972ms | 否 | http |
| 34.138.160.210:3128 | ✓ 1198ms | ✓ 1978ms | ✓ 1306ms | ✓ 1930ms | ✓ 1499ms | http |
| 133.125.237.205:8888 | ✓ 1537ms | 否 | ✓ 1504ms | ✓ 1742ms | ✓ 1632ms | http |
| 101.32.244.83:8080 | ✓ 1012ms | ✓ 1555ms | ✓ 952ms | ✓ 1374ms | ✓ 1232ms | http |
| 121.43.196.210:8222 | ✓ 1178ms | ✓ 1092ms | ✓ 845ms | ✓ 1177ms | ✓ 920ms | http |
| 121.43.196.213:8222 | ✓ 900ms | ✓ 1079ms | ✓ 1178ms | ✓ 1200ms | ✓ 971ms | http |
| 114.55.226.123:10086 | ✓ 1082ms | ✓ 1470ms | ✓ 1031ms | ✓ 1531ms | ✓ 1146ms | http |
| 3.137.167.45:93 | ✓ 1270ms | 否 | ✓ 1969ms | 否 | ✓ 1870ms | http |
| 78.12.18.91:25013 | ✓ 1533ms | 否 | ✓ 1911ms | 否 | ✓ 1874ms | http |
| 18.188.53.175:17723 | ✓ 804ms | 否 | ✓ 1590ms | ✓ 1980ms | 否 | http |
| 52.53.211.45:1111 | 否 | 否 | ✓ 1148ms | ✓ 1983ms | ✓ 1465ms | http |
| 3.121.130.230:9098 | ✓ 1731ms | 否 | 否 | ✓ 1601ms | ✓ 1217ms | http |
| 120.27.224.64:3128 | ✓ 917ms | ✓ 1060ms | ✓ 886ms | ✓ 1158ms | ✓ 969ms | http |
| 152.42.177.32:8888 | ✓ 954ms | 否 | ✓ 1020ms | ✓ 1225ms | ✓ 1258ms | http |
| 62.60.216.109:3128 | ✓ 565ms | ✓ 1615ms | ✓ 1615ms | 否 | ✓ 1389ms | http |
| 150.107.140.238:3128 | ✓ 1584ms | 否 | ✓ 976ms | 否 | ✓ 1151ms | http |
| 103.157.200.126:3128 | ✓ 1673ms | 否 | ✓ 1404ms | ✓ 1962ms | ✓ 1539ms | http |
| 223.84.151.86:30005 | ✓ 1270ms | ✓ 1336ms | ✓ 1105ms | ✓ 1364ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1190ms | 否 | ✓ 1428ms | 否 | ✓ 1586ms | http |
| 121.230.8.136:1080 | ✓ 1220ms | ✓ 1268ms | ✓ 1016ms | ✓ 1369ms | ✓ 1325ms | http |
| 183.232.248.73:7890 | ✓ 1095ms | ✓ 1157ms | ✓ 1004ms | ✓ 1438ms | ✓ 1343ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1437ms | ✓ 1471ms | ✓ 1491ms | 否 | http |
| 211.95.152.50:45046 | ✓ 1159ms | ✓ 1506ms | ✓ 1041ms | ✓ 1215ms | ✓ 1146ms | http |
| 8.211.166.184:8081 | ✓ 1480ms | ✓ 1463ms | ✓ 716ms | ✓ 834ms | ✓ 661ms | http |
| 3.121.42.224:3366 | ✓ 1351ms | 否 | ✓ 1436ms | 否 | ✓ 1761ms | http |
| 15.216.178.25:50090 | ✓ 1382ms | 否 | ✓ 1833ms | 否 | ✓ 1803ms | http |
| 52.16.215.4:42899 | ✓ 1125ms | 否 | ✓ 1444ms | ✓ 1684ms | ✓ 1589ms | http |
| 13.38.59.232:45664 | ✓ 1128ms | 否 | ✓ 1254ms | ✓ 1877ms | ✓ 1861ms | http |
| 45.88.0.116:3128 | ✓ 1211ms | ✓ 1495ms | ✓ 1136ms | 否 | ✓ 1635ms | http |
| 45.88.0.99:3128 | ✓ 1211ms | ✓ 1650ms | ✓ 987ms | 否 | ✓ 1626ms | http |
| 45.88.0.98:3128 | ✓ 1198ms | ✓ 1499ms | ✓ 1128ms | 否 | ✓ 1641ms | http |
| 91.186.217.84:3128 | ✓ 1192ms | ✓ 1820ms | ✓ 1600ms | 否 | ✓ 1922ms | http |
| 162.240.154.26:3128 | ✓ 1172ms | ✓ 1718ms | 否 | 否 | ✓ 1384ms | http |
| 178.156.224.42:3128 | ✓ 1422ms | ✓ 1806ms | ✓ 1805ms | 否 | ✓ 1682ms | http |
| 103.185.248.155:8888 | ✓ 988ms | 否 | ✓ 1719ms | ✓ 1146ms | ✓ 1270ms | http |
| 91.99.15.45:2095 | ✓ 1482ms | 否 | ✓ 1428ms | 否 | ✓ 1352ms | http |
| 45.88.0.115:3128 | ✓ 675ms | ✓ 1293ms | ✓ 1284ms | ✓ 1987ms | ✓ 1491ms | http |
| 61.52.131.172:8443 | 否 | ✓ 1214ms | ✓ 954ms | ✓ 1221ms | ✓ 957ms | http |
| 45.88.0.117:3128 | ✓ 551ms | ✓ 1517ms | ✓ 995ms | ✓ 1408ms | ✓ 1110ms | http |
| 213.220.62.62:3128 | 否 | ✓ 1524ms | ✓ 551ms | ✓ 1373ms | ✓ 1135ms | http |
| 18.222.132.180:35576 | ✓ 855ms | 否 | ✓ 1737ms | 否 | ✓ 1727ms | http |
| 3.19.213.118:40000 | ✓ 1798ms | 否 | 否 | ✓ 1768ms | ✓ 1410ms | http |
| 16.62.229.137:56768 | ✓ 1957ms | 否 | ✓ 1048ms | 否 | ✓ 1465ms | http |
| 117.236.124.166:3128 | ✓ 1719ms | 否 | ✓ 1153ms | ✓ 1927ms | 否 | http |
| 103.39.51.207:8080 | ✓ 1436ms | 否 | ✓ 1878ms | ✓ 1599ms | ✓ 1417ms | http |
| 45.140.147.82:1081 | ✓ 1478ms | ✓ 1688ms | ✓ 905ms | 否 | ✓ 1346ms | http |

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
