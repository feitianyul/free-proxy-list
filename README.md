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

最后更新：2026-03-01 06:52:44 UTC（2026-03-01 14:52:44 UTC+8）

**代理总数：81**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 148.135.85.87:1080 | ✓ 977ms | 否 | 否 | ✓ 1019ms | ✓ 907ms | http |
| 3.213.157.4:3128 | ✓ 338ms | 否 | ✓ 1244ms | ✓ 1412ms | 否 | http |
| 205.209.118.30:3138 | ✓ 1608ms | 否 | ✓ 983ms | ✓ 1382ms | ✓ 1083ms | http |
| 52.188.28.218:3128 | ✓ 567ms | 否 | ✓ 1009ms | ✓ 1904ms | 否 | http |
| 74.208.234.198:443 | ✓ 328ms | 否 | ✓ 624ms | ✓ 1039ms | ✓ 805ms | http |
| 120.92.212.16:7890 | ✓ 873ms | ✓ 1165ms | ✓ 983ms | ✓ 1173ms | ✓ 941ms | http |
| 15.204.233.75:3128 | ✓ 665ms | 否 | ✓ 1869ms | 否 | ✓ 1379ms | http |
| 177.243.209.133:999 | ✓ 1059ms | 否 | ✓ 578ms | ✓ 1271ms | ✓ 1125ms | http |
| 59.46.216.131:30001 | ✓ 1034ms | 否 | 否 | ✓ 1364ms | ✓ 1115ms | http |
| 165.227.5.10:8888 | ✓ 1286ms | ✓ 629ms | ✓ 302ms | 否 | ✓ 518ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1826ms | ✓ 758ms | ✓ 702ms | http |
| 120.92.212.16:8890 | ✓ 1192ms | ✓ 1148ms | ✓ 1116ms | ✓ 1376ms | 否 | http |
| 35.225.22.61:80 | ✓ 800ms | 否 | ✓ 432ms | 否 | ✓ 1149ms | http |
| 36.147.78.166:80 | ✓ 1648ms | ✓ 1588ms | ✓ 1875ms | ✓ 1725ms | 否 | http |
| 103.104.99.29:80 | ✓ 1792ms | 否 | ✓ 1591ms | ✓ 1453ms | ✓ 1369ms | http |
| 103.104.99.89:80 | ✓ 1792ms | 否 | ✓ 1674ms | ✓ 1406ms | ✓ 1369ms | http |
| 150.107.140.238:3128 | ✓ 1795ms | 否 | ✓ 1843ms | ✓ 1680ms | 否 | http |
| 168.235.110.63:3128 | ✓ 539ms | ✓ 1192ms | 否 | ✓ 1362ms | ✓ 1028ms | http |
| 142.171.85.32:1080 | ✓ 638ms | 否 | ✓ 1919ms | ✓ 1059ms | 否 | http |
| 81.70.169.194:80 | ✓ 1030ms | ✓ 1484ms | ✓ 846ms | ✓ 1273ms | ✓ 917ms | http |
| 101.47.73.135:3128 | ✓ 1925ms | 否 | 否 | ✓ 1072ms | ✓ 1022ms | http |
| 136.226.254.24:11719 | ✓ 1912ms | ✓ 1995ms | ✓ 1804ms | 否 | 否 | http |
| 103.113.70.189:1081 | 否 | ✓ 1982ms | 否 | ✓ 1625ms | ✓ 972ms | http |
| 136.226.254.24:11197 | ✓ 1911ms | 否 | ✓ 1210ms | ✓ 1740ms | 否 | http |
| 136.226.254.24:11933 | ✓ 1911ms | 否 | ✓ 1219ms | ✓ 1762ms | 否 | http |
| 136.226.254.24:12566 | ✓ 1910ms | 否 | ✓ 1218ms | ✓ 1771ms | 否 | http |
| 136.226.254.24:11819 | ✓ 1900ms | 否 | ✓ 1219ms | ✓ 1772ms | 否 | http |
| 136.226.254.24:11309 | ✓ 1903ms | 否 | ✓ 1645ms | ✓ 1505ms | 否 | http |
| 136.226.254.24:11909 | ✓ 1906ms | ✓ 1983ms | ✓ 1595ms | ✓ 1569ms | 否 | http |
| 101.43.255.96:80 | ✓ 948ms | ✓ 1959ms | ✓ 1580ms | ✓ 1148ms | ✓ 1872ms | http |
| 136.226.254.24:11814 | ✓ 1912ms | 否 | ✓ 1346ms | ✓ 1896ms | 否 | http |
| 136.226.254.24:12709 | ✓ 1903ms | 否 | ✓ 1219ms | ✓ 1986ms | 否 | http |
| 121.204.158.249:3128 | ✓ 1004ms | ✓ 1204ms | ✓ 1822ms | ✓ 1335ms | ✓ 946ms | http |
| 115.231.181.40:8128 | ✓ 918ms | ✓ 1807ms | 否 | ✓ 1161ms | ✓ 930ms | http |
| 195.123.209.48:3128 | ✓ 1124ms | 否 | ✓ 1405ms | 否 | ✓ 1988ms | http |
| 144.31.69.170:1080 | ✓ 1101ms | 否 | ✓ 1879ms | 否 | ✓ 1746ms | http |
| 31.59.129.75:8080 | ✓ 1111ms | ✓ 1599ms | 否 | 否 | ✓ 1054ms | http |
| 139.159.99.242:8080 | ✓ 811ms | ✓ 945ms | 否 | 否 | ✓ 769ms | http |
| 62.234.206.73:3128 | ✓ 1230ms | ✓ 1436ms | ✓ 1633ms | ✓ 1217ms | ✓ 914ms | http |
| 45.140.147.155:1081 | ✓ 692ms | 否 | ✓ 1573ms | ✓ 1806ms | ✓ 1305ms | http |
| 2.56.178.131:443 | ✓ 1076ms | 否 | ✓ 1063ms | 否 | ✓ 1840ms | http |
| 136.226.254.24:12500 | ✓ 965ms | ✓ 1946ms | ✓ 875ms | ✓ 1558ms | ✓ 1201ms | http |
| 136.226.254.24:10919 | ✓ 946ms | 否 | ✓ 873ms | ✓ 1572ms | ✓ 1164ms | http |
| 5.134.118.80:3128 | ✓ 1721ms | 否 | ✓ 1557ms | 否 | ✓ 1736ms | http |
| 136.226.254.24:10849 | ✓ 1010ms | 否 | ✓ 881ms | ✓ 1561ms | ✓ 1206ms | http |
| 136.226.254.24:12360 | ✓ 895ms | 否 | ✓ 873ms | ✓ 1532ms | 否 | http |
| 23.236.65.234:38080 | ✓ 688ms | ✓ 1209ms | ✓ 1269ms | ✓ 680ms | ✓ 653ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 778ms | ✓ 1118ms | ✓ 816ms | http |
| 198.23.236.47:1111 | 否 | ✓ 805ms | ✓ 875ms | ✓ 731ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1915ms | 否 | ✓ 1712ms | 否 | ✓ 1828ms | http |
| 185.243.218.43:49153 | ✓ 1182ms | ✓ 1876ms | ✓ 1875ms | 否 | ✓ 1878ms | http |
| 136.226.254.24:10517 | 否 | ✓ 1912ms | ✓ 1822ms | ✓ 1937ms | 否 | http |
| 136.226.254.24:12265 | 否 | 否 | ✓ 1549ms | ✓ 1925ms | ✓ 1369ms | http |
| 136.226.254.24:10504 | 否 | 否 | ✓ 1563ms | ✓ 1920ms | ✓ 1485ms | http |
| 136.226.254.24:11342 | 否 | 否 | ✓ 1834ms | ✓ 1577ms | ✓ 1688ms | http |
| 5.75.201.136:1080 | ✓ 641ms | ✓ 1506ms | 否 | 否 | ✓ 1462ms | http |
| 136.226.254.24:12501 | ✓ 1612ms | 否 | ✓ 964ms | ✓ 1710ms | ✓ 1342ms | http |
| 136.226.254.24:12457 | ✓ 1613ms | 否 | ✓ 963ms | ✓ 1758ms | ✓ 1337ms | http |
| 136.226.254.24:12511 | ✓ 1611ms | 否 | ✓ 1065ms | ✓ 1718ms | ✓ 1307ms | http |
| 136.226.254.24:11742 | ✓ 1610ms | 否 | ✓ 966ms | ✓ 1810ms | ✓ 1378ms | http |
| 147.45.216.148:1080 | ✓ 1511ms | 否 | ✓ 1542ms | 否 | ✓ 1911ms | http |
| 36.147.78.166:443 | 否 | ✓ 1574ms | 否 | ✓ 1805ms | ✓ 1672ms | http |
| 180.127.149.244:1080 | 否 | ✓ 1227ms | ✓ 1671ms | ✓ 1365ms | ✓ 1976ms | http |
| 136.226.254.24:80 | ✓ 1550ms | 否 | ✓ 1574ms | ✓ 1599ms | ✓ 1454ms | http |
| 62.113.119.14:8080 | ✓ 1099ms | 否 | ✓ 1799ms | ✓ 1692ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1156ms | ✓ 1475ms | ✓ 1391ms | ✓ 1460ms | ✓ 976ms | http |
| 103.215.36.88:15556 | ✓ 1737ms | ✓ 1860ms | ✓ 1209ms | ✓ 1289ms | 否 | http |
| 103.215.36.88:16650 | 否 | 否 | ✓ 1334ms | ✓ 1797ms | ✓ 1036ms | http |
| 103.189.139.254:8080 | ✓ 1936ms | 否 | ✓ 1658ms | ✓ 1761ms | ✓ 1769ms | http |
| 34.101.184.164:3128 | ✓ 1639ms | 否 | ✓ 1325ms | ✓ 1817ms | 否 | http |
| 35.234.17.221:8080 | 否 | ✓ 1441ms | 否 | ✓ 991ms | ✓ 935ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1583ms | ✓ 1286ms | ✓ 1777ms | http |
| 107.174.133.10:3128 | ✓ 251ms | 否 | ✓ 678ms | ✓ 793ms | 否 | http |
| 47.101.149.27:9010 | 否 | ✓ 1167ms | ✓ 1275ms | ✓ 1435ms | ✓ 1258ms | http |
| 83.219.250.8:62920 | ✓ 839ms | 否 | ✓ 1625ms | 否 | ✓ 1872ms | http |
| 57.128.188.167:9196 | ✓ 1778ms | 否 | ✓ 1852ms | 否 | ✓ 1876ms | http |
| 180.127.149.245:1080 | ✓ 1438ms | 否 | ✓ 1598ms | 否 | ✓ 1075ms | http |
| 14.56.107.244:3128 | ✓ 1577ms | ✓ 1063ms | ✓ 1173ms | ✓ 962ms | ✓ 725ms | http |
| 103.215.36.88:15088 | ✓ 1758ms | 否 | ✓ 1261ms | 否 | ✓ 1745ms | http |
| 123.20.24.166:8118 | 否 | 否 | ✓ 1298ms | ✓ 1517ms | ✓ 1042ms | http |
| 121.128.121.214:3128 | ✓ 1405ms | 否 | ✓ 1226ms | 否 | ✓ 1741ms | http |

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
