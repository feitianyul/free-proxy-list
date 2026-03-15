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

最后更新：2026-03-15 20:31:39 UTC（2026-03-16 04:31:39 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 211ms | ✓ 916ms | ✓ 1019ms | ✓ 968ms | ✓ 789ms | http |
| 38.145.203.135:8443 | ✓ 797ms | ✓ 1059ms | ✓ 1230ms | ✓ 1041ms | ✓ 739ms | http |
| 38.145.218.82:8443 | ✓ 777ms | ✓ 1292ms | ✓ 1015ms | ✓ 1127ms | ✓ 928ms | http |
| 149.50.116.240:1080 | ✓ 896ms | ✓ 1723ms | ✓ 1327ms | ✓ 1813ms | ✓ 1649ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1096ms | ✓ 1383ms | ✓ 1023ms | http |
| 92.119.127.212:6005 | ✓ 1255ms | 否 | ✓ 1643ms | ✓ 1917ms | 否 | http |
| 137.220.150.104:6005 | ✓ 1799ms | 否 | ✓ 1720ms | ✓ 1799ms | ✓ 1514ms | http |
| 113.160.132.26:8080 | ✓ 1574ms | 否 | 否 | ✓ 1759ms | ✓ 1488ms | http |
| 45.167.124.52:8080 | ✓ 1130ms | ✓ 1858ms | ✓ 947ms | ✓ 1713ms | ✓ 1342ms | http |
| 120.92.212.16:7890 | ✓ 1100ms | ✓ 1463ms | ✓ 1113ms | ✓ 1464ms | 否 | http |
| 153.126.195.219:5080 | ✓ 1514ms | ✓ 1432ms | ✓ 1290ms | ✓ 1339ms | ✓ 1167ms | http |
| 104.129.202.127:10810 | ✓ 393ms | 否 | 否 | ✓ 1909ms | ✓ 1167ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1616ms | ✓ 1251ms | ✓ 1629ms | 否 | http |
| 45.168.238.193:8443 | ✓ 739ms | ✓ 1186ms | ✓ 557ms | ✓ 1205ms | ✓ 915ms | http |
| 104.129.202.127:12354 | 否 | 否 | ✓ 305ms | ✓ 1244ms | ✓ 1416ms | http |
| 83.219.250.8:62920 | ✓ 1458ms | ✓ 1688ms | ✓ 1559ms | 否 | ✓ 1869ms | http |
| 101.43.127.100:8877 | ✓ 1070ms | ✓ 1203ms | ✓ 1064ms | ✓ 1362ms | ✓ 1058ms | http |
| 81.70.169.194:80 | ✓ 1188ms | ✓ 1530ms | ✓ 1102ms | ✓ 1592ms | ✓ 1218ms | http |
| 101.43.255.96:80 | ✓ 1113ms | ✓ 1568ms | ✓ 1237ms | ✓ 1558ms | ✓ 1163ms | http |
| 165.225.72.38:10843 | ✓ 1422ms | ✓ 1955ms | ✓ 1242ms | 否 | ✓ 1841ms | http |
| 165.225.72.38:10755 | ✓ 1421ms | ✓ 1958ms | ✓ 1239ms | 否 | 否 | http |
| 192.71.213.85:9091 | ✓ 792ms | 否 | ✓ 1564ms | ✓ 1646ms | 否 | http |
| 178.236.245.17:3128 | ✓ 948ms | 否 | ✓ 1514ms | 否 | ✓ 1298ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 1883ms | ✓ 1280ms | ✓ 1307ms | http |
| 5.102.109.41:999 | 否 | ✓ 1642ms | ✓ 1193ms | 否 | ✓ 1630ms | http |
| 165.225.72.38:10801 | ✓ 531ms | ✓ 1780ms | ✓ 966ms | ✓ 1466ms | ✓ 1108ms | http |
| 14.225.212.37:7890 | ✓ 970ms | ✓ 1781ms | ✓ 1125ms | 否 | 否 | http |
| 47.105.98.23:3128 | ✓ 1951ms | 否 | 否 | ✓ 1469ms | ✓ 1394ms | http |
| 160.238.65.7:3128 | ✓ 1065ms | 否 | 否 | ✓ 1227ms | ✓ 1553ms | http |
| 137.220.151.110:6005 | ✓ 1075ms | 否 | ✓ 1395ms | ✓ 1586ms | ✓ 1239ms | http |
| 72.11.150.178:6005 | ✓ 446ms | ✓ 1209ms | ✓ 786ms | ✓ 1365ms | ✓ 954ms | http |
| 85.198.96.242:3128 | ✓ 564ms | 否 | 否 | ✓ 1603ms | ✓ 1713ms | http |
| 178.236.245.59:3128 | ✓ 618ms | 否 | ✓ 635ms | ✓ 1600ms | ✓ 1494ms | http |
| 159.223.42.219:3128 | ✓ 1464ms | 否 | ✓ 1384ms | ✓ 1450ms | ✓ 1090ms | http |
| 85.214.190.195:3128 | ✓ 770ms | 否 | ✓ 1143ms | ✓ 1574ms | ✓ 1475ms | http |
| 62.113.119.14:8080 | ✓ 1155ms | ✓ 1480ms | ✓ 519ms | ✓ 1433ms | ✓ 1861ms | http |
| 88.80.150.82:8080 | ✓ 1172ms | 否 | ✓ 1289ms | ✓ 1938ms | ✓ 1788ms | https |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1216ms | ✓ 1452ms | ✓ 1079ms | http |
| 125.76.214.178:8091 | ✓ 1296ms | ✓ 1629ms | ✓ 1152ms | ✓ 1759ms | ✓ 1297ms | http |
| 103.235.67.190:80 | ✓ 1081ms | 否 | ✓ 1478ms | ✓ 1540ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1255ms | ✓ 1476ms | ✓ 1264ms | ✓ 1527ms | ✓ 1135ms | http |
| 165.225.72.38:33333 | ✓ 1533ms | 否 | ✓ 1345ms | ✓ 1626ms | ✓ 1327ms | http |
| 165.225.72.38:10561 | ✓ 1533ms | 否 | ✓ 1342ms | ✓ 1589ms | ✓ 1875ms | http |
| 212.192.13.76:6005 | ✓ 1612ms | 否 | ✓ 1711ms | ✓ 1688ms | ✓ 1811ms | http |
| 45.136.130.223:8443 | 否 | ✓ 885ms | ✓ 934ms | ✓ 975ms | ✓ 1013ms | http |
| 92.119.127.213:6005 | ✓ 1874ms | ✓ 1758ms | ✓ 1493ms | 否 | 否 | http |
| 160.238.65.4:3128 | ✓ 1438ms | ✓ 1899ms | ✓ 380ms | ✓ 1196ms | ✓ 1253ms | http |
| 160.238.65.2:3128 | ✓ 363ms | ✓ 1584ms | ✓ 431ms | ✓ 1625ms | ✓ 1064ms | http |
| 156.246.88.104:11111 | ✓ 1413ms | ✓ 1371ms | ✓ 1665ms | ✓ 890ms | ✓ 714ms | http |
| 47.79.40.38:55000 | ✓ 1417ms | ✓ 1498ms | ✓ 1350ms | ✓ 927ms | ✓ 798ms | http |
| 45.136.198.40:3128 | ✓ 1460ms | 否 | ✓ 1951ms | ✓ 1932ms | ✓ 1895ms | http |
| 137.220.128.149:8866 | ✓ 1419ms | 否 | ✓ 1870ms | ✓ 1900ms | 否 | http |
| 35.225.22.61:80 | ✓ 575ms | 否 | 否 | ✓ 1155ms | ✓ 977ms | http |
| 2.56.122.146:10808 | ✓ 391ms | 否 | ✓ 808ms | ✓ 1678ms | ✓ 1286ms | http |
| 172.212.68.37:3128 | ✓ 638ms | 否 | ✓ 1043ms | ✓ 1839ms | ✓ 1188ms | http |
| 185.221.215.134:7788 | ✓ 1481ms | 否 | ✓ 1194ms | 否 | ✓ 1783ms | http |
| 103.39.51.190:8080 | ✓ 1948ms | 否 | 否 | ✓ 1907ms | ✓ 1588ms | http |
| 165.225.72.38:11814 | 否 | 否 | ✓ 595ms | ✓ 1355ms | ✓ 1011ms | http |
| 165.225.72.38:10869 | 否 | 否 | ✓ 410ms | ✓ 1329ms | ✓ 1010ms | http |
| 192.71.213.85:5678 | ✓ 1221ms | 否 | ✓ 1613ms | ✓ 1801ms | 否 | http |
| 165.225.72.38:10419 | ✓ 444ms | ✓ 1501ms | ✓ 411ms | ✓ 1322ms | ✓ 1022ms | http |
| 165.225.72.38:11526 | ✓ 1572ms | ✓ 1454ms | ✓ 411ms | 否 | ✓ 999ms | http |
| 165.225.72.38:11670 | 否 | 否 | ✓ 389ms | ✓ 1276ms | ✓ 1011ms | http |
| 165.225.72.38:11093 | 否 | ✓ 1526ms | ✓ 453ms | ✓ 1296ms | ✓ 1007ms | http |
| 165.225.72.38:10514 | 否 | ✓ 1567ms | ✓ 1446ms | ✓ 1301ms | ✓ 1011ms | http |
| 165.225.72.38:10456 | ✓ 396ms | 否 | ✓ 368ms | ✓ 1320ms | ✓ 1114ms | http |
| 165.225.72.38:11023 | ✓ 474ms | 否 | ✓ 1480ms | ✓ 1277ms | ✓ 1362ms | http |
| 160.238.65.8:3128 | ✓ 821ms | ✓ 1333ms | ✓ 1119ms | ✓ 1387ms | ✓ 1208ms | http |
| 160.238.65.9:3128 | ✓ 1440ms | ✓ 1187ms | ✓ 649ms | ✓ 1386ms | ✓ 1203ms | http |
| 160.238.65.5:3128 | ✓ 819ms | 否 | ✓ 455ms | ✓ 1385ms | ✓ 1208ms | http |
| 160.238.65.6:3128 | ✓ 1442ms | ✓ 1203ms | ✓ 629ms | ✓ 1391ms | ✓ 1207ms | http |
| 160.238.65.3:3128 | ✓ 1440ms | 否 | ✓ 355ms | ✓ 1411ms | ✓ 933ms | http |
| 86.53.183.16:1080 | ✓ 850ms | ✓ 1840ms | ✓ 1086ms | ✓ 1551ms | ✓ 1204ms | http |
| 194.5.212.40:8080 | ✓ 499ms | ✓ 1723ms | ✓ 474ms | 否 | ✓ 1584ms | http |
| 120.92.212.16:8890 | ✓ 1326ms | ✓ 1724ms | 否 | 否 | ✓ 1173ms | http |
| 211.217.231.234:8080 | ✓ 1556ms | ✓ 1672ms | ✓ 1064ms | ✓ 1159ms | ✓ 1005ms | http |
| 103.84.95.54:7890 | ✓ 1394ms | 否 | ✓ 821ms | 否 | ✓ 1308ms | http |
| 85.208.108.43:2094 | ✓ 190ms | 否 | ✓ 111ms | 否 | ✓ 1766ms | http |
| 8.209.239.31:30000 | ✓ 1767ms | ✓ 940ms | ✓ 830ms | ✓ 932ms | ✓ 772ms | http |
| 121.230.8.80:1080 | ✓ 1565ms | ✓ 1778ms | ✓ 1551ms | 否 | ✓ 1658ms | http |
| 121.230.9.4:1080 | ✓ 1244ms | ✓ 1853ms | ✓ 1323ms | ✓ 1979ms | ✓ 1339ms | http |
| 121.230.8.133:1080 | ✓ 1335ms | ✓ 1663ms | ✓ 1422ms | ✓ 1608ms | ✓ 1852ms | http |
| 121.40.231.103:7890 | 否 | ✓ 1580ms | 否 | ✓ 1669ms | ✓ 1893ms | http |
| 95.3.9.78:8080 | ✓ 1960ms | 否 | ✓ 1821ms | 否 | ✓ 1696ms | http |
| 185.241.5.57:3128 | ✓ 926ms | 否 | ✓ 1895ms | 否 | ✓ 1919ms | http |
| 62.60.177.204:34094 | ✓ 348ms | ✓ 1008ms | ✓ 1022ms | ✓ 895ms | ✓ 887ms | http |

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
