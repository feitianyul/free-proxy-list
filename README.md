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

最后更新：2026-04-13 07:07:20 UTC（2026-04-13 15:07:20 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 945ms | 否 | ✓ 871ms | ✓ 1162ms | ✓ 1034ms | http |
| 5.196.101.18:3128 | ✓ 753ms | ✓ 1433ms | ✓ 1160ms | ✓ 1728ms | ✓ 1344ms | http |
| 147.161.239.240:8800 | ✓ 576ms | 否 | ✓ 1303ms | ✓ 1802ms | ✓ 1320ms | http |
| 1.231.81.166:3128 | ✓ 851ms | ✓ 1720ms | ✓ 1936ms | ✓ 1203ms | ✓ 1102ms | http |
| 5.104.87.17:8051 | ✓ 1873ms | 否 | 否 | ✓ 1679ms | ✓ 1756ms | http |
| 113.160.132.26:8080 | ✓ 1674ms | ✓ 1895ms | ✓ 1702ms | ✓ 1462ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1699ms | 否 | ✓ 1491ms | ✓ 1961ms | 否 | http |
| 45.167.124.52:8080 | 否 | ✓ 1805ms | 否 | ✓ 1637ms | ✓ 1452ms | http |
| 34.71.229.255:3128 | ✓ 347ms | 否 | ✓ 1211ms | ✓ 1617ms | ✓ 1485ms | http |
| 167.103.115.102:8800 | ✓ 1292ms | 否 | ✓ 1184ms | ✓ 1471ms | ✓ 1171ms | http |
| 46.30.46.133:3128 | ✓ 809ms | 否 | ✓ 827ms | ✓ 1789ms | 否 | http |
| 185.76.240.61:10001 | ✓ 1034ms | 否 | ✓ 1774ms | 否 | ✓ 1870ms | http |
| 43.156.132.113:3128 | ✓ 861ms | 否 | ✓ 834ms | ✓ 1202ms | ✓ 1021ms | http |
| 167.103.144.127:8800 | ✓ 1660ms | 否 | ✓ 1082ms | ✓ 1473ms | ✓ 1328ms | http |
| 36.103.198.235:7890 | ✓ 1185ms | ✓ 1536ms | 否 | ✓ 1592ms | ✓ 1976ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1365ms | ✓ 1882ms | ✓ 1564ms | http |
| 167.103.31.122:8800 | ✓ 1677ms | 否 | ✓ 1408ms | 否 | ✓ 1659ms | http |
| 159.223.225.118:8888 | ✓ 838ms | ✓ 1549ms | 否 | ✓ 1547ms | ✓ 1505ms | http |
| 45.12.151.226:2829 | ✓ 601ms | 否 | ✓ 1221ms | 否 | ✓ 1663ms | http |
| 35.225.22.61:80 | ✓ 845ms | 否 | 否 | ✓ 1094ms | ✓ 917ms | http |
| 45.167.125.21:999 | ✓ 972ms | ✓ 1921ms | 否 | ✓ 1922ms | ✓ 1339ms | http |
| 177.234.217.88:999 | ✓ 1483ms | 否 | 否 | ✓ 1906ms | ✓ 1566ms | http |
| 109.107.179.140:31000 | ✓ 1554ms | 否 | ✓ 1552ms | 否 | ✓ 1932ms | http |
| 185.76.240.201:10001 | ✓ 1112ms | 否 | ✓ 1161ms | 否 | ✓ 1717ms | http |
| 162.240.154.26:3128 | ✓ 1317ms | 否 | 否 | ✓ 887ms | ✓ 782ms | http |
| 45.129.141.143:3128 | ✓ 731ms | ✓ 1739ms | ✓ 1807ms | 否 | ✓ 1692ms | http |
| 212.58.132.5:8888 | ✓ 1806ms | 否 | ✓ 1529ms | ✓ 1571ms | ✓ 1318ms | http |
| 185.76.241.99:10001 | ✓ 630ms | 否 | ✓ 721ms | ✓ 1890ms | ✓ 1465ms | http |
| 185.76.240.19:10001 | ✓ 795ms | ✓ 1917ms | ✓ 1296ms | ✓ 1826ms | ✓ 1407ms | http |
| 185.76.240.18:10001 | ✓ 615ms | 否 | ✓ 726ms | ✓ 1874ms | ✓ 1456ms | http |
| 185.76.240.178:10001 | ✓ 681ms | 否 | ✓ 661ms | ✓ 1879ms | ✓ 1478ms | http |
| 185.76.240.12:10001 | ✓ 592ms | 否 | ✓ 1393ms | ✓ 1846ms | ✓ 1425ms | http |
| 138.68.20.38:3128 | ✓ 1455ms | ✓ 1900ms | ✓ 783ms | ✓ 846ms | ✓ 676ms | http |
| 34.85.118.216:3128 | ✓ 1697ms | ✓ 1416ms | ✓ 1313ms | ✓ 965ms | ✓ 791ms | http |
| 160.238.65.5:3128 | ✓ 514ms | ✓ 1312ms | ✓ 1657ms | ✓ 1925ms | ✓ 1361ms | http |
| 160.238.65.6:3128 | ✓ 1484ms | ✓ 1313ms | ✓ 687ms | ✓ 1929ms | ✓ 1353ms | http |
| 160.238.65.9:3128 | ✓ 1429ms | 否 | ✓ 420ms | ✓ 1582ms | ✓ 1339ms | http |
| 160.238.65.8:3128 | ✓ 1448ms | ✓ 1945ms | ✓ 456ms | ✓ 1557ms | ✓ 1363ms | http |
| 160.238.65.7:3128 | ✓ 1504ms | ✓ 1195ms | ✓ 793ms | ✓ 1922ms | ✓ 1359ms | http |
| 160.238.65.2:3128 | ✓ 1458ms | ✓ 1198ms | ✓ 828ms | ✓ 1946ms | ✓ 1339ms | http |
| 160.238.65.4:3128 | ✓ 1460ms | 否 | ✓ 410ms | ✓ 1543ms | ✓ 1357ms | http |
| 160.238.65.3:3128 | ✓ 516ms | 否 | ✓ 970ms | ✓ 1915ms | ✓ 1371ms | http |
| 43.165.195.107:3128 | 否 | 否 | ✓ 1323ms | ✓ 1379ms | ✓ 1072ms | http |
| 110.42.37.202:20005 | ✓ 1885ms | 否 | ✓ 1392ms | 否 | ✓ 1340ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1247ms | ✓ 1545ms | ✓ 1202ms | http |
| 185.191.236.162:3128 | ✓ 1502ms | 否 | ✓ 1650ms | 否 | ✓ 1790ms | http |
| 5.255.123.43:1080 | ✓ 1726ms | 否 | 否 | ✓ 1736ms | ✓ 1324ms | http |
| 200.125.171.254:999 | ✓ 923ms | 否 | ✓ 1299ms | ✓ 1366ms | ✓ 1338ms | http |
| 185.76.240.64:10001 | ✓ 681ms | 否 | ✓ 1642ms | 否 | ✓ 1724ms | http |
| 185.76.240.203:10001 | ✓ 1663ms | ✓ 1950ms | ✓ 712ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 652ms | 否 | ✓ 581ms | 否 | ✓ 1926ms | http |
| 185.76.240.177:10001 | ✓ 706ms | 否 | ✓ 1988ms | 否 | ✓ 1809ms | http |
| 8.219.97.248:80 | ✓ 1412ms | 否 | ✓ 1325ms | ✓ 1347ms | 否 | http |
| 217.76.245.80:999 | ✓ 1731ms | 否 | 否 | ✓ 1427ms | ✓ 1652ms | http |
| 114.237.77.202:1080 | ✓ 1049ms | ✓ 1393ms | ✓ 1531ms | ✓ 1684ms | ✓ 1158ms | http |
| 101.32.244.83:8080 | ✓ 1621ms | ✓ 1987ms | ✓ 1027ms | ✓ 1617ms | ✓ 1377ms | http |
| 121.43.196.210:8222 | ✓ 1059ms | ✓ 1240ms | 否 | ✓ 1269ms | ✓ 1065ms | http |
| 121.43.196.213:8222 | ✓ 1052ms | ✓ 1208ms | 否 | ✓ 1240ms | ✓ 1232ms | http |
| 114.55.226.123:10086 | ✓ 1326ms | ✓ 1631ms | ✓ 1272ms | ✓ 1535ms | ✓ 1202ms | http |
| 178.159.94.76:3128 | ✓ 1851ms | 否 | ✓ 1606ms | 否 | ✓ 1517ms | http |
| 95.214.9.93:3128 | 否 | 否 | ✓ 1861ms | ✓ 1628ms | ✓ 1741ms | http |
| 185.76.240.215:10001 | ✓ 1926ms | 否 | ✓ 986ms | 否 | ✓ 1897ms | http |
| 185.76.241.197:10001 | ✓ 1978ms | 否 | ✓ 1046ms | 否 | ✓ 1898ms | http |
| 154.18.255.131:8080 | ✓ 1960ms | 否 | ✓ 1935ms | ✓ 1595ms | ✓ 1548ms | http |
| 185.76.240.200:10001 | ✓ 1885ms | 否 | ✓ 1863ms | 否 | ✓ 1796ms | http |
| 103.157.200.126:3128 | ✓ 1271ms | 否 | ✓ 1228ms | ✓ 1648ms | ✓ 1295ms | http |
| 185.76.240.193:10001 | ✓ 1168ms | 否 | ✓ 1562ms | 否 | ✓ 1800ms | http |
| 185.76.240.32:10001 | ✓ 1744ms | 否 | ✓ 1421ms | 否 | ✓ 1766ms | http |
| 34.231.145.203:7000 | ✓ 594ms | ✓ 1037ms | ✓ 902ms | ✓ 949ms | ✓ 813ms | http |
| 223.78.91.7:7897 | ✓ 864ms | ✓ 1080ms | ✓ 877ms | ✓ 1106ms | ✓ 895ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1224ms | ✓ 984ms | ✓ 1226ms | ✓ 994ms | http |
| 152.32.132.190:7890 | ✓ 813ms | ✓ 1588ms | ✓ 886ms | ✓ 1292ms | 否 | http |
| 181.78.44.63:999 | ✓ 908ms | 否 | ✓ 1289ms | ✓ 1800ms | ✓ 1296ms | http |
| 210.223.44.230:3128 | ✓ 1538ms | ✓ 1235ms | 否 | ✓ 1484ms | ✓ 1863ms | http |
| 101.43.127.100:8877 | ✓ 1767ms | ✓ 1299ms | ✓ 1039ms | ✓ 1184ms | ✓ 1066ms | http |
| 79.132.136.58:3128 | ✓ 1693ms | 否 | ✓ 975ms | ✓ 1286ms | ✓ 1436ms | http |
| 217.77.102.18:3128 | ✓ 1723ms | ✓ 1892ms | ✓ 1300ms | 否 | 否 | http |
| 195.158.8.123:3128 | ✓ 1888ms | 否 | ✓ 1427ms | 否 | ✓ 1978ms | http |
| 202.129.206.239:3128 | ✓ 1303ms | 否 | ✓ 1581ms | ✓ 1680ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1046ms | ✓ 1367ms | ✓ 1158ms | 否 | ✓ 1177ms | http |
| 51.79.71.106:8080 | ✓ 916ms | 否 | ✓ 1737ms | ✓ 1608ms | ✓ 1216ms | http |
| 103.113.70.189:1081 | 否 | ✓ 946ms | ✓ 1142ms | ✓ 1174ms | ✓ 817ms | http |

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
