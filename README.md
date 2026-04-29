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

最后更新：2026-04-29 20:32:55 UTC（2026-04-30 04:32:55 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | 否 | 否 | ✓ 647ms | ✓ 1249ms | ✓ 1163ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1981ms | 否 | ✓ 1758ms | ✓ 1208ms | http |
| 34.71.229.255:3128 | ✓ 1022ms | 否 | ✓ 1848ms | ✓ 1743ms | ✓ 1640ms | http |
| 1.231.81.166:3128 | ✓ 1688ms | ✓ 1836ms | 否 | ✓ 1618ms | ✓ 1879ms | http |
| 45.167.124.71:999 | ✓ 876ms | ✓ 1759ms | ✓ 1512ms | ✓ 1903ms | ✓ 1790ms | http |
| 218.108.131.186:17890 | ✓ 1204ms | 否 | ✓ 1180ms | ✓ 1308ms | ✓ 1118ms | http |
| 3.121.42.224:21347 | ✓ 1049ms | 否 | ✓ 680ms | ✓ 1674ms | ✓ 1296ms | http |
| 52.47.115.41:10603 | ✓ 1064ms | 否 | ✓ 647ms | 否 | ✓ 1605ms | http |
| 35.182.12.78:30808 | ✓ 1551ms | 否 | ✓ 849ms | 否 | ✓ 1571ms | http |
| 34.246.183.20:17994 | ✓ 1262ms | 否 | ✓ 1495ms | 否 | ✓ 1751ms | http |
| 54.188.236.206:29223 | ✓ 1337ms | 否 | ✓ 1155ms | 否 | ✓ 1899ms | http |
| 35.180.75.159:8174 | 否 | 否 | ✓ 602ms | ✓ 1837ms | ✓ 1532ms | http |
| 18.100.127.123:3434 | ✓ 1076ms | 否 | ✓ 1541ms | ✓ 1709ms | 否 | http |
| 13.60.181.61:4192 | ✓ 1064ms | 否 | ✓ 1773ms | 否 | ✓ 1748ms | http |
| 207.148.127.225:1234 | ✓ 910ms | ✓ 1736ms | ✓ 1567ms | 否 | ✓ 1065ms | http |
| 77.110.119.136:3128 | ✓ 901ms | ✓ 1798ms | ✓ 1445ms | 否 | 否 | http |
| 46.101.95.183:8888 | ✓ 1203ms | ✓ 1998ms | ✓ 1337ms | 否 | 否 | http |
| 168.110.52.228:3128 | ✓ 1533ms | ✓ 1309ms | ✓ 879ms | ✓ 1060ms | ✓ 859ms | http |
| 152.32.132.190:7890 | ✓ 885ms | 否 | ✓ 903ms | 否 | ✓ 1179ms | http |
| 115.231.181.40:8128 | ✓ 1161ms | ✓ 1343ms | 否 | ✓ 1343ms | ✓ 1200ms | http |
| 8.154.21.175:3128 | ✓ 1112ms | ✓ 1260ms | ✓ 1016ms | ✓ 1383ms | ✓ 1091ms | http |
| 86.104.72.220:1082 | ✓ 1431ms | ✓ 1103ms | ✓ 724ms | ✓ 1155ms | 否 | http |
| 154.64.232.35:8080 | ✓ 1605ms | ✓ 1391ms | ✓ 981ms | 否 | 否 | http |
| 86.104.72.219:1081 | ✓ 887ms | ✓ 1271ms | ✓ 961ms | 否 | 否 | http |
| 86.104.72.220:1081 | ✓ 570ms | 否 | ✓ 444ms | ✓ 1197ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1386ms | 否 | 否 | ✓ 1636ms | ✓ 1283ms | http |
| 212.58.132.5:8888 | ✓ 1145ms | 否 | ✓ 1516ms | ✓ 1614ms | ✓ 1408ms | http |
| 101.132.61.121:8888 | ✓ 1521ms | ✓ 1472ms | ✓ 1548ms | ✓ 1679ms | ✓ 1592ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1084ms | ✓ 1047ms | ✓ 1831ms | ✓ 1600ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1427ms | ✓ 1412ms | ✓ 1891ms | http |
| 38.92.10.152:57579 | ✓ 483ms | ✓ 1192ms | ✓ 430ms | 否 | 否 | http |
| 38.92.10.98:20058 | ✓ 543ms | ✓ 1230ms | ✓ 478ms | 否 | ✓ 828ms | http |
| 38.92.10.139:33985 | ✓ 480ms | ✓ 1153ms | ✓ 527ms | 否 | 否 | http |
| 75.119.151.237:3128 | ✓ 1238ms | ✓ 1645ms | ✓ 1595ms | 否 | ✓ 1492ms | http |
| 182.204.181.166:1080 | ✓ 1321ms | ✓ 1850ms | ✓ 1259ms | ✓ 1943ms | ✓ 1349ms | http |
| 124.217.74.168:8082 | ✓ 1754ms | 否 | ✓ 1911ms | ✓ 1706ms | ✓ 1662ms | http |
| 77.110.116.224:3128 | 否 | 否 | ✓ 1201ms | ✓ 1933ms | ✓ 1528ms | http |
| 52.59.51.29:3128 | ✓ 1200ms | 否 | ✓ 1577ms | ✓ 1972ms | 否 | http |
| 13.48.13.125:443 | ✓ 1775ms | 否 | ✓ 796ms | ✓ 1928ms | ✓ 1835ms | http |
| 15.161.131.175:47 | ✓ 1270ms | 否 | ✓ 1870ms | 否 | ✓ 1309ms | http |
| 62.113.119.14:8080 | ✓ 1239ms | ✓ 1571ms | ✓ 1925ms | ✓ 1697ms | ✓ 1490ms | http |
| 3.121.130.230:35026 | ✓ 1619ms | 否 | 否 | ✓ 1461ms | ✓ 1159ms | http |
| 120.92.212.16:8890 | ✓ 1348ms | ✓ 1696ms | 否 | 否 | ✓ 1494ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1875ms | ✓ 1685ms | ✓ 1700ms | http |
| 103.189.249.133:1111 | ✓ 1976ms | 否 | 否 | ✓ 1753ms | ✓ 1682ms | http |
| 121.230.8.136:1080 | ✓ 1254ms | ✓ 1665ms | ✓ 1212ms | ✓ 1731ms | ✓ 1267ms | http |
| 8.211.166.184:8081 | ✓ 668ms | ✓ 1411ms | ✓ 904ms | ✓ 1081ms | ✓ 907ms | http |
| 51.44.97.6:3095 | ✓ 1104ms | 否 | ✓ 1270ms | ✓ 1992ms | ✓ 1633ms | http |
| 120.92.212.16:7890 | ✓ 1162ms | 否 | ✓ 1356ms | ✓ 1464ms | ✓ 1359ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1445ms | ✓ 1852ms | ✓ 1534ms | http |
| 45.140.147.155:1081 | ✓ 1782ms | ✓ 1846ms | ✓ 1468ms | 否 | 否 | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 1574ms | ✓ 1995ms | ✓ 1220ms | http |
| 142.171.157.207:3128 | ✓ 927ms | ✓ 1842ms | ✓ 1739ms | 否 | ✓ 1895ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 1790ms | ✓ 1147ms | ✓ 1482ms | http |
| 86.104.74.110:1081 | ✓ 1350ms | ✓ 1882ms | ✓ 1218ms | 否 | 否 | http |
| 86.104.74.110:1082 | ✓ 925ms | ✓ 1578ms | ✓ 1057ms | 否 | 否 | http |
| 34.96.238.40:8080 | ✓ 1031ms | ✓ 1334ms | ✓ 1068ms | 否 | 否 | http |
| 103.209.36.58:8080 | ✓ 1333ms | 否 | 否 | ✓ 1485ms | ✓ 1418ms | http |
| 59.46.216.131:30001 | ✓ 1166ms | ✓ 1672ms | 否 | ✓ 1576ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1246ms | ✓ 1820ms | ✓ 1174ms | 否 | ✓ 1313ms | http |
| 172.236.145.31:7890 | ✓ 1138ms | 否 | ✓ 1048ms | ✓ 1255ms | ✓ 1360ms | http |
| 8.219.97.248:80 | ✓ 1767ms | 否 | ✓ 1787ms | ✓ 1840ms | 否 | http |
| 149.248.76.55:10000 | ✓ 707ms | 否 | ✓ 1128ms | ✓ 1868ms | ✓ 1926ms | http |
| 130.61.174.200:1080 | ✓ 886ms | ✓ 1359ms | ✓ 527ms | ✓ 1348ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1122ms | ✓ 1416ms | ✓ 1182ms | ✓ 1449ms | ✓ 1164ms | http |
| 51.92.173.133:8069 | ✓ 1361ms | 否 | ✓ 1139ms | 否 | ✓ 1929ms | http |
| 13.53.139.178:8081 | ✓ 1922ms | 否 | ✓ 1903ms | ✓ 1839ms | ✓ 1792ms | http |
| 103.39.51.207:8080 | ✓ 1757ms | 否 | 否 | ✓ 1924ms | ✓ 1706ms | http |

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
