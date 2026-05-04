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

最后更新：2026-05-04 07:32:53 UTC（2026-05-04 15:32:53 UTC+8）

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
| 218.108.131.186:17890 | ✓ 855ms | ✓ 1065ms | ✓ 928ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 1875ms | 否 | ✓ 1581ms | ✓ 1252ms | ✓ 972ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1948ms | ✓ 1593ms | ✓ 1281ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1708ms | ✓ 1140ms | ✓ 1840ms | ✓ 1345ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1516ms | ✓ 1738ms | ✓ 1243ms | http |
| 45.167.124.71:999 | ✓ 1594ms | 否 | ✓ 1738ms | ✓ 1631ms | ✓ 1696ms | http |
| 38.95.77.85:6005 | ✓ 1908ms | ✓ 995ms | ✓ 1136ms | ✓ 1630ms | ✓ 849ms | http |
| 46.105.190.40:3128 | 否 | ✓ 1322ms | ✓ 645ms | 否 | ✓ 1532ms | http |
| 20.127.128.70:8080 | ✓ 612ms | 否 | ✓ 1689ms | ✓ 1464ms | ✓ 1858ms | http |
| 45.59.122.132:80 | 否 | 否 | ✓ 1209ms | ✓ 1713ms | ✓ 1575ms | http |
| 168.110.52.228:3128 | ✓ 1671ms | 否 | ✓ 1987ms | ✓ 1272ms | ✓ 1142ms | http |
| 1.231.81.166:3128 | ✓ 1689ms | 否 | ✓ 1369ms | ✓ 1523ms | ✓ 1266ms | http |
| 173.212.246.157:3128 | ✓ 1200ms | 否 | ✓ 1650ms | 否 | ✓ 1797ms | http |
| 103.157.200.126:3128 | ✓ 1760ms | 否 | ✓ 1331ms | 否 | ✓ 1784ms | http |
| 154.64.232.35:8080 | ✓ 1488ms | ✓ 1354ms | ✓ 1265ms | ✓ 1494ms | 否 | http |
| 81.26.190.143:1080 | ✓ 793ms | 否 | ✓ 815ms | 否 | ✓ 1493ms | http |
| 31.56.48.253:26133 | ✓ 1242ms | 否 | ✓ 384ms | 否 | ✓ 1379ms | http |
| 45.140.147.155:1081 | ✓ 1217ms | ✓ 1429ms | ✓ 1258ms | 否 | ✓ 1089ms | http |
| 37.187.109.70:10111 | ✓ 1235ms | ✓ 1916ms | ✓ 979ms | 否 | ✓ 1936ms | http |
| 120.92.108.86:7890 | ✓ 1845ms | 否 | ✓ 1776ms | ✓ 1803ms | ✓ 1559ms | http |
| 59.46.216.131:30001 | ✓ 1152ms | ✓ 1655ms | ✓ 1418ms | ✓ 1642ms | 否 | http |
| 190.9.48.193:999 | ✓ 846ms | 否 | ✓ 1557ms | ✓ 1390ms | 否 | http |
| 45.153.231.229:8080 | ✓ 1165ms | 否 | ✓ 1886ms | 否 | ✓ 1391ms | http |
| 217.52.247.70:1981 | ✓ 1834ms | 否 | ✓ 1800ms | ✓ 1499ms | ✓ 1315ms | http |
| 47.85.51.197:1080 | ✓ 469ms | ✓ 1827ms | ✓ 1758ms | ✓ 1760ms | 否 | http |
| 57.128.188.167:9174 | ✓ 1529ms | 否 | ✓ 1622ms | ✓ 1870ms | ✓ 1560ms | http |
| 45.140.147.155:1082 | ✓ 1076ms | ✓ 1599ms | ✓ 673ms | ✓ 1614ms | ✓ 1131ms | http |
| 46.105.190.38:3128 | ✓ 970ms | ✓ 1273ms | ✓ 498ms | ✓ 1903ms | ✓ 1754ms | http |
| 41.128.72.71:1981 | ✓ 1030ms | 否 | ✓ 1762ms | 否 | ✓ 1937ms | http |
| 103.109.173.171:80 | 否 | 否 | ✓ 1807ms | ✓ 1748ms | ✓ 1712ms | http |
| 47.77.216.82:1080 | ✓ 1495ms | 否 | ✓ 721ms | 否 | ✓ 1147ms | http |
| 120.92.212.16:8890 | ✓ 1388ms | ✓ 1565ms | 否 | ✓ 1917ms | ✓ 1300ms | http |
| 109.120.156.122:8090 | ✓ 990ms | 否 | ✓ 1516ms | 否 | ✓ 1599ms | http |
| 38.49.143.221:999 | ✓ 1223ms | ✓ 1382ms | ✓ 1208ms | ✓ 1264ms | ✓ 1258ms | http |
| 38.49.143.222:999 | ✓ 1222ms | ✓ 1420ms | ✓ 1170ms | ✓ 1591ms | ✓ 989ms | http |
| 38.49.143.218:999 | ✓ 1221ms | ✓ 1377ms | ✓ 1214ms | ✓ 1385ms | ✓ 1199ms | http |
| 38.49.143.220:999 | ✓ 1223ms | ✓ 1393ms | ✓ 1198ms | ✓ 1631ms | ✓ 996ms | http |
| 38.49.143.219:999 | ✓ 1222ms | ✓ 1412ms | ✓ 1179ms | ✓ 1423ms | ✓ 1420ms | http |
| 168.194.0.249:252 | ✓ 1236ms | ✓ 1519ms | ✓ 1115ms | ✓ 1290ms | ✓ 1981ms | http |
| 177.52.221.6:999 | ✓ 1212ms | 否 | ✓ 1177ms | ✓ 1524ms | ✓ 1311ms | http |
| 190.60.57.77:9992 | ✓ 1270ms | ✓ 1778ms | ✓ 801ms | 否 | 否 | http |
| 45.173.12.140:1994 | ✓ 1257ms | 否 | ✓ 1335ms | ✓ 1852ms | ✓ 1459ms | http |
| 186.96.160.220:999 | ✓ 1240ms | ✓ 1640ms | ✓ 1138ms | 否 | 否 | http |
| 45.239.48.99:999 | ✓ 1342ms | ✓ 1646ms | ✓ 1636ms | 否 | 否 | http |
| 45.239.48.101:999 | ✓ 1342ms | 否 | ✓ 1707ms | ✓ 1647ms | 否 | http |
| 45.239.48.100:999 | ✓ 1424ms | ✓ 1604ms | ✓ 1893ms | ✓ 1820ms | 否 | http |
| 45.239.48.98:999 | ✓ 1344ms | 否 | ✓ 1785ms | ✓ 1642ms | 否 | http |
| 181.119.97.24:999 | ✓ 1336ms | 否 | ✓ 1906ms | 否 | ✓ 1571ms | http |
| 45.239.48.102:999 | ✓ 1588ms | ✓ 1559ms | ✓ 1681ms | 否 | 否 | http |
| 171.234.50.203:5116 | ✓ 1900ms | 否 | ✓ 1978ms | ✓ 1878ms | ✓ 1839ms | http |
| 152.32.132.190:7890 | ✓ 1175ms | 否 | ✓ 1072ms | ✓ 1151ms | ✓ 898ms | http |
| 120.92.212.16:7890 | ✓ 1253ms | 否 | ✓ 1402ms | ✓ 1923ms | 否 | http |
| 168.222.254.136:8888 | ✓ 698ms | 否 | ✓ 1395ms | 否 | ✓ 1685ms | http |
| 101.6.65.112:10080 | ✓ 1539ms | ✓ 1358ms | ✓ 1178ms | 否 | 否 | http |
| 89.208.106.138:10808 | ✓ 1559ms | ✓ 1528ms | ✓ 1691ms | 否 | ✓ 1185ms | http |
| 86.104.72.220:1082 | ✓ 182ms | ✓ 863ms | ✓ 384ms | ✓ 1789ms | ✓ 959ms | http |
| 62.113.119.14:8080 | ✓ 817ms | 否 | ✓ 570ms | ✓ 1839ms | ✓ 1143ms | http |
| 116.80.96.162:3172 | ✓ 1451ms | 否 | ✓ 709ms | ✓ 1022ms | ✓ 895ms | http |
| 118.113.246.72:1080 | ✓ 1341ms | ✓ 1862ms | ✓ 1337ms | 否 | ✓ 1415ms | http |
| 57.128.188.167:9249 | ✓ 1658ms | 否 | ✓ 1476ms | ✓ 1700ms | ✓ 1516ms | http |
| 116.80.65.10:3172 | ✓ 1445ms | ✓ 1824ms | 否 | 否 | ✓ 1884ms | http |
| 103.227.187.3:6090 | ✓ 1625ms | 否 | ✓ 1943ms | ✓ 1729ms | 否 | http |
| 103.227.187.241:6090 | ✓ 1916ms | 否 | ✓ 1666ms | ✓ 1727ms | 否 | http |
| 38.188.247.12:999 | ✓ 1945ms | ✓ 1605ms | ✓ 1026ms | ✓ 1554ms | ✓ 1245ms | http |
| 151.245.137.49:40000 | 否 | ✓ 1858ms | 否 | ✓ 1711ms | ✓ 1791ms | http |
| 45.125.67.37:8443 | ✓ 1112ms | 否 | ✓ 1329ms | ✓ 1318ms | ✓ 1316ms | http |
| 154.90.48.209:9090 | ✓ 1859ms | 否 | 否 | ✓ 1413ms | ✓ 1222ms | http |
| 116.80.93.67:3172 | ✓ 1830ms | 否 | ✓ 1670ms | 否 | ✓ 1191ms | http |
| 116.80.49.134:3172 | ✓ 1441ms | ✓ 1839ms | ✓ 1754ms | ✓ 1245ms | 否 | http |
| 193.123.250.39:1080 | ✓ 1297ms | 否 | ✓ 1253ms | ✓ 1519ms | ✓ 1512ms | http |
| 147.45.178.211:14658 | ✓ 543ms | 否 | ✓ 1610ms | 否 | ✓ 1155ms | http |
| 85.208.87.111:3128 | ✓ 1098ms | 否 | ✓ 544ms | ✓ 1589ms | ✓ 1342ms | http |
| 116.80.96.90:3172 | ✓ 1588ms | 否 | ✓ 665ms | ✓ 1011ms | ✓ 875ms | http |
| 38.51.237.100:999 | ✓ 1118ms | 否 | ✓ 1209ms | ✓ 1497ms | ✓ 1155ms | http |
| 161.49.219.181:8082 | 否 | 否 | ✓ 1958ms | ✓ 1562ms | ✓ 1995ms | http |
| 150.107.140.238:3128 | ✓ 1079ms | 否 | ✓ 1733ms | ✓ 1446ms | ✓ 1257ms | http |
| 124.16.102.108:26001 | ✓ 1035ms | ✓ 1360ms | ✓ 1096ms | ✓ 1347ms | ✓ 1078ms | http |
| 3.101.133.120:80 | 否 | ✓ 1904ms | ✓ 1405ms | ✓ 1554ms | 否 | http |
| 152.70.91.193:40000 | 否 | 否 | ✓ 1941ms | ✓ 1742ms | ✓ 1359ms | http |
| 112.186.12.84:3128 | ✓ 1446ms | ✓ 1202ms | ✓ 986ms | ✓ 1184ms | ✓ 973ms | http |
| 112.186.12.235:3128 | ✓ 1439ms | ✓ 1197ms | ✓ 985ms | ✓ 1242ms | ✓ 935ms | http |
| 86.104.72.220:1081 | ✓ 707ms | ✓ 1702ms | ✓ 331ms | ✓ 1244ms | 否 | http |
| 104.128.138.186:1080 | ✓ 1721ms | ✓ 1900ms | ✓ 1590ms | ✓ 1740ms | ✓ 1605ms | http |
| 20.78.213.56:80 | ✓ 1573ms | ✓ 1410ms | ✓ 1392ms | ✓ 1399ms | ✓ 1186ms | http |
| 38.199.71.83:999 | ✓ 823ms | ✓ 1544ms | ✓ 1193ms | 否 | ✓ 1631ms | http |
| 217.77.102.18:3128 | ✓ 891ms | ✓ 1739ms | ✓ 1703ms | 否 | ✓ 1488ms | http |
| 62.133.60.126:24558 | ✓ 670ms | 否 | ✓ 618ms | ✓ 1404ms | 否 | http |
| 213.220.3.38:3128 | ✓ 1405ms | ✓ 1896ms | ✓ 811ms | ✓ 1608ms | ✓ 1233ms | http |
| 177.231.162.9:999 | ✓ 1154ms | ✓ 1712ms | ✓ 464ms | 否 | 否 | http |
| 181.204.113.254:11211 | ✓ 1411ms | 否 | 否 | ✓ 1593ms | ✓ 1431ms | http |
| 120.79.99.232:8099 | ✓ 1418ms | ✓ 1566ms | ✓ 1433ms | ✓ 1610ms | ✓ 1314ms | http |

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
