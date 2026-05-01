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

最后更新：2026-05-01 09:53:21 UTC（2026-05-01 17:53:21 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 825ms | ✓ 1060ms | ✓ 393ms | ✓ 997ms | ✓ 1103ms | http |
| 103.35.191.138:1082 | 否 | ✓ 1697ms | ✓ 746ms | ✓ 1153ms | ✓ 1099ms | http |
| 218.108.131.186:17890 | ✓ 911ms | ✓ 1173ms | ✓ 983ms | ✓ 1205ms | ✓ 961ms | http |
| 46.101.95.183:8888 | ✓ 1193ms | 否 | ✓ 686ms | 否 | ✓ 1271ms | http |
| 1.231.81.166:3128 | ✓ 1751ms | ✓ 1427ms | 否 | ✓ 1311ms | ✓ 948ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1443ms | ✓ 1574ms | ✓ 1597ms | ✓ 1348ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1720ms | ✓ 1523ms | ✓ 1490ms | ✓ 1089ms | http |
| 45.167.124.71:999 | ✓ 912ms | ✓ 1929ms | ✓ 1404ms | ✓ 1765ms | ✓ 1550ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1760ms | ✓ 1930ms | 否 | ✓ 1237ms | http |
| 103.35.190.182:1082 | ✓ 543ms | ✓ 1003ms | ✓ 184ms | 否 | 否 | http |
| 77.110.116.224:3128 | ✓ 629ms | 否 | ✓ 1472ms | 否 | ✓ 1528ms | http |
| 103.35.191.244:1082 | ✓ 512ms | ✓ 986ms | ✓ 873ms | 否 | ✓ 809ms | http |
| 154.64.232.35:8080 | ✓ 1454ms | 否 | ✓ 1932ms | ✓ 1914ms | ✓ 1724ms | http |
| 45.140.147.82:1081 | ✓ 1638ms | ✓ 1594ms | 否 | ✓ 1863ms | ✓ 1418ms | http |
| 152.70.91.193:40000 | ✓ 1391ms | 否 | 否 | ✓ 1365ms | ✓ 1423ms | http |
| 94.131.118.129:1081 | ✓ 1765ms | 否 | ✓ 1537ms | 否 | ✓ 1567ms | http |
| 103.35.190.69:1082 | ✓ 1424ms | 否 | 否 | ✓ 1198ms | ✓ 773ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1029ms | 否 | ✓ 1735ms | ✓ 797ms | http |
| 212.58.132.5:8888 | ✓ 1979ms | 否 | ✓ 1746ms | ✓ 1596ms | ✓ 1275ms | http |
| 159.223.225.118:8888 | ✓ 1553ms | 否 | ✓ 1867ms | ✓ 1813ms | ✓ 1468ms | http |
| 120.92.212.16:8890 | ✓ 1441ms | 否 | ✓ 1078ms | ✓ 1670ms | ✓ 1336ms | http |
| 138.124.99.216:8888 | ✓ 1112ms | 否 | ✓ 850ms | 否 | ✓ 1738ms | http |
| 34.96.238.40:8080 | ✓ 1360ms | ✓ 1195ms | 否 | 否 | ✓ 1239ms | http |
| 43.167.237.94:3128 | 否 | 否 | ✓ 616ms | ✓ 1795ms | ✓ 726ms | http |
| 150.107.140.238:3128 | ✓ 1114ms | 否 | ✓ 1063ms | ✓ 1507ms | ✓ 1408ms | http |
| 65.109.213.99:1080 | ✓ 1243ms | 否 | ✓ 1599ms | ✓ 1947ms | ✓ 1954ms | http |
| 103.35.191.174:1082 | ✓ 1759ms | 否 | ✓ 156ms | ✓ 1243ms | ✓ 783ms | http |
| 59.46.216.131:30001 | ✓ 1117ms | 否 | ✓ 1176ms | 否 | ✓ 1182ms | http |
| 103.35.191.138:1081 | ✓ 1687ms | ✓ 1100ms | ✓ 127ms | ✓ 1050ms | ✓ 832ms | http |
| 62.60.237.68:8080 | ✓ 1471ms | ✓ 1605ms | 否 | ✓ 1815ms | ✓ 1114ms | http |
| 86.104.74.110:1082 | 否 | ✓ 1209ms | ✓ 930ms | ✓ 1823ms | ✓ 1029ms | http |
| 86.104.74.110:1081 | 否 | 否 | ✓ 646ms | ✓ 1334ms | ✓ 995ms | http |
| 120.92.108.86:7890 | ✓ 1345ms | 否 | ✓ 1629ms | 否 | ✓ 1538ms | http |
| 62.60.231.71:56608 | ✓ 1084ms | 否 | ✓ 918ms | ✓ 1964ms | ✓ 1112ms | http |
| 77.110.107.80:8080 | ✓ 1036ms | 否 | ✓ 1729ms | ✓ 1728ms | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1920ms | ✓ 1087ms | 否 | ✓ 1852ms | http |
| 103.70.114.149:3128 | ✓ 1808ms | 否 | ✓ 1723ms | 否 | ✓ 1751ms | http |
| 103.35.190.182:1081 | ✓ 752ms | ✓ 1049ms | ✓ 832ms | ✓ 1092ms | ✓ 1038ms | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 818ms | ✓ 1114ms | ✓ 911ms | http |
| 103.157.200.126:3128 | ✓ 1065ms | 否 | ✓ 1111ms | ✓ 1968ms | ✓ 1342ms | http |
| 89.208.106.138:10808 | ✓ 1094ms | ✓ 1725ms | ✓ 608ms | 否 | 否 | http |
| 130.61.174.200:1080 | ✓ 489ms | 否 | ✓ 501ms | ✓ 1350ms | 否 | http |
| 138.68.153.144:3128 | ✓ 1047ms | 否 | ✓ 1195ms | ✓ 1898ms | 否 | http |
| 45.153.231.229:8080 | ✓ 1955ms | 否 | ✓ 1782ms | ✓ 1973ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1750ms | ✓ 1649ms | ✓ 1600ms | ✓ 1586ms | 否 | http |
| 43.133.44.89:8888 | ✓ 818ms | 否 | ✓ 1066ms | 否 | ✓ 881ms | http |
| 103.35.190.69:1081 | ✓ 306ms | 否 | ✓ 130ms | 否 | ✓ 975ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1892ms | ✓ 1321ms | ✓ 1270ms | http |
| 8.219.97.248:80 | ✓ 1628ms | 否 | ✓ 1572ms | 否 | ✓ 1649ms | http |
| 34.101.184.164:3128 | ✓ 1602ms | 否 | ✓ 1479ms | 否 | ✓ 1213ms | http |
| 80.92.204.47:1081 | ✓ 492ms | ✓ 1349ms | ✓ 1494ms | ✓ 1654ms | ✓ 1423ms | http |
| 120.92.212.16:7890 | ✓ 1861ms | 否 | 否 | ✓ 1268ms | ✓ 1021ms | http |
| 217.77.102.18:3128 | ✓ 1195ms | 否 | ✓ 1378ms | 否 | ✓ 1798ms | http |
| 104.248.195.47:8080 | ✓ 550ms | ✓ 1970ms | 否 | ✓ 1734ms | ✓ 1602ms | http |
| 121.230.8.55:1080 | ✓ 1111ms | ✓ 1557ms | ✓ 1237ms | 否 | ✓ 1219ms | http |
| 103.35.191.244:1081 | ✓ 1287ms | ✓ 1524ms | ✓ 884ms | ✓ 1109ms | ✓ 805ms | http |
| 72.11.151.159:6005 | ✓ 703ms | 否 | ✓ 1155ms | ✓ 1538ms | 否 | http |
| 92.119.127.211:6005 | ✓ 1450ms | ✓ 1992ms | ✓ 1462ms | ✓ 1764ms | ✓ 1487ms | http |
| 183.238.3.150:7897 | ✓ 1023ms | ✓ 1277ms | ✓ 1066ms | ✓ 1318ms | ✓ 947ms | http |
| 218.72.124.35:40000 | ✓ 992ms | 否 | ✓ 1047ms | ✓ 1305ms | ✓ 984ms | http |
| 92.119.127.213:6005 | ✓ 1327ms | 否 | ✓ 1752ms | ✓ 1848ms | ✓ 1279ms | http |
| 72.11.150.178:6005 | ✓ 407ms | ✓ 1041ms | ✓ 827ms | ✓ 1282ms | ✓ 1199ms | http |
| 185.21.11.140:1080 | ✓ 1285ms | 否 | ✓ 931ms | 否 | ✓ 1747ms | http |
| 86.104.72.220:1081 | ✓ 1329ms | 否 | ✓ 327ms | ✓ 1249ms | ✓ 1135ms | http |
| 171.5.188.106:8080 | ✓ 1516ms | 否 | ✓ 1648ms | ✓ 1765ms | ✓ 1656ms | http |
| 57.128.188.167:9147 | ✓ 1738ms | 否 | ✓ 1917ms | 否 | ✓ 1684ms | http |
| 86.104.72.219:1081 | ✓ 494ms | 否 | ✓ 1551ms | ✓ 1518ms | ✓ 1822ms | http |
| 128.199.121.61:9090 | ✓ 1869ms | 否 | ✓ 1732ms | ✓ 1606ms | ✓ 1939ms | http |
| 45.186.6.104:3128 | ✓ 1887ms | ✓ 1908ms | ✓ 1990ms | 否 | 否 | http |
| 61.52.131.172:8443 | 否 | ✓ 1939ms | ✓ 990ms | ✓ 1244ms | 否 | http |
| 116.171.106.26:3443 | 否 | 否 | ✓ 1556ms | ✓ 1887ms | ✓ 1589ms | http |
| 20.164.75.153:8080 | ✓ 1884ms | 否 | ✓ 1779ms | 否 | ✓ 1803ms | http |
| 220.197.44.36:3128 | ✓ 1590ms | ✓ 1796ms | ✓ 1959ms | ✓ 1655ms | 否 | http |
| 3.101.133.120:80 | 否 | 否 | ✓ 519ms | ✓ 1757ms | ✓ 1084ms | http |
| 103.39.51.207:8080 | ✓ 1774ms | 否 | ✓ 1353ms | ✓ 1483ms | ✓ 1418ms | http |

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
