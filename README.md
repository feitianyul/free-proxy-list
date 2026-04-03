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

最后更新：2026-04-03 21:36:16 UTC（2026-04-04 05:36:16 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 518ms | ✓ 851ms | ✓ 925ms | ✓ 864ms | ✓ 672ms | http |
| 218.108.131.186:17890 | ✓ 946ms | ✓ 1208ms | ✓ 1014ms | ✓ 1255ms | ✓ 989ms | http |
| 1.231.81.166:3128 | ✓ 885ms | ✓ 1093ms | ✓ 1711ms | ✓ 1421ms | ✓ 1057ms | http |
| 147.161.210.140:8800 | ✓ 1059ms | ✓ 1734ms | ✓ 1216ms | ✓ 1551ms | ✓ 1241ms | http |
| 111.227.254.9:22222 | ✓ 1115ms | ✓ 1507ms | ✓ 1228ms | ✓ 1495ms | ✓ 1142ms | http |
| 111.227.254.12:22222 | ✓ 1129ms | ✓ 1514ms | ✓ 1218ms | ✓ 1417ms | ✓ 1233ms | http |
| 113.160.132.26:8080 | ✓ 1801ms | ✓ 1485ms | ✓ 1200ms | ✓ 1361ms | ✓ 1260ms | http |
| 95.213.217.168:52004 | ✓ 1290ms | ✓ 1615ms | ✓ 1903ms | ✓ 1895ms | ✓ 1395ms | http |
| 167.103.115.102:8800 | ✓ 1529ms | ✓ 1846ms | ✓ 1693ms | ✓ 1440ms | ✓ 1185ms | http |
| 167.103.34.108:8800 | ✓ 1897ms | 否 | ✓ 1630ms | ✓ 1792ms | ✓ 1818ms | http |
| 45.167.124.52:8080 | ✓ 670ms | 否 | ✓ 1159ms | ✓ 1596ms | ✓ 1358ms | http |
| 35.225.22.61:80 | ✓ 188ms | 否 | 否 | ✓ 958ms | ✓ 979ms | http |
| 159.223.71.162:8080 | ✓ 829ms | 否 | 否 | ✓ 1574ms | ✓ 982ms | http |
| 5.104.87.17:8051 | ✓ 1236ms | 否 | ✓ 1657ms | ✓ 1072ms | ✓ 842ms | http |
| 167.103.144.127:8800 | ✓ 1754ms | 否 | ✓ 1597ms | ✓ 1772ms | ✓ 1643ms | http |
| 167.103.31.122:8800 | ✓ 1925ms | 否 | 否 | ✓ 1889ms | ✓ 1523ms | http |
| 120.92.212.16:7890 | ✓ 1461ms | ✓ 1330ms | 否 | 否 | ✓ 1136ms | http |
| 194.67.99.223:1080 | ✓ 1143ms | ✓ 1835ms | 否 | ✓ 1821ms | ✓ 1365ms | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1209ms | ✓ 1291ms | ✓ 1067ms | http |
| 59.46.216.131:30001 | ✓ 1127ms | ✓ 1396ms | 否 | 否 | ✓ 1182ms | http |
| 72.11.151.159:6005 | ✓ 638ms | ✓ 1166ms | ✓ 861ms | 否 | 否 | http |
| 147.161.239.240:8800 | ✓ 1129ms | ✓ 1559ms | ✓ 1267ms | ✓ 1608ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1502ms | ✓ 1166ms | ✓ 945ms | ✓ 1316ms | ✓ 1122ms | http |
| 18.201.114.187:50537 | ✓ 1173ms | 否 | ✓ 1760ms | ✓ 1990ms | 否 | http |
| 47.238.220.4:8888 | ✓ 1002ms | ✓ 1375ms | 否 | ✓ 1175ms | ✓ 1625ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1138ms | ✓ 1343ms | ✓ 1131ms | http |
| 133.26.134.100:3128 | ✓ 690ms | ✓ 1165ms | ✓ 820ms | ✓ 1126ms | ✓ 975ms | http |
| 20.2.83.243:3128 | ✓ 903ms | ✓ 1816ms | ✓ 992ms | ✓ 1332ms | ✓ 906ms | http |
| 45.149.92.147:5001 | ✓ 930ms | 否 | ✓ 935ms | ✓ 1289ms | ✓ 1018ms | http |
| 115.231.181.40:8128 | ✓ 972ms | ✓ 1213ms | ✓ 1133ms | ✓ 1308ms | 否 | http |
| 217.160.162.25:8888 | ✓ 1253ms | ✓ 1755ms | ✓ 1319ms | 否 | 否 | http |
| 195.123.209.48:3128 | ✓ 1277ms | ✓ 1667ms | ✓ 1397ms | 否 | ✓ 1753ms | http |
| 177.93.33.55:999 | ✓ 1139ms | ✓ 1957ms | ✓ 1490ms | 否 | ✓ 1634ms | http |
| 103.227.145.140:8080 | ✓ 1910ms | 否 | ✓ 1457ms | ✓ 1730ms | ✓ 1939ms | http |
| 103.39.51.207:8080 | ✓ 1480ms | 否 | ✓ 1374ms | 否 | ✓ 1559ms | http |
| 210.223.44.230:3128 | ✓ 1198ms | ✓ 1708ms | 否 | ✓ 1289ms | ✓ 1924ms | http |
| 150.249.255.91:3128 | ✓ 635ms | ✓ 1074ms | 否 | ✓ 1165ms | ✓ 805ms | http |
| 34.101.184.164:3128 | ✓ 1568ms | 否 | ✓ 1293ms | ✓ 1603ms | ✓ 1318ms | http |
| 45.12.151.226:2829 | ✓ 1578ms | 否 | ✓ 1286ms | ✓ 1573ms | 否 | http |
| 172.245.67.195:7890 | 否 | 否 | ✓ 950ms | ✓ 1295ms | ✓ 1901ms | http |
| 3.145.87.184:3015 | ✓ 1666ms | 否 | ✓ 1603ms | 否 | ✓ 1905ms | http |
| 159.223.71.162:443 | ✓ 886ms | 否 | ✓ 1121ms | ✓ 1532ms | ✓ 1589ms | http |
| 192.71.213.85:9812 | ✓ 1551ms | 否 | ✓ 1697ms | ✓ 1888ms | 否 | http |
| 3.8.3.11:3648 | ✓ 700ms | 否 | ✓ 1827ms | ✓ 1895ms | ✓ 1531ms | http |
| 8.219.97.248:80 | ✓ 1293ms | 否 | ✓ 1949ms | ✓ 1347ms | 否 | http |
| 45.77.249.199:1236 | ✓ 931ms | 否 | 否 | ✓ 1356ms | ✓ 964ms | http |
| 47.105.98.23:3128 | ✓ 1649ms | 否 | 否 | ✓ 1326ms | ✓ 1911ms | http |
| 38.34.179.24:8447 | ✓ 1741ms | 否 | ✓ 606ms | 否 | ✓ 767ms | http |
| 38.34.179.87:8447 | ✓ 994ms | 否 | ✓ 1596ms | ✓ 989ms | ✓ 1075ms | http |
| 38.145.208.209:8447 | ✓ 1719ms | ✓ 964ms | ✓ 1371ms | 否 | ✓ 821ms | http |
| 45.136.130.182:8446 | ✓ 1463ms | ✓ 1448ms | ✓ 1292ms | 否 | ✓ 829ms | http |
| 38.145.208.211:8453 | 否 | ✓ 828ms | ✓ 1224ms | 否 | ✓ 803ms | http |
| 45.136.130.178:8453 | ✓ 1463ms | ✓ 1448ms | ✓ 1291ms | 否 | ✓ 826ms | http |
| 45.136.130.183:8447 | ✓ 1716ms | ✓ 1146ms | ✓ 1783ms | 否 | ✓ 727ms | http |
| 38.34.179.154:8453 | ✓ 1970ms | ✓ 1937ms | 否 | ✓ 1607ms | ✓ 787ms | http |
| 45.136.198.40:3128 | ✓ 1509ms | ✓ 1432ms | ✓ 1436ms | 否 | ✓ 1731ms | http |
| 64.227.76.27:1080 | ✓ 768ms | 否 | ✓ 828ms | ✓ 1641ms | ✓ 1659ms | http |
| 62.113.119.14:8080 | ✓ 1052ms | ✓ 1527ms | ✓ 1225ms | ✓ 1509ms | ✓ 1097ms | http |
| 86.53.183.16:1080 | ✓ 1051ms | 否 | ✓ 1108ms | 否 | ✓ 1711ms | http |
| 103.217.224.75:3125 | 否 | 否 | ✓ 1406ms | ✓ 1555ms | ✓ 1500ms | http |
| 34.96.238.40:8080 | ✓ 1579ms | ✓ 1273ms | 否 | ✓ 1554ms | 否 | http |
| 217.77.102.18:3128 | 否 | 否 | ✓ 1516ms | ✓ 1797ms | ✓ 1616ms | http |
| 92.119.127.211:6005 | ✓ 1088ms | 否 | ✓ 1864ms | ✓ 1849ms | 否 | http |
| 154.64.230.89:3128 | ✓ 701ms | ✓ 1120ms | ✓ 877ms | ✓ 1338ms | ✓ 755ms | http |
| 174.140.109.250:3128 | ✓ 694ms | ✓ 1266ms | ✓ 944ms | ✓ 1125ms | ✓ 1055ms | http |
| 104.248.243.244:3128 | ✓ 461ms | ✓ 1477ms | ✓ 1684ms | ✓ 1662ms | ✓ 987ms | http |
| 47.74.226.8:5001 | ✓ 1107ms | ✓ 1663ms | 否 | ✓ 1648ms | 否 | http |
| 38.145.218.14:8446 | ✓ 679ms | ✓ 1211ms | ✓ 1757ms | ✓ 874ms | ✓ 1766ms | http |
| 38.145.208.185:8453 | ✓ 680ms | ✓ 840ms | 否 | ✓ 898ms | ✓ 1043ms | http |
| 38.34.179.69:8447 | ✓ 663ms | ✓ 1564ms | ✓ 1786ms | ✓ 1002ms | 否 | http |
| 61.76.95.217:40088 | ✓ 1481ms | ✓ 1616ms | ✓ 1099ms | ✓ 1413ms | ✓ 1140ms | http |
| 38.145.218.51:8444 | ✓ 668ms | ✓ 1346ms | 否 | ✓ 1302ms | 否 | http |
| 38.34.179.175:8445 | 否 | ✓ 1609ms | ✓ 657ms | 否 | ✓ 1416ms | http |
| 38.145.208.204:8446 | ✓ 1316ms | ✓ 891ms | ✓ 1809ms | ✓ 1801ms | ✓ 1984ms | http |
| 45.136.131.57:8449 | ✓ 1682ms | ✓ 1026ms | 否 | ✓ 1717ms | 否 | http |
| 121.43.189.36:28888 | ✓ 1569ms | ✓ 1599ms | ✓ 1940ms | 否 | 否 | http |
| 104.248.81.109:3128 | ✓ 603ms | ✓ 1465ms | ✓ 439ms | ✓ 1280ms | ✓ 981ms | http |
| 38.34.179.103:8448 | ✓ 1256ms | ✓ 1129ms | 否 | ✓ 1689ms | 否 | http |
| 38.145.208.227:8447 | 否 | ✓ 1310ms | ✓ 1421ms | 否 | ✓ 818ms | http |
| 38.145.208.224:8445 | ✓ 1680ms | ✓ 1541ms | 否 | ✓ 1313ms | ✓ 877ms | http |
| 18.100.127.30:23093 | ✓ 1772ms | 否 | ✓ 1763ms | 否 | ✓ 1646ms | http |
| 103.82.93.98:3128 | ✓ 1398ms | 否 | ✓ 903ms | ✓ 1927ms | ✓ 1477ms | http |
| 3.8.4.205:36785 | ✓ 1692ms | 否 | 否 | ✓ 1954ms | ✓ 1631ms | http |

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
