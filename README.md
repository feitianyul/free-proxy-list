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

最后更新：2026-03-12 12:28:31 UTC（2026-03-12 20:28:31 UTC+8）

**代理总数：71**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.175:8443 | ✓ 1314ms | ✓ 1077ms | ✓ 1149ms | ✓ 905ms | ✓ 808ms | http |
| 45.136.130.188:8443 | ✓ 1307ms | ✓ 840ms | ✓ 941ms | 否 | ✓ 700ms | http |
| 35.225.22.61:80 | ✓ 1307ms | ✓ 1890ms | ✓ 919ms | ✓ 1074ms | ✓ 1156ms | http |
| 45.136.131.47:8443 | ✓ 1308ms | ✓ 1986ms | ✓ 984ms | ✓ 1295ms | ✓ 678ms | http |
| 45.136.131.63:8443 | ✓ 1307ms | ✓ 1055ms | ✓ 726ms | ✓ 1438ms | ✓ 709ms | http |
| 1.231.81.166:3128 | ✓ 1227ms | ✓ 1207ms | ✓ 996ms | ✓ 1141ms | ✓ 867ms | http |
| 194.213.18.200:443 | ✓ 1064ms | ✓ 1810ms | ✓ 1367ms | ✓ 1965ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1886ms | ✓ 1598ms | ✓ 1518ms | ✓ 1446ms | ✓ 1178ms | http |
| 164.90.151.28:3128 | ✓ 1062ms | ✓ 1124ms | ✓ 1108ms | ✓ 941ms | ✓ 715ms | http |
| 103.166.185.54:3128 | ✓ 1619ms | 否 | ✓ 1551ms | ✓ 1569ms | ✓ 1155ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1629ms | ✓ 1677ms | ✓ 1160ms | http |
| 24.199.124.151:3128 | ✓ 1066ms | ✓ 1237ms | ✓ 1206ms | ✓ 977ms | ✓ 1331ms | http |
| 103.84.95.54:7890 | ✓ 1531ms | 否 | ✓ 867ms | 否 | ✓ 1212ms | http |
| 91.107.141.42:8081 | ✓ 1321ms | 否 | ✓ 1549ms | 否 | ✓ 1612ms | http |
| 120.92.212.16:7890 | ✓ 1886ms | 否 | ✓ 1428ms | ✓ 1743ms | ✓ 1161ms | http |
| 45.136.130.223:8443 | ✓ 798ms | ✓ 1159ms | ✓ 357ms | ✓ 914ms | ✓ 1074ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1973ms | 否 | ✓ 979ms | ✓ 731ms | http |
| 121.237.181.137:8888 | ✓ 1516ms | ✓ 1403ms | 否 | ✓ 1505ms | ✓ 1116ms | http |
| 120.92.212.16:8890 | ✓ 1191ms | ✓ 1772ms | ✓ 1944ms | ✓ 1868ms | ✓ 1621ms | http |
| 205.209.118.30:3138 | ✓ 935ms | 否 | ✓ 745ms | ✓ 1051ms | ✓ 1025ms | http |
| 46.183.25.8:443 | ✓ 792ms | 否 | ✓ 540ms | ✓ 1349ms | 否 | http |
| 172.212.68.37:3128 | 否 | 否 | ✓ 525ms | ✓ 1359ms | ✓ 1040ms | http |
| 162.248.165.72:1080 | ✓ 893ms | ✓ 1370ms | 否 | ✓ 1712ms | ✓ 1334ms | http |
| 103.3.246.71:3128 | ✓ 1383ms | 否 | ✓ 1268ms | ✓ 1525ms | ✓ 1280ms | http |
| 152.42.213.210:8080 | ✓ 1675ms | 否 | ✓ 1728ms | ✓ 1453ms | ✓ 1152ms | http |
| 185.115.74.185:8080 | ✓ 704ms | ✓ 1907ms | ✓ 1457ms | 否 | 否 | http |
| 171.251.172.78:5106 | ✓ 1512ms | 否 | 否 | ✓ 1818ms | ✓ 1608ms | http |
| 45.136.130.191:8443 | ✓ 861ms | ✓ 890ms | ✓ 484ms | ✓ 882ms | ✓ 690ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1550ms | ✓ 1728ms | ✓ 1603ms | http |
| 45.140.147.82:1082 | ✓ 450ms | ✓ 1553ms | ✓ 1168ms | ✓ 1870ms | ✓ 1196ms | http |
| 45.140.147.82:1081 | ✓ 446ms | 否 | ✓ 721ms | ✓ 1869ms | ✓ 1187ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1541ms | ✓ 1574ms | ✓ 1574ms | 否 | http |
| 47.101.159.19:8899 | ✓ 1482ms | ✓ 1998ms | ✓ 1295ms | ✓ 1346ms | ✓ 1165ms | http |
| 111.48.191.1:7890 | ✓ 954ms | ✓ 1175ms | ✓ 920ms | ✓ 1206ms | ✓ 934ms | http |
| 39.104.201.40:7890 | ✓ 1101ms | ✓ 1770ms | 否 | ✓ 1484ms | ✓ 1152ms | http |
| 59.46.216.131:30001 | ✓ 1532ms | 否 | ✓ 1399ms | ✓ 1679ms | ✓ 1303ms | http |
| 121.204.158.249:3128 | 否 | 否 | ✓ 1314ms | ✓ 1495ms | ✓ 1201ms | http |
| 45.140.147.155:1081 | ✓ 943ms | 否 | ✓ 1136ms | ✓ 1787ms | ✓ 1338ms | http |
| 24.144.86.173:1080 | ✓ 1018ms | ✓ 1587ms | ✓ 1155ms | ✓ 891ms | ✓ 807ms | http |
| 171.251.172.78:5107 | 否 | 否 | ✓ 1893ms | ✓ 1901ms | ✓ 1661ms | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1687ms | ✓ 1398ms | ✓ 1204ms | http |
| 171.251.172.78:5104 | ✓ 1775ms | 否 | ✓ 1842ms | 否 | ✓ 1610ms | http |
| 168.235.110.63:3128 | ✓ 281ms | 否 | ✓ 1765ms | ✓ 957ms | ✓ 715ms | http |
| 178.236.245.17:3128 | ✓ 661ms | ✓ 1698ms | ✓ 869ms | ✓ 1596ms | ✓ 1358ms | http |
| 178.236.245.59:3128 | ✓ 619ms | 否 | ✓ 572ms | ✓ 1624ms | 否 | http |
| 138.124.53.221:443 | ✓ 407ms | 否 | ✓ 1210ms | ✓ 1881ms | ✓ 1962ms | http |
| 124.16.111.161:7890 | ✓ 1080ms | ✓ 1334ms | ✓ 1195ms | ✓ 1332ms | ✓ 1481ms | http |
| 101.43.255.96:80 | ✓ 1268ms | ✓ 1605ms | ✓ 1172ms | ✓ 1435ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1065ms | 否 | ✓ 987ms | ✓ 1451ms | 否 | http |
| 150.249.255.91:3128 | ✓ 783ms | 否 | ✓ 1626ms | ✓ 1080ms | ✓ 897ms | http |
| 88.80.150.82:8080 | ✓ 1199ms | 否 | ✓ 1983ms | 否 | ✓ 1686ms | https |
| 192.71.213.85:9812 | ✓ 1409ms | 否 | ✓ 1557ms | ✓ 1788ms | 否 | http |
| 107.173.52.58:7890 | ✓ 1852ms | 否 | ✓ 944ms | ✓ 1624ms | ✓ 1041ms | http |
| 8.219.97.248:80 | ✓ 1196ms | 否 | ✓ 1192ms | ✓ 1862ms | 否 | http |
| 14.225.222.185:7890 | ✓ 1282ms | 否 | ✓ 1665ms | 否 | ✓ 1486ms | http |
| 46.39.105.157:8080 | ✓ 686ms | 否 | ✓ 1819ms | 否 | ✓ 1602ms | http |
| 180.127.149.244:1080 | 否 | ✓ 1479ms | ✓ 1116ms | 否 | ✓ 1174ms | http |
| 158.69.185.37:3129 | ✓ 597ms | 否 | ✓ 140ms | ✓ 1019ms | ✓ 770ms | http |
| 62.113.119.14:8080 | ✓ 938ms | ✓ 1496ms | ✓ 859ms | 否 | ✓ 1263ms | http |
| 45.136.198.40:3128 | ✓ 969ms | ✓ 1756ms | ✓ 859ms | ✓ 1874ms | ✓ 1673ms | http |
| 83.219.250.8:62920 | ✓ 1204ms | 否 | ✓ 1596ms | ✓ 1949ms | ✓ 1465ms | http |
| 45.129.141.143:3128 | ✓ 1352ms | ✓ 1875ms | ✓ 1775ms | ✓ 1795ms | ✓ 1626ms | http |
| 103.86.131.62:80 | ✓ 1711ms | 否 | 否 | ✓ 1521ms | ✓ 1219ms | http |
| 103.113.70.189:1081 | 否 | ✓ 953ms | 否 | ✓ 1072ms | ✓ 1518ms | http |
| 165.227.5.10:8888 | ✓ 1499ms | ✓ 1608ms | ✓ 795ms | ✓ 944ms | 否 | http |
| 159.223.42.219:3128 | ✓ 1682ms | 否 | ✓ 1683ms | ✓ 1409ms | ✓ 1061ms | http |
| 111.79.111.126:3128 | 否 | 否 | ✓ 1977ms | ✓ 1401ms | ✓ 1076ms | http |
| 103.39.51.190:8080 | ✓ 1931ms | 否 | 否 | ✓ 1911ms | ✓ 1575ms | http |
| 61.52.131.172:8443 | ✓ 1087ms | ✓ 1334ms | ✓ 1145ms | ✓ 1323ms | 否 | http |
| 81.70.169.194:80 | 否 | 否 | ✓ 1312ms | ✓ 1430ms | ✓ 1933ms | http |
| 190.9.109.198:999 | ✓ 1479ms | ✓ 1517ms | ✓ 1157ms | ✓ 1438ms | ✓ 1004ms | http |

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
