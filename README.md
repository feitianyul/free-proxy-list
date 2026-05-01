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

最后更新：2026-05-01 22:45:44 UTC（2026-05-02 06:45:44 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 342ms | 否 | ✓ 691ms | ✓ 1114ms | ✓ 1076ms | http |
| 1.231.81.166:3128 | ✓ 1648ms | ✓ 1239ms | ✓ 1226ms | ✓ 1454ms | ✓ 987ms | http |
| 115.231.181.40:8128 | ✓ 1202ms | ✓ 1554ms | 否 | ✓ 1246ms | ✓ 1027ms | http |
| 223.84.151.86:30005 | ✓ 1411ms | ✓ 1370ms | ✓ 1527ms | 否 | ✓ 1749ms | http |
| 45.167.124.71:999 | ✓ 1090ms | 否 | ✓ 831ms | 否 | ✓ 1489ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1262ms | ✓ 972ms | ✓ 1112ms | 否 | http |
| 206.206.126.177:2412 | ✓ 807ms | ✓ 1673ms | ✓ 1031ms | ✓ 1323ms | ✓ 888ms | http |
| 222.107.27.7:8074 | ✓ 1956ms | ✓ 1409ms | ✓ 997ms | ✓ 1127ms | ✓ 855ms | http |
| 46.101.95.183:8888 | ✓ 915ms | 否 | ✓ 999ms | ✓ 1585ms | ✓ 1401ms | http |
| 80.92.204.47:1081 | ✓ 772ms | ✓ 1322ms | ✓ 567ms | 否 | 否 | http |
| 92.119.127.208:6005 | ✓ 1162ms | ✓ 1755ms | ✓ 1637ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1861ms | ✓ 1685ms | ✓ 861ms | 否 | 否 | http |
| 72.11.151.159:6005 | ✓ 736ms | ✓ 1157ms | ✓ 1005ms | ✓ 1300ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1557ms | 否 | ✓ 1724ms | ✓ 1946ms | ✓ 1419ms | http |
| 112.65.132.181:3128 | ✓ 1742ms | ✓ 860ms | 否 | ✓ 1957ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1273ms | ✓ 1590ms | ✓ 525ms | ✓ 1959ms | ✓ 1272ms | http |
| 45.153.231.229:8080 | ✓ 898ms | ✓ 1603ms | ✓ 1479ms | 否 | 否 | http |
| 91.184.241.12:443 | 否 | 否 | ✓ 1951ms | ✓ 1858ms | ✓ 1652ms | http |
| 120.92.212.16:7890 | ✓ 1262ms | ✓ 1379ms | ✓ 1084ms | ✓ 1504ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1423ms | ✓ 1259ms | ✓ 1542ms | 否 | 否 | http |
| 119.18.146.65:10001 | 否 | 否 | ✓ 1717ms | ✓ 1656ms | ✓ 1592ms | http |
| 120.92.108.86:7890 | ✓ 1631ms | 否 | ✓ 1870ms | 否 | ✓ 1543ms | http |
| 3.238.34.111:22292 | ✓ 767ms | 否 | ✓ 1121ms | ✓ 1959ms | ✓ 1442ms | http |
| 3.8.3.11:13811 | ✓ 698ms | 否 | ✓ 1688ms | ✓ 1733ms | ✓ 1978ms | http |
| 3.121.42.224:13603 | ✓ 768ms | 否 | ✓ 1705ms | 否 | ✓ 1586ms | http |
| 51.44.97.6:20383 | ✓ 1193ms | 否 | ✓ 1908ms | 否 | ✓ 1186ms | http |
| 212.58.132.5:8888 | ✓ 1195ms | 否 | ✓ 1448ms | ✓ 1818ms | ✓ 1284ms | http |
| 148.230.4.241:999 | ✓ 1005ms | ✓ 1394ms | ✓ 638ms | ✓ 1637ms | ✓ 1445ms | http |
| 59.46.216.131:30001 | ✓ 1027ms | ✓ 1521ms | ✓ 1163ms | 否 | 否 | http |
| 218.108.131.186:17890 | ✓ 743ms | ✓ 885ms | ✓ 771ms | ✓ 949ms | ✓ 833ms | http |
| 45.129.141.143:3128 | ✓ 1218ms | ✓ 1889ms | ✓ 1287ms | 否 | ✓ 1817ms | http |
| 103.157.200.126:3128 | ✓ 1757ms | 否 | ✓ 1619ms | 否 | ✓ 1929ms | http |
| 220.197.44.36:3128 | ✓ 1862ms | 否 | ✓ 1622ms | 否 | ✓ 1643ms | http |
| 72.11.150.178:6005 | ✓ 1505ms | ✓ 1089ms | ✓ 841ms | ✓ 1142ms | ✓ 1029ms | http |
| 47.112.25.109:7890 | ✓ 1953ms | 否 | ✓ 1920ms | ✓ 1472ms | ✓ 1193ms | http |
| 38.180.2.107:3128 | ✓ 922ms | ✓ 1834ms | ✓ 1785ms | 否 | 否 | http |
| 202.129.206.239:3128 | ✓ 1571ms | 否 | ✓ 1885ms | ✓ 1673ms | ✓ 1626ms | http |
| 8.219.97.248:80 | ✓ 1438ms | 否 | ✓ 1484ms | ✓ 1509ms | ✓ 1672ms | http |
| 86.104.74.110:1081 | ✓ 940ms | ✓ 1309ms | ✓ 815ms | 否 | 否 | http |
| 86.104.72.220:1081 | ✓ 1984ms | ✓ 1003ms | ✓ 129ms | ✓ 1155ms | ✓ 890ms | http |
| 52.59.51.29:18781 | ✓ 990ms | 否 | ✓ 1913ms | 否 | ✓ 1671ms | http |
| 34.246.223.187:7592 | ✓ 676ms | 否 | ✓ 1341ms | 否 | ✓ 1622ms | http |
| 152.32.132.190:7890 | ✓ 914ms | ✓ 1954ms | 否 | ✓ 995ms | ✓ 1072ms | http |
| 8.154.21.175:3128 | ✓ 913ms | ✓ 1109ms | ✓ 1836ms | ✓ 1252ms | ✓ 972ms | http |
| 98.153.152.141:7070 | ✓ 447ms | ✓ 905ms | 否 | ✓ 1817ms | ✓ 1028ms | http |
| 185.125.203.192:3128 | ✓ 1205ms | ✓ 1625ms | ✓ 453ms | 否 | 否 | http |
| 167.160.191.204:6005 | ✓ 583ms | ✓ 1293ms | ✓ 843ms | ✓ 1328ms | ✓ 1173ms | http |
| 43.133.44.89:8888 | 否 | 否 | ✓ 1201ms | ✓ 1287ms | ✓ 1813ms | http |
| 89.208.106.138:10808 | ✓ 528ms | ✓ 1549ms | 否 | 否 | ✓ 1154ms | http |
| 86.104.72.220:1082 | 否 | ✓ 953ms | ✓ 1223ms | ✓ 1307ms | ✓ 1093ms | http |
| 128.199.121.61:9090 | ✓ 1599ms | 否 | ✓ 1736ms | ✓ 1327ms | ✓ 1355ms | http |
| 183.238.3.150:7897 | ✓ 912ms | ✓ 1233ms | ✓ 982ms | ✓ 1217ms | ✓ 949ms | http |
| 3.101.133.120:80 | 否 | ✓ 1311ms | ✓ 1842ms | ✓ 1045ms | ✓ 965ms | http |
| 5.253.43.103:3128 | 否 | 否 | ✓ 1480ms | ✓ 1962ms | ✓ 1548ms | http |
| 43.167.237.94:3128 | ✓ 737ms | ✓ 980ms | ✓ 700ms | ✓ 895ms | ✓ 1938ms | http |
| 54.229.201.146:14256 | ✓ 1070ms | 否 | ✓ 1440ms | ✓ 1794ms | ✓ 1394ms | http |
| 13.51.196.44:3629 | ✓ 1062ms | 否 | ✓ 1876ms | 否 | ✓ 1499ms | http |
| 52.59.218.12:16505 | ✓ 1063ms | 否 | ✓ 1712ms | ✓ 1857ms | 否 | http |
| 63.179.134.206:6902 | ✓ 1099ms | 否 | ✓ 1441ms | 否 | ✓ 1738ms | http |
| 81.26.190.143:1080 | ✓ 1198ms | 否 | ✓ 824ms | 否 | ✓ 1876ms | http |
| 92.119.127.211:6005 | ✓ 1082ms | ✓ 1929ms | ✓ 1240ms | ✓ 1624ms | ✓ 1803ms | http |
| 193.122.96.242:3128 | ✓ 1121ms | 否 | 否 | ✓ 1513ms | ✓ 1330ms | http |
| 160.238.65.5:3128 | ✓ 1136ms | ✓ 1880ms | 否 | 否 | ✓ 1226ms | http |
| 103.162.63.226:8082 | ✓ 1905ms | 否 | 否 | ✓ 1952ms | ✓ 1859ms | http |
| 121.230.8.136:1080 | ✓ 1092ms | 否 | ✓ 1116ms | ✓ 1409ms | 否 | http |
| 130.61.174.200:1080 | 否 | 否 | ✓ 824ms | ✓ 1506ms | ✓ 1030ms | http |
| 8.211.166.184:8081 | ✓ 573ms | 否 | ✓ 776ms | ✓ 955ms | ✓ 752ms | http |
| 152.42.177.32:8888 | ✓ 1063ms | 否 | ✓ 1116ms | ✓ 1418ms | ✓ 1148ms | http |
| 222.107.27.7:8089 | 否 | 否 | ✓ 1012ms | ✓ 1083ms | ✓ 956ms | http |
| 92.119.127.213:6005 | ✓ 1386ms | 否 | 否 | ✓ 1773ms | ✓ 1547ms | http |
| 121.230.9.160:1080 | 否 | 否 | ✓ 1202ms | ✓ 1498ms | ✓ 1168ms | http |
| 222.107.27.7:8053 | 否 | ✓ 1320ms | ✓ 1151ms | ✓ 1169ms | ✓ 988ms | http |
| 38.130.38.97:8080 | 否 | ✓ 1673ms | 否 | ✓ 1488ms | ✓ 1273ms | http |
| 45.140.147.155:1082 | ✓ 1068ms | ✓ 1629ms | 否 | ✓ 1720ms | ✓ 1558ms | http |
| 16.18.37.186:41511 | ✓ 723ms | 否 | ✓ 1431ms | ✓ 1917ms | ✓ 1692ms | http |
| 103.82.23.118:5182 | ✓ 1543ms | 否 | ✓ 1694ms | 否 | ✓ 1581ms | http |
| 101.32.243.189:80 | ✓ 1319ms | 否 | ✓ 1636ms | ✓ 1541ms | ✓ 1354ms | http |
| 103.39.51.207:8080 | ✓ 1568ms | 否 | 否 | ✓ 1486ms | ✓ 1501ms | http |
| 47.101.159.19:8899 | ✓ 1235ms | ✓ 1112ms | ✓ 956ms | ✓ 1176ms | ✓ 1009ms | http |

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
