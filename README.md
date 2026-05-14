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

最后更新：2026-05-14 20:44:21 UTC（2026-05-15 04:44:21 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1956ms | ✓ 1200ms | ✓ 1720ms | ✓ 1665ms | ✓ 1290ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1387ms | ✓ 1123ms | ✓ 1508ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1840ms | 否 | ✓ 1063ms | 否 | ✓ 1140ms | http |
| 212.58.132.5:8888 | ✓ 1769ms | 否 | ✓ 1511ms | ✓ 1531ms | ✓ 1288ms | http |
| 218.108.131.186:17890 | ✓ 1977ms | 否 | ✓ 1302ms | ✓ 1839ms | ✓ 1236ms | http |
| 129.80.217.21:444 | ✓ 1587ms | ✓ 1021ms | ✓ 1206ms | ✓ 886ms | ✓ 900ms | http |
| 59.46.216.131:30001 | ✓ 1286ms | ✓ 1687ms | ✓ 1301ms | ✓ 1651ms | 否 | http |
| 38.75.82.221:999 | ✓ 1639ms | ✓ 1723ms | ✓ 1244ms | 否 | 否 | http |
| 34.96.238.40:8080 | ✓ 1116ms | 否 | ✓ 1214ms | ✓ 1810ms | 否 | http |
| 160.238.65.6:3128 | 否 | ✓ 1364ms | 否 | ✓ 1644ms | ✓ 1141ms | http |
| 160.238.65.9:3128 | ✓ 965ms | 否 | ✓ 558ms | ✓ 1629ms | 否 | http |
| 173.212.245.136:8888 | ✓ 1893ms | 否 | ✓ 1502ms | ✓ 1917ms | ✓ 1558ms | http |
| 168.222.254.136:8888 | ✓ 1041ms | ✓ 1724ms | 否 | 否 | ✓ 1631ms | http |
| 128.199.121.61:9090 | ✓ 1842ms | 否 | 否 | ✓ 1443ms | ✓ 1073ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 1511ms | ✓ 1449ms | ✓ 1075ms | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 1413ms | ✓ 1377ms | ✓ 1053ms | http |
| 57.129.144.178:40000 | ✓ 1040ms | 否 | ✓ 1013ms | ✓ 1685ms | ✓ 1435ms | http |
| 103.147.152.12:1095 | ✓ 855ms | ✓ 1414ms | ✓ 976ms | ✓ 1429ms | 否 | http |
| 128.199.114.189:9090 | ✓ 1837ms | 否 | ✓ 940ms | ✓ 1307ms | ✓ 1036ms | http |
| 128.199.113.85:9090 | 否 | 否 | ✓ 1048ms | ✓ 1327ms | ✓ 1041ms | http |
| 146.190.80.158:9090 | ✓ 1835ms | 否 | ✓ 1098ms | ✓ 1349ms | ✓ 1159ms | http |
| 84.47.150.125:1080 | ✓ 1072ms | 否 | 否 | ✓ 1940ms | ✓ 1721ms | http |
| 45.125.67.37:8443 | ✓ 1045ms | 否 | ✓ 1528ms | ✓ 1980ms | ✓ 1361ms | http |
| 34.101.184.164:3128 | ✓ 1904ms | 否 | ✓ 1363ms | 否 | ✓ 1196ms | http |
| 103.147.152.12:1080 | ✓ 850ms | 否 | ✓ 806ms | ✓ 1451ms | 否 | http |
| 5.129.248.58:3128 | ✓ 1387ms | ✓ 1972ms | ✓ 1468ms | 否 | ✓ 1409ms | http |
| 43.167.192.85:8080 | 否 | 否 | ✓ 1568ms | ✓ 1760ms | ✓ 1340ms | http |
| 91.242.229.129:8092 | ✓ 1103ms | ✓ 1381ms | 否 | 否 | ✓ 1167ms | http |
| 160.238.65.4:3128 | ✓ 753ms | ✓ 1704ms | 否 | 否 | ✓ 1517ms | http |
| 217.182.195.221:30000 | ✓ 1328ms | 否 | ✓ 1776ms | 否 | ✓ 1481ms | http |
| 212.224.88.212:443 | 否 | ✓ 1678ms | ✓ 1454ms | 否 | ✓ 1559ms | http |
| 8.154.21.175:3128 | ✓ 1011ms | ✓ 1323ms | ✓ 1035ms | ✓ 1399ms | ✓ 1164ms | http |
| 120.79.99.232:8099 | ✓ 1294ms | ✓ 1687ms | ✓ 1435ms | ✓ 1596ms | ✓ 1319ms | http |
| 160.238.65.5:3128 | 否 | 否 | ✓ 1491ms | ✓ 1882ms | ✓ 1005ms | http |
| 185.230.191.240:3128 | ✓ 1801ms | ✓ 1692ms | ✓ 1088ms | 否 | ✓ 1905ms | http |
| 121.177.17.179:3116 | 否 | 否 | ✓ 1149ms | ✓ 1728ms | ✓ 1877ms | http |
| 212.34.146.118:3128 | ✓ 748ms | 否 | 否 | ✓ 1799ms | ✓ 1501ms | http |
| 20.164.75.153:8080 | ✓ 1196ms | 否 | ✓ 1886ms | 否 | ✓ 1756ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1829ms | ✓ 1348ms | ✓ 1462ms | ✓ 1102ms | http |
| 160.238.65.2:3128 | ✓ 1499ms | 否 | ✓ 1009ms | 否 | ✓ 1076ms | http |
| 121.230.8.91:1080 | ✓ 1450ms | ✓ 1816ms | ✓ 1536ms | 否 | ✓ 1224ms | http |
| 45.173.12.140:1994 | 否 | ✓ 1746ms | ✓ 1129ms | ✓ 1828ms | ✓ 1337ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1249ms | ✓ 1389ms | ✓ 1646ms | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 205ms | ✓ 1357ms | ✓ 1175ms | http |
| 45.146.243.133:1080 | ✓ 890ms | 否 | 否 | ✓ 1444ms | ✓ 864ms | http |
| 49.156.44.116:8080 | ✓ 1777ms | 否 | ✓ 1594ms | ✓ 1729ms | ✓ 1738ms | http |
| 5.75.139.30:1081 | ✓ 1128ms | 否 | ✓ 1676ms | 否 | ✓ 1887ms | http |
| 103.172.42.175:8082 | ✓ 1597ms | 否 | 否 | ✓ 1681ms | ✓ 1644ms | http |
| 129.212.224.122:3128 | ✓ 961ms | 否 | ✓ 966ms | ✓ 1303ms | 否 | http |
| 34.71.229.255:3128 | ✓ 500ms | ✓ 1505ms | 否 | ✓ 1373ms | ✓ 1281ms | http |
| 120.92.212.16:7890 | ✓ 1136ms | ✓ 1547ms | 否 | 否 | ✓ 1247ms | http |
| 203.150.113.253:8080 | 否 | 否 | ✓ 1922ms | ✓ 1923ms | ✓ 1586ms | http |
| 77.110.119.136:3128 | 否 | ✓ 1457ms | ✓ 215ms | ✓ 1357ms | 否 | http |
| 222.107.27.7:8064 | 否 | ✓ 1539ms | ✓ 834ms | ✓ 1433ms | ✓ 1129ms | http |
| 190.12.150.244:999 | ✓ 1449ms | ✓ 1761ms | 否 | ✓ 1669ms | ✓ 1428ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 849ms | ✓ 1074ms | ✓ 1846ms | http |
| 160.238.65.8:3128 | ✓ 1863ms | ✓ 1927ms | ✓ 1384ms | 否 | ✓ 1649ms | http |
| 103.35.191.244:1081 | ✓ 449ms | ✓ 1139ms | ✓ 1845ms | 否 | ✓ 897ms | http |
| 119.28.51.157:3128 | 否 | ✓ 1396ms | ✓ 1468ms | ✓ 1632ms | ✓ 1314ms | http |
| 104.248.151.93:9090 | ✓ 976ms | 否 | ✓ 944ms | ✓ 1285ms | ✓ 1011ms | http |
| 158.178.237.243:3128 | ✓ 1033ms | 否 | ✓ 1882ms | ✓ 1933ms | ✓ 1457ms | http |
| 62.234.206.73:3128 | ✓ 1158ms | ✓ 1496ms | ✓ 1322ms | ✓ 1621ms | ✓ 1255ms | http |
| 101.32.243.189:80 | 否 | 否 | ✓ 1901ms | ✓ 1528ms | ✓ 1628ms | http |
| 47.74.226.8:5001 | 否 | ✓ 1755ms | ✓ 1322ms | ✓ 1759ms | 否 | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1234ms | ✓ 1174ms | ✓ 926ms | http |
| 3.101.133.120:80 | ✓ 1530ms | ✓ 1554ms | 否 | ✓ 1394ms | ✓ 969ms | http |
| 113.176.92.71:3128 | ✓ 1261ms | ✓ 1540ms | ✓ 1581ms | ✓ 1779ms | ✓ 1367ms | http |
| 152.32.132.190:7890 | ✓ 1331ms | ✓ 1686ms | 否 | 否 | ✓ 1131ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1094ms | ✓ 1513ms | ✓ 1428ms | http |
| 45.59.122.132:80 | ✓ 917ms | ✓ 1456ms | ✓ 1495ms | ✓ 1494ms | ✓ 1213ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1888ms | ✓ 1842ms | ✓ 1828ms | http |
| 45.129.141.143:3128 | ✓ 998ms | ✓ 1830ms | ✓ 1544ms | ✓ 1947ms | ✓ 1592ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1422ms | ✓ 1098ms | ✓ 1548ms | ✓ 1084ms | http |
| 185.21.15.206:3128 | ✓ 1258ms | ✓ 1450ms | ✓ 1080ms | ✓ 1962ms | 否 | http |
| 128.199.254.13:9090 | ✓ 1780ms | 否 | ✓ 998ms | ✓ 1307ms | 否 | http |
| 160.238.65.7:3128 | 否 | ✓ 1348ms | ✓ 524ms | 否 | ✓ 997ms | http |
| 80.92.204.47:1082 | ✓ 795ms | 否 | ✓ 676ms | ✓ 1963ms | ✓ 1343ms | http |
| 38.211.245.131:999 | ✓ 1711ms | 否 | ✓ 1054ms | 否 | ✓ 1903ms | http |
| 158.160.215.167:8123 | ✓ 1238ms | 否 | ✓ 1889ms | 否 | ✓ 1460ms | http |
| 160.238.65.3:3128 | ✓ 1466ms | ✓ 1322ms | ✓ 1855ms | 否 | 否 | http |
| 38.211.245.18:999 | ✓ 877ms | 否 | ✓ 984ms | 否 | ✓ 1905ms | http |
| 202.154.18.80:8082 | ✓ 1764ms | 否 | 否 | ✓ 1789ms | ✓ 1797ms | http |
| 61.52.131.172:8443 | ✓ 1112ms | 否 | 否 | ✓ 1420ms | ✓ 1214ms | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1573ms | ✓ 1644ms | ✓ 1672ms | http |
| 103.158.242.58:83 | ✓ 1990ms | 否 | ✓ 1972ms | ✓ 1628ms | ✓ 1841ms | http |
| 103.157.200.126:3128 | ✓ 1639ms | 否 | 否 | ✓ 1560ms | ✓ 1234ms | http |

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
