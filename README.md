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

最后更新：2026-05-03 17:48:16 UTC（2026-05-04 01:48:16 UTC+8）

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
| 1.231.81.166:3128 | ✓ 1748ms | ✓ 1085ms | ✓ 1744ms | ✓ 1260ms | ✓ 1069ms | http |
| 206.206.126.177:2412 | ✓ 1582ms | 否 | ✓ 1095ms | ✓ 1173ms | ✓ 899ms | http |
| 113.160.132.26:8080 | ✓ 1897ms | ✓ 1724ms | 否 | ✓ 1611ms | ✓ 1437ms | http |
| 45.167.124.71:999 | ✓ 830ms | 否 | ✓ 1274ms | ✓ 1950ms | ✓ 1575ms | http |
| 154.64.232.35:8080 | 否 | ✓ 1010ms | ✓ 1266ms | ✓ 960ms | ✓ 902ms | http |
| 8.211.166.184:8081 | ✓ 626ms | 否 | ✓ 760ms | ✓ 1038ms | ✓ 804ms | http |
| 218.108.131.186:17890 | ✓ 1016ms | ✓ 1171ms | ✓ 982ms | ✓ 1258ms | ✓ 1004ms | http |
| 168.110.52.228:3128 | ✓ 747ms | 否 | ✓ 1460ms | ✓ 1116ms | ✓ 1241ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 1038ms | ✓ 1146ms | ✓ 1036ms | http |
| 46.105.190.40:3128 | ✓ 944ms | ✓ 1815ms | ✓ 762ms | 否 | 否 | http |
| 20.78.118.91:8561 | ✓ 613ms | 否 | ✓ 900ms | ✓ 1256ms | ✓ 1210ms | http |
| 20.78.26.206:8561 | ✓ 664ms | 否 | ✓ 892ms | ✓ 1260ms | ✓ 1159ms | http |
| 20.210.39.153:8561 | ✓ 663ms | ✓ 1192ms | ✓ 785ms | ✓ 1117ms | ✓ 1351ms | http |
| 109.120.156.122:8090 | ✓ 1114ms | 否 | ✓ 1007ms | 否 | ✓ 1786ms | http |
| 152.32.132.190:7890 | ✓ 1355ms | 否 | ✓ 1114ms | 否 | ✓ 1083ms | http |
| 46.105.190.38:3128 | ✓ 896ms | 否 | ✓ 1332ms | 否 | ✓ 1684ms | http |
| 212.58.132.5:8888 | ✓ 1201ms | 否 | ✓ 1408ms | ✓ 1536ms | ✓ 1185ms | http |
| 80.92.204.47:1081 | ✓ 1233ms | 否 | ✓ 731ms | 否 | ✓ 1458ms | http |
| 45.140.147.155:1081 | ✓ 683ms | ✓ 1783ms | ✓ 1180ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 679ms | ✓ 1825ms | ✓ 1136ms | 否 | 否 | http |
| 103.35.190.69:1082 | ✓ 773ms | ✓ 1198ms | ✓ 1001ms | ✓ 1147ms | ✓ 1140ms | http |
| 148.230.4.241:999 | ✓ 1470ms | ✓ 1674ms | ✓ 665ms | 否 | 否 | http |
| 178.156.224.42:3128 | ✓ 1409ms | 否 | ✓ 1766ms | 否 | ✓ 1533ms | http |
| 157.230.220.25:4857 | ✓ 180ms | 否 | ✓ 1503ms | 否 | ✓ 759ms | http |
| 34.96.238.40:8080 | ✓ 1538ms | ✓ 1693ms | ✓ 1382ms | ✓ 1551ms | ✓ 1478ms | http |
| 62.113.119.14:8080 | ✓ 972ms | ✓ 1814ms | 否 | ✓ 1865ms | ✓ 1165ms | http |
| 117.236.124.166:3128 | ✓ 1097ms | 否 | ✓ 1912ms | 否 | ✓ 1803ms | http |
| 47.77.216.82:1080 | ✓ 1470ms | 否 | ✓ 756ms | 否 | ✓ 1647ms | http |
| 193.123.250.39:1080 | ✓ 1924ms | 否 | ✓ 1876ms | ✓ 1663ms | ✓ 928ms | http |
| 8.154.21.175:3128 | ✓ 961ms | ✓ 1154ms | ✓ 950ms | ✓ 1254ms | ✓ 1001ms | http |
| 167.71.196.178:80 | ✓ 888ms | 否 | ✓ 1220ms | ✓ 1197ms | ✓ 981ms | http |
| 45.125.67.37:8443 | ✓ 1089ms | 否 | ✓ 1019ms | ✓ 1447ms | ✓ 1980ms | http |
| 103.157.200.126:3128 | ✓ 1218ms | 否 | ✓ 1186ms | ✓ 1923ms | 否 | http |
| 47.179.58.156:8080 | ✓ 1945ms | ✓ 949ms | ✓ 1163ms | ✓ 1116ms | ✓ 845ms | http |
| 103.126.86.98:7777 | ✓ 1807ms | 否 | 否 | ✓ 1593ms | ✓ 1578ms | http |
| 196.204.83.229:8080 | ✓ 1459ms | ✓ 1905ms | 否 | 否 | ✓ 1781ms | http |
| 103.87.202.198:1111 | ✓ 1805ms | 否 | ✓ 1589ms | ✓ 1882ms | ✓ 1674ms | http |
| 103.217.216.65:8181 | ✓ 1797ms | 否 | 否 | ✓ 1627ms | ✓ 1616ms | http |
| 86.104.74.110:1081 | ✓ 1566ms | ✓ 1553ms | ✓ 888ms | ✓ 1597ms | ✓ 1261ms | http |
| 8.219.97.248:80 | ✓ 1093ms | 否 | ✓ 1914ms | ✓ 1988ms | 否 | http |
| 207.254.71.62:8088 | ✓ 942ms | ✓ 1657ms | ✓ 1533ms | ✓ 1687ms | ✓ 1694ms | http |
| 101.32.243.189:80 | ✓ 1136ms | ✓ 1692ms | 否 | ✓ 1614ms | ✓ 1416ms | http |
| 152.42.177.32:8888 | ✓ 1110ms | 否 | ✓ 1352ms | ✓ 1453ms | ✓ 1459ms | http |
| 120.92.108.86:7890 | ✓ 1330ms | 否 | ✓ 1312ms | ✓ 1930ms | ✓ 1629ms | http |
| 91.233.223.147:3128 | ✓ 1349ms | 否 | ✓ 935ms | ✓ 1973ms | ✓ 1509ms | http |
| 3.101.133.120:80 | ✓ 817ms | ✓ 1438ms | ✓ 1063ms | ✓ 1364ms | ✓ 1188ms | http |
| 42.200.76.16:3888 | ✓ 856ms | 否 | ✓ 858ms | ✓ 1117ms | ✓ 881ms | http |
| 170.9.253.20:8888 | ✓ 901ms | 否 | 否 | ✓ 1367ms | ✓ 1423ms | http |
| 106.10.55.212:1121 | ✓ 1608ms | ✓ 1635ms | 否 | ✓ 1331ms | ✓ 1405ms | http |
| 37.187.109.70:10111 | ✓ 1045ms | ✓ 1654ms | ✓ 1195ms | 否 | 否 | http |
| 103.209.36.58:8080 | ✓ 1425ms | 否 | ✓ 1339ms | ✓ 1529ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1337ms | 否 | ✓ 1306ms | ✓ 1990ms | ✓ 1132ms | http |
| 45.140.147.82:1081 | ✓ 802ms | ✓ 1590ms | 否 | 否 | ✓ 1419ms | http |
| 121.130.199.80:3128 | ✓ 1521ms | ✓ 1766ms | ✓ 1566ms | ✓ 1256ms | ✓ 1308ms | http |
| 190.12.150.244:999 | ✓ 901ms | ✓ 1673ms | ✓ 1670ms | 否 | 否 | http |
| 121.230.8.91:1080 | ✓ 1346ms | ✓ 1503ms | ✓ 1285ms | ✓ 1472ms | ✓ 1245ms | http |
| 121.230.8.55:1080 | ✓ 1294ms | ✓ 1591ms | ✓ 1139ms | ✓ 1459ms | ✓ 1452ms | http |
| 113.142.152.75:8088 | ✓ 1033ms | ✓ 1447ms | ✓ 1192ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1161ms | ✓ 1241ms | ✓ 1065ms | ✓ 1336ms | ✓ 1059ms | http |
| 103.172.70.173:8080 | ✓ 1527ms | 否 | 否 | ✓ 1833ms | ✓ 1570ms | http |
| 38.121.212.98:8080 | 否 | ✓ 1689ms | 否 | ✓ 1503ms | ✓ 1547ms | http |
| 47.84.131.156:8100 | ✓ 1483ms | ✓ 1967ms | 否 | ✓ 1253ms | 否 | http |
| 129.213.162.27:17777 | ✓ 997ms | 否 | 否 | ✓ 1447ms | ✓ 1038ms | http |
| 103.176.96.213:8080 | 否 | 否 | ✓ 1880ms | ✓ 1533ms | ✓ 1523ms | http |

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
