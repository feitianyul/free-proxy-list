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

最后更新：2026-02-24 12:39:12 UTC（2026-02-24 20:39:12 UTC+8）

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
| 205.209.118.30:3138 | ✓ 1090ms | 否 | ✓ 1208ms | ✓ 1303ms | ✓ 1181ms | http |
| 103.84.95.54:7890 | ✓ 646ms | 否 | 否 | ✓ 1167ms | ✓ 729ms | http |
| 211.230.49.122:3128 | ✓ 819ms | ✓ 1461ms | ✓ 1809ms | ✓ 1251ms | ✓ 1406ms | http |
| 202.152.44.19:8081 | 否 | 否 | ✓ 1091ms | ✓ 1225ms | ✓ 994ms | http |
| 202.152.44.18:8081 | 否 | 否 | ✓ 1102ms | ✓ 1366ms | ✓ 1079ms | http |
| 14.56.107.244:3128 | ✓ 1117ms | 否 | ✓ 1397ms | 否 | ✓ 1738ms | http |
| 120.46.152.136:3128 | 否 | 否 | ✓ 1505ms | ✓ 1561ms | ✓ 1268ms | http |
| 124.16.93.70:7890 | 否 | ✓ 1069ms | ✓ 975ms | ✓ 1090ms | ✓ 851ms | http |
| 120.92.212.16:7890 | ✓ 1138ms | ✓ 1703ms | ✓ 1885ms | ✓ 1254ms | 否 | http |
| 147.45.159.213:48206 | ✓ 1303ms | 否 | 否 | ✓ 1822ms | ✓ 1321ms | http |
| 34.50.41.78:8888 | ✓ 1515ms | 否 | ✓ 1189ms | ✓ 1504ms | ✓ 1547ms | http |
| 125.128.12.94:3128 | 否 | ✓ 928ms | ✓ 1855ms | ✓ 1448ms | ✓ 984ms | http |
| 115.231.181.40:8128 | ✓ 1012ms | 否 | 否 | ✓ 1130ms | ✓ 998ms | http |
| 101.43.255.96:80 | ✓ 1213ms | ✓ 1298ms | 否 | ✓ 1294ms | ✓ 955ms | http |
| 81.70.169.194:80 | ✓ 1378ms | ✓ 1369ms | ✓ 1014ms | ✓ 1488ms | 否 | http |
| 165.227.5.10:8888 | ✓ 1825ms | 否 | 否 | ✓ 781ms | ✓ 708ms | http |
| 113.59.32.142:22222 | 否 | ✓ 1385ms | ✓ 1157ms | ✓ 1426ms | ✓ 1049ms | http |
| 195.123.209.48:3128 | ✓ 1724ms | ✓ 1957ms | 否 | 否 | ✓ 1914ms | http |
| 137.220.150.22:6005 | ✓ 1027ms | 否 | ✓ 1030ms | ✓ 1148ms | ✓ 994ms | http |
| 223.113.134.101:22222 | ✓ 1007ms | ✓ 1058ms | ✓ 715ms | 否 | 否 | http |
| 183.98.143.134:8077 | ✓ 780ms | ✓ 1381ms | ✓ 1141ms | ✓ 1012ms | ✓ 1065ms | http |
| 85.208.108.43:2094 | ✓ 514ms | 否 | ✓ 1184ms | ✓ 1520ms | ✓ 847ms | http |
| 223.113.134.105:22222 | ✓ 856ms | 否 | ✓ 672ms | 否 | ✓ 680ms | http |
| 113.59.32.161:22222 | ✓ 1100ms | ✓ 1365ms | 否 | ✓ 1266ms | ✓ 997ms | http |
| 120.232.242.119:22222 | ✓ 1195ms | ✓ 1558ms | ✓ 1260ms | 否 | ✓ 1934ms | http |
| 120.92.212.16:8890 | ✓ 953ms | 否 | ✓ 1481ms | ✓ 1236ms | ✓ 1522ms | http |
| 200.125.171.254:999 | ✓ 982ms | 否 | ✓ 1980ms | ✓ 1436ms | ✓ 1240ms | http |
| 18.229.170.122:3128 | ✓ 663ms | 否 | ✓ 1135ms | 否 | ✓ 1772ms | http |
| 152.32.255.24:27197 | ✓ 1727ms | ✓ 1993ms | 否 | 否 | ✓ 1294ms | http |
| 132.145.93.138:1080 | ✓ 1974ms | 否 | ✓ 1814ms | 否 | ✓ 1810ms | http |
| 187.216.141.46:3128 | ✓ 1507ms | 否 | ✓ 1211ms | ✓ 1974ms | ✓ 1549ms | http |
| 36.147.78.166:80 | ✓ 1675ms | ✓ 1652ms | ✓ 1637ms | ✓ 1587ms | ✓ 1621ms | http |
| 164.90.151.28:3128 | ✓ 709ms | 否 | ✓ 742ms | ✓ 714ms | ✓ 544ms | http |
| 103.82.93.100:3128 | 否 | 否 | ✓ 1380ms | ✓ 1233ms | ✓ 1019ms | http |
| 91.233.223.147:3128 | ✓ 1061ms | 否 | ✓ 1336ms | 否 | ✓ 1648ms | http |
| 181.78.77.211:999 | ✓ 1098ms | 否 | ✓ 1638ms | ✓ 1774ms | ✓ 1617ms | http |
| 181.78.77.210:999 | ✓ 1261ms | 否 | ✓ 1473ms | ✓ 1829ms | ✓ 1592ms | http |
| 78.13.231.158:3128 | ✓ 277ms | 否 | ✓ 267ms | ✓ 1114ms | ✓ 841ms | http |
| 223.113.134.84:22222 | 否 | ✓ 1697ms | ✓ 683ms | 否 | ✓ 676ms | http |
| 223.113.134.107:22222 | ✓ 656ms | ✓ 1085ms | ✓ 703ms | ✓ 894ms | ✓ 670ms | http |
| 183.249.5.213:22222 | 否 | ✓ 1088ms | ✓ 1149ms | 否 | ✓ 708ms | http |
| 117.159.239.51:22222 | ✓ 828ms | ✓ 1042ms | ✓ 878ms | ✓ 1105ms | ✓ 831ms | http |
| 121.168.58.213:3128 | ✓ 1734ms | ✓ 1852ms | ✓ 967ms | ✓ 980ms | ✓ 742ms | http |
| 113.59.32.141:22222 | ✓ 1082ms | ✓ 1372ms | ✓ 1044ms | ✓ 1346ms | ✓ 1019ms | http |
| 144.124.227.234:3128 | ✓ 607ms | ✓ 1867ms | ✓ 1355ms | ✓ 1886ms | ✓ 1179ms | http |
| 121.230.8.136:1080 | ✓ 1167ms | ✓ 1491ms | ✓ 1613ms | 否 | ✓ 1437ms | http |
| 121.230.8.34:1080 | ✓ 1279ms | 否 | ✓ 1931ms | ✓ 1393ms | 否 | http |
| 121.230.8.97:1080 | ✓ 1501ms | 否 | ✓ 1130ms | 否 | ✓ 1154ms | http |
| 222.184.48.236:22222 | ✓ 1344ms | ✓ 1795ms | ✓ 1695ms | ✓ 1304ms | ✓ 1993ms | http |
| 70.61.188.34:3128 | ✓ 876ms | ✓ 1774ms | ✓ 1844ms | ✓ 1667ms | ✓ 1300ms | http |
| 35.234.17.221:8080 | ✓ 1699ms | 否 | ✓ 1307ms | ✓ 1217ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1035ms | ✓ 1433ms | ✓ 929ms | ✓ 1413ms | ✓ 1046ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1258ms | 否 | ✓ 1366ms | ✓ 1055ms | http |
| 121.43.196.213:8222 | ✓ 939ms | ✓ 1034ms | ✓ 881ms | ✓ 1129ms | ✓ 847ms | http |
| 121.43.196.210:8222 | ✓ 927ms | ✓ 1062ms | ✓ 917ms | ✓ 1077ms | ✓ 842ms | http |
| 114.55.226.123:10086 | ✓ 1007ms | ✓ 1353ms | ✓ 1027ms | ✓ 1289ms | ✓ 977ms | http |
| 113.45.250.180:443 | ✓ 1060ms | 否 | ✓ 1000ms | ✓ 1165ms | 否 | http |
| 139.159.99.242:8080 | 否 | 否 | ✓ 830ms | ✓ 1022ms | ✓ 1569ms | http |
| 113.59.32.162:22222 | ✓ 1317ms | ✓ 1349ms | 否 | ✓ 1308ms | ✓ 948ms | http |
| 217.216.109.116:8080 | ✓ 866ms | 否 | ✓ 1060ms | ✓ 1498ms | ✓ 1613ms | http |
| 14.56.177.2:3128 | 否 | ✓ 1648ms | ✓ 1248ms | ✓ 1632ms | ✓ 1400ms | http |
| 59.46.216.131:30001 | ✓ 1112ms | 否 | ✓ 1078ms | 否 | ✓ 1032ms | http |
| 217.76.245.80:999 | ✓ 828ms | ✓ 1700ms | ✓ 1283ms | ✓ 1675ms | ✓ 1715ms | http |
| 14.56.177.34:3128 | 否 | 否 | ✓ 1796ms | ✓ 1594ms | ✓ 879ms | http |
| 18.229.201.117:3128 | ✓ 677ms | 否 | ✓ 795ms | 否 | ✓ 1631ms | http |
| 52.188.28.218:3128 | ✓ 407ms | 否 | 否 | ✓ 1142ms | ✓ 861ms | http |
| 62.113.119.14:8080 | ✓ 792ms | 否 | ✓ 1292ms | ✓ 1635ms | ✓ 1189ms | http |
| 223.113.134.98:22222 | ✓ 676ms | 否 | ✓ 724ms | 否 | ✓ 667ms | http |
| 138.124.53.25:7443 | ✓ 1189ms | 否 | ✓ 1970ms | 否 | ✓ 1587ms | http |
| 45.140.147.155:1081 | ✓ 645ms | ✓ 1608ms | ✓ 1184ms | ✓ 1672ms | ✓ 1284ms | http |
| 45.140.147.155:1082 | ✓ 646ms | ✓ 1516ms | 否 | ✓ 1446ms | ✓ 1004ms | http |
| 83.219.250.8:62920 | ✓ 1007ms | ✓ 1372ms | 否 | 否 | ✓ 1582ms | http |
| 36.147.78.166:443 | 否 | ✓ 1653ms | ✓ 1665ms | ✓ 1860ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1677ms | 否 | ✓ 1681ms | ✓ 1904ms | 否 | http |
| 14.56.118.24:3128 | ✓ 1464ms | ✓ 1397ms | ✓ 782ms | ✓ 953ms | ✓ 757ms | http |
| 14.56.118.34:3128 | ✓ 1445ms | ✓ 1038ms | ✓ 1215ms | ✓ 1003ms | ✓ 733ms | http |
| 117.159.239.54:22222 | ✓ 850ms | ✓ 1093ms | ✓ 887ms | 否 | ✓ 1117ms | http |
| 35.225.22.61:80 | ✓ 1232ms | ✓ 1677ms | ✓ 1547ms | ✓ 1228ms | ✓ 857ms | http |
| 85.208.108.43:10808 | ✓ 697ms | 否 | ✓ 564ms | ✓ 1615ms | 否 | http |
| 14.56.118.64:3128 | ✓ 635ms | ✓ 1620ms | ✓ 657ms | 否 | 否 | http |
| 103.188.252.65:1234 | ✓ 1538ms | 否 | ✓ 1669ms | ✓ 1548ms | 否 | http |
| 14.56.177.162:3128 | ✓ 1068ms | ✓ 922ms | ✓ 1433ms | ✓ 1417ms | 否 | http |
| 14.56.118.114:3128 | ✓ 1588ms | 否 | ✓ 1812ms | ✓ 1531ms | ✓ 1924ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1053ms | 否 | ✓ 1347ms | ✓ 1161ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1824ms | ✓ 1603ms | ✓ 1513ms | http |
| 183.249.5.117:22222 | ✓ 937ms | ✓ 950ms | 否 | ✓ 1189ms | ✓ 1226ms | http |
| 178.253.22.108:65431 | ✓ 901ms | 否 | ✓ 1394ms | ✓ 1541ms | ✓ 1323ms | http |
| 223.113.134.57:22222 | ✓ 765ms | ✓ 846ms | ✓ 1008ms | ✓ 1208ms | ✓ 735ms | http |
| 14.56.118.244:3128 | ✓ 639ms | 否 | ✓ 667ms | ✓ 980ms | ✓ 756ms | http |
| 223.113.134.141:22222 | ✓ 830ms | ✓ 1091ms | ✓ 685ms | ✓ 1021ms | 否 | http |
| 217.77.102.18:3128 | ✓ 1383ms | 否 | ✓ 1777ms | 否 | ✓ 1900ms | http |

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
