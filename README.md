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

最后更新：2026-05-04 00:41:05 UTC（2026-05-04 08:41:05 UTC+8）

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
| 47.85.51.197:1080 | ✓ 666ms | ✓ 1315ms | ✓ 1120ms | ✓ 1171ms | ✓ 847ms | http |
| 113.160.132.26:8080 | ✓ 1571ms | ✓ 1658ms | ✓ 1994ms | 否 | ✓ 1118ms | http |
| 45.167.124.71:999 | ✓ 940ms | 否 | ✓ 1623ms | ✓ 1696ms | ✓ 1446ms | http |
| 47.77.216.82:1080 | ✓ 1193ms | 否 | ✓ 796ms | ✓ 883ms | ✓ 1267ms | http |
| 168.110.52.228:3128 | 否 | ✓ 1221ms | ✓ 1808ms | ✓ 890ms | ✓ 806ms | http |
| 46.105.190.40:3128 | ✓ 603ms | ✓ 1477ms | ✓ 1148ms | ✓ 1958ms | ✓ 1748ms | http |
| 45.153.231.229:8080 | ✓ 940ms | 否 | ✓ 1670ms | 否 | ✓ 1901ms | http |
| 80.92.204.47:1081 | ✓ 1353ms | ✓ 1830ms | ✓ 1975ms | 否 | ✓ 1978ms | http |
| 217.76.245.80:999 | ✓ 1052ms | 否 | ✓ 1331ms | ✓ 1848ms | ✓ 1532ms | http |
| 193.123.250.39:1080 | ✓ 795ms | 否 | ✓ 1890ms | ✓ 1337ms | 否 | http |
| 203.18.158.226:3128 | ✓ 1750ms | 否 | 否 | ✓ 1868ms | ✓ 1516ms | http |
| 190.12.150.244:999 | ✓ 1073ms | 否 | ✓ 1019ms | ✓ 1987ms | ✓ 1471ms | http |
| 59.46.216.131:30001 | ✓ 1202ms | ✓ 1390ms | 否 | ✓ 1586ms | 否 | http |
| 45.59.122.132:80 | ✓ 1497ms | 否 | ✓ 1820ms | ✓ 1608ms | ✓ 1832ms | http |
| 1.231.81.166:3128 | ✓ 1632ms | ✓ 936ms | ✓ 759ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 787ms | 否 | ✓ 1197ms | ✓ 1257ms | ✓ 865ms | http |
| 103.35.190.69:1082 | ✓ 295ms | ✓ 1460ms | ✓ 301ms | ✓ 1686ms | ✓ 1845ms | http |
| 45.140.147.155:1082 | ✓ 591ms | ✓ 1468ms | ✓ 1929ms | 否 | ✓ 1581ms | http |
| 62.60.231.71:56608 | ✓ 864ms | 否 | ✓ 1608ms | 否 | ✓ 1448ms | http |
| 45.129.141.143:3128 | ✓ 1655ms | 否 | ✓ 1971ms | 否 | ✓ 1894ms | http |
| 130.61.174.200:1080 | ✓ 1792ms | ✓ 1302ms | 否 | 否 | ✓ 1383ms | http |
| 122.2.48.121:8080 | ✓ 1306ms | 否 | ✓ 1275ms | ✓ 1300ms | ✓ 1312ms | http |
| 38.188.247.12:999 | 否 | ✓ 1400ms | ✓ 517ms | ✓ 1668ms | 否 | http |
| 86.104.72.219:1081 | 否 | ✓ 1336ms | ✓ 925ms | ✓ 1155ms | ✓ 864ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 916ms | ✓ 1770ms | ✓ 1214ms | http |
| 103.157.200.126:3128 | ✓ 1696ms | 否 | ✓ 1440ms | ✓ 1982ms | 否 | http |
| 173.212.246.157:3128 | ✓ 1821ms | ✓ 1628ms | ✓ 1724ms | 否 | ✓ 1587ms | http |
| 121.230.9.148:1080 | 否 | 否 | ✓ 1082ms | ✓ 1763ms | ✓ 1211ms | http |
| 86.104.72.220:1082 | ✓ 597ms | ✓ 1079ms | ✓ 1768ms | ✓ 1227ms | 否 | http |
| 86.104.72.220:1081 | ✓ 356ms | ✓ 1041ms | ✓ 1918ms | ✓ 1298ms | ✓ 933ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1609ms | ✓ 1928ms | ✓ 748ms | http |
| 154.64.232.35:8080 | ✓ 1598ms | 否 | ✓ 932ms | 否 | ✓ 1493ms | http |
| 139.162.153.201:3128 | 否 | 否 | ✓ 1536ms | ✓ 1426ms | ✓ 1295ms | http |
| 94.131.118.129:1081 | ✓ 697ms | ✓ 1557ms | ✓ 521ms | 否 | ✓ 1314ms | http |
| 20.205.16.149:3128 | 否 | ✓ 1374ms | ✓ 1009ms | ✓ 1058ms | ✓ 847ms | http |
| 8.154.21.175:3128 | ✓ 1918ms | ✓ 1131ms | ✓ 988ms | 否 | ✓ 982ms | http |
| 107.174.80.186:3128 | 否 | ✓ 857ms | ✓ 860ms | 否 | ✓ 680ms | http |
| 154.27.201.153:999 | 否 | 否 | ✓ 1299ms | ✓ 1659ms | ✓ 1550ms | http |
| 103.22.99.137:3125 | ✓ 1410ms | 否 | ✓ 1922ms | ✓ 1871ms | 否 | http |
| 101.32.244.83:8080 | ✓ 988ms | ✓ 1792ms | ✓ 989ms | ✓ 1320ms | ✓ 1319ms | http |
| 121.43.196.210:8222 | ✓ 978ms | ✓ 1084ms | ✓ 859ms | ✓ 1162ms | ✓ 992ms | http |
| 121.43.196.213:8222 | ✓ 1041ms | ✓ 1095ms | ✓ 842ms | ✓ 1154ms | ✓ 963ms | http |
| 194.150.220.163:1082 | ✓ 1276ms | ✓ 1631ms | ✓ 1388ms | 否 | 否 | http |
| 154.90.48.209:9090 | ✓ 1083ms | 否 | 否 | ✓ 1703ms | ✓ 1904ms | http |
| 45.125.67.37:8443 | ✓ 926ms | 否 | ✓ 902ms | ✓ 1286ms | ✓ 1358ms | http |
| 45.140.147.82:1081 | ✓ 1524ms | 否 | ✓ 879ms | 否 | ✓ 1515ms | http |
| 91.217.81.131:1080 | ✓ 1914ms | 否 | ✓ 1680ms | 否 | ✓ 1783ms | http |
| 103.3.246.71:3128 | ✓ 1500ms | 否 | ✓ 1037ms | ✓ 1369ms | ✓ 1195ms | http |
| 86.104.72.219:1082 | ✓ 834ms | ✓ 1237ms | ✓ 526ms | ✓ 1267ms | ✓ 1909ms | http |
| 121.230.8.136:1080 | ✓ 1126ms | ✓ 1427ms | ✓ 1074ms | ✓ 1764ms | ✓ 1320ms | http |
| 3.101.133.120:80 | 否 | ✓ 1282ms | ✓ 1182ms | ✓ 1306ms | ✓ 901ms | http |
| 45.88.0.115:3128 | ✓ 1165ms | ✓ 1514ms | 否 | 否 | ✓ 1330ms | http |
| 91.233.223.147:3128 | ✓ 1227ms | 否 | ✓ 1332ms | 否 | ✓ 1869ms | http |
| 45.88.0.98:3128 | ✓ 571ms | ✓ 1333ms | ✓ 623ms | ✓ 1380ms | ✓ 1037ms | http |
| 45.88.0.113:3128 | ✓ 579ms | ✓ 1393ms | ✓ 592ms | ✓ 1356ms | ✓ 1070ms | http |
| 45.88.0.99:3128 | ✓ 594ms | ✓ 1470ms | ✓ 602ms | ✓ 1361ms | ✓ 1032ms | http |
| 213.220.62.63:3128 | ✓ 574ms | ✓ 1374ms | ✓ 597ms | ✓ 1368ms | ✓ 1034ms | http |
| 45.88.0.116:3128 | ✓ 576ms | ✓ 1549ms | ✓ 729ms | ✓ 1365ms | ✓ 1044ms | http |
| 45.88.0.111:3128 | ✓ 574ms | ✓ 1584ms | ✓ 710ms | ✓ 1346ms | ✓ 1048ms | http |
| 45.88.0.117:3128 | ✓ 574ms | ✓ 1275ms | ✓ 572ms | 否 | ✓ 1032ms | http |
| 172.105.118.164:3128 | 否 | 否 | ✓ 1790ms | ✓ 1709ms | ✓ 1220ms | http |
| 104.128.138.186:1080 | ✓ 1686ms | 否 | ✓ 1939ms | 否 | ✓ 1824ms | http |
| 45.140.147.82:1082 | ✓ 716ms | 否 | ✓ 1604ms | ✓ 1714ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1420ms | 否 | ✓ 1659ms | ✓ 1553ms | ✓ 1871ms | http |
| 148.230.4.241:999 | ✓ 1229ms | ✓ 1697ms | ✓ 754ms | 否 | 否 | http |
| 20.127.128.70:8080 | ✓ 1168ms | ✓ 1761ms | ✓ 1094ms | ✓ 1393ms | ✓ 1171ms | http |
| 217.182.195.221:30003 | 否 | 否 | ✓ 1108ms | ✓ 1895ms | ✓ 1947ms | http |
| 94.131.118.39:1081 | ✓ 1995ms | ✓ 1767ms | ✓ 819ms | 否 | 否 | http |
| 150.249.255.91:3128 | ✓ 1525ms | ✓ 1468ms | 否 | 否 | ✓ 992ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1817ms | ✓ 1694ms | ✓ 1412ms | http |
| 121.230.8.91:1080 | ✓ 1164ms | ✓ 1275ms | ✓ 1015ms | ✓ 1255ms | ✓ 1031ms | http |
| 121.230.9.198:1080 | ✓ 1729ms | ✓ 1303ms | ✓ 1094ms | ✓ 1840ms | ✓ 1282ms | http |
| 61.52.131.172:8443 | ✓ 904ms | 否 | ✓ 952ms | ✓ 1222ms | ✓ 1902ms | http |
| 94.131.118.129:1082 | ✓ 977ms | ✓ 1486ms | ✓ 1477ms | ✓ 1961ms | ✓ 1532ms | http |
| 37.187.109.70:10111 | ✓ 1386ms | 否 | ✓ 1720ms | 否 | ✓ 1944ms | http |
| 101.32.243.189:80 | ✓ 1186ms | 否 | 否 | ✓ 1259ms | ✓ 1234ms | http |

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
