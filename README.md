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

最后更新：2026-03-14 03:24:05 UTC（2026-03-14 11:24:05 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1219ms | 否 | ✓ 1717ms | ✓ 1133ms | ✓ 870ms | http |
| 86.53.183.16:1080 | ✓ 1117ms | 否 | ✓ 1730ms | ✓ 1939ms | ✓ 1266ms | http |
| 216.180.127.45:1080 | 否 | 否 | ✓ 1513ms | ✓ 1827ms | ✓ 1038ms | http |
| 120.240.35.173:22222 | ✓ 1125ms | ✓ 1545ms | 否 | ✓ 1756ms | 否 | http |
| 45.167.124.52:8080 | 否 | ✓ 1948ms | 否 | ✓ 1745ms | ✓ 1598ms | http |
| 183.249.5.105:22222 | ✓ 932ms | ✓ 1308ms | ✓ 871ms | ✓ 1218ms | ✓ 834ms | http |
| 120.92.212.16:7890 | ✓ 1056ms | ✓ 1361ms | ✓ 1059ms | ✓ 1360ms | ✓ 1057ms | http |
| 222.184.48.251:22222 | 否 | ✓ 1669ms | ✓ 1180ms | ✓ 1295ms | ✓ 1302ms | http |
| 47.77.193.180:1080 | ✓ 352ms | ✓ 1024ms | ✓ 517ms | 否 | 否 | http |
| 85.198.96.242:3128 | ✓ 1065ms | ✓ 1875ms | 否 | 否 | ✓ 1299ms | http |
| 146.56.182.165:3128 | ✓ 1039ms | 否 | 否 | ✓ 1996ms | ✓ 1304ms | http |
| 45.88.0.116:3128 | ✓ 1767ms | 否 | 否 | ✓ 1864ms | ✓ 1784ms | http |
| 45.88.0.115:3128 | ✓ 1766ms | 否 | 否 | ✓ 1866ms | ✓ 1799ms | http |
| 83.219.250.8:62920 | ✓ 841ms | 否 | ✓ 1346ms | ✓ 1941ms | ✓ 1275ms | http |
| 186.148.180.46:999 | ✓ 1926ms | ✓ 1594ms | ✓ 576ms | ✓ 1616ms | ✓ 1368ms | http |
| 120.92.212.16:8890 | ✓ 1052ms | ✓ 1374ms | 否 | ✓ 1612ms | ✓ 1366ms | http |
| 113.160.132.26:8080 | ✓ 1956ms | 否 | ✓ 1436ms | ✓ 1462ms | ✓ 1166ms | http |
| 35.225.22.61:80 | ✓ 953ms | 否 | ✓ 991ms | ✓ 1157ms | ✓ 1770ms | http |
| 183.249.5.117:22222 | ✓ 1076ms | ✓ 1337ms | ✓ 977ms | 否 | 否 | http |
| 45.88.0.98:3128 | ✓ 842ms | 否 | ✓ 1554ms | 否 | ✓ 1552ms | http |
| 171.251.173.39:5104 | ✓ 1767ms | 否 | 否 | ✓ 1484ms | ✓ 1525ms | http |
| 45.88.0.117:3128 | ✓ 424ms | ✓ 1311ms | ✓ 505ms | ✓ 1283ms | ✓ 980ms | http |
| 213.220.62.62:3128 | ✓ 420ms | 否 | ✓ 407ms | ✓ 1284ms | ✓ 995ms | http |
| 45.88.0.99:3128 | ✓ 428ms | 否 | ✓ 419ms | ✓ 1366ms | ✓ 975ms | http |
| 45.140.147.155:1081 | ✓ 483ms | ✓ 1960ms | ✓ 1220ms | 否 | ✓ 1429ms | http |
| 101.43.255.96:80 | 否 | ✓ 1492ms | ✓ 1523ms | ✓ 1460ms | 否 | http |
| 81.70.169.194:80 | ✓ 1865ms | ✓ 1409ms | 否 | ✓ 1793ms | 否 | http |
| 62.60.177.204:34094 | ✓ 565ms | ✓ 1237ms | ✓ 198ms | ✓ 964ms | ✓ 786ms | http |
| 165.227.5.10:8888 | ✓ 336ms | ✓ 1860ms | ✓ 1606ms | 否 | 否 | http |
| 117.159.239.54:22222 | 否 | ✓ 1305ms | ✓ 985ms | 否 | ✓ 951ms | http |
| 120.240.35.178:22222 | ✓ 1115ms | ✓ 1837ms | ✓ 1103ms | ✓ 1331ms | ✓ 1092ms | http |
| 160.250.4.245:1 | ✓ 1142ms | 否 | ✓ 1375ms | ✓ 1367ms | ✓ 1226ms | http |
| 160.250.5.22:1 | ✓ 1120ms | 否 | ✓ 1474ms | ✓ 1372ms | ✓ 1441ms | http |
| 117.159.239.51:22222 | ✓ 923ms | ✓ 1304ms | ✓ 1003ms | 否 | 否 | http |
| 120.198.141.75:22222 | ✓ 1328ms | ✓ 1693ms | ✓ 1130ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1107ms | ✓ 1601ms | ✓ 1417ms | ✓ 1847ms | 否 | http |
| 163.227.239.196:10000 | ✓ 1716ms | 否 | ✓ 1832ms | ✓ 1742ms | ✓ 1875ms | http |
| 120.238.159.229:22222 | ✓ 1052ms | ✓ 1376ms | ✓ 1024ms | 否 | ✓ 998ms | http |
| 37.187.109.70:10111 | ✓ 1902ms | 否 | ✓ 757ms | 否 | ✓ 1548ms | http |
| 45.88.0.111:3128 | ✓ 434ms | ✓ 1244ms | ✓ 946ms | 否 | 否 | http |
| 117.159.239.49:22222 | ✓ 1418ms | ✓ 1274ms | ✓ 1039ms | 否 | 否 | http |
| 24.144.86.173:1080 | ✓ 788ms | 否 | ✓ 1069ms | ✓ 1014ms | ✓ 761ms | http |
| 138.124.53.221:443 | ✓ 459ms | 否 | 否 | ✓ 1803ms | ✓ 1287ms | http |
| 210.223.44.230:3128 | ✓ 1876ms | ✓ 1570ms | ✓ 736ms | ✓ 1649ms | ✓ 1617ms | http |
| 91.247.126.241:2080 | ✓ 1346ms | 否 | ✓ 1881ms | ✓ 1959ms | ✓ 1713ms | http |
| 116.80.65.85:3172 | ✓ 1596ms | 否 | ✓ 1707ms | ✓ 1946ms | ✓ 1912ms | http |
| 116.80.65.77:3172 | 否 | 否 | ✓ 1592ms | ✓ 1937ms | ✓ 1808ms | http |
| 13.60.218.140:32047 | ✓ 1657ms | 否 | ✓ 1511ms | 否 | ✓ 1845ms | http |
| 62.113.119.14:8080 | ✓ 1299ms | 否 | ✓ 763ms | ✓ 1459ms | ✓ 1126ms | http |
| 45.88.0.113:3128 | 否 | 否 | ✓ 525ms | ✓ 1296ms | ✓ 1413ms | http |
| 45.140.147.82:1081 | ✓ 1350ms | 否 | ✓ 1248ms | 否 | ✓ 1971ms | http |
| 120.232.242.119:22222 | 否 | ✓ 1443ms | 否 | ✓ 1306ms | ✓ 984ms | http |
| 113.59.32.162:22222 | 否 | ✓ 1494ms | ✓ 1240ms | ✓ 1412ms | ✓ 1165ms | http |
| 221.202.27.194:10810 | 否 | ✓ 1668ms | ✓ 1198ms | 否 | ✓ 1478ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1407ms | ✓ 1489ms | 否 | ✓ 1499ms | http |
| 120.240.29.51:22222 | 否 | ✓ 1350ms | 否 | ✓ 1284ms | ✓ 1055ms | http |
| 38.145.218.82:8443 | ✓ 838ms | ✓ 1242ms | ✓ 1465ms | 否 | 否 | http |
| 103.3.246.71:3128 | ✓ 1070ms | 否 | ✓ 1143ms | ✓ 1588ms | 否 | http |
| 183.249.5.109:22222 | ✓ 912ms | ✓ 1047ms | ✓ 1050ms | ✓ 1072ms | ✓ 922ms | http |
| 152.42.213.210:443 | ✓ 1837ms | 否 | ✓ 1692ms | 否 | ✓ 1238ms | http |
| 152.42.213.210:8080 | ✓ 1860ms | 否 | ✓ 1665ms | 否 | ✓ 1680ms | http |
| 18.100.127.30:59204 | ✓ 1967ms | 否 | ✓ 1193ms | 否 | ✓ 1852ms | http |
| 45.129.141.143:3128 | ✓ 979ms | 否 | ✓ 1748ms | ✓ 1595ms | 否 | http |
| 150.230.249.50:1080 | ✓ 1852ms | 否 | ✓ 983ms | ✓ 1162ms | 否 | http |
| 38.145.203.135:8443 | ✓ 1840ms | ✓ 1899ms | ✓ 526ms | 否 | 否 | http |
| 183.249.5.214:22222 | 否 | ✓ 1244ms | ✓ 869ms | ✓ 1117ms | ✓ 955ms | http |
| 103.86.131.62:80 | ✓ 1358ms | 否 | 否 | ✓ 1575ms | ✓ 1503ms | http |
| 222.184.48.252:22222 | ✓ 953ms | ✓ 1307ms | ✓ 1055ms | ✓ 1293ms | 否 | http |
| 18.100.254.193:24160 | ✓ 1773ms | 否 | ✓ 1702ms | 否 | ✓ 1835ms | http |
| 120.198.141.84:22222 | ✓ 1048ms | ✓ 1731ms | 否 | ✓ 1345ms | 否 | http |
| 120.240.35.175:22222 | 否 | ✓ 1646ms | ✓ 1015ms | ✓ 1456ms | 否 | http |
| 121.40.231.103:7890 | ✓ 1419ms | 否 | ✓ 1515ms | ✓ 1964ms | ✓ 1602ms | http |
| 101.43.127.100:8877 | ✓ 1152ms | 否 | ✓ 1003ms | ✓ 1184ms | ✓ 988ms | http |
| 120.198.141.80:22222 | 否 | ✓ 1813ms | ✓ 1024ms | ✓ 1435ms | 否 | http |
| 43.143.108.2:9701 | ✓ 1152ms | ✓ 1170ms | ✓ 1384ms | ✓ 1595ms | ✓ 971ms | http |
| 45.136.198.40:3128 | ✓ 1841ms | 否 | ✓ 1909ms | 否 | ✓ 1670ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1999ms | ✓ 1526ms | ✓ 1451ms | http |
| 104.129.202.127:12354 | 否 | ✓ 960ms | ✓ 813ms | ✓ 1106ms | ✓ 823ms | http |
| 120.240.35.177:22222 | ✓ 1083ms | 否 | ✓ 1025ms | ✓ 1266ms | 否 | http |
| 104.129.202.127:10810 | 否 | 否 | ✓ 219ms | ✓ 951ms | ✓ 725ms | http |
| 172.212.68.37:3128 | ✓ 1274ms | ✓ 1730ms | ✓ 1468ms | ✓ 1540ms | ✓ 1050ms | http |
| 223.16.170.103:80 | ✓ 1755ms | 否 | ✓ 1503ms | 否 | ✓ 1249ms | http |
| 3.79.194.222:19873 | ✓ 1329ms | 否 | ✓ 1264ms | 否 | ✓ 1836ms | http |
| 94.72.109.169:8080 | 否 | 否 | ✓ 1568ms | ✓ 1985ms | ✓ 1648ms | http |
| 113.59.32.141:22222 | ✓ 1959ms | ✓ 1573ms | ✓ 1236ms | ✓ 1516ms | ✓ 1181ms | http |

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
