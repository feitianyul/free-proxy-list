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

最后更新：2026-04-14 06:19:55 UTC（2026-04-14 14:19:55 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 157.230.178.216:8088 | ✓ 779ms | ✓ 1627ms | ✓ 1121ms | 否 | ✓ 923ms | http |
| 147.161.210.140:8800 | ✓ 1187ms | ✓ 1398ms | ✓ 1314ms | ✓ 1467ms | ✓ 970ms | http |
| 94.131.118.39:1081 | ✓ 1080ms | 否 | ✓ 1448ms | ✓ 1511ms | ✓ 1210ms | http |
| 46.101.126.84:8888 | ✓ 1670ms | 否 | 否 | ✓ 1286ms | ✓ 915ms | http |
| 129.154.48.5:1080 | ✓ 1223ms | ✓ 1497ms | ✓ 1788ms | ✓ 1351ms | ✓ 931ms | http |
| 43.156.132.113:3128 | ✓ 1677ms | ✓ 1569ms | ✓ 1670ms | ✓ 1650ms | ✓ 1374ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1471ms | ✓ 1425ms | ✓ 1738ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1941ms | ✓ 1603ms | ✓ 1803ms | ✓ 1169ms | http |
| 78.11.96.22:8888 | ✓ 1698ms | 否 | ✓ 969ms | ✓ 1452ms | ✓ 1259ms | http |
| 34.71.229.255:3128 | 否 | ✓ 1419ms | ✓ 1474ms | ✓ 1196ms | ✓ 804ms | http |
| 8.219.195.129:1080 | ✓ 1034ms | ✓ 1986ms | ✓ 1085ms | ✓ 1301ms | ✓ 993ms | http |
| 13.61.155.35:9015 | ✓ 860ms | 否 | 否 | ✓ 1882ms | ✓ 1462ms | http |
| 35.180.127.14:8811 | ✓ 1610ms | 否 | ✓ 1380ms | 否 | ✓ 1741ms | http |
| 167.103.115.102:8800 | ✓ 1125ms | 否 | ✓ 1351ms | ✓ 1249ms | ✓ 1795ms | http |
| 2.27.18.184:1080 | ✓ 1777ms | 否 | ✓ 1447ms | ✓ 1689ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1690ms | 否 | ✓ 1464ms | 否 | ✓ 1920ms | http |
| 45.167.125.21:999 | ✓ 882ms | ✓ 1607ms | ✓ 1285ms | ✓ 1558ms | ✓ 1270ms | http |
| 116.80.83.23:3128 | ✓ 1714ms | 否 | ✓ 1680ms | 否 | ✓ 1774ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 219ms | ✓ 1200ms | ✓ 951ms | http |
| 147.161.239.240:8800 | ✓ 872ms | ✓ 1642ms | ✓ 1021ms | ✓ 1288ms | ✓ 1122ms | http |
| 138.124.99.216:8888 | ✓ 899ms | ✓ 1486ms | ✓ 1704ms | ✓ 1689ms | ✓ 1674ms | http |
| 167.103.31.122:8800 | ✓ 1562ms | 否 | ✓ 1281ms | ✓ 1987ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1309ms | 否 | 否 | ✓ 1527ms | ✓ 1521ms | http |
| 3.71.175.73:3129 | ✓ 918ms | 否 | ✓ 584ms | ✓ 1924ms | 否 | http |
| 52.59.51.29:3129 | ✓ 927ms | 否 | ✓ 632ms | ✓ 1877ms | ✓ 1236ms | http |
| 99.79.58.74:3129 | ✓ 941ms | 否 | ✓ 1129ms | 否 | ✓ 1593ms | http |
| 52.47.115.41:6957 | ✓ 938ms | 否 | 否 | ✓ 1439ms | ✓ 1171ms | http |
| 15.160.132.166:3129 | ✓ 929ms | 否 | ✓ 1003ms | ✓ 1910ms | ✓ 1651ms | http |
| 52.59.218.12:3129 | ✓ 939ms | 否 | 否 | ✓ 1682ms | ✓ 1235ms | http |
| 101.32.243.189:80 | ✓ 1391ms | 否 | ✓ 1766ms | 否 | ✓ 1532ms | http |
| 85.239.59.252:7890 | ✓ 951ms | ✓ 1355ms | 否 | ✓ 1885ms | 否 | http |
| 137.59.47.73:3128 | ✓ 1657ms | 否 | ✓ 1657ms | ✓ 1472ms | ✓ 1675ms | http |
| 171.227.167.109:1009 | ✓ 1676ms | 否 | ✓ 1153ms | 否 | ✓ 1747ms | http |
| 5.255.123.43:1080 | ✓ 1398ms | ✓ 1924ms | 否 | 否 | ✓ 1319ms | http |
| 185.191.236.162:3128 | ✓ 1859ms | 否 | ✓ 1962ms | 否 | ✓ 1174ms | http |
| 185.114.73.2:1080 | ✓ 1108ms | ✓ 1384ms | ✓ 1321ms | 否 | 否 | http |
| 108.131.109.106:3129 | ✓ 600ms | 否 | ✓ 1546ms | ✓ 1747ms | ✓ 1430ms | http |
| 45.140.147.155:1081 | 否 | ✓ 1505ms | ✓ 1273ms | ✓ 1526ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1876ms | ✓ 1265ms | ✓ 942ms | ✓ 1128ms | ✓ 1033ms | http |
| 45.149.92.147:5001 | ✓ 881ms | 否 | ✓ 799ms | ✓ 1142ms | ✓ 1237ms | http |
| 177.234.217.88:999 | ✓ 1106ms | 否 | ✓ 1058ms | ✓ 1721ms | ✓ 1439ms | http |
| 20.78.213.56:80 | ✓ 1491ms | ✓ 1815ms | ✓ 1465ms | ✓ 1253ms | ✓ 959ms | http |
| 79.132.136.58:3128 | ✓ 926ms | 否 | ✓ 1325ms | ✓ 1891ms | 否 | http |
| 94.131.118.129:1081 | ✓ 375ms | 否 | ✓ 912ms | ✓ 1505ms | ✓ 1406ms | http |
| 116.80.47.79:3128 | ✓ 1674ms | 否 | ✓ 1637ms | 否 | ✓ 1803ms | http |
| 116.80.95.150:3128 | 否 | ✓ 1551ms | ✓ 1279ms | ✓ 1088ms | ✓ 802ms | http |
| 116.80.96.120:3128 | 否 | ✓ 1704ms | 否 | ✓ 999ms | ✓ 824ms | http |
| 83.219.250.8:62920 | ✓ 928ms | 否 | ✓ 1976ms | ✓ 1822ms | ✓ 1247ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1623ms | ✓ 1478ms | 否 | ✓ 1202ms | http |
| 204.152.223.225:9050 | ✓ 716ms | ✓ 869ms | 否 | 否 | ✓ 735ms | http |
| 159.223.225.118:8888 | ✓ 899ms | ✓ 1764ms | ✓ 635ms | ✓ 1574ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1579ms | 否 | ✓ 1405ms | ✓ 1509ms | ✓ 1474ms | http |
| 121.43.196.210:8222 | ✓ 1125ms | ✓ 1311ms | ✓ 1072ms | ✓ 1287ms | ✓ 1054ms | http |
| 121.43.196.213:8222 | ✓ 1119ms | ✓ 1288ms | ✓ 1120ms | ✓ 1314ms | ✓ 1070ms | http |
| 114.55.226.123:10086 | ✓ 1312ms | ✓ 1718ms | ✓ 1218ms | ✓ 1513ms | ✓ 1262ms | http |
| 51.49.153.183:51366 | ✓ 805ms | 否 | ✓ 1380ms | ✓ 1849ms | ✓ 1358ms | http |
| 195.26.224.49:3128 | ✓ 824ms | 否 | ✓ 668ms | ✓ 1849ms | 否 | http |
| 18.170.25.193:3129 | ✓ 1676ms | 否 | ✓ 635ms | ✓ 1922ms | ✓ 1609ms | http |
| 13.60.163.108:3128 | ✓ 1675ms | 否 | ✓ 1752ms | ✓ 1903ms | 否 | http |
| 5.135.137.24:80 | ✓ 663ms | ✓ 1945ms | ✓ 1164ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 989ms | 否 | ✓ 1319ms | 否 | ✓ 1968ms | http |
| 162.240.154.26:3128 | ✓ 609ms | 否 | ✓ 1108ms | ✓ 1469ms | ✓ 1827ms | http |
| 34.231.145.203:7000 | 否 | ✓ 863ms | ✓ 1873ms | 否 | ✓ 936ms | http |
| 210.223.44.230:3128 | ✓ 922ms | ✓ 1293ms | ✓ 1595ms | ✓ 1853ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1739ms | 否 | ✓ 1559ms | 否 | ✓ 1622ms | http |
| 47.74.226.8:5001 | ✓ 1377ms | 否 | ✓ 1226ms | 否 | ✓ 1500ms | http |
| 77.110.113.24:40000 | 否 | 否 | ✓ 904ms | ✓ 1828ms | ✓ 1546ms | http |
| 35.180.75.159:51978 | ✓ 636ms | 否 | ✓ 1348ms | 否 | ✓ 1258ms | http |
| 52.56.167.111:11033 | ✓ 652ms | 否 | 否 | ✓ 1560ms | ✓ 1174ms | http |
| 3.79.194.222:345 | ✓ 987ms | 否 | ✓ 1002ms | 否 | ✓ 1326ms | http |
| 52.47.115.41:24427 | ✓ 639ms | 否 | ✓ 1142ms | 否 | ✓ 1639ms | http |
| 158.160.215.167:8126 | ✓ 1318ms | 否 | ✓ 1478ms | 否 | ✓ 1901ms | http |
| 217.76.245.80:999 | ✓ 844ms | ✓ 1430ms | 否 | ✓ 1316ms | ✓ 1054ms | http |
| 103.82.93.219:3128 | ✓ 1873ms | 否 | ✓ 1363ms | ✓ 1601ms | ✓ 1367ms | http |
| 185.196.61.181:8081 | 否 | ✓ 1190ms | ✓ 1232ms | 否 | ✓ 1726ms | http |
| 62.113.119.14:8080 | ✓ 761ms | ✓ 1690ms | ✓ 698ms | ✓ 1546ms | 否 | http |
| 103.137.193.55:8080 | ✓ 1698ms | 否 | ✓ 1282ms | ✓ 1726ms | ✓ 1280ms | http |
| 139.5.189.229:8888 | ✓ 1609ms | 否 | 否 | ✓ 1689ms | ✓ 1577ms | http |
| 103.159.96.195:8082 | ✓ 1991ms | 否 | ✓ 1758ms | ✓ 1745ms | ✓ 1739ms | http |
| 61.52.131.172:8443 | ✓ 1089ms | ✓ 1429ms | ✓ 1127ms | ✓ 1462ms | ✓ 1153ms | http |
| 103.85.113.66:9999 | ✓ 959ms | ✓ 1403ms | ✓ 1142ms | ✓ 1894ms | ✓ 1622ms | http |
| 52.56.167.111:8906 | ✓ 987ms | 否 | ✓ 944ms | 否 | ✓ 1617ms | http |
| 218.108.131.186:17890 | ✓ 1084ms | ✓ 1225ms | ✓ 980ms | ✓ 1352ms | ✓ 1083ms | http |
| 36.103.198.235:7890 | 否 | ✓ 1836ms | ✓ 1553ms | ✓ 1610ms | ✓ 1343ms | http |
| 36.141.21.200:7890 | ✓ 1145ms | 否 | ✓ 1979ms | ✓ 1464ms | 否 | http |
| 101.43.127.100:8877 | 否 | ✓ 1944ms | ✓ 1940ms | ✓ 1373ms | ✓ 1201ms | http |

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
