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

最后更新：2026-02-25 14:06:32 UTC（2026-02-25 22:06:32 UTC+8）

**代理总数：93**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 109.107.181.253:1080 | ✓ 1328ms | ✓ 1958ms | 否 | 否 | ✓ 1906ms | http |
| 20.78.118.91:8561 | ✓ 674ms | ✓ 1126ms | ✓ 643ms | ✓ 994ms | ✓ 769ms | http |
| 20.27.15.49:8561 | ✓ 678ms | ✓ 1217ms | ✓ 787ms | ✓ 1031ms | ✓ 787ms | http |
| 51.81.46.174:3128 | ✓ 914ms | ✓ 1288ms | ✓ 1099ms | 否 | ✓ 1708ms | http |
| 172.86.92.68:31337 | ✓ 775ms | ✓ 1677ms | ✓ 903ms | ✓ 1438ms | ✓ 1355ms | http |
| 14.56.107.244:3128 | 否 | ✓ 1937ms | ✓ 1363ms | 否 | ✓ 1584ms | http |
| 61.72.110.54:3128 | ✓ 1013ms | 否 | ✓ 1037ms | 否 | ✓ 1722ms | http |
| 104.238.30.45:59741 | ✓ 1611ms | 否 | ✓ 1643ms | 否 | ✓ 1903ms | http |
| 217.217.254.94:8080 | 否 | 否 | ✓ 1916ms | ✓ 1686ms | ✓ 1524ms | http |
| 152.32.255.24:27197 | ✓ 1921ms | 否 | ✓ 1871ms | ✓ 1955ms | ✓ 1649ms | http |
| 104.238.30.91:63900 | ✓ 1720ms | 否 | ✓ 1803ms | 否 | ✓ 1903ms | http |
| 104.238.30.58:63744 | ✓ 1602ms | 否 | ✓ 1935ms | 否 | ✓ 1904ms | http |
| 104.238.30.38:59741 | ✓ 1792ms | 否 | ✓ 1743ms | 否 | ✓ 1998ms | http |
| 165.227.5.10:8888 | ✓ 512ms | 否 | ✓ 1749ms | ✓ 1224ms | 否 | http |
| 35.234.17.221:8080 | ✓ 1108ms | ✓ 1863ms | 否 | ✓ 1418ms | 否 | http |
| 72.56.59.23:61937 | ✓ 1559ms | 否 | ✓ 1771ms | 否 | ✓ 1811ms | http |
| 104.238.30.68:63744 | ✓ 1578ms | 否 | ✓ 1680ms | 否 | ✓ 1938ms | http |
| 104.238.30.86:63900 | ✓ 1565ms | 否 | ✓ 1680ms | 否 | ✓ 1903ms | http |
| 120.92.212.16:7890 | ✓ 1098ms | 否 | ✓ 1146ms | 否 | ✓ 1154ms | http |
| 120.92.212.16:8890 | ✓ 1145ms | ✓ 1705ms | 否 | ✓ 1728ms | 否 | http |
| 168.235.110.63:3128 | ✓ 498ms | ✓ 1268ms | ✓ 1069ms | ✓ 1009ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1776ms | 否 | 否 | ✓ 1163ms | ✓ 855ms | http |
| 104.238.30.39:59741 | ✓ 1618ms | 否 | ✓ 1675ms | 否 | ✓ 1935ms | http |
| 185.191.236.162:3128 | ✓ 507ms | 否 | ✓ 473ms | ✓ 1818ms | ✓ 1033ms | http |
| 72.56.59.17:61931 | ✓ 1456ms | 否 | ✓ 1614ms | 否 | ✓ 1808ms | http |
| 81.70.169.194:80 | ✓ 1248ms | 否 | ✓ 1176ms | ✓ 1475ms | 否 | http |
| 72.56.59.62:63133 | ✓ 1897ms | 否 | ✓ 1615ms | 否 | ✓ 1775ms | http |
| 72.56.50.17:59787 | ✓ 1941ms | 否 | ✓ 1743ms | 否 | ✓ 1807ms | http |
| 72.56.59.56:63127 | ✓ 1896ms | 否 | ✓ 1614ms | 否 | ✓ 1779ms | http |
| 104.238.30.40:59741 | ✓ 1600ms | 否 | ✓ 1935ms | 否 | ✓ 1939ms | http |
| 101.43.255.96:80 | ✓ 1580ms | ✓ 1844ms | ✓ 1912ms | ✓ 1563ms | ✓ 1573ms | http |
| 61.72.110.24:3128 | ✓ 1601ms | 否 | ✓ 1211ms | 否 | ✓ 1000ms | http |
| 115.231.181.40:8128 | ✓ 1139ms | 否 | 否 | ✓ 1586ms | ✓ 1935ms | http |
| 104.238.30.50:59741 | ✓ 1596ms | 否 | ✓ 1935ms | 否 | ✓ 1935ms | http |
| 36.136.27.2:4999 | ✓ 1427ms | ✓ 1518ms | ✓ 1433ms | ✓ 1617ms | ✓ 1638ms | http |
| 20.27.15.111:8561 | ✓ 777ms | ✓ 1130ms | ✓ 782ms | ✓ 1108ms | ✓ 910ms | http |
| 20.210.39.153:8561 | ✓ 655ms | ✓ 1604ms | ✓ 615ms | ✓ 1037ms | ✓ 819ms | http |
| 20.78.26.206:8561 | ✓ 631ms | ✓ 1325ms | ✓ 892ms | ✓ 1047ms | ✓ 838ms | http |
| 20.27.14.220:8561 | ✓ 764ms | ✓ 1493ms | ✓ 678ms | ✓ 1012ms | ✓ 788ms | http |
| 20.27.11.248:8561 | ✓ 764ms | ✓ 1433ms | ✓ 757ms | ✓ 1011ms | ✓ 783ms | http |
| 20.210.76.104:8561 | ✓ 800ms | ✓ 1661ms | ✓ 883ms | ✓ 1098ms | ✓ 807ms | http |
| 20.210.76.178:8561 | ✓ 777ms | ✓ 1673ms | ✓ 893ms | ✓ 1106ms | ✓ 805ms | http |
| 20.210.76.175:8561 | ✓ 793ms | ✓ 1968ms | ✓ 764ms | ✓ 1071ms | ✓ 840ms | http |
| 202.152.44.19:8081 | ✓ 1044ms | 否 | ✓ 1140ms | ✓ 1457ms | ✓ 1162ms | http |
| 202.152.44.18:8081 | ✓ 1136ms | 否 | ✓ 1145ms | ✓ 1592ms | ✓ 1292ms | http |
| 36.147.78.166:80 | ✓ 1879ms | 否 | ✓ 1979ms | ✓ 1754ms | 否 | http |
| 104.238.30.63:63744 | ✓ 1621ms | 否 | ✓ 1870ms | 否 | ✓ 1904ms | http |
| 61.72.221.54:3128 | ✓ 1300ms | 否 | ✓ 1360ms | 否 | ✓ 1948ms | http |
| 190.9.109.207:999 | ✓ 707ms | ✓ 1475ms | ✓ 992ms | 否 | 否 | http |
| 223.16.170.103:80 | ✓ 1663ms | 否 | ✓ 1283ms | ✓ 1668ms | ✓ 1382ms | http |
| 103.153.149.229:1111 | 否 | 否 | ✓ 1624ms | ✓ 1830ms | ✓ 1768ms | http |
| 109.234.38.35:3128 | ✓ 526ms | 否 | ✓ 1444ms | ✓ 1994ms | ✓ 1612ms | http |
| 35.225.22.61:80 | ✓ 1116ms | 否 | ✓ 1214ms | ✓ 1120ms | ✓ 1112ms | http |
| 154.12.59.102:6005 | ✓ 1142ms | ✓ 1480ms | ✓ 1164ms | ✓ 1098ms | ✓ 1254ms | http |
| 61.72.110.94:3128 | ✓ 829ms | ✓ 1832ms | ✓ 1046ms | ✓ 1222ms | ✓ 988ms | http |
| 103.84.95.54:7890 | ✓ 840ms | 否 | ✓ 845ms | 否 | ✓ 853ms | http |
| 125.128.12.44:3128 | ✓ 957ms | 否 | ✓ 1293ms | 否 | ✓ 1537ms | http |
| 85.208.108.43:2094 | ✓ 1206ms | 否 | ✓ 1791ms | ✓ 1696ms | ✓ 1089ms | http |
| 91.233.223.147:3128 | ✓ 1769ms | 否 | ✓ 1083ms | 否 | ✓ 1939ms | http |
| 144.31.69.170:1080 | ✓ 1521ms | 否 | ✓ 581ms | 否 | ✓ 1755ms | http |
| 62.113.119.14:8080 | ✓ 1006ms | 否 | ✓ 1200ms | 否 | ✓ 1399ms | http |
| 165.225.222.26:12452 | ✓ 842ms | ✓ 1021ms | ✓ 1403ms | 否 | 否 | http |
| 207.254.71.62:8088 | ✓ 1799ms | 否 | ✓ 1673ms | ✓ 1848ms | ✓ 1611ms | http |
| 81.177.48.54:2080 | ✓ 1899ms | 否 | ✓ 1615ms | ✓ 1929ms | ✓ 1918ms | http |
| 124.16.93.70:7890 | 否 | ✓ 1319ms | ✓ 1175ms | ✓ 1371ms | ✓ 1086ms | http |
| 210.77.18.31:7890 | ✓ 993ms | ✓ 1576ms | ✓ 1085ms | ✓ 1345ms | ✓ 1041ms | http |
| 121.230.9.161:1080 | 否 | ✓ 1805ms | ✓ 1503ms | ✓ 1876ms | 否 | http |
| 157.100.53.103:999 | ✓ 1288ms | 否 | 否 | ✓ 1841ms | ✓ 1587ms | http |
| 157.100.53.110:999 | ✓ 1290ms | 否 | 否 | ✓ 1966ms | ✓ 1698ms | http |
| 35.212.218.202:1080 | ✓ 607ms | ✓ 1663ms | ✓ 1278ms | ✓ 1218ms | ✓ 1110ms | http |
| 120.46.152.136:3128 | ✓ 1159ms | ✓ 1431ms | ✓ 1688ms | ✓ 1679ms | ✓ 1360ms | http |
| 185.191.236.162:8080 | ✓ 912ms | 否 | ✓ 1056ms | ✓ 1564ms | ✓ 1028ms | http |
| 45.140.147.82:1081 | ✓ 452ms | 否 | ✓ 1476ms | 否 | ✓ 1388ms | http |
| 173.212.246.157:3128 | ✓ 1463ms | 否 | ✓ 1983ms | 否 | ✓ 1818ms | http |
| 121.230.9.11:1080 | ✓ 1470ms | ✓ 1880ms | 否 | ✓ 1942ms | 否 | http |
| 165.225.113.220:11807 | 否 | 否 | ✓ 995ms | ✓ 1262ms | ✓ 988ms | http |
| 165.225.113.220:10819 | 否 | 否 | ✓ 969ms | ✓ 1740ms | ✓ 996ms | http |
| 98.83.99.247:80 | ✓ 304ms | ✓ 1390ms | ✓ 980ms | 否 | 否 | http |
| 165.225.113.220:11462 | ✓ 908ms | 否 | 否 | ✓ 1257ms | ✓ 1294ms | http |
| 178.130.47.129:1082 | ✓ 1452ms | 否 | ✓ 1166ms | ✓ 946ms | 否 | http |
| 201.150.116.32:999 | 否 | 否 | ✓ 1625ms | ✓ 1279ms | ✓ 1144ms | http |
| 45.88.0.114:3128 | ✓ 774ms | 否 | ✓ 1136ms | ✓ 1343ms | 否 | http |
| 45.88.0.117:3128 | 否 | 否 | ✓ 695ms | ✓ 1314ms | ✓ 1300ms | http |
| 45.88.0.116:3128 | ✓ 563ms | 否 | ✓ 1237ms | 否 | ✓ 1596ms | http |
| 45.88.0.113:3128 | 否 | 否 | ✓ 703ms | ✓ 1318ms | ✓ 1115ms | http |
| 45.88.0.115:3128 | ✓ 578ms | 否 | ✓ 1237ms | 否 | ✓ 1345ms | http |
| 82.65.27.56:80 | ✓ 1471ms | 否 | ✓ 1137ms | ✓ 1537ms | 否 | http |
| 172.212.68.37:3128 | ✓ 320ms | 否 | ✓ 1130ms | ✓ 1296ms | ✓ 708ms | http |
| 45.129.141.143:3128 | ✓ 1465ms | 否 | ✓ 1475ms | 否 | ✓ 1491ms | http |
| 20.120.225.109:3128 | ✓ 1415ms | 否 | 否 | ✓ 962ms | ✓ 1037ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1503ms | ✓ 1666ms | ✓ 1447ms | 否 | http |
| 36.147.78.166:443 | ✓ 1943ms | ✓ 1892ms | ✓ 1925ms | ✓ 1855ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1650ms | ✓ 1629ms | 否 | 否 | ✓ 1128ms | http |

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
