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

最后更新：2026-04-14 16:10:33 UTC（2026-04-15 00:10:33 UTC+8）

**代理总数：63**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 845ms | 否 | ✓ 888ms | ✓ 1144ms | ✓ 1180ms | http |
| 147.161.239.240:8800 | ✓ 1265ms | ✓ 1907ms | ✓ 1322ms | ✓ 1893ms | ✓ 1562ms | http |
| 167.103.115.102:8800 | ✓ 1534ms | 否 | ✓ 983ms | 否 | ✓ 1261ms | http |
| 113.160.132.26:8080 | ✓ 1513ms | 否 | ✓ 1349ms | 否 | ✓ 997ms | http |
| 167.103.34.108:8800 | ✓ 1916ms | 否 | ✓ 1548ms | ✓ 1680ms | ✓ 1599ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1572ms | ✓ 1995ms | 否 | ✓ 1345ms | http |
| 34.71.229.255:3128 | ✓ 1279ms | 否 | ✓ 1098ms | ✓ 1103ms | ✓ 894ms | http |
| 167.103.144.127:8800 | ✓ 1197ms | 否 | ✓ 1084ms | ✓ 1547ms | ✓ 1440ms | http |
| 72.56.84.21:1080 | ✓ 730ms | ✓ 1803ms | ✓ 1805ms | ✓ 1495ms | ✓ 1300ms | http |
| 85.239.59.252:7890 | ✓ 1659ms | ✓ 1766ms | ✓ 1290ms | 否 | 否 | http |
| 78.11.96.22:8888 | ✓ 1426ms | 否 | ✓ 1240ms | ✓ 1891ms | ✓ 1677ms | http |
| 45.167.125.21:999 | ✓ 1361ms | ✓ 1942ms | ✓ 1384ms | ✓ 1966ms | ✓ 1658ms | http |
| 201.219.22.2:3128 | ✓ 1436ms | ✓ 1668ms | ✓ 1781ms | ✓ 1996ms | ✓ 1567ms | http |
| 144.31.188.184:3128 | 否 | ✓ 1781ms | ✓ 1660ms | 否 | ✓ 1477ms | http |
| 171.227.167.109:1008 | ✓ 916ms | 否 | ✓ 941ms | ✓ 1537ms | ✓ 956ms | http |
| 167.103.31.122:8800 | ✓ 1321ms | 否 | ✓ 1449ms | ✓ 1634ms | ✓ 1565ms | http |
| 20.27.11.248:8561 | ✓ 483ms | ✓ 1597ms | ✓ 1082ms | ✓ 1633ms | ✓ 1816ms | http |
| 171.234.72.25:17767 | ✓ 984ms | 否 | ✓ 952ms | ✓ 1145ms | ✓ 882ms | http |
| 35.225.22.61:80 | 否 | ✓ 1271ms | ✓ 1238ms | ✓ 1293ms | 否 | http |
| 138.124.99.216:8888 | ✓ 1473ms | ✓ 1843ms | ✓ 1626ms | ✓ 1980ms | ✓ 1645ms | http |
| 103.113.70.189:1081 | ✓ 1117ms | ✓ 1807ms | ✓ 535ms | 否 | ✓ 917ms | http |
| 181.78.44.63:999 | ✓ 1147ms | 否 | ✓ 1416ms | ✓ 1606ms | ✓ 1384ms | http |
| 146.59.16.105:3128 | ✓ 1786ms | 否 | ✓ 1888ms | 否 | ✓ 1702ms | http |
| 137.59.47.73:3128 | ✓ 1838ms | 否 | 否 | ✓ 1786ms | ✓ 1258ms | http |
| 103.126.87.251:8083 | ✓ 1866ms | 否 | ✓ 1650ms | ✓ 1605ms | ✓ 1351ms | http |
| 58.147.190.5:8799 | ✓ 1844ms | 否 | 否 | ✓ 1451ms | ✓ 1701ms | http |
| 120.92.108.86:7890 | ✓ 1447ms | 否 | 否 | ✓ 1967ms | ✓ 1450ms | http |
| 36.103.198.235:7890 | ✓ 930ms | ✓ 1217ms | ✓ 1741ms | 否 | 否 | http |
| 139.159.99.242:8080 | 否 | ✓ 1010ms | ✓ 795ms | ✓ 1026ms | ✓ 868ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1758ms | ✓ 1403ms | ✓ 1537ms | http |
| 210.223.44.230:3128 | ✓ 1910ms | 否 | 否 | ✓ 1853ms | ✓ 1060ms | http |
| 5.255.123.43:1080 | ✓ 726ms | 否 | ✓ 1682ms | ✓ 1807ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1010ms | 否 | ✓ 1732ms | ✓ 1860ms | 否 | http |
| 116.80.96.120:3128 | ✓ 1924ms | 否 | ✓ 1940ms | ✓ 1768ms | ✓ 1927ms | http |
| 117.86.6.244:1080 | 否 | ✓ 1152ms | 否 | ✓ 1983ms | ✓ 1695ms | http |
| 171.227.167.109:1003 | ✓ 1541ms | 否 | ✓ 1702ms | ✓ 1296ms | ✓ 1178ms | http |
| 185.76.240.95:10002 | ✓ 1017ms | 否 | ✓ 934ms | 否 | ✓ 1823ms | http |
| 126.209.18.142:8082 | ✓ 1538ms | 否 | ✓ 1458ms | ✓ 1539ms | ✓ 1541ms | http |
| 103.171.161.96:9090 | ✓ 1991ms | 否 | ✓ 1780ms | ✓ 1756ms | ✓ 1245ms | http |
| 12.89.176.82:3128 | ✓ 826ms | ✓ 1430ms | ✓ 1199ms | ✓ 1775ms | ✓ 1360ms | http |
| 157.0.142.246:10061 | 否 | ✓ 1231ms | ✓ 999ms | ✓ 1286ms | ✓ 1028ms | http |
| 171.244.130.36:3128 | ✓ 1308ms | ✓ 1920ms | ✓ 1379ms | ✓ 1564ms | ✓ 1294ms | http |
| 91.193.240.157:9877 | ✓ 1060ms | 否 | ✓ 975ms | 否 | ✓ 1711ms | http |
| 144.31.27.49:1080 | ✓ 1345ms | 否 | ✓ 1433ms | 否 | ✓ 1810ms | http |
| 93.185.156.89:3128 | ✓ 827ms | 否 | ✓ 1361ms | 否 | ✓ 1709ms | http |
| 8.219.97.248:80 | ✓ 1995ms | 否 | 否 | ✓ 1142ms | ✓ 1101ms | http |
| 223.84.151.86:30005 | ✓ 1920ms | ✓ 1678ms | ✓ 1693ms | 否 | ✓ 1792ms | http |
| 218.153.163.186:8800 | ✓ 882ms | ✓ 996ms | ✓ 844ms | ✓ 967ms | ✓ 809ms | http |
| 159.223.225.118:8888 | ✓ 902ms | ✓ 1966ms | ✓ 1930ms | 否 | 否 | http |
| 144.31.140.92:1080 | ✓ 1225ms | 否 | ✓ 1618ms | 否 | ✓ 1956ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 763ms | ✓ 1095ms | ✓ 863ms | http |
| 2.27.18.184:1080 | ✓ 1626ms | ✓ 1688ms | 否 | 否 | ✓ 1683ms | http |
| 217.217.249.160:8080 | ✓ 1636ms | 否 | ✓ 1241ms | 否 | ✓ 1216ms | http |
| 205.164.46.6:3157 | 否 | 否 | ✓ 1979ms | ✓ 1320ms | ✓ 1270ms | http |
| 121.230.9.60:1080 | ✓ 1748ms | ✓ 1328ms | ✓ 1023ms | ✓ 1360ms | ✓ 1063ms | http |
| 121.130.199.80:3128 | ✓ 1965ms | 否 | ✓ 1971ms | ✓ 1227ms | ✓ 1280ms | http |
| 101.32.243.189:80 | 否 | ✓ 1333ms | 否 | ✓ 1823ms | ✓ 1460ms | http |
| 103.159.96.195:8082 | ✓ 1882ms | 否 | ✓ 1210ms | 否 | ✓ 1392ms | http |
| 61.52.131.172:8443 | ✓ 919ms | ✓ 1169ms | ✓ 943ms | ✓ 1244ms | ✓ 958ms | http |
| 45.140.147.155:1081 | ✓ 1023ms | ✓ 1715ms | ✓ 1423ms | ✓ 1691ms | ✓ 1590ms | http |
| 62.113.119.14:8080 | ✓ 908ms | ✓ 1878ms | ✓ 1183ms | ✓ 1712ms | ✓ 1258ms | http |
| 103.39.51.207:8080 | ✓ 1342ms | 否 | 否 | ✓ 1303ms | ✓ 1312ms | http |
| 123.57.2.231:2020 | ✓ 1958ms | ✓ 1151ms | ✓ 932ms | ✓ 1157ms | ✓ 960ms | http |

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
