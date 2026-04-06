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

最后更新：2026-04-06 00:27:50 UTC（2026-04-06 08:27:50 UTC+8）

**代理总数：193**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 193 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 193 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 870ms | ✓ 839ms | ✓ 997ms | ✓ 706ms | ✓ 526ms | http |
| 35.225.22.61:80 | ✓ 877ms | ✓ 1267ms | ✓ 1018ms | 否 | ✓ 1055ms | http |
| 209.38.154.7:1080 | ✓ 870ms | ✓ 857ms | ✓ 1582ms | ✓ 732ms | 否 | http |
| 147.161.210.140:8800 | ✓ 613ms | ✓ 1544ms | ✓ 922ms | ✓ 1135ms | ✓ 948ms | http |
| 218.108.131.186:17890 | ✓ 809ms | ✓ 973ms | ✓ 860ms | ✓ 1075ms | ✓ 871ms | http |
| 115.231.181.40:8128 | ✓ 923ms | ✓ 1023ms | ✓ 976ms | ✓ 1185ms | ✓ 1007ms | http |
| 1.231.81.166:3128 | ✓ 1141ms | ✓ 967ms | ✓ 1696ms | ✓ 1042ms | ✓ 1059ms | http |
| 111.227.254.9:22222 | ✓ 1059ms | ✓ 1339ms | ✓ 1045ms | ✓ 1274ms | ✓ 1022ms | http |
| 111.227.254.10:22222 | ✓ 1075ms | ✓ 1304ms | ✓ 1072ms | ✓ 1273ms | ✓ 1097ms | http |
| 167.103.115.102:8800 | ✓ 1504ms | ✓ 1670ms | ✓ 896ms | ✓ 1294ms | ✓ 989ms | http |
| 111.227.254.12:22222 | ✓ 1370ms | ✓ 1258ms | ✓ 982ms | ✓ 1251ms | ✓ 1046ms | http |
| 111.227.254.11:22222 | ✓ 1075ms | ✓ 1351ms | ✓ 1078ms | ✓ 1389ms | ✓ 999ms | http |
| 5.104.87.17:8051 | ✓ 1493ms | 否 | ✓ 1542ms | ✓ 1290ms | ✓ 1048ms | http |
| 113.160.132.26:8080 | ✓ 1849ms | ✓ 1613ms | ✓ 1225ms | ✓ 1199ms | ✓ 931ms | http |
| 159.223.71.162:8080 | ✓ 1503ms | 否 | ✓ 1988ms | ✓ 1350ms | ✓ 823ms | http |
| 167.103.34.108:8800 | ✓ 1556ms | ✓ 1992ms | ✓ 1393ms | ✓ 1623ms | ✓ 1406ms | http |
| 159.223.71.162:443 | ✓ 1504ms | 否 | 否 | ✓ 1352ms | ✓ 1413ms | http |
| 212.58.132.5:8888 | ✓ 1501ms | 否 | ✓ 1791ms | ✓ 1557ms | ✓ 1290ms | http |
| 180.250.219.58:53281 | ✓ 1966ms | 否 | ✓ 1528ms | ✓ 1902ms | ✓ 1576ms | http |
| 138.197.68.35:4857 | 否 | ✓ 1476ms | ✓ 441ms | 否 | ✓ 1097ms | http |
| 103.113.70.189:1081 | ✓ 469ms | ✓ 1102ms | ✓ 871ms | ✓ 1184ms | ✓ 871ms | http |
| 5.104.87.17:8050 | ✓ 986ms | ✓ 1393ms | ✓ 1367ms | ✓ 824ms | ✓ 608ms | http |
| 120.92.212.16:8890 | ✓ 996ms | ✓ 1195ms | 否 | ✓ 1239ms | ✓ 987ms | http |
| 167.103.144.127:8800 | ✓ 1542ms | 否 | ✓ 1335ms | ✓ 1623ms | ✓ 1590ms | http |
| 38.145.208.185:8453 | 否 | ✓ 1199ms | ✓ 1423ms | 否 | ✓ 1040ms | http |
| 167.103.31.122:8800 | ✓ 1770ms | 否 | ✓ 1675ms | 否 | ✓ 1965ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1083ms | ✓ 771ms | ✓ 1382ms | ✓ 1456ms | http |
| 59.46.216.131:30001 | ✓ 943ms | ✓ 1267ms | ✓ 1061ms | ✓ 1328ms | ✓ 1059ms | http |
| 158.160.215.167:8124 | ✓ 1463ms | ✓ 1824ms | ✓ 1138ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 954ms | ✓ 1942ms | ✓ 952ms | ✓ 1216ms | 否 | http |
| 45.12.151.226:2829 | ✓ 961ms | ✓ 1740ms | 否 | ✓ 1806ms | 否 | http |
| 101.43.127.100:8877 | ✓ 866ms | ✓ 1140ms | ✓ 933ms | ✓ 1154ms | ✓ 923ms | http |
| 45.136.130.176:8451 | ✓ 1772ms | ✓ 1507ms | ✓ 1042ms | ✓ 1796ms | ✓ 841ms | http |
| 123.57.2.231:2020 | ✓ 1115ms | ✓ 1170ms | ✓ 1267ms | ✓ 1185ms | ✓ 1100ms | http |
| 147.161.239.240:8800 | ✓ 917ms | ✓ 1825ms | ✓ 1101ms | ✓ 1768ms | ✓ 1436ms | http |
| 101.132.61.121:8888 | 否 | 否 | ✓ 1225ms | ✓ 1479ms | ✓ 1296ms | http |
| 45.167.125.21:999 | ✓ 980ms | 否 | ✓ 1461ms | 否 | ✓ 1728ms | http |
| 103.63.26.149:8082 | ✓ 1567ms | 否 | ✓ 1306ms | ✓ 1483ms | 否 | http |
| 177.234.217.88:999 | ✓ 1326ms | ✓ 1991ms | ✓ 1937ms | 否 | ✓ 1666ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1543ms | ✓ 1399ms | ✓ 1565ms | ✓ 1259ms | http |
| 198.59.68.130:3128 | ✓ 808ms | ✓ 1042ms | ✓ 1358ms | ✓ 1840ms | ✓ 1852ms | http |
| 45.136.130.193:8447 | ✓ 1350ms | ✓ 612ms | ✓ 80ms | ✓ 1040ms | ✓ 762ms | http |
| 45.136.130.178:8453 | ✓ 1562ms | 否 | ✓ 1158ms | ✓ 1938ms | 否 | http |
| 38.145.208.211:8453 | ✓ 1335ms | ✓ 612ms | ✓ 126ms | ✓ 733ms | ✓ 952ms | http |
| 45.136.130.182:8446 | ✓ 1572ms | ✓ 863ms | ✓ 124ms | ✓ 738ms | ✓ 984ms | http |
| 38.34.179.21:8443 | ✓ 699ms | ✓ 1462ms | ✓ 467ms | ✓ 682ms | ✓ 539ms | http |
| 38.34.179.69:8447 | 否 | ✓ 1141ms | ✓ 97ms | ✓ 712ms | ✓ 1050ms | http |
| 38.34.179.87:8447 | ✓ 450ms | ✓ 1166ms | ✓ 1922ms | ✓ 948ms | ✓ 794ms | http |
| 38.34.179.150:8449 | ✓ 865ms | ✓ 1940ms | ✓ 867ms | 否 | ✓ 1267ms | http |
| 91.233.223.147:3128 | ✓ 1046ms | 否 | ✓ 968ms | 否 | ✓ 1617ms | http |
| 38.34.179.39:8452 | ✓ 613ms | 否 | ✓ 1227ms | ✓ 842ms | ✓ 1447ms | http |
| 38.34.179.82:8447 | ✓ 751ms | ✓ 1886ms | ✓ 337ms | ✓ 781ms | 否 | http |
| 38.34.179.186:8444 | ✓ 174ms | ✓ 852ms | ✓ 1017ms | ✓ 1151ms | ✓ 495ms | http |
| 38.34.179.199:8452 | ✓ 778ms | ✓ 618ms | ✓ 317ms | ✓ 917ms | ✓ 1999ms | http |
| 47.74.226.8:5001 | ✓ 1256ms | 否 | ✓ 943ms | ✓ 1295ms | 否 | http |
| 38.145.203.39:8445 | ✓ 734ms | ✓ 1060ms | ✓ 379ms | ✓ 1017ms | ✓ 1145ms | http |
| 38.145.218.9:8445 | ✓ 441ms | ✓ 1480ms | ✓ 603ms | ✓ 637ms | ✓ 627ms | http |
| 38.34.179.89:8449 | ✓ 581ms | ✓ 634ms | ✓ 507ms | ✓ 1131ms | ✓ 518ms | http |
| 38.34.179.213:8452 | ✓ 807ms | ✓ 617ms | ✓ 129ms | ✓ 719ms | ✓ 748ms | http |
| 38.145.220.27:8445 | ✓ 1286ms | ✓ 651ms | ✓ 133ms | ✓ 705ms | ✓ 1188ms | http |
| 38.145.218.234:8447 | ✓ 1721ms | ✓ 1228ms | ✓ 262ms | ✓ 1187ms | ✓ 1954ms | http |
| 38.145.208.177:8450 | ✓ 604ms | ✓ 1569ms | ✓ 717ms | ✓ 709ms | ✓ 526ms | http |
| 38.145.208.175:8447 | ✓ 1515ms | ✓ 1939ms | ✓ 95ms | ✓ 687ms | ✓ 547ms | http |
| 38.34.183.222:8453 | ✓ 605ms | ✓ 944ms | ✓ 638ms | ✓ 1833ms | ✓ 519ms | http |
| 217.182.195.221:30001 | ✓ 1304ms | 否 | ✓ 1874ms | 否 | ✓ 1656ms | http |
| 133.242.138.34:8100 | ✓ 1923ms | ✓ 1514ms | ✓ 1628ms | 否 | 否 | http |
| 180.125.216.109:8118 | ✓ 858ms | 否 | ✓ 957ms | 否 | ✓ 895ms | http |
| 38.145.208.174:8444 | ✓ 265ms | ✓ 620ms | ✓ 706ms | ✓ 1885ms | ✓ 519ms | http |
| 38.34.179.38:8447 | ✓ 377ms | ✓ 640ms | ✓ 149ms | ✓ 1226ms | 否 | http |
| 38.145.220.65:8446 | ✓ 705ms | ✓ 809ms | ✓ 426ms | ✓ 1379ms | ✓ 1704ms | http |
| 38.145.208.244:8446 | 否 | ✓ 1718ms | ✓ 94ms | ✓ 686ms | ✓ 579ms | http |
| 45.136.131.47:8452 | ✓ 848ms | 否 | ✓ 381ms | ✓ 737ms | ✓ 725ms | http |
| 150.249.255.91:3128 | ✓ 1062ms | 否 | ✓ 609ms | ✓ 1144ms | 否 | http |
| 45.140.147.155:1081 | ✓ 599ms | ✓ 1531ms | ✓ 574ms | ✓ 1690ms | 否 | http |
| 45.140.147.155:1082 | ✓ 645ms | ✓ 1721ms | ✓ 604ms | ✓ 1446ms | 否 | http |
| 38.34.179.106:8445 | 否 | ✓ 1214ms | ✓ 1593ms | ✓ 1313ms | ✓ 488ms | http |
| 180.103.19.117:1080 | ✓ 1194ms | ✓ 1606ms | ✓ 1704ms | 否 | 否 | http |
| 45.136.130.187:8449 | ✓ 1439ms | ✓ 767ms | ✓ 324ms | ✓ 1241ms | ✓ 1797ms | http |
| 38.145.220.182:8450 | ✓ 1188ms | ✓ 1979ms | ✓ 244ms | ✓ 964ms | ✓ 1030ms | http |
| 38.145.220.188:8450 | ✓ 1191ms | ✓ 1951ms | ✓ 113ms | ✓ 1129ms | ✓ 1066ms | http |
| 24.144.86.173:1080 | ✓ 462ms | ✓ 875ms | ✓ 1558ms | ✓ 1114ms | ✓ 934ms | http |
| 38.145.218.208:8446 | ✓ 197ms | ✓ 785ms | ✓ 791ms | ✓ 1769ms | ✓ 534ms | http |
| 38.34.179.179:8449 | 否 | 否 | ✓ 83ms | ✓ 664ms | ✓ 625ms | http |
| 38.34.179.18:8451 | ✓ 333ms | ✓ 938ms | ✓ 863ms | ✓ 1901ms | ✓ 709ms | http |
| 38.34.179.178:8444 | 否 | 否 | ✓ 1303ms | ✓ 1163ms | ✓ 670ms | http |
| 38.145.220.72:8445 | ✓ 963ms | ✓ 657ms | ✓ 147ms | ✓ 848ms | ✓ 1182ms | http |
| 38.145.220.20:8450 | ✓ 1395ms | 否 | ✓ 1891ms | 否 | ✓ 1871ms | http |
| 103.39.51.207:8080 | ✓ 1304ms | 否 | ✓ 1858ms | ✓ 1570ms | 否 | http |
| 111.79.111.126:3128 | ✓ 1417ms | ✓ 1687ms | ✓ 1557ms | 否 | 否 | http |
| 45.186.6.104:3128 | ✓ 1228ms | ✓ 1821ms | ✓ 1768ms | 否 | 否 | http |
| 27.254.99.183:8118 | ✓ 1018ms | 否 | ✓ 1209ms | ✓ 1274ms | ✓ 988ms | http |
| 38.34.179.86:8452 | ✓ 589ms | ✓ 866ms | ✓ 918ms | ✓ 1826ms | ✓ 659ms | http |
| 45.140.147.82:1081 | ✓ 588ms | ✓ 1382ms | ✓ 600ms | 否 | ✓ 963ms | http |
| 38.145.203.111:8450 | ✓ 1924ms | ✓ 685ms | ✓ 827ms | ✓ 1801ms | ✓ 1519ms | http |
| 38.34.179.65:8451 | ✓ 1990ms | ✓ 1519ms | ✓ 1008ms | 否 | ✓ 1988ms | http |
| 38.145.208.224:8445 | ✓ 449ms | ✓ 611ms | ✓ 230ms | ✓ 891ms | ✓ 683ms | http |
| 38.34.179.88:8446 | ✓ 478ms | ✓ 638ms | ✓ 97ms | ✓ 680ms | ✓ 1160ms | http |
| 45.136.130.171:8445 | ✓ 1186ms | ✓ 774ms | ✓ 345ms | 否 | ✓ 1403ms | http |
| 38.34.179.33:8445 | ✓ 830ms | ✓ 1622ms | ✓ 194ms | ✓ 915ms | ✓ 1025ms | http |
| 45.136.130.177:8448 | ✓ 738ms | 否 | ✓ 100ms | ✓ 724ms | ✓ 725ms | http |

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
