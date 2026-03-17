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

最后更新：2026-03-17 23:28:49 UTC（2026-03-18 07:28:49 UTC+8）

**代理总数：121**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 121 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 121 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 115.231.181.40:8128 | ✓ 871ms | ✓ 1057ms | ✓ 869ms | ✓ 1121ms | ✓ 951ms | http |
| 147.161.210.140:8800 | ✓ 1421ms | 否 | ✓ 1007ms | ✓ 968ms | ✓ 847ms | http |
| 147.161.239.240:8800 | ✓ 1114ms | ✓ 1717ms | ✓ 1542ms | ✓ 1411ms | ✓ 1222ms | http |
| 1.231.81.166:3128 | ✓ 1450ms | ✓ 1087ms | ✓ 1702ms | ✓ 1565ms | ✓ 1058ms | http |
| 192.3.203.158:1080 | ✓ 1024ms | 否 | 否 | ✓ 1952ms | ✓ 1464ms | http |
| 137.220.150.152:6005 | ✓ 1039ms | 否 | ✓ 1702ms | ✓ 1841ms | ✓ 1046ms | http |
| 113.160.132.26:8080 | ✓ 1563ms | 否 | ✓ 1399ms | 否 | ✓ 1925ms | http |
| 45.167.124.52:8080 | ✓ 1615ms | 否 | ✓ 1904ms | ✓ 1667ms | ✓ 1360ms | http |
| 38.34.179.60:8450 | ✓ 146ms | ✓ 683ms | ✓ 124ms | ✓ 674ms | ✓ 624ms | http |
| 38.145.220.33:8448 | ✓ 369ms | ✓ 746ms | ✓ 117ms | ✓ 670ms | ✓ 527ms | http |
| 38.34.179.87:8447 | ✓ 141ms | ✓ 720ms | ✓ 302ms | ✓ 1096ms | ✓ 566ms | http |
| 38.34.179.83:8448 | ✓ 143ms | ✓ 653ms | ✓ 264ms | ✓ 844ms | ✓ 898ms | http |
| 38.145.203.19:8447 | ✓ 153ms | ✓ 615ms | ✓ 588ms | ✓ 856ms | ✓ 859ms | http |
| 38.34.179.156:8447 | ✓ 400ms | ✓ 682ms | ✓ 757ms | ✓ 911ms | ✓ 525ms | http |
| 38.34.183.225:8450 | ✓ 92ms | ✓ 825ms | ✓ 116ms | ✓ 707ms | ✓ 1767ms | http |
| 38.34.183.234:8450 | ✓ 148ms | ✓ 1236ms | ✓ 286ms | ✓ 1104ms | ✓ 1128ms | http |
| 38.34.179.14:8450 | ✓ 345ms | ✓ 856ms | ✓ 706ms | ✓ 1750ms | ✓ 522ms | http |
| 38.145.208.181:8445 | ✓ 222ms | ✓ 831ms | ✓ 90ms | ✓ 645ms | ✓ 583ms | http |
| 219.117.204.211:7799 | ✓ 473ms | ✓ 1193ms | ✓ 476ms | ✓ 828ms | ✓ 638ms | http |
| 45.136.131.54:8448 | ✓ 543ms | ✓ 1321ms | ✓ 837ms | ✓ 890ms | ✓ 581ms | http |
| 38.145.208.179:8447 | ✓ 252ms | ✓ 1374ms | ✓ 134ms | ✓ 716ms | ✓ 767ms | http |
| 38.34.179.26:8450 | ✓ 811ms | ✓ 1866ms | ✓ 534ms | ✓ 974ms | ✓ 1578ms | http |
| 210.223.44.230:3128 | ✓ 734ms | ✓ 879ms | ✓ 1006ms | ✓ 1208ms | ✓ 1854ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1186ms | ✓ 944ms | ✓ 1193ms | ✓ 981ms | http |
| 120.92.212.16:8890 | ✓ 1673ms | ✓ 1464ms | ✓ 1126ms | ✓ 1463ms | ✓ 1203ms | http |
| 165.227.5.10:8888 | ✓ 113ms | ✓ 1567ms | ✓ 541ms | 否 | 否 | http |
| 138.124.53.25:7443 | ✓ 1103ms | 否 | ✓ 1852ms | ✓ 1942ms | 否 | http |
| 38.34.179.203:8450 | ✓ 801ms | ✓ 664ms | ✓ 568ms | 否 | ✓ 1290ms | http |
| 38.34.178.241:8450 | ✓ 1206ms | ✓ 1340ms | ✓ 1519ms | ✓ 1531ms | ✓ 1175ms | http |
| 45.136.131.29:8450 | ✓ 1282ms | ✓ 1801ms | ✓ 963ms | 否 | 否 | http |
| 45.136.131.25:8450 | ✓ 1312ms | ✓ 1845ms | ✓ 982ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 1162ms | ✓ 1316ms | ✓ 842ms | ✓ 1085ms | ✓ 1061ms | http |
| 103.84.95.54:7890 | ✓ 676ms | 否 | 否 | ✓ 824ms | ✓ 1049ms | http |
| 137.220.150.170:6005 | ✓ 1756ms | 否 | ✓ 870ms | ✓ 1299ms | ✓ 1254ms | http |
| 195.123.209.48:3128 | ✓ 1448ms | 否 | ✓ 1602ms | 否 | ✓ 1926ms | http |
| 185.191.236.162:3128 | ✓ 1164ms | 否 | ✓ 1665ms | 否 | ✓ 1500ms | http |
| 190.12.150.244:999 | ✓ 990ms | 否 | ✓ 1151ms | 否 | ✓ 1557ms | http |
| 116.80.65.85:3172 | ✓ 1656ms | 否 | ✓ 1541ms | ✓ 1816ms | ✓ 1685ms | http |
| 59.46.216.131:30001 | ✓ 963ms | 否 | ✓ 1171ms | 否 | ✓ 1040ms | http |
| 116.80.96.95:3172 | ✓ 1928ms | 否 | 否 | ✓ 1806ms | ✓ 1638ms | http |
| 140.227.61.201:3128 | ✓ 1661ms | 否 | ✓ 1915ms | ✓ 1943ms | 否 | http |
| 38.34.179.71:8447 | ✓ 730ms | ✓ 711ms | ✓ 347ms | ✓ 1253ms | ✓ 1707ms | http |
| 38.34.179.64:8450 | ✓ 729ms | ✓ 613ms | ✓ 444ms | ✓ 1409ms | ✓ 1564ms | http |
| 45.136.130.190:8445 | ✓ 582ms | ✓ 600ms | ✓ 608ms | ✓ 1727ms | ✓ 1532ms | http |
| 38.145.218.230:8444 | ✓ 1186ms | ✓ 1259ms | ✓ 816ms | ✓ 916ms | ✓ 612ms | http |
| 101.43.127.100:8877 | ✓ 835ms | ✓ 1081ms | ✓ 876ms | ✓ 1054ms | ✓ 817ms | http |
| 38.145.220.55:8444 | ✓ 437ms | ✓ 832ms | ✓ 1794ms | ✓ 1018ms | ✓ 510ms | http |
| 38.145.220.65:8448 | ✓ 440ms | ✓ 1183ms | ✓ 1762ms | ✓ 767ms | ✓ 513ms | http |
| 38.145.220.198:8448 | ✓ 1385ms | 否 | ✓ 1016ms | ✓ 1889ms | 否 | http |
| 106.14.203.63:3333 | ✓ 809ms | 否 | ✓ 1568ms | ✓ 1132ms | 否 | http |
| 38.145.218.218:8445 | ✓ 1592ms | ✓ 1420ms | ✓ 1340ms | ✓ 1914ms | ✓ 508ms | http |
| 38.145.218.232:8445 | ✓ 1592ms | ✓ 1412ms | ✓ 1289ms | ✓ 1969ms | ✓ 516ms | http |
| 38.145.218.234:8445 | ✓ 1588ms | ✓ 1969ms | ✓ 1263ms | ✓ 1512ms | ✓ 516ms | http |
| 86.53.183.16:1080 | ✓ 1172ms | ✓ 1918ms | ✓ 1365ms | 否 | 否 | http |
| 83.219.250.8:62920 | ✓ 1179ms | 否 | ✓ 1423ms | 否 | ✓ 1778ms | http |
| 137.220.151.110:6005 | ✓ 975ms | 否 | ✓ 1613ms | ✓ 1638ms | ✓ 915ms | http |
| 193.23.200.251:10808 | ✓ 1194ms | 否 | ✓ 1820ms | 否 | ✓ 1489ms | http |
| 38.34.179.7:8448 | ✓ 393ms | ✓ 808ms | ✓ 323ms | ✓ 758ms | ✓ 551ms | http |
| 38.34.179.17:8448 | ✓ 390ms | ✓ 814ms | ✓ 324ms | ✓ 766ms | ✓ 529ms | http |
| 38.34.179.12:8448 | ✓ 393ms | ✓ 801ms | ✓ 340ms | ✓ 768ms | ✓ 540ms | http |
| 38.145.208.229:8450 | ✓ 388ms | ✓ 676ms | ✓ 147ms | ✓ 870ms | ✓ 1235ms | http |
| 38.34.179.161:8450 | ✓ 867ms | 否 | ✓ 86ms | ✓ 688ms | ✓ 931ms | http |
| 38.34.179.154:8448 | ✓ 1131ms | ✓ 1805ms | ✓ 88ms | ✓ 700ms | ✓ 913ms | http |
| 38.34.179.46:8448 | ✓ 560ms | ✓ 608ms | ✓ 156ms | ✓ 1170ms | 否 | http |
| 211.171.114.154:3128 | ✓ 847ms | ✓ 1221ms | ✓ 1423ms | 否 | ✓ 1100ms | http |
| 1.225.116.115:1080 | ✓ 1155ms | 否 | ✓ 1328ms | ✓ 1183ms | ✓ 1452ms | http |
| 47.101.159.19:8899 | 否 | ✓ 1092ms | ✓ 865ms | 否 | ✓ 895ms | http |
| 171.229.3.99:26384 | ✓ 1457ms | 否 | ✓ 1758ms | ✓ 1545ms | ✓ 1366ms | http |
| 85.8.182.108:443 | ✓ 1800ms | 否 | 否 | ✓ 1163ms | ✓ 657ms | http |
| 222.228.171.92:8080 | ✓ 679ms | 否 | 否 | ✓ 1138ms | ✓ 850ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1759ms | ✓ 1958ms | ✓ 1206ms | http |
| 116.80.95.238:7777 | ✓ 1835ms | 否 | ✓ 1876ms | 否 | ✓ 1659ms | http |
| 159.223.42.219:3128 | ✓ 1376ms | 否 | ✓ 1534ms | ✓ 1068ms | ✓ 828ms | http |
| 112.163.160.93:3128 | 否 | 否 | ✓ 1845ms | ✓ 985ms | ✓ 781ms | http |
| 104.243.46.122:3128 | ✓ 575ms | ✓ 1383ms | ✓ 1311ms | ✓ 1632ms | ✓ 1062ms | http |
| 103.171.161.96:9090 | ✓ 1856ms | 否 | 否 | ✓ 1380ms | ✓ 1338ms | http |
| 180.127.149.228:1080 | ✓ 1034ms | ✓ 1188ms | ✓ 880ms | ✓ 1262ms | ✓ 964ms | http |
| 168.222.254.136:8888 | ✓ 1208ms | ✓ 1792ms | ✓ 1611ms | 否 | ✓ 1977ms | http |
| 38.145.203.110:8450 | ✓ 502ms | ✓ 793ms | ✓ 176ms | ✓ 809ms | ✓ 493ms | http |
| 38.145.220.81:8452 | ✓ 503ms | ✓ 913ms | ✓ 116ms | ✓ 750ms | ✓ 518ms | http |
| 45.136.131.35:8452 | ✓ 505ms | ✓ 761ms | ✓ 207ms | ✓ 667ms | ✓ 753ms | http |
| 38.34.179.193:8448 | ✓ 520ms | ✓ 657ms | ✓ 617ms | ✓ 699ms | ✓ 927ms | http |
| 38.34.179.79:8450 | ✓ 518ms | ✓ 1157ms | ✓ 111ms | ✓ 1012ms | ✓ 1967ms | http |
| 38.34.179.78:8448 | ✓ 518ms | ✓ 604ms | ✓ 654ms | ✓ 1897ms | ✓ 1234ms | http |
| 45.149.92.147:5001 | ✓ 629ms | 否 | ✓ 639ms | ✓ 794ms | ✓ 634ms | http |
| 38.145.220.9:8448 | ✓ 519ms | ✓ 1397ms | ✓ 881ms | ✓ 1569ms | ✓ 911ms | http |
| 101.47.73.135:3128 | ✓ 1680ms | 否 | 否 | ✓ 1607ms | ✓ 1228ms | http |
| 62.113.119.14:8080 | ✓ 775ms | 否 | ✓ 1151ms | ✓ 1615ms | ✓ 1217ms | http |
| 91.233.223.147:3128 | ✓ 1123ms | 否 | ✓ 1516ms | 否 | ✓ 1644ms | http |
| 38.145.203.105:8447 | ✓ 504ms | ✓ 1296ms | ✓ 80ms | ✓ 666ms | ✓ 1058ms | http |
| 180.125.216.109:8118 | ✓ 927ms | 否 | 否 | ✓ 1259ms | ✓ 942ms | http |
| 27.254.99.183:8118 | 否 | ✓ 1925ms | ✓ 1233ms | ✓ 1277ms | ✓ 1039ms | http |
| 168.235.110.63:3128 | ✓ 954ms | ✓ 1556ms | ✓ 1072ms | ✓ 1602ms | 否 | http |
| 8.219.97.248:80 | ✓ 1068ms | 否 | ✓ 1375ms | ✓ 1468ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1000ms | ✓ 1246ms | ✓ 980ms | ✓ 1270ms | ✓ 1046ms | http |
| 116.80.96.107:3172 | 否 | 否 | ✓ 1818ms | ✓ 1779ms | ✓ 1640ms | http |
| 116.80.49.167:3172 | ✓ 1632ms | 否 | ✓ 1467ms | 否 | ✓ 1666ms | http |
| 36.212.210.57:8088 | ✓ 775ms | ✓ 808ms | ✓ 773ms | ✓ 936ms | ✓ 729ms | http |
| 168.138.202.218:3128 | ✓ 686ms | 否 | ✓ 925ms | ✓ 1177ms | ✓ 1281ms | http |
| 180.127.149.244:1080 | ✓ 862ms | ✓ 1205ms | ✓ 986ms | ✓ 1174ms | ✓ 886ms | http |

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
