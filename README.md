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

最后更新：2026-05-04 21:10:00 UTC（2026-05-05 05:10:00 UTC+8）

**代理总数：61**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 126ms | 否 | ✓ 742ms | ✓ 1372ms | ✓ 1060ms | http |
| 185.191.236.162:3128 | ✓ 1101ms | ✓ 1509ms | ✓ 1830ms | 否 | ✓ 1685ms | http |
| 113.160.132.26:8080 | ✓ 1604ms | ✓ 1545ms | ✓ 1279ms | ✓ 1445ms | 否 | http |
| 34.71.229.255:3128 | ✓ 361ms | ✓ 1531ms | ✓ 937ms | 否 | ✓ 1753ms | http |
| 94.131.118.129:1081 | ✓ 1133ms | ✓ 1491ms | ✓ 1079ms | 否 | ✓ 1227ms | http |
| 165.225.113.220:8800 | ✓ 1840ms | 否 | ✓ 940ms | ✓ 1319ms | ✓ 996ms | http |
| 47.77.216.82:1080 | 否 | ✓ 1053ms | ✓ 281ms | 否 | ✓ 1619ms | http |
| 86.104.74.110:1081 | ✓ 1461ms | ✓ 1590ms | ✓ 1481ms | ✓ 1816ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1736ms | 否 | ✓ 1505ms | ✓ 1546ms | ✓ 1404ms | http |
| 91.242.229.129:8092 | ✓ 998ms | 否 | ✓ 1360ms | ✓ 1853ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1399ms | ✓ 1707ms | ✓ 1405ms | 否 | ✓ 1170ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1366ms | ✓ 1686ms | ✓ 1564ms | http |
| 99.79.58.74:9951 | ✓ 1235ms | 否 | ✓ 1019ms | 否 | ✓ 1965ms | http |
| 152.70.91.193:40000 | ✓ 1841ms | 否 | 否 | ✓ 1972ms | ✓ 1333ms | http |
| 103.157.200.126:3128 | ✓ 1714ms | 否 | ✓ 1898ms | 否 | ✓ 1778ms | http |
| 1.231.81.166:3128 | ✓ 1800ms | ✓ 1297ms | 否 | ✓ 1366ms | ✓ 1282ms | http |
| 218.108.131.186:17890 | ✓ 1263ms | ✓ 1244ms | ✓ 1019ms | ✓ 1307ms | 否 | http |
| 206.206.126.177:2412 | ✓ 1082ms | 否 | ✓ 1367ms | ✓ 1218ms | ✓ 960ms | http |
| 94.131.118.129:1082 | ✓ 993ms | ✓ 1410ms | ✓ 588ms | ✓ 1441ms | ✓ 1048ms | http |
| 160.238.65.9:3128 | ✓ 1147ms | ✓ 1297ms | ✓ 1523ms | ✓ 1871ms | ✓ 1395ms | http |
| 160.238.65.7:3128 | ✓ 1147ms | ✓ 1248ms | ✓ 1572ms | ✓ 1873ms | ✓ 1393ms | http |
| 154.64.232.35:8080 | ✓ 1124ms | 否 | 否 | ✓ 1985ms | ✓ 1106ms | http |
| 120.92.108.86:7890 | ✓ 1811ms | 否 | ✓ 1582ms | 否 | ✓ 1488ms | http |
| 49.48.110.96:8080 | ✓ 1667ms | 否 | ✓ 1728ms | ✓ 1927ms | ✓ 1608ms | http |
| 150.107.140.238:3128 | ✓ 1863ms | 否 | ✓ 1105ms | 否 | ✓ 1409ms | http |
| 181.119.97.24:999 | ✓ 1376ms | ✓ 1758ms | ✓ 1920ms | 否 | 否 | http |
| 107.173.42.121:7890 | 否 | ✓ 1108ms | ✓ 670ms | ✓ 1244ms | 否 | http |
| 103.165.138.173:8181 | ✓ 1695ms | 否 | ✓ 1919ms | ✓ 1423ms | ✓ 1149ms | http |
| 103.129.200.2:8124 | ✓ 1560ms | 否 | ✓ 1350ms | 否 | ✓ 1692ms | http |
| 51.44.97.6:46704 | ✓ 1098ms | 否 | ✓ 1620ms | 否 | ✓ 1514ms | http |
| 18.170.25.193:54929 | ✓ 1109ms | 否 | ✓ 1708ms | ✓ 1927ms | ✓ 1762ms | http |
| 52.16.215.4:8000 | ✓ 1692ms | 否 | ✓ 1028ms | ✓ 1954ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1057ms | 否 | ✓ 1238ms | ✓ 1279ms | ✓ 1313ms | http |
| 45.59.122.132:80 | ✓ 1382ms | 否 | ✓ 1173ms | 否 | ✓ 1335ms | http |
| 101.32.243.189:80 | ✓ 1408ms | 否 | ✓ 1818ms | ✓ 1745ms | ✓ 1629ms | http |
| 101.6.65.112:10080 | ✓ 1231ms | ✓ 1476ms | ✓ 1247ms | ✓ 1461ms | ✓ 1170ms | http |
| 3.101.133.120:80 | ✓ 597ms | ✓ 1643ms | ✓ 1302ms | ✓ 1659ms | ✓ 1304ms | http |
| 94.131.118.39:1082 | ✓ 865ms | ✓ 1199ms | ✓ 818ms | ✓ 1460ms | ✓ 1039ms | http |
| 99.79.58.74:16415 | ✓ 976ms | 否 | ✓ 1300ms | 否 | ✓ 1946ms | http |
| 52.59.218.12:1973 | ✓ 700ms | 否 | ✓ 1515ms | ✓ 1938ms | ✓ 1760ms | http |
| 8.211.166.184:8081 | ✓ 1563ms | ✓ 1235ms | ✓ 754ms | ✓ 1077ms | ✓ 824ms | http |
| 35.180.75.159:17928 | ✓ 1181ms | 否 | ✓ 1675ms | 否 | ✓ 1318ms | http |
| 13.51.196.44:16963 | ✓ 1195ms | 否 | ✓ 1757ms | 否 | ✓ 1579ms | http |
| 52.56.167.111:56710 | ✓ 1768ms | 否 | ✓ 1685ms | 否 | ✓ 1339ms | http |
| 94.131.118.39:1081 | ✓ 1095ms | ✓ 1304ms | ✓ 713ms | ✓ 1506ms | ✓ 1076ms | http |
| 86.104.72.220:1081 | ✓ 288ms | ✓ 1026ms | ✓ 1104ms | ✓ 1715ms | ✓ 1470ms | http |
| 194.59.247.34:10808 | ✓ 669ms | ✓ 1325ms | ✓ 1550ms | ✓ 1951ms | ✓ 1195ms | http |
| 86.104.72.220:1082 | ✓ 1153ms | ✓ 1160ms | 否 | 否 | ✓ 1957ms | http |
| 103.209.36.58:8080 | ✓ 1921ms | 否 | ✓ 1959ms | ✓ 1497ms | ✓ 1458ms | http |
| 80.92.204.47:1082 | ✓ 474ms | ✓ 1320ms | 否 | ✓ 1953ms | 否 | http |
| 45.140.147.155:1081 | ✓ 560ms | ✓ 1533ms | 否 | 否 | ✓ 1100ms | http |
| 80.92.204.47:1081 | ✓ 484ms | ✓ 1332ms | 否 | ✓ 1924ms | 否 | http |
| 103.240.6.22:16498 | ✓ 1641ms | 否 | ✓ 1341ms | ✓ 1935ms | ✓ 1678ms | http |
| 121.230.8.55:1080 | ✓ 1394ms | 否 | 否 | ✓ 1522ms | ✓ 1405ms | http |
| 91.238.105.64:2024 | ✓ 1236ms | ✓ 1552ms | ✓ 1549ms | 否 | ✓ 1713ms | http |
| 61.52.131.172:8443 | ✓ 1120ms | ✓ 1527ms | ✓ 1176ms | ✓ 1444ms | ✓ 1135ms | http |
| 45.129.141.143:3128 | ✓ 914ms | ✓ 1695ms | ✓ 1682ms | 否 | 否 | http |
| 178.156.224.42:3128 | ✓ 1952ms | 否 | ✓ 1820ms | 否 | ✓ 1320ms | http |
| 38.180.2.107:3128 | ✓ 1338ms | ✓ 1867ms | ✓ 1797ms | 否 | 否 | http |
| 3.99.158.157:44636 | ✓ 1127ms | 否 | ✓ 1980ms | 否 | ✓ 1903ms | http |
| 223.16.170.103:80 | ✓ 1376ms | 否 | ✓ 1074ms | ✓ 1329ms | ✓ 1353ms | http |

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
