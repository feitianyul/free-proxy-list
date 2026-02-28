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

最后更新：2026-02-28 07:39:27 UTC（2026-02-28 15:39:27 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 399ms | ✓ 1210ms | ✓ 685ms | 否 | 否 | http |
| 3.213.157.4:3128 | ✓ 276ms | ✓ 1629ms | ✓ 782ms | ✓ 1457ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1908ms | 否 | ✓ 1257ms | 否 | ✓ 915ms | http |
| 168.235.110.63:3128 | ✓ 549ms | ✓ 1308ms | ✓ 1218ms | ✓ 1306ms | ✓ 929ms | http |
| 120.202.127.234:10808 | ✓ 971ms | ✓ 1176ms | ✓ 1027ms | ✓ 1226ms | ✓ 1012ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1388ms | 否 | ✓ 1356ms | ✓ 1086ms | http |
| 147.45.216.148:1080 | ✓ 1149ms | 否 | ✓ 1704ms | ✓ 1928ms | ✓ 1551ms | http |
| 132.145.93.138:1080 | ✓ 1456ms | 否 | ✓ 1680ms | ✓ 1908ms | ✓ 1628ms | http |
| 59.46.216.131:30001 | ✓ 1130ms | 否 | ✓ 1177ms | ✓ 1430ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1033ms | ✓ 1345ms | ✓ 1114ms | ✓ 1648ms | 否 | http |
| 81.70.169.194:80 | ✓ 1172ms | ✓ 1536ms | 否 | 否 | ✓ 1095ms | http |
| 36.147.78.166:443 | ✓ 1791ms | ✓ 1812ms | ✓ 1840ms | 否 | 否 | http |
| 136.226.254.24:10517 | 否 | 否 | ✓ 1362ms | ✓ 1889ms | ✓ 1723ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1502ms | ✓ 1733ms | ✓ 1905ms | http |
| 35.225.22.61:80 | ✓ 582ms | ✓ 1655ms | ✓ 401ms | ✓ 920ms | ✓ 812ms | http |
| 101.43.255.96:80 | ✓ 1177ms | ✓ 1417ms | ✓ 1322ms | 否 | ✓ 1126ms | http |
| 52.188.28.218:3128 | ✓ 265ms | ✓ 1520ms | ✓ 1156ms | ✓ 1923ms | 否 | http |
| 104.238.30.40:59741 | ✓ 1689ms | 否 | ✓ 1811ms | 否 | ✓ 1995ms | http |
| 62.113.119.14:8080 | ✓ 636ms | 否 | ✓ 769ms | ✓ 1609ms | ✓ 1136ms | http |
| 136.226.254.24:9400 | ✓ 933ms | 否 | ✓ 864ms | ✓ 1416ms | ✓ 1196ms | http |
| 103.84.95.54:7890 | ✓ 1292ms | 否 | 否 | ✓ 1033ms | ✓ 791ms | http |
| 36.147.78.166:80 | 否 | ✓ 1836ms | ✓ 1811ms | ✓ 1960ms | 否 | http |
| 120.26.147.60:7890 | 否 | ✓ 1187ms | ✓ 1823ms | ✓ 1229ms | ✓ 946ms | http |
| 104.238.30.45:59741 | ✓ 1748ms | 否 | ✓ 1771ms | 否 | ✓ 1971ms | http |
| 114.231.72.214:1080 | 否 | ✓ 1331ms | ✓ 1045ms | ✓ 1356ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1151ms | 否 | ✓ 1788ms | ✓ 1652ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1390ms | 否 | ✓ 1858ms | ✓ 1950ms | 否 | http |
| 35.234.17.221:8080 | ✓ 969ms | 否 | ✓ 1147ms | ✓ 1119ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1298ms | 否 | ✓ 1831ms | ✓ 1538ms | ✓ 1970ms | http |
| 72.56.50.17:59787 | ✓ 1546ms | 否 | ✓ 1932ms | 否 | ✓ 1843ms | http |
| 47.113.95.226:8118 | 否 | 否 | ✓ 1156ms | ✓ 1610ms | ✓ 1649ms | http |
| 121.237.181.137:8888 | ✓ 1272ms | ✓ 1226ms | ✓ 1058ms | 否 | ✓ 1061ms | http |
| 85.208.108.43:2094 | ✓ 338ms | 否 | ✓ 86ms | ✓ 1086ms | ✓ 830ms | http |
| 85.208.108.43:10808 | ✓ 338ms | 否 | ✓ 96ms | ✓ 1097ms | ✓ 819ms | http |
| 111.79.111.126:3128 | ✓ 955ms | ✓ 1423ms | ✓ 1086ms | 否 | ✓ 1904ms | http |
| 150.107.140.238:3128 | ✓ 1982ms | 否 | 否 | ✓ 1884ms | ✓ 1271ms | http |
| 165.225.120.17:12497 | ✓ 1903ms | 否 | ✓ 1305ms | ✓ 1609ms | 否 | http |
| 165.225.120.17:10919 | ✓ 1905ms | 否 | ✓ 1251ms | ✓ 1555ms | 否 | http |
| 165.225.120.17:10906 | ✓ 1906ms | 否 | ✓ 1298ms | ✓ 1681ms | 否 | http |
| 165.225.120.17:11745 | ✓ 1905ms | 否 | ✓ 1303ms | ✓ 1745ms | 否 | http |
| 165.225.120.17:10880 | ✓ 1907ms | 否 | ✓ 1301ms | ✓ 1564ms | ✓ 1254ms | http |
| 165.225.120.17:11995 | ✓ 1906ms | 否 | ✓ 1199ms | ✓ 1655ms | ✓ 1395ms | http |
| 165.225.120.17:11912 | ✓ 1905ms | 否 | ✓ 1197ms | 否 | ✓ 1114ms | http |
| 121.40.231.103:7890 | ✓ 995ms | 否 | ✓ 1023ms | 否 | ✓ 960ms | http |
| 70.34.199.124:10808 | ✓ 1662ms | 否 | ✓ 1489ms | ✓ 1590ms | ✓ 1499ms | http |
| 43.161.214.161:1081 | ✓ 1412ms | 否 | ✓ 1264ms | ✓ 1814ms | ✓ 1337ms | http |
| 61.72.110.54:3128 | ✓ 1670ms | ✓ 1675ms | 否 | 否 | ✓ 938ms | http |
| 45.136.198.40:3128 | ✓ 680ms | 否 | ✓ 1545ms | 否 | ✓ 1680ms | http |
| 183.222.38.157:1082 | ✓ 1793ms | 否 | 否 | ✓ 1430ms | ✓ 1122ms | http |
| 52.201.29.25:80 | ✓ 526ms | 否 | ✓ 98ms | ✓ 1663ms | ✓ 1359ms | http |
| 95.38.195.74:1080 | ✓ 1700ms | 否 | ✓ 1948ms | 否 | ✓ 1695ms | http |
| 81.177.48.54:2080 | ✓ 1752ms | 否 | 否 | ✓ 1852ms | ✓ 1791ms | http |
| 72.56.59.56:63127 | ✓ 1565ms | 否 | ✓ 1743ms | 否 | ✓ 1936ms | http |
| 61.72.110.24:3128 | ✓ 1152ms | 否 | 否 | ✓ 1182ms | ✓ 1693ms | http |
| 104.37.184.214:1080 | ✓ 672ms | ✓ 1650ms | 否 | ✓ 1099ms | ✓ 1106ms | http |
| 45.140.147.82:1081 | ✓ 953ms | ✓ 1388ms | ✓ 1229ms | ✓ 1377ms | ✓ 890ms | http |
| 121.204.158.249:3128 | ✓ 1327ms | ✓ 1359ms | ✓ 1099ms | 否 | 否 | http |
| 45.125.67.37:8443 | ✓ 1287ms | 否 | 否 | ✓ 1496ms | ✓ 1222ms | http |
| 72.56.59.62:63133 | ✓ 1743ms | 否 | ✓ 1614ms | 否 | ✓ 1842ms | http |
| 119.46.68.228:80 | ✓ 1793ms | 否 | 否 | ✓ 1463ms | ✓ 1199ms | http |
| 52.22.64.124:80 | ✓ 174ms | 否 | ✓ 1055ms | ✓ 1965ms | ✓ 1146ms | http |
| 45.140.147.82:1082 | ✓ 935ms | 否 | ✓ 1625ms | ✓ 1392ms | ✓ 1086ms | http |
| 165.225.192.16:10919 | ✓ 814ms | 否 | ✓ 1272ms | ✓ 1963ms | ✓ 1294ms | http |
| 103.82.23.118:5242 | ✓ 1861ms | 否 | ✓ 1843ms | 否 | ✓ 1740ms | http |
| 165.227.5.10:8888 | ✓ 796ms | ✓ 1911ms | ✓ 1293ms | 否 | ✓ 897ms | http |
| 115.231.181.40:8128 | ✓ 1744ms | ✓ 1311ms | 否 | ✓ 1588ms | 否 | http |
| 104.238.30.86:63900 | ✓ 1803ms | 否 | ✓ 1839ms | 否 | ✓ 1998ms | http |
| 104.238.30.63:63744 | ✓ 1784ms | 否 | ✓ 1903ms | 否 | ✓ 1999ms | http |
| 59.8.203.55:80 | ✓ 1239ms | ✓ 1630ms | ✓ 1643ms | ✓ 1334ms | ✓ 1059ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1376ms | ✓ 1500ms | ✓ 1467ms | http |
| 104.238.30.58:63744 | ✓ 1785ms | 否 | ✓ 1807ms | 否 | ✓ 1999ms | http |
| 222.28.182.229:7890 | ✓ 1093ms | 否 | ✓ 1254ms | ✓ 1698ms | ✓ 1260ms | http |

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
