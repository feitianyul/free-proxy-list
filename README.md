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

最后更新：2026-04-22 12:02:11 UTC（2026-04-22 20:02:11 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.127.128.70:8080 | ✓ 782ms | 否 | 否 | ✓ 1754ms | ✓ 1215ms | http |
| 1.231.81.166:3128 | ✓ 1807ms | ✓ 1491ms | ✓ 1074ms | ✓ 1836ms | ✓ 769ms | http |
| 152.32.132.190:7890 | ✓ 1125ms | 否 | ✓ 860ms | 否 | ✓ 755ms | http |
| 152.42.208.139:8118 | ✓ 1303ms | 否 | ✓ 844ms | ✓ 1160ms | ✓ 928ms | http |
| 46.101.95.183:8888 | ✓ 871ms | 否 | 否 | ✓ 1815ms | ✓ 1414ms | http |
| 188.246.224.49:7890 | ✓ 806ms | ✓ 1810ms | ✓ 1704ms | 否 | ✓ 1613ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1686ms | ✓ 1413ms | ✓ 1500ms | ✓ 1105ms | http |
| 212.58.132.5:8888 | ✓ 1295ms | 否 | ✓ 1439ms | ✓ 1566ms | ✓ 1313ms | http |
| 180.191.231.47:8080 | ✓ 1909ms | 否 | ✓ 1900ms | 否 | ✓ 1721ms | http |
| 35.225.22.61:80 | ✓ 367ms | 否 | ✓ 1269ms | ✓ 1126ms | ✓ 854ms | http |
| 8.219.97.248:80 | ✓ 1199ms | 否 | ✓ 1253ms | ✓ 1679ms | 否 | http |
| 45.140.147.82:1081 | ✓ 552ms | ✓ 1529ms | ✓ 874ms | ✓ 1933ms | ✓ 1307ms | http |
| 45.76.207.177:40000 | ✓ 804ms | 否 | ✓ 1037ms | ✓ 1087ms | ✓ 1115ms | http |
| 45.140.147.82:1082 | ✓ 564ms | ✓ 1640ms | ✓ 792ms | ✓ 1881ms | ✓ 1345ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1230ms | ✓ 974ms | ✓ 1118ms | ✓ 1231ms | http |
| 213.32.85.26:3128 | ✓ 838ms | 否 | ✓ 1895ms | 否 | ✓ 1773ms | http |
| 84.47.150.126:1080 | ✓ 1147ms | 否 | ✓ 1917ms | 否 | ✓ 1943ms | http |
| 168.222.254.136:8888 | ✓ 775ms | 否 | ✓ 1806ms | 否 | ✓ 1967ms | http |
| 115.231.181.40:8128 | ✓ 1389ms | ✓ 1255ms | ✓ 1593ms | ✓ 1267ms | ✓ 1407ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1228ms | ✓ 967ms | 否 | ✓ 1392ms | http |
| 177.93.132.244:3128 | ✓ 748ms | 否 | ✓ 1856ms | 否 | ✓ 1774ms | http |
| 89.208.106.138:10808 | 否 | 否 | ✓ 1771ms | ✓ 1749ms | ✓ 1300ms | http |
| 218.108.131.186:17890 | ✓ 905ms | ✓ 1489ms | ✓ 1111ms | ✓ 1426ms | ✓ 1005ms | http |
| 91.99.15.45:2095 | ✓ 611ms | ✓ 1579ms | ✓ 1452ms | ✓ 1933ms | 否 | http |
| 85.190.99.143:443 | ✓ 1474ms | 否 | ✓ 1813ms | 否 | ✓ 1530ms | http |
| 208.87.243.199:7878 | ✓ 867ms | ✓ 1523ms | ✓ 629ms | 否 | 否 | http |
| 34.71.229.255:3128 | ✓ 587ms | 否 | ✓ 889ms | ✓ 988ms | ✓ 684ms | http |
| 62.113.119.14:8080 | ✓ 887ms | ✓ 1722ms | ✓ 800ms | ✓ 1538ms | ✓ 1166ms | http |
| 138.124.99.216:8888 | ✓ 1716ms | 否 | 否 | ✓ 1856ms | ✓ 1333ms | http |
| 144.31.27.49:1080 | 否 | ✓ 1995ms | ✓ 1622ms | 否 | ✓ 1492ms | http |
| 27.254.99.183:8118 | ✓ 1459ms | 否 | ✓ 1149ms | ✓ 1368ms | 否 | http |
| 78.11.96.22:8888 | ✓ 965ms | 否 | ✓ 1749ms | ✓ 1545ms | ✓ 1347ms | http |
| 5.253.43.103:3128 | 否 | ✓ 1928ms | ✓ 983ms | ✓ 1900ms | ✓ 1322ms | http |
| 38.180.192.119:3128 | 否 | ✓ 822ms | ✓ 732ms | ✓ 840ms | ✓ 728ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1620ms | ✓ 1454ms | ✓ 1218ms | http |
| 168.144.75.9:3128 | ✓ 1548ms | 否 | ✓ 1401ms | 否 | ✓ 1718ms | http |
| 134.209.27.75:3128 | ✓ 1680ms | 否 | ✓ 1247ms | ✓ 1885ms | ✓ 1745ms | http |
| 147.45.60.34:1082 | ✓ 431ms | ✓ 1416ms | ✓ 1040ms | 否 | 否 | http |
| 210.45.76.58:42992 | ✓ 1283ms | ✓ 1469ms | ✓ 1362ms | 否 | 否 | http |
| 84.47.150.125:1080 | ✓ 1123ms | 否 | ✓ 1728ms | 否 | ✓ 1834ms | http |
| 8.137.112.117:3128 | ✓ 1543ms | ✓ 1448ms | ✓ 1104ms | 否 | ✓ 1202ms | http |
| 77.110.113.24:40000 | ✓ 1532ms | ✓ 1833ms | ✓ 1447ms | 否 | 否 | http |
| 222.102.86.137:3008 | ✓ 1560ms | 否 | ✓ 1136ms | ✓ 1742ms | 否 | http |
| 89.111.174.221:8080 | ✓ 1252ms | 否 | ✓ 1509ms | ✓ 1942ms | ✓ 1778ms | http |
| 91.107.124.215:3128 | ✓ 1215ms | ✓ 1714ms | ✓ 781ms | 否 | ✓ 1668ms | http |
| 101.32.244.83:8080 | ✓ 1031ms | 否 | ✓ 1027ms | ✓ 1408ms | ✓ 1332ms | http |
| 47.84.73.61:1080 | ✓ 1959ms | 否 | ✓ 931ms | ✓ 1141ms | ✓ 933ms | http |
| 8.209.238.110:47701 | 否 | 否 | ✓ 897ms | ✓ 944ms | ✓ 756ms | http |
| 223.84.151.86:30005 | 否 | ✓ 1395ms | ✓ 1239ms | 否 | ✓ 1328ms | http |
| 8.222.175.80:6128 | 否 | 否 | ✓ 1251ms | ✓ 1193ms | ✓ 1428ms | http |
| 185.191.236.162:3128 | ✓ 1751ms | 否 | ✓ 1844ms | 否 | ✓ 1905ms | http |
| 45.140.147.155:1082 | ✓ 1246ms | ✓ 1902ms | ✓ 616ms | 否 | 否 | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1355ms | ✓ 1013ms | ✓ 882ms | http |
| 31.131.248.48:3129 | ✓ 1203ms | 否 | ✓ 1353ms | 否 | ✓ 1705ms | http |
| 64.181.240.152:3128 | ✓ 1027ms | ✓ 1560ms | ✓ 1846ms | ✓ 1354ms | ✓ 859ms | http |
| 59.46.216.131:30001 | ✓ 1056ms | 否 | ✓ 1229ms | ✓ 1461ms | ✓ 1824ms | http |
| 34.101.184.164:3128 | ✓ 1083ms | 否 | ✓ 1439ms | ✓ 1698ms | ✓ 1187ms | http |
| 217.76.245.80:999 | ✓ 1120ms | ✓ 1724ms | ✓ 1223ms | ✓ 1763ms | 否 | http |
| 195.26.224.49:3128 | 否 | 否 | ✓ 925ms | ✓ 1608ms | ✓ 1617ms | http |
| 161.97.184.191:8080 | ✓ 962ms | 否 | ✓ 1606ms | 否 | ✓ 1868ms | http |
| 217.77.102.18:3128 | ✓ 1035ms | 否 | ✓ 1252ms | ✓ 1942ms | ✓ 1778ms | http |
| 8.140.104.98:3128 | ✓ 973ms | ✓ 1269ms | ✓ 1051ms | ✓ 1326ms | ✓ 1041ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1831ms | ✓ 1646ms | 否 | ✓ 1424ms | http |
| 223.205.186.237:8080 | ✓ 1612ms | 否 | ✓ 1382ms | ✓ 1610ms | ✓ 1594ms | http |
| 103.137.35.2:80 | 否 | 否 | ✓ 1510ms | ✓ 1641ms | ✓ 1565ms | http |
| 200.174.198.32:8888 | ✓ 1538ms | 否 | ✓ 1573ms | 否 | ✓ 1932ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1014ms | ✓ 1804ms | ✓ 1442ms | http |
| 45.140.147.155:1081 | ✓ 562ms | ✓ 1408ms | ✓ 1412ms | ✓ 1802ms | ✓ 1221ms | http |
| 65.21.201.149:8080 | ✓ 1986ms | ✓ 1857ms | ✓ 1385ms | 否 | 否 | http |
| 103.156.16.12:8818 | ✓ 1835ms | 否 | ✓ 1436ms | ✓ 1498ms | ✓ 1514ms | http |
| 103.170.196.74:8080 | ✓ 1846ms | 否 | ✓ 1767ms | ✓ 1792ms | ✓ 1767ms | http |
| 121.43.196.213:8222 | ✓ 992ms | ✓ 1176ms | ✓ 978ms | ✓ 1308ms | ✓ 1016ms | http |
| 121.43.196.210:8222 | ✓ 1005ms | ✓ 1139ms | ✓ 960ms | ✓ 1475ms | ✓ 1015ms | http |
| 114.55.226.123:10086 | ✓ 1178ms | ✓ 1785ms | ✓ 1096ms | ✓ 1426ms | ✓ 1185ms | http |
| 61.52.131.172:8443 | ✓ 1824ms | ✓ 1259ms | 否 | ✓ 1204ms | ✓ 1035ms | http |
| 117.236.124.166:3128 | ✓ 1212ms | 否 | ✓ 1251ms | 否 | ✓ 1341ms | http |

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
