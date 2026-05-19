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

最后更新：2026-05-19 19:56:52 UTC（2026-05-20 03:56:52 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.130.126.146:6688 | ✓ 1019ms | ✓ 1133ms | ✓ 1308ms | ✓ 1911ms | 否 | http |
| 192.99.8.15:8850 | ✓ 1032ms | 否 | ✓ 1717ms | 否 | ✓ 1095ms | http |
| 161.117.86.53:8100 | ✓ 1994ms | 否 | ✓ 892ms | 否 | ✓ 1688ms | http |
| 185.200.188.234:10001 | ✓ 1597ms | 否 | ✓ 1267ms | 否 | ✓ 1850ms | http |
| 138.2.78.251:8100 | ✓ 929ms | 否 | ✓ 903ms | 否 | ✓ 1250ms | http |
| 138.2.92.70:8100 | ✓ 1564ms | 否 | 否 | ✓ 1916ms | ✓ 1156ms | http |
| 8.218.153.104:8100 | ✓ 1181ms | ✓ 1866ms | ✓ 973ms | ✓ 1330ms | ✓ 1165ms | http |
| 176.111.37.216:39811 | ✓ 1011ms | ✓ 1945ms | ✓ 1358ms | 否 | ✓ 1385ms | http |
| 176.111.37.5:39811 | ✓ 1112ms | ✓ 1929ms | ✓ 1456ms | 否 | ✓ 1428ms | http |
| 45.125.67.37:8443 | ✓ 1432ms | 否 | ✓ 1417ms | ✓ 1425ms | ✓ 1698ms | http |
| 113.160.132.26:8080 | ✓ 1913ms | ✓ 1630ms | 否 | ✓ 1271ms | ✓ 1759ms | http |
| 202.28.194.139:31280 | ✓ 1692ms | 否 | ✓ 1800ms | 否 | ✓ 1934ms | http |
| 174.137.134.182:2999 | 否 | 否 | ✓ 1021ms | ✓ 1098ms | ✓ 1912ms | http |
| 34.87.80.221:30000 | 否 | 否 | ✓ 1510ms | ✓ 1122ms | ✓ 950ms | http |
| 103.154.178.106:8080 | ✓ 1998ms | 否 | 否 | ✓ 1688ms | ✓ 1717ms | http |
| 8.210.48.83:8100 | ✓ 1025ms | 否 | 否 | ✓ 1235ms | ✓ 1147ms | http |
| 170.106.136.181:31002 | ✓ 827ms | ✓ 774ms | ✓ 693ms | ✓ 675ms | ✓ 517ms | http |
| 129.213.162.27:17777 | 否 | ✓ 1940ms | ✓ 1926ms | 否 | ✓ 1166ms | http |
| 8.154.21.175:3128 | ✓ 856ms | ✓ 1343ms | ✓ 1322ms | ✓ 1088ms | ✓ 1321ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1490ms | ✓ 772ms | ✓ 1490ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1299ms | 否 | ✓ 1505ms | ✓ 1526ms | ✓ 1229ms | http |
| 185.253.61.251:3310 | ✓ 903ms | ✓ 1658ms | ✓ 898ms | ✓ 1792ms | 否 | http |
| 45.80.231.251:3128 | ✓ 1020ms | ✓ 1891ms | ✓ 978ms | 否 | 否 | http |
| 74.208.192.81:3129 | ✓ 1482ms | 否 | ✓ 574ms | 否 | ✓ 1987ms | http |
| 152.67.191.232:6800 | ✓ 1053ms | 否 | ✓ 1069ms | ✓ 1454ms | 否 | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 878ms | ✓ 1056ms | ✓ 974ms | http |
| 47.105.98.23:3128 | ✓ 1077ms | ✓ 1168ms | 否 | 否 | ✓ 1653ms | http |
| 85.192.29.60:3128 | ✓ 1359ms | ✓ 1869ms | ✓ 975ms | ✓ 1781ms | ✓ 1496ms | http |
| 69.164.251.114:8080 | ✓ 1545ms | 否 | ✓ 1910ms | ✓ 1982ms | ✓ 1739ms | http |
| 64.188.77.26:3128 | ✓ 1498ms | 否 | 否 | ✓ 1985ms | ✓ 1622ms | http |
| 144.31.25.69:21064 | ✓ 1096ms | 否 | ✓ 936ms | ✓ 1588ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1212ms | ✓ 1234ms | ✓ 1156ms | ✓ 1376ms | http |
| 86.104.72.220:1082 | ✓ 508ms | ✓ 1064ms | ✓ 1130ms | ✓ 1198ms | ✓ 1040ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1997ms | ✓ 806ms | ✓ 1018ms | http |
| 120.92.212.16:7890 | ✓ 1024ms | 否 | ✓ 1177ms | ✓ 1347ms | ✓ 1432ms | http |
| 103.157.117.226:81 | ✓ 1983ms | 否 | ✓ 1862ms | 否 | ✓ 1344ms | http |
| 152.70.91.193:40000 | ✓ 1783ms | 否 | ✓ 1824ms | ✓ 1532ms | ✓ 1118ms | http |
| 116.80.76.252:7779 | ✓ 1582ms | 否 | 否 | ✓ 1837ms | ✓ 1756ms | http |
| 158.101.89.163:3123 | ✓ 1330ms | ✓ 1108ms | ✓ 1436ms | ✓ 867ms | ✓ 615ms | http |
| 121.230.9.225:1080 | ✓ 893ms | ✓ 1200ms | ✓ 910ms | ✓ 1267ms | ✓ 1016ms | http |
| 121.230.9.161:1080 | ✓ 874ms | ✓ 1183ms | ✓ 1007ms | ✓ 1491ms | ✓ 1153ms | http |
| 121.230.8.208:1080 | ✓ 1018ms | ✓ 1348ms | ✓ 1086ms | ✓ 1217ms | ✓ 1075ms | http |
| 172.236.140.167:3128 | ✓ 1067ms | 否 | ✓ 1788ms | ✓ 1382ms | ✓ 1323ms | http |
| 223.16.170.103:80 | ✓ 1556ms | 否 | ✓ 1187ms | ✓ 1038ms | 否 | http |
| 147.45.41.112:1080 | ✓ 1966ms | 否 | ✓ 1764ms | ✓ 1650ms | ✓ 1746ms | http |
| 128.199.116.219:9090 | ✓ 1424ms | 否 | ✓ 1941ms | 否 | ✓ 1085ms | http |
| 81.30.156.115:8080 | ✓ 1618ms | 否 | ✓ 1684ms | 否 | ✓ 1539ms | http |
| 5.252.33.13:2025 | ✓ 1503ms | 否 | ✓ 1471ms | 否 | ✓ 1803ms | http |
| 3.101.133.120:80 | 否 | ✓ 1456ms | ✓ 982ms | ✓ 1157ms | ✓ 1005ms | http |
| 147.45.78.89:1080 | ✓ 1761ms | 否 | ✓ 1135ms | 否 | ✓ 1224ms | http |
| 46.8.112.212:3128 | ✓ 1428ms | 否 | ✓ 1899ms | ✓ 1815ms | 否 | http |
| 114.214.163.108:6789 | ✓ 1057ms | ✓ 1287ms | ✓ 1206ms | ✓ 1293ms | ✓ 1092ms | http |
| 114.214.165.78:10810 | ✓ 1185ms | 否 | ✓ 1054ms | ✓ 1331ms | ✓ 1120ms | http |
| 121.130.199.80:24003 | ✓ 1679ms | ✓ 942ms | ✓ 1396ms | ✓ 1063ms | ✓ 857ms | http |
| 192.81.129.252:3136 | 否 | 否 | ✓ 1814ms | ✓ 1121ms | ✓ 996ms | http |
| 128.199.121.61:9090 | ✓ 1391ms | 否 | ✓ 1376ms | ✓ 1499ms | 否 | http |
| 157.0.142.246:10057 | ✓ 1000ms | ✓ 1303ms | ✓ 1092ms | ✓ 1359ms | 否 | http |
| 94.158.49.82:3128 | ✓ 1802ms | 否 | ✓ 1790ms | 否 | ✓ 1937ms | http |
| 168.110.52.228:3128 | ✓ 1927ms | 否 | ✓ 459ms | ✓ 1835ms | ✓ 1204ms | http |
| 190.12.150.244:999 | ✓ 1095ms | ✓ 1677ms | ✓ 1688ms | 否 | 否 | http |
| 121.230.8.237:1080 | ✓ 957ms | ✓ 1797ms | ✓ 1126ms | ✓ 1545ms | ✓ 1097ms | http |
| 1.231.81.166:3128 | ✓ 1947ms | ✓ 1530ms | 否 | ✓ 1533ms | ✓ 1003ms | http |
| 61.52.131.172:8443 | ✓ 864ms | ✓ 1242ms | ✓ 1012ms | ✓ 1248ms | ✓ 987ms | http |
| 103.172.70.173:8080 | ✓ 1725ms | ✓ 1905ms | ✓ 1714ms | ✓ 1366ms | 否 | http |
| 8.210.132.233:8100 | ✓ 1846ms | 否 | ✓ 1202ms | 否 | ✓ 1965ms | http |
| 8.218.174.172:8100 | ✓ 1430ms | 否 | ✓ 1606ms | ✓ 1248ms | ✓ 1238ms | http |
| 59.46.216.131:30001 | ✓ 1276ms | ✓ 1434ms | 否 | 否 | ✓ 1123ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1885ms | ✓ 1714ms | ✓ 1831ms | http |

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
