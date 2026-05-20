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

最后更新：2026-05-20 11:48:10 UTC（2026-05-20 19:48:10 UTC+8）

**代理总数：74**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 192.99.8.15:8850 | ✓ 798ms | 否 | ✓ 1216ms | ✓ 1366ms | ✓ 1502ms | http |
| 147.45.41.112:1080 | ✓ 1339ms | 否 | ✓ 1963ms | ✓ 1663ms | ✓ 1915ms | http |
| 113.160.132.26:8080 | ✓ 1674ms | 否 | 否 | ✓ 1555ms | ✓ 1697ms | http |
| 138.2.78.251:8100 | ✓ 1423ms | 否 | ✓ 1468ms | ✓ 1433ms | ✓ 1158ms | http |
| 138.2.92.70:8100 | ✓ 1419ms | 否 | ✓ 1806ms | ✓ 1433ms | ✓ 1950ms | http |
| 45.117.163.134:3128 | ✓ 920ms | 否 | ✓ 933ms | ✓ 1078ms | ✓ 848ms | http |
| 89.58.50.94:11140 | ✓ 1065ms | 否 | ✓ 1509ms | 否 | ✓ 1904ms | http |
| 43.130.126.146:6688 | ✓ 1761ms | 否 | ✓ 1247ms | ✓ 1661ms | 否 | http |
| 152.67.191.232:6800 | ✓ 1736ms | 否 | 否 | ✓ 1819ms | ✓ 1931ms | http |
| 45.125.67.37:8443 | ✓ 1333ms | 否 | ✓ 1169ms | ✓ 1055ms | 否 | http |
| 8.210.48.83:8100 | 否 | 否 | ✓ 1105ms | ✓ 1825ms | ✓ 1170ms | http |
| 188.34.156.126:24575 | ✓ 973ms | ✓ 1776ms | 否 | 否 | ✓ 1419ms | http |
| 144.124.227.90:21074 | ✓ 725ms | ✓ 1780ms | ✓ 1986ms | 否 | 否 | http |
| 202.28.194.139:31280 | ✓ 1968ms | 否 | ✓ 1937ms | ✓ 1959ms | ✓ 1937ms | http |
| 47.242.163.146:8100 | ✓ 1318ms | 否 | ✓ 1693ms | ✓ 1986ms | 否 | http |
| 176.111.37.5:39811 | ✓ 940ms | 否 | 否 | ✓ 1879ms | ✓ 1449ms | http |
| 8.217.78.60:8100 | ✓ 1711ms | 否 | 否 | ✓ 1251ms | ✓ 1162ms | http |
| 103.82.23.118:5249 | ✓ 1783ms | 否 | ✓ 1101ms | ✓ 1921ms | ✓ 1230ms | http |
| 212.58.132.5:8888 | ✓ 1239ms | 否 | ✓ 1320ms | ✓ 1615ms | ✓ 1300ms | http |
| 148.230.4.241:999 | ✓ 705ms | 否 | ✓ 677ms | ✓ 1705ms | ✓ 1318ms | http |
| 170.106.136.181:31002 | ✓ 720ms | ✓ 1398ms | 否 | ✓ 1410ms | ✓ 507ms | http |
| 34.87.80.221:30000 | ✓ 733ms | 否 | ✓ 1230ms | ✓ 1122ms | ✓ 833ms | http |
| 114.214.165.78:10810 | ✓ 1303ms | 否 | ✓ 1435ms | 否 | ✓ 1327ms | http |
| 171.249.232.167:4007 | ✓ 1357ms | 否 | ✓ 946ms | ✓ 1445ms | ✓ 928ms | http |
| 171.249.232.167:4002 | ✓ 1358ms | 否 | ✓ 946ms | ✓ 1555ms | ✓ 845ms | http |
| 121.230.9.5:1080 | ✓ 1083ms | 否 | 否 | ✓ 1418ms | ✓ 1120ms | http |
| 47.241.32.135:8100 | 否 | ✓ 1762ms | ✓ 1800ms | ✓ 1149ms | 否 | http |
| 47.243.206.29:8100 | ✓ 1299ms | ✓ 1901ms | ✓ 1172ms | 否 | 否 | http |
| 138.2.239.213:10010 | ✓ 528ms | 否 | ✓ 1755ms | ✓ 1018ms | ✓ 1073ms | http |
| 115.231.181.40:8128 | ✓ 1076ms | ✓ 1800ms | ✓ 1814ms | 否 | 否 | http |
| 180.191.233.125:8080 | 否 | 否 | ✓ 1705ms | ✓ 1570ms | ✓ 1551ms | http |
| 8.219.97.248:80 | ✓ 1866ms | 否 | 否 | ✓ 1467ms | ✓ 1198ms | http |
| 34.101.184.164:3128 | ✓ 1541ms | 否 | 否 | ✓ 1225ms | ✓ 1475ms | http |
| 190.12.150.244:999 | ✓ 975ms | 否 | ✓ 1199ms | ✓ 1757ms | ✓ 1497ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1953ms | 否 | ✓ 1690ms | ✓ 1262ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 938ms | ✓ 936ms | ✓ 673ms | http |
| 168.138.171.204:8100 | 否 | 否 | ✓ 955ms | ✓ 1673ms | ✓ 1918ms | http |
| 146.56.110.131:8118 | ✓ 1784ms | ✓ 894ms | 否 | ✓ 885ms | 否 | http |
| 157.0.142.246:10057 | ✓ 994ms | ✓ 1262ms | ✓ 1012ms | ✓ 1494ms | ✓ 1066ms | http |
| 121.230.8.220:1080 | ✓ 1206ms | ✓ 1256ms | ✓ 1107ms | ✓ 1864ms | ✓ 1052ms | http |
| 8.154.21.175:3128 | ✓ 907ms | ✓ 1478ms | ✓ 1099ms | 否 | ✓ 1311ms | http |
| 81.30.156.115:8080 | ✓ 1114ms | 否 | ✓ 1312ms | 否 | ✓ 1803ms | http |
| 121.230.8.109:1080 | ✓ 1278ms | ✓ 1622ms | ✓ 976ms | ✓ 1596ms | ✓ 1927ms | http |
| 121.230.9.33:1080 | ✓ 1209ms | ✓ 1568ms | ✓ 1366ms | 否 | ✓ 1225ms | http |
| 101.32.244.83:8080 | ✓ 1020ms | 否 | ✓ 991ms | ✓ 1176ms | ✓ 1264ms | http |
| 121.43.196.210:8222 | ✓ 993ms | ✓ 1033ms | ✓ 884ms | ✓ 1060ms | ✓ 912ms | http |
| 121.43.196.213:8222 | ✓ 989ms | ✓ 1082ms | ✓ 873ms | ✓ 1081ms | ✓ 902ms | http |
| 38.250.126.225:999 | ✓ 1674ms | 否 | ✓ 902ms | 否 | ✓ 1709ms | http |
| 8.210.161.8:8100 | ✓ 1109ms | 否 | ✓ 1923ms | ✓ 1326ms | 否 | http |
| 161.117.225.78:8100 | ✓ 1281ms | ✓ 1905ms | 否 | 否 | ✓ 1354ms | http |
| 119.23.68.90:9003 | ✓ 922ms | ✓ 1203ms | ✓ 1038ms | ✓ 1041ms | ✓ 862ms | http |
| 152.42.177.32:8888 | ✓ 938ms | 否 | ✓ 937ms | ✓ 1424ms | ✓ 1254ms | http |
| 138.68.101.169:3128 | ✓ 622ms | 否 | ✓ 1708ms | ✓ 1936ms | ✓ 1450ms | http |
| 3.101.133.120:80 | ✓ 805ms | ✓ 1832ms | ✓ 241ms | ✓ 1100ms | ✓ 1932ms | http |
| 8.218.174.172:8100 | ✓ 1051ms | 否 | ✓ 1142ms | 否 | ✓ 1114ms | http |
| 152.70.91.193:40000 | ✓ 1358ms | 否 | ✓ 1765ms | ✓ 1521ms | ✓ 1220ms | http |
| 185.200.188.234:10001 | ✓ 1476ms | 否 | ✓ 1178ms | 否 | ✓ 1761ms | http |
| 8.210.138.49:8100 | ✓ 1401ms | 否 | ✓ 1894ms | ✓ 1150ms | ✓ 996ms | http |
| 8.219.194.60:8100 | ✓ 1036ms | 否 | 否 | ✓ 1576ms | ✓ 1458ms | http |
| 115.198.210.72:7890 | ✓ 1304ms | ✓ 1613ms | 否 | 否 | ✓ 1029ms | http |
| 193.237.192.7:8888 | ✓ 1196ms | 否 | 否 | ✓ 1996ms | ✓ 1617ms | http |
| 114.214.163.108:6789 | ✓ 1044ms | 否 | ✓ 1307ms | ✓ 1347ms | ✓ 1076ms | http |
| 67.207.82.252:3128 | ✓ 1526ms | ✓ 1425ms | ✓ 999ms | ✓ 1153ms | 否 | http |
| 103.82.23.118:5182 | 否 | ✓ 1872ms | ✓ 1348ms | ✓ 1724ms | ✓ 1372ms | http |
| 192.81.129.252:3136 | 否 | 否 | ✓ 793ms | ✓ 1646ms | ✓ 1581ms | http |
| 217.76.245.80:999 | ✓ 1624ms | ✓ 1720ms | ✓ 1597ms | ✓ 1540ms | ✓ 1394ms | http |
| 8.210.132.233:8100 | 否 | ✓ 1368ms | ✓ 1744ms | ✓ 1263ms | ✓ 1970ms | http |
| 69.164.251.114:8080 | ✓ 1332ms | 否 | ✓ 1340ms | 否 | ✓ 1778ms | http |
| 61.52.131.172:8443 | ✓ 836ms | ✓ 1240ms | ✓ 947ms | ✓ 1247ms | ✓ 933ms | http |
| 121.230.8.250:1080 | ✓ 1194ms | ✓ 1494ms | ✓ 1703ms | ✓ 1920ms | ✓ 1540ms | http |
| 103.172.70.173:8080 | ✓ 1315ms | 否 | 否 | ✓ 1537ms | ✓ 1412ms | http |
| 223.16.170.103:80 | ✓ 1227ms | 否 | ✓ 1474ms | 否 | ✓ 1059ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1171ms | ✓ 1260ms | ✓ 1727ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1881ms | ✓ 1833ms | ✓ 1710ms | http |

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
