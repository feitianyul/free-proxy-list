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

最后更新：2026-06-07 21:03:19 UTC（2026-06-08 05:03:19 UTC+8）

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
| 34.43.46.91:80 | ✓ 708ms | ✓ 1342ms | ✓ 483ms | ✓ 1004ms | ✓ 864ms | http |
| 94.241.175.40:10808 | ✓ 1144ms | ✓ 1392ms | ✓ 1695ms | ✓ 1125ms | 否 | http |
| 81.200.154.236:48503 | ✓ 1117ms | 否 | 否 | ✓ 1630ms | ✓ 1330ms | http |
| 1.231.81.166:3128 | ✓ 1372ms | ✓ 1482ms | ✓ 1638ms | ✓ 1427ms | ✓ 1307ms | http |
| 37.49.224.15:3128 | ✓ 1109ms | 否 | ✓ 1396ms | ✓ 1936ms | ✓ 1883ms | http |
| 185.200.188.234:10001 | ✓ 1149ms | 否 | ✓ 1254ms | 否 | ✓ 1598ms | http |
| 104.161.37.187:3128 | ✓ 376ms | ✓ 936ms | ✓ 372ms | ✓ 1021ms | ✓ 1068ms | http |
| 104.154.186.48:80 | ✓ 314ms | ✓ 1175ms | ✓ 837ms | ✓ 1225ms | ✓ 945ms | http |
| 50.114.102.16:8888 | 否 | ✓ 1611ms | ✓ 1488ms | ✓ 1766ms | ✓ 1107ms | http |
| 176.111.37.216:39811 | ✓ 513ms | ✓ 1400ms | ✓ 824ms | ✓ 1634ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1562ms | ✓ 1423ms | ✓ 1020ms | ✓ 1329ms | ✓ 1049ms | http |
| 209.38.200.247:1080 | ✓ 484ms | 否 | 否 | ✓ 1955ms | ✓ 1319ms | http |
| 117.55.203.162:8899 | ✓ 1576ms | 否 | ✓ 1707ms | 否 | ✓ 1786ms | http |
| 185.233.186.88:443 | ✓ 1575ms | 否 | ✓ 1098ms | 否 | ✓ 1632ms | http |
| 84.47.150.125:1080 | ✓ 1626ms | 否 | ✓ 1925ms | 否 | ✓ 1524ms | http |
| 62.210.136.222:3128 | ✓ 575ms | 否 | ✓ 1932ms | ✓ 1757ms | ✓ 1153ms | http |
| 185.141.26.131:3128 | ✓ 1285ms | 否 | ✓ 1760ms | 否 | ✓ 1460ms | http |
| 95.3.69.222:8080 | ✓ 1746ms | ✓ 1871ms | 否 | ✓ 1709ms | ✓ 1485ms | http |
| 116.104.252.1:2105 | ✓ 1596ms | 否 | ✓ 1588ms | ✓ 1811ms | ✓ 1846ms | http |
| 129.153.7.7:60000 | 否 | ✓ 866ms | ✓ 182ms | 否 | ✓ 689ms | http |
| 169.212.15.161:5000 | ✓ 1630ms | ✓ 1839ms | ✓ 760ms | ✓ 1911ms | 否 | http |
| 188.225.58.59:443 | ✓ 1025ms | 否 | ✓ 1134ms | 否 | ✓ 1617ms | http |
| 165.227.133.230:8888 | ✓ 961ms | ✓ 1817ms | ✓ 380ms | ✓ 1536ms | 否 | http |
| 43.128.145.26:1080 | ✓ 1692ms | 否 | 否 | ✓ 1377ms | ✓ 1909ms | http |
| 85.234.100.149:8080 | 否 | 否 | ✓ 1378ms | ✓ 1765ms | ✓ 1170ms | http |
| 59.66.245.22:6382 | ✓ 1290ms | ✓ 1725ms | ✓ 1346ms | ✓ 1790ms | ✓ 1175ms | http |
| 116.104.252.1:2070 | ✓ 1586ms | 否 | ✓ 1574ms | ✓ 1831ms | 否 | http |
| 170.106.136.181:31002 | ✓ 1061ms | 否 | ✓ 956ms | ✓ 863ms | ✓ 655ms | http |
| 8.154.21.175:3128 | ✓ 1068ms | ✓ 1298ms | ✓ 1307ms | ✓ 1356ms | ✓ 1060ms | http |
| 45.88.174.195:8080 | ✓ 385ms | ✓ 1402ms | ✓ 984ms | ✓ 1884ms | ✓ 1403ms | http |
| 172.93.101.178:3128 | 否 | 否 | ✓ 798ms | ✓ 1360ms | ✓ 1232ms | http |
| 202.28.194.139:31280 | ✓ 1680ms | 否 | ✓ 1952ms | 否 | ✓ 1931ms | http |
| 195.25.20.155:3128 | ✓ 1007ms | 否 | ✓ 1712ms | 否 | ✓ 1170ms | http |
| 162.222.206.167:8080 | ✓ 1236ms | 否 | ✓ 1570ms | ✓ 1657ms | ✓ 1357ms | http |
| 146.56.133.63:1080 | 否 | 否 | ✓ 882ms | ✓ 1208ms | ✓ 1002ms | http |
| 216.9.225.157:3128 | ✓ 1716ms | 否 | ✓ 952ms | 否 | ✓ 1333ms | http |
| 117.55.203.161:8899 | ✓ 1392ms | 否 | ✓ 1770ms | 否 | ✓ 1263ms | http |
| 113.176.92.71:3128 | ✓ 1603ms | ✓ 1568ms | ✓ 1445ms | ✓ 1404ms | ✓ 1118ms | http |
| 46.8.112.212:3128 | 否 | ✓ 1980ms | 否 | ✓ 1744ms | ✓ 1493ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1944ms | ✓ 1955ms | ✓ 1212ms | http |
| 116.104.252.1:2030 | ✓ 1649ms | 否 | ✓ 1646ms | ✓ 1898ms | ✓ 1625ms | http |
| 45.186.6.104:3128 | ✓ 1517ms | ✓ 1761ms | ✓ 1850ms | 否 | 否 | http |
| 157.245.143.65:7890 | ✓ 529ms | ✓ 1188ms | ✓ 1266ms | ✓ 983ms | ✓ 661ms | http |
| 92.118.112.25:1082 | 否 | ✓ 1852ms | ✓ 951ms | ✓ 1611ms | ✓ 774ms | http |
| 43.160.236.170:8888 | ✓ 1710ms | ✓ 1786ms | ✓ 1735ms | ✓ 1381ms | ✓ 1413ms | http |
| 34.43.46.91:443 | 否 | ✓ 1006ms | ✓ 220ms | ✓ 1010ms | 否 | http |
| 92.118.112.32:1082 | ✓ 452ms | ✓ 1364ms | ✓ 190ms | ✓ 1416ms | ✓ 726ms | http |
| 2.26.87.216:1080 | ✓ 984ms | 否 | ✓ 977ms | 否 | ✓ 1514ms | http |
| 183.80.40.243:2053 | 否 | 否 | ✓ 1708ms | ✓ 1960ms | ✓ 1851ms | http |
| 137.59.47.73:3128 | ✓ 1945ms | ✓ 1701ms | 否 | ✓ 1303ms | 否 | http |
| 147.45.78.89:1080 | ✓ 807ms | 否 | ✓ 1589ms | ✓ 1465ms | 否 | http |
| 92.118.112.25:1081 | ✓ 847ms | ✓ 1630ms | 否 | 否 | ✓ 1456ms | http |
| 14.143.222.113:10158 | ✓ 1362ms | 否 | ✓ 1822ms | ✓ 1660ms | 否 | http |
| 205.215.247.164:3128 | ✓ 1629ms | ✓ 1065ms | ✓ 1290ms | 否 | 否 | http |
| 103.209.36.58:8080 | ✓ 1567ms | 否 | ✓ 1861ms | 否 | ✓ 1730ms | http |
| 43.228.215.32:8080 | 否 | 否 | ✓ 1991ms | ✓ 1826ms | ✓ 1589ms | http |
| 92.118.112.32:1081 | 否 | ✓ 1284ms | ✓ 587ms | ✓ 1922ms | ✓ 870ms | http |
| 18.180.59.181:80 | ✓ 746ms | ✓ 1181ms | ✓ 1494ms | ✓ 1491ms | ✓ 956ms | http |
| 80.150.246.98:443 | ✓ 657ms | ✓ 1306ms | ✓ 1754ms | 否 | ✓ 1550ms | http |
| 43.134.141.85:80 | ✓ 1417ms | ✓ 1751ms | ✓ 1286ms | 否 | ✓ 1535ms | http |
| 101.32.243.189:80 | ✓ 1428ms | ✓ 1678ms | ✓ 1754ms | 否 | ✓ 1544ms | http |
| es-xh-01.hpdata.click:443 | ✓ 911ms | 否 | ✓ 1398ms | 否 | ✓ 1522ms | http |
| 110.172.62.196:8080 | ✓ 1181ms | ✓ 1182ms | 否 | ✓ 1230ms | ✓ 1263ms | http |
| 168.110.52.228:3128 | ✓ 1937ms | ✓ 1869ms | ✓ 1347ms | 否 | ✓ 1050ms | http |
| 158.160.211.3:1080 | ✓ 525ms | ✓ 1612ms | ✓ 1100ms | 否 | ✓ 1738ms | http |
| 31.40.204.188:443 | ✓ 1509ms | ✓ 1960ms | 否 | ✓ 1952ms | ✓ 1897ms | http |
| 152.32.132.190:7890 | ✓ 1060ms | ✓ 1900ms | 否 | ✓ 1694ms | 否 | http |
| 112.86.205.251:7890 | ✓ 1022ms | ✓ 1323ms | ✓ 1063ms | ✓ 1367ms | ✓ 1083ms | http |
| 103.82.23.118:5249 | ✓ 1315ms | ✓ 1897ms | ✓ 1499ms | 否 | 否 | http |
| 38.191.200.174:999 | 否 | ✓ 1553ms | ✓ 1402ms | ✓ 1809ms | ✓ 1629ms | http |
| 51.79.71.106:8080 | ✓ 739ms | 否 | ✓ 1986ms | ✓ 1813ms | ✓ 1074ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1404ms | ✓ 1464ms | ✓ 1507ms | http |
| 45.189.118.192:999 | ✓ 1441ms | 否 | ✓ 1484ms | 否 | ✓ 1946ms | http |
| 169.40.6.114:3128 | ✓ 1195ms | 否 | ✓ 1128ms | 否 | ✓ 1512ms | http |
| 136.0.3.35:1234 | ✓ 415ms | 否 | ✓ 402ms | ✓ 1488ms | ✓ 694ms | http |
| 85.234.100.149:1080 | ✓ 1499ms | 否 | ✓ 939ms | ✓ 1348ms | ✓ 996ms | http |
| 43.161.239.147:8888 | ✓ 830ms | ✓ 1656ms | 否 | ✓ 1893ms | ✓ 1457ms | http |
| 61.52.131.172:8443 | ✓ 1066ms | 否 | ✓ 1089ms | ✓ 1916ms | ✓ 1132ms | http |
| 218.108.131.186:17890 | ✓ 1474ms | ✓ 1283ms | ✓ 1209ms | ✓ 1408ms | 否 | http |
| 103.39.51.207:8080 | ✓ 1726ms | 否 | 否 | ✓ 1691ms | ✓ 1672ms | http |

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
