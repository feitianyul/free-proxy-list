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

最后更新：2026-03-28 11:30:14 UTC（2026-03-28 19:30:14 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1345ms | 否 | ✓ 1103ms | ✓ 1532ms | ✓ 1024ms | http |
| 167.103.115.102:8800 | ✓ 1350ms | 否 | ✓ 998ms | ✓ 1078ms | ✓ 1030ms | http |
| 219.117.204.211:7799 | ✓ 1349ms | 否 | ✓ 1055ms | ✓ 1019ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1401ms | ✓ 1015ms | ✓ 1269ms | ✓ 975ms | http |
| 147.161.239.240:8800 | ✓ 1655ms | ✓ 1850ms | ✓ 1235ms | ✓ 1745ms | ✓ 1445ms | http |
| 167.103.34.108:8800 | ✓ 1632ms | 否 | ✓ 1525ms | ✓ 1581ms | 否 | http |
| 45.167.124.52:8080 | 否 | ✓ 1865ms | ✓ 1407ms | ✓ 1930ms | ✓ 1471ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1169ms | ✓ 1255ms | ✓ 1174ms | http |
| 43.99.54.236:5555 | ✓ 1740ms | ✓ 959ms | ✓ 641ms | ✓ 827ms | ✓ 639ms | http |
| 103.84.95.54:7890 | ✓ 738ms | 否 | ✓ 1187ms | 否 | ✓ 640ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1364ms | ✓ 1653ms | ✓ 1487ms | ✓ 1058ms | http |
| 1.231.81.166:3128 | ✓ 1192ms | ✓ 1401ms | ✓ 1047ms | ✓ 1300ms | ✓ 920ms | http |
| 167.103.144.127:8800 | ✓ 1356ms | 否 | 否 | ✓ 1497ms | ✓ 1388ms | http |
| 62.113.119.14:8080 | ✓ 1784ms | 否 | ✓ 1322ms | ✓ 1596ms | ✓ 1296ms | http |
| 180.250.219.58:53281 | ✓ 1717ms | 否 | ✓ 1660ms | ✓ 1880ms | ✓ 1803ms | http |
| 167.103.31.122:8800 | ✓ 1701ms | 否 | ✓ 1350ms | ✓ 1652ms | 否 | http |
| 45.12.151.226:2829 | ✓ 1281ms | 否 | ✓ 1227ms | ✓ 1920ms | ✓ 1504ms | http |
| 106.75.15.167:7890 | 否 | ✓ 1452ms | ✓ 1164ms | 否 | ✓ 1234ms | http |
| 208.87.243.199:7878 | ✓ 428ms | ✓ 695ms | ✓ 955ms | ✓ 665ms | ✓ 610ms | http |
| 160.250.4.245:1 | ✓ 1458ms | 否 | 否 | ✓ 1278ms | ✓ 1125ms | http |
| 128.199.114.189:9090 | ✓ 1401ms | 否 | ✓ 1143ms | ✓ 1425ms | 否 | http |
| 101.43.127.100:8877 | ✓ 942ms | ✓ 1054ms | ✓ 1506ms | ✓ 1118ms | ✓ 1204ms | http |
| 101.47.73.135:3128 | ✓ 1477ms | 否 | 否 | ✓ 1556ms | ✓ 1580ms | http |
| 45.88.0.113:3128 | ✓ 1057ms | ✓ 1576ms | ✓ 1948ms | 否 | 否 | http |
| 45.88.0.117:3128 | ✓ 1059ms | ✓ 1655ms | ✓ 1868ms | 否 | 否 | http |
| 45.88.0.116:3128 | ✓ 1060ms | ✓ 1651ms | ✓ 1865ms | 否 | ✓ 1848ms | http |
| 5.104.87.17:8051 | ✓ 776ms | 否 | ✓ 960ms | ✓ 935ms | ✓ 859ms | http |
| 160.238.65.6:3128 | ✓ 650ms | ✓ 1551ms | ✓ 1060ms | ✓ 1540ms | ✓ 1649ms | http |
| 213.220.62.62:3128 | 否 | ✓ 1708ms | ✓ 1101ms | 否 | ✓ 1258ms | http |
| 45.88.0.115:3128 | ✓ 1834ms | ✓ 1868ms | ✓ 1108ms | 否 | ✓ 1675ms | http |
| 160.238.65.2:3128 | ✓ 1833ms | 否 | ✓ 974ms | ✓ 1978ms | ✓ 1654ms | http |
| 160.238.65.3:3128 | ✓ 649ms | ✓ 1740ms | ✓ 594ms | 否 | ✓ 1187ms | http |
| 45.88.0.114:3128 | 否 | ✓ 1777ms | ✓ 1123ms | ✓ 1559ms | ✓ 1210ms | http |
| 45.88.0.99:3128 | 否 | ✓ 1743ms | ✓ 1065ms | 否 | ✓ 1647ms | http |
| 45.88.0.98:3128 | 否 | 否 | ✓ 1612ms | ✓ 1585ms | ✓ 1232ms | http |
| 160.238.65.7:3128 | ✓ 1833ms | 否 | ✓ 1464ms | ✓ 1483ms | ✓ 1642ms | http |
| 160.238.65.8:3128 | ✓ 1831ms | 否 | ✓ 1453ms | ✓ 1519ms | ✓ 1622ms | http |
| 45.88.0.111:3128 | 否 | ✓ 1718ms | ✓ 1758ms | ✓ 1613ms | ✓ 1250ms | http |
| 160.238.65.4:3128 | ✓ 649ms | 否 | ✓ 620ms | ✓ 1499ms | ✓ 1221ms | http |
| 160.238.65.5:3128 | ✓ 1834ms | 否 | ✓ 1162ms | ✓ 1523ms | ✓ 1671ms | http |
| 116.80.49.165:3172 | ✓ 1776ms | 否 | ✓ 1657ms | 否 | ✓ 1636ms | http |
| 38.145.218.51:8450 | ✓ 1415ms | ✓ 651ms | ✓ 363ms | ✓ 772ms | ✓ 1135ms | http |
| 45.136.130.251:8449 | ✓ 1355ms | ✓ 787ms | ✓ 465ms | ✓ 1490ms | ✓ 1553ms | http |
| 168.110.52.228:3128 | ✓ 698ms | ✓ 1662ms | ✓ 867ms | ✓ 942ms | ✓ 1052ms | http |
| 38.145.208.190:8444 | ✓ 613ms | 否 | ✓ 156ms | ✓ 673ms | ✓ 774ms | http |
| 128.199.254.13:9090 | ✓ 1573ms | 否 | 否 | ✓ 1911ms | ✓ 1106ms | http |
| 38.145.203.98:8449 | ✓ 477ms | ✓ 1026ms | 否 | ✓ 840ms | ✓ 810ms | http |
| 128.199.121.61:9090 | ✓ 1515ms | 否 | ✓ 1206ms | ✓ 1266ms | ✓ 835ms | http |
| 195.123.209.48:3128 | ✓ 805ms | 否 | ✓ 1437ms | 否 | ✓ 1838ms | http |
| 185.76.240.34:10001 | ✓ 1076ms | 否 | 否 | ✓ 1969ms | ✓ 1696ms | http |
| 185.76.241.124:10001 | ✓ 1063ms | 否 | ✓ 906ms | 否 | ✓ 1872ms | http |
| 45.144.232.5:11741 | ✓ 974ms | 否 | ✓ 1503ms | 否 | ✓ 1692ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 847ms | ✓ 1394ms | ✓ 1400ms | http |
| 59.46.216.131:30001 | ✓ 1885ms | ✓ 1341ms | 否 | ✓ 1375ms | ✓ 1046ms | http |
| 45.149.92.147:5001 | ✓ 1999ms | 否 | ✓ 1845ms | ✓ 1166ms | ✓ 1197ms | http |
| 113.176.92.71:3128 | ✓ 943ms | ✓ 1365ms | ✓ 1288ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1716ms | 否 | ✓ 1612ms | ✓ 1991ms | 否 | http |
| 168.63.153.150:3128 | ✓ 642ms | ✓ 1171ms | ✓ 828ms | ✓ 777ms | ✓ 641ms | http |
| 166.88.55.83:7890 | ✓ 631ms | ✓ 1069ms | ✓ 627ms | ✓ 799ms | ✓ 1675ms | http |
| 95.213.217.168:52004 | 否 | ✓ 1733ms | ✓ 834ms | ✓ 1705ms | ✓ 1257ms | http |
| 128.199.113.85:9090 | ✓ 819ms | 否 | ✓ 863ms | ✓ 1216ms | ✓ 1028ms | http |
| 165.232.146.249:3128 | 否 | ✓ 1724ms | ✓ 725ms | ✓ 722ms | ✓ 768ms | http |
| 177.234.217.88:999 | ✓ 995ms | ✓ 1754ms | ✓ 975ms | ✓ 1795ms | ✓ 1510ms | http |
| 38.145.208.223:8444 | ✓ 874ms | ✓ 988ms | 否 | ✓ 1176ms | ✓ 620ms | http |
| 45.136.131.32:8445 | ✓ 1069ms | 否 | ✓ 1254ms | 否 | ✓ 1969ms | http |
| 193.233.22.29:10808 | 否 | 否 | ✓ 1058ms | ✓ 1488ms | ✓ 1350ms | http |
| 120.92.212.16:7890 | ✓ 952ms | 否 | 否 | ✓ 1272ms | ✓ 988ms | http |
| 120.92.212.16:8890 | ✓ 1004ms | ✓ 1233ms | ✓ 1008ms | ✓ 1251ms | 否 | http |
| 106.117.208.101:7890 | ✓ 982ms | ✓ 1199ms | ✓ 1056ms | ✓ 1255ms | ✓ 1003ms | http |
| 45.136.198.40:3128 | 否 | ✓ 1869ms | ✓ 1686ms | 否 | ✓ 1969ms | http |
| 45.144.28.81:10808 | 否 | 否 | ✓ 1313ms | ✓ 1515ms | ✓ 1386ms | http |
| 8.219.97.248:80 | ✓ 1925ms | 否 | ✓ 1127ms | ✓ 1234ms | 否 | http |
| 194.67.99.223:1080 | ✓ 1740ms | 否 | ✓ 1886ms | ✓ 1858ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1492ms | ✓ 1727ms | 否 | ✓ 1135ms | ✓ 1499ms | http |
| 146.190.80.158:9090 | ✓ 1662ms | 否 | ✓ 827ms | 否 | ✓ 1764ms | http |
| 160.238.65.9:3128 | ✓ 1657ms | ✓ 1534ms | ✓ 1124ms | 否 | 否 | http |
| 88.80.150.82:8080 | ✓ 1088ms | ✓ 1954ms | 否 | 否 | ✓ 1901ms | https |
| 128.199.116.219:9090 | ✓ 784ms | 否 | ✓ 839ms | ✓ 1064ms | 否 | http |
| 64.227.76.27:1080 | ✓ 1801ms | 否 | ✓ 1640ms | ✓ 1964ms | 否 | http |
| 223.16.170.103:80 | ✓ 1104ms | 否 | ✓ 1018ms | 否 | ✓ 1047ms | http |

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
