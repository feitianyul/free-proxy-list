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

最后更新：2026-03-26 11:44:31 UTC（2026-03-26 19:44:31 UTC+8）

**代理总数：144**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 144 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 144 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1668ms | ✓ 1226ms | ✓ 1475ms | ✓ 919ms | ✓ 985ms | http |
| 38.34.179.161:8448 | 否 | ✓ 1524ms | ✓ 932ms | ✓ 661ms | ✓ 637ms | http |
| 5.104.87.17:8051 | ✓ 1108ms | 否 | ✓ 1239ms | ✓ 1030ms | 否 | http |
| 178.218.105.99:3128 | ✓ 1434ms | 否 | ✓ 1770ms | ✓ 1663ms | ✓ 1323ms | http |
| 45.167.124.52:8080 | ✓ 939ms | 否 | ✓ 1545ms | ✓ 1630ms | ✓ 1332ms | http |
| 38.34.179.40:8446 | ✓ 141ms | ✓ 620ms | ✓ 97ms | ✓ 673ms | 否 | http |
| 38.34.179.178:8445 | ✓ 299ms | ✓ 636ms | ✓ 112ms | ✓ 698ms | ✓ 1716ms | http |
| 38.34.179.174:8453 | ✓ 198ms | ✓ 709ms | ✓ 129ms | ✓ 737ms | 否 | http |
| 38.145.218.229:8450 | ✓ 346ms | ✓ 1252ms | ✓ 99ms | ✓ 708ms | 否 | http |
| 38.34.179.6:8449 | ✓ 867ms | ✓ 1644ms | ✓ 615ms | ✓ 680ms | ✓ 726ms | http |
| 38.145.208.213:8450 | ✓ 1223ms | ✓ 1533ms | ✓ 927ms | ✓ 1145ms | 否 | http |
| 219.117.204.211:7799 | 否 | ✓ 1461ms | ✓ 1034ms | ✓ 1400ms | ✓ 1363ms | http |
| 38.145.208.242:8451 | ✓ 713ms | ✓ 1393ms | ✓ 1094ms | ✓ 706ms | ✓ 692ms | http |
| 38.145.220.198:8448 | ✓ 749ms | ✓ 703ms | ✓ 280ms | 否 | ✓ 1384ms | http |
| 183.249.5.117:22222 | ✓ 1608ms | ✓ 1108ms | ✓ 1246ms | ✓ 1762ms | ✓ 1292ms | http |
| 167.103.34.108:8800 | ✓ 1130ms | 否 | ✓ 1177ms | ✓ 1340ms | ✓ 1276ms | http |
| 120.92.212.16:7890 | ✓ 976ms | ✓ 1204ms | ✓ 1641ms | 否 | ✓ 987ms | http |
| 45.136.130.191:8453 | ✓ 1090ms | ✓ 1168ms | ✓ 557ms | 否 | 否 | http |
| 167.103.144.127:8800 | 否 | 否 | ✓ 1385ms | ✓ 1342ms | ✓ 1245ms | http |
| 38.145.220.33:8448 | ✓ 895ms | ✓ 1123ms | ✓ 485ms | 否 | ✓ 1185ms | http |
| 38.145.218.228:8447 | ✓ 1129ms | ✓ 738ms | ✓ 1299ms | 否 | ✓ 1778ms | http |
| 38.34.179.172:8451 | ✓ 664ms | 否 | ✓ 1589ms | ✓ 934ms | ✓ 1328ms | http |
| 186.148.180.46:999 | ✓ 1001ms | ✓ 1950ms | ✓ 1373ms | ✓ 1846ms | ✓ 1570ms | http |
| 115.231.181.40:8128 | ✓ 896ms | ✓ 1103ms | ✓ 1569ms | ✓ 1175ms | 否 | http |
| 147.161.239.240:8800 | 否 | ✓ 1812ms | ✓ 1083ms | ✓ 1414ms | ✓ 1487ms | http |
| 222.184.48.251:22222 | ✓ 1749ms | 否 | ✓ 1912ms | ✓ 1194ms | ✓ 895ms | http |
| 35.225.22.61:80 | ✓ 783ms | 否 | ✓ 554ms | ✓ 1421ms | 否 | http |
| 8.219.97.248:80 | ✓ 1473ms | 否 | ✓ 1277ms | 否 | ✓ 1709ms | http |
| 38.145.220.102:8453 | ✓ 1556ms | 否 | ✓ 262ms | ✓ 1489ms | ✓ 662ms | http |
| 38.145.208.209:8444 | 否 | ✓ 1970ms | ✓ 231ms | ✓ 1044ms | ✓ 1654ms | http |
| 120.92.212.16:8890 | ✓ 1492ms | 否 | ✓ 1167ms | ✓ 1861ms | ✓ 982ms | http |
| 38.145.208.181:8445 | ✓ 195ms | ✓ 737ms | ✓ 785ms | 否 | 否 | http |
| 160.250.4.13:1 | ✓ 1441ms | 否 | ✓ 1122ms | 否 | ✓ 1342ms | http |
| 101.43.127.100:8877 | ✓ 1864ms | ✓ 1119ms | ✓ 1162ms | ✓ 1380ms | ✓ 1613ms | http |
| 113.160.132.26:8080 | ✓ 1424ms | 否 | 否 | ✓ 1752ms | ✓ 1045ms | http |
| 218.89.134.230:3333 | ✓ 1716ms | ✓ 1910ms | ✓ 1975ms | ✓ 1940ms | ✓ 1543ms | http |
| 38.34.183.224:8448 | ✓ 470ms | ✓ 1807ms | ✓ 1193ms | ✓ 1613ms | ✓ 721ms | http |
| 167.103.31.122:8800 | ✓ 1364ms | 否 | ✓ 1310ms | ✓ 1643ms | ✓ 1559ms | http |
| 38.34.179.150:8449 | ✓ 1641ms | ✓ 1645ms | ✓ 1361ms | ✓ 1555ms | ✓ 1445ms | http |
| 107.174.208.190:3128 | ✓ 150ms | ✓ 616ms | ✓ 654ms | ✓ 711ms | ✓ 650ms | http |
| 103.84.95.54:7890 | ✓ 977ms | ✓ 1674ms | 否 | 否 | ✓ 643ms | http |
| 38.34.179.162:8453 | ✓ 495ms | ✓ 691ms | ✓ 122ms | ✓ 837ms | ✓ 491ms | http |
| 38.34.179.189:8453 | ✓ 495ms | ✓ 682ms | ✓ 1540ms | ✓ 1834ms | ✓ 840ms | http |
| 222.184.48.252:22222 | 否 | ✓ 1146ms | ✓ 972ms | ✓ 1408ms | ✓ 926ms | http |
| 222.228.171.92:8080 | ✓ 1795ms | 否 | ✓ 1789ms | ✓ 1838ms | ✓ 1746ms | http |
| 45.136.130.168:8452 | ✓ 1148ms | ✓ 715ms | ✓ 909ms | ✓ 1137ms | ✓ 1408ms | http |
| 38.34.179.86:8452 | ✓ 410ms | ✓ 1116ms | ✓ 644ms | ✓ 1382ms | ✓ 1015ms | http |
| 101.34.21.55:90 | ✓ 1937ms | ✓ 1921ms | ✓ 813ms | ✓ 1784ms | 否 | http |
| 38.145.208.244:8448 | ✓ 569ms | ✓ 656ms | ✓ 591ms | ✓ 1743ms | ✓ 1450ms | http |
| 180.125.216.109:8118 | 否 | 否 | ✓ 997ms | ✓ 1533ms | ✓ 900ms | http |
| 45.136.130.171:8445 | ✓ 1180ms | ✓ 785ms | ✓ 1516ms | 否 | 否 | http |
| 38.145.220.182:8453 | ✓ 787ms | 否 | ✓ 917ms | ✓ 667ms | ✓ 817ms | http |
| 103.113.70.189:1081 | ✓ 862ms | 否 | ✓ 808ms | ✓ 1218ms | ✓ 848ms | http |
| 167.103.115.102:8800 | ✓ 937ms | 否 | ✓ 1975ms | ✓ 1098ms | ✓ 992ms | http |
| 62.113.119.14:8080 | ✓ 820ms | 否 | ✓ 1665ms | ✓ 1839ms | ✓ 1787ms | http |
| 38.145.208.207:8445 | ✓ 1591ms | ✓ 1371ms | 否 | ✓ 1691ms | 否 | http |
| 138.197.68.35:4857 | 否 | ✓ 1997ms | ✓ 260ms | ✓ 1374ms | ✓ 882ms | http |
| 194.67.99.223:1080 | ✓ 1267ms | 否 | ✓ 1535ms | 否 | ✓ 1536ms | http |
| 38.145.208.189:8446 | ✓ 1108ms | 否 | ✓ 116ms | ✓ 818ms | ✓ 554ms | http |
| 38.145.208.165:8449 | ✓ 1110ms | ✓ 1299ms | ✓ 257ms | ✓ 655ms | ✓ 667ms | http |
| 38.145.208.190:8444 | ✓ 1414ms | 否 | ✓ 200ms | ✓ 804ms | ✓ 1464ms | http |
| 185.41.152.110:3128 | ✓ 852ms | ✓ 1554ms | ✓ 1681ms | 否 | 否 | http |
| 5.102.109.41:999 | ✓ 1541ms | 否 | ✓ 1609ms | ✓ 1325ms | ✓ 1101ms | http |
| 59.46.216.131:30001 | ✓ 953ms | 否 | 否 | ✓ 1947ms | ✓ 1707ms | http |
| 38.145.208.241:8453 | ✓ 217ms | ✓ 1121ms | ✓ 102ms | ✓ 699ms | ✓ 527ms | http |
| 38.145.203.132:8450 | 否 | 否 | ✓ 805ms | ✓ 688ms | ✓ 482ms | http |
| 183.249.5.111:22222 | ✓ 748ms | ✓ 1084ms | ✓ 720ms | ✓ 1004ms | ✓ 1032ms | http |
| 38.34.179.83:8448 | ✓ 1340ms | ✓ 1275ms | ✓ 274ms | ✓ 776ms | ✓ 1200ms | http |
| 38.34.179.173:8452 | ✓ 1393ms | ✓ 1057ms | ✓ 151ms | 否 | ✓ 533ms | http |
| 38.34.183.130:8452 | ✓ 1101ms | 否 | 否 | ✓ 678ms | ✓ 514ms | http |
| 183.249.5.105:22222 | 否 | 否 | ✓ 1837ms | ✓ 1272ms | ✓ 952ms | http |
| 104.247.51.76:3128 | ✓ 633ms | 否 | ✓ 932ms | ✓ 1298ms | ✓ 989ms | http |
| 160.250.134.143:3128 | ✓ 1645ms | 否 | ✓ 1239ms | ✓ 1203ms | ✓ 1035ms | http |
| 38.145.208.189:8450 | ✓ 168ms | ✓ 670ms | ✓ 107ms | ✓ 696ms | ✓ 526ms | http |
| 38.145.208.193:8451 | ✓ 168ms | ✓ 693ms | ✓ 92ms | ✓ 687ms | ✓ 523ms | http |
| 38.145.208.197:8451 | ✓ 167ms | ✓ 1448ms | ✓ 109ms | ✓ 696ms | ✓ 528ms | http |
| 38.145.203.106:8448 | ✓ 567ms | ✓ 1018ms | ✓ 384ms | ✓ 659ms | ✓ 958ms | http |
| 38.145.208.191:8446 | ✓ 168ms | ✓ 753ms | ✓ 118ms | ✓ 1808ms | ✓ 805ms | http |
| 38.145.218.10:8451 | ✓ 166ms | ✓ 1199ms | ✓ 92ms | ✓ 1682ms | ✓ 539ms | http |
| 38.145.220.41:8449 | ✓ 419ms | ✓ 610ms | ✓ 98ms | ✓ 700ms | ✓ 1758ms | http |
| 45.136.131.60:8453 | ✓ 275ms | ✓ 692ms | ✓ 162ms | ✓ 687ms | 否 | http |
| 38.145.220.173:8444 | ✓ 403ms | ✓ 1155ms | ✓ 726ms | ✓ 1270ms | ✓ 849ms | http |
| 38.145.218.87:8450 | ✓ 256ms | ✓ 938ms | ✓ 593ms | 否 | ✓ 904ms | http |
| 38.145.220.35:8444 | ✓ 167ms | ✓ 1599ms | ✓ 233ms | 否 | ✓ 534ms | http |
| 38.34.179.152:8450 | ✓ 168ms | ✓ 822ms | ✓ 569ms | ✓ 867ms | 否 | http |
| 38.145.208.183:8446 | ✓ 528ms | ✓ 1236ms | ✓ 213ms | ✓ 1188ms | ✓ 670ms | http |
| 38.145.203.124:8452 | ✓ 331ms | ✓ 1256ms | ✓ 129ms | ✓ 1189ms | 否 | http |
| 38.145.203.98:8445 | ✓ 333ms | 否 | ✓ 80ms | 否 | ✓ 700ms | http |
| 38.34.179.150:8444 | ✓ 168ms | ✓ 784ms | ✓ 611ms | 否 | ✓ 1639ms | http |
| 45.136.130.247:8450 | ✓ 257ms | ✓ 1558ms | ✓ 377ms | ✓ 985ms | 否 | http |
| 45.136.130.247:8449 | ✓ 256ms | ✓ 1822ms | ✓ 86ms | 否 | ✓ 1131ms | http |
| 38.145.208.162:8452 | ✓ 254ms | ✓ 809ms | ✓ 136ms | ✓ 847ms | ✓ 1564ms | http |
| 38.34.178.141:8446 | ✓ 203ms | ✓ 1845ms | ✓ 629ms | ✓ 791ms | 否 | http |
| 38.145.218.51:8448 | ✓ 258ms | ✓ 792ms | ✓ 721ms | 否 | ✓ 966ms | http |
| 38.145.218.13:8452 | ✓ 266ms | ✓ 797ms | ✓ 701ms | 否 | ✓ 1360ms | http |
| 45.136.130.251:8444 | ✓ 252ms | 否 | ✓ 299ms | ✓ 1014ms | 否 | http |
| 38.145.218.217:8452 | ✓ 252ms | ✓ 1214ms | ✓ 221ms | 否 | 否 | http |
| 38.145.218.234:8451 | ✓ 248ms | ✓ 821ms | ✓ 617ms | 否 | 否 | http |
| 38.145.203.86:8449 | ✓ 244ms | ✓ 1602ms | ✓ 493ms | 否 | ✓ 1099ms | http |
| 38.145.208.180:8452 | ✓ 1508ms | ✓ 1225ms | ✓ 144ms | ✓ 1219ms | 否 | http |

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
