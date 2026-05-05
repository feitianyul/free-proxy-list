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

最后更新：2026-05-05 12:08:45 UTC（2026-05-05 20:08:45 UTC+8）

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
| 34.71.229.255:3128 | ✓ 453ms | 否 | ✓ 1094ms | ✓ 1597ms | ✓ 829ms | http |
| 47.85.51.197:1080 | ✓ 561ms | 否 | 否 | ✓ 1292ms | ✓ 858ms | http |
| 1.231.81.166:3128 | ✓ 1727ms | ✓ 1381ms | ✓ 1367ms | ✓ 999ms | ✓ 1039ms | http |
| 152.70.91.193:40000 | ✓ 1820ms | 否 | ✓ 1590ms | 否 | ✓ 1252ms | http |
| 185.191.236.162:3128 | ✓ 1328ms | 否 | ✓ 1873ms | 否 | ✓ 1613ms | http |
| 113.160.132.26:8080 | ✓ 1842ms | ✓ 1645ms | ✓ 1053ms | ✓ 1530ms | 否 | http |
| 38.180.192.119:3128 | ✓ 1280ms | ✓ 1822ms | 否 | ✓ 1630ms | 否 | http |
| 181.119.97.24:999 | ✓ 1514ms | 否 | ✓ 1950ms | ✓ 1943ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1031ms | ✓ 1104ms | ✓ 910ms | ✓ 1118ms | ✓ 926ms | http |
| 86.104.72.220:1081 | ✓ 309ms | 否 | ✓ 263ms | ✓ 1393ms | 否 | http |
| 45.153.231.229:8080 | ✓ 924ms | ✓ 1826ms | ✓ 1677ms | 否 | 否 | http |
| 43.133.44.89:8888 | ✓ 1738ms | 否 | ✓ 1919ms | ✓ 1042ms | ✓ 1691ms | http |
| 193.123.250.39:1080 | 否 | 否 | ✓ 1601ms | ✓ 1650ms | ✓ 1117ms | http |
| 154.64.232.35:8080 | ✓ 636ms | ✓ 1592ms | ✓ 1752ms | 否 | 否 | http |
| 217.76.245.80:999 | ✓ 1022ms | ✓ 1669ms | ✓ 1189ms | ✓ 1726ms | ✓ 1654ms | http |
| 5.63.111.238:8080 | ✓ 1294ms | 否 | ✓ 1267ms | ✓ 1952ms | ✓ 1792ms | http |
| 62.133.60.126:24558 | ✓ 1075ms | ✓ 1805ms | 否 | ✓ 1958ms | 否 | http |
| 165.225.113.220:8800 | ✓ 1401ms | 否 | 否 | ✓ 1314ms | ✓ 897ms | http |
| 91.238.104.171:2023 | ✓ 1248ms | 否 | ✓ 1647ms | ✓ 1614ms | ✓ 1226ms | http |
| 14.143.222.113:57788 | ✓ 1425ms | 否 | ✓ 1103ms | ✓ 1369ms | 否 | http |
| 86.104.72.219:1082 | ✓ 439ms | 否 | ✓ 545ms | 否 | ✓ 1739ms | http |
| 212.58.132.5:8888 | ✓ 1307ms | 否 | ✓ 1810ms | ✓ 1582ms | ✓ 1304ms | http |
| 86.104.72.219:1081 | ✓ 1053ms | 否 | 否 | ✓ 1408ms | ✓ 1721ms | http |
| 103.133.223.21:8080 | ✓ 1063ms | ✓ 1964ms | ✓ 1577ms | ✓ 1412ms | ✓ 1386ms | http |
| 62.113.119.14:8080 | ✓ 1941ms | 否 | ✓ 1722ms | ✓ 1587ms | ✓ 1148ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1892ms | ✓ 1250ms | ✓ 977ms | http |
| 217.182.195.221:30003 | ✓ 1262ms | 否 | ✓ 1496ms | ✓ 1752ms | 否 | http |
| 194.59.247.34:10808 | ✓ 1232ms | ✓ 1676ms | ✓ 1303ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 1230ms | 否 | ✓ 886ms | ✓ 1938ms | ✓ 1367ms | http |
| 106.10.55.212:1121 | ✓ 1507ms | 否 | ✓ 1717ms | ✓ 1168ms | ✓ 1051ms | http |
| 161.248.190.99:8080 | ✓ 1858ms | 否 | ✓ 1495ms | ✓ 1804ms | ✓ 1608ms | http |
| 116.171.106.26:3443 | 否 | 否 | ✓ 1488ms | ✓ 1846ms | ✓ 1751ms | http |
| 117.236.124.166:3128 | ✓ 1657ms | 否 | ✓ 1191ms | 否 | ✓ 1411ms | http |
| 103.157.200.126:3128 | ✓ 1399ms | 否 | ✓ 1331ms | 否 | ✓ 1563ms | http |
| 147.45.178.211:14658 | ✓ 970ms | ✓ 1771ms | ✓ 1520ms | 否 | ✓ 1561ms | http |
| 8.219.97.248:80 | ✓ 1580ms | 否 | ✓ 1321ms | ✓ 1896ms | ✓ 1660ms | http |
| 116.171.106.111:3443 | ✓ 1486ms | ✓ 1620ms | 否 | 否 | ✓ 1936ms | http |
| 158.160.215.167:8126 | ✓ 1552ms | ✓ 1908ms | ✓ 1785ms | ✓ 1982ms | ✓ 1689ms | http |
| 168.222.254.136:8888 | ✓ 1545ms | ✓ 1922ms | 否 | 否 | ✓ 1849ms | http |
| 141.11.93.27:8080 | ✓ 1325ms | 否 | ✓ 290ms | 否 | ✓ 1965ms | http |
| 91.217.81.131:1080 | ✓ 1374ms | 否 | ✓ 1052ms | 否 | ✓ 1578ms | http |
| 91.242.229.129:8092 | ✓ 961ms | 否 | ✓ 1708ms | 否 | ✓ 1746ms | http |
| 45.125.67.37:8443 | ✓ 724ms | 否 | ✓ 956ms | ✓ 1349ms | ✓ 1167ms | http |
| 84.47.150.125:1080 | ✓ 740ms | 否 | ✓ 1583ms | 否 | ✓ 1818ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1502ms | ✓ 1135ms | ✓ 1485ms | 否 | http |
| 202.129.206.239:3128 | ✓ 1262ms | 否 | 否 | ✓ 1732ms | ✓ 1513ms | http |
| 200.125.171.254:999 | ✓ 1091ms | ✓ 1754ms | ✓ 1505ms | ✓ 1697ms | ✓ 1516ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1410ms | ✓ 1576ms | ✓ 1003ms | http |
| 34.96.238.40:8080 | ✓ 1238ms | ✓ 1191ms | ✓ 1110ms | ✓ 1029ms | 否 | http |
| 42.200.76.16:3888 | ✓ 801ms | 否 | ✓ 769ms | ✓ 971ms | ✓ 769ms | http |
| 86.104.74.110:1081 | 否 | 否 | ✓ 925ms | ✓ 1741ms | ✓ 1044ms | http |
| 158.160.215.167:8124 | ✓ 1497ms | 否 | ✓ 1753ms | 否 | ✓ 1925ms | http |
| 86.104.74.110:1082 | ✓ 717ms | ✓ 1257ms | ✓ 556ms | ✓ 1796ms | ✓ 1542ms | http |
| 157.0.142.246:10057 | ✓ 1106ms | ✓ 1338ms | ✓ 1111ms | 否 | 否 | http |
| 222.107.27.7:8041 | 否 | 否 | ✓ 739ms | ✓ 1235ms | ✓ 824ms | http |
| 183.98.143.134:8078 | ✓ 1613ms | 否 | ✓ 698ms | 否 | ✓ 789ms | http |
| 178.63.155.151:8888 | ✓ 1433ms | ✓ 1735ms | ✓ 1774ms | 否 | ✓ 1649ms | http |
| 183.98.143.134:8037 | ✓ 1615ms | ✓ 1161ms | ✓ 793ms | 否 | 否 | http |
| 183.98.143.134:8042 | ✓ 1620ms | ✓ 1152ms | 否 | ✓ 1075ms | 否 | http |
| 103.133.254.4:3128 | 否 | 否 | ✓ 1511ms | ✓ 1923ms | ✓ 1906ms | http |
| 101.32.244.83:8080 | ✓ 990ms | 否 | ✓ 967ms | ✓ 1305ms | ✓ 1257ms | http |
| 121.43.196.210:8222 | ✓ 915ms | ✓ 1101ms | ✓ 849ms | ✓ 1166ms | ✓ 964ms | http |
| 121.43.196.213:8222 | ✓ 934ms | ✓ 1103ms | ✓ 864ms | ✓ 1148ms | ✓ 955ms | http |
| 114.231.72.27:1080 | ✓ 1199ms | ✓ 1305ms | ✓ 956ms | ✓ 1578ms | ✓ 1432ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1669ms | 否 | ✓ 1557ms | ✓ 1210ms | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 1134ms | ✓ 1121ms | ✓ 1074ms | http |
| 206.206.126.177:2412 | ✓ 1540ms | 否 | ✓ 1871ms | 否 | ✓ 1421ms | http |
| 147.45.186.28:3128 | ✓ 1314ms | 否 | ✓ 1369ms | 否 | ✓ 1666ms | http |
| 103.209.36.58:8080 | ✓ 1719ms | 否 | 否 | ✓ 1577ms | ✓ 1735ms | http |
| 77.110.119.136:3128 | 否 | 否 | ✓ 339ms | ✓ 1585ms | ✓ 928ms | http |
| 210.77.19.6:7890 | ✓ 913ms | ✓ 1217ms | ✓ 919ms | ✓ 1204ms | ✓ 1096ms | http |
| 101.6.65.112:10080 | ✓ 1065ms | ✓ 1365ms | ✓ 1399ms | ✓ 1329ms | ✓ 1167ms | http |
| 3.101.133.120:80 | ✓ 278ms | ✓ 1310ms | ✓ 493ms | ✓ 1238ms | ✓ 1147ms | http |
| 150.249.255.91:3128 | ✓ 1513ms | ✓ 942ms | 否 | ✓ 911ms | 否 | http |
| 107.173.42.121:7890 | ✓ 1890ms | 否 | ✓ 438ms | ✓ 1338ms | 否 | http |
| 168.194.0.249:252 | ✓ 1459ms | 否 | ✓ 1764ms | ✓ 1804ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1120ms | 否 | ✓ 1673ms | ✓ 1264ms | 否 | http |
| 138.124.99.216:8888 | ✓ 758ms | ✓ 1900ms | ✓ 1507ms | 否 | 否 | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1336ms | ✓ 1192ms | ✓ 1140ms | http |
| 144.124.227.88:3128 | ✓ 1297ms | 否 | ✓ 1756ms | ✓ 1962ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1335ms | ✓ 1270ms | 否 | ✓ 1245ms | http |
| 103.156.17.137:8080 | ✓ 1828ms | 否 | 否 | ✓ 1696ms | ✓ 1532ms | http |
| 61.52.131.172:8443 | ✓ 949ms | ✓ 1213ms | ✓ 922ms | ✓ 1878ms | ✓ 999ms | http |
| 195.19.217.200:3128 | ✓ 1330ms | 否 | ✓ 1924ms | 否 | ✓ 1790ms | http |
| 117.122.240.82:3338 | ✓ 921ms | ✓ 1183ms | ✓ 1726ms | ✓ 1205ms | ✓ 944ms | http |
| 1.180.87.146:22300 | ✓ 1093ms | ✓ 1464ms | ✓ 1143ms | ✓ 1679ms | ✓ 1099ms | http |

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
