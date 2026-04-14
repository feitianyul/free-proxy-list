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

最后更新：2026-04-14 18:08:57 UTC（2026-04-15 02:08:57 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 857ms | ✓ 1368ms | ✓ 985ms | ✓ 1681ms | ✓ 1198ms | http |
| 185.114.73.2:1080 | ✓ 889ms | ✓ 1912ms | ✓ 1539ms | 否 | ✓ 1300ms | http |
| 113.160.132.26:8080 | ✓ 1827ms | 否 | ✓ 1387ms | ✓ 1993ms | ✓ 1080ms | http |
| 72.56.84.21:1080 | ✓ 476ms | 否 | ✓ 1584ms | ✓ 1927ms | 否 | http |
| 78.11.96.22:8888 | ✓ 1087ms | 否 | ✓ 798ms | ✓ 1471ms | ✓ 1358ms | http |
| 218.108.131.186:17890 | ✓ 945ms | ✓ 1149ms | ✓ 976ms | ✓ 1230ms | ✓ 1024ms | http |
| 45.149.92.147:5001 | ✓ 1232ms | 否 | ✓ 799ms | ✓ 1013ms | ✓ 782ms | http |
| 36.141.21.200:7890 | ✓ 1132ms | ✓ 1297ms | ✓ 1038ms | 否 | ✓ 1098ms | http |
| 167.103.34.108:8800 | ✓ 1308ms | 否 | ✓ 1338ms | ✓ 1486ms | 否 | http |
| 36.103.198.235:7890 | ✓ 1106ms | 否 | ✓ 1258ms | ✓ 1762ms | ✓ 1549ms | http |
| 167.103.144.127:8800 | ✓ 1699ms | 否 | ✓ 1379ms | ✓ 1943ms | ✓ 1853ms | http |
| 120.92.108.86:7890 | ✓ 1350ms | 否 | 否 | ✓ 1808ms | ✓ 1485ms | http |
| 45.167.125.21:999 | ✓ 1452ms | 否 | ✓ 1367ms | ✓ 1672ms | ✓ 1498ms | http |
| 144.31.27.49:1080 | ✓ 793ms | ✓ 1984ms | 否 | ✓ 1995ms | ✓ 1851ms | http |
| 167.103.31.122:8800 | ✓ 1358ms | 否 | ✓ 1284ms | 否 | ✓ 1599ms | http |
| 185.76.240.95:10002 | ✓ 1462ms | 否 | ✓ 1128ms | ✓ 1953ms | ✓ 1447ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 871ms | ✓ 1056ms | ✓ 766ms | http |
| 45.140.147.155:1082 | ✓ 552ms | ✓ 1562ms | ✓ 764ms | 否 | 否 | http |
| 147.161.239.240:8800 | ✓ 1013ms | ✓ 1662ms | ✓ 713ms | ✓ 1594ms | ✓ 1413ms | http |
| 103.85.113.66:9999 | ✓ 1039ms | ✓ 1571ms | ✓ 1750ms | ✓ 1973ms | ✓ 1617ms | http |
| 91.193.240.157:9877 | ✓ 1072ms | 否 | ✓ 1920ms | 否 | ✓ 1740ms | http |
| 137.59.47.73:3128 | ✓ 1351ms | 否 | ✓ 1769ms | ✓ 1772ms | 否 | http |
| 140.238.242.189:8100 | ✓ 899ms | 否 | ✓ 850ms | ✓ 1781ms | ✓ 1114ms | http |
| 138.124.99.216:8888 | ✓ 707ms | ✓ 1910ms | 否 | 否 | ✓ 1869ms | http |
| 59.46.216.131:30001 | ✓ 1148ms | ✓ 1455ms | ✓ 1291ms | ✓ 1963ms | ✓ 1169ms | http |
| 114.118.82.146:80 | ✓ 1231ms | ✓ 1357ms | ✓ 1251ms | ✓ 1344ms | ✓ 1075ms | http |
| 212.58.132.5:8888 | ✓ 1843ms | 否 | ✓ 1593ms | ✓ 1671ms | ✓ 1688ms | http |
| 38.210.179.29:999 | ✓ 517ms | 否 | ✓ 735ms | ✓ 1571ms | ✓ 1146ms | http |
| 116.80.95.227:3172 | 否 | 否 | ✓ 1650ms | ✓ 1962ms | ✓ 1822ms | http |
| 125.64.244.100:8889 | ✓ 1873ms | ✓ 1897ms | ✓ 1791ms | 否 | 否 | http |
| 207.180.228.55:80 | ✓ 1154ms | 否 | ✓ 729ms | 否 | ✓ 1564ms | http |
| 167.103.115.102:8800 | ✓ 1060ms | 否 | ✓ 1177ms | ✓ 1177ms | ✓ 1157ms | http |
| 217.217.249.160:8080 | ✓ 1169ms | 否 | ✓ 1104ms | 否 | ✓ 1174ms | http |
| 1.231.81.166:3128 | ✓ 876ms | 否 | 否 | ✓ 1205ms | ✓ 1547ms | http |
| 130.61.30.221:8080 | ✓ 618ms | 否 | ✓ 718ms | ✓ 1689ms | ✓ 1994ms | http |
| 210.223.44.230:3128 | ✓ 702ms | ✓ 1477ms | ✓ 1867ms | 否 | 否 | http |
| 103.3.246.71:3128 | ✓ 1500ms | 否 | ✓ 1346ms | ✓ 1601ms | ✓ 1383ms | http |
| 8.219.195.129:1080 | ✓ 821ms | ✓ 1834ms | ✓ 1069ms | ✓ 1186ms | ✓ 928ms | http |
| 62.113.119.14:8080 | ✓ 572ms | ✓ 1612ms | ✓ 1139ms | ✓ 1540ms | ✓ 1864ms | http |
| 139.227.17.70:17890 | ✓ 943ms | ✓ 1188ms | ✓ 1003ms | ✓ 1253ms | ✓ 1044ms | http |
| 181.78.44.63:999 | 否 | ✓ 1616ms | ✓ 1153ms | ✓ 1465ms | ✓ 1258ms | http |
| 103.122.1.130:8080 | ✓ 1885ms | 否 | ✓ 1875ms | ✓ 1645ms | ✓ 1606ms | http |
| 103.171.161.96:9090 | ✓ 1803ms | ✓ 1950ms | ✓ 1786ms | ✓ 1539ms | ✓ 1471ms | http |
| 85.239.59.252:7890 | ✓ 1010ms | ✓ 1841ms | ✓ 827ms | ✓ 1981ms | 否 | http |
| 185.191.236.162:3128 | ✓ 675ms | ✓ 1663ms | ✓ 1675ms | ✓ 1832ms | 否 | http |
| 12.89.176.82:3128 | ✓ 311ms | ✓ 1085ms | ✓ 712ms | ✓ 967ms | ✓ 867ms | http |
| 94.131.118.129:1081 | ✓ 1147ms | ✓ 1938ms | ✓ 872ms | 否 | ✓ 1767ms | http |
| 45.12.151.226:2829 | ✓ 1133ms | 否 | ✓ 1782ms | 否 | ✓ 1397ms | http |
| 218.153.163.156:8123 | ✓ 938ms | ✓ 1281ms | 否 | 否 | ✓ 952ms | http |
| 147.45.214.210:1080 | ✓ 1549ms | 否 | ✓ 1549ms | 否 | ✓ 1697ms | http |
| 190.12.150.244:999 | ✓ 1313ms | ✓ 1710ms | ✓ 1724ms | 否 | 否 | http |
| 65.108.203.37:18080 | ✓ 1444ms | ✓ 1810ms | 否 | ✓ 1985ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1489ms | ✓ 1415ms | ✓ 1154ms | ✓ 1518ms | ✓ 1565ms | http |
| 94.131.118.39:1081 | 否 | 否 | ✓ 974ms | ✓ 1922ms | ✓ 1973ms | http |
| 157.230.178.216:8088 | ✓ 742ms | ✓ 1701ms | 否 | 否 | ✓ 1664ms | http |
| 103.113.70.189:1081 | ✓ 910ms | 否 | ✓ 814ms | ✓ 1069ms | ✓ 751ms | http |
| 104.168.93.120:8080 | ✓ 901ms | ✓ 1601ms | ✓ 1228ms | 否 | 否 | http |
| 139.159.99.242:8080 | ✓ 899ms | 否 | ✓ 944ms | ✓ 1195ms | ✓ 1042ms | http |
| 157.230.38.173:3128 | ✓ 1620ms | 否 | ✓ 1214ms | ✓ 1198ms | ✓ 980ms | http |
| 144.31.52.77:3128 | ✓ 1133ms | ✓ 1771ms | ✓ 1371ms | 否 | ✓ 1657ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 651ms | ✓ 937ms | ✓ 1808ms | http |
| 5.255.123.43:1080 | ✓ 1199ms | 否 | ✓ 745ms | 否 | ✓ 1258ms | http |
| 121.230.8.171:1080 | ✓ 1284ms | ✓ 1561ms | ✓ 1337ms | ✓ 1673ms | ✓ 1304ms | http |
| 101.32.243.189:80 | ✓ 1477ms | ✓ 1886ms | ✓ 1582ms | ✓ 1536ms | ✓ 1416ms | http |
| 121.230.8.153:1080 | ✓ 1026ms | ✓ 1685ms | ✓ 1255ms | ✓ 1707ms | ✓ 1712ms | http |
| 103.191.165.219:8085 | ✓ 1812ms | 否 | ✓ 1896ms | ✓ 1561ms | ✓ 1557ms | http |
| 61.52.131.172:8443 | ✓ 986ms | ✓ 1323ms | ✓ 1082ms | ✓ 1357ms | ✓ 1091ms | http |
| 177.234.217.88:999 | ✓ 1691ms | ✓ 1991ms | ✓ 1993ms | 否 | 否 | http |
| 82.114.228.67:1080 | ✓ 1151ms | ✓ 1528ms | ✓ 850ms | ✓ 1568ms | 否 | http |
| 138.2.47.198:5858 | ✓ 1680ms | 否 | ✓ 1822ms | ✓ 1874ms | ✓ 1661ms | http |
| 218.153.163.186:8800 | ✓ 1747ms | ✓ 1204ms | ✓ 974ms | ✓ 1368ms | ✓ 994ms | http |
| 103.39.51.207:8080 | ✓ 1532ms | 否 | 否 | ✓ 1704ms | ✓ 1526ms | http |

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
