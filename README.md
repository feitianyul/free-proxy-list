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

最后更新：2026-04-04 09:45:42 UTC（2026-04-04 17:45:42 UTC+8）

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
| 147.161.239.240:8800 | ✓ 1071ms | 否 | ✓ 1398ms | ✓ 1470ms | ✓ 1302ms | http |
| 147.161.210.140:8800 | ✓ 1123ms | ✓ 1799ms | ✓ 1192ms | ✓ 1314ms | ✓ 1270ms | http |
| 218.108.131.186:17890 | ✓ 1012ms | ✓ 1233ms | ✓ 1080ms | ✓ 1540ms | ✓ 1074ms | http |
| 209.38.154.7:1080 | ✓ 1957ms | 否 | ✓ 1420ms | 否 | ✓ 726ms | http |
| 159.223.71.162:8080 | ✓ 1371ms | 否 | ✓ 1494ms | ✓ 1336ms | ✓ 1039ms | http |
| 113.160.132.26:8080 | ✓ 1544ms | ✓ 1547ms | ✓ 1546ms | ✓ 1428ms | ✓ 1166ms | http |
| 167.103.115.102:8800 | ✓ 1622ms | 否 | ✓ 1230ms | ✓ 1478ms | ✓ 1434ms | http |
| 159.223.71.162:443 | ✓ 1368ms | 否 | ✓ 1621ms | 否 | ✓ 1204ms | http |
| 167.103.34.108:8800 | ✓ 1613ms | 否 | ✓ 1754ms | 否 | ✓ 1549ms | http |
| 45.167.124.52:8080 | ✓ 1613ms | ✓ 1924ms | 否 | ✓ 1937ms | ✓ 1555ms | http |
| 210.223.44.230:3128 | ✓ 1388ms | ✓ 1231ms | ✓ 1334ms | ✓ 1136ms | ✓ 979ms | http |
| 1.231.81.166:3128 | ✓ 1324ms | ✓ 1378ms | ✓ 1705ms | ✓ 1313ms | ✓ 1091ms | http |
| 95.213.217.168:52004 | ✓ 914ms | 否 | ✓ 1907ms | 否 | ✓ 1522ms | http |
| 5.104.87.17:8051 | ✓ 1536ms | 否 | ✓ 1810ms | ✓ 1673ms | ✓ 1368ms | http |
| 167.103.31.122:8800 | ✓ 1825ms | 否 | ✓ 1647ms | 否 | ✓ 1686ms | http |
| 111.227.254.9:22222 | ✓ 1192ms | ✓ 1565ms | 否 | 否 | ✓ 1256ms | http |
| 59.46.216.131:30001 | ✓ 1117ms | 否 | ✓ 1287ms | ✓ 1580ms | ✓ 1233ms | http |
| 91.233.223.147:3128 | ✓ 1322ms | 否 | ✓ 1393ms | 否 | ✓ 1561ms | http |
| 167.103.144.127:8800 | ✓ 1881ms | 否 | ✓ 1258ms | ✓ 1445ms | ✓ 1314ms | http |
| 111.227.254.12:22222 | ✓ 1159ms | ✓ 1643ms | 否 | ✓ 1853ms | ✓ 1242ms | http |
| 101.43.127.100:8877 | ✓ 1136ms | ✓ 1304ms | ✓ 1080ms | ✓ 1302ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1129ms | ✓ 1438ms | ✓ 1934ms | ✓ 1449ms | ✓ 1141ms | http |
| 101.132.61.121:8888 | 否 | ✓ 1429ms | ✓ 1570ms | ✓ 1749ms | 否 | http |
| 45.167.125.21:999 | ✓ 535ms | 否 | ✓ 1521ms | ✓ 1654ms | ✓ 1568ms | http |
| 115.231.181.40:8128 | ✓ 1159ms | 否 | ✓ 1345ms | 否 | ✓ 1138ms | http |
| 45.140.147.82:1081 | ✓ 1024ms | ✓ 1806ms | ✓ 887ms | ✓ 1562ms | ✓ 1051ms | http |
| 118.31.1.154:80 | ✓ 1025ms | ✓ 1285ms | ✓ 1000ms | ✓ 1336ms | ✓ 1092ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 582ms | ✓ 1435ms | ✓ 1059ms | http |
| 43.165.195.107:3128 | ✓ 1756ms | ✓ 1545ms | ✓ 1617ms | ✓ 1396ms | ✓ 1113ms | http |
| 64.227.76.27:1080 | ✓ 1380ms | 否 | 否 | ✓ 1581ms | ✓ 1886ms | http |
| 170.9.253.20:8888 | ✓ 834ms | 否 | ✓ 1684ms | ✓ 1729ms | 否 | http |
| 121.230.9.161:1080 | ✓ 1854ms | ✓ 1676ms | ✓ 1979ms | 否 | ✓ 1453ms | http |
| 120.92.212.16:8890 | ✓ 1096ms | ✓ 1432ms | ✓ 1139ms | ✓ 1421ms | ✓ 1187ms | http |
| 8.219.97.248:80 | ✓ 1655ms | 否 | ✓ 1486ms | ✓ 1678ms | 否 | http |
| 92.119.127.211:6005 | ✓ 924ms | ✓ 1433ms | ✓ 1788ms | ✓ 1348ms | ✓ 1329ms | http |
| 45.12.151.226:2829 | 否 | ✓ 1567ms | ✓ 922ms | ✓ 1596ms | ✓ 1148ms | http |
| 212.58.132.5:8888 | ✓ 1248ms | 否 | ✓ 1266ms | ✓ 1741ms | ✓ 1209ms | http |
| 183.61.28.254:8900 | ✓ 1498ms | ✓ 1550ms | 否 | 否 | ✓ 1563ms | http |
| 43.167.237.94:3128 | ✓ 1639ms | 否 | ✓ 1242ms | 否 | ✓ 1686ms | http |
| 93.77.179.177:8888 | 否 | 否 | ✓ 1066ms | ✓ 1800ms | ✓ 1419ms | http |
| 103.39.51.207:8080 | ✓ 1765ms | 否 | ✓ 1580ms | ✓ 1667ms | ✓ 1588ms | http |
| 104.168.92.24:3128 | ✓ 351ms | ✓ 1343ms | ✓ 1011ms | ✓ 968ms | ✓ 888ms | http |
| 45.125.67.37:443 | ✓ 1072ms | 否 | ✓ 1063ms | ✓ 1272ms | ✓ 1290ms | http |
| 201.139.115.38:8080 | ✓ 1268ms | 否 | ✓ 1384ms | 否 | ✓ 1671ms | http |
| 121.230.8.34:1080 | ✓ 1437ms | 否 | 否 | ✓ 1719ms | ✓ 1250ms | http |
| 114.237.77.219:1080 | ✓ 1506ms | ✓ 1526ms | ✓ 1133ms | ✓ 1810ms | 否 | http |
| 192.71.213.85:9812 | ✓ 1290ms | 否 | ✓ 1775ms | ✓ 1578ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1674ms | 否 | ✓ 1398ms | ✓ 1783ms | ✓ 1597ms | http |
| 181.78.44.63:999 | ✓ 1174ms | ✓ 1762ms | ✓ 676ms | ✓ 1302ms | ✓ 1021ms | http |
| 116.80.63.46:7777 | ✓ 1818ms | 否 | ✓ 1668ms | 否 | ✓ 1817ms | http |
| 38.145.220.198:8448 | ✓ 1691ms | 否 | ✓ 1059ms | 否 | ✓ 961ms | http |
| 217.217.249.160:8080 | ✓ 1076ms | 否 | ✓ 1543ms | 否 | ✓ 1739ms | http |
| 45.136.198.40:3128 | ✓ 596ms | 否 | ✓ 1868ms | ✓ 1979ms | ✓ 1641ms | http |
| 38.207.165.200:6005 | ✓ 1245ms | 否 | ✓ 1207ms | ✓ 1113ms | ✓ 1039ms | http |
| 83.166.244.230:1080 | ✓ 1105ms | ✓ 1924ms | ✓ 1906ms | 否 | ✓ 1774ms | http |
| 111.227.254.11:22222 | ✓ 1199ms | 否 | ✓ 1300ms | ✓ 1585ms | ✓ 1267ms | http |
| 35.225.22.61:80 | ✓ 758ms | ✓ 1283ms | ✓ 1362ms | ✓ 1051ms | ✓ 875ms | http |
| 103.113.70.189:1081 | ✓ 1378ms | 否 | ✓ 104ms | ✓ 1218ms | ✓ 893ms | http |
| 38.145.208.243:8447 | 否 | ✓ 838ms | ✓ 1007ms | 否 | ✓ 722ms | http |
| 5.104.87.17:8050 | ✓ 1701ms | 否 | ✓ 1907ms | 否 | ✓ 1239ms | http |
| 38.34.179.105:8449 | ✓ 366ms | ✓ 1958ms | ✓ 1487ms | ✓ 939ms | 否 | http |
| 150.249.255.91:3128 | ✓ 1395ms | ✓ 1092ms | ✓ 646ms | ✓ 1028ms | 否 | http |
| 150.241.71.15:1080 | 否 | 否 | ✓ 738ms | ✓ 1555ms | ✓ 974ms | http |
| 113.176.92.71:3128 | ✓ 1919ms | ✓ 1535ms | 否 | ✓ 1421ms | ✓ 1173ms | http |
| 45.129.141.143:3128 | ✓ 1060ms | 否 | 否 | ✓ 1905ms | ✓ 1682ms | http |
| 208.87.243.199:7878 | ✓ 1281ms | 否 | ✓ 1275ms | ✓ 1715ms | ✓ 1124ms | http |
| 120.79.99.232:8099 | ✓ 1309ms | ✓ 1731ms | ✓ 1436ms | ✓ 1675ms | ✓ 1396ms | http |
| 47.110.42.192:9003 | ✓ 1883ms | ✓ 1515ms | ✓ 1595ms | 否 | ✓ 1577ms | http |
| 34.96.238.40:8080 | ✓ 1116ms | ✓ 1424ms | ✓ 1312ms | ✓ 1199ms | 否 | http |
| 217.77.102.18:3128 | 否 | 否 | ✓ 1215ms | ✓ 1734ms | ✓ 1585ms | http |
| 80.250.165.242:3128 | ✓ 970ms | ✓ 1902ms | ✓ 1325ms | ✓ 1532ms | ✓ 1218ms | http |
| 133.242.138.34:8100 | ✓ 823ms | ✓ 1336ms | ✓ 1804ms | ✓ 1127ms | ✓ 895ms | http |
| 45.140.147.155:1081 | ✓ 562ms | ✓ 1712ms | ✓ 846ms | ✓ 1556ms | ✓ 1097ms | http |
| 45.140.147.155:1082 | ✓ 561ms | 否 | ✓ 559ms | ✓ 1556ms | ✓ 1119ms | http |
| 116.171.106.15:3443 | ✓ 1984ms | ✓ 1748ms | ✓ 1940ms | 否 | 否 | http |
| 104.243.46.122:3128 | ✓ 579ms | ✓ 1465ms | ✓ 1026ms | ✓ 1266ms | ✓ 848ms | http |
| 104.248.243.244:3128 | ✓ 421ms | ✓ 1844ms | ✓ 1580ms | ✓ 1185ms | ✓ 1321ms | http |
| 61.76.95.217:40088 | ✓ 1110ms | ✓ 1503ms | ✓ 1500ms | ✓ 1467ms | ✓ 1168ms | http |
| 185.191.236.162:3128 | ✓ 1167ms | 否 | ✓ 1932ms | 否 | ✓ 1225ms | http |
| 185.241.5.57:3128 | ✓ 1125ms | 否 | ✓ 1396ms | 否 | ✓ 1509ms | http |
| 114.237.77.244:1080 | 否 | 否 | ✓ 1156ms | ✓ 1410ms | ✓ 1485ms | http |
| 115.78.135.4:3334 | ✓ 1811ms | 否 | ✓ 1847ms | ✓ 1548ms | ✓ 1562ms | http |
| 38.34.179.78:8446 | ✓ 1544ms | 否 | ✓ 360ms | ✓ 1374ms | 否 | http |
| 38.34.179.50:8449 | ✓ 506ms | ✓ 1132ms | 否 | ✓ 1627ms | 否 | http |
| 38.145.220.11:8447 | ✓ 901ms | ✓ 1286ms | ✓ 1765ms | ✓ 1605ms | ✓ 1316ms | http |
| 41.128.72.194:1976 | ✓ 1408ms | ✓ 1906ms | 否 | 否 | ✓ 1994ms | http |

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
