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

最后更新：2026-03-06 20:36:43 UTC（2026-03-07 04:36:43 UTC+8）

**代理总数：93**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 706ms | ✓ 917ms | ✓ 775ms | ✓ 1096ms | ✓ 821ms | http |
| 1.231.81.166:3128 | ✓ 1768ms | ✓ 1164ms | ✓ 1164ms | ✓ 1085ms | ✓ 885ms | http |
| 194.213.18.200:443 | ✓ 1348ms | ✓ 1536ms | ✓ 1577ms | 否 | 否 | http |
| 178.236.245.17:3128 | ✓ 867ms | ✓ 1492ms | ✓ 1576ms | ✓ 1942ms | ✓ 1504ms | http |
| 125.128.12.144:3128 | ✓ 1773ms | ✓ 1637ms | ✓ 1525ms | 否 | 否 | http |
| 178.236.245.59:3128 | ✓ 938ms | ✓ 1600ms | ✓ 529ms | ✓ 1581ms | ✓ 1221ms | http |
| 168.235.110.63:3128 | ✓ 560ms | 否 | ✓ 805ms | ✓ 1979ms | ✓ 1729ms | http |
| 14.56.177.44:3128 | ✓ 760ms | ✓ 1271ms | ✓ 1229ms | ✓ 1374ms | ✓ 1410ms | http |
| 138.124.53.25:7443 | ✓ 1380ms | 否 | ✓ 1926ms | ✓ 1541ms | ✓ 1876ms | http |
| 211.171.114.154:3128 | ✓ 1190ms | 否 | ✓ 1727ms | ✓ 1595ms | ✓ 1697ms | http |
| 61.72.221.234:3128 | ✓ 828ms | ✓ 1494ms | ✓ 1219ms | 否 | 否 | http |
| 121.230.8.11:1080 | ✓ 1495ms | ✓ 1777ms | ✓ 1579ms | 否 | 否 | http |
| 136.49.39.94:8888 | ✓ 779ms | ✓ 1361ms | ✓ 611ms | ✓ 1763ms | ✓ 1167ms | http |
| 85.9.195.140:1080 | ✓ 253ms | 否 | ✓ 1352ms | ✓ 1807ms | ✓ 1907ms | http |
| 38.47.97.22:6005 | ✓ 1651ms | ✓ 1654ms | 否 | ✓ 1130ms | ✓ 985ms | http |
| 167.172.69.123:8080 | ✓ 1071ms | 否 | ✓ 1864ms | ✓ 1267ms | ✓ 987ms | http |
| 125.128.12.14:3128 | ✓ 1628ms | ✓ 1544ms | ✓ 1406ms | 否 | 否 | http |
| 46.249.103.192:443 | ✓ 558ms | 否 | ✓ 1638ms | ✓ 1675ms | 否 | http |
| 67.169.98.211:443 | ✓ 1783ms | 否 | ✓ 1924ms | ✓ 1543ms | 否 | http |
| 35.225.22.61:80 | ✓ 559ms | 否 | ✓ 333ms | ✓ 1008ms | ✓ 811ms | http |
| 193.108.118.190:8888 | ✓ 1474ms | 否 | ✓ 461ms | ✓ 1594ms | ✓ 1234ms | http |
| 167.172.69.123:80 | ✓ 979ms | 否 | ✓ 1838ms | ✓ 1333ms | ✓ 1047ms | http |
| 5.129.206.247:8888 | ✓ 940ms | 否 | ✓ 1696ms | 否 | ✓ 1906ms | http |
| 81.70.169.194:80 | ✓ 1169ms | ✓ 1487ms | ✓ 1304ms | 否 | ✓ 1264ms | http |
| 101.43.255.96:80 | ✓ 1131ms | ✓ 1550ms | ✓ 1815ms | 否 | ✓ 1205ms | http |
| 190.12.150.244:999 | ✓ 1148ms | 否 | ✓ 1012ms | ✓ 1884ms | ✓ 1457ms | http |
| 42.115.72.27:2046 | 否 | 否 | ✓ 1780ms | ✓ 1991ms | ✓ 1741ms | http |
| 91.193.240.157:9877 | ✓ 1723ms | 否 | ✓ 1812ms | 否 | ✓ 1752ms | http |
| 115.231.181.40:8128 | ✓ 1416ms | ✓ 1722ms | ✓ 1152ms | 否 | 否 | http |
| 101.47.73.135:3128 | ✓ 1853ms | 否 | ✓ 1170ms | 否 | ✓ 1802ms | http |
| 42.115.72.27:2102 | 否 | 否 | ✓ 1807ms | ✓ 1992ms | ✓ 1762ms | http |
| 165.227.5.10:8888 | ✓ 299ms | ✓ 1195ms | 否 | ✓ 1629ms | 否 | http |
| 46.183.25.8:443 | 否 | 否 | ✓ 1530ms | ✓ 1597ms | ✓ 981ms | http |
| 14.56.107.244:3128 | ✓ 960ms | ✓ 1689ms | ✓ 1542ms | ✓ 1312ms | 否 | http |
| 134.209.153.66:3128 | ✓ 1055ms | 否 | ✓ 1614ms | ✓ 1523ms | ✓ 1256ms | http |
| 154.90.48.209:9090 | ✓ 1666ms | 否 | ✓ 1989ms | ✓ 1880ms | ✓ 1246ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1862ms | ✓ 1888ms | ✓ 1435ms | 否 | http |
| 42.115.72.27:2049 | ✓ 1701ms | 否 | ✓ 1942ms | ✓ 1923ms | 否 | http |
| 42.115.72.27:2039 | ✓ 1698ms | 否 | ✓ 1707ms | 否 | ✓ 1760ms | http |
| 202.155.12.161:443 | ✓ 1735ms | 否 | ✓ 1817ms | ✓ 1316ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1590ms | ✓ 1135ms | ✓ 1489ms | ✓ 1131ms | ✓ 883ms | http |
| 121.230.9.64:1080 | ✓ 1499ms | ✓ 1825ms | ✓ 1232ms | ✓ 1995ms | ✓ 1628ms | http |
| 94.176.3.43:7443 | ✓ 1836ms | 否 | ✓ 1453ms | 否 | ✓ 1698ms | http |
| 61.72.110.94:3128 | ✓ 1263ms | 否 | 否 | ✓ 1374ms | ✓ 1691ms | http |
| 188.132.141.249:443 | ✓ 1471ms | 否 | ✓ 1452ms | 否 | ✓ 1984ms | http |
| 107.174.80.186:3128 | ✓ 840ms | ✓ 1052ms | ✓ 916ms | ✓ 1080ms | ✓ 911ms | http |
| 42.96.16.158:1311 | ✓ 1268ms | 否 | ✓ 1446ms | ✓ 1659ms | ✓ 1640ms | http |
| 91.233.223.147:3128 | ✓ 744ms | ✓ 1902ms | ✓ 710ms | 否 | 否 | http |
| 61.72.221.194:3128 | ✓ 793ms | 否 | ✓ 1718ms | 否 | ✓ 1456ms | http |
| 193.168.173.136:443 | ✓ 514ms | 否 | ✓ 966ms | ✓ 1667ms | 否 | http |
| 192.166.82.55:1080 | ✓ 735ms | 否 | ✓ 1456ms | 否 | ✓ 1785ms | http |
| 88.80.150.82:8080 | ✓ 970ms | ✓ 1898ms | 否 | ✓ 1533ms | 否 | https |
| 121.128.121.54:3128 | ✓ 1788ms | 否 | 否 | ✓ 1211ms | ✓ 1962ms | http |
| 69.48.179.20:3128 | ✓ 1609ms | 否 | ✓ 1403ms | ✓ 1315ms | 否 | http |
| 121.230.9.231:1080 | ✓ 1300ms | ✓ 1524ms | ✓ 1261ms | ✓ 1664ms | ✓ 1451ms | http |
| 221.127.195.224:8888 | ✓ 1476ms | 否 | ✓ 1330ms | ✓ 1570ms | ✓ 1316ms | http |
| 180.76.115.231:3128 | ✓ 1848ms | ✓ 1574ms | ✓ 1227ms | ✓ 1805ms | ✓ 1231ms | http |
| 121.230.9.205:1080 | ✓ 1317ms | 否 | ✓ 1419ms | 否 | ✓ 1605ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1718ms | ✓ 1897ms | ✓ 1449ms | http |
| 42.115.72.27:2065 | ✓ 1730ms | 否 | ✓ 1728ms | ✓ 1920ms | ✓ 1836ms | http |
| 185.191.236.162:3128 | ✓ 638ms | ✓ 1547ms | ✓ 612ms | ✓ 1502ms | 否 | http |
| 14.225.222.164:7890 | ✓ 1682ms | 否 | ✓ 1326ms | ✓ 1419ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1177ms | 否 | ✓ 1473ms | ✓ 1696ms | 否 | http |
| 107.152.32.98:1305 | ✓ 1001ms | ✓ 1962ms | ✓ 1253ms | ✓ 1699ms | ✓ 1530ms | http |
| 103.84.95.54:7890 | ✓ 863ms | 否 | ✓ 1078ms | ✓ 1224ms | 否 | http |
| 121.230.8.95:1080 | ✓ 1201ms | ✓ 1562ms | ✓ 1334ms | ✓ 1683ms | ✓ 1238ms | http |
| 45.136.198.40:3128 | ✓ 1713ms | 否 | ✓ 1170ms | 否 | ✓ 1913ms | http |
| 116.80.82.230:3172 | ✓ 1688ms | 否 | ✓ 1721ms | 否 | ✓ 1831ms | http |
| 120.92.212.16:8890 | ✓ 1107ms | 否 | ✓ 1236ms | 否 | ✓ 1943ms | http |
| 59.46.216.131:30001 | ✓ 1180ms | ✓ 1661ms | ✓ 1285ms | 否 | ✓ 1270ms | http |
| 113.176.92.71:3128 | ✓ 1755ms | ✓ 1512ms | ✓ 1377ms | ✓ 1519ms | ✓ 1077ms | http |
| 162.248.165.72:1080 | ✓ 567ms | ✓ 1575ms | ✓ 373ms | 否 | 否 | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1973ms | ✓ 1554ms | ✓ 1828ms | http |
| 103.82.23.118:5178 | ✓ 1683ms | 否 | ✓ 1437ms | 否 | ✓ 1722ms | http |
| 202.129.206.239:3128 | ✓ 1839ms | 否 | 否 | ✓ 1789ms | ✓ 1517ms | http |
| 45.140.147.155:1081 | ✓ 1362ms | ✓ 1093ms | ✓ 406ms | ✓ 1141ms | ✓ 1126ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1825ms | 否 | ✓ 1130ms | ✓ 829ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1681ms | ✓ 1759ms | ✓ 1513ms | http |
| 123.57.0.163:8888 | ✓ 1934ms | 否 | ✓ 1854ms | ✓ 1758ms | 否 | http |
| 5.252.33.13:2025 | 否 | 否 | ✓ 1714ms | ✓ 1997ms | ✓ 1619ms | http |
| 66.245.197.108:9002 | ✓ 870ms | 否 | ✓ 1226ms | ✓ 1934ms | ✓ 1731ms | http |
| 66.245.197.108:9004 | ✓ 843ms | 否 | ✓ 1239ms | 否 | ✓ 1847ms | http |
| 91.107.175.112:10801 | ✓ 812ms | 否 | ✓ 1331ms | ✓ 1776ms | ✓ 1440ms | http |
| 45.186.6.104:3128 | ✓ 1616ms | ✓ 1852ms | ✓ 1957ms | 否 | 否 | http |
| 147.75.202.36:10006 | ✓ 971ms | ✓ 997ms | 否 | ✓ 1454ms | 否 | http |
| 5.9.55.221:5000 | ✓ 582ms | 否 | ✓ 1276ms | ✓ 1936ms | 否 | http |
| 61.72.221.94:3128 | 否 | 否 | ✓ 1540ms | ✓ 1597ms | ✓ 1471ms | http |
| 194.59.204.87:9080 | ✓ 425ms | ✓ 1621ms | ✓ 1306ms | ✓ 1878ms | 否 | http |
| 116.80.82.218:3172 | ✓ 1759ms | 否 | ✓ 1642ms | 否 | ✓ 1839ms | http |
| 116.80.82.232:3172 | ✓ 1758ms | 否 | ✓ 1655ms | ✓ 1995ms | 否 | http |
| 47.101.149.27:9010 | 否 | 否 | ✓ 1578ms | ✓ 1774ms | ✓ 1615ms | http |
| 116.80.82.217:3172 | ✓ 1715ms | 否 | ✓ 1729ms | 否 | ✓ 1814ms | http |
| 159.223.42.219:3128 | ✓ 1125ms | 否 | ✓ 1032ms | ✓ 1942ms | 否 | http |

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
