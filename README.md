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

最后更新：2026-03-29 07:59:52 UTC（2026-03-29 15:59:52 UTC+8）

**代理总数：91**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 946ms | ✓ 790ms | ✓ 786ms | ✓ 1967ms | ✓ 702ms | http |
| 147.161.210.140:8800 | 否 | ✓ 1452ms | ✓ 987ms | ✓ 1162ms | ✓ 1359ms | http |
| 147.161.239.240:8800 | ✓ 1150ms | ✓ 1776ms | ✓ 1536ms | ✓ 1696ms | ✓ 1576ms | http |
| 113.160.132.26:8080 | ✓ 1523ms | 否 | ✓ 1085ms | ✓ 1447ms | ✓ 986ms | http |
| 42.96.16.158:1311 | ✓ 1541ms | 否 | ✓ 1126ms | 否 | ✓ 1063ms | http |
| 167.103.115.102:8800 | ✓ 1393ms | 否 | ✓ 1225ms | 否 | ✓ 1339ms | http |
| 45.136.198.40:3128 | ✓ 1178ms | ✓ 1745ms | ✓ 1646ms | 否 | 否 | http |
| 167.103.34.108:8800 | ✓ 1417ms | 否 | ✓ 1427ms | 否 | ✓ 1639ms | http |
| 115.231.181.40:8128 | ✓ 1928ms | ✓ 1119ms | 否 | ✓ 1596ms | ✓ 1366ms | http |
| 45.167.124.52:8080 | ✓ 1242ms | 否 | ✓ 1282ms | ✓ 1958ms | ✓ 1529ms | http |
| 180.250.219.58:53281 | ✓ 1842ms | 否 | ✓ 1804ms | ✓ 1998ms | 否 | http |
| 38.145.220.55:8443 | ✓ 533ms | ✓ 852ms | ✓ 684ms | ✓ 849ms | ✓ 605ms | http |
| 38.34.179.106:8450 | ✓ 513ms | ✓ 1191ms | ✓ 314ms | ✓ 898ms | ✓ 1068ms | http |
| 45.136.130.169:8444 | ✓ 454ms | ✓ 950ms | ✓ 575ms | ✓ 1461ms | ✓ 734ms | http |
| 43.99.54.236:5555 | ✓ 749ms | 否 | ✓ 664ms | ✓ 858ms | ✓ 693ms | http |
| 183.249.5.117:22222 | ✓ 928ms | ✓ 1353ms | 否 | ✓ 1176ms | ✓ 992ms | http |
| 95.213.217.168:52004 | ✓ 654ms | ✓ 1699ms | ✓ 759ms | ✓ 1674ms | ✓ 1260ms | http |
| 103.84.95.54:7890 | ✓ 1277ms | 否 | ✓ 731ms | ✓ 1236ms | ✓ 839ms | http |
| 45.136.130.186:8451 | ✓ 1322ms | 否 | 否 | ✓ 929ms | ✓ 853ms | http |
| 222.184.48.242:22222 | ✓ 919ms | ✓ 1217ms | 否 | ✓ 1985ms | ✓ 915ms | http |
| 167.103.144.127:8800 | ✓ 1495ms | 否 | ✓ 1497ms | 否 | ✓ 1544ms | http |
| 103.139.138.194:3128 | ✓ 1787ms | 否 | ✓ 1184ms | ✓ 1434ms | ✓ 1384ms | http |
| 45.136.130.166:8447 | ✓ 811ms | 否 | ✓ 1079ms | ✓ 1877ms | 否 | http |
| 45.136.131.35:8452 | 否 | 否 | ✓ 1844ms | ✓ 1381ms | ✓ 1296ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 528ms | ✓ 1073ms | ✓ 1140ms | http |
| 210.223.44.230:3128 | ✓ 1536ms | ✓ 1086ms | ✓ 802ms | ✓ 1758ms | ✓ 1482ms | http |
| 167.103.31.122:8800 | ✓ 1824ms | 否 | ✓ 1463ms | ✓ 1692ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1073ms | 否 | ✓ 1580ms | 否 | ✓ 1056ms | http |
| 8.219.97.248:80 | ✓ 1917ms | 否 | ✓ 1478ms | ✓ 1732ms | 否 | http |
| 222.184.48.251:22222 | ✓ 1319ms | ✓ 1613ms | 否 | 否 | ✓ 1545ms | http |
| 1.231.81.166:3128 | ✓ 1832ms | ✓ 1183ms | ✓ 1472ms | ✓ 1088ms | ✓ 916ms | http |
| 168.110.52.228:3128 | ✓ 1805ms | 否 | ✓ 1616ms | ✓ 1078ms | ✓ 829ms | http |
| 62.113.119.14:8080 | ✓ 732ms | ✓ 1829ms | ✓ 1259ms | ✓ 1653ms | ✓ 1272ms | http |
| 89.208.106.138:10808 | ✓ 975ms | ✓ 1743ms | ✓ 1379ms | ✓ 1722ms | ✓ 1432ms | http |
| 222.184.48.252:22222 | 否 | 否 | ✓ 1875ms | ✓ 1608ms | ✓ 853ms | http |
| 101.43.127.100:8877 | ✓ 1350ms | ✓ 1944ms | 否 | ✓ 1710ms | ✓ 921ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1231ms | ✓ 1549ms | 否 | ✓ 1868ms | http |
| 177.234.217.88:999 | ✓ 1750ms | 否 | ✓ 1884ms | 否 | ✓ 1812ms | http |
| 45.12.151.226:2829 | ✓ 1442ms | ✓ 1843ms | ✓ 1735ms | 否 | ✓ 1655ms | http |
| 45.136.131.28:8449 | ✓ 848ms | ✓ 737ms | ✓ 388ms | ✓ 793ms | ✓ 609ms | http |
| 45.136.131.61:8451 | ✓ 965ms | ✓ 1224ms | ✓ 336ms | ✓ 871ms | ✓ 616ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1100ms | ✓ 1378ms | ✓ 1072ms | http |
| 38.34.179.80:8453 | 否 | ✓ 1417ms | ✓ 843ms | 否 | ✓ 671ms | http |
| 120.92.212.16:8890 | ✓ 1230ms | 否 | 否 | ✓ 1268ms | ✓ 1273ms | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1322ms | ✓ 1420ms | ✓ 1167ms | http |
| 101.32.244.83:8080 | ✓ 1413ms | 否 | ✓ 1228ms | ✓ 1385ms | ✓ 1276ms | http |
| 121.43.196.210:8222 | ✓ 1002ms | ✓ 1046ms | ✓ 872ms | ✓ 1144ms | ✓ 911ms | http |
| 121.43.196.213:8222 | ✓ 1986ms | ✓ 1074ms | ✓ 815ms | ✓ 1111ms | ✓ 875ms | http |
| 114.55.226.123:10086 | ✓ 1155ms | ✓ 1681ms | ✓ 1007ms | ✓ 1249ms | ✓ 1051ms | http |
| 219.117.204.211:7799 | ✓ 1482ms | ✓ 1453ms | ✓ 1054ms | ✓ 996ms | ✓ 974ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1151ms | ✓ 971ms | ✓ 1478ms | 否 | http |
| 45.136.130.250:8450 | ✓ 1580ms | 否 | 否 | ✓ 1622ms | ✓ 1708ms | http |
| 38.145.218.233:8444 | 否 | ✓ 827ms | ✓ 533ms | ✓ 826ms | ✓ 582ms | http |
| 45.136.131.49:8444 | 否 | ✓ 814ms | ✓ 761ms | ✓ 810ms | ✓ 716ms | http |
| 45.136.131.54:8446 | 否 | ✓ 757ms | ✓ 819ms | ✓ 818ms | ✓ 720ms | http |
| 45.136.131.43:8452 | 否 | ✓ 978ms | ✓ 192ms | ✓ 874ms | ✓ 1435ms | http |
| 38.145.220.32:8450 | 否 | ✓ 889ms | ✓ 881ms | ✓ 948ms | ✓ 772ms | http |
| 45.136.130.247:8450 | 否 | ✓ 729ms | ✓ 151ms | ✓ 864ms | ✓ 828ms | http |
| 38.145.218.210:8451 | 否 | ✓ 1370ms | ✓ 687ms | ✓ 1016ms | ✓ 609ms | http |
| 38.145.203.111:8449 | 否 | ✓ 997ms | ✓ 705ms | ✓ 1364ms | ✓ 801ms | http |
| 38.145.203.106:8448 | 否 | ✓ 893ms | ✓ 819ms | ✓ 1530ms | ✓ 630ms | http |
| 38.34.179.151:8452 | 否 | ✓ 878ms | ✓ 801ms | ✓ 1080ms | ✓ 1423ms | http |
| 38.145.220.100:8450 | 否 | ✓ 941ms | ✓ 1113ms | ✓ 1197ms | ✓ 626ms | http |
| 38.145.218.208:8449 | 否 | ✓ 1579ms | ✓ 286ms | ✓ 821ms | ✓ 623ms | http |
| 38.145.218.234:8447 | 否 | ✓ 1711ms | ✓ 189ms | ✓ 807ms | ✓ 600ms | http |
| 38.145.218.218:8451 | 否 | ✓ 1493ms | ✓ 175ms | ✓ 837ms | ✓ 827ms | http |
| 45.136.130.248:8451 | 否 | ✓ 1260ms | ✓ 794ms | 否 | ✓ 569ms | http |
| 38.145.203.107:8453 | 否 | ✓ 930ms | ✓ 1193ms | ✓ 1314ms | ✓ 630ms | http |
| 38.34.183.222:8444 | 否 | ✓ 1261ms | ✓ 841ms | ✓ 776ms | ✓ 720ms | http |
| 38.34.179.13:8445 | 否 | ✓ 1504ms | ✓ 368ms | ✓ 759ms | ✓ 811ms | http |
| 38.34.183.211:8445 | 否 | 否 | ✓ 510ms | ✓ 840ms | ✓ 686ms | http |
| 38.145.208.253:8445 | 否 | ✓ 1784ms | ✓ 555ms | ✓ 1563ms | ✓ 1731ms | http |
| 38.145.220.96:8445 | 否 | ✓ 899ms | ✓ 1204ms | ✓ 1169ms | 否 | http |
| 38.145.203.108:8453 | 否 | ✓ 921ms | ✓ 1161ms | ✓ 1380ms | ✓ 790ms | http |
| 45.136.131.38:8445 | 否 | 否 | ✓ 1070ms | ✓ 1248ms | ✓ 862ms | http |
| 38.145.203.32:8452 | 否 | 否 | ✓ 1530ms | ✓ 1953ms | ✓ 601ms | http |
| 128.199.116.219:9090 | ✓ 1196ms | 否 | ✓ 1349ms | ✓ 1643ms | 否 | http |
| 38.145.218.229:8450 | 否 | 否 | ✓ 243ms | ✓ 1046ms | ✓ 877ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1401ms | ✓ 1870ms | ✓ 1151ms | http |
| 193.233.22.29:10808 | 否 | 否 | ✓ 1173ms | ✓ 1586ms | ✓ 1267ms | http |
| 114.237.77.231:1080 | ✓ 1017ms | ✓ 1158ms | ✓ 1991ms | ✓ 1285ms | ✓ 1982ms | http |
| 106.75.15.167:7890 | ✓ 1376ms | ✓ 1172ms | 否 | ✓ 1254ms | ✓ 919ms | http |
| 20.78.213.56:80 | ✓ 1488ms | ✓ 1222ms | ✓ 1888ms | ✓ 1547ms | ✓ 1554ms | http |
| 45.136.131.32:8445 | ✓ 1313ms | 否 | ✓ 975ms | 否 | ✓ 1904ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1504ms | ✓ 1702ms | ✓ 1443ms | http |
| 45.144.232.5:11741 | ✓ 1606ms | 否 | 否 | ✓ 1894ms | ✓ 1562ms | http |
| 82.146.58.184:1080 | ✓ 721ms | ✓ 1879ms | ✓ 835ms | 否 | ✓ 1785ms | http |
| 61.52.131.172:8443 | ✓ 1615ms | ✓ 1906ms | ✓ 1591ms | 否 | ✓ 1541ms | http |
| 183.249.5.105:22222 | ✓ 1092ms | ✓ 1099ms | ✓ 994ms | ✓ 962ms | ✓ 772ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1325ms | ✓ 924ms | 否 | ✓ 1320ms | http |
| 45.129.141.143:3128 | ✓ 1312ms | ✓ 1858ms | 否 | 否 | ✓ 1733ms | http |

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
