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

最后更新：2026-02-28 22:20:12 UTC（2026-03-01 06:20:12 UTC+8）

**代理总数：81**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 900ms | ✓ 989ms | ✓ 766ms | ✓ 1130ms | 否 | http |
| 168.235.110.63:3128 | ✓ 379ms | ✓ 1834ms | ✓ 824ms | ✓ 1420ms | ✓ 886ms | http |
| 3.213.157.4:3128 | ✓ 1243ms | 否 | ✓ 1655ms | ✓ 1505ms | ✓ 736ms | http |
| 148.135.85.87:1080 | ✓ 1511ms | 否 | ✓ 1395ms | ✓ 1306ms | ✓ 963ms | http |
| 62.113.119.14:8080 | ✓ 679ms | 否 | ✓ 1240ms | ✓ 1468ms | ✓ 1120ms | http |
| 36.147.78.166:80 | ✓ 1839ms | 否 | ✓ 1918ms | 否 | ✓ 1737ms | http |
| 121.237.181.137:8888 | ✓ 1057ms | ✓ 1295ms | ✓ 1013ms | ✓ 1246ms | ✓ 950ms | http |
| 59.46.216.131:30001 | ✓ 1018ms | ✓ 1489ms | ✓ 1514ms | ✓ 1580ms | 否 | http |
| 35.225.22.61:80 | ✓ 1060ms | ✓ 1136ms | ✓ 1020ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1168ms | ✓ 1380ms | ✓ 1345ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1180ms | ✓ 1348ms | ✓ 1305ms | 否 | 否 | http |
| 34.159.121.205:3128 | ✓ 436ms | ✓ 1874ms | ✓ 1302ms | ✓ 1954ms | ✓ 1506ms | http |
| 34.79.102.160:3128 | ✓ 622ms | ✓ 1798ms | ✓ 1276ms | ✓ 1745ms | ✓ 1661ms | http |
| 34.32.154.33:3128 | ✓ 528ms | ✓ 1881ms | ✓ 1279ms | ✓ 1974ms | ✓ 1747ms | http |
| 34.7.88.87:3128 | ✓ 528ms | ✓ 1874ms | ✓ 1288ms | 否 | ✓ 1745ms | http |
| 81.70.169.194:80 | ✓ 1094ms | ✓ 1399ms | ✓ 1060ms | ✓ 1422ms | ✓ 1075ms | http |
| 101.43.255.96:80 | ✓ 1117ms | ✓ 1440ms | ✓ 1116ms | ✓ 1378ms | ✓ 1170ms | http |
| 115.231.181.40:8128 | ✓ 1088ms | 否 | ✓ 1099ms | ✓ 1229ms | ✓ 1073ms | http |
| 52.188.28.218:3128 | ✓ 486ms | 否 | ✓ 128ms | 否 | ✓ 1680ms | http |
| 104.238.30.45:59741 | ✓ 1696ms | 否 | ✓ 1807ms | 否 | ✓ 1998ms | http |
| 103.84.95.54:7890 | ✓ 1042ms | 否 | ✓ 784ms | ✓ 972ms | ✓ 767ms | http |
| 34.78.177.18:3128 | ✓ 953ms | ✓ 1496ms | ✓ 692ms | ✓ 1593ms | ✓ 1251ms | http |
| 34.78.200.22:3128 | ✓ 962ms | 否 | ✓ 1190ms | ✓ 1884ms | ✓ 1190ms | http |
| 34.89.174.168:3128 | ✓ 1205ms | 否 | ✓ 1185ms | ✓ 1770ms | ✓ 1325ms | http |
| 34.185.159.217:3128 | ✓ 961ms | 否 | ✓ 1404ms | 否 | ✓ 1571ms | http |
| 91.238.104.171:2023 | ✓ 1253ms | ✓ 1887ms | ✓ 1694ms | 否 | ✓ 1703ms | http |
| 34.158.73.60:3128 | ✓ 1547ms | 否 | ✓ 1530ms | 否 | ✓ 1736ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1052ms | ✓ 1142ms | ✓ 1931ms | http |
| 210.223.44.230:3128 | ✓ 786ms | ✓ 1127ms | ✓ 771ms | ✓ 1048ms | ✓ 867ms | http |
| 194.59.204.87:9080 | ✓ 517ms | 否 | ✓ 497ms | ✓ 1832ms | ✓ 1259ms | http |
| 91.238.104.172:2024 | ✓ 608ms | ✓ 1524ms | ✓ 730ms | ✓ 1653ms | ✓ 1174ms | http |
| 94.177.131.12:3128 | ✓ 1294ms | ✓ 1947ms | ✓ 576ms | ✓ 943ms | ✓ 742ms | http |
| 194.87.43.49:8888 | ✓ 1037ms | ✓ 1985ms | ✓ 1548ms | 否 | 否 | http |
| 152.70.137.18:8888 | 否 | ✓ 1090ms | ✓ 471ms | 否 | ✓ 1002ms | http |
| 162.240.154.26:3128 | ✓ 885ms | 否 | ✓ 1174ms | ✓ 837ms | 否 | http |
| 111.79.111.126:3128 | ✓ 1271ms | 否 | ✓ 1072ms | ✓ 1424ms | ✓ 1064ms | http |
| 116.80.64.158:7777 | ✓ 1611ms | 否 | ✓ 1640ms | 否 | ✓ 1761ms | http |
| 35.234.17.221:8080 | ✓ 1476ms | ✓ 1320ms | 否 | ✓ 1675ms | 否 | http |
| 185.115.74.185:8080 | ✓ 898ms | ✓ 1744ms | ✓ 1635ms | 否 | 否 | http |
| 190.9.48.129:999 | 否 | ✓ 1856ms | ✓ 971ms | ✓ 1117ms | ✓ 1965ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1155ms | ✓ 1317ms | ✓ 1308ms | http |
| 31.59.129.75:8080 | ✓ 1991ms | ✓ 1272ms | ✓ 1261ms | 否 | 否 | http |
| 138.124.53.25:7443 | ✓ 529ms | 否 | ✓ 1737ms | 否 | ✓ 1808ms | http |
| 103.236.64.247:8888 | 否 | ✓ 1697ms | 否 | ✓ 1654ms | ✓ 1046ms | http |
| 45.140.147.82:1081 | ✓ 1926ms | 否 | ✓ 953ms | ✓ 1401ms | ✓ 1022ms | http |
| 104.238.30.63:63744 | ✓ 1715ms | 否 | ✓ 1939ms | 否 | ✓ 1994ms | http |
| 121.230.9.54:1080 | ✓ 1214ms | ✓ 1505ms | ✓ 1267ms | ✓ 1674ms | ✓ 1166ms | http |
| 180.127.149.244:1080 | 否 | 否 | ✓ 1045ms | ✓ 1421ms | ✓ 1005ms | http |
| 220.197.44.36:3128 | ✓ 1726ms | 否 | ✓ 1867ms | ✓ 1685ms | 否 | http |
| 77.83.203.6:443 | ✓ 1245ms | ✓ 1403ms | ✓ 1740ms | 否 | ✓ 1467ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 604ms | ✓ 971ms | ✓ 788ms | http |
| 150.249.255.91:3128 | ✓ 1590ms | 否 | ✓ 1975ms | 否 | ✓ 919ms | http |
| 103.215.36.88:17977 | 否 | ✓ 1451ms | ✓ 1139ms | ✓ 1316ms | ✓ 1055ms | http |
| 36.147.78.166:443 | ✓ 1863ms | ✓ 1776ms | ✓ 1878ms | 否 | ✓ 1744ms | http |
| 1.12.62.237:8080 | ✓ 1879ms | ✓ 1771ms | ✓ 1752ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 943ms | ✓ 1307ms | ✓ 659ms | ✓ 1503ms | ✓ 859ms | http |
| 104.238.30.58:63744 | ✓ 1820ms | 否 | ✓ 1935ms | 否 | ✓ 1968ms | http |
| 103.215.36.88:16316 | ✓ 1135ms | ✓ 1515ms | ✓ 1292ms | ✓ 1545ms | ✓ 1186ms | http |
| 5.75.201.136:1080 | ✓ 446ms | 否 | ✓ 1552ms | ✓ 1678ms | ✓ 1780ms | http |
| 45.140.147.82:1082 | ✓ 405ms | ✓ 1297ms | ✓ 1400ms | 否 | 否 | http |
| 104.238.30.40:59741 | ✓ 1776ms | 否 | ✓ 1937ms | 否 | ✓ 1998ms | http |
| 103.215.36.88:18574 | ✓ 1165ms | ✓ 1669ms | ✓ 1250ms | ✓ 1386ms | 否 | http |
| 121.230.9.148:1080 | 否 | ✓ 1711ms | ✓ 1339ms | ✓ 1617ms | ✓ 1844ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1861ms | 否 | ✓ 1338ms | ✓ 862ms | http |
| 172.212.68.37:3128 | ✓ 355ms | ✓ 1510ms | ✓ 1277ms | ✓ 1854ms | ✓ 1082ms | http |
| 103.39.51.190:8080 | ✓ 1903ms | 否 | 否 | ✓ 1587ms | ✓ 1567ms | http |
| 5.101.0.233:3128 | ✓ 1060ms | 否 | ✓ 1627ms | ✓ 1974ms | ✓ 1656ms | http |
| 103.82.23.118:5234 | ✓ 1715ms | 否 | ✓ 1801ms | 否 | ✓ 1595ms | http |
| 47.110.42.192:9003 | 否 | ✓ 1798ms | ✓ 1567ms | ✓ 1741ms | 否 | http |
| 45.140.147.155:1082 | ✓ 450ms | 否 | ✓ 1299ms | ✓ 1968ms | ✓ 1307ms | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 1547ms | ✓ 1179ms | ✓ 1554ms | http |
| 103.104.99.29:80 | ✓ 1905ms | 否 | ✓ 1893ms | ✓ 1905ms | ✓ 1581ms | http |
| 103.104.99.89:80 | ✓ 1908ms | 否 | ✓ 1894ms | ✓ 1900ms | ✓ 1897ms | http |
| 107.174.133.10:3128 | ✓ 785ms | ✓ 1414ms | ✓ 1193ms | ✓ 999ms | 否 | http |
| 47.106.73.57:8118 | ✓ 1297ms | 否 | ✓ 1196ms | ✓ 1305ms | ✓ 1061ms | http |
| 45.136.198.40:3128 | ✓ 998ms | ✓ 1904ms | ✓ 1815ms | ✓ 1781ms | ✓ 1746ms | http |
| 223.16.170.103:3128 | ✓ 1639ms | 否 | 否 | ✓ 1249ms | ✓ 1768ms | http |
| 45.125.67.37:8443 | ✓ 1137ms | 否 | ✓ 1170ms | ✓ 1302ms | ✓ 1054ms | http |
| 61.72.110.94:3128 | 否 | ✓ 1835ms | ✓ 1472ms | 否 | ✓ 972ms | http |
| 103.3.246.71:3128 | ✓ 1251ms | 否 | ✓ 1240ms | ✓ 1323ms | ✓ 1068ms | http |
| 175.194.173.105:3128 | 否 | ✓ 1376ms | ✓ 1196ms | ✓ 1434ms | 否 | http |

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
