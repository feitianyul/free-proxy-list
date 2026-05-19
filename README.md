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

最后更新：2026-05-19 13:33:04 UTC（2026-05-19 21:33:04 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 174.137.134.182:2999 | ✓ 451ms | ✓ 1585ms | ✓ 520ms | 否 | ✓ 1571ms | http |
| 43.130.126.146:6688 | ✓ 667ms | 否 | ✓ 819ms | ✓ 1211ms | 否 | http |
| 192.99.8.15:8850 | ✓ 444ms | 否 | ✓ 1391ms | ✓ 1457ms | ✓ 1447ms | http |
| 176.111.37.216:39811 | ✓ 1591ms | ✓ 1707ms | ✓ 1020ms | 否 | 否 | http |
| 185.200.188.234:10001 | ✓ 1628ms | 否 | ✓ 825ms | 否 | ✓ 1665ms | http |
| 113.160.132.26:8080 | ✓ 1740ms | 否 | ✓ 1751ms | ✓ 1393ms | ✓ 1199ms | http |
| 202.28.194.139:31280 | 否 | 否 | ✓ 1958ms | ✓ 1905ms | ✓ 1995ms | http |
| 161.117.86.53:8100 | ✓ 1097ms | ✓ 1460ms | ✓ 1011ms | ✓ 1397ms | ✓ 1239ms | http |
| 176.111.37.5:39811 | ✓ 830ms | 否 | ✓ 1428ms | 否 | ✓ 1432ms | http |
| 114.214.163.108:6789 | ✓ 1263ms | ✓ 1552ms | ✓ 1339ms | ✓ 1555ms | ✓ 1220ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1144ms | ✓ 1780ms | ✓ 1285ms | ✓ 1881ms | http |
| 89.58.50.94:11140 | ✓ 1304ms | 否 | ✓ 1592ms | 否 | ✓ 1401ms | http |
| 207.148.124.152:6868 | 否 | 否 | ✓ 1956ms | ✓ 1528ms | ✓ 1121ms | http |
| 47.52.134.234:36463 | ✓ 1834ms | ✓ 1663ms | ✓ 1460ms | ✓ 885ms | 否 | http |
| 138.2.78.251:8100 | ✓ 1492ms | 否 | 否 | ✓ 1461ms | ✓ 1458ms | http |
| 138.2.92.70:8100 | ✓ 1990ms | 否 | ✓ 1926ms | ✓ 1425ms | ✓ 1662ms | http |
| 45.95.203.47:6699 | 否 | 否 | ✓ 688ms | ✓ 1745ms | ✓ 1581ms | http |
| 43.154.90.238:9527 | ✓ 1885ms | 否 | ✓ 920ms | ✓ 1210ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1361ms | ✓ 1265ms | 否 | 否 | ✓ 1071ms | http |
| 5.252.33.13:2025 | ✓ 1409ms | 否 | ✓ 1400ms | 否 | ✓ 1710ms | http |
| 168.138.171.204:8100 | ✓ 1476ms | 否 | ✓ 1101ms | ✓ 1691ms | ✓ 1443ms | http |
| 121.130.199.80:24007 | 否 | ✓ 1493ms | ✓ 1117ms | ✓ 1404ms | ✓ 1023ms | http |
| 152.67.191.232:6800 | ✓ 1184ms | 否 | ✓ 1198ms | ✓ 1613ms | 否 | http |
| 148.230.4.241:999 | ✓ 1039ms | 否 | ✓ 656ms | ✓ 1678ms | ✓ 1320ms | http |
| 170.106.136.181:31002 | ✓ 683ms | ✓ 1569ms | ✓ 1699ms | ✓ 946ms | ✓ 570ms | http |
| 45.125.67.37:8443 | ✓ 881ms | 否 | ✓ 1555ms | ✓ 1834ms | ✓ 1385ms | http |
| 74.208.192.81:3129 | ✓ 872ms | 否 | ✓ 275ms | ✓ 1123ms | ✓ 832ms | http |
| 104.168.96.172:1888 | ✓ 393ms | ✓ 1020ms | ✓ 1204ms | ✓ 925ms | ✓ 717ms | http |
| 152.32.132.190:7890 | ✓ 1400ms | ✓ 1079ms | ✓ 1006ms | ✓ 900ms | 否 | http |
| 103.210.160.62:7789 | 否 | ✓ 1060ms | 否 | ✓ 1969ms | ✓ 1571ms | http |
| 38.224.150.84:999 | ✓ 1056ms | ✓ 1750ms | ✓ 1322ms | ✓ 1596ms | ✓ 1502ms | http |
| 38.224.150.86:999 | ✓ 1651ms | ✓ 1399ms | ✓ 1041ms | ✓ 1798ms | ✓ 1367ms | http |
| 38.224.150.83:999 | ✓ 819ms | ✓ 1670ms | ✓ 1594ms | ✓ 1964ms | ✓ 1241ms | http |
| 192.81.129.252:3136 | ✓ 1380ms | 否 | 否 | ✓ 1575ms | ✓ 1725ms | http |
| 152.42.177.32:8888 | ✓ 1035ms | 否 | ✓ 1029ms | ✓ 1359ms | 否 | http |
| 57.128.159.156:21 | ✓ 1389ms | 否 | ✓ 1164ms | 否 | ✓ 1820ms | http |
| 147.45.78.89:1080 | ✓ 1277ms | 否 | ✓ 1344ms | 否 | ✓ 1784ms | http |
| 185.103.252.20:3128 | ✓ 978ms | 否 | ✓ 1709ms | ✓ 1859ms | ✓ 1549ms | http |
| 3.101.133.120:80 | ✓ 395ms | 否 | ✓ 1078ms | ✓ 999ms | ✓ 1057ms | http |
| 202.40.186.222:27202 | ✓ 1629ms | 否 | ✓ 1807ms | 否 | ✓ 1742ms | http |
| 34.87.80.221:30000 | ✓ 824ms | ✓ 1367ms | ✓ 1586ms | ✓ 1262ms | ✓ 1218ms | http |
| 146.190.80.158:9090 | ✓ 1793ms | 否 | ✓ 1341ms | ✓ 1191ms | ✓ 958ms | http |
| 20.27.13.35:8561 | 否 | ✓ 1090ms | ✓ 662ms | ✓ 917ms | ✓ 869ms | http |
| 20.27.15.111:8561 | 否 | ✓ 1061ms | ✓ 654ms | ✓ 915ms | ✓ 872ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 571ms | ✓ 994ms | ✓ 725ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 627ms | ✓ 1011ms | ✓ 763ms | http |
| 20.78.118.91:8561 | 否 | ✓ 1018ms | ✓ 738ms | ✓ 926ms | ✓ 1516ms | http |
| 20.78.26.206:8561 | 否 | ✓ 1091ms | ✓ 624ms | ✓ 911ms | ✓ 1529ms | http |
| 91.107.123.158:1080 | ✓ 1320ms | 否 | ✓ 996ms | 否 | ✓ 1574ms | http |
| 8.219.97.248:80 | ✓ 1973ms | 否 | ✓ 1058ms | ✓ 1473ms | 否 | http |
| 147.45.41.112:1080 | ✓ 1083ms | 否 | ✓ 966ms | ✓ 1791ms | ✓ 1254ms | http |
| 3.15.187.17:1080 | ✓ 571ms | 否 | ✓ 1334ms | ✓ 1414ms | ✓ 1131ms | http |
| 210.223.44.230:3128 | ✓ 681ms | ✓ 1176ms | ✓ 1482ms | ✓ 1135ms | ✓ 1063ms | http |
| 85.192.29.60:3128 | ✓ 900ms | 否 | ✓ 1111ms | ✓ 1927ms | ✓ 1705ms | http |
| 34.101.184.164:3128 | ✓ 1543ms | 否 | ✓ 897ms | ✓ 1268ms | ✓ 1077ms | http |
| 147.139.141.104:8100 | ✓ 1546ms | 否 | ✓ 878ms | 否 | ✓ 1731ms | http |
| 8.217.78.60:8100 | 否 | ✓ 1974ms | 否 | ✓ 1395ms | ✓ 1801ms | http |
| 5.161.50.82:8118 | ✓ 1039ms | 否 | ✓ 1155ms | 否 | ✓ 1883ms | http |
| 168.110.52.228:3128 | ✓ 1962ms | 否 | 否 | ✓ 1182ms | ✓ 700ms | http |
| 45.80.231.251:3128 | ✓ 854ms | 否 | ✓ 813ms | 否 | ✓ 1698ms | http |
| 144.124.227.90:21074 | 否 | 否 | ✓ 1880ms | ✓ 1752ms | ✓ 1961ms | http |
| 20.210.39.153:8561 | ✓ 1460ms | ✓ 925ms | ✓ 670ms | ✓ 941ms | ✓ 767ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1444ms | ✓ 1468ms | ✓ 1409ms | ✓ 1508ms | http |
| 152.42.170.187:9090 | ✓ 1939ms | 否 | ✓ 967ms | ✓ 1232ms | 否 | http |
| 128.199.121.61:9090 | ✓ 824ms | 否 | 否 | ✓ 1632ms | ✓ 1125ms | http |
| 104.248.151.93:9090 | ✓ 832ms | 否 | ✓ 792ms | ✓ 1152ms | ✓ 928ms | http |
| 121.230.9.160:1080 | ✓ 1889ms | 否 | ✓ 1503ms | ✓ 1825ms | ✓ 1184ms | http |
| 154.27.196.243:999 | ✓ 1086ms | ✓ 1678ms | ✓ 1573ms | ✓ 1544ms | ✓ 1402ms | http |
| 154.27.196.241:999 | ✓ 1815ms | 否 | ✓ 1299ms | ✓ 1928ms | ✓ 1529ms | http |
| 61.52.131.172:8443 | ✓ 969ms | ✓ 1249ms | ✓ 1007ms | ✓ 1358ms | ✓ 1084ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 870ms | ✓ 1419ms | ✓ 944ms | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1443ms | ✓ 1750ms | ✓ 1469ms | http |
| 114.214.165.78:10810 | ✓ 1274ms | 否 | ✓ 1237ms | ✓ 1437ms | 否 | http |
| 223.16.170.103:80 | ✓ 1555ms | 否 | ✓ 1161ms | ✓ 1160ms | ✓ 1221ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1849ms | ✓ 1902ms | ✓ 1820ms | http |
| 212.58.132.5:8888 | ✓ 1869ms | 否 | ✓ 1254ms | ✓ 1566ms | 否 | http |

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
