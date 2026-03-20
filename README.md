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

最后更新：2026-03-20 05:41:50 UTC（2026-03-20 13:41:50 UTC+8）

**代理总数：148**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 148 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 148 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 154.64.243.50:7890 | ✓ 1101ms | ✓ 1563ms | ✓ 95ms | ✓ 680ms | ✓ 656ms | http |
| 103.84.95.54:7890 | ✓ 837ms | 否 | 否 | ✓ 1467ms | ✓ 637ms | http |
| 147.161.210.140:8800 | ✓ 1476ms | 否 | ✓ 820ms | ✓ 999ms | ✓ 1060ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1580ms | ✓ 917ms | ✓ 1454ms | ✓ 950ms | http |
| 38.34.179.96:8451 | ✓ 179ms | ✓ 789ms | ✓ 1386ms | 否 | ✓ 847ms | http |
| 219.117.204.211:7799 | ✓ 1479ms | 否 | ✓ 818ms | 否 | ✓ 817ms | http |
| 38.34.179.16:8451 | ✓ 1293ms | ✓ 1908ms | ✓ 1360ms | ✓ 1215ms | ✓ 523ms | http |
| 147.161.239.240:8800 | ✓ 1408ms | ✓ 1847ms | ✓ 1169ms | 否 | ✓ 1513ms | http |
| 133.242.138.34:8100 | ✓ 1498ms | 否 | 否 | ✓ 1612ms | ✓ 1074ms | http |
| 174.138.24.77:1080 | 否 | 否 | ✓ 1897ms | ✓ 1108ms | ✓ 847ms | http |
| 167.103.34.108:8800 | ✓ 1273ms | 否 | ✓ 1322ms | ✓ 1996ms | ✓ 1264ms | http |
| 38.55.104.8:6005 | ✓ 1806ms | 否 | ✓ 1413ms | 否 | ✓ 1689ms | http |
| 194.5.212.40:8080 | ✓ 1458ms | 否 | ✓ 1691ms | 否 | ✓ 1685ms | http |
| 38.55.104.99:6005 | ✓ 1912ms | 否 | ✓ 1405ms | 否 | ✓ 1977ms | http |
| 1.231.81.166:3128 | ✓ 1607ms | ✓ 1020ms | ✓ 1389ms | ✓ 1237ms | ✓ 976ms | http |
| 45.167.124.52:8080 | ✓ 1060ms | 否 | ✓ 1540ms | 否 | ✓ 1845ms | http |
| 142.171.224.229:7890 | ✓ 525ms | ✓ 1391ms | ✓ 1125ms | ✓ 758ms | ✓ 641ms | http |
| 35.225.22.61:80 | ✓ 624ms | ✓ 1473ms | 否 | ✓ 1377ms | 否 | http |
| 103.113.70.189:1081 | ✓ 443ms | ✓ 1164ms | ✓ 1392ms | ✓ 1174ms | ✓ 879ms | http |
| 38.34.183.224:8448 | 否 | ✓ 1098ms | ✓ 170ms | ✓ 797ms | ✓ 1010ms | http |
| 45.125.67.37:443 | ✓ 1885ms | 否 | ✓ 1050ms | ✓ 1104ms | ✓ 1070ms | http |
| 137.220.150.22:6005 | 否 | 否 | ✓ 712ms | ✓ 1316ms | ✓ 849ms | http |
| 38.55.104.68:6005 | ✓ 1704ms | 否 | ✓ 1392ms | ✓ 1699ms | ✓ 1211ms | http |
| 167.103.31.122:8800 | ✓ 1982ms | 否 | ✓ 1273ms | 否 | ✓ 1924ms | http |
| 167.71.60.190:8080 | ✓ 838ms | 否 | ✓ 1988ms | 否 | ✓ 1439ms | http |
| 194.59.204.87:9080 | ✓ 1400ms | ✓ 1605ms | ✓ 698ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 893ms | ✓ 1897ms | ✓ 1133ms | ✓ 1285ms | ✓ 1558ms | http |
| 45.119.85.216:3128 | ✓ 826ms | 否 | ✓ 854ms | 否 | ✓ 1441ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1950ms | ✓ 1807ms | ✓ 1010ms | http |
| 103.3.246.71:3128 | ✓ 1673ms | 否 | ✓ 1650ms | ✓ 1761ms | ✓ 1733ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1194ms | 否 | ✓ 1428ms | ✓ 956ms | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 1752ms | ✓ 1867ms | ✓ 1640ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 760ms | ✓ 1092ms | ✓ 880ms | http |
| 38.34.179.105:8449 | ✓ 1114ms | ✓ 738ms | ✓ 135ms | ✓ 849ms | 否 | http |
| 162.240.154.26:3128 | ✓ 615ms | ✓ 1548ms | ✓ 868ms | 否 | 否 | http |
| 24.144.86.173:1080 | ✓ 329ms | ✓ 1968ms | ✓ 1037ms | ✓ 1940ms | ✓ 629ms | http |
| 38.145.218.228:8447 | ✓ 773ms | ✓ 1695ms | ✓ 91ms | ✓ 658ms | 否 | http |
| 45.136.130.168:8448 | ✓ 224ms | ✓ 814ms | ✓ 462ms | ✓ 1467ms | ✓ 1695ms | http |
| 38.34.179.63:8448 | 否 | ✓ 1642ms | ✓ 83ms | ✓ 693ms | ✓ 761ms | http |
| 45.136.130.173:8448 | ✓ 220ms | ✓ 1700ms | ✓ 1305ms | ✓ 1603ms | ✓ 528ms | http |
| 38.145.220.32:8444 | ✓ 1569ms | 否 | ✓ 393ms | ✓ 1342ms | ✓ 1393ms | http |
| 38.34.179.83:8448 | ✓ 155ms | ✓ 1676ms | ✓ 1960ms | ✓ 676ms | ✓ 589ms | http |
| 38.145.208.244:8448 | ✓ 355ms | ✓ 1697ms | ✓ 1114ms | ✓ 1328ms | ✓ 684ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1230ms | 否 | ✓ 1332ms | ✓ 1311ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1261ms | 否 | ✓ 1173ms | ✓ 781ms | http |
| 91.238.104.171:2023 | ✓ 1282ms | 否 | ✓ 1759ms | 否 | ✓ 1970ms | http |
| 91.238.105.64:2024 | ✓ 1286ms | 否 | ✓ 1751ms | 否 | ✓ 1987ms | http |
| 38.145.208.168:8452 | ✓ 445ms | ✓ 798ms | ✓ 115ms | ✓ 730ms | ✓ 564ms | http |
| 38.145.208.175:8452 | ✓ 809ms | ✓ 847ms | ✓ 100ms | ✓ 676ms | ✓ 587ms | http |
| 38.145.220.39:8453 | ✓ 433ms | ✓ 826ms | ✓ 740ms | ✓ 988ms | ✓ 636ms | http |
| 45.136.131.60:8452 | ✓ 996ms | 否 | ✓ 110ms | ✓ 680ms | ✓ 858ms | http |
| 86.53.183.16:1080 | ✓ 639ms | 否 | ✓ 1438ms | 否 | ✓ 1765ms | http |
| 38.145.220.11:8445 | ✓ 1595ms | ✓ 1142ms | ✓ 278ms | ✓ 742ms | ✓ 1726ms | http |
| 120.92.212.16:7890 | ✓ 1175ms | 否 | ✓ 969ms | ✓ 1450ms | 否 | http |
| 102.134.48.240:6005 | ✓ 1130ms | ✓ 1249ms | ✓ 1543ms | ✓ 1503ms | ✓ 752ms | http |
| 8.219.97.248:80 | ✓ 1472ms | 否 | ✓ 1716ms | ✓ 1414ms | 否 | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 926ms | ✓ 1179ms | ✓ 938ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1180ms | ✓ 1141ms | ✓ 932ms | http |
| 103.139.138.194:3128 | ✓ 1715ms | 否 | ✓ 1338ms | ✓ 1543ms | ✓ 1065ms | http |
| 120.132.97.88:7897 | ✓ 980ms | ✓ 1156ms | ✓ 1097ms | ✓ 1141ms | ✓ 873ms | http |
| 152.70.84.108:8080 | ✓ 1528ms | 否 | ✓ 1301ms | ✓ 1271ms | ✓ 1094ms | http |
| 106.75.15.167:7890 | ✓ 870ms | ✓ 1788ms | ✓ 897ms | ✓ 1379ms | ✓ 1187ms | http |
| 47.101.159.19:8899 | 否 | ✓ 1674ms | ✓ 1356ms | ✓ 1911ms | 否 | http |
| 59.46.216.131:30001 | 否 | ✓ 1328ms | ✓ 1055ms | ✓ 1332ms | 否 | http |
| 112.163.160.93:3128 | 否 | 否 | ✓ 840ms | ✓ 940ms | ✓ 749ms | http |
| 8.222.175.80:6128 | ✓ 1070ms | 否 | ✓ 937ms | ✓ 1004ms | ✓ 784ms | http |
| 150.107.140.238:3128 | ✓ 735ms | 否 | ✓ 1789ms | ✓ 1114ms | ✓ 865ms | http |
| 38.145.203.108:8451 | 否 | ✓ 1698ms | ✓ 678ms | ✓ 1288ms | ✓ 1671ms | http |
| 38.34.178.152:8447 | 否 | ✓ 800ms | ✓ 1496ms | ✓ 652ms | ✓ 520ms | http |
| 137.220.150.152:6005 | ✓ 1881ms | 否 | ✓ 937ms | ✓ 1944ms | ✓ 925ms | http |
| 137.220.150.104:6005 | ✓ 1499ms | ✓ 1836ms | ✓ 1203ms | 否 | ✓ 1122ms | http |
| 38.145.220.77:8453 | ✓ 770ms | ✓ 657ms | ✓ 281ms | ✓ 923ms | ✓ 917ms | http |
| 185.41.152.110:3128 | ✓ 1135ms | 否 | ✓ 1453ms | 否 | ✓ 1653ms | http |
| 45.88.0.116:3128 | ✓ 674ms | 否 | ✓ 643ms | ✓ 1479ms | ✓ 1172ms | http |
| 45.88.0.98:3128 | ✓ 670ms | ✓ 1593ms | ✓ 969ms | ✓ 1481ms | ✓ 1245ms | http |
| 45.88.0.115:3128 | ✓ 674ms | 否 | ✓ 643ms | ✓ 1475ms | ✓ 1175ms | http |
| 45.88.0.117:3128 | ✓ 679ms | 否 | ✓ 638ms | ✓ 1442ms | ✓ 1196ms | http |
| 45.88.0.114:3128 | ✓ 680ms | ✓ 1536ms | ✓ 1021ms | ✓ 1498ms | ✓ 1221ms | http |
| 213.220.62.62:3128 | 否 | 否 | ✓ 1909ms | ✓ 1484ms | ✓ 1261ms | http |
| 45.88.0.113:3128 | ✓ 683ms | 否 | ✓ 641ms | ✓ 1483ms | ✓ 1157ms | http |
| 45.88.0.111:3128 | ✓ 684ms | 否 | ✓ 636ms | ✓ 1519ms | ✓ 1175ms | http |
| 45.88.0.99:3128 | ✓ 593ms | ✓ 1490ms | ✓ 636ms | ✓ 1447ms | ✓ 1159ms | http |
| 103.133.254.4:3128 | ✓ 1761ms | 否 | ✓ 1635ms | ✓ 1461ms | ✓ 1153ms | http |
| 101.109.168.105:8080 | ✓ 1742ms | 否 | ✓ 1214ms | 否 | ✓ 1407ms | http |
| 137.184.1.87:3128 | 否 | ✓ 1088ms | ✓ 1033ms | ✓ 645ms | ✓ 488ms | http |
| 109.120.185.119:8118 | ✓ 879ms | 否 | ✓ 1804ms | ✓ 1796ms | ✓ 1586ms | http |
| 45.136.198.40:3128 | ✓ 1132ms | 否 | ✓ 1733ms | 否 | ✓ 1651ms | http |
| 38.145.218.163:8451 | ✓ 1967ms | 否 | ✓ 1858ms | 否 | ✓ 1728ms | http |
| 38.145.208.151:8453 | ✓ 344ms | ✓ 972ms | ✓ 102ms | ✓ 677ms | ✓ 598ms | http |
| 38.34.179.35:8451 | ✓ 215ms | ✓ 897ms | ✓ 88ms | ✓ 670ms | ✓ 755ms | http |
| 38.34.179.38:8448 | ✓ 213ms | ✓ 1411ms | ✓ 172ms | ✓ 686ms | ✓ 525ms | http |
| 24.199.124.152:3128 | ✓ 228ms | ✓ 1409ms | ✓ 883ms | ✓ 651ms | ✓ 486ms | http |
| 38.34.179.77:8453 | ✓ 150ms | ✓ 618ms | ✓ 874ms | ✓ 809ms | ✓ 1277ms | http |
| 45.136.131.27:8449 | ✓ 380ms | ✓ 1258ms | ✓ 425ms | ✓ 713ms | ✓ 869ms | http |
| 38.34.179.63:8451 | ✓ 360ms | ✓ 1131ms | ✓ 510ms | ✓ 714ms | ✓ 502ms | http |
| 38.34.179.64:8450 | ✓ 774ms | 否 | ✓ 92ms | ✓ 683ms | ✓ 511ms | http |
| 38.145.218.230:8453 | ✓ 1480ms | 否 | ✓ 298ms | ✓ 808ms | ✓ 1832ms | http |
| 38.145.203.32:8451 | ✓ 1003ms | 否 | ✓ 126ms | ✓ 706ms | ✓ 768ms | http |
| 38.145.208.172:8448 | ✓ 260ms | ✓ 691ms | ✓ 980ms | ✓ 1533ms | ✓ 509ms | http |
| 122.248.45.54:8080 | ✓ 1857ms | 否 | ✓ 1182ms | 否 | ✓ 1355ms | http |

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
