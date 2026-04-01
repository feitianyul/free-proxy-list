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

最后更新：2026-04-01 17:55:01 UTC（2026-04-02 01:55:01 UTC+8）

**代理总数：96**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 96 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 96 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 1127ms | ✓ 1501ms | ✓ 678ms | ✓ 1973ms | ✓ 1319ms | http |
| 147.161.210.140:8800 | ✓ 897ms | 否 | ✓ 906ms | ✓ 1107ms | ✓ 1137ms | http |
| 95.213.217.168:52004 | ✓ 1286ms | 否 | ✓ 1165ms | 否 | ✓ 1624ms | http |
| 167.103.115.102:8800 | ✓ 1160ms | 否 | ✓ 1207ms | ✓ 1889ms | ✓ 1551ms | http |
| 167.103.34.108:8800 | ✓ 1833ms | 否 | ✓ 1477ms | ✓ 1555ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1796ms | 否 | ✓ 1444ms | ✓ 1766ms | ✓ 1671ms | http |
| 45.167.124.52:8080 | ✓ 533ms | ✓ 1480ms | ✓ 710ms | 否 | 否 | http |
| 208.87.243.199:7878 | ✓ 1587ms | ✓ 1812ms | ✓ 1699ms | ✓ 1011ms | ✓ 1961ms | http |
| 167.103.31.122:8800 | ✓ 1341ms | 否 | ✓ 1298ms | ✓ 1989ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1618ms | ✓ 1595ms | ✓ 1112ms | 否 | ✓ 1132ms | http |
| 42.96.16.158:1311 | ✓ 1676ms | 否 | ✓ 1110ms | ✓ 1410ms | ✓ 1783ms | http |
| 35.225.22.61:80 | ✓ 935ms | ✓ 1138ms | ✓ 221ms | 否 | ✓ 835ms | http |
| 43.99.54.236:5555 | ✓ 885ms | ✓ 1190ms | ✓ 861ms | ✓ 1052ms | ✓ 839ms | http |
| 207.254.71.62:8088 | ✓ 545ms | ✓ 1555ms | ✓ 1374ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 984ms | ✓ 1206ms | ✓ 1749ms | ✓ 1340ms | ✓ 1148ms | http |
| 160.250.4.245:1 | ✓ 1736ms | 否 | ✓ 1668ms | ✓ 1757ms | 否 | http |
| 45.136.130.248:8452 | ✓ 1714ms | 否 | ✓ 1104ms | ✓ 1082ms | 否 | http |
| 160.250.5.22:1 | ✓ 1738ms | 否 | ✓ 1781ms | ✓ 1553ms | ✓ 1358ms | http |
| 38.145.208.222:8446 | ✓ 1727ms | 否 | ✓ 1650ms | ✓ 1083ms | ✓ 1554ms | http |
| 212.58.132.5:8888 | ✓ 1182ms | 否 | ✓ 1784ms | ✓ 1661ms | ✓ 1282ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1805ms | ✓ 1384ms | ✓ 1355ms | ✓ 1741ms | http |
| 177.234.217.88:999 | ✓ 1572ms | ✓ 1674ms | ✓ 1733ms | 否 | 否 | http |
| 203.80.138.81:50000 | ✓ 1156ms | ✓ 1404ms | 否 | ✓ 1479ms | ✓ 1946ms | http |
| 38.34.179.65:8450 | ✓ 1793ms | 否 | ✓ 1704ms | ✓ 1685ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1068ms | 否 | ✓ 910ms | ✓ 1953ms | ✓ 1495ms | http |
| 34.101.184.164:3128 | ✓ 1726ms | 否 | ✓ 1384ms | ✓ 1521ms | ✓ 1239ms | http |
| 150.241.71.15:1080 | ✓ 1111ms | 否 | ✓ 872ms | 否 | ✓ 1417ms | http |
| 45.136.198.40:3128 | ✓ 1111ms | ✓ 1403ms | ✓ 1773ms | ✓ 1876ms | ✓ 1618ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1459ms | ✓ 1646ms | 否 | ✓ 1237ms | http |
| 45.12.151.226:2829 | ✓ 744ms | 否 | ✓ 730ms | ✓ 1811ms | ✓ 1335ms | http |
| 120.92.212.16:7890 | ✓ 1220ms | ✓ 1660ms | 否 | 否 | ✓ 1164ms | http |
| 120.92.212.16:8890 | ✓ 1132ms | 否 | ✓ 1177ms | 否 | ✓ 1191ms | http |
| 1.12.62.237:8080 | ✓ 1886ms | ✓ 1865ms | 否 | ✓ 1818ms | 否 | http |
| 38.34.179.54:8446 | 否 | ✓ 1106ms | 否 | ✓ 1238ms | ✓ 1228ms | http |
| 138.197.68.35:4857 | ✓ 468ms | ✓ 1237ms | ✓ 227ms | ✓ 1332ms | ✓ 897ms | http |
| 186.116.148.52:8080 | 否 | ✓ 1838ms | 否 | ✓ 1728ms | ✓ 1684ms | http |
| 38.188.247.12:999 | ✓ 1895ms | ✓ 1811ms | ✓ 1495ms | 否 | 否 | http |
| 133.242.138.34:8100 | ✓ 1998ms | ✓ 1652ms | ✓ 1586ms | ✓ 1530ms | 否 | http |
| 190.12.150.244:999 | ✓ 1260ms | ✓ 1811ms | ✓ 838ms | 否 | 否 | http |
| 38.34.179.179:8449 | ✓ 335ms | ✓ 1207ms | ✓ 341ms | ✓ 1332ms | ✓ 722ms | http |
| 38.145.208.172:8448 | ✓ 461ms | ✓ 1077ms | ✓ 841ms | ✓ 911ms | ✓ 695ms | http |
| 38.34.179.91:8444 | ✓ 835ms | ✓ 1967ms | ✓ 361ms | ✓ 1000ms | ✓ 820ms | http |
| 38.145.208.220:8448 | ✓ 1035ms | ✓ 1686ms | ✓ 303ms | ✓ 946ms | 否 | http |
| 45.136.130.169:8446 | ✓ 511ms | ✓ 1642ms | ✓ 896ms | ✓ 908ms | ✓ 960ms | http |
| 38.145.208.227:8451 | ✓ 575ms | ✓ 1335ms | ✓ 1303ms | ✓ 1037ms | ✓ 1133ms | http |
| 38.145.220.9:8448 | ✓ 542ms | 否 | ✓ 770ms | ✓ 1091ms | ✓ 1246ms | http |
| 38.145.208.210:8448 | ✓ 342ms | ✓ 927ms | ✓ 1375ms | ✓ 1870ms | ✓ 797ms | http |
| 38.145.208.206:8448 | ✓ 344ms | ✓ 977ms | ✓ 1603ms | ✓ 1635ms | ✓ 767ms | http |
| 38.145.220.43:8449 | ✓ 283ms | ✓ 1620ms | 否 | ✓ 1045ms | ✓ 1446ms | http |
| 45.136.130.177:8448 | 否 | 否 | ✓ 809ms | ✓ 954ms | ✓ 1307ms | http |
| 38.145.203.97:8453 | ✓ 290ms | ✓ 939ms | 否 | ✓ 1710ms | ✓ 909ms | http |
| 38.34.179.155:8448 | ✓ 915ms | ✓ 960ms | ✓ 483ms | 否 | ✓ 1175ms | http |
| 38.34.179.202:8449 | ✓ 686ms | ✓ 947ms | 否 | ✓ 1782ms | ✓ 1550ms | http |
| 38.34.179.66:8444 | ✓ 1163ms | ✓ 896ms | ✓ 1417ms | 否 | ✓ 1016ms | http |
| 45.136.131.34:8451 | ✓ 1732ms | ✓ 926ms | ✓ 1296ms | 否 | ✓ 744ms | http |
| 38.145.218.229:8450 | ✓ 791ms | 否 | ✓ 983ms | 否 | ✓ 1161ms | http |
| 38.34.179.84:8453 | 否 | ✓ 1915ms | ✓ 1633ms | ✓ 1409ms | ✓ 726ms | http |
| 38.145.208.240:8448 | ✓ 1155ms | 否 | ✓ 1940ms | 否 | ✓ 785ms | http |
| 38.145.208.246:8450 | 否 | ✓ 1376ms | ✓ 1704ms | 否 | ✓ 1534ms | http |
| 113.255.59.226:8080 | ✓ 1463ms | 否 | ✓ 1539ms | ✓ 1346ms | ✓ 1698ms | http |
| 106.117.208.101:7890 | ✓ 1596ms | 否 | ✓ 1496ms | 否 | ✓ 1362ms | http |
| 5.102.109.41:999 | ✓ 688ms | ✓ 1302ms | 否 | ✓ 1372ms | ✓ 1135ms | http |
| 185.114.73.2:1080 | ✓ 709ms | ✓ 1454ms | 否 | ✓ 1111ms | ✓ 1429ms | http |
| 140.245.66.105:8081 | ✓ 1063ms | ✓ 1529ms | ✓ 1860ms | ✓ 1389ms | ✓ 1233ms | http |
| 103.184.99.194:8080 | ✓ 1830ms | 否 | ✓ 1916ms | ✓ 1684ms | 否 | http |
| 194.67.99.223:1080 | ✓ 1267ms | ✓ 1830ms | 否 | ✓ 1892ms | ✓ 1504ms | http |
| 210.223.44.230:3128 | ✓ 1024ms | ✓ 1363ms | ✓ 1162ms | ✓ 1269ms | ✓ 1177ms | http |
| 181.78.44.63:999 | ✓ 1049ms | ✓ 1508ms | 否 | ✓ 1697ms | ✓ 1063ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1539ms | ✓ 1699ms | ✓ 1633ms | http |
| 101.34.21.55:90 | ✓ 1107ms | ✓ 1635ms | 否 | ✓ 1256ms | ✓ 1151ms | http |
| 34.141.27.50:3128 | ✓ 830ms | ✓ 1484ms | ✓ 896ms | ✓ 1622ms | ✓ 1375ms | http |
| 178.156.224.42:3128 | ✓ 1082ms | 否 | ✓ 1576ms | 否 | ✓ 1826ms | http |
| 38.34.179.27:8451 | ✓ 1850ms | ✓ 902ms | ✓ 528ms | ✓ 1521ms | ✓ 1618ms | http |
| 101.132.61.121:8888 | ✓ 1503ms | ✓ 1503ms | ✓ 1527ms | ✓ 1728ms | ✓ 1523ms | http |
| 142.171.157.207:3128 | ✓ 626ms | ✓ 970ms | ✓ 1463ms | ✓ 892ms | ✓ 688ms | http |
| 59.8.203.55:80 | 否 | ✓ 1499ms | ✓ 1305ms | ✓ 1208ms | ✓ 951ms | http |
| 61.82.227.63:3172 | ✓ 1657ms | 否 | ✓ 1944ms | ✓ 1839ms | ✓ 1834ms | http |
| 202.165.92.206:8080 | ✓ 1675ms | 否 | ✓ 1735ms | 否 | ✓ 1971ms | http |
| 217.217.249.160:8080 | ✓ 1303ms | 否 | ✓ 1675ms | 否 | ✓ 1911ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1865ms | 否 | ✓ 1638ms | ✓ 1459ms | http |
| 38.34.183.164:8444 | ✓ 694ms | ✓ 911ms | ✓ 358ms | ✓ 952ms | ✓ 745ms | http |
| 38.34.179.61:8445 | ✓ 1660ms | ✓ 966ms | ✓ 752ms | 否 | ✓ 878ms | http |
| 38.34.183.234:8450 | ✓ 1343ms | ✓ 967ms | ✓ 421ms | ✓ 1040ms | ✓ 1134ms | http |
| 38.34.183.233:8448 | ✓ 909ms | ✓ 1095ms | ✓ 572ms | ✓ 1166ms | ✓ 1342ms | http |
| 38.145.220.102:8445 | ✓ 448ms | ✓ 1315ms | ✓ 513ms | ✓ 908ms | ✓ 820ms | http |
| 45.136.131.64:8446 | ✓ 980ms | ✓ 1853ms | ✓ 288ms | ✓ 953ms | ✓ 774ms | http |
| 38.34.179.161:8448 | ✓ 630ms | ✓ 1580ms | ✓ 624ms | ✓ 952ms | ✓ 1219ms | http |
| 45.136.131.30:8451 | ✓ 434ms | ✓ 1376ms | ✓ 1670ms | ✓ 1320ms | ✓ 1471ms | http |
| 45.136.131.52:8452 | ✓ 1031ms | ✓ 890ms | ✓ 956ms | ✓ 1552ms | ✓ 1361ms | http |
| 45.136.131.50:8452 | ✓ 1020ms | ✓ 858ms | ✓ 996ms | ✓ 1584ms | ✓ 1342ms | http |
| 38.145.203.86:8446 | ✓ 854ms | ✓ 1913ms | ✓ 521ms | ✓ 1443ms | 否 | http |
| 121.230.8.34:1080 | ✓ 1813ms | 否 | ✓ 1454ms | ✓ 1707ms | ✓ 1796ms | http |
| 223.16.170.103:80 | ✓ 1705ms | 否 | ✓ 1460ms | 否 | ✓ 1375ms | http |
| 124.6.44.250:65535 | ✓ 266ms | ✓ 1274ms | ✓ 487ms | ✓ 963ms | ✓ 1006ms | http |
| 38.35.247.157:999 | ✓ 1715ms | ✓ 1344ms | 否 | ✓ 1185ms | 否 | http |
| 223.16.170.103:3128 | ✓ 1363ms | ✓ 1915ms | ✓ 1283ms | 否 | ✓ 1369ms | http |

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
