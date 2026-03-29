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

最后更新：2026-03-29 21:38:33 UTC（2026-03-30 05:38:33 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 215ms | ✓ 638ms | ✓ 934ms | ✓ 830ms | ✓ 660ms | http |
| 43.99.54.236:5555 | ✓ 682ms | ✓ 879ms | ✓ 674ms | 否 | 否 | http |
| 39.185.46.193:5911 | ✓ 634ms | ✓ 802ms | ✓ 632ms | ✓ 966ms | ✓ 734ms | http |
| 147.161.210.140:8800 | 否 | ✓ 1049ms | ✓ 866ms | ✓ 972ms | ✓ 924ms | http |
| 103.84.95.54:7890 | ✓ 645ms | 否 | ✓ 635ms | ✓ 813ms | ✓ 664ms | http |
| 113.160.132.26:8080 | ✓ 1376ms | ✓ 1302ms | ✓ 1253ms | ✓ 1295ms | ✓ 933ms | http |
| 42.96.16.158:1311 | ✓ 1370ms | 否 | ✓ 1136ms | ✓ 1356ms | ✓ 896ms | http |
| 167.103.115.102:8800 | ✓ 1013ms | 否 | ✓ 1015ms | ✓ 1383ms | ✓ 1761ms | http |
| 167.103.34.108:8800 | ✓ 1659ms | ✓ 1930ms | ✓ 1401ms | ✓ 1612ms | ✓ 1327ms | http |
| 106.75.15.167:7890 | ✓ 1089ms | ✓ 1771ms | 否 | ✓ 1104ms | ✓ 1154ms | http |
| 95.213.217.168:52004 | ✓ 1303ms | 否 | ✓ 1403ms | 否 | ✓ 1679ms | http |
| 180.250.219.58:53281 | ✓ 1481ms | 否 | ✓ 1435ms | ✓ 1900ms | ✓ 1823ms | http |
| 35.225.22.61:80 | ✓ 1024ms | ✓ 1450ms | 否 | ✓ 1182ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1239ms | 否 | ✓ 903ms | ✓ 916ms | ✓ 774ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1423ms | ✓ 992ms | ✓ 1454ms | ✓ 970ms | http |
| 167.103.31.122:8800 | ✓ 1839ms | 否 | ✓ 1416ms | 否 | ✓ 1976ms | http |
| 45.167.124.52:8080 | ✓ 713ms | ✓ 1567ms | ✓ 713ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 927ms | ✓ 1150ms | ✓ 908ms | ✓ 1081ms | ✓ 915ms | http |
| 150.107.140.238:3128 | ✓ 919ms | 否 | 否 | ✓ 1419ms | ✓ 1885ms | http |
| 116.80.49.166:3172 | 否 | 否 | ✓ 1462ms | ✓ 1802ms | ✓ 1603ms | http |
| 120.92.212.16:7890 | ✓ 916ms | ✓ 1385ms | 否 | ✓ 1375ms | ✓ 934ms | http |
| 45.12.151.226:2829 | ✓ 1412ms | 否 | ✓ 1765ms | ✓ 1825ms | ✓ 1449ms | http |
| 116.80.49.167:3172 | ✓ 1791ms | 否 | ✓ 1468ms | ✓ 1774ms | 否 | http |
| 116.80.49.162:3172 | ✓ 1792ms | ✓ 1964ms | ✓ 1442ms | 否 | 否 | http |
| 101.43.127.100:8877 | ✓ 919ms | ✓ 1096ms | ✓ 802ms | ✓ 1120ms | ✓ 858ms | http |
| 210.223.44.230:3128 | 否 | ✓ 924ms | ✓ 1005ms | ✓ 957ms | ✓ 799ms | http |
| 106.117.208.101:7890 | ✓ 929ms | ✓ 1128ms | ✓ 1010ms | ✓ 1154ms | ✓ 933ms | http |
| 147.161.239.240:8800 | ✓ 874ms | ✓ 1784ms | ✓ 1460ms | 否 | ✓ 1617ms | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 1639ms | ✓ 1259ms | ✓ 847ms | http |
| 167.103.144.127:8800 | 否 | ✓ 1648ms | ✓ 1174ms | ✓ 1563ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1285ms | 否 | 否 | ✓ 1149ms | ✓ 1534ms | http |
| 38.34.179.173:8450 | 否 | ✓ 895ms | ✓ 1898ms | ✓ 1999ms | 否 | http |
| 177.234.217.88:999 | ✓ 1450ms | ✓ 1907ms | ✓ 1792ms | 否 | 否 | http |
| 38.145.208.171:8451 | 否 | ✓ 658ms | ✓ 535ms | ✓ 1418ms | ✓ 1499ms | http |
| 38.145.208.186:8448 | 否 | 否 | ✓ 242ms | ✓ 690ms | ✓ 672ms | http |
| 38.145.218.134:8445 | 否 | 否 | ✓ 1885ms | ✓ 1873ms | ✓ 1782ms | http |
| 47.95.231.180:8084 | ✓ 827ms | ✓ 1110ms | ✓ 991ms | ✓ 1112ms | ✓ 895ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1110ms | ✓ 1097ms | ✓ 1393ms | ✓ 906ms | http |
| 222.228.171.92:8080 | 否 | ✓ 1836ms | ✓ 1518ms | ✓ 961ms | ✓ 965ms | http |
| 160.238.65.6:3128 | ✓ 1625ms | ✓ 1588ms | ✓ 1388ms | 否 | 否 | http |
| 103.39.51.207:8080 | ✓ 1510ms | 否 | 否 | ✓ 1689ms | ✓ 1389ms | http |
| 8.219.97.248:80 | ✓ 1725ms | 否 | ✓ 1202ms | ✓ 1992ms | ✓ 1029ms | http |
| 38.34.179.86:8452 | 否 | ✓ 659ms | ✓ 1420ms | ✓ 1172ms | ✓ 1269ms | http |
| 219.117.204.211:7799 | ✓ 1738ms | ✓ 1213ms | ✓ 973ms | ✓ 1010ms | 否 | http |
| 116.80.63.64:7777 | 否 | 否 | ✓ 1493ms | ✓ 1783ms | ✓ 1633ms | http |
| 88.80.150.82:8080 | ✓ 1164ms | ✓ 1952ms | 否 | 否 | ✓ 1859ms | https |
| 86.53.183.16:1080 | ✓ 1256ms | ✓ 1873ms | ✓ 1500ms | 否 | 否 | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 781ms | ✓ 1881ms | ✓ 1322ms | http |
| 209.126.84.232:8888 | ✓ 1529ms | 否 | 否 | ✓ 1914ms | ✓ 1824ms | http |
| 103.39.51.190:8080 | ✓ 1379ms | 否 | ✓ 1669ms | 否 | ✓ 1576ms | http |
| 45.129.141.143:3128 | ✓ 1400ms | ✓ 1946ms | 否 | 否 | ✓ 1659ms | http |
| 45.140.147.155:1082 | ✓ 1530ms | ✓ 1710ms | 否 | 否 | ✓ 1847ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1687ms | ✓ 1367ms | ✓ 1586ms | ✓ 1182ms | http |
| 194.59.204.87:9080 | ✓ 1304ms | ✓ 1910ms | ✓ 1718ms | 否 | 否 | http |
| 38.145.208.172:8448 | 否 | ✓ 1067ms | ✓ 1071ms | 否 | ✓ 1067ms | http |
| 38.145.220.41:8444 | ✓ 795ms | ✓ 789ms | ✓ 467ms | ✓ 1347ms | ✓ 1266ms | http |
| 38.145.220.32:8450 | ✓ 1270ms | ✓ 693ms | ✓ 470ms | ✓ 921ms | ✓ 997ms | http |
| 38.34.179.66:8446 | ✓ 1255ms | ✓ 612ms | ✓ 140ms | ✓ 780ms | ✓ 1496ms | http |
| 38.34.179.24:8453 | ✓ 1273ms | ✓ 649ms | ✓ 131ms | ✓ 704ms | ✓ 1693ms | http |
| 38.34.179.62:8453 | ✓ 1255ms | ✓ 670ms | ✓ 107ms | ✓ 755ms | ✓ 1894ms | http |
| 183.249.5.110:22222 | ✓ 722ms | ✓ 853ms | ✓ 702ms | ✓ 881ms | ✓ 715ms | http |
| 38.34.179.29:8449 | ✓ 1505ms | ✓ 644ms | ✓ 180ms | ✓ 802ms | 否 | http |
| 38.34.179.100:8452 | ✓ 1725ms | ✓ 608ms | ✓ 280ms | ✓ 1061ms | 否 | http |
| 38.145.208.169:8446 | ✓ 788ms | ✓ 1177ms | ✓ 723ms | ✓ 849ms | ✓ 657ms | http |
| 38.145.208.170:8451 | ✓ 791ms | 否 | ✓ 78ms | ✓ 648ms | ✓ 879ms | http |
| 38.145.218.227:8451 | ✓ 790ms | 否 | ✓ 90ms | ✓ 654ms | ✓ 872ms | http |
| 38.145.208.242:8444 | ✓ 1039ms | 否 | ✓ 448ms | ✓ 697ms | ✓ 956ms | http |
| 183.249.5.117:22222 | ✓ 815ms | ✓ 829ms | ✓ 642ms | ✓ 860ms | ✓ 786ms | http |
| 222.184.48.242:22222 | ✓ 835ms | ✓ 1172ms | ✓ 1121ms | ✓ 1491ms | ✓ 955ms | http |
| 222.184.48.251:22222 | 否 | ✓ 1617ms | ✓ 1395ms | ✓ 1711ms | ✓ 1409ms | http |
| 45.140.147.155:1081 | ✓ 1393ms | 否 | ✓ 1178ms | 否 | ✓ 1028ms | http |
| 183.249.5.105:22222 | ✓ 649ms | ✓ 1156ms | ✓ 634ms | ✓ 910ms | ✓ 664ms | http |
| 111.79.111.126:3128 | ✓ 1842ms | 否 | ✓ 1920ms | 否 | ✓ 1410ms | http |
| 193.233.22.29:10808 | ✓ 916ms | 否 | ✓ 1785ms | 否 | ✓ 1084ms | http |
| 45.136.198.40:3128 | ✓ 1326ms | ✓ 1870ms | 否 | 否 | ✓ 1827ms | http |
| 148.153.56.51:80 | ✓ 156ms | ✓ 642ms | ✓ 779ms | ✓ 859ms | ✓ 962ms | http |
| 59.46.216.131:30001 | ✓ 917ms | ✓ 1283ms | ✓ 1038ms | ✓ 1298ms | 否 | http |
| 180.190.187.41:5050 | ✓ 1811ms | 否 | 否 | ✓ 1424ms | ✓ 1588ms | http |
| 103.82.23.118:5171 | ✓ 1587ms | 否 | ✓ 1884ms | ✓ 1553ms | ✓ 1769ms | http |
| 38.34.179.94:8453 | ✓ 747ms | ✓ 659ms | ✓ 1068ms | 否 | ✓ 599ms | http |
| 103.18.78.250:1111 | ✓ 1705ms | 否 | ✓ 1405ms | ✓ 1370ms | ✓ 1245ms | http |
| 116.80.96.100:3172 | ✓ 1881ms | 否 | 否 | ✓ 1778ms | ✓ 1973ms | http |
| 160.238.65.7:3128 | ✓ 1406ms | 否 | 否 | ✓ 1647ms | ✓ 1296ms | http |
| 64.227.76.27:1080 | ✓ 628ms | ✓ 1627ms | 否 | ✓ 1431ms | ✓ 1615ms | http |

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
