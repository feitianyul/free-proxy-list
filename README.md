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

最后更新：2026-03-13 00:26:28 UTC（2026-03-13 08:26:28 UTC+8）

**代理总数：71**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 518ms | ✓ 820ms | ✓ 589ms | ✓ 724ms | ✓ 530ms | http |
| 205.209.118.30:3138 | ✓ 387ms | ✓ 1157ms | ✓ 1029ms | ✓ 1344ms | 否 | http |
| 202.155.12.161:443 | ✓ 1561ms | ✓ 1265ms | ✓ 861ms | ✓ 975ms | ✓ 1010ms | http |
| 1.231.81.166:3128 | ✓ 1597ms | ✓ 1002ms | ✓ 1343ms | ✓ 1043ms | ✓ 877ms | http |
| 113.160.132.26:8080 | ✓ 1502ms | ✓ 1336ms | ✓ 1340ms | ✓ 1226ms | ✓ 1024ms | http |
| 152.42.213.210:443 | 否 | 否 | ✓ 1645ms | ✓ 1063ms | ✓ 1015ms | http |
| 178.236.245.59:3128 | ✓ 1318ms | 否 | ✓ 1128ms | 否 | ✓ 1781ms | http |
| 171.251.172.78:5106 | ✓ 1501ms | 否 | ✓ 1560ms | ✓ 1489ms | 否 | http |
| 171.251.172.78:5104 | ✓ 1595ms | 否 | ✓ 1554ms | ✓ 1481ms | 否 | http |
| 103.84.95.54:7890 | ✓ 910ms | 否 | ✓ 629ms | ✓ 800ms | 否 | http |
| 120.92.212.16:8890 | ✓ 952ms | ✓ 1185ms | ✓ 1980ms | 否 | 否 | http |
| 217.76.245.80:999 | ✓ 1094ms | ✓ 1529ms | ✓ 1218ms | ✓ 1647ms | ✓ 1312ms | http |
| 186.148.180.46:999 | 否 | 否 | ✓ 1759ms | ✓ 1894ms | ✓ 1507ms | http |
| 45.136.130.175:8443 | ✓ 650ms | ✓ 612ms | ✓ 619ms | ✓ 695ms | ✓ 526ms | http |
| 193.168.173.136:443 | ✓ 1371ms | ✓ 1942ms | ✓ 976ms | 否 | 否 | http |
| 107.173.52.58:7890 | 否 | ✓ 1385ms | ✓ 998ms | ✓ 1673ms | ✓ 1335ms | http |
| 115.231.181.40:8128 | ✓ 863ms | ✓ 1182ms | ✓ 1881ms | ✓ 1144ms | ✓ 946ms | http |
| 190.9.109.198:999 | ✓ 940ms | ✓ 1536ms | ✓ 1404ms | ✓ 1430ms | ✓ 1249ms | http |
| 14.225.212.37:7890 | ✓ 810ms | ✓ 1339ms | ✓ 815ms | 否 | 否 | http |
| 46.183.25.8:443 | ✓ 710ms | 否 | ✓ 173ms | ✓ 934ms | 否 | http |
| 45.136.130.188:8443 | ✓ 705ms | ✓ 603ms | ✓ 94ms | ✓ 656ms | ✓ 517ms | http |
| 120.92.212.16:7890 | ✓ 966ms | ✓ 1200ms | ✓ 923ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1020ms | ✓ 1369ms | ✓ 1063ms | ✓ 1288ms | ✓ 1159ms | http |
| 101.43.255.96:80 | ✓ 985ms | ✓ 1333ms | ✓ 1534ms | ✓ 1272ms | ✓ 1011ms | http |
| 111.79.111.126:3128 | ✓ 1151ms | ✓ 1482ms | ✓ 1379ms | ✓ 1641ms | ✓ 1090ms | http |
| 101.47.73.135:3128 | ✓ 1251ms | 否 | ✓ 1447ms | ✓ 1480ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1071ms | ✓ 1364ms | 否 | ✓ 1397ms | ✓ 1089ms | http |
| 194.5.212.40:8080 | ✓ 590ms | ✓ 1674ms | 否 | ✓ 1937ms | 否 | http |
| 45.136.130.223:8443 | ✓ 82ms | ✓ 603ms | ✓ 81ms | ✓ 665ms | ✓ 515ms | http |
| 45.136.130.191:8443 | ✓ 315ms | ✓ 877ms | ✓ 85ms | ✓ 688ms | ✓ 514ms | http |
| 45.136.131.47:8443 | ✓ 98ms | ✓ 599ms | ✓ 737ms | ✓ 764ms | ✓ 628ms | http |
| 162.240.154.26:3128 | ✓ 355ms | ✓ 850ms | ✓ 266ms | ✓ 869ms | ✓ 673ms | http |
| 35.225.22.61:80 | ✓ 402ms | ✓ 1269ms | 否 | ✓ 1350ms | 否 | http |
| 45.168.238.193:8443 | ✓ 602ms | ✓ 1166ms | ✓ 247ms | ✓ 1348ms | ✓ 862ms | http |
| 152.42.213.210:8080 | ✓ 731ms | 否 | ✓ 1614ms | ✓ 1549ms | ✓ 1225ms | http |
| 137.184.1.87:3128 | ✓ 307ms | ✓ 810ms | ✓ 972ms | ✓ 709ms | ✓ 534ms | http |
| 24.144.86.173:1080 | ✓ 1211ms | 否 | 否 | ✓ 1090ms | ✓ 573ms | http |
| 178.236.245.17:3128 | ✓ 1776ms | 否 | ✓ 1476ms | 否 | ✓ 1760ms | http |
| 39.104.201.40:7890 | ✓ 1851ms | ✓ 1229ms | 否 | 否 | ✓ 948ms | http |
| 103.183.10.172:3125 | ✓ 1317ms | 否 | ✓ 1672ms | ✓ 1748ms | ✓ 1416ms | http |
| 138.124.53.25:7443 | ✓ 941ms | 否 | ✓ 1769ms | 否 | ✓ 1686ms | http |
| 103.82.23.118:5242 | ✓ 1519ms | 否 | ✓ 1312ms | ✓ 1921ms | ✓ 1423ms | http |
| 168.235.110.63:3128 | ✓ 767ms | ✓ 1121ms | ✓ 1213ms | ✓ 1450ms | ✓ 904ms | http |
| 43.165.195.107:3128 | ✓ 1490ms | ✓ 1963ms | ✓ 1187ms | ✓ 1218ms | ✓ 938ms | http |
| 121.230.8.97:1080 | ✓ 1517ms | ✓ 1763ms | ✓ 1134ms | ✓ 1928ms | ✓ 1192ms | http |
| 121.230.9.198:1080 | ✓ 1083ms | ✓ 1393ms | 否 | 否 | ✓ 1134ms | http |
| 103.126.87.125:8090 | ✓ 1817ms | 否 | 否 | ✓ 1542ms | ✓ 1482ms | http |
| 116.80.49.156:3172 | ✓ 1689ms | 否 | 否 | ✓ 1789ms | ✓ 1624ms | http |
| 101.32.244.83:8080 | ✓ 967ms | 否 | ✓ 927ms | ✓ 1229ms | ✓ 1267ms | http |
| 121.43.196.210:8222 | ✓ 879ms | ✓ 1058ms | ✓ 883ms | ✓ 1150ms | ✓ 880ms | http |
| 121.43.196.213:8222 | ✓ 954ms | ✓ 1064ms | ✓ 835ms | ✓ 1148ms | ✓ 902ms | http |
| 114.55.226.123:10086 | ✓ 1108ms | ✓ 1368ms | ✓ 1056ms | ✓ 1309ms | ✓ 1090ms | http |
| 150.107.140.238:3128 | ✓ 1031ms | 否 | 否 | ✓ 1882ms | ✓ 1684ms | http |
| 62.113.119.14:8080 | ✓ 1006ms | 否 | ✓ 1119ms | ✓ 1779ms | ✓ 1346ms | http |
| 106.117.208.101:7890 | ✓ 1886ms | ✓ 1830ms | ✓ 1057ms | ✓ 1279ms | ✓ 1010ms | http |
| 198.24.188.138:37800 | ✓ 1156ms | 否 | ✓ 1413ms | 否 | ✓ 1403ms | http |
| 165.227.5.10:8888 | ✓ 660ms | ✓ 1588ms | 否 | ✓ 1994ms | 否 | http |
| 202.129.206.239:3128 | ✓ 1851ms | 否 | ✓ 1532ms | ✓ 1608ms | ✓ 1662ms | http |
| 103.113.70.189:1081 | ✓ 895ms | ✓ 1100ms | 否 | ✓ 1281ms | ✓ 1008ms | http |
| 43.167.227.161:1080 | ✓ 1056ms | ✓ 828ms | ✓ 570ms | ✓ 863ms | ✓ 726ms | http |
| 180.127.149.247:1080 | ✓ 873ms | ✓ 1132ms | ✓ 970ms | ✓ 1131ms | ✓ 869ms | http |
| 121.126.185.63:25152 | 否 | 否 | ✓ 1546ms | ✓ 1768ms | ✓ 1385ms | http |
| 45.136.198.40:3128 | ✓ 821ms | ✓ 1984ms | ✓ 1698ms | 否 | ✓ 1693ms | http |
| 103.39.51.190:8080 | ✓ 1864ms | 否 | 否 | ✓ 1327ms | ✓ 1398ms | http |
| 162.248.165.72:1080 | ✓ 1598ms | 否 | ✓ 1795ms | ✓ 1660ms | 否 | http |
| 8.219.97.248:80 | ✓ 1681ms | 否 | 否 | ✓ 1655ms | ✓ 1256ms | http |
| 103.159.96.34:8085 | ✓ 1629ms | 否 | 否 | ✓ 1499ms | ✓ 1497ms | http |
| 61.52.131.172:8443 | ✓ 840ms | ✓ 1171ms | ✓ 966ms | ✓ 1232ms | ✓ 915ms | http |
| 37.187.109.70:10111 | ✓ 1536ms | ✓ 1531ms | ✓ 1152ms | ✓ 1917ms | 否 | http |
| 159.223.42.219:3128 | ✓ 1747ms | 否 | ✓ 825ms | ✓ 1118ms | ✓ 1133ms | http |
| 103.139.138.194:3128 | ✓ 1090ms | 否 | ✓ 1134ms | ✓ 1562ms | ✓ 1337ms | http |

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
