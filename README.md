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

最后更新：2026-02-28 03:07:25 UTC（2026-02-28 11:07:25 UTC+8）

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
| 205.209.118.30:3138 | ✓ 1059ms | ✓ 1887ms | ✓ 852ms | ✓ 1123ms | ✓ 868ms | http |
| 3.213.157.4:3128 | ✓ 1053ms | 否 | ✓ 364ms | ✓ 1431ms | ✓ 1232ms | http |
| 20.27.15.111:8561 | ✓ 583ms | ✓ 995ms | ✓ 672ms | ✓ 969ms | ✓ 725ms | http |
| 20.27.11.248:8561 | ✓ 593ms | ✓ 1053ms | ✓ 626ms | ✓ 970ms | ✓ 777ms | http |
| 168.235.110.63:3128 | ✓ 1489ms | ✓ 1369ms | ✓ 883ms | ✓ 1076ms | ✓ 819ms | http |
| 20.27.14.220:8561 | ✓ 578ms | 否 | ✓ 625ms | ✓ 912ms | ✓ 716ms | http |
| 20.78.118.91:8561 | ✓ 811ms | ✓ 1493ms | ✓ 734ms | ✓ 986ms | ✓ 814ms | http |
| 20.78.26.206:8561 | ✓ 849ms | ✓ 1428ms | ✓ 833ms | ✓ 979ms | ✓ 807ms | http |
| 20.210.39.153:8561 | ✓ 803ms | ✓ 1697ms | ✓ 685ms | ✓ 937ms | ✓ 783ms | http |
| 14.56.107.244:3128 | ✓ 702ms | 否 | ✓ 1728ms | ✓ 1329ms | 否 | http |
| 45.136.198.40:3128 | ✓ 769ms | ✓ 1594ms | 否 | 否 | ✓ 1616ms | http |
| 138.124.53.25:7443 | ✓ 681ms | 否 | 否 | ✓ 1952ms | ✓ 1782ms | http |
| 132.145.93.138:1080 | ✓ 1400ms | 否 | 否 | ✓ 1384ms | ✓ 1483ms | http |
| 61.72.110.24:3128 | ✓ 1746ms | ✓ 1852ms | ✓ 962ms | 否 | 否 | http |
| 104.238.30.50:59741 | ✓ 1702ms | 否 | ✓ 1836ms | 否 | ✓ 1967ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1567ms | ✓ 1060ms | ✓ 1328ms | ✓ 1079ms | http |
| 72.56.59.62:63133 | ✓ 1605ms | 否 | ✓ 1644ms | 否 | ✓ 1842ms | http |
| 104.238.30.40:59741 | ✓ 1697ms | 否 | ✓ 1871ms | 否 | ✓ 1972ms | http |
| 35.225.22.61:80 | ✓ 847ms | 否 | ✓ 362ms | 否 | ✓ 1195ms | http |
| 121.237.181.137:8888 | ✓ 1011ms | ✓ 1366ms | ✓ 958ms | ✓ 1403ms | ✓ 975ms | http |
| 195.123.209.48:3128 | ✓ 1028ms | ✓ 1913ms | ✓ 1579ms | 否 | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1438ms | ✓ 1478ms | ✓ 1361ms | http |
| 104.238.30.58:63744 | ✓ 1731ms | 否 | ✓ 1742ms | 否 | ✓ 1967ms | http |
| 72.56.59.56:63127 | ✓ 1563ms | 否 | ✓ 1616ms | 否 | ✓ 1839ms | http |
| 72.56.50.17:59787 | ✓ 1519ms | 否 | ✓ 1964ms | 否 | ✓ 1806ms | http |
| 81.177.48.54:2080 | ✓ 861ms | 否 | ✓ 780ms | ✓ 1779ms | ✓ 1808ms | http |
| 104.238.30.63:63744 | ✓ 1762ms | 否 | ✓ 1843ms | 否 | ✓ 1999ms | http |
| 147.45.216.148:1080 | ✓ 1252ms | 否 | ✓ 1933ms | 否 | ✓ 1303ms | http |
| 52.188.28.218:3128 | ✓ 1409ms | 否 | 否 | ✓ 1750ms | ✓ 1125ms | http |
| 61.72.110.54:3128 | ✓ 1716ms | 否 | 否 | ✓ 1603ms | ✓ 1344ms | http |
| 103.84.95.54:7890 | ✓ 1362ms | 否 | 否 | ✓ 1148ms | ✓ 778ms | http |
| 101.43.255.96:80 | ✓ 1079ms | ✓ 1530ms | ✓ 1090ms | ✓ 1453ms | ✓ 1129ms | http |
| 81.70.169.194:80 | ✓ 1066ms | ✓ 1449ms | ✓ 1102ms | 否 | ✓ 1155ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1389ms | ✓ 1104ms | ✓ 1632ms | ✓ 1312ms | http |
| 59.46.216.131:30001 | ✓ 1204ms | ✓ 1531ms | 否 | 否 | ✓ 1169ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1889ms | ✓ 1693ms | ✓ 1692ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1947ms | ✓ 1655ms | ✓ 1683ms | http |
| 45.125.67.37:8443 | ✓ 1630ms | 否 | ✓ 1028ms | ✓ 1290ms | ✓ 1098ms | http |
| 36.147.78.166:80 | ✓ 1956ms | 否 | 否 | ✓ 1788ms | ✓ 1822ms | http |
| 115.231.181.40:8128 | ✓ 1059ms | 否 | ✓ 1089ms | ✓ 1347ms | 否 | http |
| 104.238.30.86:63900 | ✓ 1769ms | 否 | ✓ 1780ms | 否 | ✓ 1999ms | http |
| 91.238.104.171:2023 | ✓ 639ms | ✓ 1641ms | ✓ 594ms | ✓ 1498ms | ✓ 1158ms | http |
| 103.113.70.189:1081 | 否 | ✓ 960ms | 否 | ✓ 1152ms | ✓ 807ms | http |
| 38.180.2.107:3128 | ✓ 800ms | 否 | ✓ 1608ms | 否 | ✓ 1731ms | http |
| 162.240.154.26:3128 | ✓ 1851ms | ✓ 1769ms | 否 | 否 | ✓ 1974ms | http |
| 91.238.104.172:2024 | ✓ 1806ms | 否 | ✓ 1518ms | 否 | ✓ 1770ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1501ms | 否 | ✓ 1416ms | ✓ 975ms | http |
| 45.140.147.82:1082 | ✓ 1555ms | ✓ 1268ms | ✓ 1582ms | ✓ 1717ms | ✓ 1212ms | http |
| 45.140.147.82:1081 | ✓ 871ms | 否 | 否 | ✓ 1251ms | ✓ 1200ms | http |
| 104.238.30.39:59741 | ✓ 1791ms | 否 | ✓ 1835ms | 否 | ✓ 1998ms | http |
| 104.238.30.91:63900 | ✓ 1855ms | 否 | ✓ 1839ms | 否 | ✓ 1999ms | http |
| 103.236.64.247:8888 | ✓ 1371ms | ✓ 1999ms | 否 | ✓ 1725ms | 否 | http |
| 121.230.8.246:1080 | ✓ 1213ms | ✓ 1901ms | ✓ 1594ms | ✓ 1969ms | ✓ 1251ms | http |
| 45.140.147.155:1081 | ✓ 605ms | ✓ 1212ms | ✓ 883ms | ✓ 1663ms | 否 | http |
| 103.162.63.100:3125 | 否 | 否 | ✓ 1752ms | ✓ 1821ms | ✓ 1832ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1235ms | ✓ 1749ms | ✓ 1548ms | ✓ 1605ms | http |
| 121.232.73.224:1080 | 否 | ✓ 1448ms | ✓ 1475ms | 否 | ✓ 1415ms | http |
| 121.230.8.245:1080 | ✓ 1507ms | 否 | ✓ 1491ms | 否 | ✓ 1213ms | http |
| 156.225.70.152:39151 | 否 | 否 | ✓ 890ms | ✓ 1336ms | ✓ 890ms | http |
| 45.88.0.111:3128 | ✓ 967ms | ✓ 1667ms | ✓ 1788ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 702ms | 否 | ✓ 776ms | ✓ 1570ms | ✓ 1163ms | http |
| 104.238.30.38:59741 | ✓ 1801ms | 否 | ✓ 1843ms | 否 | ✓ 1999ms | http |
| 101.255.107.34:8080 | 否 | 否 | ✓ 1772ms | ✓ 1818ms | ✓ 1484ms | http |
| 45.88.0.116:3128 | ✓ 1940ms | ✓ 1255ms | 否 | ✓ 1303ms | ✓ 1385ms | http |
| 144.31.69.170:1080 | ✓ 961ms | ✓ 1839ms | ✓ 1509ms | 否 | ✓ 1704ms | http |
| 85.89.127.80:1081 | ✓ 802ms | 否 | ✓ 1245ms | ✓ 1759ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1332ms | ✓ 1695ms | ✓ 1570ms | 否 | 否 | http |
| 47.110.42.192:9003 | ✓ 1663ms | ✓ 1497ms | ✓ 1572ms | ✓ 1778ms | ✓ 1572ms | http |
| 45.88.0.98:3128 | 否 | 否 | ✓ 1931ms | ✓ 1922ms | ✓ 1940ms | http |
| 45.88.0.115:3128 | 否 | 否 | ✓ 1837ms | ✓ 1994ms | ✓ 1950ms | http |
| 103.39.51.190:8080 | ✓ 1816ms | 否 | ✓ 1380ms | ✓ 1703ms | ✓ 1694ms | http |
| 43.161.214.161:1081 | ✓ 990ms | 否 | ✓ 1030ms | ✓ 1451ms | ✓ 1201ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1603ms | ✓ 859ms | 否 | ✓ 1599ms | http |
| 192.71.213.85:9091 | ✓ 716ms | 否 | ✓ 1071ms | ✓ 1653ms | 否 | http |
| 192.71.213.85:9812 | ✓ 470ms | 否 | ✓ 480ms | ✓ 1641ms | 否 | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1175ms | ✓ 1654ms | ✓ 1236ms | http |
| 121.230.9.248:1080 | ✓ 1323ms | ✓ 1683ms | ✓ 1263ms | ✓ 1649ms | 否 | http |
| 15.204.151.141:3128 | ✓ 1254ms | 否 | ✓ 1804ms | ✓ 1783ms | 否 | http |
| 83.219.250.8:62920 | ✓ 492ms | ✓ 1565ms | ✓ 1970ms | 否 | 否 | http |
| 165.225.34.34:12152 | ✓ 652ms | ✓ 970ms | ✓ 848ms | ✓ 1978ms | ✓ 925ms | http |
| 165.225.222.27:12267 | ✓ 648ms | ✓ 1070ms | ✓ 830ms | 否 | ✓ 926ms | http |
| 165.225.216.23:11727 | ✓ 649ms | ✓ 1934ms | ✓ 104ms | 否 | ✓ 949ms | http |
| 165.225.222.26:12452 | ✓ 649ms | 否 | ✓ 305ms | ✓ 1861ms | ✓ 791ms | http |
| 165.225.38.43:12080 | ✓ 632ms | ✓ 1949ms | ✓ 580ms | ✓ 1217ms | ✓ 1939ms | http |
| 36.147.78.166:443 | ✓ 1854ms | 否 | ✓ 1970ms | ✓ 1786ms | 否 | http |
| 85.208.108.43:2094 | ✓ 1565ms | 否 | ✓ 1126ms | ✓ 1176ms | ✓ 913ms | http |
| 188.166.208.168:9876 | ✓ 822ms | 否 | ✓ 1005ms | ✓ 1178ms | ✓ 954ms | http |
| 121.230.8.111:1080 | ✓ 1348ms | ✓ 1571ms | ✓ 1275ms | ✓ 1760ms | ✓ 1385ms | http |

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
