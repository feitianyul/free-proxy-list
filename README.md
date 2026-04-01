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

最后更新：2026-04-01 20:06:23 UTC（2026-04-02 04:06:23 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 781ms | ✓ 1038ms | ✓ 795ms | ✓ 937ms | ✓ 749ms | http |
| 1.231.81.166:3128 | ✓ 904ms | ✓ 1077ms | ✓ 1557ms | ✓ 1108ms | ✓ 880ms | http |
| 147.161.239.240:8800 | ✓ 642ms | ✓ 1637ms | ✓ 1640ms | ✓ 1767ms | ✓ 1459ms | http |
| 147.161.210.140:8800 | ✓ 936ms | 否 | ✓ 1119ms | ✓ 1344ms | ✓ 1872ms | http |
| 167.103.115.102:8800 | ✓ 1025ms | 否 | ✓ 1506ms | ✓ 1559ms | ✓ 1300ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1372ms | ✓ 1159ms | ✓ 1595ms | ✓ 1478ms | http |
| 147.45.186.28:3128 | ✓ 1051ms | 否 | ✓ 1833ms | 否 | ✓ 1758ms | http |
| 45.167.124.52:8080 | ✓ 1835ms | 否 | ✓ 1440ms | 否 | ✓ 1420ms | http |
| 38.145.220.102:8453 | ✓ 1523ms | 否 | ✓ 396ms | 否 | ✓ 1377ms | http |
| 163.5.180.103:56297 | ✓ 773ms | ✓ 1563ms | ✓ 1268ms | ✓ 1853ms | ✓ 1284ms | http |
| 95.213.217.168:52004 | ✓ 668ms | ✓ 1635ms | ✓ 1115ms | ✓ 1868ms | ✓ 1311ms | http |
| 167.103.34.108:8800 | ✓ 1230ms | 否 | ✓ 1250ms | ✓ 1436ms | ✓ 1362ms | http |
| 34.101.184.164:3128 | ✓ 1177ms | 否 | ✓ 1344ms | ✓ 1878ms | ✓ 1165ms | http |
| 167.103.144.127:8800 | ✓ 1508ms | 否 | ✓ 1478ms | ✓ 1663ms | ✓ 1841ms | http |
| 167.103.31.122:8800 | ✓ 1950ms | 否 | ✓ 1774ms | 否 | ✓ 1998ms | http |
| 180.250.219.58:53281 | ✓ 1951ms | 否 | ✓ 1577ms | 否 | ✓ 1950ms | http |
| 65.57.157.150:8080 | ✓ 456ms | ✓ 1068ms | ✓ 768ms | ✓ 1383ms | ✓ 1223ms | http |
| 133.242.138.34:8100 | ✓ 708ms | ✓ 1257ms | ✓ 1218ms | 否 | ✓ 1573ms | http |
| 101.43.127.100:8877 | ✓ 1923ms | ✓ 1630ms | ✓ 944ms | ✓ 1720ms | ✓ 998ms | http |
| 203.80.138.81:50000 | 否 | ✓ 1815ms | ✓ 1420ms | ✓ 1697ms | 否 | http |
| 177.234.217.88:999 | ✓ 1817ms | 否 | ✓ 1806ms | 否 | ✓ 1641ms | http |
| 38.34.183.130:8452 | ✓ 1485ms | ✓ 930ms | ✓ 1038ms | ✓ 1646ms | ✓ 749ms | http |
| 5.102.109.41:999 | ✓ 1429ms | ✓ 1095ms | ✓ 976ms | ✓ 1320ms | ✓ 1046ms | http |
| 183.232.248.102:7897 | ✓ 948ms | ✓ 1215ms | ✓ 1035ms | ✓ 1266ms | ✓ 1054ms | http |
| 185.40.7.206:3128 | ✓ 1168ms | ✓ 1705ms | ✓ 1723ms | 否 | ✓ 1647ms | http |
| 160.250.5.22:1 | ✓ 1512ms | 否 | ✓ 1644ms | ✓ 1324ms | ✓ 1150ms | http |
| 45.12.151.226:2829 | ✓ 1810ms | ✓ 1966ms | ✓ 1267ms | 否 | 否 | http |
| 116.80.49.166:3172 | ✓ 1606ms | 否 | ✓ 1573ms | ✓ 1948ms | ✓ 1724ms | http |
| 103.39.51.207:8080 | ✓ 1534ms | 否 | 否 | ✓ 1741ms | ✓ 1881ms | http |
| 45.140.147.155:1081 | ✓ 1154ms | ✓ 1641ms | ✓ 838ms | ✓ 1551ms | ✓ 1216ms | http |
| 38.34.179.150:8449 | ✓ 674ms | ✓ 1130ms | ✓ 1016ms | ✓ 1754ms | ✓ 744ms | http |
| 167.160.191.204:6005 | ✓ 953ms | ✓ 1953ms | 否 | ✓ 1457ms | 否 | http |
| 38.188.247.12:999 | ✓ 1868ms | 否 | ✓ 1173ms | ✓ 1993ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1729ms | 否 | ✓ 1836ms | 否 | ✓ 1997ms | http |
| 35.225.22.61:80 | ✓ 1899ms | ✓ 1181ms | ✓ 280ms | ✓ 966ms | ✓ 917ms | http |
| 198.59.68.130:3128 | ✓ 1719ms | ✓ 1559ms | ✓ 1462ms | ✓ 1990ms | ✓ 935ms | http |
| 59.46.216.131:30001 | ✓ 1028ms | ✓ 1438ms | 否 | ✓ 1459ms | ✓ 1124ms | http |
| 120.92.212.16:8890 | ✓ 1141ms | ✓ 1292ms | 否 | 否 | ✓ 1765ms | http |
| 103.113.70.189:1081 | ✓ 387ms | ✓ 1265ms | ✓ 430ms | ✓ 1294ms | ✓ 1818ms | http |
| 106.10.55.212:1121 | ✓ 1207ms | ✓ 1606ms | ✓ 1448ms | ✓ 1863ms | ✓ 1285ms | http |
| 212.58.132.5:8888 | ✓ 1661ms | 否 | ✓ 1771ms | ✓ 1558ms | ✓ 1720ms | http |
| 42.96.16.158:1311 | ✓ 1854ms | 否 | 否 | ✓ 1414ms | ✓ 1315ms | http |
| 121.230.8.111:1080 | 否 | ✓ 1735ms | ✓ 1201ms | ✓ 1756ms | ✓ 1507ms | http |
| 38.145.220.33:8448 | ✓ 487ms | ✓ 974ms | ✓ 916ms | ✓ 917ms | ✓ 852ms | http |
| 45.136.130.188:8449 | ✓ 486ms | ✓ 796ms | ✓ 1095ms | ✓ 1451ms | ✓ 864ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1385ms | ✓ 1654ms | ✓ 1141ms | http |
| 92.119.127.212:6005 | ✓ 1137ms | 否 | ✓ 1551ms | ✓ 1788ms | ✓ 1807ms | http |
| 38.145.203.135:8444 | ✓ 573ms | ✓ 898ms | ✓ 521ms | ✓ 819ms | ✓ 670ms | http |
| 38.145.208.220:8448 | ✓ 232ms | ✓ 812ms | ✓ 898ms | ✓ 954ms | ✓ 642ms | http |
| 38.34.179.21:8446 | ✓ 717ms | ✓ 897ms | ✓ 425ms | ✓ 1164ms | ✓ 934ms | http |
| 45.136.130.168:8448 | ✓ 1784ms | ✓ 1060ms | ✓ 225ms | ✓ 883ms | ✓ 1358ms | http |
| 45.136.130.173:8448 | ✓ 1817ms | ✓ 1029ms | ✓ 243ms | ✓ 921ms | ✓ 1335ms | http |
| 38.34.179.16:8451 | ✓ 1205ms | ✓ 1603ms | ✓ 499ms | ✓ 1049ms | ✓ 1461ms | http |
| 38.145.220.8:8451 | ✓ 739ms | ✓ 1673ms | ✓ 588ms | ✓ 845ms | ✓ 1383ms | http |
| 45.136.131.38:8445 | ✓ 213ms | ✓ 949ms | ✓ 1978ms | ✓ 1551ms | ✓ 807ms | http |
| 121.230.8.211:1080 | ✓ 1263ms | ✓ 1475ms | ✓ 1046ms | ✓ 1550ms | ✓ 1128ms | http |
| 38.34.179.202:8449 | ✓ 636ms | ✓ 1787ms | ✓ 1520ms | ✓ 1035ms | ✓ 1797ms | http |
| 38.34.179.100:8449 | ✓ 604ms | ✓ 883ms | ✓ 336ms | ✓ 1136ms | ✓ 811ms | http |
| 38.145.203.46:8448 | ✓ 472ms | ✓ 944ms | ✓ 240ms | ✓ 840ms | 否 | http |
| 38.145.220.41:8444 | ✓ 206ms | ✓ 823ms | ✓ 1068ms | ✓ 1686ms | 否 | http |
| 38.145.208.237:8445 | ✓ 891ms | 否 | ✓ 253ms | ✓ 1065ms | 否 | http |
| 45.136.131.34:8451 | ✓ 1078ms | ✓ 1061ms | ✓ 1048ms | 否 | 否 | http |
| 38.145.203.97:8453 | ✓ 669ms | 否 | ✓ 1274ms | ✓ 1907ms | 否 | http |
| 213.131.85.30:1976 | ✓ 1987ms | 否 | ✓ 1928ms | ✓ 1995ms | 否 | http |
| 45.149.92.147:5001 | ✓ 731ms | 否 | ✓ 769ms | ✓ 933ms | ✓ 731ms | http |
| 120.92.212.16:7890 | ✓ 1356ms | ✓ 1314ms | 否 | ✓ 1300ms | 否 | http |
| 208.87.243.199:7878 | ✓ 1449ms | ✓ 1016ms | ✓ 908ms | 否 | 否 | http |
| 38.210.179.112:999 | ✓ 684ms | ✓ 1287ms | ✓ 1204ms | 否 | 否 | http |
| 159.75.113.2:3128 | 否 | ✓ 1693ms | ✓ 1403ms | ✓ 1129ms | 否 | http |
| 38.145.208.179:8447 | ✓ 1863ms | ✓ 804ms | ✓ 622ms | ✓ 1599ms | ✓ 1495ms | http |
| 107.174.80.186:3128 | ✓ 678ms | ✓ 916ms | ✓ 1823ms | 否 | ✓ 1842ms | http |
| 38.34.179.14:8450 | 否 | ✓ 1305ms | ✓ 573ms | ✓ 1553ms | ✓ 1258ms | http |
| 113.255.59.226:8080 | 否 | 否 | ✓ 1491ms | ✓ 1213ms | ✓ 1225ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1720ms | ✓ 1748ms | ✓ 1748ms | http |
| 103.39.51.190:8080 | ✓ 1634ms | 否 | 否 | ✓ 1706ms | ✓ 1745ms | http |
| 61.52.131.172:8443 | ✓ 915ms | ✓ 1239ms | ✓ 1019ms | ✓ 1256ms | ✓ 1011ms | http |
| 45.129.141.143:3128 | ✓ 1370ms | ✓ 1751ms | 否 | 否 | ✓ 1792ms | http |
| 202.129.206.239:3128 | ✓ 1410ms | 否 | ✓ 1492ms | 否 | ✓ 1549ms | http |
| 85.208.108.43:10808 | ✓ 377ms | 否 | ✓ 1058ms | ✓ 1199ms | ✓ 856ms | http |
| 106.117.208.101:7890 | ✓ 1150ms | ✓ 1681ms | ✓ 1382ms | ✓ 1378ms | 否 | http |

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
