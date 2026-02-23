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

最后更新：2026-02-23 12:32:31 UTC（2026-02-23 20:32:31 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1100ms | 否 | ✓ 954ms | ✓ 1393ms | ✓ 1055ms | http |
| 132.145.93.138:1080 | ✓ 740ms | 否 | 否 | ✓ 1345ms | ✓ 859ms | http |
| 125.128.12.94:3128 | ✓ 743ms | 否 | 否 | ✓ 1247ms | ✓ 1516ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1165ms | ✓ 951ms | ✓ 1158ms | ✓ 937ms | http |
| 103.236.64.247:8888 | 否 | ✓ 1185ms | ✓ 1197ms | 否 | ✓ 921ms | http |
| 211.230.49.122:3128 | ✓ 832ms | 否 | ✓ 1715ms | ✓ 1192ms | ✓ 932ms | http |
| 120.92.212.16:8890 | ✓ 1806ms | 否 | ✓ 1121ms | ✓ 1457ms | ✓ 1128ms | http |
| 168.235.110.63:3128 | ✓ 1116ms | 否 | ✓ 1371ms | 否 | ✓ 1287ms | http |
| 202.152.44.19:8081 | ✓ 1765ms | 否 | ✓ 840ms | 否 | ✓ 1932ms | http |
| 35.72.254.71:3128 | 否 | 否 | ✓ 1485ms | ✓ 1210ms | ✓ 1049ms | http |
| 121.204.158.249:3128 | 否 | ✓ 1857ms | ✓ 1046ms | ✓ 1257ms | ✓ 1010ms | http |
| 120.46.152.136:3128 | ✓ 1155ms | ✓ 1962ms | 否 | ✓ 1865ms | ✓ 1901ms | http |
| 54.89.108.25:80 | ✓ 826ms | 否 | ✓ 1503ms | ✓ 1742ms | ✓ 1398ms | http |
| 202.152.44.18:8081 | ✓ 993ms | 否 | ✓ 999ms | ✓ 1307ms | ✓ 1086ms | http |
| 89.208.85.78:443 | ✓ 966ms | ✓ 1833ms | ✓ 821ms | ✓ 1902ms | 否 | http |
| 89.208.85.78:18080 | ✓ 810ms | ✓ 1905ms | ✓ 1305ms | ✓ 1850ms | ✓ 1557ms | http |
| 45.12.151.226:2828 | ✓ 1432ms | 否 | ✓ 1285ms | ✓ 1967ms | 否 | http |
| 35.234.17.221:8080 | 否 | ✓ 1498ms | ✓ 1135ms | 否 | ✓ 800ms | http |
| 59.127.212.110:4431 | ✓ 1199ms | ✓ 1301ms | ✓ 1038ms | ✓ 1575ms | ✓ 1151ms | http |
| 36.147.78.166:80 | ✓ 1740ms | ✓ 1656ms | 否 | ✓ 1852ms | ✓ 1626ms | http |
| 78.13.231.158:3128 | ✓ 682ms | 否 | ✓ 1300ms | 否 | ✓ 1317ms | http |
| 83.243.86.97:443 | ✓ 875ms | 否 | ✓ 1878ms | 否 | ✓ 1908ms | http |
| 121.230.9.96:1080 | ✓ 1264ms | 否 | ✓ 1035ms | 否 | ✓ 1211ms | http |
| 178.253.22.108:65431 | ✓ 883ms | 否 | ✓ 1354ms | ✓ 1768ms | ✓ 1507ms | http |
| 59.46.216.131:30001 | ✓ 988ms | ✓ 1361ms | 否 | ✓ 1410ms | 否 | http |
| 36.136.27.2:4999 | ✓ 1842ms | 否 | ✓ 1979ms | 否 | ✓ 1805ms | http |
| 121.40.231.103:7890 | ✓ 821ms | ✓ 1016ms | ✓ 892ms | ✓ 1115ms | ✓ 802ms | http |
| 137.220.150.22:6005 | ✓ 1787ms | 否 | ✓ 793ms | ✓ 1355ms | ✓ 1158ms | http |
| 38.47.97.22:6005 | 否 | ✓ 1403ms | ✓ 1811ms | ✓ 1585ms | ✓ 1134ms | http |
| 18.229.170.122:3128 | ✓ 1488ms | 否 | ✓ 854ms | 否 | ✓ 1701ms | http |
| 160.16.204.90:3128 | ✓ 862ms | ✓ 1536ms | 否 | ✓ 1810ms | 否 | http |
| 91.238.104.172:2024 | 否 | 否 | ✓ 1466ms | ✓ 1872ms | ✓ 1392ms | http |
| 52.188.28.218:3128 | ✓ 1305ms | 否 | ✓ 319ms | 否 | ✓ 1933ms | http |
| 116.102.242.52:10018 | ✓ 1672ms | 否 | 否 | ✓ 1947ms | ✓ 1975ms | http |
| 61.72.110.4:3128 | ✓ 1088ms | 否 | ✓ 1899ms | 否 | ✓ 1670ms | http |
| 35.212.218.202:1080 | ✓ 330ms | 否 | ✓ 1959ms | ✓ 868ms | ✓ 787ms | http |
| 152.32.255.24:27197 | ✓ 1032ms | 否 | ✓ 1513ms | ✓ 1317ms | 否 | http |
| 8.219.97.248:80 | ✓ 1720ms | 否 | ✓ 1479ms | ✓ 1316ms | ✓ 1163ms | http |
| 217.217.254.94:8080 | 否 | 否 | ✓ 1292ms | ✓ 1380ms | ✓ 1300ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1276ms | 否 | ✓ 1377ms | ✓ 1080ms | http |
| 103.3.246.71:3128 | ✓ 1059ms | 否 | ✓ 879ms | ✓ 1095ms | ✓ 896ms | http |
| 115.114.77.133:9090 | ✓ 1605ms | 否 | ✓ 1518ms | 否 | ✓ 1433ms | http |
| 62.113.119.14:8080 | ✓ 1483ms | 否 | ✓ 1293ms | ✓ 1725ms | ✓ 1396ms | http |
| 144.124.227.234:3128 | ✓ 960ms | 否 | 否 | ✓ 1792ms | ✓ 1186ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 629ms | ✓ 1785ms | ✓ 898ms | http |
| 209.38.51.97:3128 | ✓ 575ms | 否 | ✓ 1088ms | ✓ 1361ms | 否 | http |
| 124.16.93.70:7890 | 否 | ✓ 1304ms | ✓ 907ms | ✓ 1044ms | ✓ 830ms | http |
| 110.172.29.131:3128 | ✓ 1732ms | 否 | ✓ 1202ms | ✓ 1161ms | ✓ 1093ms | http |
| 121.148.239.82:3129 | ✓ 744ms | 否 | 否 | ✓ 1215ms | ✓ 778ms | http |
| 121.147.215.213:3129 | ✓ 934ms | 否 | 否 | ✓ 1343ms | ✓ 1474ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1542ms | ✓ 1745ms | ✓ 1198ms | http |
| 211.171.114.154:3128 | ✓ 779ms | ✓ 1667ms | 否 | ✓ 1063ms | ✓ 1222ms | http |
| 120.79.99.232:8099 | ✓ 1161ms | ✓ 1461ms | ✓ 1193ms | ✓ 1348ms | ✓ 1146ms | http |
| 14.56.118.44:3128 | ✓ 1378ms | 否 | 否 | ✓ 1225ms | ✓ 1782ms | http |
| 202.129.206.239:3128 | ✓ 1577ms | 否 | ✓ 1537ms | ✓ 1536ms | ✓ 1174ms | http |
| 85.208.108.43:2094 | ✓ 401ms | 否 | ✓ 1902ms | ✓ 1469ms | ✓ 1021ms | http |
| 44.209.129.187:80 | ✓ 968ms | 否 | ✓ 1539ms | ✓ 1580ms | ✓ 1378ms | http |
| 64.181.240.152:3128 | ✓ 624ms | ✓ 1846ms | 否 | ✓ 1244ms | ✓ 755ms | http |
| 188.166.208.168:9876 | ✓ 680ms | 否 | ✓ 714ms | ✓ 1003ms | ✓ 806ms | http |
| 116.102.242.52:10007 | 否 | 否 | ✓ 1457ms | ✓ 1505ms | ✓ 1481ms | http |
| 175.194.173.105:3128 | ✓ 1862ms | ✓ 1157ms | ✓ 991ms | 否 | 否 | http |
| 36.147.78.166:443 | 否 | ✓ 1592ms | 否 | ✓ 1839ms | ✓ 1554ms | http |
| 116.102.242.52:10028 | 否 | 否 | ✓ 1728ms | ✓ 1639ms | ✓ 1261ms | http |
| 139.159.99.242:8080 | ✓ 802ms | ✓ 931ms | ✓ 1193ms | 否 | 否 | http |
| 116.102.242.52:2002 | ✓ 1947ms | 否 | ✓ 1379ms | ✓ 1498ms | ✓ 1294ms | http |
| 103.82.23.118:5249 | ✓ 1751ms | 否 | ✓ 1661ms | ✓ 1761ms | ✓ 1477ms | http |
| 43.161.214.161:1081 | ✓ 1618ms | 否 | ✓ 1778ms | ✓ 1816ms | 否 | http |
| 14.56.177.182:3128 | ✓ 710ms | ✓ 1782ms | ✓ 760ms | ✓ 1264ms | ✓ 1567ms | http |
| 14.56.177.162:3128 | ✓ 782ms | 否 | ✓ 1371ms | ✓ 1798ms | ✓ 1215ms | http |
| 116.102.242.52:10008 | ✓ 1953ms | 否 | ✓ 1722ms | ✓ 1776ms | ✓ 1240ms | http |
| 45.142.36.250:3128 | ✓ 1425ms | 否 | ✓ 1637ms | 否 | ✓ 1970ms | http |
| 116.102.242.52:10029 | ✓ 1591ms | 否 | ✓ 1748ms | 否 | ✓ 1390ms | http |
| 66.228.47.125:110 | ✓ 784ms | 否 | ✓ 1601ms | ✓ 1940ms | ✓ 1307ms | http |
| 217.216.109.116:8080 | ✓ 1954ms | 否 | 否 | ✓ 1379ms | ✓ 1252ms | http |
| 61.72.221.14:3128 | ✓ 997ms | 否 | ✓ 1682ms | 否 | ✓ 1587ms | http |
| 50.17.39.43:80 | ✓ 406ms | ✓ 1849ms | ✓ 797ms | 否 | ✓ 902ms | http |
| 186.67.74.52:3128 | ✓ 765ms | 否 | ✓ 1020ms | 否 | ✓ 1581ms | http |
| 14.56.177.140:3128 | 否 | 否 | ✓ 979ms | ✓ 1757ms | ✓ 1740ms | http |
| 103.39.51.190:8080 | ✓ 1704ms | 否 | 否 | ✓ 1311ms | ✓ 1465ms | http |
| 81.177.48.54:2080 | ✓ 1107ms | 否 | ✓ 1054ms | 否 | ✓ 1451ms | http |
| 147.45.159.213:48206 | ✓ 1154ms | 否 | ✓ 1199ms | 否 | ✓ 1646ms | http |
| 14.56.177.54:3128 | 否 | 否 | ✓ 578ms | ✓ 949ms | ✓ 782ms | http |
| 18.117.46.60:3128 | ✓ 1034ms | 否 | ✓ 1981ms | ✓ 1360ms | ✓ 1224ms | http |
| 116.102.242.52:10011 | 否 | 否 | ✓ 1505ms | ✓ 1992ms | ✓ 1304ms | http |
| 165.227.5.10:8888 | ✓ 1085ms | 否 | ✓ 1151ms | ✓ 863ms | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1658ms | ✓ 1929ms | ✓ 1417ms | ✓ 1040ms | http |
| 217.216.109.116:80 | ✓ 1450ms | 否 | ✓ 1656ms | 否 | ✓ 1793ms | http |

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
