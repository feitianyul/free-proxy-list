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

最后更新：2026-02-28 19:35:07 UTC（2026-03-01 03:35:07 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 179ms | 否 | ✓ 1027ms | ✓ 1144ms | ✓ 873ms | http |
| 148.135.85.87:1080 | ✓ 777ms | ✓ 1005ms | ✓ 967ms | ✓ 1669ms | ✓ 796ms | http |
| 91.233.223.147:3128 | ✓ 1476ms | 否 | ✓ 1069ms | ✓ 1936ms | ✓ 1576ms | http |
| 168.235.110.63:3128 | ✓ 506ms | ✓ 1639ms | 否 | ✓ 1523ms | ✓ 1066ms | http |
| 62.113.119.14:8080 | ✓ 1262ms | 否 | ✓ 907ms | ✓ 1905ms | ✓ 1313ms | http |
| 104.238.30.63:63744 | ✓ 1886ms | 否 | ✓ 1875ms | 否 | ✓ 1995ms | http |
| 115.231.181.40:8128 | ✓ 940ms | ✓ 1309ms | ✓ 1078ms | 否 | 否 | http |
| 103.84.95.54:7890 | ✓ 1086ms | 否 | ✓ 1233ms | 否 | ✓ 753ms | http |
| 59.46.216.131:30001 | ✓ 1447ms | ✓ 1627ms | ✓ 1194ms | ✓ 1562ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1568ms | ✓ 1065ms | ✓ 1334ms | 否 | http |
| 121.237.181.137:8888 | ✓ 961ms | ✓ 1309ms | ✓ 969ms | ✓ 1361ms | ✓ 941ms | http |
| 81.70.169.194:80 | ✓ 1078ms | ✓ 1381ms | ✓ 1065ms | ✓ 1319ms | ✓ 1124ms | http |
| 101.43.255.96:80 | ✓ 1051ms | ✓ 1435ms | ✓ 1040ms | ✓ 1369ms | ✓ 1262ms | http |
| 3.213.157.4:3128 | ✓ 80ms | ✓ 935ms | ✓ 94ms | ✓ 951ms | ✓ 718ms | http |
| 61.72.110.24:3128 | 否 | 否 | ✓ 1315ms | ✓ 1681ms | ✓ 1625ms | http |
| 121.128.121.184:3128 | ✓ 1423ms | 否 | ✓ 1389ms | 否 | ✓ 1757ms | http |
| 14.56.107.244:3128 | 否 | ✓ 1134ms | ✓ 1304ms | ✓ 1088ms | 否 | http |
| 91.238.104.171:2023 | ✓ 1949ms | 否 | ✓ 1593ms | ✓ 1557ms | 否 | http |
| 35.225.22.61:80 | ✓ 778ms | ✓ 1971ms | ✓ 265ms | ✓ 929ms | ✓ 1073ms | http |
| 34.159.121.205:3128 | ✓ 457ms | 否 | ✓ 917ms | ✓ 1606ms | ✓ 1299ms | http |
| 183.249.5.111:22222 | ✓ 801ms | ✓ 1354ms | ✓ 848ms | ✓ 1104ms | ✓ 807ms | http |
| 35.204.245.176:3128 | ✓ 470ms | ✓ 1711ms | ✓ 1111ms | ✓ 1702ms | ✓ 1686ms | http |
| 120.238.159.228:22222 | ✓ 961ms | ✓ 1436ms | ✓ 1152ms | ✓ 1260ms | ✓ 993ms | http |
| 117.159.239.52:22222 | ✓ 1409ms | ✓ 1178ms | 否 | ✓ 1354ms | ✓ 1019ms | http |
| 109.73.195.10:8888 | ✓ 1050ms | 否 | ✓ 1405ms | ✓ 1983ms | 否 | http |
| 120.240.35.173:22222 | ✓ 1155ms | ✓ 1777ms | ✓ 1119ms | ✓ 1462ms | ✓ 1127ms | http |
| 37.187.109.70:10111 | ✓ 720ms | ✓ 1369ms | ✓ 1508ms | 否 | ✓ 1854ms | http |
| 101.47.73.135:3128 | ✓ 1776ms | 否 | ✓ 1456ms | ✓ 1535ms | ✓ 1777ms | http |
| 120.240.35.176:22222 | ✓ 1138ms | ✓ 1607ms | ✓ 1764ms | ✓ 1477ms | ✓ 1801ms | http |
| 52.188.28.218:3128 | 否 | ✓ 1322ms | ✓ 392ms | ✓ 1400ms | ✓ 1304ms | http |
| 222.228.171.92:8080 | ✓ 1003ms | 否 | 否 | ✓ 1598ms | ✓ 1012ms | http |
| 185.115.74.185:8080 | ✓ 794ms | ✓ 1703ms | ✓ 1629ms | 否 | 否 | http |
| 180.121.187.94:1080 | ✓ 1409ms | ✓ 1554ms | ✓ 1169ms | ✓ 1384ms | ✓ 1143ms | http |
| 34.78.200.22:3128 | ✓ 877ms | ✓ 1581ms | ✓ 1232ms | ✓ 1701ms | ✓ 1251ms | http |
| 34.79.102.160:3128 | ✓ 867ms | ✓ 1814ms | ✓ 1059ms | 否 | ✓ 1407ms | http |
| 34.78.177.18:3128 | ✓ 873ms | 否 | ✓ 1258ms | ✓ 1790ms | ✓ 1732ms | http |
| 34.7.88.87:3128 | ✓ 863ms | 否 | ✓ 1484ms | ✓ 1712ms | ✓ 1499ms | http |
| 90.84.188.97:8000 | ✓ 1064ms | ✓ 1462ms | ✓ 525ms | 否 | ✓ 1462ms | http |
| 45.140.147.82:1081 | ✓ 472ms | ✓ 1612ms | ✓ 1242ms | ✓ 1459ms | 否 | http |
| 210.223.44.230:3128 | ✓ 792ms | 否 | ✓ 829ms | ✓ 1227ms | ✓ 832ms | http |
| 113.59.32.145:22222 | ✓ 1192ms | 否 | ✓ 1234ms | ✓ 1408ms | ✓ 1090ms | http |
| 120.240.35.160:22222 | ✓ 1226ms | ✓ 1743ms | ✓ 1351ms | ✓ 1465ms | ✓ 1119ms | http |
| 94.177.131.12:3128 | ✓ 722ms | 否 | ✓ 850ms | ✓ 1122ms | ✓ 878ms | http |
| 162.240.154.26:3128 | ✓ 1473ms | ✓ 1402ms | ✓ 585ms | 否 | ✓ 773ms | http |
| 107.174.133.10:3128 | ✓ 680ms | ✓ 1069ms | ✓ 851ms | ✓ 1288ms | ✓ 794ms | http |
| 117.159.239.44:22222 | ✓ 930ms | ✓ 1297ms | 否 | 否 | ✓ 975ms | http |
| 222.184.48.252:22222 | ✓ 994ms | ✓ 1289ms | ✓ 1270ms | 否 | ✓ 1001ms | http |
| 222.184.48.248:22222 | 否 | ✓ 1675ms | ✓ 1105ms | ✓ 1377ms | ✓ 1033ms | http |
| 180.103.19.143:1080 | ✓ 1320ms | ✓ 1952ms | 否 | 否 | ✓ 1780ms | http |
| 213.131.85.27:1981 | ✓ 1446ms | 否 | ✓ 1802ms | ✓ 1936ms | 否 | http |
| 70.61.188.34:3128 | ✓ 1806ms | 否 | ✓ 1709ms | 否 | ✓ 1130ms | http |
| 81.177.48.54:2080 | 否 | 否 | ✓ 1890ms | ✓ 1977ms | ✓ 1594ms | http |
| 103.236.64.247:8888 | 否 | ✓ 1432ms | 否 | ✓ 1632ms | ✓ 1282ms | http |
| 61.72.110.94:3128 | ✓ 1888ms | 否 | ✓ 1626ms | ✓ 1189ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1083ms | ✓ 1361ms | ✓ 1146ms | ✓ 1339ms | ✓ 1086ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1555ms | ✓ 1215ms | 否 | ✓ 1151ms | http |
| 222.184.48.251:22222 | ✓ 1085ms | ✓ 1249ms | ✓ 971ms | ✓ 1309ms | 否 | http |
| 120.232.242.119:22222 | ✓ 1593ms | ✓ 1324ms | ✓ 1171ms | ✓ 1370ms | ✓ 1002ms | http |
| 120.240.29.51:22222 | ✓ 973ms | ✓ 1516ms | ✓ 1290ms | 否 | 否 | http |
| 104.238.30.45:59741 | ✓ 1766ms | 否 | ✓ 1839ms | 否 | ✓ 1998ms | http |
| 104.238.30.40:59741 | ✓ 1733ms | 否 | ✓ 1903ms | 否 | ✓ 1999ms | http |
| 61.171.66.158:3128 | ✓ 885ms | ✓ 1133ms | ✓ 1028ms | ✓ 1301ms | ✓ 909ms | http |
| 183.249.5.117:22222 | 否 | 否 | ✓ 1579ms | ✓ 1051ms | ✓ 833ms | http |
| 34.185.159.217:3128 | ✓ 924ms | ✓ 1834ms | ✓ 526ms | ✓ 1493ms | ✓ 1161ms | http |
| 34.158.73.60:3128 | ✓ 903ms | ✓ 1713ms | ✓ 1109ms | ✓ 1907ms | ✓ 1524ms | http |
| 34.32.154.33:3128 | ✓ 932ms | ✓ 1909ms | ✓ 1249ms | 否 | 否 | http |
| 113.59.32.161:22222 | 否 | ✓ 1418ms | ✓ 1039ms | ✓ 1323ms | ✓ 1092ms | http |
| 222.184.48.236:22222 | 否 | ✓ 1178ms | 否 | ✓ 1260ms | ✓ 974ms | http |
| 222.184.48.242:22222 | 否 | 否 | ✓ 1048ms | ✓ 1283ms | ✓ 1129ms | http |
| 103.39.51.190:8080 | ✓ 1878ms | 否 | 否 | ✓ 1690ms | ✓ 1491ms | http |
| 192.71.213.85:9091 | ✓ 958ms | 否 | ✓ 673ms | ✓ 1780ms | 否 | http |
| 36.147.78.166:80 | 否 | ✓ 1798ms | ✓ 1847ms | ✓ 1874ms | 否 | http |
| 45.140.147.155:1081 | ✓ 409ms | ✓ 1217ms | ✓ 978ms | 否 | 否 | http |
| 103.104.99.89:80 | ✓ 1989ms | 否 | ✓ 1869ms | ✓ 1789ms | 否 | http |
| 34.89.174.168:3128 | ✓ 1011ms | 否 | ✓ 1787ms | 否 | ✓ 1993ms | http |
| 103.67.46.225:3125 | ✓ 1877ms | 否 | ✓ 1921ms | 否 | ✓ 1686ms | http |
| 36.147.78.166:443 | ✓ 1901ms | ✓ 1780ms | 否 | 否 | ✓ 1501ms | http |
| 222.184.48.235:22222 | 否 | 否 | ✓ 1417ms | ✓ 1292ms | ✓ 1057ms | http |
| 165.227.5.10:8888 | ✓ 924ms | ✓ 1211ms | ✓ 1304ms | 否 | 否 | http |
| 172.212.68.37:3128 | ✓ 273ms | ✓ 1380ms | ✓ 1429ms | ✓ 1571ms | ✓ 804ms | http |
| 103.215.36.88:15088 | ✓ 1193ms | ✓ 1539ms | ✓ 1136ms | ✓ 1537ms | ✓ 1048ms | http |
| 45.140.147.155:1082 | ✓ 981ms | 否 | ✓ 1245ms | ✓ 1724ms | ✓ 1198ms | http |
| 120.238.159.230:22222 | ✓ 1664ms | ✓ 1300ms | ✓ 1157ms | ✓ 1245ms | ✓ 993ms | http |
| 5.75.201.136:1080 | ✓ 485ms | ✓ 1544ms | ✓ 1708ms | ✓ 1838ms | ✓ 1886ms | http |
| 113.59.32.141:22222 | ✓ 1165ms | ✓ 1491ms | ✓ 1026ms | ✓ 1339ms | ✓ 1051ms | http |
| 144.31.184.218:3128 | ✓ 1357ms | 否 | ✓ 1839ms | 否 | ✓ 1901ms | http |

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
