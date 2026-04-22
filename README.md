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

最后更新：2026-04-22 23:44:28 UTC（2026-04-23 07:44:28 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | 否 | ✓ 1126ms | ✓ 837ms | ✓ 1015ms | ✓ 851ms | http |
| 152.42.208.139:8118 | ✓ 1487ms | 否 | ✓ 1460ms | ✓ 1190ms | ✓ 941ms | http |
| 130.61.174.200:1080 | ✓ 1464ms | ✓ 1806ms | ✓ 1951ms | 否 | ✓ 1994ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1895ms | ✓ 1112ms | ✓ 1866ms | ✓ 1502ms | http |
| 78.11.96.22:8888 | ✓ 1478ms | ✓ 1818ms | ✓ 1651ms | ✓ 1752ms | ✓ 1815ms | http |
| 47.85.51.197:1080 | ✓ 548ms | ✓ 1276ms | ✓ 615ms | ✓ 1896ms | ✓ 1091ms | http |
| 35.225.22.61:80 | ✓ 506ms | ✓ 1673ms | ✓ 1197ms | ✓ 1011ms | ✓ 859ms | http |
| 45.59.122.132:80 | ✓ 752ms | ✓ 1776ms | ✓ 1267ms | ✓ 1681ms | ✓ 1239ms | http |
| 159.223.225.118:8888 | ✓ 869ms | 否 | ✓ 786ms | 否 | ✓ 1354ms | http |
| 62.113.119.14:8080 | ✓ 565ms | ✓ 1367ms | ✓ 787ms | 否 | 否 | http |
| 152.32.132.190:7890 | 否 | ✓ 1096ms | 否 | ✓ 995ms | ✓ 1059ms | http |
| 115.231.181.40:8128 | ✓ 990ms | 否 | ✓ 1010ms | ✓ 1336ms | 否 | http |
| 84.47.150.125:1080 | ✓ 980ms | 否 | ✓ 1968ms | 否 | ✓ 1923ms | http |
| 218.108.131.186:17890 | 否 | 否 | ✓ 1231ms | ✓ 1392ms | ✓ 1228ms | http |
| 120.92.108.86:7890 | ✓ 1628ms | 否 | ✓ 1965ms | ✓ 1800ms | ✓ 1501ms | http |
| 59.46.216.131:30001 | ✓ 1026ms | ✓ 1457ms | 否 | ✓ 1475ms | 否 | http |
| 208.87.243.199:7878 | ✓ 1518ms | ✓ 1097ms | ✓ 1445ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 1216ms | ✓ 1935ms | ✓ 742ms | ✓ 1985ms | ✓ 1537ms | http |
| 46.101.95.183:8888 | ✓ 1393ms | ✓ 1875ms | ✓ 978ms | ✓ 1838ms | 否 | http |
| 83.219.250.8:62920 | ✓ 1151ms | ✓ 1320ms | ✓ 1530ms | 否 | ✓ 1681ms | http |
| 120.92.212.16:7890 | ✓ 987ms | ✓ 1368ms | ✓ 1333ms | ✓ 1615ms | ✓ 1112ms | http |
| 120.92.212.16:8890 | ✓ 1341ms | 否 | 否 | ✓ 1656ms | ✓ 1139ms | http |
| 91.99.15.45:2095 | ✓ 930ms | ✓ 1899ms | ✓ 1248ms | 否 | 否 | http |
| 45.153.231.229:8080 | ✓ 1031ms | 否 | ✓ 1787ms | ✓ 1919ms | 否 | http |
| 103.133.254.4:3128 | ✓ 1403ms | 否 | ✓ 1852ms | 否 | ✓ 1865ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1900ms | ✓ 1908ms | ✓ 1784ms | 否 | http |
| 147.45.166.46:3128 | ✓ 924ms | 否 | ✓ 869ms | 否 | ✓ 1527ms | http |
| 20.127.128.70:8080 | ✓ 1134ms | ✓ 1972ms | ✓ 772ms | 否 | 否 | http |
| 34.71.229.255:3128 | ✓ 1110ms | ✓ 1901ms | ✓ 1763ms | ✓ 1604ms | ✓ 1502ms | http |
| 8.219.195.129:1080 | ✓ 904ms | ✓ 1770ms | ✓ 823ms | ✓ 1177ms | ✓ 943ms | http |
| 77.110.113.24:40000 | ✓ 1067ms | 否 | ✓ 721ms | 否 | ✓ 1602ms | http |
| 103.163.80.25:8081 | ✓ 1691ms | 否 | 否 | ✓ 1772ms | ✓ 1938ms | http |
| 194.150.220.163:1082 | ✓ 1585ms | ✓ 1363ms | ✓ 1599ms | ✓ 1752ms | 否 | http |
| 162.240.154.26:3128 | ✓ 1149ms | ✓ 1255ms | ✓ 1493ms | 否 | 否 | http |
| 38.180.2.107:3128 | ✓ 756ms | ✓ 1679ms | ✓ 1822ms | 否 | 否 | http |
| 168.144.75.9:3128 | ✓ 1059ms | 否 | ✓ 1937ms | 否 | ✓ 1496ms | http |
| 20.164.75.153:8080 | ✓ 1516ms | 否 | ✓ 1250ms | 否 | ✓ 1918ms | http |
| 139.159.97.82:10900 | ✓ 1342ms | 否 | ✓ 1517ms | ✓ 1619ms | ✓ 1264ms | http |
| 94.158.219.111:3128 | 否 | ✓ 1575ms | ✓ 827ms | 否 | ✓ 1618ms | http |
| 177.93.132.244:3128 | ✓ 930ms | 否 | ✓ 1353ms | 否 | ✓ 1748ms | http |
| 14.143.222.113:57788 | ✓ 1934ms | 否 | ✓ 1026ms | ✓ 1437ms | 否 | http |
| 114.237.77.245:1080 | ✓ 989ms | ✓ 1307ms | ✓ 975ms | ✓ 1382ms | ✓ 1030ms | http |
| 176.124.220.172:3128 | ✓ 1348ms | ✓ 1784ms | ✓ 1077ms | 否 | ✓ 1349ms | http |
| 121.230.8.162:1080 | ✓ 1190ms | ✓ 1490ms | ✓ 1534ms | ✓ 1744ms | ✓ 1243ms | http |
| 118.113.247.73:1080 | ✓ 1108ms | ✓ 1595ms | ✓ 1474ms | ✓ 1777ms | ✓ 1232ms | http |
| 147.45.186.28:3128 | ✓ 1036ms | ✓ 1889ms | ✓ 1350ms | ✓ 1587ms | ✓ 1480ms | http |
| 43.252.106.26:1111 | ✓ 1572ms | 否 | 否 | ✓ 1631ms | ✓ 1814ms | http |
| 47.105.98.23:3128 | ✓ 984ms | ✓ 1295ms | ✓ 1536ms | 否 | 否 | http |
| 45.229.16.49:999 | ✓ 1414ms | ✓ 1614ms | ✓ 1590ms | ✓ 1968ms | ✓ 1565ms | http |
| 68.178.167.154:9999 | ✓ 826ms | 否 | ✓ 1126ms | ✓ 1427ms | ✓ 1170ms | http |
| 94.241.173.165:1080 | ✓ 409ms | 否 | ✓ 1436ms | ✓ 1407ms | 否 | http |
| 89.208.106.138:10808 | ✓ 658ms | 否 | ✓ 1733ms | 否 | ✓ 1136ms | http |
| 85.190.99.143:443 | ✓ 1061ms | 否 | ✓ 1935ms | 否 | ✓ 1764ms | http |
| 138.124.99.216:8888 | ✓ 1629ms | ✓ 1502ms | ✓ 1298ms | ✓ 1636ms | 否 | http |
| 93.43.251.159:80 | ✓ 1082ms | ✓ 1869ms | ✓ 1762ms | ✓ 1782ms | ✓ 1550ms | http |
| 84.47.150.126:1080 | ✓ 1355ms | 否 | ✓ 1811ms | 否 | ✓ 1582ms | http |
| 192.241.132.92:80 | 否 | ✓ 994ms | ✓ 857ms | ✓ 1031ms | ✓ 747ms | http |
| 45.76.207.177:40000 | ✓ 1490ms | 否 | ✓ 778ms | ✓ 1144ms | ✓ 1148ms | http |
| 91.193.240.157:9877 | ✓ 1064ms | 否 | ✓ 919ms | 否 | ✓ 1579ms | http |
| 161.97.184.191:8080 | 否 | ✓ 1998ms | ✓ 1984ms | ✓ 1897ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1229ms | ✓ 1379ms | ✓ 1449ms | ✓ 1601ms | ✓ 1225ms | http |
| 150.136.153.231:80 | 否 | 否 | ✓ 1015ms | ✓ 1259ms | ✓ 965ms | http |
| 217.77.102.18:3128 | ✓ 1285ms | ✓ 1925ms | ✓ 1661ms | 否 | ✓ 1547ms | http |
| 45.140.147.155:1081 | ✓ 1119ms | ✓ 1515ms | ✓ 1005ms | 否 | ✓ 1550ms | http |
| 116.171.106.15:3443 | ✓ 1643ms | ✓ 1987ms | ✓ 1776ms | 否 | 否 | http |
| 82.148.18.242:443 | ✓ 1243ms | ✓ 1882ms | 否 | 否 | ✓ 1767ms | http |
| 121.230.8.220:1080 | ✓ 1156ms | ✓ 1413ms | ✓ 1115ms | ✓ 1588ms | ✓ 1150ms | http |
| 121.230.8.138:1080 | ✓ 1394ms | ✓ 1544ms | ✓ 1198ms | ✓ 1457ms | ✓ 1129ms | http |
| 103.71.22.8:3128 | ✓ 1345ms | ✓ 1482ms | ✓ 1970ms | ✓ 1896ms | ✓ 1781ms | http |
| 103.52.114.95:3128 | ✓ 1743ms | 否 | ✓ 1150ms | ✓ 1572ms | ✓ 1274ms | http |
| 121.230.9.209:1080 | ✓ 1222ms | ✓ 1712ms | ✓ 1442ms | ✓ 1728ms | ✓ 1528ms | http |
| 112.78.134.134:7777 | ✓ 1767ms | 否 | 否 | ✓ 1907ms | ✓ 1734ms | http |
| 51.79.207.21:8080 | ✓ 1419ms | 否 | ✓ 1346ms | 否 | ✓ 1528ms | http |
| 45.186.6.104:3128 | ✓ 1417ms | ✓ 1864ms | ✓ 1821ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1017ms | ✓ 1342ms | ✓ 1033ms | ✓ 1275ms | ✓ 1014ms | http |
| 121.230.8.136:1080 | ✓ 1143ms | ✓ 1512ms | ✓ 1158ms | ✓ 1668ms | ✓ 1095ms | http |

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
