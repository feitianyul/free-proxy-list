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

最后更新：2026-05-04 17:42:42 UTC（2026-05-05 01:42:42 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 234ms | 否 | ✓ 935ms | ✓ 1224ms | ✓ 888ms | http |
| 206.206.126.177:2412 | ✓ 1062ms | ✓ 1949ms | ✓ 1255ms | ✓ 1167ms | ✓ 1156ms | http |
| 185.191.236.162:3128 | ✓ 881ms | ✓ 1622ms | 否 | 否 | ✓ 1510ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1797ms | ✓ 1134ms | ✓ 1629ms | ✓ 1493ms | http |
| 181.119.97.24:999 | ✓ 1381ms | 否 | ✓ 1761ms | 否 | ✓ 1799ms | http |
| 47.85.51.197:1080 | ✓ 795ms | ✓ 1100ms | ✓ 631ms | ✓ 1061ms | 否 | http |
| 218.108.131.186:17890 | ✓ 952ms | ✓ 1194ms | 否 | ✓ 1259ms | ✓ 1014ms | http |
| 86.104.72.220:1081 | ✓ 434ms | 否 | ✓ 564ms | ✓ 1457ms | 否 | http |
| 91.242.229.129:8092 | ✓ 1469ms | ✓ 1879ms | ✓ 1488ms | 否 | 否 | http |
| 150.107.140.238:3128 | ✓ 927ms | 否 | 否 | ✓ 1487ms | ✓ 1144ms | http |
| 116.80.93.67:3172 | 否 | 否 | ✓ 1650ms | ✓ 1980ms | ✓ 1772ms | http |
| 45.153.231.229:8080 | ✓ 1432ms | 否 | ✓ 1962ms | 否 | ✓ 1898ms | http |
| 120.92.212.16:8890 | ✓ 1531ms | ✓ 1289ms | ✓ 1363ms | 否 | 否 | http |
| 35.182.12.78:41717 | ✓ 1079ms | 否 | 否 | ✓ 1691ms | ✓ 1354ms | http |
| 45.140.147.155:1081 | ✓ 983ms | ✓ 1756ms | ✓ 1957ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 1529ms | 否 | 否 | ✓ 1679ms | ✓ 1482ms | http |
| 62.60.231.71:56608 | ✓ 1356ms | 否 | ✓ 929ms | 否 | ✓ 1597ms | http |
| 104.128.138.186:1080 | ✓ 1808ms | ✓ 1972ms | 否 | ✓ 1917ms | ✓ 1533ms | http |
| 116.171.106.111:3443 | ✓ 1438ms | ✓ 1683ms | 否 | ✓ 1890ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1273ms | 否 | ✓ 1880ms | 否 | ✓ 1857ms | http |
| 193.123.250.39:1080 | 否 | 否 | ✓ 1413ms | ✓ 1301ms | ✓ 1512ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1332ms | ✓ 1137ms | ✓ 1223ms | ✓ 1005ms | http |
| 45.140.147.155:1082 | ✓ 1392ms | 否 | ✓ 1874ms | 否 | ✓ 1724ms | http |
| 120.92.108.86:7890 | ✓ 1284ms | 否 | ✓ 1429ms | 否 | ✓ 1516ms | http |
| 106.10.55.212:1121 | ✓ 1817ms | ✓ 1271ms | ✓ 1025ms | 否 | 否 | http |
| 20.27.14.220:8561 | ✓ 1700ms | 否 | 否 | ✓ 1981ms | ✓ 1879ms | http |
| 20.27.11.248:8561 | ✓ 1695ms | 否 | 否 | ✓ 1999ms | ✓ 1868ms | http |
| 154.64.232.35:8080 | ✓ 1421ms | ✓ 1344ms | ✓ 1069ms | 否 | 否 | http |
| 20.27.15.111:8561 | ✓ 1250ms | ✓ 1477ms | ✓ 1823ms | 否 | 否 | http |
| 20.27.13.35:8561 | ✓ 1268ms | ✓ 1426ms | ✓ 1839ms | 否 | 否 | http |
| 194.59.247.34:10808 | ✓ 504ms | 否 | ✓ 1343ms | ✓ 1930ms | ✓ 1326ms | http |
| 120.92.212.16:7890 | ✓ 1084ms | 否 | 否 | ✓ 1408ms | ✓ 1100ms | http |
| 8.211.166.184:8081 | ✓ 626ms | ✓ 1208ms | ✓ 809ms | ✓ 988ms | ✓ 795ms | http |
| 43.133.44.89:8888 | 否 | 否 | ✓ 868ms | ✓ 1186ms | ✓ 921ms | http |
| 45.59.122.132:80 | ✓ 1263ms | 否 | ✓ 1346ms | 否 | ✓ 1511ms | http |
| 207.254.71.62:8088 | ✓ 1227ms | 否 | ✓ 1517ms | 否 | ✓ 1720ms | http |
| 118.31.1.154:80 | ✓ 971ms | ✓ 1209ms | ✓ 1092ms | ✓ 1272ms | ✓ 1026ms | http |
| 148.230.4.241:999 | ✓ 1638ms | ✓ 1845ms | ✓ 881ms | 否 | 否 | http |
| 190.12.150.244:999 | ✓ 1497ms | 否 | ✓ 1544ms | ✓ 1683ms | ✓ 1603ms | http |
| 150.249.255.91:3128 | ✓ 1375ms | 否 | 否 | ✓ 1058ms | ✓ 1318ms | http |
| 107.173.42.121:7890 | 否 | ✓ 958ms | ✓ 107ms | ✓ 937ms | 否 | http |
| 45.173.12.140:1994 | ✓ 656ms | ✓ 1783ms | ✓ 1312ms | ✓ 1609ms | ✓ 1317ms | http |
| 62.113.119.14:8080 | ✓ 590ms | ✓ 1549ms | ✓ 569ms | ✓ 1915ms | 否 | http |
| 86.104.72.220:1082 | ✓ 758ms | 否 | ✓ 122ms | ✓ 1314ms | ✓ 829ms | http |
| 20.210.39.153:8561 | ✓ 1604ms | ✓ 1148ms | ✓ 614ms | ✓ 967ms | ✓ 722ms | http |
| 20.78.26.206:8561 | ✓ 1602ms | ✓ 1102ms | ✓ 660ms | ✓ 973ms | ✓ 750ms | http |
| 8.154.21.175:3128 | ✓ 1002ms | ✓ 1318ms | ✓ 1036ms | ✓ 1207ms | ✓ 1037ms | http |
| 20.78.118.91:8561 | ✓ 1607ms | ✓ 1201ms | ✓ 685ms | ✓ 931ms | ✓ 802ms | http |
| 114.237.77.239:1080 | ✓ 1080ms | ✓ 1399ms | ✓ 1193ms | ✓ 1635ms | ✓ 1078ms | http |
| 45.125.67.37:8443 | ✓ 1044ms | 否 | ✓ 1332ms | ✓ 1232ms | ✓ 1273ms | http |
| 34.101.184.164:3128 | ✓ 1789ms | 否 | ✓ 1504ms | ✓ 1448ms | ✓ 1265ms | http |
| 157.254.37.238:999 | ✓ 865ms | 否 | ✓ 1263ms | ✓ 1730ms | ✓ 1446ms | http |
| 193.181.35.184:8118 | ✓ 1267ms | 否 | ✓ 816ms | ✓ 1977ms | ✓ 1529ms | http |
| 103.173.139.220:8080 | ✓ 1817ms | 否 | 否 | ✓ 1575ms | ✓ 1789ms | http |
| 152.70.91.193:40000 | ✓ 1571ms | 否 | 否 | ✓ 1681ms | ✓ 1720ms | http |
| 38.180.121.135:10808 | ✓ 1296ms | ✓ 1451ms | ✓ 1311ms | 否 | 否 | http |
| 154.90.48.209:9090 | 否 | 否 | ✓ 916ms | ✓ 1303ms | ✓ 1096ms | http |
| 91.233.223.147:3128 | ✓ 910ms | 否 | ✓ 1847ms | 否 | ✓ 1916ms | http |
| 38.188.247.12:999 | ✓ 1751ms | 否 | ✓ 1557ms | ✓ 1386ms | ✓ 1367ms | http |
| 101.6.65.112:10080 | ✓ 1505ms | ✓ 1849ms | ✓ 1220ms | ✓ 1582ms | ✓ 1201ms | http |
| 3.101.133.120:80 | ✓ 310ms | ✓ 1190ms | ✓ 1315ms | ✓ 1103ms | ✓ 833ms | http |
| 80.92.204.47:1081 | ✓ 1396ms | 否 | ✓ 788ms | ✓ 1586ms | ✓ 1320ms | http |
| 139.159.97.82:10900 | ✓ 1169ms | 否 | ✓ 1362ms | ✓ 1624ms | ✓ 1240ms | http |
| 20.78.213.56:80 | ✓ 689ms | ✓ 1337ms | 否 | ✓ 1328ms | ✓ 941ms | http |
| 222.129.141.73:9000 | ✓ 1590ms | ✓ 1607ms | 否 | ✓ 1519ms | 否 | http |
| 54.229.201.146:5893 | ✓ 1477ms | 否 | ✓ 1315ms | 否 | ✓ 1845ms | http |
| 13.51.196.44:16963 | ✓ 1457ms | 否 | 否 | ✓ 1694ms | ✓ 1763ms | http |
| 103.35.190.69:1081 | ✓ 254ms | ✓ 1254ms | ✓ 103ms | ✓ 1206ms | ✓ 768ms | http |
| 144.31.25.69:21064 | ✓ 1793ms | 否 | ✓ 1168ms | 否 | ✓ 1922ms | http |
| 128.199.116.219:9090 | ✓ 1183ms | 否 | ✓ 1156ms | ✓ 1942ms | ✓ 1154ms | http |
| 13.48.13.125:39595 | ✓ 1518ms | 否 | ✓ 1131ms | 否 | ✓ 1438ms | http |
| 47.77.216.82:1080 | ✓ 390ms | ✓ 850ms | 否 | ✓ 1871ms | ✓ 754ms | http |
| 77.110.107.80:1080 | 否 | 否 | ✓ 1420ms | ✓ 1678ms | ✓ 1210ms | http |
| 20.164.75.153:8080 | ✓ 1240ms | 否 | ✓ 1991ms | 否 | ✓ 1836ms | http |
| 81.26.190.143:1080 | ✓ 1184ms | 否 | ✓ 1830ms | 否 | ✓ 1754ms | http |
| 152.42.177.32:8888 | ✓ 1076ms | 否 | ✓ 1103ms | ✓ 1452ms | ✓ 1453ms | http |
| 121.130.199.80:3128 | ✓ 1896ms | ✓ 1534ms | ✓ 1884ms | 否 | ✓ 1293ms | http |
| 121.230.8.213:1080 | ✓ 1334ms | ✓ 1915ms | ✓ 1663ms | 否 | ✓ 1471ms | http |
| 103.18.77.24:1111 | ✓ 1880ms | 否 | ✓ 1394ms | ✓ 1553ms | ✓ 1536ms | http |
| 51.79.207.21:8080 | 否 | 否 | ✓ 1757ms | ✓ 1853ms | ✓ 1144ms | http |
| 61.52.131.172:8443 | ✓ 1014ms | ✓ 1327ms | ✓ 1132ms | ✓ 1365ms | ✓ 1091ms | http |
| 137.59.47.73:3128 | ✓ 1554ms | ✓ 1609ms | ✓ 1555ms | ✓ 1244ms | ✓ 1305ms | http |
| 80.92.204.47:1082 | ✓ 1074ms | ✓ 1326ms | ✓ 778ms | ✓ 1704ms | 否 | http |
| 157.230.220.25:4857 | 否 | 否 | ✓ 1051ms | ✓ 1418ms | ✓ 993ms | http |
| 103.39.51.207:8080 | ✓ 1938ms | 否 | ✓ 1385ms | ✓ 1873ms | 否 | http |

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
