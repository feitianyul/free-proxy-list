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

最后更新：2026-04-30 08:18:33 UTC（2026-04-30 16:18:33 UTC+8）

**代理总数：74**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 833ms | ✓ 1371ms | 否 | ✓ 1153ms | ✓ 1124ms | http |
| 47.85.51.197:1080 | 否 | ✓ 1799ms | 否 | ✓ 1230ms | ✓ 651ms | http |
| 34.96.238.40:8080 | ✓ 1063ms | 否 | ✓ 1391ms | ✓ 1322ms | ✓ 1385ms | http |
| 46.101.95.183:8888 | ✓ 1272ms | ✓ 1978ms | ✓ 444ms | ✓ 1810ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1120ms | 否 | ✓ 1205ms | ✓ 1481ms | ✓ 1375ms | http |
| 1.231.81.166:3128 | ✓ 1403ms | 否 | ✓ 1644ms | ✓ 1757ms | ✓ 993ms | http |
| 113.160.132.26:8080 | ✓ 1539ms | ✓ 1774ms | ✓ 1078ms | 否 | ✓ 1119ms | http |
| 115.231.181.40:8128 | ✓ 1259ms | 否 | ✓ 1495ms | 否 | ✓ 1210ms | http |
| 154.64.232.35:8080 | ✓ 1646ms | 否 | ✓ 1680ms | 否 | ✓ 1837ms | http |
| 45.167.124.71:999 | ✓ 1255ms | ✓ 1590ms | ✓ 1319ms | ✓ 1738ms | ✓ 1404ms | http |
| 107.173.160.222:1080 | ✓ 360ms | ✓ 1495ms | ✓ 950ms | ✓ 1364ms | ✓ 904ms | http |
| 45.140.147.155:1081 | ✓ 608ms | ✓ 1442ms | ✓ 470ms | ✓ 1284ms | ✓ 1097ms | http |
| 45.140.147.155:1082 | ✓ 542ms | ✓ 1615ms | ✓ 410ms | ✓ 1183ms | ✓ 1136ms | http |
| 159.223.225.118:8888 | ✓ 558ms | 否 | ✓ 1378ms | ✓ 1717ms | ✓ 1324ms | http |
| 168.110.52.228:3128 | ✓ 958ms | 否 | ✓ 1447ms | ✓ 1091ms | ✓ 989ms | http |
| 103.3.246.71:3128 | ✓ 1178ms | 否 | ✓ 1330ms | ✓ 1424ms | ✓ 1159ms | http |
| 103.70.114.149:3128 | ✓ 1765ms | 否 | ✓ 1504ms | 否 | ✓ 1747ms | http |
| 132.226.235.199:1080 | 否 | 否 | ✓ 1544ms | ✓ 1624ms | ✓ 1780ms | http |
| 172.236.145.31:7890 | ✓ 1071ms | ✓ 1899ms | 否 | ✓ 1698ms | ✓ 1488ms | http |
| 86.104.74.110:1082 | ✓ 873ms | 否 | ✓ 1817ms | 否 | ✓ 1962ms | http |
| 218.108.131.186:17890 | ✓ 1213ms | ✓ 1285ms | ✓ 1088ms | ✓ 1312ms | ✓ 1422ms | http |
| 120.92.108.86:7890 | ✓ 1421ms | 否 | ✓ 1475ms | 否 | ✓ 1474ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1422ms | ✓ 1983ms | ✓ 1494ms | http |
| 8.154.21.175:3128 | ✓ 1049ms | ✓ 1373ms | ✓ 972ms | ✓ 1386ms | ✓ 1039ms | http |
| 65.109.213.99:1080 | ✓ 1489ms | 否 | ✓ 1289ms | 否 | ✓ 1849ms | http |
| 210.223.44.230:3128 | ✓ 1826ms | ✓ 1879ms | 否 | 否 | ✓ 1302ms | http |
| 94.72.109.214:8888 | ✓ 387ms | ✓ 1204ms | ✓ 412ms | ✓ 1195ms | ✓ 872ms | http |
| 38.92.10.152:57579 | ✓ 675ms | ✓ 977ms | ✓ 495ms | 否 | ✓ 1545ms | http |
| 206.206.126.177:2412 | ✓ 1236ms | 否 | ✓ 1469ms | ✓ 1864ms | ✓ 1154ms | http |
| 152.70.91.193:40000 | ✓ 1676ms | 否 | 否 | ✓ 1845ms | ✓ 1550ms | http |
| 103.157.200.126:3128 | ✓ 1154ms | 否 | ✓ 1142ms | ✓ 1578ms | ✓ 1246ms | http |
| 117.236.124.166:3128 | ✓ 1018ms | 否 | ✓ 1136ms | ✓ 1972ms | ✓ 1350ms | http |
| 38.180.192.119:3128 | ✓ 1558ms | ✓ 1478ms | ✓ 1175ms | ✓ 1693ms | ✓ 1584ms | http |
| 15.204.151.149:3128 | 否 | ✓ 1962ms | ✓ 982ms | ✓ 1442ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1376ms | ✓ 1951ms | ✓ 1128ms | ✓ 1374ms | ✓ 1014ms | http |
| 91.233.223.147:3128 | ✓ 1443ms | 否 | ✓ 1519ms | ✓ 1914ms | ✓ 1428ms | http |
| 130.61.174.200:1080 | ✓ 1366ms | ✓ 1644ms | 否 | 否 | ✓ 1091ms | http |
| 5.255.126.157:10001 | ✓ 409ms | 否 | ✓ 1916ms | 否 | ✓ 1010ms | http |
| 194.150.220.163:1082 | ✓ 677ms | 否 | ✓ 1821ms | ✓ 1825ms | 否 | http |
| 200.174.198.32:8888 | ✓ 1187ms | 否 | ✓ 1923ms | 否 | ✓ 1880ms | http |
| 80.92.204.47:1081 | ✓ 428ms | ✓ 1036ms | ✓ 743ms | ✓ 1568ms | ✓ 1043ms | http |
| 121.230.8.158:1080 | ✓ 1290ms | ✓ 1381ms | ✓ 1292ms | 否 | 否 | http |
| 172.208.25.199:3128 | ✓ 1530ms | 否 | ✓ 1248ms | ✓ 1667ms | ✓ 1216ms | http |
| 194.87.26.83:3128 | 否 | ✓ 1622ms | ✓ 1240ms | ✓ 1376ms | 否 | http |
| 77.110.116.224:3128 | ✓ 1059ms | 否 | ✓ 1227ms | ✓ 1684ms | ✓ 1240ms | http |
| 113.176.92.71:3128 | ✓ 1566ms | 否 | ✓ 1491ms | ✓ 1729ms | ✓ 1223ms | http |
| 34.101.184.164:3128 | ✓ 1742ms | 否 | 否 | ✓ 1663ms | ✓ 1552ms | http |
| 147.78.0.81:9443 | ✓ 1509ms | 否 | ✓ 1495ms | 否 | ✓ 1586ms | http |
| 38.92.10.98:20058 | ✓ 888ms | ✓ 1027ms | 否 | ✓ 1122ms | ✓ 937ms | http |
| 45.63.88.46:1080 | ✓ 1145ms | ✓ 1695ms | 否 | ✓ 1495ms | ✓ 1808ms | http |
| 152.32.132.190:7890 | ✓ 880ms | ✓ 1160ms | ✓ 910ms | ✓ 1155ms | ✓ 856ms | http |
| 108.131.109.106:51429 | ✓ 1033ms | 否 | ✓ 1726ms | 否 | ✓ 1657ms | http |
| 120.92.212.16:8890 | ✓ 1159ms | ✓ 1516ms | ✓ 1126ms | ✓ 1460ms | ✓ 1138ms | http |
| 208.87.243.199:7878 | 否 | ✓ 1828ms | ✓ 1911ms | 否 | ✓ 1733ms | http |
| 121.130.177.28:8888 | ✓ 1615ms | ✓ 1638ms | ✓ 1826ms | ✓ 1693ms | 否 | http |
| 209.126.84.232:8888 | ✓ 801ms | 否 | ✓ 1106ms | ✓ 1992ms | ✓ 1712ms | http |
| 86.104.74.110:1081 | ✓ 895ms | 否 | ✓ 602ms | 否 | ✓ 1131ms | http |
| 193.177.0.148:60000 | ✓ 1187ms | 否 | ✓ 1029ms | 否 | ✓ 1040ms | http |
| 185.41.152.110:3128 | ✓ 1165ms | ✓ 1795ms | ✓ 1756ms | 否 | 否 | http |
| 183.238.3.150:7897 | ✓ 1042ms | ✓ 1319ms | ✓ 1119ms | ✓ 1293ms | ✓ 1035ms | http |
| 183.232.248.73:7890 | ✓ 1024ms | ✓ 1385ms | ✓ 1033ms | ✓ 1353ms | ✓ 1065ms | http |
| 120.92.212.16:7890 | ✓ 1233ms | ✓ 1613ms | 否 | ✓ 1487ms | 否 | http |
| 94.158.219.111:3128 | ✓ 953ms | ✓ 1579ms | ✓ 824ms | ✓ 1722ms | ✓ 1525ms | http |
| 34.246.223.187:16206 | 否 | 否 | ✓ 1116ms | ✓ 1772ms | ✓ 1591ms | http |
| 77.110.119.136:3128 | ✓ 597ms | ✓ 1083ms | ✓ 284ms | ✓ 946ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1421ms | 否 | ✓ 1721ms | 否 | ✓ 1802ms | http |
| 51.79.207.21:8080 | ✓ 1089ms | 否 | 否 | ✓ 1865ms | ✓ 1591ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1872ms | ✓ 1910ms | ✓ 1623ms | http |
| 20.164.75.153:8080 | ✓ 1153ms | 否 | ✓ 1142ms | 否 | ✓ 1757ms | http |
| 38.92.10.139:33985 | ✓ 1152ms | ✓ 900ms | ✓ 1250ms | ✓ 1234ms | ✓ 898ms | http |
| 20.127.128.70:8080 | ✓ 1505ms | ✓ 1331ms | 否 | 否 | ✓ 950ms | http |
| 45.59.122.132:80 | ✓ 1628ms | 否 | ✓ 851ms | ✓ 1610ms | ✓ 1186ms | http |
| 101.32.243.189:80 | ✓ 1637ms | 否 | ✓ 1450ms | ✓ 1661ms | ✓ 1490ms | http |
| 138.197.68.35:4857 | ✓ 200ms | 否 | ✓ 1039ms | ✓ 1073ms | ✓ 843ms | http |

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
