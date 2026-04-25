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

最后更新：2026-04-25 03:38:04 UTC（2026-04-25 11:38:04 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 183ms | 否 | ✓ 675ms | ✓ 1494ms | 否 | http |
| 130.61.174.200:1080 | ✓ 571ms | ✓ 1691ms | 否 | ✓ 1973ms | 否 | http |
| 45.140.147.82:1081 | ✓ 811ms | ✓ 1856ms | ✓ 1155ms | ✓ 1602ms | 否 | http |
| 47.84.76.30:1080 | ✓ 1054ms | ✓ 1770ms | ✓ 967ms | ✓ 1160ms | ✓ 921ms | http |
| 206.206.126.177:2412 | ✓ 996ms | ✓ 1869ms | ✓ 1018ms | ✓ 1091ms | ✓ 866ms | http |
| 45.88.0.115:3128 | ✓ 1200ms | ✓ 1623ms | ✓ 1073ms | 否 | ✓ 1443ms | http |
| 45.88.0.111:3128 | ✓ 1237ms | ✓ 1531ms | ✓ 1158ms | 否 | ✓ 1427ms | http |
| 45.88.0.98:3128 | ✓ 1164ms | ✓ 1612ms | ✓ 1100ms | 否 | ✓ 1434ms | http |
| 45.88.0.117:3128 | ✓ 1200ms | ✓ 1683ms | ✓ 1020ms | 否 | ✓ 1428ms | http |
| 38.180.192.119:3128 | ✓ 774ms | 否 | ✓ 1251ms | ✓ 1097ms | ✓ 838ms | http |
| 91.217.81.131:1080 | ✓ 1323ms | 否 | ✓ 1411ms | 否 | ✓ 1560ms | http |
| 212.58.132.5:8888 | ✓ 1332ms | 否 | ✓ 1379ms | ✓ 1961ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1656ms | 否 | ✓ 1721ms | ✓ 1866ms | ✓ 1539ms | http |
| 217.76.245.80:999 | ✓ 1142ms | ✓ 1784ms | ✓ 1213ms | ✓ 1448ms | 否 | http |
| 45.88.0.99:3128 | ✓ 500ms | ✓ 1277ms | ✓ 450ms | ✓ 1360ms | ✓ 993ms | http |
| 45.88.0.114:3128 | ✓ 501ms | ✓ 1437ms | ✓ 485ms | ✓ 1313ms | ✓ 992ms | http |
| 213.220.62.63:3128 | ✓ 492ms | ✓ 1475ms | ✓ 567ms | ✓ 1359ms | ✓ 1008ms | http |
| 45.88.0.116:3128 | ✓ 498ms | ✓ 1499ms | ✓ 538ms | ✓ 1346ms | ✓ 1028ms | http |
| 213.220.62.62:3128 | ✓ 516ms | ✓ 1946ms | ✓ 449ms | ✓ 1313ms | ✓ 1002ms | http |
| 45.88.0.113:3128 | ✓ 499ms | 否 | ✓ 455ms | ✓ 1342ms | ✓ 1000ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1442ms | ✓ 1269ms | 否 | ✓ 1530ms | http |
| 2.27.54.161:1080 | ✓ 703ms | 否 | ✓ 1812ms | ✓ 1848ms | ✓ 1425ms | http |
| 45.153.231.229:8080 | ✓ 919ms | ✓ 1854ms | ✓ 1738ms | 否 | ✓ 1828ms | http |
| 160.238.65.7:3128 | ✓ 1918ms | 否 | ✓ 839ms | ✓ 1839ms | ✓ 1424ms | http |
| 20.27.13.35:8561 | ✓ 675ms | ✓ 1421ms | ✓ 1077ms | ✓ 1515ms | ✓ 1567ms | http |
| 20.27.14.220:8561 | ✓ 671ms | ✓ 1505ms | ✓ 1042ms | ✓ 1513ms | ✓ 1529ms | http |
| 20.27.11.248:8561 | ✓ 670ms | ✓ 1368ms | ✓ 1135ms | ✓ 1521ms | 否 | http |
| 20.27.15.111:8561 | ✓ 676ms | ✓ 1397ms | ✓ 1102ms | ✓ 1520ms | ✓ 1559ms | http |
| 160.238.65.2:3128 | ✓ 1891ms | ✓ 1947ms | 否 | ✓ 1859ms | 否 | http |
| 45.140.147.82:1082 | ✓ 1617ms | 否 | ✓ 561ms | ✓ 1380ms | ✓ 1216ms | http |
| 120.92.212.16:8890 | ✓ 1307ms | ✓ 1965ms | ✓ 1207ms | ✓ 1520ms | 否 | http |
| 120.92.212.16:7890 | ✓ 998ms | ✓ 1970ms | 否 | ✓ 1807ms | ✓ 999ms | http |
| 34.71.229.255:3128 | ✓ 702ms | ✓ 1257ms | ✓ 1077ms | ✓ 1180ms | ✓ 1146ms | http |
| 1.231.81.166:3128 | ✓ 1445ms | ✓ 1320ms | ✓ 1103ms | ✓ 1072ms | ✓ 824ms | http |
| 47.95.231.180:8084 | ✓ 949ms | ✓ 1281ms | ✓ 1058ms | ✓ 1252ms | ✓ 1055ms | http |
| 46.101.95.183:8888 | ✓ 1020ms | 否 | ✓ 544ms | 否 | ✓ 1564ms | http |
| 52.53.211.45:13697 | ✓ 1484ms | 否 | ✓ 1081ms | ✓ 1953ms | ✓ 1511ms | http |
| 3.71.26.7:44736 | ✓ 1033ms | 否 | ✓ 967ms | 否 | ✓ 1235ms | http |
| 3.99.158.157:49993 | ✓ 1130ms | 否 | 否 | ✓ 1835ms | ✓ 1383ms | http |
| 51.95.13.205:324 | ✓ 1224ms | 否 | ✓ 1020ms | ✓ 1821ms | ✓ 1608ms | http |
| 52.210.57.38:39455 | ✓ 1062ms | 否 | ✓ 906ms | ✓ 1986ms | 否 | http |
| 78.12.252.87:50751 | ✓ 1081ms | 否 | ✓ 1794ms | ✓ 1985ms | ✓ 1707ms | http |
| 54.188.236.206:28025 | ✓ 1152ms | 否 | ✓ 1568ms | 否 | ✓ 1817ms | http |
| 16.62.123.236:7585 | ✓ 1037ms | 否 | 否 | ✓ 1686ms | ✓ 1447ms | http |
| 44.255.8.243:59355 | ✓ 1444ms | 否 | ✓ 1580ms | 否 | ✓ 1988ms | http |
| 108.130.79.116:3128 | ✓ 1853ms | 否 | ✓ 1725ms | 否 | ✓ 1372ms | http |
| 51.84.101.19:7764 | ✓ 1326ms | 否 | ✓ 1690ms | 否 | ✓ 1971ms | http |
| 161.35.181.96:999 | ✓ 684ms | ✓ 1323ms | 否 | ✓ 1064ms | ✓ 865ms | http |
| 35.225.22.61:80 | ✓ 339ms | 否 | ✓ 1053ms | ✓ 1316ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1819ms | ✓ 1015ms | ✓ 1834ms | ✓ 1093ms | http |
| 59.46.216.131:30001 | ✓ 1253ms | 否 | ✓ 1188ms | 否 | ✓ 1166ms | http |
| 62.234.206.73:3128 | 否 | 否 | ✓ 1057ms | ✓ 1396ms | ✓ 1085ms | http |
| 36.141.21.200:7890 | ✓ 1090ms | 否 | ✓ 1090ms | 否 | ✓ 1694ms | http |
| 218.108.131.186:17890 | ✓ 1229ms | 否 | ✓ 965ms | ✓ 1197ms | ✓ 969ms | http |
| 121.230.9.203:1080 | 否 | ✓ 1610ms | ✓ 1492ms | 否 | ✓ 1702ms | http |
| 89.208.106.138:10808 | ✓ 694ms | ✓ 1809ms | ✓ 1578ms | ✓ 1619ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1119ms | 否 | ✓ 1783ms | 否 | ✓ 1512ms | http |
| 177.93.132.244:3128 | ✓ 1553ms | 否 | ✓ 922ms | 否 | ✓ 1843ms | http |
| 8.219.195.129:1080 | ✓ 833ms | 否 | ✓ 917ms | ✓ 1177ms | ✓ 921ms | http |
| 101.32.244.83:8080 | ✓ 1222ms | 否 | ✓ 1086ms | ✓ 1368ms | ✓ 1343ms | http |
| 121.43.196.213:8222 | ✓ 1009ms | ✓ 1155ms | ✓ 978ms | ✓ 1153ms | ✓ 962ms | http |
| 121.43.196.210:8222 | ✓ 1034ms | ✓ 1154ms | ✓ 946ms | ✓ 1235ms | ✓ 987ms | http |
| 114.55.226.123:10086 | ✓ 1146ms | ✓ 1791ms | ✓ 1115ms | ✓ 1426ms | ✓ 1264ms | http |
| 180.103.19.151:1080 | 否 | ✓ 1662ms | ✓ 1899ms | 否 | ✓ 1276ms | http |
| 34.246.183.20:9023 | ✓ 1090ms | 否 | ✓ 1688ms | 否 | ✓ 1455ms | http |
| 3.19.213.118:59531 | ✓ 1625ms | 否 | ✓ 1810ms | ✓ 1848ms | ✓ 1978ms | http |
| 15.160.116.45:13815 | ✓ 1048ms | 否 | ✓ 1721ms | 否 | ✓ 1509ms | http |
| 3.14.146.121:15072 | ✓ 1895ms | 否 | ✓ 1730ms | ✓ 1999ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1372ms | 否 | ✓ 1260ms | ✓ 1697ms | ✓ 1329ms | http |
| 91.193.240.157:9877 | ✓ 1299ms | 否 | ✓ 1108ms | 否 | ✓ 1824ms | http |
| 65.109.199.54:1080 | ✓ 773ms | 否 | ✓ 1391ms | ✓ 1782ms | ✓ 1337ms | http |
| 166.88.61.54:8000 | ✓ 1236ms | ✓ 1329ms | ✓ 1289ms | ✓ 1216ms | ✓ 769ms | http |
| 183.232.248.73:7890 | ✓ 1092ms | 否 | ✓ 1185ms | 否 | ✓ 1077ms | http |
| 94.131.118.129:1081 | ✓ 1634ms | 否 | ✓ 1123ms | 否 | ✓ 1708ms | http |
| 60.249.94.208:3128 | ✓ 846ms | ✓ 1042ms | ✓ 962ms | ✓ 1131ms | ✓ 995ms | http |
| 211.95.152.50:45046 | 否 | ✓ 1502ms | ✓ 1464ms | ✓ 1746ms | 否 | http |
| 34.246.223.187:49924 | ✓ 1756ms | 否 | ✓ 1739ms | ✓ 1934ms | 否 | http |
| 64.181.240.152:3128 | ✓ 1487ms | 否 | 否 | ✓ 1192ms | ✓ 876ms | http |
| 162.240.154.26:3128 | ✓ 1502ms | 否 | ✓ 1448ms | ✓ 1372ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1646ms | 否 | ✓ 1288ms | ✓ 1853ms | 否 | http |
| 218.60.0.214:80 | ✓ 1116ms | ✓ 1415ms | ✓ 1146ms | 否 | ✓ 1138ms | http |
| 45.186.6.104:3128 | ✓ 1666ms | ✓ 1906ms | ✓ 1919ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 940ms | ✓ 1260ms | ✓ 1048ms | ✓ 1663ms | 否 | http |
| 74.50.96.247:8888 | 否 | ✓ 1366ms | ✓ 823ms | ✓ 1056ms | ✓ 1195ms | http |
| 16.18.37.186:6085 | ✓ 669ms | 否 | ✓ 1478ms | 否 | ✓ 1402ms | http |
| 168.144.75.9:3128 | ✓ 1873ms | 否 | ✓ 1759ms | ✓ 1985ms | ✓ 1940ms | http |

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
