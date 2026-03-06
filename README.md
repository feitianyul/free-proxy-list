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

最后更新：2026-03-06 12:26:10 UTC（2026-03-06 20:26:10 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 295ms | 否 | ✓ 1333ms | 否 | ✓ 990ms | http |
| 1.231.81.166:3128 | ✓ 1507ms | ✓ 1749ms | ✓ 914ms | ✓ 891ms | ✓ 734ms | http |
| 152.42.195.165:8888 | ✓ 1118ms | 否 | ✓ 777ms | ✓ 1301ms | ✓ 850ms | http |
| 159.223.42.219:3128 | ✓ 1120ms | 否 | ✓ 1244ms | ✓ 1067ms | ✓ 852ms | http |
| 103.84.95.54:7890 | ✓ 716ms | 否 | ✓ 857ms | 否 | ✓ 710ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1885ms | 否 | ✓ 894ms | ✓ 729ms | http |
| 61.72.221.234:3128 | ✓ 1535ms | ✓ 959ms | ✓ 1288ms | 否 | 否 | http |
| 14.56.107.244:3128 | ✓ 738ms | 否 | ✓ 1198ms | ✓ 1020ms | 否 | http |
| 125.128.12.14:3128 | ✓ 671ms | 否 | ✓ 1805ms | ✓ 1570ms | ✓ 1226ms | http |
| 192.166.82.55:1080 | ✓ 1358ms | 否 | ✓ 1950ms | ✓ 1837ms | ✓ 1300ms | http |
| 107.174.80.186:3128 | ✓ 1326ms | 否 | 否 | ✓ 1163ms | ✓ 1020ms | http |
| 61.72.221.94:3128 | 否 | 否 | ✓ 1371ms | ✓ 1857ms | ✓ 1977ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1231ms | ✓ 1012ms | ✓ 1857ms | http |
| 35.225.22.61:80 | ✓ 579ms | ✓ 1246ms | ✓ 952ms | ✓ 1023ms | ✓ 866ms | http |
| 85.9.195.140:1080 | ✓ 527ms | 否 | 否 | ✓ 1585ms | ✓ 1427ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1111ms | ✓ 1829ms | ✓ 1008ms | ✓ 785ms | http |
| 178.236.245.17:3128 | ✓ 982ms | 否 | ✓ 673ms | ✓ 1805ms | ✓ 1392ms | http |
| 178.236.245.59:3128 | ✓ 980ms | 否 | ✓ 869ms | ✓ 1810ms | ✓ 1332ms | http |
| 167.172.69.123:8080 | 否 | 否 | ✓ 1106ms | ✓ 1047ms | ✓ 866ms | http |
| 23.224.193.46:3128 | ✓ 1083ms | ✓ 1756ms | 否 | ✓ 1638ms | ✓ 1241ms | http |
| 193.124.190.224:53294 | ✓ 1165ms | 否 | ✓ 1603ms | 否 | ✓ 1890ms | http |
| 91.107.175.112:10801 | ✓ 943ms | 否 | ✓ 1394ms | 否 | ✓ 1450ms | http |
| 14.56.177.44:3128 | 否 | 否 | ✓ 1795ms | ✓ 1846ms | ✓ 1982ms | http |
| 42.115.72.27:2039 | 否 | 否 | ✓ 1724ms | ✓ 1668ms | ✓ 1843ms | http |
| 120.92.212.16:7890 | ✓ 1807ms | ✓ 1229ms | 否 | ✓ 1520ms | 否 | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 982ms | ✓ 1591ms | ✓ 1281ms | http |
| 116.80.82.232:3172 | 否 | 否 | ✓ 1539ms | ✓ 1868ms | ✓ 1709ms | http |
| 210.223.44.230:3128 | 否 | ✓ 931ms | ✓ 752ms | 否 | ✓ 751ms | http |
| 101.43.255.96:80 | ✓ 961ms | 否 | ✓ 1087ms | ✓ 1265ms | ✓ 987ms | http |
| 61.72.221.194:3128 | ✓ 1576ms | 否 | ✓ 965ms | ✓ 1941ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1620ms | ✓ 1740ms | 否 | ✓ 1314ms | 否 | http |
| 91.193.240.157:9877 | ✓ 1155ms | 否 | ✓ 1451ms | 否 | ✓ 1955ms | http |
| 81.70.169.194:80 | ✓ 1161ms | 否 | 否 | ✓ 1310ms | ✓ 1319ms | http |
| 115.231.181.40:8128 | ✓ 880ms | 否 | ✓ 1275ms | ✓ 1255ms | ✓ 1715ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1161ms | ✓ 1323ms | ✓ 1229ms | http |
| 23.224.193.45:3128 | ✓ 889ms | 否 | ✓ 1772ms | ✓ 1541ms | ✓ 1431ms | http |
| 23.224.193.44:3128 | ✓ 900ms | 否 | ✓ 1687ms | ✓ 1636ms | 否 | http |
| 103.86.131.62:80 | ✓ 1027ms | 否 | ✓ 1810ms | ✓ 1346ms | ✓ 1085ms | http |
| 23.224.193.43:3128 | ✓ 1193ms | 否 | ✓ 1330ms | ✓ 1680ms | 否 | http |
| 162.248.165.72:1080 | ✓ 954ms | 否 | ✓ 1899ms | 否 | ✓ 1841ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1769ms | ✓ 963ms | ✓ 1263ms | 否 | http |
| 154.90.48.209:9090 | ✓ 1895ms | 否 | ✓ 1250ms | 否 | ✓ 1692ms | http |
| 103.139.138.194:3128 | ✓ 1512ms | 否 | ✓ 1095ms | ✓ 1407ms | ✓ 1061ms | http |
| 152.70.137.18:8888 | 否 | 否 | ✓ 1876ms | ✓ 1227ms | ✓ 1374ms | http |
| 59.46.216.131:30001 | ✓ 1671ms | 否 | 否 | ✓ 1814ms | ✓ 1140ms | http |
| 125.27.155.78:8080 | ✓ 1396ms | 否 | ✓ 1357ms | 否 | ✓ 1484ms | http |
| 59.153.16.105:20909 | ✓ 1695ms | 否 | ✓ 1458ms | 否 | ✓ 1584ms | http |
| 120.92.212.16:8890 | ✓ 1056ms | 否 | ✓ 1208ms | 否 | ✓ 1004ms | http |
| 61.72.110.54:3128 | ✓ 911ms | 否 | ✓ 935ms | 否 | ✓ 833ms | http |
| 168.235.110.63:3128 | ✓ 249ms | ✓ 1398ms | ✓ 526ms | ✓ 1241ms | ✓ 1736ms | http |
| 146.190.232.76:3128 | ✓ 568ms | 否 | ✓ 1837ms | 否 | ✓ 1567ms | http |
| 116.80.82.224:3172 | 否 | 否 | ✓ 1531ms | ✓ 1845ms | ✓ 1921ms | http |
| 104.129.203.244:10785 | ✓ 582ms | 否 | ✓ 108ms | ✓ 968ms | ✓ 616ms | http |
| 104.129.203.244:11522 | ✓ 774ms | 否 | ✓ 114ms | ✓ 761ms | ✓ 745ms | http |
| 104.129.203.244:11022 | ✓ 585ms | ✓ 1286ms | ✓ 1404ms | ✓ 774ms | ✓ 588ms | http |
| 104.129.203.244:10013 | ✓ 585ms | ✓ 1341ms | ✓ 691ms | ✓ 1915ms | ✓ 584ms | http |
| 104.129.203.244:11465 | ✓ 583ms | ✓ 1316ms | ✓ 837ms | ✓ 1805ms | ✓ 602ms | http |
| 167.172.69.123:80 | ✓ 737ms | 否 | ✓ 1325ms | ✓ 1071ms | ✓ 1045ms | http |
| 193.108.118.190:8888 | ✓ 654ms | 否 | ✓ 1591ms | 否 | ✓ 1611ms | http |
| 45.140.147.82:1082 | ✓ 1196ms | 否 | ✓ 702ms | 否 | ✓ 1079ms | http |
| 45.140.147.82:1081 | ✓ 1206ms | ✓ 1614ms | ✓ 627ms | 否 | 否 | http |
| 46.249.103.192:443 | ✓ 1225ms | 否 | ✓ 1414ms | ✓ 1946ms | 否 | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 1069ms | ✓ 1031ms | ✓ 794ms | http |
| 185.191.236.162:3128 | ✓ 1001ms | 否 | ✓ 1659ms | 否 | ✓ 1720ms | http |
| 103.215.36.88:18977 | ✓ 1942ms | ✓ 1927ms | 否 | ✓ 1272ms | ✓ 1052ms | http |
| 116.80.82.219:3172 | ✓ 1647ms | 否 | 否 | ✓ 1984ms | ✓ 1779ms | http |
| 61.72.110.94:3128 | 否 | 否 | ✓ 1306ms | ✓ 1037ms | ✓ 1851ms | http |
| 152.69.229.220:3128 | ✓ 1797ms | 否 | ✓ 1390ms | ✓ 1210ms | ✓ 922ms | http |
| 104.129.203.244:10571 | ✓ 490ms | ✓ 1266ms | ✓ 713ms | ✓ 761ms | ✓ 600ms | http |
| 104.129.203.244:11763 | ✓ 487ms | ✓ 902ms | ✓ 927ms | ✓ 1363ms | ✓ 609ms | http |
| 154.37.208.132:30000 | ✓ 900ms | ✓ 1709ms | 否 | ✓ 1880ms | 否 | http |
| 14.225.222.164:7890 | ✓ 856ms | 否 | ✓ 1350ms | ✓ 1845ms | ✓ 1554ms | http |
| 42.115.72.27:2038 | ✓ 1720ms | 否 | ✓ 1449ms | 否 | ✓ 1562ms | http |
| 103.215.36.88:17565 | ✓ 1140ms | 否 | 否 | ✓ 1771ms | ✓ 1351ms | http |
| 14.225.222.247:7890 | 否 | ✓ 1886ms | ✓ 873ms | ✓ 1043ms | ✓ 875ms | http |
| 116.80.82.228:3172 | ✓ 1642ms | 否 | 否 | ✓ 1834ms | ✓ 1702ms | http |
| 186.148.180.46:999 | ✓ 710ms | ✓ 1903ms | ✓ 681ms | 否 | ✓ 1483ms | http |
| 172.212.68.37:3128 | 否 | 否 | ✓ 886ms | ✓ 1505ms | ✓ 1149ms | http |
| 45.136.198.40:3128 | ✓ 853ms | 否 | ✓ 1634ms | 否 | ✓ 1719ms | http |
| 62.113.119.14:8080 | ✓ 778ms | 否 | ✓ 824ms | 否 | ✓ 1208ms | http |
| 130.61.139.145:3128 | ✓ 1092ms | 否 | 否 | ✓ 1630ms | ✓ 1775ms | http |
| 144.208.127.181:3128 | ✓ 834ms | ✓ 1360ms | ✓ 1170ms | ✓ 1623ms | 否 | http |
| 2.56.178.131:443 | ✓ 1151ms | 否 | ✓ 1177ms | ✓ 1993ms | 否 | http |
| 192.71.213.85:9091 | ✓ 742ms | 否 | ✓ 660ms | ✓ 1647ms | 否 | http |
| 103.67.46.225:3125 | ✓ 1845ms | 否 | 否 | ✓ 1644ms | ✓ 1568ms | http |
| 103.39.51.190:8080 | ✓ 1711ms | 否 | 否 | ✓ 1784ms | ✓ 1315ms | http |
| 116.80.82.225:3172 | ✓ 1893ms | 否 | ✓ 1575ms | 否 | ✓ 1772ms | http |
| 42.115.72.27:2065 | ✓ 1446ms | 否 | ✓ 1472ms | ✓ 1656ms | ✓ 1518ms | http |
| 209.38.51.97:3128 | ✓ 1068ms | 否 | 否 | ✓ 1223ms | ✓ 1023ms | http |
| 106.14.203.63:3333 | ✓ 1010ms | ✓ 1575ms | ✓ 1495ms | 否 | ✓ 1804ms | http |

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
