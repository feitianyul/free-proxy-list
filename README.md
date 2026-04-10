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

最后更新：2026-04-10 17:56:38 UTC（2026-04-11 01:56:38 UTC+8）

**代理总数：63**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 828ms | 否 | ✓ 1210ms | 否 | ✓ 919ms | http |
| 218.108.131.186:17890 | ✓ 913ms | ✓ 1133ms | ✓ 956ms | ✓ 1172ms | ✓ 967ms | http |
| 147.161.210.140:8800 | ✓ 1463ms | ✓ 1917ms | ✓ 845ms | ✓ 1210ms | ✓ 904ms | http |
| 167.103.115.102:8800 | ✓ 1167ms | 否 | ✓ 1042ms | ✓ 1124ms | ✓ 1318ms | http |
| 1.231.81.166:3128 | ✓ 1498ms | ✓ 1042ms | ✓ 1782ms | ✓ 1217ms | ✓ 1284ms | http |
| 113.160.132.26:8080 | ✓ 1519ms | ✓ 1598ms | ✓ 1537ms | ✓ 1297ms | ✓ 999ms | http |
| 202.141.161.53:10808 | ✓ 1004ms | 否 | ✓ 1296ms | ✓ 1901ms | ✓ 1070ms | http |
| 115.231.181.40:8128 | ✓ 1784ms | ✓ 1486ms | ✓ 1309ms | 否 | ✓ 1777ms | http |
| 167.103.34.108:8800 | ✓ 1760ms | 否 | 否 | ✓ 1618ms | ✓ 1750ms | http |
| 159.223.225.118:8888 | ✓ 638ms | 否 | 否 | ✓ 1979ms | ✓ 1376ms | http |
| 85.239.59.252:7890 | ✓ 682ms | 否 | ✓ 1290ms | ✓ 1735ms | 否 | http |
| 152.32.132.190:7890 | ✓ 1724ms | 否 | ✓ 992ms | 否 | ✓ 1481ms | http |
| 62.113.119.14:8080 | ✓ 1398ms | 否 | ✓ 699ms | ✓ 1627ms | ✓ 1140ms | http |
| 167.103.144.127:8800 | ✓ 1226ms | 否 | ✓ 1433ms | ✓ 1652ms | ✓ 1632ms | http |
| 45.12.151.226:2829 | ✓ 862ms | 否 | ✓ 1787ms | 否 | ✓ 1769ms | http |
| 45.167.124.52:8080 | ✓ 612ms | ✓ 1667ms | ✓ 633ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1038ms | ✓ 1420ms | ✓ 1167ms | 否 | 否 | http |
| 155.117.18.36:25388 | ✓ 1062ms | ✓ 1094ms | 否 | ✓ 1617ms | ✓ 1191ms | http |
| 167.103.31.122:8800 | ✓ 1537ms | 否 | ✓ 1425ms | ✓ 1671ms | ✓ 1580ms | http |
| 185.76.241.110:10001 | ✓ 919ms | 否 | ✓ 1018ms | ✓ 1829ms | ✓ 1386ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1729ms | ✓ 1564ms | ✓ 1710ms | http |
| 212.58.132.5:8888 | ✓ 1085ms | 否 | ✓ 1446ms | ✓ 1520ms | ✓ 1193ms | http |
| 120.92.212.16:7890 | ✓ 1802ms | 否 | ✓ 1270ms | ✓ 1312ms | ✓ 1269ms | http |
| 45.136.131.55:8451 | ✓ 1896ms | 否 | ✓ 473ms | ✓ 1634ms | ✓ 1950ms | http |
| 147.161.239.240:8800 | 否 | ✓ 1599ms | ✓ 1410ms | ✓ 1572ms | ✓ 1562ms | http |
| 45.167.125.21:999 | ✓ 1658ms | ✓ 1830ms | ✓ 1511ms | ✓ 1931ms | ✓ 1562ms | http |
| 8.209.238.110:47701 | ✓ 561ms | ✓ 1159ms | ✓ 719ms | ✓ 971ms | ✓ 763ms | http |
| 185.76.240.169:10001 | ✓ 949ms | 否 | ✓ 1112ms | 否 | ✓ 1741ms | http |
| 101.43.127.100:8877 | ✓ 878ms | ✓ 1135ms | ✓ 939ms | ✓ 1834ms | 否 | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1245ms | ✓ 1701ms | ✓ 1393ms | http |
| 207.254.71.62:8088 | ✓ 796ms | ✓ 1897ms | ✓ 1681ms | 否 | ✓ 1612ms | http |
| 45.140.147.155:1082 | ✓ 853ms | ✓ 1747ms | ✓ 971ms | ✓ 1734ms | ✓ 1079ms | http |
| 170.106.137.214:7890 | 否 | ✓ 1895ms | 否 | ✓ 1032ms | ✓ 1893ms | http |
| 38.180.2.107:3128 | ✓ 875ms | ✓ 1775ms | 否 | 否 | ✓ 1758ms | http |
| 120.92.212.16:8890 | ✓ 1294ms | ✓ 1324ms | 否 | ✓ 1336ms | ✓ 1044ms | http |
| 91.238.105.64:2024 | ✓ 1330ms | ✓ 1603ms | 否 | 否 | ✓ 1807ms | http |
| 222.228.171.92:8080 | ✓ 1948ms | 否 | 否 | ✓ 1224ms | ✓ 1409ms | http |
| 162.240.154.26:3128 | ✓ 1478ms | ✓ 1891ms | ✓ 1674ms | ✓ 1903ms | 否 | http |
| 116.80.95.231:3172 | ✓ 1570ms | 否 | ✓ 1567ms | ✓ 1949ms | ✓ 1749ms | http |
| 115.204.166.236:40000 | ✓ 988ms | ✓ 1323ms | ✓ 1002ms | ✓ 1280ms | ✓ 1131ms | http |
| 121.230.9.241:1080 | 否 | ✓ 1590ms | ✓ 1477ms | 否 | ✓ 1352ms | http |
| 178.156.224.42:3128 | ✓ 1325ms | ✓ 1983ms | ✓ 1981ms | 否 | 否 | http |
| 45.136.131.47:8452 | ✓ 1094ms | 否 | ✓ 357ms | ✓ 1223ms | ✓ 1657ms | http |
| 38.34.179.174:8453 | ✓ 404ms | ✓ 1097ms | 否 | ✓ 1128ms | ✓ 1052ms | http |
| 45.136.131.67:8453 | ✓ 705ms | 否 | ✓ 1864ms | ✓ 1260ms | 否 | http |
| 47.110.42.192:9003 | ✓ 1736ms | ✓ 1517ms | ✓ 1778ms | ✓ 1990ms | ✓ 1392ms | http |
| 3.8.3.11:9002 | ✓ 734ms | 否 | ✓ 1690ms | 否 | ✓ 1658ms | http |
| 157.0.142.246:10061 | ✓ 1091ms | ✓ 1326ms | ✓ 1072ms | ✓ 1359ms | ✓ 1127ms | http |
| 117.0.86.109:22940 | ✓ 1078ms | 否 | ✓ 1076ms | ✓ 1460ms | ✓ 1007ms | http |
| 34.101.184.164:3128 | ✓ 1586ms | 否 | ✓ 1712ms | ✓ 1580ms | ✓ 1047ms | http |
| 57.128.188.167:9157 | ✓ 1707ms | 否 | ✓ 1854ms | 否 | ✓ 1582ms | http |
| 5.104.87.17:8051 | ✓ 1831ms | 否 | ✓ 1336ms | 否 | ✓ 1459ms | http |
| 195.26.224.49:3128 | ✓ 836ms | 否 | ✓ 845ms | ✓ 1490ms | ✓ 1622ms | http |
| 150.241.106.173:8080 | ✓ 643ms | 否 | ✓ 1595ms | 否 | ✓ 1834ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1211ms | ✓ 1319ms | ✓ 1036ms | http |
| 107.172.102.234:40621 | ✓ 285ms | ✓ 1858ms | ✓ 963ms | ✓ 1014ms | 否 | http |
| 47.105.98.23:3128 | ✓ 1461ms | ✓ 1444ms | ✓ 1700ms | ✓ 1323ms | ✓ 1227ms | http |
| 104.168.93.120:8080 | ✓ 1222ms | ✓ 1576ms | ✓ 797ms | 否 | ✓ 1092ms | http |
| 121.230.8.213:1080 | ✓ 1357ms | ✓ 1508ms | ✓ 1329ms | ✓ 1467ms | ✓ 1125ms | http |
| 138.197.68.35:4857 | ✓ 777ms | ✓ 1374ms | 否 | 否 | ✓ 974ms | http |
| 180.130.80.196:9003 | ✓ 1588ms | ✓ 1542ms | ✓ 1441ms | ✓ 1621ms | ✓ 1354ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1671ms | ✓ 1427ms | ✓ 1477ms | ✓ 1087ms | http |
| 77.93.89.128:47146 | ✓ 1835ms | 否 | 否 | ✓ 1917ms | ✓ 1585ms | http |

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
