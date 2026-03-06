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

最后更新：2026-03-06 18:37:22 UTC（2026-03-07 02:37:22 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 175ms | ✓ 1817ms | ✓ 941ms | ✓ 1085ms | ✓ 807ms | http |
| 152.42.195.165:8888 | ✓ 1060ms | 否 | ✓ 1120ms | ✓ 1371ms | ✓ 1088ms | http |
| 1.231.81.166:3128 | ✓ 1821ms | ✓ 1355ms | 否 | ✓ 1074ms | ✓ 950ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1998ms | ✓ 1960ms | ✓ 1765ms | http |
| 168.235.110.63:3128 | 否 | 否 | ✓ 925ms | ✓ 1078ms | ✓ 986ms | http |
| 107.174.80.186:3128 | ✓ 785ms | ✓ 1875ms | 否 | ✓ 1739ms | 否 | http |
| 101.43.255.96:80 | ✓ 1162ms | ✓ 1489ms | 否 | ✓ 1568ms | 否 | http |
| 103.82.23.118:5185 | 否 | 否 | ✓ 1551ms | ✓ 1607ms | ✓ 1359ms | http |
| 67.169.98.211:443 | ✓ 1618ms | 否 | ✓ 1823ms | 否 | ✓ 1659ms | http |
| 46.183.25.8:443 | ✓ 1490ms | 否 | ✓ 590ms | ✓ 1053ms | 否 | http |
| 194.213.18.200:443 | ✓ 1043ms | 否 | ✓ 496ms | ✓ 929ms | ✓ 1824ms | http |
| 85.9.195.140:1080 | ✓ 769ms | ✓ 1230ms | ✓ 1423ms | ✓ 1572ms | 否 | http |
| 167.172.69.123:80 | 否 | 否 | ✓ 1514ms | ✓ 1649ms | ✓ 1423ms | http |
| 167.172.69.123:8080 | 否 | 否 | ✓ 1515ms | ✓ 1693ms | ✓ 1361ms | http |
| 81.70.169.194:80 | ✓ 1180ms | 否 | 否 | ✓ 1702ms | ✓ 1567ms | http |
| 107.152.32.98:1305 | 否 | ✓ 1839ms | ✓ 1466ms | 否 | ✓ 1707ms | http |
| 35.225.22.61:80 | ✓ 597ms | ✓ 1178ms | ✓ 1168ms | ✓ 1098ms | ✓ 891ms | http |
| 136.49.39.94:8888 | 否 | 否 | ✓ 1460ms | ✓ 1595ms | ✓ 1274ms | http |
| 91.193.240.157:9877 | ✓ 886ms | 否 | ✓ 1177ms | 否 | ✓ 1777ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 523ms | ✓ 975ms | ✓ 869ms | http |
| 113.176.92.71:3128 | ✓ 1821ms | 否 | ✓ 1435ms | ✓ 1545ms | ✓ 1162ms | http |
| 14.225.222.164:7890 | ✓ 1838ms | ✓ 1662ms | 否 | 否 | ✓ 1766ms | http |
| 5.252.33.13:2025 | 否 | 否 | ✓ 1721ms | ✓ 1955ms | ✓ 1612ms | http |
| 121.128.121.54:3128 | ✓ 1876ms | 否 | 否 | ✓ 1224ms | ✓ 979ms | http |
| 91.233.223.147:3128 | ✓ 784ms | 否 | ✓ 949ms | ✓ 1919ms | ✓ 1504ms | http |
| 88.80.150.82:8080 | ✓ 1316ms | 否 | ✓ 1960ms | ✓ 1935ms | ✓ 1673ms | https |
| 46.249.103.192:443 | ✓ 1778ms | 否 | ✓ 1455ms | ✓ 1983ms | 否 | http |
| 162.248.165.72:1080 | ✓ 870ms | 否 | ✓ 1177ms | 否 | ✓ 1912ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1818ms | 否 | ✓ 1361ms | ✓ 1284ms | http |
| 42.115.72.27:2039 | ✓ 1768ms | 否 | ✓ 1737ms | ✓ 1976ms | 否 | http |
| 42.115.72.27:2049 | ✓ 1664ms | 否 | ✓ 1905ms | ✓ 1931ms | ✓ 1675ms | http |
| 42.96.16.158:1311 | ✓ 1600ms | 否 | ✓ 1158ms | ✓ 1426ms | ✓ 1114ms | http |
| 138.124.53.25:7443 | ✓ 1630ms | 否 | ✓ 1903ms | 否 | ✓ 1698ms | http |
| 103.84.95.54:7890 | ✓ 830ms | 否 | 否 | ✓ 1626ms | ✓ 894ms | http |
| 14.56.107.244:3128 | ✓ 787ms | ✓ 1288ms | ✓ 1787ms | 否 | 否 | http |
| 125.128.12.144:3128 | ✓ 829ms | ✓ 1965ms | ✓ 1626ms | ✓ 1226ms | ✓ 972ms | http |
| 61.72.221.94:3128 | ✓ 1173ms | 否 | ✓ 1266ms | 否 | ✓ 1527ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1071ms | ✓ 1478ms | ✓ 1148ms | http |
| 42.115.72.27:2065 | ✓ 1955ms | 否 | ✓ 1733ms | 否 | ✓ 1709ms | http |
| 192.166.82.55:1080 | ✓ 662ms | ✓ 1665ms | 否 | ✓ 1388ms | ✓ 1421ms | http |
| 207.254.71.62:8088 | ✓ 792ms | 否 | ✓ 1198ms | ✓ 1826ms | 否 | http |
| 103.113.70.189:1081 | 否 | ✓ 1151ms | 否 | ✓ 1467ms | ✓ 946ms | http |
| 59.46.216.131:30001 | ✓ 1133ms | ✓ 1593ms | ✓ 1341ms | 否 | 否 | http |
| 118.113.246.177:1080 | ✓ 1425ms | ✓ 1788ms | ✓ 1366ms | 否 | ✓ 1546ms | http |
| 201.144.25.226:3128 | ✓ 765ms | ✓ 1408ms | ✓ 1226ms | ✓ 1320ms | 否 | http |
| 178.213.25.221:7890 | ✓ 1264ms | 否 | ✓ 1111ms | ✓ 1715ms | ✓ 1267ms | http |
| 157.230.38.173:3128 | 否 | 否 | ✓ 1306ms | ✓ 1265ms | ✓ 1007ms | http |
| 103.139.138.194:3128 | ✓ 1939ms | 否 | ✓ 1346ms | ✓ 1647ms | ✓ 1277ms | http |
| 180.76.115.231:3128 | ✓ 1265ms | ✓ 1490ms | 否 | ✓ 1629ms | 否 | http |
| 39.104.201.40:7890 | ✓ 1851ms | 否 | ✓ 1124ms | 否 | ✓ 1392ms | http |
| 45.136.198.40:3128 | ✓ 968ms | ✓ 1420ms | ✓ 873ms | 否 | ✓ 1824ms | http |
| 162.240.154.26:3128 | ✓ 1178ms | 否 | ✓ 1221ms | ✓ 1492ms | 否 | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1503ms | ✓ 1347ms | ✓ 1227ms | http |
| 120.92.212.16:7890 | ✓ 1134ms | ✓ 1449ms | 否 | 否 | ✓ 1155ms | http |
| 45.140.147.155:1081 | ✓ 469ms | ✓ 1642ms | ✓ 1454ms | ✓ 1835ms | ✓ 1461ms | http |
| 61.72.110.94:3128 | 否 | 否 | ✓ 1338ms | ✓ 1295ms | ✓ 1964ms | http |
| 193.168.173.136:443 | ✓ 929ms | ✓ 1730ms | 否 | 否 | ✓ 1616ms | http |
| 103.82.23.118:5247 | 否 | 否 | ✓ 1449ms | ✓ 1623ms | ✓ 1408ms | http |
| 103.215.36.88:19655 | ✓ 1569ms | ✓ 1357ms | ✓ 1214ms | ✓ 1466ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1494ms | 否 | ✓ 1839ms | ✓ 1427ms | 否 | http |
| 45.129.141.143:3128 | ✓ 624ms | 否 | ✓ 1605ms | ✓ 1690ms | ✓ 1465ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1638ms | ✓ 1424ms | ✓ 1065ms | http |
| 103.39.51.190:8080 | ✓ 1904ms | 否 | 否 | ✓ 1634ms | ✓ 1624ms | http |
| 2.56.178.131:443 | ✓ 1024ms | 否 | ✓ 1043ms | ✓ 1750ms | 否 | http |
| 94.176.3.43:7443 | ✓ 1631ms | ✓ 1985ms | ✓ 1630ms | 否 | ✓ 1772ms | http |
| 188.132.141.249:443 | ✓ 760ms | 否 | ✓ 1270ms | 否 | ✓ 1986ms | http |
| 103.104.99.89:80 | ✓ 1964ms | 否 | ✓ 1895ms | ✓ 1855ms | ✓ 1759ms | http |
| 109.234.38.35:3128 | ✓ 537ms | 否 | ✓ 1908ms | ✓ 1887ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1180ms | 否 | ✓ 1993ms | ✓ 1939ms | 否 | http |
| 190.12.150.244:999 | ✓ 1352ms | ✓ 1728ms | ✓ 1155ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1025ms | ✓ 1360ms | ✓ 1088ms | ✓ 1418ms | ✓ 1069ms | http |
| 61.72.110.54:3128 | ✓ 1712ms | 否 | 否 | ✓ 1915ms | ✓ 1568ms | http |
| 35.180.127.14:1001 | ✓ 1126ms | 否 | ✓ 1981ms | 否 | ✓ 1993ms | http |
| 103.215.36.88:17597 | ✓ 1134ms | ✓ 1415ms | ✓ 1232ms | ✓ 1466ms | ✓ 1113ms | http |
| 103.215.36.88:17780 | ✓ 1322ms | ✓ 1391ms | ✓ 1218ms | ✓ 1431ms | ✓ 1129ms | http |

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
