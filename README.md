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

最后更新：2026-04-10 12:50:44 UTC（2026-04-10 20:50:44 UTC+8）

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
| 218.108.131.186:17890 | ✓ 835ms | ✓ 720ms | ✓ 625ms | ✓ 798ms | ✓ 648ms | http |
| 147.161.210.140:8800 | ✓ 557ms | 否 | ✓ 725ms | ✓ 1051ms | ✓ 917ms | http |
| 113.160.132.26:8080 | ✓ 980ms | ✓ 1402ms | ✓ 1171ms | ✓ 1686ms | ✓ 1320ms | http |
| 167.103.115.102:8800 | ✓ 1848ms | 否 | ✓ 978ms | ✓ 1482ms | ✓ 945ms | http |
| 167.103.34.108:8800 | ✓ 1923ms | 否 | ✓ 1666ms | ✓ 1794ms | ✓ 1976ms | http |
| 38.92.10.152:57579 | ✓ 407ms | ✓ 1144ms | ✓ 377ms | ✓ 1168ms | ✓ 675ms | http |
| 167.103.144.127:8800 | ✓ 1471ms | ✓ 1819ms | ✓ 1424ms | ✓ 1626ms | ✓ 1532ms | http |
| 120.92.212.16:8890 | ✓ 1172ms | ✓ 1162ms | ✓ 1180ms | ✓ 1761ms | 否 | http |
| 202.141.161.53:10808 | ✓ 1937ms | ✓ 1849ms | ✓ 1331ms | 否 | ✓ 1156ms | http |
| 167.103.31.122:8800 | ✓ 1848ms | 否 | ✓ 1559ms | 否 | ✓ 1812ms | http |
| 59.46.216.131:30001 | ✓ 866ms | ✓ 1250ms | 否 | ✓ 1310ms | 否 | http |
| 147.161.239.240:8800 | ✓ 985ms | 否 | ✓ 1105ms | ✓ 1908ms | ✓ 1536ms | http |
| 38.92.10.139:33985 | ✓ 869ms | ✓ 922ms | ✓ 792ms | ✓ 950ms | 否 | http |
| 161.35.70.36:8888 | ✓ 1856ms | 否 | ✓ 704ms | ✓ 1896ms | 否 | http |
| 45.12.151.226:2829 | ✓ 867ms | 否 | ✓ 864ms | 否 | ✓ 1605ms | http |
| 155.117.18.36:25388 | ✓ 1075ms | ✓ 1051ms | ✓ 1334ms | ✓ 1429ms | ✓ 1243ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 607ms | ✓ 1259ms | ✓ 1093ms | http |
| 5.104.87.17:8051 | ✓ 1208ms | 否 | 否 | ✓ 1137ms | ✓ 697ms | http |
| 43.99.25.221:46509 | ✓ 662ms | ✓ 1077ms | ✓ 737ms | ✓ 733ms | ✓ 594ms | http |
| 1.231.81.166:3128 | ✓ 827ms | ✓ 1363ms | ✓ 1717ms | 否 | 否 | http |
| 101.43.127.100:8877 | ✓ 783ms | ✓ 1055ms | ✓ 976ms | ✓ 1391ms | ✓ 905ms | http |
| 38.34.179.106:8445 | 否 | 否 | ✓ 420ms | ✓ 807ms | ✓ 665ms | http |
| 115.231.181.40:8128 | ✓ 1029ms | ✓ 1868ms | 否 | 否 | ✓ 1106ms | http |
| 38.145.203.34:8444 | 否 | 否 | ✓ 1654ms | ✓ 1386ms | ✓ 1439ms | http |
| 91.238.104.171:2023 | ✓ 1476ms | ✓ 1941ms | ✓ 1684ms | 否 | ✓ 1874ms | http |
| 120.92.212.16:7890 | ✓ 1766ms | ✓ 1139ms | 否 | ✓ 1439ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1786ms | 否 | ✓ 1752ms | ✓ 1794ms | ✓ 1793ms | http |
| 45.136.130.247:8448 | 否 | 否 | ✓ 1446ms | ✓ 1000ms | ✓ 1957ms | http |
| 45.167.125.21:999 | ✓ 1285ms | 否 | ✓ 1766ms | 否 | ✓ 1636ms | http |
| 103.214.251.52:8080 | ✓ 1874ms | 否 | ✓ 1841ms | ✓ 1989ms | 否 | http |
| 38.34.179.46:8448 | ✓ 178ms | 否 | ✓ 1811ms | ✓ 842ms | ✓ 709ms | http |
| 47.74.226.8:5001 | ✓ 1185ms | ✓ 1264ms | ✓ 813ms | 否 | 否 | http |
| 38.92.10.98:20058 | ✓ 1682ms | ✓ 1125ms | ✓ 834ms | 否 | 否 | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1344ms | ✓ 1304ms | ✓ 1010ms | http |
| 38.147.160.208:24239 | ✓ 895ms | ✓ 752ms | ✓ 1161ms | ✓ 1406ms | 否 | http |
| 8.219.97.248:80 | ✓ 1267ms | 否 | ✓ 867ms | 否 | ✓ 1961ms | http |
| 8.209.238.110:47701 | ✓ 519ms | 否 | ✓ 426ms | ✓ 808ms | ✓ 640ms | http |
| 107.172.102.234:40621 | ✓ 755ms | 否 | ✓ 926ms | ✓ 1026ms | ✓ 1102ms | http |
| 34.96.238.40:8080 | ✓ 973ms | 否 | ✓ 951ms | ✓ 1353ms | 否 | http |
| 38.34.183.224:8448 | 否 | 否 | ✓ 1125ms | ✓ 792ms | ✓ 690ms | http |
| 91.238.105.64:2024 | ✓ 933ms | ✓ 1801ms | ✓ 1081ms | 否 | ✓ 1685ms | http |
| 201.139.115.38:8080 | 否 | 否 | ✓ 832ms | ✓ 1551ms | ✓ 1404ms | http |
| 168.110.52.228:3128 | ✓ 1559ms | 否 | 否 | ✓ 971ms | ✓ 745ms | http |
| 170.106.137.214:7890 | ✓ 992ms | ✓ 1556ms | 否 | 否 | ✓ 1843ms | http |
| 116.80.82.219:3172 | 否 | 否 | ✓ 1473ms | ✓ 1736ms | ✓ 1591ms | http |
| 94.72.109.214:8888 | ✓ 1468ms | 否 | ✓ 1097ms | 否 | ✓ 1755ms | http |
| 38.145.218.161:8445 | ✓ 1050ms | 否 | ✓ 591ms | ✓ 915ms | ✓ 1087ms | http |
| 103.157.200.126:3128 | ✓ 1601ms | 否 | ✓ 1514ms | ✓ 1960ms | ✓ 1529ms | http |
| 223.16.170.103:80 | ✓ 1022ms | 否 | ✓ 898ms | ✓ 960ms | ✓ 989ms | http |
| 38.145.218.227:8447 | 否 | 否 | ✓ 722ms | ✓ 755ms | ✓ 691ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1354ms | ✓ 1356ms | 否 | ✓ 1778ms | http |
| 77.93.89.128:47146 | ✓ 1253ms | 否 | ✓ 1511ms | ✓ 1092ms | ✓ 1035ms | http |
| 45.140.147.155:1082 | ✓ 1420ms | 否 | ✓ 1320ms | 否 | ✓ 1800ms | http |
| 5.255.123.43:1080 | ✓ 1401ms | 否 | ✓ 1356ms | ✓ 1973ms | 否 | http |
| 185.76.241.110:10001 | ✓ 1150ms | 否 | ✓ 1152ms | 否 | ✓ 1762ms | http |
| 121.230.9.125:1080 | ✓ 1429ms | ✓ 1843ms | 否 | 否 | ✓ 1114ms | http |
| 38.34.179.39:8452 | ✓ 1167ms | ✓ 1933ms | ✓ 1910ms | ✓ 1441ms | 否 | http |
| 34.101.184.164:3128 | ✓ 744ms | 否 | ✓ 1439ms | ✓ 1249ms | ✓ 1059ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1809ms | ✓ 1394ms | ✓ 870ms | http |
| 61.52.131.172:8443 | ✓ 640ms | ✓ 864ms | ✓ 707ms | ✓ 939ms | ✓ 863ms | http |
| 38.145.220.173:8450 | ✓ 673ms | 否 | ✓ 1846ms | ✓ 854ms | ✓ 1953ms | http |
| 38.34.179.63:8448 | ✓ 878ms | 否 | ✓ 219ms | ✓ 790ms | ✓ 917ms | http |
| 38.145.218.87:8451 | ✓ 1365ms | 否 | ✓ 1057ms | 否 | ✓ 1550ms | http |
| 38.145.220.39:8452 | ✓ 1001ms | 否 | ✓ 1349ms | ✓ 1421ms | 否 | http |
| 38.34.179.51:8449 | ✓ 1497ms | 否 | ✓ 1542ms | ✓ 1154ms | ✓ 1389ms | http |
| 185.76.240.254:10001 | ✓ 1553ms | 否 | ✓ 1108ms | 否 | ✓ 1697ms | http |
| 185.76.240.169:10001 | ✓ 1556ms | 否 | ✓ 1100ms | 否 | ✓ 1726ms | http |
| 101.32.75.4:8888 | ✓ 1299ms | ✓ 976ms | ✓ 1076ms | ✓ 1350ms | ✓ 884ms | http |

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
