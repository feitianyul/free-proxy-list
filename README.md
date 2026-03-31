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

最后更新：2026-03-31 06:02:27 UTC（2026-03-31 14:02:27 UTC+8）

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
| 208.87.243.199:7878 | ✓ 377ms | ✓ 870ms | ✓ 633ms | ✓ 900ms | ✓ 665ms | http |
| 95.213.217.168:52004 | ✓ 637ms | ✓ 1568ms | ✓ 556ms | ✓ 1549ms | ✓ 1233ms | http |
| 39.185.46.193:5911 | ✓ 980ms | ✓ 993ms | ✓ 793ms | ✓ 1251ms | ✓ 825ms | http |
| 147.161.210.140:8800 | ✓ 1688ms | 否 | ✓ 758ms | ✓ 936ms | ✓ 991ms | http |
| 1.231.81.166:3128 | ✓ 1715ms | ✓ 1212ms | ✓ 1336ms | ✓ 1368ms | ✓ 1290ms | http |
| 5.104.87.17:8051 | ✓ 973ms | 否 | 否 | ✓ 1806ms | ✓ 919ms | http |
| 167.103.115.102:8800 | ✓ 1815ms | 否 | ✓ 1065ms | ✓ 1168ms | ✓ 1142ms | http |
| 113.160.132.26:8080 | ✓ 1504ms | 否 | 否 | ✓ 1420ms | ✓ 1256ms | http |
| 34.101.184.164:3128 | ✓ 1653ms | 否 | ✓ 885ms | ✓ 1480ms | ✓ 1043ms | http |
| 167.103.34.108:8800 | ✓ 1858ms | 否 | ✓ 1285ms | ✓ 1545ms | ✓ 1481ms | http |
| 101.47.73.135:3128 | ✓ 895ms | 否 | ✓ 1133ms | ✓ 1476ms | ✓ 1360ms | http |
| 45.167.124.52:8080 | ✓ 542ms | 否 | ✓ 551ms | ✓ 1562ms | ✓ 1301ms | http |
| 35.225.22.61:80 | ✓ 1021ms | 否 | ✓ 411ms | ✓ 1125ms | ✓ 657ms | http |
| 43.99.54.236:5555 | ✓ 1467ms | 否 | ✓ 765ms | ✓ 960ms | ✓ 817ms | http |
| 120.92.212.16:7890 | ✓ 1063ms | 否 | ✓ 1093ms | ✓ 1377ms | ✓ 1106ms | http |
| 217.217.249.160:8080 | ✓ 1124ms | 否 | ✓ 1102ms | 否 | ✓ 1227ms | http |
| 31.192.106.135:8010 | ✓ 1216ms | 否 | ✓ 1971ms | ✓ 1748ms | ✓ 1677ms | http |
| 103.84.95.54:7890 | ✓ 1476ms | 否 | ✓ 906ms | 否 | ✓ 1031ms | http |
| 150.107.140.238:3128 | ✓ 995ms | 否 | ✓ 1984ms | ✓ 1270ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1529ms | 否 | ✓ 1185ms | ✓ 1449ms | ✓ 1295ms | http |
| 42.96.16.158:1311 | ✓ 1609ms | 否 | ✓ 1095ms | ✓ 1366ms | ✓ 1019ms | http |
| 45.12.151.226:2829 | ✓ 712ms | ✓ 1959ms | ✓ 1138ms | 否 | 否 | http |
| 167.103.31.122:8800 | ✓ 1878ms | 否 | ✓ 1376ms | 否 | ✓ 1761ms | http |
| 209.126.84.232:8888 | ✓ 1474ms | ✓ 1554ms | ✓ 1312ms | 否 | 否 | http |
| 91.238.123.111:8000 | ✓ 458ms | ✓ 1409ms | ✓ 1209ms | 否 | 否 | http |
| 91.238.123.230:8000 | ✓ 441ms | ✓ 1402ms | ✓ 1233ms | 否 | 否 | http |
| 133.242.138.34:8100 | 否 | ✓ 1169ms | ✓ 750ms | ✓ 1370ms | ✓ 1634ms | http |
| 40.177.212.89:30309 | ✓ 1461ms | 否 | ✓ 1104ms | 否 | ✓ 1897ms | http |
| 35.181.173.74:7598 | ✓ 1751ms | 否 | ✓ 610ms | 否 | ✓ 1588ms | http |
| 101.43.127.100:8877 | 否 | 否 | ✓ 1838ms | ✓ 1625ms | ✓ 1241ms | http |
| 164.90.201.255:3128 | 否 | ✓ 1741ms | ✓ 1936ms | ✓ 1604ms | ✓ 1313ms | http |
| 38.34.179.100:8449 | ✓ 1545ms | ✓ 968ms | ✓ 649ms | ✓ 1753ms | ✓ 754ms | http |
| 5.102.109.41:999 | 否 | ✓ 1652ms | ✓ 1841ms | 否 | ✓ 1795ms | http |
| 195.123.213.129:1080 | ✓ 1598ms | 否 | ✓ 1570ms | 否 | ✓ 1196ms | http |
| 103.158.62.186:8088 | ✓ 1800ms | 否 | ✓ 1848ms | ✓ 1877ms | ✓ 1901ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1195ms | ✓ 1784ms | ✓ 1022ms | http |
| 120.92.212.16:8890 | ✓ 1153ms | ✓ 1360ms | 否 | ✓ 1324ms | ✓ 1097ms | http |
| 8.219.97.248:80 | ✓ 1453ms | 否 | 否 | ✓ 1900ms | ✓ 1872ms | http |
| 167.71.196.28:8080 | ✓ 1471ms | 否 | ✓ 1272ms | 否 | ✓ 952ms | http |
| 147.161.239.240:8800 | ✓ 1771ms | ✓ 1596ms | ✓ 995ms | ✓ 1727ms | ✓ 1402ms | http |
| 38.34.179.202:8452 | ✓ 1044ms | ✓ 1529ms | ✓ 640ms | ✓ 1598ms | ✓ 1266ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 764ms | ✓ 1031ms | ✓ 903ms | http |
| 121.230.9.101:1080 | ✓ 1438ms | 否 | ✓ 1371ms | ✓ 1853ms | ✓ 1331ms | http |
| 181.78.44.63:999 | ✓ 1681ms | 否 | ✓ 1095ms | ✓ 1644ms | ✓ 1426ms | http |
| 103.210.119.107:7777 | ✓ 1903ms | 否 | 否 | ✓ 1841ms | ✓ 1644ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 931ms | ✓ 1398ms | ✓ 1007ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1480ms | ✓ 1122ms | ✓ 1485ms | 否 | http |
| 177.234.217.88:999 | ✓ 1574ms | ✓ 1908ms | ✓ 1893ms | 否 | 否 | http |
| 24.199.124.151:3128 | ✓ 426ms | 否 | ✓ 555ms | ✓ 910ms | ✓ 696ms | http |
| 121.230.8.111:1080 | 否 | ✓ 1884ms | ✓ 1483ms | 否 | ✓ 1209ms | http |
| 185.191.236.162:3128 | ✓ 1221ms | 否 | ✓ 1526ms | 否 | ✓ 1717ms | http |
| 47.74.226.8:5001 | 否 | 否 | ✓ 1188ms | ✓ 1483ms | ✓ 1420ms | http |
| 106.10.55.212:1121 | ✓ 1982ms | 否 | ✓ 1877ms | 否 | ✓ 1295ms | http |
| 65.108.203.35:18080 | ✓ 1024ms | ✓ 1774ms | ✓ 969ms | 否 | 否 | http |
| 185.114.73.2:1080 | ✓ 1015ms | ✓ 1533ms | 否 | ✓ 1558ms | ✓ 1232ms | http |
| 37.187.109.70:10111 | ✓ 1255ms | ✓ 1335ms | ✓ 730ms | 否 | ✓ 1351ms | http |
| 158.160.215.167:8124 | ✓ 700ms | ✓ 1893ms | ✓ 1993ms | 否 | 否 | http |
| 121.126.185.63:25152 | ✓ 1972ms | 否 | ✓ 1489ms | 否 | ✓ 1508ms | http |
| 38.34.179.74:8449 | ✓ 1256ms | ✓ 1928ms | ✓ 1333ms | ✓ 1271ms | ✓ 1515ms | http |
| 197.164.101.11:1976 | ✓ 1551ms | 否 | ✓ 1288ms | ✓ 1999ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1264ms | 否 | ✓ 1356ms | 否 | ✓ 1599ms | http |
| 45.140.147.155:1081 | ✓ 1398ms | 否 | 否 | ✓ 1156ms | ✓ 1885ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1909ms | ✓ 1437ms | 否 | ✓ 1734ms | http |
| 104.243.46.122:3128 | ✓ 1388ms | 否 | ✓ 1723ms | ✓ 1647ms | ✓ 1258ms | http |
| 103.154.146.58:8080 | ✓ 1896ms | 否 | ✓ 1563ms | ✓ 1823ms | ✓ 1808ms | http |
| 157.230.220.25:4857 | 否 | 否 | ✓ 1569ms | ✓ 996ms | ✓ 1012ms | http |
| 187.245.214.11:999 | 否 | ✓ 1509ms | ✓ 1519ms | ✓ 1757ms | 否 | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1816ms | ✓ 1914ms | ✓ 1801ms | http |
| 3.99.169.21:27755 | ✓ 1358ms | 否 | ✓ 1481ms | 否 | ✓ 1733ms | http |
| 45.15.158.60:2222 | ✓ 727ms | 否 | ✓ 1542ms | 否 | ✓ 1576ms | http |
| 45.129.141.143:3128 | 否 | ✓ 1741ms | ✓ 1770ms | 否 | ✓ 1811ms | http |
| 197.164.101.13:1981 | 否 | 否 | ✓ 1228ms | ✓ 1972ms | ✓ 1857ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1541ms | ✓ 1275ms | ✓ 1558ms | http |
| 103.143.197.218:8008 | 否 | 否 | ✓ 1700ms | ✓ 1634ms | ✓ 1559ms | http |

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
