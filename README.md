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

最后更新：2026-04-01 10:08:14 UTC（2026-04-01 18:08:14 UTC+8）

**代理总数：101**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 101 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 101 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | 否 | 否 | ✓ 1096ms | ✓ 1384ms | ✓ 1193ms | http |
| 1.231.81.166:3128 | ✓ 1601ms | ✓ 1023ms | ✓ 1574ms | ✓ 1027ms | ✓ 796ms | http |
| 167.103.115.102:8800 | ✓ 1400ms | 否 | ✓ 943ms | ✓ 1025ms | ✓ 990ms | http |
| 147.161.210.140:8800 | ✓ 1565ms | 否 | ✓ 914ms | ✓ 1366ms | ✓ 1328ms | http |
| 147.161.239.240:8800 | ✓ 1155ms | ✓ 1804ms | ✓ 1438ms | ✓ 1521ms | ✓ 1543ms | http |
| 95.213.217.168:52004 | ✓ 1239ms | ✓ 1598ms | ✓ 1672ms | ✓ 1982ms | ✓ 1757ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1096ms | ✓ 1759ms | ✓ 1081ms | http |
| 167.103.34.108:8800 | ✓ 1683ms | 否 | ✓ 1473ms | ✓ 1602ms | ✓ 1659ms | http |
| 180.250.219.58:53281 | ✓ 1838ms | 否 | ✓ 1800ms | ✓ 1975ms | ✓ 1935ms | http |
| 45.136.198.40:3128 | ✓ 797ms | ✓ 1717ms | ✓ 771ms | ✓ 1678ms | ✓ 1262ms | http |
| 167.103.144.127:8800 | ✓ 1378ms | 否 | ✓ 1094ms | 否 | ✓ 1105ms | http |
| 167.103.31.122:8800 | ✓ 1657ms | 否 | ✓ 1512ms | ✓ 1992ms | ✓ 1936ms | http |
| 42.96.16.158:1311 | ✓ 1450ms | 否 | ✓ 1080ms | ✓ 1247ms | 否 | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1307ms | ✓ 1050ms | ✓ 1060ms | http |
| 121.230.8.213:1080 | ✓ 1363ms | ✓ 1501ms | ✓ 1141ms | 否 | ✓ 1170ms | http |
| 120.92.212.16:7890 | ✓ 1037ms | ✓ 1721ms | ✓ 1226ms | 否 | 否 | http |
| 203.80.138.81:50000 | ✓ 887ms | ✓ 1060ms | ✓ 839ms | ✓ 1001ms | ✓ 781ms | http |
| 208.87.243.199:7878 | ✓ 698ms | 否 | 否 | ✓ 833ms | ✓ 629ms | http |
| 160.250.5.22:1 | ✓ 1530ms | 否 | ✓ 1402ms | ✓ 1343ms | ✓ 1068ms | http |
| 101.43.127.100:8877 | ✓ 1611ms | 否 | ✓ 1598ms | ✓ 1825ms | ✓ 1117ms | http |
| 181.78.44.63:999 | ✓ 913ms | ✓ 1612ms | ✓ 1233ms | ✓ 1462ms | ✓ 1360ms | http |
| 177.234.217.88:999 | ✓ 1399ms | 否 | ✓ 1749ms | 否 | ✓ 1999ms | http |
| 45.12.151.226:2829 | ✓ 870ms | 否 | ✓ 1140ms | ✓ 1810ms | 否 | http |
| 38.145.220.65:8449 | ✓ 542ms | ✓ 1817ms | ✓ 117ms | ✓ 697ms | ✓ 931ms | http |
| 120.92.212.16:8890 | ✓ 1298ms | 否 | ✓ 1497ms | 否 | ✓ 1849ms | http |
| 209.126.84.232:8888 | 否 | 否 | ✓ 1169ms | ✓ 1965ms | ✓ 1415ms | http |
| 160.250.4.245:1 | ✓ 1706ms | 否 | ✓ 1407ms | ✓ 1295ms | ✓ 1070ms | http |
| 103.113.70.189:1081 | ✓ 402ms | 否 | ✓ 1587ms | 否 | ✓ 884ms | http |
| 103.179.218.5:8080 | ✓ 1695ms | 否 | ✓ 1507ms | ✓ 1522ms | ✓ 1503ms | http |
| 38.34.179.54:8446 | ✓ 490ms | ✓ 947ms | ✓ 424ms | ✓ 1058ms | ✓ 519ms | http |
| 38.34.179.174:8453 | ✓ 1549ms | 否 | ✓ 110ms | ✓ 836ms | ✓ 525ms | http |
| 38.34.183.13:8449 | ✓ 387ms | ✓ 874ms | ✓ 503ms | ✓ 749ms | ✓ 544ms | http |
| 38.34.179.14:8450 | ✓ 1190ms | ✓ 617ms | ✓ 95ms | ✓ 851ms | ✓ 1028ms | http |
| 38.145.220.188:8444 | ✓ 112ms | ✓ 602ms | ✓ 622ms | ✓ 896ms | ✓ 1264ms | http |
| 38.34.183.47:8452 | ✓ 559ms | ✓ 648ms | ✓ 102ms | ✓ 1785ms | ✓ 1889ms | http |
| 38.145.220.33:8448 | ✓ 648ms | 否 | ✓ 1743ms | ✓ 1088ms | 否 | http |
| 20.120.225.109:3128 | ✓ 630ms | ✓ 1400ms | 否 | ✓ 1144ms | ✓ 862ms | http |
| 38.145.208.162:8448 | ✓ 854ms | ✓ 1447ms | ✓ 387ms | 否 | ✓ 687ms | http |
| 133.242.138.34:8100 | ✓ 1615ms | ✓ 1352ms | 否 | 否 | ✓ 1124ms | http |
| 212.58.132.5:8888 | ✓ 1461ms | 否 | ✓ 1273ms | 否 | ✓ 1294ms | http |
| 115.231.181.40:8128 | ✓ 1130ms | ✓ 1089ms | ✓ 1131ms | 否 | ✓ 1678ms | http |
| 24.144.86.173:1080 | 否 | 否 | ✓ 1619ms | ✓ 961ms | ✓ 1105ms | http |
| 38.145.220.102:8453 | ✓ 1834ms | ✓ 1712ms | ✓ 1045ms | ✓ 1602ms | ✓ 1269ms | http |
| 38.34.179.21:8443 | ✓ 538ms | ✓ 776ms | ✓ 597ms | ✓ 1198ms | ✓ 971ms | http |
| 38.34.179.86:8452 | ✓ 488ms | ✓ 860ms | ✓ 538ms | ✓ 1228ms | ✓ 1108ms | http |
| 38.34.179.88:8446 | ✓ 479ms | ✓ 797ms | ✓ 601ms | ✓ 1230ms | ✓ 1157ms | http |
| 167.160.191.204:6005 | ✓ 1400ms | ✓ 1736ms | ✓ 1226ms | 否 | 否 | http |
| 38.34.179.40:8446 | ✓ 1087ms | ✓ 897ms | ✓ 287ms | 否 | ✓ 1715ms | http |
| 38.34.179.54:8453 | ✓ 520ms | ✓ 1370ms | ✓ 285ms | ✓ 1176ms | ✓ 1005ms | http |
| 38.34.179.39:8452 | ✓ 1092ms | ✓ 1019ms | ✓ 413ms | ✓ 1428ms | 否 | http |
| 45.136.131.54:8448 | ✓ 269ms | 否 | ✓ 1872ms | ✓ 1337ms | ✓ 1685ms | http |
| 38.34.179.162:8451 | ✓ 1084ms | ✓ 1718ms | ✓ 307ms | 否 | ✓ 1897ms | http |
| 45.136.131.36:8450 | ✓ 1265ms | ✓ 822ms | ✓ 447ms | 否 | ✓ 1428ms | http |
| 45.140.147.155:1082 | ✓ 1352ms | 否 | ✓ 1434ms | ✓ 1951ms | 否 | http |
| 45.136.131.25:8450 | ✓ 687ms | ✓ 602ms | ✓ 81ms | ✓ 668ms | ✓ 638ms | http |
| 45.136.130.177:8448 | ✓ 667ms | ✓ 656ms | ✓ 119ms | ✓ 786ms | ✓ 643ms | http |
| 45.136.130.169:8444 | ✓ 700ms | ✓ 676ms | ✓ 207ms | ✓ 831ms | ✓ 551ms | http |
| 45.136.131.28:8447 | ✓ 691ms | ✓ 1102ms | ✓ 82ms | ✓ 679ms | ✓ 531ms | http |
| 38.145.218.229:8450 | ✓ 683ms | ✓ 727ms | ✓ 282ms | ✓ 686ms | ✓ 538ms | http |
| 45.136.130.167:8444 | ✓ 696ms | ✓ 715ms | ✓ 202ms | ✓ 1323ms | ✓ 507ms | http |
| 45.136.130.186:8451 | ✓ 681ms | ✓ 613ms | ✓ 239ms | ✓ 1013ms | ✓ 571ms | http |
| 45.136.131.56:8450 | ✓ 699ms | ✓ 665ms | ✓ 376ms | ✓ 955ms | ✓ 528ms | http |
| 38.145.208.207:8445 | ✓ 706ms | ✓ 1094ms | ✓ 170ms | ✓ 689ms | ✓ 883ms | http |
| 45.136.130.166:8447 | ✓ 676ms | ✓ 1701ms | ✓ 122ms | ✓ 695ms | ✓ 518ms | http |
| 38.34.179.106:8450 | ✓ 680ms | ✓ 713ms | ✓ 96ms | ✓ 719ms | ✓ 871ms | http |
| 38.34.179.155:8448 | ✓ 676ms | ✓ 733ms | ✓ 102ms | ✓ 736ms | ✓ 1143ms | http |
| 38.34.179.74:8449 | ✓ 702ms | ✓ 1408ms | ✓ 803ms | ✓ 718ms | ✓ 566ms | http |
| 45.136.130.191:8453 | ✓ 917ms | ✓ 669ms | ✓ 203ms | ✓ 777ms | ✓ 1455ms | http |
| 45.136.130.193:8444 | ✓ 915ms | ✓ 626ms | ✓ 603ms | 否 | ✓ 1107ms | http |
| 38.34.178.153:8453 | ✓ 698ms | ✓ 1766ms | ✓ 1247ms | ✓ 1415ms | ✓ 493ms | http |
| 38.34.179.192:8450 | ✓ 1178ms | 否 | ✓ 218ms | ✓ 783ms | ✓ 890ms | http |
| 38.145.208.208:8453 | ✓ 704ms | ✓ 700ms | ✓ 661ms | ✓ 1455ms | ✓ 507ms | http |
| 38.145.208.206:8453 | ✓ 688ms | ✓ 703ms | ✓ 673ms | ✓ 1904ms | ✓ 785ms | http |
| 38.145.208.203:8451 | ✓ 702ms | ✓ 710ms | ✓ 732ms | ✓ 1700ms | ✓ 569ms | http |
| 38.145.208.211:8453 | ✓ 767ms | ✓ 683ms | ✓ 709ms | ✓ 1696ms | ✓ 1483ms | http |
| 38.145.208.212:8451 | ✓ 692ms | ✓ 740ms | ✓ 848ms | ✓ 1565ms | ✓ 591ms | http |
| 38.34.179.61:8445 | ✓ 1197ms | ✓ 684ms | ✓ 552ms | 否 | ✓ 1072ms | http |
| 217.76.245.80:999 | ✓ 796ms | 否 | ✓ 1219ms | ✓ 1767ms | ✓ 1346ms | http |
| 38.145.220.77:8447 | ✓ 932ms | 否 | ✓ 747ms | ✓ 884ms | ✓ 995ms | http |
| 165.232.146.249:3128 | ✓ 987ms | ✓ 1184ms | 否 | ✓ 1005ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1462ms | 否 | ✓ 932ms | ✓ 1326ms | ✓ 1247ms | http |
| 121.232.73.214:1080 | 否 | 否 | ✓ 871ms | ✓ 1303ms | ✓ 1138ms | http |
| 121.43.196.213:8222 | ✓ 905ms | ✓ 1091ms | ✓ 865ms | ✓ 1101ms | ✓ 916ms | http |
| 121.43.196.210:8222 | ✓ 923ms | ✓ 1029ms | ✓ 919ms | ✓ 1252ms | ✓ 922ms | http |
| 114.55.226.123:10086 | ✓ 1021ms | ✓ 1705ms | ✓ 1034ms | ✓ 1296ms | ✓ 1048ms | http |
| 211.95.152.50:45046 | ✓ 1214ms | ✓ 1379ms | ✓ 1280ms | ✓ 1424ms | ✓ 1109ms | http |
| 59.46.216.131:30001 | ✓ 1970ms | ✓ 1995ms | ✓ 998ms | ✓ 1595ms | 否 | http |
| 106.117.208.101:7890 | 否 | ✓ 1471ms | ✓ 1385ms | ✓ 1515ms | 否 | http |
| 5.102.109.41:999 | ✓ 836ms | ✓ 1277ms | ✓ 1050ms | ✓ 1398ms | ✓ 1344ms | http |
| 194.59.204.87:9080 | ✓ 1320ms | ✓ 1738ms | ✓ 1503ms | 否 | 否 | http |
| 183.98.143.134:8083 | ✓ 1102ms | ✓ 1304ms | 否 | 否 | ✓ 880ms | http |
| 210.223.44.230:3128 | ✓ 1592ms | 否 | ✓ 948ms | 否 | ✓ 1930ms | http |
| 62.113.119.14:8080 | ✓ 800ms | 否 | ✓ 901ms | ✓ 1699ms | ✓ 1277ms | http |
| 103.39.51.190:8080 | ✓ 1809ms | 否 | 否 | ✓ 1627ms | ✓ 1321ms | http |
| 217.217.249.160:8080 | ✓ 1541ms | 否 | ✓ 1617ms | 否 | ✓ 1573ms | http |
| 34.96.238.40:8080 | ✓ 899ms | 否 | 否 | ✓ 984ms | ✓ 995ms | http |
| 103.67.46.225:3125 | ✓ 1813ms | 否 | ✓ 1970ms | ✓ 1700ms | ✓ 1587ms | http |
| 61.52.131.172:8443 | ✓ 916ms | ✓ 1243ms | ✓ 1036ms | ✓ 1208ms | ✓ 926ms | http |
| 34.101.184.164:3128 | ✓ 1756ms | 否 | ✓ 1289ms | ✓ 1225ms | ✓ 1082ms | http |
| 180.191.23.171:8081 | ✓ 1649ms | 否 | ✓ 1448ms | 否 | ✓ 1389ms | http |

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
