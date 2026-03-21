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

最后更新：2026-03-21 09:29:28 UTC（2026-03-21 17:29:28 UTC+8）

**代理总数：138**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 137 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 138 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1483ms | 否 | ✓ 1075ms | ✓ 989ms | ✓ 844ms | http |
| 137.220.151.110:6005 | ✓ 1068ms | 否 | ✓ 893ms | ✓ 1259ms | ✓ 1013ms | http |
| 147.161.239.240:8800 | ✓ 1248ms | ✓ 1655ms | ✓ 1163ms | ✓ 1600ms | ✓ 1604ms | http |
| 113.160.132.26:8080 | ✓ 1570ms | ✓ 1947ms | ✓ 1197ms | ✓ 1258ms | ✓ 1031ms | http |
| 133.242.138.34:8100 | ✓ 1951ms | 否 | ✓ 1131ms | ✓ 1742ms | ✓ 964ms | http |
| 137.220.150.104:6005 | ✓ 987ms | 否 | 否 | ✓ 1173ms | ✓ 1107ms | http |
| 45.167.124.52:8080 | ✓ 845ms | ✓ 1873ms | ✓ 858ms | ✓ 1818ms | ✓ 1475ms | http |
| 38.145.203.19:8449 | 否 | 否 | ✓ 505ms | ✓ 1236ms | ✓ 1568ms | http |
| 137.220.150.22:6005 | ✓ 949ms | 否 | ✓ 930ms | ✓ 1317ms | ✓ 1010ms | http |
| 38.34.179.186:8444 | 否 | 否 | ✓ 1870ms | ✓ 1283ms | ✓ 981ms | http |
| 101.47.73.135:3128 | ✓ 1820ms | 否 | ✓ 1488ms | ✓ 1332ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1229ms | 否 | ✓ 1228ms | ✓ 1409ms | ✓ 1332ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1309ms | ✓ 1102ms | ✓ 1757ms | ✓ 1667ms | http |
| 103.3.246.71:3128 | ✓ 1264ms | 否 | ✓ 1471ms | ✓ 1634ms | ✓ 1070ms | http |
| 167.103.31.122:8800 | ✓ 1608ms | 否 | ✓ 1685ms | 否 | ✓ 1964ms | http |
| 38.34.179.74:8449 | 否 | 否 | ✓ 723ms | ✓ 1834ms | ✓ 1919ms | http |
| 137.184.1.87:3128 | ✓ 403ms | ✓ 1088ms | ✓ 711ms | ✓ 820ms | ✓ 639ms | http |
| 219.117.204.211:7799 | ✓ 1312ms | 否 | ✓ 573ms | ✓ 914ms | ✓ 748ms | http |
| 66.151.40.248:3128 | ✓ 1053ms | 否 | ✓ 1027ms | ✓ 1769ms | 否 | http |
| 38.34.179.78:8445 | ✓ 770ms | 否 | ✓ 1437ms | ✓ 952ms | 否 | http |
| 35.225.22.61:80 | ✓ 321ms | ✓ 1330ms | ✓ 1939ms | ✓ 1036ms | ✓ 871ms | http |
| 8.219.97.248:80 | ✓ 1605ms | 否 | ✓ 1422ms | 否 | ✓ 1806ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1717ms | ✓ 1577ms | ✓ 1453ms | http |
| 38.34.183.234:8450 | ✓ 317ms | ✓ 920ms | ✓ 426ms | ✓ 942ms | ✓ 647ms | http |
| 38.34.179.173:8452 | ✓ 388ms | ✓ 843ms | ✓ 477ms | ✓ 1104ms | 否 | http |
| 139.159.99.242:8080 | ✓ 992ms | ✓ 1111ms | ✓ 924ms | 否 | 否 | http |
| 38.34.179.172:8451 | ✓ 315ms | ✓ 904ms | ✓ 392ms | ✓ 953ms | ✓ 791ms | http |
| 38.34.178.245:8446 | ✓ 251ms | ✓ 1016ms | ✓ 578ms | ✓ 1066ms | ✓ 619ms | http |
| 142.171.224.229:7890 | ✓ 477ms | ✓ 993ms | ✓ 977ms | ✓ 805ms | ✓ 657ms | http |
| 45.136.131.39:8451 | ✓ 292ms | ✓ 1511ms | ✓ 211ms | ✓ 900ms | ✓ 1060ms | http |
| 38.34.183.222:8453 | ✓ 241ms | ✓ 928ms | ✓ 1074ms | ✓ 1653ms | 否 | http |
| 38.145.208.171:8449 | ✓ 559ms | ✓ 798ms | ✓ 391ms | ✓ 787ms | ✓ 734ms | http |
| 38.34.179.25:8448 | ✓ 554ms | ✓ 1034ms | ✓ 213ms | ✓ 821ms | ✓ 632ms | http |
| 38.34.179.21:8452 | ✓ 553ms | 否 | ✓ 242ms | ✓ 876ms | ✓ 802ms | http |
| 38.145.220.81:8446 | ✓ 674ms | ✓ 820ms | ✓ 1007ms | ✓ 1046ms | ✓ 1428ms | http |
| 45.136.130.168:8452 | ✓ 943ms | ✓ 1021ms | ✓ 1265ms | ✓ 948ms | ✓ 683ms | http |
| 38.145.203.32:8452 | ✓ 554ms | 否 | ✓ 1123ms | ✓ 813ms | ✓ 632ms | http |
| 172.212.68.37:3128 | ✓ 957ms | ✓ 1394ms | ✓ 806ms | ✓ 1874ms | ✓ 1424ms | http |
| 192.71.213.85:9091 | ✓ 1026ms | 否 | ✓ 1404ms | ✓ 1963ms | 否 | http |
| 91.238.105.64:2024 | ✓ 998ms | 否 | ✓ 1876ms | 否 | ✓ 1565ms | http |
| 217.174.244.117:3129 | ✓ 1001ms | 否 | ✓ 1623ms | 否 | ✓ 1683ms | http |
| 101.43.127.100:8877 | ✓ 1696ms | ✓ 1380ms | ✓ 939ms | 否 | ✓ 924ms | http |
| 45.88.0.113:3128 | ✓ 1080ms | ✓ 1669ms | ✓ 922ms | ✓ 1786ms | 否 | http |
| 45.88.0.117:3128 | ✓ 1087ms | 否 | ✓ 582ms | 否 | ✓ 1080ms | http |
| 45.88.0.98:3128 | 否 | ✓ 1889ms | ✓ 579ms | ✓ 1348ms | ✓ 1442ms | http |
| 106.75.15.167:7890 | 否 | 否 | ✓ 1171ms | ✓ 1508ms | ✓ 1000ms | http |
| 38.34.179.179:8449 | ✓ 353ms | ✓ 778ms | ✓ 197ms | ✓ 856ms | ✓ 636ms | http |
| 38.145.203.34:8444 | ✓ 214ms | ✓ 757ms | ✓ 211ms | ✓ 840ms | ✓ 658ms | http |
| 38.145.208.224:8445 | ✓ 205ms | ✓ 758ms | ✓ 206ms | ✓ 842ms | ✓ 794ms | http |
| 38.145.208.193:8452 | ✓ 279ms | ✓ 984ms | ✓ 267ms | ✓ 809ms | ✓ 611ms | http |
| 38.34.179.190:8451 | ✓ 418ms | ✓ 1169ms | ✓ 207ms | ✓ 918ms | ✓ 984ms | http |
| 93.183.69.162:1080 | 否 | 否 | ✓ 952ms | ✓ 1701ms | ✓ 1624ms | http |
| 181.78.44.63:999 | ✓ 951ms | 否 | ✓ 502ms | ✓ 1410ms | ✓ 1149ms | http |
| 47.101.159.19:8899 | ✓ 966ms | ✓ 1178ms | ✓ 1141ms | ✓ 1196ms | ✓ 935ms | http |
| 103.82.23.118:5247 | ✓ 1794ms | 否 | ✓ 1493ms | ✓ 1908ms | ✓ 1754ms | http |
| 148.153.56.51:80 | ✓ 1801ms | 否 | 否 | ✓ 1861ms | ✓ 1533ms | http |
| 88.80.150.82:8080 | ✓ 1473ms | 否 | 否 | ✓ 1605ms | ✓ 1646ms | https |
| 38.34.183.47:8452 | ✓ 1368ms | ✓ 831ms | ✓ 1367ms | 否 | 否 | http |
| 38.145.220.198:8448 | ✓ 550ms | ✓ 794ms | ✓ 694ms | 否 | 否 | http |
| 38.34.179.20:8445 | ✓ 1426ms | ✓ 1729ms | ✓ 846ms | 否 | 否 | http |
| 38.34.179.98:8453 | ✓ 1423ms | 否 | ✓ 1465ms | ✓ 1270ms | ✓ 742ms | http |
| 137.220.150.170:6005 | ✓ 977ms | 否 | ✓ 1838ms | ✓ 1252ms | ✓ 1425ms | http |
| 45.136.131.53:8452 | ✓ 1109ms | 否 | ✓ 333ms | ✓ 1248ms | 否 | http |
| 45.136.131.54:8448 | ✓ 949ms | ✓ 1946ms | ✓ 614ms | ✓ 1977ms | ✓ 1350ms | http |
| 167.71.60.190:8080 | ✓ 831ms | 否 | 否 | ✓ 1873ms | ✓ 1950ms | http |
| 38.34.179.16:8451 | ✓ 780ms | ✓ 1198ms | ✓ 838ms | ✓ 805ms | ✓ 762ms | http |
| 38.145.220.33:8448 | ✓ 1041ms | ✓ 804ms | ✓ 381ms | ✓ 1583ms | ✓ 1577ms | http |
| 38.34.183.130:8452 | ✓ 241ms | ✓ 784ms | ✓ 192ms | ✓ 866ms | ✓ 643ms | http |
| 24.144.86.173:1080 | ✓ 816ms | 否 | ✓ 1310ms | ✓ 1268ms | ✓ 663ms | http |
| 210.223.44.230:3128 | ✓ 735ms | ✓ 971ms | ✓ 992ms | ✓ 1439ms | ✓ 806ms | http |
| 45.144.28.81:10808 | ✓ 600ms | ✓ 1345ms | ✓ 496ms | 否 | 否 | http |
| 47.77.193.180:1080 | ✓ 532ms | ✓ 1088ms | ✓ 173ms | ✓ 819ms | ✓ 681ms | http |
| 38.145.208.242:8451 | ✓ 1399ms | ✓ 1664ms | ✓ 1066ms | 否 | ✓ 653ms | http |
| 38.145.208.241:8453 | ✓ 1833ms | 否 | 否 | ✓ 1146ms | ✓ 685ms | http |
| 45.136.130.174:8450 | ✓ 226ms | ✓ 1281ms | ✓ 401ms | ✓ 1223ms | ✓ 742ms | http |
| 38.145.203.132:8452 | ✓ 827ms | ✓ 1834ms | ✓ 937ms | ✓ 911ms | ✓ 626ms | http |
| 38.145.218.101:8448 | ✓ 450ms | ✓ 1625ms | ✓ 318ms | ✓ 816ms | ✓ 849ms | http |
| 45.136.130.171:8452 | ✓ 315ms | 否 | ✓ 1092ms | ✓ 796ms | ✓ 692ms | http |
| 38.34.179.165:8446 | ✓ 893ms | ✓ 1944ms | ✓ 357ms | ✓ 982ms | ✓ 1251ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 1720ms | ✓ 1540ms | ✓ 1164ms | http |
| 213.220.62.62:3128 | ✓ 1004ms | ✓ 1676ms | ✓ 643ms | 否 | ✓ 1371ms | http |
| 162.240.154.26:3128 | ✓ 1420ms | 否 | 否 | ✓ 890ms | ✓ 1752ms | http |
| 38.34.179.162:8451 | ✓ 341ms | 否 | ✓ 1645ms | ✓ 1090ms | ✓ 995ms | http |
| 5.104.87.17:8050 | ✓ 774ms | 否 | 否 | ✓ 1409ms | ✓ 933ms | http |
| 45.88.0.116:3128 | 否 | 否 | ✓ 1974ms | ✓ 1341ms | ✓ 1031ms | http |
| 103.145.34.67:8080 | 否 | 否 | ✓ 1407ms | ✓ 1489ms | ✓ 1438ms | http |
| 39.100.88.235:3256 | ✓ 1009ms | ✓ 1377ms | ✓ 1032ms | 否 | 否 | http |
| 166.88.55.83:7890 | ✓ 799ms | ✓ 1186ms | ✓ 713ms | 否 | 否 | http |
| 154.64.240.85:1080 | ✓ 1008ms | ✓ 1275ms | ✓ 1019ms | ✓ 959ms | ✓ 791ms | http |
| 103.84.95.54:7890 | ✓ 740ms | 否 | ✓ 957ms | ✓ 969ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1545ms | ✓ 1952ms | ✓ 1504ms | 否 | 否 | http |
| 38.145.203.135:8451 | ✓ 332ms | ✓ 790ms | ✓ 262ms | ✓ 967ms | ✓ 1704ms | http |
| 38.145.218.113:8444 | ✓ 627ms | ✓ 1387ms | ✓ 392ms | ✓ 1172ms | ✓ 624ms | http |
| 38.145.208.215:8453 | ✓ 341ms | ✓ 834ms | ✓ 875ms | ✓ 872ms | ✓ 939ms | http |
| 38.145.208.253:8445 | ✓ 463ms | ✓ 1137ms | ✓ 1222ms | ✓ 1491ms | ✓ 624ms | http |
| 45.136.131.43:8449 | ✓ 806ms | 否 | ✓ 396ms | ✓ 1144ms | ✓ 1224ms | http |
| 38.34.179.160:8453 | ✓ 590ms | ✓ 948ms | ✓ 508ms | ✓ 1105ms | ✓ 678ms | http |
| 45.136.130.173:8448 | ✓ 307ms | ✓ 1777ms | ✓ 877ms | ✓ 1868ms | ✓ 668ms | http |
| 38.145.220.72:8444 | ✓ 920ms | 否 | ✓ 973ms | 否 | ✓ 805ms | http |
| 38.145.220.175:8452 | ✓ 470ms | ✓ 1390ms | ✓ 885ms | ✓ 1256ms | ✓ 1977ms | http |

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
