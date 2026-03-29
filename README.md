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

最后更新：2026-03-29 00:27:33 UTC（2026-03-29 08:27:33 UTC+8）

**代理总数：93**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 859ms | ✓ 1128ms | ✓ 828ms | ✓ 1043ms | ✓ 826ms | http |
| 39.185.46.193:5911 | ✓ 830ms | ✓ 1042ms | ✓ 1074ms | ✓ 1135ms | ✓ 875ms | http |
| 95.213.217.168:52004 | ✓ 1068ms | ✓ 1701ms | ✓ 1164ms | ✓ 1500ms | ✓ 1080ms | http |
| 147.161.210.140:8800 | ✓ 1682ms | 否 | ✓ 959ms | ✓ 1021ms | ✓ 929ms | http |
| 219.117.204.211:7799 | 否 | ✓ 1469ms | ✓ 1049ms | 否 | ✓ 837ms | http |
| 167.103.115.102:8800 | ✓ 1814ms | ✓ 1893ms | ✓ 1426ms | ✓ 1284ms | ✓ 1212ms | http |
| 103.84.95.54:7890 | ✓ 1057ms | 否 | ✓ 860ms | ✓ 1783ms | ✓ 1835ms | http |
| 167.103.34.108:8800 | ✓ 1846ms | 否 | ✓ 1397ms | ✓ 1829ms | 否 | http |
| 103.9.78.2:3128 | 否 | ✓ 1954ms | ✓ 1743ms | ✓ 1439ms | ✓ 1589ms | http |
| 45.167.124.52:8080 | ✓ 1006ms | ✓ 1886ms | ✓ 1280ms | ✓ 1881ms | 否 | http |
| 35.225.22.61:80 | ✓ 439ms | ✓ 1646ms | 否 | ✓ 970ms | ✓ 921ms | http |
| 167.103.144.127:8800 | ✓ 1344ms | 否 | ✓ 1558ms | ✓ 1457ms | ✓ 1352ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1593ms | ✓ 1337ms | 否 | ✓ 1173ms | http |
| 167.103.31.122:8800 | ✓ 1651ms | 否 | ✓ 1241ms | ✓ 1493ms | ✓ 1404ms | http |
| 45.144.232.5:11741 | ✓ 900ms | 否 | ✓ 1503ms | 否 | ✓ 1598ms | http |
| 45.12.151.226:2829 | ✓ 1833ms | ✓ 1681ms | 否 | ✓ 1947ms | ✓ 1326ms | http |
| 120.92.212.16:8890 | ✓ 1876ms | 否 | ✓ 1489ms | ✓ 1437ms | ✓ 1824ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1529ms | ✓ 1300ms | 否 | ✓ 1214ms | http |
| 38.145.218.232:8445 | ✓ 606ms | ✓ 864ms | ✓ 943ms | ✓ 1048ms | ✓ 878ms | http |
| 38.145.208.242:8445 | ✓ 1006ms | ✓ 1058ms | ✓ 1815ms | ✓ 1464ms | ✓ 1057ms | http |
| 38.145.220.173:8453 | ✓ 1380ms | 否 | ✓ 1356ms | ✓ 1924ms | ✓ 1965ms | http |
| 208.87.243.199:7878 | ✓ 546ms | ✓ 897ms | ✓ 1148ms | ✓ 1093ms | ✓ 838ms | http |
| 101.43.127.100:8877 | ✓ 1111ms | ✓ 1293ms | ✓ 1080ms | ✓ 1376ms | ✓ 1061ms | http |
| 128.199.254.13:9090 | ✓ 910ms | 否 | ✓ 1788ms | ✓ 1560ms | ✓ 1071ms | http |
| 128.199.121.61:9090 | ✓ 953ms | 否 | ✓ 1493ms | ✓ 1573ms | ✓ 1053ms | http |
| 45.15.158.60:2222 | ✓ 783ms | ✓ 1924ms | ✓ 1663ms | 否 | ✓ 1715ms | http |
| 128.199.114.189:9090 | ✓ 1257ms | 否 | ✓ 1370ms | ✓ 1660ms | ✓ 1140ms | http |
| 45.8.157.38:3128 | ✓ 624ms | ✓ 1609ms | ✓ 1307ms | ✓ 1432ms | 否 | http |
| 106.75.15.167:7890 | ✓ 1503ms | 否 | ✓ 1188ms | ✓ 1502ms | 否 | http |
| 177.234.217.88:999 | ✓ 1196ms | 否 | ✓ 1591ms | 否 | ✓ 1651ms | http |
| 160.238.65.8:3128 | 否 | ✓ 1842ms | ✓ 507ms | ✓ 1264ms | ✓ 960ms | http |
| 160.238.65.9:3128 | 否 | ✓ 1771ms | ✓ 579ms | ✓ 1260ms | ✓ 959ms | http |
| 160.238.65.4:3128 | 否 | ✓ 1501ms | ✓ 1032ms | 否 | ✓ 1061ms | http |
| 160.238.65.3:3128 | ✓ 1633ms | 否 | ✓ 717ms | 否 | ✓ 970ms | http |
| 160.238.65.7:3128 | ✓ 1638ms | 否 | ✓ 712ms | ✓ 1260ms | 否 | http |
| 193.233.22.29:10808 | ✓ 256ms | 否 | ✓ 806ms | ✓ 1331ms | ✓ 896ms | http |
| 147.161.239.240:8800 | ✓ 847ms | ✓ 1552ms | ✓ 825ms | ✓ 1639ms | ✓ 1269ms | http |
| 128.199.116.219:9090 | ✓ 1546ms | 否 | ✓ 1308ms | ✓ 1279ms | ✓ 1045ms | http |
| 146.190.80.158:9090 | ✓ 1551ms | 否 | ✓ 1534ms | ✓ 1276ms | ✓ 1036ms | http |
| 128.199.113.85:9090 | ✓ 1549ms | 否 | ✓ 1374ms | ✓ 1442ms | ✓ 1042ms | http |
| 160.238.65.6:3128 | ✓ 1274ms | 否 | ✓ 395ms | ✓ 1183ms | 否 | http |
| 194.59.204.87:9080 | ✓ 848ms | ✓ 1264ms | ✓ 1099ms | 否 | 否 | http |
| 181.41.201.85:3128 | ✓ 629ms | ✓ 1960ms | ✓ 871ms | 否 | 否 | http |
| 93.43.251.159:80 | ✓ 1273ms | ✓ 1902ms | ✓ 1284ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 1542ms | ✓ 1854ms | ✓ 1977ms | 否 | ✓ 1323ms | http |
| 88.80.150.82:8080 | ✓ 1891ms | ✓ 1897ms | 否 | ✓ 1969ms | ✓ 1546ms | https |
| 120.92.212.16:7890 | 否 | ✓ 1404ms | 否 | ✓ 1448ms | ✓ 1160ms | http |
| 166.249.54.61:7234 | ✓ 650ms | ✓ 1728ms | 否 | ✓ 1500ms | ✓ 1545ms | http |
| 103.113.70.189:1081 | ✓ 751ms | ✓ 886ms | ✓ 130ms | ✓ 1101ms | ✓ 721ms | http |
| 45.140.147.155:1081 | ✓ 419ms | 否 | ✓ 1310ms | ✓ 1336ms | 否 | http |
| 121.230.9.252:1080 | ✓ 1389ms | ✓ 1831ms | ✓ 1253ms | ✓ 1545ms | ✓ 1175ms | http |
| 115.231.181.40:8128 | ✓ 1207ms | ✓ 1346ms | 否 | 否 | ✓ 1997ms | http |
| 103.171.183.255:1111 | 否 | 否 | ✓ 1987ms | ✓ 1639ms | ✓ 1797ms | http |
| 194.67.99.223:1080 | ✓ 1093ms | 否 | 否 | ✓ 1843ms | ✓ 1420ms | http |
| 160.238.65.5:3128 | ✓ 1082ms | 否 | ✓ 731ms | ✓ 1234ms | ✓ 1301ms | http |
| 45.136.130.187:8452 | ✓ 1490ms | ✓ 1623ms | ✓ 1128ms | ✓ 1159ms | ✓ 853ms | http |
| 45.136.130.182:8445 | ✓ 345ms | ✓ 958ms | ✓ 379ms | ✓ 1067ms | ✓ 809ms | http |
| 38.145.220.102:8445 | ✓ 343ms | ✓ 879ms | ✓ 623ms | ✓ 1335ms | ✓ 851ms | http |
| 38.145.220.96:8445 | ✓ 342ms | ✓ 879ms | ✓ 662ms | ✓ 1316ms | ✓ 837ms | http |
| 45.136.130.253:8444 | ✓ 348ms | ✓ 1148ms | ✓ 370ms | ✓ 1054ms | ✓ 1062ms | http |
| 38.145.203.76:8452 | ✓ 371ms | ✓ 905ms | ✓ 760ms | ✓ 1356ms | ✓ 724ms | http |
| 45.136.130.172:8451 | ✓ 701ms | ✓ 947ms | ✓ 583ms | ✓ 910ms | ✓ 944ms | http |
| 38.145.218.27:8446 | ✓ 1066ms | ✓ 1108ms | ✓ 965ms | ✓ 962ms | ✓ 1023ms | http |
| 38.145.220.82:8448 | ✓ 375ms | ✓ 988ms | ✓ 1008ms | ✓ 1261ms | ✓ 761ms | http |
| 38.145.220.35:8444 | ✓ 550ms | ✓ 904ms | ✓ 770ms | ✓ 1741ms | ✓ 833ms | http |
| 38.34.178.193:8446 | ✓ 435ms | ✓ 884ms | ✓ 614ms | ✓ 1853ms | ✓ 779ms | http |
| 195.123.209.48:3128 | ✓ 869ms | ✓ 1753ms | ✓ 1344ms | ✓ 1952ms | ✓ 1675ms | http |
| 38.34.179.86:8452 | 否 | 否 | ✓ 1431ms | ✓ 1726ms | ✓ 1209ms | http |
| 5.104.87.17:8051 | ✓ 1955ms | 否 | ✓ 1356ms | ✓ 1582ms | ✓ 1337ms | http |
| 160.238.65.2:3128 | ✓ 451ms | 否 | ✓ 473ms | ✓ 1406ms | 否 | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1053ms | ✓ 1386ms | ✓ 1093ms | http |
| 185.114.73.2:1080 | 否 | 否 | ✓ 1576ms | ✓ 1617ms | ✓ 1429ms | http |
| 45.140.147.155:1082 | ✓ 938ms | ✓ 1132ms | ✓ 919ms | 否 | ✓ 1148ms | http |
| 74.242.169.16:3128 | ✓ 1045ms | ✓ 1836ms | 否 | ✓ 1471ms | ✓ 1216ms | http |
| 210.223.44.230:3128 | ✓ 1445ms | ✓ 1209ms | ✓ 1139ms | 否 | ✓ 903ms | http |
| 89.208.106.138:10808 | ✓ 397ms | ✓ 1397ms | ✓ 1559ms | ✓ 1743ms | 否 | http |
| 38.34.179.79:8451 | ✓ 1141ms | ✓ 1248ms | ✓ 946ms | ✓ 1608ms | ✓ 1110ms | http |
| 83.219.250.8:62920 | ✓ 644ms | ✓ 1801ms | ✓ 1268ms | 否 | ✓ 1688ms | http |
| 180.125.216.109:8118 | ✓ 1071ms | ✓ 1309ms | ✓ 1040ms | 否 | 否 | http |
| 103.139.138.194:3128 | ✓ 1215ms | 否 | ✓ 1667ms | ✓ 1655ms | ✓ 1333ms | http |
| 186.148.180.46:999 | 否 | 否 | ✓ 1593ms | ✓ 1835ms | ✓ 1372ms | http |
| 223.16.170.103:80 | ✓ 1659ms | 否 | 否 | ✓ 1309ms | ✓ 1674ms | http |
| 45.136.198.40:3128 | ✓ 1151ms | ✓ 1484ms | ✓ 1491ms | ✓ 1856ms | ✓ 1667ms | http |
| 45.129.141.143:3128 | ✓ 1142ms | ✓ 1564ms | ✓ 1215ms | ✓ 1969ms | ✓ 1575ms | http |
| 45.186.6.104:3128 | ✓ 1397ms | ✓ 1757ms | ✓ 1689ms | 否 | 否 | http |
| 104.248.151.93:9090 | ✓ 908ms | 否 | ✓ 1696ms | ✓ 1290ms | ✓ 995ms | http |
| 62.113.119.14:8080 | ✓ 506ms | ✓ 1635ms | ✓ 521ms | ✓ 1495ms | ✓ 1140ms | http |
| 103.93.93.77:8050 | ✓ 1941ms | 否 | ✓ 1574ms | ✓ 1823ms | ✓ 1791ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1984ms | ✓ 1578ms | ✓ 1646ms | http |
| 152.70.84.108:8080 | ✓ 879ms | ✓ 1980ms | ✓ 1568ms | ✓ 1231ms | ✓ 869ms | http |
| 61.52.131.172:8443 | ✓ 881ms | ✓ 1096ms | ✓ 958ms | ✓ 1227ms | ✓ 926ms | http |
| 121.230.8.250:1080 | ✓ 1435ms | ✓ 1852ms | ✓ 1313ms | ✓ 1823ms | ✓ 1383ms | http |
| 45.140.147.82:1081 | ✓ 464ms | ✓ 1395ms | ✓ 1275ms | ✓ 1236ms | ✓ 1338ms | http |

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
