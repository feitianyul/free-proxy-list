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

最后更新：2026-04-11 16:31:36 UTC（2026-04-12 00:31:36 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | 否 | 否 | ✓ 1596ms | ✓ 1218ms | ✓ 1029ms | http |
| 167.103.115.102:8800 | ✓ 1770ms | 否 | ✓ 1302ms | ✓ 1197ms | ✓ 1296ms | http |
| 167.103.34.108:8800 | ✓ 1840ms | 否 | ✓ 1738ms | ✓ 1781ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1099ms | ✓ 1887ms | 否 | ✓ 1664ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1820ms | ✓ 1034ms | ✓ 1353ms | ✓ 1102ms | http |
| 38.34.179.24:8447 | ✓ 868ms | ✓ 1894ms | 否 | 否 | ✓ 811ms | http |
| 167.103.144.127:8800 | ✓ 1438ms | 否 | ✓ 1412ms | ✓ 1774ms | ✓ 1778ms | http |
| 45.167.125.21:999 | ✓ 1209ms | 否 | ✓ 1215ms | ✓ 1767ms | ✓ 1359ms | http |
| 43.135.138.31:59394 | ✓ 466ms | 否 | ✓ 416ms | ✓ 850ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1340ms | 否 | ✓ 1452ms | 否 | ✓ 1488ms | http |
| 155.117.18.36:25388 | 否 | 否 | ✓ 724ms | ✓ 1495ms | ✓ 1506ms | http |
| 38.145.220.82:8448 | ✓ 617ms | ✓ 1601ms | ✓ 723ms | 否 | 否 | http |
| 147.161.239.240:8800 | ✓ 793ms | ✓ 1535ms | ✓ 1169ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1126ms | ✓ 1400ms | 否 | ✓ 1411ms | 否 | http |
| 101.43.127.100:8877 | 否 | ✓ 1444ms | ✓ 1018ms | ✓ 1977ms | ✓ 1359ms | http |
| 59.46.216.131:30001 | ✓ 1665ms | 否 | ✓ 1244ms | ✓ 1617ms | ✓ 1600ms | http |
| 186.96.111.214:999 | ✓ 1316ms | 否 | ✓ 1760ms | ✓ 1927ms | ✓ 1754ms | http |
| 38.145.220.188:8451 | ✓ 317ms | 否 | ✓ 343ms | ✓ 923ms | ✓ 1050ms | http |
| 45.136.130.176:8451 | ✓ 537ms | 否 | ✓ 366ms | ✓ 1070ms | ✓ 811ms | http |
| 38.145.220.81:8445 | ✓ 330ms | 否 | ✓ 498ms | ✓ 1215ms | ✓ 703ms | http |
| 38.34.179.21:8452 | ✓ 499ms | 否 | ✓ 414ms | ✓ 1036ms | ✓ 868ms | http |
| 38.34.179.173:8451 | ✓ 455ms | 否 | ✓ 445ms | ✓ 1119ms | ✓ 827ms | http |
| 38.34.179.167:8451 | ✓ 478ms | ✓ 1218ms | ✓ 1121ms | ✓ 1192ms | ✓ 826ms | http |
| 38.145.218.14:8446 | ✓ 567ms | 否 | ✓ 572ms | ✓ 1252ms | ✓ 781ms | http |
| 38.34.179.26:8448 | ✓ 733ms | ✓ 1748ms | ✓ 420ms | ✓ 1470ms | ✓ 825ms | http |
| 38.145.208.211:8453 | ✓ 445ms | 否 | ✓ 899ms | ✓ 1114ms | ✓ 811ms | http |
| 38.145.220.72:8451 | ✓ 300ms | ✓ 1880ms | ✓ 481ms | ✓ 1010ms | 否 | http |
| 38.145.208.220:8448 | ✓ 566ms | ✓ 1407ms | ✓ 537ms | ✓ 1136ms | ✓ 1628ms | http |
| 38.34.179.199:8452 | ✓ 506ms | ✓ 1385ms | ✓ 591ms | ✓ 1371ms | ✓ 805ms | http |
| 38.34.179.86:8447 | ✓ 1558ms | 否 | ✓ 415ms | ✓ 1155ms | ✓ 987ms | http |
| 38.34.179.87:8447 | ✓ 1436ms | ✓ 1408ms | ✓ 932ms | ✓ 1489ms | ✓ 815ms | http |
| 91.238.105.43:2023 | ✓ 1106ms | 否 | ✓ 1637ms | ✓ 1708ms | ✓ 1243ms | http |
| 83.219.250.8:62920 | ✓ 798ms | ✓ 1986ms | ✓ 1505ms | 否 | ✓ 1690ms | http |
| 202.141.161.53:10808 | ✓ 1432ms | ✓ 1365ms | ✓ 1262ms | ✓ 1572ms | ✓ 1288ms | http |
| 62.113.119.14:8080 | ✓ 952ms | 否 | ✓ 806ms | 否 | ✓ 1215ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1373ms | ✓ 1347ms | ✓ 1122ms | http |
| 38.145.220.182:8450 | ✓ 953ms | ✓ 1595ms | ✓ 1478ms | 否 | 否 | http |
| 38.34.179.65:8448 | ✓ 741ms | ✓ 1249ms | ✓ 544ms | ✓ 1099ms | ✓ 1019ms | http |
| 38.145.218.161:8445 | ✓ 420ms | 否 | ✓ 1038ms | ✓ 1031ms | ✓ 820ms | http |
| 38.145.208.241:8447 | ✓ 485ms | 否 | ✓ 666ms | ✓ 935ms | ✓ 674ms | http |
| 38.145.208.242:8444 | ✓ 481ms | 否 | ✓ 660ms | ✓ 1194ms | ✓ 874ms | http |
| 115.231.181.40:8128 | ✓ 980ms | 否 | 否 | ✓ 1344ms | ✓ 1052ms | http |
| 38.34.179.42:8449 | ✓ 1550ms | 否 | ✓ 863ms | ✓ 1450ms | ✓ 1080ms | http |
| 165.225.194.14:10610 | ✓ 941ms | ✓ 1725ms | ✓ 1103ms | ✓ 1490ms | ✓ 1245ms | http |
| 185.46.212.32:11201 | ✓ 934ms | 否 | ✓ 908ms | ✓ 1615ms | ✓ 1137ms | http |
| 165.225.200.17:11276 | ✓ 942ms | 否 | ✓ 831ms | 否 | ✓ 1194ms | http |
| 170.85.30.13:9480 | ✓ 948ms | ✓ 1737ms | ✓ 1095ms | ✓ 1970ms | ✓ 1283ms | http |
| 147.161.130.13:11267 | ✓ 946ms | 否 | ✓ 919ms | ✓ 1789ms | ✓ 1424ms | http |
| 165.225.202.12:9480 | ✓ 946ms | ✓ 1568ms | ✓ 1263ms | 否 | ✓ 1385ms | http |
| 165.225.196.12:9400 | ✓ 949ms | 否 | ✓ 1142ms | 否 | ✓ 1485ms | http |
| 147.161.168.62:11195 | ✓ 956ms | ✓ 1569ms | ✓ 1258ms | 否 | ✓ 1657ms | http |
| 147.161.130.13:11101 | ✓ 949ms | 否 | ✓ 828ms | ✓ 1769ms | 否 | http |
| 165.225.196.12:10011 | ✓ 948ms | 否 | ✓ 1440ms | ✓ 1884ms | ✓ 1430ms | http |
| 165.225.77.180:10945 | ✓ 952ms | ✓ 1614ms | ✓ 1343ms | 否 | 否 | http |
| 147.161.248.13:10013 | ✓ 1205ms | 否 | ✓ 1311ms | 否 | ✓ 1305ms | http |
| 147.161.186.17:10625 | ✓ 959ms | 否 | ✓ 1375ms | ✓ 1674ms | 否 | http |
| 136.226.198.13:9401 | ✓ 946ms | ✓ 1567ms | ✓ 1264ms | 否 | 否 | http |
| 121.230.9.225:1080 | ✓ 1964ms | 否 | ✓ 1086ms | 否 | ✓ 1171ms | http |
| 114.237.77.245:1080 | 否 | 否 | ✓ 1357ms | ✓ 1379ms | ✓ 1103ms | http |
| 35.225.22.61:80 | ✓ 268ms | ✓ 1327ms | 否 | ✓ 1192ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1053ms | ✓ 1360ms | 否 | 否 | ✓ 1110ms | http |
| 147.161.232.13:10011 | ✓ 889ms | 否 | ✓ 417ms | ✓ 1249ms | ✓ 1711ms | http |
| 203.150.128.10:8080 | ✓ 1704ms | 否 | ✓ 1417ms | ✓ 1920ms | 否 | http |
| 38.34.179.71:8453 | ✓ 870ms | ✓ 1586ms | ✓ 481ms | ✓ 1353ms | ✓ 1544ms | http |
| 162.240.154.26:3128 | ✓ 1514ms | ✓ 923ms | 否 | ✓ 1967ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1385ms | 否 | ✓ 1649ms | ✓ 1477ms | 否 | http |
| 121.230.9.125:1080 | ✓ 1406ms | 否 | ✓ 1398ms | 否 | ✓ 1344ms | http |
| 43.99.54.236:5555 | ✓ 795ms | ✓ 1105ms | ✓ 754ms | ✓ 961ms | ✓ 842ms | http |
| 46.30.46.133:3128 | ✓ 823ms | 否 | 否 | ✓ 1827ms | ✓ 1347ms | http |
| 103.113.70.189:1081 | ✓ 852ms | ✓ 1276ms | 否 | 否 | ✓ 779ms | http |
| 152.32.132.190:7890 | ✓ 854ms | 否 | ✓ 1031ms | ✓ 967ms | ✓ 1164ms | http |
| 104.129.202.127:10810 | ✓ 423ms | ✓ 1733ms | ✓ 900ms | ✓ 1107ms | ✓ 775ms | http |
| 104.129.202.127:12354 | ✓ 425ms | 否 | ✓ 807ms | ✓ 959ms | ✓ 891ms | http |
| 218.108.131.186:17890 | ✓ 943ms | ✓ 1217ms | ✓ 964ms | ✓ 1245ms | ✓ 973ms | http |
| 103.242.105.76:80 | ✓ 1933ms | 否 | ✓ 1745ms | ✓ 1876ms | ✓ 1776ms | http |

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
