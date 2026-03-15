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

最后更新：2026-03-15 07:55:00 UTC（2026-03-15 15:55:00 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 305ms | 否 | ✓ 1212ms | ✓ 1245ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1594ms | 否 | 否 | ✓ 1919ms | ✓ 1507ms | http |
| 45.140.147.155:1081 | ✓ 637ms | ✓ 1181ms | ✓ 1398ms | ✓ 1616ms | ✓ 1997ms | http |
| 38.145.203.135:8443 | 否 | ✓ 1963ms | ✓ 322ms | ✓ 1076ms | ✓ 1474ms | http |
| 113.160.132.26:8080 | ✓ 1473ms | ✓ 1411ms | 否 | ✓ 1269ms | ✓ 1119ms | http |
| 34.96.238.40:8080 | ✓ 1285ms | ✓ 1541ms | ✓ 1755ms | ✓ 1081ms | 否 | http |
| 86.53.183.16:1080 | ✓ 660ms | ✓ 1879ms | ✓ 1664ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 942ms | ✓ 1253ms | ✓ 914ms | http |
| 72.11.150.178:6005 | ✓ 739ms | ✓ 1407ms | ✓ 1641ms | 否 | ✓ 1054ms | http |
| 38.55.105.94:6005 | ✓ 1155ms | 否 | ✓ 1168ms | ✓ 1030ms | ✓ 820ms | http |
| 137.220.150.152:6005 | ✓ 1569ms | 否 | ✓ 970ms | ✓ 1165ms | ✓ 972ms | http |
| 137.220.150.104:6005 | ✓ 1570ms | 否 | ✓ 1051ms | ✓ 1223ms | ✓ 930ms | http |
| 137.220.151.110:6005 | ✓ 1577ms | 否 | ✓ 1614ms | ✓ 1253ms | ✓ 932ms | http |
| 165.227.5.10:8888 | ✓ 1440ms | 否 | ✓ 1702ms | 否 | ✓ 1144ms | http |
| 81.70.169.194:80 | 否 | ✓ 1620ms | ✓ 1424ms | 否 | ✓ 1966ms | http |
| 120.92.212.16:7890 | ✓ 1065ms | ✓ 1290ms | 否 | ✓ 1302ms | ✓ 1051ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1542ms | ✓ 1106ms | ✓ 1302ms | ✓ 1271ms | http |
| 43.167.227.161:1080 | ✓ 1846ms | 否 | ✓ 1124ms | 否 | ✓ 845ms | http |
| 101.43.255.96:80 | ✓ 1077ms | ✓ 1384ms | ✓ 1580ms | ✓ 1383ms | ✓ 1925ms | http |
| 62.60.177.204:34094 | ✓ 272ms | 否 | ✓ 846ms | ✓ 1016ms | ✓ 779ms | http |
| 64.188.90.36:1080 | ✓ 679ms | ✓ 1777ms | ✓ 776ms | ✓ 1888ms | ✓ 1163ms | http |
| 85.198.96.242:3128 | ✓ 616ms | 否 | ✓ 1800ms | ✓ 1696ms | 否 | http |
| 38.55.106.208:6005 | ✓ 1177ms | 否 | ✓ 1130ms | ✓ 1103ms | ✓ 838ms | http |
| 92.119.127.212:6005 | ✓ 1028ms | 否 | ✓ 1611ms | ✓ 1798ms | ✓ 1867ms | http |
| 212.192.13.76:6005 | ✓ 1394ms | 否 | ✓ 1555ms | ✓ 1546ms | ✓ 1060ms | http |
| 116.105.21.153:9045 | ✓ 1221ms | 否 | ✓ 1310ms | ✓ 1668ms | ✓ 1469ms | http |
| 45.136.198.40:3128 | ✓ 779ms | ✓ 1599ms | 否 | 否 | ✓ 1670ms | http |
| 45.136.131.28:8447 | 否 | 否 | ✓ 880ms | ✓ 1173ms | ✓ 823ms | http |
| 146.19.128.135:1080 | ✓ 707ms | 否 | ✓ 763ms | ✓ 1603ms | ✓ 1381ms | http |
| 1.225.116.115:1080 | ✓ 1397ms | 否 | ✓ 1827ms | ✓ 1382ms | ✓ 1596ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1704ms | ✓ 1407ms | ✓ 1668ms | ✓ 1439ms | http |
| 38.34.183.130:8443 | ✓ 817ms | ✓ 1870ms | ✓ 304ms | ✓ 990ms | ✓ 728ms | http |
| 45.136.130.155:8443 | ✓ 420ms | ✓ 904ms | ✓ 332ms | ✓ 968ms | ✓ 744ms | http |
| 45.136.130.162:8443 | ✓ 421ms | ✓ 1838ms | ✓ 310ms | ✓ 925ms | ✓ 728ms | http |
| 38.145.220.25:8447 | ✓ 906ms | ✓ 1149ms | ✓ 344ms | ✓ 945ms | ✓ 874ms | http |
| 43.165.195.107:3128 | ✓ 1549ms | 否 | ✓ 1269ms | ✓ 1260ms | ✓ 1090ms | http |
| 88.80.150.82:8080 | ✓ 1965ms | 否 | ✓ 844ms | 否 | ✓ 1988ms | https |
| 124.16.111.161:7890 | ✓ 925ms | ✓ 1219ms | ✓ 1101ms | ✓ 1195ms | ✓ 960ms | http |
| 121.230.9.26:1080 | ✓ 1084ms | ✓ 1983ms | 否 | 否 | ✓ 1532ms | http |
| 178.115.230.100:8080 | ✓ 1892ms | 否 | ✓ 1694ms | 否 | ✓ 1616ms | http |
| 24.144.86.173:1080 | ✓ 376ms | ✓ 989ms | ✓ 1543ms | ✓ 873ms | ✓ 760ms | http |
| 209.126.10.139:3128 | 否 | ✓ 1240ms | ✓ 1104ms | ✓ 983ms | ✓ 1539ms | http |
| 38.145.218.101:8447 | 否 | ✓ 1096ms | ✓ 486ms | ✓ 939ms | ✓ 747ms | http |
| 45.136.130.214:8443 | 否 | ✓ 1113ms | ✓ 464ms | ✓ 969ms | ✓ 806ms | http |
| 45.136.130.215:8443 | 否 | ✓ 1646ms | ✓ 307ms | ✓ 1975ms | ✓ 738ms | http |
| 64.31.49.174:3128 | ✓ 1015ms | 否 | 否 | ✓ 1335ms | ✓ 1896ms | http |
| 150.249.255.91:3128 | ✓ 1597ms | ✓ 1746ms | ✓ 1915ms | 否 | 否 | http |
| 47.101.149.27:9010 | 否 | ✓ 1330ms | 否 | ✓ 1564ms | ✓ 1613ms | http |
| 172.212.68.37:3128 | ✓ 467ms | 否 | ✓ 808ms | ✓ 1330ms | ✓ 1670ms | http |
| 45.136.130.160:8443 | 否 | ✓ 1644ms | ✓ 548ms | ✓ 1071ms | 否 | http |
| 45.136.130.159:8443 | 否 | ✓ 1097ms | ✓ 1091ms | ✓ 1089ms | 否 | http |
| 106.14.203.63:3333 | 否 | ✓ 1129ms | ✓ 1881ms | ✓ 1168ms | 否 | http |
| 38.145.218.82:8443 | ✓ 1064ms | ✓ 1018ms | ✓ 1020ms | ✓ 1179ms | ✓ 1959ms | http |
| 164.92.148.68:3128 | ✓ 566ms | ✓ 1888ms | ✓ 1970ms | ✓ 1863ms | ✓ 1502ms | http |
| 210.223.44.230:3128 | ✓ 1460ms | 否 | 否 | ✓ 1279ms | ✓ 971ms | http |
| 103.84.95.54:7890 | ✓ 756ms | 否 | ✓ 939ms | 否 | ✓ 1320ms | http |
| 211.171.114.154:3128 | ✓ 1715ms | 否 | ✓ 1823ms | 否 | ✓ 1213ms | http |
| 101.43.127.100:8877 | ✓ 973ms | ✓ 1096ms | ✓ 1318ms | 否 | ✓ 1715ms | http |
| 45.136.131.42:8447 | ✓ 981ms | ✓ 1113ms | 否 | ✓ 932ms | ✓ 947ms | http |
| 103.39.51.190:8080 | ✓ 1616ms | 否 | ✓ 1479ms | ✓ 1395ms | ✓ 1723ms | http |
| 45.136.130.230:8443 | ✓ 561ms | ✓ 1647ms | ✓ 417ms | ✓ 960ms | ✓ 868ms | http |
| 38.145.208.132:8443 | ✓ 561ms | ✓ 1121ms | ✓ 620ms | ✓ 1138ms | ✓ 1026ms | http |
| 45.136.130.217:8443 | ✓ 614ms | ✓ 1866ms | ✓ 334ms | ✓ 954ms | ✓ 743ms | http |
| 45.136.130.231:8443 | ✓ 299ms | ✓ 897ms | ✓ 320ms | ✓ 962ms | ✓ 709ms | http |
| 38.145.208.103:8443 | ✓ 325ms | ✓ 931ms | ✓ 450ms | ✓ 960ms | ✓ 755ms | http |
| 38.145.208.137:8443 | ✓ 322ms | 否 | ✓ 335ms | ✓ 1879ms | 否 | http |
| 38.145.208.138:8447 | ✓ 427ms | 否 | ✓ 356ms | ✓ 1073ms | ✓ 767ms | http |
| 38.145.208.102:8443 | ✓ 329ms | 否 | ✓ 318ms | ✓ 990ms | ✓ 722ms | http |
| 59.46.216.131:30001 | ✓ 1133ms | ✓ 1524ms | ✓ 1492ms | ✓ 1869ms | ✓ 1157ms | http |
| 95.3.9.78:8080 | ✓ 1587ms | 否 | 否 | ✓ 1655ms | ✓ 1272ms | http |
| 45.136.130.233:8443 | ✓ 673ms | ✓ 932ms | ✓ 358ms | ✓ 1243ms | ✓ 727ms | http |
| 45.136.130.232:8443 | ✓ 638ms | ✓ 959ms | ✓ 349ms | ✓ 1206ms | ✓ 753ms | http |
| 38.145.208.37:8443 | ✓ 626ms | ✓ 970ms | ✓ 334ms | ✓ 1200ms | ✓ 757ms | http |
| 38.145.203.235:8443 | ✓ 637ms | ✓ 1003ms | ✓ 364ms | ✓ 1158ms | ✓ 741ms | http |
| 38.145.208.135:8443 | ✓ 657ms | ✓ 1468ms | ✓ 316ms | ✓ 935ms | ✓ 732ms | http |
| 45.136.130.156:8443 | ✓ 645ms | ✓ 1323ms | ✓ 348ms | ✓ 1201ms | ✓ 712ms | http |
| 45.136.130.223:8443 | ✓ 686ms | ✓ 970ms | ✓ 912ms | ✓ 974ms | ✓ 734ms | http |
| 45.136.130.158:8443 | ✓ 668ms | ✓ 885ms | ✓ 781ms | ✓ 1197ms | ✓ 739ms | http |
| 38.145.203.162:8443 | ✓ 689ms | ✓ 1394ms | ✓ 544ms | ✓ 1137ms | ✓ 765ms | http |
| 38.145.203.246:8443 | ✓ 649ms | ✓ 1838ms | ✓ 317ms | ✓ 984ms | ✓ 744ms | http |
| 38.145.208.94:8443 | ✓ 656ms | ✓ 1379ms | ✓ 548ms | ✓ 1154ms | ✓ 777ms | http |
| 38.145.208.96:8443 | ✓ 640ms | ✓ 1372ms | ✓ 572ms | ✓ 1121ms | ✓ 785ms | http |
| 38.145.208.95:8443 | ✓ 674ms | ✓ 1764ms | ✓ 371ms | ✓ 1008ms | ✓ 753ms | http |
| 38.145.208.93:8443 | ✓ 649ms | ✓ 1908ms | ✓ 351ms | ✓ 984ms | ✓ 759ms | http |
| 38.145.208.98:8443 | ✓ 675ms | 否 | ✓ 339ms | ✓ 1001ms | ✓ 742ms | http |
| 38.145.203.161:8443 | ✓ 684ms | 否 | ✓ 349ms | ✓ 967ms | ✓ 734ms | http |
| 38.145.208.99:8443 | ✓ 677ms | 否 | ✓ 339ms | ✓ 954ms | ✓ 1843ms | http |
| 20.120.225.109:3128 | ✓ 677ms | 否 | ✓ 568ms | ✓ 1412ms | 否 | http |
| 2.56.122.146:10808 | ✓ 493ms | 否 | ✓ 1116ms | ✓ 1891ms | ✓ 1624ms | http |
| 178.236.245.59:3128 | ✓ 675ms | ✓ 1726ms | ✓ 779ms | ✓ 1872ms | 否 | http |
| 160.250.4.245:1 | ✓ 1409ms | 否 | ✓ 1404ms | ✓ 1404ms | ✓ 1123ms | http |
| 45.93.29.147:6005 | ✓ 1708ms | 否 | 否 | ✓ 1332ms | ✓ 1057ms | http |

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
