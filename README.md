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

最后更新：2026-03-11 12:28:35 UTC（2026-03-11 20:28:35 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.175:8443 | 否 | ✓ 1079ms | ✓ 1477ms | ✓ 984ms | ✓ 732ms | http |
| 45.136.131.47:8443 | 否 | 否 | ✓ 550ms | ✓ 1105ms | ✓ 736ms | http |
| 45.136.131.63:8443 | 否 | 否 | ✓ 998ms | ✓ 1014ms | ✓ 710ms | http |
| 91.107.141.42:8081 | ✓ 1171ms | ✓ 1818ms | ✓ 1469ms | 否 | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1883ms | ✓ 1383ms | ✓ 1506ms | ✓ 1070ms | http |
| 62.113.119.14:8080 | ✓ 559ms | 否 | ✓ 1194ms | 否 | ✓ 1060ms | http |
| 168.235.110.63:3128 | ✓ 676ms | 否 | ✓ 1767ms | ✓ 1894ms | 否 | http |
| 45.140.147.82:1081 | ✓ 394ms | 否 | ✓ 1496ms | 否 | ✓ 1060ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1760ms | ✓ 1909ms | 否 | ✓ 1748ms | http |
| 217.76.245.80:999 | ✓ 915ms | 否 | 否 | ✓ 1439ms | ✓ 1022ms | http |
| 194.213.18.200:443 | ✓ 710ms | 否 | 否 | ✓ 1950ms | ✓ 1652ms | http |
| 162.248.165.72:1080 | ✓ 1044ms | 否 | ✓ 1857ms | 否 | ✓ 1406ms | http |
| 103.84.95.54:7890 | ✓ 891ms | ✓ 1898ms | ✓ 1147ms | ✓ 1678ms | ✓ 1595ms | http |
| 190.9.109.198:999 | ✓ 634ms | 否 | ✓ 1308ms | ✓ 1200ms | ✓ 1014ms | http |
| 150.107.140.238:3128 | ✓ 1945ms | 否 | ✓ 1885ms | 否 | ✓ 1576ms | http |
| 165.227.5.10:8888 | ✓ 394ms | 否 | 否 | ✓ 1017ms | ✓ 1788ms | http |
| 120.92.212.16:7890 | ✓ 1205ms | ✓ 1678ms | 否 | 否 | ✓ 1114ms | http |
| 36.155.100.217:8080 | 否 | ✓ 1290ms | ✓ 1294ms | ✓ 1567ms | ✓ 1333ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1168ms | 否 | ✓ 1124ms | ✓ 726ms | http |
| 46.183.25.8:443 | ✓ 921ms | 否 | ✓ 443ms | ✓ 1161ms | ✓ 939ms | http |
| 45.136.130.223:8443 | ✓ 927ms | ✓ 1914ms | ✓ 272ms | ✓ 1074ms | ✓ 740ms | http |
| 202.155.12.161:443 | ✓ 1731ms | 否 | ✓ 1378ms | ✓ 1311ms | 否 | http |
| 95.3.9.78:3128 | ✓ 675ms | ✓ 1813ms | ✓ 620ms | ✓ 1588ms | ✓ 1810ms | http |
| 45.136.130.188:8443 | ✓ 433ms | ✓ 923ms | ✓ 856ms | ✓ 934ms | ✓ 1511ms | http |
| 205.209.118.30:3138 | ✓ 974ms | 否 | ✓ 1635ms | 否 | ✓ 1821ms | http |
| 45.136.130.191:8443 | 否 | ✓ 1948ms | ✓ 282ms | ✓ 931ms | ✓ 695ms | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1591ms | ✓ 1406ms | ✓ 1241ms | http |
| 160.238.65.8:3128 | ✓ 1551ms | ✓ 1319ms | ✓ 400ms | 否 | 否 | http |
| 160.238.65.4:3128 | ✓ 1541ms | ✓ 1310ms | ✓ 388ms | 否 | 否 | http |
| 160.238.65.7:3128 | ✓ 1554ms | ✓ 1749ms | ✓ 384ms | 否 | 否 | http |
| 160.238.65.3:3128 | ✓ 1767ms | ✓ 1932ms | ✓ 406ms | 否 | 否 | http |
| 160.238.65.2:3128 | ✓ 1536ms | ✓ 1444ms | ✓ 429ms | 否 | 否 | http |
| 160.238.65.9:3128 | ✓ 1558ms | ✓ 1367ms | ✓ 387ms | 否 | 否 | http |
| 94.176.3.43:7443 | ✓ 1759ms | 否 | ✓ 1117ms | ✓ 1587ms | 否 | http |
| 37.139.33.145:1080 | 否 | ✓ 1829ms | ✓ 1891ms | 否 | ✓ 1870ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1472ms | ✓ 1269ms | 否 | ✓ 1195ms | http |
| 180.121.187.93:1080 | 否 | ✓ 1389ms | ✓ 1632ms | ✓ 1396ms | 否 | http |
| 81.70.169.194:80 | ✓ 1201ms | 否 | ✓ 1207ms | ✓ 1919ms | ✓ 1632ms | http |
| 152.42.213.210:8080 | ✓ 945ms | 否 | ✓ 1593ms | ✓ 1344ms | ✓ 990ms | http |
| 160.238.65.5:3128 | ✓ 824ms | ✓ 1795ms | 否 | 否 | ✓ 1274ms | http |
| 190.6.54.12:6969 | ✓ 1435ms | 否 | ✓ 1431ms | 否 | ✓ 1879ms | http |
| 95.3.9.78:8080 | ✓ 660ms | ✓ 1751ms | ✓ 1366ms | 否 | ✓ 1313ms | http |
| 158.69.185.37:3129 | ✓ 304ms | ✓ 1439ms | ✓ 1878ms | ✓ 1018ms | ✓ 840ms | http |
| 101.32.244.83:8080 | 否 | 否 | ✓ 1138ms | ✓ 1720ms | ✓ 1458ms | http |
| 121.43.196.210:8222 | ✓ 1213ms | ✓ 1306ms | ✓ 1122ms | ✓ 1413ms | ✓ 1076ms | http |
| 121.43.196.213:8222 | 否 | ✓ 1265ms | ✓ 992ms | ✓ 1383ms | ✓ 1034ms | http |
| 114.55.226.123:10086 | ✓ 1237ms | 否 | ✓ 1352ms | ✓ 1445ms | ✓ 1221ms | http |
| 24.199.124.151:3128 | 否 | ✓ 1219ms | ✓ 1329ms | ✓ 961ms | ✓ 769ms | http |
| 88.80.150.82:8080 | ✓ 970ms | 否 | 否 | ✓ 1885ms | ✓ 1621ms | https |
| 35.225.22.61:80 | ✓ 1820ms | 否 | 否 | ✓ 1069ms | ✓ 1032ms | http |
| 101.43.255.96:80 | ✓ 1584ms | ✓ 1504ms | ✓ 1304ms | 否 | 否 | http |
| 123.57.0.163:8888 | ✓ 1714ms | ✓ 1751ms | 否 | ✓ 1954ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1051ms | 否 | ✓ 1332ms | ✓ 1250ms | ✓ 992ms | http |
| 45.140.147.155:1081 | ✓ 1463ms | ✓ 1308ms | ✓ 883ms | ✓ 1402ms | 否 | http |
| 107.172.125.217:3128 | ✓ 482ms | 否 | 否 | ✓ 1008ms | ✓ 854ms | http |
| 185.191.236.162:3128 | ✓ 535ms | 否 | ✓ 1001ms | ✓ 1785ms | ✓ 1003ms | http |
| 47.101.149.27:9010 | 否 | 否 | ✓ 1575ms | ✓ 1507ms | ✓ 1554ms | http |
| 85.208.108.43:10808 | ✓ 492ms | 否 | ✓ 625ms | 否 | ✓ 932ms | http |
| 152.42.213.210:443 | 否 | 否 | ✓ 1797ms | ✓ 1421ms | ✓ 1143ms | http |
| 162.240.154.26:3128 | ✓ 968ms | 否 | ✓ 1468ms | ✓ 1411ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1784ms | ✓ 1693ms | ✓ 662ms | ✓ 1954ms | ✓ 1556ms | http |
| 59.46.216.131:30001 | ✓ 1649ms | ✓ 1683ms | 否 | 否 | ✓ 1363ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1269ms | ✓ 1379ms | ✓ 1522ms | http |
| 39.104.201.40:7890 | ✓ 1070ms | ✓ 1407ms | ✓ 1108ms | ✓ 1449ms | ✓ 1151ms | http |
| 45.186.6.104:3128 | ✓ 1467ms | ✓ 1858ms | ✓ 1747ms | 否 | 否 | http |
| 14.225.212.37:7890 | ✓ 1948ms | 否 | ✓ 1055ms | ✓ 1352ms | 否 | http |
| 152.69.229.220:3128 | ✓ 1567ms | 否 | ✓ 1913ms | ✓ 1891ms | 否 | http |
| 61.52.131.172:8443 | 否 | ✓ 1317ms | ✓ 1133ms | ✓ 1369ms | ✓ 1112ms | http |
| 178.236.245.59:3128 | ✓ 1321ms | 否 | ✓ 911ms | ✓ 1778ms | ✓ 1288ms | http |
| 152.70.98.46:8888 | ✓ 1696ms | 否 | ✓ 1474ms | ✓ 1232ms | ✓ 1302ms | http |
| 14.225.222.213:7890 | 否 | ✓ 1879ms | 否 | ✓ 1393ms | ✓ 1888ms | http |
| 14.225.222.164:7890 | 否 | 否 | ✓ 1626ms | ✓ 1354ms | ✓ 1691ms | http |
| 66.228.47.125:110 | 否 | 否 | ✓ 1631ms | ✓ 1749ms | ✓ 1597ms | http |

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
