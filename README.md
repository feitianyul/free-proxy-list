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

最后更新：2026-03-12 11:36:10 UTC（2026-03-12 19:36:10 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.188:8443 | 否 | 否 | ✓ 530ms | ✓ 887ms | ✓ 674ms | http |
| 46.183.25.8:443 | ✓ 1399ms | 否 | ✓ 1476ms | ✓ 1297ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1449ms | ✓ 1159ms | ✓ 1902ms | ✓ 1138ms | ✓ 886ms | http |
| 115.231.181.40:8128 | ✓ 1756ms | 否 | ✓ 1141ms | ✓ 1824ms | ✓ 1146ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1849ms | ✓ 1568ms | ✓ 1499ms | ✓ 1148ms | http |
| 202.141.161.53:10808 | ✓ 1156ms | ✓ 1542ms | ✓ 1439ms | ✓ 1322ms | ✓ 1268ms | http |
| 178.236.245.59:3128 | ✓ 607ms | 否 | ✓ 1024ms | ✓ 1763ms | ✓ 1574ms | http |
| 178.236.245.17:3128 | ✓ 649ms | ✓ 1676ms | ✓ 1239ms | ✓ 1865ms | ✓ 1608ms | http |
| 120.92.212.16:7890 | ✓ 1136ms | ✓ 1435ms | ✓ 1151ms | ✓ 1462ms | ✓ 1163ms | http |
| 211.171.114.154:3128 | ✓ 1665ms | ✓ 1657ms | 否 | ✓ 1622ms | ✓ 1432ms | http |
| 194.213.18.200:443 | ✓ 402ms | ✓ 1167ms | ✓ 164ms | ✓ 1878ms | ✓ 879ms | http |
| 45.136.130.191:8443 | 否 | ✓ 879ms | ✓ 247ms | ✓ 904ms | ✓ 818ms | http |
| 45.136.130.175:8443 | 否 | ✓ 887ms | ✓ 845ms | ✓ 1023ms | ✓ 685ms | http |
| 45.136.131.47:8443 | 否 | ✓ 1767ms | ✓ 270ms | ✓ 904ms | ✓ 855ms | http |
| 45.136.131.63:8443 | 否 | ✓ 1875ms | ✓ 1806ms | ✓ 1964ms | ✓ 698ms | http |
| 107.173.52.58:7890 | ✓ 721ms | 否 | ✓ 1233ms | ✓ 1843ms | ✓ 1149ms | http |
| 91.107.141.42:8081 | ✓ 1692ms | 否 | ✓ 1236ms | 否 | ✓ 1706ms | http |
| 107.155.65.87:13428 | ✓ 942ms | 否 | 否 | ✓ 1271ms | ✓ 1053ms | http |
| 168.235.110.63:3128 | ✓ 922ms | ✓ 1184ms | ✓ 1784ms | ✓ 1595ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1160ms | ✓ 1462ms | ✓ 1113ms | 否 | 否 | http |
| 45.136.130.223:8443 | ✓ 380ms | ✓ 912ms | ✓ 869ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 519ms | ✓ 1500ms | ✓ 660ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 556ms | ✓ 1928ms | ✓ 265ms | 否 | 否 | http |
| 121.237.181.137:8888 | ✓ 1156ms | 否 | ✓ 1093ms | ✓ 1886ms | ✓ 1024ms | http |
| 171.251.172.78:5106 | ✓ 1743ms | 否 | 否 | ✓ 1841ms | ✓ 1647ms | http |
| 45.168.238.193:8443 | ✓ 421ms | ✓ 1875ms | ✓ 418ms | 否 | ✓ 868ms | http |
| 81.70.169.194:80 | ✓ 1127ms | ✓ 1571ms | ✓ 1157ms | ✓ 1791ms | 否 | http |
| 101.43.255.96:80 | ✓ 1315ms | 否 | 否 | ✓ 1574ms | ✓ 1162ms | http |
| 171.251.172.78:5104 | 否 | 否 | ✓ 1729ms | ✓ 1809ms | ✓ 1652ms | http |
| 103.113.70.189:1081 | ✓ 315ms | ✓ 1376ms | 否 | ✓ 1247ms | ✓ 805ms | http |
| 94.72.109.169:8080 | ✓ 844ms | 否 | ✓ 1510ms | ✓ 1630ms | ✓ 1570ms | http |
| 124.16.111.161:7890 | ✓ 1111ms | ✓ 1331ms | ✓ 1213ms | ✓ 1367ms | ✓ 1061ms | http |
| 8.219.97.248:80 | ✓ 1156ms | 否 | ✓ 1569ms | ✓ 1466ms | 否 | http |
| 158.69.185.37:3129 | ✓ 834ms | ✓ 1587ms | ✓ 872ms | ✓ 1251ms | ✓ 814ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1764ms | ✓ 1488ms | ✓ 1254ms | http |
| 160.250.5.22:1 | ✓ 1087ms | 否 | ✓ 1777ms | ✓ 1537ms | ✓ 1202ms | http |
| 91.233.223.147:3128 | 否 | 否 | ✓ 1237ms | ✓ 1953ms | ✓ 1531ms | http |
| 111.48.191.1:7890 | ✓ 921ms | ✓ 1148ms | ✓ 942ms | ✓ 1178ms | ✓ 1013ms | http |
| 14.225.222.185:7890 | 否 | ✓ 1890ms | ✓ 1389ms | ✓ 1580ms | 否 | http |
| 103.82.23.118:5207 | ✓ 1626ms | 否 | ✓ 1700ms | 否 | ✓ 1790ms | http |
| 59.46.216.131:30001 | ✓ 1156ms | 否 | ✓ 1892ms | 否 | ✓ 1316ms | http |
| 171.251.172.78:5107 | 否 | 否 | ✓ 1974ms | ✓ 1888ms | ✓ 1641ms | http |
| 47.101.149.27:9010 | ✓ 1595ms | 否 | 否 | ✓ 1868ms | ✓ 1571ms | http |
| 137.184.6.117:3128 | ✓ 504ms | ✓ 1102ms | ✓ 1211ms | ✓ 917ms | ✓ 751ms | http |
| 27.133.238.94:80 | ✓ 1763ms | ✓ 1620ms | ✓ 1517ms | ✓ 1367ms | ✓ 878ms | http |
| 205.209.118.30:3138 | ✓ 1364ms | 否 | ✓ 1724ms | 否 | ✓ 1965ms | http |
| 202.155.12.161:443 | ✓ 1605ms | ✓ 1632ms | ✓ 1469ms | ✓ 1336ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1003ms | ✓ 1087ms | ✓ 843ms | http |
| 88.80.150.82:8080 | ✓ 1290ms | ✓ 1913ms | ✓ 1997ms | ✓ 1923ms | ✓ 1539ms | https |
| 45.186.6.104:3128 | ✓ 1069ms | ✓ 1823ms | ✓ 1941ms | 否 | 否 | http |
| 165.227.5.10:8888 | ✓ 302ms | 否 | 否 | ✓ 1414ms | ✓ 773ms | http |
| 47.77.193.180:1080 | ✓ 772ms | ✓ 1737ms | ✓ 355ms | ✓ 909ms | ✓ 696ms | http |
| 38.180.2.107:3128 | ✓ 1424ms | 否 | ✓ 1940ms | ✓ 1999ms | ✓ 1833ms | http |
| 162.248.165.72:1080 | ✓ 1355ms | 否 | ✓ 1676ms | ✓ 1980ms | 否 | http |
| 106.117.208.101:7890 | 否 | ✓ 1576ms | ✓ 1270ms | ✓ 1669ms | ✓ 1189ms | http |
| 45.136.198.40:3128 | ✓ 1036ms | ✓ 1431ms | ✓ 1346ms | ✓ 1992ms | ✓ 1344ms | http |
| 34.101.184.164:3128 | ✓ 1190ms | 否 | ✓ 1202ms | ✓ 1863ms | ✓ 1463ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1022ms | 否 | ✓ 981ms | ✓ 703ms | http |
| 39.104.201.40:7890 | ✓ 1110ms | ✓ 1454ms | ✓ 1142ms | ✓ 1445ms | ✓ 1148ms | http |
| 152.70.98.46:8888 | ✓ 1673ms | 否 | ✓ 1917ms | ✓ 1254ms | ✓ 965ms | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1349ms | ✓ 1266ms | ✓ 1006ms | http |
| 103.166.185.54:3128 | ✓ 1421ms | ✓ 1859ms | ✓ 1313ms | ✓ 1497ms | ✓ 1194ms | http |
| 162.240.154.26:3128 | ✓ 738ms | ✓ 1316ms | ✓ 1089ms | ✓ 1326ms | ✓ 961ms | http |
| 103.76.108.227:80 | 否 | 否 | ✓ 1907ms | ✓ 1911ms | ✓ 1702ms | http |
| 109.234.38.35:3128 | 否 | 否 | ✓ 1148ms | ✓ 1482ms | ✓ 1385ms | http |
| 152.42.213.210:443 | ✓ 930ms | 否 | ✓ 965ms | ✓ 1288ms | ✓ 1402ms | http |
| 45.140.147.82:1081 | ✓ 406ms | 否 | ✓ 1052ms | ✓ 1717ms | ✓ 1210ms | http |
| 202.129.206.239:3128 | ✓ 1791ms | 否 | ✓ 1833ms | 否 | ✓ 1434ms | http |
| 43.167.227.161:1080 | 否 | 否 | ✓ 1252ms | ✓ 1156ms | ✓ 1587ms | http |
| 45.140.147.155:1081 | ✓ 415ms | ✓ 1424ms | ✓ 1166ms | ✓ 1811ms | ✓ 1563ms | http |
| 61.52.131.172:8443 | ✓ 1135ms | ✓ 1359ms | ✓ 1058ms | ✓ 1405ms | ✓ 1104ms | http |
| 164.90.151.28:3128 | ✓ 757ms | ✓ 1079ms | ✓ 1164ms | ✓ 1019ms | ✓ 739ms | http |
| 111.79.111.126:3128 | 否 | 否 | ✓ 1659ms | ✓ 1598ms | ✓ 1342ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1501ms | ✓ 1428ms | ✓ 1125ms | http |
| 172.212.68.37:3128 | ✓ 204ms | 否 | 否 | ✓ 1231ms | ✓ 909ms | http |
| 113.255.59.226:8080 | ✓ 1409ms | 否 | 否 | ✓ 1387ms | ✓ 1417ms | http |
| 103.3.246.71:3128 | ✓ 1424ms | 否 | ✓ 1292ms | ✓ 1405ms | ✓ 1117ms | http |

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
