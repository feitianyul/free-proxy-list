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

最后更新：2026-04-13 09:53:48 UTC（2026-04-13 17:53:48 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 746ms | 否 | ✓ 1247ms | ✓ 1203ms | 否 | http |
| 147.161.239.240:8800 | ✓ 597ms | ✓ 1551ms | ✓ 961ms | ✓ 1554ms | ✓ 1201ms | http |
| 34.71.229.255:3128 | ✓ 746ms | 否 | ✓ 1247ms | ✓ 1231ms | ✓ 1259ms | http |
| 147.161.210.140:8800 | ✓ 1764ms | 否 | ✓ 892ms | ✓ 1261ms | ✓ 1155ms | http |
| 167.103.34.108:8800 | ✓ 1613ms | 否 | ✓ 1748ms | 否 | ✓ 1872ms | http |
| 113.160.132.26:8080 | ✓ 1824ms | ✓ 1958ms | ✓ 1410ms | 否 | ✓ 1391ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1751ms | 否 | ✓ 1905ms | ✓ 1348ms | http |
| 5.196.101.18:3128 | ✓ 1811ms | ✓ 1368ms | ✓ 764ms | 否 | 否 | http |
| 95.214.9.93:3128 | ✓ 1276ms | 否 | ✓ 1735ms | 否 | ✓ 1447ms | http |
| 167.103.115.102:8800 | ✓ 1116ms | 否 | ✓ 1369ms | ✓ 1273ms | ✓ 1869ms | http |
| 5.104.87.17:8051 | ✓ 1981ms | 否 | ✓ 1905ms | ✓ 1341ms | ✓ 1195ms | http |
| 120.92.108.86:7890 | ✓ 1417ms | 否 | 否 | ✓ 1980ms | ✓ 1721ms | http |
| 167.103.31.122:8800 | ✓ 1985ms | 否 | ✓ 1646ms | 否 | ✓ 1908ms | http |
| 45.167.125.21:999 | ✓ 979ms | ✓ 1787ms | ✓ 1288ms | ✓ 1622ms | ✓ 1467ms | http |
| 43.156.132.113:3128 | ✓ 921ms | 否 | ✓ 1044ms | ✓ 1254ms | ✓ 1010ms | http |
| 167.103.144.127:8800 | ✓ 1167ms | 否 | ✓ 1303ms | ✓ 1493ms | ✓ 1347ms | http |
| 46.30.46.133:3128 | ✓ 1555ms | ✓ 1203ms | 否 | ✓ 1642ms | ✓ 1307ms | http |
| 138.197.68.35:4857 | ✓ 341ms | ✓ 1478ms | ✓ 966ms | ✓ 1086ms | ✓ 786ms | http |
| 164.92.148.68:3128 | ✓ 1973ms | ✓ 1424ms | ✓ 910ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 389ms | 否 | ✓ 708ms | ✓ 1837ms | 否 | http |
| 181.78.44.63:999 | ✓ 799ms | 否 | 否 | ✓ 1486ms | ✓ 1126ms | http |
| 5.255.123.43:1080 | ✓ 1483ms | ✓ 1363ms | ✓ 1384ms | 否 | 否 | http |
| 178.128.243.121:3128 | ✓ 968ms | 否 | 否 | ✓ 1204ms | ✓ 936ms | http |
| 139.59.222.40:3128 | ✓ 918ms | 否 | ✓ 1525ms | ✓ 1252ms | ✓ 1003ms | http |
| 173.212.246.157:3128 | ✓ 775ms | 否 | ✓ 1319ms | 否 | ✓ 1903ms | http |
| 103.157.200.126:3128 | ✓ 1550ms | 否 | 否 | ✓ 1536ms | ✓ 1239ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1310ms | 否 | ✓ 1388ms | ✓ 1175ms | http |
| 180.103.19.123:1080 | ✓ 1398ms | 否 | ✓ 1873ms | ✓ 1904ms | ✓ 1482ms | http |
| 45.77.246.231:80 | ✓ 1438ms | 否 | ✓ 1371ms | ✓ 1328ms | ✓ 1016ms | http |
| 171.227.167.109:1004 | ✓ 1641ms | 否 | 否 | ✓ 1456ms | ✓ 1123ms | http |
| 8.219.64.245:3128 | 否 | 否 | ✓ 926ms | ✓ 1291ms | ✓ 1031ms | http |
| 171.227.167.109:1008 | 否 | 否 | ✓ 1478ms | ✓ 1523ms | ✓ 1319ms | http |
| 45.140.147.155:1082 | ✓ 928ms | 否 | ✓ 974ms | ✓ 1440ms | ✓ 862ms | http |
| 45.12.151.226:2829 | ✓ 1098ms | 否 | ✓ 1669ms | 否 | ✓ 1736ms | http |
| 171.227.167.109:1006 | ✓ 1257ms | 否 | ✓ 1572ms | ✓ 1518ms | ✓ 1293ms | http |
| 160.238.65.2:3128 | ✓ 454ms | ✓ 1393ms | ✓ 1313ms | ✓ 1944ms | 否 | http |
| 160.238.65.7:3128 | ✓ 459ms | ✓ 1913ms | ✓ 789ms | ✓ 1942ms | 否 | http |
| 213.131.85.28:1976 | ✓ 965ms | ✓ 1828ms | ✓ 1818ms | ✓ 1892ms | ✓ 1934ms | http |
| 59.46.216.131:30001 | ✓ 1292ms | 否 | ✓ 1385ms | ✓ 1513ms | 否 | http |
| 34.85.118.216:3128 | ✓ 741ms | ✓ 1237ms | ✓ 1174ms | ✓ 1232ms | ✓ 843ms | http |
| 81.169.170.253:3128 | ✓ 911ms | 否 | ✓ 1456ms | 否 | ✓ 1984ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1421ms | ✓ 1630ms | ✓ 1391ms | 否 | http |
| 45.140.147.155:1081 | ✓ 891ms | 否 | ✓ 1473ms | ✓ 1534ms | ✓ 1055ms | http |
| 187.216.141.46:3128 | ✓ 1677ms | ✓ 1861ms | ✓ 1062ms | 否 | ✓ 1961ms | http |
| 47.74.226.8:5001 | ✓ 1556ms | 否 | ✓ 1305ms | ✓ 1540ms | 否 | http |
| 121.230.8.89:1080 | 否 | ✓ 1780ms | ✓ 1481ms | ✓ 1775ms | ✓ 1344ms | http |
| 160.238.65.6:3128 | ✓ 506ms | 否 | ✓ 1847ms | 否 | ✓ 1360ms | http |
| 160.238.65.4:3128 | ✓ 505ms | 否 | ✓ 1851ms | 否 | ✓ 1364ms | http |
| 160.238.65.3:3128 | ✓ 506ms | 否 | ✓ 1846ms | 否 | ✓ 1366ms | http |
| 160.238.65.5:3128 | ✓ 366ms | 否 | ✓ 358ms | 否 | ✓ 915ms | http |
| 157.180.50.70:3128 | ✓ 891ms | ✓ 1833ms | ✓ 1830ms | 否 | 否 | http |
| 61.76.95.217:40088 | ✓ 1226ms | ✓ 1493ms | ✓ 1737ms | ✓ 1865ms | ✓ 1379ms | http |
| 121.230.9.125:1080 | ✓ 1645ms | ✓ 1547ms | ✓ 1283ms | ✓ 1795ms | ✓ 1289ms | http |
| 185.76.240.203:10001 | ✓ 541ms | ✓ 1858ms | ✓ 666ms | ✓ 1992ms | ✓ 1391ms | http |
| 185.76.240.64:10001 | ✓ 578ms | ✓ 1745ms | ✓ 721ms | 否 | ✓ 1429ms | http |
| 185.76.240.61:10001 | ✓ 570ms | ✓ 1875ms | ✓ 716ms | ✓ 1922ms | ✓ 1429ms | http |
| 62.113.119.14:8080 | ✓ 581ms | ✓ 1517ms | ✓ 1321ms | ✓ 1816ms | 否 | http |
| 159.223.225.118:8888 | ✓ 1429ms | ✓ 1813ms | 否 | 否 | ✓ 1317ms | http |
| 36.103.198.235:7890 | 否 | ✓ 1899ms | ✓ 1824ms | 否 | ✓ 1626ms | http |
| 168.110.52.228:3128 | ✓ 1788ms | 否 | 否 | ✓ 1300ms | ✓ 1044ms | http |
| 91.233.223.147:3128 | ✓ 793ms | 否 | ✓ 1445ms | 否 | ✓ 1885ms | http |
| 24.144.86.173:1080 | ✓ 1700ms | 否 | ✓ 1346ms | ✓ 1733ms | ✓ 833ms | http |
| 109.107.179.140:31000 | ✓ 1313ms | 否 | ✓ 1020ms | 否 | ✓ 1449ms | http |
| 217.77.102.18:3128 | 否 | ✓ 1831ms | ✓ 949ms | ✓ 1664ms | ✓ 1496ms | http |
| 146.196.97.193:57413 | 否 | 否 | ✓ 1696ms | ✓ 1757ms | ✓ 1744ms | http |
| 107.172.102.234:40621 | ✓ 565ms | ✓ 1098ms | ✓ 1331ms | ✓ 1021ms | ✓ 786ms | http |
| 46.39.105.157:8080 | ✓ 660ms | 否 | ✓ 1152ms | 否 | ✓ 1489ms | http |
| 8.219.97.248:80 | ✓ 1188ms | 否 | ✓ 1377ms | ✓ 1772ms | ✓ 1507ms | http |
| 158.160.215.167:8126 | ✓ 1640ms | ✓ 1789ms | ✓ 1083ms | 否 | 否 | http |
| 162.240.154.26:3128 | ✓ 1241ms | 否 | ✓ 1823ms | ✓ 1354ms | 否 | http |
| 121.135.144.141:8088 | ✓ 1002ms | ✓ 1447ms | ✓ 1239ms | ✓ 1426ms | ✓ 983ms | http |
| 121.230.8.138:1080 | ✓ 1216ms | ✓ 1713ms | ✓ 1271ms | ✓ 1613ms | ✓ 1371ms | http |
| 180.103.19.219:1080 | 否 | 否 | ✓ 1337ms | ✓ 1532ms | ✓ 1437ms | http |
| 62.234.206.73:3128 | 否 | 否 | ✓ 1476ms | ✓ 1991ms | ✓ 1932ms | http |
| 138.124.99.216:8888 | ✓ 1633ms | ✓ 1778ms | ✓ 1822ms | 否 | 否 | http |
| 81.12.96.60:8118 | ✓ 1335ms | 否 | ✓ 1420ms | 否 | ✓ 1857ms | http |
| 82.114.228.67:1080 | 否 | 否 | ✓ 852ms | ✓ 1539ms | ✓ 1145ms | http |
| 5.104.87.17:8050 | ✓ 1326ms | 否 | ✓ 1594ms | ✓ 1419ms | ✓ 926ms | http |
| 103.39.51.207:8080 | ✓ 1735ms | 否 | ✓ 1940ms | 否 | ✓ 1609ms | http |

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
