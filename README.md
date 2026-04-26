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

最后更新：2026-04-26 04:16:56 UTC（2026-04-26 12:16:56 UTC+8）

**代理总数：53**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 46.101.95.183:8888 | ✓ 666ms | 否 | ✓ 1532ms | ✓ 1767ms | ✓ 1230ms | http |
| 80.92.204.47:1081 | ✓ 814ms | ✓ 1488ms | ✓ 1936ms | 否 | ✓ 1842ms | http |
| 47.85.51.197:1080 | ✓ 645ms | ✓ 1128ms | ✓ 379ms | ✓ 1744ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1593ms | ✓ 1450ms | 否 | ✓ 1355ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1234ms | ✓ 1814ms | ✓ 1261ms | 否 | ✓ 1517ms | http |
| 91.99.15.45:2095 | ✓ 1757ms | ✓ 1893ms | 否 | ✓ 1958ms | 否 | http |
| 43.133.90.161:8888 | ✓ 690ms | 否 | ✓ 1691ms | ✓ 1433ms | 否 | http |
| 218.108.131.186:17890 | ✓ 934ms | ✓ 1154ms | ✓ 1013ms | ✓ 1205ms | ✓ 1008ms | http |
| 36.141.21.200:7890 | ✓ 1047ms | ✓ 1307ms | ✓ 1089ms | ✓ 1349ms | ✓ 1097ms | http |
| 103.157.200.126:3128 | ✓ 1602ms | 否 | 否 | ✓ 1682ms | ✓ 1295ms | http |
| 47.95.231.180:8084 | ✓ 1003ms | ✓ 1323ms | ✓ 1033ms | 否 | 否 | http |
| 94.131.106.231:1081 | ✓ 1237ms | 否 | ✓ 1524ms | ✓ 1807ms | 否 | http |
| 117.236.124.166:3128 | ✓ 1570ms | 否 | ✓ 1965ms | ✓ 1947ms | ✓ 1345ms | http |
| 62.113.119.14:8080 | ✓ 613ms | 否 | ✓ 577ms | ✓ 1511ms | ✓ 1247ms | http |
| 206.206.126.177:2412 | ✓ 832ms | ✓ 1698ms | ✓ 871ms | ✓ 1148ms | ✓ 895ms | http |
| 45.140.147.82:1081 | ✓ 1398ms | ✓ 1782ms | ✓ 1886ms | 否 | 否 | http |
| 179.1.113.129:999 | ✓ 994ms | 否 | ✓ 1465ms | ✓ 1877ms | ✓ 1592ms | http |
| 217.52.247.69:1976 | ✓ 1956ms | 否 | ✓ 1799ms | 否 | ✓ 1825ms | http |
| 8.219.195.129:1080 | ✓ 1398ms | ✓ 1890ms | ✓ 956ms | ✓ 1218ms | ✓ 933ms | http |
| 2.27.54.161:1080 | ✓ 1289ms | 否 | ✓ 630ms | ✓ 1892ms | ✓ 1496ms | http |
| 120.92.108.86:7890 | ✓ 1502ms | 否 | ✓ 1724ms | 否 | ✓ 1514ms | http |
| 183.232.248.73:7890 | ✓ 1905ms | ✓ 1965ms | 否 | 否 | ✓ 1562ms | http |
| 120.92.212.16:8890 | ✓ 1161ms | ✓ 1940ms | ✓ 1139ms | ✓ 1292ms | ✓ 1068ms | http |
| 8.209.238.110:47701 | ✓ 621ms | 否 | ✓ 684ms | ✓ 1012ms | ✓ 817ms | http |
| 124.16.93.233:7890 | ✓ 1032ms | ✓ 1331ms | ✓ 1138ms | ✓ 1339ms | ✓ 1070ms | http |
| 194.31.87.77:3128 | ✓ 1605ms | 否 | ✓ 1659ms | 否 | ✓ 1698ms | http |
| 114.237.77.207:1080 | ✓ 1009ms | ✓ 1409ms | ✓ 1059ms | ✓ 1445ms | ✓ 1107ms | http |
| 43.133.44.89:8888 | ✓ 1925ms | 否 | ✓ 1034ms | 否 | ✓ 1175ms | http |
| 34.101.184.164:3128 | ✓ 907ms | 否 | ✓ 1825ms | ✓ 1357ms | ✓ 1071ms | http |
| 121.230.8.136:1080 | ✓ 1198ms | ✓ 1395ms | ✓ 1973ms | ✓ 1756ms | 否 | http |
| 128.199.113.85:9090 | ✓ 1614ms | 否 | ✓ 1375ms | ✓ 1246ms | ✓ 1192ms | http |
| 47.101.159.19:8899 | ✓ 970ms | ✓ 1208ms | ✓ 953ms | ✓ 1279ms | ✓ 1015ms | http |
| 168.144.75.9:3128 | ✓ 1438ms | 否 | ✓ 1414ms | ✓ 1954ms | ✓ 1833ms | http |
| 128.199.116.219:9090 | ✓ 997ms | 否 | 否 | ✓ 1669ms | ✓ 986ms | http |
| 42.200.76.16:3888 | ✓ 858ms | 否 | ✓ 876ms | ✓ 1107ms | ✓ 907ms | http |
| 128.199.114.189:9090 | ✓ 1016ms | 否 | ✓ 979ms | ✓ 1898ms | ✓ 1030ms | http |
| 128.199.254.13:9090 | ✓ 1197ms | 否 | ✓ 1526ms | ✓ 1649ms | ✓ 1480ms | http |
| 20.210.39.153:8561 | ✓ 841ms | ✓ 1217ms | ✓ 1389ms | ✓ 1626ms | ✓ 1518ms | http |
| 20.78.118.91:8561 | ✓ 832ms | ✓ 1371ms | ✓ 1228ms | ✓ 1628ms | ✓ 1513ms | http |
| 20.78.26.206:8561 | ✓ 850ms | ✓ 1304ms | ✓ 1305ms | ✓ 1617ms | ✓ 1540ms | http |
| 91.233.223.147:3128 | ✓ 1185ms | 否 | ✓ 1507ms | 否 | ✓ 1512ms | http |
| 160.238.65.9:3128 | 否 | 否 | ✓ 829ms | ✓ 1964ms | ✓ 1794ms | http |
| 121.230.9.160:1080 | ✓ 1424ms | ✓ 1743ms | ✓ 1651ms | 否 | ✓ 1028ms | http |
| 208.87.243.199:7878 | ✓ 1213ms | ✓ 1165ms | ✓ 503ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1110ms | 否 | ✓ 1053ms | 否 | ✓ 1179ms | http |
| 23.224.193.46:3128 | ✓ 1603ms | ✓ 1720ms | 否 | ✓ 1978ms | 否 | http |
| 148.153.56.51:80 | ✓ 675ms | ✓ 796ms | ✓ 1100ms | ✓ 1008ms | ✓ 809ms | http |
| 128.199.121.61:9090 | ✓ 1720ms | 否 | ✓ 1131ms | ✓ 1369ms | 否 | http |
| 146.190.80.158:9090 | 否 | 否 | ✓ 1058ms | ✓ 1560ms | ✓ 1629ms | http |
| 101.132.61.121:8888 | ✓ 1870ms | 否 | ✓ 1442ms | ✓ 1635ms | ✓ 1426ms | http |
| 105.159.136.255:4566 | ✓ 1183ms | ✓ 1696ms | ✓ 1437ms | 否 | ✓ 1797ms | http |
| 61.52.131.172:8443 | ✓ 1007ms | ✓ 1317ms | ✓ 1049ms | ✓ 1322ms | ✓ 1096ms | http |
| 201.144.20.238:3128 | ✓ 619ms | ✓ 1330ms | ✓ 1081ms | ✓ 1345ms | ✓ 1076ms | http |

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
