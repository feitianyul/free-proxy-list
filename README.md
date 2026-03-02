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

最后更新：2026-03-02 06:48:19 UTC（2026-03-02 14:48:19 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | 否 | ✓ 1590ms | ✓ 968ms | ✓ 1127ms | ✓ 812ms | http |
| 142.171.85.32:1080 | ✓ 687ms | 否 | ✓ 1171ms | 否 | ✓ 1245ms | http |
| 85.198.84.77:10808 | ✓ 1147ms | ✓ 1827ms | ✓ 1575ms | 否 | ✓ 1487ms | http |
| 45.125.67.37:8443 | ✓ 1082ms | 否 | ✓ 1733ms | ✓ 1289ms | ✓ 1292ms | http |
| 61.72.110.54:3128 | ✓ 1924ms | 否 | 否 | ✓ 1656ms | ✓ 1020ms | http |
| 115.76.5.32:10006 | ✓ 1778ms | 否 | ✓ 1808ms | ✓ 1896ms | 否 | http |
| 45.190.78.20:999 | ✓ 1679ms | 否 | ✓ 1219ms | ✓ 1443ms | 否 | http |
| 35.225.22.61:80 | ✓ 926ms | 否 | ✓ 974ms | ✓ 1120ms | ✓ 1076ms | http |
| 5.75.196.26:40000 | ✓ 427ms | ✓ 1356ms | 否 | 否 | ✓ 1090ms | http |
| 162.240.154.26:3128 | ✓ 1162ms | 否 | ✓ 1737ms | ✓ 1558ms | 否 | http |
| 121.237.181.137:8888 | ✓ 1130ms | ✓ 1330ms | ✓ 1128ms | 否 | ✓ 1035ms | http |
| 185.115.74.185:8080 | ✓ 675ms | ✓ 1909ms | ✓ 1362ms | 否 | 否 | http |
| 165.227.5.10:8888 | ✓ 543ms | ✓ 1234ms | ✓ 1035ms | ✓ 1939ms | ✓ 914ms | http |
| 120.92.212.16:8890 | ✓ 1185ms | ✓ 1480ms | ✓ 1403ms | 否 | ✓ 1129ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1681ms | ✓ 1123ms | ✓ 1707ms | ✓ 1124ms | http |
| 210.223.44.230:3128 | ✓ 1762ms | 否 | ✓ 1479ms | 否 | ✓ 1934ms | http |
| 35.234.17.221:8080 | ✓ 1045ms | 否 | ✓ 1403ms | ✓ 1246ms | 否 | http |
| 70.22.175.232:3128 | 否 | 否 | ✓ 113ms | ✓ 984ms | ✓ 781ms | http |
| 107.174.133.10:3128 | ✓ 710ms | ✓ 1088ms | ✓ 985ms | ✓ 1101ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1094ms | 否 | 否 | ✓ 1083ms | ✓ 962ms | http |
| 81.70.169.194:80 | ✓ 1226ms | ✓ 1495ms | ✓ 1140ms | ✓ 1515ms | ✓ 1147ms | http |
| 222.28.182.229:7890 | ✓ 1250ms | ✓ 1434ms | ✓ 1250ms | 否 | 否 | http |
| 138.124.53.25:7443 | ✓ 1760ms | ✓ 1401ms | ✓ 1034ms | 否 | ✓ 1925ms | http |
| 103.113.70.189:1081 | 否 | ✓ 911ms | 否 | ✓ 1069ms | ✓ 1162ms | http |
| 74.208.234.198:443 | ✓ 714ms | ✓ 1830ms | 否 | ✓ 1648ms | 否 | http |
| 5.129.228.225:1080 | ✓ 1732ms | ✓ 1456ms | 否 | ✓ 1241ms | 否 | http |
| 38.180.2.107:3128 | ✓ 793ms | 否 | ✓ 1700ms | 否 | ✓ 1598ms | http |
| 45.136.198.40:3128 | ✓ 1240ms | ✓ 1532ms | ✓ 1651ms | 否 | ✓ 1662ms | http |
| 46.249.103.192:443 | ✓ 517ms | 否 | ✓ 1028ms | ✓ 1882ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1307ms | ✓ 1418ms | ✓ 1136ms | ✓ 1378ms | ✓ 1013ms | http |
| 120.202.127.234:10808 | ✓ 1623ms | 否 | ✓ 999ms | ✓ 1342ms | ✓ 1319ms | http |
| 121.230.8.61:1080 | ✓ 1563ms | 否 | ✓ 1707ms | 否 | ✓ 1361ms | http |
| 34.101.184.164:3128 | ✓ 1104ms | 否 | ✓ 1185ms | ✓ 1589ms | ✓ 1245ms | http |
| 115.76.5.32:10005 | ✓ 1564ms | 否 | ✓ 1441ms | 否 | ✓ 1540ms | http |
| 91.238.104.172:2024 | ✓ 577ms | ✓ 1443ms | ✓ 836ms | ✓ 1417ms | ✓ 1107ms | http |
| 91.238.104.171:2023 | ✓ 554ms | 否 | ✓ 583ms | ✓ 1473ms | ✓ 1108ms | http |
| 188.166.208.168:9876 | 否 | 否 | ✓ 1120ms | ✓ 1412ms | ✓ 1426ms | http |
| 177.93.33.55:999 | ✓ 1195ms | 否 | ✓ 1399ms | ✓ 1604ms | 否 | http |
| 90.84.188.97:8000 | ✓ 1994ms | 否 | ✓ 889ms | 否 | ✓ 1727ms | http |
| 36.147.78.166:80 | ✓ 1875ms | 否 | ✓ 1956ms | 否 | ✓ 1844ms | http |
| 211.171.114.154:3128 | ✓ 1500ms | ✓ 1910ms | ✓ 1497ms | ✓ 1878ms | 否 | http |
| 62.113.119.14:8080 | ✓ 785ms | ✓ 1536ms | ✓ 545ms | ✓ 1411ms | ✓ 1067ms | http |
| 85.208.108.43:2094 | ✓ 377ms | 否 | ✓ 1084ms | ✓ 1118ms | ✓ 890ms | http |
| 45.140.147.155:1081 | ✓ 1124ms | 否 | ✓ 1370ms | ✓ 1371ms | ✓ 1158ms | http |
| 59.46.216.131:30001 | ✓ 1302ms | ✓ 1789ms | 否 | 否 | ✓ 1277ms | http |
| 5.101.0.233:3128 | ✓ 1890ms | 否 | ✓ 1449ms | 否 | ✓ 1607ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1711ms | ✓ 1604ms | ✓ 1912ms | http |
| 217.77.102.18:3128 | ✓ 860ms | 否 | ✓ 1928ms | 否 | ✓ 1747ms | http |
| 103.3.246.71:3128 | ✓ 1104ms | 否 | ✓ 1183ms | ✓ 1397ms | ✓ 1131ms | http |
| 103.215.36.88:19328 | 否 | ✓ 1586ms | 否 | ✓ 1633ms | ✓ 1248ms | http |
| 171.234.62.116:10002 | ✓ 1796ms | 否 | ✓ 1891ms | ✓ 1606ms | ✓ 1692ms | http |
| 115.76.5.32:10007 | ✓ 1550ms | 否 | ✓ 1724ms | 否 | ✓ 1846ms | http |
| 144.31.25.69:21064 | ✓ 1626ms | 否 | ✓ 1985ms | 否 | ✓ 1980ms | http |
| 217.217.254.94:8080 | ✓ 1327ms | 否 | ✓ 1760ms | 否 | ✓ 1876ms | http |
| 1.225.116.115:1080 | ✓ 1210ms | 否 | ✓ 1884ms | ✓ 1213ms | ✓ 1051ms | http |
| 129.212.226.87:443 | ✓ 1712ms | 否 | ✓ 1135ms | ✓ 1402ms | ✓ 1001ms | http |
| 45.140.147.82:1081 | ✓ 378ms | ✓ 1785ms | ✓ 1177ms | ✓ 1655ms | ✓ 1662ms | http |
| 185.242.233.230:3128 | 否 | 否 | ✓ 878ms | ✓ 1102ms | ✓ 827ms | http |
| 223.16.170.103:3128 | ✓ 1347ms | 否 | ✓ 1043ms | ✓ 1410ms | ✓ 1372ms | http |
| 103.39.51.190:8080 | ✓ 1950ms | 否 | 否 | ✓ 1652ms | ✓ 1624ms | http |
| 95.85.252.153:21064 | ✓ 911ms | ✓ 1649ms | ✓ 1379ms | 否 | 否 | http |
| 167.160.184.231:6005 | ✓ 498ms | ✓ 1851ms | ✓ 1942ms | ✓ 1221ms | ✓ 1243ms | http |
| 103.236.64.247:8888 | 否 | 否 | ✓ 1293ms | ✓ 1429ms | ✓ 1172ms | http |
| 193.32.178.160:57329 | ✓ 1792ms | 否 | ✓ 1347ms | ✓ 1424ms | 否 | http |
| 5.102.109.41:999 | ✓ 608ms | ✓ 1485ms | ✓ 1097ms | ✓ 1958ms | 否 | http |
| 121.230.8.49:1080 | 否 | 否 | ✓ 1347ms | ✓ 1832ms | ✓ 1326ms | http |
| 115.76.5.32:10010 | 否 | 否 | ✓ 1449ms | ✓ 1981ms | ✓ 1862ms | http |
| 172.212.68.37:3128 | ✓ 509ms | 否 | ✓ 791ms | ✓ 1863ms | ✓ 770ms | http |
| 122.2.48.121:8080 | ✓ 1508ms | 否 | ✓ 1679ms | ✓ 1801ms | ✓ 1584ms | http |
| 121.230.9.184:1080 | ✓ 1347ms | 否 | ✓ 1354ms | 否 | ✓ 1497ms | http |
| 115.76.5.32:10009 | ✓ 1688ms | 否 | ✓ 1946ms | 否 | ✓ 1621ms | http |
| 36.147.78.166:443 | ✓ 1937ms | ✓ 1847ms | 否 | 否 | ✓ 1607ms | http |
| 103.82.23.118:5171 | ✓ 1896ms | 否 | ✓ 1418ms | ✓ 1894ms | ✓ 1511ms | http |
| 113.255.59.226:8080 | ✓ 1784ms | 否 | 否 | ✓ 1364ms | ✓ 1445ms | http |
| 45.88.0.115:3128 | ✓ 547ms | ✓ 1800ms | 否 | 否 | ✓ 1346ms | http |
| 45.88.0.116:3128 | ✓ 869ms | 否 | ✓ 1392ms | ✓ 1245ms | ✓ 1080ms | http |
| 45.88.0.113:3128 | ✓ 836ms | 否 | ✓ 1390ms | ✓ 1250ms | ✓ 1078ms | http |
| 45.88.0.99:3128 | ✓ 834ms | 否 | ✓ 1386ms | ✓ 1250ms | ✓ 1078ms | http |
| 45.88.0.114:3128 | ✓ 365ms | ✓ 1737ms | ✓ 639ms | ✓ 1252ms | ✓ 945ms | http |
| 45.88.0.98:3128 | ✓ 1424ms | ✓ 1254ms | ✓ 561ms | ✓ 1251ms | ✓ 967ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1088ms | ✓ 1487ms | ✓ 1092ms | http |
| 180.103.19.123:1080 | ✓ 1171ms | ✓ 1551ms | ✓ 1303ms | 否 | ✓ 1835ms | http |
| 45.88.0.111:3128 | ✓ 359ms | ✓ 1173ms | ✓ 502ms | ✓ 1188ms | ✓ 1212ms | http |
| 121.230.9.205:1080 | 否 | ✓ 1697ms | 否 | ✓ 1770ms | ✓ 1413ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1672ms | ✓ 1689ms | ✓ 1634ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1425ms | ✓ 1612ms | ✓ 1542ms | http |

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
