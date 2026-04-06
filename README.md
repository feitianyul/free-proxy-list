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

最后更新：2026-04-06 06:26:32 UTC（2026-04-06 14:26:32 UTC+8）

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
| 147.161.210.140:8800 | ✓ 1378ms | ✓ 1575ms | ✓ 928ms | ✓ 1083ms | ✓ 727ms | http |
| 167.103.115.102:8800 | ✓ 939ms | 否 | ✓ 926ms | ✓ 1761ms | ✓ 1075ms | http |
| 35.225.22.61:80 | 否 | ✓ 1633ms | 否 | ✓ 1511ms | ✓ 1098ms | http |
| 147.161.239.240:8800 | ✓ 1277ms | 否 | ✓ 1372ms | ✓ 1709ms | ✓ 1658ms | http |
| 8.219.97.248:80 | ✓ 1087ms | 否 | ✓ 1240ms | 否 | ✓ 1549ms | http |
| 167.103.34.108:8800 | ✓ 1533ms | 否 | ✓ 1485ms | ✓ 1473ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1939ms | ✓ 1706ms | ✓ 1441ms | ✓ 1457ms | ✓ 1610ms | http |
| 111.227.254.11:22222 | ✓ 1133ms | 否 | ✓ 1912ms | ✓ 1354ms | ✓ 1848ms | http |
| 111.227.254.10:22222 | 否 | ✓ 1943ms | 否 | ✓ 1341ms | ✓ 1119ms | http |
| 111.227.254.12:22222 | ✓ 1698ms | ✓ 1601ms | ✓ 1049ms | ✓ 1350ms | 否 | http |
| 1.225.116.115:1080 | ✓ 1794ms | 否 | 否 | ✓ 1485ms | ✓ 1338ms | http |
| 45.167.124.52:8080 | ✓ 1067ms | 否 | ✓ 1567ms | ✓ 1935ms | 否 | http |
| 159.223.71.162:8080 | ✓ 758ms | 否 | ✓ 1671ms | ✓ 1404ms | ✓ 851ms | http |
| 1.231.81.166:3128 | ✓ 1042ms | 否 | ✓ 701ms | ✓ 967ms | ✓ 718ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1759ms | ✓ 1276ms | ✓ 1197ms | ✓ 991ms | http |
| 167.103.144.127:8800 | ✓ 1656ms | 否 | ✓ 1411ms | 否 | ✓ 1951ms | http |
| 217.76.245.80:999 | ✓ 910ms | ✓ 1763ms | ✓ 1227ms | ✓ 1630ms | ✓ 1252ms | http |
| 45.167.125.21:999 | ✓ 962ms | ✓ 1933ms | ✓ 1423ms | ✓ 1690ms | ✓ 1456ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1191ms | ✓ 1081ms | ✓ 1031ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1763ms | 否 | ✓ 1448ms | ✓ 1664ms | ✓ 1569ms | http |
| 209.38.154.7:1080 | ✓ 1145ms | ✓ 1932ms | 否 | ✓ 1374ms | 否 | http |
| 104.129.203.245:10139 | 否 | 否 | ✓ 107ms | ✓ 789ms | ✓ 567ms | http |
| 104.129.203.245:10733 | 否 | 否 | ✓ 79ms | ✓ 763ms | ✓ 633ms | http |
| 104.129.203.245:10026 | 否 | ✓ 752ms | ✓ 890ms | ✓ 1817ms | ✓ 603ms | http |
| 101.43.127.100:8877 | ✓ 1011ms | ✓ 1962ms | 否 | ✓ 1081ms | ✓ 930ms | http |
| 159.223.71.162:443 | ✓ 1535ms | 否 | ✓ 1637ms | 否 | ✓ 838ms | http |
| 111.227.254.9:22222 | ✓ 1225ms | ✓ 1318ms | ✓ 1723ms | 否 | ✓ 1752ms | http |
| 177.234.217.88:999 | ✓ 1331ms | 否 | ✓ 1806ms | ✓ 1925ms | ✓ 1621ms | http |
| 110.76.145.119:8080 | ✓ 1777ms | 否 | 否 | ✓ 1580ms | ✓ 1490ms | http |
| 62.113.119.14:8080 | ✓ 1625ms | 否 | ✓ 1116ms | ✓ 1653ms | ✓ 1338ms | http |
| 34.101.184.164:3128 | ✓ 1487ms | 否 | ✓ 1130ms | ✓ 1357ms | ✓ 1027ms | http |
| 212.58.132.5:8888 | ✓ 1364ms | 否 | ✓ 1810ms | ✓ 1585ms | ✓ 1280ms | http |
| 121.230.8.158:1080 | 否 | 否 | ✓ 1207ms | ✓ 1519ms | ✓ 1550ms | http |
| 59.46.216.131:30001 | ✓ 968ms | ✓ 1253ms | 否 | 否 | ✓ 1064ms | http |
| 218.108.131.186:17890 | ✓ 825ms | ✓ 1015ms | ✓ 816ms | ✓ 1077ms | ✓ 839ms | http |
| 76.169.128.104:8080 | ✓ 1404ms | ✓ 917ms | 否 | ✓ 1072ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1028ms | ✓ 1643ms | ✓ 944ms | ✓ 1245ms | ✓ 1261ms | http |
| 121.43.196.213:8222 | ✓ 1019ms | ✓ 1054ms | ✓ 886ms | ✓ 1128ms | ✓ 907ms | http |
| 121.43.196.210:8222 | ✓ 1016ms | ✓ 1046ms | ✓ 864ms | ✓ 1123ms | ✓ 933ms | http |
| 114.55.226.123:10086 | ✓ 1141ms | ✓ 1393ms | ✓ 1010ms | ✓ 1308ms | ✓ 1056ms | http |
| 91.233.223.147:3128 | ✓ 1540ms | 否 | ✓ 1229ms | 否 | ✓ 1651ms | http |
| 154.56.114.91:8082 | ✓ 1561ms | 否 | ✓ 1451ms | ✓ 1665ms | 否 | http |
| 4.216.195.194:3128 | ✓ 1310ms | ✓ 1098ms | ✓ 772ms | ✓ 1183ms | ✓ 744ms | http |
| 43.167.237.94:3128 | ✓ 1473ms | 否 | 否 | ✓ 1218ms | ✓ 1881ms | http |
| 5.104.87.17:8051 | ✓ 1298ms | 否 | ✓ 1121ms | ✓ 1142ms | 否 | http |
| 106.117.208.101:7890 | 否 | ✓ 1951ms | 否 | ✓ 1811ms | ✓ 1698ms | http |
| 103.113.70.189:1081 | ✓ 373ms | ✓ 1119ms | ✓ 283ms | ✓ 1245ms | ✓ 878ms | http |
| 24.144.86.173:1080 | ✓ 311ms | ✓ 746ms | ✓ 1447ms | ✓ 1122ms | ✓ 1641ms | http |
| 38.34.179.98:8451 | ✓ 1675ms | ✓ 1236ms | ✓ 450ms | ✓ 1498ms | ✓ 1920ms | http |
| 183.238.3.150:7897 | ✓ 906ms | ✓ 1107ms | ✓ 948ms | ✓ 1061ms | 否 | http |
| 116.233.160.106:7897 | ✓ 901ms | ✓ 1104ms | ✓ 1043ms | ✓ 1194ms | ✓ 1749ms | http |
| 16.78.119.130:443 | ✓ 1960ms | 否 | ✓ 1973ms | ✓ 1598ms | ✓ 1570ms | http |
| 217.76.46.230:8080 | 否 | 否 | ✓ 1153ms | ✓ 1907ms | ✓ 1747ms | http |
| 209.126.10.139:3128 | ✓ 587ms | ✓ 1116ms | 否 | ✓ 1144ms | ✓ 881ms | http |
| 195.123.209.48:3128 | ✓ 850ms | 否 | ✓ 1670ms | 否 | ✓ 1891ms | http |
| 181.78.44.63:999 | ✓ 1003ms | ✓ 1315ms | ✓ 509ms | 否 | ✓ 1835ms | http |
| 38.191.122.74:8181 | 否 | 否 | ✓ 1905ms | ✓ 1680ms | ✓ 1878ms | http |
| 38.34.179.66:8444 | 否 | ✓ 686ms | 否 | ✓ 1773ms | ✓ 1306ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1851ms | ✓ 1145ms | ✓ 900ms | http |
| 64.227.76.27:1080 | ✓ 971ms | 否 | ✓ 1075ms | 否 | ✓ 1683ms | http |
| 150.241.71.15:1080 | ✓ 627ms | 否 | ✓ 817ms | ✓ 1571ms | 否 | http |
| 194.67.99.223:1080 | ✓ 834ms | 否 | 否 | ✓ 1836ms | ✓ 1645ms | http |
| 89.208.106.138:10808 | ✓ 621ms | ✓ 1546ms | ✓ 1329ms | 否 | 否 | http |
| 103.135.102.161:8080 | ✓ 1033ms | 否 | ✓ 1120ms | ✓ 784ms | ✓ 1067ms | http |
| 101.132.61.121:8888 | ✓ 1455ms | ✓ 1217ms | ✓ 1325ms | ✓ 1380ms | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1177ms | ✓ 1192ms | ✓ 1221ms | ✓ 991ms | http |
| 103.191.92.157:1009 | ✓ 1849ms | 否 | ✓ 1153ms | ✓ 1263ms | ✓ 1294ms | http |
| 103.39.51.207:8080 | ✓ 1430ms | 否 | 否 | ✓ 1686ms | ✓ 1885ms | http |
| 120.92.212.16:8890 | ✓ 1052ms | ✓ 1180ms | ✓ 914ms | ✓ 1201ms | ✓ 984ms | http |

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
