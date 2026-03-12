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

最后更新：2026-03-12 07:55:13 UTC（2026-03-12 15:55:13 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.150.4.3:10000 | ✓ 846ms | ✓ 1488ms | ✓ 1017ms | ✓ 891ms | ✓ 1644ms | http |
| 45.136.130.175:8443 | ✓ 1807ms | 否 | ✓ 283ms | ✓ 988ms | ✓ 716ms | http |
| 107.172.125.217:3128 | 否 | 否 | ✓ 654ms | ✓ 949ms | ✓ 755ms | http |
| 45.136.131.63:8443 | ✓ 826ms | ✓ 1521ms | ✓ 1012ms | ✓ 1716ms | ✓ 736ms | http |
| 46.183.25.8:443 | ✓ 905ms | 否 | ✓ 1277ms | ✓ 1556ms | ✓ 1536ms | http |
| 103.84.95.54:7890 | ✓ 819ms | 否 | ✓ 987ms | ✓ 1067ms | ✓ 899ms | http |
| 45.136.130.239:8443 | 否 | ✓ 997ms | 否 | ✓ 1573ms | ✓ 933ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1728ms | ✓ 994ms | ✓ 1156ms | ✓ 909ms | http |
| 45.136.131.47:8443 | ✓ 1070ms | ✓ 1417ms | ✓ 938ms | 否 | 否 | http |
| 14.225.222.185:7890 | ✓ 1513ms | 否 | ✓ 1061ms | ✓ 1404ms | ✓ 1179ms | http |
| 107.173.52.58:7890 | ✓ 893ms | 否 | 否 | ✓ 1711ms | ✓ 1339ms | http |
| 113.160.132.26:8080 | ✓ 1492ms | 否 | ✓ 1474ms | ✓ 1418ms | ✓ 1173ms | http |
| 171.251.172.78:5105 | ✓ 1517ms | 否 | ✓ 1756ms | ✓ 1855ms | ✓ 1606ms | http |
| 171.251.172.78:5104 | ✓ 1530ms | 否 | ✓ 1826ms | ✓ 1834ms | ✓ 1607ms | http |
| 171.251.172.78:5102 | ✓ 1920ms | 否 | ✓ 1860ms | ✓ 1766ms | ✓ 1594ms | http |
| 47.77.193.180:1080 | ✓ 332ms | ✓ 1849ms | ✓ 436ms | ✓ 966ms | ✓ 671ms | http |
| 62.113.119.14:8080 | ✓ 870ms | 否 | ✓ 1105ms | ✓ 1404ms | ✓ 1227ms | http |
| 120.238.159.228:22222 | ✓ 1094ms | ✓ 1339ms | ✓ 1244ms | ✓ 1440ms | ✓ 1062ms | http |
| 222.184.48.251:22222 | 否 | ✓ 1291ms | ✓ 1179ms | ✓ 1381ms | ✓ 1121ms | http |
| 120.238.159.230:22222 | ✓ 1162ms | ✓ 1466ms | ✓ 1309ms | 否 | ✓ 1056ms | http |
| 210.223.44.230:3128 | ✓ 1924ms | 否 | 否 | ✓ 1157ms | ✓ 882ms | http |
| 217.76.245.80:999 | ✓ 695ms | 否 | ✓ 1214ms | ✓ 1418ms | ✓ 1196ms | http |
| 194.213.18.200:443 | ✓ 557ms | ✓ 1131ms | ✓ 781ms | ✓ 1386ms | 否 | http |
| 45.136.130.223:8443 | ✓ 863ms | ✓ 1073ms | ✓ 1264ms | ✓ 1425ms | ✓ 719ms | http |
| 86.53.183.16:1080 | ✓ 462ms | 否 | ✓ 1328ms | 否 | ✓ 1496ms | http |
| 115.231.181.40:8128 | ✓ 1660ms | 否 | ✓ 1182ms | 否 | ✓ 1454ms | http |
| 171.251.172.78:5107 | ✓ 1902ms | 否 | ✓ 1478ms | ✓ 1805ms | 否 | http |
| 171.251.172.78:5106 | 否 | 否 | ✓ 1502ms | ✓ 1910ms | ✓ 1736ms | http |
| 35.225.22.61:80 | 否 | ✓ 1483ms | ✓ 1059ms | ✓ 1060ms | 否 | http |
| 101.43.255.96:80 | ✓ 1209ms | 否 | ✓ 1572ms | ✓ 1470ms | ✓ 1229ms | http |
| 81.70.169.194:80 | 否 | ✓ 1419ms | ✓ 1444ms | ✓ 1416ms | ✓ 1225ms | http |
| 120.92.212.16:8890 | ✓ 1566ms | 否 | ✓ 1180ms | ✓ 1705ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1245ms | ✓ 1797ms | 否 | ✓ 1663ms | ✓ 1602ms | http |
| 91.107.141.42:8081 | ✓ 1777ms | 否 | ✓ 1660ms | ✓ 1806ms | ✓ 1679ms | http |
| 202.155.12.161:443 | ✓ 1842ms | 否 | ✓ 1295ms | 否 | ✓ 1542ms | http |
| 45.136.198.40:3128 | ✓ 1384ms | 否 | ✓ 1411ms | ✓ 1884ms | ✓ 1602ms | http |
| 120.238.159.189:22222 | ✓ 1045ms | ✓ 1848ms | ✓ 1152ms | ✓ 1382ms | ✓ 1037ms | http |
| 185.115.74.185:8080 | ✓ 804ms | ✓ 1991ms | ✓ 1303ms | 否 | 否 | http |
| 117.159.239.56:22222 | ✓ 1938ms | 否 | 否 | ✓ 1301ms | ✓ 1057ms | http |
| 138.124.53.25:7443 | ✓ 1294ms | 否 | ✓ 1600ms | ✓ 1362ms | 否 | http |
| 190.9.109.198:999 | ✓ 1051ms | 否 | ✓ 1168ms | ✓ 1249ms | ✓ 1051ms | http |
| 103.113.70.189:1081 | ✓ 391ms | 否 | 否 | ✓ 1238ms | ✓ 973ms | http |
| 120.240.35.160:22222 | ✓ 1306ms | 否 | ✓ 1200ms | 否 | ✓ 1081ms | http |
| 120.240.35.173:22222 | ✓ 1120ms | ✓ 1519ms | ✓ 1011ms | 否 | ✓ 1099ms | http |
| 113.59.32.161:22222 | ✓ 1227ms | ✓ 1522ms | 否 | ✓ 1538ms | ✓ 1123ms | http |
| 152.42.213.210:443 | ✓ 1564ms | 否 | ✓ 1451ms | ✓ 1292ms | ✓ 1055ms | http |
| 152.42.213.210:8080 | ✓ 1762ms | 否 | ✓ 1435ms | ✓ 1515ms | ✓ 1144ms | http |
| 152.70.98.46:8888 | 否 | 否 | ✓ 1779ms | ✓ 1219ms | ✓ 1202ms | http |
| 160.250.5.22:1 | ✓ 1116ms | 否 | ✓ 1575ms | ✓ 1728ms | ✓ 1250ms | http |
| 160.250.4.245:1 | ✓ 1425ms | 否 | ✓ 1468ms | ✓ 1595ms | ✓ 1214ms | http |
| 144.31.25.69:21064 | ✓ 784ms | 否 | ✓ 1969ms | 否 | ✓ 1909ms | http |
| 103.82.23.118:5234 | ✓ 1724ms | 否 | ✓ 1956ms | ✓ 1933ms | ✓ 1806ms | http |
| 222.184.48.252:22222 | 否 | ✓ 1399ms | ✓ 1267ms | ✓ 1429ms | 否 | http |
| 183.249.5.109:22222 | ✓ 904ms | ✓ 1228ms | ✓ 948ms | 否 | ✓ 1064ms | http |
| 117.159.239.45:22222 | ✓ 1170ms | ✓ 1241ms | ✓ 998ms | ✓ 1310ms | ✓ 1033ms | http |
| 34.101.184.164:3128 | ✓ 1743ms | 否 | ✓ 1429ms | ✓ 1451ms | ✓ 1733ms | http |
| 111.48.191.1:7890 | ✓ 881ms | ✓ 1167ms | ✓ 921ms | ✓ 1167ms | ✓ 953ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1458ms | ✓ 1162ms | 否 | ✓ 1455ms | http |
| 18.100.127.30:9002 | ✓ 1481ms | 否 | ✓ 1844ms | 否 | ✓ 1857ms | http |
| 88.80.150.82:8080 | ✓ 1092ms | 否 | ✓ 1993ms | 否 | ✓ 1943ms | https |
| 18.100.126.55:49035 | ✓ 1860ms | 否 | ✓ 1541ms | 否 | ✓ 1875ms | http |
| 101.32.244.83:8080 | ✓ 1603ms | 否 | ✓ 1167ms | ✓ 1714ms | ✓ 1501ms | http |
| 121.43.196.210:8222 | ✓ 1139ms | ✓ 1326ms | ✓ 1038ms | ✓ 1320ms | ✓ 1025ms | http |
| 121.43.196.213:8222 | ✓ 1139ms | ✓ 1263ms | ✓ 1082ms | ✓ 1346ms | ✓ 1018ms | http |
| 114.55.226.123:10086 | ✓ 1396ms | ✓ 1891ms | ✓ 1186ms | ✓ 1481ms | ✓ 1225ms | http |
| 20.120.225.109:3128 | 否 | ✓ 1438ms | ✓ 1041ms | ✓ 1425ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1190ms | ✓ 1430ms | 否 | 否 | ✓ 1108ms | http |
| 8.219.97.248:80 | ✓ 1750ms | 否 | ✓ 1999ms | ✓ 1775ms | 否 | http |
| 117.159.239.49:22222 | ✓ 1035ms | ✓ 1302ms | ✓ 972ms | ✓ 1267ms | ✓ 1003ms | http |
| 45.140.147.155:1081 | ✓ 390ms | 否 | ✓ 393ms | ✓ 1381ms | ✓ 1211ms | http |
| 45.136.130.191:8443 | 否 | 否 | ✓ 1671ms | ✓ 1994ms | ✓ 1753ms | http |
| 45.136.130.188:8443 | ✓ 1828ms | 否 | ✓ 1216ms | ✓ 1931ms | ✓ 1860ms | http |
| 120.238.159.229:22222 | ✓ 1203ms | ✓ 1523ms | ✓ 1156ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1222ms | 否 | ✓ 1703ms | 否 | ✓ 1173ms | http |
| 205.209.118.30:3138 | ✓ 62ms | ✓ 984ms | ✓ 950ms | ✓ 1026ms | ✓ 787ms | http |
| 45.186.6.104:3128 | ✓ 1998ms | ✓ 1883ms | ✓ 1822ms | 否 | 否 | http |
| 121.138.61.193:8085 | 否 | 否 | ✓ 1474ms | ✓ 1503ms | ✓ 1295ms | http |
| 200.174.198.32:8888 | ✓ 912ms | 否 | ✓ 1763ms | 否 | ✓ 1989ms | http |
| 120.240.35.178:22222 | ✓ 1144ms | ✓ 1456ms | ✓ 1144ms | ✓ 1316ms | ✓ 1143ms | http |
| 120.198.141.79:22222 | ✓ 1157ms | ✓ 1432ms | ✓ 1197ms | ✓ 1420ms | ✓ 1121ms | http |
| 120.240.35.177:22222 | ✓ 1078ms | 否 | 否 | ✓ 1654ms | ✓ 1098ms | http |
| 103.39.51.190:8080 | ✓ 1989ms | 否 | 否 | ✓ 1831ms | ✓ 1692ms | http |
| 210.77.23.212:7890 | ✓ 1290ms | 否 | ✓ 1310ms | ✓ 1706ms | ✓ 1161ms | http |
| 61.52.131.172:8443 | ✓ 1099ms | ✓ 1386ms | ✓ 1064ms | ✓ 1361ms | ✓ 1066ms | http |
| 113.59.32.141:22222 | ✓ 1207ms | ✓ 1611ms | 否 | ✓ 1670ms | ✓ 1212ms | http |
| 113.59.32.145:22222 | ✓ 1464ms | ✓ 1635ms | ✓ 1138ms | ✓ 1451ms | ✓ 1155ms | http |
| 222.184.48.242:22222 | 否 | ✓ 1803ms | ✓ 1058ms | ✓ 1465ms | ✓ 1075ms | http |
| 120.240.35.161:22222 | 否 | ✓ 1482ms | ✓ 1182ms | 否 | ✓ 1043ms | http |
| 222.184.48.235:22222 | ✓ 1110ms | ✓ 1562ms | ✓ 1026ms | 否 | ✓ 1139ms | http |

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
