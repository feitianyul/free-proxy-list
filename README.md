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

最后更新：2026-03-21 10:27:33 UTC（2026-03-21 18:27:33 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1560ms | 否 | ✓ 767ms | ✓ 913ms | ✓ 1572ms | http |
| 147.161.239.240:8800 | ✓ 1408ms | 否 | ✓ 1379ms | ✓ 1500ms | ✓ 1291ms | http |
| 167.103.34.108:8800 | ✓ 1956ms | 否 | ✓ 1420ms | ✓ 1502ms | 否 | http |
| 137.220.150.22:6005 | ✓ 709ms | 否 | ✓ 1633ms | 否 | ✓ 848ms | http |
| 137.220.150.152:6005 | ✓ 1331ms | 否 | ✓ 798ms | 否 | ✓ 1412ms | http |
| 185.114.73.2:1080 | ✓ 1365ms | 否 | 否 | ✓ 1775ms | ✓ 1548ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 991ms | ✓ 1524ms | ✓ 1366ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1419ms | 否 | ✓ 1185ms | ✓ 1168ms | http |
| 219.117.204.211:7799 | ✓ 1394ms | 否 | ✓ 1854ms | ✓ 781ms | ✓ 623ms | http |
| 35.225.22.61:80 | ✓ 638ms | ✓ 1302ms | 否 | ✓ 1369ms | ✓ 967ms | http |
| 59.46.216.131:30001 | ✓ 911ms | ✓ 1290ms | ✓ 1071ms | 否 | 否 | http |
| 5.104.87.17:8050 | ✓ 1171ms | 否 | ✓ 1298ms | ✓ 929ms | ✓ 767ms | http |
| 137.220.151.110:6005 | ✓ 1185ms | 否 | ✓ 829ms | ✓ 1384ms | ✓ 1112ms | http |
| 137.220.150.104:6005 | ✓ 1185ms | 否 | ✓ 1033ms | ✓ 1074ms | ✓ 1225ms | http |
| 167.103.31.122:8800 | ✓ 1457ms | 否 | ✓ 1338ms | ✓ 1589ms | 否 | http |
| 202.38.72.235:26001 | ✓ 1864ms | ✓ 1592ms | ✓ 1739ms | ✓ 1693ms | ✓ 1349ms | http |
| 120.92.212.16:7890 | ✓ 997ms | ✓ 1429ms | ✓ 971ms | 否 | 否 | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1334ms | ✓ 1472ms | ✓ 1505ms | http |
| 101.43.127.100:8877 | ✓ 1110ms | ✓ 1035ms | 否 | ✓ 1455ms | ✓ 1547ms | http |
| 160.250.4.245:1 | ✓ 1424ms | 否 | ✓ 1057ms | ✓ 1389ms | ✓ 1002ms | http |
| 142.171.224.229:7890 | ✓ 1771ms | ✓ 1478ms | ✓ 354ms | ✓ 668ms | ✓ 493ms | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 1065ms | ✓ 973ms | ✓ 1733ms | http |
| 133.242.138.34:8100 | 否 | ✓ 1946ms | ✓ 844ms | ✓ 1197ms | ✓ 1294ms | http |
| 106.75.15.167:7890 | ✓ 1131ms | ✓ 1704ms | 否 | ✓ 1574ms | 否 | http |
| 150.249.255.91:3128 | 否 | ✓ 1635ms | ✓ 651ms | ✓ 908ms | ✓ 1028ms | http |
| 103.113.70.189:1081 | ✓ 939ms | 否 | ✓ 651ms | ✓ 1737ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1371ms | 否 | ✓ 1108ms | 否 | ✓ 1416ms | http |
| 166.88.55.83:7890 | ✓ 718ms | ✓ 1442ms | ✓ 607ms | ✓ 755ms | ✓ 593ms | http |
| 115.231.181.40:8128 | ✓ 849ms | ✓ 1172ms | 否 | ✓ 1091ms | ✓ 886ms | http |
| 45.140.147.155:1082 | ✓ 582ms | 否 | ✓ 1203ms | ✓ 1830ms | ✓ 1555ms | http |
| 38.34.179.179:8449 | 否 | ✓ 824ms | ✓ 98ms | ✓ 690ms | ✓ 522ms | http |
| 38.145.208.224:8445 | 否 | ✓ 1702ms | ✓ 878ms | ✓ 875ms | ✓ 531ms | http |
| 38.145.203.32:8452 | 否 | 否 | ✓ 395ms | ✓ 1047ms | ✓ 571ms | http |
| 45.136.130.171:8452 | 否 | 否 | ✓ 751ms | ✓ 1556ms | ✓ 1312ms | http |
| 181.78.44.63:999 | ✓ 1192ms | 否 | ✓ 1536ms | ✓ 1547ms | ✓ 1401ms | http |
| 38.34.179.191:8453 | ✓ 225ms | ✓ 635ms | ✓ 117ms | ✓ 773ms | ✓ 523ms | http |
| 45.136.131.50:8445 | ✓ 133ms | ✓ 1151ms | ✓ 352ms | ✓ 696ms | ✓ 554ms | http |
| 38.145.203.97:8449 | ✓ 203ms | ✓ 1663ms | ✓ 103ms | ✓ 689ms | ✓ 548ms | http |
| 38.145.208.192:8448 | ✓ 151ms | 否 | ✓ 88ms | ✓ 677ms | ✓ 518ms | http |
| 45.136.131.43:8444 | ✓ 170ms | ✓ 740ms | ✓ 670ms | ✓ 691ms | ✓ 565ms | http |
| 38.145.208.189:8446 | ✓ 145ms | 否 | ✓ 117ms | ✓ 701ms | ✓ 569ms | http |
| 45.136.130.179:8444 | ✓ 418ms | 否 | ✓ 140ms | ✓ 689ms | ✓ 1325ms | http |
| 38.34.179.160:8453 | ✓ 166ms | ✓ 1317ms | ✓ 655ms | ✓ 998ms | ✓ 503ms | http |
| 38.34.179.68:8451 | ✓ 1418ms | ✓ 678ms | ✓ 90ms | ✓ 681ms | ✓ 528ms | http |
| 38.145.220.72:8444 | ✓ 1339ms | ✓ 745ms | ✓ 97ms | ✓ 755ms | ✓ 1344ms | http |
| 45.136.131.39:8451 | ✓ 378ms | ✓ 1792ms | ✓ 400ms | ✓ 1282ms | ✓ 1523ms | http |
| 38.34.179.102:8449 | ✓ 327ms | ✓ 1448ms | ✓ 108ms | ✓ 671ms | ✓ 594ms | http |
| 38.145.208.253:8445 | ✓ 297ms | ✓ 667ms | ✓ 753ms | ✓ 1842ms | ✓ 673ms | http |
| 38.145.208.245:8444 | ✓ 314ms | ✓ 688ms | ✓ 717ms | ✓ 1825ms | ✓ 744ms | http |
| 45.136.131.54:8448 | ✓ 678ms | ✓ 1861ms | ✓ 619ms | ✓ 658ms | ✓ 541ms | http |
| 38.145.203.135:8451 | ✓ 385ms | 否 | ✓ 1293ms | ✓ 1000ms | ✓ 999ms | http |
| 45.136.130.174:8450 | ✓ 501ms | 否 | ✓ 846ms | ✓ 773ms | ✓ 588ms | http |
| 217.174.244.117:3129 | ✓ 996ms | 否 | ✓ 1361ms | 否 | ✓ 1553ms | http |
| 91.238.105.64:2024 | 否 | 否 | ✓ 1085ms | ✓ 1747ms | ✓ 1313ms | http |
| 38.145.218.161:8444 | ✓ 716ms | 否 | ✓ 256ms | ✓ 1096ms | ✓ 1572ms | http |
| 38.145.218.230:8453 | ✓ 1219ms | 否 | ✓ 237ms | ✓ 725ms | ✓ 1634ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1266ms | ✓ 1510ms | ✓ 1631ms | 否 | http |
| 116.80.65.80:3172 | ✓ 1441ms | 否 | ✓ 1454ms | ✓ 1797ms | 否 | http |
| 45.136.131.53:8452 | 否 | ✓ 1416ms | ✓ 91ms | ✓ 662ms | ✓ 507ms | http |
| 24.144.86.173:1080 | ✓ 1329ms | 否 | ✓ 1426ms | ✓ 1470ms | ✓ 1121ms | http |
| 210.223.44.230:3128 | 否 | ✓ 979ms | ✓ 1170ms | 否 | ✓ 1055ms | http |
| 45.136.130.171:8445 | 否 | ✓ 1039ms | ✓ 263ms | ✓ 757ms | ✓ 696ms | http |
| 114.237.77.231:1080 | ✓ 872ms | ✓ 1184ms | ✓ 966ms | ✓ 1422ms | ✓ 916ms | http |
| 202.141.161.53:30001 | 否 | 否 | ✓ 1303ms | ✓ 1499ms | ✓ 1086ms | http |
| 210.45.70.16:7895 | ✓ 1915ms | ✓ 1277ms | 否 | ✓ 1715ms | 否 | http |
| 106.117.208.101:7890 | ✓ 990ms | ✓ 1907ms | 否 | ✓ 1277ms | ✓ 890ms | http |
| 27.254.99.183:8118 | ✓ 1594ms | ✓ 1987ms | 否 | ✓ 1242ms | 否 | http |
| 144.31.79.117:8888 | ✓ 1363ms | 否 | ✓ 1446ms | 否 | ✓ 1747ms | http |
| 5.102.109.41:999 | ✓ 1723ms | 否 | 否 | ✓ 1403ms | ✓ 1259ms | http |
| 20.120.225.109:3128 | ✓ 312ms | ✓ 1143ms | ✓ 1781ms | ✓ 884ms | ✓ 873ms | http |
| 38.34.179.98:8453 | ✓ 983ms | 否 | ✓ 169ms | ✓ 934ms | 否 | http |
| 194.67.99.223:1080 | ✓ 1288ms | 否 | ✓ 738ms | ✓ 1905ms | ✓ 1501ms | http |
| 134.209.153.66:3128 | ✓ 1249ms | 否 | ✓ 1741ms | ✓ 1529ms | ✓ 1262ms | http |
| 38.34.183.222:8453 | ✓ 232ms | ✓ 715ms | ✓ 146ms | ✓ 698ms | ✓ 736ms | http |
| 45.168.238.193:8443 | ✓ 369ms | 否 | ✓ 332ms | ✓ 1238ms | ✓ 1198ms | http |
| 38.34.178.245:8446 | ✓ 1994ms | 否 | ✓ 1906ms | 否 | ✓ 1504ms | http |
| 66.228.47.125:110 | ✓ 1656ms | ✓ 1915ms | ✓ 1916ms | 否 | ✓ 1612ms | http |
| 45.129.141.143:3128 | ✓ 1266ms | 否 | ✓ 1788ms | 否 | ✓ 1650ms | http |
| 103.39.51.190:8080 | ✓ 1777ms | 否 | 否 | ✓ 1609ms | ✓ 1558ms | http |
| 64.181.240.152:3128 | ✓ 815ms | ✓ 1145ms | ✓ 218ms | ✓ 669ms | ✓ 623ms | http |
| 180.125.216.109:8118 | 否 | 否 | ✓ 839ms | ✓ 1533ms | ✓ 1601ms | http |
| 47.101.159.19:8899 | ✓ 861ms | ✓ 1022ms | ✓ 914ms | ✓ 1504ms | ✓ 1646ms | http |
| 45.186.6.104:3128 | ✓ 1355ms | ✓ 1801ms | ✓ 1638ms | 否 | 否 | http |
| 171.227.173.219:10001 | 否 | 否 | ✓ 1627ms | ✓ 1500ms | ✓ 1340ms | http |
| 222.109.119.178:3128 | 否 | 否 | ✓ 944ms | ✓ 956ms | ✓ 1011ms | http |
| 103.67.46.225:3125 | ✓ 1845ms | 否 | ✓ 1707ms | 否 | ✓ 1646ms | http |
| 61.52.131.172:8443 | ✓ 883ms | ✓ 1104ms | ✓ 948ms | ✓ 1163ms | ✓ 910ms | http |
| 101.47.73.135:3128 | ✓ 1155ms | 否 | ✓ 1780ms | ✓ 1594ms | ✓ 959ms | http |
| 77.110.113.24:40000 | ✓ 1320ms | 否 | ✓ 1668ms | 否 | ✓ 1596ms | http |

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
