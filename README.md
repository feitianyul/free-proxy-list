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

最后更新：2026-05-02 21:40:06 UTC（2026-05-03 05:40:06 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:3128 | ✓ 930ms | ✓ 1955ms | 否 | ✓ 1532ms | 否 | http |
| 8.211.166.184:8081 | ✓ 1595ms | ✓ 1138ms | ✓ 828ms | ✓ 1043ms | ✓ 856ms | http |
| 47.77.216.82:1080 | ✓ 510ms | ✓ 1261ms | 否 | 否 | ✓ 773ms | http |
| 47.85.51.197:1080 | ✓ 870ms | ✓ 1069ms | ✓ 491ms | ✓ 1717ms | 否 | http |
| 218.108.131.186:17890 | ✓ 913ms | ✓ 1212ms | ✓ 1013ms | ✓ 1227ms | ✓ 994ms | http |
| 45.167.124.71:999 | ✓ 1715ms | ✓ 1872ms | ✓ 1373ms | ✓ 1953ms | ✓ 1449ms | http |
| 152.32.132.190:7890 | ✓ 801ms | ✓ 1296ms | 否 | 否 | ✓ 1050ms | http |
| 212.58.132.5:8888 | ✓ 1174ms | 否 | ✓ 1431ms | ✓ 1561ms | ✓ 1263ms | http |
| 109.120.156.122:8090 | ✓ 1430ms | 否 | ✓ 1623ms | 否 | ✓ 1870ms | http |
| 113.160.132.26:8080 | ✓ 1059ms | 否 | ✓ 1421ms | ✓ 1853ms | ✓ 1569ms | http |
| 206.206.126.177:2412 | ✓ 1269ms | 否 | ✓ 821ms | ✓ 1142ms | ✓ 893ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1388ms | ✓ 1561ms | ✓ 1117ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1383ms | 否 | ✓ 1404ms | ✓ 1344ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1392ms | ✓ 1018ms | 否 | ✓ 1383ms | http |
| 120.92.108.86:7890 | ✓ 1473ms | 否 | ✓ 1441ms | ✓ 1825ms | 否 | http |
| 103.35.190.69:1082 | ✓ 790ms | ✓ 993ms | ✓ 558ms | 否 | 否 | http |
| 149.51.42.10:8080 | ✓ 516ms | ✓ 1489ms | 否 | ✓ 1556ms | 否 | http |
| 103.240.6.22:16498 | ✓ 1316ms | 否 | ✓ 1308ms | 否 | ✓ 1600ms | http |
| 72.11.150.178:6005 | ✓ 324ms | ✓ 1103ms | ✓ 1055ms | ✓ 1259ms | ✓ 986ms | http |
| 223.84.151.86:30005 | ✓ 1494ms | ✓ 1798ms | 否 | 否 | ✓ 1593ms | http |
| 91.108.243.203:3128 | ✓ 736ms | 否 | ✓ 1232ms | ✓ 1938ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1903ms | 否 | 否 | ✓ 1883ms | ✓ 1343ms | http |
| 1.231.81.166:3128 | ✓ 778ms | ✓ 1002ms | ✓ 964ms | ✓ 1187ms | ✓ 905ms | http |
| 62.113.119.14:8080 | 否 | ✓ 1394ms | ✓ 1027ms | 否 | ✓ 1260ms | http |
| 121.230.8.237:1080 | ✓ 1401ms | ✓ 1691ms | ✓ 1063ms | ✓ 1738ms | ✓ 1278ms | http |
| 62.60.231.71:56608 | 否 | ✓ 1943ms | ✓ 1151ms | ✓ 1847ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1162ms | 否 | ✓ 1218ms | ✓ 1470ms | 否 | http |
| 80.92.204.47:1081 | ✓ 987ms | ✓ 1792ms | ✓ 443ms | 否 | 否 | http |
| 86.104.72.220:1082 | ✓ 796ms | 否 | ✓ 1078ms | ✓ 1122ms | ✓ 829ms | http |
| 110.172.28.217:3128 | ✓ 1543ms | 否 | ✓ 959ms | ✓ 1335ms | ✓ 1235ms | http |
| 180.103.19.151:1080 | ✓ 1144ms | ✓ 1642ms | ✓ 1417ms | ✓ 1748ms | ✓ 1147ms | http |
| 105.159.147.94:5928 | ✓ 1217ms | ✓ 1702ms | ✓ 1486ms | ✓ 1803ms | ✓ 1892ms | http |
| 105.159.147.132:4399 | ✓ 1217ms | ✓ 1895ms | ✓ 1187ms | ✓ 1886ms | ✓ 1922ms | http |
| 105.159.147.132:5952 | ✓ 1213ms | ✓ 1847ms | ✓ 1268ms | 否 | ✓ 1789ms | http |
| 105.159.146.139:4961 | ✓ 1214ms | ✓ 1734ms | ✓ 1362ms | 否 | 否 | http |
| 105.159.147.132:4166 | ✓ 1215ms | ✓ 1532ms | ✓ 1664ms | 否 | ✓ 1658ms | http |
| 105.159.150.184:5664 | ✓ 1216ms | ✓ 1754ms | ✓ 1325ms | ✓ 1973ms | ✓ 1831ms | http |
| 105.159.154.224:4146 | ✓ 1215ms | ✓ 1833ms | ✓ 1251ms | ✓ 1977ms | ✓ 1827ms | http |
| 105.159.154.224:4010 | ✓ 1217ms | ✓ 1738ms | ✓ 1419ms | ✓ 1825ms | ✓ 1954ms | http |
| 105.159.154.4:4635 | ✓ 1216ms | ✓ 1700ms | ✓ 1383ms | 否 | ✓ 1859ms | http |
| 105.159.147.94:4150 | ✓ 1217ms | ✓ 1578ms | ✓ 1545ms | ✓ 1856ms | 否 | http |
| 105.159.148.106:5678 | ✓ 1213ms | ✓ 1920ms | ✓ 1164ms | 否 | ✓ 1964ms | http |
| 105.159.148.42:5678 | ✓ 1248ms | ✓ 1834ms | ✓ 1319ms | 否 | ✓ 1868ms | http |
| 105.159.150.184:4701 | ✓ 1214ms | ✓ 1617ms | ✓ 1512ms | 否 | ✓ 1979ms | http |
| 101.32.243.189:80 | ✓ 1337ms | ✓ 1737ms | ✓ 1672ms | ✓ 1603ms | 否 | http |
| 8.154.21.175:3128 | ✓ 1254ms | ✓ 1819ms | 否 | ✓ 1284ms | ✓ 1067ms | http |
| 62.133.60.126:24558 | ✓ 879ms | 否 | ✓ 1312ms | ✓ 1944ms | 否 | http |
| 34.96.238.40:8080 | ✓ 939ms | 否 | ✓ 954ms | 否 | ✓ 1138ms | http |
| 92.119.127.208:6005 | ✓ 1309ms | ✓ 1584ms | ✓ 1714ms | ✓ 1823ms | 否 | http |
| 147.45.178.211:14658 | ✓ 625ms | ✓ 1687ms | ✓ 1755ms | 否 | 否 | http |
| 121.230.8.136:1080 | ✓ 1826ms | 否 | ✓ 1782ms | 否 | ✓ 1620ms | http |
| 3.101.133.120:80 | ✓ 282ms | ✓ 1270ms | ✓ 1014ms | ✓ 1404ms | ✓ 1212ms | http |
| 47.112.25.109:7890 | ✓ 1087ms | 否 | ✓ 1542ms | ✓ 1951ms | 否 | http |
| 117.236.124.166:3128 | ✓ 1003ms | 否 | ✓ 1036ms | 否 | ✓ 1533ms | http |
| 119.195.17.15:3056 | ✓ 1991ms | ✓ 1320ms | 否 | 否 | ✓ 1811ms | http |
| 86.104.74.110:1081 | ✓ 1108ms | ✓ 1713ms | ✓ 794ms | ✓ 1994ms | ✓ 1258ms | http |
| 86.104.74.110:1082 | ✓ 1107ms | ✓ 1747ms | ✓ 769ms | ✓ 1974ms | ✓ 1209ms | http |
| 210.223.44.230:3128 | ✓ 1359ms | ✓ 1495ms | ✓ 1475ms | ✓ 1215ms | ✓ 1627ms | http |
| 86.104.72.219:1082 | ✓ 383ms | ✓ 1061ms | ✓ 1184ms | 否 | 否 | http |
| 86.104.72.219:1081 | ✓ 388ms | ✓ 1056ms | ✓ 1081ms | ✓ 1377ms | ✓ 1916ms | http |
| 121.230.9.225:1080 | ✓ 1001ms | ✓ 1351ms | ✓ 975ms | ✓ 1309ms | ✓ 998ms | http |
| 121.230.9.231:1080 | ✓ 937ms | ✓ 1221ms | ✓ 1072ms | ✓ 1381ms | ✓ 1102ms | http |
| 121.230.8.162:1080 | ✓ 989ms | ✓ 1367ms | ✓ 978ms | ✓ 1353ms | ✓ 1077ms | http |
| 154.27.196.3:8080 | ✓ 1245ms | ✓ 1550ms | ✓ 1295ms | 否 | ✓ 1341ms | http |
| 61.52.131.172:8443 | ✓ 981ms | ✓ 1316ms | ✓ 1069ms | ✓ 1369ms | ✓ 1044ms | http |
| 43.133.44.89:8888 | ✓ 1705ms | 否 | ✓ 1826ms | 否 | ✓ 1199ms | http |
| 103.172.70.173:8080 | ✓ 1463ms | 否 | 否 | ✓ 1734ms | ✓ 1476ms | http |
| 148.230.4.241:999 | ✓ 974ms | ✓ 1662ms | ✓ 689ms | ✓ 1362ms | ✓ 1213ms | http |
| 103.39.51.207:8080 | ✓ 1509ms | 否 | 否 | ✓ 1986ms | ✓ 1944ms | http |

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
