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

最后更新：2026-06-13 17:05:08 UTC（2026-06-14 01:05:08 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 159.198.35.187:1080 | ✓ 1897ms | 否 | ✓ 508ms | ✓ 1389ms | ✓ 950ms | http |
| 138.124.114.42:7443 | ✓ 1313ms | 否 | ✓ 1753ms | ✓ 1439ms | ✓ 1601ms | http |
| 169.212.15.161:5000 | 否 | ✓ 1575ms | ✓ 1715ms | ✓ 1510ms | ✓ 1101ms | http |
| 185.200.188.234:10001 | ✓ 1357ms | 否 | ✓ 1793ms | 否 | ✓ 1633ms | http |
| 113.160.132.26:8080 | ✓ 1717ms | ✓ 1739ms | ✓ 1093ms | ✓ 1814ms | ✓ 1244ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1589ms | ✓ 1531ms | 否 | ✓ 1276ms | http |
| 95.140.154.156:1080 | ✓ 1350ms | 否 | ✓ 1848ms | 否 | ✓ 1600ms | http |
| 202.28.194.139:31280 | 否 | 否 | ✓ 1862ms | ✓ 1884ms | ✓ 1990ms | http |
| 170.106.136.181:31002 | ✓ 862ms | ✓ 1012ms | ✓ 660ms | ✓ 883ms | ✓ 648ms | http |
| 34.43.46.91:443 | ✓ 701ms | 否 | ✓ 669ms | ✓ 1150ms | ✓ 1115ms | http |
| 34.43.46.91:80 | ✓ 646ms | 否 | ✓ 669ms | ✓ 1324ms | ✓ 1075ms | http |
| 150.241.116.167:443 | ✓ 442ms | 否 | ✓ 966ms | ✓ 1809ms | ✓ 1394ms | http |
| 185.141.26.131:3128 | ✓ 738ms | ✓ 1519ms | ✓ 1318ms | 否 | ✓ 1401ms | http |
| 82.97.247.37:80 | 否 | 否 | ✓ 1294ms | ✓ 1306ms | ✓ 981ms | http |
| 91.107.172.30:82 | ✓ 937ms | 否 | ✓ 1356ms | ✓ 1680ms | 否 | http |
| 213.165.42.185:7443 | ✓ 1175ms | 否 | ✓ 1401ms | 否 | ✓ 1708ms | http |
| 185.11.134.227:8443 | ✓ 1399ms | ✓ 1459ms | ✓ 725ms | ✓ 1753ms | 否 | http |
| 52.188.28.218:3128 | ✓ 1511ms | ✓ 1380ms | ✓ 854ms | ✓ 1633ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1252ms | 否 | ✓ 1356ms | ✓ 1579ms | ✓ 1548ms | http |
| 84.47.150.125:1080 | ✓ 1185ms | 否 | ✓ 1496ms | 否 | ✓ 1569ms | http |
| 95.3.69.222:8080 | ✓ 1969ms | 否 | ✓ 1577ms | 否 | ✓ 1769ms | http |
| 91.107.168.255:82 | ✓ 1009ms | 否 | ✓ 828ms | ✓ 1395ms | ✓ 1224ms | http |
| 176.111.37.5:39811 | ✓ 981ms | 否 | ✓ 988ms | 否 | ✓ 1219ms | http |
| 81.200.154.236:48503 | ✓ 1727ms | 否 | ✓ 1334ms | ✓ 1175ms | ✓ 853ms | http |
| 144.31.134.29:1050 | ✓ 1383ms | 否 | ✓ 1091ms | ✓ 1652ms | ✓ 1350ms | http |
| 85.192.28.47:7443 | ✓ 652ms | 否 | 否 | ✓ 1973ms | ✓ 1179ms | http |
| 180.2.108.38:8080 | ✓ 1531ms | 否 | ✓ 886ms | ✓ 1120ms | ✓ 860ms | http |
| 138.197.68.35:4857 | ✓ 436ms | ✓ 917ms | ✓ 55ms | ✓ 1350ms | ✓ 855ms | http |
| 31.57.172.220:10808 | ✓ 707ms | 否 | ✓ 1470ms | 否 | ✓ 1086ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 981ms | ✓ 1645ms | ✓ 1230ms | http |
| 3.137.86.220:1080 | ✓ 929ms | 否 | ✓ 1896ms | ✓ 1690ms | 否 | http |
| 85.192.28.65:7443 | ✓ 974ms | 否 | 否 | ✓ 1766ms | ✓ 1091ms | http |
| 116.63.160.98:8899 | ✓ 1201ms | ✓ 1451ms | ✓ 1201ms | 否 | 否 | http |
| 157.245.100.190:442 | 否 | 否 | ✓ 1490ms | ✓ 1319ms | ✓ 1054ms | http |
| 91.107.182.124:82 | ✓ 994ms | 否 | 否 | ✓ 1765ms | ✓ 1777ms | http |
| 34.204.185.68:3128 | ✓ 206ms | ✓ 1600ms | ✓ 816ms | 否 | ✓ 1136ms | http |
| 18.180.59.181:80 | ✓ 736ms | ✓ 1025ms | 否 | ✓ 1392ms | ✓ 917ms | http |
| 144.172.114.214:1080 | ✓ 622ms | 否 | 否 | ✓ 1596ms | ✓ 1169ms | http |
| 151.243.180.211:2080 | ✓ 1772ms | 否 | ✓ 594ms | ✓ 1992ms | ✓ 1704ms | http |
| 79.137.205.130:7443 | ✓ 1385ms | 否 | ✓ 1618ms | 否 | ✓ 976ms | http |
| 47.80.103.120:8080 | 否 | 否 | ✓ 1735ms | ✓ 1739ms | ✓ 1306ms | http |
| 82.102.11.164:3460 | ✓ 926ms | ✓ 1788ms | ✓ 1769ms | ✓ 1813ms | 否 | http |
| 3.137.86.220:443 | ✓ 1019ms | 否 | ✓ 1928ms | 否 | ✓ 1309ms | http |
| 104.154.186.48:80 | ✓ 465ms | ✓ 1076ms | ✓ 170ms | ✓ 1178ms | ✓ 685ms | http |
| 209.141.35.94:28080 | ✓ 1298ms | 否 | ✓ 1129ms | 否 | ✓ 1570ms | http |
| 8.215.25.3:2080 | ✓ 1693ms | 否 | ✓ 1019ms | ✓ 1469ms | ✓ 1086ms | http |
| 157.230.220.25:4857 | ✓ 317ms | ✓ 1142ms | ✓ 1873ms | ✓ 1247ms | ✓ 716ms | http |
| 89.127.207.174:18080 | ✓ 523ms | ✓ 1756ms | ✓ 732ms | ✓ 1594ms | ✓ 1545ms | http |
| 144.31.73.173:3128 | ✓ 1227ms | 否 | 否 | ✓ 1804ms | ✓ 1678ms | http |
| 14.143.222.113:57788 | ✓ 1280ms | 否 | ✓ 1917ms | ✓ 1667ms | 否 | http |
| 34.71.229.255:3128 | 否 | 否 | ✓ 1442ms | ✓ 1387ms | ✓ 1629ms | http |
| 85.192.28.62:7443 | 否 | 否 | ✓ 838ms | ✓ 1915ms | ✓ 1660ms | http |
| 85.192.60.187:7443 | ✓ 1257ms | 否 | ✓ 1029ms | ✓ 1726ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1076ms | 否 | 否 | ✓ 1179ms | ✓ 1276ms | http |
| 159.223.87.50:443 | 否 | 否 | ✓ 1383ms | ✓ 1794ms | ✓ 1947ms | http |
| 91.186.213.124:1081 | ✓ 1025ms | 否 | ✓ 1028ms | 否 | ✓ 1032ms | http |
| 92.118.112.25:1082 | 否 | 否 | ✓ 133ms | ✓ 1288ms | ✓ 780ms | http |
| 77.110.116.93:7443 | ✓ 967ms | 否 | ✓ 1312ms | ✓ 1857ms | 否 | http |
| 164.92.165.209:18080 | ✓ 1898ms | 否 | ✓ 1421ms | 否 | ✓ 1581ms | http |
| 62.133.62.207:1081 | 否 | ✓ 1958ms | ✓ 1478ms | 否 | ✓ 1313ms | http |
| 92.118.112.32:1082 | 否 | ✓ 1255ms | ✓ 125ms | ✓ 1253ms | ✓ 1754ms | http |
| 144.202.14.153:50000 | 否 | ✓ 1084ms | ✓ 154ms | ✓ 1051ms | 否 | http |
| 195.25.20.155:3128 | ✓ 912ms | ✓ 1663ms | ✓ 642ms | 否 | ✓ 1145ms | http |
| 45.88.174.195:8080 | ✓ 1463ms | ✓ 1487ms | ✓ 1496ms | ✓ 1553ms | 否 | http |
| 140.245.238.56:53 | ✓ 1527ms | 否 | ✓ 1408ms | ✓ 1833ms | ✓ 1426ms | http |
| 92.118.112.32:1081 | ✓ 94ms | ✓ 1692ms | ✓ 1013ms | 否 | 否 | http |
| 20.78.118.91:8561 | ✓ 1999ms | ✓ 1923ms | ✓ 1073ms | ✓ 1514ms | ✓ 1361ms | http |
| 20.210.39.153:8561 | 否 | ✓ 1638ms | ✓ 1312ms | ✓ 1579ms | ✓ 1356ms | http |
| 2.26.87.216:1080 | ✓ 1869ms | 否 | ✓ 1287ms | 否 | ✓ 1816ms | http |
| 151.243.153.157:8118 | ✓ 1652ms | ✓ 1921ms | ✓ 799ms | 否 | 否 | http |
| 139.59.105.64:8080 | 否 | 否 | ✓ 1646ms | ✓ 1803ms | ✓ 1589ms | http |
| 45.84.222.25:1080 | ✓ 726ms | 否 | ✓ 1737ms | 否 | ✓ 1294ms | http |
| 185.245.35.88:9443 | ✓ 1433ms | 否 | ✓ 954ms | ✓ 1188ms | ✓ 1833ms | http |
| 62.133.62.17:1081 | ✓ 1387ms | 否 | ✓ 872ms | ✓ 1962ms | 否 | http |
| 116.102.185.188:80 | ✓ 1685ms | 否 | ✓ 1474ms | ✓ 1368ms | ✓ 1097ms | http |
| 154.206.67.83:9000 | ✓ 833ms | 否 | 否 | ✓ 1018ms | ✓ 1862ms | http |
| 119.28.16.210:3128 | ✓ 1111ms | ✓ 1355ms | ✓ 1917ms | 否 | 否 | http |
| 103.69.84.106:3131 | 否 | 否 | ✓ 1501ms | ✓ 1460ms | ✓ 1323ms | http |
| 61.52.131.172:8443 | ✓ 1075ms | ✓ 1966ms | ✓ 1068ms | ✓ 1446ms | ✓ 1143ms | http |
| 144.31.134.103:1080 | 否 | 否 | ✓ 1494ms | ✓ 1433ms | ✓ 1598ms | http |
| 50.114.102.16:8888 | ✓ 1467ms | ✓ 1713ms | 否 | 否 | ✓ 1382ms | http |
| 217.154.155.115:8080 | ✓ 1674ms | ✓ 1498ms | ✓ 1334ms | 否 | 否 | http |
| 89.125.68.33:10000 | ✓ 1561ms | 否 | 否 | ✓ 1708ms | ✓ 1905ms | http |
| 83.147.36.155:8080 | ✓ 1206ms | ✓ 1266ms | ✓ 1077ms | 否 | ✓ 951ms | http |
| 103.39.51.207:8080 | ✓ 1667ms | 否 | 否 | ✓ 1907ms | ✓ 1791ms | http |
| 103.189.250.23:8082 | ✓ 1676ms | 否 | ✓ 1595ms | ✓ 1818ms | ✓ 1714ms | http |
| 111.230.27.213:3128 | ✓ 1144ms | ✓ 1458ms | ✓ 1153ms | ✓ 1327ms | ✓ 1069ms | http |
| 103.56.80.61:8282 | ✓ 1991ms | 否 | ✓ 1840ms | ✓ 1684ms | ✓ 1759ms | http |

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
