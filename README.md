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

最后更新：2026-03-12 06:46:48 UTC（2026-03-12 14:46:48 UTC+8）

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
| 45.136.131.47:8443 | ✓ 648ms | ✓ 1100ms | ✓ 631ms | ✓ 1148ms | ✓ 770ms | http |
| 45.136.131.63:8443 | ✓ 654ms | 否 | ✓ 374ms | ✓ 1058ms | ✓ 747ms | http |
| 45.136.130.175:8443 | ✓ 646ms | ✓ 891ms | ✓ 860ms | ✓ 1115ms | ✓ 1764ms | http |
| 194.213.18.200:443 | ✓ 445ms | 否 | ✓ 897ms | ✓ 1652ms | ✓ 966ms | http |
| 1.231.81.166:3128 | ✓ 1272ms | ✓ 1567ms | ✓ 983ms | ✓ 1107ms | ✓ 899ms | http |
| 113.160.132.26:8080 | ✓ 1613ms | ✓ 1561ms | ✓ 1126ms | ✓ 1758ms | ✓ 1126ms | http |
| 91.107.141.42:8081 | ✓ 598ms | 否 | ✓ 1316ms | ✓ 1711ms | 否 | http |
| 171.251.172.78:5105 | ✓ 1837ms | 否 | ✓ 1939ms | ✓ 1954ms | ✓ 1696ms | http |
| 171.251.172.78:5102 | ✓ 1924ms | 否 | ✓ 1885ms | 否 | ✓ 1702ms | http |
| 171.251.172.78:5107 | ✓ 1836ms | 否 | ✓ 1860ms | ✓ 1897ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1210ms | ✓ 1130ms | ✓ 850ms | http |
| 205.209.118.30:3138 | ✓ 233ms | ✓ 1806ms | ✓ 834ms | ✓ 1271ms | ✓ 1491ms | http |
| 46.183.25.8:443 | ✓ 837ms | 否 | ✓ 1205ms | ✓ 1205ms | 否 | http |
| 115.231.181.40:8128 | ✓ 947ms | ✓ 1250ms | ✓ 981ms | 否 | ✓ 1912ms | http |
| 59.46.216.131:30001 | ✓ 1434ms | ✓ 1505ms | ✓ 1374ms | 否 | 否 | http |
| 107.173.52.58:7890 | ✓ 366ms | 否 | ✓ 949ms | 否 | ✓ 820ms | http |
| 144.31.25.69:21064 | ✓ 935ms | 否 | ✓ 410ms | 否 | ✓ 1792ms | http |
| 171.251.172.78:5104 | 否 | 否 | ✓ 1816ms | ✓ 1815ms | ✓ 1608ms | http |
| 62.113.119.14:8080 | ✓ 1147ms | 否 | ✓ 558ms | 否 | ✓ 1284ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1011ms | ✓ 1073ms | ✓ 872ms | http |
| 138.124.53.25:7443 | ✓ 994ms | 否 | 否 | ✓ 1994ms | ✓ 1611ms | http |
| 190.9.109.198:999 | ✓ 1001ms | ✓ 1423ms | ✓ 1211ms | ✓ 1487ms | ✓ 1331ms | http |
| 104.248.25.131:3128 | ✓ 384ms | 否 | ✓ 1972ms | ✓ 1545ms | ✓ 1218ms | http |
| 45.136.130.188:8443 | 否 | 否 | ✓ 308ms | ✓ 942ms | ✓ 780ms | http |
| 202.155.12.161:443 | ✓ 1463ms | 否 | 否 | ✓ 1316ms | ✓ 1143ms | http |
| 14.225.222.185:7890 | 否 | ✓ 1662ms | ✓ 1901ms | 否 | ✓ 1132ms | http |
| 162.240.154.26:3128 | ✓ 883ms | 否 | 否 | ✓ 1789ms | ✓ 1461ms | http |
| 45.136.130.191:8443 | ✓ 292ms | ✓ 933ms | ✓ 317ms | ✓ 1020ms | ✓ 719ms | http |
| 123.57.0.163:8888 | 否 | ✓ 1676ms | 否 | ✓ 1885ms | ✓ 1692ms | http |
| 111.79.111.126:3128 | ✓ 1142ms | ✓ 1399ms | ✓ 1132ms | 否 | 否 | http |
| 5.252.33.13:2025 | ✓ 1281ms | 否 | ✓ 1198ms | ✓ 1984ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1192ms | ✓ 1816ms | ✓ 1975ms | 否 | 否 | http |
| 107.172.125.217:3128 | ✓ 1579ms | 否 | ✓ 1057ms | ✓ 1168ms | ✓ 880ms | http |
| 194.5.212.40:8080 | ✓ 918ms | 否 | ✓ 1267ms | ✓ 1684ms | 否 | http |
| 81.70.169.194:80 | ✓ 1166ms | ✓ 1414ms | ✓ 1104ms | ✓ 1279ms | ✓ 1570ms | http |
| 101.43.255.96:80 | ✓ 1451ms | ✓ 1418ms | ✓ 1356ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 396ms | 否 | ✓ 1376ms | ✓ 1426ms | ✓ 1184ms | http |
| 120.92.212.16:7890 | ✓ 1010ms | ✓ 1279ms | ✓ 1036ms | ✓ 1375ms | ✓ 1023ms | http |
| 39.104.201.40:7890 | ✓ 1097ms | ✓ 1407ms | ✓ 1205ms | ✓ 1509ms | ✓ 1177ms | http |
| 162.248.165.72:1080 | ✓ 1486ms | 否 | ✓ 929ms | 否 | ✓ 1962ms | http |
| 120.92.212.16:8890 | ✓ 1009ms | ✓ 1334ms | ✓ 1040ms | ✓ 1570ms | ✓ 1058ms | http |
| 160.250.4.245:1 | ✓ 1301ms | 否 | ✓ 1668ms | ✓ 1455ms | ✓ 1202ms | http |
| 160.250.5.22:1 | ✓ 1300ms | 否 | ✓ 1635ms | ✓ 1563ms | 否 | http |
| 168.235.110.63:3128 | ✓ 520ms | ✓ 1358ms | ✓ 1228ms | 否 | ✓ 864ms | http |
| 111.48.191.1:7890 | ✓ 888ms | ✓ 1151ms | ✓ 940ms | ✓ 1154ms | ✓ 923ms | http |
| 120.92.211.211:7890 | 否 | ✓ 1929ms | ✓ 1295ms | ✓ 1949ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1049ms | ✓ 1317ms | ✓ 1430ms | 否 | ✓ 1238ms | http |
| 103.113.70.189:1081 | ✓ 277ms | ✓ 1186ms | 否 | ✓ 1803ms | 否 | http |
| 45.136.198.40:3128 | ✓ 890ms | 否 | 否 | ✓ 1805ms | ✓ 1551ms | http |
| 210.223.44.230:3128 | ✓ 1946ms | ✓ 1193ms | ✓ 844ms | ✓ 1994ms | ✓ 1016ms | http |
| 45.136.130.223:8443 | ✓ 432ms | ✓ 1568ms | ✓ 939ms | 否 | 否 | http |
| 95.3.9.78:3128 | ✓ 1535ms | 否 | 否 | ✓ 1941ms | ✓ 1620ms | http |
| 95.3.9.78:8080 | ✓ 1534ms | 否 | 否 | ✓ 1937ms | ✓ 1623ms | http |
| 43.165.195.107:3128 | ✓ 1732ms | 否 | ✓ 1361ms | ✓ 1715ms | 否 | http |
| 86.53.183.16:1080 | ✓ 883ms | 否 | ✓ 1092ms | 否 | ✓ 1347ms | http |
| 8.219.97.248:80 | ✓ 1438ms | 否 | ✓ 1723ms | 否 | ✓ 1530ms | http |
| 152.42.213.210:443 | 否 | 否 | ✓ 1564ms | ✓ 1545ms | ✓ 1115ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 966ms | ✓ 1975ms | ✓ 1145ms | http |
| 152.42.213.210:8080 | ✓ 1885ms | 否 | ✓ 1915ms | 否 | ✓ 1117ms | http |
| 171.251.172.78:5106 | ✓ 1764ms | 否 | ✓ 1861ms | ✓ 1670ms | ✓ 1805ms | http |
| 4.233.138.204:8888 | ✓ 1528ms | ✓ 1765ms | 否 | 否 | ✓ 1627ms | http |
| 38.150.4.3:10000 | ✓ 607ms | ✓ 1882ms | ✓ 965ms | ✓ 1037ms | ✓ 824ms | http |
| 200.174.198.32:8888 | ✓ 935ms | 否 | ✓ 1641ms | 否 | ✓ 1638ms | http |
| 121.230.9.5:1080 | ✓ 955ms | ✓ 1499ms | ✓ 1347ms | 否 | ✓ 1452ms | http |
| 121.230.8.245:1080 | ✓ 1173ms | ✓ 1363ms | ✓ 1261ms | ✓ 1534ms | ✓ 1054ms | http |
| 72.56.104.188:1080 | 否 | ✓ 1660ms | 否 | ✓ 1418ms | ✓ 1933ms | http |
| 74.208.234.198:443 | ✓ 1232ms | 否 | ✓ 1951ms | ✓ 1705ms | ✓ 1080ms | http |
| 106.117.208.101:7890 | ✓ 1983ms | ✓ 1432ms | ✓ 1152ms | ✓ 1407ms | 否 | http |

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
