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

最后更新：2026-03-11 08:44:06 UTC（2026-03-11 16:44:06 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 848ms | 否 | ✓ 359ms | ✓ 646ms | ✓ 492ms | http |
| 45.136.131.47:8443 | ✓ 939ms | 否 | ✓ 596ms | ✓ 680ms | ✓ 504ms | http |
| 35.225.22.61:80 | ✓ 612ms | 否 | ✓ 572ms | ✓ 1237ms | ✓ 1039ms | http |
| 205.209.118.30:3138 | ✓ 349ms | 否 | ✓ 991ms | ✓ 1664ms | ✓ 1474ms | http |
| 158.69.185.37:3129 | 否 | 否 | ✓ 1763ms | ✓ 1966ms | ✓ 962ms | http |
| 202.155.12.161:443 | ✓ 1890ms | 否 | 否 | ✓ 1964ms | ✓ 1276ms | http |
| 14.225.222.213:7890 | 否 | 否 | ✓ 949ms | ✓ 1143ms | ✓ 902ms | http |
| 190.212.131.238:3128 | ✓ 946ms | ✓ 1890ms | 否 | ✓ 1982ms | ✓ 1440ms | http |
| 160.238.65.9:3128 | ✓ 1302ms | 否 | 否 | ✓ 1627ms | ✓ 1750ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1519ms | ✓ 1603ms | ✓ 1674ms | ✓ 1660ms | http |
| 45.136.130.175:8443 | 否 | 否 | ✓ 1465ms | ✓ 1725ms | ✓ 1922ms | http |
| 14.225.222.164:7890 | ✓ 1541ms | ✓ 1893ms | ✓ 889ms | 否 | 否 | http |
| 190.9.109.198:999 | ✓ 928ms | 否 | ✓ 1305ms | 否 | ✓ 1359ms | http |
| 103.82.23.118:5247 | ✓ 1609ms | 否 | ✓ 1173ms | ✓ 1541ms | ✓ 1260ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1039ms | ✓ 1633ms | ✓ 1008ms | http |
| 46.183.25.8:443 | ✓ 1468ms | 否 | ✓ 1731ms | ✓ 1448ms | 否 | http |
| 120.92.212.16:7890 | ✓ 980ms | 否 | ✓ 972ms | ✓ 1222ms | 否 | http |
| 194.213.18.200:443 | ✓ 1662ms | ✓ 1213ms | ✓ 1072ms | 否 | 否 | http |
| 81.70.169.194:80 | 否 | ✓ 1253ms | 否 | ✓ 1278ms | ✓ 1014ms | http |
| 101.43.255.96:80 | ✓ 974ms | 否 | 否 | ✓ 1136ms | ✓ 925ms | http |
| 115.231.181.40:8128 | ✓ 893ms | 否 | ✓ 1180ms | 否 | ✓ 888ms | http |
| 34.101.184.164:3128 | ✓ 1435ms | 否 | ✓ 992ms | ✓ 1227ms | ✓ 949ms | http |
| 42.96.16.158:1311 | ✓ 1957ms | 否 | ✓ 1416ms | ✓ 1307ms | ✓ 1297ms | http |
| 45.136.130.191:8443 | ✓ 562ms | ✓ 643ms | ✓ 569ms | ✓ 650ms | ✓ 875ms | http |
| 45.136.130.188:8443 | ✓ 561ms | ✓ 1675ms | ✓ 81ms | ✓ 666ms | ✓ 510ms | http |
| 165.227.5.10:8888 | ✓ 552ms | ✓ 1170ms | ✓ 1628ms | ✓ 1974ms | ✓ 491ms | http |
| 95.3.9.78:8080 | 否 | 否 | ✓ 1168ms | ✓ 1916ms | ✓ 1464ms | http |
| 95.3.9.78:3128 | ✓ 1335ms | 否 | 否 | ✓ 1883ms | ✓ 1495ms | http |
| 121.230.8.34:1080 | 否 | ✓ 1503ms | ✓ 1771ms | ✓ 1749ms | ✓ 1102ms | http |
| 39.104.201.40:7890 | ✓ 894ms | ✓ 1163ms | ✓ 974ms | ✓ 1170ms | ✓ 910ms | http |
| 14.225.212.37:7890 | 否 | 否 | ✓ 1137ms | ✓ 1867ms | ✓ 1018ms | http |
| 91.107.141.42:8081 | ✓ 1769ms | 否 | 否 | ✓ 1828ms | ✓ 1680ms | http |
| 24.199.124.151:3128 | 否 | ✓ 1023ms | ✓ 1007ms | ✓ 658ms | ✓ 484ms | http |
| 107.172.125.217:3128 | ✓ 716ms | 否 | ✓ 1595ms | ✓ 840ms | ✓ 489ms | http |
| 120.92.212.16:8890 | ✓ 971ms | ✓ 1186ms | 否 | 否 | ✓ 946ms | http |
| 37.139.33.145:1080 | ✓ 804ms | 否 | ✓ 1734ms | 否 | ✓ 1874ms | http |
| 152.70.98.46:8888 | ✓ 654ms | 否 | ✓ 594ms | ✓ 905ms | ✓ 800ms | http |
| 116.236.189.93:29999 | ✓ 840ms | 否 | ✓ 829ms | 否 | ✓ 841ms | http |
| 138.124.90.140:1080 | ✓ 1214ms | ✓ 1715ms | ✓ 1768ms | ✓ 1552ms | 否 | http |
| 160.238.65.5:3128 | ✓ 1972ms | 否 | 否 | ✓ 1630ms | ✓ 1889ms | http |
| 45.136.198.40:3128 | ✓ 961ms | ✓ 1719ms | 否 | 否 | ✓ 1780ms | http |
| 86.53.183.16:1080 | ✓ 1340ms | 否 | ✓ 1496ms | 否 | ✓ 1607ms | http |
| 103.82.23.118:5171 | ✓ 1918ms | 否 | ✓ 1622ms | ✓ 1649ms | ✓ 1409ms | http |
| 185.191.236.162:3128 | ✓ 1836ms | 否 | ✓ 1270ms | 否 | ✓ 1359ms | http |
| 121.232.73.224:1080 | ✓ 888ms | ✓ 1211ms | 否 | ✓ 1499ms | ✓ 883ms | http |
| 103.84.95.54:7890 | ✓ 603ms | 否 | ✓ 630ms | ✓ 776ms | ✓ 659ms | http |
| 45.140.147.82:1082 | ✓ 1140ms | 否 | ✓ 1256ms | 否 | ✓ 1230ms | http |
| 103.39.51.190:8080 | ✓ 1813ms | 否 | ✓ 1729ms | ✓ 1679ms | ✓ 1523ms | http |
| 45.136.130.223:8443 | ✓ 610ms | ✓ 603ms | ✓ 252ms | ✓ 800ms | ✓ 508ms | http |
| 1.225.116.115:1080 | ✓ 1783ms | ✓ 1484ms | ✓ 1480ms | ✓ 1637ms | ✓ 965ms | http |
| 210.223.44.230:3128 | ✓ 1625ms | 否 | ✓ 1005ms | 否 | ✓ 1423ms | http |
| 45.186.6.104:3128 | ✓ 1084ms | ✓ 1801ms | ✓ 1784ms | 否 | 否 | http |
| 101.47.73.135:3128 | ✓ 1181ms | 否 | ✓ 1672ms | ✓ 1267ms | 否 | http |
| 190.6.54.12:6969 | ✓ 837ms | ✓ 1551ms | ✓ 1004ms | ✓ 1867ms | ✓ 1491ms | http |
| 113.176.92.71:3128 | ✓ 1920ms | ✓ 1778ms | 否 | 否 | ✓ 1869ms | http |
| 61.52.131.172:8443 | ✓ 849ms | ✓ 1065ms | ✓ 913ms | ✓ 1089ms | ✓ 872ms | http |
| 103.156.233.137:8080 | ✓ 1264ms | 否 | ✓ 1914ms | ✓ 1337ms | ✓ 1423ms | http |
| 62.113.119.14:8080 | ✓ 1292ms | 否 | ✓ 1435ms | ✓ 1658ms | ✓ 1210ms | http |
| 221.122.91.36:11195 | ✓ 1216ms | ✓ 1401ms | ✓ 995ms | ✓ 1196ms | ✓ 955ms | http |
| 221.122.91.36:11273 | ✓ 867ms | ✓ 1113ms | ✓ 832ms | ✓ 1139ms | ✓ 877ms | http |
| 45.140.147.155:1082 | ✓ 590ms | ✓ 1340ms | ✓ 1628ms | ✓ 1802ms | ✓ 1249ms | http |
| 116.80.49.166:3172 | ✓ 1490ms | 否 | ✓ 1518ms | ✓ 1777ms | 否 | http |
| 160.238.65.2:3128 | ✓ 681ms | ✓ 1703ms | ✓ 1830ms | ✓ 1746ms | 否 | http |
| 160.238.65.7:3128 | ✓ 673ms | 否 | ✓ 1350ms | ✓ 1598ms | ✓ 1811ms | http |
| 160.238.65.8:3128 | ✓ 677ms | 否 | ✓ 804ms | ✓ 1640ms | ✓ 1199ms | http |
| 160.238.65.6:3128 | ✓ 644ms | 否 | ✓ 1383ms | ✓ 1661ms | ✓ 1917ms | http |
| 160.238.65.3:3128 | ✓ 654ms | 否 | ✓ 1343ms | ✓ 1615ms | ✓ 1909ms | http |
| 160.238.65.4:3128 | ✓ 629ms | 否 | ✓ 1205ms | 否 | ✓ 1322ms | http |
| 152.42.213.210:8080 | ✓ 1304ms | 否 | ✓ 1311ms | ✓ 1329ms | ✓ 1083ms | http |
| 94.176.3.43:7443 | ✓ 1656ms | ✓ 1992ms | ✓ 1205ms | 否 | 否 | http |
| 138.124.53.25:7443 | ✓ 933ms | 否 | ✓ 1397ms | ✓ 1910ms | ✓ 1745ms | http |
| 57.128.188.167:8182 | ✓ 1773ms | 否 | ✓ 1904ms | 否 | ✓ 1947ms | http |
| 202.129.206.239:3128 | ✓ 1840ms | 否 | ✓ 1259ms | ✓ 1792ms | ✓ 1780ms | http |

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
