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

最后更新：2026-05-18 13:28:46 UTC（2026-05-18 21:28:46 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 1063ms | ✓ 1305ms | ✓ 1060ms | ✓ 1331ms | ✓ 1088ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1807ms | ✓ 1540ms | 否 | ✓ 1320ms | http |
| 170.106.136.181:31002 | ✓ 788ms | ✓ 950ms | ✓ 420ms | ✓ 862ms | ✓ 666ms | http |
| 46.39.105.157:8080 | ✓ 662ms | 否 | ✓ 1239ms | ✓ 1757ms | ✓ 1332ms | http |
| 1.231.81.166:3128 | ✓ 1751ms | 否 | ✓ 1510ms | ✓ 1252ms | ✓ 1001ms | http |
| 185.200.188.234:10001 | ✓ 946ms | 否 | ✓ 1777ms | 否 | ✓ 1573ms | http |
| 103.134.85.167:3128 | 否 | 否 | ✓ 1189ms | ✓ 1574ms | ✓ 1238ms | http |
| 51.161.50.166:3128 | 否 | 否 | ✓ 1739ms | ✓ 1672ms | ✓ 1507ms | http |
| 103.35.190.69:1082 | ✓ 794ms | ✓ 1870ms | ✓ 1623ms | 否 | 否 | http |
| 119.28.51.157:3128 | ✓ 1082ms | ✓ 1521ms | ✓ 1160ms | ✓ 1272ms | ✓ 1319ms | http |
| 190.12.150.244:999 | ✓ 883ms | ✓ 1569ms | ✓ 883ms | ✓ 1625ms | ✓ 1351ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1622ms | ✓ 1626ms | ✓ 1249ms | http |
| 115.231.181.40:8128 | ✓ 1257ms | 否 | 否 | ✓ 1696ms | ✓ 1156ms | http |
| 20.164.75.153:8080 | ✓ 1193ms | 否 | ✓ 1159ms | 否 | ✓ 1828ms | http |
| 86.104.72.220:1082 | ✓ 416ms | ✓ 1209ms | ✓ 1704ms | 否 | 否 | http |
| 129.80.217.21:444 | ✓ 268ms | ✓ 1074ms | ✓ 858ms | ✓ 1146ms | ✓ 675ms | http |
| 128.199.121.61:9090 | 否 | 否 | ✓ 1809ms | ✓ 1377ms | ✓ 1098ms | http |
| 129.80.238.83:444 | 否 | ✓ 937ms | ✓ 856ms | ✓ 894ms | ✓ 669ms | http |
| 152.42.170.187:9090 | ✓ 904ms | 否 | 否 | ✓ 1771ms | ✓ 1056ms | http |
| 128.199.114.189:9090 | ✓ 962ms | 否 | ✓ 1114ms | ✓ 1360ms | ✓ 1058ms | http |
| 148.230.4.241:999 | ✓ 1114ms | 否 | ✓ 603ms | ✓ 1355ms | ✓ 1195ms | http |
| 207.254.71.62:8088 | ✓ 537ms | ✓ 1728ms | ✓ 1550ms | 否 | 否 | http |
| 8.154.21.175:3128 | ✓ 1060ms | ✓ 1281ms | ✓ 1073ms | ✓ 1369ms | ✓ 1158ms | http |
| 103.147.152.12:1095 | ✓ 817ms | ✓ 1485ms | ✓ 1112ms | 否 | ✓ 1173ms | http |
| 64.227.174.131:9090 | ✓ 1859ms | 否 | ✓ 1319ms | 否 | ✓ 1430ms | http |
| 106.10.55.212:1121 | ✓ 1008ms | ✓ 1595ms | 否 | ✓ 1760ms | ✓ 1243ms | http |
| 116.80.79.244:3172 | ✓ 1559ms | 否 | 否 | ✓ 1088ms | ✓ 910ms | http |
| 84.47.150.125:1080 | 否 | 否 | ✓ 990ms | ✓ 1812ms | ✓ 1841ms | http |
| 210.223.44.230:3128 | ✓ 1596ms | ✓ 1505ms | ✓ 1544ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1110ms | 否 | ✓ 1353ms | ✓ 1921ms | ✓ 1235ms | http |
| 158.160.215.167:8123 | ✓ 1120ms | ✓ 1814ms | ✓ 996ms | 否 | ✓ 1744ms | http |
| 158.160.215.167:8124 | ✓ 866ms | 否 | ✓ 1881ms | 否 | ✓ 1759ms | http |
| 107.175.85.198:1080 | ✓ 1509ms | ✓ 1208ms | ✓ 699ms | 否 | 否 | http |
| 173.212.246.157:3128 | ✓ 933ms | 否 | ✓ 892ms | ✓ 1985ms | ✓ 1660ms | http |
| 181.119.97.24:999 | ✓ 1151ms | ✓ 1729ms | ✓ 1753ms | 否 | 否 | http |
| 192.73.243.65:3128 | ✓ 927ms | 否 | 否 | ✓ 1532ms | ✓ 718ms | http |
| 121.230.8.138:1080 | ✓ 1319ms | 否 | ✓ 1217ms | ✓ 1450ms | ✓ 1279ms | http |
| 103.147.152.12:1080 | 否 | ✓ 1405ms | ✓ 1421ms | 否 | ✓ 1150ms | http |
| 185.21.15.206:3128 | 否 | 否 | ✓ 501ms | ✓ 1633ms | ✓ 1659ms | http |
| 103.53.79.115:8050 | ✓ 1359ms | 否 | 否 | ✓ 1712ms | ✓ 1643ms | http |
| 5.252.33.13:2025 | ✓ 1418ms | 否 | ✓ 1860ms | 否 | ✓ 1970ms | http |
| 34.101.184.164:3128 | ✓ 1770ms | 否 | ✓ 1180ms | ✓ 1394ms | ✓ 1358ms | http |
| 8.219.97.248:80 | ✓ 1299ms | 否 | 否 | ✓ 1788ms | ✓ 1538ms | http |
| 128.199.116.219:9090 | ✓ 931ms | 否 | 否 | ✓ 1349ms | ✓ 1066ms | http |
| 128.199.113.85:9090 | ✓ 1195ms | 否 | ✓ 1167ms | ✓ 1360ms | ✓ 1038ms | http |
| 101.32.243.189:80 | ✓ 1404ms | ✓ 1523ms | 否 | ✓ 1771ms | ✓ 1549ms | http |
| 121.130.177.28:8888 | ✓ 1587ms | 否 | 否 | ✓ 1801ms | ✓ 1740ms | http |
| 43.167.192.85:8080 | ✓ 1277ms | ✓ 1229ms | ✓ 874ms | ✓ 1129ms | ✓ 992ms | http |
| 160.250.4.13:1 | ✓ 1644ms | 否 | ✓ 1756ms | 否 | ✓ 1677ms | http |
| 3.101.133.120:80 | ✓ 826ms | ✓ 1654ms | ✓ 1050ms | ✓ 1376ms | ✓ 1122ms | http |
| 77.110.107.80:8080 | ✓ 705ms | ✓ 1461ms | ✓ 1143ms | ✓ 1799ms | 否 | http |
| 77.110.107.80:1080 | ✓ 747ms | 否 | ✓ 968ms | ✓ 1438ms | 否 | http |
| 168.110.52.228:3128 | ✓ 1584ms | 否 | ✓ 660ms | 否 | ✓ 818ms | http |
| 114.214.165.78:10810 | ✓ 1397ms | 否 | ✓ 1464ms | ✓ 1695ms | ✓ 1338ms | http |
| 45.174.175.26:999 | ✓ 1011ms | 否 | ✓ 1098ms | ✓ 1466ms | ✓ 1476ms | http |
| 158.160.215.167:8126 | ✓ 1436ms | ✓ 1800ms | ✓ 1972ms | 否 | 否 | http |
| 27.254.99.183:8118 | ✓ 1085ms | 否 | ✓ 1287ms | 否 | ✓ 1326ms | http |
| 57.129.144.178:40000 | ✓ 588ms | 否 | ✓ 1147ms | ✓ 1847ms | ✓ 1498ms | http |
| 20.120.225.109:3128 | ✓ 1124ms | ✓ 1803ms | ✓ 1840ms | 否 | 否 | http |
| 178.63.155.151:8898 | ✓ 1391ms | 否 | ✓ 1434ms | ✓ 1993ms | 否 | http |
| 20.27.11.248:8561 | ✓ 1509ms | ✓ 1409ms | ✓ 717ms | ✓ 1072ms | ✓ 790ms | http |
| 20.27.14.220:8561 | ✓ 1506ms | 否 | ✓ 641ms | ✓ 980ms | ✓ 774ms | http |
| 221.122.91.36:11273 | ✓ 1115ms | ✓ 1405ms | ✓ 1180ms | ✓ 1413ms | ✓ 1135ms | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1260ms | ✓ 1303ms | ✓ 1040ms | http |
| 146.190.80.158:9090 | ✓ 968ms | 否 | 否 | ✓ 1731ms | ✓ 1072ms | http |
| 104.248.151.93:9090 | ✓ 951ms | 否 | ✓ 1395ms | ✓ 1311ms | ✓ 1050ms | http |
| 20.27.15.111:8561 | ✓ 1513ms | ✓ 1522ms | ✓ 991ms | ✓ 1013ms | ✓ 825ms | http |
| 20.78.118.91:8561 | ✓ 1522ms | ✓ 1538ms | ✓ 732ms | ✓ 1069ms | ✓ 846ms | http |
| 20.210.39.153:8561 | ✓ 1529ms | 否 | ✓ 686ms | ✓ 1054ms | ✓ 848ms | http |
| 20.78.26.206:8561 | ✓ 1532ms | 否 | ✓ 676ms | ✓ 1081ms | ✓ 850ms | http |
| 146.56.52.57:20201 | ✓ 1805ms | 否 | ✓ 1730ms | 否 | ✓ 1639ms | http |
| 139.198.113.42:10023 | 否 | ✓ 1991ms | 否 | ✓ 1582ms | ✓ 1457ms | http |
| 150.107.140.238:3128 | ✓ 1538ms | 否 | 否 | ✓ 1943ms | ✓ 1366ms | http |
| 203.150.113.164:8080 | ✓ 1908ms | 否 | 否 | ✓ 1741ms | ✓ 1956ms | http |
| 152.42.177.32:8888 | ✓ 1690ms | 否 | ✓ 920ms | ✓ 1578ms | 否 | http |
| 159.223.41.216:9090 | ✓ 1690ms | 否 | ✓ 1596ms | ✓ 1325ms | ✓ 1034ms | http |
| 61.52.131.172:8443 | ✓ 1160ms | ✓ 1436ms | ✓ 1115ms | ✓ 1422ms | ✓ 1145ms | http |
| 158.255.212.55:3256 | ✓ 1393ms | 否 | ✓ 1344ms | ✓ 1688ms | 否 | http |
| 158.255.212.55:7839 | ✓ 1388ms | 否 | ✓ 1339ms | ✓ 1701ms | 否 | http |
| 158.255.212.55:9005 | ✓ 1397ms | 否 | ✓ 1339ms | ✓ 1701ms | 否 | http |
| 158.255.212.55:9480 | ✓ 1393ms | 否 | ✓ 1344ms | ✓ 1703ms | 否 | http |
| 158.255.212.55:7497 | ✓ 1393ms | 否 | ✓ 1339ms | ✓ 1707ms | 否 | http |
| 20.210.76.178:8561 | 否 | ✓ 1599ms | ✓ 1148ms | ✓ 1207ms | ✓ 1058ms | http |
| 20.210.76.104:8561 | ✓ 1395ms | ✓ 1691ms | ✓ 707ms | ✓ 1380ms | ✓ 912ms | http |
| 20.210.76.175:8561 | ✓ 939ms | ✓ 1158ms | ✓ 771ms | ✓ 1064ms | ✓ 920ms | http |
| 47.84.131.156:8100 | ✓ 1173ms | ✓ 1930ms | ✓ 922ms | ✓ 1307ms | 否 | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1432ms | ✓ 1252ms | ✓ 1274ms | http |
| 139.135.170.12:8082 | ✓ 1692ms | 否 | 否 | ✓ 1713ms | ✓ 1623ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1527ms | 否 | ✓ 1465ms | ✓ 1237ms | http |

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
