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

最后更新：2026-03-01 08:42:08 UTC（2026-03-01 16:42:08 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 148.135.85.87:1080 | ✓ 774ms | ✓ 1263ms | ✓ 1509ms | ✓ 985ms | ✓ 516ms | http |
| 3.213.157.4:3128 | ✓ 231ms | 否 | ✓ 263ms | ✓ 1129ms | ✓ 852ms | http |
| 101.47.73.135:3128 | ✓ 981ms | 否 | ✓ 1025ms | ✓ 1984ms | 否 | http |
| 14.56.107.244:3128 | 否 | ✓ 820ms | ✓ 752ms | 否 | ✓ 757ms | http |
| 217.76.245.80:999 | ✓ 704ms | 否 | ✓ 1205ms | ✓ 1386ms | ✓ 1446ms | http |
| 52.188.28.218:3128 | ✓ 940ms | 否 | ✓ 331ms | ✓ 1124ms | ✓ 1864ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1300ms | ✓ 1536ms | ✓ 1242ms | http |
| 115.231.181.40:8128 | ✓ 1280ms | 否 | ✓ 971ms | ✓ 1261ms | 否 | http |
| 177.243.209.133:999 | ✓ 957ms | 否 | ✓ 510ms | ✓ 1242ms | ✓ 1317ms | http |
| 211.171.114.154:3128 | ✓ 1227ms | 否 | ✓ 1585ms | ✓ 1975ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1905ms | ✓ 1303ms | ✓ 1970ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1311ms | ✓ 1322ms | ✓ 956ms | 否 | 否 | http |
| 81.70.169.194:80 | 否 | ✓ 1279ms | ✓ 955ms | ✓ 1573ms | ✓ 1012ms | http |
| 101.43.255.96:80 | ✓ 946ms | ✓ 1562ms | ✓ 1965ms | ✓ 1285ms | ✓ 1047ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1348ms | ✓ 1141ms | 否 | ✓ 1135ms | http |
| 45.140.147.82:1082 | ✓ 683ms | ✓ 1815ms | ✓ 1122ms | ✓ 1302ms | ✓ 1018ms | http |
| 117.159.239.54:22222 | 否 | ✓ 1713ms | ✓ 1351ms | ✓ 1769ms | 否 | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 802ms | ✓ 1250ms | ✓ 988ms | http |
| 45.205.28.107:8080 | ✓ 899ms | ✓ 1326ms | ✓ 688ms | 否 | 否 | http |
| 205.209.118.30:3138 | ✓ 330ms | 否 | ✓ 366ms | ✓ 1570ms | ✓ 1043ms | http |
| 113.59.32.141:22222 | ✓ 1343ms | ✓ 1378ms | 否 | 否 | ✓ 953ms | http |
| 120.240.35.176:22222 | ✓ 1155ms | ✓ 1512ms | ✓ 1039ms | ✓ 1276ms | ✓ 1049ms | http |
| 217.119.129.86:2222 | ✓ 788ms | 否 | ✓ 736ms | 否 | ✓ 1753ms | http |
| 36.147.78.166:80 | ✓ 1668ms | ✓ 1678ms | ✓ 1644ms | 否 | ✓ 1606ms | http |
| 5.75.201.136:1080 | ✓ 586ms | ✓ 1616ms | ✓ 1034ms | 否 | ✓ 1998ms | http |
| 35.234.17.221:8080 | ✓ 811ms | ✓ 1656ms | ✓ 1440ms | 否 | ✓ 1310ms | http |
| 35.225.22.61:80 | ✓ 565ms | 否 | ✓ 1005ms | ✓ 1124ms | ✓ 1144ms | http |
| 200.125.171.254:999 | ✓ 1157ms | 否 | 否 | ✓ 1438ms | ✓ 1609ms | http |
| 142.171.85.32:1080 | ✓ 489ms | ✓ 1003ms | ✓ 1811ms | ✓ 878ms | ✓ 909ms | http |
| 103.84.95.54:7890 | ✓ 827ms | ✓ 1903ms | ✓ 1129ms | 否 | ✓ 1521ms | http |
| 74.208.234.198:443 | ✓ 991ms | ✓ 1350ms | 否 | ✓ 1092ms | ✓ 997ms | http |
| 222.228.171.92:8080 | ✓ 1740ms | ✓ 1766ms | 否 | 否 | ✓ 1229ms | http |
| 113.59.32.162:22222 | 否 | ✓ 1311ms | ✓ 1144ms | ✓ 1192ms | ✓ 947ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1656ms | ✓ 1604ms | ✓ 1456ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1468ms | ✓ 1481ms | ✓ 1459ms | http |
| 120.238.159.230:22222 | ✓ 929ms | ✓ 1330ms | ✓ 956ms | ✓ 1134ms | ✓ 967ms | http |
| 141.11.210.35:1080 | ✓ 847ms | 否 | ✓ 1119ms | ✓ 776ms | ✓ 609ms | http |
| 113.59.32.142:22222 | ✓ 1089ms | ✓ 1302ms | 否 | ✓ 1292ms | ✓ 959ms | http |
| 120.240.35.178:22222 | ✓ 1071ms | 否 | 否 | ✓ 1609ms | ✓ 1099ms | http |
| 120.240.35.160:22222 | 否 | ✓ 1935ms | 否 | ✓ 1288ms | ✓ 1077ms | http |
| 114.231.73.123:1080 | ✓ 908ms | 否 | ✓ 1495ms | ✓ 1898ms | ✓ 978ms | http |
| 101.32.244.83:8080 | ✓ 1006ms | 否 | ✓ 938ms | ✓ 1573ms | ✓ 1064ms | http |
| 121.43.196.213:8222 | ✓ 973ms | ✓ 1078ms | ✓ 903ms | ✓ 1132ms | ✓ 896ms | http |
| 121.43.196.210:8222 | ✓ 1016ms | ✓ 1158ms | ✓ 848ms | ✓ 1110ms | ✓ 869ms | http |
| 114.55.226.123:10086 | ✓ 1522ms | 否 | ✓ 989ms | ✓ 1285ms | ✓ 1020ms | http |
| 45.136.198.40:3128 | ✓ 829ms | 否 | ✓ 1616ms | 否 | ✓ 1880ms | http |
| 42.115.247.250:10031 | 否 | 否 | ✓ 1421ms | ✓ 1701ms | ✓ 1587ms | http |
| 45.125.67.37:8443 | ✓ 1115ms | 否 | ✓ 988ms | ✓ 1590ms | ✓ 877ms | http |
| 185.243.218.43:49153 | ✓ 1458ms | 否 | ✓ 1815ms | ✓ 1898ms | ✓ 1772ms | http |
| 36.147.78.166:443 | ✓ 1707ms | ✓ 1683ms | 否 | 否 | ✓ 1615ms | http |
| 103.215.36.88:15968 | 否 | ✓ 1296ms | ✓ 1597ms | 否 | ✓ 1006ms | http |
| 167.160.184.231:6005 | ✓ 1699ms | ✓ 1949ms | ✓ 1061ms | ✓ 1229ms | ✓ 1130ms | http |
| 183.249.5.214:22222 | ✓ 928ms | ✓ 867ms | ✓ 724ms | 否 | 否 | http |
| 123.57.0.163:8888 | 否 | ✓ 1750ms | ✓ 1892ms | ✓ 1645ms | ✓ 1411ms | http |
| 103.236.64.247:8888 | ✓ 1916ms | 否 | ✓ 1015ms | ✓ 1724ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1870ms | 否 | ✓ 923ms | ✓ 1732ms | ✓ 1419ms | http |
| 117.159.239.52:22222 | ✓ 901ms | ✓ 1090ms | 否 | ✓ 1045ms | ✓ 848ms | http |
| 120.198.141.75:22222 | ✓ 1085ms | ✓ 1815ms | ✓ 1028ms | ✓ 1242ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1180ms | 否 | 否 | ✓ 1541ms | ✓ 1400ms | http |
| 91.233.223.147:3128 | ✓ 964ms | 否 | ✓ 1292ms | 否 | ✓ 1632ms | http |
| 217.77.102.18:3128 | ✓ 1569ms | 否 | ✓ 1926ms | 否 | ✓ 1671ms | http |
| 113.59.32.163:22222 | ✓ 1403ms | ✓ 1683ms | ✓ 1261ms | ✓ 1305ms | 否 | http |
| 172.212.68.37:3128 | ✓ 260ms | 否 | ✓ 1488ms | ✓ 1828ms | 否 | http |
| 120.240.35.173:22222 | ✓ 1069ms | ✓ 1500ms | ✓ 1167ms | ✓ 1432ms | 否 | http |
| 183.249.5.109:22222 | ✓ 974ms | ✓ 924ms | 否 | 否 | ✓ 809ms | http |
| 222.102.86.137:3040 | 否 | 否 | ✓ 1966ms | ✓ 995ms | ✓ 774ms | http |
| 31.59.129.75:8080 | ✓ 1732ms | 否 | ✓ 1485ms | 否 | ✓ 1540ms | http |
| 103.39.51.190:8080 | ✓ 1651ms | 否 | 否 | ✓ 1513ms | ✓ 1364ms | http |
| 107.174.133.10:3128 | 否 | ✓ 780ms | ✓ 1203ms | ✓ 1145ms | 否 | http |

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
