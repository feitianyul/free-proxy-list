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

最后更新：2026-04-01 06:19:24 UTC（2026-04-01 14:19:24 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 487ms | ✓ 865ms | 否 | ✓ 1193ms | ✓ 779ms | http |
| 43.99.54.236:5555 | ✓ 754ms | ✓ 1359ms | ✓ 748ms | ✓ 919ms | ✓ 744ms | http |
| 147.161.210.140:8800 | ✓ 783ms | 否 | ✓ 790ms | ✓ 1198ms | ✓ 981ms | http |
| 147.161.239.240:8800 | ✓ 610ms | ✓ 1690ms | ✓ 925ms | ✓ 1680ms | ✓ 1433ms | http |
| 1.231.81.166:3128 | ✓ 1096ms | ✓ 1465ms | 否 | ✓ 1153ms | ✓ 843ms | http |
| 133.242.138.34:8100 | ✓ 995ms | ✓ 1646ms | ✓ 911ms | ✓ 1363ms | 否 | http |
| 13.214.186.22:3128 | ✓ 1456ms | ✓ 1589ms | ✓ 1093ms | ✓ 1366ms | ✓ 921ms | http |
| 82.114.228.67:1080 | ✓ 743ms | ✓ 1823ms | ✓ 983ms | 否 | ✓ 1493ms | http |
| 167.103.115.102:8800 | ✓ 1454ms | 否 | ✓ 1155ms | 否 | ✓ 1148ms | http |
| 95.213.217.168:52004 | ✓ 725ms | ✓ 1886ms | ✓ 1834ms | 否 | ✓ 1941ms | http |
| 113.160.132.26:8080 | ✓ 1882ms | ✓ 1505ms | ✓ 1344ms | ✓ 1412ms | ✓ 1288ms | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1780ms | ✓ 1699ms | ✓ 1200ms | http |
| 167.103.34.108:8800 | ✓ 1661ms | 否 | 否 | ✓ 1609ms | ✓ 1771ms | http |
| 35.225.22.61:80 | ✓ 485ms | 否 | ✓ 1401ms | ✓ 1029ms | ✓ 1062ms | http |
| 167.160.191.204:6005 | ✓ 816ms | 否 | ✓ 1154ms | ✓ 1062ms | ✓ 1894ms | http |
| 45.12.151.226:2829 | ✓ 837ms | 否 | ✓ 1208ms | 否 | ✓ 1344ms | http |
| 167.103.144.127:8800 | ✓ 1518ms | ✓ 1928ms | ✓ 1134ms | 否 | ✓ 1658ms | http |
| 167.103.31.122:8800 | ✓ 1833ms | 否 | ✓ 1726ms | 否 | ✓ 1714ms | http |
| 45.167.124.52:8080 | ✓ 856ms | ✓ 1536ms | ✓ 618ms | 否 | 否 | http |
| 165.232.146.249:3128 | ✓ 648ms | ✓ 1392ms | ✓ 1365ms | ✓ 969ms | ✓ 663ms | http |
| 185.191.236.162:3128 | ✓ 695ms | ✓ 1641ms | ✓ 1538ms | 否 | ✓ 1156ms | http |
| 31.192.106.135:8010 | ✓ 1471ms | 否 | ✓ 1126ms | 否 | ✓ 1627ms | http |
| 195.19.217.200:3128 | ✓ 1535ms | 否 | ✓ 1284ms | 否 | ✓ 1639ms | http |
| 180.250.219.58:53281 | ✓ 1573ms | 否 | ✓ 1561ms | 否 | ✓ 1966ms | http |
| 116.80.63.64:7777 | ✓ 1548ms | 否 | ✓ 1589ms | 否 | ✓ 1760ms | http |
| 160.250.5.22:1 | ✓ 1451ms | 否 | ✓ 1351ms | ✓ 1379ms | ✓ 1092ms | http |
| 177.234.217.88:999 | ✓ 1659ms | ✓ 1904ms | ✓ 1934ms | 否 | ✓ 1609ms | http |
| 181.78.44.63:999 | ✓ 929ms | ✓ 1564ms | ✓ 1289ms | 否 | ✓ 1194ms | http |
| 59.46.216.131:30001 | ✓ 1069ms | ✓ 1448ms | 否 | ✓ 1433ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1120ms | ✓ 1803ms | 否 | ✓ 1527ms | ✓ 1663ms | http |
| 150.241.71.15:1080 | ✓ 777ms | 否 | ✓ 1927ms | ✓ 1772ms | ✓ 1134ms | http |
| 120.92.212.16:7890 | ✓ 1887ms | ✓ 1568ms | 否 | 否 | ✓ 1080ms | http |
| 129.213.162.27:17777 | 否 | ✓ 1789ms | ✓ 1849ms | 否 | ✓ 1536ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1641ms | ✓ 220ms | 否 | ✓ 929ms | http |
| 1.12.62.237:8080 | 否 | ✓ 1858ms | ✓ 1603ms | ✓ 1758ms | ✓ 1887ms | http |
| 3.99.169.21:48021 | ✓ 1406ms | 否 | ✓ 1733ms | 否 | ✓ 1617ms | http |
| 91.233.223.147:3128 | ✓ 1171ms | 否 | ✓ 794ms | 否 | ✓ 1880ms | http |
| 47.74.226.8:5001 | ✓ 1088ms | 否 | 否 | ✓ 1319ms | ✓ 1339ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1535ms | ✓ 1541ms | ✓ 1088ms | http |
| 72.11.150.178:6005 | ✓ 364ms | 否 | ✓ 874ms | ✓ 1264ms | ✓ 1039ms | http |
| 203.80.138.81:50000 | ✓ 925ms | ✓ 1212ms | ✓ 1031ms | ✓ 1072ms | ✓ 920ms | http |
| 38.34.178.152:8444 | ✓ 1343ms | ✓ 1794ms | ✓ 1849ms | ✓ 969ms | ✓ 811ms | http |
| 8.222.175.80:6128 | ✓ 847ms | 否 | ✓ 810ms | ✓ 1146ms | ✓ 927ms | http |
| 106.10.55.212:1121 | 否 | ✓ 1731ms | ✓ 1367ms | ✓ 1360ms | ✓ 1141ms | http |
| 92.119.127.212:6005 | ✓ 534ms | 否 | ✓ 1637ms | ✓ 1928ms | 否 | http |
| 180.103.19.124:1080 | ✓ 1320ms | ✓ 1805ms | 否 | ✓ 1561ms | ✓ 1261ms | http |
| 217.217.249.160:8080 | ✓ 1331ms | 否 | ✓ 1171ms | 否 | ✓ 1373ms | http |
| 183.98.143.134:8042 | ✓ 834ms | ✓ 1404ms | ✓ 1042ms | ✓ 1111ms | ✓ 841ms | http |
| 86.53.183.16:1080 | ✓ 1460ms | 否 | ✓ 1399ms | ✓ 1833ms | ✓ 1623ms | http |
| 178.156.224.42:3128 | ✓ 890ms | 否 | ✓ 1404ms | 否 | ✓ 1562ms | http |
| 146.190.80.158:9090 | ✓ 1284ms | 否 | ✓ 1455ms | ✓ 1333ms | ✓ 953ms | http |
| 128.199.121.61:9090 | ✓ 1283ms | 否 | ✓ 1216ms | ✓ 1259ms | ✓ 1257ms | http |
| 128.199.254.13:9090 | ✓ 1290ms | 否 | 否 | ✓ 1813ms | ✓ 1223ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 970ms | ✓ 1219ms | ✓ 1227ms | http |
| 104.248.151.93:9090 | ✓ 1295ms | 否 | ✓ 896ms | ✓ 1175ms | ✓ 914ms | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1120ms | ✓ 1202ms | ✓ 1582ms | http |
| 103.125.118.172:3125 | 否 | 否 | ✓ 1454ms | ✓ 1693ms | ✓ 1612ms | http |
| 103.217.224.75:3125 | 否 | 否 | ✓ 1838ms | ✓ 1510ms | ✓ 1488ms | http |
| 128.199.113.85:9090 | ✓ 1300ms | 否 | ✓ 830ms | ✓ 1138ms | 否 | http |
| 128.199.114.189:9090 | ✓ 1278ms | 否 | ✓ 829ms | ✓ 1196ms | ✓ 963ms | http |
| 128.199.116.219:9090 | ✓ 1303ms | 否 | ✓ 1203ms | ✓ 1157ms | ✓ 932ms | http |
| 45.140.147.155:1081 | 否 | ✓ 1350ms | ✓ 1498ms | ✓ 1679ms | 否 | http |
| 203.175.102.162:1111 | 否 | 否 | ✓ 1633ms | ✓ 1551ms | ✓ 1753ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 980ms | ✓ 1687ms | ✓ 1055ms | http |
| 168.222.254.136:8888 | 否 | ✓ 1863ms | ✓ 1321ms | 否 | ✓ 1708ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 650ms | ✓ 986ms | ✓ 767ms | http |
| 209.126.84.232:8888 | ✓ 1712ms | 否 | 否 | ✓ 1670ms | ✓ 1450ms | http |
| 62.113.119.14:8080 | ✓ 1510ms | 否 | ✓ 1122ms | ✓ 1785ms | ✓ 1849ms | http |
| 148.153.56.51:80 | 否 | 否 | ✓ 953ms | ✓ 1000ms | ✓ 833ms | http |
| 43.225.185.4:8000 | 否 | 否 | ✓ 1482ms | ✓ 1415ms | ✓ 1389ms | http |
| 103.191.92.157:1009 | ✓ 1663ms | 否 | ✓ 1592ms | ✓ 1359ms | ✓ 1416ms | http |
| 45.136.198.40:3128 | ✓ 726ms | ✓ 1502ms | ✓ 1377ms | ✓ 1977ms | ✓ 1473ms | http |
| 45.140.147.155:1082 | ✓ 1883ms | ✓ 1422ms | 否 | ✓ 1691ms | 否 | http |
| 8.219.97.248:80 | ✓ 1713ms | 否 | ✓ 1683ms | ✓ 1732ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1839ms | 否 | 否 | ✓ 1759ms | ✓ 1524ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1217ms | 否 | ✓ 1379ms | ✓ 952ms | http |
| 160.250.4.245:1 | ✓ 1920ms | 否 | ✓ 1764ms | ✓ 1376ms | ✓ 1369ms | http |

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
