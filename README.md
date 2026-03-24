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

最后更新：2026-03-24 03:21:42 UTC（2026-03-24 11:21:42 UTC+8）

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
| 8.217.106.71:8888 | ✓ 609ms | 否 | ✓ 620ms | ✓ 985ms | ✓ 623ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 971ms | ✓ 1116ms | ✓ 975ms | http |
| 147.161.210.140:8800 | ✓ 1467ms | 否 | ✓ 755ms | ✓ 1392ms | ✓ 874ms | http |
| 113.160.132.26:8080 | ✓ 1423ms | 否 | ✓ 1223ms | ✓ 1344ms | ✓ 941ms | http |
| 167.103.34.108:8800 | ✓ 1349ms | 否 | ✓ 1174ms | ✓ 1543ms | ✓ 1224ms | http |
| 45.167.124.52:8080 | ✓ 1606ms | 否 | ✓ 1900ms | 否 | ✓ 1569ms | http |
| 20.27.15.49:8561 | ✓ 441ms | ✓ 1145ms | ✓ 436ms | ✓ 747ms | ✓ 577ms | http |
| 20.210.76.175:8561 | ✓ 456ms | ✓ 881ms | ✓ 608ms | ✓ 840ms | ✓ 581ms | http |
| 20.210.76.104:8561 | ✓ 448ms | ✓ 1184ms | ✓ 464ms | ✓ 726ms | ✓ 593ms | http |
| 20.210.76.178:8561 | ✓ 458ms | ✓ 1907ms | ✓ 434ms | ✓ 739ms | ✓ 585ms | http |
| 155.212.132.241:3128 | ✓ 727ms | 否 | ✓ 1138ms | ✓ 1962ms | ✓ 1500ms | http |
| 167.103.31.122:8800 | ✓ 1531ms | 否 | ✓ 1379ms | ✓ 1610ms | 否 | http |
| 45.88.0.117:3128 | ✓ 1242ms | ✓ 1891ms | ✓ 1715ms | 否 | 否 | http |
| 45.88.0.116:3128 | ✓ 1243ms | ✓ 1906ms | ✓ 1702ms | 否 | 否 | http |
| 116.80.65.79:3172 | ✓ 1561ms | 否 | ✓ 1475ms | 否 | ✓ 1602ms | http |
| 120.92.212.16:7890 | ✓ 1001ms | ✓ 1145ms | ✓ 1168ms | ✓ 1163ms | ✓ 914ms | http |
| 120.92.212.16:8890 | ✓ 1139ms | ✓ 1179ms | ✓ 959ms | ✓ 1361ms | ✓ 1133ms | http |
| 142.171.224.229:7890 | ✓ 1458ms | ✓ 1099ms | ✓ 1114ms | 否 | ✓ 640ms | http |
| 219.117.204.211:7799 | ✓ 1460ms | 否 | ✓ 472ms | 否 | ✓ 851ms | http |
| 38.34.179.150:8449 | ✓ 270ms | ✓ 1043ms | ✓ 1538ms | ✓ 702ms | ✓ 894ms | http |
| 150.249.255.91:3128 | ✓ 1461ms | ✓ 991ms | ✓ 592ms | 否 | 否 | http |
| 147.161.239.240:8800 | ✓ 1203ms | 否 | ✓ 1322ms | ✓ 1936ms | ✓ 1298ms | http |
| 103.84.95.54:7890 | ✓ 713ms | 否 | 否 | ✓ 1338ms | ✓ 723ms | http |
| 45.88.0.99:3128 | ✓ 1841ms | ✓ 1505ms | ✓ 1442ms | 否 | ✓ 1897ms | http |
| 45.88.0.98:3128 | ✓ 1289ms | 否 | ✓ 1930ms | ✓ 1934ms | 否 | http |
| 213.220.62.62:3128 | ✓ 1294ms | ✓ 1829ms | ✓ 1664ms | 否 | ✓ 1864ms | http |
| 45.88.0.115:3128 | ✓ 751ms | ✓ 1530ms | ✓ 1122ms | 否 | ✓ 1184ms | http |
| 5.102.109.41:999 | ✓ 535ms | ✓ 1338ms | ✓ 940ms | 否 | 否 | http |
| 116.80.49.156:3172 | 否 | 否 | ✓ 1457ms | ✓ 1869ms | ✓ 1711ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1225ms | ✓ 1141ms | ✓ 1197ms | http |
| 101.43.127.100:8877 | ✓ 911ms | ✓ 965ms | ✓ 853ms | ✓ 1299ms | ✓ 776ms | http |
| 59.8.203.55:80 | 否 | ✓ 1259ms | ✓ 981ms | ✓ 933ms | ✓ 751ms | http |
| 167.103.115.102:8800 | ✓ 979ms | 否 | ✓ 1232ms | ✓ 1027ms | ✓ 1209ms | http |
| 115.231.181.40:8128 | ✓ 1545ms | ✓ 1089ms | ✓ 1121ms | ✓ 1679ms | 否 | http |
| 137.220.150.22:6005 | ✓ 721ms | 否 | ✓ 720ms | ✓ 1134ms | ✓ 935ms | http |
| 218.89.134.230:3333 | ✓ 1428ms | ✓ 1592ms | ✓ 1450ms | 否 | ✓ 1234ms | http |
| 211.95.152.50:45046 | ✓ 1311ms | ✓ 1317ms | ✓ 1266ms | 否 | 否 | http |
| 118.113.246.72:1080 | ✓ 1485ms | ✓ 1651ms | ✓ 1596ms | ✓ 1478ms | ✓ 1482ms | http |
| 200.125.171.254:999 | ✓ 999ms | 否 | ✓ 1429ms | ✓ 1814ms | ✓ 1638ms | http |
| 1.231.81.166:3128 | ✓ 1690ms | 否 | ✓ 1852ms | ✓ 1798ms | ✓ 1434ms | http |
| 166.88.55.83:7890 | ✓ 772ms | 否 | 否 | ✓ 1260ms | ✓ 610ms | http |
| 144.31.79.117:8888 | ✓ 826ms | 否 | ✓ 1362ms | 否 | ✓ 1521ms | http |
| 150.107.140.238:3128 | ✓ 1522ms | 否 | 否 | ✓ 1140ms | ✓ 936ms | http |
| 45.144.28.81:10808 | ✓ 944ms | 否 | ✓ 1162ms | ✓ 1720ms | ✓ 1747ms | http |
| 217.76.245.80:999 | ✓ 1030ms | 否 | ✓ 1236ms | ✓ 1673ms | ✓ 1206ms | http |
| 59.46.216.131:30001 | ✓ 929ms | 否 | ✓ 1112ms | ✓ 1294ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1457ms | ✓ 1662ms | ✓ 894ms | ✓ 1378ms | ✓ 1164ms | http |
| 121.43.196.213:8222 | ✓ 889ms | ✓ 1065ms | ✓ 818ms | ✓ 1089ms | ✓ 813ms | http |
| 121.43.196.210:8222 | ✓ 896ms | ✓ 997ms | ✓ 884ms | ✓ 1043ms | ✓ 863ms | http |
| 114.55.226.123:10086 | 否 | ✓ 1564ms | ✓ 969ms | ✓ 1254ms | ✓ 1026ms | http |
| 172.212.68.37:3128 | ✓ 894ms | 否 | 否 | ✓ 1707ms | ✓ 1431ms | http |
| 181.41.201.85:3128 | ✓ 1386ms | 否 | ✓ 738ms | 否 | ✓ 1719ms | http |
| 47.95.231.180:8084 | ✓ 843ms | ✓ 1171ms | ✓ 840ms | ✓ 1547ms | 否 | http |
| 202.141.161.53:30001 | ✓ 1040ms | ✓ 1357ms | ✓ 1378ms | ✓ 1182ms | ✓ 1007ms | http |
| 45.140.147.155:1082 | ✓ 1152ms | 否 | ✓ 1150ms | ✓ 1955ms | 否 | http |
| 137.220.151.110:6005 | ✓ 1670ms | 否 | ✓ 771ms | ✓ 1085ms | ✓ 835ms | http |
| 210.223.44.230:3128 | ✓ 1462ms | 否 | ✓ 1201ms | ✓ 1850ms | ✓ 662ms | http |
| 194.67.99.223:1080 | ✓ 1758ms | 否 | ✓ 1789ms | ✓ 1917ms | ✓ 1470ms | http |
| 137.220.150.104:6005 | ✓ 1791ms | ✓ 1953ms | ✓ 812ms | ✓ 1270ms | ✓ 1413ms | http |
| 20.210.39.153:8561 | ✓ 1277ms | ✓ 1099ms | ✓ 496ms | ✓ 768ms | ✓ 575ms | http |
| 20.78.26.206:8561 | ✓ 1275ms | ✓ 1474ms | ✓ 429ms | ✓ 728ms | ✓ 570ms | http |
| 103.39.51.190:8080 | ✓ 1790ms | 否 | 否 | ✓ 1464ms | ✓ 1482ms | http |
| 103.3.246.71:3128 | ✓ 1056ms | 否 | 否 | ✓ 1208ms | ✓ 994ms | http |
| 160.250.5.22:1 | ✓ 1496ms | 否 | ✓ 1439ms | ✓ 1267ms | ✓ 1041ms | http |
| 160.250.4.13:1 | ✓ 1505ms | 否 | ✓ 1315ms | ✓ 1418ms | ✓ 1360ms | http |
| 222.228.171.92:8080 | ✓ 1928ms | ✓ 1852ms | ✓ 725ms | 否 | ✓ 878ms | http |
| 20.78.118.91:8561 | ✓ 1131ms | ✓ 678ms | ✓ 515ms | ✓ 722ms | ✓ 576ms | http |
| 8.219.97.248:80 | ✓ 891ms | 否 | ✓ 1020ms | 否 | ✓ 1185ms | http |
| 43.99.54.236:5555 | ✓ 609ms | ✓ 1306ms | ✓ 624ms | ✓ 780ms | ✓ 622ms | http |
| 111.79.111.126:3128 | ✓ 1402ms | 否 | ✓ 1807ms | 否 | ✓ 1178ms | http |
| 45.136.131.61:8453 | ✓ 147ms | ✓ 780ms | ✓ 291ms | ✓ 731ms | ✓ 581ms | http |
| 45.136.131.46:8446 | ✓ 316ms | ✓ 759ms | ✓ 748ms | ✓ 960ms | ✓ 507ms | http |
| 38.145.203.132:8450 | ✓ 145ms | ✓ 899ms | ✓ 324ms | ✓ 646ms | ✓ 687ms | http |
| 45.136.131.39:8453 | ✓ 322ms | ✓ 1032ms | ✓ 1209ms | ✓ 718ms | ✓ 921ms | http |
| 38.34.179.38:8449 | ✓ 909ms | ✓ 1613ms | ✓ 87ms | ✓ 723ms | ✓ 721ms | http |
| 47.101.149.27:9010 | ✓ 1243ms | ✓ 1229ms | 否 | ✓ 1373ms | 否 | http |
| 38.34.179.150:8444 | 否 | 否 | ✓ 816ms | ✓ 1913ms | ✓ 859ms | http |
| 45.136.131.53:8445 | ✓ 234ms | ✓ 1665ms | ✓ 1228ms | ✓ 1995ms | ✓ 510ms | http |
| 121.230.8.211:1080 | ✓ 1979ms | ✓ 1330ms | ✓ 898ms | ✓ 1397ms | ✓ 1456ms | http |
| 121.230.8.41:1080 | 否 | ✓ 1620ms | ✓ 1217ms | ✓ 1496ms | ✓ 1219ms | http |
| 38.34.179.178:8445 | ✓ 1602ms | 否 | ✓ 1218ms | 否 | ✓ 617ms | http |
| 58.220.95.8:10174 | ✓ 886ms | ✓ 1771ms | ✓ 1948ms | 否 | ✓ 1734ms | http |
| 45.136.198.40:3128 | ✓ 1870ms | 否 | ✓ 1210ms | 否 | ✓ 1421ms | http |
| 38.145.218.76:8448 | ✓ 700ms | ✓ 1089ms | ✓ 266ms | ✓ 695ms | ✓ 515ms | http |
| 45.136.131.25:8444 | ✓ 696ms | ✓ 632ms | ✓ 718ms | ✓ 729ms | ✓ 665ms | http |
| 38.145.220.27:8446 | ✓ 988ms | ✓ 1721ms | ✓ 333ms | ✓ 1105ms | 否 | http |
| 210.45.70.16:7895 | ✓ 1018ms | ✓ 1576ms | ✓ 1844ms | 否 | ✓ 980ms | http |
| 106.117.208.101:7890 | ✓ 1111ms | ✓ 1801ms | 否 | ✓ 1934ms | 否 | http |
| 103.113.70.189:1081 | ✓ 416ms | ✓ 1187ms | ✓ 558ms | ✓ 1165ms | ✓ 1077ms | http |
| 38.34.179.173:8452 | ✓ 937ms | ✓ 1072ms | ✓ 198ms | ✓ 777ms | ✓ 721ms | http |
| 38.145.218.229:8450 | ✓ 413ms | ✓ 1197ms | ✓ 752ms | ✓ 714ms | ✓ 614ms | http |

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
