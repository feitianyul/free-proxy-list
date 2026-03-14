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

最后更新：2026-03-14 05:39:44 UTC（2026-03-14 13:39:44 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.203.135:8443 | ✓ 1135ms | ✓ 647ms | ✓ 1062ms | ✓ 652ms | ✓ 710ms | http |
| 205.209.118.30:3138 | ✓ 426ms | ✓ 1235ms | ✓ 1331ms | ✓ 1387ms | ✓ 1100ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1900ms | ✓ 1402ms | ✓ 1225ms | 否 | http |
| 216.180.127.45:1080 | ✓ 1388ms | 否 | ✓ 1632ms | 否 | ✓ 1925ms | http |
| 45.88.0.111:3128 | ✓ 1126ms | 否 | ✓ 1790ms | 否 | ✓ 1798ms | http |
| 45.88.0.114:3128 | ✓ 1126ms | 否 | ✓ 1789ms | 否 | ✓ 1782ms | http |
| 45.88.0.117:3128 | ✓ 1126ms | 否 | ✓ 1791ms | 否 | ✓ 1767ms | http |
| 45.88.0.98:3128 | ✓ 1126ms | 否 | ✓ 1784ms | 否 | ✓ 1804ms | http |
| 45.88.0.116:3128 | ✓ 1134ms | 否 | 否 | ✓ 1905ms | ✓ 1670ms | http |
| 45.88.0.99:3128 | ✓ 1126ms | 否 | ✓ 1784ms | 否 | ✓ 1793ms | http |
| 45.88.0.113:3128 | ✓ 1133ms | 否 | ✓ 1782ms | 否 | ✓ 1783ms | http |
| 45.88.0.115:3128 | ✓ 1126ms | 否 | ✓ 1790ms | 否 | ✓ 1780ms | http |
| 45.167.124.52:8080 | ✓ 1572ms | 否 | ✓ 1440ms | ✓ 1808ms | ✓ 1599ms | http |
| 120.92.212.16:7890 | ✓ 915ms | ✓ 1183ms | ✓ 1226ms | ✓ 1184ms | 否 | http |
| 183.249.5.105:22222 | ✓ 872ms | ✓ 1230ms | ✓ 964ms | ✓ 1133ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1362ms | 否 | ✓ 1394ms | ✓ 965ms | http |
| 165.227.5.10:8888 | ✓ 1002ms | ✓ 1836ms | ✓ 358ms | 否 | 否 | http |
| 85.198.96.242:3128 | ✓ 1330ms | ✓ 1960ms | 否 | ✓ 1931ms | 否 | http |
| 116.80.65.79:3172 | 否 | 否 | ✓ 1519ms | ✓ 1931ms | ✓ 1612ms | http |
| 120.240.35.173:22222 | ✓ 942ms | ✓ 1268ms | ✓ 1136ms | ✓ 1396ms | ✓ 850ms | http |
| 117.159.239.49:22222 | ✓ 829ms | ✓ 1005ms | ✓ 801ms | ✓ 1044ms | 否 | http |
| 62.60.177.204:34094 | ✓ 380ms | 否 | ✓ 1059ms | ✓ 1109ms | ✓ 1040ms | http |
| 113.59.32.162:22222 | 否 | 否 | ✓ 973ms | ✓ 1165ms | ✓ 1064ms | http |
| 62.113.119.14:8080 | ✓ 1329ms | ✓ 1961ms | ✓ 1284ms | ✓ 1680ms | ✓ 1247ms | http |
| 101.43.255.96:80 | 否 | ✓ 1979ms | ✓ 1741ms | 否 | ✓ 978ms | http |
| 150.230.249.50:1080 | 否 | 否 | ✓ 1915ms | ✓ 1971ms | ✓ 1687ms | http |
| 213.220.62.62:3128 | ✓ 1008ms | 否 | ✓ 633ms | ✓ 1525ms | ✓ 1179ms | http |
| 222.184.48.235:22222 | 否 | ✓ 1226ms | ✓ 945ms | 否 | ✓ 912ms | http |
| 211.171.114.154:3128 | 否 | 否 | ✓ 1564ms | ✓ 1738ms | ✓ 1540ms | http |
| 171.251.173.39:5104 | 否 | 否 | ✓ 1771ms | ✓ 1422ms | ✓ 1270ms | http |
| 103.84.95.54:7890 | ✓ 1346ms | ✓ 1865ms | 否 | 否 | ✓ 614ms | http |
| 81.70.169.194:80 | ✓ 1582ms | ✓ 1250ms | 否 | 否 | ✓ 1040ms | http |
| 117.159.239.44:22222 | ✓ 805ms | 否 | ✓ 862ms | ✓ 1032ms | ✓ 815ms | http |
| 117.159.239.51:22222 | 否 | ✓ 992ms | ✓ 787ms | 否 | ✓ 901ms | http |
| 117.159.239.52:22222 | ✓ 782ms | ✓ 997ms | ✓ 919ms | ✓ 1281ms | ✓ 816ms | http |
| 113.59.32.141:22222 | 否 | ✓ 1558ms | ✓ 1299ms | ✓ 1247ms | ✓ 972ms | http |
| 160.250.5.22:1 | ✓ 1000ms | 否 | ✓ 1270ms | ✓ 1170ms | 否 | http |
| 160.250.4.245:1 | ✓ 975ms | 否 | ✓ 1145ms | ✓ 1197ms | ✓ 1332ms | http |
| 152.42.213.210:8080 | ✓ 1450ms | 否 | ✓ 1679ms | ✓ 1093ms | ✓ 870ms | http |
| 152.42.213.210:443 | ✓ 1374ms | 否 | ✓ 1953ms | ✓ 1112ms | ✓ 886ms | http |
| 35.225.22.61:80 | ✓ 1056ms | 否 | ✓ 1128ms | 否 | ✓ 965ms | http |
| 120.240.29.51:22222 | 否 | ✓ 1177ms | ✓ 991ms | ✓ 1130ms | ✓ 878ms | http |
| 120.238.159.228:22222 | ✓ 1697ms | ✓ 1164ms | ✓ 832ms | ✓ 1059ms | ✓ 900ms | http |
| 120.232.242.119:22222 | ✓ 856ms | ✓ 1174ms | ✓ 894ms | ✓ 1106ms | ✓ 850ms | http |
| 45.136.131.39:8443 | ✓ 624ms | ✓ 1318ms | ✓ 354ms | ✓ 641ms | ✓ 505ms | http |
| 222.184.48.251:22222 | 否 | ✓ 1542ms | 否 | ✓ 1275ms | ✓ 828ms | http |
| 120.238.159.230:22222 | 否 | ✓ 1166ms | ✓ 944ms | ✓ 1098ms | ✓ 947ms | http |
| 117.159.239.54:22222 | ✓ 877ms | ✓ 998ms | ✓ 785ms | ✓ 1062ms | ✓ 924ms | http |
| 120.240.35.161:22222 | ✓ 896ms | ✓ 1133ms | ✓ 1048ms | ✓ 1114ms | ✓ 853ms | http |
| 24.144.86.173:1080 | ✓ 532ms | ✓ 1953ms | ✓ 1667ms | ✓ 1448ms | ✓ 847ms | http |
| 210.223.44.230:3128 | ✓ 1808ms | 否 | ✓ 1551ms | ✓ 1860ms | 否 | http |
| 45.136.131.42:8447 | ✓ 73ms | ✓ 1075ms | ✓ 95ms | ✓ 869ms | ✓ 496ms | http |
| 45.136.130.159:8443 | ✓ 77ms | ✓ 644ms | ✓ 98ms | ✓ 641ms | ✓ 820ms | http |
| 45.136.130.165:8443 | ✓ 84ms | ✓ 635ms | ✓ 81ms | ✓ 641ms | ✓ 860ms | http |
| 45.136.130.162:8443 | ✓ 89ms | ✓ 1081ms | ✓ 76ms | ✓ 660ms | 否 | http |
| 45.136.130.155:8443 | ✓ 83ms | ✓ 1531ms | ✓ 285ms | ✓ 1981ms | 否 | http |
| 45.136.130.158:8443 | ✓ 80ms | ✓ 1677ms | ✓ 351ms | 否 | 否 | http |
| 120.238.159.229:22222 | ✓ 1050ms | ✓ 1180ms | ✓ 947ms | ✓ 1108ms | ✓ 839ms | http |
| 121.40.231.103:7890 | ✓ 1334ms | 否 | 否 | ✓ 1809ms | ✓ 1572ms | http |
| 183.249.5.109:22222 | 否 | ✓ 1256ms | ✓ 917ms | ✓ 1099ms | ✓ 835ms | http |
| 120.198.141.84:22222 | 否 | ✓ 1611ms | ✓ 1165ms | ✓ 1305ms | 否 | http |
| 45.136.198.40:3128 | 否 | ✓ 1768ms | ✓ 831ms | ✓ 1720ms | ✓ 1315ms | http |
| 45.136.130.157:8443 | ✓ 1278ms | ✓ 639ms | ✓ 81ms | ✓ 657ms | ✓ 471ms | http |
| 45.136.130.161:8443 | ✓ 1277ms | ✓ 1005ms | ✓ 78ms | ✓ 646ms | ✓ 495ms | http |
| 106.117.208.101:7890 | ✓ 970ms | ✓ 1202ms | ✓ 939ms | 否 | 否 | http |
| 27.254.99.183:8118 | ✓ 1591ms | 否 | ✓ 919ms | ✓ 1261ms | 否 | http |
| 86.53.183.16:1080 | ✓ 1158ms | 否 | ✓ 1505ms | 否 | ✓ 1559ms | http |
| 54.222.174.194:80 | 否 | ✓ 1669ms | ✓ 1716ms | ✓ 1765ms | 否 | http |
| 114.214.208.153:10808 | ✓ 1164ms | ✓ 1469ms | ✓ 1515ms | ✓ 1388ms | ✓ 1083ms | http |
| 183.249.5.117:22222 | ✓ 785ms | ✓ 1209ms | ✓ 811ms | ✓ 1381ms | ✓ 802ms | http |
| 120.240.35.177:22222 | ✓ 1092ms | ✓ 1099ms | ✓ 1013ms | ✓ 1199ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1725ms | ✓ 1705ms | ✓ 1269ms | http |
| 172.212.68.37:3128 | ✓ 815ms | 否 | ✓ 1753ms | ✓ 1478ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1904ms | 否 | ✓ 1925ms | ✓ 1678ms | 否 | http |
| 222.184.48.252:22222 | ✓ 954ms | ✓ 1066ms | 否 | 否 | ✓ 850ms | http |
| 113.59.32.142:22222 | ✓ 1071ms | ✓ 1315ms | ✓ 1025ms | 否 | ✓ 1000ms | http |
| 59.46.216.131:30001 | ✓ 1054ms | 否 | ✓ 1004ms | 否 | ✓ 1071ms | http |
| 85.208.108.43:2094 | ✓ 635ms | 否 | ✓ 1543ms | ✓ 1534ms | 否 | http |
| 45.136.130.163:8443 | 否 | ✓ 1078ms | ✓ 224ms | ✓ 658ms | ✓ 771ms | http |
| 121.126.185.63:25152 | 否 | 否 | ✓ 1836ms | ✓ 1971ms | ✓ 1665ms | http |
| 101.34.21.55:90 | ✓ 1775ms | 否 | ✓ 1144ms | 否 | ✓ 1822ms | http |
| 101.43.127.100:8877 | ✓ 1758ms | 否 | ✓ 1577ms | 否 | ✓ 1090ms | http |
| 103.67.46.225:3125 | ✓ 1719ms | 否 | ✓ 1580ms | ✓ 1563ms | ✓ 1634ms | http |
| 61.52.131.172:8443 | ✓ 910ms | ✓ 1133ms | ✓ 913ms | ✓ 1107ms | ✓ 850ms | http |
| 91.233.223.147:3128 | ✓ 1570ms | 否 | ✓ 930ms | 否 | ✓ 1713ms | http |
| 43.165.195.107:3128 | ✓ 1532ms | ✓ 1365ms | ✓ 901ms | ✓ 1124ms | 否 | http |
| 183.249.5.110:22222 | ✓ 741ms | ✓ 988ms | ✓ 871ms | ✓ 955ms | 否 | http |
| 222.184.48.248:22222 | ✓ 1163ms | 否 | 否 | ✓ 1211ms | ✓ 1047ms | http |
| 183.249.5.214:22222 | ✓ 883ms | ✓ 1140ms | ✓ 816ms | ✓ 1351ms | 否 | http |

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
