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

最后更新：2026-04-16 11:11:17 UTC（2026-04-16 19:11:17 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 905ms | ✓ 1211ms | ✓ 926ms | ✓ 1117ms | ✓ 1165ms | http |
| 147.161.239.240:8800 | ✓ 1100ms | ✓ 1589ms | ✓ 1108ms | ✓ 1409ms | ✓ 1242ms | http |
| 147.161.210.140:8800 | 否 | ✓ 1580ms | ✓ 818ms | ✓ 1218ms | ✓ 1036ms | http |
| 157.230.178.216:8088 | ✓ 814ms | 否 | ✓ 967ms | ✓ 1229ms | ✓ 949ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1296ms | ✓ 1554ms | ✓ 1706ms | ✓ 1306ms | http |
| 167.103.115.102:8800 | ✓ 1708ms | 否 | ✓ 1285ms | ✓ 1333ms | ✓ 1442ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1507ms | ✓ 1322ms | ✓ 1781ms | ✓ 1315ms | http |
| 5.104.87.17:8051 | ✓ 1674ms | 否 | ✓ 1436ms | ✓ 1594ms | 否 | http |
| 78.11.96.22:8888 | ✓ 1134ms | ✓ 1896ms | ✓ 1075ms | ✓ 1853ms | ✓ 1684ms | http |
| 35.225.22.61:80 | ✓ 381ms | ✓ 1083ms | ✓ 776ms | 否 | ✓ 863ms | http |
| 62.113.119.14:8080 | ✓ 706ms | 否 | ✓ 593ms | 否 | ✓ 1902ms | http |
| 130.61.30.221:8080 | ✓ 705ms | 否 | ✓ 1455ms | ✓ 1893ms | ✓ 1634ms | http |
| 167.103.34.108:8800 | ✓ 1360ms | 否 | ✓ 1242ms | ✓ 1460ms | 否 | http |
| 27.71.24.102:3128 | ✓ 1413ms | 否 | ✓ 1323ms | ✓ 1225ms | ✓ 1153ms | http |
| 20.127.128.70:8080 | ✓ 1160ms | 否 | ✓ 1241ms | 否 | ✓ 1712ms | http |
| 34.101.184.164:3128 | ✓ 1730ms | 否 | ✓ 1410ms | ✓ 1784ms | ✓ 1492ms | http |
| 167.103.144.127:8800 | 否 | 否 | ✓ 1437ms | ✓ 1873ms | ✓ 1915ms | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 958ms | ✓ 1156ms | ✓ 998ms | http |
| 177.234.217.88:999 | ✓ 1192ms | 否 | ✓ 1943ms | 否 | ✓ 1676ms | http |
| 140.238.242.189:8100 | ✓ 1589ms | 否 | ✓ 1122ms | ✓ 1868ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1109ms | ✓ 1611ms | ✓ 1188ms | ✓ 1404ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1368ms | 否 | ✓ 1950ms | ✓ 1957ms | ✓ 1725ms | http |
| 167.103.31.122:8800 | ✓ 1597ms | 否 | ✓ 1351ms | ✓ 1686ms | ✓ 1562ms | http |
| 190.12.150.244:999 | ✓ 891ms | 否 | ✓ 859ms | ✓ 1615ms | ✓ 1401ms | http |
| 38.43.93.161:8000 | ✓ 782ms | ✓ 1102ms | ✓ 1293ms | 否 | 否 | http |
| 116.80.63.194:3172 | 否 | 否 | ✓ 1617ms | ✓ 1968ms | ✓ 1770ms | http |
| 104.248.211.46:7890 | ✓ 831ms | ✓ 1247ms | ✓ 386ms | 否 | 否 | http |
| 103.113.70.189:1082 | ✓ 803ms | 否 | ✓ 108ms | ✓ 1114ms | ✓ 1469ms | http |
| 140.245.66.105:8081 | ✓ 1577ms | ✓ 1288ms | ✓ 881ms | ✓ 1165ms | ✓ 926ms | http |
| 103.85.113.66:9999 | ✓ 1222ms | ✓ 1671ms | ✓ 1380ms | ✓ 1838ms | ✓ 1434ms | http |
| 103.113.70.189:1081 | ✓ 1540ms | ✓ 1079ms | ✓ 515ms | ✓ 1173ms | ✓ 777ms | http |
| 181.78.44.63:999 | ✓ 916ms | ✓ 1523ms | ✓ 1173ms | ✓ 1339ms | ✓ 1351ms | http |
| 45.167.125.21:999 | ✓ 1271ms | ✓ 1926ms | ✓ 1433ms | ✓ 1868ms | ✓ 1610ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1955ms | ✓ 1301ms | 否 | ✓ 1304ms | http |
| 185.191.236.162:3128 | ✓ 1830ms | 否 | ✓ 1271ms | ✓ 1965ms | ✓ 1379ms | http |
| 144.31.27.49:1080 | 否 | ✓ 1916ms | 否 | ✓ 1675ms | ✓ 1622ms | http |
| 158.160.215.167:8124 | ✓ 1550ms | 否 | ✓ 1923ms | 否 | ✓ 1944ms | http |
| 64.181.240.152:3128 | ✓ 1327ms | ✓ 1299ms | ✓ 827ms | ✓ 966ms | ✓ 726ms | http |
| 34.96.238.40:8080 | ✓ 1302ms | ✓ 1439ms | 否 | ✓ 1127ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1792ms | 否 | ✓ 1522ms | ✓ 1222ms | ✓ 825ms | http |
| 185.132.178.178:1080 | ✓ 1643ms | ✓ 1664ms | ✓ 898ms | ✓ 1808ms | ✓ 1917ms | http |
| 45.140.147.82:1081 | ✓ 1645ms | ✓ 1278ms | ✓ 1281ms | ✓ 1798ms | ✓ 1539ms | http |
| 168.144.75.9:3128 | ✓ 1154ms | 否 | ✓ 1952ms | ✓ 1855ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1695ms | 否 | ✓ 1891ms | ✓ 1784ms | ✓ 1232ms | http |
| 195.26.224.49:3128 | 否 | 否 | ✓ 632ms | ✓ 1708ms | ✓ 1574ms | http |
| 188.246.224.49:7890 | ✓ 1187ms | ✓ 1955ms | ✓ 1740ms | 否 | ✓ 1545ms | http |
| 43.132.188.134:443 | ✓ 1789ms | ✓ 1901ms | 否 | ✓ 1360ms | ✓ 783ms | http |
| 85.239.59.252:7890 | ✓ 1455ms | ✓ 1896ms | 否 | 否 | ✓ 1136ms | http |
| 103.138.70.165:3129 | ✓ 1958ms | 否 | ✓ 1950ms | ✓ 1606ms | ✓ 1523ms | http |
| 157.230.38.173:3128 | ✓ 933ms | 否 | ✓ 1306ms | ✓ 1281ms | ✓ 1012ms | http |
| 114.237.77.200:1080 | 否 | 否 | ✓ 1144ms | ✓ 1465ms | ✓ 1163ms | http |
| 168.119.38.41:5286 | ✓ 1977ms | 否 | ✓ 1796ms | 否 | ✓ 1597ms | http |
| 121.230.8.17:1080 | ✓ 1194ms | ✓ 1553ms | 否 | ✓ 1511ms | 否 | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1192ms | ✓ 1451ms | ✓ 1281ms | http |
| 91.107.124.215:3128 | ✓ 949ms | 否 | ✓ 666ms | ✓ 1962ms | 否 | http |
| 45.12.151.226:2829 | ✓ 935ms | ✓ 1624ms | ✓ 541ms | ✓ 1694ms | ✓ 1267ms | http |
| 52.59.51.29:8008 | ✓ 1954ms | 否 | ✓ 1789ms | 否 | ✓ 1730ms | http |
| 185.230.190.195:3128 | ✓ 1158ms | ✓ 1691ms | ✓ 814ms | 否 | ✓ 1599ms | http |
| 12.89.176.82:3128 | ✓ 374ms | 否 | ✓ 1876ms | ✓ 996ms | 否 | http |
| 101.32.243.189:80 | ✓ 1107ms | 否 | ✓ 1866ms | ✓ 1718ms | ✓ 1537ms | http |
| 42.101.8.101:8888 | 否 | 否 | ✓ 1233ms | ✓ 1619ms | ✓ 1201ms | http |
| 117.236.124.166:3128 | ✓ 1842ms | 否 | ✓ 1035ms | ✓ 1961ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1536ms | ✓ 1539ms | ✓ 1832ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1647ms | ✓ 1279ms | ✓ 1557ms | 否 | http |
| 147.45.166.46:3128 | 否 | 否 | ✓ 1602ms | ✓ 1753ms | ✓ 1352ms | http |
| 114.118.82.146:80 | ✓ 1353ms | ✓ 1394ms | ✓ 1300ms | ✓ 1485ms | ✓ 1156ms | http |
| 84.47.150.126:1080 | 否 | ✓ 1870ms | ✓ 1313ms | 否 | ✓ 1850ms | http |
| 202.141.161.53:10808 | ✓ 1238ms | ✓ 1595ms | ✓ 1243ms | ✓ 1441ms | ✓ 1225ms | http |
| 47.74.226.8:5001 | ✓ 1559ms | ✓ 1761ms | ✓ 1139ms | 否 | 否 | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1395ms | ✓ 1617ms | ✓ 1360ms | http |
| 60.249.94.208:3128 | ✓ 1076ms | ✓ 1188ms | 否 | ✓ 1024ms | ✓ 853ms | http |
| 94.131.118.129:1081 | ✓ 1139ms | 否 | ✓ 749ms | ✓ 1957ms | ✓ 1515ms | http |
| 158.160.85.248:3128 | 否 | 否 | ✓ 1041ms | ✓ 1998ms | ✓ 1920ms | http |
| 147.45.214.210:1080 | 否 | ✓ 1820ms | ✓ 1714ms | 否 | ✓ 1290ms | http |
| 152.32.132.190:7890 | ✓ 816ms | 否 | 否 | ✓ 1780ms | ✓ 798ms | http |
| 107.172.102.234:40621 | ✓ 1486ms | ✓ 965ms | ✓ 888ms | ✓ 1303ms | ✓ 834ms | http |
| 160.238.65.5:3128 | ✓ 1109ms | 否 | ✓ 545ms | 否 | ✓ 1537ms | http |
| 160.238.65.4:3128 | ✓ 1109ms | ✓ 1210ms | ✓ 1335ms | 否 | 否 | http |
| 160.238.65.8:3128 | ✓ 1109ms | ✓ 1362ms | ✓ 1183ms | 否 | 否 | http |
| 160.238.65.3:3128 | ✓ 1591ms | ✓ 1501ms | ✓ 586ms | 否 | 否 | http |
| 160.238.65.6:3128 | ✓ 1109ms | 否 | ✓ 545ms | 否 | ✓ 1534ms | http |
| 160.238.65.7:3128 | ✓ 1109ms | 否 | ✓ 544ms | 否 | ✓ 1539ms | http |
| 160.238.65.9:3128 | ✓ 1109ms | ✓ 1455ms | ✓ 1091ms | 否 | 否 | http |
| 147.45.167.84:3128 | ✓ 1367ms | 否 | ✓ 673ms | ✓ 1625ms | 否 | http |
| 181.78.49.177:999 | ✓ 1427ms | ✓ 1737ms | ✓ 1795ms | ✓ 1799ms | 否 | http |
| 8.217.90.107:22520 | ✓ 784ms | ✓ 1289ms | ✓ 785ms | 否 | 否 | http |
| 129.213.162.27:17777 | ✓ 509ms | ✓ 1527ms | 否 | 否 | ✓ 1321ms | http |
| 103.4.77.33:8082 | 否 | 否 | ✓ 1798ms | ✓ 1577ms | ✓ 1972ms | http |

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
