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

最后更新：2026-03-15 18:31:06 UTC（2026-03-16 02:31:06 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 862ms | 否 | ✓ 1032ms | 否 | ✓ 1008ms | http |
| 205.209.118.30:3138 | ✓ 640ms | ✓ 987ms | ✓ 1114ms | ✓ 1001ms | ✓ 773ms | http |
| 149.50.116.240:1080 | ✓ 963ms | ✓ 1742ms | ✓ 1636ms | ✓ 1928ms | 否 | http |
| 194.87.43.49:8888 | ✓ 1292ms | 否 | ✓ 1424ms | 否 | ✓ 1962ms | http |
| 5.129.206.247:8888 | ✓ 1040ms | 否 | ✓ 1828ms | 否 | ✓ 1948ms | http |
| 115.231.181.40:8128 | ✓ 1415ms | 否 | ✓ 1330ms | ✓ 1349ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1299ms | 否 | ✓ 1238ms | ✓ 1670ms | ✓ 1524ms | http |
| 104.129.202.127:12354 | ✓ 708ms | ✓ 1531ms | ✓ 1289ms | ✓ 1252ms | ✓ 850ms | http |
| 160.238.65.3:3128 | ✓ 916ms | ✓ 1187ms | ✓ 807ms | ✓ 1249ms | ✓ 972ms | http |
| 113.160.132.26:8080 | ✓ 1358ms | ✓ 1479ms | ✓ 1003ms | ✓ 1388ms | ✓ 1119ms | http |
| 72.11.150.178:6005 | ✓ 439ms | ✓ 1967ms | ✓ 1775ms | ✓ 1395ms | ✓ 776ms | http |
| 101.43.255.96:80 | ✓ 1068ms | ✓ 1672ms | ✓ 1111ms | 否 | ✓ 1087ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1235ms | ✓ 1057ms | ✓ 1236ms | ✓ 1796ms | http |
| 38.145.218.82:8443 | ✓ 468ms | ✓ 958ms | ✓ 948ms | ✓ 1091ms | ✓ 840ms | http |
| 160.238.65.2:3128 | ✓ 840ms | ✓ 1330ms | ✓ 919ms | ✓ 1317ms | ✓ 971ms | http |
| 160.238.65.9:3128 | ✓ 837ms | ✓ 1306ms | ✓ 946ms | ✓ 1313ms | ✓ 984ms | http |
| 160.238.65.8:3128 | ✓ 811ms | ✓ 1247ms | ✓ 424ms | ✓ 1418ms | ✓ 967ms | http |
| 160.238.65.5:3128 | ✓ 815ms | ✓ 1247ms | ✓ 438ms | ✓ 1425ms | ✓ 984ms | http |
| 160.238.65.7:3128 | ✓ 814ms | ✓ 1989ms | ✓ 443ms | ✓ 1318ms | ✓ 989ms | http |
| 137.220.150.152:6005 | ✓ 1033ms | 否 | ✓ 889ms | ✓ 1311ms | ✓ 1203ms | http |
| 137.220.150.104:6005 | ✓ 1032ms | 否 | ✓ 1137ms | ✓ 1656ms | ✓ 1263ms | http |
| 81.70.169.194:80 | ✓ 1475ms | ✓ 1486ms | ✓ 1293ms | ✓ 1465ms | 否 | http |
| 95.3.9.78:8080 | ✓ 1799ms | 否 | ✓ 1630ms | ✓ 1966ms | ✓ 1632ms | http |
| 45.136.130.223:8443 | ✓ 297ms | ✓ 1108ms | ✓ 331ms | ✓ 1078ms | ✓ 741ms | http |
| 160.238.65.6:3128 | ✓ 1139ms | ✓ 1451ms | ✓ 1180ms | ✓ 1259ms | ✓ 1249ms | http |
| 160.238.65.4:3128 | ✓ 1134ms | ✓ 1473ms | ✓ 1161ms | ✓ 1245ms | ✓ 1271ms | http |
| 95.3.9.78:3128 | ✓ 680ms | ✓ 1705ms | ✓ 690ms | ✓ 1616ms | ✓ 1253ms | http |
| 165.225.72.38:10801 | ✓ 584ms | 否 | ✓ 822ms | ✓ 1369ms | ✓ 1078ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1218ms | ✓ 436ms | ✓ 891ms | 否 | http |
| 109.173.19.42:3128 | ✓ 1162ms | 否 | ✓ 1024ms | 否 | ✓ 1760ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1785ms | ✓ 1024ms | ✓ 1354ms | ✓ 1588ms | http |
| 92.119.127.212:6005 | ✓ 918ms | 否 | ✓ 1510ms | ✓ 1997ms | 否 | http |
| 85.198.96.242:3128 | ✓ 1736ms | 否 | ✓ 1050ms | ✓ 1928ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1147ms | ✓ 1481ms | ✓ 1191ms | ✓ 1553ms | ✓ 1157ms | http |
| 137.220.151.110:6005 | ✓ 968ms | 否 | ✓ 940ms | ✓ 1381ms | 否 | http |
| 104.129.202.127:10810 | ✓ 1801ms | 否 | ✓ 1963ms | ✓ 1882ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1079ms | 否 | 否 | ✓ 1364ms | ✓ 1094ms | http |
| 157.0.142.246:10057 | ✓ 1050ms | ✓ 1319ms | ✓ 1191ms | ✓ 1381ms | ✓ 1048ms | http |
| 43.225.151.194:43030 | ✓ 1688ms | 否 | ✓ 1967ms | 否 | ✓ 1991ms | http |
| 121.230.9.19:1080 | ✓ 1663ms | ✓ 1589ms | ✓ 1281ms | ✓ 1612ms | ✓ 1465ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1574ms | ✓ 1012ms | ✓ 1072ms | 否 | http |
| 178.236.245.59:3128 | 否 | 否 | ✓ 1950ms | ✓ 1811ms | ✓ 1306ms | http |
| 120.92.212.16:8890 | ✓ 1235ms | ✓ 1319ms | 否 | ✓ 1389ms | ✓ 1070ms | http |
| 178.156.224.42:3128 | ✓ 1179ms | 否 | ✓ 1946ms | 否 | ✓ 1623ms | http |
| 83.219.250.8:62920 | ✓ 760ms | 否 | ✓ 765ms | 否 | ✓ 1548ms | http |
| 88.80.150.82:8080 | ✓ 1319ms | 否 | ✓ 1984ms | 否 | ✓ 1803ms | https |
| 142.93.195.158:80 | ✓ 1429ms | ✓ 1226ms | ✓ 1074ms | 否 | ✓ 916ms | http |
| 172.212.68.37:3128 | ✓ 857ms | ✓ 1392ms | 否 | ✓ 1370ms | ✓ 1734ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 722ms | ✓ 1125ms | ✓ 1102ms | http |
| 45.149.92.147:5001 | ✓ 743ms | 否 | ✓ 786ms | ✓ 989ms | ✓ 748ms | http |
| 62.113.119.14:8080 | ✓ 1110ms | 否 | ✓ 986ms | ✓ 1499ms | ✓ 1172ms | http |
| 38.180.2.107:3128 | ✓ 1252ms | 否 | ✓ 1849ms | 否 | ✓ 1988ms | http |
| 2.56.122.146:10808 | ✓ 1471ms | ✓ 1260ms | ✓ 1020ms | ✓ 1643ms | ✓ 1214ms | http |
| 47.84.131.156:8100 | ✓ 1327ms | ✓ 1817ms | ✓ 874ms | ✓ 1182ms | 否 | http |
| 103.139.138.194:3128 | ✓ 1405ms | 否 | ✓ 1209ms | ✓ 1507ms | ✓ 1186ms | http |
| 118.31.1.154:80 | 否 | 否 | ✓ 957ms | ✓ 1388ms | ✓ 1241ms | http |
| 165.225.72.38:10960 | ✓ 903ms | 否 | ✓ 1061ms | 否 | ✓ 1595ms | http |
| 45.136.198.40:3128 | ✓ 1620ms | ✓ 1878ms | 否 | 否 | ✓ 1974ms | http |
| 103.39.51.190:8080 | ✓ 1889ms | 否 | 否 | ✓ 1544ms | ✓ 1612ms | http |
| 38.145.203.135:8443 | ✓ 359ms | ✓ 939ms | ✓ 510ms | ✓ 966ms | ✓ 715ms | http |
| 164.92.148.68:3128 | ✓ 624ms | ✓ 1412ms | ✓ 1867ms | ✓ 1772ms | ✓ 1329ms | http |
| 46.39.105.157:8080 | ✓ 748ms | 否 | 否 | ✓ 1937ms | ✓ 1675ms | http |
| 47.95.231.180:8084 | 否 | ✓ 1618ms | 否 | ✓ 1298ms | ✓ 1033ms | http |
| 121.40.231.103:7890 | ✓ 1346ms | 否 | ✓ 1298ms | ✓ 1398ms | 否 | http |
| 62.234.206.73:3128 | 否 | ✓ 1439ms | ✓ 1158ms | 否 | ✓ 1018ms | http |
| 207.254.71.62:8088 | ✓ 583ms | 否 | ✓ 1919ms | ✓ 1717ms | ✓ 1895ms | http |
| 45.207.200.120:1080 | ✓ 1149ms | ✓ 1764ms | ✓ 1166ms | 否 | 否 | http |
| 45.186.6.104:3128 | ✓ 1200ms | ✓ 1637ms | ✓ 1618ms | 否 | 否 | http |
| 192.71.213.85:9091 | ✓ 811ms | 否 | ✓ 720ms | ✓ 1958ms | 否 | http |
| 192.71.213.85:9812 | ✓ 477ms | 否 | ✓ 479ms | ✓ 1318ms | 否 | http |
| 218.60.0.214:80 | 否 | 否 | ✓ 1183ms | ✓ 1444ms | ✓ 1130ms | http |
| 47.105.98.23:3128 | ✓ 1005ms | 否 | 否 | ✓ 1654ms | ✓ 1579ms | http |
| 198.24.188.138:38050 | ✓ 721ms | ✓ 1493ms | ✓ 1395ms | ✓ 1742ms | 否 | http |
| 103.247.13.131:8085 | ✓ 1910ms | 否 | 否 | ✓ 1677ms | ✓ 1633ms | http |
| 162.240.154.26:3128 | ✓ 775ms | ✓ 1454ms | 否 | ✓ 1578ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1165ms | ✓ 1450ms | ✓ 1121ms | ✓ 1446ms | ✓ 1083ms | http |
| 45.168.238.193:8443 | ✓ 390ms | ✓ 1178ms | ✓ 355ms | ✓ 1243ms | ✓ 897ms | http |
| 103.113.70.189:1081 | ✓ 229ms | ✓ 1831ms | 否 | ✓ 1032ms | ✓ 725ms | http |
| 45.136.130.245:8447 | 否 | ✓ 1578ms | ✓ 861ms | ✓ 1063ms | ✓ 1113ms | http |
| 143.244.140.119:3128 | ✓ 1391ms | 否 | ✓ 1826ms | ✓ 1722ms | 否 | http |

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
