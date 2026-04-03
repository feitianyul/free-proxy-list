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

最后更新：2026-04-03 15:46:41 UTC（2026-04-03 23:46:41 UTC+8）

**代理总数：94**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 94 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 94 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 203.80.138.81:50000 | ✓ 870ms | ✓ 1029ms | ✓ 959ms | ✓ 854ms | ✓ 777ms | http |
| 218.108.131.186:17890 | ✓ 835ms | ✓ 1010ms | ✓ 801ms | ✓ 1044ms | ✓ 876ms | http |
| 147.161.210.140:8800 | ✓ 663ms | 否 | ✓ 1462ms | ✓ 882ms | ✓ 898ms | http |
| 192.3.181.90:1234 | ✓ 573ms | 否 | 否 | ✓ 1372ms | ✓ 1188ms | http |
| 159.223.71.162:8080 | ✓ 1330ms | 否 | ✓ 1100ms | ✓ 1033ms | ✓ 813ms | http |
| 159.223.71.162:443 | ✓ 1333ms | 否 | ✓ 1089ms | ✓ 1045ms | ✓ 813ms | http |
| 167.103.115.102:8800 | ✓ 1328ms | 否 | ✓ 936ms | ✓ 1515ms | ✓ 1168ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1293ms | ✓ 1584ms | ✓ 933ms | http |
| 95.213.217.168:52004 | ✓ 1417ms | 否 | ✓ 1455ms | 否 | ✓ 1914ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1819ms | ✓ 1749ms | ✓ 1376ms | ✓ 917ms | http |
| 167.103.34.108:8800 | ✓ 1075ms | 否 | ✓ 1159ms | ✓ 1293ms | ✓ 1222ms | http |
| 167.103.144.127:8800 | ✓ 1285ms | 否 | ✓ 1324ms | 否 | ✓ 1304ms | http |
| 45.167.124.52:8080 | ✓ 719ms | ✓ 1750ms | ✓ 888ms | 否 | 否 | http |
| 167.103.31.122:8800 | ✓ 1393ms | 否 | ✓ 1259ms | 否 | ✓ 1483ms | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 1453ms | ✓ 1919ms | ✓ 805ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 943ms | ✓ 1157ms | ✓ 944ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1175ms | ✓ 1170ms | ✓ 954ms | http |
| 111.227.254.12:22222 | 否 | ✓ 1480ms | 否 | ✓ 1359ms | ✓ 1651ms | http |
| 165.232.146.249:3128 | ✓ 355ms | 否 | ✓ 730ms | ✓ 892ms | 否 | http |
| 147.161.239.240:8800 | ✓ 758ms | 否 | ✓ 1554ms | ✓ 1717ms | ✓ 1625ms | http |
| 101.43.127.100:8877 | ✓ 931ms | ✓ 983ms | 否 | ✓ 1094ms | ✓ 1795ms | http |
| 111.227.254.9:22222 | ✓ 1011ms | ✓ 1963ms | 否 | ✓ 1258ms | ✓ 1295ms | http |
| 1.231.81.166:3128 | ✓ 809ms | ✓ 927ms | ✓ 1262ms | ✓ 943ms | ✓ 1061ms | http |
| 5.104.87.17:8050 | ✓ 505ms | 否 | ✓ 1202ms | ✓ 996ms | ✓ 625ms | http |
| 59.11.138.229:3128 | ✓ 658ms | ✓ 906ms | ✓ 543ms | ✓ 919ms | ✓ 720ms | http |
| 46.39.105.157:8080 | 否 | 否 | ✓ 1923ms | ✓ 1858ms | ✓ 1634ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1001ms | ✓ 1272ms | ✓ 1850ms | http |
| 35.194.4.51:3128 | 否 | 否 | ✓ 1342ms | ✓ 1003ms | ✓ 1168ms | http |
| 51.79.71.106:8080 | ✓ 816ms | 否 | 否 | ✓ 1985ms | ✓ 1398ms | http |
| 45.12.151.226:2829 | ✓ 1234ms | 否 | ✓ 1844ms | 否 | ✓ 1536ms | http |
| 150.249.255.91:3128 | 否 | ✓ 859ms | ✓ 460ms | 否 | ✓ 593ms | http |
| 72.11.151.159:6005 | ✓ 498ms | 否 | ✓ 859ms | ✓ 1361ms | ✓ 1243ms | http |
| 5.104.87.17:8051 | ✓ 973ms | 否 | ✓ 1244ms | ✓ 1335ms | ✓ 1238ms | http |
| 64.227.76.27:1080 | ✓ 907ms | 否 | ✓ 1193ms | ✓ 1627ms | 否 | http |
| 103.184.99.194:8080 | ✓ 1876ms | 否 | ✓ 1984ms | 否 | ✓ 1477ms | http |
| 62.113.119.14:8080 | 否 | ✓ 1908ms | ✓ 1501ms | 否 | ✓ 1332ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 501ms | ✓ 1138ms | ✓ 925ms | http |
| 110.77.244.74:8080 | ✓ 1458ms | 否 | ✓ 1657ms | ✓ 1565ms | 否 | http |
| 138.197.68.35:4857 | ✓ 677ms | 否 | ✓ 1892ms | ✓ 1535ms | ✓ 1126ms | http |
| 43.165.195.107:3128 | ✓ 1554ms | 否 | ✓ 1118ms | ✓ 1284ms | ✓ 927ms | http |
| 38.207.164.82:6005 | ✓ 1116ms | 否 | ✓ 1326ms | ✓ 934ms | ✓ 705ms | http |
| 34.101.184.164:3128 | ✓ 1574ms | 否 | ✓ 1320ms | ✓ 1551ms | ✓ 1235ms | http |
| 121.43.189.36:28888 | ✓ 1489ms | ✓ 1286ms | ✓ 1606ms | ✓ 1457ms | 否 | http |
| 103.124.138.149:8080 | ✓ 1615ms | 否 | ✓ 1965ms | ✓ 1405ms | ✓ 1418ms | http |
| 72.240.9.63:80 | ✓ 738ms | 否 | 否 | ✓ 1522ms | ✓ 1077ms | http |
| 62.234.206.73:3128 | ✓ 1139ms | 否 | ✓ 876ms | ✓ 1951ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1245ms | 否 | 否 | ✓ 1932ms | ✓ 1841ms | http |
| 61.171.66.158:3128 | ✓ 868ms | ✓ 1677ms | ✓ 1762ms | ✓ 1011ms | ✓ 912ms | http |
| 20.205.16.149:3128 | ✓ 903ms | ✓ 1203ms | ✓ 1711ms | ✓ 1619ms | ✓ 1937ms | http |
| 210.223.44.230:3128 | ✓ 1169ms | ✓ 938ms | ✓ 1111ms | ✓ 980ms | ✓ 808ms | http |
| 8.219.97.248:80 | ✓ 1654ms | ✓ 1818ms | ✓ 1827ms | 否 | 否 | http |
| 113.59.113.4:1088 | ✓ 1199ms | 否 | ✓ 1390ms | ✓ 1256ms | 否 | http |
| 185.114.73.2:1080 | ✓ 581ms | 否 | 否 | ✓ 1688ms | ✓ 1360ms | http |
| 38.34.183.219:8446 | ✓ 1535ms | 否 | ✓ 1646ms | ✓ 666ms | ✓ 971ms | http |
| 192.73.243.65:3128 | ✓ 510ms | 否 | ✓ 352ms | ✓ 1158ms | ✓ 1148ms | http |
| 128.199.254.13:9090 | ✓ 1823ms | 否 | 否 | ✓ 1170ms | ✓ 1039ms | http |
| 116.171.106.15:3443 | ✓ 1419ms | ✓ 1578ms | ✓ 1567ms | 否 | 否 | http |
| 128.199.116.219:9090 | ✓ 832ms | 否 | ✓ 1206ms | ✓ 1060ms | ✓ 931ms | http |
| 61.76.95.217:40088 | ✓ 1071ms | ✓ 1376ms | ✓ 1124ms | ✓ 1618ms | ✓ 1460ms | http |
| 103.82.23.118:5247 | ✓ 1042ms | 否 | ✓ 1003ms | ✓ 1579ms | ✓ 1305ms | http |
| 180.103.19.151:1080 | ✓ 863ms | ✓ 1514ms | ✓ 1528ms | ✓ 1345ms | ✓ 1794ms | http |
| 128.199.121.61:9090 | 否 | 否 | ✓ 1099ms | ✓ 1676ms | ✓ 1475ms | http |
| 86.53.183.16:1080 | ✓ 883ms | ✓ 1734ms | ✓ 1675ms | 否 | 否 | http |
| 103.235.67.190:80 | ✓ 1526ms | 否 | ✓ 1894ms | ✓ 1435ms | ✓ 1074ms | http |
| 103.52.114.95:3128 | ✓ 1524ms | 否 | ✓ 1243ms | ✓ 1970ms | ✓ 1585ms | http |
| 180.103.19.81:1080 | ✓ 1541ms | ✓ 1866ms | ✓ 1048ms | 否 | ✓ 1199ms | http |
| 103.22.99.43:8085 | ✓ 1535ms | 否 | ✓ 1991ms | ✓ 1880ms | ✓ 1449ms | http |
| 38.34.179.102:8444 | ✓ 559ms | ✓ 663ms | ✓ 291ms | ✓ 753ms | ✓ 1585ms | http |
| 38.145.208.215:8444 | ✓ 1032ms | ✓ 848ms | ✓ 1465ms | 否 | ✓ 515ms | http |
| 38.34.179.178:8444 | ✓ 891ms | ✓ 1854ms | ✓ 624ms | ✓ 1180ms | ✓ 1247ms | http |
| 38.34.179.105:8449 | ✓ 1360ms | 否 | ✓ 147ms | ✓ 890ms | ✓ 1350ms | http |
| 38.34.179.95:8444 | ✓ 559ms | ✓ 1395ms | ✓ 507ms | 否 | ✓ 509ms | http |
| 38.34.183.164:8444 | ✓ 1906ms | ✓ 701ms | ✓ 470ms | ✓ 1390ms | 否 | http |
| 38.34.179.78:8448 | ✓ 166ms | ✓ 682ms | ✓ 1888ms | ✓ 1789ms | ✓ 664ms | http |
| 203.207.56.57:8080 | ✓ 1401ms | 否 | ✓ 1525ms | ✓ 1317ms | ✓ 1320ms | http |
| 38.145.203.39:8445 | ✓ 1078ms | ✓ 1640ms | 否 | ✓ 1098ms | ✓ 596ms | http |
| 38.34.179.89:8449 | ✓ 707ms | ✓ 1536ms | ✓ 176ms | ✓ 692ms | 否 | http |
| 38.145.220.8:8452 | ✓ 660ms | 否 | ✓ 889ms | ✓ 1718ms | ✓ 1698ms | http |
| 38.145.208.214:8446 | ✓ 337ms | ✓ 1887ms | ✓ 111ms | ✓ 738ms | ✓ 1248ms | http |
| 38.145.208.210:8448 | ✓ 367ms | ✓ 1688ms | ✓ 1732ms | ✓ 732ms | ✓ 1529ms | http |
| 38.145.203.35:8450 | ✓ 1206ms | ✓ 1419ms | ✓ 119ms | ✓ 901ms | ✓ 758ms | http |
| 38.34.183.221:8452 | ✓ 1623ms | ✓ 1706ms | ✓ 926ms | 否 | ✓ 623ms | http |
| 20.210.76.175:8561 | 否 | 否 | ✓ 534ms | ✓ 1071ms | ✓ 1360ms | http |
| 106.117.208.101:7890 | 否 | 否 | ✓ 1170ms | ✓ 1611ms | ✓ 1682ms | http |
| 38.34.179.61:8445 | 否 | 否 | ✓ 214ms | ✓ 1028ms | ✓ 1215ms | http |
| 45.136.130.181:8445 | ✓ 1600ms | ✓ 890ms | ✓ 441ms | ✓ 1545ms | ✓ 1344ms | http |
| 38.34.179.98:8451 | 否 | 否 | ✓ 346ms | ✓ 1180ms | ✓ 1692ms | http |
| 38.145.220.13:8450 | ✓ 886ms | 否 | ✓ 1622ms | ✓ 683ms | ✓ 749ms | http |
| 38.145.220.93:8445 | ✓ 1399ms | ✓ 1695ms | 否 | ✓ 820ms | ✓ 806ms | http |
| 38.145.220.72:8445 | 否 | ✓ 953ms | ✓ 1745ms | 否 | ✓ 954ms | http |
| 38.34.179.150:8449 | 否 | ✓ 774ms | ✓ 1066ms | ✓ 1776ms | ✓ 996ms | http |
| 45.136.130.188:8449 | 否 | 否 | ✓ 1810ms | ✓ 1362ms | ✓ 572ms | http |
| 110.235.136.71:8081 | 否 | 否 | ✓ 1166ms | ✓ 1210ms | ✓ 1222ms | http |
| 38.34.179.79:8451 | ✓ 725ms | ✓ 782ms | 否 | ✓ 1315ms | ✓ 1161ms | http |

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
