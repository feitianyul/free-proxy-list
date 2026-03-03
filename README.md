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

最后更新：2026-03-03 22:39:25 UTC（2026-03-04 06:39:25 UTC+8）

**代理总数：93**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 918ms | 否 | ✓ 834ms | ✓ 1210ms | ✓ 1077ms | http |
| 3.213.157.4:3128 | ✓ 702ms | 否 | ✓ 681ms | ✓ 1512ms | ✓ 1289ms | http |
| 166.0.192.117:8888 | ✓ 966ms | ✓ 1826ms | ✓ 1780ms | ✓ 998ms | ✓ 906ms | http |
| 192.166.82.55:1080 | ✓ 944ms | ✓ 1367ms | ✓ 1472ms | ✓ 1437ms | ✓ 1105ms | http |
| 121.128.121.54:3128 | ✓ 1364ms | ✓ 1412ms | ✓ 1034ms | ✓ 1082ms | ✓ 860ms | http |
| 14.56.177.44:3128 | ✓ 1372ms | ✓ 1665ms | ✓ 969ms | ✓ 1104ms | ✓ 973ms | http |
| 205.209.118.30:3138 | ✓ 711ms | ✓ 1021ms | ✓ 811ms | ✓ 1398ms | ✓ 905ms | http |
| 125.128.12.144:3128 | ✓ 1368ms | ✓ 1652ms | ✓ 796ms | ✓ 1071ms | ✓ 1722ms | http |
| 14.56.107.244:3128 | ✓ 1969ms | ✓ 1430ms | ✓ 1302ms | ✓ 1399ms | 否 | http |
| 45.140.147.82:1081 | ✓ 546ms | 否 | ✓ 1228ms | 否 | ✓ 1374ms | http |
| 125.128.12.14:3128 | ✓ 708ms | ✓ 925ms | ✓ 1092ms | ✓ 1048ms | ✓ 823ms | http |
| 183.249.5.111:22222 | ✓ 1951ms | ✓ 1247ms | 否 | ✓ 1057ms | ✓ 792ms | http |
| 120.232.242.119:22222 | ✓ 956ms | ✓ 1356ms | ✓ 952ms | ✓ 1211ms | 否 | http |
| 120.240.29.51:22222 | ✓ 1322ms | 否 | ✓ 1122ms | ✓ 1309ms | ✓ 990ms | http |
| 61.72.110.54:3128 | ✓ 1398ms | 否 | 否 | ✓ 1464ms | ✓ 871ms | http |
| 59.46.216.131:30001 | ✓ 1122ms | ✓ 1408ms | ✓ 1114ms | ✓ 1443ms | ✓ 1168ms | http |
| 2.56.178.131:443 | ✓ 1059ms | 否 | ✓ 1722ms | 否 | ✓ 1877ms | http |
| 61.72.221.94:3128 | ✓ 697ms | 否 | ✓ 1995ms | ✓ 1391ms | ✓ 878ms | http |
| 61.72.221.234:3128 | ✓ 695ms | ✓ 1820ms | ✓ 1137ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1313ms | 否 | ✓ 1051ms | ✓ 1360ms | ✓ 1050ms | http |
| 5.75.196.26:40000 | ✓ 527ms | ✓ 1365ms | ✓ 1584ms | 否 | 否 | http |
| 61.72.110.94:3128 | ✓ 851ms | 否 | 否 | ✓ 1104ms | ✓ 863ms | http |
| 61.72.221.194:3128 | ✓ 1008ms | ✓ 1969ms | ✓ 903ms | 否 | 否 | http |
| 120.240.35.177:22222 | 否 | ✓ 1257ms | ✓ 1134ms | ✓ 1276ms | ✓ 968ms | http |
| 74.208.234.198:443 | ✓ 1055ms | ✓ 1536ms | 否 | ✓ 1530ms | 否 | http |
| 81.70.169.194:80 | ✓ 1190ms | ✓ 1452ms | ✓ 1008ms | ✓ 1320ms | ✓ 1096ms | http |
| 101.43.255.96:80 | ✓ 1101ms | ✓ 1419ms | ✓ 1111ms | ✓ 1322ms | ✓ 1113ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1189ms | 否 | ✓ 1330ms | ✓ 965ms | http |
| 222.184.48.235:22222 | ✓ 990ms | 否 | ✓ 1031ms | ✓ 1292ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1134ms | 否 | ✓ 1048ms | ✓ 1523ms | ✓ 1258ms | http |
| 103.82.23.118:5234 | ✓ 1573ms | ✓ 1897ms | ✓ 1254ms | 否 | ✓ 1689ms | http |
| 144.31.69.170:1080 | ✓ 829ms | 否 | ✓ 980ms | ✓ 1978ms | ✓ 1389ms | http |
| 91.193.240.157:9877 | ✓ 884ms | 否 | ✓ 875ms | ✓ 1964ms | ✓ 1400ms | http |
| 117.159.239.52:22222 | ✓ 1024ms | ✓ 1141ms | ✓ 926ms | ✓ 1169ms | ✓ 929ms | http |
| 117.159.239.51:22222 | ✓ 1046ms | ✓ 1131ms | ✓ 944ms | ✓ 1198ms | ✓ 921ms | http |
| 120.240.35.161:22222 | ✓ 985ms | ✓ 1302ms | ✓ 1117ms | ✓ 1232ms | ✓ 957ms | http |
| 47.95.231.180:8084 | ✓ 1092ms | ✓ 1379ms | ✓ 943ms | ✓ 1289ms | ✓ 1012ms | http |
| 113.59.32.142:22222 | ✓ 1141ms | ✓ 1401ms | ✓ 1017ms | ✓ 1257ms | ✓ 1060ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1063ms | ✓ 1342ms | ✓ 1847ms | http |
| 113.59.32.161:22222 | ✓ 1097ms | ✓ 1400ms | ✓ 962ms | ✓ 1373ms | ✓ 1061ms | http |
| 120.79.99.232:8099 | ✓ 1410ms | ✓ 1632ms | ✓ 1301ms | ✓ 1523ms | ✓ 1265ms | http |
| 185.243.218.43:49153 | ✓ 1463ms | ✓ 1700ms | ✓ 1744ms | 否 | ✓ 1711ms | http |
| 103.86.131.62:80 | ✓ 1944ms | 否 | 否 | ✓ 1352ms | ✓ 1089ms | http |
| 47.105.98.23:3128 | ✓ 938ms | ✓ 1224ms | 否 | 否 | ✓ 1024ms | http |
| 150.107.140.238:3128 | ✓ 917ms | 否 | ✓ 1329ms | ✓ 1535ms | 否 | http |
| 171.234.62.116:10008 | ✓ 1909ms | 否 | ✓ 1821ms | ✓ 1501ms | ✓ 1346ms | http |
| 35.234.17.221:8080 | ✓ 1660ms | ✓ 1174ms | ✓ 1341ms | ✓ 1683ms | 否 | http |
| 192.71.213.85:5555 | ✓ 1880ms | 否 | ✓ 1816ms | ✓ 1769ms | 否 | http |
| 121.230.8.11:1080 | ✓ 1392ms | ✓ 1561ms | ✓ 1401ms | ✓ 1643ms | ✓ 1497ms | http |
| 90.84.188.97:8000 | ✓ 1478ms | ✓ 1751ms | 否 | 否 | ✓ 1574ms | http |
| 103.84.95.54:7890 | ✓ 730ms | 否 | 否 | ✓ 1197ms | ✓ 1314ms | http |
| 113.59.32.162:22222 | ✓ 1164ms | ✓ 1372ms | ✓ 1107ms | ✓ 1336ms | ✓ 1069ms | http |
| 121.204.158.249:3128 | ✓ 1174ms | ✓ 1342ms | ✓ 1128ms | ✓ 1353ms | ✓ 1062ms | http |
| 34.101.184.164:3128 | ✓ 1216ms | 否 | ✓ 1043ms | ✓ 1479ms | ✓ 1050ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1550ms | 否 | ✓ 1632ms | ✓ 1912ms | http |
| 182.253.160.168:1452 | ✓ 1894ms | 否 | ✓ 1089ms | ✓ 1493ms | ✓ 1479ms | http |
| 94.177.131.12:3128 | ✓ 606ms | ✓ 1597ms | ✓ 1140ms | ✓ 1031ms | ✓ 988ms | http |
| 37.27.100.112:443 | ✓ 1699ms | ✓ 1533ms | ✓ 686ms | 否 | 否 | http |
| 47.101.149.27:9010 | 否 | ✓ 1878ms | ✓ 1918ms | ✓ 1551ms | 否 | http |
| 106.14.203.63:3333 | ✓ 981ms | ✓ 1624ms | ✓ 999ms | 否 | ✓ 951ms | http |
| 157.230.220.25:4857 | ✓ 751ms | 否 | ✓ 1319ms | 否 | ✓ 794ms | http |
| 24.199.124.151:3128 | ✓ 366ms | 否 | ✓ 993ms | ✓ 915ms | ✓ 1003ms | http |
| 165.227.5.10:8888 | ✓ 1547ms | ✓ 916ms | ✓ 529ms | 否 | ✓ 650ms | http |
| 159.89.31.62:8080 | ✓ 1099ms | ✓ 1752ms | ✓ 669ms | ✓ 1305ms | ✓ 1028ms | http |
| 77.83.203.6:443 | ✓ 1113ms | ✓ 1802ms | ✓ 1795ms | 否 | ✓ 1929ms | http |
| 45.136.198.40:3128 | ✓ 760ms | 否 | ✓ 1651ms | 否 | ✓ 1591ms | http |
| 171.234.62.116:10002 | 否 | 否 | ✓ 1558ms | ✓ 1755ms | ✓ 1580ms | http |
| 186.148.180.46:999 | ✓ 1299ms | 否 | ✓ 1405ms | ✓ 1861ms | ✓ 1443ms | http |
| 103.215.36.88:16345 | ✓ 1072ms | ✓ 1362ms | ✓ 1160ms | ✓ 1433ms | ✓ 1118ms | http |
| 70.22.175.232:3128 | 否 | 否 | ✓ 298ms | ✓ 1154ms | ✓ 809ms | http |
| 74.48.78.224:2080 | 否 | ✓ 947ms | ✓ 511ms | ✓ 984ms | 否 | http |
| 154.90.48.209:9090 | ✓ 1320ms | 否 | ✓ 1707ms | ✓ 1628ms | ✓ 1301ms | http |
| 91.233.223.147:3128 | ✓ 881ms | 否 | ✓ 939ms | ✓ 1974ms | ✓ 1572ms | http |
| 120.55.163.237:10086 | ✓ 956ms | ✓ 1143ms | ✓ 981ms | ✓ 1255ms | ✓ 997ms | http |
| 103.39.51.190:8080 | ✓ 1961ms | 否 | 否 | ✓ 1707ms | ✓ 1604ms | http |
| 103.118.102.98:80 | ✓ 1970ms | 否 | ✓ 1999ms | ✓ 1769ms | ✓ 1715ms | http |
| 117.159.239.54:22222 | ✓ 951ms | ✓ 1145ms | ✓ 922ms | ✓ 1186ms | ✓ 926ms | http |
| 113.59.32.141:22222 | ✓ 1150ms | ✓ 1403ms | ✓ 1103ms | ✓ 1340ms | 否 | http |
| 103.215.36.88:18147 | ✓ 1074ms | ✓ 1299ms | ✓ 1108ms | ✓ 1421ms | ✓ 1087ms | http |
| 150.249.255.91:3128 | 否 | ✓ 935ms | 否 | ✓ 1980ms | ✓ 808ms | http |
| 8.219.97.248:80 | ✓ 1643ms | 否 | ✓ 1233ms | ✓ 1920ms | 否 | http |
| 46.249.103.192:443 | ✓ 818ms | 否 | ✓ 775ms | ✓ 1997ms | 否 | http |
| 121.230.9.160:1080 | ✓ 1319ms | ✓ 1597ms | ✓ 1160ms | ✓ 1456ms | ✓ 1253ms | http |
| 117.159.239.50:22222 | ✓ 901ms | ✓ 1154ms | ✓ 947ms | ✓ 1291ms | ✓ 929ms | http |
| 120.240.35.174:22222 | 否 | ✓ 1661ms | 否 | ✓ 1435ms | ✓ 1703ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1576ms | 否 | ✓ 1843ms | ✓ 1201ms | http |
| 103.3.246.71:3128 | ✓ 1178ms | 否 | ✓ 1069ms | ✓ 1373ms | ✓ 1014ms | http |
| 172.212.68.37:3128 | ✓ 343ms | ✓ 1460ms | ✓ 1040ms | ✓ 1563ms | ✓ 1139ms | http |
| 120.198.141.79:22222 | ✓ 1320ms | ✓ 1636ms | ✓ 1206ms | ✓ 1349ms | ✓ 1189ms | http |
| 116.80.63.67:7777 | ✓ 1708ms | 否 | ✓ 1604ms | 否 | ✓ 1715ms | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1107ms | ✓ 1205ms | ✓ 1422ms | http |
| 37.27.100.107:443 | ✓ 1004ms | ✓ 1549ms | 否 | 否 | ✓ 1917ms | http |
| 120.240.35.160:22222 | 否 | ✓ 1654ms | ✓ 1236ms | 否 | ✓ 1135ms | http |

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
