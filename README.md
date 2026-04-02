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

最后更新：2026-04-02 08:15:10 UTC（2026-04-02 16:15:10 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 715ms | ✓ 1939ms | ✓ 1108ms | 否 | ✓ 1097ms | http |
| 203.80.138.81:50000 | ✓ 846ms | ✓ 1154ms | ✓ 900ms | ✓ 1076ms | ✓ 859ms | http |
| 1.231.81.166:3128 | ✓ 825ms | ✓ 1021ms | ✓ 1448ms | ✓ 1178ms | ✓ 856ms | http |
| 147.161.210.140:8800 | ✓ 692ms | 否 | ✓ 846ms | ✓ 1567ms | ✓ 904ms | http |
| 159.223.71.162:443 | ✓ 1233ms | 否 | ✓ 858ms | ✓ 1093ms | ✓ 883ms | http |
| 167.103.115.102:8800 | ✓ 1233ms | 否 | ✓ 1012ms | ✓ 1100ms | ✓ 1028ms | http |
| 34.96.238.40:8080 | ✓ 1124ms | 否 | ✓ 1269ms | ✓ 1047ms | ✓ 1183ms | http |
| 147.161.239.240:8800 | ✓ 1463ms | ✓ 1743ms | ✓ 1099ms | ✓ 1740ms | ✓ 1511ms | http |
| 113.160.132.26:8080 | ✓ 1531ms | ✓ 1686ms | ✓ 960ms | ✓ 1595ms | ✓ 1178ms | http |
| 5.104.87.17:8051 | ✓ 1615ms | 否 | ✓ 1710ms | ✓ 1677ms | ✓ 1260ms | http |
| 95.213.217.168:52004 | ✓ 1499ms | ✓ 1988ms | ✓ 1508ms | 否 | ✓ 1790ms | http |
| 167.103.34.108:8800 | ✓ 1759ms | 否 | ✓ 1625ms | ✓ 1525ms | ✓ 1457ms | http |
| 180.250.219.58:53281 | ✓ 1890ms | 否 | ✓ 1538ms | 否 | ✓ 1930ms | http |
| 45.140.147.155:1082 | ✓ 609ms | 否 | ✓ 1027ms | 否 | ✓ 1298ms | http |
| 45.149.92.147:5001 | 否 | 否 | ✓ 727ms | ✓ 1062ms | ✓ 709ms | http |
| 34.101.184.164:3128 | ✓ 1319ms | 否 | ✓ 1245ms | ✓ 1488ms | ✓ 1460ms | http |
| 167.103.144.127:8800 | ✓ 1532ms | 否 | ✓ 1342ms | 否 | ✓ 1706ms | http |
| 45.167.124.52:8080 | ✓ 635ms | 否 | ✓ 836ms | ✓ 1647ms | ✓ 1388ms | http |
| 45.12.151.226:2829 | ✓ 1221ms | 否 | ✓ 745ms | ✓ 1712ms | ✓ 1391ms | http |
| 42.96.16.158:1311 | ✓ 1451ms | 否 | ✓ 1539ms | ✓ 1252ms | ✓ 1109ms | http |
| 167.103.31.122:8800 | ✓ 1658ms | 否 | ✓ 1457ms | ✓ 1751ms | ✓ 1649ms | http |
| 43.99.54.236:5555 | ✓ 705ms | ✓ 1517ms | ✓ 846ms | ✓ 863ms | ✓ 674ms | http |
| 208.87.243.199:7878 | ✓ 432ms | ✓ 1393ms | ✓ 819ms | ✓ 1039ms | 否 | http |
| 202.141.161.53:10808 | ✓ 1054ms | ✓ 1397ms | ✓ 1119ms | 否 | ✓ 1130ms | http |
| 92.124.132.106:3128 | ✓ 1312ms | 否 | ✓ 1884ms | 否 | ✓ 1456ms | http |
| 166.0.192.117:8888 | ✓ 534ms | 否 | ✓ 1528ms | ✓ 1552ms | ✓ 985ms | http |
| 146.190.80.158:9090 | ✓ 1195ms | 否 | ✓ 803ms | ✓ 1075ms | ✓ 847ms | http |
| 128.199.254.13:9090 | ✓ 1192ms | 否 | ✓ 795ms | ✓ 1112ms | ✓ 864ms | http |
| 128.199.121.61:9090 | ✓ 1196ms | 否 | ✓ 848ms | ✓ 1127ms | ✓ 851ms | http |
| 120.92.212.16:8890 | ✓ 1346ms | ✓ 1213ms | ✓ 955ms | ✓ 1564ms | 否 | http |
| 160.250.5.22:1 | ✓ 1431ms | 否 | ✓ 1729ms | ✓ 1462ms | ✓ 1306ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1026ms | ✓ 1231ms | ✓ 1822ms | http |
| 45.167.125.21:999 | ✓ 1686ms | 否 | ✓ 1355ms | 否 | ✓ 1639ms | http |
| 159.223.71.162:8080 | ✓ 753ms | 否 | ✓ 786ms | ✓ 1085ms | ✓ 832ms | http |
| 128.199.113.85:9090 | ✓ 890ms | 否 | ✓ 1020ms | ✓ 1163ms | ✓ 880ms | http |
| 128.199.116.219:9090 | ✓ 737ms | 否 | ✓ 1037ms | ✓ 1076ms | ✓ 891ms | http |
| 106.10.55.212:1121 | 否 | ✓ 1370ms | ✓ 1165ms | ✓ 1654ms | ✓ 1026ms | http |
| 177.234.217.88:999 | ✓ 1671ms | 否 | ✓ 1803ms | 否 | ✓ 1710ms | http |
| 167.160.184.231:6005 | ✓ 845ms | ✓ 1269ms | ✓ 795ms | ✓ 1119ms | 否 | http |
| 160.250.4.245:1 | ✓ 1723ms | 否 | 否 | ✓ 1328ms | ✓ 1122ms | http |
| 209.126.84.232:8888 | 否 | 否 | ✓ 1118ms | ✓ 1943ms | ✓ 1484ms | http |
| 212.58.132.5:8888 | ✓ 1972ms | 否 | 否 | ✓ 1668ms | ✓ 1765ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1407ms | ✓ 1563ms | ✓ 1389ms | ✓ 1291ms | http |
| 38.145.220.55:8451 | 否 | ✓ 949ms | ✓ 704ms | 否 | ✓ 950ms | http |
| 167.160.191.204:6005 | ✓ 657ms | 否 | ✓ 1408ms | ✓ 1115ms | ✓ 1031ms | http |
| 209.97.149.157:80 | ✓ 1175ms | ✓ 1640ms | 否 | 否 | ✓ 1854ms | http |
| 174.140.109.250:3128 | ✓ 758ms | ✓ 1713ms | ✓ 325ms | ✓ 1121ms | ✓ 1067ms | http |
| 112.124.34.143:7500 | ✓ 865ms | ✓ 1993ms | ✓ 1678ms | ✓ 1397ms | ✓ 1054ms | http |
| 165.232.146.249:3128 | ✓ 765ms | ✓ 672ms | 否 | ✓ 1005ms | 否 | http |
| 164.90.151.28:3128 | ✓ 1021ms | ✓ 1995ms | ✓ 898ms | ✓ 734ms | ✓ 621ms | http |
| 45.174.241.172:999 | 否 | 否 | ✓ 584ms | ✓ 1534ms | ✓ 1386ms | http |
| 116.80.65.79:3172 | ✓ 1666ms | 否 | ✓ 1981ms | 否 | ✓ 1994ms | http |
| 128.199.114.189:9090 | ✓ 1752ms | 否 | ✓ 760ms | ✓ 1088ms | ✓ 865ms | http |
| 101.43.127.100:8877 | ✓ 1842ms | ✓ 1166ms | ✓ 946ms | ✓ 1104ms | ✓ 900ms | http |
| 38.145.208.241:8453 | ✓ 1027ms | 否 | ✓ 857ms | ✓ 836ms | ✓ 1973ms | http |
| 38.145.208.213:8449 | ✓ 536ms | ✓ 719ms | ✓ 1062ms | 否 | ✓ 617ms | http |
| 45.136.130.166:8449 | ✓ 470ms | ✓ 1716ms | ✓ 1689ms | ✓ 942ms | 否 | http |
| 45.136.130.186:8451 | ✓ 585ms | ✓ 889ms | ✓ 1270ms | 否 | ✓ 578ms | http |
| 38.145.220.41:8444 | ✓ 696ms | ✓ 709ms | ✓ 538ms | 否 | ✓ 1227ms | http |
| 38.34.179.68:8452 | ✓ 1197ms | ✓ 1953ms | ✓ 1126ms | ✓ 1917ms | ✓ 789ms | http |
| 38.34.179.106:8450 | ✓ 1511ms | ✓ 1078ms | ✓ 1282ms | 否 | ✓ 595ms | http |
| 38.34.179.21:8446 | ✓ 1627ms | ✓ 1189ms | ✓ 1118ms | 否 | ✓ 571ms | http |
| 38.34.179.100:8449 | 否 | ✓ 1017ms | ✓ 1456ms | 否 | ✓ 756ms | http |
| 38.145.203.45:8452 | ✓ 1553ms | 否 | ✓ 322ms | 否 | ✓ 1742ms | http |
| 121.230.9.225:1080 | 否 | ✓ 1771ms | ✓ 1125ms | ✓ 1998ms | ✓ 1826ms | http |
| 103.113.70.189:1081 | ✓ 1140ms | 否 | ✓ 784ms | ✓ 1118ms | ✓ 853ms | http |
| 38.34.179.51:8449 | 否 | ✓ 761ms | ✓ 1306ms | 否 | ✓ 582ms | http |
| 158.160.215.167:8124 | ✓ 1311ms | ✓ 1916ms | 否 | 否 | ✓ 1710ms | http |
| 195.19.217.200:3128 | ✓ 1287ms | 否 | ✓ 1954ms | 否 | ✓ 1903ms | http |
| 38.145.218.232:8448 | ✓ 1367ms | 否 | ✓ 959ms | ✓ 914ms | 否 | http |
| 181.78.44.63:999 | ✓ 944ms | 否 | ✓ 1280ms | ✓ 1453ms | ✓ 1380ms | http |
| 104.248.151.93:9090 | ✓ 852ms | 否 | ✓ 819ms | ✓ 1254ms | ✓ 922ms | http |
| 198.59.68.130:3128 | ✓ 787ms | ✓ 1218ms | ✓ 1222ms | ✓ 1264ms | ✓ 1040ms | http |
| 38.34.179.105:8449 | ✓ 1772ms | ✓ 1019ms | ✓ 384ms | ✓ 1177ms | 否 | http |
| 38.34.183.224:8448 | 否 | ✓ 1280ms | ✓ 327ms | ✓ 1264ms | 否 | http |
| 109.234.38.35:3128 | 否 | ✓ 1825ms | ✓ 1835ms | 否 | ✓ 1629ms | http |
| 59.46.216.131:30001 | ✓ 1005ms | ✓ 1411ms | ✓ 1679ms | 否 | 否 | http |
| 116.80.65.77:3172 | ✓ 1510ms | 否 | 否 | ✓ 1824ms | ✓ 1701ms | http |
| 103.69.84.106:3131 | ✓ 1821ms | 否 | ✓ 1577ms | ✓ 1201ms | ✓ 985ms | http |
| 103.67.46.225:3125 | ✓ 1809ms | 否 | ✓ 1724ms | ✓ 1667ms | ✓ 1692ms | http |
| 61.52.131.172:8443 | ✓ 948ms | ✓ 1144ms | ✓ 969ms | ✓ 1157ms | ✓ 954ms | http |
| 5.102.109.41:999 | 否 | ✓ 1298ms | ✓ 296ms | ✓ 1228ms | 否 | http |
| 107.174.80.186:3128 | ✓ 703ms | 否 | ✓ 341ms | ✓ 1053ms | 否 | http |
| 116.80.63.64:7777 | 否 | 否 | ✓ 1507ms | ✓ 1835ms | ✓ 1712ms | http |
| 106.117.208.101:7890 | ✓ 1240ms | 否 | ✓ 1376ms | ✓ 1457ms | ✓ 1056ms | http |
| 45.136.198.40:3128 | ✓ 821ms | 否 | ✓ 1675ms | 否 | ✓ 1699ms | http |

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
