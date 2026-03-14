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

最后更新：2026-03-14 00:23:56 UTC（2026-03-14 08:23:56 UTC+8）

**代理总数：74**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.218.82:8443 | ✓ 681ms | ✓ 761ms | ✓ 198ms | ✓ 730ms | ✓ 691ms | http |
| 38.145.203.135:8443 | ✓ 670ms | ✓ 852ms | ✓ 279ms | ✓ 726ms | ✓ 579ms | http |
| 205.209.118.30:3138 | ✓ 378ms | ✓ 1934ms | ✓ 936ms | ✓ 1281ms | ✓ 998ms | http |
| 216.180.127.45:1080 | ✓ 503ms | 否 | ✓ 931ms | ✓ 1227ms | 否 | http |
| 86.53.183.16:1080 | ✓ 1224ms | 否 | ✓ 1315ms | 否 | ✓ 1505ms | http |
| 113.160.132.26:8080 | ✓ 1845ms | 否 | ✓ 1760ms | ✓ 1209ms | ✓ 944ms | http |
| 38.145.208.95:8443 | ✓ 646ms | ✓ 675ms | ✓ 319ms | ✓ 728ms | ✓ 995ms | http |
| 38.145.203.235:8443 | ✓ 672ms | ✓ 672ms | ✓ 301ms | ✓ 740ms | ✓ 1857ms | http |
| 38.145.208.139:8443 | ✓ 670ms | ✓ 705ms | ✓ 266ms | ✓ 752ms | 否 | http |
| 38.145.208.143:8443 | ✓ 670ms | ✓ 696ms | ✓ 276ms | ✓ 800ms | 否 | http |
| 38.145.208.144:8443 | ✓ 680ms | ✓ 719ms | ✓ 244ms | ✓ 798ms | 否 | http |
| 38.145.220.192:8443 | ✓ 670ms | ✓ 690ms | ✓ 281ms | ✓ 943ms | 否 | http |
| 38.145.208.201:8447 | ✓ 646ms | ✓ 1305ms | ✓ 147ms | ✓ 772ms | 否 | http |
| 45.136.130.215:8443 | ✓ 889ms | 否 | ✓ 928ms | ✓ 1379ms | ✓ 1520ms | http |
| 38.145.218.189:8447 | ✓ 669ms | ✓ 1410ms | ✓ 776ms | 否 | 否 | http |
| 62.60.177.204:34094 | ✓ 475ms | ✓ 1285ms | ✓ 1272ms | ✓ 1166ms | ✓ 771ms | http |
| 120.92.212.16:7890 | ✓ 1114ms | ✓ 1223ms | ✓ 1270ms | 否 | ✓ 1230ms | http |
| 38.145.208.138:8447 | ✓ 525ms | ✓ 798ms | ✓ 495ms | 否 | 否 | http |
| 38.145.203.245:8443 | ✓ 523ms | ✓ 690ms | ✓ 755ms | 否 | 否 | http |
| 45.136.130.211:8447 | ✓ 523ms | ✓ 734ms | ✓ 558ms | 否 | ✓ 629ms | http |
| 38.145.208.98:8443 | ✓ 1312ms | 否 | ✓ 172ms | ✓ 773ms | ✓ 696ms | http |
| 85.198.96.242:3128 | ✓ 663ms | ✓ 1982ms | ✓ 642ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 962ms | ✓ 1315ms | ✓ 1064ms | ✓ 1222ms | ✓ 967ms | http |
| 81.70.169.194:80 | ✓ 1052ms | ✓ 1325ms | ✓ 1090ms | ✓ 1216ms | ✓ 1017ms | http |
| 171.251.173.39:5104 | 否 | 否 | ✓ 1615ms | ✓ 1679ms | ✓ 1526ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1364ms | 否 | ✓ 1757ms | ✓ 964ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1197ms | ✓ 1467ms | ✓ 1132ms | http |
| 35.225.22.61:80 | ✓ 1131ms | 否 | ✓ 594ms | ✓ 1146ms | 否 | http |
| 95.3.9.78:8080 | ✓ 1839ms | 否 | 否 | ✓ 1744ms | ✓ 1351ms | http |
| 95.3.9.78:3128 | ✓ 1830ms | 否 | 否 | ✓ 1793ms | ✓ 1365ms | http |
| 120.92.212.16:8890 | ✓ 1347ms | 否 | ✓ 953ms | 否 | ✓ 981ms | http |
| 38.145.218.101:8447 | ✓ 1127ms | ✓ 674ms | ✓ 213ms | ✓ 735ms | ✓ 788ms | http |
| 45.136.131.28:8447 | ✓ 780ms | ✓ 685ms | ✓ 1046ms | ✓ 1597ms | ✓ 850ms | http |
| 45.136.131.30:8447 | ✓ 783ms | ✓ 697ms | ✓ 1031ms | ✓ 1601ms | ✓ 863ms | http |
| 62.113.119.14:8080 | ✓ 778ms | ✓ 1580ms | ✓ 868ms | 否 | 否 | http |
| 45.136.131.42:8447 | ✓ 1213ms | ✓ 1626ms | ✓ 1419ms | ✓ 1212ms | 否 | http |
| 45.136.130.245:8447 | 否 | 否 | ✓ 1026ms | ✓ 890ms | ✓ 586ms | http |
| 38.145.220.25:8447 | 否 | 否 | ✓ 461ms | ✓ 849ms | ✓ 598ms | http |
| 170.78.208.245:999 | ✓ 1896ms | ✓ 1960ms | ✓ 1275ms | 否 | 否 | http |
| 38.145.220.60:8447 | ✓ 342ms | ✓ 1299ms | ✓ 1627ms | ✓ 948ms | ✓ 959ms | http |
| 103.84.95.54:7890 | ✓ 1069ms | 否 | ✓ 669ms | 否 | ✓ 865ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1129ms | ✓ 1949ms | ✓ 572ms | http |
| 101.47.73.135:3128 | ✓ 1342ms | 否 | ✓ 1917ms | ✓ 1092ms | ✓ 1149ms | http |
| 45.136.130.217:8443 | ✓ 1763ms | ✓ 1589ms | 否 | ✓ 1366ms | 否 | http |
| 45.136.130.236:8443 | ✓ 1163ms | ✓ 1133ms | ✓ 210ms | ✓ 766ms | ✓ 598ms | http |
| 3.9.173.53:5555 | 否 | 否 | ✓ 1851ms | ✓ 1974ms | ✓ 1647ms | http |
| 210.223.44.230:3128 | ✓ 1494ms | ✓ 1890ms | ✓ 1520ms | ✓ 1272ms | 否 | http |
| 18.201.114.187:52179 | 否 | 否 | ✓ 1862ms | ✓ 1749ms | ✓ 1809ms | http |
| 141.98.197.133:18791 | ✓ 1290ms | ✓ 1334ms | ✓ 1215ms | ✓ 1196ms | ✓ 858ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1788ms | ✓ 1860ms | ✓ 1884ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1474ms | ✓ 1771ms | ✓ 1387ms | ✓ 1599ms | ✓ 1428ms | http |
| 150.230.249.50:1080 | ✓ 1687ms | 否 | ✓ 1196ms | 否 | ✓ 1884ms | http |
| 38.145.218.162:8443 | ✓ 525ms | ✓ 728ms | ✓ 878ms | ✓ 749ms | ✓ 581ms | http |
| 38.145.218.161:8443 | ✓ 513ms | ✓ 716ms | ✓ 888ms | ✓ 751ms | ✓ 579ms | http |
| 45.136.131.35:8443 | ✓ 1397ms | ✓ 830ms | ✓ 220ms | ✓ 766ms | ✓ 682ms | http |
| 45.136.131.31:8443 | ✓ 1427ms | ✓ 810ms | ✓ 146ms | ✓ 862ms | ✓ 681ms | http |
| 45.136.131.33:8443 | ✓ 1440ms | ✓ 809ms | ✓ 146ms | ✓ 851ms | ✓ 681ms | http |
| 38.145.218.163:8443 | ✓ 181ms | ✓ 710ms | ✓ 142ms | ✓ 734ms | ✓ 568ms | http |
| 116.80.62.22:3128 | 否 | 否 | ✓ 1562ms | ✓ 1842ms | ✓ 1691ms | http |
| 193.148.58.165:3128 | ✓ 687ms | 否 | ✓ 1535ms | ✓ 1904ms | ✓ 1670ms | http |
| 211.171.114.154:3128 | ✓ 1812ms | 否 | ✓ 1088ms | ✓ 1448ms | 否 | http |
| 137.184.6.117:3128 | ✓ 1761ms | ✓ 899ms | ✓ 1084ms | ✓ 763ms | ✓ 565ms | http |
| 45.129.141.143:3128 | ✓ 1248ms | 否 | ✓ 1924ms | ✓ 1940ms | ✓ 1767ms | http |
| 137.184.14.135:3128 | ✓ 316ms | ✓ 1312ms | ✓ 1000ms | ✓ 767ms | ✓ 562ms | http |
| 45.136.198.40:3128 | ✓ 1959ms | ✓ 1689ms | ✓ 775ms | ✓ 1619ms | ✓ 1242ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1800ms | ✓ 1427ms | ✓ 1441ms | http |
| 116.80.49.170:3172 | ✓ 1951ms | 否 | 否 | ✓ 1989ms | ✓ 1689ms | http |
| 190.52.104.214:999 | ✓ 1197ms | 否 | ✓ 1478ms | ✓ 1751ms | 否 | http |
| 45.136.131.32:8443 | ✓ 676ms | ✓ 719ms | ✓ 140ms | ✓ 731ms | ✓ 865ms | http |
| 45.136.131.39:8443 | ✓ 676ms | ✓ 1290ms | ✓ 635ms | ✓ 741ms | ✓ 730ms | http |
| 54.90.159.174:27740 | ✓ 1975ms | 否 | ✓ 1054ms | 否 | ✓ 1857ms | http |
| 61.52.131.172:8443 | ✓ 876ms | ✓ 1142ms | 否 | 否 | ✓ 978ms | http |
| 172.212.68.37:3128 | ✓ 1032ms | ✓ 1576ms | ✓ 1051ms | ✓ 1474ms | ✓ 1693ms | http |
| 43.167.227.161:1080 | 否 | 否 | ✓ 1398ms | ✓ 1264ms | ✓ 680ms | http |

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
