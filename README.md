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

最后更新：2026-03-24 23:28:30 UTC（2026-03-25 07:28:30 UTC+8）

**代理总数：285**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 284 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 285 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.208.213:8450 | ✓ 337ms | ✓ 705ms | ✓ 1800ms | ✓ 1288ms | ✓ 955ms | http |
| 147.161.210.140:8800 | ✓ 1411ms | ✓ 1926ms | ✓ 742ms | ✓ 1161ms | ✓ 1141ms | http |
| 167.103.115.102:8800 | ✓ 1366ms | 否 | ✓ 928ms | ✓ 1087ms | ✓ 1048ms | http |
| 167.103.34.108:8800 | ✓ 1419ms | 否 | ✓ 1450ms | ✓ 1425ms | ✓ 1489ms | http |
| 38.34.179.174:8453 | ✓ 628ms | ✓ 605ms | ✓ 647ms | ✓ 908ms | ✓ 1404ms | http |
| 38.34.179.173:8452 | ✓ 824ms | ✓ 629ms | ✓ 469ms | ✓ 1015ms | ✓ 1317ms | http |
| 38.145.220.198:8448 | ✓ 868ms | ✓ 807ms | ✓ 156ms | ✓ 731ms | ✓ 762ms | http |
| 45.136.131.42:8447 | ✓ 626ms | ✓ 610ms | ✓ 282ms | ✓ 1043ms | ✓ 1740ms | http |
| 38.34.179.161:8448 | ✓ 799ms | ✓ 612ms | ✓ 756ms | ✓ 969ms | ✓ 1182ms | http |
| 35.225.22.61:80 | ✓ 650ms | ✓ 1303ms | 否 | ✓ 1141ms | ✓ 1170ms | http |
| 154.12.59.102:6005 | ✓ 880ms | ✓ 688ms | 否 | ✓ 1058ms | ✓ 749ms | http |
| 43.99.54.236:5555 | ✓ 1033ms | ✓ 890ms | ✓ 621ms | ✓ 793ms | ✓ 653ms | http |
| 38.145.218.228:8447 | ✓ 1038ms | ✓ 634ms | ✓ 1197ms | 否 | ✓ 615ms | http |
| 45.136.131.36:8450 | ✓ 615ms | ✓ 601ms | ✓ 339ms | ✓ 1238ms | 否 | http |
| 38.145.218.229:8450 | ✓ 860ms | ✓ 1131ms | ✓ 209ms | ✓ 700ms | ✓ 1024ms | http |
| 38.34.179.57:8453 | ✓ 1077ms | ✓ 877ms | ✓ 805ms | ✓ 1267ms | ✓ 1709ms | http |
| 38.34.179.51:8449 | ✓ 1236ms | ✓ 704ms | ✓ 770ms | ✓ 1106ms | ✓ 1839ms | http |
| 38.145.208.209:8444 | ✓ 784ms | ✓ 697ms | ✓ 1060ms | ✓ 1769ms | ✓ 503ms | http |
| 38.145.220.33:8448 | ✓ 1148ms | ✓ 788ms | ✓ 1280ms | 否 | ✓ 534ms | http |
| 38.34.179.186:8444 | ✓ 1734ms | ✓ 1884ms | ✓ 707ms | ✓ 1014ms | ✓ 1230ms | http |
| 38.34.179.14:8450 | ✓ 869ms | 否 | ✓ 229ms | ✓ 840ms | ✓ 959ms | http |
| 38.34.179.20:8445 | ✓ 793ms | ✓ 1961ms | ✓ 272ms | ✓ 818ms | ✓ 697ms | http |
| 38.34.179.40:8446 | ✓ 858ms | ✓ 1214ms | 否 | ✓ 664ms | ✓ 606ms | http |
| 38.34.183.234:8450 | ✓ 880ms | ✓ 1098ms | ✓ 1333ms | ✓ 881ms | ✓ 816ms | http |
| 38.34.183.211:8445 | ✓ 820ms | 否 | ✓ 417ms | ✓ 883ms | ✓ 1017ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1429ms | ✓ 1233ms | ✓ 1211ms | ✓ 961ms | http |
| 38.34.179.39:8452 | ✓ 835ms | ✓ 1203ms | 否 | ✓ 713ms | ✓ 751ms | http |
| 45.136.131.39:8443 | ✓ 636ms | ✓ 627ms | ✓ 325ms | ✓ 1282ms | ✓ 1967ms | http |
| 155.212.132.241:3128 | ✓ 726ms | ✓ 1934ms | ✓ 806ms | ✓ 1951ms | ✓ 1433ms | http |
| 120.92.212.16:7890 | ✓ 1953ms | ✓ 1398ms | ✓ 1123ms | ✓ 1639ms | ✓ 1378ms | http |
| 113.160.132.26:8080 | ✓ 920ms | ✓ 1395ms | ✓ 1093ms | ✓ 1278ms | 否 | http |
| 38.34.183.47:8452 | ✓ 1801ms | ✓ 812ms | ✓ 384ms | ✓ 1769ms | ✓ 1679ms | http |
| 150.241.71.15:1080 | ✓ 1168ms | 否 | ✓ 1785ms | 否 | ✓ 1473ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1071ms | ✓ 1864ms | ✓ 970ms | http |
| 45.136.130.168:8452 | ✓ 1704ms | ✓ 1709ms | ✓ 1677ms | 否 | ✓ 873ms | http |
| 167.103.31.122:8800 | ✓ 1522ms | 否 | ✓ 1725ms | 否 | ✓ 1929ms | http |
| 103.84.95.54:7890 | ✓ 1068ms | 否 | ✓ 621ms | ✓ 1374ms | ✓ 692ms | http |
| 222.228.171.92:8080 | 否 | ✓ 1751ms | ✓ 659ms | 否 | ✓ 1863ms | http |
| 223.16.170.103:80 | ✓ 1909ms | 否 | ✓ 1213ms | ✓ 1022ms | ✓ 1037ms | http |
| 101.43.127.100:8877 | ✓ 785ms | ✓ 1091ms | ✓ 903ms | ✓ 1051ms | ✓ 878ms | http |
| 34.101.184.164:3128 | ✓ 1540ms | 否 | ✓ 1079ms | ✓ 1270ms | ✓ 1013ms | http |
| 160.250.4.13:1 | ✓ 1534ms | 否 | ✓ 1750ms | ✓ 1529ms | ✓ 1045ms | http |
| 137.220.151.110:6005 | ✓ 1775ms | ✓ 1894ms | ✓ 1849ms | 否 | 否 | http |
| 137.220.150.104:6005 | ✓ 1791ms | ✓ 1949ms | ✓ 809ms | ✓ 1096ms | ✓ 1058ms | http |
| 218.89.134.230:3333 | 否 | ✓ 1552ms | 否 | ✓ 1556ms | ✓ 1236ms | http |
| 150.107.140.238:3128 | ✓ 747ms | 否 | 否 | ✓ 1482ms | ✓ 892ms | http |
| 194.67.99.223:1080 | ✓ 751ms | ✓ 1866ms | ✓ 1298ms | ✓ 1892ms | 否 | http |
| 181.41.201.85:3128 | ✓ 908ms | 否 | ✓ 964ms | 否 | ✓ 1735ms | http |
| 38.145.208.221:8453 | ✓ 146ms | ✓ 596ms | ✓ 91ms | ✓ 683ms | ✓ 530ms | http |
| 38.145.203.86:8449 | ✓ 322ms | ✓ 591ms | ✓ 744ms | ✓ 682ms | ✓ 715ms | http |
| 147.161.239.240:8800 | ✓ 990ms | ✓ 1748ms | ✓ 1333ms | ✓ 1850ms | ✓ 1704ms | http |
| 193.233.22.29:10808 | ✓ 867ms | 否 | ✓ 1813ms | ✓ 1674ms | ✓ 1245ms | http |
| 47.77.193.180:1080 | ✓ 652ms | ✓ 720ms | ✓ 379ms | ✓ 740ms | ✓ 574ms | http |
| 142.171.224.229:7890 | ✓ 822ms | ✓ 673ms | ✓ 972ms | ✓ 736ms | ✓ 598ms | http |
| 38.145.220.102:8453 | ✓ 1132ms | ✓ 635ms | ✓ 443ms | ✓ 868ms | ✓ 1296ms | http |
| 45.136.130.177:8448 | ✓ 1766ms | ✓ 1479ms | ✓ 417ms | ✓ 1369ms | ✓ 662ms | http |
| 150.249.255.91:3128 | 否 | ✓ 785ms | ✓ 612ms | ✓ 941ms | ✓ 698ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1305ms | ✓ 1146ms | ✓ 981ms | ✓ 862ms | http |
| 160.238.65.4:3128 | ✓ 655ms | ✓ 1522ms | ✓ 1443ms | 否 | ✓ 1590ms | http |
| 160.238.65.3:3128 | ✓ 658ms | ✓ 1539ms | ✓ 1407ms | 否 | ✓ 1618ms | http |
| 160.238.65.2:3128 | ✓ 657ms | ✓ 1407ms | ✓ 1537ms | 否 | ✓ 1620ms | http |
| 198.244.188.127:3128 | ✓ 989ms | ✓ 1862ms | ✓ 1593ms | 否 | ✓ 1827ms | http |
| 38.145.203.110:8448 | ✓ 776ms | ✓ 1473ms | 否 | ✓ 692ms | ✓ 492ms | http |
| 195.123.213.129:1080 | ✓ 1623ms | 否 | ✓ 1639ms | 否 | ✓ 1336ms | http |
| 45.205.28.107:8080 | ✓ 349ms | ✓ 651ms | ✓ 815ms | 否 | 否 | http |
| 38.34.179.150:8449 | 否 | ✓ 732ms | ✓ 455ms | ✓ 1216ms | ✓ 554ms | http |
| 45.136.130.178:8449 | 否 | ✓ 1693ms | ✓ 208ms | ✓ 873ms | ✓ 1611ms | http |
| 137.220.150.22:6005 | 否 | ✓ 1964ms | ✓ 1168ms | ✓ 1501ms | ✓ 971ms | http |
| 38.145.208.241:8453 | 否 | ✓ 1150ms | ✓ 1886ms | ✓ 1453ms | ✓ 1314ms | http |
| 121.126.185.63:25152 | 否 | ✓ 1680ms | ✓ 1978ms | 否 | ✓ 1495ms | http |
| 137.220.150.152:6005 | ✓ 1884ms | 否 | ✓ 1569ms | ✓ 1946ms | ✓ 1254ms | http |
| 115.231.181.40:8128 | ✓ 882ms | 否 | ✓ 910ms | ✓ 1098ms | 否 | http |
| 38.34.178.186:8451 | 否 | ✓ 802ms | ✓ 148ms | ✓ 822ms | ✓ 922ms | http |
| 106.75.15.167:7890 | ✓ 1287ms | ✓ 1697ms | ✓ 1669ms | 否 | ✓ 902ms | http |
| 38.34.183.224:8448 | ✓ 573ms | ✓ 1341ms | ✓ 123ms | ✓ 906ms | ✓ 540ms | http |
| 45.136.130.191:8453 | ✓ 143ms | ✓ 748ms | ✓ 884ms | ✓ 1431ms | ✓ 547ms | http |
| 45.136.131.58:8449 | ✓ 578ms | ✓ 1476ms | ✓ 186ms | ✓ 710ms | ✓ 1058ms | http |
| 45.8.157.38:3128 | 否 | 否 | ✓ 1045ms | ✓ 1421ms | ✓ 1672ms | http |
| 38.34.183.13:8449 | ✓ 187ms | ✓ 635ms | ✓ 530ms | ✓ 1460ms | ✓ 1747ms | http |
| 36.92.51.132:8080 | ✓ 1829ms | 否 | ✓ 1476ms | 否 | ✓ 1412ms | http |
| 160.238.65.7:3128 | ✓ 1459ms | ✓ 1644ms | 否 | ✓ 1479ms | ✓ 1531ms | http |
| 160.238.65.9:3128 | ✓ 1457ms | ✓ 1550ms | ✓ 1376ms | 否 | ✓ 1714ms | http |
| 160.238.65.5:3128 | ✓ 1457ms | ✓ 1484ms | ✓ 1443ms | 否 | ✓ 1720ms | http |
| 160.238.65.6:3128 | ✓ 1454ms | ✓ 1499ms | ✓ 1430ms | 否 | ✓ 1705ms | http |
| 160.238.65.8:3128 | ✓ 1458ms | ✓ 1475ms | ✓ 1452ms | 否 | ✓ 1718ms | http |
| 88.80.150.82:8080 | ✓ 1438ms | ✓ 1899ms | 否 | 否 | ✓ 1980ms | https |
| 103.82.23.118:5253 | 否 | 否 | ✓ 1262ms | ✓ 1682ms | ✓ 1698ms | http |
| 37.187.109.70:10111 | ✓ 1219ms | ✓ 1756ms | ✓ 1483ms | ✓ 1907ms | 否 | http |
| 38.34.179.23:8448 | ✓ 309ms | ✓ 870ms | ✓ 766ms | 否 | 否 | http |
| 137.220.150.170:6005 | ✓ 1620ms | ✓ 1852ms | ✓ 1043ms | ✓ 1269ms | ✓ 1767ms | http |
| 38.145.208.176:8449 | ✓ 506ms | ✓ 767ms | ✓ 588ms | ✓ 737ms | ✓ 493ms | http |
| 45.136.130.170:8448 | ✓ 503ms | ✓ 843ms | ✓ 967ms | 否 | ✓ 1010ms | http |
| 45.136.130.168:8448 | ✓ 502ms | ✓ 845ms | ✓ 1120ms | 否 | ✓ 1094ms | http |
| 45.136.130.173:8448 | ✓ 503ms | ✓ 847ms | ✓ 1119ms | 否 | ✓ 1122ms | http |
| 38.145.208.185:8452 | ✓ 122ms | ✓ 614ms | ✓ 299ms | ✓ 716ms | ✓ 540ms | http |
| 38.145.208.180:8445 | ✓ 129ms | ✓ 620ms | ✓ 286ms | ✓ 652ms | ✓ 626ms | http |
| 38.145.218.217:8452 | ✓ 179ms | ✓ 609ms | ✓ 247ms | ✓ 672ms | ✓ 515ms | http |
| 38.145.203.107:8453 | ✓ 344ms | ✓ 771ms | ✓ 128ms | ✓ 746ms | ✓ 544ms | http |
| 38.145.218.218:8446 | ✓ 164ms | ✓ 613ms | ✓ 261ms | ✓ 671ms | ✓ 541ms | http |
| 38.145.220.65:8450 | ✓ 132ms | ✓ 737ms | ✓ 463ms | ✓ 692ms | ✓ 554ms | http |

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
