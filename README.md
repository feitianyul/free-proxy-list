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

最后更新：2026-03-15 19:41:27 UTC（2026-03-16 03:41:27 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 607ms | 否 | ✓ 1107ms | 否 | ✓ 793ms | http |
| 149.50.116.240:1080 | ✓ 1307ms | 否 | ✓ 1286ms | ✓ 1764ms | ✓ 1285ms | http |
| 137.220.151.110:6005 | ✓ 847ms | 否 | ✓ 922ms | ✓ 1962ms | ✓ 1263ms | http |
| 137.220.150.104:6005 | ✓ 1311ms | 否 | ✓ 1708ms | ✓ 1262ms | ✓ 974ms | http |
| 137.220.150.152:6005 | ✓ 1055ms | 否 | ✓ 1022ms | ✓ 1331ms | ✓ 1048ms | http |
| 113.160.132.26:8080 | ✓ 1827ms | ✓ 1734ms | ✓ 1500ms | ✓ 1987ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 787ms | ✓ 959ms | ✓ 1356ms | http |
| 14.225.212.37:7890 | ✓ 1488ms | 否 | ✓ 927ms | 否 | ✓ 956ms | http |
| 45.167.124.52:8080 | ✓ 1183ms | 否 | 否 | ✓ 1576ms | ✓ 1369ms | http |
| 2.56.122.146:10808 | ✓ 426ms | ✓ 1905ms | ✓ 867ms | 否 | 否 | http |
| 101.43.127.100:8877 | ✓ 963ms | ✓ 1251ms | ✓ 1010ms | ✓ 1161ms | ✓ 938ms | http |
| 115.231.181.40:8128 | ✓ 1024ms | ✓ 1221ms | ✓ 1065ms | ✓ 1372ms | ✓ 929ms | http |
| 81.70.169.194:80 | ✓ 1144ms | ✓ 1434ms | ✓ 1139ms | ✓ 1447ms | ✓ 1185ms | http |
| 160.238.65.3:3128 | ✓ 904ms | ✓ 1180ms | ✓ 1824ms | ✓ 1766ms | ✓ 1807ms | http |
| 160.238.65.8:3128 | ✓ 902ms | ✓ 1220ms | ✓ 1785ms | ✓ 1776ms | ✓ 1772ms | http |
| 160.238.65.4:3128 | ✓ 900ms | ✓ 1332ms | ✓ 1679ms | ✓ 1794ms | ✓ 1768ms | http |
| 38.145.218.82:8443 | ✓ 637ms | ✓ 1056ms | ✓ 532ms | ✓ 1084ms | ✓ 890ms | http |
| 160.238.65.2:3128 | ✓ 1079ms | 否 | ✓ 429ms | ✓ 1262ms | ✓ 982ms | http |
| 120.92.212.16:7890 | ✓ 1029ms | ✓ 1345ms | ✓ 1067ms | ✓ 1350ms | ✓ 1055ms | http |
| 160.238.65.7:3128 | 否 | ✓ 1501ms | ✓ 1623ms | ✓ 1633ms | ✓ 988ms | http |
| 120.92.212.16:8890 | ✓ 1320ms | ✓ 1336ms | ✓ 1031ms | ✓ 1584ms | ✓ 1307ms | http |
| 45.119.85.216:3128 | ✓ 1666ms | 否 | ✓ 1895ms | 否 | ✓ 1322ms | http |
| 35.225.22.61:80 | ✓ 922ms | ✓ 1184ms | ✓ 939ms | 否 | 否 | http |
| 160.238.65.5:3128 | ✓ 928ms | ✓ 1345ms | ✓ 497ms | ✓ 1341ms | ✓ 1000ms | http |
| 160.238.65.9:3128 | ✓ 941ms | ✓ 1334ms | ✓ 510ms | ✓ 1338ms | ✓ 983ms | http |
| 160.238.65.6:3128 | ✓ 937ms | ✓ 1797ms | ✓ 462ms | ✓ 1313ms | ✓ 988ms | http |
| 101.43.255.96:80 | ✓ 1149ms | ✓ 1480ms | ✓ 1023ms | ✓ 1359ms | ✓ 1148ms | http |
| 57.128.188.167:9158 | ✓ 1581ms | 否 | ✓ 1605ms | 否 | ✓ 1607ms | http |
| 38.145.203.135:8443 | ✓ 734ms | ✓ 1509ms | ✓ 777ms | ✓ 995ms | 否 | http |
| 85.198.96.242:3128 | ✓ 614ms | 否 | ✓ 945ms | ✓ 1641ms | 否 | http |
| 47.105.98.23:3128 | ✓ 969ms | ✓ 1284ms | 否 | 否 | ✓ 1054ms | http |
| 143.244.140.119:3128 | ✓ 1160ms | 否 | 否 | ✓ 1607ms | ✓ 1556ms | http |
| 59.46.216.131:30001 | ✓ 1281ms | 否 | 否 | ✓ 1553ms | ✓ 1225ms | http |
| 116.80.96.107:3172 | ✓ 1662ms | 否 | ✓ 1665ms | ✓ 1944ms | 否 | http |
| 92.119.127.212:6005 | ✓ 1217ms | 否 | ✓ 1341ms | ✓ 1685ms | ✓ 1523ms | http |
| 138.124.81.12:8888 | ✓ 1516ms | 否 | ✓ 1651ms | 否 | ✓ 1773ms | http |
| 72.11.150.178:6005 | 否 | ✓ 1130ms | ✓ 1027ms | ✓ 1142ms | ✓ 905ms | http |
| 104.129.202.127:10810 | ✓ 1101ms | ✓ 1216ms | ✓ 1234ms | ✓ 1380ms | ✓ 1300ms | http |
| 199.68.217.2:3128 | 否 | 否 | ✓ 1513ms | ✓ 1035ms | ✓ 976ms | http |
| 104.129.202.127:12354 | ✓ 1365ms | 否 | 否 | ✓ 1264ms | ✓ 1813ms | http |
| 121.230.9.19:1080 | ✓ 1462ms | ✓ 1695ms | ✓ 1690ms | ✓ 1754ms | ✓ 1286ms | http |
| 121.230.9.33:1080 | ✓ 1257ms | ✓ 1919ms | ✓ 1259ms | 否 | ✓ 1319ms | http |
| 121.232.73.214:1080 | 否 | ✓ 1731ms | ✓ 1482ms | ✓ 1405ms | ✓ 1342ms | http |
| 178.236.245.17:3128 | ✓ 1056ms | 否 | ✓ 735ms | ✓ 1709ms | ✓ 1295ms | http |
| 45.136.198.40:3128 | ✓ 895ms | 否 | ✓ 1018ms | ✓ 1754ms | ✓ 1323ms | http |
| 106.117.208.101:7890 | ✓ 1405ms | ✓ 1375ms | ✓ 1212ms | ✓ 1445ms | ✓ 1089ms | http |
| 45.207.200.120:1080 | ✓ 1517ms | 否 | ✓ 1511ms | 否 | ✓ 792ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1895ms | ✓ 1399ms | ✓ 1361ms | ✓ 758ms | http |
| 83.219.250.8:62920 | ✓ 624ms | ✓ 1432ms | ✓ 968ms | ✓ 1512ms | 否 | http |
| 95.3.9.78:8080 | ✓ 712ms | ✓ 1857ms | ✓ 1926ms | ✓ 1598ms | ✓ 1230ms | http |
| 95.3.9.78:3128 | ✓ 721ms | ✓ 1723ms | 否 | ✓ 1660ms | ✓ 1264ms | http |
| 200.174.198.32:8888 | ✓ 1904ms | 否 | ✓ 1850ms | 否 | ✓ 1720ms | http |
| 172.212.68.37:3128 | ✓ 943ms | ✓ 1545ms | ✓ 729ms | ✓ 1321ms | ✓ 767ms | http |
| 144.31.137.23:8080 | ✓ 1719ms | ✓ 1759ms | 否 | 否 | ✓ 1603ms | http |
| 62.234.206.73:3128 | ✓ 1073ms | ✓ 1330ms | ✓ 959ms | ✓ 1401ms | ✓ 1039ms | http |
| 45.149.92.147:5001 | ✓ 753ms | 否 | ✓ 930ms | ✓ 1633ms | ✓ 754ms | http |
| 207.244.244.178:3128 | ✓ 216ms | ✓ 1558ms | ✓ 392ms | ✓ 864ms | ✓ 707ms | http |
| 45.136.130.223:8443 | ✓ 1278ms | ✓ 879ms | ✓ 306ms | ✓ 970ms | ✓ 855ms | http |
| 164.90.155.209:3128 | 否 | 否 | ✓ 1300ms | ✓ 1323ms | ✓ 675ms | http |
| 103.39.51.190:8080 | ✓ 1815ms | 否 | 否 | ✓ 1536ms | ✓ 1529ms | http |
| 210.223.44.230:3128 | ✓ 1584ms | ✓ 1051ms | 否 | 否 | ✓ 1842ms | http |
| 192.71.213.85:9091 | ✓ 837ms | 否 | ✓ 675ms | ✓ 1962ms | 否 | http |
| 192.71.213.85:9812 | ✓ 479ms | 否 | ✓ 502ms | ✓ 1385ms | 否 | http |
| 194.5.212.40:8080 | ✓ 914ms | ✓ 1733ms | ✓ 1645ms | 否 | ✓ 1677ms | http |
| 38.145.208.138:8447 | ✓ 685ms | ✓ 987ms | ✓ 825ms | ✓ 1020ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1081ms | ✓ 1260ms | ✓ 997ms | ✓ 1270ms | ✓ 1025ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1090ms | ✓ 1502ms | ✓ 1341ms | http |
| 113.255.59.226:8080 | ✓ 1336ms | 否 | ✓ 1283ms | 否 | ✓ 1306ms | http |
| 88.80.150.82:8080 | ✓ 987ms | ✓ 1870ms | ✓ 1532ms | 否 | 否 | https |
| 113.176.92.71:3128 | 否 | ✓ 1418ms | ✓ 1632ms | ✓ 1673ms | ✓ 1041ms | http |

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
