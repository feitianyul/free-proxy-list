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

最后更新：2026-03-23 14:08:46 UTC（2026-03-23 22:08:46 UTC+8）

**代理总数：260**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 260 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 260 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1662ms | 否 | ✓ 986ms | ✓ 1160ms | ✓ 1094ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1376ms | ✓ 1813ms | ✓ 1706ms | http |
| 103.113.70.189:1081 | ✓ 697ms | 否 | ✓ 433ms | ✓ 984ms | ✓ 739ms | http |
| 104.129.202.127:12354 | ✓ 1273ms | ✓ 952ms | ✓ 1143ms | ✓ 962ms | ✓ 855ms | http |
| 142.171.224.229:7890 | ✓ 741ms | 否 | ✓ 976ms | ✓ 1017ms | ✓ 978ms | http |
| 104.129.202.127:10810 | ✓ 569ms | 否 | ✓ 1822ms | ✓ 908ms | ✓ 737ms | http |
| 46.101.190.71:3128 | ✓ 564ms | ✓ 1680ms | ✓ 1039ms | ✓ 1599ms | ✓ 1183ms | http |
| 167.103.34.108:8800 | ✓ 1324ms | 否 | ✓ 1354ms | ✓ 1553ms | ✓ 1401ms | http |
| 167.103.31.122:8800 | ✓ 1822ms | 否 | ✓ 1810ms | 否 | ✓ 1947ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 139ms | ✓ 1181ms | ✓ 794ms | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 1043ms | ✓ 1135ms | ✓ 768ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 790ms | ✓ 1019ms | ✓ 1515ms | http |
| 103.183.10.169:3125 | ✓ 1937ms | 否 | 否 | ✓ 1888ms | ✓ 1636ms | http |
| 59.46.216.131:30001 | ✓ 986ms | ✓ 1304ms | ✓ 1003ms | 否 | 否 | http |
| 172.212.68.37:3128 | ✓ 1287ms | 否 | ✓ 1457ms | ✓ 1891ms | ✓ 1032ms | http |
| 147.161.239.240:8800 | ✓ 1206ms | 否 | ✓ 867ms | 否 | ✓ 1449ms | http |
| 74.242.169.16:3128 | ✓ 1227ms | ✓ 1837ms | ✓ 1888ms | ✓ 1744ms | ✓ 1443ms | http |
| 113.160.132.26:8080 | ✓ 997ms | ✓ 1621ms | ✓ 1049ms | 否 | ✓ 1137ms | http |
| 115.231.181.40:8128 | ✓ 1893ms | ✓ 1122ms | 否 | ✓ 1201ms | 否 | http |
| 20.27.11.248:8561 | ✓ 1291ms | ✓ 1087ms | ✓ 590ms | ✓ 906ms | ✓ 859ms | http |
| 20.27.15.111:8561 | ✓ 1293ms | ✓ 1320ms | ✓ 571ms | ✓ 923ms | ✓ 713ms | http |
| 20.78.118.91:8561 | ✓ 1293ms | ✓ 1254ms | ✓ 616ms | ✓ 945ms | ✓ 845ms | http |
| 20.210.39.153:8561 | ✓ 1290ms | ✓ 1680ms | ✓ 600ms | ✓ 914ms | ✓ 727ms | http |
| 20.27.14.220:8561 | ✓ 1289ms | 否 | ✓ 583ms | ✓ 950ms | ✓ 729ms | http |
| 20.27.13.35:8561 | ✓ 1291ms | 否 | ✓ 586ms | ✓ 996ms | ✓ 728ms | http |
| 20.78.26.206:8561 | ✓ 1296ms | 否 | ✓ 585ms | ✓ 945ms | ✓ 716ms | http |
| 114.237.77.244:1080 | 否 | ✓ 1175ms | ✓ 1927ms | ✓ 1320ms | 否 | http |
| 114.237.77.216:1080 | 否 | 否 | ✓ 1270ms | ✓ 1119ms | ✓ 898ms | http |
| 45.136.130.187:8452 | ✓ 1757ms | 否 | ✓ 1736ms | ✓ 927ms | ✓ 1766ms | http |
| 38.145.208.242:8451 | ✓ 1768ms | 否 | ✓ 1835ms | 否 | ✓ 667ms | http |
| 38.34.179.98:8453 | ✓ 856ms | ✓ 931ms | ✓ 254ms | ✓ 1020ms | ✓ 690ms | http |
| 38.34.179.106:8446 | ✓ 880ms | 否 | ✓ 847ms | ✓ 855ms | ✓ 1878ms | http |
| 38.34.179.167:8450 | ✓ 594ms | ✓ 861ms | ✓ 400ms | ✓ 947ms | ✓ 673ms | http |
| 38.145.208.235:8453 | ✓ 912ms | ✓ 859ms | ✓ 260ms | ✓ 855ms | ✓ 686ms | http |
| 38.145.220.22:8445 | ✓ 1209ms | ✓ 1885ms | ✓ 1366ms | 否 | ✓ 704ms | http |
| 38.34.179.103:8453 | ✓ 591ms | 否 | ✓ 1575ms | ✓ 1115ms | 否 | http |
| 137.220.150.22:6005 | ✓ 1714ms | 否 | ✓ 1197ms | ✓ 1702ms | 否 | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1662ms | ✓ 1582ms | ✓ 1842ms | http |
| 38.34.179.26:8445 | ✓ 1041ms | ✓ 1969ms | ✓ 1239ms | 否 | 否 | http |
| 207.254.71.62:8088 | ✓ 1414ms | ✓ 1930ms | ✓ 1593ms | ✓ 1416ms | ✓ 1176ms | http |
| 94.79.152.14:80 | ✓ 1219ms | 否 | ✓ 1179ms | ✓ 1629ms | ✓ 1400ms | http |
| 120.92.212.16:8890 | ✓ 1641ms | 否 | ✓ 896ms | 否 | ✓ 898ms | http |
| 38.34.179.16:8451 | ✓ 905ms | ✓ 1144ms | ✓ 711ms | ✓ 1760ms | ✓ 1203ms | http |
| 38.145.218.228:8447 | ✓ 1320ms | 否 | ✓ 887ms | 否 | ✓ 969ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1713ms | ✓ 1956ms | ✓ 1499ms | http |
| 38.145.208.162:8443 | ✓ 1983ms | 否 | 否 | ✓ 1192ms | ✓ 690ms | http |
| 38.145.208.172:8448 | 否 | 否 | ✓ 457ms | ✓ 860ms | ✓ 660ms | http |
| 38.34.183.224:8448 | ✓ 1118ms | 否 | ✓ 273ms | ✓ 1452ms | ✓ 1884ms | http |
| 38.145.203.97:8444 | ✓ 686ms | ✓ 966ms | ✓ 513ms | ✓ 1010ms | ✓ 666ms | http |
| 38.34.179.162:8451 | ✓ 1262ms | ✓ 871ms | ✓ 334ms | ✓ 1556ms | ✓ 863ms | http |
| 120.92.212.16:7890 | ✓ 915ms | ✓ 1139ms | ✓ 928ms | ✓ 1391ms | ✓ 1208ms | http |
| 38.145.203.105:8452 | ✓ 307ms | ✓ 853ms | ✓ 265ms | ✓ 930ms | ✓ 663ms | http |
| 38.145.218.227:8449 | ✓ 441ms | ✓ 900ms | ✓ 259ms | ✓ 927ms | ✓ 677ms | http |
| 38.34.179.38:8452 | ✓ 314ms | ✓ 975ms | ✓ 275ms | ✓ 982ms | ✓ 811ms | http |
| 45.136.130.188:8445 | ✓ 462ms | ✓ 928ms | ✓ 267ms | ✓ 900ms | ✓ 717ms | http |
| 45.136.130.172:8451 | ✓ 312ms | ✓ 985ms | ✓ 251ms | ✓ 1099ms | ✓ 684ms | http |
| 45.136.130.171:8445 | ✓ 317ms | ✓ 1032ms | ✓ 259ms | ✓ 1230ms | ✓ 686ms | http |
| 45.136.131.56:8445 | ✓ 411ms | ✓ 893ms | ✓ 249ms | ✓ 1202ms | ✓ 822ms | http |
| 38.34.179.56:8447 | ✓ 249ms | ✓ 1447ms | ✓ 280ms | ✓ 917ms | ✓ 691ms | http |
| 38.145.220.32:8449 | ✓ 533ms | ✓ 802ms | ✓ 756ms | ✓ 895ms | ✓ 720ms | http |
| 38.34.179.39:8452 | ✓ 332ms | ✓ 1469ms | ✓ 246ms | ✓ 908ms | ✓ 917ms | http |
| 45.136.131.26:8445 | ✓ 269ms | ✓ 1325ms | ✓ 502ms | ✓ 948ms | ✓ 738ms | http |
| 38.145.208.167:8445 | ✓ 466ms | ✓ 1883ms | ✓ 246ms | ✓ 828ms | ✓ 636ms | http |
| 38.145.220.20:8449 | ✓ 319ms | ✓ 1719ms | ✓ 294ms | ✓ 1022ms | ✓ 682ms | http |
| 38.145.208.165:8451 | ✓ 464ms | ✓ 1857ms | ✓ 242ms | ✓ 857ms | ✓ 660ms | http |
| 38.145.220.39:8453 | ✓ 334ms | ✓ 1802ms | ✓ 296ms | ✓ 901ms | ✓ 714ms | http |
| 38.145.203.98:8453 | ✓ 280ms | 否 | ✓ 260ms | ✓ 884ms | ✓ 718ms | http |
| 38.34.179.23:8448 | ✓ 464ms | ✓ 1594ms | ✓ 351ms | ✓ 896ms | ✓ 718ms | http |
| 38.145.208.163:8449 | ✓ 480ms | 否 | ✓ 250ms | ✓ 841ms | ✓ 679ms | http |
| 38.145.218.206:8453 | ✓ 297ms | 否 | ✓ 243ms | ✓ 903ms | ✓ 732ms | http |
| 38.145.203.106:8448 | ✓ 286ms | ✓ 1750ms | ✓ 506ms | ✓ 1023ms | ✓ 652ms | http |
| 38.34.179.31:8445 | ✓ 579ms | ✓ 1230ms | ✓ 533ms | ✓ 1019ms | ✓ 727ms | http |
| 38.145.203.109:8450 | ✓ 549ms | ✓ 1878ms | ✓ 221ms | ✓ 916ms | ✓ 840ms | http |
| 38.145.203.97:8448 | ✓ 304ms | 否 | ✓ 446ms | ✓ 979ms | ✓ 661ms | http |
| 38.145.208.187:8451 | ✓ 451ms | ✓ 1780ms | ✓ 380ms | ✓ 942ms | ✓ 916ms | http |
| 38.145.218.208:8448 | ✓ 238ms | 否 | ✓ 754ms | ✓ 870ms | ✓ 658ms | http |
| 38.145.218.234:8445 | ✓ 272ms | ✓ 1896ms | ✓ 813ms | ✓ 865ms | ✓ 659ms | http |
| 38.34.179.25:8449 | ✓ 687ms | ✓ 863ms | ✓ 799ms | ✓ 1264ms | ✓ 1117ms | http |
| 45.136.130.193:8450 | ✓ 459ms | ✓ 1474ms | ✓ 358ms | ✓ 1016ms | ✓ 1276ms | http |
| 45.136.130.247:8450 | ✓ 249ms | 否 | ✓ 774ms | ✓ 891ms | ✓ 814ms | http |
| 38.145.208.226:8452 | ✓ 331ms | 否 | ✓ 645ms | ✓ 874ms | ✓ 703ms | http |
| 38.34.179.99:8444 | ✓ 333ms | 否 | ✓ 821ms | ✓ 895ms | ✓ 748ms | http |
| 38.34.179.101:8446 | ✓ 333ms | 否 | ✓ 823ms | ✓ 898ms | ✓ 758ms | http |
| 38.145.218.211:8451 | ✓ 525ms | ✓ 1421ms | ✓ 1246ms | ✓ 888ms | ✓ 691ms | http |
| 38.34.179.193:8451 | ✓ 332ms | ✓ 1223ms | ✓ 1441ms | ✓ 865ms | ✓ 1034ms | http |
| 38.34.179.96:8451 | ✓ 885ms | 否 | ✓ 290ms | ✓ 894ms | ✓ 923ms | http |
| 38.34.179.181:8448 | ✓ 322ms | ✓ 1399ms | ✓ 1257ms | ✓ 900ms | ✓ 1040ms | http |
| 38.34.183.224:8445 | ✓ 462ms | 否 | ✓ 757ms | ✓ 858ms | ✓ 692ms | http |
| 45.136.130.248:8447 | ✓ 707ms | ✓ 1476ms | ✓ 835ms | ✓ 963ms | ✓ 970ms | http |
| 38.145.208.185:8449 | ✓ 501ms | ✓ 1675ms | ✓ 395ms | ✓ 1661ms | ✓ 844ms | http |
| 38.34.179.38:8451 | ✓ 659ms | 否 | ✓ 703ms | ✓ 874ms | ✓ 667ms | http |
| 45.136.130.251:8445 | ✓ 249ms | 否 | ✓ 769ms | ✓ 1086ms | ✓ 994ms | http |
| 38.34.179.190:8446 | ✓ 318ms | 否 | ✓ 666ms | ✓ 1465ms | ✓ 680ms | http |
| 38.145.220.40:8446 | ✓ 315ms | ✓ 1908ms | ✓ 270ms | ✓ 899ms | ✓ 1770ms | http |
| 45.136.130.184:8448 | ✓ 548ms | ✓ 1046ms | ✓ 509ms | ✓ 1135ms | ✓ 934ms | http |
| 45.136.130.192:8451 | ✓ 852ms | ✓ 872ms | ✓ 472ms | ✓ 1784ms | ✓ 1409ms | http |
| 45.136.130.173:8448 | ✓ 560ms | ✓ 1560ms | ✓ 467ms | ✓ 864ms | ✓ 720ms | http |
| 45.136.130.189:8451 | ✓ 860ms | ✓ 903ms | ✓ 553ms | 否 | ✓ 1104ms | http |
| 38.145.203.35:8448 | ✓ 738ms | 否 | ✓ 285ms | ✓ 1012ms | ✓ 1492ms | http |
| 45.136.131.64:8452 | ✓ 264ms | ✓ 1532ms | 否 | ✓ 1051ms | ✓ 769ms | http |

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
