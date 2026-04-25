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

最后更新：2026-04-25 10:48:40 UTC（2026-04-25 18:48:40 UTC+8）

**代理总数：54**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1926ms | ✓ 1715ms | ✓ 1080ms | ✓ 1557ms | ✓ 994ms | http |
| 47.85.51.197:1080 | ✓ 221ms | 否 | ✓ 855ms | ✓ 1766ms | ✓ 854ms | http |
| 47.84.76.30:1080 | ✓ 777ms | ✓ 1714ms | ✓ 791ms | ✓ 1071ms | ✓ 848ms | http |
| 206.206.126.177:2412 | ✓ 736ms | 否 | ✓ 1318ms | ✓ 1021ms | ✓ 803ms | http |
| 45.140.147.82:1081 | ✓ 1299ms | ✓ 1919ms | ✓ 1594ms | ✓ 1798ms | ✓ 1555ms | http |
| 106.10.55.212:1121 | ✓ 1259ms | ✓ 1333ms | ✓ 1464ms | ✓ 1210ms | 否 | http |
| 130.61.174.200:1080 | ✓ 709ms | 否 | ✓ 1709ms | ✓ 1664ms | 否 | http |
| 210.77.22.138:7890 | ✓ 1290ms | ✓ 1689ms | ✓ 1399ms | 否 | 否 | http |
| 46.101.95.183:8888 | ✓ 1361ms | 否 | ✓ 1285ms | 否 | ✓ 1564ms | http |
| 2.27.54.161:1080 | ✓ 1321ms | ✓ 1885ms | ✓ 1403ms | 否 | ✓ 1681ms | http |
| 59.46.216.131:30001 | ✓ 1853ms | ✓ 1625ms | ✓ 1508ms | 否 | 否 | http |
| 218.108.131.186:17890 | 否 | ✓ 1201ms | ✓ 1003ms | ✓ 1297ms | ✓ 1096ms | http |
| 166.88.61.54:8000 | ✓ 1083ms | ✓ 1309ms | ✓ 1382ms | ✓ 1307ms | ✓ 1112ms | http |
| 161.35.181.96:999 | ✓ 1499ms | 否 | 否 | ✓ 1173ms | ✓ 1848ms | http |
| 94.131.106.231:1081 | 否 | ✓ 1855ms | ✓ 1496ms | ✓ 1957ms | ✓ 1613ms | http |
| 45.140.147.82:1082 | ✓ 1105ms | 否 | ✓ 1676ms | ✓ 1583ms | 否 | http |
| 86.104.74.110:1081 | 否 | ✓ 1815ms | ✓ 794ms | 否 | ✓ 1317ms | http |
| 35.225.22.61:80 | ✓ 849ms | 否 | ✓ 407ms | ✓ 1211ms | 否 | http |
| 66.228.47.125:110 | ✓ 712ms | ✓ 1584ms | ✓ 1042ms | 否 | ✓ 1677ms | http |
| 45.88.0.115:3128 | ✓ 585ms | 否 | ✓ 1751ms | ✓ 1880ms | 否 | http |
| 201.144.20.238:3128 | ✓ 1330ms | ✓ 1322ms | ✓ 1156ms | ✓ 1171ms | ✓ 1022ms | http |
| 45.140.147.155:1081 | ✓ 929ms | 否 | ✓ 1848ms | 否 | ✓ 1625ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1548ms | ✓ 1482ms | ✓ 1203ms | http |
| 177.93.132.244:3128 | ✓ 1413ms | 否 | ✓ 955ms | 否 | ✓ 1793ms | http |
| 213.220.62.63:3128 | ✓ 1495ms | 否 | ✓ 1507ms | 否 | ✓ 1723ms | http |
| 45.88.0.116:3128 | ✓ 598ms | ✓ 1251ms | ✓ 1871ms | ✓ 1809ms | ✓ 1681ms | http |
| 45.88.0.114:3128 | ✓ 578ms | 否 | ✓ 1138ms | ✓ 1900ms | ✓ 1596ms | http |
| 213.220.62.62:3128 | ✓ 583ms | 否 | ✓ 1130ms | ✓ 1841ms | ✓ 1660ms | http |
| 45.88.0.117:3128 | ✓ 579ms | ✓ 1598ms | ✓ 1536ms | ✓ 1825ms | ✓ 1788ms | http |
| 45.88.0.99:3128 | ✓ 576ms | ✓ 1421ms | ✓ 1718ms | ✓ 1829ms | ✓ 1683ms | http |
| 45.88.0.98:3128 | ✓ 589ms | ✓ 1415ms | ✓ 1712ms | ✓ 1835ms | ✓ 1673ms | http |
| 45.88.0.113:3128 | ✓ 594ms | ✓ 1417ms | ✓ 1702ms | ✓ 1832ms | ✓ 1685ms | http |
| 45.88.0.111:3128 | ✓ 575ms | ✓ 1663ms | ✓ 1476ms | ✓ 1828ms | ✓ 1685ms | http |
| 94.131.122.128:1081 | ✓ 803ms | ✓ 1843ms | ✓ 1583ms | 否 | 否 | http |
| 45.129.141.143:3128 | ✓ 1370ms | 否 | 否 | ✓ 1989ms | ✓ 1841ms | http |
| 47.74.226.8:5001 | ✓ 1834ms | ✓ 1520ms | 否 | 否 | ✓ 1298ms | http |
| 43.133.44.89:8888 | 否 | ✓ 1557ms | 否 | ✓ 1167ms | ✓ 1557ms | http |
| 120.79.99.232:8099 | 否 | 否 | ✓ 1574ms | ✓ 1655ms | ✓ 1310ms | http |
| 113.11.127.179:64300 | ✓ 1839ms | 否 | ✓ 1408ms | ✓ 1906ms | 否 | http |
| 45.76.207.177:40000 | 否 | 否 | ✓ 1387ms | ✓ 1341ms | ✓ 1199ms | http |
| 80.92.204.47:1081 | ✓ 1161ms | 否 | ✓ 1225ms | ✓ 1695ms | ✓ 1223ms | http |
| 8.219.195.129:1080 | ✓ 864ms | 否 | ✓ 959ms | ✓ 1147ms | ✓ 946ms | http |
| 92.113.149.172:1080 | ✓ 876ms | ✓ 1623ms | ✓ 1553ms | 否 | 否 | http |
| 121.230.9.160:1080 | 否 | ✓ 1326ms | ✓ 1048ms | ✓ 1397ms | ✓ 1806ms | http |
| 62.113.119.14:8080 | ✓ 1409ms | 否 | ✓ 705ms | ✓ 1900ms | ✓ 1780ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1432ms | ✓ 1233ms | 否 | ✓ 1359ms | http |
| 159.89.31.62:8080 | ✓ 598ms | 否 | ✓ 1993ms | ✓ 1827ms | ✓ 1830ms | http |
| 1.231.81.166:3128 | ✓ 1683ms | 否 | ✓ 1011ms | ✓ 1066ms | ✓ 857ms | http |
| 121.230.8.91:1080 | ✓ 1205ms | ✓ 1236ms | ✓ 1190ms | ✓ 1650ms | 否 | http |
| 185.191.236.162:8080 | ✓ 1778ms | ✓ 1690ms | ✓ 1858ms | 否 | ✓ 1750ms | http |
| 103.126.238.13:8081 | ✓ 1826ms | 否 | 否 | ✓ 1601ms | ✓ 1695ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1821ms | ✓ 1770ms | ✓ 1675ms | http |
| 94.131.122.129:1081 | ✓ 569ms | 否 | ✓ 1879ms | 否 | ✓ 1974ms | http |
| 103.39.51.207:8080 | ✓ 1432ms | 否 | ✓ 1707ms | ✓ 1513ms | 否 | http |

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
