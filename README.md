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

最后更新：2026-06-11 22:14:28 UTC（2026-06-12 06:14:28 UTC+8）

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
| 34.43.46.91:443 | ✓ 450ms | ✓ 1486ms | 否 | 否 | ✓ 1490ms | http |
| 37.49.224.15:3128 | ✓ 732ms | 否 | 否 | ✓ 1928ms | ✓ 1731ms | http |
| 34.43.46.91:80 | ✓ 499ms | ✓ 1520ms | ✓ 941ms | ✓ 1560ms | ✓ 1075ms | http |
| 83.147.36.155:8080 | ✓ 1069ms | ✓ 905ms | ✓ 936ms | ✓ 1705ms | 否 | http |
| 47.80.103.120:8080 | ✓ 1091ms | ✓ 1989ms | ✓ 974ms | ✓ 1305ms | ✓ 1025ms | http |
| 217.154.155.115:8080 | ✓ 908ms | 否 | ✓ 1410ms | 否 | ✓ 1720ms | http |
| 47.80.112.92:8080 | 否 | ✓ 1803ms | 否 | ✓ 1222ms | ✓ 1018ms | http |
| 8.219.97.248:80 | ✓ 950ms | 否 | ✓ 1071ms | ✓ 1447ms | 否 | http |
| 203.30.9.8:8443 | ✓ 1218ms | 否 | ✓ 1595ms | ✓ 1670ms | ✓ 1587ms | http |
| 217.209.35.22:443 | ✓ 1314ms | 否 | ✓ 1769ms | 否 | ✓ 1660ms | http |
| 148.224.7.58:999 | ✓ 626ms | ✓ 1370ms | ✓ 1198ms | 否 | ✓ 1200ms | http |
| 116.80.80.225:3172 | ✓ 1867ms | ✓ 1950ms | ✓ 1998ms | ✓ 1792ms | ✓ 1628ms | http |
| 116.80.92.228:3172 | ✓ 1846ms | 否 | ✓ 1718ms | 否 | ✓ 1728ms | http |
| 95.3.69.222:8080 | ✓ 1348ms | ✓ 1990ms | ✓ 1194ms | ✓ 1933ms | ✓ 1652ms | http |
| 23.95.76.201:8443 | ✓ 1017ms | ✓ 1988ms | ✓ 1281ms | ✓ 1952ms | 否 | http |
| 203.150.113.199:8080 | ✓ 1631ms | 否 | ✓ 1814ms | ✓ 1749ms | ✓ 1578ms | http |
| 159.198.35.187:1080 | ✓ 835ms | 否 | ✓ 505ms | ✓ 1078ms | 否 | http |
| 57.129.144.178:40000 | ✓ 1356ms | 否 | ✓ 1010ms | ✓ 1639ms | ✓ 1806ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1097ms | ✓ 1971ms | ✓ 931ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1278ms | ✓ 1387ms | 否 | ✓ 1568ms | 否 | http |
| 176.111.37.216:39811 | ✓ 1659ms | ✓ 1746ms | ✓ 1812ms | 否 | ✓ 1858ms | http |
| 116.80.60.245:3172 | ✓ 1866ms | 否 | 否 | ✓ 1792ms | ✓ 1628ms | http |
| 116.80.63.158:3172 | ✓ 1865ms | 否 | ✓ 1533ms | ✓ 1805ms | ✓ 1616ms | http |
| 45.186.6.104:3128 | ✓ 1564ms | ✓ 1841ms | ✓ 1937ms | 否 | 否 | http |
| 185.200.188.234:10001 | ✓ 1895ms | 否 | ✓ 1153ms | ✓ 1841ms | ✓ 1678ms | http |
| 223.204.176.6:3128 | ✓ 1629ms | 否 | ✓ 1369ms | 否 | ✓ 1518ms | http |
| 217.76.245.80:999 | ✓ 712ms | ✓ 1724ms | ✓ 1303ms | ✓ 1660ms | ✓ 1396ms | http |
| 38.145.218.142:3128 | ✓ 186ms | ✓ 595ms | ✓ 574ms | ✓ 1075ms | ✓ 517ms | http |
| 187.102.236.161:8080 | ✓ 605ms | ✓ 1103ms | ✓ 332ms | 否 | 否 | http |
| 107.172.139.2:3128 | ✓ 941ms | ✓ 1419ms | ✓ 980ms | ✓ 1312ms | ✓ 1256ms | http |
| 180.2.108.38:8080 | ✓ 1624ms | 否 | ✓ 905ms | ✓ 833ms | ✓ 1379ms | http |
| 84.47.150.125:1080 | ✓ 1347ms | 否 | ✓ 1353ms | 否 | ✓ 1489ms | http |
| 167.99.173.119:3128 | ✓ 716ms | ✓ 958ms | ✓ 917ms | ✓ 1358ms | 否 | http |
| 116.80.64.3:3172 | ✓ 1613ms | 否 | ✓ 1482ms | 否 | ✓ 1620ms | http |
| 36.70.154.124:8080 | ✓ 1763ms | 否 | ✓ 1659ms | ✓ 1940ms | ✓ 1870ms | http |
| 190.60.62.20:999 | ✓ 1513ms | ✓ 1545ms | ✓ 1733ms | 否 | 否 | http |
| 34.71.229.255:3128 | ✓ 245ms | ✓ 1293ms | ✓ 240ms | ✓ 1246ms | ✓ 791ms | http |
| 194.59.204.87:9080 | ✓ 1090ms | ✓ 1933ms | ✓ 1697ms | 否 | 否 | http |
| 91.149.222.102:22335 | ✓ 1391ms | 否 | ✓ 1693ms | 否 | ✓ 1827ms | http |
| 136.226.254.24:11814 | 否 | ✓ 1927ms | ✓ 1518ms | ✓ 1927ms | ✓ 1480ms | http |
| 122.224.198.218:808 | ✓ 1837ms | 否 | ✓ 1984ms | 否 | ✓ 1986ms | http |
| 103.176.97.119:8082 | ✓ 1482ms | 否 | ✓ 1743ms | ✓ 1473ms | 否 | http |
| 157.230.220.25:4857 | 否 | ✓ 1543ms | ✓ 992ms | 否 | ✓ 937ms | http |
| 119.28.16.210:3128 | ✓ 890ms | 否 | ✓ 1134ms | ✓ 1095ms | 否 | http |
| 154.206.67.83:9000 | ✓ 620ms | 否 | ✓ 618ms | ✓ 787ms | ✓ 623ms | http |
| 136.226.254.24:10285 | ✓ 1609ms | ✓ 1990ms | ✓ 1317ms | ✓ 1479ms | ✓ 1343ms | http |
| 136.226.254.24:10517 | ✓ 1684ms | 否 | ✓ 1227ms | ✓ 1571ms | ✓ 1272ms | http |
| 136.226.254.24:11276 | ✓ 1590ms | ✓ 1972ms | ✓ 1496ms | ✓ 1610ms | ✓ 1198ms | http |
| 136.226.254.24:11197 | ✓ 1408ms | ✓ 1986ms | ✓ 1552ms | ✓ 1535ms | ✓ 1418ms | http |
| 136.226.254.24:10919 | ✓ 1503ms | ✓ 1904ms | ✓ 1371ms | ✓ 1746ms | ✓ 1395ms | http |
| 136.226.254.24:9400 | ✓ 1640ms | ✓ 1993ms | ✓ 1315ms | ✓ 1524ms | ✓ 1596ms | http |
| 58.69.124.137:5050 | ✓ 1328ms | 否 | ✓ 1658ms | ✓ 1390ms | ✓ 1405ms | http |
| 116.80.91.144:3172 | ✓ 1792ms | 否 | 否 | ✓ 1852ms | ✓ 1794ms | http |
| 41.216.191.177:8080 | 否 | 否 | ✓ 1531ms | ✓ 1549ms | ✓ 1529ms | http |
| 121.43.196.210:8222 | ✓ 1044ms | ✓ 1068ms | ✓ 793ms | ✓ 1106ms | ✓ 913ms | http |
| 121.43.196.213:8222 | ✓ 976ms | ✓ 1031ms | ✓ 863ms | ✓ 1107ms | ✓ 968ms | http |
| 199.127.62.89:3129 | ✓ 1232ms | 否 | ✓ 1368ms | 否 | ✓ 1692ms | http |
| 104.154.186.48:80 | ✓ 649ms | ✓ 1570ms | ✓ 756ms | ✓ 1743ms | ✓ 1084ms | http |
| 152.67.191.232:6800 | ✓ 1648ms | 否 | ✓ 1090ms | ✓ 1504ms | ✓ 1170ms | http |
| 103.157.117.226:81 | ✓ 1266ms | 否 | ✓ 1227ms | ✓ 1386ms | 否 | http |
| 18.180.59.181:80 | ✓ 594ms | ✓ 851ms | ✓ 1181ms | ✓ 826ms | ✓ 727ms | http |
| 43.156.228.168:80 | ✓ 1419ms | ✓ 1334ms | ✓ 1606ms | ✓ 1493ms | ✓ 1220ms | http |
| 185.219.41.70:2002 | 否 | 否 | ✓ 1185ms | ✓ 1834ms | ✓ 1781ms | http |
| 123.112.208.42:8888 | ✓ 1185ms | 否 | ✓ 1564ms | ✓ 1839ms | 否 | http |
| 106.10.55.212:1121 | ✓ 1639ms | ✓ 1309ms | ✓ 1003ms | ✓ 1345ms | ✓ 1003ms | http |
| 181.79.91.101:999 | ✓ 1714ms | ✓ 1873ms | ✓ 844ms | ✓ 1914ms | ✓ 1480ms | http |
| 217.142.238.66:1080 | ✓ 1659ms | 否 | ✓ 729ms | 否 | ✓ 713ms | http |
| 116.80.92.196:3172 | ✓ 1483ms | 否 | 否 | ✓ 1803ms | ✓ 1644ms | http |
| 193.29.224.20:3128 | ✓ 1591ms | 否 | ✓ 1356ms | 否 | ✓ 1848ms | http |
| 116.80.96.213:3172 | ✓ 1628ms | ✓ 1969ms | ✓ 1500ms | 否 | 否 | http |
| 24.199.113.242:443 | ✓ 788ms | 否 | ✓ 947ms | 否 | ✓ 561ms | http |
| 34.87.80.221:30000 | ✓ 801ms | ✓ 1162ms | ✓ 1555ms | ✓ 1040ms | 否 | http |
| 147.45.170.190:3128 | 否 | ✓ 1728ms | 否 | ✓ 1556ms | ✓ 1415ms | http |
| 169.40.6.114:3128 | ✓ 1018ms | 否 | ✓ 1260ms | 否 | ✓ 1821ms | http |
| 190.147.40.192:9021 | ✓ 1232ms | ✓ 1913ms | ✓ 1652ms | ✓ 1785ms | ✓ 1279ms | http |
| 116.80.91.203:3172 | 否 | 否 | ✓ 1549ms | ✓ 1772ms | ✓ 1645ms | http |
| 170.106.136.181:31002 | ✓ 1455ms | ✓ 1229ms | 否 | 否 | ✓ 1526ms | http |
| 1.231.81.166:3128 | ✓ 1692ms | ✓ 1380ms | 否 | ✓ 1546ms | ✓ 1856ms | http |
| 61.52.131.172:8443 | ✓ 883ms | ✓ 1103ms | ✓ 899ms | ✓ 1133ms | ✓ 913ms | http |
| 129.213.162.27:17777 | ✓ 922ms | 否 | 否 | ✓ 1962ms | ✓ 1345ms | http |
| 103.172.70.173:8080 | ✓ 1357ms | ✓ 1934ms | ✓ 1276ms | ✓ 1383ms | ✓ 1351ms | http |
| 199.68.217.2:3128 | ✓ 463ms | ✓ 895ms | 否 | ✓ 895ms | ✓ 812ms | http |
| 103.157.117.116:8080 | ✓ 1931ms | 否 | ✓ 1499ms | ✓ 1700ms | ✓ 1850ms | http |
| 103.39.51.207:8080 | ✓ 1582ms | 否 | 否 | ✓ 1604ms | ✓ 1987ms | http |

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
