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

最后更新：2026-06-07 17:08:33 UTC（2026-06-08 01:08:33 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.43.46.91:80 | ✓ 550ms | ✓ 1292ms | ✓ 268ms | ✓ 1171ms | ✓ 1135ms | http |
| 104.154.186.48:80 | ✓ 310ms | ✓ 1190ms | ✓ 889ms | ✓ 1358ms | ✓ 912ms | http |
| 104.161.37.187:3128 | ✓ 555ms | 否 | ✓ 299ms | 否 | ✓ 1310ms | http |
| 185.200.188.234:10001 | ✓ 942ms | 否 | ✓ 1115ms | 否 | ✓ 1505ms | http |
| 169.212.15.161:5000 | 否 | 否 | ✓ 863ms | ✓ 1294ms | ✓ 1028ms | http |
| 176.111.37.216:39811 | ✓ 650ms | ✓ 1596ms | 否 | 否 | ✓ 1886ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1458ms | ✓ 1944ms | ✓ 1357ms | ✓ 1310ms | http |
| 113.160.132.26:8080 | ✓ 1713ms | 否 | ✓ 1404ms | ✓ 1464ms | ✓ 1462ms | http |
| 38.123.220.147:999 | ✓ 676ms | ✓ 1907ms | ✓ 1222ms | ✓ 1398ms | ✓ 1266ms | http |
| 129.153.7.7:60000 | ✓ 329ms | ✓ 859ms | ✓ 422ms | 否 | 否 | http |
| 117.55.203.162:8899 | ✓ 1497ms | 否 | ✓ 1428ms | ✓ 1779ms | ✓ 1532ms | http |
| 207.211.161.235:8888 | 否 | 否 | ✓ 1240ms | ✓ 1072ms | ✓ 1649ms | http |
| 185.141.26.131:3128 | ✓ 609ms | 否 | ✓ 1277ms | 否 | ✓ 1833ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1577ms | ✓ 1380ms | ✓ 1180ms | ✓ 1527ms | http |
| 94.241.175.40:10808 | ✓ 454ms | 否 | ✓ 814ms | ✓ 1129ms | 否 | http |
| 188.225.58.59:443 | ✓ 1144ms | 否 | ✓ 1511ms | 否 | ✓ 1859ms | http |
| 45.159.79.101:3128 | ✓ 748ms | 否 | ✓ 1323ms | ✓ 1233ms | ✓ 1193ms | http |
| 209.38.200.247:1080 | ✓ 908ms | ✓ 1961ms | ✓ 512ms | 否 | ✓ 1921ms | http |
| 3.110.181.109:3128 | ✓ 979ms | ✓ 1580ms | ✓ 1008ms | ✓ 1632ms | ✓ 1261ms | http |
| 84.47.150.125:1080 | ✓ 1168ms | 否 | ✓ 1443ms | 否 | ✓ 1687ms | http |
| 81.200.154.236:48503 | ✓ 918ms | ✓ 1658ms | ✓ 844ms | ✓ 1842ms | 否 | http |
| 202.28.194.139:31280 | ✓ 1816ms | 否 | 否 | ✓ 1866ms | ✓ 1948ms | http |
| 137.59.47.73:3128 | ✓ 1682ms | 否 | ✓ 1141ms | 否 | ✓ 1069ms | http |
| 170.106.136.181:31002 | ✓ 745ms | ✓ 995ms | ✓ 660ms | ✓ 1003ms | ✓ 1870ms | http |
| 216.9.225.157:3128 | ✓ 713ms | 否 | ✓ 868ms | 否 | ✓ 1423ms | http |
| 34.87.80.221:30000 | ✓ 1762ms | ✓ 1696ms | ✓ 1475ms | ✓ 1955ms | ✓ 1084ms | http |
| 2.26.68.177:8080 | ✓ 1149ms | 否 | ✓ 1558ms | ✓ 1598ms | ✓ 1545ms | http |
| 150.241.100.167:8443 | ✓ 418ms | ✓ 1478ms | ✓ 1320ms | 否 | ✓ 1276ms | http |
| 8.219.97.248:80 | ✓ 1788ms | 否 | ✓ 1293ms | ✓ 1525ms | 否 | http |
| 117.55.203.161:8899 | ✓ 843ms | 否 | ✓ 846ms | ✓ 1587ms | ✓ 1265ms | http |
| 110.172.62.196:8080 | ✓ 1159ms | ✓ 1127ms | ✓ 1412ms | ✓ 1339ms | ✓ 1109ms | http |
| 43.160.236.170:8888 | ✓ 925ms | 否 | ✓ 938ms | ✓ 1491ms | ✓ 1083ms | http |
| 85.234.100.149:8080 | 否 | 否 | ✓ 1292ms | ✓ 1665ms | ✓ 972ms | http |
| 37.49.224.15:3128 | ✓ 616ms | 否 | ✓ 1566ms | ✓ 1805ms | 否 | http |
| 85.234.100.149:1080 | 否 | 否 | ✓ 652ms | ✓ 1578ms | ✓ 1136ms | http |
| 185.233.186.88:443 | ✓ 1480ms | 否 | ✓ 995ms | ✓ 1972ms | ✓ 1861ms | http |
| 183.80.40.243:2051 | 否 | 否 | ✓ 1685ms | ✓ 1913ms | ✓ 1882ms | http |
| 59.66.28.115:6382 | ✓ 1688ms | ✓ 1628ms | ✓ 1442ms | 否 | 否 | http |
| 43.128.145.26:1080 | ✓ 1672ms | 否 | ✓ 880ms | 否 | ✓ 1597ms | http |
| 36.50.205.70:8080 | ✓ 1594ms | 否 | 否 | ✓ 1621ms | ✓ 1405ms | http |
| 195.25.20.155:3128 | ✓ 403ms | 否 | ✓ 1213ms | 否 | ✓ 1140ms | http |
| 50.114.102.16:8888 | ✓ 1941ms | ✓ 1562ms | ✓ 414ms | ✓ 1523ms | ✓ 1427ms | http |
| 165.227.133.230:8888 | ✓ 411ms | 否 | ✓ 378ms | 否 | ✓ 1232ms | http |
| 157.245.143.65:7890 | ✓ 1344ms | ✓ 1872ms | ✓ 469ms | ✓ 932ms | 否 | http |
| 14.143.222.113:10158 | ✓ 952ms | 否 | ✓ 939ms | ✓ 1374ms | 否 | http |
| 121.230.8.55:1080 | ✓ 1321ms | 否 | ✓ 1240ms | ✓ 1908ms | ✓ 1391ms | http |
| 121.43.196.210:8222 | ✓ 1072ms | ✓ 1288ms | ✓ 1060ms | ✓ 1373ms | ✓ 1163ms | http |
| 121.43.196.213:8222 | ✓ 1092ms | ✓ 1343ms | ✓ 1023ms | ✓ 1388ms | ✓ 1093ms | http |
| 160.238.65.5:3128 | ✓ 1059ms | ✓ 1655ms | 否 | 否 | ✓ 1737ms | http |
| 43.228.215.32:8080 | ✓ 1878ms | 否 | ✓ 964ms | ✓ 1463ms | ✓ 1606ms | http |
| 116.104.252.1:2113 | ✓ 1915ms | 否 | ✓ 1701ms | 否 | ✓ 1710ms | http |
| 116.104.252.1:2045 | ✓ 1659ms | 否 | ✓ 1564ms | 否 | ✓ 1729ms | http |
| 18.180.59.181:80 | ✓ 808ms | ✓ 947ms | ✓ 1203ms | ✓ 1460ms | ✓ 1078ms | http |
| 80.150.246.98:443 | ✓ 461ms | ✓ 1343ms | ✓ 1930ms | 否 | ✓ 1583ms | http |
| 8.154.21.175:3128 | ✓ 1104ms | ✓ 1238ms | ✓ 1062ms | ✓ 1291ms | ✓ 1107ms | http |
| 43.134.141.85:80 | ✓ 1435ms | 否 | 否 | ✓ 1780ms | ✓ 1496ms | http |
| es-xh-01.hpdata.click:443 | ✓ 940ms | ✓ 1983ms | ✓ 1537ms | 否 | ✓ 1560ms | http |
| 103.209.36.58:8080 | ✓ 1438ms | 否 | 否 | ✓ 1726ms | ✓ 1464ms | http |
| 58.187.104.56:2104 | 否 | 否 | ✓ 1730ms | ✓ 1917ms | ✓ 1990ms | http |
| 45.88.174.195:8080 | ✓ 917ms | ✓ 1874ms | ✓ 944ms | 否 | ✓ 1653ms | http |
| 212.58.132.5:8888 | ✓ 1584ms | 否 | ✓ 1413ms | ✓ 1470ms | ✓ 1183ms | http |
| 86.104.74.110:1082 | ✓ 784ms | ✓ 1944ms | ✓ 995ms | ✓ 1478ms | ✓ 1111ms | http |
| 169.40.6.114:3128 | 否 | 否 | ✓ 784ms | ✓ 1826ms | ✓ 1542ms | http |
| 116.104.252.1:2091 | ✓ 1668ms | 否 | ✓ 1624ms | 否 | ✓ 1919ms | http |
| 103.28.37.131:3128 | ✓ 1039ms | 否 | ✓ 1185ms | ✓ 1345ms | 否 | http |
| 92.118.112.32:1082 | ✓ 425ms | ✓ 977ms | ✓ 1686ms | ✓ 1234ms | 否 | http |
| 172.234.65.27:3128 | ✓ 425ms | 否 | 否 | ✓ 1366ms | ✓ 1093ms | http |
| 80.64.17.112:18080 | ✓ 915ms | ✓ 1521ms | ✓ 1381ms | ✓ 1783ms | 否 | http |
| 46.39.105.157:8080 | ✓ 1204ms | 否 | ✓ 1378ms | 否 | ✓ 1289ms | http |
| 92.118.112.32:1081 | ✓ 430ms | ✓ 998ms | ✓ 1665ms | ✓ 1234ms | ✓ 1284ms | http |
| 217.76.245.80:999 | 否 | ✓ 1683ms | ✓ 983ms | ✓ 1319ms | ✓ 1355ms | http |
| 152.32.132.190:7890 | ✓ 839ms | 否 | 否 | ✓ 1028ms | ✓ 1707ms | http |
| 136.0.3.35:1234 | ✓ 1300ms | 否 | ✓ 564ms | ✓ 1302ms | ✓ 870ms | http |
| 91.149.222.102:22335 | ✓ 1244ms | ✓ 1793ms | ✓ 1177ms | 否 | ✓ 1886ms | http |
| 92.118.112.25:1082 | ✓ 1614ms | 否 | ✓ 1323ms | ✓ 1501ms | 否 | http |
| 45.84.222.25:1080 | ✓ 420ms | 否 | ✓ 949ms | ✓ 1302ms | ✓ 1405ms | http |
| 121.230.8.220:1080 | 否 | ✓ 1550ms | ✓ 1363ms | ✓ 1717ms | 否 | http |
| 116.171.106.111:3443 | ✓ 1735ms | ✓ 1810ms | ✓ 1450ms | 否 | 否 | http |
| 121.230.8.251:1080 | ✓ 1189ms | ✓ 1875ms | 否 | ✓ 1683ms | 否 | http |
| 121.230.8.99:1080 | 否 | 否 | ✓ 1294ms | ✓ 1674ms | ✓ 1943ms | http |
| 116.104.252.1:2087 | ✓ 1609ms | 否 | ✓ 1571ms | ✓ 1846ms | ✓ 1832ms | http |
| 85.121.51.40:9091 | 否 | 否 | ✓ 1064ms | ✓ 1604ms | ✓ 903ms | http |
| 138.124.93.170:3129 | ✓ 918ms | 否 | 否 | ✓ 1921ms | ✓ 1688ms | http |
| 218.108.131.186:17890 | ✓ 1044ms | ✓ 1314ms | ✓ 987ms | ✓ 1431ms | ✓ 1098ms | http |
| 122.233.222.209:2222 | 否 | ✓ 1359ms | 否 | ✓ 1601ms | ✓ 1173ms | http |
| 138.2.239.213:10010 | ✓ 1149ms | 否 | 否 | ✓ 1499ms | ✓ 1202ms | http |
| 123.57.0.163:8888 | ✓ 1703ms | ✓ 1677ms | ✓ 1690ms | 否 | 否 | http |

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
