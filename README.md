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

最后更新：2026-06-14 12:49:31 UTC（2026-06-14 20:49:31 UTC+8）

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
| 159.198.35.187:1080 | ✓ 1127ms | 否 | ✓ 1172ms | 否 | ✓ 904ms | http |
| 34.43.46.91:80 | 否 | ✓ 1684ms | ✓ 229ms | ✓ 1111ms | ✓ 1437ms | http |
| 34.43.46.91:443 | 否 | ✓ 1400ms | ✓ 262ms | ✓ 1315ms | 否 | http |
| 104.154.186.48:80 | 否 | 否 | ✓ 945ms | ✓ 1240ms | ✓ 946ms | http |
| 212.118.52.196:1080 | ✓ 973ms | 否 | ✓ 1003ms | ✓ 1641ms | ✓ 1358ms | http |
| 91.107.168.255:83 | 否 | 否 | ✓ 1455ms | ✓ 1249ms | ✓ 943ms | http |
| 185.200.188.234:10001 | ✓ 1007ms | 否 | ✓ 706ms | 否 | ✓ 1482ms | http |
| 91.107.168.255:82 | 否 | 否 | ✓ 423ms | ✓ 1985ms | ✓ 1432ms | http |
| 64.188.125.187:3128 | ✓ 992ms | 否 | ✓ 696ms | 否 | ✓ 1987ms | http |
| 176.111.37.216:39811 | ✓ 1008ms | 否 | ✓ 1334ms | 否 | ✓ 1724ms | http |
| 1.231.81.166:3128 | ✓ 1725ms | ✓ 1507ms | ✓ 1598ms | ✓ 1283ms | ✓ 1486ms | http |
| 47.79.119.13:8080 | 否 | 否 | ✓ 972ms | ✓ 1294ms | ✓ 1024ms | http |
| 82.97.247.37:80 | ✓ 997ms | ✓ 1674ms | ✓ 1040ms | 否 | ✓ 1387ms | http |
| 81.200.154.236:48503 | ✓ 464ms | ✓ 1666ms | ✓ 411ms | ✓ 1487ms | ✓ 1125ms | http |
| 170.106.136.181:31002 | 否 | ✓ 1569ms | ✓ 415ms | ✓ 886ms | ✓ 657ms | http |
| 52.188.28.218:3128 | ✓ 818ms | 否 | ✓ 779ms | ✓ 1660ms | ✓ 1402ms | http |
| 176.111.37.5:39811 | ✓ 445ms | ✓ 1551ms | ✓ 451ms | ✓ 1629ms | 否 | http |
| 144.172.114.214:1080 | 否 | 否 | ✓ 390ms | ✓ 1158ms | ✓ 1306ms | http |
| 138.124.114.42:7443 | ✓ 424ms | 否 | ✓ 742ms | ✓ 1494ms | ✓ 1760ms | http |
| 185.141.26.131:3128 | ✓ 494ms | 否 | ✓ 902ms | ✓ 1773ms | ✓ 1501ms | http |
| 57.129.144.178:40000 | ✓ 600ms | ✓ 1365ms | ✓ 1640ms | ✓ 1973ms | ✓ 1518ms | http |
| 191.40.5.0:8888 | 否 | ✓ 1517ms | ✓ 821ms | ✓ 1852ms | ✓ 1171ms | http |
| 213.165.42.185:7443 | ✓ 1934ms | 否 | ✓ 391ms | ✓ 1800ms | ✓ 1414ms | http |
| 91.107.182.124:82 | ✓ 1278ms | 否 | ✓ 1272ms | 否 | ✓ 1767ms | http |
| 89.169.53.40:7443 | ✓ 1001ms | ✓ 1632ms | ✓ 1966ms | 否 | 否 | http |
| 151.243.180.211:2080 | ✓ 1501ms | ✓ 1524ms | ✓ 526ms | 否 | ✓ 1554ms | http |
| 84.47.150.125:1080 | ✓ 1708ms | 否 | ✓ 1606ms | 否 | ✓ 1683ms | http |
| 85.192.28.62:7443 | ✓ 1761ms | 否 | 否 | ✓ 1890ms | ✓ 1807ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1402ms | ✓ 1896ms | ✓ 1167ms | http |
| 82.102.11.164:3460 | ✓ 934ms | ✓ 1483ms | ✓ 1714ms | ✓ 1768ms | ✓ 1498ms | http |
| 185.11.134.227:8443 | ✓ 1364ms | ✓ 1559ms | ✓ 1145ms | 否 | ✓ 1784ms | http |
| 91.186.213.124:1081 | ✓ 1096ms | 否 | ✓ 750ms | 否 | ✓ 1886ms | http |
| 51.250.86.201:2080 | ✓ 971ms | 否 | ✓ 1304ms | ✓ 1956ms | 否 | http |
| 144.31.134.103:1080 | ✓ 1311ms | 否 | ✓ 550ms | ✓ 1587ms | 否 | http |
| 159.223.87.50:443 | ✓ 935ms | 否 | ✓ 949ms | ✓ 1396ms | ✓ 1302ms | http |
| 95.3.69.222:8080 | ✓ 1392ms | ✓ 1732ms | ✓ 1177ms | ✓ 1676ms | ✓ 1455ms | http |
| 77.246.104.106:4433 | ✓ 1249ms | 否 | ✓ 1501ms | ✓ 1471ms | ✓ 1231ms | http |
| 3.137.86.220:443 | 否 | 否 | ✓ 402ms | ✓ 1766ms | ✓ 1329ms | http |
| 217.154.155.115:8080 | ✓ 950ms | 否 | 否 | ✓ 1913ms | ✓ 1392ms | http |
| 8.154.21.175:3128 | ✓ 1028ms | ✓ 1579ms | ✓ 1134ms | 否 | ✓ 1117ms | http |
| 157.245.100.190:442 | 否 | 否 | ✓ 1974ms | ✓ 1245ms | ✓ 979ms | http |
| 152.32.132.190:7890 | ✓ 1351ms | 否 | ✓ 1425ms | ✓ 1312ms | 否 | http |
| 192.9.182.6:20172 | ✓ 1708ms | 否 | ✓ 1299ms | ✓ 1555ms | 否 | http |
| 188.137.252.194:1080 | ✓ 878ms | 否 | 否 | ✓ 1859ms | ✓ 1001ms | http |
| 147.45.111.209:2080 | ✓ 1086ms | 否 | ✓ 1662ms | ✓ 1215ms | 否 | http |
| 158.255.212.55:3256 | ✓ 575ms | 否 | ✓ 1583ms | ✓ 1542ms | 否 | http |
| 158.255.212.55:7497 | ✓ 575ms | 否 | ✓ 1584ms | ✓ 1545ms | 否 | http |
| 158.255.212.55:9300 | ✓ 573ms | 否 | ✓ 1584ms | ✓ 1551ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1267ms | 否 | ✓ 1069ms | ✓ 1645ms | ✓ 1306ms | http |
| 62.133.62.249:1082 | ✓ 930ms | 否 | ✓ 1255ms | 否 | ✓ 1807ms | http |
| 103.99.115.75:40000 | 否 | ✓ 1427ms | ✓ 331ms | ✓ 1224ms | ✓ 1267ms | http |
| 62.133.62.17:1082 | ✓ 1317ms | 否 | ✓ 812ms | ✓ 1975ms | ✓ 1315ms | http |
| 85.192.28.47:7443 | ✓ 989ms | 否 | ✓ 1317ms | ✓ 1880ms | 否 | http |
| 91.233.223.147:3128 | ✓ 996ms | 否 | ✓ 1763ms | 否 | ✓ 1842ms | http |
| 169.212.15.161:5000 | 否 | ✓ 1680ms | 否 | ✓ 1262ms | ✓ 1230ms | http |
| 79.137.205.130:7443 | ✓ 420ms | 否 | ✓ 718ms | ✓ 1508ms | ✓ 1066ms | http |
| 91.107.172.30:82 | ✓ 444ms | 否 | ✓ 701ms | 否 | ✓ 1659ms | http |
| 45.88.174.195:8080 | ✓ 423ms | 否 | ✓ 1296ms | 否 | ✓ 1434ms | http |
| 37.49.224.15:3128 | ✓ 1207ms | 否 | ✓ 1468ms | ✓ 1661ms | ✓ 1674ms | http |
| 212.58.132.5:8888 | ✓ 1483ms | 否 | 否 | ✓ 1595ms | ✓ 1718ms | http |
| 154.206.67.83:9000 | ✓ 1626ms | 否 | ✓ 827ms | ✓ 1046ms | 否 | http |
| 34.87.80.221:30000 | ✓ 1066ms | 否 | ✓ 1552ms | ✓ 1304ms | ✓ 1095ms | http |
| 202.28.194.139:31280 | 否 | 否 | ✓ 1808ms | ✓ 1991ms | ✓ 1883ms | http |
| 62.246.118.139:3128 | ✓ 1151ms | 否 | ✓ 1646ms | 否 | ✓ 1663ms | http |
| 149.104.4.88:10809 | ✓ 1034ms | 否 | ✓ 1005ms | ✓ 1101ms | ✓ 858ms | http |
| 18.180.59.181:80 | ✓ 825ms | ✓ 1076ms | ✓ 1739ms | ✓ 1212ms | ✓ 900ms | http |
| 150.241.116.167:443 | ✓ 612ms | 否 | ✓ 679ms | ✓ 1511ms | 否 | http |
| 105.156.113.242:5555 | ✓ 1359ms | ✓ 1858ms | ✓ 1784ms | 否 | 否 | http |
| 95.140.154.156:1080 | ✓ 1925ms | 否 | ✓ 1575ms | ✓ 1550ms | ✓ 1818ms | http |
| 85.192.60.187:7443 | ✓ 1832ms | 否 | ✓ 1262ms | 否 | ✓ 1945ms | http |
| 85.234.100.149:8080 | ✓ 902ms | 否 | ✓ 1306ms | ✓ 1640ms | ✓ 1147ms | http |
| 160.238.65.2:3128 | ✓ 1505ms | ✓ 1914ms | ✓ 760ms | ✓ 1321ms | ✓ 1298ms | http |
| 83.147.36.155:8080 | ✓ 1987ms | 否 | ✓ 1016ms | ✓ 1573ms | 否 | http |
| 160.238.65.7:3128 | ✓ 1246ms | 否 | 否 | ✓ 1344ms | ✓ 1318ms | http |
| 160.238.65.8:3128 | ✓ 1245ms | 否 | 否 | ✓ 1344ms | ✓ 1246ms | http |
| 194.87.55.194:2053 | 否 | 否 | ✓ 425ms | ✓ 1268ms | ✓ 1282ms | http |
| 85.234.100.149:1080 | ✓ 1695ms | 否 | 否 | ✓ 1241ms | ✓ 1624ms | http |
| 89.127.207.174:18080 | ✓ 974ms | 否 | ✓ 1989ms | ✓ 1731ms | ✓ 1326ms | http |
| 45.84.222.25:1080 | ✓ 1279ms | 否 | ✓ 1568ms | 否 | ✓ 1718ms | http |
| 160.238.65.9:3128 | ✓ 1425ms | ✓ 1437ms | ✓ 419ms | 否 | ✓ 969ms | http |
| 160.238.65.5:3128 | ✓ 748ms | ✓ 1433ms | 否 | ✓ 1242ms | ✓ 1299ms | http |
| 160.238.65.6:3128 | ✓ 432ms | ✓ 1833ms | ✓ 436ms | 否 | 否 | http |
| 160.238.65.3:3128 | ✓ 420ms | ✓ 1743ms | 否 | 否 | ✓ 965ms | http |
| 160.238.65.4:3128 | ✓ 1760ms | ✓ 1638ms | ✓ 1522ms | ✓ 1267ms | 否 | http |
| 62.133.62.17:1081 | ✓ 1877ms | 否 | ✓ 852ms | ✓ 1753ms | 否 | http |
| 144.202.14.153:50000 | 否 | 否 | ✓ 114ms | ✓ 1336ms | ✓ 837ms | http |
| 61.52.131.172:8443 | ✓ 1074ms | ✓ 1401ms | ✓ 1202ms | ✓ 1412ms | ✓ 1166ms | http |
| 36.147.78.166:443 | 否 | ✓ 1990ms | ✓ 1946ms | 否 | ✓ 1899ms | http |
| 103.172.70.173:8080 | ✓ 1810ms | 否 | 否 | ✓ 1481ms | ✓ 1731ms | http |
| 45.82.153.23:80 | ✓ 700ms | 否 | ✓ 1671ms | 否 | ✓ 1989ms | http |
| 144.31.134.29:1050 | ✓ 992ms | 否 | ✓ 999ms | ✓ 1630ms | 否 | http |

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
