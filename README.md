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

最后更新：2026-05-16 08:10:21 UTC（2026-05-16 16:10:21 UTC+8）

**代理总数：93**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 129.80.217.21:444 | ✓ 1637ms | ✓ 1196ms | ✓ 1365ms | ✓ 1241ms | ✓ 1416ms | http |
| 129.80.238.83:444 | 否 | 否 | ✓ 406ms | ✓ 1192ms | ✓ 1244ms | http |
| 1.231.81.166:3128 | ✓ 1763ms | 否 | ✓ 1299ms | ✓ 971ms | ✓ 857ms | http |
| 103.147.152.12:1080 | ✓ 1080ms | 否 | ✓ 1216ms | 否 | ✓ 1377ms | http |
| 114.214.170.41:27890 | ✓ 982ms | ✓ 1144ms | ✓ 1074ms | ✓ 1144ms | ✓ 927ms | http |
| 116.254.118.180:80 | ✓ 1642ms | ✓ 1993ms | ✓ 1324ms | ✓ 1267ms | ✓ 974ms | http |
| 212.58.132.5:8888 | ✓ 1110ms | 否 | ✓ 1119ms | ✓ 1559ms | ✓ 1201ms | http |
| 8.219.97.248:80 | ✓ 1002ms | 否 | ✓ 1533ms | 否 | ✓ 1610ms | http |
| 185.200.188.234:10001 | ✓ 1963ms | 否 | ✓ 1090ms | 否 | ✓ 1960ms | http |
| 113.160.132.26:8080 | ✓ 1672ms | 否 | ✓ 1549ms | ✓ 1626ms | 否 | http |
| 45.88.0.113:3128 | ✓ 1985ms | 否 | ✓ 1945ms | 否 | ✓ 1367ms | http |
| 181.78.17.131:999 | ✓ 1274ms | 否 | 否 | ✓ 1683ms | ✓ 1448ms | http |
| 45.88.0.115:3128 | ✓ 1995ms | 否 | ✓ 1948ms | 否 | ✓ 1386ms | http |
| 103.134.85.167:3128 | ✓ 1131ms | 否 | ✓ 832ms | ✓ 1200ms | ✓ 968ms | http |
| 84.47.150.125:1080 | ✓ 1134ms | 否 | ✓ 1511ms | 否 | ✓ 1809ms | http |
| 158.160.215.167:8126 | ✓ 1700ms | 否 | ✓ 865ms | 否 | ✓ 1639ms | http |
| 91.242.229.129:8092 | ✓ 834ms | 否 | ✓ 1688ms | 否 | ✓ 1908ms | http |
| 218.108.131.186:17890 | ✓ 841ms | ✓ 1041ms | ✓ 815ms | ✓ 1100ms | 否 | http |
| 59.46.216.131:30001 | 否 | ✓ 1358ms | ✓ 1335ms | 否 | ✓ 1422ms | http |
| 115.231.181.40:8128 | ✓ 926ms | ✓ 1082ms | ✓ 929ms | 否 | 否 | http |
| 103.21.220.141:3128 | ✓ 633ms | 否 | ✓ 616ms | ✓ 774ms | ✓ 634ms | http |
| 8.154.21.175:3128 | ✓ 807ms | ✓ 985ms | ✓ 833ms | ✓ 1036ms | ✓ 858ms | http |
| 128.199.113.85:9090 | ✓ 1330ms | 否 | ✓ 717ms | ✓ 1013ms | ✓ 811ms | http |
| 152.42.170.187:9090 | ✓ 1321ms | 否 | ✓ 725ms | ✓ 1111ms | ✓ 815ms | http |
| 128.199.114.189:9090 | ✓ 1320ms | 否 | ✓ 726ms | ✓ 1338ms | ✓ 821ms | http |
| 128.199.116.219:9090 | ✓ 1322ms | 否 | ✓ 906ms | ✓ 1060ms | ✓ 805ms | http |
| 77.110.107.80:8080 | ✓ 905ms | 否 | ✓ 1568ms | 否 | ✓ 1450ms | http |
| 45.59.122.132:80 | ✓ 1781ms | ✓ 1945ms | ✓ 1150ms | 否 | ✓ 1195ms | http |
| 148.230.4.241:999 | ✓ 630ms | ✓ 1604ms | ✓ 675ms | ✓ 1352ms | ✓ 1302ms | http |
| 42.114.172.179:2083 | ✓ 1519ms | 否 | ✓ 1516ms | ✓ 1639ms | ✓ 1461ms | http |
| 137.59.47.73:3128 | ✓ 1055ms | ✓ 1225ms | ✓ 1285ms | ✓ 1424ms | ✓ 851ms | http |
| 106.10.55.212:1121 | ✓ 1599ms | 否 | 否 | ✓ 1307ms | ✓ 1021ms | http |
| 5.252.33.13:2025 | ✓ 1816ms | 否 | ✓ 1544ms | 否 | ✓ 1930ms | http |
| 34.71.229.255:3128 | ✓ 305ms | ✓ 1287ms | ✓ 275ms | ✓ 1084ms | ✓ 796ms | http |
| 146.190.80.158:9090 | ✓ 738ms | 否 | ✓ 694ms | ✓ 1031ms | 否 | http |
| 128.199.254.13:9090 | ✓ 717ms | 否 | ✓ 714ms | ✓ 1007ms | ✓ 820ms | http |
| 128.199.121.61:9090 | ✓ 711ms | 否 | ✓ 914ms | ✓ 1037ms | ✓ 870ms | http |
| 57.129.144.178:40000 | ✓ 825ms | 否 | ✓ 1488ms | ✓ 1617ms | ✓ 1677ms | http |
| 160.238.65.6:3128 | ✓ 1234ms | ✓ 1846ms | ✓ 707ms | ✓ 1568ms | ✓ 1277ms | http |
| 160.238.65.4:3128 | ✓ 1216ms | 否 | 否 | ✓ 1628ms | ✓ 1372ms | http |
| 160.238.65.2:3128 | ✓ 1694ms | 否 | 否 | ✓ 1613ms | ✓ 1785ms | http |
| 160.238.65.8:3128 | ✓ 1222ms | 否 | 否 | ✓ 1627ms | ✓ 1370ms | http |
| 160.238.65.3:3128 | ✓ 1702ms | 否 | ✓ 720ms | ✓ 1609ms | ✓ 1727ms | http |
| 107.175.85.198:1080 | ✓ 1247ms | ✓ 1463ms | ✓ 1127ms | ✓ 1479ms | ✓ 1186ms | http |
| 185.191.236.162:3128 | ✓ 821ms | 否 | ✓ 838ms | ✓ 1873ms | ✓ 1252ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1756ms | ✓ 1155ms | ✓ 1461ms | http |
| 213.220.62.63:3128 | ✓ 1792ms | 否 | ✓ 719ms | 否 | ✓ 1297ms | http |
| 158.160.215.167:8124 | ✓ 1324ms | 否 | ✓ 1315ms | 否 | ✓ 1710ms | http |
| 42.114.172.179:2045 | ✓ 1463ms | 否 | ✓ 1484ms | ✓ 1515ms | ✓ 1370ms | http |
| 42.114.172.179:2088 | ✓ 1503ms | 否 | ✓ 1632ms | ✓ 1587ms | ✓ 1480ms | http |
| 103.144.146.2:8080 | 否 | 否 | ✓ 1409ms | ✓ 1439ms | ✓ 1436ms | http |
| 47.74.226.8:5001 | 否 | 否 | ✓ 1031ms | ✓ 1852ms | ✓ 1259ms | http |
| 210.223.44.230:3128 | ✓ 1144ms | ✓ 1614ms | 否 | ✓ 1818ms | ✓ 1404ms | http |
| 212.34.146.118:3128 | ✓ 1734ms | 否 | ✓ 1791ms | ✓ 1686ms | 否 | http |
| 45.167.126.21:999 | ✓ 1439ms | 否 | 否 | ✓ 1761ms | ✓ 1491ms | http |
| 120.92.212.16:8890 | ✓ 1932ms | 否 | ✓ 1910ms | 否 | ✓ 1651ms | http |
| 120.92.212.16:7890 | ✓ 862ms | ✓ 1384ms | ✓ 1219ms | 否 | ✓ 1624ms | http |
| 38.188.247.12:999 | 否 | 否 | ✓ 1879ms | ✓ 1422ms | ✓ 1479ms | http |
| 166.88.55.83:7890 | ✓ 620ms | ✓ 1225ms | ✓ 612ms | ✓ 768ms | ✓ 613ms | http |
| 45.13.116.188:21537 | ✓ 1289ms | ✓ 1729ms | ✓ 1508ms | ✓ 1522ms | ✓ 1172ms | http |
| 190.60.34.250:999 | ✓ 1306ms | ✓ 1873ms | ✓ 1691ms | ✓ 1998ms | ✓ 1821ms | http |
| 160.238.65.7:3128 | ✓ 1710ms | ✓ 1683ms | 否 | 否 | ✓ 1996ms | http |
| 152.70.91.193:40000 | ✓ 1665ms | 否 | 否 | ✓ 1746ms | ✓ 1636ms | http |
| 121.130.177.28:8888 | ✓ 937ms | ✓ 1018ms | 否 | ✓ 1569ms | 否 | http |
| 87.120.222.214:444 | ✓ 1118ms | 否 | ✓ 1063ms | 否 | ✓ 1761ms | http |
| 3.101.133.120:80 | ✓ 908ms | ✓ 1341ms | ✓ 1016ms | ✓ 835ms | ✓ 842ms | http |
| 200.174.198.32:8888 | ✓ 1421ms | 否 | ✓ 1038ms | 否 | ✓ 1713ms | http |
| 168.110.52.228:3128 | ✓ 1462ms | 否 | ✓ 506ms | ✓ 760ms | ✓ 1581ms | http |
| 112.163.160.93:3128 | 否 | 否 | ✓ 1949ms | ✓ 964ms | ✓ 742ms | http |
| 179.43.159.98:1080 | ✓ 1753ms | 否 | ✓ 1773ms | 否 | ✓ 1595ms | http |
| 49.144.29.62:8082 | ✓ 1298ms | 否 | ✓ 1075ms | ✓ 1160ms | ✓ 1182ms | http |
| 45.88.0.99:3128 | ✓ 1942ms | ✓ 1606ms | ✓ 1997ms | 否 | 否 | http |
| 103.147.152.12:1095 | ✓ 652ms | 否 | ✓ 1059ms | ✓ 1740ms | ✓ 1369ms | http |
| 61.144.152.160:9000 | ✓ 1554ms | ✓ 1255ms | ✓ 1201ms | ✓ 1813ms | ✓ 1203ms | http |
| 168.222.254.136:8888 | ✓ 1170ms | 否 | ✓ 1815ms | 否 | ✓ 1574ms | http |
| 86.104.72.219:1081 | ✓ 402ms | ✓ 1307ms | ✓ 522ms | 否 | 否 | http |
| 103.227.187.241:6090 | ✓ 1800ms | 否 | ✓ 1287ms | 否 | ✓ 1496ms | http |
| 121.230.8.136:1080 | 否 | 否 | ✓ 1215ms | ✓ 1190ms | ✓ 1083ms | http |
| 104.248.151.93:9090 | ✓ 891ms | 否 | ✓ 893ms | ✓ 1052ms | ✓ 825ms | http |
| 193.160.209.58:1080 | ✓ 1622ms | 否 | ✓ 999ms | 否 | ✓ 1826ms | http |
| 190.93.224.32:999 | ✓ 1160ms | 否 | ✓ 1296ms | 否 | ✓ 1844ms | http |
| 86.104.72.220:1082 | ✓ 1047ms | ✓ 1229ms | ✓ 470ms | ✓ 1577ms | ✓ 1111ms | http |
| 185.21.15.206:3128 | 否 | 否 | ✓ 841ms | ✓ 1744ms | ✓ 1734ms | http |
| 43.167.192.85:8080 | ✓ 1233ms | 否 | ✓ 1332ms | 否 | ✓ 1399ms | http |
| 158.160.215.167:8127 | ✓ 1308ms | 否 | ✓ 1843ms | 否 | ✓ 1451ms | http |
| 147.45.186.28:3128 | ✓ 1430ms | 否 | ✓ 989ms | ✓ 1846ms | ✓ 1513ms | http |
| 152.32.132.190:7890 | ✓ 1020ms | 否 | ✓ 1210ms | ✓ 785ms | ✓ 1694ms | http |
| 158.160.215.167:8125 | ✓ 1629ms | 否 | ✓ 1633ms | 否 | ✓ 1880ms | http |
| 190.12.150.244:999 | ✓ 1049ms | ✓ 1899ms | ✓ 1105ms | 否 | 否 | http |
| 86.104.72.219:1082 | ✓ 305ms | 否 | ✓ 891ms | ✓ 1274ms | ✓ 1729ms | http |
| 61.52.131.172:8443 | ✓ 888ms | ✓ 1162ms | ✓ 966ms | ✓ 1291ms | ✓ 929ms | http |
| 103.172.70.173:8080 | ✓ 1247ms | 否 | ✓ 1880ms | ✓ 1459ms | 否 | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1708ms | ✓ 1683ms | ✓ 1810ms | http |

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
