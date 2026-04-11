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

最后更新：2026-04-11 19:43:41 UTC（2026-04-12 03:43:41 UTC+8）

**代理总数：242**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 242 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 242 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.34.179.174:8453 | ✓ 576ms | ✓ 984ms | ✓ 355ms | ✓ 805ms | ✓ 579ms | http |
| 147.161.210.140:8800 | ✓ 664ms | 否 | ✓ 1115ms | ✓ 912ms | ✓ 748ms | http |
| 5.196.101.18:3128 | ✓ 625ms | ✓ 1569ms | ✓ 1301ms | 否 | ✓ 1618ms | http |
| 159.223.225.118:8888 | ✓ 890ms | 否 | ✓ 1273ms | 否 | ✓ 1354ms | http |
| 167.103.115.102:8800 | ✓ 923ms | 否 | ✓ 1192ms | 否 | ✓ 1020ms | http |
| 137.59.47.73:3128 | ✓ 1398ms | 否 | ✓ 1262ms | ✓ 1425ms | ✓ 901ms | http |
| 167.103.34.108:8800 | ✓ 1420ms | 否 | ✓ 1441ms | ✓ 1366ms | ✓ 1288ms | http |
| 38.145.208.241:8447 | ✓ 914ms | ✓ 1963ms | ✓ 970ms | ✓ 1552ms | ✓ 1469ms | http |
| 45.136.130.247:8448 | ✓ 581ms | ✓ 1401ms | ✓ 1536ms | ✓ 1373ms | ✓ 1105ms | http |
| 113.160.132.26:8080 | ✓ 1418ms | ✓ 1362ms | ✓ 1563ms | 否 | ✓ 1494ms | http |
| 38.34.179.105:8447 | ✓ 458ms | ✓ 814ms | ✓ 189ms | ✓ 750ms | ✓ 551ms | http |
| 38.145.220.39:8449 | ✓ 329ms | ✓ 897ms | ✓ 252ms | ✓ 920ms | ✓ 714ms | http |
| 38.145.220.81:8445 | ✓ 418ms | ✓ 917ms | ✓ 236ms | ✓ 907ms | ✓ 695ms | http |
| 38.145.218.51:8444 | ✓ 521ms | ✓ 1008ms | ✓ 409ms | ✓ 923ms | ✓ 747ms | http |
| 38.34.179.46:8448 | ✓ 391ms | ✓ 859ms | ✓ 754ms | ✓ 1050ms | ✓ 1039ms | http |
| 43.99.86.50:30219 | ✓ 675ms | ✓ 933ms | ✓ 662ms | ✓ 841ms | ✓ 676ms | http |
| 43.99.83.201:44467 | ✓ 719ms | ✓ 946ms | ✓ 655ms | ✓ 865ms | ✓ 678ms | http |
| 38.145.203.111:8450 | ✓ 581ms | ✓ 1016ms | ✓ 591ms | ✓ 1385ms | ✓ 1286ms | http |
| 38.145.208.224:8445 | ✓ 327ms | ✓ 1956ms | ✓ 147ms | ✓ 854ms | ✓ 877ms | http |
| 38.145.208.242:8444 | ✓ 181ms | ✓ 1677ms | ✓ 1266ms | ✓ 831ms | ✓ 1253ms | http |
| 38.34.179.24:8447 | ✓ 818ms | ✓ 1367ms | ✓ 231ms | ✓ 876ms | ✓ 1172ms | http |
| 38.34.179.77:8448 | ✓ 556ms | ✓ 1400ms | ✓ 865ms | ✓ 810ms | ✓ 1102ms | http |
| 59.46.216.131:30001 | ✓ 1037ms | ✓ 1375ms | ✓ 1152ms | ✓ 1398ms | ✓ 1198ms | http |
| 167.103.144.127:8800 | ✓ 1731ms | 否 | ✓ 1251ms | ✓ 1541ms | ✓ 1301ms | http |
| 45.136.131.28:8449 | ✓ 1130ms | ✓ 1549ms | ✓ 775ms | ✓ 1525ms | ✓ 1436ms | http |
| 212.58.132.5:8888 | ✓ 1624ms | 否 | ✓ 1293ms | ✓ 1529ms | ✓ 1242ms | http |
| 38.145.208.223:8450 | ✓ 383ms | ✓ 883ms | ✓ 1140ms | ✓ 758ms | ✓ 900ms | http |
| 38.145.208.217:8450 | ✓ 365ms | ✓ 880ms | ✓ 1167ms | ✓ 757ms | ✓ 876ms | http |
| 5.104.87.17:8051 | ✓ 1799ms | 否 | ✓ 1912ms | ✓ 1688ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1228ms | 否 | ✓ 1462ms | 否 | ✓ 1624ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1855ms | ✓ 1466ms | ✓ 1876ms | 否 | http |
| 217.76.245.80:999 | ✓ 897ms | ✓ 1614ms | ✓ 1189ms | ✓ 1779ms | ✓ 1431ms | http |
| 45.167.125.21:999 | ✓ 904ms | ✓ 1916ms | ✓ 1785ms | ✓ 1850ms | ✓ 1467ms | http |
| 147.161.239.240:8800 | ✓ 1235ms | ✓ 1855ms | ✓ 1076ms | ✓ 1849ms | ✓ 1364ms | http |
| 34.101.184.164:3128 | ✓ 914ms | 否 | ✓ 1478ms | ✓ 1602ms | ✓ 1460ms | http |
| 91.238.105.43:2023 | ✓ 1295ms | 否 | ✓ 1736ms | 否 | ✓ 1728ms | http |
| 167.103.31.122:8800 | ✓ 1350ms | 否 | ✓ 1362ms | ✓ 1660ms | ✓ 1575ms | http |
| 190.12.150.244:999 | ✓ 1544ms | 否 | ✓ 1149ms | ✓ 1987ms | ✓ 1776ms | http |
| 155.117.18.36:25388 | ✓ 494ms | ✓ 1037ms | ✓ 1291ms | ✓ 1217ms | ✓ 927ms | http |
| 120.92.212.16:8890 | ✓ 964ms | ✓ 1227ms | ✓ 1186ms | 否 | ✓ 1261ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1458ms | ✓ 1890ms | ✓ 1480ms | ✓ 1286ms | http |
| 35.225.22.61:80 | ✓ 695ms | 否 | ✓ 837ms | ✓ 1165ms | ✓ 971ms | http |
| 38.145.208.204:8446 | ✓ 1116ms | ✓ 867ms | ✓ 1216ms | 否 | 否 | http |
| 38.145.208.181:8445 | ✓ 1116ms | ✓ 912ms | ✓ 1401ms | ✓ 883ms | ✓ 782ms | http |
| 38.145.208.187:8449 | ✓ 995ms | ✓ 967ms | ✓ 1245ms | ✓ 971ms | ✓ 888ms | http |
| 38.145.218.14:8446 | 否 | ✓ 1862ms | ✓ 304ms | ✓ 1042ms | ✓ 1047ms | http |
| 38.34.179.61:8445 | 否 | ✓ 1888ms | ✓ 208ms | ✓ 816ms | ✓ 1125ms | http |
| 38.34.179.101:8453 | ✓ 1558ms | ✓ 1958ms | ✓ 159ms | ✓ 808ms | ✓ 680ms | http |
| 103.113.70.189:1082 | ✓ 1752ms | 否 | ✓ 961ms | ✓ 1966ms | ✓ 1049ms | http |
| 38.145.203.35:8450 | ✓ 1130ms | ✓ 1619ms | ✓ 1192ms | ✓ 1642ms | ✓ 1368ms | http |
| 38.145.220.182:8450 | 否 | ✓ 1832ms | ✓ 1557ms | ✓ 1733ms | ✓ 793ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1859ms | ✓ 861ms | ✓ 1879ms | 否 | http |
| 38.145.208.226:8448 | ✓ 1369ms | 否 | ✓ 728ms | ✓ 1898ms | ✓ 1119ms | http |
| 38.34.179.69:8447 | ✓ 1541ms | ✓ 1898ms | ✓ 513ms | ✓ 1575ms | ✓ 1985ms | http |
| 38.34.179.11:8447 | 否 | ✓ 1628ms | ✓ 895ms | ✓ 1526ms | ✓ 1274ms | http |
| 38.34.179.104:8447 | 否 | ✓ 1914ms | ✓ 238ms | ✓ 995ms | ✓ 1581ms | http |
| 38.34.179.19:8448 | ✓ 1215ms | ✓ 1548ms | ✓ 1184ms | ✓ 1992ms | ✓ 1476ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1886ms | ✓ 854ms | ✓ 1232ms | 否 | http |
| 38.145.220.77:8453 | ✓ 1093ms | ✓ 1365ms | ✓ 837ms | ✓ 1343ms | 否 | http |
| 38.34.179.178:8444 | ✓ 1756ms | ✓ 1889ms | ✓ 1066ms | 否 | 否 | http |
| 38.34.179.17:8446 | ✓ 1399ms | ✓ 1827ms | ✓ 1731ms | ✓ 1456ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1121ms | ✓ 1348ms | ✓ 1376ms | 否 | 否 | http |
| 38.145.208.246:8445 | ✓ 1026ms | ✓ 852ms | ✓ 1459ms | 否 | 否 | http |
| 177.234.217.88:999 | ✓ 1881ms | 否 | ✓ 1867ms | 否 | ✓ 1583ms | http |
| 38.34.179.39:8452 | ✓ 1752ms | ✓ 1667ms | ✓ 1847ms | 否 | 否 | http |
| 38.34.179.65:8448 | ✓ 630ms | ✓ 1020ms | ✓ 647ms | ✓ 829ms | ✓ 596ms | http |
| 38.34.179.67:8446 | ✓ 631ms | ✓ 1013ms | ✓ 910ms | ✓ 776ms | ✓ 589ms | http |
| 38.34.179.105:8449 | ✓ 629ms | ✓ 1127ms | ✓ 872ms | ✓ 886ms | ✓ 740ms | http |
| 45.136.130.176:8451 | ✓ 638ms | ✓ 1046ms | ✓ 782ms | ✓ 1157ms | ✓ 724ms | http |
| 38.145.218.229:8444 | ✓ 635ms | ✓ 1034ms | ✓ 964ms | ✓ 1232ms | ✓ 591ms | http |
| 45.136.131.40:8449 | ✓ 867ms | ✓ 1732ms | ✓ 311ms | ✓ 801ms | ✓ 607ms | http |
| 38.145.208.245:8449 | ✓ 629ms | ✓ 1044ms | ✓ 627ms | ✓ 1142ms | ✓ 1341ms | http |
| 38.145.208.209:8447 | ✓ 865ms | ✓ 1706ms | ✓ 568ms | ✓ 953ms | ✓ 673ms | http |
| 38.145.218.210:8448 | ✓ 1580ms | ✓ 798ms | ✓ 257ms | ✓ 943ms | ✓ 802ms | http |
| 38.145.208.241:8453 | ✓ 631ms | ✓ 1038ms | ✓ 627ms | ✓ 1050ms | ✓ 1443ms | http |
| 38.145.218.87:8451 | ✓ 873ms | ✓ 848ms | ✓ 1164ms | ✓ 1845ms | ✓ 618ms | http |
| 38.145.220.72:8451 | ✓ 1124ms | ✓ 1452ms | ✓ 264ms | ✓ 767ms | ✓ 726ms | http |
| 38.34.179.99:8446 | ✓ 657ms | 否 | ✓ 1244ms | ✓ 818ms | ✓ 658ms | http |
| 38.145.220.55:8444 | ✓ 871ms | ✓ 1309ms | ✓ 837ms | ✓ 1207ms | ✓ 673ms | http |
| 38.145.203.98:8446 | ✓ 661ms | ✓ 1592ms | ✓ 670ms | ✓ 1265ms | ✓ 889ms | http |
| 38.145.220.168:8453 | ✓ 867ms | ✓ 1947ms | ✓ 353ms | ✓ 1104ms | ✓ 572ms | http |
| 38.34.179.164:8448 | ✓ 1105ms | ✓ 1294ms | ✓ 411ms | ✓ 1204ms | ✓ 1173ms | http |
| 45.136.131.61:8444 | ✓ 1102ms | ✓ 1597ms | ✓ 273ms | ✓ 869ms | ✓ 1132ms | http |
| 38.34.179.81:8443 | ✓ 898ms | 否 | ✓ 1093ms | ✓ 1228ms | ✓ 641ms | http |
| 45.136.131.30:8446 | ✓ 633ms | ✓ 1804ms | ✓ 1009ms | ✓ 736ms | ✓ 584ms | http |
| 38.34.179.176:8446 | ✓ 1106ms | ✓ 1748ms | ✓ 692ms | ✓ 1434ms | ✓ 1247ms | http |
| 38.145.220.188:8451 | ✓ 1125ms | ✓ 1888ms | ✓ 302ms | ✓ 784ms | ✓ 746ms | http |
| 38.145.220.41:8444 | ✓ 869ms | ✓ 1364ms | ✓ 1073ms | ✓ 1248ms | ✓ 725ms | http |
| 47.238.220.4:8888 | ✓ 900ms | ✓ 1666ms | ✓ 674ms | ✓ 851ms | ✓ 878ms | http |
| 101.43.127.100:8877 | ✓ 992ms | ✓ 1073ms | ✓ 1601ms | ✓ 1199ms | ✓ 949ms | http |
| 38.145.218.163:8451 | 否 | 否 | ✓ 1865ms | ✓ 787ms | ✓ 588ms | http |
| 45.136.130.249:8447 | ✓ 1590ms | 否 | ✓ 1114ms | ✓ 1387ms | ✓ 967ms | http |
| 38.145.203.19:8447 | ✓ 1129ms | ✓ 1614ms | ✓ 1607ms | ✓ 1305ms | ✓ 626ms | http |
| 46.39.105.157:8080 | ✓ 857ms | ✓ 1993ms | ✓ 1743ms | ✓ 1717ms | ✓ 1481ms | http |
| 38.145.203.124:8444 | ✓ 1978ms | ✓ 1802ms | 否 | ✓ 958ms | ✓ 1296ms | http |
| 202.141.161.53:10808 | ✓ 1074ms | ✓ 1333ms | ✓ 1219ms | ✓ 1183ms | ✓ 1241ms | http |
| 181.78.44.63:999 | ✓ 1051ms | ✓ 1703ms | ✓ 1196ms | ✓ 1424ms | ✓ 1286ms | http |
| 38.145.220.82:8448 | ✓ 417ms | ✓ 1139ms | ✓ 664ms | ✓ 932ms | ✓ 670ms | http |
| 38.145.220.40:8450 | ✓ 417ms | ✓ 1158ms | ✓ 662ms | ✓ 933ms | ✓ 658ms | http |
| 38.34.179.102:8449 | ✓ 215ms | ✓ 1684ms | ✓ 653ms | ✓ 993ms | ✓ 619ms | http |

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
