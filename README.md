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

最后更新：2026-04-04 06:56:32 UTC（2026-04-04 14:56:32 UTC+8）

**代理总数：81**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 963ms | 否 | ✓ 1013ms | ✓ 984ms | ✓ 991ms | http |
| 95.213.217.168:52004 | ✓ 1150ms | 否 | ✓ 1496ms | ✓ 1916ms | ✓ 1176ms | http |
| 159.223.71.162:8080 | ✓ 917ms | 否 | ✓ 1661ms | ✓ 1698ms | ✓ 1015ms | http |
| 178.156.224.42:3128 | ✓ 1409ms | 否 | ✓ 1839ms | 否 | ✓ 1777ms | http |
| 113.160.132.26:8080 | ✓ 1990ms | 否 | ✓ 1099ms | ✓ 1422ms | ✓ 1099ms | http |
| 167.103.115.102:8800 | ✓ 1149ms | 否 | ✓ 1185ms | 否 | ✓ 1667ms | http |
| 111.227.254.12:22222 | 否 | ✓ 1541ms | ✓ 1935ms | ✓ 1557ms | ✓ 1279ms | http |
| 111.227.254.9:22222 | ✓ 1999ms | ✓ 1577ms | 否 | ✓ 1967ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1296ms | ✓ 1776ms | 否 | ✓ 1543ms | ✓ 1263ms | http |
| 167.103.34.108:8800 | ✓ 1228ms | 否 | ✓ 1206ms | ✓ 1356ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1857ms | 否 | ✓ 1245ms | 否 | ✓ 1904ms | http |
| 167.103.144.127:8800 | ✓ 1824ms | 否 | ✓ 1673ms | ✓ 1908ms | ✓ 1792ms | http |
| 147.161.239.240:8800 | ✓ 970ms | 否 | ✓ 1373ms | ✓ 1371ms | ✓ 1350ms | http |
| 185.118.51.163:3128 | ✓ 1420ms | 否 | ✓ 1350ms | 否 | ✓ 1986ms | http |
| 45.167.125.21:999 | ✓ 915ms | 否 | ✓ 1439ms | ✓ 1721ms | ✓ 1591ms | http |
| 35.225.22.61:80 | 否 | ✓ 1803ms | 否 | ✓ 1038ms | ✓ 916ms | http |
| 177.234.217.88:999 | ✓ 1278ms | 否 | ✓ 874ms | ✓ 1769ms | ✓ 1479ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1482ms | 否 | ✓ 1697ms | ✓ 1183ms | http |
| 1.231.81.166:3128 | ✓ 1047ms | ✓ 1707ms | 否 | ✓ 1371ms | ✓ 1014ms | http |
| 38.145.208.209:8447 | ✓ 1742ms | 否 | ✓ 587ms | ✓ 948ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1156ms | 否 | ✓ 1469ms | ✓ 1488ms | ✓ 1234ms | http |
| 218.108.131.186:17890 | ✓ 899ms | ✓ 957ms | ✓ 872ms | ✓ 1061ms | ✓ 845ms | http |
| 47.110.42.192:9003 | ✓ 1938ms | 否 | ✓ 1798ms | 否 | ✓ 1578ms | http |
| 101.43.127.100:8877 | ✓ 1083ms | ✓ 1223ms | ✓ 1838ms | ✓ 1333ms | 否 | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 839ms | ✓ 1515ms | ✓ 1559ms | http |
| 64.227.76.27:1080 | ✓ 1118ms | 否 | ✓ 1759ms | 否 | ✓ 1434ms | http |
| 162.243.149.86:31028 | 否 | ✓ 980ms | ✓ 1572ms | ✓ 1985ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1332ms | 否 | ✓ 1324ms | ✓ 1769ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1643ms | 否 | ✓ 1460ms | ✓ 1452ms | ✓ 1347ms | http |
| 45.136.198.40:3128 | ✓ 1002ms | ✓ 1425ms | ✓ 1761ms | ✓ 1971ms | ✓ 1701ms | http |
| 43.167.237.94:3128 | ✓ 625ms | ✓ 1649ms | ✓ 620ms | ✓ 1001ms | ✓ 1403ms | http |
| 45.136.130.176:8451 | 否 | ✓ 834ms | ✓ 768ms | 否 | ✓ 1197ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 700ms | ✓ 1251ms | ✓ 907ms | http |
| 38.145.208.241:8447 | 否 | 否 | ✓ 430ms | ✓ 933ms | ✓ 1073ms | http |
| 38.145.208.175:8447 | ✓ 268ms | ✓ 1215ms | 否 | ✓ 1242ms | ✓ 1855ms | http |
| 38.145.208.243:8447 | 否 | 否 | ✓ 410ms | ✓ 1011ms | ✓ 1332ms | http |
| 38.145.208.203:8449 | 否 | ✓ 1137ms | ✓ 1579ms | ✓ 1150ms | ✓ 949ms | http |
| 133.242.138.34:8100 | ✓ 1030ms | ✓ 1673ms | ✓ 1615ms | ✓ 1243ms | ✓ 1120ms | http |
| 120.92.212.16:7890 | ✓ 1314ms | ✓ 1465ms | 否 | ✓ 1534ms | ✓ 1155ms | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 1846ms | ✓ 1113ms | ✓ 835ms | http |
| 3.8.4.205:443 | ✓ 1326ms | 否 | ✓ 1780ms | ✓ 1972ms | ✓ 1688ms | http |
| 159.223.71.162:443 | ✓ 1483ms | 否 | ✓ 1367ms | ✓ 1274ms | ✓ 1378ms | http |
| 5.104.87.17:8050 | 否 | 否 | ✓ 1456ms | ✓ 1304ms | ✓ 1499ms | http |
| 181.78.44.63:999 | ✓ 1279ms | 否 | 否 | ✓ 1977ms | ✓ 1389ms | http |
| 148.135.116.20:8118 | 否 | ✓ 1100ms | ✓ 1227ms | ✓ 969ms | ✓ 998ms | http |
| 59.8.203.55:80 | ✓ 935ms | ✓ 1609ms | ✓ 1930ms | ✓ 1201ms | ✓ 946ms | http |
| 38.145.208.207:8453 | ✓ 932ms | ✓ 1360ms | 否 | ✓ 1471ms | 否 | http |
| 38.34.179.24:8447 | 否 | ✓ 1039ms | ✓ 1972ms | ✓ 1748ms | ✓ 1420ms | http |
| 114.237.77.239:1080 | ✓ 1073ms | ✓ 1441ms | ✓ 1399ms | ✓ 1389ms | 否 | http |
| 57.128.188.167:9163 | ✓ 1510ms | 否 | ✓ 1732ms | ✓ 1879ms | ✓ 1598ms | http |
| 38.34.179.105:8449 | ✓ 1901ms | 否 | ✓ 1520ms | ✓ 1956ms | ✓ 1894ms | http |
| 101.132.61.121:8888 | 否 | ✓ 1457ms | 否 | ✓ 1788ms | ✓ 1477ms | http |
| 47.121.114.42:3129 | 否 | ✓ 1933ms | 否 | ✓ 1869ms | ✓ 1829ms | http |
| 203.150.128.106:8080 | ✓ 1738ms | 否 | ✓ 1783ms | ✓ 1709ms | 否 | http |
| 62.113.119.14:8080 | ✓ 541ms | 否 | ✓ 1004ms | ✓ 1667ms | 否 | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1341ms | ✓ 1479ms | ✓ 1757ms | http |
| 37.187.92.9:1029 | ✓ 1873ms | 否 | ✓ 1235ms | 否 | ✓ 1819ms | http |
| 45.125.67.37:443 | ✓ 1067ms | 否 | ✓ 1073ms | ✓ 1459ms | ✓ 1322ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1689ms | ✓ 1433ms | ✓ 1144ms | http |
| 106.117.208.101:7890 | ✓ 1507ms | 否 | ✓ 1343ms | ✓ 1466ms | 否 | http |
| 82.114.228.67:1080 | ✓ 1666ms | 否 | ✓ 999ms | 否 | ✓ 1159ms | http |
| 24.144.86.173:1080 | ✓ 634ms | ✓ 1832ms | ✓ 1845ms | 否 | 否 | http |
| 150.241.71.15:1080 | 否 | 否 | ✓ 1495ms | ✓ 1763ms | ✓ 1334ms | http |
| 154.16.41.18:5555 | ✓ 1002ms | ✓ 1208ms | ✓ 1212ms | ✓ 1501ms | ✓ 771ms | http |
| 103.191.92.157:1009 | ✓ 1857ms | 否 | 否 | ✓ 1435ms | ✓ 1171ms | http |
| 45.126.250.38:3125 | ✓ 1857ms | 否 | ✓ 1900ms | ✓ 1660ms | ✓ 1670ms | http |
| 38.34.183.130:8452 | ✓ 1166ms | ✓ 1543ms | ✓ 946ms | ✓ 976ms | ✓ 755ms | http |
| 38.145.208.215:8444 | ✓ 871ms | 否 | ✓ 1572ms | ✓ 996ms | 否 | http |
| 65.21.90.27:1019 | ✓ 1082ms | 否 | ✓ 1230ms | 否 | ✓ 1838ms | http |
| 217.217.249.160:8080 | ✓ 1183ms | 否 | ✓ 1083ms | 否 | ✓ 1793ms | http |
| 38.34.179.103:8448 | ✓ 889ms | ✓ 1880ms | ✓ 1569ms | 否 | 否 | http |
| 38.145.208.244:8446 | 否 | ✓ 1539ms | ✓ 539ms | ✓ 941ms | ✓ 1077ms | http |
| 38.145.220.175:8449 | 否 | ✓ 1791ms | ✓ 1372ms | ✓ 965ms | 否 | http |
| 209.38.154.7:1080 | ✓ 1543ms | ✓ 1388ms | 否 | ✓ 1085ms | ✓ 698ms | http |
| 213.165.186.45:18998 | ✓ 1058ms | ✓ 1711ms | 否 | 否 | ✓ 1872ms | http |
| 158.101.113.18:80 | ✓ 223ms | ✓ 849ms | ✓ 1555ms | ✓ 910ms | ✓ 670ms | http |
| 183.232.248.73:7890 | ✓ 1284ms | ✓ 1542ms | 否 | ✓ 1297ms | 否 | http |
| 187.250.33.95:3128 | ✓ 612ms | 否 | 否 | ✓ 1963ms | ✓ 1100ms | http |
| 38.34.179.186:8444 | ✓ 512ms | 否 | ✓ 1033ms | ✓ 965ms | 否 | http |
| 38.34.179.177:8453 | ✓ 762ms | ✓ 983ms | ✓ 1292ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1336ms | ✓ 1623ms | 否 | 否 | ✓ 1525ms | http |

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
