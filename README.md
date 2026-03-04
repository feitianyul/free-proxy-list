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

最后更新：2026-03-04 22:40:20 UTC（2026-03-05 06:40:20 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 104ms | ✓ 1707ms | ✓ 741ms | ✓ 1206ms | ✓ 799ms | http |
| 35.225.22.61:80 | ✓ 872ms | ✓ 1767ms | ✓ 1256ms | ✓ 1400ms | ✓ 924ms | http |
| 138.124.53.25:7443 | ✓ 385ms | ✓ 1778ms | ✓ 1757ms | 否 | ✓ 1263ms | http |
| 14.56.107.244:3128 | ✓ 1095ms | ✓ 1421ms | ✓ 1121ms | 否 | ✓ 962ms | http |
| 125.128.12.14:3128 | ✓ 967ms | ✓ 1526ms | ✓ 1150ms | 否 | ✓ 1042ms | http |
| 14.56.177.44:3128 | ✓ 971ms | 否 | ✓ 1165ms | ✓ 1332ms | ✓ 1184ms | http |
| 125.128.12.144:3128 | ✓ 1341ms | ✓ 1882ms | 否 | ✓ 1385ms | ✓ 969ms | http |
| 61.72.221.234:3128 | ✓ 1380ms | ✓ 1877ms | ✓ 1049ms | 否 | ✓ 1496ms | http |
| 61.72.110.54:3128 | 否 | ✓ 1553ms | ✓ 811ms | 否 | ✓ 974ms | http |
| 121.128.121.54:3128 | ✓ 751ms | ✓ 1027ms | ✓ 734ms | ✓ 1222ms | ✓ 928ms | http |
| 162.240.154.26:3128 | ✓ 1977ms | ✓ 1596ms | ✓ 1209ms | ✓ 1662ms | ✓ 1285ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1432ms | 否 | ✓ 1418ms | ✓ 1146ms | http |
| 61.72.110.94:3128 | 否 | ✓ 1652ms | 否 | ✓ 1549ms | ✓ 1768ms | http |
| 90.84.188.97:8000 | ✓ 1017ms | ✓ 1731ms | 否 | 否 | ✓ 1859ms | http |
| 222.228.171.92:8080 | ✓ 1253ms | 否 | ✓ 1949ms | ✓ 1248ms | ✓ 1399ms | http |
| 211.171.114.154:3128 | ✓ 1336ms | ✓ 1598ms | ✓ 1826ms | 否 | ✓ 1471ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 836ms | ✓ 1963ms | ✓ 1293ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1698ms | 否 | ✓ 1547ms | ✓ 778ms | http |
| 91.233.223.147:3128 | ✓ 908ms | 否 | ✓ 1973ms | ✓ 1957ms | 否 | http |
| 172.212.68.37:3128 | ✓ 187ms | ✓ 1216ms | ✓ 1280ms | ✓ 1278ms | ✓ 994ms | http |
| 192.166.82.55:1080 | ✓ 1022ms | ✓ 1403ms | ✓ 1480ms | 否 | 否 | http |
| 160.238.65.6:3128 | 否 | 否 | ✓ 529ms | ✓ 1348ms | ✓ 975ms | http |
| 160.238.65.2:3128 | 否 | 否 | ✓ 516ms | ✓ 1291ms | ✓ 953ms | http |
| 91.193.240.157:9877 | ✓ 934ms | ✓ 1906ms | ✓ 851ms | 否 | ✓ 1395ms | http |
| 52.47.164.226:11725 | ✓ 1229ms | 否 | ✓ 1513ms | ✓ 1926ms | 否 | http |
| 199.38.85.122:40004 | 否 | ✓ 1921ms | ✓ 1265ms | ✓ 1676ms | ✓ 1345ms | http |
| 188.132.141.249:443 | ✓ 1278ms | 否 | ✓ 1851ms | 否 | ✓ 1784ms | http |
| 168.235.110.63:3128 | ✓ 571ms | ✓ 1186ms | ✓ 1769ms | 否 | ✓ 767ms | http |
| 101.43.255.96:80 | ✓ 1187ms | ✓ 1459ms | ✓ 1105ms | ✓ 1535ms | ✓ 1206ms | http |
| 210.223.44.230:3128 | ✓ 901ms | ✓ 1003ms | ✓ 791ms | ✓ 1126ms | ✓ 985ms | http |
| 121.230.9.183:1080 | ✓ 1355ms | ✓ 1552ms | ✓ 1168ms | ✓ 1632ms | ✓ 1212ms | http |
| 61.72.221.94:3128 | 否 | ✓ 1869ms | 否 | ✓ 1243ms | ✓ 1248ms | http |
| 120.92.212.16:8890 | ✓ 1114ms | ✓ 1653ms | ✓ 1089ms | 否 | 否 | http |
| 160.238.65.4:3128 | ✓ 875ms | ✓ 1990ms | ✓ 645ms | ✓ 1270ms | ✓ 945ms | http |
| 160.238.65.7:3128 | ✓ 870ms | ✓ 1263ms | ✓ 1381ms | ✓ 1342ms | ✓ 985ms | http |
| 160.238.65.5:3128 | ✓ 870ms | ✓ 1142ms | ✓ 1502ms | ✓ 1250ms | ✓ 963ms | http |
| 160.238.65.9:3128 | ✓ 876ms | ✓ 1327ms | ✓ 1308ms | ✓ 1253ms | 否 | http |
| 160.238.65.3:3128 | ✓ 870ms | ✓ 1161ms | ✓ 1484ms | ✓ 1266ms | 否 | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1467ms | ✓ 1688ms | ✓ 1353ms | http |
| 160.238.65.8:3128 | ✓ 870ms | ✓ 1700ms | ✓ 945ms | ✓ 1244ms | ✓ 1994ms | http |
| 81.70.169.194:80 | ✓ 1212ms | ✓ 1426ms | ✓ 1187ms | ✓ 1531ms | ✓ 1294ms | http |
| 185.213.20.105:3128 | ✓ 1055ms | ✓ 1351ms | ✓ 1284ms | ✓ 1810ms | 否 | http |
| 173.212.246.157:3128 | ✓ 588ms | 否 | ✓ 1862ms | ✓ 1967ms | ✓ 1397ms | http |
| 46.249.103.192:443 | ✓ 505ms | 否 | ✓ 1319ms | ✓ 1880ms | 否 | http |
| 154.90.48.209:9090 | 否 | 否 | ✓ 1465ms | ✓ 1907ms | ✓ 1724ms | http |
| 213.131.85.27:1981 | 否 | 否 | ✓ 909ms | ✓ 1565ms | ✓ 1641ms | http |
| 154.12.231.32:80 | ✓ 474ms | ✓ 1155ms | ✓ 1457ms | ✓ 1364ms | ✓ 936ms | http |
| 20.120.225.109:3128 | ✓ 704ms | ✓ 1476ms | ✓ 1217ms | ✓ 1590ms | ✓ 1052ms | http |
| 199.38.85.122:40001 | ✓ 1035ms | 否 | ✓ 1211ms | 否 | ✓ 1441ms | http |
| 121.230.9.248:1080 | ✓ 1355ms | ✓ 1799ms | ✓ 1611ms | ✓ 1875ms | ✓ 1278ms | http |
| 35.234.17.221:8080 | ✓ 1497ms | ✓ 1878ms | ✓ 1477ms | ✓ 1258ms | ✓ 1030ms | http |
| 148.153.56.51:80 | ✓ 561ms | ✓ 1540ms | ✓ 1294ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1042ms | ✓ 1483ms | ✓ 911ms | ✓ 1509ms | ✓ 1247ms | http |
| 176.100.39.18:3128 | ✓ 996ms | ✓ 1636ms | ✓ 969ms | ✓ 1547ms | ✓ 1259ms | http |
| 45.136.198.40:3128 | 否 | ✓ 1807ms | ✓ 1830ms | ✓ 1919ms | ✓ 1553ms | http |
| 193.181.35.26:8118 | ✓ 875ms | 否 | 否 | ✓ 1916ms | ✓ 1445ms | http |
| 45.140.147.155:1081 | 否 | 否 | ✓ 1103ms | ✓ 1425ms | ✓ 1100ms | http |
| 47.101.149.27:9010 | ✓ 1566ms | ✓ 1564ms | 否 | 否 | ✓ 1598ms | http |
| 103.215.36.88:16988 | ✓ 1278ms | ✓ 1498ms | ✓ 1152ms | ✓ 1556ms | ✓ 1261ms | http |
| 150.249.255.91:3128 | ✓ 838ms | ✓ 1676ms | 否 | ✓ 1948ms | ✓ 1179ms | http |
| 74.48.78.224:2080 | ✓ 463ms | ✓ 901ms | ✓ 458ms | ✓ 1064ms | ✓ 803ms | http |
| 120.55.163.237:10086 | ✓ 1020ms | ✓ 1301ms | ✓ 1093ms | ✓ 1325ms | ✓ 1094ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1946ms | 否 | ✓ 1126ms | ✓ 951ms | http |
| 185.243.218.43:49153 | ✓ 548ms | ✓ 1825ms | ✓ 1382ms | ✓ 1810ms | ✓ 1669ms | http |
| 103.215.36.88:18561 | ✓ 1355ms | 否 | ✓ 1248ms | ✓ 1517ms | ✓ 1223ms | http |
| 1.225.116.115:1080 | ✓ 1326ms | ✓ 1521ms | 否 | 否 | ✓ 1132ms | http |
| 103.39.51.190:8080 | ✓ 1981ms | 否 | 否 | ✓ 1609ms | ✓ 1587ms | http |
| 165.227.5.10:8888 | ✓ 830ms | ✓ 1822ms | ✓ 565ms | ✓ 1550ms | ✓ 978ms | http |
| 5.75.196.26:40000 | 否 | 否 | ✓ 1624ms | ✓ 1995ms | ✓ 1943ms | http |
| 88.80.150.82:8080 | ✓ 1410ms | ✓ 1586ms | 否 | 否 | ✓ 1667ms | https |
| 91.217.179.174:8080 | 否 | ✓ 1835ms | ✓ 1741ms | ✓ 1885ms | 否 | http |
| 121.230.8.153:1080 | ✓ 1213ms | ✓ 1662ms | ✓ 1321ms | 否 | 否 | http |
| 8.137.112.117:3128 | ✓ 1091ms | ✓ 1683ms | ✓ 1132ms | ✓ 1583ms | ✓ 1206ms | http |
| 163.44.126.97:3128 | ✓ 1183ms | 否 | ✓ 1608ms | ✓ 1847ms | ✓ 1264ms | http |
| 157.125.220.80:8080 | ✓ 1436ms | 否 | ✓ 1184ms | ✓ 1900ms | 否 | http |
| 125.64.244.100:8889 | ✓ 1864ms | 否 | ✓ 1865ms | 否 | ✓ 1969ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1784ms | ✓ 1819ms | ✓ 1719ms | http |
| 103.215.36.88:16345 | ✓ 1100ms | ✓ 1372ms | ✓ 1140ms | ✓ 1397ms | ✓ 1126ms | http |

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
