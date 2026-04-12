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

最后更新：2026-04-12 11:32:12 UTC（2026-04-12 19:32:12 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 805ms | 否 | 否 | ✓ 1176ms | ✓ 944ms | http |
| 147.161.210.140:8800 | ✓ 1076ms | ✓ 1622ms | ✓ 1022ms | ✓ 1334ms | ✓ 1220ms | http |
| 167.103.115.102:8800 | ✓ 1725ms | 否 | ✓ 1464ms | 否 | ✓ 1478ms | http |
| 45.167.124.52:8080 | ✓ 1223ms | 否 | ✓ 1510ms | ✓ 1685ms | ✓ 1372ms | http |
| 167.103.34.108:8800 | ✓ 1480ms | 否 | ✓ 1608ms | ✓ 1568ms | ✓ 1622ms | http |
| 159.223.225.118:8888 | ✓ 1222ms | ✓ 1414ms | ✓ 847ms | ✓ 1455ms | ✓ 1159ms | http |
| 46.30.46.133:3128 | ✓ 1234ms | 否 | ✓ 551ms | ✓ 1604ms | ✓ 1147ms | http |
| 167.103.144.127:8800 | ✓ 1379ms | 否 | 否 | ✓ 1484ms | ✓ 1316ms | http |
| 167.103.31.122:8800 | 否 | 否 | ✓ 1761ms | ✓ 1512ms | ✓ 1645ms | http |
| 162.240.154.26:3128 | ✓ 1585ms | ✓ 1243ms | 否 | 否 | ✓ 937ms | http |
| 171.227.167.109:1004 | ✓ 1511ms | 否 | 否 | ✓ 1906ms | ✓ 1626ms | http |
| 113.160.132.26:8080 | ✓ 1845ms | ✓ 1784ms | ✓ 1165ms | 否 | 否 | http |
| 147.161.239.240:8800 | ✓ 623ms | ✓ 1487ms | ✓ 894ms | ✓ 1375ms | ✓ 1421ms | http |
| 5.196.101.18:3128 | 否 | 否 | ✓ 968ms | ✓ 1967ms | ✓ 1732ms | http |
| 45.136.130.169:8444 | ✓ 1177ms | 否 | ✓ 618ms | ✓ 1374ms | 否 | http |
| 38.145.203.39:8445 | ✓ 1153ms | 否 | ✓ 478ms | ✓ 1575ms | 否 | http |
| 38.34.183.47:8452 | ✓ 1195ms | 否 | ✓ 1323ms | ✓ 1178ms | ✓ 1562ms | http |
| 38.145.208.229:8453 | ✓ 1150ms | ✓ 1953ms | ✓ 1529ms | ✓ 1210ms | ✓ 1846ms | http |
| 38.34.179.178:8444 | ✓ 1154ms | ✓ 1666ms | ✓ 878ms | 否 | ✓ 1800ms | http |
| 38.145.220.39:8449 | ✓ 1153ms | ✓ 1672ms | ✓ 1999ms | ✓ 1072ms | ✓ 1196ms | http |
| 45.136.131.28:8449 | ✓ 1171ms | 否 | ✓ 969ms | ✓ 1655ms | ✓ 1500ms | http |
| 38.145.203.41:8453 | ✓ 1150ms | 否 | ✓ 479ms | ✓ 1977ms | 否 | http |
| 45.140.147.155:1082 | ✓ 479ms | 否 | ✓ 1266ms | ✓ 1092ms | 否 | http |
| 46.39.105.157:8080 | ✓ 614ms | ✓ 1801ms | ✓ 1791ms | ✓ 1521ms | ✓ 1307ms | http |
| 5.104.87.17:8051 | ✓ 1971ms | 否 | ✓ 1846ms | ✓ 1877ms | ✓ 1367ms | http |
| 223.84.151.86:30005 | ✓ 1779ms | ✓ 1906ms | ✓ 1694ms | 否 | 否 | http |
| 103.113.70.189:1081 | ✓ 281ms | ✓ 1555ms | ✓ 59ms | 否 | ✓ 963ms | http |
| 38.145.208.209:8447 | 否 | ✓ 1591ms | 否 | ✓ 1792ms | ✓ 1339ms | http |
| 109.107.179.140:31000 | 否 | 否 | ✓ 949ms | ✓ 1850ms | ✓ 1431ms | http |
| 120.92.108.86:7890 | ✓ 1637ms | 否 | 否 | ✓ 1888ms | ✓ 1756ms | http |
| 59.46.216.131:30001 | ✓ 1249ms | ✓ 1602ms | ✓ 1303ms | 否 | 否 | http |
| 38.34.179.13:8451 | ✓ 1419ms | 否 | ✓ 1992ms | ✓ 997ms | ✓ 1477ms | http |
| 38.34.179.40:8446 | ✓ 1590ms | 否 | ✓ 1002ms | 否 | ✓ 983ms | http |
| 103.113.70.189:1082 | ✓ 283ms | 否 | ✓ 923ms | ✓ 1196ms | ✓ 761ms | http |
| 45.186.6.104:3128 | ✓ 1321ms | ✓ 1952ms | ✓ 1733ms | 否 | 否 | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1879ms | ✓ 1728ms | ✓ 1096ms | http |
| 108.181.0.167:8080 | ✓ 478ms | ✓ 1306ms | ✓ 1194ms | 否 | ✓ 721ms | http |
| 139.159.99.242:8080 | ✓ 984ms | ✓ 1239ms | ✓ 1014ms | ✓ 1288ms | ✓ 1090ms | http |
| 34.85.118.216:3128 | ✓ 704ms | ✓ 1296ms | ✓ 1570ms | ✓ 1133ms | ✓ 821ms | http |
| 178.159.94.76:3128 | ✓ 612ms | ✓ 1744ms | ✓ 1538ms | 否 | ✓ 1235ms | http |
| 95.214.9.93:3128 | ✓ 972ms | 否 | ✓ 1746ms | ✓ 1849ms | ✓ 1374ms | http |
| 20.78.213.56:80 | ✓ 801ms | 否 | 否 | ✓ 1755ms | ✓ 1112ms | http |
| 36.141.21.200:7890 | ✓ 1690ms | 否 | ✓ 1479ms | 否 | ✓ 1521ms | http |
| 38.180.2.107:3128 | ✓ 1781ms | ✓ 1630ms | ✓ 1654ms | 否 | ✓ 1994ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1379ms | ✓ 1308ms | ✓ 1336ms | http |
| 8.219.97.248:80 | ✓ 1313ms | 否 | ✓ 1200ms | 否 | ✓ 1519ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1171ms | ✓ 1514ms | ✓ 1190ms | http |
| 202.141.161.53:10808 | ✓ 1241ms | ✓ 1636ms | ✓ 1606ms | 否 | ✓ 1273ms | http |
| 114.237.77.231:1080 | ✓ 1243ms | ✓ 1465ms | ✓ 1533ms | 否 | ✓ 1978ms | http |
| 36.103.198.235:7890 | ✓ 1182ms | ✓ 1529ms | 否 | ✓ 1636ms | ✓ 1291ms | http |
| 160.187.174.249:80 | 否 | 否 | ✓ 1965ms | ✓ 1672ms | ✓ 1682ms | http |
| 5.255.123.43:1080 | 否 | ✓ 1865ms | ✓ 1058ms | 否 | ✓ 1854ms | http |
| 115.231.181.40:8128 | ✓ 1847ms | 否 | ✓ 1566ms | 否 | ✓ 1162ms | http |
| 218.108.131.186:17890 | ✓ 1762ms | ✓ 1228ms | ✓ 1073ms | ✓ 1294ms | ✓ 1079ms | http |
| 212.58.132.5:8888 | ✓ 1068ms | 否 | 否 | ✓ 1493ms | ✓ 1187ms | http |
| 34.96.238.40:8080 | ✓ 1282ms | ✓ 1351ms | ✓ 1357ms | 否 | 否 | http |
| 185.76.240.167:10001 | ✓ 788ms | ✓ 1949ms | ✓ 770ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 966ms | 否 | ✓ 1619ms | ✓ 1391ms | ✓ 1190ms | http |
| 34.101.184.164:3128 | ✓ 1678ms | 否 | ✓ 1504ms | ✓ 1499ms | 否 | http |
| 181.78.44.63:999 | ✓ 769ms | ✓ 1922ms | ✓ 1748ms | ✓ 1759ms | ✓ 1779ms | http |

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
