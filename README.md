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

最后更新：2026-05-01 00:47:13 UTC（2026-05-01 08:47:13 UTC+8）

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
| 218.108.131.186:17890 | ✓ 880ms | ✓ 1143ms | ✓ 1107ms | ✓ 1209ms | ✓ 994ms | http |
| 1.231.81.166:3128 | ✓ 1544ms | ✓ 1346ms | 否 | ✓ 1084ms | ✓ 905ms | http |
| 212.58.132.5:8888 | ✓ 1531ms | 否 | ✓ 1623ms | ✓ 1505ms | ✓ 1211ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1345ms | ✓ 968ms | ✓ 1489ms | ✓ 978ms | http |
| 43.167.237.94:3128 | ✓ 1514ms | ✓ 1884ms | ✓ 973ms | ✓ 862ms | ✓ 1362ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1276ms | ✓ 991ms | ✓ 1228ms | 否 | http |
| 45.167.124.71:999 | ✓ 1261ms | ✓ 1849ms | ✓ 1616ms | ✓ 1986ms | ✓ 1540ms | http |
| 8.211.166.184:8081 | ✓ 571ms | ✓ 1336ms | ✓ 749ms | ✓ 1016ms | ✓ 738ms | http |
| 77.110.116.224:3128 | ✓ 583ms | 否 | ✓ 1423ms | ✓ 1900ms | 否 | http |
| 206.206.126.177:2412 | ✓ 791ms | 否 | ✓ 1037ms | ✓ 1128ms | ✓ 853ms | http |
| 223.84.151.86:30005 | ✓ 1316ms | ✓ 1308ms | ✓ 1290ms | ✓ 1544ms | ✓ 1300ms | http |
| 154.64.232.35:8080 | ✓ 1173ms | ✓ 1683ms | ✓ 1452ms | ✓ 1999ms | ✓ 1584ms | http |
| 46.101.95.183:8888 | ✓ 1333ms | 否 | ✓ 1566ms | 否 | ✓ 1853ms | http |
| 103.35.190.182:1082 | ✓ 326ms | 否 | 否 | ✓ 1011ms | ✓ 1000ms | http |
| 34.96.238.40:8080 | ✓ 957ms | 否 | ✓ 1297ms | 否 | ✓ 1321ms | http |
| 107.173.160.222:1080 | ✓ 1779ms | ✓ 1632ms | ✓ 1132ms | 否 | 否 | http |
| 103.70.114.149:3128 | ✓ 1498ms | 否 | ✓ 1585ms | ✓ 1619ms | ✓ 1796ms | http |
| 80.92.204.47:1081 | ✓ 472ms | ✓ 1326ms | ✓ 844ms | ✓ 1394ms | ✓ 1322ms | http |
| 103.35.191.138:1082 | ✓ 203ms | 否 | ✓ 356ms | ✓ 1614ms | 否 | http |
| 146.56.164.121:3128 | ✓ 1321ms | 否 | ✓ 1491ms | ✓ 1254ms | ✓ 1143ms | http |
| 37.187.109.70:10111 | ✓ 1847ms | 否 | ✓ 1549ms | 否 | ✓ 1489ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1955ms | ✓ 1020ms | ✓ 983ms | 否 | http |
| 103.35.191.244:1082 | ✓ 143ms | ✓ 877ms | ✓ 132ms | ✓ 1734ms | ✓ 802ms | http |
| 103.35.191.244:1081 | ✓ 197ms | ✓ 902ms | ✓ 129ms | ✓ 1035ms | ✓ 1783ms | http |
| 120.92.212.16:8890 | ✓ 1288ms | ✓ 1561ms | ✓ 1180ms | ✓ 1469ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1229ms | 否 | ✓ 1299ms | 否 | ✓ 1378ms | http |
| 45.153.231.229:8080 | ✓ 932ms | ✓ 1944ms | ✓ 1724ms | ✓ 1923ms | ✓ 1894ms | http |
| 120.92.212.16:7890 | ✓ 1131ms | 否 | ✓ 1338ms | 否 | ✓ 1710ms | http |
| 20.164.75.153:8080 | ✓ 1663ms | 否 | ✓ 1272ms | 否 | ✓ 1890ms | http |
| 47.85.51.197:1080 | ✓ 1244ms | ✓ 1267ms | ✓ 828ms | 否 | 否 | http |
| 130.61.174.200:1080 | ✓ 1870ms | ✓ 1820ms | 否 | 否 | ✓ 1489ms | http |
| 8.154.21.175:3128 | ✓ 898ms | ✓ 1448ms | ✓ 899ms | ✓ 1187ms | ✓ 980ms | http |
| 200.174.198.32:8888 | ✓ 1138ms | 否 | ✓ 1702ms | 否 | ✓ 1974ms | http |
| 103.157.200.126:3128 | ✓ 1880ms | 否 | ✓ 1990ms | 否 | ✓ 1949ms | http |
| 8.219.97.248:80 | ✓ 1384ms | ✓ 1965ms | 否 | 否 | ✓ 1940ms | http |
| 20.127.128.70:8080 | ✓ 1207ms | ✓ 1912ms | ✓ 637ms | 否 | 否 | http |
| 62.60.231.71:56608 | ✓ 1281ms | 否 | ✓ 1290ms | 否 | ✓ 1346ms | http |
| 45.129.141.143:3128 | ✓ 713ms | ✓ 1816ms | ✓ 1539ms | ✓ 1915ms | ✓ 1721ms | http |
| 106.10.55.212:1121 | ✓ 1355ms | 否 | ✓ 1979ms | ✓ 1519ms | 否 | http |
| 47.121.114.42:3129 | ✓ 1627ms | ✓ 1731ms | 否 | ✓ 1747ms | ✓ 1848ms | http |
| 62.60.237.68:8080 | ✓ 1284ms | 否 | ✓ 1396ms | 否 | ✓ 1578ms | http |
| 59.46.216.131:30001 | ✓ 1075ms | 否 | ✓ 1130ms | 否 | ✓ 1133ms | http |
| 180.103.19.40:1080 | 否 | ✓ 1553ms | ✓ 1170ms | ✓ 1770ms | ✓ 1041ms | http |
| 62.113.119.14:8080 | ✓ 1073ms | 否 | ✓ 1501ms | ✓ 1911ms | 否 | http |
| 173.212.245.136:8888 | ✓ 1127ms | ✓ 1849ms | ✓ 1981ms | 否 | 否 | http |
| 77.239.114.179:3128 | ✓ 1099ms | 否 | ✓ 1506ms | 否 | ✓ 1811ms | http |
| 83.229.73.113:13554 | ✓ 1117ms | ✓ 1919ms | ✓ 1503ms | 否 | 否 | http |
| 34.101.184.164:3128 | ✓ 1094ms | 否 | ✓ 1370ms | ✓ 1390ms | ✓ 1013ms | http |
| 185.41.152.110:3128 | ✓ 1105ms | 否 | ✓ 1475ms | 否 | ✓ 1744ms | http |
| 45.168.238.193:8443 | ✓ 880ms | ✓ 1024ms | ✓ 850ms | ✓ 1451ms | ✓ 1249ms | http |
| 220.197.44.36:3128 | ✓ 1301ms | ✓ 1545ms | ✓ 1377ms | ✓ 1433ms | ✓ 1168ms | http |
| 101.36.105.101:9128 | ✓ 1469ms | 否 | 否 | ✓ 1286ms | ✓ 1716ms | http |
| 101.32.243.189:80 | ✓ 1544ms | 否 | ✓ 1673ms | ✓ 1507ms | ✓ 1711ms | http |
| 150.107.140.238:3128 | ✓ 1320ms | 否 | ✓ 1190ms | ✓ 1210ms | ✓ 950ms | http |
| 65.109.213.99:1080 | ✓ 1226ms | ✓ 1820ms | ✓ 631ms | 否 | 否 | http |
| 183.238.3.150:7897 | ✓ 938ms | ✓ 1269ms | ✓ 1043ms | ✓ 1213ms | ✓ 991ms | http |
| 218.72.124.35:40000 | ✓ 1116ms | ✓ 1225ms | ✓ 989ms | ✓ 1302ms | ✓ 1024ms | http |
| 103.35.191.174:1082 | ✓ 641ms | ✓ 1293ms | ✓ 886ms | 否 | 否 | http |
| 3.101.133.120:80 | ✓ 322ms | ✓ 1341ms | ✓ 1437ms | ✓ 1230ms | ✓ 990ms | http |
| 119.95.170.51:8082 | ✓ 1476ms | 否 | ✓ 1712ms | ✓ 1625ms | ✓ 1552ms | http |
| 77.110.107.80:1080 | ✓ 1449ms | ✓ 1728ms | ✓ 614ms | ✓ 1889ms | ✓ 1224ms | http |
| 139.162.153.201:3128 | ✓ 862ms | 否 | ✓ 1711ms | ✓ 1807ms | ✓ 1488ms | http |
| 103.35.190.182:1081 | ✓ 187ms | 否 | ✓ 456ms | 否 | ✓ 804ms | http |
| 185.230.190.195:3128 | 否 | ✓ 1774ms | ✓ 888ms | ✓ 1927ms | ✓ 1831ms | http |
| 14.143.222.113:57788 | ✓ 1332ms | 否 | ✓ 1929ms | ✓ 1464ms | 否 | http |
| 44.201.32.14:31163 | ✓ 661ms | 否 | ✓ 1670ms | 否 | ✓ 1251ms | http |
| 18.170.25.193:927 | ✓ 725ms | 否 | ✓ 1427ms | 否 | ✓ 1257ms | http |
| 40.177.89.65:26257 | ✓ 1264ms | 否 | ✓ 1328ms | 否 | ✓ 1703ms | http |
| 52.59.218.12:45074 | ✓ 673ms | 否 | ✓ 1479ms | 否 | ✓ 1897ms | http |
| 13.41.196.179:50565 | ✓ 1013ms | 否 | 否 | ✓ 1929ms | ✓ 1373ms | http |
| 103.35.190.69:1081 | 否 | ✓ 1986ms | ✓ 1383ms | 否 | ✓ 1181ms | http |
| 2.78.60.10:3129 | ✓ 1791ms | 否 | ✓ 1563ms | 否 | ✓ 1975ms | http |
| 89.208.106.138:10808 | 否 | ✓ 1881ms | ✓ 1543ms | 否 | ✓ 1571ms | http |
| 143.198.223.214:1084 | ✓ 1216ms | 否 | ✓ 1189ms | ✓ 1139ms | ✓ 911ms | http |
| 185.4.132.222:3128 | ✓ 1019ms | ✓ 1562ms | ✓ 1629ms | ✓ 1771ms | ✓ 1481ms | http |
| 202.129.206.239:3128 | ✓ 1250ms | 否 | ✓ 1250ms | ✓ 1723ms | ✓ 1541ms | http |
| 86.104.72.220:1082 | ✓ 625ms | ✓ 1387ms | ✓ 424ms | ✓ 1414ms | 否 | http |
| 86.104.72.220:1081 | ✓ 575ms | ✓ 1260ms | ✓ 175ms | ✓ 1416ms | 否 | http |
| 137.59.47.73:3128 | 否 | ✓ 1616ms | ✓ 1892ms | 否 | ✓ 1474ms | http |
| 91.217.81.131:1080 | ✓ 1609ms | ✓ 1850ms | 否 | ✓ 1900ms | 否 | http |
| 86.104.72.219:1081 | ✓ 936ms | ✓ 1887ms | ✓ 1053ms | 否 | 否 | http |
| 20.120.225.109:3128 | ✓ 385ms | ✓ 1253ms | ✓ 859ms | ✓ 1293ms | ✓ 605ms | http |
| 185.234.66.87:1081 | ✓ 1788ms | 否 | ✓ 1938ms | ✓ 1751ms | 否 | http |
| 61.52.131.172:8443 | ✓ 993ms | ✓ 1236ms | ✓ 929ms | ✓ 1310ms | ✓ 1023ms | http |
| 103.172.70.173:8080 | ✓ 1449ms | 否 | ✓ 1753ms | ✓ 1463ms | ✓ 1558ms | http |
| 3.121.130.230:52971 | ✓ 1794ms | 否 | ✓ 1442ms | ✓ 1863ms | ✓ 1628ms | http |
| 13.60.163.108:25477 | ✓ 1830ms | 否 | ✓ 1938ms | ✓ 1908ms | ✓ 1761ms | http |
| 173.212.246.157:3128 | ✓ 1011ms | ✓ 1932ms | ✓ 935ms | 否 | 否 | http |
| 223.16.170.103:80 | ✓ 1219ms | 否 | ✓ 1544ms | ✓ 1629ms | ✓ 1200ms | http |
| 103.39.51.207:8080 | ✓ 1478ms | 否 | ✓ 1861ms | 否 | ✓ 1428ms | http |
| 43.133.44.89:8888 | ✓ 1831ms | 否 | ✓ 825ms | ✓ 1091ms | ✓ 1734ms | http |

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
