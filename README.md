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

最后更新：2026-03-04 19:53:35 UTC（2026-03-05 03:53:35 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 942ms | ✓ 1472ms | ✓ 880ms | ✓ 1390ms | ✓ 1126ms | http |
| 162.240.154.26:3128 | ✓ 971ms | ✓ 1749ms | ✓ 1783ms | ✓ 1602ms | ✓ 1322ms | http |
| 61.72.110.94:3128 | ✓ 896ms | 否 | ✓ 1739ms | ✓ 1558ms | ✓ 1216ms | http |
| 61.72.221.94:3128 | ✓ 1339ms | 否 | 否 | ✓ 1299ms | ✓ 981ms | http |
| 125.128.12.144:3128 | 否 | 否 | ✓ 1684ms | ✓ 956ms | ✓ 808ms | http |
| 35.225.22.61:80 | ✓ 400ms | 否 | ✓ 1090ms | ✓ 1139ms | ✓ 1184ms | http |
| 14.56.107.244:3128 | ✓ 1641ms | ✓ 1341ms | ✓ 570ms | 否 | ✓ 745ms | http |
| 45.140.147.82:1082 | ✓ 639ms | ✓ 1229ms | ✓ 1265ms | 否 | ✓ 1195ms | http |
| 45.140.147.82:1081 | ✓ 629ms | 否 | ✓ 640ms | ✓ 1859ms | ✓ 1836ms | http |
| 61.72.110.54:3128 | ✓ 1315ms | 否 | ✓ 1636ms | 否 | ✓ 814ms | http |
| 61.72.221.234:3128 | ✓ 931ms | ✓ 1208ms | ✓ 1260ms | ✓ 1035ms | 否 | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 819ms | ✓ 1172ms | ✓ 833ms | http |
| 121.128.121.54:3128 | 否 | ✓ 1791ms | ✓ 537ms | 否 | ✓ 1767ms | http |
| 14.56.177.44:3128 | ✓ 695ms | ✓ 1333ms | ✓ 558ms | ✓ 912ms | ✓ 751ms | http |
| 125.128.12.14:3128 | ✓ 887ms | ✓ 990ms | ✓ 908ms | ✓ 991ms | ✓ 769ms | http |
| 103.84.95.54:7890 | ✓ 1616ms | 否 | ✓ 1244ms | ✓ 1591ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1386ms | 否 | ✓ 1663ms | 否 | ✓ 1825ms | http |
| 165.227.5.10:8888 | ✓ 47ms | ✓ 667ms | ✓ 1096ms | ✓ 1096ms | ✓ 495ms | http |
| 101.43.255.96:80 | ✓ 1018ms | ✓ 1280ms | ✓ 938ms | ✓ 1262ms | ✓ 1022ms | http |
| 81.70.169.194:80 | ✓ 919ms | ✓ 1297ms | ✓ 970ms | ✓ 1322ms | ✓ 1030ms | http |
| 160.238.65.5:3128 | 否 | ✓ 1611ms | ✓ 1546ms | ✓ 1931ms | ✓ 1684ms | http |
| 160.238.65.4:3128 | 否 | 否 | ✓ 1168ms | ✓ 1925ms | ✓ 1664ms | http |
| 160.238.65.7:3128 | 否 | ✓ 1652ms | ✓ 1510ms | ✓ 1924ms | ✓ 1687ms | http |
| 160.238.65.2:3128 | 否 | 否 | ✓ 1182ms | ✓ 1915ms | ✓ 1705ms | http |
| 160.238.65.9:3128 | ✓ 1848ms | ✓ 1740ms | ✓ 1568ms | ✓ 1916ms | ✓ 1719ms | http |
| 120.92.212.16:7890 | ✓ 893ms | ✓ 1373ms | 否 | 否 | ✓ 1167ms | http |
| 160.238.65.3:3128 | 否 | ✓ 1697ms | ✓ 1469ms | ✓ 1927ms | ✓ 1710ms | http |
| 160.238.65.6:3128 | 否 | ✓ 1689ms | ✓ 1487ms | ✓ 1921ms | ✓ 1729ms | http |
| 160.238.65.8:3128 | 否 | ✓ 1675ms | ✓ 1481ms | ✓ 1892ms | ✓ 1758ms | http |
| 120.92.212.16:8890 | ✓ 949ms | ✓ 1566ms | 否 | 否 | ✓ 953ms | http |
| 35.206.88.200:8888 | 否 | 否 | ✓ 1275ms | ✓ 1315ms | ✓ 1294ms | http |
| 172.104.63.237:3128 | ✓ 1040ms | 否 | ✓ 1993ms | ✓ 1051ms | ✓ 1525ms | http |
| 91.193.240.157:9877 | ✓ 1403ms | 否 | ✓ 1581ms | 否 | ✓ 1599ms | http |
| 103.125.118.172:3125 | ✓ 1406ms | 否 | ✓ 1478ms | ✓ 1854ms | ✓ 1802ms | http |
| 192.166.82.55:1080 | ✓ 1330ms | ✓ 1842ms | ✓ 1502ms | ✓ 1487ms | 否 | http |
| 210.223.44.230:3128 | ✓ 811ms | ✓ 1009ms | ✓ 1119ms | ✓ 1495ms | ✓ 663ms | http |
| 74.48.78.224:2080 | ✓ 304ms | ✓ 820ms | ✓ 681ms | ✓ 1124ms | ✓ 873ms | http |
| 45.136.198.40:3128 | ✓ 994ms | 否 | ✓ 1073ms | 否 | ✓ 1865ms | http |
| 35.234.17.221:8080 | ✓ 1360ms | ✓ 1458ms | ✓ 1073ms | 否 | 否 | http |
| 18.201.114.187:48277 | ✓ 1833ms | 否 | ✓ 1579ms | 否 | ✓ 1819ms | http |
| 103.82.23.118:5216 | ✓ 1933ms | 否 | ✓ 1913ms | ✓ 1797ms | ✓ 1564ms | http |
| 162.43.36.42:8080 | ✓ 1536ms | 否 | ✓ 1464ms | ✓ 1873ms | ✓ 1212ms | http |
| 150.107.140.238:3128 | ✓ 946ms | 否 | 否 | ✓ 1654ms | ✓ 1902ms | http |
| 199.38.85.122:40001 | ✓ 584ms | ✓ 1581ms | ✓ 1487ms | ✓ 1655ms | ✓ 1359ms | http |
| 199.38.85.122:40004 | ✓ 691ms | 否 | ✓ 1602ms | 否 | ✓ 1292ms | http |
| 121.230.9.248:1080 | ✓ 1077ms | ✓ 1516ms | ✓ 1085ms | ✓ 1560ms | ✓ 1065ms | http |
| 106.14.205.114:483 | 否 | ✓ 1925ms | ✓ 1404ms | ✓ 1100ms | ✓ 818ms | http |
| 8.137.112.117:3128 | 否 | ✓ 1380ms | ✓ 1001ms | ✓ 1340ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1315ms | 否 | ✓ 1153ms | ✓ 1680ms | ✓ 1862ms | http |
| 77.83.203.6:443 | ✓ 960ms | ✓ 1735ms | ✓ 1211ms | 否 | ✓ 1650ms | http |
| 222.228.171.92:8080 | ✓ 472ms | 否 | ✓ 1493ms | ✓ 1928ms | ✓ 629ms | http |
| 91.233.223.147:3128 | ✓ 1066ms | 否 | ✓ 1090ms | 否 | ✓ 1722ms | http |
| 103.139.138.194:3128 | ✓ 1672ms | 否 | ✓ 1512ms | ✓ 1350ms | ✓ 1041ms | http |
| 116.80.63.64:7777 | ✓ 1467ms | 否 | ✓ 1455ms | ✓ 1791ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1666ms | 否 | 否 | ✓ 1505ms | ✓ 1416ms | http |
| 120.55.163.237:10086 | ✓ 776ms | ✓ 1027ms | ✓ 952ms | ✓ 1041ms | ✓ 830ms | http |
| 83.219.250.8:62920 | ✓ 1033ms | ✓ 1380ms | ✓ 1245ms | ✓ 1998ms | ✓ 1701ms | http |
| 103.218.241.170:8118 | ✓ 695ms | 否 | ✓ 1221ms | 否 | ✓ 999ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1333ms | 否 | ✓ 1510ms | ✓ 991ms | http |
| 103.67.46.225:3125 | ✓ 1849ms | 否 | ✓ 1734ms | ✓ 1588ms | ✓ 1447ms | http |
| 186.148.180.46:999 | ✓ 973ms | 否 | ✓ 1657ms | 否 | ✓ 1881ms | http |
| 103.215.36.88:17013 | ✓ 1557ms | ✓ 1656ms | ✓ 1393ms | ✓ 1409ms | ✓ 1461ms | http |
| 88.80.150.82:8080 | ✓ 1587ms | ✓ 1852ms | ✓ 1569ms | ✓ 1898ms | ✓ 1568ms | https |
| 103.35.188.243:3128 | 否 | ✓ 1261ms | 否 | ✓ 1344ms | ✓ 1061ms | http |
| 116.80.60.44:7777 | ✓ 1502ms | 否 | ✓ 1635ms | 否 | ✓ 1651ms | http |

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
