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

最后更新：2026-03-22 10:27:40 UTC（2026-03-22 18:27:40 UTC+8）

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
| 1.231.81.166:3128 | 否 | ✓ 1225ms | ✓ 1611ms | ✓ 1129ms | ✓ 1006ms | http |
| 113.160.132.26:8080 | ✓ 1690ms | ✓ 1403ms | ✓ 1286ms | ✓ 1159ms | ✓ 1261ms | http |
| 137.220.150.104:6005 | ✓ 1003ms | 否 | 否 | ✓ 1428ms | ✓ 809ms | http |
| 147.161.239.240:8800 | ✓ 934ms | 否 | ✓ 1541ms | ✓ 1511ms | ✓ 1672ms | http |
| 167.103.34.108:8800 | ✓ 1784ms | 否 | ✓ 1533ms | ✓ 1548ms | ✓ 1543ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1865ms | ✓ 1625ms | ✓ 1827ms | http |
| 45.167.124.52:8080 | ✓ 1560ms | 否 | ✓ 915ms | ✓ 1740ms | ✓ 1445ms | http |
| 142.171.224.229:7890 | ✓ 366ms | ✓ 1688ms | ✓ 702ms | ✓ 679ms | ✓ 532ms | http |
| 147.161.210.140:8800 | ✓ 534ms | ✓ 1998ms | ✓ 536ms | ✓ 725ms | ✓ 773ms | http |
| 35.225.22.61:80 | ✓ 506ms | 否 | ✓ 1060ms | ✓ 1249ms | ✓ 966ms | http |
| 172.93.42.71:3128 | ✓ 242ms | 否 | 否 | ✓ 668ms | ✓ 512ms | http |
| 137.220.150.22:6005 | ✓ 787ms | 否 | ✓ 860ms | ✓ 1181ms | ✓ 995ms | http |
| 120.92.212.16:8890 | ✓ 894ms | ✓ 1177ms | ✓ 926ms | ✓ 1178ms | ✓ 941ms | http |
| 77.232.135.22:1080 | ✓ 921ms | 否 | ✓ 1973ms | ✓ 1607ms | ✓ 1224ms | http |
| 144.31.79.117:8888 | ✓ 776ms | 否 | ✓ 1247ms | 否 | ✓ 1829ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1414ms | ✓ 1793ms | ✓ 1432ms | ✓ 1126ms | http |
| 45.168.238.193:8443 | ✓ 308ms | ✓ 1034ms | ✓ 270ms | ✓ 1211ms | ✓ 925ms | http |
| 104.168.158.236:10808 | ✓ 470ms | ✓ 1092ms | ✓ 491ms | ✓ 1122ms | ✓ 783ms | http |
| 101.47.73.135:3128 | ✓ 1062ms | 否 | ✓ 1236ms | ✓ 1084ms | ✓ 974ms | http |
| 167.103.31.122:8800 | ✓ 1329ms | 否 | ✓ 1353ms | ✓ 1616ms | 否 | http |
| 137.220.150.170:6005 | ✓ 1017ms | 否 | ✓ 996ms | 否 | ✓ 1054ms | http |
| 103.113.70.189:1081 | ✓ 1295ms | 否 | ✓ 522ms | ✓ 1195ms | 否 | http |
| 106.75.15.167:7890 | ✓ 871ms | ✓ 1110ms | ✓ 1579ms | 否 | ✓ 911ms | http |
| 116.80.65.79:3172 | 否 | 否 | ✓ 1445ms | ✓ 1781ms | ✓ 1649ms | http |
| 103.84.95.54:7890 | ✓ 792ms | 否 | ✓ 604ms | ✓ 1397ms | 否 | http |
| 160.250.4.245:1 | ✓ 1446ms | 否 | ✓ 1113ms | ✓ 1564ms | ✓ 1117ms | http |
| 115.231.181.40:8128 | ✓ 914ms | ✓ 1663ms | ✓ 1418ms | 否 | ✓ 1931ms | http |
| 137.220.150.152:6005 | ✓ 1578ms | 否 | ✓ 888ms | ✓ 1261ms | ✓ 1798ms | http |
| 222.228.171.92:8080 | ✓ 1850ms | 否 | ✓ 1524ms | ✓ 831ms | ✓ 701ms | http |
| 218.89.134.230:3333 | ✓ 1386ms | ✓ 1564ms | ✓ 1463ms | ✓ 1523ms | ✓ 1247ms | http |
| 49.156.44.114:8080 | ✓ 1661ms | 否 | ✓ 1408ms | ✓ 1373ms | ✓ 1350ms | http |
| 113.255.59.226:8080 | ✓ 1163ms | 否 | ✓ 1382ms | ✓ 1478ms | 否 | http |
| 181.78.194.249:999 | 否 | 否 | ✓ 1828ms | ✓ 1979ms | ✓ 1779ms | http |
| 194.67.99.223:1080 | ✓ 943ms | 否 | 否 | ✓ 1950ms | ✓ 1510ms | http |
| 8.219.97.248:80 | 否 | ✓ 1874ms | ✓ 1151ms | ✓ 1117ms | 否 | http |
| 85.198.96.242:3128 | ✓ 903ms | 否 | ✓ 1923ms | ✓ 1856ms | ✓ 1522ms | http |
| 45.144.28.81:10808 | ✓ 585ms | 否 | ✓ 1450ms | 否 | ✓ 1437ms | http |
| 162.240.154.26:3128 | ✓ 1032ms | ✓ 1623ms | ✓ 1986ms | 否 | 否 | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1326ms | ✓ 1719ms | ✓ 1906ms | http |
| 5.102.109.41:999 | ✓ 1862ms | 否 | ✓ 1094ms | 否 | ✓ 1064ms | http |
| 101.32.244.83:8080 | ✓ 952ms | ✓ 1864ms | ✓ 883ms | ✓ 1303ms | ✓ 1194ms | http |
| 121.43.196.210:8222 | ✓ 940ms | ✓ 1026ms | ✓ 844ms | ✓ 1056ms | ✓ 840ms | http |
| 121.43.196.213:8222 | ✓ 930ms | ✓ 1001ms | ✓ 881ms | ✓ 1087ms | ✓ 846ms | http |
| 114.55.226.123:10086 | ✓ 1013ms | ✓ 1292ms | ✓ 970ms | ✓ 1235ms | ✓ 1028ms | http |
| 150.107.140.238:3128 | ✓ 1591ms | 否 | ✓ 984ms | ✓ 1140ms | ✓ 981ms | http |
| 91.238.105.64:2024 | ✓ 957ms | 否 | ✓ 1721ms | 否 | ✓ 1876ms | http |
| 27.254.99.183:8118 | 否 | ✓ 1908ms | ✓ 1171ms | 否 | ✓ 998ms | http |
| 185.241.5.57:3128 | ✓ 1037ms | 否 | ✓ 1251ms | 否 | ✓ 1812ms | http |
| 38.34.183.233:8448 | ✓ 413ms | ✓ 1221ms | ✓ 482ms | ✓ 1237ms | ✓ 1058ms | http |
| 38.34.183.224:8448 | ✓ 497ms | ✓ 656ms | ✓ 417ms | 否 | ✓ 1418ms | http |
| 62.113.119.14:8080 | ✓ 1412ms | 否 | ✓ 826ms | ✓ 1648ms | ✓ 1277ms | http |
| 38.145.218.228:8447 | ✓ 790ms | ✓ 1234ms | ✓ 1811ms | ✓ 1381ms | ✓ 552ms | http |
| 166.88.55.83:7890 | ✓ 598ms | ✓ 1049ms | ✓ 598ms | ✓ 745ms | ✓ 603ms | http |
| 202.141.161.53:30001 | ✓ 1019ms | ✓ 1416ms | 否 | ✓ 1149ms | ✓ 1020ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1913ms | ✓ 1079ms | ✓ 1278ms | ✓ 977ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1316ms | ✓ 1790ms | ✓ 1339ms | 否 | http |
| 45.151.183.183:1080 | ✓ 1019ms | 否 | ✓ 776ms | ✓ 1883ms | ✓ 1427ms | http |
| 45.136.198.40:3128 | ✓ 1335ms | 否 | ✓ 1368ms | 否 | ✓ 1949ms | http |
| 45.140.147.155:1081 | ✓ 605ms | ✓ 1648ms | ✓ 976ms | ✓ 1815ms | 否 | http |
| 172.212.68.37:3128 | 否 | 否 | ✓ 748ms | ✓ 1707ms | ✓ 1333ms | http |
| 181.78.44.63:999 | ✓ 973ms | ✓ 1682ms | ✓ 1970ms | ✓ 1564ms | ✓ 1615ms | http |
| 101.43.127.100:8877 | ✓ 903ms | ✓ 1044ms | ✓ 895ms | ✓ 1123ms | ✓ 1630ms | http |
| 103.39.51.190:8080 | ✓ 1689ms | 否 | ✓ 1670ms | ✓ 1464ms | ✓ 1331ms | http |
| 150.241.77.172:1080 | ✓ 801ms | 否 | ✓ 1715ms | 否 | ✓ 1577ms | http |
| 34.150.20.6:8888 | ✓ 1198ms | ✓ 1702ms | ✓ 1076ms | ✓ 1196ms | ✓ 1451ms | http |
| 103.139.138.194:3128 | ✓ 1891ms | 否 | 否 | ✓ 1431ms | ✓ 1893ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1321ms | ✓ 1274ms | 否 | ✓ 1161ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1845ms | ✓ 893ms | 否 | ✓ 1988ms | http |
| 47.77.193.180:1080 | ✓ 715ms | ✓ 1392ms | ✓ 242ms | ✓ 638ms | ✓ 546ms | http |
| 85.208.108.43:2094 | ✓ 1222ms | 否 | ✓ 1420ms | ✓ 1284ms | ✓ 839ms | http |
| 85.208.108.43:10808 | ✓ 1207ms | 否 | ✓ 1421ms | ✓ 1269ms | ✓ 848ms | http |
| 38.34.179.186:8444 | ✓ 1080ms | ✓ 1821ms | ✓ 383ms | ✓ 691ms | ✓ 735ms | http |
| 38.34.179.16:8451 | ✓ 1347ms | 否 | ✓ 695ms | 否 | ✓ 1695ms | http |

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
