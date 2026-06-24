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

最后更新：2026-06-24 14:12:04 UTC（2026-06-24 22:12:04 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.43.46.91:80 | ✓ 767ms | 否 | ✓ 1053ms | ✓ 1144ms | ✓ 907ms | http |
| 47.84.204.82:443 | ✓ 853ms | 否 | ✓ 890ms | ✓ 1219ms | ✓ 967ms | http |
| 3.211.120.181:443 | ✓ 218ms | ✓ 1554ms | ✓ 863ms | 否 | 否 | http |
| 45.88.174.195:8080 | ✓ 1012ms | 否 | ✓ 1632ms | ✓ 1646ms | ✓ 1019ms | http |
| 1.231.81.166:3128 | ✓ 1658ms | 否 | ✓ 866ms | ✓ 1172ms | ✓ 1161ms | http |
| 103.139.103.140:1080 | ✓ 1032ms | 否 | ✓ 998ms | ✓ 1290ms | ✓ 1078ms | http |
| 176.111.37.216:39811 | ✓ 1005ms | 否 | ✓ 1156ms | 否 | ✓ 1427ms | http |
| 185.200.188.234:10001 | ✓ 1032ms | 否 | ✓ 932ms | ✓ 1927ms | ✓ 1521ms | http |
| 114.94.148.37:18080 | ✓ 1907ms | 否 | ✓ 1868ms | ✓ 1457ms | ✓ 1299ms | http |
| 34.101.184.164:3128 | ✓ 1761ms | 否 | ✓ 1005ms | ✓ 1298ms | ✓ 1039ms | http |
| 152.42.243.23:10808 | ✓ 921ms | 否 | 否 | ✓ 1189ms | ✓ 931ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1223ms | ✓ 1450ms | ✓ 1193ms | http |
| 91.107.182.124:82 | ✓ 1166ms | 否 | ✓ 1560ms | ✓ 1928ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1204ms | 否 | ✓ 1506ms | ✓ 1494ms | ✓ 1264ms | http |
| 95.3.69.222:8080 | ✓ 1093ms | ✓ 1918ms | ✓ 1899ms | ✓ 1701ms | ✓ 1467ms | http |
| 116.101.9.20:2055 | ✓ 1666ms | 否 | ✓ 1760ms | 否 | ✓ 1620ms | http |
| 45.153.4.154:3128 | ✓ 609ms | 否 | ✓ 662ms | ✓ 1386ms | ✓ 1133ms | http |
| 174.137.134.182:2999 | ✓ 1128ms | 否 | 否 | ✓ 1006ms | ✓ 855ms | http |
| 8.215.25.3:2080 | ✓ 945ms | 否 | ✓ 897ms | ✓ 1300ms | ✓ 1001ms | http |
| 117.236.124.166:3128 | ✓ 1529ms | 否 | ✓ 1046ms | ✓ 1579ms | 否 | http |
| 202.28.194.139:31280 | ✓ 1824ms | 否 | 否 | ✓ 1929ms | ✓ 1865ms | http |
| 185.191.239.248:3128 | ✓ 793ms | 否 | ✓ 1704ms | ✓ 1661ms | ✓ 1640ms | http |
| 116.101.75.173:2054 | ✓ 1644ms | 否 | ✓ 1713ms | ✓ 1711ms | 否 | http |
| 34.43.46.91:443 | ✓ 399ms | ✓ 1049ms | ✓ 812ms | ✓ 836ms | ✓ 632ms | http |
| 82.97.247.37:80 | ✓ 843ms | ✓ 1463ms | ✓ 561ms | ✓ 1392ms | ✓ 1058ms | http |
| 85.234.100.149:8080 | ✓ 452ms | 否 | ✓ 828ms | 否 | ✓ 1187ms | http |
| 116.101.9.20:2047 | ✓ 1689ms | 否 | ✓ 1734ms | ✓ 1977ms | ✓ 1653ms | http |
| 104.194.146.9:80 | ✓ 614ms | ✓ 1776ms | ✓ 1484ms | ✓ 1628ms | ✓ 1122ms | http |
| 104.194.148.204:80 | ✓ 721ms | 否 | ✓ 1153ms | ✓ 1893ms | ✓ 1190ms | http |
| 138.124.114.42:7443 | ✓ 1819ms | 否 | 否 | ✓ 1670ms | ✓ 1963ms | http |
| 116.101.9.20:2040 | ✓ 1645ms | 否 | ✓ 1654ms | 否 | ✓ 1515ms | http |
| 144.202.14.153:50000 | ✓ 1278ms | 否 | ✓ 409ms | ✓ 1221ms | ✓ 1344ms | http |
| 3.137.86.220:443 | ✓ 1203ms | ✓ 1951ms | 否 | 否 | ✓ 1716ms | http |
| 47.84.204.82:80 | 否 | 否 | ✓ 1659ms | ✓ 1239ms | ✓ 965ms | http |
| 34.87.80.221:30000 | 否 | 否 | ✓ 847ms | ✓ 1586ms | ✓ 1108ms | http |
| 47.84.204.82:1080 | ✓ 1476ms | 否 | 否 | ✓ 1246ms | ✓ 997ms | http |
| 3.137.86.220:1080 | ✓ 282ms | 否 | ✓ 1915ms | 否 | ✓ 702ms | http |
| 104.154.186.48:80 | ✓ 386ms | 否 | ✓ 840ms | ✓ 1114ms | ✓ 978ms | http |
| 129.226.92.241:80 | ✓ 846ms | 否 | ✓ 884ms | ✓ 1195ms | ✓ 949ms | http |
| 36.147.78.166:80 | ✓ 1948ms | 否 | 否 | ✓ 1897ms | ✓ 1883ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1639ms | ✓ 1556ms | ✓ 1927ms | http |
| 116.171.106.78:3443 | 否 | 否 | ✓ 1691ms | ✓ 1907ms | ✓ 1673ms | http |
| 64.188.125.187:3128 | ✓ 667ms | 否 | ✓ 804ms | ✓ 1866ms | ✓ 1850ms | http |
| 116.171.106.15:3443 | 否 | ✓ 1705ms | ✓ 1845ms | ✓ 1919ms | ✓ 1710ms | http |
| 116.101.9.20:2067 | ✓ 1740ms | 否 | ✓ 1626ms | ✓ 1723ms | ✓ 1540ms | http |
| 195.25.20.155:3128 | ✓ 1424ms | 否 | ✓ 1040ms | 否 | ✓ 1614ms | http |
| 125.129.15.95:3128 | ✓ 905ms | 否 | ✓ 848ms | ✓ 1132ms | ✓ 1762ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1538ms | ✓ 1799ms | ✓ 1720ms | http |
| 199.247.29.193:50000 | ✓ 599ms | ✓ 1705ms | ✓ 1061ms | ✓ 1736ms | ✓ 1547ms | http |
| 178.156.224.42:3128 | ✓ 1554ms | 否 | 否 | ✓ 1891ms | ✓ 1382ms | http |
| 151.243.153.157:8118 | ✓ 1021ms | ✓ 852ms | 否 | ✓ 1413ms | 否 | http |
| 91.107.182.124:84 | ✓ 1117ms | ✓ 1949ms | ✓ 1358ms | 否 | ✓ 1737ms | http |
| 111.48.191.1:7890 | ✓ 1818ms | ✓ 1099ms | ✓ 836ms | 否 | ✓ 886ms | http |
| 205.164.192.115:999 | ✓ 1196ms | ✓ 1892ms | 否 | 否 | ✓ 1774ms | http |
| 54.153.56.243:80 | ✓ 433ms | 否 | ✓ 774ms | 否 | ✓ 1718ms | http |
| 194.87.130.170:8443 | ✓ 660ms | 否 | 否 | ✓ 1412ms | ✓ 1525ms | https |
| 47.79.119.13:8080 | ✓ 968ms | 否 | ✓ 895ms | ✓ 1204ms | ✓ 962ms | http |
| 34.84.162.206:38080 | ✓ 1411ms | 否 | ✓ 1598ms | ✓ 1693ms | 否 | http |
| 103.157.117.226:81 | 否 | 否 | ✓ 1390ms | ✓ 1753ms | ✓ 1482ms | http |
| 8.154.21.175:3128 | 否 | ✓ 1691ms | ✓ 976ms | ✓ 1287ms | ✓ 1038ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1315ms | ✓ 1335ms | ✓ 1133ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1162ms | ✓ 1409ms | ✓ 1413ms | http |
| 85.234.100.149:1080 | ✓ 1533ms | 否 | ✓ 434ms | ✓ 1696ms | ✓ 1157ms | http |
| 43.110.40.117:8888 | ✓ 1366ms | ✓ 1087ms | ✓ 1056ms | ✓ 867ms | ✓ 776ms | http |
| 91.186.213.124:1081 | ✓ 1361ms | 否 | ✓ 1216ms | 否 | ✓ 1677ms | http |
| 46.8.112.212:3128 | ✓ 1781ms | 否 | ✓ 642ms | ✓ 1546ms | ✓ 1475ms | http |
| 190.2.145.103:3128 | 否 | ✓ 1870ms | 否 | ✓ 1815ms | ✓ 1484ms | http |
| 158.101.18.102:8888 | 否 | 否 | ✓ 1410ms | ✓ 1206ms | ✓ 935ms | http |
| 43.167.192.85:8080 | ✓ 1159ms | 否 | 否 | ✓ 931ms | ✓ 750ms | http |
| 103.77.243.229:1 | 否 | 否 | ✓ 1951ms | ✓ 1793ms | ✓ 1328ms | http |
| 66.175.236.184:1080 | 否 | 否 | ✓ 1648ms | ✓ 1413ms | ✓ 1266ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 596ms | ✓ 1045ms | ✓ 839ms | http |
| 149.104.4.88:10809 | ✓ 1927ms | 否 | ✓ 1858ms | ✓ 1186ms | ✓ 1997ms | http |
| 14.143.222.113:57748 | 否 | 否 | ✓ 1067ms | ✓ 1499ms | ✓ 1197ms | http |
| 38.41.0.116:999 | 否 | 否 | ✓ 1836ms | ✓ 1447ms | ✓ 1197ms | http |
| 176.12.65.24:443 | ✓ 1893ms | 否 | ✓ 1642ms | ✓ 1778ms | 否 | http |
| 157.230.220.25:4857 | 否 | 否 | ✓ 1669ms | ✓ 960ms | ✓ 901ms | http |
| 85.209.135.240:1081 | ✓ 1854ms | 否 | 否 | ✓ 1960ms | ✓ 1170ms | http |
| 144.178.199.118:8443 | ✓ 1051ms | 否 | ✓ 1676ms | 否 | ✓ 1492ms | http |
| 8.219.97.248:80 | ✓ 1089ms | 否 | ✓ 1296ms | 否 | ✓ 1403ms | http |
| 103.72.68.111:1 | ✓ 1915ms | 否 | 否 | ✓ 1670ms | ✓ 1161ms | http |
| 61.52.131.172:8443 | ✓ 1017ms | ✓ 1294ms | ✓ 1078ms | ✓ 1347ms | ✓ 1041ms | http |
| 18.170.25.193:2977 | ✓ 1315ms | 否 | ✓ 1537ms | ✓ 1831ms | 否 | http |
| 159.195.49.27:8888 | ✓ 1000ms | 否 | ✓ 459ms | ✓ 1904ms | ✓ 1543ms | http |

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
