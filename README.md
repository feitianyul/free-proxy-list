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

最后更新：2026-03-30 16:05:04 UTC（2026-03-31 00:05:04 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 558ms | ✓ 1137ms | ✓ 1192ms | ✓ 924ms | ✓ 1082ms | http |
| 39.185.46.193:5911 | ✓ 862ms | ✓ 1045ms | ✓ 926ms | ✓ 1128ms | ✓ 849ms | http |
| 147.161.239.240:8800 | ✓ 628ms | 否 | ✓ 1222ms | 否 | ✓ 1311ms | http |
| 147.161.210.140:8800 | ✓ 1572ms | 否 | ✓ 1101ms | ✓ 983ms | ✓ 994ms | http |
| 95.213.217.168:52004 | ✓ 567ms | ✓ 1736ms | ✓ 1576ms | ✓ 1749ms | ✓ 1865ms | http |
| 1.231.81.166:3128 | ✓ 1596ms | ✓ 1393ms | ✓ 1352ms | ✓ 1362ms | ✓ 1012ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1556ms | ✓ 1590ms | ✓ 1771ms | ✓ 1208ms | http |
| 167.103.115.102:8800 | ✓ 1938ms | 否 | ✓ 1640ms | ✓ 1577ms | ✓ 1469ms | http |
| 43.99.54.236:5555 | ✓ 1699ms | ✓ 1235ms | ✓ 1928ms | ✓ 1324ms | ✓ 1893ms | http |
| 45.136.198.40:3128 | ✓ 619ms | ✓ 1537ms | ✓ 1568ms | ✓ 1936ms | ✓ 1509ms | http |
| 167.103.34.108:8800 | ✓ 1151ms | 否 | ✓ 1216ms | ✓ 1450ms | ✓ 1322ms | http |
| 120.92.212.16:7890 | ✓ 1144ms | ✓ 1499ms | 否 | ✓ 1499ms | ✓ 1135ms | http |
| 34.101.184.164:3128 | ✓ 1066ms | 否 | ✓ 1223ms | ✓ 1647ms | ✓ 1164ms | http |
| 167.103.31.122:8800 | ✓ 1788ms | 否 | ✓ 1885ms | 否 | ✓ 1920ms | http |
| 46.39.105.157:8080 | ✓ 1697ms | 否 | ✓ 999ms | ✓ 1854ms | ✓ 1326ms | http |
| 167.103.144.127:8800 | ✓ 1362ms | 否 | ✓ 1305ms | ✓ 1553ms | ✓ 1458ms | http |
| 160.238.65.8:3128 | ✓ 1833ms | ✓ 1684ms | 否 | 否 | ✓ 1690ms | http |
| 120.92.212.16:8890 | ✓ 1476ms | ✓ 1423ms | ✓ 1150ms | ✓ 1765ms | 否 | http |
| 5.104.87.17:8050 | ✓ 702ms | ✓ 1616ms | 否 | ✓ 1091ms | ✓ 863ms | http |
| 209.126.84.232:8888 | 否 | ✓ 1830ms | ✓ 1339ms | 否 | ✓ 1073ms | http |
| 8.219.97.248:80 | ✓ 1608ms | 否 | ✓ 1840ms | 否 | ✓ 1522ms | http |
| 150.241.71.15:1080 | ✓ 997ms | ✓ 1441ms | 否 | ✓ 1265ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1174ms | 否 | ✓ 1150ms | ✓ 1420ms | ✓ 1398ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1283ms | ✓ 1852ms | ✓ 1779ms | ✓ 1176ms | http |
| 177.234.217.88:999 | 否 | 否 | ✓ 1805ms | ✓ 1769ms | ✓ 1649ms | http |
| 5.104.87.17:8051 | ✓ 1564ms | 否 | ✓ 1641ms | 否 | ✓ 1750ms | http |
| 45.12.151.226:2829 | ✓ 983ms | 否 | ✓ 605ms | 否 | ✓ 1397ms | http |
| 150.107.140.238:3128 | ✓ 1671ms | 否 | 否 | ✓ 1388ms | ✓ 1063ms | http |
| 138.197.68.35:4857 | ✓ 906ms | ✓ 1698ms | ✓ 625ms | 否 | 否 | http |
| 101.32.244.83:8080 | ✓ 1198ms | ✓ 1937ms | ✓ 1161ms | ✓ 1546ms | ✓ 1505ms | http |
| 121.43.196.213:8222 | ✓ 1106ms | ✓ 1297ms | ✓ 1090ms | ✓ 1339ms | ✓ 1057ms | http |
| 121.43.196.210:8222 | ✓ 1433ms | ✓ 1236ms | ✓ 1081ms | ✓ 1305ms | ✓ 1024ms | http |
| 114.55.226.123:10086 | ✓ 1210ms | ✓ 1594ms | ✓ 1183ms | ✓ 1474ms | 否 | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1639ms | ✓ 1813ms | ✓ 1536ms | http |
| 144.124.227.88:3128 | ✓ 493ms | 否 | 否 | ✓ 1985ms | ✓ 1640ms | http |
| 103.113.70.189:1081 | 否 | 否 | ✓ 1142ms | ✓ 978ms | ✓ 1773ms | http |
| 5.102.109.41:999 | ✓ 1593ms | 否 | ✓ 1993ms | 否 | ✓ 1893ms | http |
| 219.117.204.211:7799 | ✓ 763ms | 否 | ✓ 1259ms | ✓ 1374ms | ✓ 1465ms | http |
| 31.192.106.135:8010 | ✓ 1813ms | 否 | ✓ 1015ms | 否 | ✓ 1633ms | http |
| 185.175.156.62:8080 | ✓ 1833ms | 否 | 否 | ✓ 1854ms | ✓ 1691ms | http |
| 103.84.95.54:7890 | ✓ 1232ms | 否 | 否 | ✓ 1146ms | ✓ 1080ms | http |
| 86.53.183.16:1080 | ✓ 1087ms | 否 | ✓ 1467ms | 否 | ✓ 1841ms | http |
| 121.126.185.63:25152 | ✓ 1660ms | 否 | ✓ 1457ms | 否 | ✓ 1676ms | http |
| 107.174.208.190:3128 | ✓ 340ms | 否 | ✓ 1948ms | ✓ 1225ms | ✓ 896ms | http |
| 194.87.73.134:3128 | ✓ 489ms | ✓ 1164ms | ✓ 1377ms | ✓ 1789ms | ✓ 1426ms | http |
| 23.95.76.201:8443 | ✓ 1153ms | ✓ 1225ms | 否 | ✓ 1624ms | ✓ 1031ms | http |
| 165.22.241.190:3128 | ✓ 1284ms | 否 | ✓ 1533ms | ✓ 1267ms | ✓ 1129ms | http |
| 195.123.213.129:1080 | ✓ 1158ms | ✓ 1765ms | ✓ 1440ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1177ms | 否 | 否 | ✓ 1549ms | ✓ 1193ms | http |
| 217.76.245.80:999 | ✓ 1731ms | ✓ 1227ms | ✓ 1227ms | ✓ 1289ms | ✓ 1385ms | http |
| 45.140.147.155:1081 | ✓ 1189ms | ✓ 1324ms | ✓ 1486ms | 否 | 否 | http |
| 208.87.243.199:7878 | ✓ 918ms | ✓ 1555ms | ✓ 389ms | ✓ 1673ms | ✓ 1095ms | http |
| 45.140.147.155:1082 | ✓ 954ms | ✓ 1354ms | ✓ 671ms | ✓ 1120ms | ✓ 937ms | http |
| 59.11.138.229:3128 | ✓ 811ms | 否 | ✓ 1184ms | ✓ 1223ms | ✓ 1064ms | http |
| 137.184.1.87:3128 | 否 | 否 | ✓ 1124ms | ✓ 1150ms | ✓ 751ms | http |
| 147.75.34.93:9400 | ✓ 1085ms | ✓ 1303ms | ✓ 718ms | ✓ 1502ms | ✓ 1293ms | http |
| 147.75.34.93:9480 | ✓ 1085ms | ✓ 1426ms | ✓ 687ms | ✓ 1509ms | ✓ 1234ms | http |
| 147.75.34.93:9401 | ✓ 1091ms | ✓ 1573ms | ✓ 585ms | ✓ 1537ms | ✓ 1254ms | http |
| 147.75.34.93:80 | ✓ 1090ms | 否 | ✓ 596ms | ✓ 1472ms | ✓ 1313ms | http |
| 180.103.19.53:1080 | ✓ 1421ms | 否 | ✓ 1493ms | 否 | ✓ 1367ms | http |
| 101.108.172.74:8080 | ✓ 1699ms | 否 | ✓ 1959ms | ✓ 1804ms | ✓ 1775ms | http |
| 170.78.208.251:999 | ✓ 751ms | ✓ 1190ms | ✓ 1530ms | 否 | 否 | http |
| 20.120.225.109:3128 | ✓ 1079ms | ✓ 1583ms | 否 | ✓ 1589ms | ✓ 708ms | http |
| 202.43.122.156:1111 | ✓ 1398ms | 否 | ✓ 1421ms | 否 | ✓ 1952ms | http |
| 103.39.51.190:8080 | ✓ 1945ms | 否 | 否 | ✓ 1776ms | ✓ 1670ms | http |
| 45.140.147.82:1081 | ✓ 577ms | ✓ 1388ms | ✓ 1342ms | ✓ 1110ms | ✓ 1002ms | http |
| 45.129.141.143:3128 | ✓ 1525ms | ✓ 1699ms | ✓ 1877ms | 否 | ✓ 1438ms | http |

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
