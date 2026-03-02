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

最后更新：2026-03-02 19:46:18 UTC（2026-03-03 03:46:18 UTC+8）

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
| 185.115.74.185:8080 | ✓ 907ms | ✓ 1953ms | ✓ 1820ms | 否 | 否 | http |
| 77.83.203.6:443 | ✓ 1624ms | ✓ 1767ms | 否 | 否 | ✓ 1513ms | http |
| 217.76.245.80:999 | ✓ 661ms | 否 | ✓ 1431ms | ✓ 1370ms | ✓ 1143ms | http |
| 35.234.17.221:8080 | ✓ 962ms | 否 | ✓ 1245ms | ✓ 1365ms | 否 | http |
| 91.99.99.83:9000 | ✓ 788ms | ✓ 1626ms | ✓ 1660ms | ✓ 1906ms | ✓ 1835ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1607ms | ✓ 1654ms | ✓ 998ms | http |
| 45.88.0.114:3128 | ✓ 784ms | ✓ 1595ms | ✓ 769ms | 否 | 否 | http |
| 115.76.5.32:10008 | ✓ 1852ms | 否 | ✓ 1725ms | ✓ 1695ms | ✓ 1870ms | http |
| 35.225.22.61:80 | ✓ 728ms | ✓ 1132ms | ✓ 1003ms | ✓ 943ms | ✓ 792ms | http |
| 81.70.169.194:80 | ✓ 1078ms | ✓ 1485ms | ✓ 1027ms | ✓ 1356ms | ✓ 1170ms | http |
| 101.43.255.96:80 | ✓ 1134ms | ✓ 1374ms | 否 | ✓ 1369ms | ✓ 1580ms | http |
| 45.129.141.143:3128 | ✓ 1489ms | 否 | 否 | ✓ 1940ms | ✓ 1530ms | http |
| 74.48.78.224:2080 | ✓ 1017ms | ✓ 1738ms | ✓ 1362ms | ✓ 1211ms | ✓ 1000ms | http |
| 38.207.165.2:6005 | ✓ 1837ms | 否 | ✓ 1494ms | 否 | ✓ 1392ms | http |
| 5.75.196.26:40000 | ✓ 555ms | ✓ 1715ms | ✓ 472ms | 否 | 否 | http |
| 205.209.118.30:3138 | ✓ 1183ms | ✓ 1986ms | ✓ 126ms | 否 | 否 | http |
| 121.128.121.54:3128 | ✓ 832ms | ✓ 1516ms | ✓ 846ms | ✓ 1736ms | 否 | http |
| 190.9.48.129:999 | ✓ 603ms | ✓ 1268ms | ✓ 1120ms | 否 | 否 | http |
| 115.76.5.32:10010 | ✓ 1567ms | 否 | ✓ 1544ms | ✓ 1806ms | 否 | http |
| 170.78.208.251:999 | ✓ 860ms | ✓ 1973ms | ✓ 1696ms | 否 | 否 | http |
| 115.76.5.32:10005 | ✓ 1582ms | 否 | ✓ 1377ms | 否 | ✓ 1897ms | http |
| 103.39.51.207:8080 | ✓ 1938ms | 否 | ✓ 1375ms | 否 | ✓ 1465ms | http |
| 34.101.184.164:3128 | ✓ 1784ms | 否 | ✓ 943ms | ✓ 1435ms | ✓ 1121ms | http |
| 171.234.62.116:10001 | ✓ 1910ms | 否 | ✓ 1784ms | ✓ 1942ms | ✓ 1633ms | http |
| 171.234.62.116:10002 | ✓ 1905ms | 否 | 否 | ✓ 1754ms | ✓ 1666ms | http |
| 47.113.95.226:8118 | 否 | 否 | ✓ 1022ms | ✓ 1581ms | ✓ 1189ms | http |
| 38.180.2.107:3128 | ✓ 1870ms | ✓ 1974ms | ✓ 1912ms | 否 | ✓ 1706ms | http |
| 45.136.198.40:3128 | ✓ 938ms | ✓ 1801ms | 否 | ✓ 1910ms | ✓ 1613ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1320ms | ✓ 1066ms | 否 | ✓ 1822ms | http |
| 115.76.5.32:10009 | ✓ 1939ms | 否 | ✓ 1797ms | 否 | ✓ 1972ms | http |
| 160.238.65.5:3128 | ✓ 1316ms | ✓ 1457ms | 否 | ✓ 1847ms | ✓ 1515ms | http |
| 160.238.65.9:3128 | ✓ 1316ms | ✓ 1557ms | 否 | ✓ 1726ms | ✓ 1541ms | http |
| 160.238.65.6:3128 | ✓ 1314ms | ✓ 1658ms | 否 | ✓ 1623ms | ✓ 1540ms | http |
| 160.238.65.8:3128 | ✓ 1315ms | 否 | 否 | ✓ 1297ms | ✓ 1618ms | http |
| 160.238.65.2:3128 | ✓ 1315ms | 否 | 否 | ✓ 1278ms | ✓ 1908ms | http |
| 171.234.62.116:10004 | 否 | 否 | ✓ 1736ms | ✓ 1872ms | ✓ 1668ms | http |
| 106.14.205.114:483 | ✓ 1154ms | ✓ 1127ms | ✓ 1285ms | ✓ 1222ms | ✓ 976ms | http |
| 160.238.65.4:3128 | 否 | 否 | ✓ 473ms | ✓ 1258ms | ✓ 1534ms | http |
| 142.171.131.38:7890 | 否 | ✓ 822ms | ✓ 991ms | ✓ 990ms | ✓ 730ms | http |
| 121.230.8.208:1080 | ✓ 1191ms | ✓ 1811ms | ✓ 1169ms | ✓ 1636ms | ✓ 1348ms | http |
| 45.88.0.111:3128 | ✓ 1685ms | 否 | ✓ 1182ms | 否 | ✓ 1010ms | http |
| 45.88.0.116:3128 | ✓ 885ms | 否 | ✓ 1538ms | ✓ 1281ms | ✓ 1008ms | http |
| 45.88.0.98:3128 | ✓ 885ms | 否 | ✓ 878ms | ✓ 1298ms | 否 | http |
| 45.88.0.117:3128 | ✓ 1282ms | 否 | ✓ 483ms | ✓ 1893ms | 否 | http |
| 121.230.9.161:1080 | ✓ 1323ms | ✓ 1579ms | 否 | ✓ 1963ms | ✓ 1688ms | http |
| 101.32.244.83:8080 | ✓ 1353ms | ✓ 1569ms | ✓ 1056ms | ✓ 1527ms | ✓ 1398ms | http |
| 114.55.226.123:10086 | ✓ 1190ms | ✓ 1480ms | ✓ 1113ms | ✓ 1397ms | ✓ 1137ms | http |
| 14.56.177.44:3128 | ✓ 1127ms | ✓ 1270ms | ✓ 1070ms | ✓ 1296ms | ✓ 1022ms | http |
| 157.230.38.173:3128 | ✓ 1569ms | 否 | ✓ 984ms | ✓ 1210ms | ✓ 921ms | http |
| 157.0.142.246:10061 | ✓ 1126ms | ✓ 1378ms | ✓ 1109ms | ✓ 1501ms | ✓ 1131ms | http |
| 103.139.138.194:3128 | ✓ 1929ms | 否 | ✓ 1466ms | ✓ 1558ms | ✓ 1257ms | http |
| 222.228.171.92:8080 | ✓ 1267ms | 否 | ✓ 1732ms | ✓ 1221ms | 否 | http |
| 125.128.12.84:3128 | ✓ 1804ms | 否 | ✓ 763ms | 否 | ✓ 904ms | http |
| 45.140.147.82:1081 | 否 | 否 | ✓ 1034ms | ✓ 1856ms | ✓ 1326ms | http |
| 115.76.5.32:10007 | ✓ 1714ms | 否 | ✓ 1929ms | 否 | ✓ 1960ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1677ms | 否 | ✓ 1192ms | ✓ 1631ms | http |
| 138.124.53.25:7443 | ✓ 1657ms | 否 | ✓ 1796ms | ✓ 1796ms | ✓ 1684ms | http |
| 113.255.59.226:8080 | ✓ 1329ms | 否 | ✓ 1201ms | ✓ 1254ms | ✓ 1303ms | http |
| 150.249.255.91:3128 | ✓ 1400ms | 否 | ✓ 867ms | ✓ 972ms | ✓ 781ms | http |
| 186.148.180.46:999 | ✓ 976ms | ✓ 1840ms | ✓ 1818ms | ✓ 1788ms | ✓ 1330ms | http |
| 114.231.72.60:1080 | ✓ 1016ms | ✓ 1336ms | ✓ 1032ms | ✓ 1283ms | ✓ 1074ms | http |
| 61.72.221.194:3128 | ✓ 772ms | 否 | ✓ 1181ms | ✓ 1291ms | 否 | http |
| 221.127.195.224:8888 | ✓ 1311ms | 否 | ✓ 1246ms | ✓ 1440ms | ✓ 1243ms | http |
| 85.208.108.43:2094 | ✓ 817ms | 否 | ✓ 1262ms | ✓ 1160ms | 否 | http |
| 185.243.218.43:49153 | ✓ 1127ms | 否 | ✓ 1733ms | ✓ 1831ms | 否 | http |
| 165.225.113.220:10958 | ✓ 835ms | 否 | ✓ 1186ms | ✓ 1449ms | ✓ 1160ms | http |
| 165.225.113.220:11462 | ✓ 828ms | 否 | ✓ 1030ms | 否 | ✓ 951ms | http |
| 45.140.147.82:1082 | ✓ 419ms | ✓ 1596ms | ✓ 1233ms | ✓ 1949ms | ✓ 1460ms | http |
| 120.92.212.16:7890 | ✓ 1056ms | 否 | ✓ 1063ms | 否 | ✓ 1802ms | http |
| 202.129.206.239:3128 | ✓ 1585ms | 否 | 否 | ✓ 1843ms | ✓ 1947ms | http |
| 91.107.148.58:53967 | ✓ 477ms | ✓ 1779ms | ✓ 1907ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1949ms | 否 | 否 | ✓ 1915ms | ✓ 1509ms | http |
| 45.140.147.155:1081 | ✓ 448ms | ✓ 1886ms | ✓ 839ms | ✓ 1715ms | ✓ 1265ms | http |
| 125.128.12.114:3128 | ✓ 1527ms | 否 | ✓ 753ms | 否 | ✓ 970ms | http |
| 125.128.12.194:3128 | ✓ 1525ms | 否 | ✓ 1897ms | ✓ 1222ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1168ms | ✓ 1246ms | ✓ 1222ms | ✓ 1133ms | ✓ 888ms | http |
| 37.187.109.70:10111 | ✓ 859ms | 否 | ✓ 1933ms | 否 | ✓ 1819ms | http |
| 103.131.19.42:8181 | ✓ 1486ms | 否 | ✓ 1733ms | ✓ 1583ms | 否 | http |
| 121.230.9.148:1080 | 否 | ✓ 1867ms | ✓ 1300ms | 否 | ✓ 1504ms | http |
| 61.72.110.54:3128 | ✓ 1610ms | 否 | 否 | ✓ 1507ms | ✓ 1293ms | http |
| 61.72.110.94:3128 | 否 | ✓ 1748ms | ✓ 1474ms | ✓ 1909ms | ✓ 1786ms | http |
| 210.223.44.230:3128 | ✓ 940ms | ✓ 973ms | ✓ 1052ms | ✓ 1206ms | ✓ 1153ms | http |
| 94.177.131.12:3128 | ✓ 755ms | 否 | ✓ 1544ms | ✓ 1008ms | ✓ 822ms | http |
| 171.234.62.116:10010 | 否 | 否 | ✓ 1871ms | ✓ 1561ms | ✓ 1362ms | http |
| 61.72.221.94:3128 | ✓ 744ms | ✓ 1474ms | 否 | 否 | ✓ 1982ms | http |

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
