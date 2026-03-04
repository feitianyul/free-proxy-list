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

最后更新：2026-03-04 00:32:24 UTC（2026-03-04 08:32:24 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 14.56.177.44:3128 | ✓ 1160ms | ✓ 969ms | ✓ 609ms | ✓ 978ms | ✓ 724ms | http |
| 125.128.12.144:3128 | ✓ 1167ms | ✓ 1126ms | ✓ 893ms | ✓ 1095ms | ✓ 1358ms | http |
| 61.72.110.94:3128 | ✓ 1164ms | ✓ 1543ms | ✓ 1340ms | ✓ 1016ms | ✓ 1124ms | http |
| 3.213.157.4:3128 | ✓ 315ms | 否 | ✓ 1716ms | ✓ 1590ms | ✓ 1440ms | http |
| 61.72.221.94:3128 | ✓ 1749ms | ✓ 1452ms | ✓ 1010ms | ✓ 1148ms | 否 | http |
| 166.0.192.117:8888 | ✓ 1039ms | ✓ 1951ms | ✓ 1962ms | ✓ 1215ms | ✓ 1103ms | http |
| 61.72.110.54:3128 | ✓ 1164ms | ✓ 1444ms | ✓ 1980ms | ✓ 978ms | ✓ 1922ms | http |
| 121.128.121.54:3128 | ✓ 1159ms | 否 | ✓ 1858ms | ✓ 944ms | ✓ 714ms | http |
| 35.225.22.61:80 | ✓ 1098ms | ✓ 1302ms | 否 | 否 | ✓ 937ms | http |
| 125.128.12.14:3128 | ✓ 562ms | ✓ 1634ms | ✓ 583ms | ✓ 1170ms | 否 | http |
| 61.72.221.194:3128 | 否 | ✓ 1102ms | ✓ 1295ms | ✓ 942ms | ✓ 765ms | http |
| 205.209.118.30:3138 | ✓ 988ms | 否 | 否 | ✓ 1385ms | ✓ 1231ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1349ms | ✓ 1128ms | 否 | ✓ 1085ms | http |
| 59.46.216.131:30001 | ✓ 842ms | 否 | ✓ 975ms | 否 | ✓ 1089ms | http |
| 186.148.180.46:999 | ✓ 1383ms | 否 | ✓ 1702ms | 否 | ✓ 1765ms | http |
| 14.56.107.244:3128 | ✓ 933ms | ✓ 774ms | ✓ 616ms | ✓ 955ms | ✓ 702ms | http |
| 113.176.92.71:3128 | ✓ 1671ms | ✓ 1204ms | ✓ 1218ms | ✓ 1098ms | ✓ 878ms | http |
| 74.208.234.198:443 | ✓ 1057ms | 否 | ✓ 1913ms | 否 | ✓ 1189ms | http |
| 120.92.212.16:8890 | ✓ 1317ms | 否 | ✓ 960ms | ✓ 1581ms | ✓ 1107ms | http |
| 165.227.5.10:8888 | ✓ 1460ms | ✓ 714ms | ✓ 963ms | 否 | 否 | http |
| 61.72.221.234:3128 | ✓ 1004ms | ✓ 1841ms | ✓ 1583ms | 否 | ✓ 799ms | http |
| 103.3.246.71:3128 | 否 | 否 | ✓ 1197ms | ✓ 1132ms | ✓ 913ms | http |
| 8.219.97.248:80 | ✓ 1271ms | 否 | ✓ 1221ms | ✓ 1569ms | 否 | http |
| 101.43.255.96:80 | ✓ 967ms | ✓ 1105ms | ✓ 884ms | ✓ 1197ms | ✓ 906ms | http |
| 81.70.169.194:80 | ✓ 962ms | ✓ 1112ms | ✓ 982ms | ✓ 1229ms | ✓ 915ms | http |
| 121.204.158.249:3128 | ✓ 914ms | ✓ 1120ms | ✓ 950ms | ✓ 1109ms | ✓ 984ms | http |
| 1.225.116.115:1080 | ✓ 1023ms | ✓ 1125ms | ✓ 1434ms | ✓ 1062ms | ✓ 975ms | http |
| 106.14.203.63:3333 | ✓ 1520ms | ✓ 1668ms | ✓ 1130ms | ✓ 1025ms | ✓ 1576ms | http |
| 185.115.74.185:8080 | ✓ 1063ms | ✓ 1941ms | ✓ 1924ms | 否 | 否 | http |
| 181.78.194.249:999 | ✓ 1540ms | 否 | ✓ 1773ms | 否 | ✓ 1968ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 575ms | ✓ 770ms | ✓ 588ms | http |
| 157.230.220.25:4857 | 否 | 否 | ✓ 1008ms | ✓ 1436ms | ✓ 1039ms | http |
| 159.89.31.62:8080 | ✓ 1102ms | 否 | ✓ 1335ms | ✓ 1870ms | ✓ 1754ms | http |
| 91.193.240.157:9877 | ✓ 1633ms | 否 | ✓ 920ms | 否 | ✓ 1871ms | http |
| 115.231.181.40:8128 | ✓ 1047ms | ✓ 1941ms | ✓ 1164ms | ✓ 1389ms | ✓ 1795ms | http |
| 154.90.48.209:9090 | ✓ 1110ms | 否 | ✓ 1431ms | ✓ 1365ms | ✓ 918ms | http |
| 45.140.147.82:1081 | ✓ 695ms | ✓ 1727ms | ✓ 1128ms | ✓ 1982ms | 否 | http |
| 157.66.36.77:8181 | ✓ 1136ms | 否 | ✓ 1083ms | ✓ 1250ms | ✓ 1220ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1982ms | ✓ 1573ms | ✓ 1849ms | ✓ 1660ms | http |
| 110.232.94.49:8080 | ✓ 1341ms | 否 | ✓ 1357ms | ✓ 1553ms | ✓ 1743ms | http |
| 192.166.82.55:1080 | ✓ 1769ms | 否 | ✓ 1642ms | ✓ 1922ms | 否 | http |
| 150.107.140.238:3128 | ✓ 977ms | 否 | ✓ 991ms | 否 | ✓ 1056ms | http |
| 90.84.188.97:8000 | ✓ 1126ms | 否 | ✓ 1910ms | ✓ 1965ms | ✓ 1623ms | http |
| 94.176.3.43:7443 | ✓ 1267ms | 否 | ✓ 1691ms | 否 | ✓ 1755ms | http |
| 34.96.238.40:8080 | ✓ 743ms | ✓ 1433ms | ✓ 1045ms | 否 | 否 | http |
| 5.75.196.26:40000 | ✓ 642ms | ✓ 1416ms | 否 | ✓ 1816ms | ✓ 1049ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1711ms | ✓ 1139ms | ✓ 1532ms | 否 | http |
| 47.74.226.8:5001 | ✓ 1108ms | ✓ 1625ms | ✓ 933ms | 否 | 否 | http |
| 138.197.68.35:4857 | ✓ 526ms | ✓ 1535ms | ✓ 1085ms | 否 | 否 | http |
| 103.139.138.194:3128 | ✓ 1499ms | 否 | ✓ 1449ms | ✓ 1494ms | ✓ 1225ms | http |
| 91.233.223.147:3128 | ✓ 1549ms | 否 | ✓ 1272ms | 否 | ✓ 1825ms | http |
| 103.159.96.75:8081 | 否 | 否 | ✓ 1513ms | ✓ 1345ms | ✓ 1968ms | http |
| 46.249.103.192:443 | ✓ 985ms | 否 | ✓ 1274ms | ✓ 1992ms | 否 | http |
| 35.234.17.221:8080 | ✓ 1154ms | ✓ 1902ms | ✓ 780ms | 否 | 否 | http |
| 47.84.131.156:8100 | ✓ 1423ms | ✓ 1898ms | ✓ 977ms | 否 | 否 | http |
| 45.140.147.82:1082 | ✓ 1618ms | 否 | ✓ 713ms | ✓ 1426ms | 否 | http |
| 103.112.163.46:8080 | ✓ 1568ms | 否 | ✓ 1570ms | 否 | ✓ 1626ms | http |
| 121.141.161.55:1080 | ✓ 921ms | ✓ 1838ms | ✓ 982ms | ✓ 957ms | ✓ 719ms | http |
| 181.78.44.63:999 | ✓ 1400ms | ✓ 1847ms | ✓ 1421ms | ✓ 1665ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1237ms | 否 | ✓ 1108ms | ✓ 1659ms | ✓ 1254ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1919ms | ✓ 1457ms | ✓ 1911ms | ✓ 1665ms | http |
| 95.85.252.153:21064 | ✓ 639ms | ✓ 1708ms | ✓ 1311ms | 否 | ✓ 1631ms | http |
| 211.171.114.154:3128 | ✓ 1981ms | ✓ 1508ms | ✓ 1844ms | ✓ 1611ms | ✓ 1032ms | http |
| 103.215.36.88:15247 | ✓ 1058ms | ✓ 1248ms | ✓ 1080ms | ✓ 1407ms | ✓ 1042ms | http |
| 132.145.148.124:3128 | ✓ 680ms | ✓ 1936ms | 否 | 否 | ✓ 1397ms | http |
| 27.73.57.47:10002 | ✓ 1616ms | 否 | ✓ 1258ms | ✓ 1511ms | ✓ 1307ms | http |
| 182.253.160.168:1452 | ✓ 1411ms | 否 | ✓ 1112ms | 否 | ✓ 1292ms | http |
| 27.73.57.47:10004 | ✓ 1616ms | 否 | ✓ 1533ms | ✓ 1523ms | ✓ 1296ms | http |
| 70.22.175.232:3128 | ✓ 441ms | ✓ 1267ms | 否 | 否 | ✓ 1131ms | http |
| 143.189.3.198:8080 | ✓ 832ms | ✓ 1127ms | ✓ 437ms | ✓ 749ms | ✓ 571ms | http |
| 45.136.198.40:3128 | ✓ 1075ms | ✓ 1728ms | ✓ 1902ms | 否 | ✓ 1955ms | http |
| 103.39.51.190:8080 | ✓ 1916ms | 否 | 否 | ✓ 1337ms | ✓ 1262ms | http |
| 103.215.36.88:15887 | ✓ 899ms | ✓ 1650ms | 否 | ✓ 1119ms | ✓ 974ms | http |
| 47.110.42.192:9003 | ✓ 1555ms | ✓ 1514ms | ✓ 1296ms | ✓ 1393ms | ✓ 1510ms | http |
| 172.212.68.37:3128 | ✓ 820ms | 否 | ✓ 783ms | ✓ 1503ms | ✓ 1098ms | http |
| 121.230.9.251:1080 | ✓ 1066ms | ✓ 1571ms | ✓ 1435ms | ✓ 1516ms | ✓ 1416ms | http |
| 45.77.249.199:1236 | 否 | 否 | ✓ 720ms | ✓ 921ms | ✓ 843ms | http |
| 212.175.29.184:8080 | ✓ 1489ms | 否 | ✓ 1740ms | 否 | ✓ 1758ms | http |
| 45.140.147.155:1081 | ✓ 611ms | 否 | ✓ 943ms | 否 | ✓ 1284ms | http |
| 34.101.184.164:3128 | ✓ 878ms | 否 | ✓ 1422ms | ✓ 1541ms | ✓ 1509ms | http |
| 180.103.19.123:1080 | ✓ 1403ms | ✓ 1971ms | ✓ 1503ms | 否 | 否 | http |
| 121.230.8.11:1080 | ✓ 1044ms | 否 | ✓ 1462ms | ✓ 1424ms | ✓ 1398ms | http |
| 144.124.227.90:21074 | ✓ 867ms | ✓ 1711ms | ✓ 1353ms | 否 | 否 | http |

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
