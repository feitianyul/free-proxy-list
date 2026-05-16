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

最后更新：2026-05-16 23:41:13 UTC（2026-05-17 07:41:13 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 212.58.132.5:8888 | ✓ 1652ms | 否 | ✓ 1565ms | ✓ 1543ms | ✓ 1139ms | http |
| 185.200.188.234:10001 | ✓ 1283ms | 否 | ✓ 1885ms | 否 | ✓ 1719ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1982ms | ✓ 1218ms | 否 | ✓ 1817ms | http |
| 168.110.52.228:3128 | ✓ 839ms | ✓ 1516ms | ✓ 1469ms | ✓ 1286ms | ✓ 1057ms | http |
| 157.22.231.253:1080 | ✓ 582ms | 否 | ✓ 1633ms | 否 | ✓ 1631ms | http |
| 94.241.169.176:1080 | ✓ 1102ms | ✓ 1771ms | ✓ 1786ms | ✓ 1992ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1155ms | ✓ 1657ms | 否 | 否 | ✓ 1325ms | http |
| 150.107.140.238:3128 | ✓ 964ms | 否 | ✓ 1030ms | ✓ 1373ms | ✓ 1066ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1198ms | ✓ 1155ms | ✓ 1509ms | http |
| 45.88.0.98:3128 | ✓ 855ms | ✓ 1853ms | 否 | 否 | ✓ 1271ms | http |
| 51.161.50.166:3128 | 否 | ✓ 1451ms | ✓ 826ms | ✓ 1775ms | 否 | http |
| 170.106.136.181:31002 | 否 | ✓ 1943ms | 否 | ✓ 907ms | ✓ 1467ms | http |
| 91.242.229.129:8092 | ✓ 1240ms | ✓ 1824ms | ✓ 1120ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1685ms | ✓ 1315ms | 否 | 否 | ✓ 1205ms | http |
| 103.21.220.141:3128 | ✓ 847ms | 否 | ✓ 874ms | ✓ 1090ms | ✓ 892ms | http |
| 129.212.224.122:3128 | ✓ 923ms | 否 | ✓ 1033ms | ✓ 1332ms | 否 | http |
| 8.154.21.175:3128 | ✓ 1067ms | ✓ 1302ms | ✓ 1082ms | ✓ 1393ms | ✓ 1141ms | http |
| 35.194.4.51:3128 | 否 | ✓ 1496ms | ✓ 944ms | 否 | ✓ 1320ms | http |
| 148.230.4.241:999 | ✓ 651ms | ✓ 1589ms | ✓ 781ms | ✓ 1596ms | ✓ 1381ms | http |
| 190.12.150.244:999 | ✓ 1252ms | ✓ 1987ms | ✓ 1799ms | ✓ 1586ms | ✓ 1395ms | http |
| 185.40.77.94:1080 | ✓ 1310ms | ✓ 1837ms | 否 | 否 | ✓ 1913ms | http |
| 137.184.0.30:3128 | ✓ 559ms | ✓ 1221ms | ✓ 1196ms | ✓ 1054ms | ✓ 774ms | http |
| 84.47.150.125:1080 | ✓ 1119ms | 否 | ✓ 1955ms | 否 | ✓ 1631ms | http |
| 158.160.215.167:8126 | ✓ 1114ms | ✓ 1826ms | ✓ 1878ms | 否 | 否 | http |
| 110.172.28.217:3128 | 否 | 否 | ✓ 1515ms | ✓ 1425ms | ✓ 1529ms | http |
| 218.108.131.186:17890 | ✓ 1039ms | ✓ 1321ms | ✓ 1101ms | ✓ 1367ms | ✓ 1087ms | http |
| 106.10.55.212:1121 | ✓ 961ms | ✓ 1473ms | ✓ 1683ms | ✓ 1287ms | 否 | http |
| 192.210.140.36:7913 | ✓ 350ms | ✓ 1297ms | ✓ 958ms | ✓ 1236ms | ✓ 1060ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1261ms | ✓ 867ms | ✓ 1428ms | ✓ 996ms | http |
| 196.204.248.140:8080 | ✓ 1390ms | ✓ 1729ms | ✓ 1009ms | ✓ 1909ms | ✓ 1564ms | http |
| 5.129.248.58:3128 | ✓ 987ms | 否 | 否 | ✓ 1474ms | ✓ 1635ms | http |
| 116.171.106.78:3443 | 否 | ✓ 1765ms | ✓ 1769ms | ✓ 1984ms | 否 | http |
| 180.125.216.109:8118 | ✓ 1153ms | 否 | 否 | ✓ 1450ms | ✓ 1118ms | http |
| 64.188.77.221:3128 | ✓ 1298ms | 否 | ✓ 640ms | ✓ 1526ms | ✓ 1429ms | http |
| 64.188.77.26:3128 | 否 | ✓ 1805ms | 否 | ✓ 1454ms | ✓ 1248ms | http |
| 14.242.161.183:5106 | ✓ 1690ms | 否 | ✓ 1509ms | ✓ 1827ms | 否 | http |
| 165.101.43.67:8080 | 否 | 否 | ✓ 1893ms | ✓ 1777ms | ✓ 1730ms | http |
| 128.199.114.189:9090 | ✓ 1040ms | 否 | 否 | ✓ 1412ms | ✓ 1100ms | http |
| 62.60.149.161:3128 | ✓ 1869ms | 否 | ✓ 1513ms | 否 | ✓ 1772ms | http |
| 159.89.31.62:8080 | ✓ 396ms | ✓ 1389ms | ✓ 1273ms | ✓ 1984ms | ✓ 1257ms | http |
| 157.0.142.246:10057 | ✓ 1226ms | ✓ 1492ms | 否 | ✓ 1522ms | 否 | http |
| 129.80.217.21:444 | ✓ 734ms | 否 | ✓ 168ms | 否 | ✓ 648ms | http |
| 159.223.41.216:9090 | ✓ 952ms | 否 | ✓ 1350ms | ✓ 1325ms | ✓ 1080ms | http |
| 91.233.223.147:3128 | ✓ 860ms | ✓ 1796ms | ✓ 1113ms | 否 | ✓ 1982ms | http |
| 152.42.177.32:8888 | ✓ 1182ms | 否 | ✓ 1226ms | ✓ 1604ms | ✓ 1580ms | http |
| 101.32.243.189:80 | ✓ 1506ms | ✓ 1693ms | ✓ 1884ms | ✓ 1754ms | ✓ 1927ms | http |
| 152.70.91.193:40000 | ✓ 1832ms | 否 | ✓ 1728ms | ✓ 1681ms | ✓ 1836ms | http |
| 121.130.177.28:8888 | ✓ 1587ms | 否 | ✓ 1696ms | ✓ 1689ms | 否 | http |
| 43.156.90.221:10808 | ✓ 1725ms | 否 | ✓ 1912ms | 否 | ✓ 1000ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1464ms | ✓ 1532ms | 否 | ✓ 1202ms | http |
| 3.101.133.120:80 | ✓ 1085ms | ✓ 1380ms | ✓ 1354ms | ✓ 1336ms | ✓ 1184ms | http |
| 190.97.242.233:999 | ✓ 1288ms | 否 | ✓ 608ms | ✓ 1706ms | ✓ 1642ms | http |
| 103.153.247.110:8080 | ✓ 1718ms | 否 | ✓ 1997ms | ✓ 1934ms | ✓ 1923ms | http |
| 120.92.212.16:8890 | ✓ 1195ms | ✓ 1500ms | ✓ 1135ms | 否 | ✓ 1555ms | http |
| 178.156.224.42:3128 | ✓ 1051ms | 否 | ✓ 1385ms | 否 | ✓ 1709ms | http |
| 1.231.81.166:3128 | ✓ 1473ms | ✓ 1551ms | ✓ 1607ms | ✓ 1946ms | ✓ 1657ms | http |
| 190.93.224.32:999 | 否 | 否 | ✓ 1983ms | ✓ 1940ms | ✓ 1694ms | http |
| 213.220.62.63:3128 | ✓ 931ms | 否 | ✓ 1972ms | ✓ 1536ms | 否 | http |
| 45.88.0.115:3128 | ✓ 925ms | 否 | ✓ 1971ms | ✓ 1821ms | 否 | http |
| 134.209.153.66:3128 | ✓ 1020ms | 否 | ✓ 1478ms | ✓ 1903ms | ✓ 1590ms | http |
| 146.56.164.121:3128 | ✓ 1861ms | 否 | ✓ 1678ms | ✓ 1628ms | ✓ 1076ms | http |
| 103.154.214.50:3128 | ✓ 1162ms | 否 | ✓ 996ms | ✓ 1441ms | ✓ 1135ms | http |
| 40.90.163.168:3128 | ✓ 1201ms | ✓ 1479ms | ✓ 1192ms | ✓ 1504ms | ✓ 1204ms | http |
| 139.162.46.62:3128 | ✓ 1471ms | 否 | ✓ 931ms | ✓ 1293ms | 否 | http |
| 77.110.107.80:1080 | 否 | ✓ 1682ms | ✓ 1879ms | ✓ 1799ms | ✓ 1631ms | http |
| 114.214.170.41:27890 | 否 | ✓ 1642ms | ✓ 1434ms | ✓ 1594ms | ✓ 1336ms | http |
| 222.127.55.155:8082 | ✓ 1997ms | 否 | ✓ 1626ms | 否 | ✓ 1642ms | http |
| 139.198.113.42:10023 | 否 | 否 | ✓ 1910ms | ✓ 1476ms | ✓ 1489ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1558ms | ✓ 1945ms | ✓ 1809ms | http |
| 193.160.209.58:1080 | ✓ 917ms | 否 | ✓ 1464ms | ✓ 1882ms | ✓ 1633ms | http |
| 185.230.191.240:3128 | ✓ 691ms | 否 | ✓ 938ms | ✓ 1871ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1446ms | ✓ 1496ms | ✓ 1253ms | 否 | 否 | http |
| 20.164.75.153:8080 | ✓ 1231ms | 否 | ✓ 1834ms | 否 | ✓ 1814ms | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1359ms | ✓ 1503ms | ✓ 1470ms | http |
| 152.42.170.187:9090 | ✓ 1158ms | 否 | ✓ 1248ms | ✓ 1881ms | ✓ 1189ms | http |
| 103.122.142.174:8080 | ✓ 1728ms | 否 | ✓ 1454ms | 否 | ✓ 1716ms | http |
| 61.52.131.172:8443 | ✓ 1128ms | ✓ 1363ms | ✓ 1088ms | ✓ 1478ms | ✓ 1154ms | http |
| 144.124.227.88:3128 | ✓ 1360ms | 否 | ✓ 1849ms | ✓ 1870ms | ✓ 1784ms | http |
| 103.147.152.12:1095 | ✓ 1569ms | 否 | ✓ 1500ms | 否 | ✓ 1520ms | http |

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
