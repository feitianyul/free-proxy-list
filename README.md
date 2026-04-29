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

最后更新：2026-04-29 21:58:11 UTC（2026-04-30 05:58:11 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | 否 | ✓ 1067ms | ✓ 1440ms | ✓ 1020ms | ✓ 905ms | http |
| 34.71.229.255:3128 | ✓ 1300ms | 否 | ✓ 1562ms | ✓ 1344ms | ✓ 1597ms | http |
| 113.160.132.26:8080 | ✓ 1443ms | ✓ 1327ms | ✓ 1075ms | ✓ 1244ms | ✓ 1038ms | http |
| 45.167.124.71:999 | ✓ 1361ms | ✓ 1807ms | ✓ 1638ms | ✓ 1948ms | ✓ 1892ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1187ms | ✓ 1203ms | 否 | ✓ 1108ms | http |
| 45.88.0.115:3128 | ✓ 1132ms | ✓ 1290ms | ✓ 1464ms | 否 | ✓ 1652ms | http |
| 45.88.0.113:3128 | ✓ 1139ms | ✓ 1482ms | ✓ 1266ms | 否 | ✓ 1690ms | http |
| 45.88.0.117:3128 | ✓ 1138ms | ✓ 1494ms | ✓ 1260ms | 否 | ✓ 1684ms | http |
| 45.88.0.116:3128 | ✓ 1137ms | ✓ 1468ms | ✓ 1287ms | 否 | ✓ 1649ms | http |
| 213.220.62.63:3128 | ✓ 1132ms | ✓ 1486ms | ✓ 1273ms | 否 | ✓ 1647ms | http |
| 45.88.0.114:3128 | ✓ 1132ms | ✓ 1348ms | ✓ 1411ms | 否 | ✓ 1648ms | http |
| 45.88.0.111:3128 | ✓ 1137ms | ✓ 1476ms | ✓ 1273ms | 否 | ✓ 1641ms | http |
| 213.220.62.62:3128 | ✓ 1137ms | ✓ 1294ms | ✓ 1460ms | 否 | ✓ 1644ms | http |
| 45.88.0.99:3128 | ✓ 1132ms | ✓ 1472ms | ✓ 1287ms | 否 | ✓ 1687ms | http |
| 45.88.0.98:3128 | ✓ 1133ms | ✓ 1302ms | ✓ 1453ms | 否 | ✓ 1691ms | http |
| 47.85.51.197:1080 | ✓ 724ms | ✓ 1135ms | ✓ 609ms | 否 | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1224ms | 否 | ✓ 1453ms | ✓ 1021ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1641ms | ✓ 1703ms | ✓ 1465ms | http |
| 218.108.131.186:17890 | ✓ 1135ms | 否 | 否 | ✓ 1391ms | ✓ 1017ms | http |
| 152.32.132.190:7890 | ✓ 1921ms | ✓ 1085ms | ✓ 1466ms | ✓ 993ms | ✓ 1634ms | http |
| 8.154.21.175:3128 | ✓ 919ms | ✓ 1116ms | ✓ 926ms | ✓ 1155ms | ✓ 946ms | http |
| 34.246.223.187:44458 | ✓ 1378ms | 否 | ✓ 670ms | 否 | ✓ 1241ms | http |
| 185.41.152.110:3128 | ✓ 1003ms | ✓ 1511ms | ✓ 1704ms | ✓ 1825ms | ✓ 1862ms | http |
| 34.241.123.181:25193 | ✓ 1342ms | 否 | ✓ 1851ms | 否 | ✓ 1783ms | http |
| 16.62.123.236:8088 | ✓ 1360ms | 否 | ✓ 1940ms | 否 | ✓ 1416ms | http |
| 3.121.42.224:21347 | ✓ 1261ms | 否 | 否 | ✓ 1809ms | ✓ 1722ms | http |
| 46.101.95.183:8888 | ✓ 1591ms | 否 | ✓ 843ms | 否 | ✓ 1949ms | http |
| 52.59.51.29:3128 | ✓ 1669ms | 否 | ✓ 672ms | ✓ 1611ms | ✓ 1449ms | http |
| 130.61.174.200:1080 | ✓ 1380ms | 否 | 否 | ✓ 1721ms | ✓ 1171ms | http |
| 123.139.36.231:8118 | 否 | ✓ 1573ms | ✓ 1149ms | 否 | ✓ 1489ms | http |
| 159.223.225.118:8888 | 否 | 否 | ✓ 1201ms | ✓ 1856ms | ✓ 1449ms | http |
| 212.58.132.5:8888 | ✓ 1935ms | 否 | ✓ 1355ms | ✓ 1530ms | ✓ 1323ms | http |
| 206.206.126.177:2412 | ✓ 829ms | 否 | ✓ 995ms | 否 | ✓ 1327ms | http |
| 120.92.212.16:7890 | ✓ 973ms | 否 | ✓ 1027ms | 否 | ✓ 1046ms | http |
| 103.157.200.126:3128 | ✓ 1538ms | 否 | ✓ 1683ms | 否 | ✓ 1624ms | http |
| 86.104.74.110:1081 | 否 | ✓ 1535ms | 否 | ✓ 1836ms | ✓ 1409ms | http |
| 77.110.116.224:3128 | ✓ 693ms | 否 | ✓ 819ms | ✓ 1536ms | ✓ 1452ms | http |
| 154.64.232.35:8080 | ✓ 950ms | 否 | ✓ 1255ms | ✓ 1324ms | ✓ 1865ms | http |
| 183.238.3.150:7897 | ✓ 1981ms | ✓ 1221ms | ✓ 1030ms | 否 | 否 | http |
| 103.122.64.163:8080 | ✓ 1780ms | 否 | ✓ 1520ms | ✓ 1426ms | 否 | http |
| 121.130.177.28:8888 | ✓ 999ms | ✓ 1295ms | ✓ 1525ms | ✓ 1557ms | ✓ 1228ms | http |
| 86.104.74.110:1082 | ✓ 1310ms | ✓ 1431ms | ✓ 1017ms | ✓ 1789ms | ✓ 1561ms | http |
| 120.92.212.16:8890 | ✓ 1629ms | 否 | ✓ 1638ms | ✓ 1278ms | ✓ 962ms | http |
| 34.101.184.164:3128 | ✓ 1522ms | 否 | ✓ 1811ms | ✓ 1644ms | ✓ 1110ms | http |
| 183.232.248.73:7890 | ✓ 993ms | ✓ 1219ms | ✓ 1040ms | ✓ 1097ms | ✓ 980ms | http |
| 103.3.246.71:3128 | ✓ 1302ms | 否 | ✓ 1335ms | ✓ 1234ms | ✓ 1015ms | http |
| 172.236.145.31:7890 | ✓ 835ms | 否 | ✓ 772ms | ✓ 1076ms | ✓ 887ms | http |
| 120.92.108.86:7890 | ✓ 1464ms | 否 | ✓ 1570ms | 否 | ✓ 1311ms | http |
| 8.209.238.110:47701 | ✓ 1058ms | ✓ 1116ms | ✓ 1238ms | ✓ 905ms | ✓ 695ms | http |
| 150.249.255.91:3128 | ✓ 1087ms | 否 | ✓ 723ms | 否 | ✓ 1546ms | http |
| 118.113.247.154:1080 | ✓ 1283ms | ✓ 1566ms | ✓ 1360ms | ✓ 1590ms | ✓ 1311ms | http |
| 3.121.42.224:49441 | ✓ 894ms | 否 | ✓ 1198ms | 否 | ✓ 1897ms | http |
| 16.62.229.137:53444 | ✓ 912ms | 否 | ✓ 1208ms | ✓ 1939ms | 否 | http |
| 15.160.116.45:57416 | ✓ 1003ms | 否 | ✓ 1217ms | 否 | ✓ 1981ms | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 678ms | ✓ 1158ms | ✓ 1052ms | http |
| 18.222.132.180:35031 | ✓ 1105ms | 否 | ✓ 1657ms | ✓ 1727ms | ✓ 1334ms | http |
| 3.238.34.111:44677 | ✓ 1126ms | 否 | ✓ 1817ms | ✓ 1981ms | ✓ 1580ms | http |
| 13.60.181.61:4192 | ✓ 1209ms | 否 | ✓ 1864ms | ✓ 1981ms | ✓ 1550ms | http |
| 47.105.98.23:3128 | ✓ 898ms | ✓ 1181ms | ✓ 1361ms | ✓ 1262ms | ✓ 1133ms | http |
| 157.0.142.246:10057 | 否 | ✓ 1324ms | ✓ 1080ms | ✓ 1328ms | ✓ 1074ms | http |
| 42.200.76.16:3888 | ✓ 1826ms | 否 | ✓ 1107ms | ✓ 968ms | ✓ 769ms | http |
| 91.233.223.147:3128 | ✓ 1202ms | 否 | ✓ 1002ms | 否 | ✓ 1626ms | http |
| 217.182.195.221:30001 | ✓ 1030ms | 否 | ✓ 1650ms | ✓ 1989ms | 否 | http |
| 59.46.216.131:30001 | ✓ 998ms | 否 | ✓ 1063ms | ✓ 1341ms | ✓ 1039ms | http |
| 45.153.231.229:8080 | ✓ 1184ms | ✓ 1860ms | ✓ 1689ms | ✓ 1825ms | 否 | http |
| 121.230.9.160:1080 | 否 | 否 | ✓ 1337ms | ✓ 1765ms | ✓ 1702ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1335ms | ✓ 1317ms | ✓ 1056ms | http |
| 20.127.128.70:8080 | ✓ 1856ms | 否 | ✓ 1598ms | 否 | ✓ 1832ms | http |
| 86.104.72.220:1081 | ✓ 913ms | ✓ 1447ms | ✓ 387ms | ✓ 1571ms | ✓ 1258ms | http |
| 86.104.72.220:1082 | 否 | 否 | ✓ 609ms | ✓ 1307ms | ✓ 1919ms | http |
| 61.52.131.172:8443 | 否 | ✓ 1260ms | ✓ 1682ms | ✓ 1229ms | ✓ 1025ms | http |
| 54.189.56.88:45841 | ✓ 1957ms | 否 | ✓ 1444ms | 否 | ✓ 1827ms | http |
| 3.71.26.7:16789 | ✓ 1815ms | 否 | ✓ 1470ms | ✓ 1646ms | ✓ 1539ms | http |
| 13.41.196.179:23714 | ✓ 1248ms | 否 | ✓ 1993ms | 否 | ✓ 1676ms | http |
| 13.51.196.44:27491 | ✓ 1712ms | 否 | ✓ 1251ms | 否 | ✓ 1927ms | http |
| 51.95.13.205:57352 | ✓ 1979ms | 否 | ✓ 1165ms | 否 | ✓ 1819ms | http |
| 51.84.101.19:46454 | ✓ 1603ms | 否 | ✓ 1582ms | 否 | ✓ 1939ms | http |
| 210.223.44.230:3128 | ✓ 1585ms | 否 | ✓ 980ms | 否 | ✓ 775ms | http |
| 20.164.75.153:8080 | ✓ 1818ms | 否 | ✓ 1884ms | 否 | ✓ 1901ms | http |
| 101.32.243.189:80 | ✓ 1186ms | ✓ 1480ms | ✓ 1992ms | ✓ 1509ms | ✓ 1296ms | http |
| 103.39.51.207:8080 | ✓ 1336ms | 否 | ✓ 1282ms | ✓ 1803ms | ✓ 1547ms | http |
| 38.92.10.98:20058 | ✓ 614ms | ✓ 996ms | ✓ 1054ms | 否 | 否 | http |

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
