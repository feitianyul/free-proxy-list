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

最后更新：2026-03-21 15:35:58 UTC（2026-03-21 23:35:58 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.27.13.35:8561 | 否 | ✓ 874ms | ✓ 656ms | ✓ 850ms | ✓ 611ms | http |
| 20.27.15.111:8561 | 否 | ✓ 1162ms | ✓ 476ms | ✓ 762ms | ✓ 587ms | http |
| 38.34.179.27:8451 | ✓ 1227ms | ✓ 1923ms | ✓ 332ms | ✓ 1146ms | ✓ 1352ms | http |
| 210.76.193.248:10808 | ✓ 1062ms | ✓ 1083ms | ✓ 1039ms | ✓ 1210ms | ✓ 971ms | http |
| 147.161.210.140:8800 | 否 | 否 | ✓ 880ms | ✓ 951ms | ✓ 869ms | http |
| 113.160.132.26:8080 | ✓ 1535ms | 否 | ✓ 1014ms | ✓ 1478ms | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1904ms | ✓ 1302ms | ✓ 1188ms | http |
| 167.103.34.108:8800 | ✓ 1839ms | 否 | ✓ 1438ms | 否 | ✓ 1739ms | http |
| 45.136.130.178:8449 | ✓ 573ms | 否 | ✓ 235ms | ✓ 864ms | ✓ 1746ms | http |
| 45.136.130.198:8449 | ✓ 818ms | 否 | ✓ 154ms | ✓ 890ms | ✓ 1116ms | http |
| 45.136.130.188:8449 | ✓ 811ms | 否 | ✓ 155ms | ✓ 917ms | ✓ 1121ms | http |
| 38.34.183.8:8450 | ✓ 592ms | ✓ 1121ms | ✓ 967ms | ✓ 928ms | ✓ 718ms | http |
| 219.117.204.211:7799 | ✓ 890ms | 否 | ✓ 937ms | ✓ 1022ms | ✓ 751ms | http |
| 38.34.179.75:8453 | ✓ 245ms | ✓ 1270ms | ✓ 1827ms | ✓ 1927ms | ✓ 1074ms | http |
| 137.220.150.22:6005 | ✓ 803ms | 否 | ✓ 980ms | ✓ 1258ms | ✓ 1092ms | http |
| 85.208.108.43:2094 | ✓ 1484ms | 否 | ✓ 1699ms | ✓ 1323ms | ✓ 939ms | http |
| 38.34.179.86:8452 | ✓ 325ms | ✓ 1666ms | ✓ 1766ms | 否 | 否 | http |
| 103.113.70.189:1081 | 否 | ✓ 1180ms | 否 | ✓ 1374ms | ✓ 1915ms | http |
| 91.233.223.147:3128 | ✓ 1453ms | 否 | ✓ 1110ms | 否 | ✓ 1670ms | http |
| 115.231.181.40:8128 | ✓ 861ms | 否 | ✓ 1536ms | ✓ 1994ms | 否 | http |
| 137.184.1.87:3128 | ✓ 1084ms | 否 | ✓ 1490ms | ✓ 1690ms | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 422ms | ✓ 1365ms | ✓ 952ms | http |
| 223.16.170.103:80 | ✓ 1456ms | 否 | ✓ 1223ms | ✓ 1008ms | 否 | http |
| 202.38.72.235:26001 | 否 | ✓ 1591ms | ✓ 1433ms | ✓ 1974ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1353ms | 否 | ✓ 1421ms | ✓ 1572ms | ✓ 1494ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1384ms | ✓ 1176ms | 否 | ✓ 926ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1430ms | ✓ 1123ms | 否 | ✓ 850ms | http |
| 120.92.212.16:8890 | ✓ 932ms | ✓ 1171ms | 否 | ✓ 1423ms | 否 | http |
| 139.159.99.242:8080 | ✓ 852ms | ✓ 1001ms | ✓ 1327ms | 否 | 否 | http |
| 142.171.224.229:7890 | 否 | ✓ 1571ms | ✓ 557ms | ✓ 757ms | ✓ 522ms | http |
| 101.43.127.100:8877 | ✓ 1397ms | 否 | ✓ 830ms | ✓ 1167ms | ✓ 1187ms | http |
| 137.220.150.170:6005 | ✓ 1904ms | 否 | ✓ 1257ms | ✓ 1867ms | ✓ 1907ms | http |
| 49.156.44.114:8080 | ✓ 1984ms | 否 | ✓ 1559ms | ✓ 1348ms | 否 | http |
| 45.136.131.59:8450 | ✓ 1257ms | 否 | ✓ 1643ms | ✓ 1110ms | ✓ 1632ms | http |
| 133.242.138.34:8100 | ✓ 1545ms | ✓ 1165ms | 否 | ✓ 1175ms | ✓ 902ms | http |
| 38.34.179.182:8446 | ✓ 395ms | ✓ 653ms | ✓ 93ms | ✓ 942ms | ✓ 677ms | http |
| 24.199.124.152:3128 | ✓ 226ms | ✓ 1327ms | ✓ 541ms | ✓ 627ms | ✓ 501ms | http |
| 106.14.203.63:3333 | ✓ 1629ms | ✓ 1468ms | ✓ 1673ms | ✓ 1870ms | ✓ 848ms | http |
| 24.199.124.151:3128 | ✓ 1214ms | ✓ 622ms | ✓ 223ms | ✓ 649ms | ✓ 494ms | http |
| 167.71.196.28:8080 | ✓ 684ms | 否 | ✓ 971ms | 否 | ✓ 795ms | http |
| 45.149.92.147:5001 | ✓ 794ms | 否 | ✓ 596ms | ✓ 757ms | ✓ 634ms | http |
| 147.161.239.240:8800 | ✓ 1261ms | 否 | ✓ 1313ms | ✓ 1468ms | ✓ 1282ms | http |
| 45.136.130.171:8445 | ✓ 820ms | ✓ 844ms | ✓ 812ms | 否 | 否 | http |
| 222.228.171.92:8080 | ✓ 1307ms | 否 | 否 | ✓ 1283ms | ✓ 1695ms | http |
| 38.34.179.150:8449 | 否 | 否 | ✓ 692ms | ✓ 1926ms | ✓ 628ms | http |
| 20.27.14.220:8561 | ✓ 1294ms | ✓ 1420ms | ✓ 634ms | ✓ 812ms | ✓ 837ms | http |
| 20.27.11.248:8561 | ✓ 1226ms | ✓ 1188ms | ✓ 442ms | ✓ 762ms | ✓ 661ms | http |
| 137.220.150.104:6005 | ✓ 1361ms | 否 | ✓ 771ms | ✓ 1127ms | ✓ 1356ms | http |
| 116.80.65.79:3172 | ✓ 1601ms | 否 | ✓ 1501ms | 否 | ✓ 1607ms | http |
| 47.74.226.8:5001 | ✓ 1564ms | 否 | ✓ 1034ms | ✓ 1217ms | 否 | http |
| 91.238.105.64:2024 | ✓ 1072ms | 否 | ✓ 1778ms | 否 | ✓ 1957ms | http |
| 23.148.244.206:20103 | ✓ 646ms | 否 | ✓ 239ms | 否 | ✓ 677ms | http |
| 137.220.151.110:6005 | ✓ 840ms | 否 | ✓ 712ms | ✓ 1012ms | ✓ 805ms | http |
| 38.34.183.47:8452 | ✓ 360ms | ✓ 1209ms | ✓ 660ms | ✓ 962ms | ✓ 522ms | http |
| 38.34.178.186:8451 | ✓ 512ms | ✓ 1318ms | ✓ 270ms | ✓ 673ms | ✓ 639ms | http |
| 45.136.131.35:8452 | ✓ 918ms | ✓ 1405ms | ✓ 807ms | ✓ 1132ms | ✓ 1506ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 1331ms | ✓ 1980ms | ✓ 1314ms | http |
| 38.34.178.245:8446 | ✓ 1000ms | 否 | ✓ 1315ms | ✓ 1972ms | ✓ 1380ms | http |
| 85.208.108.43:10808 | ✓ 1252ms | 否 | ✓ 1256ms | ✓ 1318ms | ✓ 961ms | http |
| 38.34.179.186:8444 | ✓ 453ms | ✓ 1786ms | ✓ 260ms | ✓ 1070ms | ✓ 806ms | http |
| 128.199.121.61:9090 | ✓ 777ms | 否 | 否 | ✓ 1234ms | ✓ 1353ms | http |
| 45.136.130.174:8450 | ✓ 733ms | 否 | ✓ 202ms | ✓ 1900ms | 否 | http |
| 172.212.68.37:3128 | ✓ 847ms | 否 | ✓ 1632ms | ✓ 1860ms | ✓ 1610ms | http |
| 38.34.179.191:8453 | ✓ 441ms | ✓ 1712ms | ✓ 177ms | ✓ 1263ms | ✓ 809ms | http |
| 61.76.95.217:40088 | ✓ 1143ms | 否 | ✓ 1391ms | ✓ 1745ms | ✓ 1204ms | http |
| 222.109.119.178:3128 | 否 | 否 | ✓ 999ms | ✓ 1169ms | ✓ 752ms | http |
| 194.67.99.223:1080 | ✓ 1259ms | 否 | ✓ 975ms | ✓ 1840ms | ✓ 1287ms | http |
| 166.88.55.83:7890 | ✓ 602ms | ✓ 1037ms | ✓ 607ms | ✓ 754ms | ✓ 596ms | http |
| 59.46.216.131:30001 | ✓ 906ms | 否 | 否 | ✓ 1964ms | ✓ 966ms | http |
| 38.34.179.17:8446 | ✓ 899ms | ✓ 1961ms | ✓ 332ms | ✓ 1344ms | ✓ 929ms | http |
| 104.168.158.236:10808 | ✓ 556ms | 否 | ✓ 1268ms | 否 | ✓ 1092ms | http |
| 45.136.198.40:3128 | ✓ 1282ms | 否 | ✓ 1366ms | 否 | ✓ 1573ms | http |
| 137.184.1.155:3128 | 否 | 否 | ✓ 214ms | ✓ 639ms | ✓ 480ms | http |
| 45.136.131.28:8445 | ✓ 1883ms | 否 | ✓ 901ms | ✓ 1192ms | 否 | http |
| 137.184.6.37:3128 | 否 | ✓ 580ms | ✓ 906ms | ✓ 622ms | ✓ 481ms | http |
| 143.244.196.234:3128 | ✓ 1251ms | ✓ 1882ms | 否 | ✓ 1855ms | ✓ 1562ms | http |
| 38.34.178.154:8445 | ✓ 675ms | ✓ 617ms | ✓ 385ms | ✓ 1698ms | 否 | http |
| 8.219.97.248:80 | ✓ 1547ms | 否 | ✓ 1285ms | 否 | ✓ 1198ms | http |
| 103.39.51.190:8080 | ✓ 1649ms | 否 | ✓ 1863ms | ✓ 1887ms | ✓ 1473ms | http |
| 45.186.6.104:3128 | ✓ 1492ms | ✓ 1764ms | ✓ 1776ms | 否 | 否 | http |
| 147.45.67.148:8080 | ✓ 974ms | 否 | 否 | ✓ 1814ms | ✓ 1469ms | http |
| 106.75.15.167:7890 | 否 | ✓ 1627ms | 否 | ✓ 1222ms | ✓ 901ms | http |
| 45.136.131.54:8448 | ✓ 626ms | 否 | ✓ 709ms | 否 | ✓ 1755ms | http |
| 45.136.131.53:8452 | ✓ 629ms | ✓ 1724ms | ✓ 417ms | 否 | 否 | http |
| 45.136.130.177:8448 | ✓ 673ms | 否 | ✓ 371ms | ✓ 1300ms | ✓ 1440ms | http |
| 38.145.208.175:8451 | ✓ 762ms | ✓ 1993ms | ✓ 301ms | ✓ 1639ms | ✓ 954ms | http |
| 45.93.29.147:6005 | 否 | ✓ 1887ms | 否 | ✓ 1860ms | ✓ 1930ms | http |
| 121.230.9.26:1080 | 否 | 否 | ✓ 1684ms | ✓ 1646ms | ✓ 1071ms | http |
| 150.241.77.172:1080 | ✓ 1041ms | 否 | ✓ 1641ms | ✓ 1786ms | ✓ 1289ms | http |
| 114.237.77.231:1080 | 否 | ✓ 1761ms | ✓ 1953ms | 否 | ✓ 1592ms | http |
| 1.231.81.166:3128 | ✓ 889ms | ✓ 1492ms | ✓ 1282ms | ✓ 881ms | ✓ 690ms | http |
| 103.155.197.103:8080 | 否 | 否 | ✓ 1394ms | ✓ 1322ms | ✓ 1662ms | http |

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
