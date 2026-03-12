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

最后更新：2026-03-12 23:29:02 UTC（2026-03-13 07:29:02 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 202.155.12.161:443 | ✓ 1121ms | ✓ 1377ms | ✓ 808ms | ✓ 1021ms | 否 | http |
| 205.209.118.30:3138 | ✓ 280ms | 否 | ✓ 961ms | ✓ 1283ms | ✓ 974ms | http |
| 1.231.81.166:3128 | ✓ 1180ms | ✓ 1206ms | ✓ 1054ms | ✓ 972ms | ✓ 833ms | http |
| 113.160.132.26:8080 | ✓ 1549ms | ✓ 1572ms | ✓ 1444ms | ✓ 1537ms | ✓ 982ms | http |
| 193.168.173.136:443 | ✓ 1798ms | 否 | ✓ 1290ms | 否 | ✓ 1814ms | http |
| 171.251.172.78:5104 | ✓ 1651ms | 否 | ✓ 1621ms | ✓ 1908ms | ✓ 1472ms | http |
| 171.251.172.78:5106 | ✓ 1550ms | 否 | ✓ 1603ms | ✓ 1548ms | ✓ 1380ms | http |
| 185.115.74.185:8080 | ✓ 1800ms | ✓ 1952ms | ✓ 1805ms | 否 | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1211ms | ✓ 955ms | 否 | ✓ 981ms | http |
| 217.76.245.80:999 | ✓ 954ms | ✓ 1523ms | ✓ 1453ms | ✓ 1475ms | ✓ 1061ms | http |
| 45.136.130.175:8443 | ✓ 435ms | ✓ 1389ms | ✓ 144ms | ✓ 745ms | ✓ 575ms | http |
| 115.231.181.40:8128 | ✓ 1397ms | ✓ 1184ms | ✓ 900ms | ✓ 1175ms | ✓ 1188ms | http |
| 84.247.149.172:3128 | ✓ 1451ms | 否 | ✓ 1738ms | ✓ 1361ms | ✓ 933ms | http |
| 190.9.109.198:999 | ✓ 755ms | ✓ 1517ms | ✓ 1261ms | ✓ 1422ms | ✓ 1427ms | http |
| 116.80.49.169:3172 | ✓ 1943ms | 否 | 否 | ✓ 1851ms | ✓ 1686ms | http |
| 138.124.53.25:7443 | ✓ 868ms | 否 | ✓ 1681ms | ✓ 1670ms | ✓ 1573ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1402ms | 否 | ✓ 1405ms | ✓ 1022ms | http |
| 46.183.25.8:443 | ✓ 343ms | 否 | ✓ 478ms | ✓ 988ms | ✓ 896ms | http |
| 152.42.213.210:8080 | ✓ 748ms | 否 | ✓ 1348ms | ✓ 1518ms | 否 | http |
| 120.92.212.16:8890 | ✓ 991ms | ✓ 1900ms | ✓ 967ms | ✓ 1447ms | 否 | http |
| 45.136.130.188:8443 | ✓ 149ms | ✓ 898ms | ✓ 627ms | ✓ 905ms | ✓ 556ms | http |
| 81.70.169.194:80 | ✓ 1005ms | ✓ 1236ms | ✓ 1054ms | ✓ 1226ms | ✓ 963ms | http |
| 101.43.255.96:80 | ✓ 954ms | ✓ 1301ms | ✓ 940ms | ✓ 1301ms | ✓ 1033ms | http |
| 178.236.245.59:3128 | ✓ 942ms | ✓ 1554ms | ✓ 770ms | 否 | ✓ 1624ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 696ms | ✓ 1021ms | ✓ 1258ms | http |
| 83.219.250.8:62920 | ✓ 928ms | 否 | ✓ 1779ms | 否 | ✓ 1785ms | http |
| 1.225.116.115:1080 | ✓ 1999ms | 否 | ✓ 1987ms | ✓ 1958ms | 否 | http |
| 111.79.111.126:3128 | ✓ 1619ms | ✓ 1617ms | ✓ 1974ms | 否 | 否 | http |
| 190.242.157.215:8080 | ✓ 972ms | 否 | ✓ 698ms | ✓ 1851ms | ✓ 1516ms | http |
| 45.136.130.223:8443 | ✓ 572ms | ✓ 1222ms | ✓ 629ms | ✓ 833ms | ✓ 857ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1945ms | 否 | ✓ 1819ms | ✓ 1288ms | http |
| 24.144.86.173:1080 | ✓ 394ms | ✓ 793ms | ✓ 856ms | ✓ 873ms | ✓ 568ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1990ms | ✓ 1038ms | ✓ 1975ms | ✓ 970ms | http |
| 116.80.49.167:3172 | ✓ 1533ms | 否 | ✓ 1807ms | 否 | ✓ 1673ms | http |
| 121.230.8.111:1080 | 否 | ✓ 1628ms | ✓ 1518ms | ✓ 1762ms | ✓ 1568ms | http |
| 168.235.110.63:3128 | ✓ 425ms | 否 | ✓ 1331ms | ✓ 1199ms | ✓ 1023ms | http |
| 217.52.247.87:1976 | 否 | 否 | ✓ 1897ms | ✓ 1795ms | ✓ 1541ms | http |
| 165.227.5.10:8888 | ✓ 922ms | ✓ 1022ms | ✓ 1433ms | 否 | ✓ 1270ms | http |
| 138.124.53.221:443 | ✓ 728ms | 否 | ✓ 1421ms | ✓ 1888ms | ✓ 1177ms | http |
| 121.230.9.205:1080 | ✓ 1276ms | ✓ 1447ms | ✓ 958ms | ✓ 1522ms | ✓ 1312ms | http |
| 201.150.116.32:999 | 否 | ✓ 1457ms | ✓ 1355ms | 否 | ✓ 1324ms | http |
| 116.80.49.170:3172 | ✓ 1712ms | 否 | ✓ 1786ms | ✓ 1830ms | 否 | http |
| 45.136.131.47:8443 | ✓ 423ms | ✓ 872ms | ✓ 134ms | ✓ 884ms | ✓ 679ms | http |
| 45.136.130.191:8443 | ✓ 424ms | ✓ 863ms | ✓ 147ms | ✓ 864ms | ✓ 705ms | http |
| 45.136.131.63:8443 | ✓ 426ms | ✓ 893ms | ✓ 232ms | ✓ 837ms | ✓ 578ms | http |
| 35.225.22.61:80 | 否 | ✓ 1178ms | ✓ 331ms | ✓ 1033ms | 否 | http |
| 8.219.97.248:80 | ✓ 1077ms | 否 | 否 | ✓ 1374ms | ✓ 1275ms | http |
| 107.173.52.58:7890 | ✓ 970ms | ✓ 1404ms | ✓ 1201ms | 否 | ✓ 1926ms | http |
| 14.225.212.37:7890 | ✓ 1447ms | 否 | ✓ 1185ms | 否 | ✓ 1124ms | http |
| 45.140.147.82:1081 | ✓ 617ms | ✓ 1808ms | ✓ 1422ms | 否 | ✓ 1839ms | http |
| 88.80.150.82:8080 | ✓ 1426ms | ✓ 1918ms | 否 | 否 | ✓ 1910ms | https |
| 101.32.244.83:8080 | ✓ 965ms | 否 | ✓ 949ms | ✓ 1288ms | ✓ 1269ms | http |
| 103.235.67.190:80 | ✓ 870ms | 否 | ✓ 1245ms | ✓ 1252ms | ✓ 1007ms | http |
| 121.230.8.55:1080 | ✓ 1074ms | ✓ 1400ms | ✓ 1105ms | ✓ 1888ms | ✓ 1710ms | http |
| 121.43.196.210:8222 | ✓ 926ms | ✓ 1083ms | ✓ 890ms | ✓ 1125ms | ✓ 877ms | http |
| 121.43.196.213:8222 | ✓ 926ms | ✓ 1192ms | ✓ 879ms | ✓ 1149ms | ✓ 939ms | http |
| 114.55.226.123:10086 | ✓ 1047ms | ✓ 1418ms | ✓ 1084ms | ✓ 1269ms | ✓ 1052ms | http |
| 86.109.3.24:10007 | ✓ 393ms | ✓ 983ms | ✓ 365ms | ✓ 1116ms | ✓ 853ms | http |
| 86.109.3.24:10010 | ✓ 977ms | ✓ 946ms | ✓ 452ms | ✓ 988ms | ✓ 878ms | http |
| 34.101.184.164:3128 | ✓ 1491ms | 否 | ✓ 1572ms | ✓ 1577ms | ✓ 981ms | http |
| 162.240.154.26:3128 | ✓ 1614ms | ✓ 1950ms | ✓ 1554ms | ✓ 1552ms | ✓ 1249ms | http |
| 180.76.115.231:3128 | ✓ 1034ms | ✓ 1714ms | ✓ 1557ms | ✓ 1994ms | ✓ 1110ms | http |
| 86.109.3.24:10000 | 否 | ✓ 1956ms | ✓ 533ms | ✓ 1075ms | ✓ 871ms | http |
| 152.42.213.210:443 | ✓ 746ms | 否 | ✓ 1174ms | ✓ 1060ms | ✓ 1041ms | http |
| 178.236.245.17:3128 | ✓ 743ms | ✓ 1875ms | ✓ 958ms | ✓ 1712ms | ✓ 1359ms | http |
| 103.179.253.11:8181 | ✓ 1438ms | 否 | 否 | ✓ 1765ms | ✓ 1788ms | http |
| 103.113.70.189:1081 | ✓ 422ms | ✓ 1257ms | 否 | ✓ 1235ms | ✓ 1031ms | http |
| 36.212.210.57:8088 | ✓ 746ms | ✓ 872ms | ✓ 900ms | ✓ 957ms | ✓ 753ms | http |
| 119.18.145.49:20326 | ✓ 1941ms | 否 | ✓ 1828ms | 否 | ✓ 1942ms | http |
| 106.117.208.101:7890 | ✓ 1003ms | ✓ 1322ms | ✓ 1069ms | ✓ 1254ms | ✓ 1045ms | http |
| 45.186.6.104:3128 | ✓ 1457ms | ✓ 1649ms | ✓ 1880ms | 否 | 否 | http |
| 185.191.236.162:3128 | ✓ 1708ms | ✓ 1632ms | ✓ 803ms | ✓ 1696ms | ✓ 1410ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1702ms | ✓ 1684ms | 否 | ✓ 957ms | http |
| 45.136.198.40:3128 | ✓ 857ms | ✓ 1880ms | ✓ 1798ms | 否 | ✓ 1842ms | http |
| 103.39.51.190:8080 | ✓ 1845ms | 否 | 否 | ✓ 1472ms | ✓ 1646ms | http |
| 116.80.64.44:7777 | ✓ 1874ms | 否 | ✓ 1549ms | 否 | ✓ 1926ms | http |
| 121.230.8.62:1080 | 否 | 否 | ✓ 1234ms | ✓ 1617ms | ✓ 1011ms | http |
| 118.113.246.173:1080 | ✓ 1277ms | ✓ 1763ms | ✓ 1533ms | ✓ 1810ms | ✓ 1656ms | http |
| 59.153.16.105:20909 | ✓ 1317ms | 否 | 否 | ✓ 1750ms | ✓ 1720ms | http |
| 170.78.208.245:999 | ✓ 736ms | 否 | ✓ 1139ms | ✓ 1275ms | ✓ 1260ms | http |
| 61.76.95.217:40088 | 否 | ✓ 1578ms | ✓ 1591ms | ✓ 1492ms | ✓ 961ms | http |
| 61.52.131.172:8443 | ✓ 934ms | ✓ 1188ms | ✓ 1049ms | ✓ 1209ms | ✓ 918ms | http |
| 45.207.200.85:1080 | ✓ 1121ms | ✓ 1536ms | ✓ 1364ms | 否 | 否 | http |
| 109.234.38.35:3128 | ✓ 618ms | ✓ 1837ms | ✓ 1709ms | ✓ 1400ms | ✓ 1077ms | http |
| 172.212.68.37:3128 | ✓ 408ms | 否 | 否 | ✓ 1450ms | ✓ 1527ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1930ms | ✓ 1545ms | ✓ 1412ms | http |

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
