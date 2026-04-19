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

最后更新：2026-04-19 12:30:33 UTC（2026-04-19 20:30:33 UTC+8）

**代理总数：95**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 95 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 95 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:3128 | ✓ 724ms | ✓ 1363ms | 否 | ✓ 1601ms | 否 | http |
| 149.51.42.10:8080 | ✓ 720ms | ✓ 1372ms | 否 | ✓ 1369ms | 否 | http |
| 218.108.131.186:17890 | ✓ 849ms | ✓ 1045ms | ✓ 855ms | ✓ 1123ms | ✓ 916ms | http |
| 1.231.81.166:3128 | ✓ 1266ms | ✓ 1825ms | ✓ 710ms | ✓ 988ms | ✓ 737ms | http |
| 46.101.95.183:8888 | ✓ 889ms | 否 | ✓ 1088ms | ✓ 1851ms | ✓ 1352ms | http |
| 152.42.208.139:8118 | ✓ 1461ms | 否 | ✓ 1382ms | ✓ 1090ms | ✓ 852ms | http |
| 113.160.132.26:8080 | ✓ 1562ms | ✓ 1451ms | ✓ 1393ms | ✓ 1213ms | ✓ 980ms | http |
| 166.88.61.54:8000 | ✓ 1044ms | ✓ 1573ms | 否 | 否 | ✓ 711ms | http |
| 14.247.76.52:8080 | ✓ 1567ms | ✓ 1751ms | 否 | ✓ 1222ms | ✓ 1066ms | http |
| 162.19.253.202:8443 | ✓ 1013ms | 否 | ✓ 1486ms | 否 | ✓ 1729ms | http |
| 188.246.224.49:7890 | ✓ 668ms | 否 | 否 | ✓ 1712ms | ✓ 1707ms | http |
| 81.30.156.115:8080 | ✓ 789ms | ✓ 1927ms | ✓ 1541ms | ✓ 1640ms | ✓ 1572ms | http |
| 159.223.225.118:8888 | ✓ 817ms | ✓ 1944ms | ✓ 1890ms | ✓ 1691ms | 否 | http |
| 43.132.188.134:443 | 否 | 否 | ✓ 1997ms | ✓ 1076ms | ✓ 1718ms | http |
| 35.225.22.61:80 | ✓ 587ms | ✓ 1170ms | 否 | ✓ 1406ms | ✓ 977ms | http |
| 38.180.192.119:3128 | ✓ 608ms | 否 | ✓ 820ms | ✓ 933ms | ✓ 598ms | http |
| 103.113.70.189:1081 | ✓ 566ms | 否 | ✓ 872ms | ✓ 1468ms | ✓ 1031ms | http |
| 192.3.248.190:8014 | ✓ 701ms | 否 | ✓ 918ms | ✓ 1846ms | ✓ 895ms | http |
| 168.110.52.228:3128 | ✓ 610ms | 否 | ✓ 1387ms | ✓ 1358ms | ✓ 840ms | http |
| 177.93.132.244:3128 | ✓ 902ms | 否 | ✓ 772ms | 否 | ✓ 1740ms | http |
| 185.138.116.150:8080 | ✓ 687ms | 否 | ✓ 776ms | 否 | ✓ 1252ms | http |
| 91.99.15.45:2095 | ✓ 1944ms | ✓ 1497ms | ✓ 1169ms | 否 | ✓ 1729ms | http |
| 59.46.216.131:30001 | ✓ 1204ms | ✓ 1408ms | ✓ 1105ms | 否 | ✓ 1781ms | http |
| 62.113.119.14:8080 | ✓ 1187ms | 否 | ✓ 1758ms | 否 | ✓ 1482ms | http |
| 45.88.0.115:3128 | ✓ 641ms | 否 | ✓ 810ms | ✓ 1451ms | ✓ 1172ms | http |
| 45.88.0.117:3128 | ✓ 641ms | 否 | ✓ 808ms | ✓ 1423ms | ✓ 1211ms | http |
| 213.220.62.63:3128 | ✓ 636ms | 否 | ✓ 814ms | ✓ 1440ms | ✓ 1184ms | http |
| 213.220.62.62:3128 | ✓ 641ms | 否 | ✓ 808ms | ✓ 1411ms | ✓ 1219ms | http |
| 45.88.0.99:3128 | ✓ 641ms | 否 | ✓ 806ms | ✓ 1457ms | ✓ 1191ms | http |
| 42.200.76.16:3888 | ✓ 1338ms | 否 | ✓ 775ms | ✓ 947ms | ✓ 757ms | http |
| 45.88.0.116:3128 | ✓ 641ms | ✓ 1594ms | ✓ 1573ms | ✓ 1412ms | ✓ 1118ms | http |
| 45.88.0.114:3128 | ✓ 641ms | ✓ 1601ms | ✓ 1563ms | ✓ 1420ms | ✓ 1150ms | http |
| 45.88.0.98:3128 | ✓ 641ms | ✓ 1806ms | ✓ 1533ms | ✓ 1374ms | ✓ 1114ms | http |
| 45.88.0.113:3128 | ✓ 639ms | ✓ 1815ms | ✓ 1530ms | ✓ 1395ms | 否 | http |
| 167.71.196.178:80 | ✓ 1490ms | 否 | ✓ 1167ms | ✓ 1048ms | ✓ 854ms | http |
| 161.97.184.191:8080 | 否 | ✓ 1702ms | ✓ 849ms | 否 | ✓ 1531ms | http |
| 147.45.166.46:3128 | ✓ 1699ms | ✓ 1818ms | ✓ 647ms | 否 | 否 | http |
| 38.55.106.208:6005 | ✓ 1162ms | 否 | ✓ 1058ms | ✓ 873ms | ✓ 712ms | http |
| 120.92.108.86:7890 | ✓ 1253ms | 否 | ✓ 1917ms | 否 | ✓ 1321ms | http |
| 78.11.96.22:8888 | ✓ 1027ms | ✓ 1777ms | ✓ 1148ms | ✓ 1645ms | ✓ 1456ms | http |
| 115.231.181.40:8128 | ✓ 1897ms | ✓ 1907ms | 否 | 否 | ✓ 1764ms | http |
| 125.64.244.100:8889 | 否 | ✓ 1593ms | ✓ 1672ms | ✓ 1855ms | ✓ 1784ms | http |
| 208.87.243.199:7878 | ✓ 878ms | 否 | ✓ 1835ms | ✓ 1548ms | ✓ 1055ms | http |
| 223.84.151.86:30005 | ✓ 1328ms | ✓ 1932ms | ✓ 1660ms | 否 | 否 | http |
| 202.141.161.53:10808 | ✓ 1330ms | ✓ 1890ms | ✓ 1057ms | ✓ 1153ms | ✓ 1011ms | http |
| 84.47.150.126:1080 | ✓ 1070ms | ✓ 1977ms | ✓ 1983ms | 否 | ✓ 1514ms | http |
| 103.82.23.118:5207 | ✓ 1324ms | 否 | 否 | ✓ 1793ms | ✓ 1616ms | http |
| 103.82.23.118:5185 | ✓ 1535ms | 否 | 否 | ✓ 1646ms | ✓ 1599ms | http |
| 84.47.150.125:1080 | ✓ 1964ms | 否 | 否 | ✓ 1794ms | ✓ 1545ms | http |
| 82.114.228.67:1080 | ✓ 1864ms | ✓ 1762ms | 否 | ✓ 1671ms | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1012ms | ✓ 1314ms | ✓ 1518ms | ✓ 1871ms | http |
| 85.190.99.143:443 | ✓ 1417ms | 否 | ✓ 802ms | 否 | ✓ 1730ms | http |
| 194.104.9.38:3128 | ✓ 1122ms | 否 | ✓ 1796ms | ✓ 1873ms | ✓ 1613ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1225ms | ✓ 1719ms | ✓ 1345ms | http |
| 34.101.184.164:3128 | ✓ 1873ms | 否 | ✓ 1740ms | ✓ 1734ms | ✓ 1605ms | http |
| 201.144.20.238:3128 | ✓ 681ms | ✓ 1331ms | ✓ 1375ms | ✓ 1456ms | ✓ 1213ms | http |
| 120.92.212.16:7890 | ✓ 1058ms | 否 | ✓ 1386ms | ✓ 1539ms | ✓ 1759ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1259ms | ✓ 1007ms | 否 | ✓ 949ms | http |
| 102.134.48.240:6005 | ✓ 572ms | ✓ 1025ms | ✓ 1420ms | ✓ 910ms | ✓ 816ms | http |
| 102.134.49.165:6005 | ✓ 667ms | ✓ 1056ms | ✓ 1297ms | ✓ 1238ms | ✓ 896ms | http |
| 103.229.126.221:7890 | ✓ 805ms | ✓ 1582ms | ✓ 1145ms | ✓ 913ms | ✓ 732ms | http |
| 38.55.107.137:6005 | ✓ 1791ms | 否 | ✓ 1091ms | ✓ 891ms | ✓ 713ms | http |
| 38.55.106.206:6005 | 否 | 否 | ✓ 1019ms | ✓ 941ms | ✓ 719ms | http |
| 38.55.107.254:6005 | 否 | 否 | ✓ 1351ms | ✓ 1247ms | ✓ 708ms | http |
| 38.55.104.182:6005 | 否 | 否 | ✓ 959ms | ✓ 1261ms | ✓ 1083ms | http |
| 38.55.104.8:6005 | 否 | 否 | ✓ 1059ms | ✓ 878ms | ✓ 714ms | http |
| 103.133.254.4:3128 | ✓ 1996ms | 否 | ✓ 1294ms | ✓ 1595ms | ✓ 1260ms | http |
| 45.93.29.147:6005 | ✓ 772ms | ✓ 1024ms | ✓ 1417ms | ✓ 1260ms | ✓ 1297ms | http |
| 45.93.30.241:6005 | 否 | ✓ 1016ms | ✓ 1381ms | ✓ 1465ms | ✓ 1088ms | http |
| 182.204.179.113:1080 | 否 | ✓ 1452ms | ✓ 1295ms | ✓ 1961ms | ✓ 1590ms | http |
| 45.93.31.93:6005 | ✓ 1841ms | 否 | ✓ 1491ms | ✓ 1229ms | ✓ 1161ms | http |
| 45.93.30.177:6005 | ✓ 848ms | ✓ 1687ms | 否 | ✓ 1708ms | ✓ 982ms | http |
| 45.93.28.159:6005 | ✓ 1967ms | ✓ 1290ms | ✓ 1673ms | ✓ 1588ms | ✓ 1330ms | http |
| 116.171.106.15:3443 | 否 | ✓ 1688ms | ✓ 1452ms | ✓ 1923ms | 否 | http |
| 77.110.113.24:40000 | ✓ 1434ms | 否 | ✓ 1637ms | 否 | ✓ 1966ms | http |
| 185.214.108.46:40000 | ✓ 1725ms | ✓ 1702ms | ✓ 1616ms | 否 | 否 | http |
| 159.89.191.221:3128 | ✓ 382ms | 否 | ✓ 1409ms | ✓ 1712ms | ✓ 1608ms | http |
| 103.82.23.118:5221 | ✓ 1340ms | 否 | ✓ 1833ms | 否 | ✓ 1877ms | http |
| 74.50.96.247:8888 | 否 | 否 | ✓ 436ms | ✓ 894ms | ✓ 1421ms | http |
| 38.55.104.99:6005 | ✓ 1029ms | 否 | ✓ 1011ms | ✓ 1119ms | ✓ 717ms | http |
| 38.55.104.68:6005 | ✓ 1063ms | 否 | 否 | ✓ 894ms | ✓ 707ms | http |
| 121.135.144.234:8546 | ✓ 1122ms | 否 | ✓ 1281ms | ✓ 1629ms | ✓ 1993ms | http |
| 20.127.128.70:8080 | ✓ 1173ms | 否 | ✓ 1177ms | 否 | ✓ 1730ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1032ms | 否 | ✓ 1579ms | ✓ 712ms | http |
| 207.177.122.144:8080 | ✓ 1352ms | 否 | ✓ 1688ms | 否 | ✓ 980ms | http |
| 130.61.30.221:8080 | ✓ 719ms | 否 | ✓ 1776ms | 否 | ✓ 1704ms | http |
| 101.32.243.189:80 | ✓ 1215ms | ✓ 1796ms | ✓ 1433ms | ✓ 1242ms | ✓ 1277ms | http |
| 152.42.177.32:8888 | ✓ 1008ms | 否 | ✓ 1202ms | ✓ 1468ms | ✓ 1245ms | http |
| 218.153.163.186:8181 | ✓ 1607ms | ✓ 1524ms | 否 | ✓ 1248ms | ✓ 848ms | http |
| 218.153.163.186:8623 | ✓ 1930ms | 否 | ✓ 1665ms | ✓ 1221ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1535ms | ✓ 1150ms | ✓ 874ms | ✓ 1165ms | ✓ 1227ms | http |
| 106.10.55.212:1121 | ✓ 1128ms | 否 | 否 | ✓ 1369ms | ✓ 1907ms | http |
| 103.113.70.189:1082 | ✓ 1246ms | ✓ 1960ms | ✓ 1119ms | ✓ 1713ms | ✓ 1346ms | http |
| 103.157.117.116:8080 | ✓ 1852ms | 否 | 否 | ✓ 1905ms | ✓ 1910ms | http |
| 207.254.71.62:8088 | ✓ 749ms | ✓ 1988ms | ✓ 1340ms | 否 | 否 | http |

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
