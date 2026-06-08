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

最后更新：2026-06-08 01:00:02 UTC（2026-06-08 09:00:02 UTC+8）

**代理总数：109**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 109 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 109 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.43.46.91:80 | ✓ 910ms | ✓ 1460ms | ✓ 141ms | ✓ 1328ms | ✓ 1211ms | http |
| 34.43.46.91:443 | ✓ 913ms | ✓ 1142ms | ✓ 419ms | ✓ 1544ms | ✓ 1205ms | http |
| 104.161.37.187:3128 | ✓ 531ms | ✓ 1106ms | ✓ 1079ms | 否 | ✓ 927ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1424ms | ✓ 947ms | ✓ 1351ms | ✓ 1117ms | http |
| 152.32.132.190:7890 | ✓ 708ms | 否 | 否 | ✓ 1116ms | ✓ 1463ms | http |
| 169.212.15.161:5000 | 否 | ✓ 1243ms | ✓ 733ms | ✓ 1811ms | 否 | http |
| 117.55.203.161:8899 | ✓ 1358ms | 否 | ✓ 1498ms | 否 | ✓ 1432ms | http |
| 167.86.95.198:3128 | ✓ 1221ms | ✓ 1681ms | ✓ 1583ms | 否 | ✓ 1820ms | http |
| 117.55.203.162:8899 | ✓ 1358ms | ✓ 1836ms | ✓ 1663ms | 否 | 否 | http |
| 185.233.186.88:443 | ✓ 1263ms | 否 | ✓ 1301ms | 否 | ✓ 1703ms | http |
| 185.200.188.234:10001 | ✓ 1117ms | 否 | ✓ 1759ms | 否 | ✓ 1836ms | http |
| 188.225.58.59:443 | ✓ 1268ms | 否 | ✓ 1297ms | 否 | ✓ 1693ms | http |
| 43.228.215.32:8080 | 否 | 否 | ✓ 1460ms | ✓ 1836ms | ✓ 964ms | http |
| 104.154.186.48:80 | ✓ 461ms | ✓ 1138ms | ✓ 840ms | ✓ 1132ms | ✓ 965ms | http |
| 129.153.7.7:60000 | 否 | ✓ 1467ms | ✓ 206ms | ✓ 1395ms | ✓ 785ms | http |
| 57.129.144.178:40000 | ✓ 666ms | 否 | ✓ 1085ms | ✓ 1772ms | ✓ 1419ms | http |
| 209.38.200.247:1080 | ✓ 638ms | ✓ 1834ms | ✓ 539ms | ✓ 1253ms | ✓ 1289ms | http |
| 165.227.133.230:8888 | ✓ 726ms | 否 | ✓ 920ms | 否 | ✓ 1570ms | http |
| 176.111.37.216:39811 | ✓ 529ms | ✓ 1873ms | ✓ 525ms | ✓ 1759ms | ✓ 1590ms | http |
| 147.45.78.89:1080 | ✓ 633ms | 否 | ✓ 1334ms | 否 | ✓ 1620ms | http |
| 62.210.136.222:3128 | ✓ 1889ms | ✓ 1512ms | ✓ 1487ms | ✓ 1628ms | ✓ 1612ms | http |
| 185.141.26.131:3128 | ✓ 1179ms | 否 | ✓ 619ms | 否 | ✓ 1677ms | http |
| 110.172.62.196:8080 | ✓ 1425ms | ✓ 1133ms | ✓ 1320ms | ✓ 1169ms | ✓ 1128ms | http |
| 2.26.87.216:1080 | ✓ 758ms | 否 | ✓ 1182ms | ✓ 1904ms | 否 | http |
| 202.28.194.139:31280 | ✓ 1782ms | 否 | ✓ 1973ms | 否 | ✓ 1921ms | http |
| 170.106.136.181:31002 | 否 | ✓ 887ms | ✓ 699ms | ✓ 744ms | ✓ 567ms | http |
| 94.241.175.40:10808 | ✓ 488ms | ✓ 1440ms | ✓ 429ms | ✓ 1176ms | 否 | http |
| 31.40.204.188:443 | ✓ 1193ms | 否 | ✓ 1513ms | ✓ 1883ms | 否 | http |
| 95.3.69.222:8080 | ✓ 1184ms | 否 | ✓ 1114ms | ✓ 1824ms | ✓ 1600ms | http |
| 81.200.154.236:48503 | ✓ 502ms | 否 | ✓ 915ms | ✓ 1706ms | ✓ 898ms | http |
| 199.127.62.89:3129 | ✓ 1034ms | 否 | ✓ 1118ms | 否 | ✓ 1654ms | http |
| 45.76.78.247:10001 | ✓ 1198ms | 否 | ✓ 1006ms | ✓ 1942ms | ✓ 1830ms | http |
| 85.234.100.149:8080 | ✓ 590ms | 否 | ✓ 825ms | 否 | ✓ 1410ms | http |
| 116.104.252.1:2105 | ✓ 1825ms | 否 | ✓ 1436ms | ✓ 1584ms | ✓ 1432ms | http |
| 45.84.222.25:1080 | 否 | 否 | ✓ 1170ms | ✓ 1378ms | ✓ 1032ms | http |
| 50.114.102.16:8888 | ✓ 603ms | ✓ 1644ms | ✓ 1312ms | 否 | 否 | http |
| 8.154.21.175:3128 | ✓ 922ms | ✓ 1136ms | ✓ 918ms | ✓ 1175ms | ✓ 992ms | http |
| 1.231.81.166:3128 | ✓ 1398ms | ✓ 1065ms | ✓ 1609ms | ✓ 1444ms | ✓ 929ms | http |
| 43.128.145.26:1080 | 否 | 否 | ✓ 1281ms | ✓ 1859ms | ✓ 790ms | http |
| 116.104.252.1:2045 | ✓ 1467ms | 否 | 否 | ✓ 1635ms | ✓ 1452ms | http |
| 116.80.50.48:3172 | ✓ 1566ms | 否 | ✓ 1586ms | 否 | ✓ 1731ms | http |
| 47.74.226.8:5001 | 否 | 否 | ✓ 1149ms | ✓ 1375ms | ✓ 1345ms | http |
| 195.25.20.155:3128 | ✓ 814ms | 否 | ✓ 574ms | ✓ 1676ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1947ms | 否 | ✓ 1579ms | ✓ 1448ms | ✓ 1330ms | http |
| 37.49.224.15:3128 | ✓ 1080ms | 否 | ✓ 1749ms | ✓ 1832ms | ✓ 1787ms | http |
| 147.45.170.190:3128 | ✓ 1870ms | 否 | ✓ 955ms | ✓ 1683ms | 否 | http |
| 207.211.161.235:8888 | ✓ 1185ms | ✓ 928ms | ✓ 884ms | 否 | ✓ 1373ms | http |
| 84.47.150.125:1080 | ✓ 999ms | 否 | ✓ 1899ms | 否 | ✓ 1565ms | http |
| 81.168.119.85:443 | ✓ 1261ms | 否 | ✓ 1541ms | 否 | ✓ 1803ms | http |
| 136.0.3.35:1234 | ✓ 1988ms | 否 | ✓ 298ms | ✓ 1089ms | 否 | http |
| 14.143.222.113:10158 | ✓ 1018ms | 否 | ✓ 1007ms | ✓ 1453ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1341ms | 否 | ✓ 1661ms | 否 | ✓ 1721ms | http |
| 116.104.252.1:2070 | 否 | 否 | ✓ 1448ms | ✓ 1662ms | ✓ 1462ms | http |
| 216.9.225.157:3128 | ✓ 1486ms | 否 | ✓ 373ms | ✓ 1795ms | 否 | http |
| 45.88.174.195:8080 | ✓ 1015ms | ✓ 1447ms | ✓ 1723ms | ✓ 1970ms | ✓ 1555ms | http |
| 138.2.239.213:10010 | ✓ 1156ms | 否 | ✓ 1793ms | ✓ 1894ms | ✓ 1303ms | http |
| 138.68.101.169:3128 | ✓ 1965ms | 否 | 否 | ✓ 1791ms | ✓ 1663ms | http |
| 103.157.117.226:81 | ✓ 1319ms | 否 | ✓ 1324ms | ✓ 1466ms | ✓ 1454ms | http |
| 103.125.117.106:8080 | ✓ 1988ms | 否 | ✓ 1329ms | ✓ 1356ms | ✓ 1623ms | http |
| 49.128.211.157:3128 | ✓ 768ms | ✓ 1605ms | ✓ 642ms | ✓ 1084ms | ✓ 792ms | http |
| 18.180.59.181:80 | ✓ 680ms | ✓ 1063ms | ✓ 1503ms | ✓ 1229ms | ✓ 1116ms | http |
| 218.108.131.186:17890 | ✓ 988ms | ✓ 1115ms | ✓ 1003ms | ✓ 1287ms | ✓ 983ms | http |
| 123.57.2.231:2020 | ✓ 1093ms | ✓ 1251ms | ✓ 951ms | ✓ 1249ms | ✓ 1236ms | http |
| 91.217.149.240:8080 | ✓ 704ms | 否 | ✓ 1056ms | ✓ 1885ms | ✓ 1540ms | http |
| 43.134.141.85:80 | ✓ 1402ms | ✓ 1501ms | ✓ 1495ms | ✓ 1581ms | ✓ 1371ms | http |
| 43.156.228.168:80 | ✓ 1402ms | 否 | ✓ 1130ms | ✓ 1315ms | ✓ 1648ms | http |
| 46.8.112.212:3128 | ✓ 1144ms | 否 | 否 | ✓ 1969ms | ✓ 1901ms | http |
| es-xh-01.hpdata.click:443 | ✓ 1021ms | 否 | ✓ 1278ms | 否 | ✓ 1596ms | http |
| fr-xh-01.hpdata.click:443 | ✓ 955ms | ✓ 1719ms | ✓ 1542ms | 否 | 否 | http |
| uk-xh-01.hpdata.click:443 | ✓ 1056ms | 否 | ✓ 1591ms | 否 | ✓ 1907ms | http |
| ch-xh-01.hpdata.click:443 | ✓ 1530ms | 否 | ✓ 1197ms | 否 | ✓ 1919ms | http |
| 85.234.100.149:1080 | ✓ 1039ms | 否 | ✓ 911ms | ✓ 1566ms | 否 | http |
| 116.104.250.118:2026 | 否 | 否 | ✓ 1480ms | ✓ 1649ms | ✓ 1432ms | http |
| 103.250.128.8:8082 | ✓ 1473ms | 否 | ✓ 1409ms | 否 | ✓ 1674ms | http |
| 103.214.102.154:8083 | 否 | 否 | ✓ 1460ms | ✓ 1880ms | ✓ 1441ms | http |
| 157.245.143.65:7890 | ✓ 275ms | ✓ 1891ms | ✓ 272ms | ✓ 1539ms | ✓ 979ms | http |
| 169.40.6.114:3128 | ✓ 860ms | 否 | ✓ 1674ms | 否 | ✓ 1839ms | http |
| 59.66.245.22:6382 | ✓ 1040ms | ✓ 1190ms | ✓ 1304ms | ✓ 1208ms | ✓ 952ms | http |
| 36.147.78.166:443 | 否 | 否 | ✓ 1813ms | ✓ 1999ms | ✓ 1697ms | http |
| 162.222.206.167:8080 | ✓ 1576ms | 否 | ✓ 925ms | 否 | ✓ 1219ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1342ms | ✓ 1417ms | ✓ 1426ms | ✓ 969ms | http |
| 164.92.165.209:18080 | ✓ 1901ms | 否 | ✓ 1886ms | 否 | ✓ 1993ms | http |
| 139.59.105.64:8080 | ✓ 1080ms | 否 | ✓ 1119ms | ✓ 1383ms | ✓ 1358ms | http |
| 216.8.184.119:80 | ✓ 761ms | ✓ 1606ms | ✓ 1098ms | ✓ 1114ms | ✓ 898ms | http |
| 217.209.35.22:443 | ✓ 1231ms | 否 | ✓ 1014ms | 否 | ✓ 1677ms | http |
| 223.204.177.1:3128 | ✓ 1775ms | 否 | ✓ 1106ms | 否 | ✓ 1595ms | http |
| 213.165.43.205:3128 | ✓ 1588ms | 否 | ✓ 963ms | ✓ 1400ms | ✓ 1278ms | http |
| 125.129.15.95:3128 | ✓ 842ms | ✓ 1774ms | ✓ 979ms | ✓ 1520ms | 否 | http |
| 34.84.162.206:38080 | ✓ 1274ms | ✓ 1301ms | ✓ 1694ms | 否 | 否 | http |
| 139.162.46.62:3128 | 否 | 否 | ✓ 1775ms | ✓ 1690ms | ✓ 913ms | http |
| 116.104.252.1:2091 | ✓ 1640ms | 否 | ✓ 1401ms | ✓ 1890ms | ✓ 1439ms | http |
| 168.235.80.215:3128 | 否 | 否 | ✓ 1232ms | ✓ 1085ms | ✓ 802ms | http |
| 116.104.252.1:2030 | ✓ 1468ms | 否 | ✓ 1437ms | ✓ 1596ms | 否 | http |
| 92.118.112.32:1081 | 否 | 否 | ✓ 1272ms | ✓ 1223ms | ✓ 1194ms | http |
| 92.118.112.25:1081 | ✓ 1163ms | 否 | ✓ 1868ms | ✓ 1996ms | 否 | http |
| 72.56.72.238:40001 | ✓ 476ms | ✓ 1965ms | 否 | 否 | ✓ 1108ms | http |
| 92.118.112.25:1082 | ✓ 1271ms | 否 | ✓ 708ms | ✓ 1031ms | ✓ 804ms | http |
| 89.22.230.26:1080 | ✓ 1112ms | 否 | ✓ 1692ms | 否 | ✓ 1916ms | http |
| 103.18.77.14:1111 | 否 | 否 | ✓ 1417ms | ✓ 1589ms | ✓ 1562ms | http |
| 42.200.76.16:3888 | 否 | 否 | ✓ 770ms | ✓ 1124ms | ✓ 1717ms | http |

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
