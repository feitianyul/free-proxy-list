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

最后更新：2026-04-20 12:11:29 UTC（2026-04-20 20:11:29 UTC+8）

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
| 46.101.95.183:8888 | ✓ 1142ms | 否 | ✓ 1214ms | ✓ 1770ms | ✓ 1704ms | http |
| 188.246.224.49:7890 | ✓ 1165ms | ✓ 1694ms | ✓ 1968ms | 否 | ✓ 1577ms | http |
| 152.42.208.139:8118 | ✓ 1651ms | 否 | ✓ 1369ms | ✓ 1290ms | ✓ 1366ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1908ms | ✓ 1114ms | ✓ 1003ms | http |
| 81.30.156.115:8080 | ✓ 777ms | ✓ 1598ms | ✓ 722ms | 否 | ✓ 1688ms | http |
| 91.99.15.45:2095 | 否 | 否 | ✓ 1666ms | ✓ 1877ms | ✓ 1913ms | http |
| 185.138.116.150:8080 | ✓ 1355ms | ✓ 1609ms | ✓ 607ms | 否 | 否 | http |
| 103.125.181.135:9999 | ✓ 1911ms | 否 | ✓ 1699ms | ✓ 1871ms | ✓ 1183ms | http |
| 113.160.132.26:8080 | ✓ 1789ms | ✓ 1609ms | ✓ 1309ms | ✓ 1441ms | 否 | http |
| 217.77.102.18:3128 | ✓ 1266ms | 否 | 否 | ✓ 1782ms | ✓ 1886ms | http |
| 59.46.216.131:30001 | ✓ 1228ms | 否 | ✓ 1297ms | ✓ 1660ms | ✓ 1264ms | http |
| 223.84.151.86:30005 | ✓ 1554ms | ✓ 1528ms | ✓ 1251ms | ✓ 1729ms | ✓ 1582ms | http |
| 138.124.99.216:8888 | ✓ 683ms | 否 | 否 | ✓ 1857ms | ✓ 1789ms | http |
| 45.12.151.226:2829 | ✓ 572ms | 否 | ✓ 814ms | ✓ 1664ms | 否 | http |
| 149.51.42.10:3128 | ✓ 879ms | ✓ 1242ms | 否 | ✓ 1295ms | 否 | http |
| 149.51.42.10:8080 | ✓ 1466ms | ✓ 1176ms | 否 | ✓ 1222ms | 否 | http |
| 107.150.41.226:18080 | ✓ 380ms | ✓ 1794ms | ✓ 222ms | ✓ 1494ms | ✓ 1092ms | http |
| 34.71.229.255:3128 | ✓ 583ms | 否 | ✓ 1014ms | ✓ 925ms | ✓ 1229ms | http |
| 168.222.254.136:8888 | ✓ 881ms | ✓ 1780ms | ✓ 1999ms | 否 | 否 | http |
| 152.70.91.193:40000 | ✓ 1807ms | 否 | ✓ 1816ms | ✓ 1190ms | ✓ 1415ms | http |
| 168.144.75.9:3128 | ✓ 973ms | 否 | ✓ 1602ms | ✓ 1999ms | 否 | http |
| 84.47.150.125:1080 | ✓ 924ms | 否 | ✓ 1929ms | 否 | ✓ 1619ms | http |
| 45.140.147.82:1081 | ✓ 1606ms | ✓ 1475ms | ✓ 710ms | ✓ 1954ms | ✓ 1325ms | http |
| 77.110.113.24:40000 | ✓ 1170ms | 否 | ✓ 743ms | 否 | ✓ 1939ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1424ms | ✓ 1562ms | ✓ 1424ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1677ms | 否 | ✓ 1090ms | ✓ 1608ms | ✓ 1080ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1860ms | ✓ 1390ms | ✓ 1741ms | http |
| 78.11.96.22:8888 | 否 | ✓ 1535ms | ✓ 924ms | ✓ 1404ms | ✓ 1242ms | http |
| 168.222.254.26:8888 | ✓ 569ms | ✓ 1592ms | 否 | 否 | ✓ 1538ms | http |
| 45.153.231.229:8080 | ✓ 774ms | 否 | ✓ 1537ms | 否 | ✓ 1673ms | http |
| 139.159.97.82:10900 | ✓ 1861ms | 否 | ✓ 1588ms | ✓ 1715ms | ✓ 1335ms | http |
| 159.89.191.221:3128 | ✓ 199ms | ✓ 1959ms | ✓ 1315ms | ✓ 877ms | ✓ 680ms | http |
| 117.236.124.166:3128 | ✓ 1718ms | 否 | ✓ 1062ms | ✓ 1886ms | 否 | http |
| 161.97.184.191:8080 | ✓ 893ms | 否 | 否 | ✓ 1826ms | ✓ 1532ms | http |
| 85.190.99.143:443 | ✓ 1693ms | 否 | ✓ 1710ms | 否 | ✓ 1916ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1869ms | ✓ 168ms | ✓ 1277ms | ✓ 745ms | http |
| 192.3.248.190:8014 | ✓ 646ms | ✓ 1432ms | ✓ 626ms | ✓ 1645ms | ✓ 1050ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1728ms | ✓ 1549ms | ✓ 931ms | http |
| 121.230.8.91:1080 | ✓ 1182ms | ✓ 1639ms | ✓ 1441ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1091ms | ✓ 1898ms | ✓ 1152ms | ✓ 901ms | http |
| 208.87.243.199:7878 | 否 | ✓ 1522ms | 否 | ✓ 1598ms | ✓ 1067ms | http |
| 102.134.49.165:6005 | ✓ 1470ms | ✓ 1537ms | ✓ 1791ms | ✓ 1180ms | ✓ 1029ms | http |
| 160.250.134.143:3128 | ✓ 1404ms | 否 | ✓ 1145ms | ✓ 1477ms | ✓ 1170ms | http |
| 57.128.188.167:9173 | ✓ 1525ms | 否 | ✓ 1513ms | 否 | ✓ 1685ms | http |
| 82.114.228.67:1080 | ✓ 716ms | ✓ 1643ms | ✓ 1136ms | 否 | ✓ 1100ms | http |
| 159.223.225.118:8888 | ✓ 1839ms | 否 | ✓ 1281ms | ✓ 1747ms | 否 | http |
| 27.71.24.102:3128 | ✓ 1614ms | 否 | ✓ 1611ms | ✓ 1338ms | ✓ 1128ms | http |
| 144.31.25.69:21064 | ✓ 992ms | 否 | ✓ 1906ms | 否 | ✓ 1993ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1457ms | ✓ 1524ms | 否 | ✓ 907ms | http |
| 157.230.178.216:8088 | ✓ 1974ms | 否 | ✓ 1858ms | ✓ 1487ms | ✓ 1336ms | http |
| 91.193.240.157:9877 | ✓ 1322ms | 否 | ✓ 1898ms | 否 | ✓ 1459ms | http |
| 101.32.243.189:80 | ✓ 1577ms | 否 | ✓ 1776ms | 否 | ✓ 1847ms | http |
| 120.92.212.16:7890 | ✓ 1161ms | 否 | ✓ 1314ms | 否 | ✓ 1097ms | http |
| 177.93.132.244:3128 | ✓ 1253ms | 否 | ✓ 1084ms | 否 | ✓ 1627ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1995ms | ✓ 1681ms | ✓ 1780ms | http |
| 147.45.186.28:3128 | 否 | 否 | ✓ 1932ms | ✓ 1700ms | ✓ 1361ms | http |
| 130.61.30.221:8080 | ✓ 910ms | 否 | ✓ 1336ms | ✓ 1559ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1096ms | ✓ 1344ms | ✓ 1111ms | ✓ 1200ms | ✓ 1378ms | http |
| 91.107.124.215:3128 | ✓ 743ms | ✓ 1599ms | ✓ 919ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 1361ms | 否 | 否 | ✓ 1656ms | ✓ 1905ms | http |
| 20.205.16.149:3128 | ✓ 1702ms | 否 | 否 | ✓ 1497ms | ✓ 1098ms | http |
| 194.104.9.38:3128 | ✓ 390ms | 否 | ✓ 1166ms | ✓ 1506ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1231ms | ✓ 1468ms | 否 | 否 | ✓ 1228ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1800ms | ✓ 1388ms | ✓ 1897ms | http |
| 20.120.225.109:3128 | ✓ 1201ms | 否 | 否 | ✓ 1283ms | ✓ 1279ms | http |
| 172.86.66.149:8080 | ✓ 1458ms | 否 | ✓ 705ms | ✓ 1574ms | ✓ 1330ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1070ms | ✓ 1464ms | ✓ 1704ms | http |
| 42.101.8.101:8888 | ✓ 1310ms | 否 | ✓ 1663ms | 否 | ✓ 1328ms | http |
| 220.197.44.36:3128 | ✓ 1374ms | ✓ 1886ms | ✓ 1524ms | ✓ 1696ms | ✓ 1238ms | http |
| 212.58.132.5:8888 | ✓ 1207ms | 否 | ✓ 1501ms | ✓ 1571ms | ✓ 1595ms | http |
| 121.230.8.55:1080 | ✓ 1349ms | ✓ 1769ms | 否 | ✓ 1699ms | 否 | http |
| 15.204.151.147:3128 | 否 | 否 | ✓ 1391ms | ✓ 1779ms | ✓ 1553ms | http |
| 61.52.131.172:8443 | ✓ 1062ms | ✓ 1353ms | ✓ 1156ms | ✓ 1399ms | ✓ 1106ms | http |
| 60.249.94.208:3128 | ✓ 995ms | 否 | ✓ 912ms | ✓ 1340ms | ✓ 985ms | http |
| 103.85.113.66:9999 | ✓ 1050ms | ✓ 1801ms | ✓ 1829ms | ✓ 1840ms | 否 | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1708ms | ✓ 1842ms | ✓ 1202ms | http |

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
