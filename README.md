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

最后更新：2026-05-19 10:43:39 UTC（2026-05-19 18:43:39 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.130.126.146:6688 | ✓ 286ms | ✓ 839ms | ✓ 1012ms | ✓ 1592ms | 否 | http |
| 192.99.8.15:8850 | ✓ 301ms | 否 | ✓ 1266ms | ✓ 1218ms | ✓ 1089ms | http |
| 174.137.134.182:2999 | ✓ 310ms | ✓ 1972ms | 否 | ✓ 1110ms | ✓ 1097ms | http |
| 104.168.96.172:1888 | ✓ 729ms | ✓ 1072ms | ✓ 1118ms | ✓ 1090ms | ✓ 1178ms | http |
| 89.58.50.94:11140 | ✓ 623ms | ✓ 1655ms | ✓ 1361ms | ✓ 1875ms | ✓ 1522ms | http |
| 185.200.188.234:10001 | ✓ 946ms | 否 | ✓ 993ms | 否 | ✓ 1568ms | http |
| 176.111.37.5:39811 | ✓ 788ms | ✓ 1411ms | ✓ 1824ms | 否 | ✓ 1923ms | http |
| 176.111.37.216:39811 | ✓ 866ms | 否 | ✓ 1504ms | 否 | ✓ 1680ms | http |
| 47.52.134.234:36463 | ✓ 1726ms | 否 | 否 | ✓ 1063ms | ✓ 840ms | http |
| 113.160.132.26:8080 | ✓ 1762ms | 否 | ✓ 1898ms | ✓ 1926ms | ✓ 1856ms | http |
| 3.15.187.17:1080 | ✓ 1077ms | ✓ 1564ms | ✓ 1297ms | ✓ 1405ms | ✓ 1149ms | http |
| 45.95.203.47:6699 | ✓ 556ms | ✓ 1816ms | ✓ 610ms | ✓ 1599ms | ✓ 1528ms | http |
| 43.154.90.238:9527 | 否 | 否 | ✓ 1095ms | ✓ 1203ms | ✓ 1299ms | http |
| 47.237.181.222:8100 | ✓ 1162ms | 否 | 否 | ✓ 1336ms | ✓ 1079ms | http |
| 170.106.136.181:31002 | ✓ 881ms | ✓ 1249ms | 否 | 否 | ✓ 1683ms | http |
| 45.125.67.37:8443 | ✓ 1172ms | 否 | ✓ 1602ms | ✓ 1322ms | ✓ 1360ms | http |
| 115.231.181.40:8128 | ✓ 1461ms | 否 | ✓ 1126ms | ✓ 1494ms | 否 | http |
| 34.87.80.221:30000 | ✓ 1521ms | ✓ 1971ms | ✓ 1554ms | ✓ 1575ms | 否 | http |
| 5.252.33.13:2025 | ✓ 1760ms | ✓ 1940ms | ✓ 1964ms | 否 | 否 | http |
| 202.28.194.139:31280 | ✓ 1739ms | 否 | 否 | ✓ 1923ms | ✓ 1896ms | http |
| 161.117.86.53:8100 | ✓ 1773ms | ✓ 1673ms | ✓ 1162ms | ✓ 1545ms | ✓ 1337ms | http |
| 152.67.191.232:6800 | ✓ 1388ms | 否 | ✓ 1342ms | ✓ 1637ms | 否 | http |
| 138.2.78.251:8100 | 否 | 否 | ✓ 1288ms | ✓ 1712ms | ✓ 1776ms | http |
| 114.214.163.108:6789 | ✓ 1287ms | ✓ 1647ms | ✓ 1682ms | ✓ 1638ms | ✓ 1327ms | http |
| 128.199.113.85:9090 | 否 | 否 | ✓ 1453ms | ✓ 1488ms | ✓ 1166ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 977ms | ✓ 1381ms | ✓ 1170ms | http |
| 147.45.41.112:1080 | ✓ 988ms | 否 | ✓ 947ms | ✓ 1581ms | ✓ 1431ms | http |
| 128.199.121.61:9090 | ✓ 1608ms | 否 | ✓ 1010ms | ✓ 1324ms | ✓ 1045ms | http |
| 128.199.254.13:9090 | ✓ 1604ms | 否 | ✓ 1253ms | ✓ 1454ms | ✓ 1062ms | http |
| 148.230.4.241:999 | ✓ 1816ms | ✓ 1880ms | ✓ 646ms | ✓ 1666ms | ✓ 1478ms | http |
| 122.2.48.121:8080 | ✓ 1533ms | 否 | ✓ 1452ms | ✓ 1546ms | ✓ 1557ms | http |
| 168.138.171.204:8100 | ✓ 1603ms | 否 | 否 | ✓ 1888ms | ✓ 1471ms | http |
| 152.32.132.190:7890 | ✓ 1072ms | ✓ 1892ms | ✓ 867ms | ✓ 1709ms | ✓ 1626ms | http |
| 84.47.150.125:1080 | ✓ 638ms | 否 | ✓ 1525ms | ✓ 1933ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1433ms | ✓ 1815ms | ✓ 1685ms | 否 | 否 | http |
| 138.2.92.70:8100 | 否 | 否 | ✓ 1385ms | ✓ 1792ms | ✓ 1307ms | http |
| 45.129.141.143:3128 | ✓ 707ms | ✓ 1713ms | ✓ 1602ms | ✓ 1753ms | ✓ 1912ms | http |
| 158.160.215.167:8126 | ✓ 612ms | ✓ 1429ms | ✓ 764ms | 否 | ✓ 1929ms | http |
| 57.129.144.178:40000 | ✓ 863ms | 否 | ✓ 1261ms | ✓ 1530ms | ✓ 1647ms | http |
| 57.128.188.167:9634 | ✓ 911ms | 否 | 否 | ✓ 1982ms | ✓ 1331ms | http |
| 158.255.212.55:7497 | ✓ 926ms | 否 | ✓ 1134ms | ✓ 1843ms | 否 | http |
| 8.154.21.175:3128 | ✓ 1086ms | ✓ 1331ms | 否 | ✓ 1360ms | ✓ 1131ms | http |
| 158.247.250.156:9898 | ✓ 1549ms | 否 | ✓ 1455ms | ✓ 1100ms | ✓ 973ms | http |
| 57.128.188.167:9172 | ✓ 1767ms | 否 | ✓ 1703ms | 否 | ✓ 1892ms | http |
| 1.231.81.166:3128 | ✓ 1991ms | ✓ 1752ms | ✓ 1273ms | ✓ 1526ms | ✓ 1543ms | http |
| 8.218.153.104:8100 | ✓ 1854ms | 否 | ✓ 1217ms | 否 | ✓ 1188ms | http |
| 103.210.160.62:7789 | 否 | ✓ 1367ms | 否 | ✓ 1551ms | ✓ 1202ms | http |
| 34.84.162.206:38080 | ✓ 1089ms | ✓ 1284ms | 否 | ✓ 1453ms | ✓ 1005ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1652ms | ✓ 1688ms | ✓ 1702ms | 否 | http |
| 38.180.2.107:3128 | ✓ 1578ms | ✓ 1603ms | ✓ 1880ms | 否 | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1842ms | ✓ 1241ms | 否 | ✓ 1809ms | http |
| 31.172.78.12:3128 | 否 | 否 | ✓ 1658ms | ✓ 1402ms | ✓ 1449ms | http |
| 78.186.118.164:3311 | ✓ 1286ms | ✓ 1785ms | ✓ 1897ms | ✓ 1997ms | ✓ 1846ms | http |
| 114.214.165.78:10810 | 否 | 否 | ✓ 1347ms | ✓ 1635ms | ✓ 1358ms | http |
| 81.30.156.115:8080 | ✓ 1093ms | 否 | ✓ 1139ms | ✓ 1723ms | ✓ 1309ms | http |
| 51.79.71.106:8080 | ✓ 1677ms | 否 | 否 | ✓ 1610ms | ✓ 1365ms | http |
| 74.208.192.81:3129 | ✓ 1376ms | 否 | ✓ 227ms | ✓ 1082ms | ✓ 1199ms | http |
| 8.219.97.248:80 | ✓ 1644ms | 否 | ✓ 1471ms | ✓ 1482ms | 否 | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 1890ms | ✓ 1525ms | ✓ 1093ms | http |
| 159.223.41.216:9090 | ✓ 964ms | 否 | ✓ 1163ms | ✓ 1312ms | ✓ 1023ms | http |
| 152.70.91.193:40000 | ✓ 1699ms | 否 | ✓ 1718ms | ✓ 1954ms | ✓ 1725ms | http |
| 3.101.133.120:80 | ✓ 794ms | ✓ 1544ms | ✓ 881ms | ✓ 1410ms | ✓ 1223ms | http |
| 207.148.124.152:6868 | ✓ 1060ms | 否 | ✓ 1143ms | ✓ 1506ms | ✓ 1878ms | http |
| 120.92.212.16:8890 | ✓ 1225ms | ✓ 1497ms | 否 | 否 | ✓ 1103ms | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 990ms | ✓ 1381ms | ✓ 1048ms | http |
| 146.190.80.158:9090 | ✓ 1459ms | 否 | ✓ 980ms | ✓ 1333ms | ✓ 1058ms | http |
| 178.63.155.151:8898 | ✓ 1171ms | 否 | ✓ 1799ms | ✓ 1545ms | ✓ 1372ms | http |
| 45.59.122.132:80 | ✓ 434ms | 否 | ✓ 733ms | ✓ 1529ms | ✓ 1235ms | http |
| 88.248.130.191:3311 | ✓ 1546ms | 否 | ✓ 1698ms | 否 | ✓ 1944ms | http |
| 20.164.75.153:8080 | ✓ 1955ms | 否 | ✓ 1824ms | ✓ 1963ms | ✓ 1821ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1549ms | ✓ 1151ms | ✓ 800ms | http |
| 158.255.212.55:7839 | ✓ 1100ms | 否 | ✓ 1725ms | ✓ 1758ms | 否 | http |
| 158.255.212.55:9005 | ✓ 1108ms | 否 | ✓ 1723ms | ✓ 1762ms | 否 | http |
| 158.255.212.55:3256 | ✓ 1108ms | 否 | ✓ 1722ms | ✓ 1763ms | 否 | http |
| 158.255.212.55:9480 | ✓ 1106ms | 否 | ✓ 1724ms | ✓ 1765ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1699ms | 否 | ✓ 1556ms | ✓ 1532ms | ✓ 1242ms | http |
| 104.248.151.93:9090 | 否 | 否 | ✓ 1024ms | ✓ 1364ms | ✓ 1032ms | http |
| 62.113.119.14:8080 | ✓ 1296ms | 否 | ✓ 1246ms | ✓ 1789ms | ✓ 1525ms | http |
| 82.114.228.67:1080 | 否 | ✓ 1680ms | ✓ 902ms | ✓ 1679ms | 否 | http |
| 158.160.215.167:8124 | ✓ 920ms | 否 | ✓ 1849ms | 否 | ✓ 1730ms | http |
| 147.45.186.28:3128 | ✓ 894ms | ✓ 1454ms | ✓ 1290ms | ✓ 1912ms | 否 | http |
| 192.81.129.252:3136 | 否 | ✓ 1347ms | ✓ 1282ms | 否 | ✓ 1854ms | http |
| 103.69.84.106:3131 | 否 | 否 | ✓ 1495ms | ✓ 1368ms | ✓ 1104ms | http |
| 61.52.131.172:8443 | ✓ 1161ms | ✓ 1449ms | ✓ 1156ms | ✓ 1485ms | ✓ 1199ms | http |

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
