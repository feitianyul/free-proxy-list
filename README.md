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

最后更新：2026-04-12 21:30:42 UTC（2026-04-13 05:30:42 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 962ms | ✓ 1069ms | ✓ 841ms | ✓ 1060ms | ✓ 894ms | http |
| 147.161.210.140:8800 | ✓ 1717ms | 否 | ✓ 1220ms | ✓ 1163ms | ✓ 1013ms | http |
| 167.103.115.102:8800 | 否 | ✓ 1889ms | ✓ 1913ms | ✓ 1263ms | ✓ 1200ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1596ms | ✓ 1679ms | ✓ 1292ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1776ms | ✓ 1269ms | 否 | ✓ 1409ms | http |
| 35.225.22.61:80 | ✓ 277ms | 否 | 否 | ✓ 1129ms | ✓ 810ms | http |
| 34.71.229.255:3128 | ✓ 700ms | ✓ 1412ms | ✓ 1238ms | ✓ 1034ms | ✓ 796ms | http |
| 46.30.46.133:3128 | ✓ 365ms | ✓ 1165ms | ✓ 1456ms | 否 | ✓ 1076ms | http |
| 83.219.250.8:62920 | ✓ 771ms | ✓ 1672ms | ✓ 995ms | 否 | ✓ 1277ms | http |
| 95.214.9.93:3128 | ✓ 1665ms | ✓ 1591ms | 否 | ✓ 1536ms | ✓ 1371ms | http |
| 80.250.165.242:3128 | ✓ 1023ms | ✓ 1845ms | ✓ 1444ms | ✓ 1543ms | 否 | http |
| 168.110.52.228:3128 | ✓ 725ms | 否 | 否 | ✓ 1008ms | ✓ 809ms | http |
| 167.103.144.127:8800 | ✓ 1507ms | ✓ 1991ms | ✓ 1517ms | 否 | ✓ 1622ms | http |
| 167.103.31.122:8800 | ✓ 1575ms | 否 | ✓ 1784ms | 否 | ✓ 1791ms | http |
| 120.92.108.86:7890 | ✓ 1422ms | 否 | ✓ 1961ms | ✓ 1910ms | ✓ 1531ms | http |
| 24.144.86.173:1080 | ✓ 1153ms | 否 | ✓ 1876ms | ✓ 1255ms | ✓ 1825ms | http |
| 45.167.125.21:999 | ✓ 739ms | ✓ 1803ms | ✓ 1325ms | ✓ 1592ms | ✓ 1303ms | http |
| 173.212.246.157:3128 | ✓ 1296ms | ✓ 1676ms | ✓ 805ms | 否 | ✓ 1787ms | http |
| 222.228.171.92:8080 | ✓ 1561ms | 否 | ✓ 1881ms | 否 | ✓ 1318ms | http |
| 130.61.30.221:8080 | ✓ 604ms | ✓ 1285ms | ✓ 1126ms | ✓ 1860ms | ✓ 1376ms | http |
| 36.103.198.235:7890 | ✓ 1269ms | ✓ 1455ms | ✓ 1937ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1221ms | 否 | 否 | ✓ 1538ms | ✓ 1256ms | http |
| 8.209.238.110:47701 | ✓ 674ms | 否 | ✓ 854ms | ✓ 1020ms | ✓ 837ms | http |
| 5.104.87.17:8051 | ✓ 931ms | 否 | ✓ 1685ms | 否 | ✓ 935ms | http |
| 37.187.109.70:10111 | ✓ 1308ms | 否 | ✓ 1758ms | 否 | ✓ 1790ms | http |
| 147.161.239.240:8800 | ✓ 624ms | ✓ 1535ms | ✓ 1696ms | ✓ 1467ms | ✓ 1314ms | http |
| 8.219.195.129:1080 | ✓ 1009ms | ✓ 1950ms | ✓ 926ms | ✓ 1265ms | ✓ 991ms | http |
| 36.141.21.200:7890 | ✓ 1072ms | ✓ 1366ms | ✓ 1113ms | ✓ 1451ms | ✓ 1161ms | http |
| 20.118.221.52:3128 | ✓ 937ms | ✓ 1909ms | 否 | ✓ 1840ms | 否 | http |
| 2.27.32.81:3128 | ✓ 1201ms | ✓ 1941ms | ✓ 938ms | 否 | 否 | http |
| 113.160.132.26:8080 | ✓ 1085ms | ✓ 1840ms | 否 | ✓ 1410ms | 否 | http |
| 152.32.132.190:7890 | ✓ 1938ms | 否 | ✓ 1573ms | ✓ 1729ms | ✓ 1974ms | http |
| 45.140.147.155:1082 | ✓ 545ms | ✓ 1309ms | ✓ 1281ms | ✓ 1553ms | ✓ 1001ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1073ms | ✓ 1576ms | ✓ 1173ms | http |
| 180.76.115.231:3128 | 否 | 否 | ✓ 1461ms | ✓ 1482ms | ✓ 1503ms | http |
| 168.222.254.136:8888 | ✓ 689ms | ✓ 1529ms | ✓ 1475ms | 否 | ✓ 1308ms | http |
| 177.234.217.88:999 | ✓ 1429ms | ✓ 1979ms | ✓ 1696ms | ✓ 1875ms | ✓ 1950ms | http |
| 222.129.141.73:9000 | ✓ 1336ms | 否 | ✓ 1084ms | ✓ 1363ms | 否 | http |
| 109.107.179.140:31000 | ✓ 1628ms | 否 | ✓ 765ms | ✓ 1914ms | 否 | http |
| 104.129.203.244:11465 | 否 | ✓ 1136ms | ✓ 1189ms | ✓ 1134ms | ✓ 1016ms | http |
| 194.87.85.207:1080 | ✓ 852ms | 否 | ✓ 1014ms | ✓ 1811ms | ✓ 1560ms | http |
| 217.217.249.160:8080 | ✓ 1702ms | 否 | ✓ 1467ms | 否 | ✓ 1300ms | http |
| 1.231.81.166:3128 | ✓ 1392ms | 否 | 否 | ✓ 1374ms | ✓ 1714ms | http |
| 212.58.132.5:8888 | ✓ 1691ms | 否 | ✓ 1523ms | ✓ 1812ms | ✓ 1594ms | http |
| 185.212.119.154:3128 | ✓ 877ms | ✓ 1641ms | ✓ 1407ms | ✓ 1726ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1306ms | ✓ 1945ms | 否 | ✓ 1912ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1734ms | ✓ 1305ms | 否 | ✓ 1423ms | 否 | http |
| 158.160.215.167:8127 | 否 | ✓ 1497ms | ✓ 920ms | 否 | ✓ 1571ms | http |
| 1.20.169.239:8080 | ✓ 1980ms | 否 | 否 | ✓ 1802ms | ✓ 1913ms | http |
| 116.63.160.98:8899 | ✓ 1028ms | ✓ 1222ms | ✓ 1028ms | ✓ 1313ms | ✓ 1063ms | http |
| 103.235.67.190:80 | ✓ 1055ms | 否 | ✓ 1113ms | ✓ 1449ms | 否 | http |
| 121.230.8.111:1080 | ✓ 1362ms | ✓ 1628ms | ✓ 1393ms | ✓ 1953ms | ✓ 1370ms | http |
| 45.186.6.104:3128 | ✓ 1693ms | ✓ 1871ms | ✓ 1976ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1067ms | 否 | ✓ 1407ms | ✓ 1700ms | 否 | http |
| 171.227.167.109:1006 | ✓ 1895ms | 否 | 否 | ✓ 1925ms | ✓ 1263ms | http |
| 103.155.199.68:7777 | ✓ 1646ms | 否 | ✓ 1610ms | ✓ 1797ms | 否 | http |
| 114.237.77.231:1080 | ✓ 1125ms | ✓ 1403ms | ✓ 1127ms | ✓ 1435ms | ✓ 1146ms | http |
| 213.3.34.39:443 | ✓ 820ms | ✓ 1744ms | ✓ 710ms | ✓ 1973ms | ✓ 1293ms | http |
| 168.222.254.88:3128 | ✓ 583ms | ✓ 1719ms | ✓ 1316ms | ✓ 1830ms | 否 | http |
| 113.176.92.71:3128 | 否 | ✓ 1488ms | ✓ 1421ms | ✓ 1706ms | ✓ 1426ms | http |
| 206.27.173.58:80 | ✓ 710ms | ✓ 1205ms | ✓ 1836ms | ✓ 1502ms | ✓ 929ms | http |
| 93.77.181.116:8888 | ✓ 548ms | ✓ 1743ms | ✓ 503ms | ✓ 1692ms | ✓ 1342ms | http |
| 217.182.195.221:30003 | ✓ 914ms | 否 | ✓ 846ms | ✓ 1845ms | ✓ 1639ms | http |
| 147.45.180.91:8888 | ✓ 654ms | 否 | ✓ 1660ms | ✓ 1633ms | ✓ 1572ms | http |
| 109.120.135.118:3128 | ✓ 897ms | ✓ 1936ms | ✓ 1442ms | 否 | ✓ 1791ms | http |
| 85.239.59.252:7890 | 否 | ✓ 1463ms | ✓ 512ms | ✓ 1612ms | 否 | http |
| 158.160.215.167:8124 | ✓ 1055ms | 否 | ✓ 636ms | ✓ 1841ms | ✓ 1410ms | http |
| 195.26.224.49:3128 | ✓ 1000ms | 否 | ✓ 1030ms | ✓ 1448ms | ✓ 1471ms | http |
| 147.45.186.28:3128 | ✓ 1015ms | ✓ 1398ms | ✓ 1720ms | 否 | ✓ 1819ms | http |
| 5.196.101.18:3128 | ✓ 1450ms | ✓ 1565ms | ✓ 1092ms | 否 | ✓ 1396ms | http |
| 167.71.196.178:80 | ✓ 1101ms | 否 | ✓ 1148ms | ✓ 1258ms | ✓ 997ms | http |
| 138.197.68.35:4857 | 否 | ✓ 1293ms | ✓ 422ms | ✓ 1113ms | 否 | http |
| 45.129.141.143:3128 | ✓ 966ms | ✓ 1713ms | ✓ 1866ms | ✓ 1865ms | ✓ 1735ms | http |
| 180.130.80.196:9003 | 否 | 否 | ✓ 1393ms | ✓ 1409ms | ✓ 1126ms | http |
| 38.180.2.107:3128 | ✓ 1206ms | ✓ 1859ms | ✓ 990ms | 否 | 否 | http |
| 162.240.154.26:3128 | ✓ 1024ms | ✓ 1720ms | ✓ 1842ms | 否 | 否 | http |
| 104.248.151.93:9090 | ✓ 1380ms | 否 | 否 | ✓ 1965ms | ✓ 1979ms | http |
| 62.113.119.14:8080 | ✓ 1133ms | ✓ 1504ms | ✓ 806ms | ✓ 1558ms | ✓ 1045ms | http |
| 45.12.151.226:2829 | ✓ 1491ms | ✓ 1766ms | ✓ 1575ms | 否 | 否 | http |
| 158.160.215.167:8123 | 否 | 否 | ✓ 1759ms | ✓ 1725ms | ✓ 1422ms | http |
| 103.135.102.161:8081 | ✓ 1939ms | 否 | ✓ 1759ms | 否 | ✓ 1702ms | http |
| 103.67.46.225:3125 | ✓ 1805ms | 否 | 否 | ✓ 1877ms | ✓ 1901ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1623ms | ✓ 1855ms | ✓ 1851ms | ✓ 1841ms | http |
| 103.39.51.207:8080 | ✓ 1668ms | 否 | ✓ 1783ms | ✓ 1660ms | ✓ 1961ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1937ms | 否 | ✓ 1929ms | ✓ 1470ms | http |

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
