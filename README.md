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

最后更新：2026-06-29 14:56:08 UTC（2026-06-29 22:56:08 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.43.46.91:443 | ✓ 435ms | ✓ 1442ms | 否 | ✓ 1522ms | ✓ 1248ms | http |
| 34.43.46.91:80 | ✓ 745ms | 否 | ✓ 1123ms | ✓ 1466ms | ✓ 1123ms | http |
| 176.111.37.216:39811 | ✓ 445ms | ✓ 1463ms | ✓ 1159ms | ✓ 1506ms | ✓ 1085ms | http |
| 43.133.187.166:3128 | ✓ 1819ms | ✓ 1577ms | ✓ 727ms | ✓ 1321ms | ✓ 1301ms | http |
| 37.49.224.15:3128 | ✓ 766ms | 否 | ✓ 1042ms | ✓ 1813ms | ✓ 1832ms | http |
| 83.171.224.165:8080 | ✓ 997ms | 否 | ✓ 832ms | 否 | ✓ 1599ms | http |
| 185.200.188.234:10001 | ✓ 971ms | 否 | ✓ 662ms | 否 | ✓ 1584ms | http |
| 45.153.4.154:3128 | ✓ 1033ms | ✓ 1583ms | ✓ 1516ms | ✓ 1856ms | ✓ 1449ms | http |
| 176.111.37.5:39811 | ✓ 624ms | 否 | ✓ 1397ms | 否 | ✓ 1479ms | http |
| 43.133.169.167:3128 | 否 | 否 | ✓ 1293ms | ✓ 1067ms | ✓ 932ms | http |
| 104.194.148.188:3128 | ✓ 960ms | 否 | ✓ 702ms | 否 | ✓ 1640ms | http |
| 54.38.138.60:3128 | 否 | 否 | ✓ 1994ms | ✓ 1968ms | ✓ 1853ms | http |
| 104.154.186.48:80 | 否 | 否 | ✓ 1775ms | ✓ 1310ms | ✓ 1884ms | http |
| 43.133.1.198:3128 | ✓ 1790ms | ✓ 1946ms | ✓ 1532ms | ✓ 1167ms | ✓ 1285ms | http |
| 43.153.199.126:8888 | ✓ 940ms | 否 | ✓ 986ms | ✓ 1291ms | 否 | http |
| 174.137.134.182:2999 | ✓ 1458ms | ✓ 1554ms | 否 | 否 | ✓ 1475ms | http |
| 43.133.15.47:3128 | ✓ 1794ms | 否 | ✓ 1471ms | ✓ 1309ms | ✓ 1294ms | http |
| 43.133.30.18:3128 | ✓ 1757ms | ✓ 1555ms | 否 | 否 | ✓ 851ms | http |
| 47.84.204.82:443 | ✓ 1416ms | 否 | ✓ 922ms | ✓ 1985ms | ✓ 1657ms | http |
| 1.231.81.166:3128 | ✓ 1762ms | 否 | 否 | ✓ 1518ms | ✓ 1337ms | http |
| 212.58.132.5:8888 | ✓ 1389ms | 否 | ✓ 1660ms | ✓ 1586ms | ✓ 1335ms | http |
| 43.167.245.99:3128 | ✓ 1771ms | ✓ 1969ms | ✓ 1649ms | 否 | 否 | http |
| 110.49.66.210:8080 | ✓ 1624ms | 否 | ✓ 1995ms | 否 | ✓ 1330ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1956ms | ✓ 1808ms | 否 | ✓ 1185ms | http |
| 139.162.91.253:1080 | ✓ 1818ms | 否 | ✓ 1455ms | 否 | ✓ 966ms | http |
| 120.92.212.16:7890 | ✓ 1234ms | 否 | 否 | ✓ 1395ms | ✓ 1465ms | http |
| 43.161.239.147:8888 | ✓ 1434ms | 否 | ✓ 1467ms | ✓ 1672ms | ✓ 1728ms | http |
| 152.32.132.190:7890 | ✓ 1975ms | 否 | ✓ 1533ms | ✓ 1891ms | 否 | http |
| 86.127.175.42:8080 | ✓ 1118ms | ✓ 1986ms | ✓ 808ms | 否 | ✓ 1760ms | http |
| 3.211.120.181:443 | ✓ 621ms | 否 | ✓ 370ms | ✓ 1633ms | 否 | http |
| 117.236.124.166:3128 | ✓ 1034ms | 否 | ✓ 1037ms | ✓ 1422ms | ✓ 1129ms | http |
| 159.195.49.27:8888 | ✓ 1049ms | 否 | ✓ 457ms | ✓ 1586ms | ✓ 1855ms | http |
| 47.84.204.82:1080 | ✓ 1149ms | 否 | ✓ 1135ms | ✓ 1330ms | ✓ 1057ms | http |
| 65.109.179.84:8443 | ✓ 733ms | 否 | ✓ 1440ms | 否 | ✓ 1667ms | http |
| 185.191.239.248:3128 | ✓ 1317ms | 否 | ✓ 1656ms | 否 | ✓ 1662ms | http |
| 176.12.65.24:443 | ✓ 1395ms | 否 | ✓ 1782ms | 否 | ✓ 1874ms | http |
| 124.156.230.244:3128 | ✓ 1873ms | 否 | ✓ 1720ms | 否 | ✓ 1518ms | http |
| 8.208.85.87:8080 | ✓ 1087ms | 否 | ✓ 616ms | ✓ 1529ms | 否 | http |
| 72.56.238.99:9090 | ✓ 578ms | 否 | ✓ 537ms | 否 | ✓ 1530ms | http |
| 114.94.148.37:18080 | ✓ 1533ms | 否 | ✓ 1564ms | ✓ 1471ms | ✓ 1634ms | http |
| 54.38.139.182:3128 | ✓ 481ms | 否 | ✓ 1193ms | 否 | ✓ 1827ms | http |
| 72.56.238.99:1080 | ✓ 648ms | 否 | 否 | ✓ 1791ms | ✓ 1618ms | http |
| 112.28.149.152:8443 | ✓ 1589ms | 否 | ✓ 1762ms | ✓ 1992ms | 否 | http |
| 185.196.61.251:8081 | ✓ 1496ms | 否 | ✓ 852ms | ✓ 1657ms | 否 | http |
| 157.180.84.115:443 | ✓ 532ms | ✓ 1638ms | ✓ 1623ms | 否 | 否 | http |
| 202.28.194.139:31280 | 否 | 否 | ✓ 1786ms | ✓ 1921ms | ✓ 1914ms | http |
| 3.85.42.63:3128 | ✓ 100ms | 否 | ✓ 1341ms | ✓ 1040ms | ✓ 1279ms | http |
| 34.87.80.221:30000 | ✓ 1939ms | 否 | ✓ 1652ms | ✓ 1483ms | ✓ 1101ms | http |
| 85.234.100.149:8080 | ✓ 437ms | 否 | ✓ 510ms | 否 | ✓ 1115ms | http |
| 8.154.21.175:3128 | ✓ 1390ms | ✓ 1323ms | ✓ 1069ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1505ms | ✓ 1630ms | ✓ 1197ms | 否 | ✓ 1181ms | http |
| 144.202.14.153:50000 | ✓ 379ms | ✓ 1192ms | ✓ 472ms | ✓ 1173ms | ✓ 774ms | http |
| 199.182.135.85:3128 | ✓ 275ms | ✓ 1590ms | ✓ 529ms | ✓ 1615ms | ✓ 1069ms | http |
| 188.225.46.163:10808 | ✓ 1562ms | 否 | ✓ 1400ms | 否 | ✓ 1766ms | http |
| 170.106.136.181:31002 | ✓ 448ms | ✓ 802ms | ✓ 828ms | ✓ 885ms | ✓ 654ms | http |
| 82.146.38.71:443 | ✓ 1141ms | 否 | ✓ 1582ms | ✓ 1995ms | 否 | http |
| 14.143.222.113:57748 | 否 | 否 | ✓ 1390ms | ✓ 1899ms | ✓ 1693ms | http |
| 111.235.151.112:8443 | ✓ 1973ms | 否 | 否 | ✓ 1762ms | ✓ 1456ms | http |
| 117.55.203.158:8899 | ✓ 1362ms | 否 | 否 | ✓ 1113ms | ✓ 1160ms | http |
| 103.21.220.141:3128 | ✓ 847ms | 否 | ✓ 844ms | ✓ 1109ms | ✓ 877ms | http |
| 34.71.229.255:3128 | ✓ 1317ms | ✓ 1500ms | ✓ 994ms | ✓ 1294ms | ✓ 1085ms | http |
| 178.62.184.67:3128 | 否 | 否 | ✓ 410ms | ✓ 1394ms | ✓ 947ms | http |
| 144.178.199.118:8443 | ✓ 973ms | ✓ 1452ms | ✓ 1733ms | 否 | ✓ 1594ms | http |
| 116.62.60.22:3128 | ✓ 1130ms | ✓ 1278ms | ✓ 1224ms | ✓ 1391ms | ✓ 1132ms | http |
| 199.247.29.193:50000 | ✓ 1464ms | 否 | ✓ 774ms | ✓ 1663ms | 否 | http |
| 205.215.247.164:3128 | ✓ 817ms | 否 | ✓ 1827ms | 否 | ✓ 904ms | http |
| 110.172.62.196:8080 | ✓ 1074ms | 否 | ✓ 1616ms | ✓ 1346ms | ✓ 956ms | http |
| 91.188.213.143:1080 | ✓ 1009ms | 否 | ✓ 1994ms | 否 | ✓ 1418ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1454ms | ✓ 1802ms | ✓ 1497ms | ✓ 1521ms | http |
| 18.180.59.181:80 | ✓ 811ms | ✓ 1000ms | ✓ 719ms | ✓ 1103ms | ✓ 1114ms | http |
| 158.101.18.102:8888 | ✓ 808ms | 否 | 否 | ✓ 1405ms | ✓ 1005ms | http |
| 95.190.193.74:3128 | ✓ 1497ms | ✓ 1861ms | 否 | 否 | ✓ 1772ms | http |
| 49.48.110.157:8081 | ✓ 1687ms | 否 | ✓ 1919ms | ✓ 1835ms | ✓ 1821ms | http |
| 64.188.77.26:3128 | ✓ 805ms | 否 | 否 | ✓ 1807ms | ✓ 1959ms | http |
| 206.189.144.164:10808 | ✓ 1028ms | 否 | ✓ 1728ms | ✓ 1689ms | 否 | http |
| 62.133.62.3:1082 | ✓ 1871ms | 否 | ✓ 856ms | ✓ 1841ms | ✓ 1168ms | http |
| 45.186.6.104:3128 | ✓ 1334ms | ✓ 1887ms | ✓ 1984ms | 否 | 否 | http |
| 117.55.203.163:8899 | 否 | 否 | ✓ 531ms | ✓ 1120ms | ✓ 854ms | http |
| 217.154.155.115:8080 | 否 | 否 | ✓ 1713ms | ✓ 1587ms | ✓ 1926ms | http |
| 112.28.149.156:8443 | 否 | 否 | ✓ 1020ms | ✓ 1442ms | ✓ 1159ms | http |
| 8.140.104.98:3128 | ✓ 933ms | 否 | ✓ 992ms | ✓ 1272ms | ✓ 1047ms | http |
| 92.118.112.32:1082 | ✓ 1232ms | 否 | ✓ 116ms | ✓ 1663ms | 否 | http |
| 106.10.55.212:1121 | ✓ 1516ms | ✓ 1725ms | ✓ 841ms | ✓ 1442ms | ✓ 1147ms | http |
| 91.107.182.124:82 | 否 | ✓ 1717ms | ✓ 1093ms | ✓ 1767ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1205ms | 否 | ✓ 1437ms | ✓ 1549ms | ✓ 1220ms | http |
| 61.52.131.172:8443 | ✓ 1597ms | ✓ 1195ms | ✓ 899ms | ✓ 1783ms | ✓ 943ms | http |

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
