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

最后更新：2026-06-11 15:28:55 UTC（2026-06-11 23:28:55 UTC+8）

**代理总数：94**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 94 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 94 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.43.46.91:443 | ✓ 1608ms | 否 | ✓ 439ms | ✓ 1292ms | 否 | http |
| 34.43.46.91:80 | ✓ 1611ms | 否 | ✓ 973ms | ✓ 1100ms | ✓ 834ms | http |
| 159.198.35.187:1080 | ✓ 828ms | ✓ 1845ms | ✓ 1121ms | ✓ 1318ms | ✓ 1931ms | http |
| 37.49.224.15:3128 | ✓ 1428ms | 否 | ✓ 1799ms | ✓ 1793ms | 否 | http |
| 47.85.51.197:1080 | ✓ 306ms | 否 | 否 | ✓ 1934ms | ✓ 718ms | http |
| 176.111.37.5:39811 | ✓ 597ms | 否 | ✓ 544ms | ✓ 1633ms | ✓ 1241ms | http |
| 47.238.106.120:666 | ✓ 1031ms | ✓ 1293ms | ✓ 1019ms | ✓ 1105ms | ✓ 892ms | http |
| 185.141.26.131:3128 | ✓ 581ms | ✓ 1421ms | ✓ 1247ms | ✓ 1771ms | ✓ 1272ms | http |
| 169.212.15.161:5000 | ✓ 844ms | ✓ 1081ms | ✓ 959ms | ✓ 1331ms | ✓ 965ms | http |
| 167.86.95.198:3128 | ✓ 650ms | 否 | ✓ 920ms | 否 | ✓ 1859ms | http |
| 185.200.188.234:10001 | ✓ 1051ms | 否 | ✓ 739ms | 否 | ✓ 1622ms | http |
| 34.87.80.221:30000 | ✓ 1045ms | 否 | ✓ 1443ms | ✓ 1374ms | ✓ 1293ms | http |
| 185.106.183.85:8080 | ✓ 1536ms | 否 | ✓ 607ms | 否 | ✓ 1542ms | http |
| 47.80.103.120:8080 | ✓ 1599ms | 否 | ✓ 1402ms | 否 | ✓ 1073ms | http |
| 152.32.132.190:7890 | ✓ 1905ms | 否 | ✓ 1588ms | ✓ 1027ms | 否 | http |
| 47.80.112.92:8080 | ✓ 1082ms | 否 | ✓ 1319ms | ✓ 1419ms | ✓ 1147ms | http |
| 1.231.81.166:3128 | ✓ 730ms | ✓ 1052ms | ✓ 911ms | ✓ 1100ms | ✓ 898ms | http |
| 107.150.61.226:8886 | ✓ 277ms | ✓ 955ms | ✓ 215ms | ✓ 1032ms | ✓ 1512ms | http |
| 176.111.37.216:39811 | ✓ 1086ms | ✓ 1611ms | ✓ 1218ms | 否 | ✓ 1426ms | http |
| 84.47.150.125:1080 | ✓ 1086ms | 否 | ✓ 1400ms | 否 | ✓ 1586ms | http |
| 174.137.134.182:2999 | 否 | 否 | ✓ 670ms | ✓ 843ms | ✓ 1375ms | http |
| 14.143.222.113:10144 | ✓ 1018ms | 否 | ✓ 965ms | ✓ 1416ms | 否 | http |
| 152.42.239.142:3128 | ✓ 1104ms | 否 | 否 | ✓ 1644ms | ✓ 1017ms | http |
| 81.177.214.151:8080 | ✓ 1251ms | 否 | ✓ 1712ms | ✓ 1748ms | ✓ 1466ms | http |
| 95.3.69.222:8080 | 否 | ✓ 1853ms | 否 | ✓ 1621ms | ✓ 1425ms | http |
| 20.78.26.206:8561 | ✓ 644ms | ✓ 1419ms | ✓ 714ms | ✓ 1002ms | ✓ 785ms | http |
| 91.186.213.124:1081 | ✓ 974ms | 否 | ✓ 560ms | ✓ 1849ms | 否 | http |
| 20.27.11.248:8561 | ✓ 1763ms | ✓ 1433ms | ✓ 1008ms | 否 | 否 | http |
| 138.124.106.119:8080 | ✓ 484ms | 否 | ✓ 1046ms | 否 | ✓ 1250ms | http |
| 34.71.229.255:3128 | ✓ 494ms | ✓ 1309ms | ✓ 1609ms | ✓ 1186ms | ✓ 962ms | http |
| 144.31.73.173:3128 | ✓ 833ms | 否 | ✓ 1164ms | ✓ 1643ms | ✓ 1662ms | http |
| 81.200.154.236:48503 | ✓ 413ms | ✓ 1857ms | ✓ 1448ms | ✓ 1908ms | 否 | http |
| 139.226.3.89:60000 | ✓ 1039ms | 否 | ✓ 995ms | ✓ 1272ms | ✓ 1054ms | http |
| 116.80.80.225:3172 | ✓ 1991ms | 否 | ✓ 1659ms | 否 | ✓ 1828ms | http |
| 116.80.63.158:3172 | ✓ 1712ms | 否 | ✓ 1742ms | 否 | ✓ 1803ms | http |
| 20.78.118.91:8561 | 否 | ✓ 1427ms | ✓ 719ms | ✓ 1155ms | ✓ 950ms | http |
| 52.188.28.218:3128 | ✓ 35ms | 否 | ✓ 340ms | ✓ 839ms | ✓ 1403ms | http |
| 152.53.52.47:1080 | ✓ 242ms | 否 | ✓ 858ms | ✓ 1111ms | ✓ 938ms | http |
| 20.27.14.220:8561 | ✓ 1541ms | ✓ 1354ms | ✓ 835ms | ✓ 1047ms | ✓ 769ms | http |
| 43.160.236.170:8118 | 否 | 否 | ✓ 1376ms | ✓ 1291ms | ✓ 1041ms | http |
| 68.183.53.204:3128 | ✓ 194ms | ✓ 1020ms | ✓ 172ms | 否 | ✓ 818ms | http |
| 20.27.15.111:8561 | ✓ 754ms | ✓ 1192ms | ✓ 651ms | ✓ 956ms | ✓ 801ms | http |
| 180.2.108.38:8080 | ✓ 739ms | 否 | ✓ 1477ms | ✓ 1109ms | ✓ 912ms | http |
| 50.114.102.16:8888 | ✓ 594ms | ✓ 1843ms | ✓ 1598ms | ✓ 1727ms | ✓ 1196ms | http |
| 45.88.174.195:8080 | ✓ 582ms | 否 | ✓ 1239ms | 否 | ✓ 1477ms | http |
| 80.150.246.98:443 | ✓ 512ms | ✓ 1323ms | 否 | ✓ 1939ms | ✓ 1742ms | http |
| 47.80.112.92:80 | ✓ 1022ms | 否 | ✓ 978ms | ✓ 1369ms | ✓ 1153ms | http |
| 47.80.103.120:80 | ✓ 1106ms | 否 | ✓ 971ms | ✓ 1389ms | ✓ 1054ms | http |
| 18.180.59.181:80 | ✓ 703ms | 否 | 否 | ✓ 1377ms | ✓ 1836ms | http |
| 43.134.141.85:80 | ✓ 1445ms | ✓ 1642ms | ✓ 1560ms | ✓ 1736ms | ✓ 1530ms | http |
| 152.42.177.32:8888 | ✓ 1183ms | 否 | ✓ 1717ms | ✓ 1766ms | ✓ 1530ms | http |
| 101.32.243.189:80 | ✓ 1481ms | ✓ 1644ms | ✓ 1800ms | ✓ 1753ms | ✓ 1536ms | http |
| 43.156.228.168:80 | ✓ 1489ms | 否 | ✓ 1514ms | ✓ 1723ms | ✓ 1536ms | http |
| 203.30.9.8:8443 | ✓ 1445ms | 否 | ✓ 1926ms | 否 | ✓ 1510ms | http |
| 170.106.136.181:31002 | ✓ 469ms | ✓ 1564ms | ✓ 754ms | ✓ 881ms | ✓ 652ms | http |
| 144.172.114.214:1080 | ✓ 556ms | 否 | ✓ 1272ms | ✓ 1402ms | 否 | http |
| 2.26.87.216:1080 | ✓ 997ms | 否 | ✓ 1783ms | ✓ 1955ms | ✓ 1481ms | http |
| 147.102.16.32:8080 | ✓ 1289ms | 否 | 否 | ✓ 1791ms | ✓ 1208ms | http |
| 117.55.203.165:8899 | ✓ 1416ms | 否 | ✓ 1699ms | 否 | ✓ 1728ms | http |
| 62.133.62.17:1081 | ✓ 974ms | 否 | ✓ 1375ms | 否 | ✓ 1649ms | http |
| 144.31.156.131:1080 | ✓ 1003ms | 否 | ✓ 1570ms | 否 | ✓ 1536ms | http |
| 151.243.153.157:8118 | ✓ 832ms | 否 | ✓ 1311ms | 否 | ✓ 1327ms | http |
| 202.28.194.139:31280 | ✓ 1796ms | 否 | 否 | ✓ 1883ms | ✓ 1936ms | http |
| 212.58.132.5:8888 | ✓ 1627ms | 否 | ✓ 1550ms | ✓ 1492ms | ✓ 1369ms | http |
| 20.27.13.35:8561 | ✓ 698ms | ✓ 1452ms | ✓ 888ms | ✓ 1086ms | ✓ 916ms | http |
| 45.84.222.25:1080 | ✓ 1594ms | 否 | ✓ 1568ms | ✓ 1846ms | 否 | http |
| 62.133.62.184:1082 | ✓ 871ms | 否 | ✓ 780ms | ✓ 1749ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1084ms | ✓ 1665ms | ✓ 1229ms | ✓ 1727ms | ✓ 1268ms | http |
| 85.234.100.149:8080 | ✓ 416ms | 否 | 否 | ✓ 1901ms | ✓ 1098ms | http |
| 20.210.39.153:8561 | ✓ 676ms | 否 | ✓ 660ms | ✓ 1926ms | ✓ 879ms | http |
| 70.35.196.194:8085 | ✓ 1516ms | ✓ 1707ms | ✓ 1590ms | 否 | ✓ 1770ms | http |
| 106.10.55.212:1121 | ✓ 1902ms | 否 | ✓ 1234ms | 否 | ✓ 1936ms | http |
| 217.154.155.115:8080 | ✓ 1192ms | 否 | ✓ 608ms | ✓ 1248ms | ✓ 937ms | http |
| 5.78.60.155:8080 | ✓ 1095ms | 否 | ✓ 1357ms | 否 | ✓ 1467ms | http |
| 120.92.212.16:7890 | ✓ 1623ms | 否 | ✓ 1461ms | 否 | ✓ 1727ms | http |
| 62.133.62.12:1081 | ✓ 1010ms | 否 | ✓ 1132ms | 否 | ✓ 1754ms | http |
| 83.147.36.155:8080 | ✓ 652ms | 否 | 否 | ✓ 1320ms | ✓ 1360ms | http |
| 129.80.238.83:444 | ✓ 223ms | ✓ 1447ms | ✓ 1560ms | ✓ 956ms | 否 | http |
| 132.243.160.171:3128 | 否 | 否 | ✓ 626ms | ✓ 1819ms | ✓ 1556ms | http |
| 8.154.21.175:3128 | ✓ 1057ms | ✓ 1359ms | ✓ 1181ms | ✓ 1373ms | ✓ 1062ms | http |
| 121.138.61.78:8304 | 否 | 否 | ✓ 1241ms | ✓ 1263ms | ✓ 976ms | http |
| 38.180.2.107:3128 | ✓ 1211ms | ✓ 1824ms | ✓ 1309ms | 否 | 否 | http |
| 94.241.175.40:10808 | ✓ 531ms | 否 | ✓ 1628ms | 否 | ✓ 1943ms | http |
| 57.129.144.178:40000 | ✓ 1154ms | 否 | ✓ 1072ms | ✓ 1808ms | ✓ 1225ms | http |
| 157.245.143.65:7890 | ✓ 629ms | ✓ 1861ms | ✓ 1395ms | 否 | 否 | http |
| 193.29.224.20:3128 | 否 | 否 | ✓ 1039ms | ✓ 1909ms | ✓ 1424ms | http |
| 213.21.254.26:1081 | 否 | 否 | ✓ 550ms | ✓ 1572ms | ✓ 1174ms | http |
| 178.63.155.151:8890 | ✓ 1153ms | 否 | ✓ 1261ms | ✓ 1953ms | 否 | http |
| 152.67.191.232:6800 | ✓ 1468ms | 否 | 否 | ✓ 1372ms | ✓ 1058ms | http |
| 61.52.131.172:8443 | ✓ 1093ms | ✓ 1384ms | ✓ 1110ms | 否 | ✓ 1093ms | http |
| 199.127.62.89:3129 | ✓ 1363ms | 否 | ✓ 954ms | 否 | ✓ 1993ms | http |
| 117.55.203.163:8899 | ✓ 647ms | 否 | ✓ 1104ms | 否 | ✓ 1954ms | http |
| 20.210.76.178:8561 | 否 | ✓ 1608ms | ✓ 1989ms | ✓ 1656ms | ✓ 964ms | http |
| 20.210.76.175:8561 | 否 | 否 | ✓ 1965ms | ✓ 1493ms | ✓ 803ms | http |

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
