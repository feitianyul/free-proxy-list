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

最后更新：2026-03-01 18:26:24 UTC（2026-03-02 02:26:24 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 148.135.85.87:1080 | ✓ 1112ms | 否 | ✓ 1498ms | ✓ 1198ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1248ms | ✓ 1027ms | ✓ 1696ms | 否 | ✓ 1369ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1226ms | 否 | ✓ 1255ms | ✓ 1040ms | http |
| 61.72.110.54:3128 | ✓ 1414ms | 否 | ✓ 1300ms | 否 | ✓ 1212ms | http |
| 35.234.17.221:8080 | ✓ 1235ms | 否 | ✓ 1209ms | 否 | ✓ 993ms | http |
| 59.46.216.131:30001 | ✓ 1083ms | ✓ 1456ms | ✓ 1089ms | 否 | 否 | http |
| 205.209.118.30:3138 | ✓ 336ms | ✓ 983ms | ✓ 194ms | 否 | 否 | http |
| 141.11.210.35:1080 | ✓ 707ms | 否 | ✓ 1197ms | ✓ 870ms | ✓ 719ms | http |
| 81.70.169.194:80 | ✓ 956ms | ✓ 1438ms | ✓ 997ms | ✓ 1342ms | ✓ 1708ms | http |
| 120.92.212.16:8890 | ✓ 1042ms | ✓ 1332ms | 否 | ✓ 1300ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1315ms | ✓ 1304ms | 否 | 否 | ✓ 1277ms | http |
| 101.43.255.96:80 | ✓ 982ms | ✓ 1285ms | ✓ 1044ms | ✓ 1541ms | ✓ 1662ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1827ms | ✓ 1069ms | ✓ 771ms | http |
| 16.78.119.130:443 | ✓ 1816ms | 否 | ✓ 1584ms | 否 | ✓ 1690ms | http |
| 35.225.22.61:80 | 否 | ✓ 1969ms | ✓ 1118ms | ✓ 1009ms | ✓ 974ms | http |
| 45.140.147.155:1081 | 否 | 否 | ✓ 1415ms | ✓ 1777ms | ✓ 1236ms | http |
| 95.85.252.153:21064 | ✓ 536ms | ✓ 1549ms | ✓ 1471ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 1710ms | 否 | 否 | ✓ 1610ms | ✓ 1136ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 432ms | ✓ 874ms | ✓ 1697ms | http |
| 91.238.104.171:2023 | ✓ 1130ms | 否 | ✓ 666ms | ✓ 1971ms | 否 | http |
| 111.79.111.126:3128 | ✓ 1210ms | ✓ 1404ms | ✓ 1312ms | ✓ 1514ms | ✓ 1284ms | http |
| 103.104.99.89:80 | ✓ 1601ms | 否 | ✓ 1398ms | ✓ 1571ms | ✓ 1545ms | http |
| 103.104.99.29:80 | ✓ 1457ms | 否 | ✓ 1417ms | ✓ 1589ms | ✓ 1558ms | http |
| 14.56.177.44:3128 | ✓ 683ms | ✓ 1212ms | ✓ 657ms | ✓ 1123ms | ✓ 902ms | http |
| 74.208.234.198:443 | ✓ 1053ms | ✓ 1267ms | ✓ 943ms | 否 | ✓ 1116ms | http |
| 101.47.73.135:3128 | ✓ 1634ms | 否 | ✓ 1611ms | 否 | ✓ 1939ms | http |
| 5.75.201.136:1080 | ✓ 516ms | ✓ 1483ms | 否 | ✓ 1880ms | 否 | http |
| 103.113.70.189:1081 | 否 | ✓ 1075ms | 否 | ✓ 1423ms | ✓ 1731ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1548ms | ✓ 1380ms | ✓ 1120ms | http |
| 107.174.133.10:3128 | ✓ 719ms | 否 | ✓ 971ms | 否 | ✓ 1914ms | http |
| 121.128.121.54:3128 | ✓ 1144ms | 否 | 否 | ✓ 1426ms | ✓ 1625ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1812ms | 否 | ✓ 1390ms | ✓ 1141ms | http |
| 34.101.184.164:3128 | ✓ 1930ms | 否 | ✓ 1401ms | 否 | ✓ 1147ms | http |
| 103.191.169.218:1111 | 否 | 否 | ✓ 1593ms | ✓ 1861ms | ✓ 1803ms | http |
| 37.27.100.107:443 | ✓ 804ms | ✓ 1926ms | ✓ 1462ms | 否 | 否 | http |
| 167.160.184.231:6005 | ✓ 849ms | ✓ 1260ms | ✓ 1455ms | ✓ 1401ms | ✓ 1262ms | http |
| 142.171.85.32:1080 | 否 | ✓ 1226ms | ✓ 1021ms | ✓ 999ms | ✓ 1940ms | http |
| 103.84.95.54:7890 | ✓ 988ms | 否 | ✓ 761ms | ✓ 1888ms | ✓ 942ms | http |
| 121.128.121.184:3128 | 否 | ✓ 1599ms | ✓ 1592ms | ✓ 1265ms | ✓ 869ms | http |
| 183.128.208.68:7890 | ✓ 980ms | ✓ 1254ms | ✓ 1702ms | 否 | ✓ 1108ms | http |
| 116.80.48.217:7777 | ✓ 1841ms | 否 | 否 | ✓ 1916ms | ✓ 1735ms | http |
| 36.147.78.166:80 | ✓ 1930ms | ✓ 1745ms | ✓ 1760ms | 否 | 否 | http |
| 36.147.78.166:443 | ✓ 1748ms | 否 | 否 | ✓ 1971ms | ✓ 1752ms | http |
| 61.72.110.94:3128 | ✓ 1382ms | 否 | ✓ 1657ms | 否 | ✓ 1750ms | http |
| 207.254.71.62:8088 | ✓ 898ms | ✓ 1609ms | ✓ 1420ms | 否 | ✓ 1761ms | http |
| 129.213.139.179:8080 | ✓ 537ms | 否 | ✓ 571ms | 否 | ✓ 1519ms | http |
| 103.82.23.118:5178 | ✓ 1922ms | ✓ 1969ms | ✓ 1287ms | 否 | ✓ 1567ms | http |
| 157.245.194.13:8888 | ✓ 793ms | 否 | ✓ 1365ms | ✓ 1155ms | ✓ 957ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1933ms | 否 | ✓ 1057ms | ✓ 1832ms | http |
| 14.103.233.245:3128 | 否 | ✓ 1078ms | ✓ 1123ms | 否 | ✓ 1969ms | http |
| 180.127.149.247:1080 | ✓ 1011ms | ✓ 1319ms | ✓ 1037ms | ✓ 1259ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1425ms | 否 | ✓ 1232ms | ✓ 1649ms | ✓ 1160ms | http |
| 207.244.244.178:3128 | ✓ 1554ms | ✓ 1504ms | ✓ 1152ms | 否 | 否 | http |
| 8.217.147.173:8080 | ✓ 1098ms | ✓ 1713ms | ✓ 1104ms | ✓ 1232ms | ✓ 1133ms | http |
| 103.39.51.190:8080 | ✓ 1958ms | 否 | 否 | ✓ 1753ms | ✓ 1957ms | http |
| 47.95.231.180:8084 | ✓ 965ms | ✓ 1326ms | ✓ 970ms | ✓ 1268ms | ✓ 992ms | http |
| 77.83.203.5:443 | ✓ 922ms | 否 | ✓ 1063ms | 否 | ✓ 1669ms | http |
| 45.136.198.40:3128 | ✓ 1070ms | 否 | ✓ 1382ms | 否 | ✓ 1995ms | http |
| 92.118.228.31:3128 | 否 | ✓ 1644ms | ✓ 1731ms | 否 | ✓ 1365ms | http |
| 209.126.10.139:3128 | ✓ 1035ms | ✓ 1054ms | ✓ 1042ms | 否 | ✓ 1437ms | http |
| 194.59.204.87:9080 | ✓ 533ms | ✓ 1434ms | ✓ 646ms | 否 | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1056ms | ✓ 781ms | 否 | ✓ 810ms | http |
| 37.27.100.112:443 | ✓ 1036ms | 否 | ✓ 1132ms | 否 | ✓ 1783ms | http |
| 211.171.114.154:3128 | ✓ 1324ms | ✓ 1893ms | ✓ 1948ms | 否 | 否 | http |
| 24.199.124.152:3128 | ✓ 1083ms | ✓ 1201ms | ✓ 1231ms | ✓ 840ms | ✓ 664ms | http |
| 103.153.202.120:8080 | ✓ 1816ms | 否 | ✓ 1683ms | ✓ 1778ms | ✓ 1805ms | http |
| 177.243.209.133:999 | ✓ 1011ms | ✓ 1659ms | ✓ 880ms | ✓ 1916ms | ✓ 1708ms | http |
| 121.40.231.103:7890 | ✓ 928ms | ✓ 1165ms | ✓ 980ms | ✓ 1175ms | ✓ 908ms | http |
| 188.166.208.168:9876 | ✓ 1515ms | 否 | ✓ 828ms | ✓ 1165ms | ✓ 910ms | http |
| 37.27.100.108:443 | 否 | 否 | ✓ 1617ms | ✓ 1599ms | ✓ 1996ms | http |
| 121.230.8.80:1080 | ✓ 1281ms | ✓ 1452ms | ✓ 1202ms | ✓ 1431ms | ✓ 1197ms | http |
| 121.230.8.11:1080 | ✓ 1264ms | ✓ 1742ms | ✓ 1344ms | ✓ 1463ms | ✓ 1178ms | http |
| 121.230.8.49:1080 | ✓ 1752ms | ✓ 1901ms | ✓ 1618ms | ✓ 1544ms | ✓ 1112ms | http |
| 94.177.131.12:3128 | ✓ 731ms | 否 | ✓ 1015ms | ✓ 981ms | ✓ 784ms | http |
| 222.228.171.92:8080 | ✓ 1333ms | 否 | 否 | ✓ 1286ms | ✓ 1012ms | http |
| 103.157.78.122:3125 | ✓ 1733ms | 否 | ✓ 1442ms | ✓ 1364ms | ✓ 1413ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1200ms | ✓ 990ms | 否 | ✓ 1303ms | http |

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
