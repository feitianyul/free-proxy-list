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

最后更新：2026-04-19 00:34:32 UTC（2026-04-19 08:34:32 UTC+8）

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
| 149.51.42.10:3128 | ✓ 1076ms | ✓ 1319ms | 否 | ✓ 1253ms | 否 | http |
| 149.51.42.10:8080 | ✓ 1076ms | ✓ 1600ms | 否 | ✓ 1236ms | 否 | http |
| 1.231.81.166:3128 | ✓ 935ms | ✓ 1240ms | ✓ 1090ms | ✓ 1112ms | ✓ 894ms | http |
| 185.138.116.150:8080 | ✓ 1215ms | ✓ 1729ms | ✓ 1139ms | ✓ 1362ms | ✓ 1013ms | http |
| 218.108.131.186:17890 | ✓ 991ms | ✓ 1264ms | ✓ 980ms | ✓ 1492ms | ✓ 1053ms | http |
| 194.104.9.38:3128 | ✓ 970ms | ✓ 1620ms | ✓ 1704ms | ✓ 1821ms | ✓ 1358ms | http |
| 188.246.224.49:7890 | ✓ 993ms | ✓ 1548ms | ✓ 1595ms | ✓ 1530ms | ✓ 1676ms | http |
| 117.122.240.82:3338 | ✓ 1041ms | ✓ 1405ms | ✓ 1034ms | ✓ 1377ms | ✓ 1086ms | http |
| 46.101.95.183:8888 | ✓ 1962ms | 否 | ✓ 1397ms | ✓ 1506ms | ✓ 1564ms | http |
| 81.30.156.115:8080 | ✓ 971ms | ✓ 1299ms | ✓ 1289ms | 否 | ✓ 1297ms | http |
| 152.42.208.139:8118 | ✓ 1519ms | 否 | ✓ 1906ms | ✓ 1262ms | ✓ 1073ms | http |
| 14.247.76.52:8080 | ✓ 1714ms | 否 | ✓ 1113ms | ✓ 1424ms | ✓ 1210ms | http |
| 168.144.75.9:3128 | ✓ 1661ms | 否 | ✓ 1934ms | ✓ 1763ms | ✓ 1386ms | http |
| 78.11.96.22:8888 | ✓ 1093ms | ✓ 1418ms | ✓ 1062ms | ✓ 1412ms | ✓ 1250ms | http |
| 62.113.119.14:8080 | ✓ 559ms | ✓ 1467ms | ✓ 539ms | ✓ 1440ms | ✓ 1058ms | http |
| 91.99.15.45:2095 | ✓ 770ms | ✓ 1657ms | ✓ 1049ms | ✓ 1772ms | ✓ 1562ms | http |
| 77.110.113.24:40000 | ✓ 884ms | 否 | ✓ 562ms | 否 | ✓ 1433ms | http |
| 113.160.132.26:8080 | ✓ 1013ms | 否 | ✓ 1053ms | ✓ 1388ms | ✓ 1230ms | http |
| 84.47.150.125:1080 | ✓ 918ms | 否 | ✓ 1931ms | ✓ 1741ms | ✓ 1583ms | http |
| 117.236.124.166:3128 | ✓ 1158ms | 否 | ✓ 1272ms | ✓ 1935ms | 否 | http |
| 45.76.207.177:40000 | 否 | 否 | ✓ 1686ms | ✓ 1980ms | ✓ 1392ms | http |
| 59.46.216.131:30001 | ✓ 1155ms | ✓ 1532ms | ✓ 1309ms | ✓ 1569ms | 否 | http |
| 159.89.191.221:3128 | ✓ 1784ms | 否 | ✓ 473ms | ✓ 1119ms | ✓ 734ms | http |
| 149.56.24.51:3128 | ✓ 582ms | ✓ 1768ms | ✓ 1533ms | ✓ 1309ms | ✓ 1116ms | http |
| 42.200.76.16:3888 | ✓ 864ms | 否 | ✓ 879ms | ✓ 1085ms | ✓ 876ms | http |
| 38.180.192.119:3128 | 否 | 否 | ✓ 1918ms | ✓ 1653ms | ✓ 802ms | http |
| 84.47.150.126:1080 | ✓ 1634ms | ✓ 1623ms | ✓ 589ms | 否 | 否 | http |
| 85.190.99.143:443 | ✓ 973ms | 否 | ✓ 634ms | 否 | ✓ 1712ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 844ms | ✓ 1620ms | ✓ 1663ms | http |
| 35.225.22.61:80 | ✓ 349ms | 否 | 否 | ✓ 1575ms | ✓ 957ms | http |
| 192.3.248.190:8014 | ✓ 736ms | ✓ 1500ms | ✓ 384ms | ✓ 1309ms | ✓ 1068ms | http |
| 115.231.181.40:8128 | ✓ 1177ms | 否 | ✓ 1066ms | 否 | ✓ 1107ms | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 1081ms | ✓ 1544ms | ✓ 785ms | http |
| 45.153.231.229:8080 | ✓ 654ms | 否 | ✓ 1171ms | ✓ 1903ms | ✓ 1470ms | http |
| 82.148.18.242:443 | ✓ 782ms | ✓ 1969ms | ✓ 1910ms | 否 | 否 | http |
| 162.19.253.202:8443 | ✓ 956ms | ✓ 1779ms | ✓ 1217ms | 否 | ✓ 1768ms | http |
| 161.97.184.191:8080 | ✓ 1188ms | ✓ 1352ms | 否 | ✓ 1870ms | ✓ 1560ms | http |
| 152.32.132.190:7890 | ✓ 1177ms | 否 | 否 | ✓ 1613ms | ✓ 1794ms | http |
| 91.198.220.93:443 | ✓ 1075ms | 否 | ✓ 1311ms | 否 | ✓ 1743ms | http |
| 223.84.151.86:30005 | ✓ 1637ms | ✓ 1671ms | ✓ 1584ms | 否 | ✓ 1725ms | http |
| 194.67.127.23:10808 | ✓ 652ms | 否 | ✓ 907ms | 否 | ✓ 1458ms | http |
| 170.205.39.31:1080 | ✓ 1047ms | 否 | ✓ 844ms | ✓ 1012ms | ✓ 1907ms | http |
| 147.45.186.28:3128 | 否 | ✓ 1712ms | ✓ 602ms | ✓ 1337ms | ✓ 1131ms | http |
| 82.114.228.67:1080 | ✓ 1139ms | 否 | ✓ 1413ms | ✓ 1476ms | 否 | http |
| 144.31.27.49:1080 | 否 | ✓ 1919ms | ✓ 1105ms | 否 | ✓ 1756ms | http |
| 45.140.147.82:1081 | ✓ 1008ms | ✓ 1204ms | ✓ 1242ms | ✓ 1213ms | ✓ 1095ms | http |
| 45.140.147.82:1082 | ✓ 1002ms | ✓ 1195ms | ✓ 1247ms | ✓ 1214ms | ✓ 1097ms | http |
| 8.219.195.129:1080 | ✓ 1480ms | ✓ 1890ms | ✓ 1077ms | ✓ 1254ms | ✓ 1049ms | http |
| 185.114.73.2:1080 | ✓ 1644ms | ✓ 1541ms | ✓ 1971ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1138ms | ✓ 1492ms | ✓ 1198ms | ✓ 1475ms | ✓ 1170ms | http |
| 45.12.151.226:2829 | ✓ 994ms | 否 | ✓ 1337ms | 否 | ✓ 1226ms | http |
| 120.92.212.16:7890 | ✓ 1773ms | 否 | ✓ 1101ms | 否 | ✓ 1210ms | http |
| 210.223.44.230:3128 | ✓ 1958ms | 否 | ✓ 1026ms | 否 | ✓ 1962ms | http |
| 84.247.171.137:3128 | ✓ 428ms | 否 | ✓ 1135ms | ✓ 1829ms | ✓ 1337ms | http |
| 177.93.132.244:3128 | ✓ 660ms | 否 | ✓ 635ms | 否 | ✓ 1571ms | http |
| 94.131.118.39:1081 | ✓ 525ms | 否 | ✓ 1141ms | 否 | ✓ 1999ms | http |
| 43.132.188.134:443 | ✓ 1692ms | ✓ 1446ms | 否 | 否 | ✓ 1717ms | http |
| 213.32.85.26:3128 | ✓ 837ms | ✓ 1668ms | ✓ 504ms | 否 | ✓ 1475ms | http |
| 212.58.132.5:8888 | ✓ 1827ms | 否 | ✓ 1556ms | ✓ 1602ms | ✓ 1307ms | http |
| 147.45.60.34:1082 | ✓ 419ms | ✓ 1555ms | 否 | ✓ 1061ms | 否 | http |
| 121.230.8.22:1080 | ✓ 1476ms | ✓ 1644ms | ✓ 1180ms | ✓ 1696ms | ✓ 1150ms | http |
| 103.113.70.189:1081 | ✓ 663ms | ✓ 1251ms | ✓ 93ms | ✓ 1220ms | ✓ 779ms | http |
| 103.113.70.189:1082 | ✓ 662ms | 否 | ✓ 205ms | ✓ 1109ms | ✓ 744ms | http |
| 168.110.52.228:3128 | ✓ 740ms | 否 | 否 | ✓ 1145ms | ✓ 896ms | http |
| 101.32.243.189:80 | ✓ 1480ms | ✓ 1481ms | 否 | ✓ 1421ms | ✓ 1515ms | http |
| 91.193.240.157:9877 | ✓ 1742ms | 否 | ✓ 896ms | 否 | ✓ 1949ms | http |
| 187.216.141.46:3128 | ✓ 604ms | ✓ 1466ms | ✓ 925ms | ✓ 1433ms | 否 | http |
| 83.219.250.8:62920 | ✓ 1269ms | ✓ 1252ms | ✓ 664ms | 否 | 否 | http |
| 14.143.222.113:10155 | ✓ 1404ms | 否 | ✓ 1038ms | ✓ 1369ms | 否 | http |
| 144.124.227.88:3128 | ✓ 434ms | 否 | ✓ 894ms | ✓ 1976ms | ✓ 1900ms | http |
| 34.96.238.40:8080 | ✓ 1394ms | ✓ 1466ms | ✓ 1338ms | ✓ 1254ms | 否 | http |
| 166.88.61.54:8000 | ✓ 1095ms | ✓ 1692ms | ✓ 1030ms | ✓ 1004ms | ✓ 804ms | http |
| 91.107.124.215:3128 | ✓ 774ms | ✓ 1708ms | ✓ 840ms | ✓ 1866ms | 否 | http |
| 72.56.105.251:3128 | ✓ 1097ms | ✓ 1721ms | 否 | 否 | ✓ 1312ms | http |
| 220.197.44.36:3128 | 否 | ✓ 1781ms | ✓ 1793ms | 否 | ✓ 1774ms | http |
| 116.171.106.26:3443 | ✓ 1707ms | ✓ 1709ms | ✓ 1681ms | ✓ 1980ms | 否 | http |
| 103.138.70.165:3129 | ✓ 1545ms | 否 | ✓ 1827ms | 否 | ✓ 1831ms | http |
| 45.186.6.104:3128 | ✓ 1224ms | ✓ 1748ms | ✓ 1651ms | 否 | 否 | http |
| 178.63.155.151:9000 | ✓ 950ms | ✓ 1403ms | ✓ 1159ms | 否 | 否 | http |
| 202.141.161.53:10808 | ✓ 1089ms | ✓ 1605ms | ✓ 1282ms | ✓ 1499ms | ✓ 1225ms | http |
| 121.135.144.234:8879 | ✓ 1201ms | ✓ 1762ms | ✓ 1632ms | ✓ 1633ms | ✓ 1259ms | http |
| 37.187.109.70:10111 | ✓ 1123ms | 否 | ✓ 1901ms | 否 | ✓ 1785ms | http |
| 45.140.147.155:1082 | ✓ 855ms | ✓ 1037ms | ✓ 1003ms | ✓ 1314ms | 否 | http |
| 178.156.224.42:3128 | ✓ 930ms | ✓ 1665ms | ✓ 1589ms | 否 | ✓ 1729ms | http |
| 160.238.65.8:3128 | ✓ 413ms | 否 | ✓ 1454ms | ✓ 1556ms | ✓ 1267ms | http |
| 218.60.0.214:80 | ✓ 1799ms | 否 | ✓ 1648ms | ✓ 1727ms | ✓ 1286ms | http |
| 45.88.0.115:3128 | ✓ 470ms | ✓ 1292ms | ✓ 1439ms | ✓ 1911ms | ✓ 1447ms | http |
| 45.140.147.155:1081 | ✓ 408ms | 否 | ✓ 472ms | 否 | ✓ 1519ms | http |
| 61.52.131.172:8443 | ✓ 1079ms | ✓ 1346ms | ✓ 1055ms | ✓ 1340ms | ✓ 1116ms | http |
| 20.205.16.149:3128 | ✓ 1098ms | ✓ 1650ms | 否 | ✓ 1885ms | ✓ 1044ms | http |
| 20.127.128.70:8080 | ✓ 967ms | ✓ 1990ms | ✓ 1494ms | 否 | 否 | http |
| 45.88.0.117:3128 | ✓ 1207ms | ✓ 1247ms | ✓ 1561ms | ✓ 1663ms | ✓ 1486ms | http |
| 121.230.8.136:1080 | ✓ 1244ms | ✓ 1641ms | ✓ 1346ms | ✓ 1708ms | ✓ 1409ms | http |

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
