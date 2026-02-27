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

最后更新：2026-02-27 10:46:17 UTC（2026-02-27 18:46:17 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1881ms | 否 | ✓ 913ms | ✓ 1337ms | ✓ 869ms | http |
| 72.56.59.62:63133 | ✓ 1500ms | 否 | ✓ 1810ms | 否 | ✓ 1900ms | http |
| 72.56.59.56:63127 | ✓ 1598ms | 否 | ✓ 1677ms | 否 | ✓ 1873ms | http |
| 45.88.0.98:3128 | ✓ 722ms | 否 | ✓ 1934ms | 否 | ✓ 1576ms | http |
| 85.208.108.43:10808 | ✓ 432ms | 否 | ✓ 843ms | ✓ 1351ms | ✓ 1428ms | http |
| 113.59.32.162:22222 | ✓ 1079ms | ✓ 1412ms | 否 | ✓ 1288ms | ✓ 963ms | http |
| 120.232.242.119:22222 | 否 | ✓ 1320ms | ✓ 1092ms | 否 | ✓ 1062ms | http |
| 83.219.250.8:62920 | ✓ 707ms | 否 | ✓ 1321ms | ✓ 1908ms | ✓ 1618ms | http |
| 147.45.216.148:1080 | ✓ 483ms | 否 | ✓ 1529ms | ✓ 1775ms | ✓ 1568ms | http |
| 72.56.59.23:61937 | ✓ 1503ms | 否 | ✓ 1611ms | 否 | ✓ 1871ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1280ms | ✓ 1423ms | ✓ 1110ms | http |
| 52.188.28.218:3128 | 否 | 否 | ✓ 1156ms | ✓ 948ms | ✓ 1769ms | http |
| 104.238.30.38:59741 | ✓ 1728ms | 否 | ✓ 1871ms | 否 | ✓ 1967ms | http |
| 104.238.30.37:59741 | ✓ 1732ms | 否 | ✓ 1867ms | 否 | ✓ 1966ms | http |
| 103.84.95.54:7890 | ✓ 861ms | 否 | ✓ 1109ms | 否 | ✓ 879ms | http |
| 223.113.134.102:22222 | ✓ 857ms | ✓ 1988ms | ✓ 896ms | ✓ 1292ms | ✓ 790ms | http |
| 104.238.30.86:63900 | ✓ 1753ms | 否 | ✓ 1871ms | 否 | ✓ 1967ms | http |
| 104.238.30.58:63744 | ✓ 1697ms | 否 | ✓ 1807ms | 否 | ✓ 1967ms | http |
| 120.240.35.173:22222 | ✓ 1104ms | 否 | ✓ 1219ms | ✓ 1436ms | ✓ 1090ms | http |
| 104.238.30.91:63900 | ✓ 1703ms | 否 | ✓ 1744ms | 否 | ✓ 1967ms | http |
| 59.46.216.131:30001 | ✓ 1057ms | ✓ 1623ms | ✓ 1319ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1631ms | ✓ 1540ms | ✓ 1175ms | ✓ 1695ms | ✓ 1034ms | http |
| 34.142.0.1:10808 | ✓ 985ms | ✓ 1550ms | ✓ 1721ms | 否 | ✓ 1468ms | http |
| 72.56.59.17:61931 | ✓ 1586ms | 否 | ✓ 1903ms | 否 | ✓ 1909ms | http |
| 101.47.73.135:3128 | ✓ 1744ms | 否 | ✓ 1179ms | ✓ 1537ms | ✓ 1334ms | http |
| 72.56.50.17:59787 | ✓ 1552ms | 否 | ✓ 1867ms | 否 | ✓ 1901ms | http |
| 223.113.134.103:22222 | ✓ 799ms | ✓ 1006ms | ✓ 826ms | ✓ 1057ms | ✓ 815ms | http |
| 168.235.110.63:3128 | ✓ 1028ms | ✓ 1936ms | ✓ 797ms | ✓ 1295ms | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1366ms | ✓ 1574ms | ✓ 1408ms | ✓ 1064ms | http |
| 61.72.110.24:3128 | ✓ 1904ms | ✓ 1333ms | 否 | ✓ 1573ms | 否 | http |
| 104.238.30.45:59741 | ✓ 1757ms | 否 | ✓ 1807ms | 否 | ✓ 1999ms | http |
| 35.225.22.61:80 | ✓ 504ms | ✓ 1217ms | ✓ 261ms | ✓ 965ms | ✓ 720ms | http |
| 45.125.67.37:8443 | ✓ 1249ms | 否 | ✓ 1334ms | ✓ 1298ms | 否 | http |
| 223.113.134.141:22222 | ✓ 888ms | ✓ 1254ms | ✓ 924ms | ✓ 1052ms | ✓ 787ms | http |
| 120.238.159.229:22222 | 否 | ✓ 1494ms | 否 | ✓ 1305ms | ✓ 1025ms | http |
| 183.249.5.109:22222 | ✓ 829ms | ✓ 1197ms | ✓ 1091ms | ✓ 1662ms | ✓ 1109ms | http |
| 195.123.209.48:3128 | ✓ 1006ms | 否 | ✓ 1554ms | 否 | ✓ 1748ms | http |
| 185.246.90.163:10808 | ✓ 1683ms | 否 | ✓ 990ms | 否 | ✓ 1680ms | http |
| 138.124.53.25:7443 | ✓ 1012ms | 否 | ✓ 1634ms | ✓ 1942ms | ✓ 1768ms | http |
| 104.238.30.39:59741 | ✓ 1750ms | 否 | ✓ 1771ms | 否 | ✓ 1999ms | http |
| 120.92.212.16:7890 | ✓ 1107ms | ✓ 1396ms | 否 | ✓ 1642ms | ✓ 1053ms | http |
| 43.252.106.26:1111 | ✓ 1585ms | 否 | 否 | ✓ 1641ms | ✓ 1576ms | http |
| 85.208.108.43:2094 | ✓ 452ms | 否 | ✓ 161ms | ✓ 1305ms | ✓ 828ms | http |
| 223.113.134.105:22222 | 否 | ✓ 1048ms | ✓ 832ms | ✓ 1056ms | ✓ 832ms | http |
| 183.249.5.111:22222 | ✓ 845ms | ✓ 1050ms | ✓ 844ms | 否 | ✓ 854ms | http |
| 183.249.5.117:22222 | 否 | ✓ 1019ms | ✓ 1072ms | ✓ 1116ms | 否 | http |
| 223.113.134.92:22222 | ✓ 907ms | ✓ 1245ms | ✓ 780ms | ✓ 1084ms | ✓ 810ms | http |
| 120.198.141.75:22222 | 否 | 否 | ✓ 1130ms | ✓ 1351ms | ✓ 1001ms | http |
| 113.59.32.161:22222 | ✓ 1109ms | ✓ 1395ms | ✓ 976ms | ✓ 1270ms | ✓ 936ms | http |
| 120.240.35.160:22222 | ✓ 1153ms | ✓ 1611ms | ✓ 1133ms | 否 | ✓ 1056ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 920ms | ✓ 1107ms | ✓ 907ms | http |
| 113.59.32.160:22222 | ✓ 1126ms | ✓ 1320ms | ✓ 971ms | 否 | ✓ 1557ms | http |
| 120.240.35.175:22222 | ✓ 1120ms | 否 | ✓ 1524ms | ✓ 1431ms | ✓ 1267ms | http |
| 35.234.17.221:8080 | ✓ 960ms | 否 | 否 | ✓ 1285ms | ✓ 1015ms | http |
| 45.140.147.82:1081 | ✓ 407ms | 否 | ✓ 981ms | ✓ 1580ms | ✓ 1295ms | http |
| 120.232.242.120:22222 | ✓ 1000ms | ✓ 1330ms | ✓ 997ms | 否 | ✓ 988ms | http |
| 185.243.218.43:49153 | ✓ 630ms | 否 | ✓ 1193ms | 否 | ✓ 1860ms | http |
| 61.72.110.94:3128 | ✓ 1773ms | 否 | 否 | ✓ 1188ms | ✓ 1690ms | http |
| 103.113.70.189:1081 | 否 | ✓ 868ms | 否 | ✓ 1262ms | ✓ 817ms | http |
| 103.117.100.127:13082 | ✓ 977ms | 否 | ✓ 805ms | ✓ 923ms | ✓ 770ms | http |
| 45.136.198.40:3128 | ✓ 1025ms | 否 | ✓ 1286ms | 否 | ✓ 1776ms | http |
| 104.238.30.50:59741 | ✓ 1751ms | 否 | ✓ 1839ms | 否 | ✓ 1999ms | http |
| 211.171.114.154:3128 | ✓ 1670ms | 否 | ✓ 1941ms | ✓ 1371ms | ✓ 1152ms | http |
| 172.212.68.37:3128 | ✓ 219ms | 否 | ✓ 946ms | 否 | ✓ 1019ms | http |
| 113.59.32.142:22222 | ✓ 1081ms | ✓ 1433ms | ✓ 1154ms | ✓ 1322ms | ✓ 998ms | http |
| 8.219.97.248:80 | ✓ 1967ms | 否 | ✓ 1515ms | ✓ 1554ms | 否 | http |
| 45.140.147.155:1082 | ✓ 532ms | 否 | ✓ 1069ms | 否 | ✓ 1450ms | http |
| 147.45.251.242:8888 | ✓ 1212ms | 否 | ✓ 1677ms | 否 | ✓ 1912ms | http |
| 152.32.255.24:27197 | ✓ 1830ms | 否 | ✓ 1487ms | ✓ 1378ms | ✓ 1080ms | http |
| 120.240.29.53:22222 | 否 | ✓ 1350ms | 否 | ✓ 1329ms | ✓ 1082ms | http |
| 34.101.184.164:3128 | ✓ 1791ms | 否 | 否 | ✓ 1553ms | ✓ 1230ms | http |
| 103.215.36.88:19239 | ✓ 1332ms | ✓ 1939ms | ✓ 1333ms | ✓ 1570ms | 否 | http |
| 104.238.30.40:59741 | ✓ 1713ms | 否 | ✓ 1835ms | 否 | ✓ 1999ms | http |
| 45.140.147.155:1081 | ✓ 443ms | ✓ 1304ms | 否 | 否 | ✓ 1073ms | http |
| 120.240.35.178:22222 | 否 | ✓ 1387ms | 否 | ✓ 1244ms | ✓ 986ms | http |
| 81.177.48.54:2080 | ✓ 1674ms | 否 | ✓ 792ms | 否 | ✓ 1956ms | http |
| 113.59.32.148:22222 | ✓ 1128ms | ✓ 1341ms | 否 | 否 | ✓ 1133ms | http |
| 104.37.184.214:1080 | 否 | 否 | ✓ 1607ms | ✓ 1659ms | ✓ 1779ms | http |
| 37.27.100.112:443 | ✓ 1996ms | 否 | ✓ 639ms | ✓ 1550ms | 否 | http |
| 121.237.181.137:8888 | ✓ 977ms | ✓ 1204ms | ✓ 1583ms | 否 | 否 | http |
| 103.215.36.88:13763 | ✓ 1050ms | ✓ 1552ms | ✓ 1271ms | ✓ 1555ms | ✓ 1180ms | http |
| 121.128.121.244:3128 | ✓ 1960ms | 否 | ✓ 1993ms | 否 | ✓ 921ms | http |

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
