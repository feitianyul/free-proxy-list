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

最后更新：2026-04-15 17:59:37 UTC（2026-04-16 01:59:37 UTC+8）

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
| 147.161.210.140:8800 | ✓ 802ms | 否 | ✓ 1006ms | ✓ 911ms | ✓ 920ms | http |
| 167.103.115.102:8800 | ✓ 962ms | 否 | ✓ 1094ms | ✓ 1336ms | ✓ 1259ms | http |
| 167.103.34.108:8800 | ✓ 1303ms | 否 | ✓ 1536ms | ✓ 1429ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1516ms | ✓ 1827ms | ✓ 1539ms | ✓ 1363ms | ✓ 1105ms | http |
| 202.141.161.53:10808 | 否 | ✓ 1254ms | ✓ 1431ms | ✓ 1539ms | ✓ 1013ms | http |
| 83.219.250.8:62920 | ✓ 1783ms | 否 | ✓ 866ms | 否 | ✓ 1773ms | http |
| 45.149.92.147:5001 | ✓ 669ms | 否 | ✓ 660ms | 否 | ✓ 704ms | http |
| 1.231.81.166:3128 | ✓ 1782ms | ✓ 1836ms | 否 | ✓ 1477ms | ✓ 1396ms | http |
| 59.46.216.131:30001 | ✓ 986ms | ✓ 1373ms | ✓ 1068ms | ✓ 1401ms | ✓ 1143ms | http |
| 217.76.245.80:999 | ✓ 941ms | ✓ 1799ms | 否 | ✓ 1728ms | ✓ 1249ms | http |
| 45.167.125.21:999 | ✓ 946ms | 否 | ✓ 1604ms | ✓ 1800ms | ✓ 1521ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1056ms | ✓ 1342ms | ✓ 1083ms | http |
| 167.103.144.127:8800 | ✓ 1733ms | 否 | ✓ 943ms | ✓ 1338ms | ✓ 1177ms | http |
| 147.161.239.240:8800 | ✓ 841ms | 否 | ✓ 1236ms | ✓ 1734ms | ✓ 1527ms | http |
| 157.230.178.216:8088 | 否 | 否 | ✓ 1139ms | ✓ 1573ms | ✓ 1636ms | http |
| 207.254.71.62:8088 | ✓ 797ms | ✓ 1727ms | ✓ 1765ms | ✓ 1992ms | ✓ 1911ms | http |
| 167.103.31.122:8800 | ✓ 1339ms | 否 | ✓ 1372ms | ✓ 1679ms | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1486ms | ✓ 1168ms | ✓ 1494ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1290ms | 否 | ✓ 1045ms | ✓ 1501ms | ✓ 1211ms | http |
| 185.76.241.159:10002 | ✓ 793ms | ✓ 1998ms | ✓ 1350ms | ✓ 1919ms | ✓ 1484ms | http |
| 78.11.96.22:8888 | 否 | 否 | ✓ 1482ms | ✓ 1786ms | ✓ 1455ms | http |
| 212.58.132.5:8888 | ✓ 1179ms | 否 | ✓ 1643ms | ✓ 1486ms | ✓ 1112ms | http |
| 223.84.151.86:30005 | 否 | ✓ 1956ms | ✓ 1921ms | 否 | ✓ 1748ms | http |
| 185.76.240.190:10002 | ✓ 1327ms | 否 | ✓ 688ms | ✓ 1955ms | ✓ 1474ms | http |
| 27.71.24.102:3128 | ✓ 1814ms | 否 | ✓ 1222ms | ✓ 1172ms | ✓ 1065ms | http |
| 82.114.228.67:1080 | ✓ 848ms | ✓ 1646ms | ✓ 805ms | 否 | 否 | http |
| 5.255.123.43:1080 | ✓ 1850ms | 否 | ✓ 904ms | 否 | ✓ 1531ms | http |
| 20.78.26.206:8561 | ✓ 501ms | ✓ 844ms | ✓ 610ms | ✓ 892ms | ✓ 852ms | http |
| 20.210.39.153:8561 | ✓ 503ms | ✓ 1125ms | ✓ 583ms | ✓ 921ms | ✓ 766ms | http |
| 84.47.150.125:1080 | ✓ 1343ms | 否 | ✓ 1680ms | 否 | ✓ 1687ms | http |
| 121.232.73.214:1080 | ✓ 1211ms | ✓ 1212ms | ✓ 1026ms | ✓ 1230ms | ✓ 1383ms | http |
| 45.140.147.82:1081 | ✓ 1039ms | ✓ 1369ms | ✓ 1943ms | 否 | ✓ 1362ms | http |
| 104.248.211.46:7890 | 否 | ✓ 1718ms | 否 | ✓ 1437ms | ✓ 1178ms | http |
| 8.219.97.248:80 | ✓ 1531ms | 否 | ✓ 1103ms | ✓ 1204ms | 否 | http |
| 85.239.59.252:7890 | ✓ 1155ms | ✓ 1754ms | ✓ 818ms | 否 | 否 | http |
| 107.173.42.121:7890 | ✓ 903ms | ✓ 1373ms | ✓ 854ms | ✓ 1253ms | 否 | http |
| 162.240.154.26:3128 | ✓ 773ms | ✓ 1777ms | ✓ 1132ms | ✓ 1406ms | 否 | http |
| 164.163.42.25:10000 | ✓ 1607ms | 否 | ✓ 1228ms | 否 | ✓ 1907ms | http |
| 164.163.42.27:10000 | ✓ 1402ms | 否 | ✓ 1275ms | 否 | ✓ 1865ms | http |
| 185.132.178.178:1080 | ✓ 848ms | ✓ 1832ms | ✓ 1561ms | 否 | ✓ 1303ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1217ms | ✓ 1224ms | ✓ 1249ms | http |
| 34.96.238.40:8080 | ✓ 944ms | 否 | ✓ 1520ms | ✓ 1388ms | 否 | http |
| 72.56.105.251:3128 | ✓ 794ms | 否 | 否 | ✓ 1660ms | ✓ 1490ms | http |
| 121.230.8.39:1080 | ✓ 1547ms | ✓ 1765ms | ✓ 1348ms | ✓ 1776ms | ✓ 1371ms | http |
| 150.230.249.50:1080 | 否 | ✓ 1062ms | ✓ 1963ms | ✓ 1001ms | ✓ 794ms | http |
| 138.124.99.216:8888 | 否 | ✓ 1881ms | ✓ 1808ms | 否 | ✓ 1830ms | http |
| 101.32.244.83:8080 | ✓ 999ms | 否 | ✓ 976ms | ✓ 1446ms | ✓ 1294ms | http |
| 121.43.196.213:8222 | ✓ 911ms | ✓ 1072ms | ✓ 1046ms | ✓ 1136ms | ✓ 952ms | http |
| 121.43.196.210:8222 | ✓ 958ms | ✓ 1106ms | ✓ 846ms | ✓ 1174ms | ✓ 1031ms | http |
| 114.55.226.123:10086 | ✓ 1211ms | ✓ 1492ms | ✓ 1068ms | ✓ 1367ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1148ms | ✓ 1719ms | ✓ 978ms | ✓ 1290ms | 否 | http |
| 159.223.225.118:8888 | ✓ 1069ms | 否 | ✓ 1892ms | 否 | ✓ 1430ms | http |
| 210.223.44.230:3128 | ✓ 1846ms | 否 | ✓ 1340ms | ✓ 1842ms | ✓ 1389ms | http |
| 185.212.119.154:3128 | ✓ 1366ms | 否 | ✓ 1364ms | 否 | ✓ 1858ms | http |
| 62.113.119.14:8080 | ✓ 1358ms | 否 | ✓ 1949ms | 否 | ✓ 1199ms | http |
| 147.45.214.210:1080 | 否 | ✓ 1783ms | ✓ 1688ms | 否 | ✓ 1770ms | http |
| 211.95.152.50:45046 | ✓ 1610ms | ✓ 1549ms | ✓ 1365ms | ✓ 1674ms | 否 | http |
| 195.26.224.49:3128 | ✓ 680ms | 否 | ✓ 592ms | ✓ 1579ms | ✓ 1686ms | http |
| 36.103.198.235:7890 | 否 | 否 | ✓ 1073ms | ✓ 1408ms | ✓ 1142ms | http |
| 120.92.108.86:7890 | ✓ 1612ms | 否 | ✓ 1618ms | 否 | ✓ 1918ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 848ms | ✓ 1640ms | ✓ 965ms | http |
| 144.31.27.49:1080 | ✓ 1668ms | ✓ 1981ms | ✓ 1016ms | 否 | ✓ 1267ms | http |
| 103.163.80.25:8081 | ✓ 1225ms | 否 | ✓ 1217ms | 否 | ✓ 1232ms | http |
| 45.12.151.226:2829 | ✓ 1641ms | 否 | ✓ 632ms | ✓ 1876ms | 否 | http |
| 46.39.105.157:8080 | ✓ 694ms | ✓ 1971ms | ✓ 685ms | ✓ 1491ms | ✓ 1329ms | http |
| 140.238.242.189:8100 | ✓ 1964ms | 否 | ✓ 1413ms | 否 | ✓ 1528ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1399ms | ✓ 1687ms | ✓ 1377ms | http |
| 43.252.238.218:8080 | 否 | 否 | ✓ 1447ms | ✓ 1440ms | ✓ 1387ms | http |
| 37.187.109.70:10111 | ✓ 1052ms | ✓ 1380ms | ✓ 1157ms | 否 | ✓ 1960ms | http |
| 94.131.118.39:1081 | ✓ 956ms | 否 | ✓ 590ms | ✓ 1855ms | ✓ 1355ms | http |
| 200.125.171.254:999 | ✓ 1575ms | ✓ 1694ms | ✓ 1400ms | ✓ 1537ms | ✓ 1232ms | http |
| 12.89.176.82:3128 | ✓ 587ms | ✓ 1194ms | ✓ 1177ms | ✓ 1251ms | ✓ 1221ms | http |
| 121.230.9.113:1080 | ✓ 1161ms | 否 | ✓ 1079ms | 否 | ✓ 1165ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1239ms | ✓ 920ms | ✓ 1828ms | ✓ 1013ms | http |
| 117.84.29.218:1080 | ✓ 1341ms | 否 | ✓ 1349ms | ✓ 1709ms | 否 | http |
| 103.113.70.189:1081 | ✓ 549ms | ✓ 1350ms | ✓ 968ms | ✓ 1157ms | ✓ 888ms | http |
| 181.78.44.63:999 | ✓ 1088ms | ✓ 1800ms | ✓ 1326ms | 否 | 否 | http |
| 34.71.229.255:3128 | ✓ 970ms | ✓ 1441ms | ✓ 964ms | ✓ 1380ms | ✓ 1158ms | http |
| 20.127.128.70:8080 | ✓ 1531ms | ✓ 1848ms | ✓ 748ms | ✓ 1982ms | ✓ 1574ms | http |
| 152.32.132.190:7890 | ✓ 1366ms | ✓ 1401ms | ✓ 927ms | 否 | ✓ 981ms | http |
| 103.39.51.207:8080 | ✓ 1287ms | 否 | 否 | ✓ 1840ms | ✓ 1524ms | http |
| 47.93.216.160:1081 | ✓ 1126ms | ✓ 1830ms | ✓ 902ms | 否 | 否 | http |

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
