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

最后更新：2026-03-13 07:56:46 UTC（2026-03-13 15:56:46 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 831ms | ✓ 1108ms | ✓ 1129ms | ✓ 1103ms | ✓ 884ms | http |
| 205.209.118.30:3138 | ✓ 1930ms | ✓ 981ms | ✓ 1807ms | ✓ 1212ms | ✓ 857ms | http |
| 113.160.132.26:8080 | ✓ 1705ms | ✓ 1565ms | ✓ 1312ms | ✓ 1851ms | ✓ 1122ms | http |
| 147.45.251.242:8888 | ✓ 1402ms | 否 | ✓ 1976ms | 否 | ✓ 1847ms | http |
| 45.140.147.155:1081 | ✓ 444ms | 否 | ✓ 1829ms | ✓ 1766ms | 否 | http |
| 165.225.113.220:11462 | ✓ 1344ms | 否 | 否 | ✓ 1199ms | ✓ 940ms | http |
| 165.225.113.220:10958 | ✓ 1333ms | 否 | ✓ 848ms | 否 | ✓ 943ms | http |
| 46.183.25.8:443 | ✓ 1271ms | 否 | ✓ 1066ms | ✓ 1328ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1102ms | ✓ 1199ms | ✓ 1368ms | ✓ 1791ms | ✓ 903ms | http |
| 14.225.211.139:7890 | ✓ 1845ms | 否 | ✓ 1909ms | 否 | ✓ 1688ms | http |
| 103.113.70.189:1081 | ✓ 452ms | ✓ 1081ms | 否 | ✓ 1111ms | ✓ 853ms | http |
| 81.70.169.194:80 | 否 | ✓ 1392ms | ✓ 1180ms | ✓ 1400ms | 否 | http |
| 171.251.172.78:5102 | 否 | 否 | ✓ 1577ms | ✓ 1797ms | ✓ 1815ms | http |
| 45.136.130.191:8443 | ✓ 742ms | 否 | ✓ 245ms | ✓ 871ms | ✓ 680ms | http |
| 115.231.181.40:8128 | ✓ 1463ms | 否 | ✓ 1059ms | 否 | ✓ 1439ms | http |
| 45.136.130.188:8443 | ✓ 248ms | ✓ 894ms | ✓ 248ms | ✓ 1059ms | ✓ 669ms | http |
| 45.136.131.63:8443 | ✓ 258ms | ✓ 1109ms | ✓ 256ms | ✓ 1032ms | ✓ 1704ms | http |
| 45.136.130.175:8443 | ✓ 269ms | ✓ 1554ms | ✓ 807ms | ✓ 903ms | ✓ 884ms | http |
| 168.235.110.63:3128 | 否 | ✓ 1364ms | ✓ 1139ms | ✓ 1047ms | ✓ 969ms | http |
| 152.42.213.210:8080 | ✓ 1488ms | 否 | ✓ 965ms | ✓ 1226ms | ✓ 967ms | http |
| 192.71.213.85:5678 | ✓ 1380ms | 否 | ✓ 1697ms | ✓ 1864ms | 否 | http |
| 190.9.109.198:999 | 否 | ✓ 1525ms | ✓ 1222ms | ✓ 1327ms | ✓ 1091ms | http |
| 171.251.172.78:5107 | ✓ 1906ms | 否 | 否 | ✓ 1716ms | ✓ 1546ms | http |
| 45.136.131.47:8443 | 否 | ✓ 855ms | ✓ 1137ms | 否 | ✓ 660ms | http |
| 210.223.44.230:3128 | ✓ 1978ms | ✓ 1135ms | ✓ 797ms | ✓ 1044ms | ✓ 835ms | http |
| 45.136.198.40:3128 | ✓ 1091ms | ✓ 1819ms | ✓ 1669ms | ✓ 1928ms | 否 | http |
| 116.80.49.161:3172 | ✓ 1942ms | 否 | ✓ 1840ms | ✓ 1963ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1902ms | 否 | ✓ 1514ms | 否 | ✓ 1842ms | http |
| 138.124.53.25:7443 | ✓ 564ms | 否 | ✓ 1827ms | ✓ 1345ms | 否 | http |
| 160.250.5.22:1 | ✓ 1048ms | 否 | ✓ 1748ms | ✓ 1431ms | ✓ 1415ms | http |
| 162.248.165.72:1080 | ✓ 1012ms | 否 | ✓ 1917ms | 否 | ✓ 1704ms | http |
| 160.250.4.13:1 | 否 | 否 | ✓ 1438ms | ✓ 1860ms | ✓ 1687ms | http |
| 165.225.113.220:11596 | ✓ 839ms | 否 | ✓ 866ms | 否 | ✓ 950ms | http |
| 35.225.22.61:80 | ✓ 972ms | ✓ 1597ms | ✓ 452ms | 否 | ✓ 769ms | http |
| 120.92.212.16:8890 | ✓ 1077ms | 否 | ✓ 1060ms | ✓ 1419ms | 否 | http |
| 45.136.130.223:8443 | ✓ 444ms | ✓ 1693ms | ✓ 751ms | ✓ 875ms | ✓ 657ms | http |
| 165.227.5.10:8888 | ✓ 833ms | ✓ 1474ms | ✓ 482ms | ✓ 1289ms | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1630ms | ✓ 1056ms | ✓ 1384ms | 否 | http |
| 101.43.255.96:80 | ✓ 1293ms | 否 | ✓ 1158ms | ✓ 1430ms | ✓ 1454ms | http |
| 59.46.216.131:30001 | ✓ 1458ms | 否 | ✓ 1245ms | ✓ 1549ms | 否 | http |
| 178.236.245.59:3128 | ✓ 961ms | ✓ 1541ms | ✓ 717ms | ✓ 1892ms | ✓ 1236ms | http |
| 178.236.245.17:3128 | ✓ 1000ms | 否 | ✓ 638ms | ✓ 1666ms | ✓ 1209ms | http |
| 152.42.213.210:443 | ✓ 1557ms | 否 | ✓ 1588ms | ✓ 1527ms | ✓ 997ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 1598ms | ✓ 1242ms | ✓ 1006ms | http |
| 103.139.138.194:3128 | ✓ 1844ms | 否 | 否 | ✓ 1605ms | ✓ 1368ms | http |
| 201.150.116.32:999 | ✓ 946ms | 否 | 否 | ✓ 1427ms | ✓ 1294ms | http |
| 162.240.154.26:3128 | ✓ 1345ms | ✓ 1578ms | 否 | ✓ 1909ms | 否 | http |
| 67.169.98.211:443 | ✓ 1329ms | 否 | ✓ 1746ms | ✓ 1623ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1578ms | ✓ 1650ms | ✓ 634ms | ✓ 1760ms | 否 | http |
| 165.225.113.220:10819 | ✓ 1522ms | 否 | 否 | ✓ 1186ms | ✓ 951ms | http |
| 171.251.172.78:5104 | ✓ 1783ms | 否 | 否 | ✓ 1891ms | ✓ 1681ms | http |
| 160.238.65.7:3128 | ✓ 894ms | ✓ 1615ms | 否 | ✓ 1332ms | 否 | http |
| 144.124.227.90:21074 | ✓ 899ms | 否 | ✓ 1774ms | 否 | ✓ 1747ms | http |
| 172.212.68.37:3128 | ✓ 355ms | ✓ 1754ms | ✓ 1535ms | ✓ 1233ms | ✓ 1008ms | http |
| 204.48.31.203:80 | ✓ 779ms | ✓ 1471ms | ✓ 1213ms | ✓ 1456ms | ✓ 1155ms | http |
| 185.191.236.162:3128 | ✓ 1262ms | 否 | ✓ 1718ms | 否 | ✓ 1718ms | http |
| 43.167.227.161:1080 | ✓ 1579ms | 否 | ✓ 840ms | 否 | ✓ 870ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1497ms | ✓ 1515ms | ✓ 1678ms | ✓ 1525ms | http |
| 51.79.207.21:8080 | ✓ 1051ms | 否 | ✓ 1899ms | ✓ 1896ms | ✓ 1742ms | http |
| 45.186.6.104:3128 | ✓ 1831ms | ✓ 1692ms | ✓ 1658ms | 否 | 否 | http |
| 138.124.53.221:443 | ✓ 1030ms | ✓ 1665ms | ✓ 533ms | ✓ 1395ms | ✓ 1089ms | http |
| 45.136.130.239:8443 | 否 | ✓ 936ms | 否 | ✓ 1141ms | ✓ 859ms | http |
| 185.191.236.162:8080 | ✓ 1033ms | 否 | ✓ 1152ms | ✓ 1669ms | ✓ 1687ms | http |
| 160.238.65.5:3128 | ✓ 1028ms | ✓ 1285ms | 否 | 否 | ✓ 1368ms | http |
| 160.238.65.8:3128 | ✓ 1030ms | ✓ 1224ms | 否 | 否 | ✓ 1433ms | http |
| 160.238.65.9:3128 | ✓ 1028ms | ✓ 1763ms | 否 | ✓ 1694ms | ✓ 1216ms | http |
| 160.238.65.3:3128 | ✓ 1027ms | ✓ 1219ms | 否 | 否 | ✓ 1444ms | http |
| 160.238.65.2:3128 | ✓ 1024ms | 否 | 否 | ✓ 1464ms | ✓ 1198ms | http |
| 160.238.65.6:3128 | ✓ 1027ms | 否 | 否 | ✓ 1439ms | ✓ 1224ms | http |
| 160.238.65.4:3128 | ✓ 1024ms | 否 | 否 | ✓ 1438ms | ✓ 1245ms | http |
| 61.76.95.217:40088 | ✓ 1098ms | 否 | ✓ 1203ms | ✓ 1470ms | ✓ 1277ms | http |
| 103.52.115.171:3128 | ✓ 1619ms | 否 | ✓ 966ms | ✓ 1330ms | ✓ 1169ms | http |
| 193.122.96.242:3128 | ✓ 1091ms | 否 | 否 | ✓ 1406ms | ✓ 1090ms | http |
| 39.104.201.40:7890 | ✓ 1004ms | 否 | ✓ 1298ms | 否 | ✓ 1112ms | http |
| 121.230.9.5:1080 | 否 | ✓ 1897ms | ✓ 1814ms | ✓ 1925ms | ✓ 1470ms | http |
| 211.171.114.154:3128 | ✓ 927ms | 否 | ✓ 1881ms | 否 | ✓ 1112ms | http |
| 14.225.222.213:7890 | ✓ 1444ms | 否 | ✓ 1118ms | ✓ 1336ms | ✓ 1035ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1598ms | ✓ 1161ms | ✓ 1464ms | ✓ 1174ms | http |
| 18.100.127.123:15805 | ✓ 1882ms | 否 | ✓ 1680ms | 否 | ✓ 1895ms | http |
| 165.225.113.220:11845 | ✓ 1436ms | 否 | ✓ 835ms | ✓ 1172ms | 否 | http |
| 62.234.206.73:3128 | ✓ 1065ms | 否 | 否 | ✓ 1807ms | ✓ 1196ms | http |
| 165.225.113.220:11083 | 否 | 否 | ✓ 848ms | ✓ 1188ms | ✓ 932ms | http |

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
