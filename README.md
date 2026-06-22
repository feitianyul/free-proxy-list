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

最后更新：2026-06-22 11:31:04 UTC（2026-06-22 19:31:04 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.43.46.91:443 | ✓ 1043ms | ✓ 1383ms | ✓ 1485ms | ✓ 1360ms | ✓ 813ms | http |
| 3.211.120.181:443 | ✓ 546ms | ✓ 1662ms | ✓ 474ms | ✓ 1977ms | ✓ 1605ms | http |
| 66.175.236.184:1080 | ✓ 876ms | ✓ 1670ms | ✓ 1521ms | 否 | ✓ 904ms | http |
| 176.111.37.216:39811 | ✓ 805ms | ✓ 1763ms | ✓ 928ms | 否 | ✓ 1887ms | http |
| 37.49.224.15:3128 | ✓ 1009ms | 否 | 否 | ✓ 1973ms | ✓ 1514ms | http |
| 113.160.132.26:8080 | ✓ 1632ms | ✓ 1505ms | ✓ 1080ms | ✓ 1523ms | ✓ 1066ms | http |
| 47.79.119.13:8080 | ✓ 1620ms | 否 | ✓ 749ms | 否 | ✓ 853ms | http |
| 54.255.119.3:9090 | ✓ 1616ms | ✓ 1735ms | ✓ 814ms | ✓ 1333ms | ✓ 1161ms | http |
| 34.43.46.91:80 | ✓ 693ms | ✓ 1700ms | ✓ 903ms | ✓ 1711ms | ✓ 1095ms | http |
| 47.84.204.82:80 | ✓ 782ms | 否 | 否 | ✓ 1775ms | ✓ 865ms | http |
| 185.200.188.234:10001 | ✓ 1163ms | 否 | ✓ 1712ms | 否 | ✓ 1684ms | http |
| 203.90.233.11:8080 | ✓ 1488ms | 否 | ✓ 1757ms | 否 | ✓ 1717ms | http |
| 176.111.37.5:39811 | ✓ 1850ms | ✓ 1849ms | ✓ 1179ms | 否 | ✓ 1757ms | http |
| 47.84.204.82:443 | ✓ 1445ms | ✓ 1761ms | ✓ 1335ms | 否 | ✓ 965ms | http |
| 64.188.125.187:3128 | ✓ 1293ms | 否 | ✓ 855ms | ✓ 1713ms | ✓ 1691ms | http |
| 117.236.124.166:3128 | ✓ 1254ms | 否 | ✓ 1305ms | ✓ 1612ms | 否 | http |
| 190.2.145.103:3128 | ✓ 1857ms | ✓ 1749ms | 否 | ✓ 1777ms | 否 | http |
| 95.3.69.222:8080 | ✓ 1263ms | 否 | ✓ 1158ms | ✓ 1823ms | ✓ 1661ms | http |
| 1.231.81.166:3128 | ✓ 1444ms | 否 | ✓ 1680ms | ✓ 1275ms | ✓ 923ms | http |
| 114.94.148.37:18080 | ✓ 1170ms | 否 | 否 | ✓ 1728ms | ✓ 1837ms | http |
| 79.137.205.130:7443 | ✓ 933ms | 否 | ✓ 1022ms | ✓ 1609ms | ✓ 1476ms | http |
| 144.178.199.118:8443 | ✓ 543ms | ✓ 1698ms | ✓ 852ms | ✓ 1568ms | ✓ 1099ms | http |
| 47.84.204.82:1080 | ✓ 743ms | ✓ 1698ms | 否 | 否 | ✓ 1792ms | http |
| 194.87.187.20:8118 | ✓ 1289ms | 否 | ✓ 936ms | ✓ 1836ms | ✓ 1388ms | http |
| 3.137.86.220:443 | 否 | 否 | ✓ 286ms | ✓ 1774ms | ✓ 873ms | http |
| 212.58.132.5:8888 | ✓ 1543ms | 否 | ✓ 1856ms | ✓ 1589ms | ✓ 1257ms | http |
| 167.103.6.46:11197 | ✓ 1870ms | 否 | ✓ 1231ms | ✓ 1389ms | 否 | http |
| 144.31.73.173:3128 | ✓ 1070ms | 否 | ✓ 917ms | 否 | ✓ 1969ms | http |
| 77.221.156.241:7443 | ✓ 1080ms | 否 | ✓ 1566ms | 否 | ✓ 1913ms | http |
| 205.164.192.115:999 | ✓ 901ms | ✓ 1623ms | 否 | ✓ 1146ms | ✓ 1013ms | http |
| 103.1.51.177:8181 | ✓ 1841ms | 否 | ✓ 1488ms | ✓ 1528ms | 否 | http |
| 199.247.29.193:50000 | ✓ 913ms | ✓ 1919ms | ✓ 1122ms | ✓ 1907ms | ✓ 1462ms | http |
| 34.101.184.164:3128 | ✓ 1645ms | 否 | ✓ 1567ms | ✓ 1367ms | ✓ 1353ms | http |
| 116.101.75.173:2111 | 否 | 否 | ✓ 1658ms | ✓ 1797ms | ✓ 1404ms | http |
| 41.216.182.232:10808 | ✓ 829ms | ✓ 1904ms | ✓ 1389ms | 否 | 否 | http |
| 84.47.150.125:1080 | ✓ 1674ms | 否 | ✓ 1428ms | 否 | ✓ 1723ms | http |
| 82.97.247.37:80 | ✓ 1587ms | 否 | ✓ 783ms | ✓ 1546ms | 否 | http |
| 180.125.216.109:8118 | ✓ 943ms | 否 | ✓ 1371ms | 否 | ✓ 1648ms | http |
| 138.124.114.42:7443 | ✓ 728ms | 否 | ✓ 1685ms | ✓ 1590ms | ✓ 1239ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1681ms | ✓ 1271ms | ✓ 1421ms | http |
| 51.91.208.57:3128 | ✓ 1126ms | ✓ 1640ms | ✓ 1715ms | 否 | ✓ 1970ms | http |
| 103.28.37.131:3128 | ✓ 1819ms | 否 | ✓ 1854ms | ✓ 1662ms | 否 | http |
| 129.226.92.241:80 | ✓ 1461ms | 否 | ✓ 738ms | ✓ 1685ms | 否 | http |
| 14.143.222.113:57748 | ✓ 1229ms | 否 | ✓ 1414ms | 否 | ✓ 1356ms | http |
| 45.129.141.143:3128 | ✓ 1545ms | 否 | ✓ 1759ms | 否 | ✓ 1846ms | http |
| 121.230.8.136:1080 | ✓ 1720ms | ✓ 1602ms | 否 | 否 | ✓ 1511ms | http |
| 8.154.21.175:3128 | 否 | ✓ 1300ms | ✓ 1676ms | 否 | ✓ 1431ms | http |
| 176.12.69.2:443 | ✓ 1975ms | 否 | ✓ 1482ms | 否 | ✓ 1389ms | http |
| 45.88.174.195:8080 | ✓ 550ms | 否 | 否 | ✓ 1739ms | ✓ 1053ms | http |
| 103.167.61.162:3128 | ✓ 668ms | 否 | ✓ 968ms | ✓ 1105ms | ✓ 853ms | http |
| 34.71.229.255:3128 | ✓ 1032ms | 否 | ✓ 643ms | ✓ 1514ms | ✓ 1017ms | http |
| 104.194.148.204:80 | ✓ 837ms | ✓ 1583ms | ✓ 543ms | ✓ 1356ms | ✓ 1284ms | http |
| 104.194.146.9:80 | ✓ 887ms | ✓ 1842ms | ✓ 555ms | ✓ 1395ms | ✓ 1053ms | http |
| 34.200.8.40:3128 | ✓ 345ms | 否 | ✓ 902ms | ✓ 1970ms | ✓ 1483ms | http |
| 150.136.239.172:3128 | ✓ 814ms | 否 | ✓ 1875ms | ✓ 1417ms | ✓ 1151ms | http |
| 167.86.95.198:3128 | ✓ 1044ms | ✓ 1765ms | ✓ 974ms | 否 | ✓ 1301ms | http |
| 45.59.122.132:80 | ✓ 672ms | ✓ 1797ms | ✓ 874ms | 否 | ✓ 1094ms | http |
| 160.238.65.6:3128 | ✓ 1263ms | ✓ 1573ms | 否 | 否 | ✓ 1337ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1186ms | 否 | ✓ 1924ms | ✓ 974ms | http |
| 160.238.65.7:3128 | ✓ 1589ms | ✓ 1454ms | 否 | 否 | ✓ 1356ms | http |
| 72.56.238.99:9090 | ✓ 753ms | 否 | ✓ 743ms | 否 | ✓ 1709ms | http |
| 116.101.9.20:2050 | ✓ 1535ms | 否 | ✓ 1488ms | ✓ 1544ms | ✓ 1490ms | http |
| 120.92.212.16:8890 | ✓ 1712ms | 否 | ✓ 1606ms | ✓ 1728ms | ✓ 1338ms | http |
| 116.101.9.20:2055 | ✓ 1525ms | 否 | ✓ 1579ms | ✓ 1559ms | ✓ 1444ms | http |
| 120.92.212.16:7890 | ✓ 1921ms | ✓ 1451ms | ✓ 1319ms | ✓ 1269ms | 否 | http |
| 100.52.49.226:3128 | ✓ 1247ms | 否 | ✓ 1936ms | ✓ 1890ms | 否 | http |
| 116.101.9.20:2040 | 否 | 否 | ✓ 1595ms | ✓ 1569ms | ✓ 1455ms | http |
| 199.127.62.89:3129 | ✓ 1199ms | 否 | ✓ 1762ms | 否 | ✓ 1712ms | http |
| 62.133.62.17:1082 | ✓ 1561ms | 否 | ✓ 1486ms | 否 | ✓ 1510ms | http |
| 202.28.194.139:31280 | ✓ 1962ms | 否 | ✓ 1777ms | 否 | ✓ 1941ms | http |
| 116.101.75.173:2054 | ✓ 1469ms | 否 | ✓ 1343ms | ✓ 1605ms | 否 | http |
| 103.69.84.106:3131 | 否 | 否 | ✓ 1683ms | ✓ 1266ms | ✓ 991ms | http |
| 154.9.237.250:10810 | ✓ 989ms | ✓ 657ms | ✓ 988ms | ✓ 897ms | ✓ 556ms | http |
| 111.48.191.1:7890 | ✓ 1775ms | 否 | ✓ 1507ms | ✓ 989ms | ✓ 775ms | http |
| 61.52.131.172:8443 | ✓ 874ms | ✓ 1628ms | ✓ 934ms | ✓ 1184ms | ✓ 975ms | http |
| 158.255.212.55:9005 | ✓ 1291ms | 否 | ✓ 1878ms | ✓ 1854ms | 否 | http |
| 158.255.212.55:9480 | ✓ 1290ms | 否 | ✓ 1874ms | ✓ 1870ms | 否 | http |
| 158.255.212.55:3256 | ✓ 1288ms | 否 | ✓ 1876ms | ✓ 1861ms | 否 | http |
| 158.255.212.55:7497 | ✓ 1556ms | 否 | ✓ 1609ms | ✓ 1875ms | 否 | http |
| 152.32.132.190:7890 | ✓ 1321ms | ✓ 1093ms | ✓ 861ms | ✓ 1724ms | ✓ 1588ms | http |

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
