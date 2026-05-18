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

最后更新：2026-05-18 09:39:48 UTC（2026-05-18 17:39:48 UTC+8）

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
| 51.161.50.166:3128 | ✓ 392ms | 否 | ✓ 1107ms | ✓ 1311ms | ✓ 1189ms | http |
| 185.200.188.234:10001 | ✓ 1761ms | 否 | ✓ 1665ms | 否 | ✓ 1721ms | http |
| 212.58.132.5:8888 | ✓ 1945ms | 否 | ✓ 1692ms | ✓ 1690ms | ✓ 1379ms | http |
| 218.108.131.186:17890 | ✓ 1662ms | ✓ 1907ms | 否 | ✓ 1492ms | 否 | http |
| 86.104.72.219:1082 | ✓ 258ms | ✓ 1091ms | ✓ 100ms | ✓ 1250ms | ✓ 740ms | http |
| 86.104.72.219:1081 | ✓ 180ms | ✓ 1754ms | ✓ 102ms | ✓ 1010ms | ✓ 746ms | http |
| 107.175.85.198:1080 | 否 | 否 | ✓ 1762ms | ✓ 1600ms | ✓ 1020ms | http |
| 103.147.152.12:1095 | ✓ 854ms | 否 | ✓ 847ms | ✓ 1534ms | ✓ 1181ms | http |
| 34.101.184.164:3128 | ✓ 1851ms | 否 | ✓ 1851ms | ✓ 1550ms | ✓ 1772ms | http |
| 113.160.132.26:8080 | ✓ 1387ms | 否 | ✓ 1005ms | 否 | ✓ 1216ms | http |
| 91.242.229.129:8092 | ✓ 1372ms | 否 | ✓ 1370ms | ✓ 1636ms | ✓ 1198ms | http |
| 8.154.21.175:3128 | ✓ 1874ms | ✓ 1221ms | ✓ 1085ms | ✓ 1291ms | 否 | http |
| 114.214.165.78:10810 | ✓ 1558ms | 否 | ✓ 1552ms | ✓ 1712ms | ✓ 1411ms | http |
| 148.230.4.241:999 | ✓ 796ms | ✓ 1882ms | ✓ 748ms | ✓ 1802ms | ✓ 1272ms | http |
| 122.2.48.121:8080 | ✓ 1467ms | 否 | ✓ 1595ms | ✓ 1418ms | ✓ 1416ms | http |
| 5.252.33.13:2025 | ✓ 1417ms | 否 | ✓ 1472ms | 否 | ✓ 1825ms | http |
| 103.147.152.12:1080 | ✓ 1888ms | ✓ 1479ms | 否 | ✓ 1521ms | 否 | http |
| 170.106.136.181:31002 | ✓ 978ms | ✓ 703ms | ✓ 369ms | ✓ 824ms | ✓ 601ms | http |
| 84.47.150.125:1080 | ✓ 1742ms | 否 | 否 | ✓ 1981ms | ✓ 1731ms | http |
| 59.46.216.131:30001 | ✓ 1648ms | 否 | ✓ 1342ms | 否 | ✓ 1241ms | http |
| 114.214.170.41:27890 | ✓ 1381ms | 否 | ✓ 1530ms | ✓ 1631ms | ✓ 1393ms | http |
| 150.107.140.238:3128 | ✓ 1101ms | 否 | 否 | ✓ 1792ms | ✓ 1348ms | http |
| 128.199.114.189:9090 | ✓ 1842ms | 否 | ✓ 904ms | ✓ 1523ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1214ms | ✓ 1244ms | 否 | ✓ 1327ms | ✓ 1120ms | http |
| 120.92.212.16:8890 | ✓ 1555ms | ✓ 1773ms | ✓ 1893ms | 否 | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1757ms | ✓ 1960ms | ✓ 1490ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1899ms | ✓ 1099ms | ✓ 794ms | ✓ 1095ms | ✓ 890ms | http |
| 199.38.85.122:40010 | 否 | 否 | ✓ 1281ms | ✓ 1955ms | ✓ 1548ms | http |
| 199.38.85.122:40014 | ✓ 535ms | 否 | ✓ 1384ms | ✓ 1759ms | ✓ 1562ms | http |
| 77.110.107.80:1080 | ✓ 1348ms | ✓ 1891ms | ✓ 765ms | 否 | 否 | http |
| 129.80.238.83:444 | ✓ 1384ms | 否 | ✓ 1318ms | ✓ 1195ms | 否 | http |
| 82.114.228.67:1080 | ✓ 1392ms | ✓ 1754ms | ✓ 927ms | ✓ 1641ms | 否 | http |
| 183.98.143.134:8059 | ✓ 1669ms | ✓ 1047ms | ✓ 835ms | 否 | ✓ 899ms | http |
| 158.255.212.55:3256 | ✓ 907ms | 否 | ✓ 1715ms | ✓ 1847ms | 否 | http |
| 158.255.212.55:7497 | ✓ 943ms | 否 | ✓ 1685ms | ✓ 1842ms | 否 | http |
| 158.255.212.55:9480 | ✓ 572ms | 否 | ✓ 632ms | ✓ 1385ms | 否 | http |
| 128.199.113.85:9090 | 否 | 否 | ✓ 931ms | ✓ 1272ms | ✓ 1040ms | http |
| 43.153.28.68:3128 | ✓ 393ms | 否 | ✓ 722ms | ✓ 825ms | 否 | http |
| 128.199.116.219:9090 | ✓ 879ms | 否 | ✓ 973ms | ✓ 1345ms | ✓ 1088ms | http |
| 128.199.254.13:9090 | ✓ 1547ms | 否 | 否 | ✓ 1929ms | ✓ 1244ms | http |
| 57.129.144.178:40000 | ✓ 1081ms | 否 | ✓ 1387ms | ✓ 1737ms | ✓ 1532ms | http |
| 128.199.121.61:9090 | ✓ 1589ms | 否 | ✓ 1302ms | ✓ 1745ms | 否 | http |
| 104.248.151.93:9090 | ✓ 1548ms | 否 | ✓ 906ms | ✓ 1296ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1732ms | 否 | ✓ 1707ms | 否 | ✓ 1809ms | http |
| 121.130.177.28:8888 | ✓ 1608ms | ✓ 1288ms | 否 | ✓ 1524ms | ✓ 1759ms | http |
| 3.101.133.120:80 | ✓ 307ms | ✓ 1293ms | ✓ 748ms | ✓ 1343ms | ✓ 979ms | http |
| 62.113.119.14:8080 | ✓ 618ms | ✓ 1515ms | ✓ 1226ms | ✓ 1762ms | ✓ 1143ms | http |
| 147.45.78.89:1080 | ✓ 1187ms | 否 | ✓ 1236ms | ✓ 1776ms | ✓ 1936ms | http |
| 152.42.170.187:9090 | ✓ 1441ms | 否 | ✓ 986ms | ✓ 1810ms | 否 | http |
| 42.200.76.16:3888 | 否 | 否 | ✓ 874ms | ✓ 1108ms | ✓ 880ms | http |
| 216.106.179.216:49152 | ✓ 1758ms | ✓ 807ms | ✓ 875ms | 否 | 否 | http |
| 183.98.143.134:8048 | ✓ 1741ms | ✓ 1676ms | ✓ 1305ms | ✓ 1191ms | ✓ 944ms | http |
| 183.98.143.134:8069 | ✓ 1751ms | 否 | ✓ 1201ms | ✓ 1169ms | ✓ 928ms | http |
| 129.80.217.21:444 | ✓ 759ms | 否 | ✓ 960ms | ✓ 931ms | ✓ 739ms | http |
| 91.243.192.17:3128 | ✓ 1261ms | 否 | 否 | ✓ 1894ms | ✓ 1890ms | http |
| 20.164.75.153:8080 | ✓ 1524ms | 否 | ✓ 1701ms | 否 | ✓ 1850ms | http |
| 146.190.80.158:9090 | ✓ 1976ms | 否 | ✓ 1261ms | ✓ 1334ms | ✓ 1222ms | http |
| 180.125.216.109:8118 | ✓ 1162ms | 否 | ✓ 1112ms | 否 | ✓ 1208ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1877ms | ✓ 1870ms | ✓ 1776ms | http |
| 61.52.131.172:8443 | ✓ 1208ms | ✓ 1333ms | ✓ 1112ms | ✓ 1293ms | ✓ 1141ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1171ms | 否 | ✓ 1199ms | ✓ 949ms | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1708ms | ✓ 1536ms | ✓ 1632ms | http |
| 144.124.227.90:21074 | ✓ 961ms | 否 | ✓ 1436ms | 否 | ✓ 1820ms | http |
| 103.157.117.116:8080 | ✓ 1969ms | 否 | 否 | ✓ 1896ms | ✓ 1880ms | http |

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
