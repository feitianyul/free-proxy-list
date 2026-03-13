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

最后更新：2026-03-13 11:36:36 UTC（2026-03-13 19:36:36 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.175:8443 | ✓ 365ms | ✓ 889ms | ✓ 274ms | ✓ 662ms | ✓ 500ms | http |
| 45.136.130.188:8443 | ✓ 493ms | ✓ 1343ms | ✓ 1955ms | 否 | ✓ 1735ms | http |
| 45.167.124.52:8080 | ✓ 1454ms | ✓ 1796ms | ✓ 1574ms | ✓ 1832ms | ✓ 1476ms | http |
| 47.77.193.180:1080 | ✓ 210ms | ✓ 1412ms | ✓ 248ms | ✓ 762ms | ✓ 524ms | http |
| 1.225.116.115:1080 | ✓ 782ms | ✓ 1496ms | ✓ 869ms | ✓ 887ms | ✓ 810ms | http |
| 205.209.118.30:3138 | ✓ 655ms | 否 | ✓ 995ms | ✓ 1300ms | ✓ 1028ms | http |
| 47.101.149.27:9010 | ✓ 1248ms | 否 | ✓ 1304ms | 否 | ✓ 1303ms | http |
| 217.76.245.80:999 | ✓ 904ms | 否 | ✓ 1231ms | ✓ 1618ms | ✓ 1541ms | http |
| 113.160.132.26:8080 | ✓ 1369ms | ✓ 1327ms | ✓ 1769ms | ✓ 1197ms | ✓ 1031ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1173ms | ✓ 1307ms | ✓ 1757ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1147ms | ✓ 1228ms | ✓ 1165ms | http |
| 45.136.131.47:8443 | ✓ 560ms | ✓ 1132ms | ✓ 765ms | ✓ 791ms | ✓ 545ms | http |
| 137.184.0.30:3128 | ✓ 364ms | 否 | ✓ 1194ms | 否 | ✓ 1270ms | http |
| 46.183.25.8:443 | ✓ 593ms | 否 | ✓ 941ms | 否 | ✓ 1603ms | http |
| 62.113.119.14:8080 | ✓ 909ms | 否 | ✓ 901ms | 否 | ✓ 1785ms | http |
| 45.136.130.191:8443 | ✓ 126ms | ✓ 1253ms | ✓ 258ms | 否 | 否 | http |
| 86.109.3.24:10007 | 否 | ✓ 1155ms | ✓ 536ms | ✓ 1125ms | ✓ 940ms | http |
| 157.230.38.173:3128 | ✓ 1650ms | 否 | ✓ 1110ms | ✓ 1051ms | ✓ 819ms | http |
| 101.43.255.96:80 | 否 | ✓ 1228ms | ✓ 1312ms | 否 | ✓ 1090ms | http |
| 8.219.97.248:80 | ✓ 1347ms | 否 | ✓ 1409ms | ✓ 1538ms | 否 | http |
| 120.92.212.16:7890 | ✓ 948ms | 否 | 否 | ✓ 1193ms | ✓ 982ms | http |
| 45.136.131.63:8443 | ✓ 386ms | ✓ 1320ms | ✓ 1508ms | ✓ 668ms | ✓ 516ms | http |
| 45.136.130.223:8443 | ✓ 402ms | ✓ 1664ms | ✓ 698ms | ✓ 927ms | ✓ 1602ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1998ms | 否 | ✓ 1159ms | ✓ 1672ms | http |
| 149.62.191.202:3128 | ✓ 899ms | 否 | ✓ 1973ms | 否 | ✓ 1553ms | http |
| 86.109.3.24:10012 | ✓ 457ms | ✓ 1776ms | ✓ 466ms | ✓ 1142ms | 否 | http |
| 14.225.211.139:7890 | ✓ 891ms | 否 | ✓ 1536ms | ✓ 1381ms | ✓ 1820ms | http |
| 159.223.42.219:3128 | ✓ 837ms | 否 | ✓ 1178ms | ✓ 1334ms | ✓ 1298ms | http |
| 43.167.227.161:1080 | 否 | ✓ 1241ms | ✓ 1159ms | ✓ 1582ms | ✓ 713ms | http |
| 138.124.53.25:7443 | ✓ 1073ms | 否 | 否 | ✓ 1764ms | ✓ 1700ms | http |
| 45.88.0.116:3128 | ✓ 1103ms | 否 | ✓ 1544ms | 否 | ✓ 1275ms | http |
| 45.88.0.99:3128 | 否 | 否 | ✓ 1838ms | ✓ 1548ms | ✓ 1177ms | http |
| 81.70.169.194:80 | 否 | ✓ 1294ms | ✓ 1028ms | 否 | ✓ 1739ms | http |
| 178.236.245.17:3128 | ✓ 1439ms | 否 | ✓ 1187ms | ✓ 1750ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1068ms | 否 | ✓ 942ms | ✓ 1154ms | ✓ 1258ms | http |
| 121.43.196.210:8222 | ✓ 1008ms | ✓ 1053ms | ✓ 954ms | ✓ 1144ms | ✓ 853ms | http |
| 121.43.196.213:8222 | ✓ 1117ms | ✓ 1026ms | ✓ 863ms | ✓ 1149ms | ✓ 965ms | http |
| 114.55.226.123:10086 | ✓ 1111ms | ✓ 1391ms | ✓ 1020ms | ✓ 1294ms | ✓ 1057ms | http |
| 24.199.124.152:3128 | ✓ 309ms | ✓ 1478ms | ✓ 1200ms | ✓ 721ms | ✓ 538ms | http |
| 201.150.116.3:999 | 否 | 否 | ✓ 1369ms | ✓ 1172ms | ✓ 1953ms | http |
| 45.88.0.114:3128 | 否 | ✓ 1583ms | ✓ 1172ms | 否 | ✓ 1166ms | http |
| 106.51.76.128:3127 | ✓ 1276ms | 否 | ✓ 1909ms | 否 | ✓ 1916ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1430ms | ✓ 1734ms | ✓ 1216ms | http |
| 106.117.208.101:7890 | ✓ 1951ms | 否 | 否 | ✓ 1379ms | ✓ 1003ms | http |
| 45.140.147.82:1082 | ✓ 587ms | 否 | ✓ 1672ms | ✓ 1575ms | ✓ 1258ms | http |
| 45.140.147.82:1081 | ✓ 562ms | 否 | ✓ 1167ms | ✓ 1676ms | ✓ 1234ms | http |
| 14.29.168.215:1080 | ✓ 946ms | ✓ 1941ms | ✓ 1159ms | ✓ 1188ms | ✓ 942ms | http |
| 213.220.62.62:3128 | 否 | 否 | ✓ 1073ms | ✓ 1697ms | ✓ 1199ms | http |
| 45.136.198.40:3128 | ✓ 782ms | 否 | ✓ 1195ms | 否 | ✓ 1294ms | http |
| 46.39.105.157:8080 | ✓ 821ms | 否 | ✓ 1918ms | ✓ 1841ms | ✓ 1558ms | http |
| 45.140.147.155:1081 | ✓ 695ms | ✓ 1908ms | ✓ 1770ms | ✓ 1811ms | 否 | http |
| 128.199.120.45:9090 | ✓ 910ms | 否 | ✓ 1009ms | ✓ 1067ms | ✓ 859ms | http |
| 104.248.151.93:9090 | ✓ 748ms | 否 | ✓ 910ms | ✓ 1362ms | ✓ 887ms | http |
| 35.225.22.61:80 | ✓ 741ms | ✓ 1861ms | 否 | ✓ 1114ms | 否 | http |
| 116.80.49.172:3172 | 否 | 否 | ✓ 1990ms | ✓ 1855ms | ✓ 1704ms | http |
| 115.231.181.40:8128 | ✓ 916ms | 否 | ✓ 883ms | 否 | ✓ 1942ms | http |
| 152.53.194.38:7890 | ✓ 292ms | ✓ 1909ms | ✓ 551ms | ✓ 1171ms | ✓ 912ms | http |
| 137.184.6.117:3128 | ✓ 295ms | 否 | 否 | ✓ 919ms | ✓ 555ms | http |
| 27.133.238.94:80 | ✓ 1438ms | ✓ 823ms | ✓ 1424ms | ✓ 1223ms | ✓ 835ms | http |
| 42.84.157.30:10808 | ✓ 989ms | ✓ 1748ms | ✓ 1058ms | ✓ 1245ms | ✓ 1000ms | http |
| 176.122.64.111:44000 | ✓ 1728ms | ✓ 1795ms | ✓ 1600ms | 否 | ✓ 1633ms | http |
| 103.67.46.225:3125 | ✓ 1867ms | 否 | ✓ 1670ms | 否 | ✓ 1588ms | http |
| 103.39.51.190:8080 | ✓ 1677ms | 否 | ✓ 1389ms | ✓ 1437ms | ✓ 1315ms | http |
| 211.171.114.154:3128 | ✓ 1916ms | ✓ 1304ms | 否 | ✓ 1393ms | ✓ 1058ms | http |

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
