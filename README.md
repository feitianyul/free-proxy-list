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

最后更新：2026-04-23 20:06:33 UTC（2026-04-24 04:06:33 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 448ms | ✓ 900ms | ✓ 465ms | ✓ 869ms | ✓ 719ms | http |
| 46.101.95.183:8888 | ✓ 572ms | 否 | ✓ 1297ms | ✓ 1665ms | ✓ 1171ms | http |
| 1.231.81.166:3128 | ✓ 1414ms | ✓ 1246ms | ✓ 1199ms | 否 | ✓ 885ms | http |
| 113.160.132.26:8080 | ✓ 1907ms | ✓ 1452ms | ✓ 1072ms | ✓ 1391ms | ✓ 1250ms | http |
| 212.58.132.5:8888 | ✓ 1707ms | 否 | ✓ 1612ms | ✓ 1450ms | ✓ 1234ms | http |
| 35.225.22.61:80 | ✓ 317ms | ✓ 1394ms | 否 | ✓ 1124ms | 否 | http |
| 62.113.119.14:8080 | ✓ 548ms | ✓ 1664ms | ✓ 901ms | ✓ 1509ms | ✓ 1316ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1294ms | ✓ 1035ms | ✓ 1400ms | ✓ 1070ms | http |
| 115.231.181.40:8128 | ✓ 1162ms | ✓ 1232ms | 否 | ✓ 1451ms | 否 | http |
| 45.153.231.229:8080 | ✓ 1214ms | 否 | ✓ 1988ms | ✓ 1905ms | ✓ 1946ms | http |
| 120.92.108.86:7890 | ✓ 1775ms | 否 | ✓ 1911ms | ✓ 1989ms | ✓ 1935ms | http |
| 120.92.212.16:8890 | ✓ 1227ms | 否 | ✓ 1572ms | 否 | ✓ 1355ms | http |
| 20.127.128.70:8080 | 否 | 否 | ✓ 181ms | ✓ 861ms | ✓ 807ms | http |
| 210.45.76.58:42992 | ✓ 1169ms | ✓ 1473ms | ✓ 1455ms | ✓ 1574ms | ✓ 1245ms | http |
| 223.84.151.86:30005 | ✓ 1575ms | 否 | ✓ 1458ms | 否 | ✓ 1532ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1770ms | 否 | ✓ 1652ms | ✓ 1840ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1621ms | ✓ 1313ms | ✓ 1551ms | ✓ 1183ms | http |
| 45.76.207.177:40000 | 否 | 否 | ✓ 1103ms | ✓ 1046ms | ✓ 1081ms | http |
| 34.71.229.255:3128 | ✓ 467ms | ✓ 1390ms | ✓ 1047ms | ✓ 944ms | ✓ 791ms | http |
| 34.101.184.164:3128 | ✓ 1660ms | 否 | ✓ 1686ms | ✓ 1569ms | ✓ 1781ms | http |
| 38.180.2.107:3128 | ✓ 1441ms | ✓ 1682ms | ✓ 1583ms | 否 | 否 | http |
| 84.47.150.125:1080 | ✓ 967ms | 否 | 否 | ✓ 1927ms | ✓ 1435ms | http |
| 130.61.174.200:1080 | ✓ 447ms | 否 | ✓ 1372ms | ✓ 1851ms | 否 | http |
| 168.144.75.9:3128 | ✓ 1874ms | 否 | ✓ 1363ms | ✓ 1914ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1171ms | 否 | ✓ 1481ms | 否 | ✓ 1487ms | http |
| 91.99.15.45:2095 | ✓ 709ms | ✓ 1661ms | ✓ 1793ms | 否 | 否 | http |
| 8.209.238.110:47701 | ✓ 699ms | ✓ 1210ms | ✓ 803ms | ✓ 1167ms | ✓ 953ms | http |
| 94.131.118.129:1081 | ✓ 974ms | 否 | ✓ 1201ms | 否 | ✓ 1306ms | http |
| 45.81.130.20:8888 | ✓ 1652ms | ✓ 1322ms | ✓ 958ms | ✓ 1172ms | ✓ 895ms | http |
| 178.63.155.151:8888 | 否 | ✓ 1472ms | ✓ 1601ms | 否 | ✓ 1295ms | http |
| 121.230.9.54:1080 | ✓ 1289ms | ✓ 1537ms | ✓ 1205ms | ✓ 1515ms | ✓ 1250ms | http |
| 121.230.8.158:1080 | ✓ 1253ms | ✓ 1703ms | ✓ 1433ms | ✓ 1510ms | ✓ 1298ms | http |
| 121.230.9.203:1080 | ✓ 1187ms | ✓ 1762ms | ✓ 1444ms | ✓ 1645ms | ✓ 1426ms | http |
| 121.230.8.41:1080 | ✓ 1722ms | ✓ 1994ms | ✓ 1660ms | ✓ 1532ms | 否 | http |
| 162.240.154.26:3128 | ✓ 1475ms | ✓ 1716ms | ✓ 1951ms | ✓ 1836ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1811ms | ✓ 1868ms | ✓ 1749ms | 否 | 否 | http |
| 168.222.254.136:8888 | ✓ 1361ms | ✓ 1891ms | ✓ 1647ms | 否 | 否 | http |
| 177.93.132.244:3128 | ✓ 1135ms | 否 | ✓ 660ms | 否 | ✓ 1731ms | http |
| 57.128.188.167:8183 | ✓ 1959ms | 否 | ✓ 1418ms | 否 | ✓ 1915ms | http |
| 42.101.8.101:8888 | 否 | ✓ 1613ms | ✓ 1527ms | 否 | ✓ 1352ms | http |
| 152.32.132.190:7890 | ✓ 1781ms | ✓ 1714ms | ✓ 1836ms | 否 | 否 | http |
| 20.120.225.109:3128 | 否 | ✓ 1294ms | ✓ 1043ms | 否 | ✓ 1016ms | http |
| 45.140.147.82:1081 | ✓ 575ms | ✓ 1244ms | 否 | 否 | ✓ 873ms | http |
| 103.82.23.118:5261 | ✓ 1851ms | 否 | ✓ 1489ms | ✓ 1904ms | ✓ 1939ms | http |
| 38.79.118.202:33858 | ✓ 1306ms | 否 | ✓ 1260ms | 否 | ✓ 1446ms | http |
| 64.188.77.26:3128 | 否 | ✓ 1434ms | ✓ 775ms | 否 | ✓ 1353ms | http |
| 64.188.77.221:3128 | ✓ 1088ms | ✓ 1998ms | 否 | 否 | ✓ 1282ms | http |
| 47.84.73.61:1080 | ✓ 1202ms | ✓ 1850ms | ✓ 1112ms | ✓ 1286ms | ✓ 1023ms | http |
| 94.131.118.129:1082 | ✓ 1212ms | ✓ 1829ms | ✓ 1333ms | 否 | ✓ 1373ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1976ms | ✓ 1755ms | ✓ 1751ms | http |
| 20.164.75.153:8080 | ✓ 1277ms | 否 | ✓ 1903ms | 否 | ✓ 1835ms | http |
| 208.87.243.199:7878 | 否 | ✓ 1639ms | ✓ 1071ms | 否 | ✓ 1590ms | http |
| 85.190.99.143:443 | ✓ 1382ms | 否 | ✓ 1026ms | 否 | ✓ 1852ms | http |
| 37.187.109.70:10111 | ✓ 1822ms | ✓ 1446ms | 否 | 否 | ✓ 1528ms | http |
| 61.52.131.172:8443 | ✓ 1060ms | ✓ 1325ms | ✓ 1072ms | ✓ 1385ms | ✓ 1159ms | http |
| 168.110.52.228:3128 | ✓ 683ms | ✓ 1392ms | ✓ 1359ms | ✓ 1247ms | ✓ 1023ms | http |
| 223.16.170.103:80 | ✓ 1715ms | ✓ 1968ms | ✓ 1694ms | 否 | ✓ 1332ms | http |
| 103.20.184.66:1111 | ✓ 1664ms | 否 | ✓ 1483ms | ✓ 1562ms | ✓ 1553ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1424ms | ✓ 1813ms | ✓ 1444ms | ✓ 1654ms | http |

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
