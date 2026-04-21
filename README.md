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

最后更新：2026-04-21 12:02:27 UTC（2026-04-21 20:02:27 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:8080 | ✓ 546ms | ✓ 1997ms | 否 | ✓ 1498ms | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1095ms | ✓ 837ms | ✓ 833ms | ✓ 659ms | http |
| 116.171.106.26:3443 | 否 | ✓ 1273ms | ✓ 1953ms | ✓ 1599ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1142ms | ✓ 1390ms | ✓ 1072ms | ✓ 1318ms | ✓ 989ms | http |
| 152.42.208.139:8118 | ✓ 693ms | 否 | ✓ 1420ms | ✓ 977ms | ✓ 1147ms | http |
| 113.160.132.26:8080 | ✓ 1400ms | ✓ 1627ms | ✓ 895ms | ✓ 1655ms | ✓ 897ms | http |
| 36.141.21.200:7890 | ✓ 917ms | 否 | ✓ 1742ms | 否 | ✓ 944ms | http |
| 46.101.95.183:8888 | ✓ 1063ms | 否 | ✓ 1686ms | 否 | ✓ 1384ms | http |
| 152.32.132.190:7890 | ✓ 1332ms | ✓ 1432ms | ✓ 1649ms | ✓ 1967ms | 否 | http |
| 162.19.253.202:8443 | ✓ 1215ms | ✓ 1942ms | 否 | 否 | ✓ 1967ms | http |
| 218.108.131.186:17890 | ✓ 705ms | ✓ 745ms | ✓ 629ms | ✓ 836ms | ✓ 688ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1052ms | ✓ 1245ms | ✓ 1059ms | http |
| 47.84.73.61:1080 | ✓ 815ms | 否 | ✓ 870ms | ✓ 1045ms | ✓ 813ms | http |
| 188.253.125.38:28798 | ✓ 746ms | ✓ 1179ms | ✓ 1411ms | ✓ 1025ms | ✓ 1027ms | http |
| 220.197.44.36:3128 | ✓ 1523ms | ✓ 1481ms | ✓ 1143ms | ✓ 1263ms | ✓ 1039ms | http |
| 168.144.75.9:3128 | ✓ 1103ms | 否 | ✓ 1156ms | ✓ 1942ms | 否 | http |
| 159.89.191.221:3128 | ✓ 1125ms | 否 | 否 | ✓ 1327ms | ✓ 1177ms | http |
| 188.246.224.49:7890 | ✓ 1352ms | ✓ 1868ms | ✓ 1134ms | 否 | ✓ 1519ms | http |
| 212.58.132.5:8888 | ✓ 1116ms | 否 | ✓ 1079ms | ✓ 1497ms | ✓ 1212ms | http |
| 42.101.8.101:8888 | ✓ 1320ms | 否 | 否 | ✓ 1364ms | ✓ 1118ms | http |
| 43.132.188.134:443 | 否 | ✓ 1814ms | 否 | ✓ 1686ms | ✓ 827ms | http |
| 59.46.216.131:30001 | ✓ 882ms | 否 | ✓ 1348ms | ✓ 1243ms | ✓ 1997ms | http |
| 78.11.96.22:8888 | ✓ 1268ms | 否 | ✓ 1278ms | ✓ 1729ms | ✓ 1595ms | http |
| 177.93.132.244:3128 | ✓ 868ms | 否 | ✓ 868ms | 否 | ✓ 1868ms | http |
| 91.99.15.45:2095 | ✓ 1382ms | 否 | ✓ 1577ms | 否 | ✓ 1962ms | http |
| 45.153.231.229:8080 | ✓ 1441ms | 否 | ✓ 967ms | 否 | ✓ 1982ms | http |
| 115.231.181.40:8128 | ✓ 814ms | ✓ 767ms | ✓ 987ms | 否 | 否 | http |
| 81.30.156.115:8080 | ✓ 1357ms | 否 | ✓ 1203ms | ✓ 1894ms | ✓ 1697ms | http |
| 185.138.116.150:8080 | ✓ 1402ms | 否 | ✓ 1040ms | ✓ 1791ms | 否 | http |
| 158.160.215.167:8127 | ✓ 1533ms | 否 | ✓ 1504ms | 否 | ✓ 1749ms | http |
| 45.140.147.82:1082 | ✓ 1297ms | ✓ 1398ms | ✓ 603ms | ✓ 1500ms | ✓ 1143ms | http |
| 149.51.42.10:3128 | ✓ 1746ms | ✓ 1402ms | 否 | ✓ 1435ms | 否 | http |
| 45.12.151.226:2829 | ✓ 836ms | ✓ 1898ms | ✓ 1493ms | 否 | ✓ 1649ms | http |
| 45.76.207.177:40000 | 否 | 否 | ✓ 1047ms | ✓ 1474ms | ✓ 1337ms | http |
| 91.233.223.147:3128 | ✓ 998ms | 否 | ✓ 1159ms | 否 | ✓ 1737ms | http |
| 157.230.178.216:8088 | ✓ 581ms | 否 | ✓ 1413ms | 否 | ✓ 1464ms | http |
| 20.127.128.70:8080 | ✓ 1014ms | 否 | ✓ 895ms | 否 | ✓ 1647ms | http |
| 89.208.106.138:10808 | ✓ 692ms | ✓ 1852ms | 否 | 否 | ✓ 1188ms | http |
| 120.92.212.16:7890 | ✓ 775ms | ✓ 1130ms | 否 | ✓ 1096ms | ✓ 893ms | http |
| 103.82.23.118:5261 | ✓ 1489ms | ✓ 1845ms | ✓ 1397ms | 否 | ✓ 1668ms | http |
| 120.92.212.16:8890 | ✓ 1477ms | ✓ 996ms | 否 | ✓ 1358ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1482ms | ✓ 1202ms | ✓ 1035ms | ✓ 1464ms | 否 | http |
| 34.138.160.210:3128 | ✓ 1783ms | 否 | ✓ 1998ms | 否 | ✓ 1999ms | http |
| 208.87.243.199:7878 | ✓ 246ms | ✓ 1249ms | ✓ 1356ms | ✓ 1139ms | ✓ 664ms | http |
| 64.181.240.152:3128 | ✓ 557ms | ✓ 1418ms | 否 | ✓ 975ms | ✓ 755ms | http |
| 210.223.44.230:3128 | ✓ 591ms | ✓ 1456ms | 否 | 否 | ✓ 654ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1183ms | ✓ 1609ms | ✓ 1342ms | http |
| 103.113.70.189:1081 | ✓ 490ms | 否 | ✓ 381ms | ✓ 1743ms | ✓ 978ms | http |
| 84.47.150.125:1080 | 否 | 否 | ✓ 1572ms | ✓ 1676ms | ✓ 1755ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 958ms | ✓ 1751ms | ✓ 1652ms | http |
| 178.63.155.151:9002 | ✓ 1422ms | 否 | 否 | ✓ 1967ms | ✓ 1832ms | http |
| 111.79.111.126:3128 | 否 | 否 | ✓ 1292ms | ✓ 1171ms | ✓ 1316ms | http |
| 139.135.79.218:8080 | 否 | 否 | ✓ 1707ms | ✓ 1929ms | ✓ 1844ms | http |
| 159.203.220.84:3128 | 否 | ✓ 1820ms | ✓ 881ms | ✓ 687ms | ✓ 855ms | http |
| 103.187.146.151:3128 | ✓ 1536ms | ✓ 1961ms | ✓ 812ms | ✓ 1195ms | ✓ 836ms | http |
| 210.77.29.244:6478 | 否 | ✓ 1027ms | 否 | ✓ 1018ms | ✓ 829ms | http |
| 113.176.92.71:3128 | ✓ 1843ms | ✓ 1532ms | ✓ 1400ms | ✓ 1595ms | ✓ 1107ms | http |
| 162.240.154.26:3128 | ✓ 1065ms | ✓ 1846ms | ✓ 1497ms | 否 | 否 | http |
| 192.3.248.190:8014 | ✓ 1119ms | 否 | ✓ 1310ms | ✓ 1087ms | ✓ 835ms | http |
| 128.199.121.61:9090 | ✓ 1325ms | 否 | ✓ 892ms | ✓ 1097ms | ✓ 882ms | http |
| 152.70.91.193:40000 | ✓ 1329ms | 否 | ✓ 1720ms | ✓ 1286ms | ✓ 955ms | http |
| 152.42.177.32:8888 | ✓ 905ms | 否 | ✓ 907ms | ✓ 1191ms | ✓ 1201ms | http |
| 57.128.188.167:9088 | ✓ 1668ms | 否 | ✓ 1617ms | 否 | ✓ 1916ms | http |
| 62.113.119.14:8080 | ✓ 1312ms | 否 | ✓ 1187ms | ✓ 1864ms | 否 | http |
| 144.31.27.49:1080 | ✓ 1002ms | 否 | ✓ 1361ms | 否 | ✓ 1698ms | http |
| 193.177.0.148:60000 | ✓ 930ms | 否 | ✓ 895ms | ✓ 1735ms | ✓ 1929ms | http |
| 104.129.203.245:10139 | ✓ 738ms | ✓ 1250ms | ✓ 526ms | ✓ 729ms | 否 | http |
| 104.129.203.245:10026 | ✓ 738ms | 否 | ✓ 29ms | ✓ 663ms | 否 | http |
| 47.74.226.8:5001 | 否 | ✓ 1323ms | ✓ 997ms | ✓ 1208ms | 否 | http |
| 217.182.195.221:30000 | ✓ 1279ms | 否 | ✓ 1939ms | 否 | ✓ 1548ms | http |
| 103.113.70.189:1082 | ✓ 1238ms | 否 | ✓ 1783ms | 否 | ✓ 1228ms | http |
| 14.143.222.113:57788 | ✓ 1567ms | 否 | ✓ 1002ms | ✓ 1318ms | 否 | http |
| 178.63.155.151:8888 | ✓ 1220ms | 否 | ✓ 1339ms | 否 | ✓ 1731ms | http |
| 160.238.65.5:3128 | ✓ 1686ms | 否 | ✓ 1019ms | 否 | ✓ 1755ms | http |
| 160.238.65.3:3128 | ✓ 1209ms | ✓ 1731ms | ✓ 1765ms | 否 | ✓ 1776ms | http |
| 160.238.65.7:3128 | ✓ 1686ms | ✓ 1475ms | ✓ 1544ms | 否 | ✓ 1785ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1686ms | 否 | ✓ 958ms | ✓ 815ms | http |
| 121.230.9.96:1080 | ✓ 1363ms | ✓ 1781ms | ✓ 885ms | ✓ 1919ms | 否 | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1736ms | ✓ 1866ms | ✓ 1581ms | http |
| 61.52.131.172:8443 | ✓ 660ms | ✓ 1902ms | ✓ 711ms | ✓ 1336ms | ✓ 716ms | http |
| 128.199.116.219:9090 | ✓ 1539ms | 否 | ✓ 1351ms | ✓ 1824ms | 否 | http |
| 146.190.80.158:9090 | ✓ 1630ms | 否 | ✓ 1681ms | 否 | ✓ 1325ms | http |
| 108.131.109.106:5074 | ✓ 864ms | 否 | ✓ 1565ms | 否 | ✓ 1677ms | http |
| 103.156.248.53:8080 | 否 | 否 | ✓ 1859ms | ✓ 1455ms | ✓ 1397ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1806ms | ✓ 1948ms | ✓ 1987ms | http |

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
