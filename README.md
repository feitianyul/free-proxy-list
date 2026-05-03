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

最后更新：2026-05-03 21:41:54 UTC（2026-05-04 05:41:54 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 91.233.223.147:3128 | ✓ 1610ms | 否 | ✓ 1733ms | 否 | ✓ 1986ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1925ms | ✓ 1924ms | ✓ 1613ms | ✓ 1577ms | http |
| 148.230.4.241:999 | ✓ 1326ms | ✓ 1913ms | ✓ 601ms | ✓ 1653ms | ✓ 1288ms | http |
| 45.167.124.71:999 | ✓ 1347ms | 否 | ✓ 1387ms | ✓ 1937ms | ✓ 1466ms | http |
| 168.110.52.228:3128 | ✓ 1407ms | ✓ 1120ms | ✓ 742ms | ✓ 954ms | ✓ 802ms | http |
| 8.211.166.184:8081 | ✓ 1437ms | ✓ 1326ms | ✓ 700ms | ✓ 885ms | ✓ 658ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 1996ms | ✓ 1181ms | ✓ 819ms | http |
| 206.206.126.177:2412 | ✓ 1472ms | ✓ 1362ms | ✓ 1055ms | ✓ 1042ms | ✓ 796ms | http |
| 1.231.81.166:3128 | ✓ 1510ms | ✓ 1027ms | ✓ 1482ms | ✓ 1096ms | ✓ 1086ms | http |
| 86.104.74.110:1081 | ✓ 693ms | ✓ 1945ms | ✓ 757ms | ✓ 1807ms | ✓ 1520ms | http |
| 218.108.131.186:17890 | ✓ 826ms | ✓ 1957ms | 否 | ✓ 1169ms | ✓ 916ms | http |
| 46.105.190.40:3128 | ✓ 1046ms | 否 | ✓ 677ms | ✓ 1914ms | 否 | http |
| 103.3.246.71:3128 | ✓ 943ms | 否 | ✓ 1277ms | 否 | ✓ 960ms | http |
| 45.153.231.229:8080 | ✓ 876ms | ✓ 1928ms | ✓ 1650ms | 否 | 否 | http |
| 34.101.184.164:3128 | ✓ 1677ms | 否 | ✓ 1287ms | ✓ 1853ms | ✓ 1692ms | http |
| 47.77.216.82:1080 | ✓ 106ms | 否 | ✓ 713ms | ✓ 1147ms | ✓ 747ms | http |
| 154.64.232.35:8080 | ✓ 810ms | 否 | ✓ 907ms | ✓ 1219ms | 否 | http |
| 20.164.75.153:8080 | ✓ 1960ms | 否 | ✓ 1854ms | 否 | ✓ 1965ms | http |
| 8.219.97.248:80 | ✓ 1580ms | 否 | ✓ 1841ms | 否 | ✓ 1259ms | http |
| 152.32.132.190:7890 | ✓ 767ms | ✓ 1933ms | ✓ 840ms | 否 | 否 | http |
| 45.140.147.82:1082 | ✓ 1174ms | ✓ 1677ms | ✓ 676ms | 否 | ✓ 1522ms | http |
| 117.236.124.166:3128 | ✓ 1648ms | 否 | ✓ 1106ms | 否 | ✓ 1746ms | http |
| 193.123.250.39:1080 | ✓ 1588ms | 否 | ✓ 1796ms | ✓ 1671ms | ✓ 1079ms | http |
| 103.157.200.126:3128 | ✓ 1369ms | 否 | ✓ 1806ms | 否 | ✓ 1482ms | http |
| 38.188.247.12:999 | 否 | 否 | ✓ 433ms | ✓ 1560ms | ✓ 1497ms | http |
| 190.12.150.244:999 | ✓ 1663ms | 否 | ✓ 1212ms | ✓ 1768ms | ✓ 1738ms | http |
| 116.171.106.111:3443 | ✓ 1398ms | ✓ 1432ms | ✓ 1509ms | ✓ 1678ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1193ms | ✓ 1736ms | ✓ 1321ms | ✓ 1611ms | ✓ 1873ms | http |
| 59.46.216.131:30001 | ✓ 1016ms | ✓ 1475ms | ✓ 1076ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 725ms | ✓ 1411ms | ✓ 1297ms | ✓ 1543ms | ✓ 1148ms | http |
| 109.120.156.122:8090 | ✓ 1211ms | 否 | ✓ 1786ms | 否 | ✓ 1799ms | http |
| 103.240.6.22:16498 | ✓ 1481ms | 否 | ✓ 1233ms | ✓ 1764ms | ✓ 1406ms | http |
| 80.92.204.47:1081 | ✓ 712ms | ✓ 1496ms | ✓ 1162ms | 否 | ✓ 1830ms | http |
| 120.92.212.16:8890 | ✓ 1768ms | ✓ 1254ms | ✓ 1659ms | 否 | ✓ 1175ms | http |
| 86.104.72.219:1081 | ✓ 744ms | ✓ 1337ms | ✓ 1557ms | ✓ 1831ms | 否 | http |
| 86.104.72.219:1082 | ✓ 735ms | ✓ 1460ms | ✓ 900ms | 否 | 否 | http |
| 8.154.21.175:3128 | ✓ 877ms | ✓ 1087ms | ✓ 867ms | ✓ 1146ms | ✓ 924ms | http |
| 121.230.8.34:1080 | ✓ 1173ms | ✓ 1589ms | ✓ 958ms | ✓ 1728ms | ✓ 1125ms | http |
| 45.125.67.37:8443 | ✓ 917ms | 否 | ✓ 868ms | ✓ 1124ms | ✓ 1205ms | http |
| 103.166.182.144:3128 | ✓ 1655ms | 否 | ✓ 968ms | ✓ 1191ms | ✓ 992ms | http |
| 77.110.107.80:8080 | ✓ 1272ms | ✓ 1588ms | ✓ 1994ms | 否 | 否 | http |
| 120.92.108.86:7890 | ✓ 1836ms | 否 | ✓ 1734ms | ✓ 1868ms | ✓ 1436ms | http |
| 103.155.199.15:8080 | ✓ 1346ms | ✓ 1858ms | ✓ 1733ms | ✓ 1443ms | ✓ 1445ms | http |
| 119.95.173.101:8082 | ✓ 1338ms | 否 | ✓ 1245ms | ✓ 1817ms | ✓ 1425ms | http |
| 103.171.245.165:1080 | ✓ 1284ms | 否 | ✓ 1297ms | ✓ 1565ms | 否 | http |
| 103.217.216.65:8181 | ✓ 1342ms | 否 | ✓ 1854ms | 否 | ✓ 1458ms | http |
| 103.87.202.198:1111 | ✓ 1980ms | 否 | ✓ 1696ms | ✓ 1513ms | ✓ 1461ms | http |
| 103.172.121.61:7778 | ✓ 1856ms | 否 | 否 | ✓ 1956ms | ✓ 1539ms | http |
| 120.92.212.16:7890 | ✓ 1071ms | ✓ 1358ms | ✓ 1022ms | ✓ 1514ms | ✓ 1016ms | http |
| 43.133.44.89:8888 | ✓ 917ms | 否 | ✓ 1054ms | ✓ 1033ms | ✓ 1588ms | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1197ms | ✓ 1426ms | ✓ 1325ms | http |
| 125.76.214.178:8091 | 否 | 否 | ✓ 1122ms | ✓ 1448ms | ✓ 1126ms | http |
| 154.90.48.209:9090 | ✓ 1427ms | 否 | ✓ 1299ms | ✓ 1904ms | ✓ 1773ms | http |
| 101.32.243.189:80 | ✓ 1297ms | 否 | ✓ 1661ms | ✓ 1424ms | 否 | http |
| 23.253.80.88:80 | ✓ 1152ms | ✓ 1595ms | 否 | 否 | ✓ 1724ms | http |
| 86.104.72.220:1082 | ✓ 780ms | ✓ 1775ms | ✓ 242ms | ✓ 1141ms | ✓ 1804ms | http |
| 3.101.133.120:80 | ✓ 243ms | ✓ 1116ms | ✓ 1392ms | ✓ 1379ms | ✓ 848ms | http |
| 86.104.72.220:1081 | ✓ 859ms | ✓ 1017ms | ✓ 255ms | ✓ 1212ms | ✓ 878ms | http |
| 116.254.118.180:80 | ✓ 1562ms | 否 | ✓ 1077ms | ✓ 1294ms | ✓ 1012ms | http |
| 46.105.190.38:3128 | ✓ 1392ms | ✓ 1396ms | ✓ 662ms | 否 | 否 | http |
| 139.159.97.82:10900 | ✓ 1569ms | 否 | ✓ 1374ms | ✓ 1756ms | 否 | http |
| 45.59.122.132:80 | ✓ 1169ms | 否 | ✓ 1798ms | 否 | ✓ 1733ms | http |
| 103.209.36.58:8080 | 否 | 否 | ✓ 1525ms | ✓ 1689ms | ✓ 1573ms | http |
| 45.129.141.143:3128 | ✓ 1233ms | ✓ 1892ms | ✓ 1435ms | 否 | ✓ 1677ms | http |
| 138.124.73.27:10808 | ✓ 1715ms | 否 | ✓ 1546ms | ✓ 1617ms | ✓ 1769ms | http |
| 104.128.138.186:1080 | ✓ 1281ms | 否 | ✓ 1890ms | ✓ 1909ms | ✓ 1874ms | http |
| 62.133.60.126:24558 | ✓ 682ms | ✓ 1825ms | 否 | ✓ 1778ms | 否 | http |
| 210.48.154.94:80 | ✓ 1868ms | 否 | ✓ 1456ms | ✓ 1338ms | ✓ 1060ms | http |
| 37.187.109.70:10111 | ✓ 1426ms | ✓ 1649ms | ✓ 1959ms | 否 | 否 | http |
| 185.21.11.140:1080 | ✓ 1075ms | ✓ 1970ms | ✓ 1298ms | ✓ 1815ms | ✓ 1365ms | http |
| 121.230.8.220:1080 | ✓ 1270ms | ✓ 1232ms | ✓ 1110ms | ✓ 1527ms | ✓ 1198ms | http |
| 103.67.46.225:3125 | ✓ 1752ms | 否 | 否 | ✓ 1641ms | ✓ 1602ms | http |
| 121.230.8.250:1080 | ✓ 1468ms | ✓ 1630ms | ✓ 1460ms | ✓ 1408ms | ✓ 1701ms | http |
| 61.52.131.172:8443 | ✓ 886ms | ✓ 1222ms | ✓ 1064ms | ✓ 1254ms | ✓ 1006ms | http |
| 60.249.94.208:3128 | ✓ 1013ms | ✓ 1001ms | ✓ 789ms | 否 | 否 | http |
| 168.222.254.136:8888 | 否 | ✓ 1907ms | ✓ 1755ms | 否 | ✓ 1753ms | http |

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
