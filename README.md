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

最后更新：2026-05-08 10:16:35 UTC（2026-05-08 18:16:35 UTC+8）

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
| 103.147.152.12:1080 | ✓ 1531ms | ✓ 1469ms | ✓ 1012ms | 否 | ✓ 1758ms | http |
| 212.224.88.212:443 | ✓ 1149ms | 否 | ✓ 1039ms | 否 | ✓ 1585ms | http |
| 185.221.237.57:8443 | ✓ 1158ms | 否 | ✓ 1619ms | ✓ 1937ms | ✓ 1725ms | http |
| 65.109.125.111:8443 | ✓ 1248ms | ✓ 1863ms | ✓ 1732ms | 否 | 否 | http |
| 193.160.209.58:1080 | ✓ 1791ms | 否 | ✓ 1402ms | 否 | ✓ 1671ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1980ms | 否 | ✓ 1621ms | ✓ 1341ms | http |
| 181.119.97.24:999 | ✓ 1413ms | 否 | ✓ 1879ms | ✓ 1997ms | ✓ 1762ms | http |
| 107.173.42.121:7890 | 否 | ✓ 932ms | ✓ 557ms | ✓ 1054ms | 否 | http |
| 185.125.100.115:40000 | ✓ 1160ms | 否 | ✓ 1326ms | ✓ 1973ms | ✓ 1917ms | http |
| 45.125.67.37:8443 | ✓ 1214ms | 否 | ✓ 1191ms | 否 | ✓ 1446ms | http |
| 218.108.131.186:17890 | ✓ 966ms | ✓ 1201ms | ✓ 1218ms | 否 | 否 | http |
| 38.194.254.134:999 | ✓ 1851ms | ✓ 1392ms | ✓ 1155ms | 否 | ✓ 1567ms | http |
| 91.242.229.129:8092 | ✓ 969ms | ✓ 1386ms | ✓ 1687ms | 否 | ✓ 1559ms | http |
| 62.133.60.126:24558 | ✓ 944ms | 否 | ✓ 1228ms | 否 | ✓ 1640ms | http |
| 147.45.178.211:14658 | ✓ 1643ms | ✓ 1676ms | ✓ 1943ms | 否 | 否 | http |
| 79.137.205.44:40000 | ✓ 1177ms | 否 | ✓ 1401ms | ✓ 1781ms | 否 | http |
| 77.110.119.136:3128 | ✓ 1577ms | 否 | 否 | ✓ 1449ms | ✓ 1010ms | http |
| 115.231.181.40:8128 | ✓ 1138ms | 否 | ✓ 1378ms | ✓ 1385ms | 否 | http |
| 103.158.242.58:83 | 否 | 否 | ✓ 1731ms | ✓ 1799ms | ✓ 1819ms | http |
| 185.221.237.57:443 | ✓ 674ms | 否 | ✓ 1595ms | 否 | ✓ 1608ms | http |
| 38.211.245.18:999 | ✓ 1078ms | 否 | ✓ 1030ms | 否 | ✓ 1890ms | http |
| 120.92.108.86:7890 | ✓ 1467ms | 否 | ✓ 1914ms | 否 | ✓ 1935ms | http |
| 116.80.49.97:3172 | 否 | 否 | ✓ 1612ms | ✓ 1965ms | ✓ 1801ms | http |
| 194.59.247.34:10808 | ✓ 1040ms | ✓ 1373ms | ✓ 1392ms | 否 | 否 | http |
| 43.156.132.113:3128 | 否 | ✓ 1825ms | ✓ 1711ms | ✓ 1193ms | ✓ 937ms | http |
| 15.161.131.175:59656 | ✓ 1827ms | 否 | 否 | ✓ 1790ms | ✓ 1393ms | http |
| 31.131.248.48:3129 | ✓ 904ms | ✓ 1770ms | 否 | 否 | ✓ 1782ms | http |
| 148.230.4.241:999 | ✓ 1446ms | ✓ 1928ms | ✓ 651ms | ✓ 1618ms | ✓ 1443ms | http |
| 84.47.150.125:1080 | ✓ 1101ms | 否 | 否 | ✓ 1699ms | ✓ 1571ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1902ms | 否 | ✓ 1921ms | ✓ 1753ms | http |
| 137.59.47.73:3128 | ✓ 1525ms | 否 | ✓ 1073ms | ✓ 1723ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1435ms | 否 | 否 | ✓ 1590ms | ✓ 1247ms | http |
| 206.206.126.177:2412 | ✓ 1771ms | 否 | 否 | ✓ 1956ms | ✓ 1623ms | http |
| 103.209.36.58:8080 | ✓ 1578ms | 否 | 否 | ✓ 1908ms | ✓ 1778ms | http |
| 59.46.216.131:30001 | ✓ 1103ms | 否 | 否 | ✓ 1535ms | ✓ 1162ms | http |
| 118.113.244.225:1080 | 否 | ✓ 1760ms | ✓ 1405ms | ✓ 1951ms | ✓ 1576ms | http |
| 8.209.238.110:47701 | ✓ 1421ms | 否 | ✓ 885ms | ✓ 1003ms | ✓ 814ms | http |
| 45.153.231.229:8080 | ✓ 1347ms | 否 | ✓ 875ms | 否 | ✓ 1820ms | http |
| 210.223.44.230:3128 | ✓ 1794ms | ✓ 1397ms | 否 | 否 | ✓ 1796ms | http |
| 95.217.103.18:4567 | ✓ 1034ms | 否 | ✓ 1610ms | ✓ 1987ms | 否 | http |
| 86.104.72.219:1082 | ✓ 182ms | ✓ 1586ms | 否 | ✓ 1357ms | ✓ 1105ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1736ms | ✓ 1454ms | ✓ 1093ms | http |
| 38.188.247.12:999 | ✓ 955ms | ✓ 1634ms | 否 | ✓ 1428ms | ✓ 1367ms | http |
| 86.104.72.220:1081 | ✓ 396ms | ✓ 962ms | ✓ 1209ms | ✓ 1015ms | ✓ 944ms | http |
| 45.116.14.87:8080 | ✓ 1754ms | ✓ 1960ms | ✓ 1042ms | ✓ 1167ms | ✓ 846ms | http |
| 49.65.127.215:3128 | 否 | ✓ 1302ms | ✓ 1250ms | ✓ 1348ms | ✓ 977ms | http |
| 121.230.8.237:1080 | ✓ 1302ms | 否 | ✓ 1098ms | ✓ 1542ms | ✓ 1347ms | http |
| 8.154.21.175:3128 | ✓ 1722ms | 否 | ✓ 1010ms | 否 | ✓ 1021ms | http |
| 110.172.28.217:3128 | ✓ 1652ms | 否 | ✓ 1790ms | ✓ 1340ms | ✓ 1107ms | http |
| 34.101.184.164:3128 | ✓ 1916ms | 否 | ✓ 1296ms | 否 | ✓ 1322ms | http |
| 152.32.132.190:7890 | ✓ 1336ms | 否 | 否 | ✓ 1032ms | ✓ 1313ms | http |
| 86.104.74.110:1082 | 否 | ✓ 1295ms | ✓ 768ms | ✓ 1563ms | ✓ 1220ms | http |
| 86.104.74.110:1081 | 否 | ✓ 1307ms | ✓ 755ms | 否 | ✓ 910ms | http |
| 168.110.52.228:3128 | ✓ 1276ms | 否 | ✓ 685ms | ✓ 990ms | ✓ 748ms | http |
| 101.32.243.189:80 | ✓ 1312ms | ✓ 1591ms | ✓ 1717ms | ✓ 1648ms | ✓ 1560ms | http |
| 103.157.200.126:3128 | ✓ 1162ms | 否 | ✓ 1090ms | ✓ 1721ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1751ms | ✓ 1401ms | ✓ 711ms | ✓ 1503ms | ✓ 1152ms | http |
| 104.161.23.122:5036 | ✓ 1892ms | 否 | ✓ 1809ms | 否 | ✓ 1518ms | http |
| 3.101.133.120:80 | ✓ 1776ms | ✓ 1335ms | ✓ 409ms | ✓ 1074ms | ✓ 987ms | http |
| 178.63.155.151:8888 | ✓ 1132ms | ✓ 1594ms | 否 | 否 | ✓ 1543ms | http |
| 147.45.186.28:3128 | ✓ 663ms | 否 | ✓ 803ms | ✓ 1919ms | 否 | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1491ms | ✓ 1328ms | ✓ 1157ms | http |
| 138.197.68.35:4857 | ✓ 416ms | ✓ 1068ms | 否 | ✓ 1131ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1167ms | 否 | ✓ 1795ms | 否 | ✓ 1844ms | http |
| 45.186.6.104:3128 | ✓ 1691ms | ✓ 1841ms | ✓ 1909ms | 否 | 否 | http |
| 185.121.13.73:3128 | ✓ 980ms | 否 | ✓ 887ms | 否 | ✓ 1538ms | http |
| 77.110.107.80:1080 | 否 | 否 | ✓ 1465ms | ✓ 1589ms | ✓ 1838ms | http |
| 77.110.107.80:8080 | 否 | 否 | ✓ 1606ms | ✓ 1478ms | ✓ 1834ms | http |
| 47.112.25.109:7890 | 否 | 否 | ✓ 1451ms | ✓ 1282ms | ✓ 1267ms | http |
| 91.217.81.131:1080 | ✓ 779ms | ✓ 1905ms | ✓ 1856ms | 否 | 否 | http |
| 196.204.138.244:1976 | ✓ 1864ms | ✓ 1832ms | 否 | 否 | ✓ 1884ms | http |
| 65.108.203.36:18080 | ✓ 728ms | ✓ 1736ms | ✓ 783ms | 否 | 否 | http |
| 137.184.0.30:3128 | ✓ 431ms | ✓ 971ms | ✓ 421ms | ✓ 880ms | ✓ 724ms | http |
| 116.171.106.111:3443 | ✓ 1623ms | ✓ 1648ms | 否 | 否 | ✓ 1405ms | http |
| 2.27.32.81:3128 | ✓ 1287ms | 否 | ✓ 913ms | 否 | ✓ 1922ms | http |
| 5.42.127.197:3128 | ✓ 1073ms | 否 | ✓ 1393ms | 否 | ✓ 1487ms | http |
| 61.52.131.172:8443 | ✓ 1051ms | ✓ 1327ms | ✓ 1094ms | ✓ 1341ms | ✓ 1111ms | http |
| 20.2.83.243:3128 | ✓ 1795ms | ✓ 1682ms | ✓ 998ms | ✓ 1030ms | ✓ 836ms | http |
| 150.107.140.238:3128 | ✓ 1852ms | 否 | ✓ 1015ms | ✓ 1585ms | 否 | http |
| 121.230.8.136:1080 | 否 | 否 | ✓ 1397ms | ✓ 1969ms | ✓ 1496ms | http |

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
