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

最后更新：2026-04-06 16:48:58 UTC（2026-04-07 00:48:58 UTC+8）

**代理总数：91**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1552ms | ✓ 1243ms | ✓ 803ms | ✓ 1852ms | ✓ 880ms | http |
| 167.103.115.102:8800 | ✓ 923ms | 否 | ✓ 1598ms | ✓ 1066ms | ✓ 1015ms | http |
| 113.160.132.26:8080 | ✓ 1804ms | ✓ 1486ms | ✓ 1030ms | ✓ 1208ms | ✓ 1034ms | http |
| 159.223.71.162:8080 | ✓ 746ms | 否 | ✓ 1676ms | ✓ 1093ms | ✓ 1401ms | http |
| 159.223.71.162:443 | ✓ 752ms | 否 | ✓ 1855ms | ✓ 1149ms | ✓ 1186ms | http |
| 111.227.254.9:22222 | ✓ 1069ms | ✓ 1396ms | ✓ 1040ms | 否 | ✓ 1159ms | http |
| 1.231.81.166:3128 | ✓ 1578ms | ✓ 1329ms | ✓ 1858ms | ✓ 1230ms | ✓ 918ms | http |
| 111.227.254.12:22222 | ✓ 1117ms | ✓ 1441ms | 否 | 否 | ✓ 1100ms | http |
| 167.103.34.108:8800 | ✓ 1505ms | 否 | ✓ 1552ms | 否 | ✓ 1668ms | http |
| 170.106.137.214:7890 | ✓ 1041ms | 否 | 否 | ✓ 722ms | ✓ 526ms | http |
| 218.108.131.186:17890 | ✓ 858ms | ✓ 1041ms | ✓ 904ms | ✓ 1227ms | ✓ 891ms | http |
| 38.145.208.211:8453 | 否 | ✓ 1437ms | ✓ 155ms | ✓ 771ms | ✓ 1404ms | http |
| 167.103.144.127:8800 | ✓ 1628ms | 否 | ✓ 1476ms | ✓ 1898ms | 否 | http |
| 59.46.216.131:30001 | ✓ 971ms | ✓ 1433ms | 否 | ✓ 1380ms | ✓ 1081ms | http |
| 212.58.132.5:8888 | ✓ 1190ms | 否 | ✓ 1730ms | ✓ 1900ms | ✓ 1571ms | http |
| 35.225.22.61:80 | ✓ 305ms | 否 | 否 | ✓ 1175ms | ✓ 1180ms | http |
| 167.103.31.122:8800 | ✓ 1734ms | 否 | ✓ 1373ms | 否 | ✓ 1531ms | http |
| 116.80.95.226:3172 | ✓ 1534ms | 否 | ✓ 1507ms | ✓ 1887ms | ✓ 1724ms | http |
| 115.231.181.40:8128 | ✓ 1973ms | ✓ 1637ms | 否 | ✓ 1274ms | ✓ 1904ms | http |
| 116.80.65.221:3172 | ✓ 1530ms | 否 | 否 | ✓ 1946ms | ✓ 1715ms | http |
| 111.227.254.11:22222 | ✓ 1097ms | 否 | 否 | ✓ 1313ms | ✓ 1166ms | http |
| 120.92.212.16:8890 | ✓ 1884ms | 否 | ✓ 987ms | ✓ 1250ms | ✓ 1032ms | http |
| 20.78.26.206:8561 | ✓ 793ms | ✓ 1084ms | ✓ 895ms | ✓ 1410ms | ✓ 1984ms | http |
| 20.210.39.153:8561 | ✓ 798ms | ✓ 1071ms | ✓ 900ms | ✓ 1410ms | ✓ 1987ms | http |
| 111.227.254.10:22222 | ✓ 1127ms | ✓ 1531ms | ✓ 1091ms | 否 | 否 | http |
| 217.76.245.80:999 | ✓ 944ms | ✓ 1539ms | ✓ 1449ms | 否 | ✓ 1358ms | http |
| 147.161.239.240:8800 | 否 | ✓ 1698ms | ✓ 1384ms | ✓ 1709ms | ✓ 1568ms | http |
| 101.43.127.100:8877 | ✓ 914ms | ✓ 1997ms | ✓ 1244ms | ✓ 1185ms | ✓ 978ms | http |
| 45.167.125.21:999 | ✓ 1499ms | ✓ 1945ms | ✓ 1620ms | ✓ 1784ms | ✓ 1573ms | http |
| 38.145.208.204:8446 | ✓ 1622ms | ✓ 1216ms | ✓ 1150ms | ✓ 1737ms | ✓ 601ms | http |
| 45.136.130.188:8449 | ✓ 967ms | ✓ 775ms | ✓ 738ms | 否 | ✓ 1145ms | http |
| 89.208.106.138:10808 | ✓ 1216ms | ✓ 1670ms | ✓ 1846ms | ✓ 1966ms | ✓ 1415ms | http |
| 150.107.29.165:3128 | ✓ 1348ms | 否 | ✓ 1322ms | ✓ 1662ms | ✓ 1273ms | http |
| 34.101.184.164:3128 | ✓ 1571ms | 否 | 否 | ✓ 1554ms | ✓ 1902ms | http |
| 121.43.196.210:8222 | ✓ 935ms | ✓ 1084ms | ✓ 924ms | ✓ 1146ms | ✓ 974ms | http |
| 120.92.212.16:7890 | ✓ 979ms | ✓ 1445ms | ✓ 936ms | ✓ 1263ms | ✓ 997ms | http |
| 62.113.119.14:8080 | ✓ 1018ms | 否 | ✓ 1257ms | ✓ 1581ms | ✓ 1262ms | http |
| 141.11.210.159:8888 | ✓ 600ms | ✓ 1214ms | ✓ 899ms | ✓ 900ms | ✓ 1761ms | http |
| 195.133.9.136:3128 | ✓ 1493ms | ✓ 1877ms | 否 | 否 | ✓ 1802ms | http |
| 20.78.118.91:8561 | 否 | ✓ 1406ms | ✓ 504ms | ✓ 797ms | ✓ 649ms | http |
| 95.181.174.211:1080 | ✓ 1352ms | 否 | ✓ 1148ms | ✓ 1924ms | ✓ 1516ms | http |
| 209.38.154.7:1080 | ✓ 590ms | ✓ 1806ms | 否 | ✓ 1856ms | 否 | http |
| 158.160.215.167:8127 | ✓ 747ms | 否 | ✓ 921ms | ✓ 1835ms | 否 | http |
| 38.145.220.27:8445 | ✓ 663ms | 否 | ✓ 183ms | ✓ 1148ms | ✓ 1217ms | http |
| 38.34.179.57:8448 | ✓ 513ms | 否 | ✓ 188ms | ✓ 752ms | ✓ 659ms | http |
| 124.83.110.111:8085 | ✓ 1858ms | 否 | 否 | ✓ 1478ms | ✓ 1295ms | http |
| 49.48.68.15:8080 | ✓ 1555ms | 否 | ✓ 1941ms | 否 | ✓ 1506ms | http |
| 49.146.247.229:8080 | ✓ 1892ms | 否 | ✓ 1883ms | ✓ 1636ms | ✓ 1950ms | http |
| 54.37.72.89:80 | ✓ 552ms | ✓ 1891ms | ✓ 1870ms | 否 | 否 | http |
| 101.132.61.121:8888 | ✓ 1431ms | ✓ 1218ms | ✓ 1268ms | ✓ 1515ms | ✓ 1277ms | http |
| 1.32.48.93:8081 | ✓ 1932ms | 否 | ✓ 1212ms | ✓ 1579ms | ✓ 1308ms | http |
| 1.20.248.184:8080 | ✓ 1951ms | 否 | ✓ 1412ms | ✓ 1545ms | ✓ 1468ms | http |
| 101.32.244.83:8080 | ✓ 1045ms | ✓ 1558ms | ✓ 970ms | ✓ 1283ms | ✓ 1241ms | http |
| 121.43.196.213:8222 | ✓ 1030ms | ✓ 1079ms | ✓ 882ms | ✓ 1115ms | ✓ 979ms | http |
| 114.55.226.123:10086 | ✓ 1112ms | ✓ 1435ms | ✓ 1068ms | ✓ 1365ms | ✓ 1136ms | http |
| 178.128.24.162:8080 | ✓ 1368ms | 否 | ✓ 1028ms | ✓ 1098ms | ✓ 855ms | http |
| 159.223.225.118:8888 | ✓ 589ms | 否 | ✓ 1059ms | ✓ 1915ms | ✓ 1370ms | http |
| 165.22.57.158:8080 | ✓ 1363ms | 否 | ✓ 1137ms | ✓ 1254ms | ✓ 812ms | http |
| 38.145.208.209:8447 | ✓ 989ms | ✓ 1244ms | ✓ 1747ms | ✓ 1852ms | ✓ 618ms | http |
| 147.45.167.84:3128 | ✓ 1020ms | 否 | ✓ 926ms | ✓ 1872ms | ✓ 1749ms | http |
| 38.34.179.99:8449 | ✓ 566ms | 否 | ✓ 848ms | ✓ 1017ms | ✓ 1452ms | http |
| 38.145.218.229:8449 | ✓ 143ms | ✓ 1040ms | ✓ 1184ms | 否 | ✓ 1206ms | http |
| 61.76.95.217:40088 | ✓ 1027ms | ✓ 1310ms | ✓ 1408ms | ✓ 1353ms | ✓ 1778ms | http |
| 195.123.209.48:3128 | ✓ 1069ms | ✓ 1770ms | ✓ 1551ms | 否 | ✓ 1681ms | http |
| 43.167.237.94:3128 | ✓ 1523ms | 否 | ✓ 1415ms | ✓ 815ms | ✓ 1100ms | http |
| 210.223.44.230:3128 | ✓ 1497ms | ✓ 1196ms | ✓ 1653ms | 否 | 否 | http |
| 116.80.60.44:7777 | ✓ 1855ms | 否 | ✓ 1526ms | ✓ 1873ms | 否 | http |
| 47.74.226.8:5001 | ✓ 960ms | ✓ 1416ms | 否 | ✓ 1319ms | 否 | http |
| 148.135.22.89:1080 | ✓ 277ms | 否 | ✓ 1707ms | ✓ 1081ms | ✓ 1238ms | http |
| 152.32.132.190:7890 | ✓ 1646ms | ✓ 1986ms | ✓ 907ms | ✓ 928ms | ✓ 905ms | http |
| 171.43.65.238:7890 | ✓ 965ms | ✓ 1358ms | ✓ 1208ms | ✓ 1287ms | ✓ 968ms | http |
| 194.87.94.82:7890 | ✓ 1939ms | 否 | 否 | ✓ 1525ms | ✓ 1165ms | http |
| 223.205.76.72:8080 | 否 | 否 | ✓ 1916ms | ✓ 1551ms | ✓ 1511ms | http |
| 47.251.25.177:443 | ✓ 606ms | ✓ 1272ms | ✓ 138ms | ✓ 776ms | ✓ 638ms | http |
| 211.95.152.50:45046 | ✓ 1015ms | ✓ 1380ms | ✓ 1100ms | 否 | ✓ 1661ms | http |
| 146.19.56.212:7890 | ✓ 919ms | 否 | ✓ 1382ms | ✓ 1722ms | ✓ 594ms | http |
| 45.186.6.104:3128 | ✓ 1921ms | ✓ 1697ms | ✓ 1703ms | 否 | 否 | http |
| 123.57.2.231:2020 | ✓ 1164ms | ✓ 1198ms | ✓ 1068ms | ✓ 1229ms | ✓ 1138ms | http |
| 46.39.105.157:8080 | ✓ 1227ms | 否 | ✓ 1631ms | 否 | ✓ 1518ms | http |
| 124.221.164.32:7890 | 否 | 否 | ✓ 1780ms | ✓ 1233ms | ✓ 888ms | http |
| 194.67.119.10:10808 | ✓ 1413ms | ✓ 1977ms | ✓ 1821ms | 否 | 否 | http |
| 192.71.213.85:9091 | ✓ 775ms | 否 | ✓ 1754ms | ✓ 1883ms | 否 | http |
| 185.76.240.203:10001 | ✓ 1492ms | 否 | ✓ 1366ms | 否 | ✓ 1546ms | http |
| 185.76.240.61:10001 | ✓ 1486ms | 否 | ✓ 1375ms | 否 | ✓ 1661ms | http |
| 114.232.71.166:7890 | ✓ 981ms | ✓ 1151ms | ✓ 943ms | ✓ 1206ms | ✓ 939ms | http |
| 38.34.179.53:8451 | ✓ 485ms | ✓ 782ms | ✓ 1153ms | 否 | ✓ 615ms | http |
| 61.52.131.172:8443 | ✓ 994ms | ✓ 1184ms | ✓ 969ms | ✓ 1264ms | ✓ 993ms | http |
| 38.34.179.85:8444 | ✓ 1680ms | ✓ 781ms | 否 | ✓ 1436ms | ✓ 707ms | http |
| 38.34.179.94:8444 | ✓ 1679ms | ✓ 896ms | 否 | ✓ 1350ms | ✓ 723ms | http |
| 106.117.208.101:7890 | ✓ 1413ms | ✓ 1520ms | ✓ 1511ms | ✓ 1243ms | ✓ 1099ms | http |
| 103.39.51.207:8080 | ✓ 1338ms | 否 | 否 | ✓ 1444ms | ✓ 1784ms | http |

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
