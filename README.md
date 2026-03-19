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

最后更新：2026-03-19 15:45:31 UTC（2026-03-19 23:45:31 UTC+8）

**代理总数：71**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 164.90.151.28:3128 | ✓ 1150ms | ✓ 1757ms | ✓ 1147ms | ✓ 873ms | ✓ 686ms | http |
| 202.155.12.161:443 | ✓ 1704ms | 否 | ✓ 1192ms | ✓ 1598ms | 否 | http |
| 147.161.210.140:8800 | ✓ 1707ms | 否 | ✓ 902ms | ✓ 1316ms | ✓ 1089ms | http |
| 147.161.239.240:8800 | ✓ 568ms | ✓ 1598ms | ✓ 1858ms | ✓ 1693ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1672ms | ✓ 1510ms | ✓ 1474ms | ✓ 1375ms | ✓ 1151ms | http |
| 174.138.24.77:1080 | ✓ 911ms | 否 | ✓ 1998ms | ✓ 1246ms | ✓ 955ms | http |
| 45.167.124.52:8080 | ✓ 1636ms | 否 | ✓ 1308ms | ✓ 1687ms | 否 | http |
| 85.198.96.242:3128 | ✓ 1424ms | 否 | ✓ 754ms | ✓ 1974ms | ✓ 1436ms | http |
| 137.220.150.170:6005 | ✓ 860ms | 否 | ✓ 885ms | ✓ 1416ms | 否 | http |
| 35.225.22.61:80 | ✓ 418ms | 否 | 否 | ✓ 1008ms | ✓ 845ms | http |
| 137.184.6.37:3128 | ✓ 548ms | 否 | 否 | ✓ 886ms | ✓ 688ms | http |
| 219.117.204.211:7799 | ✓ 1433ms | 否 | ✓ 617ms | ✓ 1408ms | ✓ 1084ms | http |
| 1.231.81.166:3128 | ✓ 1517ms | ✓ 1756ms | ✓ 1724ms | ✓ 1567ms | ✓ 995ms | http |
| 38.145.208.172:8448 | ✓ 1192ms | 否 | ✓ 1930ms | ✓ 897ms | 否 | http |
| 101.43.127.100:8877 | 否 | ✓ 1303ms | ✓ 1792ms | ✓ 1304ms | ✓ 958ms | http |
| 133.242.138.34:8100 | ✓ 1453ms | 否 | 否 | ✓ 1511ms | ✓ 1752ms | http |
| 162.243.149.86:31028 | ✓ 1147ms | ✓ 1407ms | 否 | ✓ 1890ms | 否 | http |
| 160.238.65.6:3128 | ✓ 851ms | 否 | ✓ 651ms | ✓ 1676ms | ✓ 1068ms | http |
| 160.238.65.9:3128 | ✓ 858ms | 否 | ✓ 644ms | ✓ 1998ms | 否 | http |
| 160.238.65.8:3128 | ✓ 876ms | ✓ 1728ms | ✓ 1600ms | 否 | ✓ 1129ms | http |
| 160.238.65.5:3128 | ✓ 1175ms | ✓ 1293ms | ✓ 762ms | ✓ 1683ms | 否 | http |
| 160.238.65.7:3128 | ✓ 1171ms | 否 | ✓ 1507ms | ✓ 1364ms | 否 | http |
| 160.238.65.3:3128 | ✓ 875ms | 否 | ✓ 933ms | 否 | ✓ 1759ms | http |
| 149.56.24.51:3128 | ✓ 1043ms | 否 | 否 | ✓ 1578ms | ✓ 893ms | http |
| 137.220.150.104:6005 | ✓ 1780ms | 否 | ✓ 1106ms | ✓ 1436ms | ✓ 1049ms | http |
| 115.231.181.40:8128 | ✓ 1170ms | 否 | ✓ 1775ms | ✓ 1728ms | 否 | http |
| 47.101.159.19:8899 | 否 | ✓ 1754ms | ✓ 1070ms | 否 | ✓ 1532ms | http |
| 160.238.65.2:3128 | ✓ 489ms | ✓ 1393ms | ✓ 1855ms | 否 | 否 | http |
| 5.102.109.41:999 | ✓ 1017ms | 否 | ✓ 681ms | ✓ 1300ms | 否 | http |
| 180.103.19.99:1080 | ✓ 1905ms | ✓ 1445ms | ✓ 1249ms | ✓ 1911ms | ✓ 1287ms | http |
| 45.125.67.37:443 | ✓ 1602ms | 否 | ✓ 1547ms | ✓ 1389ms | ✓ 1383ms | http |
| 14.225.212.37:7890 | ✓ 1050ms | 否 | ✓ 990ms | ✓ 1277ms | ✓ 999ms | http |
| 137.220.151.110:6005 | ✓ 1429ms | 否 | ✓ 1174ms | ✓ 1707ms | ✓ 1757ms | http |
| 38.145.220.81:8452 | ✓ 1563ms | 否 | ✓ 439ms | ✓ 1286ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1082ms | ✓ 1380ms | ✓ 1127ms | ✓ 1707ms | ✓ 1128ms | http |
| 164.90.155.209:3128 | ✓ 434ms | ✓ 859ms | ✓ 675ms | ✓ 899ms | ✓ 692ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1332ms | ✓ 1543ms | ✓ 1390ms | http |
| 168.235.110.63:3128 | ✓ 268ms | ✓ 1926ms | 否 | ✓ 1194ms | 否 | http |
| 45.136.131.61:8451 | ✓ 1745ms | 否 | ✓ 1830ms | 否 | ✓ 1594ms | http |
| 202.35.251.72:8080 | 否 | ✓ 1468ms | 否 | ✓ 1499ms | ✓ 1152ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1189ms | ✓ 1608ms | ✓ 1197ms | http |
| 138.124.53.25:7443 | ✓ 571ms | 否 | ✓ 1297ms | 否 | ✓ 1665ms | http |
| 45.136.198.40:3128 | ✓ 1242ms | 否 | 否 | ✓ 1905ms | ✓ 1651ms | http |
| 106.117.208.101:7890 | ✓ 1893ms | ✓ 1542ms | 否 | 否 | ✓ 1243ms | http |
| 137.220.150.152:6005 | ✓ 1680ms | 否 | ✓ 1001ms | ✓ 1227ms | ✓ 973ms | http |
| 38.55.107.254:6005 | 否 | 否 | ✓ 1638ms | ✓ 1779ms | ✓ 1494ms | http |
| 104.248.81.109:3128 | ✓ 475ms | 否 | ✓ 1667ms | ✓ 1542ms | ✓ 1225ms | http |
| 114.237.77.231:1080 | ✓ 1093ms | 否 | 否 | ✓ 1448ms | ✓ 1044ms | http |
| 103.39.51.190:8080 | ✓ 1884ms | 否 | 否 | ✓ 1914ms | ✓ 1582ms | http |
| 106.14.203.63:3333 | ✓ 952ms | 否 | ✓ 1048ms | ✓ 1270ms | ✓ 1524ms | http |
| 34.141.27.50:3128 | ✓ 1081ms | ✓ 1584ms | ✓ 1481ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1095ms | ✓ 1363ms | 否 | ✓ 1391ms | ✓ 1103ms | http |
| 20.120.225.109:3128 | ✓ 776ms | ✓ 1008ms | ✓ 1001ms | ✓ 1391ms | ✓ 997ms | http |
| 62.113.119.14:8080 | ✓ 594ms | ✓ 1783ms | ✓ 1253ms | ✓ 1753ms | ✓ 1169ms | http |
| 38.34.179.178:8445 | 否 | ✓ 1341ms | ✓ 276ms | ✓ 926ms | ✓ 653ms | http |
| 45.136.130.185:8445 | 否 | ✓ 1119ms | ✓ 715ms | 否 | ✓ 717ms | http |
| 45.136.131.47:8449 | 否 | ✓ 1998ms | ✓ 363ms | ✓ 1173ms | 否 | http |
| 217.174.244.117:3129 | ✓ 513ms | 否 | ✓ 1801ms | 否 | ✓ 1482ms | http |
| 38.145.203.108:8450 | 否 | 否 | ✓ 790ms | ✓ 880ms | ✓ 976ms | http |
| 103.53.231.142:24608 | 否 | 否 | ✓ 1089ms | ✓ 1624ms | ✓ 1318ms | http |
| 114.237.77.220:1080 | ✓ 1089ms | 否 | ✓ 1995ms | 否 | ✓ 1128ms | http |
| 46.151.27.7:3128 | ✓ 519ms | 否 | ✓ 1728ms | ✓ 1488ms | 否 | http |
| 103.155.64.217:8080 | ✓ 1861ms | 否 | ✓ 1710ms | ✓ 1783ms | ✓ 1310ms | http |
| 49.156.44.114:8080 | ✓ 1876ms | 否 | ✓ 1617ms | ✓ 1760ms | ✓ 1703ms | http |
| 38.145.208.247:8444 | ✓ 1781ms | ✓ 909ms | ✓ 934ms | 否 | ✓ 686ms | http |
| 101.47.73.135:3128 | ✓ 1108ms | 否 | ✓ 1719ms | ✓ 1496ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1988ms | ✓ 1949ms | ✓ 1894ms | 否 | 否 | http |
| 38.34.179.203:8450 | ✓ 1026ms | ✓ 1099ms | ✓ 1143ms | ✓ 1927ms | ✓ 1179ms | http |
| 88.80.150.82:8080 | ✓ 1457ms | 否 | 否 | ✓ 1973ms | ✓ 1491ms | https |
| 172.212.68.37:3128 | ✓ 733ms | ✓ 1416ms | ✓ 1837ms | ✓ 1237ms | ✓ 1166ms | http |
| 103.113.70.189:1081 | ✓ 430ms | 否 | ✓ 180ms | ✓ 996ms | ✓ 758ms | http |

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
