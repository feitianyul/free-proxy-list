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

最后更新：2026-03-19 16:47:24 UTC（2026-03-20 00:47:24 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1604ms | 否 | ✓ 1888ms | ✓ 1308ms | ✓ 1004ms | http |
| 174.138.24.77:1080 | 否 | 否 | ✓ 1063ms | ✓ 1397ms | ✓ 1016ms | http |
| 113.160.132.26:8080 | ✓ 1512ms | ✓ 1654ms | ✓ 1998ms | ✓ 1338ms | ✓ 1382ms | http |
| 147.161.239.240:8800 | ✓ 1742ms | 否 | ✓ 1812ms | ✓ 1914ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1677ms | 否 | ✓ 1518ms | ✓ 1234ms | 否 | http |
| 45.167.124.52:8080 | 否 | 否 | ✓ 1882ms | ✓ 1793ms | ✓ 1942ms | http |
| 85.198.96.242:3128 | ✓ 710ms | 否 | ✓ 1025ms | 否 | ✓ 1347ms | http |
| 219.117.204.211:7799 | ✓ 1946ms | 否 | ✓ 1095ms | ✓ 1202ms | ✓ 880ms | http |
| 137.220.150.152:6005 | ✓ 889ms | 否 | ✓ 1061ms | 否 | ✓ 937ms | http |
| 137.220.150.104:6005 | ✓ 867ms | 否 | ✓ 902ms | ✓ 1468ms | ✓ 1853ms | http |
| 45.125.67.37:443 | ✓ 1368ms | 否 | ✓ 1665ms | ✓ 1362ms | ✓ 1438ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1820ms | ✓ 1794ms | ✓ 1477ms | http |
| 137.220.151.110:6005 | ✓ 1441ms | 否 | ✓ 1098ms | ✓ 1231ms | ✓ 975ms | http |
| 133.242.138.34:8100 | ✓ 1855ms | 否 | ✓ 1858ms | ✓ 1969ms | 否 | http |
| 45.136.131.61:8451 | ✓ 621ms | ✓ 1256ms | ✓ 741ms | ✓ 1307ms | ✓ 649ms | http |
| 38.145.208.247:8444 | ✓ 1822ms | ✓ 995ms | ✓ 535ms | ✓ 1384ms | ✓ 1273ms | http |
| 38.145.203.34:8445 | ✓ 1424ms | 否 | ✓ 888ms | ✓ 1654ms | ✓ 1308ms | http |
| 101.43.127.100:8877 | ✓ 825ms | ✓ 992ms | ✓ 771ms | ✓ 973ms | ✓ 1444ms | http |
| 38.145.218.163:8451 | ✓ 1176ms | 否 | ✓ 1811ms | ✓ 854ms | ✓ 1201ms | http |
| 49.156.44.114:8080 | ✓ 1815ms | 否 | ✓ 1316ms | ✓ 1540ms | ✓ 1514ms | http |
| 38.145.208.167:8450 | ✓ 962ms | ✓ 1225ms | ✓ 593ms | ✓ 924ms | ✓ 882ms | http |
| 38.34.179.29:8452 | ✓ 934ms | ✓ 903ms | ✓ 943ms | ✓ 854ms | ✓ 765ms | http |
| 14.225.212.37:7890 | 否 | 否 | ✓ 1130ms | ✓ 1177ms | ✓ 1447ms | http |
| 38.145.208.229:8450 | ✓ 918ms | 否 | ✓ 779ms | ✓ 1060ms | 否 | http |
| 38.145.203.132:8450 | ✓ 919ms | 否 | ✓ 761ms | ✓ 1084ms | 否 | http |
| 45.136.131.47:8449 | ✓ 1548ms | 否 | ✓ 1033ms | 否 | ✓ 959ms | http |
| 154.16.41.18:5555 | ✓ 1647ms | ✓ 1413ms | ✓ 1550ms | ✓ 1330ms | ✓ 1908ms | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1367ms | ✓ 1818ms | ✓ 1335ms | http |
| 140.245.66.105:8081 | ✓ 1456ms | 否 | 否 | ✓ 1299ms | ✓ 1349ms | http |
| 152.69.229.220:3128 | ✓ 1450ms | 否 | 否 | ✓ 1368ms | ✓ 1435ms | http |
| 38.55.106.206:6005 | ✓ 1874ms | 否 | ✓ 1685ms | 否 | ✓ 1803ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1045ms | ✓ 1498ms | ✓ 1654ms | http |
| 59.46.216.131:30001 | ✓ 1858ms | 否 | ✓ 1189ms | ✓ 1385ms | 否 | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 1138ms | ✓ 1831ms | ✓ 1128ms | http |
| 210.223.44.230:3128 | ✓ 806ms | ✓ 1461ms | 否 | 否 | ✓ 837ms | http |
| 45.140.147.155:1081 | ✓ 521ms | 否 | ✓ 1303ms | 否 | ✓ 1408ms | http |
| 45.140.147.155:1082 | ✓ 538ms | 否 | ✓ 1282ms | 否 | ✓ 1408ms | http |
| 35.225.22.61:80 | ✓ 1029ms | 否 | 否 | ✓ 1336ms | ✓ 782ms | http |
| 160.238.65.3:3128 | 否 | 否 | ✓ 1214ms | ✓ 1317ms | ✓ 1018ms | http |
| 160.238.65.2:3128 | 否 | 否 | ✓ 1216ms | ✓ 1330ms | ✓ 1019ms | http |
| 160.238.65.6:3128 | 否 | 否 | ✓ 1217ms | ✓ 1340ms | ✓ 1034ms | http |
| 160.238.65.5:3128 | 否 | 否 | ✓ 1206ms | ✓ 1357ms | ✓ 1050ms | http |
| 160.238.65.7:3128 | 否 | 否 | ✓ 1210ms | ✓ 1341ms | ✓ 1048ms | http |
| 160.238.65.4:3128 | 否 | 否 | ✓ 1217ms | ✓ 1366ms | ✓ 1021ms | http |
| 160.238.65.9:3128 | 否 | 否 | ✓ 1210ms | ✓ 1365ms | ✓ 1029ms | http |
| 160.238.65.8:3128 | 否 | 否 | ✓ 1210ms | ✓ 1340ms | ✓ 1037ms | http |
| 45.136.198.40:3128 | 否 | ✓ 1974ms | ✓ 1291ms | ✓ 1919ms | ✓ 1466ms | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 1201ms | ✓ 1497ms | ✓ 1351ms | http |
| 120.92.212.16:8890 | ✓ 931ms | ✓ 1211ms | 否 | ✓ 1227ms | ✓ 954ms | http |
| 38.145.208.234:8453 | ✓ 812ms | ✓ 1665ms | ✓ 938ms | ✓ 844ms | ✓ 1657ms | http |
| 38.34.178.153:8449 | ✓ 1306ms | 否 | 否 | ✓ 1023ms | ✓ 1820ms | http |
| 120.92.212.16:7890 | ✓ 949ms | ✓ 1208ms | 否 | ✓ 1187ms | ✓ 1817ms | http |
| 113.176.92.71:3128 | ✓ 1422ms | ✓ 1393ms | ✓ 1476ms | ✓ 1891ms | ✓ 1017ms | http |
| 181.78.51.157:999 | ✓ 938ms | ✓ 1623ms | ✓ 1873ms | 否 | 否 | http |
| 64.31.49.174:3128 | 否 | ✓ 1444ms | 否 | ✓ 1733ms | ✓ 1074ms | http |
| 38.145.208.181:8445 | ✓ 803ms | ✓ 916ms | ✓ 1388ms | ✓ 798ms | ✓ 1037ms | http |
| 115.190.91.223:7897 | 否 | ✓ 1194ms | ✓ 1984ms | ✓ 1564ms | 否 | http |
| 20.120.225.109:3128 | ✓ 808ms | ✓ 1382ms | ✓ 1038ms | ✓ 855ms | ✓ 1545ms | http |
| 38.145.220.198:8448 | ✓ 906ms | ✓ 1173ms | ✓ 641ms | 否 | ✓ 921ms | http |
| 38.34.179.35:8445 | ✓ 297ms | ✓ 1218ms | ✓ 1052ms | ✓ 953ms | ✓ 678ms | http |
| 45.136.131.39:8451 | ✓ 420ms | ✓ 1005ms | 否 | ✓ 990ms | ✓ 877ms | http |
| 193.23.200.251:10808 | ✓ 1431ms | 否 | ✓ 1514ms | 否 | ✓ 1500ms | http |
| 38.34.179.63:8448 | ✓ 519ms | ✓ 1171ms | ✓ 731ms | ✓ 843ms | ✓ 646ms | http |
| 38.34.179.60:8450 | ✓ 449ms | 否 | ✓ 1760ms | ✓ 1141ms | 否 | http |
| 38.145.218.212:8448 | ✓ 873ms | ✓ 791ms | ✓ 662ms | ✓ 1053ms | ✓ 1184ms | http |
| 38.145.218.232:8448 | ✓ 1475ms | ✓ 1176ms | ✓ 725ms | 否 | ✓ 632ms | http |
| 38.145.218.235:8445 | ✓ 1795ms | ✓ 838ms | ✓ 1117ms | ✓ 1579ms | 否 | http |
| 137.220.150.22:6005 | ✓ 840ms | 否 | ✓ 1201ms | ✓ 1206ms | ✓ 938ms | http |
| 38.145.220.60:8444 | ✓ 299ms | ✓ 794ms | ✓ 257ms | ✓ 982ms | ✓ 608ms | http |
| 38.34.179.61:8445 | ✓ 1583ms | ✓ 1028ms | ✓ 472ms | ✓ 1046ms | ✓ 1754ms | http |
| 38.145.218.102:8444 | ✓ 312ms | ✓ 972ms | ✓ 704ms | ✓ 837ms | ✓ 839ms | http |
| 45.136.131.51:8451 | ✓ 235ms | ✓ 1179ms | ✓ 861ms | ✓ 835ms | ✓ 1101ms | http |
| 45.136.130.177:8448 | ✓ 530ms | ✓ 1212ms | ✓ 1561ms | ✓ 1402ms | ✓ 1845ms | http |
| 45.136.131.63:8451 | ✓ 512ms | ✓ 1285ms | ✓ 766ms | ✓ 801ms | ✓ 684ms | http |
| 45.136.131.44:8452 | ✓ 1252ms | ✓ 1321ms | ✓ 316ms | ✓ 1573ms | 否 | http |
| 45.136.130.186:8451 | ✓ 1712ms | 否 | ✓ 319ms | ✓ 850ms | ✓ 825ms | http |
| 38.34.179.190:8450 | 否 | ✓ 1358ms | ✓ 862ms | 否 | ✓ 1076ms | http |
| 45.136.131.65:8451 | ✓ 203ms | ✓ 783ms | ✓ 329ms | ✓ 909ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1739ms | 否 | ✓ 1598ms | ✓ 1612ms | 否 | http |
| 38.34.179.65:8448 | ✓ 1320ms | ✓ 1777ms | ✓ 347ms | ✓ 1129ms | 否 | http |
| 45.136.130.176:8444 | ✓ 1471ms | ✓ 823ms | ✓ 838ms | 否 | 否 | http |
| 45.136.130.171:8445 | ✓ 1826ms | 否 | ✓ 217ms | ✓ 1037ms | ✓ 644ms | http |
| 38.34.179.78:8448 | ✓ 905ms | 否 | ✓ 987ms | ✓ 1068ms | 否 | http |
| 38.34.179.192:8450 | 否 | ✓ 1674ms | ✓ 632ms | 否 | ✓ 1299ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1283ms | 否 | ✓ 1456ms | ✓ 1107ms | http |
| 88.80.150.82:8080 | ✓ 1733ms | 否 | ✓ 1953ms | 否 | ✓ 1373ms | https |
| 45.136.131.54:8448 | ✓ 943ms | 否 | ✓ 472ms | ✓ 1070ms | ✓ 1722ms | http |

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
