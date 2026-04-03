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

最后更新：2026-04-03 10:55:07 UTC（2026-04-03 18:55:07 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 461ms | 否 | ✓ 1412ms | ✓ 1638ms | ✓ 1294ms | http |
| 203.80.138.81:50000 | ✓ 1319ms | ✓ 1401ms | ✓ 1220ms | ✓ 1208ms | ✓ 1126ms | http |
| 1.231.81.166:3128 | ✓ 1093ms | ✓ 1209ms | ✓ 1833ms | ✓ 1476ms | ✓ 1084ms | http |
| 147.161.210.140:8800 | ✓ 938ms | 否 | ✓ 1000ms | ✓ 1307ms | 否 | http |
| 167.103.115.102:8800 | ✓ 1281ms | ✓ 1906ms | ✓ 1156ms | ✓ 1348ms | ✓ 1352ms | http |
| 167.103.34.108:8800 | ✓ 1556ms | 否 | ✓ 1558ms | 否 | ✓ 1603ms | http |
| 113.160.132.26:8080 | ✓ 1817ms | 否 | ✓ 1164ms | ✓ 1843ms | 否 | http |
| 159.223.71.162:8080 | ✓ 1580ms | 否 | ✓ 1635ms | ✓ 1276ms | ✓ 1037ms | http |
| 95.213.217.168:52004 | ✓ 808ms | ✓ 1532ms | 否 | ✓ 1973ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1557ms | 否 | ✓ 1724ms | ✓ 1837ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1577ms | 否 | ✓ 1933ms | ✓ 1798ms | 否 | http |
| 217.76.245.80:999 | ✓ 767ms | 否 | ✓ 1049ms | ✓ 1486ms | 否 | http |
| 45.167.125.21:999 | ✓ 1019ms | ✓ 1849ms | 否 | ✓ 1743ms | ✓ 1448ms | http |
| 206.81.27.105:3128 | ✓ 733ms | 否 | ✓ 1712ms | 否 | ✓ 1709ms | http |
| 59.46.216.131:30001 | ✓ 1211ms | ✓ 1555ms | ✓ 1205ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 928ms | ✓ 1309ms | 否 | ✓ 954ms | 否 | http |
| 101.43.127.100:8877 | ✓ 996ms | ✓ 1381ms | ✓ 1021ms | ✓ 1389ms | ✓ 1084ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1392ms | ✓ 1634ms | 否 | ✓ 1220ms | http |
| 177.234.217.88:999 | ✓ 1446ms | 否 | ✓ 1802ms | ✓ 1984ms | ✓ 1801ms | http |
| 103.113.70.189:1081 | ✓ 350ms | ✓ 959ms | 否 | ✓ 1251ms | ✓ 903ms | http |
| 38.34.179.150:8449 | ✓ 926ms | ✓ 965ms | ✓ 708ms | ✓ 1561ms | ✓ 870ms | http |
| 38.34.179.106:8445 | ✓ 925ms | ✓ 1039ms | ✓ 639ms | ✓ 1531ms | ✓ 893ms | http |
| 38.145.220.72:8445 | ✓ 904ms | ✓ 1252ms | 否 | ✓ 1013ms | ✓ 1301ms | http |
| 38.34.178.154:8445 | ✓ 444ms | ✓ 970ms | ✓ 304ms | ✓ 937ms | ✓ 808ms | http |
| 38.145.220.82:8448 | ✓ 621ms | ✓ 894ms | ✓ 323ms | ✓ 1086ms | ✓ 950ms | http |
| 38.145.208.215:8444 | ✓ 912ms | 否 | ✓ 316ms | ✓ 968ms | ✓ 1440ms | http |
| 38.145.208.220:8448 | ✓ 1100ms | ✓ 1189ms | ✓ 1705ms | ✓ 946ms | ✓ 1070ms | http |
| 38.145.208.229:8453 | ✓ 326ms | ✓ 975ms | ✓ 914ms | ✓ 1499ms | ✓ 756ms | http |
| 38.145.220.39:8452 | ✓ 1218ms | ✓ 1121ms | ✓ 1725ms | ✓ 1155ms | ✓ 1078ms | http |
| 46.39.105.157:8080 | ✓ 1267ms | 否 | ✓ 1704ms | ✓ 1360ms | 否 | http |
| 38.145.218.227:8445 | ✓ 842ms | 否 | ✓ 362ms | ✓ 927ms | ✓ 686ms | http |
| 38.145.208.207:8445 | ✓ 836ms | ✓ 911ms | ✓ 1438ms | ✓ 1391ms | ✓ 714ms | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 996ms | ✓ 1617ms | ✓ 1174ms | http |
| 38.34.179.105:8449 | ✓ 1257ms | 否 | ✓ 717ms | ✓ 1644ms | ✓ 761ms | http |
| 208.87.243.199:7878 | ✓ 584ms | 否 | ✓ 1172ms | ✓ 968ms | ✓ 724ms | http |
| 72.11.151.159:6005 | ✓ 406ms | 否 | ✓ 881ms | ✓ 1045ms | ✓ 797ms | http |
| 38.34.179.40:8446 | ✓ 1467ms | 否 | ✓ 328ms | ✓ 962ms | ✓ 1493ms | http |
| 38.145.203.35:8450 | ✓ 976ms | ✓ 1316ms | ✓ 1765ms | ✓ 1036ms | 否 | http |
| 38.145.203.41:8453 | ✓ 968ms | ✓ 1558ms | ✓ 1513ms | ✓ 1029ms | ✓ 1959ms | http |
| 38.34.179.192:8446 | ✓ 1724ms | 否 | ✓ 1155ms | ✓ 950ms | ✓ 1492ms | http |
| 38.34.179.79:8451 | ✓ 687ms | ✓ 1643ms | ✓ 1818ms | ✓ 1388ms | 否 | http |
| 45.136.130.169:8444 | ✓ 583ms | 否 | ✓ 1772ms | ✓ 1394ms | 否 | http |
| 38.145.203.132:8450 | ✓ 978ms | ✓ 1791ms | 否 | ✓ 1388ms | ✓ 1997ms | http |
| 45.136.131.61:8444 | ✓ 1135ms | ✓ 1142ms | 否 | ✓ 1494ms | ✓ 1498ms | http |
| 45.136.130.180:8452 | ✓ 1172ms | ✓ 1205ms | ✓ 1870ms | ✓ 1540ms | ✓ 1644ms | http |
| 38.145.218.216:8449 | ✓ 1739ms | ✓ 1705ms | 否 | ✓ 955ms | ✓ 1921ms | http |
| 38.34.183.8:8450 | 否 | ✓ 1796ms | ✓ 712ms | 否 | ✓ 1184ms | http |
| 38.145.208.242:8451 | ✓ 904ms | 否 | ✓ 1368ms | ✓ 1537ms | 否 | http |
| 38.34.179.50:8444 | ✓ 723ms | 否 | ✓ 1517ms | 否 | ✓ 1768ms | http |
| 45.136.131.49:8445 | ✓ 837ms | 否 | ✓ 1243ms | ✓ 1454ms | 否 | http |
| 38.145.218.87:8451 | ✓ 878ms | ✓ 1072ms | ✓ 1478ms | 否 | 否 | http |
| 45.136.130.247:8448 | ✓ 749ms | ✓ 1868ms | ✓ 1415ms | ✓ 1440ms | ✓ 924ms | http |
| 38.145.218.101:8450 | ✓ 1076ms | ✓ 1473ms | ✓ 1886ms | ✓ 1332ms | ✓ 969ms | http |
| 210.223.44.230:3128 | ✓ 1861ms | ✓ 1347ms | ✓ 841ms | ✓ 1150ms | ✓ 1968ms | http |
| 101.32.244.83:8080 | ✓ 1437ms | 否 | ✓ 1191ms | ✓ 1496ms | ✓ 1592ms | http |
| 121.43.196.213:8222 | ✓ 1086ms | ✓ 1310ms | ✓ 1104ms | ✓ 1323ms | ✓ 1096ms | http |
| 121.43.196.210:8222 | ✓ 1109ms | ✓ 1263ms | ✓ 1122ms | ✓ 1314ms | ✓ 1108ms | http |
| 114.55.226.123:10086 | ✓ 1271ms | ✓ 1938ms | ✓ 1222ms | ✓ 1510ms | ✓ 1256ms | http |
| 122.45.51.68:11113 | ✓ 1983ms | 否 | ✓ 1351ms | ✓ 1151ms | 否 | http |
| 195.123.209.48:3128 | ✓ 1061ms | ✓ 1451ms | ✓ 1142ms | ✓ 1871ms | ✓ 1728ms | http |
| 34.96.238.40:8080 | ✓ 1472ms | ✓ 1651ms | ✓ 1394ms | ✓ 1359ms | ✓ 1445ms | http |
| 165.232.146.249:3128 | ✓ 613ms | ✓ 1421ms | ✓ 1945ms | ✓ 1425ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1286ms | 否 | ✓ 1918ms | ✓ 1733ms | 否 | http |
| 150.5.166.194:1080 | ✓ 958ms | 否 | ✓ 958ms | ✓ 1110ms | ✓ 900ms | http |
| 139.59.222.40:3128 | ✓ 890ms | 否 | ✓ 1013ms | ✓ 1274ms | ✓ 996ms | http |
| 38.242.212.16:3128 | ✓ 900ms | 否 | ✓ 1583ms | ✓ 1815ms | ✓ 1676ms | http |
| 5.253.43.103:3128 | ✓ 1273ms | 否 | ✓ 1410ms | 否 | ✓ 1385ms | http |
| 194.147.90.23:3128 | ✓ 1136ms | 否 | ✓ 678ms | 否 | ✓ 1817ms | http |
| 165.22.63.252:3128 | ✓ 882ms | 否 | ✓ 892ms | ✓ 1260ms | ✓ 1008ms | http |
| 45.140.147.82:1081 | ✓ 987ms | ✓ 1243ms | ✓ 1302ms | ✓ 1785ms | ✓ 1272ms | http |
| 104.248.243.244:3128 | ✓ 931ms | ✓ 1691ms | ✓ 1771ms | ✓ 1696ms | ✓ 1304ms | http |
| 185.40.7.206:3128 | ✓ 1027ms | ✓ 1552ms | ✓ 1406ms | 否 | 否 | http |
| 199.68.217.2:3128 | ✓ 1133ms | 否 | 否 | ✓ 1613ms | ✓ 1233ms | http |
| 176.100.39.18:8080 | ✓ 999ms | 否 | ✓ 979ms | 否 | ✓ 1722ms | http |
| 49.7.179.70:3333 | ✓ 1270ms | 否 | ✓ 1472ms | 否 | ✓ 1256ms | http |
| 140.245.66.105:8081 | ✓ 973ms | ✓ 1480ms | ✓ 1340ms | ✓ 1730ms | ✓ 1112ms | http |
| 150.249.255.91:3128 | ✓ 660ms | ✓ 1183ms | 否 | 否 | ✓ 822ms | http |
| 1.225.116.115:1080 | ✓ 1031ms | ✓ 1746ms | ✓ 1064ms | ✓ 1199ms | ✓ 888ms | http |
| 20.118.221.52:3128 | ✓ 846ms | ✓ 1389ms | ✓ 954ms | ✓ 1138ms | 否 | http |
| 51.79.71.106:8080 | ✓ 540ms | 否 | ✓ 1795ms | ✓ 1486ms | ✓ 1109ms | http |
| 125.76.214.178:8091 | 否 | ✓ 1440ms | ✓ 1121ms | ✓ 1517ms | ✓ 1276ms | http |
| 38.145.220.60:8447 | ✓ 1243ms | 否 | ✓ 1501ms | ✓ 1179ms | 否 | http |
| 8.219.97.248:80 | ✓ 1751ms | 否 | ✓ 1479ms | ✓ 1832ms | 否 | http |

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
