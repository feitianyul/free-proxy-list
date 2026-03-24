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

最后更新：2026-03-24 21:39:37 UTC（2026-03-25 05:39:37 UTC+8）

**代理总数：95**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 94 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 95 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 682ms | ✓ 971ms | ✓ 824ms | ✓ 852ms | ✓ 673ms | http |
| 147.161.210.140:8800 | ✓ 1679ms | ✓ 974ms | ✓ 1022ms | ✓ 1142ms | ✓ 1101ms | http |
| 167.103.115.102:8800 | ✓ 920ms | ✓ 1653ms | ✓ 1087ms | ✓ 1279ms | ✓ 1753ms | http |
| 45.167.124.52:8080 | ✓ 1736ms | 否 | 否 | ✓ 1987ms | ✓ 1724ms | http |
| 142.171.224.229:7890 | ✓ 307ms | ✓ 958ms | ✓ 816ms | ✓ 835ms | ✓ 581ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 687ms | ✓ 834ms | ✓ 671ms | http |
| 35.225.22.61:80 | 否 | ✓ 1811ms | ✓ 1408ms | ✓ 1096ms | ✓ 1037ms | http |
| 167.103.34.108:8800 | ✓ 1160ms | 否 | ✓ 1167ms | ✓ 1358ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1694ms | ✓ 1297ms | ✓ 1203ms | ✓ 1248ms | ✓ 1263ms | http |
| 38.34.179.150:8449 | 否 | ✓ 1818ms | ✓ 1390ms | 否 | ✓ 582ms | http |
| 106.75.15.167:7890 | ✓ 1132ms | 否 | 否 | ✓ 1243ms | ✓ 956ms | http |
| 59.46.216.131:30001 | ✓ 923ms | ✓ 1368ms | ✓ 1109ms | 否 | 否 | http |
| 181.41.201.85:3128 | ✓ 779ms | 否 | ✓ 643ms | 否 | ✓ 1591ms | http |
| 115.231.181.40:8128 | ✓ 1890ms | ✓ 1314ms | ✓ 1690ms | 否 | ✓ 1363ms | http |
| 116.80.49.165:3172 | 否 | 否 | ✓ 1607ms | ✓ 1851ms | ✓ 1660ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1220ms | ✓ 1723ms | 否 | ✓ 989ms | http |
| 167.103.31.122:8800 | ✓ 1655ms | 否 | ✓ 1446ms | ✓ 1729ms | ✓ 1567ms | http |
| 120.92.212.16:8890 | ✓ 1192ms | 否 | ✓ 1230ms | ✓ 1235ms | ✓ 1697ms | http |
| 45.136.130.177:8448 | ✓ 492ms | ✓ 846ms | ✓ 747ms | ✓ 1121ms | ✓ 1432ms | http |
| 101.43.127.100:8877 | ✓ 862ms | ✓ 1108ms | ✓ 939ms | ✓ 1092ms | ✓ 879ms | http |
| 222.228.171.92:8080 | ✓ 1499ms | ✓ 1651ms | ✓ 936ms | ✓ 1277ms | ✓ 945ms | http |
| 147.161.239.240:8800 | ✓ 1567ms | ✓ 1701ms | ✓ 1587ms | ✓ 1684ms | ✓ 1515ms | http |
| 91.233.223.147:3128 | ✓ 1627ms | 否 | ✓ 1458ms | 否 | ✓ 1713ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1969ms | ✓ 1616ms | ✓ 1414ms | http |
| 218.89.134.230:3333 | ✓ 1450ms | ✓ 1683ms | ✓ 1577ms | ✓ 1589ms | ✓ 1297ms | http |
| 38.145.208.172:8448 | ✓ 1925ms | ✓ 1706ms | ✓ 1474ms | 否 | 否 | http |
| 155.212.132.241:3128 | ✓ 710ms | ✓ 1870ms | ✓ 762ms | ✓ 1774ms | ✓ 1399ms | http |
| 38.145.220.173:8450 | ✓ 1250ms | ✓ 1612ms | ✓ 720ms | ✓ 968ms | ✓ 1582ms | http |
| 38.34.179.70:8452 | ✓ 1439ms | ✓ 1867ms | ✓ 564ms | 否 | ✓ 1059ms | http |
| 45.136.130.190:8450 | ✓ 1248ms | ✓ 1819ms | ✓ 944ms | ✓ 1422ms | ✓ 1610ms | http |
| 38.145.208.213:8450 | ✓ 1747ms | ✓ 726ms | ✓ 482ms | ✓ 1240ms | ✓ 1740ms | http |
| 103.82.23.118:5207 | ✓ 1404ms | 否 | ✓ 1334ms | 否 | ✓ 1539ms | http |
| 120.26.147.60:7890 | 否 | ✓ 1473ms | ✓ 1634ms | ✓ 1808ms | 否 | http |
| 185.115.74.185:8080 | ✓ 1047ms | ✓ 1841ms | ✓ 1437ms | 否 | 否 | http |
| 121.43.196.213:8222 | ✓ 1264ms | ✓ 1166ms | ✓ 906ms | ✓ 1247ms | ✓ 929ms | http |
| 137.220.150.104:6005 | ✓ 1224ms | ✓ 1927ms | ✓ 993ms | ✓ 1214ms | ✓ 1022ms | http |
| 137.220.150.22:6005 | ✓ 1450ms | 否 | ✓ 1049ms | ✓ 1263ms | 否 | http |
| 85.208.51.165:443 | ✓ 1769ms | ✓ 1653ms | ✓ 912ms | 否 | 否 | http |
| 194.67.99.223:1080 | ✓ 670ms | ✓ 1829ms | ✓ 1425ms | ✓ 1509ms | ✓ 1631ms | http |
| 101.47.73.135:3128 | ✓ 1303ms | 否 | ✓ 1575ms | ✓ 1254ms | ✓ 1145ms | http |
| 162.243.149.86:31028 | ✓ 306ms | ✓ 1968ms | 否 | 否 | ✓ 1606ms | http |
| 38.34.183.224:8448 | 否 | ✓ 1061ms | ✓ 960ms | 否 | ✓ 1987ms | http |
| 116.80.59.67:3128 | 否 | 否 | ✓ 1522ms | ✓ 1842ms | ✓ 1695ms | http |
| 47.77.193.180:1080 | ✓ 155ms | ✓ 809ms | ✓ 295ms | ✓ 733ms | ✓ 577ms | http |
| 185.41.152.110:3128 | ✓ 1543ms | ✓ 1593ms | ✓ 1418ms | ✓ 1997ms | ✓ 1836ms | http |
| 217.76.245.80:999 | ✓ 760ms | ✓ 1928ms | ✓ 1185ms | ✓ 1531ms | ✓ 1412ms | http |
| 101.32.244.83:8080 | ✓ 988ms | ✓ 1474ms | ✓ 1804ms | ✓ 1193ms | ✓ 1231ms | http |
| 121.43.196.210:8222 | ✓ 940ms | ✓ 1098ms | ✓ 900ms | ✓ 1130ms | ✓ 937ms | http |
| 114.55.226.123:10086 | ✓ 1144ms | ✓ 1406ms | ✓ 1001ms | ✓ 1256ms | ✓ 1067ms | http |
| 137.184.1.87:3128 | ✓ 473ms | ✓ 720ms | ✓ 979ms | ✓ 735ms | ✓ 620ms | http |
| 134.209.153.66:3128 | ✓ 1261ms | 否 | ✓ 1573ms | ✓ 1543ms | ✓ 1189ms | http |
| 88.80.150.82:8080 | ✓ 1394ms | 否 | ✓ 1687ms | 否 | ✓ 1778ms | https |
| 164.92.137.207:3128 | ✓ 800ms | ✓ 1993ms | 否 | ✓ 1879ms | ✓ 1892ms | http |
| 45.136.198.40:3128 | 否 | ✓ 1728ms | ✓ 1903ms | 否 | ✓ 1892ms | http |
| 111.79.111.126:3128 | ✓ 1159ms | ✓ 1451ms | 否 | ✓ 1712ms | 否 | http |
| 45.136.130.173:8448 | ✓ 1509ms | ✓ 731ms | ✓ 983ms | 否 | 否 | http |
| 15.204.151.141:3128 | ✓ 1066ms | ✓ 1894ms | ✓ 1767ms | 否 | 否 | http |
| 45.186.6.104:3128 | ✓ 1794ms | ✓ 1676ms | ✓ 1783ms | 否 | 否 | http |
| 137.220.150.152:6005 | ✓ 918ms | ✓ 1983ms | ✓ 907ms | ✓ 1214ms | ✓ 1111ms | http |
| 137.220.151.110:6005 | ✓ 744ms | ✓ 1785ms | ✓ 754ms | ✓ 1104ms | ✓ 1840ms | http |
| 45.136.130.168:8448 | ✓ 1076ms | ✓ 879ms | ✓ 796ms | 否 | ✓ 818ms | http |
| 45.136.130.170:8448 | ✓ 1085ms | ✓ 876ms | ✓ 792ms | 否 | ✓ 831ms | http |
| 38.145.208.176:8449 | ✓ 812ms | ✓ 768ms | ✓ 1176ms | ✓ 1365ms | ✓ 1151ms | http |
| 38.34.179.186:8451 | ✓ 1458ms | ✓ 1059ms | ✓ 522ms | ✓ 941ms | ✓ 840ms | http |
| 38.34.179.99:8444 | ✓ 908ms | ✓ 880ms | ✓ 1253ms | ✓ 1493ms | ✓ 875ms | http |
| 45.8.157.38:3128 | ✓ 368ms | ✓ 1286ms | ✓ 1546ms | 否 | ✓ 1422ms | http |
| 38.34.179.30:8449 | ✓ 908ms | 否 | ✓ 889ms | ✓ 885ms | ✓ 1063ms | http |
| 38.34.179.29:8449 | ✓ 907ms | 否 | ✓ 889ms | ✓ 887ms | ✓ 1082ms | http |
| 38.34.179.104:8450 | ✓ 907ms | ✓ 1261ms | ✓ 1748ms | ✓ 796ms | ✓ 890ms | http |
| 38.34.179.97:8450 | ✓ 906ms | ✓ 1335ms | ✓ 1723ms | ✓ 757ms | ✓ 889ms | http |
| 38.145.208.227:8453 | 否 | ✓ 866ms | ✓ 582ms | ✓ 1598ms | ✓ 1253ms | http |
| 38.34.183.222:8444 | ✓ 1805ms | ✓ 864ms | ✓ 1064ms | ✓ 1556ms | ✓ 1390ms | http |
| 38.145.208.227:8452 | 否 | ✓ 863ms | ✓ 485ms | ✓ 1185ms | ✓ 1701ms | http |
| 38.145.220.173:8444 | 否 | ✓ 1772ms | ✓ 1902ms | 否 | ✓ 1357ms | http |
| 202.141.161.53:30001 | ✓ 1038ms | ✓ 1411ms | ✓ 1222ms | ✓ 1300ms | ✓ 1199ms | http |
| 8.219.97.248:80 | ✓ 1308ms | 否 | ✓ 1677ms | 否 | ✓ 1276ms | http |
| 45.129.141.143:3128 | ✓ 1505ms | ✓ 1905ms | ✓ 1466ms | ✓ 1972ms | ✓ 1786ms | http |
| 103.39.51.190:8080 | ✓ 1876ms | 否 | 否 | ✓ 1637ms | ✓ 1429ms | http |
| 160.250.4.13:1 | ✓ 1732ms | 否 | ✓ 1685ms | ✓ 1650ms | 否 | http |
| 144.31.137.23:8080 | ✓ 1442ms | ✓ 1837ms | ✓ 1376ms | ✓ 1519ms | 否 | http |
| 137.220.150.170:6005 | ✓ 1941ms | 否 | ✓ 1942ms | ✓ 1987ms | ✓ 1132ms | http |
| 38.145.218.229:8450 | ✓ 596ms | ✓ 993ms | ✓ 329ms | ✓ 879ms | ✓ 859ms | http |
| 38.34.179.73:8446 | 否 | ✓ 1163ms | ✓ 1238ms | ✓ 1835ms | ✓ 598ms | http |
| 38.34.179.83:8448 | ✓ 1635ms | ✓ 1203ms | 否 | ✓ 1692ms | 否 | http |
| 65.108.203.37:18080 | ✓ 791ms | ✓ 1951ms | ✓ 836ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 894ms | ✓ 1105ms | ✓ 1069ms | ✓ 1247ms | ✓ 966ms | http |
| 45.136.130.191:8453 | ✓ 667ms | ✓ 725ms | ✓ 363ms | ✓ 788ms | ✓ 963ms | http |
| 38.34.179.16:8451 | ✓ 1644ms | ✓ 773ms | ✓ 597ms | 否 | ✓ 1452ms | http |
| 121.126.185.63:25152 | ✓ 1979ms | 否 | 否 | ✓ 1933ms | ✓ 1750ms | http |
| 106.117.208.101:7890 | ✓ 1136ms | ✓ 1711ms | ✓ 1336ms | ✓ 1256ms | ✓ 1010ms | http |
| 45.140.147.82:1081 | ✓ 796ms | ✓ 1719ms | ✓ 987ms | ✓ 1532ms | ✓ 1309ms | http |
| 59.8.203.55:80 | 否 | ✓ 1661ms | 否 | ✓ 1557ms | ✓ 1387ms | http |
| 185.104.249.25:3128 | ✓ 757ms | 否 | ✓ 1339ms | 否 | ✓ 1884ms | http |
| 37.187.109.70:10111 | ✓ 1607ms | 否 | ✓ 1859ms | ✓ 1791ms | ✓ 1387ms | http |
| 62.234.206.73:3128 | ✓ 928ms | ✓ 1234ms | ✓ 988ms | ✓ 1186ms | ✓ 942ms | http |

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
