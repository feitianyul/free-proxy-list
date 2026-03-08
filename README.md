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

最后更新：2026-03-08 03:24:34 UTC（2026-03-08 11:24:34 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 202ms | 否 | ✓ 1977ms | ✓ 1295ms | ✓ 1057ms | http |
| 152.42.213.210:8080 | ✓ 1092ms | 否 | ✓ 1724ms | ✓ 1274ms | ✓ 1383ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1640ms | 否 | ✓ 1287ms | ✓ 1069ms | http |
| 114.4.251.26:8080 | ✓ 1958ms | 否 | ✓ 1461ms | ✓ 1636ms | ✓ 1601ms | http |
| 125.128.12.14:3128 | ✓ 830ms | 否 | ✓ 1236ms | 否 | ✓ 1642ms | http |
| 61.72.221.194:3128 | ✓ 777ms | 否 | ✓ 1976ms | 否 | ✓ 1016ms | http |
| 121.128.121.54:3128 | ✓ 847ms | ✓ 1350ms | 否 | ✓ 1617ms | 否 | http |
| 181.78.194.249:999 | ✓ 1424ms | 否 | ✓ 1753ms | 否 | ✓ 1686ms | http |
| 85.9.195.140:1080 | ✓ 959ms | ✓ 1971ms | ✓ 1778ms | 否 | 否 | http |
| 167.172.69.123:8080 | ✓ 1613ms | 否 | ✓ 1886ms | ✓ 1348ms | ✓ 1128ms | http |
| 171.7.58.223:8213 | 否 | 否 | ✓ 1483ms | ✓ 1787ms | ✓ 1443ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1994ms | 否 | ✓ 1207ms | ✓ 991ms | http |
| 61.72.110.54:3128 | ✓ 1612ms | 否 | ✓ 765ms | 否 | ✓ 907ms | http |
| 165.227.5.10:8888 | ✓ 1368ms | ✓ 1680ms | 否 | ✓ 1267ms | 否 | http |
| 45.88.0.116:3128 | ✓ 1543ms | ✓ 1721ms | ✓ 1066ms | 否 | 否 | http |
| 181.78.79.155:999 | ✓ 1575ms | ✓ 1861ms | ✓ 1375ms | 否 | 否 | http |
| 202.155.12.161:443 | ✓ 1957ms | 否 | 否 | ✓ 1472ms | ✓ 1491ms | http |
| 14.56.177.44:3128 | ✓ 1796ms | 否 | ✓ 1394ms | ✓ 1208ms | ✓ 1088ms | http |
| 46.183.25.8:443 | ✓ 646ms | 否 | ✓ 602ms | ✓ 1253ms | ✓ 1040ms | http |
| 103.84.95.54:7890 | 否 | ✓ 1998ms | ✓ 1058ms | ✓ 1116ms | ✓ 855ms | http |
| 159.89.31.62:8080 | 否 | 否 | ✓ 1840ms | ✓ 1625ms | ✓ 1487ms | http |
| 81.70.169.194:80 | ✓ 1225ms | ✓ 1548ms | ✓ 1535ms | 否 | ✓ 1173ms | http |
| 14.225.222.164:7890 | ✓ 1226ms | ✓ 1843ms | 否 | ✓ 1351ms | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1543ms | ✓ 1090ms | ✓ 1487ms | 否 | http |
| 45.88.0.115:3128 | ✓ 1495ms | ✓ 1723ms | ✓ 545ms | 否 | 否 | http |
| 45.88.0.111:3128 | ✓ 1497ms | ✓ 1681ms | ✓ 520ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1430ms | ✓ 968ms | 否 | ✓ 866ms | http |
| 62.113.119.14:8080 | ✓ 1122ms | ✓ 1680ms | ✓ 1293ms | ✓ 1475ms | ✓ 1084ms | http |
| 162.248.165.72:1080 | ✓ 1107ms | 否 | ✓ 1887ms | ✓ 1627ms | 否 | http |
| 159.223.42.219:3128 | ✓ 1641ms | 否 | ✓ 1186ms | ✓ 1258ms | ✓ 1001ms | http |
| 167.172.69.123:80 | ✓ 1878ms | 否 | ✓ 1041ms | ✓ 1302ms | ✓ 1331ms | http |
| 194.59.204.87:9080 | ✓ 389ms | ✓ 1312ms | ✓ 488ms | ✓ 1789ms | 否 | http |
| 168.235.110.63:3128 | ✓ 314ms | ✓ 1764ms | ✓ 612ms | 否 | 否 | http |
| 178.236.245.59:3128 | ✓ 743ms | 否 | ✓ 643ms | ✓ 1633ms | ✓ 1263ms | http |
| 88.80.150.82:8080 | ✓ 1182ms | ✓ 1790ms | 否 | ✓ 1907ms | ✓ 1579ms | https |
| 178.236.245.17:3128 | ✓ 1839ms | 否 | ✓ 1697ms | ✓ 1819ms | ✓ 1313ms | http |
| 121.40.231.103:7890 | ✓ 1119ms | ✓ 1335ms | ✓ 1274ms | ✓ 1321ms | ✓ 1032ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 951ms | ✓ 1484ms | ✓ 1276ms | http |
| 5.252.33.13:2025 | ✓ 1345ms | 否 | ✓ 1158ms | 否 | ✓ 1548ms | http |
| 180.76.115.231:3128 | ✓ 1279ms | ✓ 1543ms | ✓ 1798ms | 否 | 否 | http |
| 34.101.184.164:3128 | ✓ 1255ms | 否 | ✓ 1768ms | ✓ 1796ms | ✓ 1223ms | http |
| 45.88.0.113:3128 | 否 | 否 | ✓ 1099ms | ✓ 1339ms | ✓ 1041ms | http |
| 103.183.10.172:3125 | ✓ 1593ms | 否 | 否 | ✓ 1719ms | ✓ 1700ms | http |
| 14.56.107.244:3128 | ✓ 1570ms | ✓ 1780ms | ✓ 1947ms | ✓ 1161ms | 否 | http |
| 14.225.211.139:7890 | ✓ 1697ms | 否 | ✓ 1260ms | 否 | ✓ 1051ms | http |
| 200.174.198.32:8888 | ✓ 1962ms | 否 | ✓ 649ms | 否 | ✓ 1713ms | http |
| 112.65.132.181:3128 | 否 | ✓ 958ms | ✓ 1789ms | ✓ 981ms | 否 | http |
| 103.183.10.169:3125 | ✓ 1548ms | 否 | ✓ 1556ms | ✓ 1743ms | ✓ 1667ms | http |
| 67.169.98.211:443 | ✓ 1138ms | 否 | ✓ 1230ms | 否 | ✓ 1631ms | http |
| 46.249.103.192:443 | ✓ 1593ms | ✓ 1929ms | ✓ 1638ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1235ms | ✓ 1664ms | ✓ 1406ms | 否 | ✓ 1255ms | http |
| 103.183.10.203:3125 | ✓ 1559ms | 否 | ✓ 1536ms | ✓ 1680ms | ✓ 1637ms | http |
| 160.25.35.34:8080 | ✓ 1823ms | 否 | ✓ 1473ms | ✓ 1635ms | ✓ 1586ms | http |
| 163.44.126.97:3128 | ✓ 1767ms | 否 | ✓ 1989ms | 否 | ✓ 1185ms | http |
| 165.232.183.96:3128 | 否 | 否 | ✓ 1661ms | ✓ 1916ms | ✓ 1102ms | http |
| 150.107.140.238:3128 | ✓ 1766ms | 否 | 否 | ✓ 1421ms | ✓ 1421ms | http |
| 152.42.213.210:80 | ✓ 895ms | 否 | 否 | ✓ 1255ms | ✓ 1069ms | http |
| 112.78.187.186:8080 | ✓ 1285ms | 否 | ✓ 1467ms | ✓ 1638ms | ✓ 1623ms | http |
| 185.243.218.43:49153 | ✓ 529ms | ✓ 1852ms | ✓ 1405ms | ✓ 1901ms | ✓ 1551ms | http |
| 114.231.75.125:1080 | ✓ 1464ms | ✓ 1466ms | ✓ 1470ms | ✓ 1505ms | ✓ 1107ms | http |
| 193.181.35.182:8118 | ✓ 679ms | 否 | ✓ 1230ms | ✓ 1901ms | ✓ 1482ms | http |
| 4.213.222.169:3128 | ✓ 1624ms | ✓ 1630ms | ✓ 1381ms | ✓ 1687ms | ✓ 1429ms | http |
| 89.185.85.138:1080 | ✓ 402ms | 否 | ✓ 1871ms | ✓ 1848ms | 否 | http |
| 61.72.221.94:3128 | 否 | 否 | ✓ 1016ms | ✓ 1864ms | ✓ 992ms | http |
| 213.220.62.62:3128 | ✓ 393ms | 否 | ✓ 393ms | 否 | ✓ 1001ms | http |
| 45.88.0.114:3128 | ✓ 411ms | 否 | ✓ 702ms | 否 | ✓ 1034ms | http |
| 45.88.0.117:3128 | 否 | ✓ 1334ms | ✓ 1450ms | ✓ 1326ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1151ms | 否 | ✓ 1975ms | ✓ 1686ms | ✓ 1433ms | http |
| 138.124.53.25:7443 | ✓ 1199ms | ✓ 1504ms | ✓ 1598ms | ✓ 1672ms | ✓ 1220ms | http |
| 120.92.212.16:7890 | ✓ 1197ms | ✓ 1695ms | ✓ 1224ms | 否 | ✓ 1359ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1622ms | ✓ 1211ms | ✓ 1550ms | http |
| 45.186.6.104:3128 | ✓ 1999ms | ✓ 1950ms | ✓ 1664ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 494ms | 否 | ✓ 904ms | ✓ 1169ms | ✓ 868ms | http |
| 47.77.193.180:1080 | ✓ 568ms | ✓ 1527ms | ✓ 274ms | ✓ 954ms | ✓ 655ms | http |
| 107.175.115.167:3128 | ✓ 944ms | 否 | ✓ 669ms | ✓ 928ms | ✓ 1755ms | http |
| 45.136.198.40:3128 | ✓ 1208ms | ✓ 1705ms | ✓ 1540ms | ✓ 1897ms | ✓ 1594ms | http |
| 39.104.201.40:7890 | ✓ 989ms | ✓ 1195ms | ✓ 1234ms | 否 | ✓ 1629ms | http |
| 222.102.86.137:3040 | ✓ 1415ms | 否 | ✓ 1362ms | 否 | ✓ 1442ms | http |
| 103.215.36.88:17690 | ✓ 1326ms | 否 | ✓ 1539ms | ✓ 1488ms | 否 | http |
| 103.215.36.88:19870 | 否 | 否 | ✓ 1109ms | ✓ 1820ms | ✓ 1181ms | http |
| 188.132.141.249:443 | ✓ 675ms | 否 | ✓ 1923ms | 否 | ✓ 1704ms | http |
| 103.39.51.190:8080 | ✓ 1900ms | 否 | 否 | ✓ 1662ms | ✓ 1715ms | http |
| 103.215.36.88:19475 | ✓ 1348ms | ✓ 1517ms | ✓ 1207ms | 否 | ✓ 1204ms | http |
| 14.225.222.185:7890 | 否 | ✓ 1781ms | 否 | ✓ 1989ms | ✓ 1974ms | http |
| 172.212.68.37:3128 | ✓ 1575ms | ✓ 1210ms | ✓ 504ms | ✓ 1330ms | ✓ 949ms | http |

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
