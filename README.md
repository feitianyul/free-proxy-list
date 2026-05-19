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

最后更新：2026-05-19 22:15:08 UTC（2026-05-20 06:15:08 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 192.99.8.15:8850 | ✓ 1070ms | 否 | ✓ 1495ms | ✓ 1537ms | ✓ 1494ms | http |
| 43.130.126.146:6688 | ✓ 863ms | 否 | 否 | ✓ 1859ms | ✓ 1057ms | http |
| 176.111.37.5:39811 | ✓ 1411ms | ✓ 1386ms | ✓ 1987ms | 否 | ✓ 1731ms | http |
| 176.111.37.216:39811 | ✓ 1412ms | ✓ 1643ms | ✓ 1731ms | 否 | ✓ 1815ms | http |
| 185.200.188.234:10001 | ✓ 1446ms | 否 | ✓ 1639ms | 否 | ✓ 1563ms | http |
| 1.231.81.166:3128 | ✓ 1863ms | ✓ 1485ms | ✓ 1988ms | ✓ 1581ms | ✓ 1231ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1729ms | ✓ 1615ms | 否 | ✓ 1229ms | http |
| 202.28.194.139:31280 | 否 | 否 | ✓ 1852ms | ✓ 1989ms | ✓ 1918ms | http |
| 138.2.78.251:8100 | ✓ 948ms | 否 | ✓ 1820ms | ✓ 1499ms | ✓ 1234ms | http |
| 167.179.101.174:8890 | ✓ 670ms | ✓ 1371ms | ✓ 1319ms | ✓ 1072ms | ✓ 865ms | http |
| 144.31.25.69:21064 | 否 | 否 | ✓ 721ms | ✓ 1821ms | ✓ 1359ms | http |
| 174.137.134.182:2999 | 否 | 否 | ✓ 748ms | ✓ 1532ms | ✓ 1838ms | http |
| 147.45.78.89:1080 | ✓ 921ms | 否 | ✓ 1616ms | 否 | ✓ 1541ms | http |
| 45.125.67.37:8443 | ✓ 1568ms | 否 | ✓ 1351ms | ✓ 1373ms | ✓ 1389ms | http |
| 5.252.33.13:2025 | ✓ 1263ms | 否 | 否 | ✓ 1959ms | ✓ 1710ms | http |
| 147.45.41.112:1080 | ✓ 1576ms | 否 | ✓ 1435ms | ✓ 1837ms | ✓ 1222ms | http |
| 207.148.124.152:6868 | ✓ 1111ms | 否 | ✓ 1639ms | ✓ 1943ms | 否 | http |
| 138.2.92.70:8100 | ✓ 1439ms | 否 | ✓ 1147ms | ✓ 1887ms | 否 | http |
| 192.81.129.252:3136 | ✓ 596ms | ✓ 1102ms | ✓ 1268ms | ✓ 977ms | 否 | http |
| 74.208.192.81:3129 | ✓ 466ms | 否 | ✓ 1769ms | 否 | ✓ 1281ms | http |
| 89.58.50.94:11140 | ✓ 971ms | 否 | ✓ 926ms | ✓ 1987ms | ✓ 1498ms | http |
| 158.255.212.55:9300 | ✓ 996ms | 否 | ✓ 1173ms | ✓ 1884ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1000ms | 否 | ✓ 1437ms | 否 | ✓ 1978ms | http |
| 8.154.21.175:3128 | 否 | ✓ 1251ms | ✓ 1330ms | 否 | ✓ 1646ms | http |
| 170.106.136.181:31002 | ✓ 1199ms | ✓ 1541ms | ✓ 997ms | ✓ 915ms | ✓ 1742ms | http |
| 146.190.80.158:9090 | ✓ 1357ms | 否 | 否 | ✓ 1459ms | ✓ 1091ms | http |
| 86.104.72.220:1081 | ✓ 1064ms | ✓ 1549ms | ✓ 1130ms | 否 | 否 | http |
| 158.160.215.167:8127 | ✓ 1223ms | ✓ 1798ms | ✓ 1692ms | 否 | ✓ 1806ms | http |
| 148.230.4.241:999 | ✓ 955ms | ✓ 1616ms | ✓ 643ms | ✓ 1873ms | ✓ 1376ms | http |
| 78.186.118.164:3311 | ✓ 1209ms | 否 | ✓ 1719ms | ✓ 1690ms | ✓ 1317ms | http |
| 152.67.191.232:6800 | ✓ 1706ms | 否 | 否 | ✓ 1811ms | ✓ 1731ms | http |
| 47.84.93.78:8100 | 否 | 否 | ✓ 1213ms | ✓ 1345ms | ✓ 1648ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1317ms | ✓ 1583ms | ✓ 1317ms | http |
| 34.87.80.221:30000 | ✓ 989ms | 否 | ✓ 1225ms | ✓ 1358ms | ✓ 1057ms | http |
| 104.168.96.172:1888 | 否 | ✓ 1823ms | ✓ 1780ms | ✓ 1433ms | ✓ 1398ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1816ms | ✓ 1069ms | 否 | ✓ 1115ms | http |
| 138.2.239.213:10010 | ✓ 875ms | ✓ 1692ms | 否 | ✓ 1696ms | ✓ 880ms | http |
| 43.159.128.164:10000 | 否 | 否 | ✓ 1188ms | ✓ 954ms | ✓ 752ms | http |
| 167.99.173.119:3128 | 否 | 否 | ✓ 1003ms | ✓ 1185ms | ✓ 755ms | http |
| 103.189.197.43:7778 | ✓ 1855ms | 否 | ✓ 1222ms | ✓ 1740ms | 否 | http |
| 81.30.156.115:8080 | ✓ 1161ms | 否 | ✓ 1452ms | ✓ 1877ms | ✓ 1774ms | http |
| 121.230.9.33:1080 | ✓ 1313ms | ✓ 1929ms | ✓ 1458ms | 否 | ✓ 1343ms | http |
| 103.110.10.174:3300 | ✓ 1576ms | 否 | ✓ 1672ms | ✓ 1714ms | ✓ 1939ms | http |
| 45.80.231.251:3128 | ✓ 1175ms | 否 | ✓ 1361ms | ✓ 1860ms | ✓ 1675ms | http |
| 85.192.29.60:3128 | ✓ 1109ms | ✓ 1661ms | ✓ 1608ms | ✓ 1540ms | ✓ 1310ms | http |
| 31.172.78.12:3128 | 否 | 否 | ✓ 1085ms | ✓ 1522ms | ✓ 1330ms | http |
| 38.19.41.227:999 | 否 | 否 | ✓ 879ms | ✓ 1531ms | ✓ 1322ms | http |
| 3.101.133.120:80 | ✓ 879ms | ✓ 1508ms | ✓ 1637ms | ✓ 1405ms | ✓ 1193ms | http |
| 217.52.247.75:1981 | ✓ 1856ms | 否 | ✓ 1745ms | ✓ 1924ms | 否 | http |
| 41.33.219.130:1976 | ✓ 1366ms | 否 | ✓ 1958ms | 否 | ✓ 1839ms | http |
| 94.131.122.128:1081 | 否 | ✓ 1072ms | ✓ 1064ms | ✓ 1944ms | ✓ 1425ms | http |
| 45.153.231.229:8080 | ✓ 1039ms | ✓ 1634ms | 否 | 否 | ✓ 1597ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1515ms | 否 | ✓ 1164ms | ✓ 1684ms | http |
| 20.164.75.153:8080 | ✓ 1403ms | 否 | ✓ 1295ms | 否 | ✓ 1755ms | http |
| 158.160.215.167:8123 | ✓ 1033ms | ✓ 1837ms | ✓ 1332ms | 否 | 否 | http |
| 212.34.146.118:3128 | ✓ 1923ms | 否 | ✓ 894ms | ✓ 1675ms | ✓ 1670ms | http |
| 45.117.163.134:3128 | ✓ 1660ms | 否 | ✓ 1019ms | ✓ 1332ms | ✓ 1054ms | http |
| 114.214.163.108:6789 | ✓ 1307ms | ✓ 1607ms | ✓ 1643ms | ✓ 1607ms | ✓ 1318ms | http |
| 114.214.165.78:10810 | ✓ 1555ms | ✓ 1925ms | ✓ 1566ms | ✓ 1628ms | ✓ 1599ms | http |
| 223.16.170.103:80 | ✓ 1144ms | 否 | ✓ 1099ms | 否 | ✓ 1365ms | http |
| 103.157.117.226:81 | 否 | 否 | ✓ 1498ms | ✓ 1609ms | ✓ 1597ms | http |
| 168.110.52.228:3128 | ✓ 667ms | ✓ 1193ms | ✓ 695ms | ✓ 1033ms | ✓ 794ms | http |
| 94.131.122.128:1082 | ✓ 368ms | ✓ 1330ms | ✓ 415ms | 否 | 否 | http |
| 103.82.23.118:5195 | ✓ 1955ms | 否 | ✓ 1992ms | 否 | ✓ 1624ms | http |
| 64.188.77.26:3128 | ✓ 1179ms | 否 | ✓ 1667ms | 否 | ✓ 1383ms | http |
| 121.130.199.80:24007 | ✓ 1920ms | ✓ 1734ms | 否 | ✓ 1686ms | ✓ 1335ms | http |
| 103.157.78.85:8080 | ✓ 1655ms | 否 | ✓ 1942ms | ✓ 1690ms | ✓ 1809ms | http |
| 61.144.152.5:9000 | 否 | ✓ 1698ms | ✓ 1547ms | ✓ 1508ms | ✓ 1579ms | http |
| 190.12.150.244:999 | ✓ 1192ms | ✓ 1647ms | ✓ 941ms | 否 | 否 | http |
| 77.110.107.80:8080 | 否 | 否 | ✓ 1613ms | ✓ 1735ms | ✓ 1902ms | http |
| 104.248.151.93:9090 | ✓ 1324ms | 否 | ✓ 931ms | ✓ 1394ms | ✓ 1168ms | http |
| 129.212.224.122:3128 | ✓ 933ms | 否 | ✓ 918ms | ✓ 1313ms | ✓ 1017ms | http |
| 8.219.97.248:80 | ✓ 1283ms | 否 | ✓ 1353ms | 否 | ✓ 1555ms | http |
| 91.233.223.147:3128 | ✓ 1238ms | 否 | ✓ 1291ms | ✓ 1890ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1136ms | 否 | ✓ 1542ms | ✓ 1509ms | ✓ 1266ms | http |
| 168.138.171.204:8100 | 否 | 否 | ✓ 1830ms | ✓ 1862ms | ✓ 1658ms | http |
| 64.188.77.221:3128 | ✓ 1156ms | 否 | ✓ 1725ms | ✓ 1946ms | 否 | http |
| 103.69.84.106:3131 | 否 | 否 | ✓ 1562ms | ✓ 1470ms | ✓ 1136ms | http |
| 213.131.85.26:1976 | 否 | 否 | ✓ 1919ms | ✓ 1617ms | ✓ 1341ms | http |
| 121.230.9.217:1080 | ✓ 1382ms | ✓ 1973ms | ✓ 1202ms | ✓ 1506ms | ✓ 1108ms | http |
| 159.223.41.216:9090 | ✓ 1190ms | 否 | ✓ 1223ms | ✓ 1930ms | ✓ 1210ms | http |
| 121.230.8.97:1080 | ✓ 1235ms | ✓ 1896ms | ✓ 1526ms | ✓ 1629ms | ✓ 1395ms | http |
| 51.79.71.106:8080 | 否 | ✓ 1767ms | ✓ 1273ms | 否 | ✓ 1206ms | http |
| 116.63.160.98:8899 | ✓ 1486ms | ✓ 1565ms | ✓ 1217ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1162ms | ✓ 1391ms | ✓ 1144ms | ✓ 1454ms | ✓ 1106ms | http |
| 82.114.228.67:1080 | 否 | ✓ 1732ms | ✓ 1267ms | ✓ 1563ms | 否 | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1988ms | ✓ 1650ms | ✓ 1567ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1769ms | ✓ 1548ms | ✓ 853ms | http |
| 27.254.99.183:8118 | ✓ 1101ms | 否 | ✓ 1305ms | ✓ 1558ms | 否 | http |
| 103.164.180.185:8060 | 否 | 否 | ✓ 1572ms | ✓ 1785ms | ✓ 1726ms | http |

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
