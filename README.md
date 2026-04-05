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

最后更新：2026-04-05 08:08:40 UTC（2026-04-05 16:08:40 UTC+8）

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
| 147.161.210.140:8800 | ✓ 633ms | 否 | ✓ 887ms | ✓ 1004ms | ✓ 926ms | http |
| 1.231.81.166:3128 | ✓ 822ms | ✓ 929ms | ✓ 1803ms | ✓ 1152ms | ✓ 853ms | http |
| 111.227.254.11:22222 | ✓ 1037ms | ✓ 1339ms | ✓ 1019ms | ✓ 1349ms | ✓ 1062ms | http |
| 167.103.115.102:8800 | ✓ 1026ms | 否 | ✓ 1195ms | ✓ 1381ms | ✓ 1052ms | http |
| 159.223.71.162:8080 | ✓ 861ms | 否 | 否 | ✓ 1006ms | ✓ 967ms | http |
| 159.223.71.162:443 | ✓ 767ms | 否 | 否 | ✓ 1024ms | ✓ 1056ms | http |
| 147.161.239.240:8800 | ✓ 1233ms | ✓ 1846ms | ✓ 1172ms | ✓ 1891ms | ✓ 1632ms | http |
| 113.160.132.26:8080 | ✓ 1689ms | 否 | ✓ 1036ms | ✓ 1666ms | ✓ 1131ms | http |
| 209.38.154.7:1080 | 否 | 否 | ✓ 1230ms | ✓ 678ms | ✓ 677ms | http |
| 34.101.184.164:3128 | ✓ 1745ms | 否 | ✓ 1435ms | ✓ 1200ms | ✓ 1351ms | http |
| 111.227.254.9:22222 | ✓ 1016ms | ✓ 1328ms | 否 | ✓ 1374ms | ✓ 1032ms | http |
| 95.213.217.168:52004 | ✓ 1271ms | 否 | ✓ 1511ms | 否 | ✓ 1985ms | http |
| 86.53.183.16:1080 | ✓ 1239ms | 否 | ✓ 947ms | 否 | ✓ 1882ms | http |
| 167.103.34.108:8800 | ✓ 1656ms | 否 | ✓ 1630ms | 否 | ✓ 1554ms | http |
| 133.242.138.34:8100 | ✓ 908ms | 否 | 否 | ✓ 1215ms | ✓ 1773ms | http |
| 111.227.254.12:22222 | ✓ 1040ms | ✓ 1317ms | 否 | 否 | ✓ 1719ms | http |
| 38.34.179.36:8453 | ✓ 1736ms | ✓ 1565ms | ✓ 1852ms | ✓ 1153ms | ✓ 1994ms | http |
| 168.110.52.228:3128 | ✓ 589ms | 否 | ✓ 1327ms | ✓ 960ms | ✓ 1137ms | http |
| 43.167.237.94:3128 | ✓ 1634ms | 否 | ✓ 442ms | ✓ 756ms | ✓ 1904ms | http |
| 167.103.144.127:8800 | ✓ 1572ms | 否 | ✓ 1267ms | ✓ 1554ms | ✓ 1573ms | http |
| 45.167.124.52:8080 | ✓ 755ms | 否 | ✓ 710ms | ✓ 1746ms | ✓ 1491ms | http |
| 201.144.20.238:3128 | ✓ 733ms | ✓ 1775ms | ✓ 1030ms | 否 | 否 | http |
| 185.191.236.162:3128 | ✓ 1245ms | 否 | ✓ 998ms | ✓ 1737ms | 否 | http |
| 198.59.68.130:3128 | ✓ 1916ms | ✓ 1835ms | ✓ 1519ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1607ms | ✓ 1283ms | ✓ 1257ms | ✓ 1036ms | http |
| 8.219.97.248:80 | ✓ 1370ms | 否 | ✓ 1473ms | ✓ 1377ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1993ms | 否 | ✓ 1388ms | ✓ 1253ms | 否 | http |
| 45.12.151.226:2829 | ✓ 1557ms | 否 | ✓ 1299ms | ✓ 1861ms | ✓ 1428ms | http |
| 167.103.31.122:8800 | ✓ 1395ms | 否 | ✓ 1365ms | ✓ 1684ms | ✓ 1640ms | http |
| 38.145.220.188:8450 | ✓ 891ms | 否 | ✓ 1077ms | ✓ 954ms | 否 | http |
| 38.145.220.182:8450 | ✓ 890ms | ✓ 1600ms | ✓ 1479ms | ✓ 966ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1524ms | ✓ 1093ms | ✓ 1627ms | ✓ 1653ms | ✓ 1127ms | http |
| 45.167.125.21:999 | ✓ 731ms | 否 | ✓ 732ms | ✓ 1901ms | ✓ 1533ms | http |
| 157.20.252.154:1111 | 否 | 否 | ✓ 1441ms | ✓ 1682ms | ✓ 1401ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1214ms | ✓ 1160ms | ✓ 1492ms | http |
| 38.145.208.182:8452 | ✓ 527ms | ✓ 829ms | ✓ 300ms | ✓ 1174ms | ✓ 1330ms | http |
| 111.227.254.10:22222 | ✓ 975ms | ✓ 1655ms | ✓ 1527ms | 否 | ✓ 1058ms | http |
| 38.145.220.39:8449 | ✓ 1312ms | 否 | ✓ 357ms | ✓ 1229ms | 否 | http |
| 38.145.220.40:8449 | ✓ 1316ms | 否 | ✓ 356ms | ✓ 1232ms | 否 | http |
| 208.87.243.199:7878 | ✓ 914ms | ✓ 1776ms | ✓ 388ms | ✓ 753ms | ✓ 593ms | http |
| 65.108.203.35:18080 | ✓ 918ms | 否 | ✓ 1840ms | 否 | ✓ 1767ms | http |
| 65.108.203.37:18080 | ✓ 949ms | 否 | ✓ 1819ms | 否 | ✓ 1421ms | http |
| 120.28.216.197:8082 | ✓ 1383ms | 否 | ✓ 1461ms | ✓ 1488ms | 否 | http |
| 210.223.44.230:3128 | ✓ 676ms | ✓ 1996ms | ✓ 1062ms | 否 | ✓ 1702ms | http |
| 64.227.76.27:1080 | ✓ 1889ms | ✓ 1862ms | ✓ 695ms | ✓ 1740ms | 否 | http |
| 116.80.49.168:3172 | ✓ 1626ms | 否 | 否 | ✓ 1794ms | ✓ 1639ms | http |
| 62.113.119.14:8080 | ✓ 827ms | ✓ 1691ms | ✓ 850ms | 否 | 否 | http |
| 182.204.177.139:1080 | ✓ 1203ms | 否 | ✓ 1149ms | ✓ 1459ms | ✓ 1120ms | http |
| 106.10.55.212:1121 | ✓ 1339ms | ✓ 1237ms | ✓ 1238ms | 否 | ✓ 1086ms | http |
| 120.92.212.16:8890 | ✓ 1529ms | ✓ 1386ms | 否 | ✓ 1443ms | ✓ 960ms | http |
| 120.92.212.16:7890 | ✓ 1518ms | ✓ 1378ms | 否 | ✓ 1202ms | ✓ 959ms | http |
| 47.74.226.8:5001 | ✓ 926ms | ✓ 1330ms | 否 | ✓ 1435ms | 否 | http |
| 103.219.160.145:6565 | 否 | 否 | ✓ 1945ms | ✓ 1919ms | ✓ 1803ms | http |
| 82.114.228.67:1080 | ✓ 870ms | 否 | ✓ 835ms | ✓ 1804ms | ✓ 1311ms | http |
| 89.208.106.138:10808 | ✓ 1307ms | 否 | ✓ 1298ms | ✓ 1497ms | ✓ 1938ms | http |
| 103.39.51.207:8080 | ✓ 1234ms | 否 | ✓ 1181ms | ✓ 1959ms | 否 | http |
| 103.69.84.106:3131 | 否 | 否 | ✓ 1143ms | ✓ 1115ms | ✓ 875ms | http |
| 123.57.0.163:8888 | ✓ 1357ms | ✓ 1411ms | ✓ 1384ms | 否 | 否 | http |
| 121.230.8.111:1080 | ✓ 1154ms | 否 | ✓ 1816ms | ✓ 1583ms | ✓ 1985ms | http |
| 177.234.217.88:999 | ✓ 1440ms | ✓ 1967ms | ✓ 1680ms | 否 | ✓ 1795ms | http |
| 115.231.181.40:8128 | ✓ 1199ms | ✓ 1822ms | ✓ 896ms | 否 | ✓ 1855ms | http |
| 212.58.132.5:8888 | ✓ 1175ms | 否 | ✓ 1113ms | ✓ 1834ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1270ms | 否 | ✓ 1557ms | 否 | ✓ 1899ms | http |
| 38.145.220.102:8453 | ✓ 1154ms | ✓ 713ms | ✓ 356ms | ✓ 951ms | ✓ 1148ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1135ms | ✓ 556ms | ✓ 929ms | ✓ 666ms | http |
| 5.104.87.17:8050 | ✓ 1362ms | 否 | ✓ 1305ms | ✓ 1165ms | ✓ 962ms | http |
| 45.88.0.113:3128 | ✓ 1280ms | 否 | ✓ 1780ms | 否 | ✓ 1234ms | http |
| 45.88.0.117:3128 | ✓ 1309ms | ✓ 1701ms | ✓ 1097ms | 否 | 否 | http |
| 45.88.0.99:3128 | ✓ 1303ms | 否 | ✓ 1740ms | ✓ 1647ms | 否 | http |
| 218.108.131.186:17890 | ✓ 789ms | ✓ 1017ms | ✓ 782ms | ✓ 1029ms | ✓ 853ms | http |
| 93.77.181.116:8888 | ✓ 1022ms | 否 | ✓ 1352ms | ✓ 1989ms | ✓ 1580ms | http |
| 38.145.218.217:8450 | 否 | ✓ 1332ms | ✓ 112ms | ✓ 770ms | ✓ 1220ms | http |
| 38.145.218.13:8446 | ✓ 1970ms | ✓ 1063ms | ✓ 639ms | ✓ 1357ms | ✓ 1106ms | http |
| 45.136.131.39:8451 | ✓ 432ms | ✓ 1192ms | ✓ 1886ms | ✓ 840ms | ✓ 939ms | http |
| 38.145.220.175:8449 | ✓ 417ms | 否 | ✓ 1355ms | ✓ 731ms | ✓ 586ms | http |
| 168.222.254.26:8888 | ✓ 1196ms | 否 | ✓ 1737ms | 否 | ✓ 1893ms | http |
| 38.34.179.179:8449 | 否 | 否 | ✓ 1293ms | ✓ 1780ms | ✓ 1931ms | http |
| 197.164.101.11:1976 | ✓ 1375ms | 否 | ✓ 1567ms | ✓ 1858ms | ✓ 1800ms | http |
| 38.145.218.234:8447 | ✓ 842ms | 否 | ✓ 803ms | ✓ 1748ms | ✓ 1283ms | http |
| 37.187.109.70:10111 | ✓ 964ms | 否 | ✓ 1274ms | 否 | ✓ 1743ms | http |
| 45.88.0.115:3128 | ✓ 725ms | ✓ 1684ms | ✓ 723ms | 否 | 否 | http |
| 181.78.44.63:999 | ✓ 1233ms | 否 | ✓ 1425ms | ✓ 1591ms | ✓ 1570ms | http |
| 38.34.179.56:8450 | ✓ 562ms | 否 | ✓ 1503ms | ✓ 771ms | ✓ 1635ms | http |
| 38.145.220.8:8447 | 否 | 否 | ✓ 1790ms | ✓ 806ms | ✓ 632ms | http |
| 45.136.131.68:8446 | ✓ 499ms | 否 | ✓ 1351ms | ✓ 693ms | ✓ 938ms | http |
| 45.136.131.64:8453 | ✓ 520ms | ✓ 1373ms | ✓ 1677ms | ✓ 751ms | ✓ 792ms | http |
| 45.136.131.37:8447 | ✓ 587ms | ✓ 1140ms | ✓ 1565ms | ✓ 1315ms | ✓ 1598ms | http |
| 38.34.179.79:8451 | 否 | 否 | ✓ 1820ms | ✓ 1910ms | ✓ 514ms | http |
| 45.136.130.188:8449 | ✓ 1160ms | ✓ 688ms | ✓ 579ms | ✓ 1874ms | ✓ 1389ms | http |
| 103.67.46.225:3125 | ✓ 1810ms | 否 | 否 | ✓ 1579ms | ✓ 1641ms | http |
| 61.52.131.172:8443 | ✓ 862ms | ✓ 1103ms | ✓ 1050ms | ✓ 1136ms | ✓ 882ms | http |
| 109.234.38.35:3128 | ✓ 933ms | ✓ 1806ms | ✓ 1974ms | 否 | ✓ 1628ms | http |

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
