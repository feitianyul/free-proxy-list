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

最后更新：2026-05-07 17:45:49 UTC（2026-05-08 01:45:49 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1547ms | ✓ 1461ms | 否 | ✓ 1217ms | ✓ 1145ms | http |
| 1.231.81.166:3128 | ✓ 1772ms | ✓ 1297ms | 否 | ✓ 1094ms | ✓ 1232ms | http |
| 45.125.67.37:8443 | ✓ 909ms | 否 | ✓ 924ms | ✓ 1083ms | ✓ 1184ms | http |
| 113.176.92.71:3128 | ✓ 1052ms | ✓ 1387ms | ✓ 1619ms | ✓ 1340ms | ✓ 1130ms | http |
| 115.231.181.40:8128 | ✓ 1006ms | 否 | ✓ 972ms | ✓ 1225ms | ✓ 1051ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1181ms | 否 | ✓ 1208ms | ✓ 1102ms | http |
| 103.147.152.12:1080 | ✓ 1068ms | ✓ 1776ms | 否 | 否 | ✓ 1369ms | http |
| 77.110.107.80:1080 | ✓ 1481ms | 否 | ✓ 726ms | ✓ 1918ms | 否 | http |
| 138.197.68.35:4857 | ✓ 488ms | ✓ 1640ms | 否 | 否 | ✓ 1030ms | http |
| 8.219.97.248:80 | ✓ 1545ms | 否 | ✓ 1584ms | ✓ 1874ms | 否 | http |
| 43.133.44.89:8888 | ✓ 1870ms | 否 | ✓ 1129ms | 否 | ✓ 1781ms | http |
| 91.242.229.129:8092 | 否 | ✓ 1487ms | 否 | ✓ 1797ms | ✓ 1144ms | http |
| 45.59.122.132:80 | ✓ 1169ms | 否 | ✓ 951ms | 否 | ✓ 1087ms | http |
| 212.58.132.5:8888 | ✓ 1244ms | 否 | ✓ 1583ms | ✓ 1497ms | ✓ 1183ms | http |
| 120.92.212.16:8890 | ✓ 1193ms | 否 | ✓ 1579ms | 否 | ✓ 1029ms | http |
| 158.160.215.167:8126 | ✓ 706ms | 否 | ✓ 707ms | 否 | ✓ 1899ms | http |
| 120.92.108.86:7890 | ✓ 1964ms | 否 | ✓ 1850ms | ✓ 1753ms | ✓ 1513ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1434ms | 否 | ✓ 1496ms | ✓ 1084ms | http |
| 43.156.132.113:3128 | ✓ 1342ms | 否 | ✓ 797ms | ✓ 1067ms | ✓ 875ms | http |
| 38.180.192.119:3128 | ✓ 917ms | 否 | ✓ 1869ms | ✓ 1234ms | ✓ 1399ms | http |
| 45.63.88.46:1080 | ✓ 1515ms | ✓ 1696ms | 否 | 否 | ✓ 1905ms | http |
| 86.104.74.110:1081 | ✓ 661ms | ✓ 1498ms | ✓ 933ms | ✓ 1775ms | ✓ 1178ms | http |
| 121.130.199.80:24007 | ✓ 863ms | ✓ 1558ms | ✓ 1191ms | ✓ 1268ms | ✓ 1266ms | http |
| 77.110.119.136:3128 | ✓ 748ms | 否 | ✓ 386ms | 否 | ✓ 1135ms | http |
| 45.153.231.229:8080 | ✓ 962ms | ✓ 1982ms | ✓ 1487ms | 否 | 否 | http |
| 129.213.162.27:17777 | ✓ 1467ms | ✓ 1350ms | 否 | ✓ 1623ms | ✓ 1187ms | http |
| 47.112.25.109:7890 | ✓ 1149ms | ✓ 1179ms | ✓ 971ms | ✓ 1091ms | ✓ 889ms | http |
| 86.104.74.110:1082 | ✓ 1315ms | 否 | ✓ 957ms | ✓ 1943ms | ✓ 1000ms | http |
| 223.16.170.103:80 | ✓ 1118ms | 否 | ✓ 898ms | 否 | ✓ 1145ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1245ms | ✓ 1330ms | ✓ 1258ms | http |
| 193.160.209.58:1080 | ✓ 1222ms | 否 | ✓ 1186ms | ✓ 1845ms | ✓ 1776ms | http |
| 86.104.72.219:1081 | ✓ 906ms | ✓ 1423ms | ✓ 1146ms | ✓ 1545ms | 否 | http |
| 47.79.39.142:30000 | 否 | 否 | ✓ 1202ms | ✓ 1035ms | ✓ 785ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1037ms | ✓ 1285ms | ✓ 1833ms | 否 | http |
| 195.19.217.200:3128 | ✓ 1561ms | 否 | ✓ 1670ms | 否 | ✓ 1627ms | http |
| 37.187.109.70:10111 | ✓ 1518ms | 否 | ✓ 1429ms | 否 | ✓ 1962ms | http |
| 158.160.215.167:8124 | ✓ 1698ms | 否 | ✓ 1718ms | 否 | ✓ 1873ms | http |
| 181.119.97.24:999 | ✓ 1767ms | ✓ 1964ms | ✓ 1820ms | 否 | 否 | http |
| 185.191.236.162:3128 | ✓ 1707ms | ✓ 1743ms | ✓ 1706ms | 否 | ✓ 1984ms | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 1981ms | ✓ 1902ms | ✓ 1675ms | http |
| 101.32.243.189:80 | ✓ 1060ms | ✓ 1597ms | ✓ 1577ms | ✓ 1383ms | ✓ 1308ms | http |
| 43.99.54.236:5555 | ✓ 717ms | ✓ 956ms | ✓ 687ms | ✓ 872ms | ✓ 681ms | http |
| 8.154.21.175:3128 | ✓ 854ms | ✓ 1133ms | ✓ 857ms | ✓ 1096ms | ✓ 922ms | http |
| 77.93.89.128:47146 | ✓ 1340ms | 否 | ✓ 1276ms | ✓ 1116ms | ✓ 913ms | http |
| 62.113.119.14:8080 | ✓ 1469ms | ✓ 1521ms | ✓ 1572ms | ✓ 1645ms | ✓ 1173ms | http |
| 121.230.9.217:1080 | ✓ 1247ms | ✓ 1540ms | ✓ 1022ms | ✓ 1686ms | ✓ 1092ms | http |
| 118.113.244.152:1080 | ✓ 1108ms | ✓ 1596ms | ✓ 1803ms | ✓ 1564ms | ✓ 1998ms | http |
| 121.230.9.19:1080 | 否 | ✓ 1529ms | ✓ 1444ms | ✓ 1681ms | ✓ 1123ms | http |
| 220.197.44.36:3128 | 否 | 否 | ✓ 1747ms | ✓ 1917ms | ✓ 1507ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1627ms | 否 | ✓ 1549ms | ✓ 1304ms | http |
| 220.121.143.33:3128 | 否 | ✓ 1159ms | ✓ 888ms | ✓ 993ms | 否 | http |
| 20.120.225.109:3128 | ✓ 662ms | ✓ 1257ms | ✓ 1660ms | ✓ 1357ms | ✓ 760ms | http |
| 207.254.71.62:8088 | ✓ 1216ms | ✓ 1846ms | 否 | 否 | ✓ 1710ms | http |
| 218.108.131.186:17890 | ✓ 866ms | ✓ 1076ms | ✓ 846ms | ✓ 1443ms | ✓ 930ms | http |
| 152.70.91.193:40000 | ✓ 1267ms | 否 | 否 | ✓ 1931ms | ✓ 1333ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1718ms | ✓ 1042ms | ✓ 1321ms | ✓ 960ms | http |
| 64.188.77.26:3128 | ✓ 1107ms | 否 | 否 | ✓ 1632ms | ✓ 1406ms | http |
| 3.101.133.120:80 | ✓ 1849ms | ✓ 1157ms | ✓ 271ms | ✓ 778ms | ✓ 596ms | http |
| 157.230.220.25:4857 | ✓ 316ms | 否 | 否 | ✓ 1398ms | ✓ 869ms | http |
| 84.47.150.125:1080 | ✓ 822ms | 否 | ✓ 1937ms | 否 | ✓ 1661ms | http |
| 20.164.75.153:8080 | ✓ 1821ms | 否 | ✓ 1444ms | 否 | ✓ 1999ms | http |
| 116.171.106.111:3443 | 否 | 否 | ✓ 1605ms | ✓ 1808ms | ✓ 1439ms | http |
| 94.131.118.129:1081 | ✓ 1353ms | 否 | ✓ 819ms | ✓ 1790ms | ✓ 1241ms | http |
| 148.230.4.241:999 | ✓ 1146ms | ✓ 1717ms | ✓ 576ms | 否 | 否 | http |
| 86.104.72.219:1082 | ✓ 509ms | ✓ 1247ms | ✓ 853ms | ✓ 1477ms | ✓ 892ms | http |
| 128.199.121.61:9090 | ✓ 1588ms | 否 | ✓ 1880ms | ✓ 1914ms | ✓ 1105ms | http |
| 38.194.254.134:999 | ✓ 1005ms | ✓ 1559ms | ✓ 1347ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 1190ms | 否 | ✓ 875ms | 否 | ✓ 1613ms | http |
| 116.80.49.134:3172 | ✓ 1599ms | 否 | ✓ 1587ms | ✓ 1845ms | 否 | http |
| 94.131.118.129:1082 | ✓ 676ms | 否 | ✓ 742ms | ✓ 1614ms | ✓ 1033ms | http |
| 129.226.81.110:7890 | 否 | 否 | ✓ 804ms | ✓ 1137ms | ✓ 1101ms | http |
| 94.131.118.39:1081 | ✓ 1586ms | 否 | ✓ 1945ms | 否 | ✓ 1734ms | http |
| 103.67.46.225:3125 | ✓ 1780ms | 否 | 否 | ✓ 1654ms | ✓ 1964ms | http |
| 61.52.131.172:8443 | ✓ 917ms | ✓ 1213ms | ✓ 955ms | ✓ 1235ms | ✓ 940ms | http |
| 45.236.129.64:3128 | ✓ 1464ms | 否 | 否 | ✓ 1982ms | ✓ 1534ms | http |
| 8.140.104.98:3128 | ✓ 917ms | ✓ 1731ms | ✓ 1001ms | ✓ 1373ms | ✓ 1071ms | http |
| 80.92.204.47:1081 | ✓ 579ms | ✓ 1912ms | ✓ 1085ms | 否 | 否 | http |
| 103.172.70.173:8080 | ✓ 1289ms | ✓ 1906ms | ✓ 1546ms | ✓ 1666ms | ✓ 1333ms | http |
| 52.59.51.29:9443 | ✓ 1303ms | 否 | ✓ 1408ms | 否 | ✓ 1810ms | http |
| 3.71.26.7:9050 | 否 | 否 | ✓ 1713ms | ✓ 1915ms | ✓ 1611ms | http |
| 103.157.117.116:8080 | ✓ 1943ms | 否 | 否 | ✓ 1723ms | ✓ 1816ms | http |
| 103.109.96.129:6321 | ✓ 1880ms | 否 | 否 | ✓ 1834ms | ✓ 1950ms | http |
| 118.113.247.175:1080 | ✓ 1656ms | ✓ 1514ms | ✓ 1195ms | ✓ 1535ms | ✓ 1551ms | http |

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
