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

最后更新：2026-03-01 21:23:45 UTC（2026-03-02 05:23:45 UTC+8）

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
| 5.129.206.247:8888 | ✓ 1387ms | ✓ 1702ms | ✓ 1638ms | ✓ 1975ms | 否 | http |
| 194.87.43.49:8888 | ✓ 1388ms | 否 | ✓ 1427ms | 否 | ✓ 1957ms | http |
| 103.84.95.54:7890 | ✓ 851ms | 否 | ✓ 1003ms | 否 | ✓ 1108ms | http |
| 165.225.113.220:10958 | ✓ 1773ms | 否 | 否 | ✓ 1442ms | ✓ 1143ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1220ms | ✓ 1192ms | ✓ 953ms | http |
| 61.72.221.234:3128 | ✓ 1553ms | ✓ 1483ms | ✓ 1484ms | 否 | 否 | http |
| 14.56.107.244:3128 | ✓ 1772ms | ✓ 1462ms | ✓ 1140ms | 否 | ✓ 1869ms | http |
| 217.76.245.80:999 | ✓ 864ms | ✓ 1401ms | ✓ 1062ms | ✓ 1385ms | ✓ 1264ms | http |
| 211.95.152.50:45046 | ✓ 1085ms | ✓ 1389ms | ✓ 1103ms | ✓ 1331ms | ✓ 1085ms | http |
| 115.231.181.40:8128 | ✓ 1063ms | ✓ 1361ms | ✓ 1153ms | 否 | ✓ 1103ms | http |
| 61.72.110.54:3128 | ✓ 1178ms | 否 | ✓ 1332ms | ✓ 1880ms | 否 | http |
| 165.227.5.10:8888 | ✓ 368ms | ✓ 1122ms | 否 | ✓ 1146ms | 否 | http |
| 205.209.118.30:3138 | ✓ 764ms | ✓ 1559ms | ✓ 1701ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 882ms | 否 | ✓ 453ms | ✓ 1156ms | ✓ 869ms | http |
| 120.92.212.16:8890 | ✓ 1980ms | ✓ 1505ms | ✓ 1382ms | 否 | 否 | http |
| 190.9.109.196:999 | ✓ 1243ms | ✓ 1448ms | ✓ 1396ms | ✓ 1428ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1242ms | 否 | 否 | ✓ 1528ms | ✓ 1063ms | http |
| 101.43.255.96:80 | ✓ 1237ms | ✓ 1494ms | ✓ 1140ms | ✓ 1539ms | ✓ 1221ms | http |
| 81.70.169.194:80 | ✓ 1234ms | ✓ 1546ms | ✓ 1172ms | ✓ 1567ms | ✓ 1185ms | http |
| 168.235.110.63:3128 | ✓ 358ms | ✓ 1915ms | ✓ 1738ms | 否 | ✓ 1055ms | http |
| 14.56.177.44:3128 | ✓ 1457ms | 否 | ✓ 1421ms | ✓ 1276ms | ✓ 1001ms | http |
| 120.92.212.16:7890 | ✓ 1122ms | ✓ 1691ms | 否 | ✓ 1499ms | 否 | http |
| 66.228.47.125:110 | ✓ 282ms | 否 | ✓ 1632ms | ✓ 1511ms | ✓ 1095ms | http |
| 150.249.255.91:3128 | ✓ 694ms | ✓ 1055ms | 否 | ✓ 1501ms | 否 | http |
| 47.105.98.23:3128 | ✓ 1333ms | ✓ 1328ms | ✓ 1157ms | ✓ 1491ms | ✓ 1129ms | http |
| 36.147.78.166:80 | ✓ 1904ms | ✓ 1837ms | 否 | 否 | ✓ 1688ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1257ms | ✓ 992ms | ✓ 1868ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1226ms | ✓ 1618ms | ✓ 1444ms | ✓ 1265ms | ✓ 1963ms | http |
| 8.219.97.248:80 | ✓ 1408ms | 否 | ✓ 1580ms | 否 | ✓ 1825ms | http |
| 74.208.234.198:443 | ✓ 600ms | ✓ 1862ms | ✓ 1278ms | ✓ 1542ms | ✓ 836ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1339ms | ✓ 1567ms | ✓ 1290ms | http |
| 37.27.100.102:443 | 否 | ✓ 1638ms | ✓ 1810ms | ✓ 1788ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1387ms | 否 | 否 | ✓ 1826ms | ✓ 1390ms | http |
| 77.83.203.5:443 | ✓ 1249ms | ✓ 1681ms | ✓ 1255ms | ✓ 1974ms | ✓ 1476ms | http |
| 90.84.188.97:8000 | ✓ 381ms | ✓ 1649ms | ✓ 1108ms | 否 | ✓ 1362ms | http |
| 34.101.184.164:3128 | ✓ 1757ms | 否 | ✓ 1554ms | ✓ 1731ms | ✓ 1553ms | http |
| 35.234.17.221:8080 | ✓ 1641ms | ✓ 1378ms | 否 | 否 | ✓ 1041ms | http |
| 144.31.69.170:1080 | ✓ 1013ms | 否 | ✓ 1404ms | 否 | ✓ 1962ms | http |
| 121.40.231.103:7890 | 否 | ✓ 1313ms | ✓ 1024ms | ✓ 1320ms | ✓ 1026ms | http |
| 2.56.178.131:443 | ✓ 1157ms | 否 | ✓ 1883ms | 否 | ✓ 1586ms | http |
| 103.74.192.242:7890 | 否 | 否 | ✓ 1550ms | ✓ 1457ms | ✓ 946ms | http |
| 103.104.99.29:80 | ✓ 1786ms | 否 | ✓ 1913ms | ✓ 1816ms | ✓ 1694ms | http |
| 103.104.99.89:80 | ✓ 1812ms | 否 | ✓ 1909ms | ✓ 1818ms | ✓ 1720ms | http |
| 95.85.252.153:21064 | ✓ 639ms | ✓ 1403ms | ✓ 1122ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 854ms | ✓ 1827ms | ✓ 1649ms | 否 | 否 | http |
| 149.56.24.51:3128 | ✓ 179ms | 否 | ✓ 143ms | ✓ 1037ms | ✓ 761ms | http |
| 167.160.184.231:6005 | ✓ 716ms | 否 | ✓ 795ms | ✓ 1345ms | ✓ 1029ms | http |
| 91.238.104.171:2023 | ✓ 1199ms | ✓ 1593ms | ✓ 898ms | ✓ 1437ms | ✓ 1172ms | http |
| 91.238.104.172:2024 | ✓ 1194ms | 否 | ✓ 673ms | ✓ 1644ms | ✓ 1144ms | http |
| 139.159.99.242:8080 | ✓ 1201ms | 否 | ✓ 1027ms | 否 | ✓ 1533ms | http |
| 123.20.24.166:8118 | 否 | 否 | ✓ 1495ms | ✓ 1871ms | ✓ 1357ms | http |
| 180.127.149.244:1080 | ✓ 1123ms | ✓ 1412ms | ✓ 1494ms | ✓ 1462ms | ✓ 1074ms | http |
| 101.32.244.83:8080 | ✓ 1289ms | 否 | ✓ 1159ms | ✓ 1537ms | ✓ 1458ms | http |
| 121.43.196.210:8222 | ✓ 1153ms | ✓ 1403ms | ✓ 1056ms | ✓ 1429ms | ✓ 1064ms | http |
| 121.43.196.213:8222 | ✓ 1495ms | ✓ 1267ms | ✓ 1007ms | ✓ 1363ms | ✓ 986ms | http |
| 114.55.226.123:10086 | ✓ 1252ms | ✓ 1803ms | ✓ 1184ms | ✓ 1461ms | ✓ 1214ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1236ms | ✓ 1434ms | ✓ 1193ms | http |
| 85.208.108.43:2094 | ✓ 378ms | 否 | ✓ 1309ms | ✓ 1367ms | ✓ 961ms | http |
| 62.113.119.14:8080 | ✓ 1007ms | 否 | ✓ 560ms | ✓ 1864ms | ✓ 1131ms | http |
| 45.140.147.155:1081 | ✓ 1991ms | ✓ 1495ms | ✓ 1497ms | ✓ 1606ms | 否 | http |
| 195.123.209.48:3128 | ✓ 952ms | ✓ 1649ms | ✓ 1601ms | ✓ 1752ms | ✓ 1735ms | http |
| 36.68.150.7:3128 | ✓ 1769ms | 否 | ✓ 1680ms | ✓ 1785ms | 否 | http |
| 61.72.221.194:3128 | ✓ 1084ms | 否 | ✓ 1505ms | 否 | ✓ 1366ms | http |
| 103.82.23.118:5247 | 否 | 否 | ✓ 1988ms | ✓ 1873ms | ✓ 1861ms | http |
| 45.129.141.143:3128 | ✓ 997ms | 否 | ✓ 1737ms | ✓ 1899ms | ✓ 1736ms | http |
| 120.55.163.237:10086 | ✓ 1052ms | 否 | 否 | ✓ 1346ms | ✓ 1120ms | http |
| 36.147.78.166:443 | ✓ 1869ms | ✓ 1892ms | 否 | 否 | ✓ 1873ms | http |
| 24.199.124.152:3128 | 否 | 否 | ✓ 1149ms | ✓ 1244ms | ✓ 1043ms | http |
| 103.39.51.190:8080 | ✓ 1678ms | 否 | ✓ 1908ms | ✓ 1739ms | ✓ 1828ms | http |
| 23.95.191.199:3128 | ✓ 505ms | ✓ 1336ms | ✓ 250ms | ✓ 1159ms | ✓ 967ms | http |
| 45.136.198.40:3128 | ✓ 776ms | ✓ 1562ms | ✓ 1569ms | ✓ 1996ms | ✓ 1536ms | http |
| 103.236.64.247:8888 | ✓ 1426ms | 否 | ✓ 1191ms | ✓ 1399ms | 否 | http |
| 94.177.131.12:3128 | ✓ 646ms | 否 | ✓ 808ms | ✓ 1039ms | ✓ 822ms | http |
| 212.175.29.184:8080 | 否 | 否 | ✓ 1715ms | ✓ 1905ms | ✓ 1462ms | http |
| 85.198.84.77:10808 | ✓ 1106ms | 否 | ✓ 1813ms | 否 | ✓ 1914ms | http |
| 200.125.171.254:999 | ✓ 925ms | ✓ 1500ms | ✓ 1342ms | ✓ 1278ms | ✓ 1252ms | http |
| 194.59.204.87:9080 | ✓ 410ms | ✓ 1358ms | ✓ 923ms | 否 | 否 | http |
| 47.77.180.205:1080 | ✓ 557ms | ✓ 1546ms | ✓ 725ms | ✓ 1034ms | ✓ 836ms | http |

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
