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

最后更新：2026-04-01 14:40:09 UTC（2026-04-01 22:40:09 UTC+8）

**代理总数：134**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 134 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 134 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | 否 | 否 | ✓ 996ms | ✓ 1355ms | ✓ 878ms | http |
| 208.87.243.199:7878 | ✓ 541ms | ✓ 1343ms | ✓ 1213ms | ✓ 1286ms | ✓ 896ms | http |
| 147.161.239.240:8800 | ✓ 618ms | 否 | ✓ 669ms | ✓ 1342ms | 否 | http |
| 147.161.210.140:8800 | ✓ 866ms | 否 | ✓ 1198ms | 否 | ✓ 1173ms | http |
| 203.80.138.81:50000 | 否 | 否 | ✓ 1179ms | ✓ 1117ms | ✓ 1045ms | http |
| 95.213.217.168:52004 | ✓ 891ms | 否 | ✓ 1956ms | 否 | ✓ 1598ms | http |
| 185.191.236.162:3128 | ✓ 637ms | ✓ 1752ms | ✓ 1570ms | ✓ 1977ms | ✓ 1512ms | http |
| 167.103.115.102:8800 | ✓ 1555ms | 否 | ✓ 1498ms | ✓ 1400ms | ✓ 1808ms | http |
| 167.103.34.108:8800 | ✓ 1598ms | 否 | ✓ 1695ms | ✓ 1823ms | ✓ 1483ms | http |
| 1.231.81.166:3128 | ✓ 813ms | ✓ 1106ms | ✓ 865ms | ✓ 1111ms | ✓ 888ms | http |
| 163.5.180.103:56297 | ✓ 430ms | ✓ 1433ms | ✓ 1452ms | ✓ 1667ms | ✓ 1306ms | http |
| 38.34.179.27:8451 | ✓ 560ms | 否 | ✓ 1374ms | ✓ 1036ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1142ms | 否 | ✓ 1490ms | 否 | ✓ 1858ms | http |
| 167.103.144.127:8800 | ✓ 1427ms | 否 | ✓ 1522ms | 否 | ✓ 1709ms | http |
| 113.160.132.26:8080 | ✓ 1119ms | 否 | 否 | ✓ 1612ms | ✓ 1197ms | http |
| 45.167.124.52:8080 | ✓ 520ms | 否 | ✓ 502ms | ✓ 1547ms | ✓ 1272ms | http |
| 38.34.179.54:8446 | ✓ 903ms | ✓ 1963ms | ✓ 534ms | ✓ 968ms | ✓ 767ms | http |
| 45.12.151.226:2829 | ✓ 670ms | ✓ 1842ms | ✓ 660ms | ✓ 1588ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1604ms | 否 | ✓ 1545ms | 否 | ✓ 1491ms | http |
| 192.71.213.85:9090 | ✓ 1172ms | 否 | ✓ 1613ms | ✓ 1918ms | 否 | http |
| 160.250.5.22:1 | ✓ 1923ms | 否 | ✓ 1425ms | ✓ 1620ms | ✓ 1176ms | http |
| 101.43.127.100:8877 | ✓ 1385ms | ✓ 1322ms | ✓ 1335ms | ✓ 1314ms | ✓ 1817ms | http |
| 181.78.44.63:999 | ✓ 1063ms | ✓ 1321ms | ✓ 990ms | ✓ 1376ms | ✓ 1183ms | http |
| 177.234.217.88:999 | ✓ 1376ms | 否 | ✓ 1703ms | 否 | ✓ 1783ms | http |
| 31.192.106.135:8010 | ✓ 1087ms | 否 | ✓ 1083ms | 否 | ✓ 1607ms | http |
| 62.113.119.14:8080 | ✓ 585ms | ✓ 1723ms | ✓ 574ms | ✓ 1676ms | 否 | http |
| 38.145.220.8:8451 | 否 | ✓ 1076ms | ✓ 981ms | ✓ 981ms | ✓ 752ms | http |
| 38.145.208.222:8446 | 否 | ✓ 1927ms | ✓ 399ms | ✓ 1341ms | ✓ 988ms | http |
| 38.145.218.238:8451 | 否 | ✓ 1000ms | ✓ 1082ms | ✓ 1299ms | ✓ 1301ms | http |
| 38.145.208.240:8448 | 否 | ✓ 999ms | ✓ 515ms | ✓ 1081ms | ✓ 1768ms | http |
| 45.136.130.248:8452 | 否 | 否 | ✓ 1241ms | ✓ 1281ms | ✓ 954ms | http |
| 45.136.131.41:8451 | 否 | 否 | ✓ 476ms | ✓ 1686ms | ✓ 903ms | http |
| 45.136.131.38:8445 | 否 | 否 | ✓ 472ms | ✓ 980ms | ✓ 1648ms | http |
| 38.145.220.33:8448 | 否 | ✓ 1987ms | ✓ 1019ms | ✓ 1011ms | ✓ 1405ms | http |
| 45.136.130.186:8451 | 否 | 否 | ✓ 783ms | ✓ 1872ms | ✓ 1457ms | http |
| 160.250.4.245:1 | ✓ 1721ms | 否 | 否 | ✓ 1713ms | ✓ 1378ms | http |
| 38.145.220.188:8444 | ✓ 515ms | 否 | ✓ 474ms | 否 | ✓ 1567ms | http |
| 34.101.184.164:3128 | ✓ 1670ms | 否 | ✓ 1588ms | ✓ 1502ms | ✓ 1751ms | http |
| 212.58.132.5:8888 | ✓ 1790ms | 否 | ✓ 1742ms | ✓ 1508ms | ✓ 1269ms | http |
| 167.160.191.204:6005 | ✓ 1890ms | 否 | 否 | ✓ 1633ms | ✓ 1830ms | http |
| 34.96.238.40:8080 | ✓ 1783ms | ✓ 1362ms | ✓ 1263ms | ✓ 1270ms | ✓ 1646ms | http |
| 120.92.212.16:8890 | ✓ 1205ms | 否 | ✓ 1481ms | 否 | ✓ 1187ms | http |
| 38.34.179.88:8446 | ✓ 516ms | ✓ 1048ms | ✓ 355ms | ✓ 984ms | ✓ 766ms | http |
| 38.34.179.86:8452 | ✓ 787ms | ✓ 1550ms | ✓ 338ms | ✓ 1035ms | ✓ 1174ms | http |
| 38.145.208.242:8451 | ✓ 828ms | ✓ 1979ms | ✓ 663ms | ✓ 984ms | ✓ 743ms | http |
| 38.145.208.246:8450 | ✓ 828ms | ✓ 1015ms | ✓ 878ms | ✓ 1793ms | ✓ 972ms | http |
| 38.34.179.162:8451 | ✓ 859ms | ✓ 981ms | ✓ 1021ms | ✓ 1897ms | ✓ 749ms | http |
| 38.145.208.213:8450 | ✓ 846ms | ✓ 1516ms | ✓ 1007ms | ✓ 1208ms | 否 | http |
| 38.34.183.130:8452 | ✓ 1303ms | ✓ 1802ms | ✓ 391ms | ✓ 1289ms | 否 | http |
| 38.34.179.150:8449 | ✓ 678ms | 否 | ✓ 738ms | ✓ 1582ms | ✓ 766ms | http |
| 42.96.16.158:1311 | ✓ 1812ms | 否 | ✓ 1277ms | ✓ 1494ms | ✓ 1140ms | http |
| 45.136.130.189:8451 | ✓ 1185ms | 否 | ✓ 1005ms | 否 | ✓ 1993ms | http |
| 20.27.15.49:8561 | ✓ 1046ms | ✓ 1479ms | ✓ 911ms | ✓ 1079ms | 否 | http |
| 38.145.203.97:8448 | ✓ 1131ms | ✓ 1169ms | ✓ 321ms | ✓ 1213ms | ✓ 729ms | http |
| 45.136.130.177:8448 | ✓ 873ms | 否 | ✓ 470ms | ✓ 975ms | ✓ 1184ms | http |
| 45.136.130.168:8448 | ✓ 1118ms | ✓ 889ms | ✓ 812ms | ✓ 1407ms | ✓ 1380ms | http |
| 38.145.208.210:8448 | ✓ 873ms | ✓ 1336ms | ✓ 414ms | ✓ 1001ms | ✓ 1574ms | http |
| 38.145.218.232:8448 | ✓ 868ms | ✓ 1139ms | ✓ 1226ms | ✓ 1887ms | ✓ 726ms | http |
| 38.145.208.172:8448 | ✓ 875ms | ✓ 902ms | ✓ 693ms | ✓ 1634ms | ✓ 1113ms | http |
| 45.136.130.173:8448 | ✓ 1120ms | ✓ 1700ms | ✓ 673ms | ✓ 1985ms | ✓ 750ms | http |
| 38.34.179.66:8444 | ✓ 873ms | ✓ 1408ms | ✓ 899ms | ✓ 963ms | ✓ 1165ms | http |
| 120.92.212.16:7890 | ✓ 1140ms | ✓ 1451ms | ✓ 1501ms | ✓ 1476ms | ✓ 1216ms | http |
| 38.34.183.222:8453 | ✓ 1907ms | ✓ 1973ms | 否 | ✓ 1375ms | 否 | http |
| 38.145.220.9:8448 | ✓ 1886ms | 否 | ✓ 1877ms | ✓ 1762ms | 否 | http |
| 38.34.183.233:8448 | ✓ 1885ms | 否 | ✓ 1893ms | ✓ 1847ms | 否 | http |
| 38.34.179.27:8452 | ✓ 325ms | ✓ 970ms | ✓ 546ms | ✓ 944ms | ✓ 753ms | http |
| 38.34.179.84:8453 | ✓ 314ms | ✓ 996ms | ✓ 535ms | ✓ 981ms | ✓ 823ms | http |
| 38.34.179.135:8446 | ✓ 389ms | ✓ 1012ms | ✓ 447ms | ✓ 976ms | ✓ 824ms | http |
| 38.145.208.182:8450 | ✓ 396ms | ✓ 1015ms | ✓ 436ms | ✓ 1073ms | ✓ 976ms | http |
| 45.136.130.180:8453 | ✓ 304ms | ✓ 1416ms | ✓ 311ms | ✓ 1056ms | ✓ 719ms | http |
| 38.34.179.70:8453 | ✓ 318ms | ✓ 959ms | ✓ 568ms | ✓ 1165ms | ✓ 1058ms | http |
| 38.145.218.208:8453 | ✓ 294ms | ✓ 1536ms | ✓ 299ms | ✓ 975ms | ✓ 1174ms | http |
| 38.145.203.106:8448 | ✓ 424ms | ✓ 865ms | ✓ 552ms | ✓ 1189ms | ✓ 716ms | http |
| 38.34.179.193:8452 | ✓ 374ms | ✓ 1803ms | ✓ 332ms | ✓ 982ms | ✓ 798ms | http |
| 38.34.179.179:8449 | ✓ 393ms | ✓ 1232ms | ✓ 317ms | ✓ 991ms | ✓ 896ms | http |
| 38.145.208.227:8451 | ✓ 624ms | ✓ 1993ms | ✓ 427ms | ✓ 974ms | ✓ 863ms | http |
| 38.145.208.209:8444 | ✓ 383ms | 否 | ✓ 309ms | ✓ 928ms | ✓ 722ms | http |
| 38.34.179.18:8444 | ✓ 507ms | ✓ 1150ms | ✓ 1712ms | ✓ 1129ms | ✓ 778ms | http |
| 38.145.218.229:8450 | ✓ 836ms | ✓ 906ms | ✓ 391ms | ✓ 1343ms | ✓ 1966ms | http |
| 38.145.203.132:8450 | ✓ 731ms | ✓ 1463ms | ✓ 848ms | ✓ 1213ms | ✓ 1220ms | http |
| 38.145.218.210:8451 | ✓ 492ms | ✓ 1129ms | ✓ 821ms | ✓ 918ms | ✓ 759ms | http |
| 38.34.183.224:8451 | ✓ 510ms | 否 | ✓ 311ms | ✓ 1020ms | ✓ 829ms | http |
| 38.145.208.213:8449 | ✓ 299ms | ✓ 1020ms | ✓ 680ms | ✓ 1128ms | ✓ 721ms | http |
| 38.145.208.206:8448 | ✓ 583ms | ✓ 1219ms | ✓ 591ms | ✓ 974ms | ✓ 826ms | http |
| 38.145.220.65:8449 | ✓ 570ms | ✓ 1589ms | ✓ 569ms | ✓ 928ms | ✓ 749ms | http |
| 38.145.208.206:8453 | ✓ 617ms | 否 | ✓ 304ms | ✓ 954ms | ✓ 1438ms | http |
| 38.145.220.43:8449 | ✓ 536ms | 否 | ✓ 270ms | ✓ 970ms | ✓ 724ms | http |
| 38.145.203.107:8446 | ✓ 749ms | ✓ 1140ms | ✓ 866ms | ✓ 927ms | ✓ 797ms | http |
| 38.34.179.155:8448 | ✓ 416ms | ✓ 986ms | ✓ 899ms | ✓ 1830ms | ✓ 767ms | http |
| 38.145.208.179:8447 | ✓ 557ms | 否 | ✓ 1053ms | ✓ 1041ms | 否 | http |
| 38.34.179.95:8444 | 否 | 否 | ✓ 483ms | ✓ 982ms | ✓ 838ms | http |
| 38.34.179.61:8444 | 否 | 否 | ✓ 490ms | ✓ 935ms | ✓ 858ms | http |
| 38.34.179.68:8452 | ✓ 341ms | ✓ 1245ms | ✓ 1847ms | ✓ 1160ms | ✓ 1220ms | http |
| 195.123.209.48:3128 | ✓ 965ms | ✓ 1711ms | ✓ 1373ms | 否 | ✓ 1737ms | http |
| 38.34.179.8:8449 | ✓ 1422ms | 否 | ✓ 1265ms | ✓ 1555ms | ✓ 711ms | http |
| 47.101.159.19:8899 | ✓ 1350ms | ✓ 1366ms | ✓ 1809ms | 否 | ✓ 1337ms | http |
| 20.210.76.175:8561 | ✓ 1438ms | ✓ 1210ms | ✓ 676ms | ✓ 1085ms | ✓ 905ms | http |
| 20.210.76.104:8561 | ✓ 1430ms | ✓ 1136ms | ✓ 751ms | ✓ 1093ms | ✓ 896ms | http |
| 20.210.76.178:8561 | ✓ 1440ms | ✓ 1467ms | ✓ 762ms | ✓ 1042ms | ✓ 870ms | http |
| 38.34.183.47:8452 | ✓ 1720ms | ✓ 1308ms | ✓ 1381ms | ✓ 1436ms | ✓ 706ms | http |

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
