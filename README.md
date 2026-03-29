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

最后更新：2026-03-29 20:42:50 UTC（2026-03-30 04:42:50 UTC+8）

**代理总数：122**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 122 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 122 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 39.185.46.193:5911 | ✓ 753ms | ✓ 923ms | ✓ 843ms | ✓ 982ms | ✓ 758ms | http |
| 147.161.239.240:8800 | ✓ 610ms | ✓ 1676ms | ✓ 1271ms | ✓ 1316ms | ✓ 1148ms | http |
| 103.84.95.54:7890 | 否 | ✓ 1844ms | ✓ 757ms | ✓ 997ms | ✓ 1024ms | http |
| 147.161.210.140:8800 | ✓ 1841ms | ✓ 1213ms | ✓ 878ms | ✓ 1258ms | ✓ 1399ms | http |
| 95.213.217.168:52004 | ✓ 712ms | ✓ 1704ms | ✓ 1185ms | 否 | ✓ 1608ms | http |
| 167.103.115.102:8800 | ✓ 1925ms | ✓ 1738ms | ✓ 1082ms | ✓ 1358ms | ✓ 1090ms | http |
| 167.103.34.108:8800 | ✓ 1976ms | 否 | ✓ 1469ms | ✓ 1459ms | ✓ 1564ms | http |
| 222.228.171.92:8080 | ✓ 1923ms | 否 | ✓ 1327ms | ✓ 1742ms | 否 | http |
| 35.225.22.61:80 | ✓ 610ms | ✓ 1273ms | ✓ 446ms | ✓ 1185ms | ✓ 866ms | http |
| 43.99.54.236:5555 | ✓ 785ms | ✓ 1066ms | ✓ 738ms | ✓ 935ms | ✓ 743ms | http |
| 183.249.5.117:22222 | ✓ 831ms | ✓ 957ms | ✓ 740ms | ✓ 1041ms | ✓ 804ms | http |
| 222.184.48.242:22222 | ✓ 987ms | ✓ 1210ms | ✓ 999ms | ✓ 1311ms | ✓ 979ms | http |
| 120.92.212.16:8890 | ✓ 1045ms | ✓ 1326ms | ✓ 1716ms | ✓ 1363ms | ✓ 1097ms | http |
| 167.103.144.127:8800 | ✓ 1326ms | 否 | ✓ 1225ms | ✓ 1717ms | ✓ 1527ms | http |
| 8.219.97.248:80 | ✓ 1229ms | 否 | ✓ 1337ms | 否 | ✓ 1357ms | http |
| 5.104.87.17:8051 | ✓ 1570ms | 否 | ✓ 1658ms | 否 | ✓ 1487ms | http |
| 167.103.31.122:8800 | ✓ 1755ms | 否 | ✓ 1721ms | 否 | ✓ 1935ms | http |
| 180.250.219.58:53281 | ✓ 1843ms | 否 | ✓ 1694ms | ✓ 1999ms | 否 | http |
| 208.87.243.199:7878 | ✓ 1014ms | ✓ 951ms | ✓ 1075ms | ✓ 1065ms | ✓ 694ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1214ms | ✓ 1352ms | ✓ 1150ms | http |
| 42.96.16.158:1311 | ✓ 1520ms | 否 | ✓ 975ms | ✓ 1301ms | ✓ 988ms | http |
| 115.231.181.40:8128 | ✓ 1961ms | ✓ 1316ms | ✓ 1038ms | ✓ 1276ms | ✓ 1086ms | http |
| 45.12.151.226:2829 | ✓ 737ms | ✓ 1754ms | ✓ 736ms | ✓ 1685ms | ✓ 1386ms | http |
| 219.117.204.211:7799 | ✓ 1828ms | ✓ 1767ms | ✓ 895ms | ✓ 1187ms | ✓ 1691ms | http |
| 222.184.48.251:22222 | ✓ 1034ms | ✓ 1297ms | ✓ 947ms | ✓ 1262ms | ✓ 1651ms | http |
| 120.92.212.16:7890 | ✓ 1031ms | ✓ 1317ms | ✓ 985ms | ✓ 1296ms | ✓ 1021ms | http |
| 38.34.179.67:8451 | ✓ 587ms | ✓ 1063ms | ✓ 1460ms | 否 | 否 | http |
| 38.145.220.102:8453 | ✓ 656ms | ✓ 1206ms | ✓ 1084ms | ✓ 813ms | 否 | http |
| 45.136.131.38:8445 | ✓ 201ms | ✓ 821ms | ✓ 400ms | ✓ 1146ms | ✓ 1032ms | http |
| 101.43.127.100:8877 | ✓ 1115ms | ✓ 1214ms | ✓ 1173ms | ✓ 1212ms | ✓ 922ms | http |
| 5.102.109.41:999 | ✓ 363ms | ✓ 1051ms | ✓ 1945ms | ✓ 1423ms | ✓ 1115ms | http |
| 177.234.217.88:999 | ✓ 1885ms | ✓ 1656ms | ✓ 1749ms | ✓ 1894ms | ✓ 1727ms | http |
| 116.80.65.80:3172 | ✓ 1899ms | 否 | ✓ 1986ms | 否 | ✓ 1978ms | http |
| 38.145.203.86:8452 | ✓ 832ms | ✓ 1925ms | ✓ 202ms | ✓ 814ms | ✓ 752ms | http |
| 45.136.131.65:8452 | ✓ 224ms | ✓ 827ms | ✓ 374ms | ✓ 970ms | ✓ 709ms | http |
| 38.34.179.37:8451 | ✓ 480ms | ✓ 1106ms | ✓ 402ms | ✓ 1107ms | ✓ 792ms | http |
| 38.34.179.39:8452 | ✓ 743ms | ✓ 1537ms | ✓ 751ms | ✓ 1109ms | ✓ 845ms | http |
| 38.145.220.11:8446 | ✓ 494ms | ✓ 1286ms | ✓ 1112ms | ✓ 836ms | ✓ 765ms | http |
| 38.34.183.130:8452 | ✓ 820ms | ✓ 1311ms | ✓ 822ms | ✓ 1002ms | ✓ 1270ms | http |
| 38.34.179.64:8451 | ✓ 896ms | ✓ 1210ms | ✓ 1090ms | 否 | ✓ 930ms | http |
| 38.34.179.88:8446 | ✓ 1960ms | ✓ 1326ms | ✓ 1424ms | ✓ 1962ms | 否 | http |
| 183.249.5.105:22222 | ✓ 812ms | ✓ 933ms | ✓ 772ms | ✓ 1358ms | ✓ 783ms | http |
| 183.249.5.110:22222 | ✓ 975ms | ✓ 1224ms | ✓ 769ms | ✓ 1046ms | ✓ 777ms | http |
| 222.184.48.252:22222 | ✓ 1541ms | ✓ 1675ms | ✓ 1196ms | ✓ 1307ms | ✓ 1140ms | http |
| 181.41.201.85:3128 | ✓ 1888ms | 否 | ✓ 1982ms | 否 | ✓ 1897ms | http |
| 160.238.65.2:3128 | ✓ 1696ms | 否 | ✓ 1545ms | 否 | ✓ 1406ms | http |
| 160.238.65.3:3128 | 否 | ✓ 1553ms | ✓ 1677ms | 否 | ✓ 1373ms | http |
| 103.18.78.250:1111 | ✓ 1789ms | 否 | ✓ 1570ms | ✓ 1485ms | ✓ 1466ms | http |
| 38.34.179.101:8446 | 否 | ✓ 1150ms | ✓ 816ms | 否 | ✓ 1090ms | http |
| 38.34.179.86:8452 | ✓ 625ms | ✓ 933ms | ✓ 1258ms | ✓ 1302ms | ✓ 905ms | http |
| 38.145.218.228:8447 | ✓ 1106ms | 否 | ✓ 416ms | ✓ 1196ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1241ms | 否 | ✓ 1773ms | 否 | ✓ 1176ms | http |
| 38.34.179.25:8444 | ✓ 397ms | ✓ 902ms | ✓ 541ms | ✓ 1199ms | ✓ 1188ms | http |
| 38.34.179.23:8444 | ✓ 405ms | ✓ 951ms | ✓ 565ms | ✓ 1183ms | ✓ 1156ms | http |
| 210.223.44.230:3128 | ✓ 1394ms | ✓ 966ms | 否 | ✓ 979ms | ✓ 756ms | http |
| 38.34.179.71:8444 | ✓ 403ms | ✓ 917ms | ✓ 1161ms | ✓ 1789ms | ✓ 872ms | http |
| 38.34.179.66:8444 | ✓ 393ms | ✓ 908ms | ✓ 1072ms | ✓ 1960ms | ✓ 854ms | http |
| 38.34.179.14:8450 | ✓ 1697ms | ✓ 907ms | ✓ 978ms | 否 | ✓ 711ms | http |
| 194.59.204.87:9080 | ✓ 1125ms | ✓ 1710ms | ✓ 1795ms | 否 | 否 | http |
| 59.8.203.55:80 | ✓ 1678ms | ✓ 1265ms | ✓ 1718ms | ✓ 1070ms | ✓ 838ms | http |
| 121.230.9.252:1080 | ✓ 1184ms | ✓ 1455ms | ✓ 1533ms | ✓ 1632ms | ✓ 1204ms | http |
| 209.126.84.232:8888 | ✓ 1828ms | 否 | ✓ 1903ms | ✓ 1728ms | ✓ 1565ms | http |
| 45.136.131.47:8445 | ✓ 206ms | ✓ 862ms | ✓ 407ms | 否 | 否 | http |
| 38.145.218.10:8450 | ✓ 1429ms | ✓ 796ms | ✓ 1078ms | 否 | ✓ 853ms | http |
| 45.136.131.35:8443 | ✓ 577ms | ✓ 971ms | 否 | ✓ 1424ms | ✓ 1027ms | http |
| 45.136.131.31:8443 | ✓ 613ms | ✓ 972ms | 否 | ✓ 1463ms | ✓ 996ms | http |
| 160.238.65.9:3128 | ✓ 646ms | 否 | ✓ 520ms | ✓ 1710ms | ✓ 1795ms | http |
| 45.136.131.32:8443 | ✓ 843ms | ✓ 761ms | ✓ 907ms | 否 | 否 | http |
| 38.34.179.173:8450 | 否 | ✓ 1022ms | ✓ 1246ms | ✓ 1365ms | ✓ 792ms | http |
| 38.145.203.135:8444 | 否 | 否 | ✓ 227ms | ✓ 806ms | ✓ 923ms | http |
| 38.34.179.16:8451 | 否 | 否 | ✓ 334ms | ✓ 1016ms | ✓ 912ms | http |
| 38.145.203.76:8446 | 否 | ✓ 1001ms | ✓ 1144ms | ✓ 1433ms | ✓ 773ms | http |
| 38.145.208.244:8448 | 否 | ✓ 1827ms | ✓ 323ms | ✓ 819ms | ✓ 1020ms | http |
| 38.145.218.227:8451 | ✓ 256ms | ✓ 747ms | ✓ 379ms | ✓ 806ms | ✓ 632ms | http |
| 38.145.208.169:8446 | ✓ 274ms | ✓ 838ms | ✓ 261ms | ✓ 805ms | ✓ 642ms | http |
| 38.145.208.242:8444 | ✓ 223ms | ✓ 809ms | ✓ 199ms | ✓ 898ms | ✓ 666ms | http |
| 38.145.208.217:8449 | ✓ 360ms | ✓ 821ms | ✓ 942ms | ✓ 1144ms | ✓ 980ms | http |
| 38.34.179.24:8453 | ✓ 424ms | ✓ 941ms | ✓ 553ms | ✓ 1252ms | ✓ 737ms | http |
| 38.145.208.191:8446 | ✓ 485ms | ✓ 870ms | ✓ 280ms | ✓ 1138ms | 否 | http |
| 38.145.203.111:8449 | ✓ 1067ms | ✓ 1081ms | ✓ 243ms | ✓ 844ms | ✓ 939ms | http |
| 38.34.179.66:8446 | ✓ 557ms | ✓ 1594ms | ✓ 380ms | ✓ 967ms | ✓ 756ms | http |
| 38.34.179.29:8449 | ✓ 537ms | ✓ 1281ms | ✓ 828ms | ✓ 935ms | ✓ 770ms | http |
| 38.34.179.62:8453 | ✓ 369ms | ✓ 1115ms | ✓ 947ms | ✓ 980ms | ✓ 888ms | http |
| 38.34.179.100:8452 | ✓ 606ms | 否 | ✓ 593ms | ✓ 999ms | ✓ 1499ms | http |
| 45.8.157.38:3128 | ✓ 1242ms | 否 | ✓ 946ms | ✓ 1411ms | ✓ 1268ms | http |
| 38.34.179.13:8445 | ✓ 1250ms | ✓ 1042ms | 否 | ✓ 953ms | ✓ 964ms | http |
| 38.34.179.92:8448 | ✓ 517ms | ✓ 1517ms | ✓ 1655ms | ✓ 1318ms | ✓ 1608ms | http |
| 38.34.179.20:8452 | ✓ 1250ms | ✓ 1004ms | 否 | ✓ 1000ms | ✓ 1037ms | http |
| 195.123.209.48:3128 | ✓ 931ms | ✓ 1587ms | ✓ 1493ms | 否 | ✓ 1789ms | http |
| 38.34.179.35:8448 | ✓ 1200ms | ✓ 967ms | ✓ 1901ms | ✓ 1922ms | ✓ 815ms | http |
| 91.233.223.147:3128 | ✓ 940ms | 否 | ✓ 1662ms | ✓ 1953ms | ✓ 1522ms | http |
| 121.230.8.250:1080 | ✓ 1143ms | ✓ 1673ms | ✓ 1189ms | ✓ 1662ms | ✓ 1504ms | http |
| 38.145.220.198:8451 | ✓ 1374ms | ✓ 1611ms | ✓ 1057ms | 否 | ✓ 888ms | http |
| 62.171.161.88:2018 | ✓ 1980ms | ✓ 1606ms | ✓ 604ms | ✓ 1596ms | ✓ 1586ms | http |
| 38.145.208.237:8445 | ✓ 878ms | ✓ 1599ms | 否 | ✓ 1611ms | 否 | http |
| 65.108.203.35:18080 | 否 | 否 | ✓ 1933ms | ✓ 1958ms | ✓ 1936ms | http |
| 59.46.216.131:30001 | ✓ 1043ms | 否 | ✓ 1140ms | 否 | ✓ 1089ms | http |
| 106.75.15.167:7890 | ✓ 1522ms | ✓ 1509ms | ✓ 1098ms | 否 | 否 | http |
| 180.125.216.109:8118 | 否 | ✓ 1168ms | ✓ 945ms | ✓ 1362ms | 否 | http |
| 116.80.65.78:3172 | ✓ 1616ms | 否 | ✓ 1566ms | 否 | ✓ 1728ms | http |

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
