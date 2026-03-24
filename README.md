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

最后更新：2026-03-24 05:48:13 UTC（2026-03-24 13:48:13 UTC+8）

**代理总数：167**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 167 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 167 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1574ms | ✓ 1725ms | ✓ 1025ms | 否 | ✓ 1038ms | http |
| 167.103.115.102:8800 | ✓ 1468ms | 否 | ✓ 1078ms | ✓ 1353ms | ✓ 1093ms | http |
| 8.217.106.71:8888 | ✓ 1978ms | ✓ 1561ms | ✓ 1251ms | ✓ 915ms | ✓ 1603ms | http |
| 113.160.132.26:8080 | ✓ 1482ms | ✓ 1794ms | ✓ 1237ms | ✓ 1457ms | ✓ 1101ms | http |
| 101.47.73.135:3128 | ✓ 1981ms | 否 | 否 | ✓ 1749ms | ✓ 1944ms | http |
| 167.103.34.108:8800 | ✓ 1509ms | 否 | 否 | ✓ 1819ms | ✓ 1590ms | http |
| 45.167.124.52:8080 | ✓ 1326ms | 否 | ✓ 1556ms | 否 | ✓ 1505ms | http |
| 43.99.54.236:5555 | ✓ 811ms | ✓ 1386ms | ✓ 742ms | ✓ 909ms | ✓ 730ms | http |
| 193.233.22.29:10808 | ✓ 1139ms | 否 | ✓ 329ms | ✓ 1435ms | ✓ 1946ms | http |
| 38.34.179.150:8449 | ✓ 1689ms | 否 | ✓ 762ms | ✓ 1291ms | ✓ 1350ms | http |
| 38.34.179.27:8451 | ✓ 1622ms | ✓ 1923ms | ✓ 611ms | ✓ 1031ms | ✓ 1731ms | http |
| 155.212.132.241:3128 | ✓ 640ms | ✓ 1804ms | ✓ 603ms | ✓ 1824ms | ✓ 1418ms | http |
| 185.114.73.2:1080 | ✓ 1160ms | 否 | ✓ 1358ms | ✓ 1723ms | ✓ 1197ms | http |
| 219.117.204.211:7799 | ✓ 700ms | 否 | ✓ 1888ms | ✓ 1780ms | ✓ 1317ms | http |
| 38.34.183.130:8452 | ✓ 1913ms | 否 | 否 | ✓ 1288ms | ✓ 940ms | http |
| 38.34.179.39:8452 | ✓ 1657ms | 否 | 否 | ✓ 1511ms | ✓ 806ms | http |
| 120.92.212.16:7890 | ✓ 1010ms | ✓ 1309ms | 否 | ✓ 1348ms | 否 | http |
| 160.238.65.2:3128 | ✓ 1074ms | 否 | ✓ 436ms | ✓ 1915ms | ✓ 1577ms | http |
| 160.238.65.8:3128 | ✓ 1068ms | ✓ 1827ms | ✓ 606ms | ✓ 1947ms | ✓ 1566ms | http |
| 160.238.65.6:3128 | ✓ 1071ms | 否 | ✓ 436ms | ✓ 1903ms | ✓ 1604ms | http |
| 160.238.65.3:3128 | ✓ 1074ms | ✓ 1395ms | ✓ 1019ms | ✓ 1943ms | ✓ 1587ms | http |
| 160.238.65.7:3128 | ✓ 1074ms | 否 | ✓ 448ms | ✓ 1922ms | ✓ 1572ms | http |
| 160.238.65.4:3128 | ✓ 1068ms | 否 | ✓ 449ms | ✓ 1902ms | ✓ 1590ms | http |
| 160.238.65.9:3128 | ✓ 1074ms | ✓ 1637ms | ✓ 765ms | ✓ 1973ms | ✓ 1579ms | http |
| 160.238.65.5:3128 | ✓ 1070ms | 否 | ✓ 452ms | ✓ 1924ms | ✓ 1556ms | http |
| 213.220.62.62:3128 | ✓ 1076ms | 否 | ✓ 545ms | ✓ 1772ms | ✓ 1178ms | http |
| 45.88.0.113:3128 | ✓ 1074ms | 否 | ✓ 938ms | ✓ 1398ms | ✓ 1148ms | http |
| 167.103.31.122:8800 | ✓ 1748ms | 否 | ✓ 1366ms | ✓ 1650ms | ✓ 1569ms | http |
| 45.88.0.115:3128 | ✓ 1075ms | 否 | ✓ 585ms | ✓ 1796ms | ✓ 1147ms | http |
| 45.88.0.117:3128 | ✓ 1075ms | 否 | ✓ 575ms | ✓ 1543ms | ✓ 1112ms | http |
| 45.88.0.111:3128 | ✓ 1073ms | 否 | ✓ 940ms | ✓ 1431ms | 否 | http |
| 45.88.0.114:3128 | ✓ 1076ms | ✓ 1848ms | ✓ 562ms | 否 | ✓ 1732ms | http |
| 45.88.0.99:3128 | 否 | 否 | ✓ 1652ms | ✓ 1555ms | ✓ 1430ms | http |
| 45.88.0.116:3128 | ✓ 1074ms | ✓ 1745ms | 否 | ✓ 1379ms | 否 | http |
| 45.88.0.98:3128 | ✓ 1791ms | 否 | 否 | ✓ 1453ms | ✓ 1159ms | http |
| 35.225.22.61:80 | ✓ 560ms | 否 | ✓ 1287ms | ✓ 992ms | ✓ 936ms | http |
| 106.75.15.167:7890 | ✓ 1337ms | ✓ 1300ms | 否 | 否 | ✓ 1293ms | http |
| 38.145.208.185:8453 | ✓ 301ms | ✓ 1211ms | ✓ 254ms | ✓ 1039ms | ✓ 638ms | http |
| 38.145.208.234:8445 | ✓ 709ms | ✓ 1242ms | ✓ 228ms | ✓ 971ms | ✓ 916ms | http |
| 38.145.208.166:8450 | ✓ 284ms | 否 | ✓ 210ms | ✓ 801ms | ✓ 772ms | http |
| 38.145.220.41:8449 | ✓ 1277ms | ✓ 1488ms | ✓ 273ms | ✓ 998ms | ✓ 811ms | http |
| 45.136.131.64:8451 | ✓ 1112ms | ✓ 1985ms | ✓ 418ms | ✓ 865ms | ✓ 864ms | http |
| 64.227.76.27:1080 | ✓ 498ms | 否 | ✓ 1649ms | ✓ 1657ms | ✓ 1120ms | http |
| 45.136.131.62:8446 | ✓ 1832ms | ✓ 756ms | ✓ 309ms | ✓ 1250ms | 否 | http |
| 38.34.179.29:8449 | ✓ 1164ms | ✓ 1961ms | ✓ 262ms | ✓ 864ms | ✓ 1565ms | http |
| 38.34.179.62:8453 | ✓ 506ms | 否 | ✓ 257ms | ✓ 832ms | ✓ 858ms | http |
| 38.145.218.218:8446 | ✓ 654ms | ✓ 798ms | ✓ 387ms | ✓ 1402ms | ✓ 916ms | http |
| 38.34.179.40:8446 | ✓ 953ms | 否 | ✓ 505ms | ✓ 894ms | ✓ 699ms | http |
| 137.220.150.22:6005 | ✓ 903ms | 否 | ✓ 819ms | ✓ 1208ms | 否 | http |
| 137.220.150.104:6005 | ✓ 1016ms | 否 | ✓ 1696ms | 否 | ✓ 964ms | http |
| 120.92.212.16:8890 | ✓ 1056ms | ✓ 1342ms | ✓ 1038ms | ✓ 1292ms | ✓ 1053ms | http |
| 45.136.131.54:8448 | ✓ 1007ms | 否 | ✓ 467ms | ✓ 1320ms | ✓ 1049ms | http |
| 45.144.28.81:10808 | ✓ 472ms | ✓ 1970ms | ✓ 1090ms | ✓ 1904ms | ✓ 1387ms | http |
| 137.220.150.152:6005 | ✓ 1164ms | 否 | ✓ 1134ms | ✓ 1233ms | ✓ 1101ms | http |
| 1.231.81.166:3128 | ✓ 1431ms | ✓ 1974ms | ✓ 1436ms | ✓ 1235ms | ✓ 1106ms | http |
| 38.34.179.61:8445 | 否 | 否 | ✓ 1462ms | ✓ 893ms | ✓ 1738ms | http |
| 147.161.239.240:8800 | ✓ 1439ms | ✓ 1658ms | ✓ 1077ms | ✓ 1471ms | ✓ 1285ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1462ms | ✓ 1142ms | ✓ 1674ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1015ms | 否 | ✓ 1672ms | ✓ 1999ms | ✓ 1533ms | http |
| 45.136.198.40:3128 | ✓ 757ms | ✓ 1903ms | ✓ 1425ms | 否 | ✓ 1756ms | http |
| 150.249.255.91:3128 | ✓ 1945ms | 否 | ✓ 690ms | 否 | ✓ 1198ms | http |
| 115.231.181.40:8128 | ✓ 955ms | ✓ 1713ms | 否 | 否 | ✓ 1921ms | http |
| 8.219.97.248:80 | ✓ 1274ms | 否 | ✓ 1578ms | ✓ 1841ms | 否 | http |
| 38.145.208.172:8448 | ✓ 414ms | 否 | ✓ 980ms | ✓ 929ms | ✓ 766ms | http |
| 45.136.130.173:8448 | ✓ 862ms | 否 | ✓ 597ms | ✓ 924ms | ✓ 860ms | http |
| 45.136.130.168:8448 | ✓ 383ms | ✓ 1771ms | ✓ 404ms | ✓ 897ms | ✓ 1241ms | http |
| 103.113.70.189:1081 | ✓ 379ms | ✓ 1802ms | ✓ 1109ms | ✓ 1180ms | ✓ 1861ms | http |
| 38.34.179.52:8445 | ✓ 834ms | ✓ 950ms | ✓ 263ms | ✓ 863ms | ✓ 716ms | http |
| 38.145.208.179:8446 | ✓ 893ms | ✓ 874ms | ✓ 755ms | ✓ 1286ms | ✓ 706ms | http |
| 38.145.218.212:8444 | ✓ 885ms | ✓ 948ms | ✓ 240ms | ✓ 921ms | ✓ 959ms | http |
| 38.145.220.32:8447 | ✓ 885ms | ✓ 957ms | ✓ 737ms | ✓ 1531ms | ✓ 641ms | http |
| 38.145.208.172:8444 | ✓ 1137ms | ✓ 1519ms | ✓ 376ms | ✓ 961ms | ✓ 774ms | http |
| 38.145.220.43:8444 | ✓ 839ms | ✓ 1405ms | ✓ 924ms | ✓ 817ms | ✓ 758ms | http |
| 38.34.179.67:8449 | ✓ 835ms | ✓ 997ms | ✓ 1190ms | ✓ 1666ms | ✓ 629ms | http |
| 38.34.179.62:8449 | ✓ 837ms | ✓ 1973ms | ✓ 1227ms | ✓ 892ms | ✓ 650ms | http |
| 38.145.220.27:8444 | ✓ 884ms | ✓ 1166ms | ✓ 1413ms | ✓ 1286ms | ✓ 965ms | http |
| 38.34.179.51:8449 | ✓ 839ms | ✓ 1448ms | ✓ 1465ms | ✓ 1452ms | ✓ 670ms | http |
| 45.136.130.181:8446 | ✓ 836ms | 否 | ✓ 977ms | 否 | ✓ 632ms | http |
| 38.34.179.96:8449 | ✓ 836ms | ✓ 967ms | ✓ 1557ms | 否 | ✓ 1013ms | http |
| 45.136.130.172:8453 | ✓ 1454ms | ✓ 1598ms | ✓ 933ms | 否 | ✓ 621ms | http |
| 38.145.203.87:8445 | ✓ 1468ms | 否 | ✓ 1420ms | ✓ 1588ms | ✓ 649ms | http |
| 103.84.95.54:7890 | ✓ 1429ms | 否 | ✓ 747ms | ✓ 1249ms | 否 | http |
| 38.145.218.233:8452 | ✓ 1467ms | 否 | 否 | ✓ 1096ms | ✓ 629ms | http |
| 38.34.179.54:8444 | ✓ 837ms | 否 | ✓ 1723ms | ✓ 1855ms | 否 | http |
| 38.145.218.232:8451 | ✓ 549ms | ✓ 791ms | ✓ 318ms | ✓ 822ms | ✓ 694ms | http |
| 38.34.179.37:8451 | ✓ 311ms | ✓ 814ms | ✓ 803ms | ✓ 1196ms | ✓ 644ms | http |
| 38.145.218.14:8450 | ✓ 404ms | ✓ 931ms | ✓ 341ms | ✓ 888ms | ✓ 702ms | http |
| 38.34.179.36:8451 | ✓ 289ms | ✓ 777ms | ✓ 865ms | ✓ 1205ms | ✓ 636ms | http |
| 38.34.179.39:8445 | ✓ 298ms | ✓ 1290ms | ✓ 331ms | ✓ 1368ms | ✓ 630ms | http |
| 38.145.208.237:8446 | ✓ 219ms | 否 | ✓ 269ms | ✓ 807ms | ✓ 724ms | http |
| 38.145.208.216:8446 | ✓ 571ms | ✓ 1327ms | ✓ 730ms | ✓ 811ms | ✓ 731ms | http |
| 38.145.208.163:8448 | ✓ 271ms | 否 | ✓ 257ms | ✓ 849ms | ✓ 848ms | http |
| 38.145.208.180:8445 | ✓ 313ms | 否 | ✓ 283ms | ✓ 887ms | ✓ 928ms | http |
| 38.145.208.204:8451 | ✓ 707ms | ✓ 1267ms | ✓ 645ms | ✓ 861ms | ✓ 896ms | http |
| 38.145.208.222:8452 | ✓ 1529ms | ✓ 842ms | ✓ 254ms | ✓ 858ms | ✓ 881ms | http |
| 38.145.208.218:8453 | ✓ 807ms | ✓ 1863ms | ✓ 245ms | ✓ 833ms | ✓ 853ms | http |
| 38.145.208.221:8453 | ✓ 823ms | 否 | ✓ 252ms | ✓ 900ms | ✓ 683ms | http |
| 38.34.179.21:8449 | ✓ 773ms | ✓ 833ms | ✓ 374ms | ✓ 1222ms | ✓ 1398ms | http |
| 38.145.218.229:8448 | ✓ 720ms | ✓ 1562ms | ✓ 409ms | ✓ 1012ms | ✓ 921ms | http |
| 38.145.203.43:8451 | ✓ 840ms | ✓ 1472ms | ✓ 560ms | ✓ 1015ms | ✓ 1080ms | http |

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
