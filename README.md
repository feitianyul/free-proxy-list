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

最后更新：2026-05-19 17:12:52 UTC（2026-05-20 01:12:52 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 174.137.134.182:2999 | ✓ 557ms | ✓ 1366ms | 否 | ✓ 1564ms | ✓ 1810ms | http |
| 192.99.8.15:8850 | ✓ 438ms | 否 | 否 | ✓ 1683ms | ✓ 1233ms | http |
| 1.231.81.166:3128 | ✓ 1872ms | 否 | ✓ 1417ms | ✓ 1250ms | ✓ 1009ms | http |
| 113.160.132.26:8080 | ✓ 1561ms | ✓ 1679ms | 否 | ✓ 1513ms | ✓ 1215ms | http |
| 176.111.37.216:39811 | ✓ 1245ms | ✓ 1950ms | ✓ 1978ms | 否 | ✓ 1823ms | http |
| 185.200.188.234:10001 | ✓ 1793ms | 否 | ✓ 1405ms | 否 | ✓ 1940ms | http |
| 47.74.226.8:5001 | ✓ 1126ms | 否 | ✓ 1145ms | ✓ 1566ms | 否 | http |
| 176.111.37.5:39811 | ✓ 667ms | ✓ 1999ms | ✓ 626ms | 否 | ✓ 1331ms | http |
| 59.46.216.131:30001 | ✓ 1119ms | ✓ 1881ms | 否 | 否 | ✓ 1134ms | http |
| 168.110.52.228:3128 | 否 | ✓ 1931ms | ✓ 628ms | ✓ 944ms | ✓ 809ms | http |
| 152.67.191.232:6800 | ✓ 1175ms | 否 | ✓ 1244ms | ✓ 1679ms | ✓ 1237ms | http |
| 34.101.184.164:3128 | ✓ 1725ms | 否 | ✓ 952ms | ✓ 1708ms | ✓ 1245ms | http |
| 8.210.161.8:8100 | ✓ 1324ms | 否 | ✓ 1917ms | 否 | ✓ 1819ms | http |
| 138.2.92.70:8100 | ✓ 1482ms | 否 | ✓ 1720ms | 否 | ✓ 1961ms | http |
| 89.58.50.94:11140 | ✓ 923ms | ✓ 1782ms | ✓ 1643ms | 否 | 否 | http |
| 43.130.126.146:6688 | ✓ 716ms | 否 | ✓ 903ms | 否 | ✓ 1892ms | http |
| 8.218.153.104:8100 | ✓ 1238ms | 否 | 否 | ✓ 1287ms | ✓ 1828ms | http |
| 8.154.21.175:3128 | ✓ 1120ms | ✓ 1817ms | ✓ 961ms | ✓ 1672ms | ✓ 1444ms | http |
| 152.228.163.79:80 | ✓ 1803ms | 否 | ✓ 1396ms | 否 | ✓ 1896ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1911ms | 否 | ✓ 1798ms | ✓ 1765ms | http |
| 45.125.67.37:8443 | ✓ 1503ms | 否 | 否 | ✓ 1941ms | ✓ 1873ms | http |
| 74.208.192.81:3129 | ✓ 607ms | 否 | ✓ 631ms | 否 | ✓ 1131ms | http |
| 86.104.72.219:1081 | ✓ 230ms | ✓ 1956ms | ✓ 1270ms | ✓ 1274ms | ✓ 810ms | http |
| 170.106.136.181:31002 | ✓ 1764ms | 否 | ✓ 1349ms | ✓ 1773ms | ✓ 842ms | http |
| 202.28.194.139:31280 | ✓ 1673ms | 否 | ✓ 1916ms | ✓ 1897ms | ✓ 1978ms | http |
| 5.252.33.13:2025 | ✓ 1410ms | 否 | ✓ 1993ms | 否 | ✓ 1889ms | http |
| 138.2.239.213:10010 | ✓ 1959ms | ✓ 1995ms | ✓ 1904ms | ✓ 1501ms | ✓ 995ms | http |
| 104.168.96.172:1888 | ✓ 558ms | ✓ 1724ms | ✓ 1293ms | 否 | ✓ 1751ms | http |
| 223.16.170.103:3128 | ✓ 1354ms | ✓ 1818ms | ✓ 1456ms | ✓ 1158ms | ✓ 1192ms | http |
| 34.87.80.221:30000 | ✓ 1528ms | 否 | ✓ 1522ms | ✓ 1213ms | ✓ 984ms | http |
| 138.2.78.251:8100 | 否 | 否 | ✓ 1199ms | ✓ 1625ms | ✓ 1655ms | http |
| 103.157.117.226:81 | ✓ 1754ms | 否 | 否 | ✓ 1383ms | ✓ 1449ms | http |
| 116.254.118.180:80 | ✓ 1683ms | 否 | ✓ 1110ms | ✓ 1338ms | ✓ 1105ms | http |
| 81.30.156.115:8080 | ✓ 1090ms | 否 | ✓ 1217ms | ✓ 1900ms | ✓ 1792ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1239ms | ✓ 1304ms | ✓ 1036ms | http |
| 161.117.86.53:8100 | 否 | 否 | ✓ 1852ms | ✓ 1332ms | ✓ 1168ms | http |
| 3.101.133.120:80 | ✓ 1095ms | ✓ 1435ms | ✓ 1341ms | ✓ 1188ms | ✓ 1023ms | http |
| 38.224.172.232:999 | ✓ 1285ms | ✓ 1995ms | ✓ 1543ms | ✓ 1947ms | ✓ 1896ms | http |
| 86.104.72.220:1081 | 否 | ✓ 1048ms | ✓ 1411ms | ✓ 1684ms | 否 | http |
| 64.188.77.221:3128 | ✓ 1239ms | 否 | ✓ 1158ms | 否 | ✓ 1502ms | http |
| 114.214.163.108:6789 | ✓ 1045ms | ✓ 1303ms | ✓ 1277ms | ✓ 1290ms | ✓ 1058ms | http |
| 192.81.129.252:3136 | 否 | 否 | ✓ 1992ms | ✓ 1070ms | ✓ 689ms | http |
| 45.174.92.44:8087 | ✓ 1640ms | ✓ 1841ms | ✓ 1170ms | 否 | ✓ 1160ms | http |
| 152.32.132.190:7890 | ✓ 1978ms | 否 | 否 | ✓ 1656ms | ✓ 943ms | http |
| 173.212.245.136:8888 | ✓ 1598ms | 否 | ✓ 1908ms | 否 | ✓ 1693ms | http |
| 128.199.116.219:9090 | ✓ 1980ms | 否 | ✓ 1822ms | ✓ 1485ms | 否 | http |
| 85.192.29.60:3128 | 否 | 否 | ✓ 979ms | ✓ 1704ms | ✓ 1964ms | http |
| 84.47.150.125:1080 | ✓ 1383ms | 否 | 否 | ✓ 1850ms | ✓ 1708ms | http |
| 218.108.131.186:17890 | ✓ 1919ms | ✓ 1570ms | 否 | 否 | ✓ 1664ms | http |
| 103.81.195.222:3125 | ✓ 1903ms | 否 | ✓ 1351ms | ✓ 1550ms | ✓ 1705ms | http |
| 8.218.174.172:8100 | 否 | 否 | ✓ 1276ms | ✓ 1972ms | ✓ 1312ms | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 869ms | ✓ 1202ms | ✓ 951ms | http |
| 5.102.109.41:999 | ✓ 694ms | ✓ 1322ms | ✓ 1689ms | 否 | 否 | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1319ms | ✓ 1670ms | ✓ 1329ms | http |
| 8.210.138.49:8100 | ✓ 1357ms | ✓ 1869ms | ✓ 1741ms | 否 | ✓ 1191ms | http |
| 207.254.71.62:8088 | ✓ 599ms | 否 | ✓ 633ms | ✓ 1591ms | ✓ 1993ms | http |
| 2.27.32.81:3128 | ✓ 1758ms | 否 | ✓ 1504ms | 否 | ✓ 1952ms | http |
| 129.80.217.21:444 | 否 | ✓ 971ms | ✓ 1032ms | ✓ 1986ms | ✓ 787ms | http |
| 129.80.238.83:444 | ✓ 1338ms | 否 | 否 | ✓ 1013ms | ✓ 754ms | http |
| 64.188.77.26:3128 | ✓ 689ms | ✓ 1585ms | ✓ 738ms | ✓ 1770ms | 否 | http |
| 128.199.121.61:9090 | 否 | 否 | ✓ 1608ms | ✓ 1669ms | ✓ 1153ms | http |
| 20.210.76.175:8561 | 否 | ✓ 1155ms | ✓ 1298ms | ✓ 1933ms | ✓ 907ms | http |
| 61.52.131.172:8443 | ✓ 986ms | ✓ 1200ms | 否 | ✓ 1321ms | ✓ 1040ms | http |
| 120.92.212.16:7890 | ✓ 1787ms | 否 | ✓ 1025ms | ✓ 1908ms | 否 | http |
| 114.214.165.78:10810 | ✓ 1143ms | 否 | ✓ 1169ms | ✓ 1340ms | ✓ 1137ms | http |
| 121.230.8.136:1080 | ✓ 1401ms | 否 | ✓ 1765ms | ✓ 1570ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1049ms | 否 | ✓ 1997ms | 否 | ✓ 1330ms | http |
| 157.0.142.246:10057 | ✓ 1632ms | ✓ 1378ms | ✓ 1144ms | 否 | 否 | http |
| 147.45.78.89:1080 | ✓ 1499ms | 否 | ✓ 1285ms | 否 | ✓ 1557ms | http |

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
