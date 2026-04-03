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

最后更新：2026-04-03 16:43:26 UTC（2026-04-04 00:43:26 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 203.80.138.81:50000 | 否 | ✓ 1358ms | ✓ 1161ms | ✓ 1220ms | ✓ 1053ms | http |
| 167.103.115.102:8800 | ✓ 1090ms | 否 | ✓ 1164ms | ✓ 1202ms | ✓ 1219ms | http |
| 1.231.81.166:3128 | ✓ 1135ms | 否 | ✓ 1762ms | ✓ 1249ms | ✓ 1024ms | http |
| 159.223.71.162:8080 | ✓ 866ms | 否 | ✓ 1796ms | ✓ 1260ms | ✓ 1012ms | http |
| 159.223.71.162:443 | ✓ 864ms | 否 | ✓ 1791ms | ✓ 1272ms | ✓ 1052ms | http |
| 115.231.181.40:8128 | ✓ 1106ms | ✓ 1511ms | ✓ 1404ms | ✓ 1343ms | ✓ 1832ms | http |
| 111.227.254.9:22222 | 否 | ✓ 1553ms | 否 | ✓ 1556ms | ✓ 1183ms | http |
| 95.213.217.168:52004 | ✓ 553ms | ✓ 1515ms | ✓ 594ms | ✓ 1636ms | ✓ 1232ms | http |
| 218.108.131.186:17890 | ✓ 940ms | ✓ 1170ms | ✓ 1139ms | ✓ 1261ms | ✓ 1002ms | http |
| 38.145.220.72:8445 | ✓ 958ms | 否 | ✓ 1016ms | ✓ 1377ms | 否 | http |
| 38.34.179.61:8445 | ✓ 931ms | 否 | ✓ 1605ms | ✓ 1403ms | 否 | http |
| 38.34.179.98:8451 | ✓ 964ms | 否 | ✓ 1594ms | ✓ 1180ms | 否 | http |
| 38.145.208.215:8444 | ✓ 914ms | ✓ 1841ms | ✓ 1818ms | ✓ 975ms | 否 | http |
| 72.11.151.159:6005 | ✓ 317ms | ✓ 1336ms | ✓ 996ms | ✓ 1108ms | ✓ 1809ms | http |
| 113.160.132.26:8080 | ✓ 1005ms | ✓ 1883ms | ✓ 1092ms | ✓ 1362ms | ✓ 1141ms | http |
| 167.103.34.108:8800 | ✓ 1368ms | 否 | ✓ 1297ms | ✓ 1494ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1319ms | 否 | ✓ 1465ms | 否 | ✓ 1786ms | http |
| 38.145.208.207:8445 | 否 | ✓ 933ms | ✓ 1017ms | 否 | ✓ 928ms | http |
| 38.145.208.214:8446 | 否 | ✓ 893ms | ✓ 1036ms | ✓ 1897ms | ✓ 1630ms | http |
| 147.161.210.140:8800 | ✓ 1375ms | 否 | ✓ 1408ms | ✓ 1385ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1731ms | 否 | ✓ 1508ms | ✓ 1678ms | 否 | http |
| 111.227.254.12:22222 | ✓ 1084ms | ✓ 1755ms | 否 | ✓ 1775ms | ✓ 1869ms | http |
| 38.34.179.178:8444 | 否 | ✓ 1693ms | 否 | ✓ 1735ms | ✓ 1726ms | http |
| 167.103.31.122:8800 | ✓ 1397ms | 否 | 否 | ✓ 1844ms | ✓ 1893ms | http |
| 222.184.48.251:22222 | ✓ 1430ms | ✓ 1238ms | 否 | ✓ 1356ms | ✓ 1102ms | http |
| 86.53.183.16:1080 | ✓ 1973ms | ✓ 1870ms | ✓ 1603ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1078ms | 否 | ✓ 1539ms | ✓ 1840ms | ✓ 1995ms | http |
| 120.92.212.16:8890 | ✓ 1103ms | ✓ 1350ms | ✓ 1126ms | 否 | 否 | http |
| 192.3.181.90:1234 | 否 | ✓ 1587ms | 否 | ✓ 1252ms | ✓ 761ms | http |
| 45.136.131.58:8445 | ✓ 1870ms | 否 | ✓ 1896ms | ✓ 1433ms | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1429ms | ✓ 1115ms | ✓ 968ms | 否 | http |
| 208.87.243.199:7878 | ✓ 1152ms | ✓ 1343ms | ✓ 1933ms | ✓ 1233ms | ✓ 1053ms | http |
| 209.126.84.232:8888 | ✓ 1692ms | 否 | ✓ 1186ms | 否 | ✓ 1893ms | http |
| 177.234.217.88:999 | ✓ 1248ms | 否 | ✓ 1734ms | ✓ 1836ms | ✓ 1584ms | http |
| 147.161.239.240:8800 | ✓ 1040ms | 否 | ✓ 846ms | ✓ 1637ms | ✓ 1307ms | http |
| 64.227.76.27:1080 | ✓ 1140ms | 否 | ✓ 712ms | ✓ 1283ms | ✓ 1941ms | http |
| 45.12.151.226:2829 | 否 | ✓ 1860ms | ✓ 947ms | ✓ 1782ms | ✓ 1316ms | http |
| 34.96.238.40:8080 | ✓ 1260ms | ✓ 1236ms | ✓ 1170ms | 否 | 否 | http |
| 38.34.179.105:8449 | ✓ 574ms | ✓ 1980ms | ✓ 1609ms | ✓ 968ms | ✓ 1040ms | http |
| 45.136.130.181:8445 | ✓ 575ms | ✓ 1362ms | 否 | ✓ 966ms | 否 | http |
| 38.145.218.216:8449 | ✓ 578ms | 否 | ✓ 1692ms | ✓ 900ms | ✓ 1739ms | http |
| 212.58.132.5:8888 | ✓ 1570ms | 否 | 否 | ✓ 1501ms | ✓ 1249ms | http |
| 38.34.179.79:8451 | ✓ 1280ms | 否 | ✓ 444ms | ✓ 1864ms | ✓ 1879ms | http |
| 62.234.206.73:3128 | ✓ 1951ms | 否 | 否 | ✓ 1361ms | ✓ 1384ms | http |
| 222.184.48.242:22222 | ✓ 1844ms | 否 | ✓ 1057ms | ✓ 1315ms | 否 | http |
| 222.184.48.252:22222 | ✓ 1061ms | 否 | 否 | ✓ 1483ms | ✓ 1141ms | http |
| 165.232.146.249:3128 | ✓ 600ms | 否 | ✓ 1029ms | 否 | ✓ 638ms | http |
| 114.231.73.92:1080 | ✓ 1407ms | ✓ 1511ms | ✓ 1085ms | 否 | 否 | http |
| 34.101.184.164:3128 | ✓ 1009ms | 否 | ✓ 1715ms | 否 | ✓ 1148ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1500ms | ✓ 1543ms | 否 | ✓ 1183ms | http |
| 38.34.179.14:8451 | 否 | ✓ 940ms | ✓ 1982ms | ✓ 1835ms | ✓ 1281ms | http |
| 103.113.70.189:1081 | ✓ 768ms | ✓ 1034ms | ✓ 164ms | ✓ 1432ms | ✓ 922ms | http |
| 38.145.208.242:8451 | 否 | ✓ 1643ms | 否 | ✓ 1282ms | ✓ 1202ms | http |
| 45.136.130.195:8446 | ✓ 890ms | ✓ 1760ms | ✓ 1009ms | ✓ 907ms | ✓ 1564ms | http |
| 38.145.208.181:8445 | ✓ 1780ms | 否 | ✓ 329ms | ✓ 1450ms | 否 | http |
| 38.145.208.235:8452 | ✓ 1773ms | 否 | ✓ 1951ms | ✓ 878ms | 否 | http |
| 46.39.105.157:8080 | 否 | 否 | ✓ 1533ms | ✓ 1595ms | ✓ 1444ms | http |
| 38.34.183.164:8444 | ✓ 747ms | 否 | ✓ 1096ms | ✓ 1346ms | 否 | http |
| 38.145.208.210:8448 | ✓ 773ms | 否 | ✓ 1654ms | ✓ 973ms | 否 | http |
| 144.124.232.204:3128 | ✓ 1045ms | 否 | ✓ 1265ms | ✓ 1741ms | 否 | http |
| 103.133.254.4:3128 | ✓ 1604ms | 否 | ✓ 1735ms | ✓ 1834ms | ✓ 1493ms | http |
| 38.145.220.168:8444 | ✓ 1060ms | ✓ 1615ms | ✓ 1070ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1221ms | ✓ 1080ms | ✓ 1033ms | ✓ 1077ms | ✓ 973ms | http |
| 38.145.218.161:8445 | ✓ 1159ms | 否 | ✓ 1176ms | ✓ 970ms | ✓ 1288ms | http |
| 45.136.130.197:8452 | ✓ 1757ms | ✓ 1225ms | ✓ 453ms | ✓ 1710ms | ✓ 1804ms | http |
| 101.43.127.100:8877 | ✓ 1911ms | 否 | 否 | ✓ 1565ms | ✓ 979ms | http |
| 45.140.147.155:1081 | ✓ 718ms | ✓ 1595ms | ✓ 754ms | ✓ 1706ms | 否 | http |
| 193.148.58.165:3128 | ✓ 1308ms | 否 | ✓ 1348ms | 否 | ✓ 1586ms | http |
| 129.213.162.27:17777 | ✓ 402ms | ✓ 1867ms | 否 | 否 | ✓ 1724ms | http |
| 195.123.209.48:3128 | ✓ 747ms | ✓ 1731ms | ✓ 1141ms | 否 | 否 | http |
| 157.0.142.246:10061 | ✓ 1101ms | ✓ 1434ms | ✓ 1135ms | ✓ 1471ms | ✓ 1177ms | http |
| 157.0.142.246:10057 | ✓ 1141ms | ✓ 1369ms | ✓ 1173ms | ✓ 1498ms | ✓ 1175ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 963ms | ✓ 1943ms | ✓ 1269ms | http |
| 38.34.179.34:8445 | ✓ 1298ms | ✓ 1600ms | ✓ 839ms | ✓ 1973ms | ✓ 932ms | http |
| 38.34.179.35:8444 | ✓ 1446ms | 否 | ✓ 1762ms | ✓ 1503ms | 否 | http |
| 45.136.130.246:8445 | ✓ 1741ms | 否 | ✓ 656ms | 否 | ✓ 1047ms | http |
| 38.145.220.188:8451 | 否 | ✓ 1923ms | ✓ 1785ms | ✓ 1140ms | ✓ 1410ms | http |
| 150.241.71.15:1080 | ✓ 1946ms | 否 | ✓ 507ms | ✓ 1925ms | 否 | http |
| 83.219.250.8:62920 | ✓ 1014ms | ✓ 1385ms | 否 | 否 | ✓ 1695ms | http |
| 43.99.54.236:5555 | ✓ 1839ms | ✓ 1203ms | ✓ 777ms | ✓ 1036ms | ✓ 769ms | http |

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
