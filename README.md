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

最后更新：2026-03-26 14:25:34 UTC（2026-03-26 22:25:34 UTC+8）

**代理总数：56**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 609ms | 否 | ✓ 815ms | ✓ 942ms | ✓ 935ms | http |
| 167.103.115.102:8800 | ✓ 1003ms | 否 | ✓ 1234ms | ✓ 1405ms | ✓ 1515ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1325ms | ✓ 1176ms | ✓ 1707ms | ✓ 1062ms | http |
| 167.103.34.108:8800 | ✓ 1756ms | 否 | ✓ 1505ms | ✓ 1651ms | ✓ 1551ms | http |
| 165.232.146.249:3128 | 否 | ✓ 1098ms | ✓ 1371ms | ✓ 1195ms | ✓ 1021ms | http |
| 35.225.22.61:80 | ✓ 1317ms | ✓ 1313ms | ✓ 1134ms | ✓ 1133ms | ✓ 1094ms | http |
| 101.47.73.135:3128 | ✓ 1083ms | 否 | ✓ 965ms | ✓ 1290ms | ✓ 995ms | http |
| 167.103.144.127:8800 | ✓ 1198ms | 否 | ✓ 956ms | ✓ 1855ms | ✓ 1613ms | http |
| 125.26.4.219:8080 | 否 | 否 | ✓ 1239ms | ✓ 1501ms | ✓ 1311ms | http |
| 167.103.31.122:8800 | ✓ 1462ms | 否 | ✓ 1296ms | ✓ 1981ms | ✓ 1476ms | http |
| 103.84.95.54:7890 | ✓ 1050ms | 否 | ✓ 860ms | 否 | ✓ 796ms | http |
| 186.148.180.46:999 | 否 | 否 | ✓ 1597ms | ✓ 1742ms | ✓ 1385ms | http |
| 5.104.87.17:8051 | ✓ 1787ms | 否 | 否 | ✓ 1641ms | ✓ 1477ms | http |
| 160.250.4.13:1 | 否 | 否 | ✓ 1284ms | ✓ 1544ms | ✓ 1191ms | http |
| 101.43.127.100:8877 | ✓ 937ms | 否 | ✓ 1091ms | ✓ 1340ms | ✓ 897ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1815ms | 否 | ✓ 1185ms | ✓ 1175ms | http |
| 147.161.239.240:8800 | ✓ 1129ms | ✓ 1823ms | ✓ 1527ms | ✓ 1504ms | ✓ 1296ms | http |
| 115.231.181.40:8128 | ✓ 1190ms | ✓ 1621ms | ✓ 872ms | 否 | 否 | http |
| 115.231.37.166:11518 | 否 | 否 | ✓ 1138ms | ✓ 1941ms | ✓ 1197ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1849ms | ✓ 1003ms | ✓ 1850ms | 否 | http |
| 89.208.106.138:10808 | ✓ 999ms | ✓ 1791ms | 否 | 否 | ✓ 1689ms | http |
| 20.210.39.153:8561 | ✓ 559ms | ✓ 794ms | ✓ 539ms | ✓ 939ms | ✓ 663ms | http |
| 20.78.118.91:8561 | ✓ 561ms | ✓ 1030ms | ✓ 437ms | ✓ 866ms | ✓ 668ms | http |
| 20.27.11.248:8561 | ✓ 795ms | ✓ 1371ms | ✓ 924ms | ✓ 1079ms | ✓ 707ms | http |
| 20.27.15.49:8561 | ✓ 1242ms | ✓ 1801ms | ✓ 1769ms | 否 | ✓ 1929ms | http |
| 20.27.13.35:8561 | ✓ 1325ms | ✓ 843ms | ✓ 884ms | ✓ 1183ms | ✓ 1097ms | http |
| 62.113.119.14:8080 | ✓ 1764ms | ✓ 1875ms | ✓ 1251ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1239ms | 否 | ✓ 1415ms | ✓ 1904ms | 否 | http |
| 133.18.110.87:1081 | ✓ 651ms | ✓ 1190ms | ✓ 686ms | ✓ 1017ms | 否 | http |
| 45.136.131.38:8444 | ✓ 1356ms | 否 | ✓ 1807ms | ✓ 1224ms | 否 | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1299ms | ✓ 1059ms | ✓ 1749ms | http |
| 38.34.178.186:8451 | ✓ 1711ms | 否 | ✓ 1654ms | ✓ 698ms | ✓ 1110ms | http |
| 114.237.77.231:1080 | 否 | 否 | ✓ 830ms | ✓ 1859ms | ✓ 948ms | http |
| 85.208.108.43:2094 | ✓ 700ms | 否 | ✓ 1183ms | 否 | ✓ 887ms | http |
| 20.78.26.206:8561 | ✓ 1839ms | ✓ 1700ms | ✓ 1358ms | ✓ 1626ms | ✓ 1394ms | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 1939ms | ✓ 1108ms | ✓ 1321ms | http |
| 103.113.70.189:1081 | ✓ 316ms | ✓ 1193ms | ✓ 668ms | ✓ 1359ms | ✓ 897ms | http |
| 167.71.196.28:8080 | ✓ 1397ms | 否 | ✓ 1461ms | ✓ 1008ms | 否 | http |
| 193.233.22.29:10808 | ✓ 1136ms | 否 | ✓ 1493ms | 否 | ✓ 1450ms | http |
| 106.75.15.167:7890 | ✓ 938ms | ✓ 1089ms | 否 | ✓ 1257ms | ✓ 884ms | http |
| 217.77.102.18:3128 | ✓ 1373ms | 否 | ✓ 1686ms | 否 | ✓ 1965ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 840ms | ✓ 1068ms | ✓ 1884ms | http |
| 116.80.49.159:3172 | ✓ 1891ms | 否 | 否 | ✓ 1764ms | ✓ 1851ms | http |
| 24.199.124.152:3128 | ✓ 247ms | 否 | ✓ 958ms | ✓ 623ms | ✓ 489ms | http |
| 45.140.147.82:1081 | ✓ 779ms | ✓ 1404ms | ✓ 1658ms | ✓ 1514ms | ✓ 1305ms | http |
| 210.223.44.230:3128 | ✓ 1270ms | 否 | ✓ 1301ms | 否 | ✓ 1255ms | http |
| 20.27.14.220:8561 | ✓ 831ms | 否 | ✓ 1302ms | ✓ 1583ms | ✓ 1134ms | http |
| 115.231.37.249:10198 | ✓ 1379ms | 否 | ✓ 1210ms | 否 | ✓ 1409ms | http |
| 5.104.87.17:8050 | 否 | 否 | ✓ 905ms | ✓ 1041ms | ✓ 700ms | http |
| 45.136.198.40:3128 | ✓ 868ms | 否 | ✓ 956ms | ✓ 1738ms | ✓ 1411ms | http |
| 101.255.119.26:8080 | 否 | 否 | ✓ 1668ms | ✓ 1490ms | ✓ 1335ms | http |
| 160.238.65.2:3128 | 否 | ✓ 1965ms | ✓ 1293ms | ✓ 1888ms | 否 | http |
| 218.89.134.230:3333 | ✓ 1848ms | 否 | ✓ 1991ms | 否 | ✓ 1756ms | http |
| 5.102.109.41:999 | ✓ 695ms | ✓ 1989ms | ✓ 1737ms | 否 | 否 | http |
| 38.34.179.40:8446 | 否 | 否 | ✓ 1859ms | ✓ 919ms | ✓ 1573ms | http |
| 20.210.76.175:8561 | ✓ 1141ms | ✓ 1530ms | ✓ 1338ms | ✓ 1683ms | ✓ 1508ms | http |

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
