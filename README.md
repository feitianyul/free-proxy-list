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

最后更新：2026-03-06 16:35:02 UTC（2026-03-07 00:35:02 UTC+8）

**代理总数：74**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 337ms | ✓ 1960ms | ✓ 958ms | ✓ 1045ms | ✓ 790ms | http |
| 136.49.39.94:8888 | ✓ 736ms | 否 | ✓ 1368ms | ✓ 1694ms | ✓ 1259ms | http |
| 152.42.195.165:8888 | ✓ 898ms | 否 | ✓ 961ms | ✓ 1341ms | ✓ 1019ms | http |
| 1.231.81.166:3128 | ✓ 1846ms | ✓ 1117ms | ✓ 1566ms | ✓ 1165ms | ✓ 880ms | http |
| 178.236.245.17:3128 | ✓ 1469ms | 否 | ✓ 1761ms | 否 | ✓ 1670ms | http |
| 35.225.22.61:80 | ✓ 562ms | 否 | ✓ 333ms | 否 | ✓ 1056ms | http |
| 125.128.12.144:3128 | ✓ 888ms | 否 | ✓ 885ms | ✓ 1532ms | ✓ 1120ms | http |
| 159.223.42.219:3128 | ✓ 895ms | 否 | ✓ 1316ms | ✓ 1283ms | ✓ 1013ms | http |
| 14.56.107.244:3128 | ✓ 771ms | ✓ 1430ms | 否 | ✓ 1188ms | ✓ 1973ms | http |
| 61.72.221.94:3128 | ✓ 828ms | 否 | ✓ 1498ms | 否 | ✓ 1835ms | http |
| 107.174.80.186:3128 | 否 | 否 | ✓ 1364ms | ✓ 1054ms | ✓ 921ms | http |
| 178.236.245.59:3128 | ✓ 1783ms | 否 | ✓ 1800ms | 否 | ✓ 1978ms | http |
| 186.148.180.46:999 | ✓ 1115ms | 否 | ✓ 1444ms | 否 | ✓ 1357ms | http |
| 154.37.208.132:30000 | ✓ 884ms | 否 | 否 | ✓ 1370ms | ✓ 1394ms | http |
| 167.172.69.123:8080 | ✓ 951ms | 否 | ✓ 1299ms | ✓ 1279ms | ✓ 1021ms | http |
| 121.128.121.54:3128 | ✓ 1994ms | ✓ 1384ms | ✓ 1065ms | ✓ 1428ms | 否 | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1428ms | ✓ 1561ms | ✓ 1427ms | http |
| 103.84.95.54:7890 | ✓ 1133ms | 否 | ✓ 844ms | ✓ 1143ms | 否 | http |
| 168.235.110.63:3128 | ✓ 1179ms | 否 | ✓ 859ms | ✓ 1593ms | ✓ 756ms | http |
| 91.193.240.157:9877 | ✓ 970ms | 否 | ✓ 1238ms | 否 | ✓ 1576ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1757ms | ✓ 1967ms | ✓ 1673ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1901ms | ✓ 1779ms | ✓ 1744ms | http |
| 23.94.182.50:12345 | ✓ 1471ms | 否 | ✓ 1116ms | 否 | ✓ 1281ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1442ms | ✓ 1298ms | ✓ 1212ms | ✓ 1081ms | http |
| 159.89.31.62:8080 | ✓ 970ms | 否 | ✓ 1762ms | 否 | ✓ 1494ms | http |
| 167.172.69.123:80 | ✓ 943ms | 否 | ✓ 1584ms | ✓ 1401ms | ✓ 1449ms | http |
| 101.43.255.96:80 | ✓ 1137ms | ✓ 1545ms | ✓ 1620ms | 否 | ✓ 1533ms | http |
| 194.59.204.87:9080 | ✓ 435ms | 否 | ✓ 423ms | ✓ 1729ms | 否 | http |
| 2.83.243.148:7777 | 否 | 否 | ✓ 1566ms | ✓ 1917ms | ✓ 1455ms | http |
| 103.82.23.118:5207 | ✓ 1625ms | 否 | ✓ 1575ms | ✓ 1947ms | ✓ 1340ms | http |
| 103.139.138.194:3128 | ✓ 1299ms | 否 | 否 | ✓ 1899ms | ✓ 1459ms | http |
| 42.96.16.158:1311 | ✓ 1029ms | 否 | ✓ 1134ms | ✓ 1390ms | ✓ 1074ms | http |
| 139.99.238.95:8080 | ✓ 1820ms | 否 | ✓ 1176ms | ✓ 1486ms | ✓ 1143ms | http |
| 45.140.147.155:1082 | ✓ 979ms | 否 | ✓ 1057ms | ✓ 1378ms | ✓ 860ms | http |
| 45.140.147.155:1081 | ✓ 979ms | ✓ 1386ms | ✓ 1078ms | ✓ 1477ms | 否 | http |
| 89.185.85.138:1080 | ✓ 479ms | 否 | ✓ 1373ms | ✓ 1340ms | ✓ 1176ms | http |
| 120.92.212.16:7890 | ✓ 1118ms | 否 | ✓ 1379ms | ✓ 1454ms | ✓ 1110ms | http |
| 193.108.118.190:8888 | ✓ 472ms | ✓ 1395ms | ✓ 830ms | ✓ 1572ms | ✓ 1205ms | http |
| 103.215.36.88:17853 | 否 | ✓ 1651ms | ✓ 1364ms | 否 | ✓ 1312ms | http |
| 120.92.212.16:8890 | ✓ 1145ms | ✓ 1477ms | ✓ 1153ms | 否 | ✓ 1397ms | http |
| 74.48.78.224:2080 | ✓ 1060ms | ✓ 1764ms | ✓ 1042ms | 否 | 否 | http |
| 8.217.147.173:8080 | ✓ 1262ms | ✓ 1987ms | ✓ 1266ms | 否 | ✓ 1925ms | http |
| 85.9.195.140:1080 | ✓ 317ms | 否 | 否 | ✓ 1397ms | ✓ 1393ms | http |
| 162.248.165.72:1080 | 否 | 否 | ✓ 1693ms | ✓ 1990ms | ✓ 1647ms | http |
| 107.152.32.98:1305 | ✓ 1161ms | 否 | 否 | ✓ 1888ms | ✓ 1568ms | http |
| 185.191.236.162:3128 | ✓ 1515ms | ✓ 1791ms | ✓ 1445ms | ✓ 1927ms | ✓ 1290ms | http |
| 138.124.53.25:7443 | ✓ 676ms | ✓ 1865ms | 否 | 否 | ✓ 1620ms | http |
| 61.72.110.94:3128 | ✓ 1758ms | 否 | ✓ 1298ms | 否 | ✓ 1417ms | http |
| 70.22.175.232:3128 | ✓ 220ms | 否 | ✓ 1830ms | ✓ 943ms | ✓ 726ms | http |
| 221.127.195.224:8888 | ✓ 1363ms | 否 | ✓ 1423ms | ✓ 1659ms | ✓ 1383ms | http |
| 14.56.177.44:3128 | ✓ 1673ms | 否 | ✓ 1476ms | 否 | ✓ 988ms | http |
| 104.243.46.122:3128 | ✓ 562ms | ✓ 1726ms | ✓ 1181ms | ✓ 1521ms | 否 | http |
| 47.101.149.27:9010 | 否 | ✓ 1544ms | ✓ 1550ms | 否 | ✓ 1551ms | http |
| 192.166.82.55:1080 | ✓ 1016ms | 否 | ✓ 1851ms | ✓ 1378ms | 否 | http |
| 61.72.221.194:3128 | ✓ 866ms | 否 | 否 | ✓ 1301ms | ✓ 1009ms | http |
| 46.249.103.192:443 | ✓ 1016ms | 否 | ✓ 1472ms | 否 | ✓ 1867ms | http |
| 38.180.2.107:3128 | ✓ 954ms | 否 | ✓ 1944ms | ✓ 1958ms | ✓ 1613ms | http |
| 125.128.12.14:3128 | ✓ 834ms | 否 | 否 | ✓ 1254ms | ✓ 1255ms | http |
| 61.72.221.234:3128 | ✓ 1538ms | 否 | ✓ 1721ms | ✓ 1284ms | 否 | http |
| 172.212.68.37:3128 | ✓ 164ms | 否 | ✓ 700ms | ✓ 1562ms | ✓ 1053ms | http |
| 45.136.198.40:3128 | ✓ 596ms | 否 | ✓ 1473ms | ✓ 1984ms | ✓ 1728ms | http |
| 62.113.119.14:8080 | ✓ 1544ms | 否 | ✓ 1182ms | 否 | ✓ 1647ms | http |
| 81.70.169.194:80 | 否 | ✓ 1480ms | ✓ 1538ms | ✓ 1416ms | ✓ 1938ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1806ms | ✓ 1831ms | ✓ 1483ms | http |
| 120.79.99.232:8099 | ✓ 1362ms | ✓ 1832ms | ✓ 1457ms | 否 | 否 | http |
| 193.168.173.136:443 | ✓ 658ms | 否 | ✓ 1972ms | ✓ 1710ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1962ms | 否 | 否 | ✓ 1836ms | ✓ 1705ms | http |
| 61.52.131.172:8443 | ✓ 1030ms | ✓ 1409ms | ✓ 1232ms | ✓ 1426ms | ✓ 1125ms | http |
| 67.169.98.211:443 | ✓ 1542ms | 否 | ✓ 1705ms | ✓ 1480ms | ✓ 1066ms | http |
| 194.213.18.200:443 | ✓ 668ms | 否 | ✓ 1631ms | 否 | ✓ 1138ms | http |
| 83.219.250.8:62920 | ✓ 847ms | 否 | ✓ 1530ms | 否 | ✓ 1860ms | http |
| 88.80.150.82:8080 | ✓ 664ms | 否 | ✓ 784ms | ✓ 1536ms | ✓ 1197ms | https |
| 106.14.203.63:3333 | 否 | ✓ 1866ms | ✓ 1682ms | ✓ 1337ms | ✓ 1750ms | http |
| 59.46.216.131:30001 | ✓ 1250ms | ✓ 1921ms | ✓ 1228ms | 否 | 否 | http |

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
