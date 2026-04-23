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

最后更新：2026-04-23 00:39:36 UTC（2026-04-23 08:39:36 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 99ms | ✓ 1117ms | ✓ 534ms | ✓ 1072ms | ✓ 866ms | http |
| 218.108.131.186:17890 | ✓ 771ms | ✓ 955ms | ✓ 787ms | 否 | 否 | http |
| 152.42.208.139:8118 | ✓ 987ms | 否 | ✓ 1642ms | ✓ 1275ms | ✓ 941ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1638ms | ✓ 1595ms | 否 | ✓ 1089ms | http |
| 1.231.81.166:3128 | ✓ 1810ms | ✓ 1247ms | ✓ 1114ms | ✓ 1059ms | ✓ 856ms | http |
| 35.225.22.61:80 | 否 | ✓ 1321ms | 否 | ✓ 1083ms | ✓ 904ms | http |
| 59.46.216.131:30001 | ✓ 1144ms | ✓ 1483ms | ✓ 1241ms | ✓ 1583ms | ✓ 1209ms | http |
| 152.32.132.190:7890 | ✓ 1284ms | 否 | ✓ 1293ms | ✓ 1008ms | 否 | http |
| 147.45.60.34:1082 | 否 | ✓ 1169ms | ✓ 336ms | ✓ 1237ms | ✓ 1341ms | http |
| 103.35.191.173:1081 | ✓ 832ms | 否 | ✓ 1452ms | ✓ 1517ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1923ms | 否 | 否 | ✓ 1931ms | ✓ 1725ms | http |
| 45.76.207.177:40000 | ✓ 1632ms | 否 | ✓ 996ms | ✓ 1541ms | ✓ 1413ms | http |
| 45.153.231.229:8080 | ✓ 1063ms | ✓ 1786ms | ✓ 1247ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 998ms | ✓ 1260ms | ✓ 1007ms | 否 | 否 | http |
| 130.61.174.200:1080 | 否 | 否 | ✓ 940ms | ✓ 1964ms | ✓ 1255ms | http |
| 2.27.40.180:1080 | ✓ 1719ms | 否 | ✓ 1110ms | 否 | ✓ 1973ms | http |
| 91.99.15.45:2095 | ✓ 1689ms | ✓ 1968ms | ✓ 1219ms | ✓ 1891ms | ✓ 1897ms | http |
| 120.92.212.16:7890 | ✓ 1120ms | ✓ 1638ms | ✓ 1161ms | ✓ 1556ms | ✓ 1730ms | http |
| 120.92.212.16:8890 | ✓ 1068ms | 否 | ✓ 1265ms | ✓ 1508ms | 否 | http |
| 159.223.225.118:8888 | ✓ 1990ms | 否 | ✓ 1831ms | 否 | ✓ 1486ms | http |
| 46.101.95.183:8888 | ✓ 1624ms | 否 | ✓ 1438ms | 否 | ✓ 1445ms | http |
| 78.11.96.22:8888 | ✓ 883ms | ✓ 1481ms | ✓ 1308ms | ✓ 1502ms | ✓ 1320ms | http |
| 84.47.150.125:1080 | ✓ 1042ms | ✓ 1803ms | ✓ 1557ms | 否 | ✓ 1524ms | http |
| 217.77.102.18:3128 | ✓ 1193ms | ✓ 1998ms | 否 | ✓ 1916ms | ✓ 1388ms | http |
| 139.159.97.82:10900 | ✓ 1667ms | 否 | ✓ 1501ms | ✓ 1625ms | ✓ 1312ms | http |
| 168.144.75.9:3128 | ✓ 1657ms | 否 | ✓ 1846ms | 否 | ✓ 1990ms | http |
| 8.219.195.129:1080 | ✓ 944ms | ✓ 1862ms | ✓ 995ms | ✓ 1200ms | ✓ 957ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1811ms | 否 | ✓ 1859ms | ✓ 1523ms | http |
| 212.58.132.5:8888 | ✓ 1585ms | 否 | ✓ 1253ms | ✓ 1555ms | ✓ 1285ms | http |
| 8.209.238.110:47701 | ✓ 662ms | ✓ 1998ms | ✓ 999ms | ✓ 1765ms | ✓ 1247ms | http |
| 14.143.222.113:57788 | ✓ 978ms | 否 | ✓ 1006ms | ✓ 1369ms | 否 | http |
| 20.127.128.70:8080 | ✓ 1267ms | 否 | ✓ 1445ms | ✓ 1842ms | ✓ 1117ms | http |
| 85.190.99.143:443 | ✓ 1172ms | 否 | ✓ 1297ms | 否 | ✓ 1811ms | http |
| 72.56.87.46:3128 | ✓ 1080ms | ✓ 1660ms | ✓ 813ms | 否 | 否 | http |
| 129.213.162.27:17777 | ✓ 1588ms | ✓ 1217ms | 否 | 否 | ✓ 1411ms | http |
| 34.71.229.255:3128 | ✓ 1110ms | ✓ 1506ms | ✓ 877ms | ✓ 1343ms | ✓ 939ms | http |
| 60.249.94.208:3128 | 否 | ✓ 1149ms | ✓ 819ms | ✓ 1083ms | ✓ 879ms | http |
| 177.93.132.244:3128 | ✓ 1868ms | 否 | ✓ 708ms | 否 | ✓ 1992ms | http |
| 168.222.254.136:8888 | 否 | ✓ 1650ms | 否 | ✓ 1903ms | ✓ 1637ms | http |
| 38.180.192.119:3128 | ✓ 1501ms | ✓ 1109ms | ✓ 830ms | ✓ 1035ms | ✓ 696ms | http |
| 121.230.9.209:1080 | 否 | 否 | ✓ 1196ms | ✓ 1939ms | ✓ 1380ms | http |
| 223.84.151.86:30005 | ✓ 1473ms | ✓ 1874ms | ✓ 1468ms | ✓ 1820ms | ✓ 1647ms | http |
| 103.160.182.80:1111 | ✓ 1956ms | 否 | ✓ 1507ms | 否 | ✓ 1736ms | http |
| 68.178.167.154:9999 | 否 | 否 | ✓ 1578ms | ✓ 1182ms | ✓ 1195ms | http |
| 49.48.54.172:8080 | ✓ 1542ms | 否 | 否 | ✓ 1873ms | ✓ 1664ms | http |
| 83.219.250.8:62920 | ✓ 1238ms | 否 | ✓ 1470ms | ✓ 1583ms | ✓ 1412ms | http |
| 91.233.223.147:3128 | ✓ 1218ms | ✓ 1942ms | ✓ 1175ms | ✓ 1870ms | ✓ 1506ms | http |
| 194.150.220.163:1082 | ✓ 1709ms | ✓ 1477ms | ✓ 1945ms | 否 | 否 | http |
| 91.193.240.157:9877 | ✓ 1040ms | 否 | ✓ 1672ms | 否 | ✓ 1753ms | http |
| 51.79.71.106:8080 | ✓ 1798ms | 否 | ✓ 1793ms | ✓ 1941ms | ✓ 1519ms | http |
| 62.113.119.14:8080 | ✓ 829ms | ✓ 1533ms | ✓ 603ms | ✓ 1512ms | ✓ 1085ms | http |
| 45.59.122.132:80 | ✓ 1076ms | ✓ 1827ms | ✓ 1280ms | ✓ 1513ms | ✓ 1348ms | http |
| 103.56.112.222:7890 | ✓ 869ms | ✓ 1755ms | ✓ 1326ms | ✓ 1341ms | ✓ 1022ms | http |
| 85.208.51.165:443 | ✓ 1274ms | 否 | ✓ 1609ms | 否 | ✓ 1798ms | http |
| 45.88.0.115:3128 | ✓ 755ms | ✓ 1285ms | ✓ 1859ms | 否 | ✓ 1422ms | http |
| 144.31.25.69:21064 | ✓ 759ms | 否 | ✓ 1759ms | 否 | ✓ 1879ms | http |
| 45.88.0.113:3128 | ✓ 624ms | ✓ 1146ms | ✓ 710ms | ✓ 1258ms | ✓ 959ms | http |
| 45.88.0.98:3128 | ✓ 427ms | ✓ 1165ms | ✓ 402ms | ✓ 1246ms | ✓ 956ms | http |
| 45.88.0.117:3128 | ✓ 392ms | ✓ 1129ms | ✓ 429ms | ✓ 1270ms | ✓ 983ms | http |
| 45.88.0.116:3128 | ✓ 416ms | ✓ 1240ms | ✓ 403ms | ✓ 1248ms | ✓ 978ms | http |
| 45.88.0.111:3128 | ✓ 411ms | ✓ 1241ms | ✓ 406ms | ✓ 1275ms | ✓ 962ms | http |
| 208.87.243.199:7878 | ✓ 581ms | ✓ 1106ms | ✓ 446ms | ✓ 1255ms | ✓ 1093ms | http |
| 117.122.240.82:3338 | ✓ 1833ms | ✓ 1772ms | 否 | ✓ 1073ms | ✓ 886ms | http |
| 47.74.226.8:5001 | ✓ 1485ms | ✓ 1584ms | 否 | ✓ 1677ms | 否 | http |
| 45.129.141.143:3128 | ✓ 929ms | ✓ 1934ms | ✓ 1573ms | 否 | ✓ 1734ms | http |
| 45.186.6.104:3128 | ✓ 1202ms | ✓ 1928ms | ✓ 1726ms | 否 | 否 | http |
| 45.88.0.99:3128 | ✓ 508ms | ✓ 1311ms | ✓ 1273ms | 否 | ✓ 1516ms | http |
| 138.124.108.176:3128 | ✓ 1519ms | ✓ 1463ms | ✓ 1476ms | 否 | ✓ 1175ms | http |
| 47.101.159.19:8899 | 否 | ✓ 1203ms | ✓ 1092ms | ✓ 1202ms | ✓ 1015ms | http |
| 8.140.104.98:3128 | 否 | ✓ 1139ms | ✓ 1890ms | ✓ 1260ms | ✓ 1012ms | http |
| 103.82.93.98:3128 | 否 | 否 | ✓ 905ms | ✓ 1616ms | ✓ 1174ms | http |
| 47.110.42.192:9003 | 否 | ✓ 1624ms | ✓ 1422ms | ✓ 1809ms | ✓ 1496ms | http |
| 77.223.107.108:3128 | 否 | ✓ 1855ms | ✓ 1912ms | 否 | ✓ 1565ms | http |
| 202.141.161.53:10808 | ✓ 1931ms | ✓ 1327ms | 否 | 否 | ✓ 1071ms | http |
| 20.78.213.56:80 | 否 | ✓ 1250ms | ✓ 1604ms | ✓ 1346ms | ✓ 847ms | http |
| 109.234.38.35:3128 | 否 | ✓ 1535ms | ✓ 1255ms | ✓ 1592ms | ✓ 1258ms | http |
| 161.97.184.191:8080 | ✓ 1161ms | ✓ 1972ms | 否 | 否 | ✓ 1956ms | http |
| 128.199.247.154:8080 | 否 | 否 | ✓ 1782ms | ✓ 1404ms | ✓ 1076ms | http |
| 148.251.86.68:16379 | 否 | 否 | ✓ 669ms | ✓ 1749ms | ✓ 1514ms | http |
| 213.220.62.63:3128 | ✓ 917ms | ✓ 1605ms | ✓ 943ms | ✓ 1544ms | ✓ 1203ms | http |
| 116.171.106.26:3443 | ✓ 1527ms | ✓ 1532ms | ✓ 1742ms | ✓ 1819ms | 否 | http |
| 162.240.154.26:3128 | 否 | ✓ 1793ms | ✓ 1910ms | 否 | ✓ 1763ms | http |
| 89.208.106.138:10808 | ✓ 579ms | 否 | 否 | ✓ 1854ms | ✓ 1244ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1953ms | ✓ 1847ms | ✓ 1723ms | http |
| 61.52.131.172:8443 | ✓ 1239ms | ✓ 1030ms | 否 | 否 | ✓ 853ms | http |
| 104.248.195.47:8080 | ✓ 923ms | ✓ 1697ms | ✓ 1602ms | ✓ 1862ms | 否 | http |

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
