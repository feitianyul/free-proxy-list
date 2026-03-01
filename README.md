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

最后更新：2026-03-01 23:25:08 UTC（2026-03-02 07:25:08 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 103.84.95.54:7890 | ✓ 749ms | 否 | ✓ 759ms | ✓ 943ms | ✓ 881ms | http |
| 205.209.118.30:3138 | ✓ 392ms | ✓ 971ms | ✓ 729ms | ✓ 1297ms | ✓ 1612ms | http |
| 35.225.22.61:80 | ✓ 698ms | 否 | ✓ 876ms | ✓ 967ms | ✓ 1074ms | http |
| 62.113.119.14:8080 | ✓ 598ms | ✓ 1662ms | ✓ 609ms | 否 | ✓ 1083ms | http |
| 91.238.104.171:2023 | ✓ 788ms | ✓ 1807ms | ✓ 1695ms | 否 | ✓ 1980ms | http |
| 1.225.116.115:1080 | ✓ 1945ms | ✓ 1122ms | ✓ 973ms | ✓ 1172ms | ✓ 830ms | http |
| 143.198.37.6:8888 | ✓ 318ms | ✓ 1122ms | ✓ 618ms | 否 | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1495ms | ✓ 985ms | 否 | ✓ 1008ms | http |
| 91.238.104.172:2024 | ✓ 1330ms | ✓ 1857ms | ✓ 1442ms | 否 | 否 | http |
| 5.129.206.247:8888 | ✓ 1399ms | 否 | ✓ 1579ms | 否 | ✓ 1965ms | http |
| 101.43.255.96:80 | ✓ 1073ms | ✓ 1325ms | ✓ 1008ms | ✓ 1296ms | ✓ 1095ms | http |
| 81.70.169.194:80 | ✓ 1121ms | ✓ 1474ms | ✓ 1095ms | ✓ 1310ms | ✓ 1062ms | http |
| 121.40.231.103:7890 | 否 | ✓ 1164ms | ✓ 944ms | ✓ 1222ms | ✓ 1911ms | http |
| 165.227.5.10:8888 | 否 | ✓ 960ms | 否 | ✓ 1365ms | ✓ 973ms | http |
| 190.9.109.202:999 | ✓ 769ms | ✓ 1402ms | ✓ 1180ms | 否 | 否 | http |
| 190.9.109.196:999 | ✓ 774ms | ✓ 1471ms | ✓ 1174ms | 否 | 否 | http |
| 190.9.109.194:999 | ✓ 766ms | ✓ 1504ms | ✓ 1183ms | 否 | 否 | http |
| 150.249.255.91:3128 | 否 | ✓ 1107ms | ✓ 753ms | 否 | ✓ 778ms | http |
| 34.101.184.164:3128 | ✓ 1002ms | 否 | ✓ 1479ms | ✓ 1627ms | ✓ 1250ms | http |
| 77.83.203.5:443 | ✓ 1372ms | ✓ 1676ms | ✓ 1093ms | 否 | 否 | http |
| 36.147.78.166:80 | ✓ 1834ms | ✓ 1788ms | 否 | ✓ 1998ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1055ms | 否 | ✓ 1060ms | 否 | ✓ 1870ms | http |
| 103.215.36.88:18225 | ✓ 1165ms | ✓ 1381ms | ✓ 1132ms | ✓ 1722ms | ✓ 1379ms | http |
| 45.140.147.82:1081 | ✓ 465ms | ✓ 1186ms | ✓ 1106ms | ✓ 1474ms | ✓ 1432ms | http |
| 142.171.85.32:1080 | ✓ 867ms | ✓ 1428ms | ✓ 1032ms | ✓ 1169ms | ✓ 1747ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1085ms | ✓ 1208ms | ✓ 1873ms | http |
| 210.223.44.230:3128 | ✓ 1887ms | 否 | 否 | ✓ 1376ms | ✓ 1782ms | http |
| 61.72.110.54:3128 | ✓ 1022ms | 否 | ✓ 1891ms | 否 | ✓ 1790ms | http |
| 121.128.121.54:3128 | ✓ 1045ms | 否 | ✓ 1499ms | ✓ 1560ms | ✓ 1462ms | http |
| 59.46.216.131:30001 | ✓ 1161ms | ✓ 1576ms | ✓ 1185ms | ✓ 1541ms | 否 | http |
| 162.240.154.26:3128 | ✓ 1862ms | ✓ 1601ms | 否 | ✓ 1607ms | 否 | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1955ms | ✓ 1719ms | ✓ 1574ms | http |
| 138.124.53.25:7443 | ✓ 1521ms | 否 | ✓ 1477ms | ✓ 1622ms | ✓ 1390ms | http |
| 168.235.110.63:3128 | ✓ 645ms | 否 | ✓ 1725ms | 否 | ✓ 1275ms | http |
| 35.234.17.221:8080 | ✓ 1025ms | ✓ 1610ms | ✓ 1292ms | 否 | 否 | http |
| 103.104.99.29:80 | ✓ 1929ms | 否 | ✓ 1615ms | ✓ 1737ms | ✓ 1595ms | http |
| 90.84.188.97:8000 | ✓ 954ms | 否 | 否 | ✓ 1749ms | ✓ 1668ms | http |
| 156.225.70.152:39151 | ✓ 1504ms | ✓ 1256ms | ✓ 710ms | 否 | 否 | http |
| 203.2.151.24:8080 | ✓ 1456ms | 否 | 否 | ✓ 1554ms | ✓ 1491ms | http |
| 37.111.53.31:8080 | 否 | 否 | ✓ 1680ms | ✓ 1830ms | ✓ 1773ms | http |
| 47.77.180.205:1080 | ✓ 749ms | ✓ 958ms | ✓ 594ms | ✓ 963ms | ✓ 749ms | http |
| 167.160.184.231:6005 | ✓ 347ms | ✓ 1909ms | ✓ 901ms | ✓ 1146ms | ✓ 795ms | http |
| 121.230.8.80:1080 | ✓ 1292ms | ✓ 1585ms | ✓ 1193ms | ✓ 1559ms | ✓ 1201ms | http |
| 120.79.99.232:8099 | ✓ 1330ms | ✓ 1651ms | ✓ 1429ms | ✓ 1594ms | ✓ 1329ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1305ms | ✓ 1541ms | ✓ 1640ms | ✓ 1918ms | http |
| 14.56.107.244:3128 | ✓ 1389ms | ✓ 1869ms | 否 | 否 | ✓ 1085ms | http |
| 121.230.9.198:1080 | ✓ 1162ms | ✓ 1602ms | ✓ 1312ms | ✓ 1774ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1397ms | 否 | ✓ 1538ms | ✓ 1154ms | ✓ 1064ms | http |
| 120.92.212.16:7890 | ✓ 1464ms | 否 | 否 | ✓ 1328ms | ✓ 1343ms | http |
| 8.219.97.248:80 | ✓ 1063ms | 否 | 否 | ✓ 1364ms | ✓ 1691ms | http |
| 85.198.84.77:10808 | ✓ 911ms | 否 | ✓ 1782ms | 否 | ✓ 1655ms | http |
| 74.208.234.198:443 | ✓ 1190ms | 否 | ✓ 1346ms | ✓ 1387ms | ✓ 933ms | http |
| 27.254.99.183:8118 | 否 | 否 | ✓ 1391ms | ✓ 1467ms | ✓ 1141ms | http |
| 195.123.209.48:3128 | ✓ 807ms | 否 | ✓ 1388ms | 否 | ✓ 1849ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1692ms | ✓ 1213ms | ✓ 1032ms | http |
| 95.85.252.153:21064 | ✓ 961ms | ✓ 1754ms | ✓ 1362ms | 否 | 否 | http |
| 211.171.114.154:3128 | ✓ 1144ms | ✓ 1571ms | 否 | ✓ 1491ms | ✓ 1306ms | http |
| 103.215.36.88:10101 | ✓ 1133ms | ✓ 1566ms | ✓ 1438ms | 否 | ✓ 1241ms | http |
| 46.249.103.192:443 | ✓ 954ms | 否 | ✓ 1101ms | ✓ 1867ms | 否 | http |
| 172.212.68.37:3128 | 否 | 否 | ✓ 1440ms | ✓ 1434ms | ✓ 773ms | http |
| 103.100.159.145:80 | ✓ 820ms | ✓ 1762ms | ✓ 1098ms | ✓ 1388ms | ✓ 826ms | http |
| 185.213.20.105:3128 | ✓ 761ms | 否 | ✓ 1508ms | ✓ 1930ms | 否 | http |
| 121.230.9.148:1080 | ✓ 1378ms | ✓ 1531ms | ✓ 1103ms | ✓ 1660ms | ✓ 1267ms | http |
| 45.136.198.40:3128 | ✓ 1010ms | ✓ 1841ms | ✓ 1673ms | ✓ 1988ms | ✓ 1549ms | http |
| 45.140.147.155:1081 | ✓ 872ms | ✓ 1179ms | ✓ 1364ms | ✓ 1588ms | ✓ 1171ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1389ms | ✓ 1538ms | ✓ 1479ms | http |
| 45.140.147.155:1082 | ✓ 396ms | ✓ 1149ms | ✓ 1470ms | ✓ 1405ms | ✓ 1239ms | http |
| 2.56.178.131:443 | ✓ 1439ms | 否 | ✓ 1555ms | 否 | ✓ 1867ms | http |
| 37.27.100.108:443 | ✓ 1080ms | 否 | ✓ 1704ms | 否 | ✓ 1801ms | http |
| 121.230.8.49:1080 | ✓ 1817ms | 否 | ✓ 1462ms | 否 | ✓ 1896ms | http |
| 43.161.253.215:1080 | ✓ 1583ms | 否 | ✓ 1762ms | 否 | ✓ 1881ms | http |
| 107.174.133.10:3128 | ✓ 910ms | 否 | ✓ 906ms | 否 | ✓ 993ms | http |
| 101.47.73.135:3128 | ✓ 1948ms | 否 | 否 | ✓ 1918ms | ✓ 1329ms | http |
| 120.55.163.237:10086 | ✓ 1604ms | ✓ 1834ms | ✓ 1406ms | ✓ 1569ms | ✓ 1167ms | http |
| 103.236.64.247:8888 | ✓ 1874ms | 否 | ✓ 1091ms | ✓ 1349ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1133ms | ✓ 1786ms | ✓ 1723ms | 否 | 否 | http |
| 121.230.8.118:1080 | ✓ 1228ms | ✓ 1755ms | ✓ 1662ms | ✓ 1863ms | ✓ 1677ms | http |
| 212.175.29.184:8080 | ✓ 1021ms | ✓ 1550ms | ✓ 1517ms | 否 | ✓ 1458ms | http |

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
