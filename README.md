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

最后更新：2026-02-28 23:22:15 UTC（2026-03-01 07:22:15 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 148.135.85.87:1080 | ✓ 819ms | ✓ 1325ms | ✓ 1261ms | ✓ 811ms | ✓ 714ms | http |
| 205.209.118.30:3138 | ✓ 339ms | ✓ 1958ms | ✓ 1040ms | ✓ 1603ms | ✓ 1045ms | http |
| 101.47.73.135:3128 | ✓ 1630ms | 否 | ✓ 1486ms | ✓ 1427ms | ✓ 1275ms | http |
| 34.32.154.33:3128 | ✓ 1073ms | 否 | ✓ 1384ms | 否 | ✓ 1335ms | http |
| 34.79.102.160:3128 | ✓ 1072ms | ✓ 1808ms | ✓ 1368ms | 否 | ✓ 1787ms | http |
| 34.158.73.60:3128 | ✓ 1075ms | 否 | ✓ 1199ms | 否 | ✓ 1973ms | http |
| 34.89.174.168:3128 | ✓ 1091ms | 否 | ✓ 1520ms | 否 | ✓ 1978ms | http |
| 34.7.88.87:3128 | ✓ 1068ms | ✓ 1900ms | ✓ 1308ms | ✓ 1943ms | ✓ 1914ms | http |
| 34.159.121.205:3128 | ✓ 1064ms | 否 | ✓ 1552ms | 否 | ✓ 1932ms | http |
| 3.213.157.4:3128 | ✓ 341ms | 否 | ✓ 1003ms | ✓ 1466ms | ✓ 1324ms | http |
| 34.78.200.22:3128 | ✓ 691ms | 否 | ✓ 1251ms | 否 | ✓ 1464ms | http |
| 34.185.159.217:3128 | ✓ 912ms | 否 | ✓ 1235ms | 否 | ✓ 1740ms | http |
| 14.56.107.244:3128 | ✓ 545ms | ✓ 1316ms | ✓ 928ms | ✓ 1924ms | ✓ 707ms | http |
| 103.84.95.54:7890 | ✓ 647ms | 否 | 否 | ✓ 759ms | ✓ 760ms | http |
| 85.208.108.43:2094 | ✓ 1129ms | 否 | ✓ 1094ms | ✓ 1309ms | ✓ 848ms | http |
| 120.92.212.16:8890 | ✓ 906ms | 否 | ✓ 1154ms | ✓ 1379ms | ✓ 1133ms | http |
| 35.234.17.221:8080 | ✓ 1315ms | 否 | ✓ 1002ms | ✓ 1160ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1226ms | ✓ 1521ms | ✓ 1452ms | ✓ 1640ms | ✓ 1244ms | http |
| 81.70.169.194:80 | ✓ 977ms | ✓ 1191ms | ✓ 1000ms | ✓ 1192ms | ✓ 993ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1111ms | ✓ 886ms | ✓ 1115ms | ✓ 862ms | http |
| 101.43.255.96:80 | ✓ 885ms | ✓ 1215ms | ✓ 869ms | ✓ 1172ms | ✓ 1002ms | http |
| 211.171.114.154:3128 | ✓ 1009ms | ✓ 1416ms | ✓ 1856ms | ✓ 1760ms | 否 | http |
| 59.46.216.131:30001 | ✓ 916ms | 否 | 否 | ✓ 1313ms | ✓ 1062ms | http |
| 210.223.44.230:3128 | ✓ 641ms | ✓ 751ms | ✓ 564ms | ✓ 870ms | ✓ 987ms | http |
| 143.198.37.6:8888 | 否 | 否 | ✓ 1038ms | ✓ 1377ms | ✓ 1240ms | http |
| 120.92.212.16:7890 | ✓ 888ms | ✓ 1145ms | ✓ 987ms | ✓ 1378ms | ✓ 923ms | http |
| 121.230.9.148:1080 | 否 | ✓ 1347ms | ✓ 956ms | 否 | ✓ 914ms | http |
| 116.80.63.67:7777 | 否 | 否 | ✓ 1566ms | ✓ 1801ms | ✓ 1607ms | http |
| 185.115.74.185:8080 | ✓ 1127ms | ✓ 1928ms | ✓ 1802ms | 否 | 否 | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1608ms | ✓ 1412ms | ✓ 1349ms | http |
| 35.225.22.61:80 | ✓ 676ms | ✓ 1460ms | ✓ 323ms | ✓ 1210ms | 否 | http |
| 103.236.64.247:8888 | 否 | ✓ 1195ms | 否 | ✓ 1139ms | ✓ 1280ms | http |
| 61.72.110.24:3128 | ✓ 1451ms | 否 | ✓ 1206ms | ✓ 1291ms | 否 | http |
| 138.124.53.25:7443 | ✓ 636ms | 否 | ✓ 1619ms | ✓ 1671ms | ✓ 1433ms | http |
| 115.231.181.40:8128 | ✓ 981ms | ✓ 1123ms | ✓ 1733ms | 否 | 否 | http |
| 103.104.99.29:80 | ✓ 1909ms | 否 | ✓ 1623ms | ✓ 1584ms | ✓ 1408ms | http |
| 5.101.0.233:3128 | ✓ 912ms | ✓ 1817ms | ✓ 1759ms | 否 | 否 | http |
| 36.147.78.166:80 | 否 | ✓ 1552ms | ✓ 1611ms | ✓ 1775ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1083ms | 否 | ✓ 861ms | 否 | ✓ 1007ms | http |
| 91.238.104.172:2024 | ✓ 927ms | ✓ 1769ms | ✓ 1095ms | ✓ 1770ms | ✓ 1356ms | http |
| 52.188.28.218:3128 | 否 | 否 | ✓ 974ms | ✓ 1668ms | ✓ 1294ms | http |
| 111.79.111.126:3128 | ✓ 1053ms | ✓ 1250ms | ✓ 1364ms | ✓ 1482ms | 否 | http |
| 14.143.222.113:10158 | ✓ 1996ms | 否 | ✓ 949ms | ✓ 1881ms | 否 | http |
| 222.228.171.92:8080 | ✓ 1760ms | ✓ 1562ms | 否 | 否 | ✓ 1632ms | http |
| 43.161.214.161:1080 | ✓ 840ms | ✓ 1172ms | ✓ 830ms | ✓ 1022ms | ✓ 1025ms | http |
| 168.235.110.63:3128 | ✓ 1349ms | ✓ 1560ms | ✓ 1234ms | ✓ 1403ms | ✓ 1017ms | http |
| 180.127.149.225:1080 | ✓ 1918ms | ✓ 1170ms | ✓ 942ms | ✓ 1242ms | ✓ 832ms | http |
| 121.230.9.248:1080 | ✓ 1087ms | ✓ 1388ms | ✓ 1192ms | ✓ 1264ms | ✓ 1194ms | http |
| 121.230.8.181:1080 | ✓ 1197ms | ✓ 1970ms | ✓ 1471ms | ✓ 1704ms | ✓ 1229ms | http |
| 121.230.8.97:1080 | ✓ 1219ms | ✓ 1717ms | 否 | ✓ 1633ms | 否 | http |
| 180.127.149.244:1080 | ✓ 960ms | ✓ 1128ms | ✓ 924ms | ✓ 1135ms | ✓ 1208ms | http |
| 34.78.177.18:3128 | ✓ 1167ms | 否 | ✓ 1754ms | 否 | ✓ 1749ms | http |
| 61.72.110.54:3128 | ✓ 1938ms | 否 | ✓ 1333ms | 否 | ✓ 1752ms | http |
| 8.219.97.248:80 | ✓ 1221ms | 否 | ✓ 1910ms | 否 | ✓ 1608ms | http |
| 121.230.9.54:1080 | 否 | ✓ 1354ms | ✓ 1105ms | ✓ 1407ms | 否 | http |
| 157.0.142.246:10057 | ✓ 975ms | ✓ 1231ms | ✓ 978ms | ✓ 1225ms | ✓ 982ms | http |
| 162.240.154.26:3128 | ✓ 1406ms | 否 | ✓ 1852ms | 否 | ✓ 1313ms | http |
| 36.147.78.166:443 | ✓ 1571ms | ✓ 1509ms | 否 | ✓ 1653ms | 否 | http |
| 175.194.173.105:3128 | ✓ 1711ms | ✓ 974ms | ✓ 979ms | 否 | 否 | http |
| 103.215.36.88:15247 | ✓ 964ms | ✓ 1111ms | ✓ 933ms | ✓ 1277ms | ✓ 905ms | http |
| 8.217.147.173:8080 | ✓ 1386ms | 否 | ✓ 1327ms | ✓ 1120ms | ✓ 932ms | http |
| 167.71.20.25:3128 | ✓ 394ms | ✓ 1425ms | ✓ 1009ms | ✓ 1224ms | ✓ 914ms | http |
| 217.119.129.86:2222 | ✓ 822ms | 否 | ✓ 1327ms | 否 | ✓ 1833ms | http |
| 160.238.65.9:3128 | ✓ 1295ms | 否 | ✓ 1681ms | 否 | ✓ 1776ms | http |
| 103.39.51.190:8080 | ✓ 1723ms | 否 | 否 | ✓ 1549ms | ✓ 1971ms | http |
| 121.128.121.54:3128 | ✓ 1651ms | ✓ 1969ms | ✓ 1172ms | ✓ 1884ms | 否 | http |
| 61.72.110.94:3128 | 否 | 否 | ✓ 1215ms | ✓ 1911ms | ✓ 1110ms | http |
| 120.92.211.211:7890 | ✓ 1092ms | 否 | ✓ 1829ms | ✓ 1753ms | 否 | http |
| 121.40.231.103:7890 | ✓ 1917ms | 否 | 否 | ✓ 1989ms | ✓ 1583ms | http |
| 31.59.129.75:8080 | ✓ 668ms | ✓ 1883ms | ✓ 1655ms | 否 | 否 | http |
| 185.226.195.218:2222 | ✓ 1649ms | 否 | ✓ 1700ms | 否 | ✓ 1693ms | http |
| 45.136.198.40:3128 | ✓ 1172ms | ✓ 1697ms | ✓ 1883ms | 否 | ✓ 1951ms | http |
| 77.83.203.6:443 | ✓ 1614ms | 否 | ✓ 1954ms | 否 | ✓ 1819ms | http |
| 172.212.68.37:3128 | ✓ 527ms | ✓ 1929ms | ✓ 1553ms | ✓ 1594ms | ✓ 1014ms | http |
| 182.204.180.146:1080 | ✓ 1095ms | ✓ 1456ms | ✓ 1145ms | 否 | 否 | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1713ms | ✓ 1771ms | ✓ 1440ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1028ms | ✓ 1457ms | ✓ 1455ms | ✓ 783ms | http |
| 45.140.147.155:1081 | ✓ 1006ms | ✓ 1730ms | ✓ 1509ms | 否 | 否 | http |
| 5.75.201.136:1080 | ✓ 613ms | 否 | ✓ 1820ms | ✓ 1927ms | 否 | http |
| 121.230.8.111:1080 | ✓ 985ms | ✓ 1372ms | ✓ 1016ms | ✓ 1527ms | ✓ 1134ms | http |
| 209.38.54.154:8443 | 否 | 否 | ✓ 869ms | ✓ 1505ms | ✓ 1476ms | http |
| 91.238.104.171:2023 | ✓ 1151ms | 否 | ✓ 812ms | ✓ 1804ms | ✓ 1345ms | http |
| 120.55.163.237:10086 | ✓ 1529ms | ✓ 1527ms | 否 | 否 | ✓ 1949ms | http |

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
