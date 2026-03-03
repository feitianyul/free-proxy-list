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

最后更新：2026-03-03 05:46:53 UTC（2026-03-03 13:46:53 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 256ms | 否 | 否 | ✓ 1353ms | ✓ 1978ms | http |
| 120.92.212.16:8890 | ✓ 1819ms | 否 | ✓ 1138ms | 否 | ✓ 1872ms | http |
| 3.213.157.4:3128 | ✓ 233ms | 否 | ✓ 91ms | ✓ 951ms | ✓ 715ms | http |
| 166.0.192.117:8888 | ✓ 228ms | 否 | 否 | ✓ 1421ms | ✓ 911ms | http |
| 125.128.12.14:3128 | ✓ 712ms | 否 | 否 | ✓ 1116ms | ✓ 885ms | http |
| 35.225.22.61:80 | ✓ 580ms | ✓ 1215ms | ✓ 878ms | ✓ 851ms | ✓ 690ms | http |
| 5.75.196.26:40000 | 否 | ✓ 1825ms | 否 | ✓ 1964ms | ✓ 1998ms | http |
| 207.180.228.55:80 | ✓ 458ms | ✓ 1835ms | ✓ 1459ms | 否 | ✓ 1454ms | http |
| 81.70.169.194:80 | ✓ 1063ms | ✓ 1562ms | ✓ 1171ms | ✓ 1368ms | ✓ 1119ms | http |
| 101.43.255.96:80 | ✓ 1469ms | ✓ 1693ms | ✓ 1855ms | ✓ 1394ms | ✓ 1195ms | http |
| 91.238.104.171:2023 | ✓ 1077ms | ✓ 1966ms | 否 | 否 | ✓ 1920ms | http |
| 70.22.175.232:3128 | ✓ 196ms | ✓ 1078ms | ✓ 160ms | ✓ 1047ms | ✓ 793ms | http |
| 183.249.5.111:22222 | ✓ 1463ms | 否 | 否 | ✓ 1797ms | ✓ 852ms | http |
| 222.184.48.236:22222 | ✓ 1277ms | 否 | ✓ 1019ms | ✓ 1553ms | 否 | http |
| 91.99.99.83:9000 | ✓ 510ms | 否 | 否 | ✓ 1880ms | ✓ 1613ms | http |
| 222.184.48.252:22222 | ✓ 1443ms | ✓ 1278ms | ✓ 1498ms | 否 | 否 | http |
| 160.250.5.22:1 | ✓ 1867ms | 否 | ✓ 1406ms | ✓ 1315ms | ✓ 1053ms | http |
| 160.250.4.245:1 | ✓ 1874ms | 否 | ✓ 1393ms | ✓ 1370ms | ✓ 1091ms | http |
| 125.128.12.144:3128 | 否 | 否 | ✓ 1710ms | ✓ 1133ms | ✓ 879ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 1967ms | ✓ 1539ms | ✓ 1253ms | http |
| 14.56.107.244:3128 | ✓ 745ms | 否 | ✓ 799ms | 否 | ✓ 1928ms | http |
| 103.84.95.54:7890 | ✓ 766ms | 否 | 否 | ✓ 1028ms | ✓ 839ms | http |
| 74.208.234.198:443 | ✓ 1050ms | ✓ 1773ms | ✓ 1426ms | 否 | 否 | http |
| 61.72.110.54:3128 | ✓ 1395ms | ✓ 1924ms | ✓ 1418ms | 否 | ✓ 1773ms | http |
| 95.85.252.153:21064 | ✓ 494ms | ✓ 1886ms | ✓ 1307ms | 否 | 否 | http |
| 212.175.29.184:8080 | ✓ 1371ms | 否 | ✓ 1630ms | 否 | ✓ 1677ms | http |
| 120.92.212.16:7890 | ✓ 1769ms | 否 | 否 | ✓ 1364ms | ✓ 1046ms | http |
| 47.94.228.56:8090 | ✓ 1291ms | ✓ 1357ms | ✓ 1138ms | ✓ 1461ms | ✓ 1175ms | http |
| 120.232.242.119:22222 | 否 | ✓ 1386ms | ✓ 1113ms | ✓ 1274ms | ✓ 1002ms | http |
| 120.198.141.79:22222 | ✓ 1037ms | ✓ 1403ms | 否 | ✓ 1276ms | ✓ 1022ms | http |
| 160.238.65.3:3128 | ✓ 1406ms | 否 | 否 | ✓ 1371ms | ✓ 1075ms | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 1571ms | ✓ 1263ms | ✓ 1330ms | http |
| 103.175.228.142:8090 | 否 | 否 | ✓ 1922ms | ✓ 1618ms | ✓ 1585ms | http |
| 183.249.5.214:22222 | 否 | ✓ 1283ms | 否 | ✓ 1635ms | ✓ 1958ms | http |
| 117.159.239.44:22222 | ✓ 1680ms | 否 | 否 | ✓ 1787ms | ✓ 1969ms | http |
| 120.198.141.80:22222 | ✓ 1098ms | ✓ 1599ms | ✓ 1347ms | ✓ 1431ms | ✓ 1049ms | http |
| 121.128.121.54:3128 | ✓ 1060ms | 否 | ✓ 1134ms | 否 | ✓ 887ms | http |
| 121.230.9.160:1080 | ✓ 1363ms | 否 | ✓ 1805ms | 否 | ✓ 1261ms | http |
| 202.129.206.239:3128 | ✓ 1879ms | 否 | ✓ 1392ms | ✓ 1716ms | ✓ 1656ms | http |
| 91.233.223.147:3128 | 否 | 否 | ✓ 1619ms | ✓ 1959ms | ✓ 1842ms | http |
| 192.71.213.85:9812 | ✓ 1434ms | 否 | ✓ 1754ms | ✓ 1859ms | 否 | http |
| 45.136.198.40:3128 | ✓ 984ms | 否 | ✓ 1579ms | 否 | ✓ 1500ms | http |
| 61.72.110.94:3128 | ✓ 1101ms | 否 | ✓ 818ms | 否 | ✓ 937ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1917ms | 否 | ✓ 1650ms | ✓ 1737ms | http |
| 45.140.147.155:1081 | ✓ 469ms | 否 | ✓ 1317ms | ✓ 1706ms | ✓ 1125ms | http |
| 160.238.65.9:3128 | ✓ 864ms | 否 | ✓ 1111ms | 否 | ✓ 1336ms | http |
| 91.238.104.172:2024 | ✓ 897ms | 否 | 否 | ✓ 1514ms | ✓ 1245ms | http |
| 113.59.32.162:22222 | ✓ 1530ms | ✓ 1770ms | 否 | ✓ 1946ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1870ms | 否 | 否 | ✓ 1989ms | ✓ 1825ms | http |
| 117.159.239.51:22222 | ✓ 909ms | ✓ 1205ms | 否 | ✓ 1289ms | 否 | http |
| 195.123.209.48:3128 | ✓ 1247ms | ✓ 1904ms | ✓ 1223ms | 否 | ✓ 1747ms | http |
| 119.134.178.238:7890 | ✓ 1086ms | ✓ 1423ms | 否 | ✓ 1491ms | ✓ 1079ms | http |
| 49.151.179.53:8082 | ✓ 1792ms | 否 | ✓ 1446ms | ✓ 1861ms | ✓ 1504ms | http |
| 103.166.33.51:3125 | ✓ 1897ms | 否 | 否 | ✓ 1897ms | ✓ 1762ms | http |
| 2.56.178.131:443 | ✓ 965ms | 否 | ✓ 1541ms | 否 | ✓ 1545ms | http |
| 62.113.119.14:8080 | ✓ 1172ms | 否 | ✓ 761ms | ✓ 1632ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1803ms | 否 | ✓ 1002ms | 否 | ✓ 1948ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1630ms | ✓ 878ms | ✓ 1130ms | ✓ 848ms | http |
| 138.124.53.25:7443 | ✓ 1298ms | 否 | ✓ 1245ms | ✓ 1907ms | 否 | http |
| 74.48.78.224:2080 | ✓ 1381ms | 否 | ✓ 1581ms | ✓ 1411ms | 否 | http |
| 103.131.19.42:8181 | ✓ 1630ms | 否 | ✓ 1702ms | ✓ 1553ms | ✓ 1517ms | http |
| 61.52.131.172:8443 | 否 | 否 | ✓ 1002ms | ✓ 1259ms | ✓ 1036ms | http |
| 120.238.159.228:22222 | ✓ 997ms | ✓ 1401ms | ✓ 965ms | ✓ 1272ms | ✓ 994ms | http |
| 46.249.103.192:443 | ✓ 1589ms | 否 | ✓ 1377ms | 否 | ✓ 1872ms | http |

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
