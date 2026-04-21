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

最后更新：2026-04-21 10:39:06 UTC（2026-04-21 18:39:06 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:8080 | ✓ 1015ms | ✓ 1985ms | 否 | ✓ 1372ms | 否 | http |
| 149.51.42.10:3128 | ✓ 1015ms | ✓ 1948ms | 否 | ✓ 1440ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1209ms | 否 | 否 | ✓ 1380ms | ✓ 1572ms | http |
| 46.101.95.183:8888 | ✓ 1045ms | 否 | ✓ 1140ms | ✓ 1647ms | ✓ 1214ms | http |
| 89.208.106.138:10808 | ✓ 429ms | ✓ 1865ms | ✓ 1266ms | ✓ 1661ms | 否 | http |
| 162.19.253.202:8443 | ✓ 721ms | 否 | ✓ 1158ms | 否 | ✓ 1632ms | http |
| 1.231.81.166:3128 | ✓ 1642ms | 否 | ✓ 1025ms | ✓ 1165ms | ✓ 956ms | http |
| 188.246.224.49:7890 | ✓ 679ms | ✓ 1922ms | ✓ 1981ms | 否 | 否 | http |
| 152.42.208.139:8118 | ✓ 1699ms | 否 | ✓ 1410ms | ✓ 1244ms | ✓ 1083ms | http |
| 161.97.184.191:8080 | ✓ 1464ms | 否 | ✓ 1348ms | 否 | ✓ 1975ms | http |
| 113.160.132.26:8080 | ✓ 1813ms | ✓ 1650ms | ✓ 1439ms | ✓ 1422ms | ✓ 1610ms | http |
| 81.30.156.115:8080 | ✓ 742ms | ✓ 1723ms | ✓ 1375ms | 否 | 否 | http |
| 14.143.222.113:57788 | ✓ 1750ms | 否 | ✓ 1942ms | ✓ 1917ms | 否 | http |
| 35.225.22.61:80 | ✓ 572ms | ✓ 1840ms | 否 | ✓ 1335ms | 否 | http |
| 20.127.128.70:8080 | ✓ 551ms | 否 | ✓ 499ms | ✓ 1475ms | ✓ 1290ms | http |
| 103.113.70.189:1081 | ✓ 1469ms | ✓ 1249ms | ✓ 247ms | ✓ 1138ms | ✓ 1201ms | http |
| 108.181.201.118:1234 | 否 | 否 | ✓ 156ms | ✓ 1185ms | ✓ 1119ms | http |
| 45.76.207.177:40000 | ✓ 950ms | 否 | ✓ 1130ms | ✓ 1358ms | ✓ 1173ms | http |
| 91.99.15.45:2095 | ✓ 543ms | ✓ 1982ms | ✓ 1369ms | ✓ 1781ms | ✓ 1673ms | http |
| 84.47.150.125:1080 | ✓ 913ms | 否 | ✓ 1707ms | 否 | ✓ 1609ms | http |
| 85.190.99.143:443 | ✓ 1071ms | 否 | ✓ 1777ms | 否 | ✓ 1949ms | http |
| 185.138.116.150:8080 | ✓ 673ms | 否 | ✓ 991ms | 否 | ✓ 1502ms | http |
| 208.87.243.199:7878 | ✓ 611ms | ✓ 1180ms | ✓ 1174ms | ✓ 1707ms | ✓ 1298ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1991ms | ✓ 1042ms | ✓ 1617ms | ✓ 1994ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1991ms | ✓ 1439ms | ✓ 1592ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1577ms | ✓ 1534ms | ✓ 1308ms | http |
| 193.177.0.148:60000 | ✓ 902ms | 否 | ✓ 1298ms | ✓ 1517ms | ✓ 1689ms | http |
| 159.89.191.221:3128 | ✓ 1187ms | ✓ 1885ms | ✓ 709ms | 否 | ✓ 642ms | http |
| 177.93.132.244:3128 | ✓ 1347ms | 否 | ✓ 821ms | 否 | ✓ 1639ms | http |
| 152.32.132.190:7890 | ✓ 1159ms | 否 | ✓ 1051ms | 否 | ✓ 1404ms | http |
| 120.92.108.86:7890 | ✓ 1425ms | 否 | ✓ 1470ms | ✓ 1907ms | ✓ 1639ms | http |
| 45.153.231.229:8080 | ✓ 1229ms | 否 | ✓ 1290ms | 否 | ✓ 1899ms | http |
| 51.95.13.205:30771 | ✓ 1186ms | 否 | ✓ 1646ms | 否 | ✓ 1892ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1907ms | ✓ 1090ms | ✓ 1429ms | http |
| 152.70.91.193:40000 | ✓ 1906ms | 否 | ✓ 1800ms | ✓ 1711ms | ✓ 1806ms | http |
| 116.171.106.26:3443 | ✓ 1631ms | ✓ 1635ms | ✓ 1701ms | 否 | ✓ 1796ms | http |
| 157.230.178.216:8088 | ✓ 1324ms | ✓ 1212ms | ✓ 1397ms | ✓ 1176ms | ✓ 824ms | http |
| 84.47.150.126:1080 | ✓ 984ms | ✓ 1650ms | ✓ 1590ms | 否 | ✓ 1560ms | http |
| 218.60.0.214:80 | ✓ 1538ms | ✓ 1789ms | 否 | 否 | ✓ 1210ms | http |
| 47.84.73.61:1080 | ✓ 1541ms | 否 | ✓ 1196ms | ✓ 1300ms | ✓ 1026ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1970ms | ✓ 1603ms | ✓ 1710ms | http |
| 223.84.151.86:30005 | 否 | ✓ 1290ms | ✓ 1871ms | ✓ 1873ms | ✓ 1812ms | http |
| 168.144.75.9:3128 | ✓ 939ms | 否 | ✓ 1389ms | ✓ 1917ms | 否 | http |
| 45.12.151.226:2829 | ✓ 975ms | ✓ 1520ms | ✓ 626ms | 否 | 否 | http |
| 42.101.8.101:8888 | ✓ 1681ms | ✓ 1790ms | ✓ 1523ms | 否 | ✓ 1344ms | http |
| 59.46.216.131:30001 | ✓ 1977ms | ✓ 1950ms | 否 | ✓ 1645ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1380ms | 否 | ✓ 1452ms | ✓ 1548ms | 否 | http |
| 188.253.125.38:28798 | ✓ 1386ms | ✓ 1429ms | ✓ 1568ms | ✓ 1505ms | ✓ 1214ms | http |
| 159.203.220.84:3128 | ✓ 800ms | ✓ 1098ms | ✓ 1841ms | ✓ 1127ms | ✓ 1065ms | http |
| 101.32.244.83:8080 | ✓ 1371ms | 否 | ✓ 1169ms | ✓ 1444ms | ✓ 1529ms | http |
| 121.43.196.210:8222 | ✓ 1157ms | ✓ 1318ms | ✓ 1110ms | ✓ 1319ms | ✓ 1125ms | http |
| 121.43.196.213:8222 | ✓ 1131ms | ✓ 1275ms | 否 | ✓ 1329ms | ✓ 1134ms | http |
| 114.55.226.123:10086 | ✓ 1496ms | ✓ 1639ms | ✓ 1272ms | ✓ 1524ms | ✓ 1249ms | http |
| 111.79.111.126:3128 | ✓ 1746ms | ✓ 1985ms | ✓ 1379ms | 否 | 否 | http |
| 192.3.248.190:8014 | ✓ 993ms | 否 | ✓ 888ms | ✓ 1521ms | ✓ 1042ms | http |
| 91.233.223.147:3128 | ✓ 852ms | 否 | ✓ 1979ms | 否 | ✓ 1560ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1724ms | ✓ 1374ms | ✓ 1142ms | http |
| 174.114.24.95:3128 | ✓ 675ms | ✓ 1250ms | ✓ 521ms | ✓ 1777ms | ✓ 1590ms | http |
| 91.193.240.157:9877 | ✓ 960ms | 否 | ✓ 988ms | 否 | ✓ 1430ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 605ms | ✓ 1472ms | ✓ 1093ms | http |
| 172.86.66.149:8080 | ✓ 711ms | ✓ 1698ms | ✓ 947ms | ✓ 1850ms | ✓ 1672ms | http |
| 43.132.188.134:443 | ✓ 1068ms | ✓ 1828ms | 否 | 否 | ✓ 1955ms | http |
| 210.223.44.230:3128 | ✓ 1138ms | ✓ 1221ms | ✓ 1221ms | ✓ 1260ms | 否 | http |
| 217.76.245.80:999 | ✓ 1278ms | ✓ 1574ms | ✓ 1048ms | ✓ 1391ms | ✓ 1242ms | http |
| 103.113.70.189:1082 | ✓ 553ms | 否 | ✓ 477ms | ✓ 1708ms | ✓ 1261ms | http |
| 144.31.52.77:3128 | ✓ 638ms | ✓ 1736ms | ✓ 943ms | 否 | ✓ 1702ms | http |
| 78.11.96.22:8888 | ✓ 785ms | ✓ 1485ms | ✓ 1144ms | 否 | 否 | http |
| 195.26.243.76:3128 | ✓ 1855ms | 否 | ✓ 1375ms | ✓ 1174ms | 否 | http |
| 65.108.203.35:18080 | ✓ 669ms | 否 | ✓ 1788ms | 否 | ✓ 1529ms | http |
| 65.108.203.37:18080 | ✓ 1723ms | 否 | ✓ 1795ms | ✓ 1750ms | 否 | http |
| 45.129.141.143:3128 | ✓ 717ms | ✓ 1865ms | ✓ 1863ms | ✓ 1955ms | ✓ 1926ms | http |
| 38.180.2.107:3128 | ✓ 739ms | ✓ 1834ms | ✓ 1966ms | 否 | 否 | http |
| 45.186.6.104:3128 | ✓ 1897ms | ✓ 1803ms | ✓ 1553ms | 否 | 否 | http |
| 2.63.162.206:3128 | ✓ 1057ms | 否 | ✓ 1030ms | 否 | ✓ 1925ms | http |
| 178.63.155.151:8888 | ✓ 752ms | 否 | ✓ 933ms | ✓ 1753ms | ✓ 1555ms | http |
| 103.56.115.156:7890 | ✓ 1448ms | 否 | ✓ 1078ms | ✓ 1108ms | ✓ 1289ms | http |
| 116.171.106.78:3443 | 否 | ✓ 1757ms | ✓ 1625ms | 否 | ✓ 1873ms | http |
| 52.163.56.148:80 | ✓ 994ms | ✓ 1536ms | ✓ 1248ms | ✓ 1442ms | ✓ 1570ms | http |
| 202.141.161.53:10808 | 否 | ✓ 1928ms | ✓ 1385ms | ✓ 1536ms | ✓ 1227ms | http |
| 91.107.124.215:3128 | ✓ 1336ms | 否 | ✓ 1626ms | 否 | ✓ 1669ms | http |
| 138.124.99.216:8888 | 否 | 否 | ✓ 1253ms | ✓ 1945ms | ✓ 1029ms | http |
| 120.92.212.16:7890 | ✓ 1258ms | ✓ 1901ms | 否 | ✓ 1528ms | 否 | http |
| 217.182.195.221:30000 | ✓ 1800ms | 否 | ✓ 1932ms | 否 | ✓ 1688ms | http |
| 121.230.8.220:1080 | 否 | ✓ 1480ms | ✓ 1372ms | 否 | ✓ 1358ms | http |
| 103.67.46.225:3125 | ✓ 1983ms | 否 | 否 | ✓ 1749ms | ✓ 1830ms | http |
| 61.52.131.172:8443 | ✓ 1033ms | ✓ 1319ms | ✓ 1053ms | ✓ 1412ms | ✓ 1115ms | http |
| 170.9.224.115:3128 | ✓ 407ms | 否 | 否 | ✓ 1294ms | ✓ 959ms | http |
| 5.104.87.17:8050 | 否 | 否 | ✓ 1779ms | ✓ 1395ms | ✓ 1510ms | http |
| 210.45.76.58:42992 | 否 | ✓ 1647ms | ✓ 1698ms | ✓ 1592ms | ✓ 1297ms | http |

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
