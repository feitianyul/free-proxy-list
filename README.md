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

最后更新：2026-03-02 05:41:19 UTC（2026-03-02 13:41:19 UTC+8）

**代理总数：91**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 251ms | ✓ 978ms | ✓ 918ms | ✓ 1270ms | ✓ 789ms | http |
| 5.75.196.26:40000 | ✓ 389ms | ✓ 1745ms | 否 | 否 | ✓ 1555ms | http |
| 101.47.73.135:3128 | ✓ 1106ms | 否 | ✓ 1138ms | ✓ 1618ms | 否 | http |
| 183.128.208.68:7890 | ✓ 1139ms | 否 | ✓ 1222ms | ✓ 1361ms | ✓ 1013ms | http |
| 116.99.49.187:10002 | ✓ 1572ms | 否 | ✓ 1578ms | ✓ 1893ms | ✓ 1367ms | http |
| 115.76.5.32:10005 | 否 | 否 | ✓ 1737ms | ✓ 1812ms | ✓ 1703ms | http |
| 118.68.216.26:10006 | 否 | 否 | ✓ 1949ms | ✓ 1967ms | ✓ 1914ms | http |
| 85.198.84.77:10808 | ✓ 1283ms | 否 | ✓ 1546ms | 否 | ✓ 1513ms | http |
| 118.68.216.26:10007 | ✓ 1712ms | 否 | ✓ 1772ms | 否 | ✓ 1725ms | http |
| 59.46.216.131:30001 | ✓ 1867ms | ✓ 1722ms | ✓ 1208ms | 否 | 否 | http |
| 74.208.234.198:443 | ✓ 543ms | 否 | ✓ 1664ms | ✓ 1496ms | ✓ 1054ms | http |
| 35.234.17.221:8080 | ✓ 1009ms | ✓ 1690ms | 否 | 否 | ✓ 1144ms | http |
| 120.202.127.234:10808 | ✓ 1093ms | ✓ 1470ms | ✓ 1094ms | 否 | ✓ 1047ms | http |
| 101.43.255.96:80 | ✓ 1250ms | 否 | ✓ 1877ms | ✓ 1472ms | ✓ 1219ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 894ms | ✓ 1036ms | ✓ 840ms | http |
| 35.225.22.61:80 | ✓ 291ms | 否 | ✓ 236ms | ✓ 1147ms | ✓ 987ms | http |
| 95.85.252.153:21064 | ✓ 776ms | ✓ 1687ms | 否 | ✓ 1774ms | 否 | http |
| 81.70.169.194:80 | 否 | ✓ 1421ms | ✓ 1096ms | ✓ 1432ms | ✓ 1159ms | http |
| 120.92.212.16:7890 | ✓ 1072ms | 否 | 否 | ✓ 1659ms | ✓ 1134ms | http |
| 143.198.37.6:8888 | ✓ 144ms | ✓ 1777ms | ✓ 154ms | ✓ 920ms | ✓ 731ms | http |
| 107.174.133.10:3128 | 否 | ✓ 1311ms | ✓ 1070ms | ✓ 1139ms | ✓ 1098ms | http |
| 45.125.67.37:8443 | ✓ 1072ms | 否 | ✓ 1278ms | ✓ 1485ms | ✓ 1047ms | http |
| 121.40.231.103:7890 | 否 | 否 | ✓ 1023ms | ✓ 1304ms | ✓ 1062ms | http |
| 45.140.147.82:1081 | ✓ 751ms | ✓ 1710ms | ✓ 1025ms | ✓ 1817ms | ✓ 1050ms | http |
| 116.108.130.196:4000 | ✓ 1857ms | 否 | ✓ 1048ms | ✓ 1612ms | ✓ 1202ms | http |
| 116.108.138.138:4000 | ✓ 1692ms | 否 | ✓ 1184ms | ✓ 1956ms | ✓ 1118ms | http |
| 120.79.99.232:8099 | ✓ 1399ms | ✓ 1757ms | ✓ 1513ms | ✓ 1668ms | ✓ 1391ms | http |
| 121.237.181.137:8888 | ✓ 1797ms | ✓ 1271ms | ✓ 1384ms | ✓ 1683ms | ✓ 1122ms | http |
| 160.250.5.22:1 | ✓ 1682ms | 否 | 否 | ✓ 1433ms | ✓ 1500ms | http |
| 115.231.181.40:8128 | ✓ 1058ms | ✓ 1268ms | ✓ 1469ms | 否 | ✓ 1398ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1438ms | ✓ 1121ms | ✓ 1729ms | ✓ 1475ms | http |
| 91.238.104.171:2023 | 否 | 否 | ✓ 1464ms | ✓ 1709ms | ✓ 1633ms | http |
| 123.28.12.2:10003 | ✓ 1577ms | ✓ 1864ms | ✓ 1529ms | ✓ 1609ms | ✓ 1586ms | http |
| 115.76.5.32:10006 | ✓ 1856ms | 否 | ✓ 1800ms | ✓ 1883ms | ✓ 1615ms | http |
| 2.56.178.131:443 | ✓ 1833ms | 否 | ✓ 1356ms | ✓ 1682ms | 否 | http |
| 37.27.100.80:443 | ✓ 952ms | ✓ 1870ms | 否 | ✓ 1892ms | 否 | http |
| 209.38.51.97:3128 | 否 | ✓ 1735ms | 否 | ✓ 1122ms | ✓ 1300ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1803ms | ✓ 1681ms | ✓ 1579ms | ✓ 1603ms | http |
| 36.147.78.166:80 | ✓ 1973ms | ✓ 1851ms | 否 | 否 | ✓ 1963ms | http |
| 123.112.211.191:9000 | 否 | 否 | ✓ 1699ms | ✓ 1570ms | ✓ 1608ms | http |
| 195.90.215.186:3443 | ✓ 373ms | 否 | ✓ 1629ms | ✓ 1869ms | 否 | http |
| 106.14.205.114:483 | ✓ 1250ms | ✓ 1550ms | ✓ 1299ms | ✓ 1305ms | ✓ 1000ms | http |
| 45.136.198.40:3128 | ✓ 912ms | ✓ 1972ms | ✓ 1701ms | ✓ 1803ms | ✓ 1568ms | http |
| 49.7.179.70:3333 | 否 | 否 | ✓ 1136ms | ✓ 1488ms | ✓ 1332ms | http |
| 167.160.184.231:6005 | ✓ 512ms | 否 | ✓ 1005ms | ✓ 1066ms | ✓ 1181ms | http |
| 83.229.73.113:13554 | ✓ 874ms | 否 | ✓ 1716ms | 否 | ✓ 1690ms | http |
| 150.107.140.238:3128 | ✓ 1794ms | 否 | ✓ 1219ms | ✓ 1410ms | ✓ 1647ms | http |
| 165.227.5.10:8888 | ✓ 896ms | 否 | ✓ 852ms | ✓ 1332ms | 否 | http |
| 222.228.171.92:8080 | ✓ 1555ms | 否 | ✓ 1851ms | 否 | ✓ 1298ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1236ms | ✓ 1982ms | 否 | ✓ 1505ms | http |
| 170.78.208.245:999 | ✓ 1812ms | 否 | ✓ 414ms | ✓ 1826ms | ✓ 1019ms | http |
| 201.150.116.32:999 | ✓ 1874ms | ✓ 1971ms | ✓ 1032ms | ✓ 1341ms | ✓ 1508ms | http |
| 104.37.184.214:1080 | ✓ 1635ms | ✓ 1029ms | ✓ 446ms | ✓ 1248ms | ✓ 1131ms | http |
| 142.171.85.32:1080 | ✓ 846ms | ✓ 1850ms | 否 | 否 | ✓ 1103ms | http |
| 103.35.188.243:3128 | 否 | ✓ 927ms | 否 | ✓ 1264ms | ✓ 887ms | http |
| 180.191.16.105:8082 | ✓ 1696ms | 否 | ✓ 1959ms | ✓ 1932ms | 否 | http |
| 103.3.246.71:3128 | ✓ 1519ms | 否 | ✓ 1170ms | ✓ 1410ms | ✓ 1090ms | http |
| 123.20.24.166:8118 | 否 | 否 | ✓ 1519ms | ✓ 1712ms | ✓ 1333ms | http |
| 113.165.202.31:1001 | ✓ 1674ms | 否 | 否 | ✓ 1937ms | ✓ 1698ms | http |
| 202.129.206.239:3128 | ✓ 1305ms | 否 | ✓ 1447ms | 否 | ✓ 1661ms | http |
| 118.68.216.26:10003 | ✓ 1761ms | 否 | ✓ 1711ms | ✓ 1952ms | ✓ 1762ms | http |
| 113.162.163.104:10006 | ✓ 1851ms | ✓ 1974ms | ✓ 1944ms | ✓ 1917ms | ✓ 1644ms | http |
| 113.162.163.104:10005 | ✓ 1889ms | ✓ 1938ms | ✓ 1929ms | 否 | ✓ 1962ms | http |
| 70.22.175.232:3128 | ✓ 253ms | ✓ 1693ms | ✓ 1885ms | ✓ 977ms | ✓ 730ms | http |
| 138.124.53.25:7443 | ✓ 616ms | ✓ 1893ms | 否 | 否 | ✓ 1681ms | http |
| 121.230.9.81:1080 | ✓ 1406ms | ✓ 1577ms | ✓ 1717ms | 否 | ✓ 1384ms | http |
| 213.131.85.27:1976 | ✓ 1928ms | 否 | ✓ 1534ms | ✓ 1897ms | 否 | http |
| 197.164.101.11:1981 | 否 | 否 | ✓ 1756ms | ✓ 1681ms | ✓ 1253ms | http |
| 103.236.64.247:8888 | 否 | 否 | ✓ 1732ms | ✓ 1712ms | ✓ 1115ms | http |
| 85.208.108.43:2094 | ✓ 975ms | 否 | ✓ 604ms | ✓ 1182ms | ✓ 1257ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1584ms | ✓ 1365ms | 否 | ✓ 1984ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1922ms | ✓ 1882ms | ✓ 1912ms | http |
| 103.39.51.190:8080 | ✓ 1963ms | 否 | 否 | ✓ 1989ms | ✓ 1753ms | http |
| 222.28.182.229:7890 | ✓ 1228ms | ✓ 1990ms | ✓ 1603ms | 否 | ✓ 1184ms | http |
| 103.74.192.243:7890 | ✓ 1643ms | 否 | ✓ 1206ms | ✓ 1578ms | ✓ 1816ms | http |
| 158.178.237.243:3128 | ✓ 1006ms | 否 | 否 | ✓ 1577ms | ✓ 1265ms | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 1035ms | ✓ 1242ms | ✓ 952ms | http |
| 115.76.5.32:10010 | ✓ 1780ms | 否 | 否 | ✓ 1977ms | ✓ 1932ms | http |
| 103.165.234.46:8080 | ✓ 1544ms | 否 | ✓ 1491ms | ✓ 1660ms | 否 | http |
| 5.129.228.225:1080 | ✓ 663ms | 否 | ✓ 975ms | ✓ 1424ms | 否 | http |
| 45.140.147.155:1082 | ✓ 991ms | ✓ 1201ms | ✓ 1087ms | ✓ 1860ms | ✓ 1336ms | http |
| 103.113.70.189:1081 | 否 | ✓ 930ms | 否 | ✓ 1374ms | ✓ 841ms | http |
| 121.230.9.148:1080 | ✓ 1275ms | 否 | ✓ 1680ms | 否 | ✓ 1739ms | http |
| 118.68.216.26:10009 | 否 | 否 | ✓ 1849ms | ✓ 1964ms | ✓ 1735ms | http |
| 121.230.9.168:1080 | ✓ 1206ms | ✓ 1536ms | ✓ 1982ms | 否 | ✓ 1190ms | http |
| 103.104.99.29:80 | ✓ 1770ms | 否 | ✓ 1805ms | ✓ 1562ms | 否 | http |
| 103.104.99.89:80 | ✓ 1789ms | 否 | ✓ 1682ms | ✓ 1668ms | 否 | http |
| 36.147.78.166:443 | 否 | ✓ 1861ms | ✓ 1978ms | 否 | ✓ 1880ms | http |
| 172.212.68.37:3128 | ✓ 288ms | ✓ 1326ms | ✓ 270ms | ✓ 814ms | ✓ 617ms | http |
| 113.255.59.226:8080 | ✓ 1407ms | 否 | ✓ 1332ms | ✓ 1350ms | ✓ 1369ms | http |
| 46.249.103.192:443 | ✓ 1158ms | 否 | ✓ 1459ms | ✓ 1931ms | 否 | http |

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
