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

最后更新：2026-03-22 03:29:48 UTC（2026-03-22 11:29:48 UTC+8）

**代理总数：181**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 180 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 181 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 172.93.42.71:3128 | ✓ 981ms | ✓ 889ms | ✓ 1995ms | ✓ 905ms | ✓ 692ms | http |
| 45.136.130.171:8445 | ✓ 980ms | 否 | ✓ 884ms | ✓ 1044ms | ✓ 703ms | http |
| 103.84.95.54:7890 | ✓ 958ms | 否 | ✓ 892ms | ✓ 944ms | 否 | http |
| 147.161.210.140:8800 | ✓ 1766ms | ✓ 1965ms | ✓ 929ms | ✓ 1226ms | ✓ 1179ms | http |
| 38.34.183.13:8449 | ✓ 1187ms | 否 | ✓ 1017ms | ✓ 1671ms | ✓ 1273ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1447ms | ✓ 1672ms | ✓ 1738ms | http |
| 106.75.15.167:7890 | ✓ 1315ms | ✓ 1362ms | ✓ 971ms | 否 | ✓ 1623ms | http |
| 38.34.179.155:8452 | ✓ 310ms | ✓ 1175ms | ✓ 735ms | ✓ 928ms | ✓ 698ms | http |
| 142.171.224.229:7890 | ✓ 511ms | ✓ 1480ms | ✓ 984ms | ✓ 828ms | ✓ 644ms | http |
| 103.113.70.189:1081 | ✓ 327ms | 否 | ✓ 903ms | ✓ 990ms | ✓ 772ms | http |
| 20.78.118.91:8561 | ✓ 595ms | ✓ 1206ms | ✓ 591ms | ✓ 1005ms | ✓ 722ms | http |
| 38.34.179.162:8451 | ✓ 1377ms | ✓ 1744ms | ✓ 292ms | ✓ 874ms | ✓ 680ms | http |
| 38.34.179.165:8446 | ✓ 690ms | ✓ 1593ms | ✓ 653ms | ✓ 1060ms | ✓ 1240ms | http |
| 20.27.15.111:8561 | ✓ 1610ms | 否 | ✓ 744ms | ✓ 943ms | ✓ 737ms | http |
| 20.210.39.153:8561 | ✓ 612ms | ✓ 1632ms | ✓ 589ms | ✓ 939ms | ✓ 748ms | http |
| 38.145.208.185:8449 | ✓ 395ms | 否 | ✓ 901ms | ✓ 866ms | ✓ 1261ms | http |
| 20.27.13.35:8561 | ✓ 687ms | ✓ 1607ms | ✓ 626ms | ✓ 979ms | ✓ 740ms | http |
| 38.34.179.27:8451 | ✓ 408ms | 否 | ✓ 302ms | ✓ 1124ms | ✓ 1034ms | http |
| 20.27.11.248:8561 | ✓ 654ms | 否 | ✓ 613ms | ✓ 907ms | ✓ 734ms | http |
| 38.145.218.228:8447 | ✓ 608ms | 否 | ✓ 274ms | ✓ 966ms | ✓ 1787ms | http |
| 38.34.179.49:8450 | ✓ 575ms | ✓ 1720ms | ✓ 633ms | ✓ 1053ms | ✓ 1076ms | http |
| 3.137.216.199:1080 | 否 | ✓ 837ms | ✓ 1044ms | ✓ 1479ms | ✓ 1112ms | http |
| 20.78.26.206:8561 | ✓ 1651ms | ✓ 1482ms | ✓ 596ms | ✓ 906ms | ✓ 717ms | http |
| 43.99.54.236:5555 | ✓ 830ms | 否 | ✓ 762ms | ✓ 959ms | ✓ 766ms | http |
| 38.34.178.245:8446 | ✓ 392ms | ✓ 1974ms | ✓ 1506ms | ✓ 892ms | ✓ 667ms | http |
| 45.136.131.54:8448 | ✓ 613ms | 否 | ✓ 337ms | ✓ 1207ms | ✓ 844ms | http |
| 45.136.130.162:8443 | 否 | 否 | ✓ 1075ms | ✓ 1093ms | ✓ 1391ms | http |
| 38.34.179.81:8446 | ✓ 603ms | 否 | ✓ 1652ms | ✓ 983ms | ✓ 752ms | http |
| 38.34.179.16:8451 | ✓ 469ms | 否 | ✓ 1649ms | ✓ 910ms | ✓ 1311ms | http |
| 38.34.179.57:8453 | 否 | ✓ 1927ms | ✓ 1160ms | ✓ 1824ms | ✓ 689ms | http |
| 144.31.79.117:8888 | ✓ 730ms | ✓ 1804ms | ✓ 1363ms | ✓ 1796ms | ✓ 1591ms | http |
| 77.232.135.22:1080 | ✓ 883ms | 否 | ✓ 1486ms | 否 | ✓ 1307ms | http |
| 45.136.130.177:8448 | ✓ 1017ms | ✓ 1091ms | ✓ 1001ms | ✓ 1905ms | ✓ 878ms | http |
| 38.34.178.186:8451 | ✓ 1477ms | ✓ 1604ms | ✓ 1370ms | ✓ 1840ms | ✓ 774ms | http |
| 38.145.208.241:8453 | ✓ 586ms | 否 | ✓ 1222ms | ✓ 917ms | ✓ 1218ms | http |
| 120.92.212.16:7890 | ✓ 1156ms | ✓ 1600ms | 否 | 否 | ✓ 1930ms | http |
| 38.145.208.242:8451 | ✓ 715ms | ✓ 1779ms | ✓ 1344ms | ✓ 856ms | ✓ 1077ms | http |
| 45.167.124.52:8080 | ✓ 1641ms | 否 | ✓ 533ms | ✓ 1541ms | ✓ 1326ms | http |
| 38.145.208.244:8448 | ✓ 677ms | ✓ 1873ms | ✓ 1256ms | 否 | 否 | http |
| 20.27.14.220:8561 | ✓ 1428ms | 否 | ✓ 576ms | ✓ 907ms | ✓ 753ms | http |
| 104.168.158.236:10808 | ✓ 455ms | 否 | ✓ 555ms | 否 | ✓ 848ms | http |
| 38.34.183.222:8453 | ✓ 1030ms | ✓ 1356ms | 否 | ✓ 1220ms | ✓ 1274ms | http |
| 167.103.31.122:8800 | ✓ 1740ms | 否 | ✓ 1348ms | ✓ 1580ms | ✓ 1468ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1463ms | ✓ 1841ms | ✓ 1009ms | http |
| 219.117.204.211:7799 | ✓ 1607ms | 否 | ✓ 643ms | ✓ 1003ms | ✓ 766ms | http |
| 120.92.212.16:8890 | ✓ 1229ms | ✓ 1383ms | 否 | ✓ 1414ms | 否 | http |
| 38.145.208.181:8445 | ✓ 674ms | ✓ 813ms | ✓ 681ms | ✓ 976ms | ✓ 679ms | http |
| 45.136.131.68:8444 | ✓ 386ms | 否 | ✓ 278ms | ✓ 1118ms | ✓ 1056ms | http |
| 38.34.179.105:8449 | ✓ 1433ms | ✓ 998ms | ✓ 442ms | ✓ 979ms | ✓ 1117ms | http |
| 38.34.179.75:8453 | ✓ 588ms | ✓ 1139ms | ✓ 1056ms | ✓ 1028ms | 否 | http |
| 38.34.179.54:8446 | ✓ 499ms | ✓ 1952ms | ✓ 287ms | ✓ 916ms | ✓ 778ms | http |
| 38.34.179.173:8452 | ✓ 439ms | ✓ 1932ms | ✓ 941ms | 否 | 否 | http |
| 38.34.178.155:8448 | ✓ 584ms | ✓ 1605ms | ✓ 340ms | ✓ 888ms | ✓ 733ms | http |
| 38.34.179.150:8453 | ✓ 604ms | ✓ 874ms | ✓ 571ms | ✓ 1444ms | ✓ 696ms | http |
| 38.34.183.130:8451 | ✓ 602ms | ✓ 868ms | ✓ 589ms | ✓ 1364ms | ✓ 1375ms | http |
| 38.34.179.101:8446 | 否 | ✓ 1480ms | ✓ 294ms | ✓ 971ms | ✓ 687ms | http |
| 113.160.132.26:8080 | ✓ 1965ms | ✓ 1496ms | ✓ 1167ms | ✓ 1428ms | ✓ 1114ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1856ms | ✓ 1480ms | ✓ 1458ms | http |
| 137.220.150.22:6005 | 否 | 否 | ✓ 1789ms | ✓ 1785ms | ✓ 1976ms | http |
| 101.43.127.100:8877 | ✓ 1867ms | ✓ 1214ms | 否 | ✓ 1498ms | ✓ 986ms | http |
| 45.136.130.156:8450 | ✓ 586ms | 否 | ✓ 821ms | ✓ 1018ms | ✓ 1065ms | http |
| 38.145.203.106:8448 | ✓ 1301ms | ✓ 1880ms | ✓ 442ms | ✓ 1385ms | ✓ 759ms | http |
| 49.156.44.114:8080 | 否 | 否 | ✓ 1875ms | ✓ 1737ms | ✓ 1716ms | http |
| 38.34.179.72:8445 | ✓ 428ms | 否 | ✓ 864ms | ✓ 881ms | ✓ 677ms | http |
| 45.140.147.155:1082 | ✓ 407ms | ✓ 1163ms | ✓ 1303ms | ✓ 1600ms | ✓ 1049ms | http |
| 147.161.239.240:8800 | ✓ 707ms | 否 | ✓ 1060ms | ✓ 1568ms | 否 | http |
| 34.158.226.196:443 | ✓ 694ms | 否 | ✓ 1487ms | 否 | ✓ 1823ms | http |
| 38.34.183.233:8448 | ✓ 1863ms | 否 | ✓ 959ms | ✓ 874ms | ✓ 1139ms | http |
| 35.225.22.61:80 | ✓ 518ms | ✓ 1127ms | ✓ 288ms | ✓ 1167ms | ✓ 974ms | http |
| 38.34.179.69:8448 | ✓ 604ms | ✓ 823ms | ✓ 620ms | ✓ 896ms | ✓ 776ms | http |
| 38.34.183.13:8444 | ✓ 248ms | ✓ 1884ms | ✓ 242ms | ✓ 961ms | ✓ 1167ms | http |
| 45.136.130.174:8451 | ✓ 707ms | 否 | ✓ 409ms | ✓ 870ms | ✓ 959ms | http |
| 38.145.220.193:8449 | ✓ 603ms | ✓ 1331ms | ✓ 825ms | ✓ 938ms | ✓ 872ms | http |
| 38.34.178.241:8444 | ✓ 603ms | ✓ 1800ms | ✓ 338ms | ✓ 1047ms | ✓ 811ms | http |
| 137.220.150.152:6005 | ✓ 1610ms | 否 | ✓ 1017ms | ✓ 1384ms | ✓ 1314ms | http |
| 181.78.44.63:999 | ✓ 1071ms | 否 | ✓ 832ms | ✓ 1434ms | ✓ 1718ms | http |
| 217.76.243.2:999 | 否 | 否 | ✓ 1040ms | ✓ 1424ms | ✓ 1299ms | http |
| 38.145.220.11:8445 | 否 | ✓ 1903ms | ✓ 406ms | ✓ 1144ms | ✓ 996ms | http |
| 45.136.131.59:8450 | ✓ 1025ms | 否 | ✓ 350ms | ✓ 899ms | ✓ 1513ms | http |
| 45.136.131.36:8450 | ✓ 800ms | ✓ 1568ms | ✓ 1294ms | ✓ 886ms | ✓ 1218ms | http |
| 38.34.179.98:8453 | 否 | ✓ 1633ms | ✓ 1569ms | ✓ 1165ms | ✓ 756ms | http |
| 45.136.131.35:8452 | ✓ 1542ms | ✓ 1043ms | ✓ 1844ms | ✓ 1191ms | ✓ 828ms | http |
| 38.34.179.106:8446 | 否 | ✓ 1619ms | ✓ 1142ms | ✓ 1101ms | ✓ 687ms | http |
| 38.34.183.130:8452 | 否 | ✓ 926ms | ✓ 931ms | ✓ 1719ms | ✓ 1284ms | http |
| 150.241.77.172:1080 | ✓ 1853ms | 否 | ✓ 1018ms | 否 | ✓ 1338ms | http |
| 38.145.203.19:8449 | ✓ 1563ms | ✓ 921ms | ✓ 719ms | 否 | ✓ 1229ms | http |
| 38.34.179.61:8445 | ✓ 1365ms | ✓ 882ms | ✓ 410ms | ✓ 1087ms | ✓ 1222ms | http |
| 38.34.179.174:8453 | ✓ 655ms | ✓ 1328ms | ✓ 1263ms | ✓ 910ms | ✓ 1081ms | http |
| 38.34.178.7:8452 | ✓ 1983ms | 否 | ✓ 267ms | ✓ 915ms | ✓ 1040ms | http |
| 38.34.178.141:8453 | ✓ 1803ms | ✓ 1867ms | ✓ 570ms | ✓ 920ms | ✓ 1052ms | http |
| 38.34.179.26:8450 | ✓ 728ms | 否 | ✓ 681ms | ✓ 1365ms | ✓ 786ms | http |
| 45.136.130.188:8449 | ✓ 1623ms | 否 | ✓ 1535ms | ✓ 919ms | ✓ 710ms | http |
| 38.34.179.83:8448 | ✓ 1773ms | 否 | ✓ 1761ms | ✓ 894ms | ✓ 1117ms | http |
| 38.34.179.150:8449 | ✓ 938ms | 否 | ✓ 338ms | ✓ 1070ms | ✓ 1386ms | http |
| 137.220.150.104:6005 | ✓ 1290ms | 否 | ✓ 843ms | ✓ 1470ms | ✓ 1756ms | http |
| 38.34.183.8:8450 | ✓ 727ms | ✓ 832ms | ✓ 411ms | ✓ 871ms | ✓ 694ms | http |
| 38.34.179.51:8449 | ✓ 1379ms | ✓ 821ms | ✓ 591ms | ✓ 1283ms | ✓ 1726ms | http |
| 38.34.179.48:8449 | ✓ 1530ms | ✓ 860ms | ✓ 772ms | ✓ 1837ms | ✓ 1028ms | http |
| 38.34.179.172:8451 | ✓ 1388ms | 否 | ✓ 911ms | ✓ 932ms | ✓ 731ms | http |
| 137.220.151.110:6005 | ✓ 899ms | 否 | ✓ 1811ms | ✓ 1434ms | ✓ 1356ms | http |

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
