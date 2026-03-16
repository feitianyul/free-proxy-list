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

最后更新：2026-03-16 21:32:48 UTC（2026-03-17 05:32:48 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.209.239.31:30000 | ✓ 1438ms | ✓ 1695ms | ✓ 577ms | ✓ 858ms | ✓ 1051ms | http |
| 147.161.210.140:8800 | ✓ 1439ms | ✓ 1113ms | ✓ 1018ms | ✓ 1330ms | ✓ 1025ms | http |
| 202.155.12.161:443 | ✓ 1461ms | 否 | ✓ 1374ms | ✓ 1899ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1466ms | ✓ 935ms | ✓ 1790ms | ✓ 1074ms | ✓ 861ms | http |
| 113.160.132.26:8080 | ✓ 1382ms | ✓ 1371ms | ✓ 1225ms | ✓ 1608ms | ✓ 1509ms | http |
| 137.220.150.104:6005 | ✓ 1170ms | 否 | ✓ 1391ms | ✓ 1268ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1565ms | ✓ 1917ms | 否 | 否 | ✓ 1494ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1483ms | ✓ 1015ms | ✓ 1589ms | ✓ 1197ms | http |
| 45.140.147.155:1081 | ✓ 643ms | 否 | ✓ 1334ms | ✓ 1683ms | ✓ 1403ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1131ms | 否 | ✓ 1318ms | ✓ 878ms | http |
| 137.220.151.110:6005 | ✓ 1804ms | 否 | ✓ 1748ms | ✓ 1378ms | ✓ 897ms | http |
| 59.46.216.131:30001 | ✓ 953ms | ✓ 1285ms | ✓ 1019ms | ✓ 1409ms | 否 | http |
| 137.220.150.152:6005 | ✓ 785ms | 否 | ✓ 1039ms | ✓ 1471ms | 否 | http |
| 168.235.110.63:3128 | ✓ 629ms | ✓ 1356ms | ✓ 262ms | ✓ 1202ms | 否 | http |
| 101.43.127.100:8877 | ✓ 860ms | ✓ 1106ms | ✓ 830ms | ✓ 1148ms | ✓ 914ms | http |
| 120.92.212.16:8890 | ✓ 1284ms | 否 | 否 | ✓ 1486ms | ✓ 989ms | http |
| 14.225.212.37:7890 | ✓ 1353ms | ✓ 1553ms | ✓ 1738ms | 否 | ✓ 1925ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 636ms | ✓ 804ms | ✓ 632ms | http |
| 38.34.179.98:8447 | ✓ 762ms | ✓ 607ms | ✓ 583ms | ✓ 628ms | ✓ 475ms | http |
| 35.225.22.61:80 | ✓ 798ms | 否 | ✓ 1125ms | 否 | ✓ 1149ms | http |
| 38.34.178.153:8448 | ✓ 780ms | ✓ 915ms | ✓ 906ms | ✓ 1520ms | ✓ 589ms | http |
| 178.236.245.17:3128 | ✓ 1329ms | 否 | ✓ 1532ms | 否 | ✓ 1750ms | http |
| 38.34.179.14:8450 | ✓ 1981ms | ✓ 885ms | ✓ 1158ms | 否 | ✓ 505ms | http |
| 59.8.203.55:80 | ✓ 1353ms | 否 | ✓ 844ms | ✓ 945ms | ✓ 791ms | http |
| 212.192.12.90:6005 | ✓ 1308ms | 否 | ✓ 1075ms | ✓ 869ms | ✓ 762ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1323ms | 否 | ✓ 996ms | ✓ 843ms | http |
| 72.11.151.159:6005 | ✓ 683ms | 否 | ✓ 1022ms | ✓ 1350ms | ✓ 1131ms | http |
| 137.220.150.170:6005 | ✓ 1522ms | ✓ 1815ms | ✓ 731ms | ✓ 1447ms | ✓ 824ms | http |
| 38.55.107.137:6005 | 否 | 否 | ✓ 973ms | ✓ 1137ms | ✓ 661ms | http |
| 106.14.203.63:3333 | ✓ 1833ms | ✓ 1220ms | ✓ 1962ms | ✓ 1123ms | ✓ 895ms | http |
| 86.53.183.16:1080 | ✓ 1926ms | 否 | ✓ 1289ms | ✓ 1892ms | ✓ 1527ms | http |
| 159.223.42.219:3128 | ✓ 1512ms | 否 | 否 | ✓ 1561ms | ✓ 1657ms | http |
| 37.187.109.70:10111 | ✓ 1523ms | 否 | ✓ 1526ms | 否 | ✓ 1912ms | http |
| 138.124.53.25:7443 | ✓ 1916ms | 否 | ✓ 1803ms | 否 | ✓ 1311ms | http |
| 49.7.179.70:3333 | 否 | ✓ 1275ms | ✓ 956ms | ✓ 1348ms | 否 | http |
| 91.107.148.58:53967 | ✓ 1167ms | 否 | ✓ 1993ms | 否 | ✓ 1627ms | http |
| 45.136.130.165:8443 | ✓ 95ms | ✓ 779ms | ✓ 1205ms | ✓ 968ms | ✓ 528ms | http |
| 2.56.122.146:10808 | ✓ 557ms | 否 | ✓ 604ms | 否 | ✓ 1007ms | http |
| 47.77.193.180:1080 | 否 | ✓ 1743ms | ✓ 775ms | ✓ 773ms | ✓ 1332ms | http |
| 101.47.73.135:3128 | ✓ 1304ms | 否 | ✓ 1865ms | 否 | ✓ 1284ms | http |
| 178.236.245.59:3128 | ✓ 1226ms | 否 | ✓ 1815ms | ✓ 1844ms | ✓ 1418ms | http |
| 152.70.137.18:8888 | ✓ 455ms | ✓ 1451ms | ✓ 1596ms | ✓ 1564ms | 否 | http |
| 180.103.19.40:1080 | ✓ 1564ms | ✓ 1805ms | ✓ 1239ms | ✓ 1841ms | ✓ 1356ms | http |
| 149.50.116.240:1080 | ✓ 1839ms | 否 | ✓ 1637ms | ✓ 1874ms | 否 | http |
| 194.5.212.40:8080 | ✓ 1575ms | ✓ 1973ms | ✓ 917ms | 否 | ✓ 1412ms | http |
| 38.145.218.210:8448 | ✓ 494ms | ✓ 672ms | ✓ 622ms | ✓ 632ms | ✓ 596ms | http |
| 38.145.218.234:8447 | ✓ 494ms | ✓ 629ms | ✓ 664ms | ✓ 664ms | ✓ 887ms | http |
| 158.69.185.37:3129 | ✓ 1177ms | ✓ 1890ms | ✓ 1138ms | ✓ 1280ms | ✓ 1018ms | http |
| 219.117.204.211:7799 | ✓ 1418ms | 否 | ✓ 1034ms | ✓ 940ms | ✓ 845ms | http |
| 133.242.138.34:8100 | ✓ 1432ms | ✓ 1475ms | ✓ 1314ms | 否 | 否 | http |
| 211.171.114.154:3128 | ✓ 1385ms | ✓ 1939ms | ✓ 1452ms | ✓ 1207ms | 否 | http |
| 38.145.218.101:8447 | ✓ 676ms | ✓ 1390ms | ✓ 1454ms | ✓ 998ms | ✓ 629ms | http |
| 38.34.183.224:8448 | 否 | ✓ 1958ms | ✓ 719ms | ✓ 1293ms | ✓ 515ms | http |
| 45.136.130.163:8443 | ✓ 1183ms | ✓ 1157ms | ✓ 1022ms | ✓ 1744ms | ✓ 1014ms | http |
| 45.136.130.164:8443 | ✓ 1181ms | ✓ 1157ms | ✓ 1025ms | ✓ 1792ms | ✓ 962ms | http |
| 193.23.200.251:10808 | ✓ 1070ms | 否 | ✓ 1188ms | 否 | ✓ 1754ms | http |
| 38.145.220.33:8448 | ✓ 554ms | ✓ 656ms | ✓ 1086ms | ✓ 699ms | 否 | http |
| 38.34.183.234:8450 | ✓ 717ms | ✓ 678ms | ✓ 562ms | ✓ 1371ms | ✓ 1627ms | http |
| 85.198.96.242:3128 | ✓ 1764ms | 否 | ✓ 1522ms | ✓ 1992ms | ✓ 1844ms | http |
| 45.136.131.30:8447 | 否 | ✓ 1322ms | 否 | ✓ 1899ms | ✓ 1400ms | http |
| 45.136.131.28:8447 | 否 | ✓ 1335ms | 否 | ✓ 1880ms | ✓ 1416ms | http |
| 116.80.65.81:3172 | ✓ 1886ms | 否 | ✓ 1519ms | ✓ 1930ms | 否 | http |
| 45.136.130.160:8447 | 否 | 否 | ✓ 851ms | ✓ 777ms | ✓ 516ms | http |
| 172.212.68.37:3128 | ✓ 371ms | ✓ 1857ms | ✓ 1731ms | 否 | 否 | http |
| 116.80.49.170:3172 | ✓ 1554ms | 否 | ✓ 1993ms | ✓ 1972ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1434ms | ✓ 1556ms | ✓ 1195ms | ✓ 1500ms | ✓ 1000ms | http |
| 24.144.86.173:1080 | ✓ 388ms | ✓ 804ms | ✓ 977ms | ✓ 1265ms | ✓ 1161ms | http |
| 45.136.198.40:3128 | ✓ 1866ms | 否 | 否 | ✓ 1698ms | ✓ 1341ms | http |
| 8.219.97.248:80 | ✓ 961ms | 否 | ✓ 1692ms | 否 | ✓ 1383ms | http |
| 38.34.179.60:8450 | 否 | 否 | ✓ 1535ms | ✓ 1023ms | ✓ 890ms | http |
| 38.145.208.190:8450 | ✓ 189ms | ✓ 621ms | ✓ 548ms | ✓ 953ms | ✓ 518ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1424ms | 否 | ✓ 1929ms | ✓ 1352ms | http |
| 61.52.131.172:8443 | ✓ 1511ms | ✓ 1212ms | ✓ 994ms | 否 | ✓ 1379ms | http |
| 45.136.130.156:8450 | 否 | ✓ 1817ms | ✓ 784ms | ✓ 1130ms | ✓ 1537ms | http |
| 116.80.49.169:3172 | ✓ 1546ms | 否 | ✓ 1495ms | 否 | ✓ 1683ms | http |
| 38.34.179.203:8450 | 否 | ✓ 1683ms | ✓ 666ms | ✓ 980ms | ✓ 540ms | http |
| 103.113.70.189:1081 | ✓ 1038ms | ✓ 1963ms | 否 | ✓ 1172ms | ✓ 880ms | http |

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
