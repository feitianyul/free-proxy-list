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

最后更新：2026-03-03 21:42:58 UTC（2026-03-04 05:42:58 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 483ms | 否 | ✓ 1199ms | ✓ 1301ms | ✓ 1112ms | http |
| 125.128.12.144:3128 | ✓ 1367ms | ✓ 1312ms | ✓ 957ms | ✓ 1219ms | ✓ 960ms | http |
| 14.56.107.244:3128 | ✓ 1886ms | ✓ 1669ms | ✓ 1094ms | ✓ 1354ms | ✓ 984ms | http |
| 166.0.192.117:8888 | 否 | ✓ 1374ms | ✓ 354ms | 否 | ✓ 1216ms | http |
| 61.72.110.54:3128 | ✓ 774ms | ✓ 1885ms | ✓ 767ms | ✓ 1557ms | ✓ 966ms | http |
| 121.128.121.54:3128 | ✓ 747ms | 否 | ✓ 734ms | 否 | ✓ 950ms | http |
| 61.72.110.94:3128 | ✓ 817ms | 否 | ✓ 1065ms | ✓ 1237ms | ✓ 1026ms | http |
| 205.209.118.30:3138 | ✓ 222ms | ✓ 924ms | 否 | ✓ 1055ms | ✓ 797ms | http |
| 61.72.221.234:3128 | ✓ 838ms | 否 | 否 | ✓ 1506ms | ✓ 1444ms | http |
| 45.140.147.82:1081 | ✓ 396ms | ✓ 1315ms | 否 | 否 | ✓ 1027ms | http |
| 186.148.180.46:999 | ✓ 1161ms | ✓ 1733ms | ✓ 1293ms | ✓ 1615ms | ✓ 1788ms | http |
| 14.56.177.44:3128 | ✓ 1403ms | 否 | ✓ 742ms | ✓ 1200ms | ✓ 925ms | http |
| 125.128.12.14:3128 | ✓ 1410ms | 否 | ✓ 1024ms | ✓ 1205ms | ✓ 981ms | http |
| 150.107.140.238:3128 | ✓ 955ms | 否 | ✓ 1115ms | ✓ 1341ms | ✓ 1044ms | http |
| 23.224.193.43:3128 | ✓ 1877ms | ✓ 1571ms | ✓ 1352ms | ✓ 1527ms | ✓ 1430ms | http |
| 23.224.193.45:3128 | ✓ 1875ms | ✓ 1512ms | ✓ 1412ms | ✓ 1540ms | ✓ 1419ms | http |
| 103.139.138.194:3128 | ✓ 1333ms | 否 | ✓ 1728ms | 否 | ✓ 1629ms | http |
| 121.204.158.249:3128 | ✓ 1178ms | ✓ 1451ms | ✓ 1135ms | 否 | ✓ 1158ms | http |
| 74.208.234.198:443 | ✓ 1219ms | ✓ 1979ms | 否 | ✓ 1446ms | 否 | http |
| 103.4.76.105:8589 | ✓ 1580ms | 否 | 否 | ✓ 1861ms | ✓ 1877ms | http |
| 103.84.95.54:7890 | ✓ 877ms | 否 | ✓ 1071ms | 否 | ✓ 823ms | http |
| 59.46.216.131:30001 | ✓ 1132ms | 否 | ✓ 1236ms | ✓ 1609ms | ✓ 1481ms | http |
| 35.225.22.61:80 | ✓ 326ms | ✓ 1423ms | ✓ 834ms | 否 | ✓ 668ms | http |
| 2.56.178.131:443 | ✓ 1024ms | 否 | 否 | ✓ 1798ms | ✓ 1729ms | http |
| 165.227.5.10:8888 | ✓ 381ms | ✓ 1666ms | ✓ 790ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1963ms | 否 | ✓ 1368ms | 否 | ✓ 1132ms | http |
| 120.92.212.16:7890 | ✓ 1107ms | ✓ 1664ms | ✓ 1170ms | 否 | ✓ 1156ms | http |
| 158.160.215.167:8127 | ✓ 1246ms | 否 | ✓ 1424ms | 否 | ✓ 1246ms | http |
| 61.72.221.94:3128 | 否 | ✓ 1659ms | ✓ 1595ms | ✓ 1790ms | 否 | http |
| 101.43.255.96:80 | ✓ 1187ms | ✓ 1460ms | ✓ 1096ms | ✓ 1542ms | ✓ 1166ms | http |
| 81.70.169.194:80 | ✓ 1131ms | ✓ 1548ms | ✓ 1218ms | ✓ 1452ms | ✓ 1174ms | http |
| 90.84.188.97:8000 | ✓ 946ms | 否 | 否 | ✓ 1883ms | ✓ 1608ms | http |
| 147.45.251.242:8888 | ✓ 1227ms | 否 | ✓ 1658ms | 否 | ✓ 1994ms | http |
| 159.89.31.62:8080 | ✓ 1043ms | ✓ 1676ms | 否 | ✓ 1596ms | ✓ 1176ms | http |
| 91.233.223.147:3128 | ✓ 1075ms | 否 | ✓ 877ms | ✓ 1840ms | ✓ 1450ms | http |
| 91.193.240.157:9877 | ✓ 1080ms | 否 | ✓ 1467ms | 否 | ✓ 1785ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1218ms | ✓ 1341ms | ✓ 1057ms | http |
| 62.113.119.14:8080 | ✓ 620ms | 否 | ✓ 1421ms | 否 | ✓ 1457ms | http |
| 47.105.98.23:3128 | ✓ 1981ms | 否 | ✓ 1152ms | ✓ 1428ms | ✓ 1107ms | http |
| 120.79.99.232:8099 | ✓ 1449ms | ✓ 1752ms | ✓ 1470ms | ✓ 1672ms | ✓ 1402ms | http |
| 77.83.203.5:443 | ✓ 647ms | 否 | 否 | ✓ 1420ms | ✓ 1824ms | http |
| 185.115.74.185:8080 | ✓ 775ms | ✓ 1633ms | ✓ 1690ms | 否 | 否 | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 938ms | ✓ 1983ms | ✓ 1898ms | http |
| 61.72.221.194:3128 | ✓ 1377ms | ✓ 1143ms | 否 | ✓ 1538ms | 否 | http |
| 192.166.82.55:1080 | ✓ 461ms | 否 | ✓ 935ms | 否 | ✓ 1382ms | http |
| 34.101.184.164:3128 | ✓ 923ms | 否 | ✓ 1174ms | ✓ 1862ms | ✓ 1201ms | http |
| 46.249.103.192:443 | ✓ 1567ms | 否 | ✓ 1587ms | ✓ 1806ms | 否 | http |
| 158.160.215.167:8125 | ✓ 1000ms | ✓ 1853ms | 否 | 否 | ✓ 1739ms | http |
| 5.75.196.26:40000 | ✓ 970ms | 否 | ✓ 1778ms | 否 | ✓ 1626ms | http |
| 5.101.0.233:3128 | ✓ 1586ms | 否 | ✓ 1392ms | ✓ 1924ms | ✓ 1427ms | http |
| 118.31.1.154:80 | 否 | ✓ 1685ms | ✓ 1806ms | ✓ 1389ms | ✓ 1043ms | http |
| 24.199.124.151:3128 | ✓ 545ms | ✓ 1861ms | ✓ 1189ms | ✓ 1000ms | ✓ 767ms | http |
| 47.243.241.13:3128 | ✓ 810ms | ✓ 1631ms | 否 | 否 | ✓ 1640ms | http |
| 124.121.2.241:8080 | 否 | 否 | ✓ 1938ms | ✓ 1774ms | ✓ 1812ms | http |
| 121.230.8.74:1080 | ✓ 1355ms | ✓ 1733ms | ✓ 1438ms | 否 | 否 | http |
| 121.230.8.110:1080 | ✓ 1381ms | ✓ 1428ms | ✓ 1126ms | ✓ 1737ms | ✓ 1181ms | http |
| 8.217.147.173:8080 | ✓ 1402ms | ✓ 1343ms | ✓ 1263ms | 否 | 否 | http |
| 171.234.62.116:10001 | 否 | 否 | ✓ 1971ms | ✓ 1922ms | ✓ 1728ms | http |
| 209.38.51.97:3128 | 否 | ✓ 1090ms | ✓ 1166ms | ✓ 1156ms | ✓ 893ms | http |
| 163.5.128.210:14270 | ✓ 1251ms | ✓ 1068ms | ✓ 981ms | 否 | 否 | http |
| 117.86.6.35:1080 | ✓ 1061ms | ✓ 1337ms | ✓ 1077ms | ✓ 1391ms | 否 | http |
| 106.14.203.63:3333 | ✓ 996ms | ✓ 1268ms | ✓ 1042ms | 否 | 否 | http |
| 121.230.8.114:1080 | ✓ 1305ms | 否 | 否 | ✓ 1712ms | ✓ 1300ms | http |
| 121.230.9.251:1080 | ✓ 1174ms | 否 | ✓ 1367ms | ✓ 1646ms | ✓ 1505ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1234ms | 否 | ✓ 1065ms | ✓ 802ms | http |
| 45.136.198.40:3128 | ✓ 1013ms | ✓ 1689ms | ✓ 1694ms | ✓ 1799ms | ✓ 1760ms | http |
| 121.230.9.205:1080 | ✓ 1448ms | ✓ 1756ms | ✓ 1147ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 1532ms | 否 | 否 | ✓ 1785ms | ✓ 1803ms | http |
| 172.212.68.37:3128 | ✓ 285ms | ✓ 1303ms | ✓ 710ms | ✓ 968ms | ✓ 850ms | http |
| 121.230.9.168:1080 | ✓ 1657ms | ✓ 1906ms | ✓ 1359ms | 否 | 否 | http |
| 121.230.9.179:1080 | ✓ 1171ms | ✓ 1687ms | ✓ 1319ms | ✓ 1649ms | ✓ 1301ms | http |
| 147.45.60.34:1082 | ✓ 407ms | 否 | ✓ 921ms | ✓ 1121ms | ✓ 899ms | http |
| 45.129.141.143:3128 | ✓ 1371ms | 否 | ✓ 1536ms | 否 | ✓ 1598ms | http |
| 38.180.2.107:3128 | ✓ 1370ms | ✓ 1876ms | ✓ 1574ms | 否 | ✓ 1855ms | http |
| 94.176.3.43:7443 | ✓ 626ms | 否 | ✓ 1088ms | ✓ 1750ms | ✓ 1222ms | http |
| 188.166.208.168:9876 | ✓ 1329ms | 否 | ✓ 931ms | ✓ 1259ms | ✓ 1092ms | http |
| 101.47.73.135:3128 | ✓ 1345ms | 否 | ✓ 1982ms | ✓ 1493ms | 否 | http |
| 35.234.17.221:8080 | 否 | ✓ 1844ms | 否 | ✓ 1159ms | ✓ 1053ms | http |
| 182.253.160.168:1452 | ✓ 1949ms | 否 | 否 | ✓ 1637ms | ✓ 1604ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1523ms | ✓ 1935ms | ✓ 1480ms | http |
| 144.31.69.170:1080 | ✓ 1104ms | 否 | ✓ 1327ms | 否 | ✓ 1182ms | http |
| 37.27.100.107:443 | ✓ 1739ms | 否 | ✓ 1093ms | ✓ 1474ms | 否 | http |
| 45.22.209.157:8888 | ✓ 682ms | 否 | ✓ 299ms | 否 | ✓ 1947ms | http |
| 47.74.226.8:5001 | ✓ 1507ms | 否 | ✓ 1284ms | ✓ 1490ms | 否 | http |
| 77.83.203.6:443 | ✓ 1043ms | ✓ 1751ms | ✓ 1249ms | 否 | ✓ 1857ms | http |
| 122.248.45.54:8080 | ✓ 1892ms | 否 | ✓ 1941ms | ✓ 1974ms | 否 | http |
| 121.230.9.132:1080 | ✓ 1613ms | ✓ 1679ms | ✓ 1284ms | ✓ 1828ms | ✓ 1219ms | http |
| 37.27.100.108:443 | 否 | ✓ 1643ms | ✓ 1803ms | 否 | ✓ 1288ms | http |
| 121.230.8.11:1080 | ✓ 1386ms | ✓ 1751ms | ✓ 1327ms | ✓ 1724ms | ✓ 1260ms | http |
| 47.101.149.27:9010 | ✓ 1593ms | 否 | 否 | ✓ 1793ms | ✓ 1980ms | http |
| 152.70.137.18:8888 | 否 | ✓ 1000ms | ✓ 435ms | ✓ 1267ms | 否 | http |
| 61.76.95.217:40088 | ✓ 1295ms | ✓ 1453ms | ✓ 1246ms | ✓ 1709ms | 否 | http |

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
