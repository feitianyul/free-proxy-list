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

最后更新：2026-03-13 10:37:25 UTC（2026-03-13 18:37:25 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 571ms | ✓ 1101ms | ✓ 1005ms | ✓ 1142ms | ✓ 692ms | http |
| 45.136.130.188:8443 | ✓ 565ms | ✓ 1855ms | ✓ 1243ms | ✓ 883ms | ✓ 704ms | http |
| 205.209.118.30:3138 | ✓ 1403ms | ✓ 1150ms | 否 | ✓ 1067ms | ✓ 814ms | http |
| 46.183.25.8:443 | ✓ 1483ms | 否 | ✓ 1923ms | ✓ 1693ms | ✓ 1673ms | http |
| 178.236.245.17:3128 | ✓ 1319ms | 否 | ✓ 1601ms | 否 | ✓ 1390ms | http |
| 113.160.132.26:8080 | ✓ 1616ms | ✓ 1873ms | ✓ 1349ms | ✓ 1416ms | ✓ 1143ms | http |
| 5.129.206.247:8888 | ✓ 1317ms | 否 | ✓ 1661ms | ✓ 1938ms | 否 | http |
| 45.167.124.52:8080 | ✓ 942ms | ✓ 1803ms | ✓ 1435ms | ✓ 1713ms | ✓ 1397ms | http |
| 185.115.74.185:8080 | ✓ 1293ms | ✓ 1984ms | ✓ 1418ms | 否 | 否 | http |
| 160.238.65.7:3128 | ✓ 1061ms | ✓ 1261ms | ✓ 1819ms | ✓ 1849ms | ✓ 1313ms | http |
| 120.92.212.16:8890 | ✓ 1169ms | ✓ 1505ms | ✓ 1095ms | ✓ 1484ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1219ms | ✓ 1549ms | ✓ 1089ms | ✓ 1441ms | 否 | http |
| 45.136.130.175:8443 | ✓ 459ms | ✓ 844ms | ✓ 871ms | ✓ 918ms | ✓ 712ms | http |
| 45.136.131.63:8443 | ✓ 276ms | 否 | ✓ 787ms | ✓ 883ms | ✓ 691ms | http |
| 81.70.169.194:80 | 否 | ✓ 1603ms | ✓ 1287ms | ✓ 1896ms | ✓ 1882ms | http |
| 101.43.255.96:80 | ✓ 1206ms | 否 | ✓ 1229ms | ✓ 1795ms | ✓ 1243ms | http |
| 115.231.181.40:8128 | ✓ 1810ms | 否 | ✓ 1448ms | 否 | ✓ 1056ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1867ms | ✓ 1006ms | ✓ 952ms | http |
| 47.77.193.180:1080 | 否 | ✓ 1675ms | ✓ 325ms | ✓ 940ms | ✓ 682ms | http |
| 45.136.130.223:8443 | 否 | ✓ 923ms | ✓ 301ms | ✓ 882ms | ✓ 750ms | http |
| 45.136.130.191:8443 | 否 | ✓ 1533ms | ✓ 302ms | ✓ 905ms | ✓ 690ms | http |
| 47.101.149.27:9010 | ✓ 1518ms | 否 | ✓ 1537ms | ✓ 1585ms | 否 | http |
| 210.223.44.230:3128 | ✓ 769ms | ✓ 1672ms | ✓ 1097ms | ✓ 1112ms | ✓ 1958ms | http |
| 160.238.65.6:3128 | ✓ 655ms | ✓ 1461ms | 否 | 否 | ✓ 1642ms | http |
| 160.238.65.3:3128 | ✓ 655ms | ✓ 1481ms | 否 | 否 | ✓ 1623ms | http |
| 157.230.38.173:3128 | ✓ 1358ms | 否 | ✓ 891ms | ✓ 1265ms | ✓ 1021ms | http |
| 160.250.4.245:1 | ✓ 1062ms | 否 | ✓ 1611ms | ✓ 1634ms | ✓ 1160ms | http |
| 201.150.116.3:999 | ✓ 527ms | ✓ 1814ms | ✓ 483ms | ✓ 1244ms | ✓ 1083ms | http |
| 24.144.86.173:1080 | ✓ 579ms | ✓ 1660ms | 否 | 否 | ✓ 1976ms | http |
| 137.184.6.117:3128 | ✓ 490ms | 否 | ✓ 1145ms | ✓ 967ms | ✓ 1742ms | http |
| 14.225.8.195:3218 | ✓ 1413ms | 否 | ✓ 1193ms | ✓ 1592ms | ✓ 1143ms | http |
| 213.220.62.62:3128 | ✓ 387ms | ✓ 1209ms | ✓ 1501ms | ✓ 1188ms | ✓ 1388ms | http |
| 45.88.0.111:3128 | ✓ 362ms | ✓ 1254ms | ✓ 428ms | ✓ 1187ms | ✓ 924ms | http |
| 45.88.0.99:3128 | ✓ 361ms | ✓ 1140ms | ✓ 440ms | ✓ 1220ms | ✓ 917ms | http |
| 45.88.0.117:3128 | ✓ 363ms | ✓ 1782ms | ✓ 355ms | ✓ 1242ms | ✓ 964ms | http |
| 45.88.0.113:3128 | ✓ 362ms | 否 | ✓ 370ms | ✓ 1190ms | ✓ 916ms | http |
| 45.88.0.114:3128 | ✓ 376ms | 否 | ✓ 366ms | ✓ 1183ms | ✓ 912ms | http |
| 45.88.0.116:3128 | ✓ 366ms | ✓ 1968ms | ✓ 399ms | ✓ 1196ms | ✓ 950ms | http |
| 45.88.0.115:3128 | ✓ 355ms | 否 | ✓ 389ms | ✓ 1199ms | ✓ 948ms | http |
| 45.88.0.98:3128 | ✓ 367ms | 否 | ✓ 373ms | ✓ 1190ms | ✓ 961ms | http |
| 88.80.150.82:8080 | ✓ 898ms | 否 | 否 | ✓ 1945ms | ✓ 1561ms | https |
| 14.225.211.139:7890 | 否 | 否 | ✓ 1225ms | ✓ 1360ms | ✓ 1115ms | http |
| 152.42.213.210:8080 | ✓ 1790ms | 否 | ✓ 1936ms | ✓ 1411ms | ✓ 1461ms | http |
| 185.191.236.162:3128 | ✓ 1990ms | 否 | ✓ 1568ms | 否 | ✓ 1898ms | http |
| 192.71.213.85:9812 | ✓ 1188ms | 否 | ✓ 1476ms | ✓ 1730ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1568ms | 否 | ✓ 1316ms | 否 | ✓ 1742ms | http |
| 103.113.70.189:1081 | ✓ 286ms | ✓ 1926ms | ✓ 869ms | ✓ 1074ms | ✓ 870ms | http |
| 45.136.198.40:3128 | ✓ 587ms | ✓ 1462ms | ✓ 611ms | ✓ 1817ms | ✓ 1492ms | http |
| 103.84.95.54:7890 | ✓ 1008ms | 否 | 否 | ✓ 1859ms | ✓ 805ms | http |
| 59.46.216.131:30001 | ✓ 1240ms | 否 | ✓ 1284ms | 否 | ✓ 1233ms | http |
| 106.117.208.101:7890 | ✓ 1219ms | 否 | ✓ 1194ms | ✓ 1464ms | ✓ 1169ms | http |
| 5.102.109.41:999 | ✓ 1848ms | 否 | 否 | ✓ 1340ms | ✓ 1167ms | http |
| 223.16.170.103:80 | ✓ 1317ms | 否 | ✓ 1320ms | ✓ 1307ms | ✓ 1350ms | http |
| 45.186.6.104:3128 | ✓ 1461ms | ✓ 1923ms | ✓ 1856ms | 否 | 否 | http |
| 120.238.159.234:22222 | ✓ 1120ms | ✓ 1398ms | 否 | 否 | ✓ 1163ms | http |
| 192.71.213.85:5141 | ✓ 1157ms | 否 | ✓ 1876ms | ✓ 1622ms | 否 | http |
| 83.219.250.8:62920 | ✓ 840ms | 否 | ✓ 1577ms | ✓ 1805ms | 否 | http |
| 152.42.213.210:443 | ✓ 1291ms | 否 | ✓ 1762ms | ✓ 1825ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1960ms | 否 | 否 | ✓ 1868ms | ✓ 1648ms | http |
| 103.183.10.172:3125 | ✓ 1966ms | 否 | ✓ 1577ms | ✓ 1709ms | ✓ 1700ms | http |
| 61.52.131.172:8443 | ✓ 1058ms | ✓ 1373ms | ✓ 1075ms | ✓ 1362ms | ✓ 1079ms | http |
| 167.71.196.28:8080 | ✓ 1501ms | 否 | ✓ 907ms | ✓ 1798ms | 否 | http |

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
