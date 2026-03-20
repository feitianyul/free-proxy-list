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

最后更新：2026-03-20 18:31:34 UTC（2026-03-21 02:31:34 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 371ms | ✓ 1805ms | ✓ 815ms | 否 | ✓ 1158ms | http |
| 178.156.187.185:10001 | ✓ 255ms | 否 | ✓ 1066ms | ✓ 1443ms | ✓ 1578ms | http |
| 147.161.239.240:8800 | ✓ 1025ms | 否 | ✓ 1146ms | 否 | ✓ 1407ms | http |
| 167.103.34.108:8800 | ✓ 1779ms | 否 | ✓ 1658ms | ✓ 1926ms | 否 | http |
| 174.138.24.77:1080 | ✓ 1401ms | 否 | ✓ 1812ms | 否 | ✓ 1290ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 907ms | ✓ 1431ms | ✓ 1381ms | http |
| 45.167.124.52:8080 | ✓ 740ms | ✓ 1738ms | ✓ 1143ms | ✓ 1955ms | ✓ 1420ms | http |
| 147.161.210.140:8800 | ✓ 1746ms | 否 | ✓ 980ms | 否 | ✓ 1178ms | http |
| 137.220.151.110:6005 | ✓ 1656ms | 否 | ✓ 1070ms | ✓ 1438ms | ✓ 1120ms | http |
| 113.160.132.26:8080 | ✓ 1721ms | 否 | ✓ 1201ms | ✓ 1716ms | ✓ 1123ms | http |
| 38.145.220.11:8445 | 否 | 否 | ✓ 1785ms | ✓ 908ms | ✓ 1163ms | http |
| 167.103.31.122:8800 | ✓ 1732ms | 否 | ✓ 1675ms | ✓ 1923ms | 否 | http |
| 137.220.150.104:6005 | ✓ 1330ms | 否 | ✓ 1437ms | 否 | ✓ 1363ms | http |
| 133.242.138.34:8100 | ✓ 1812ms | 否 | 否 | ✓ 1905ms | ✓ 1783ms | http |
| 137.220.150.22:6005 | ✓ 952ms | 否 | ✓ 935ms | ✓ 1379ms | ✓ 1132ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 774ms | ✓ 1082ms | ✓ 808ms | http |
| 139.159.99.242:8080 | 否 | 否 | ✓ 1005ms | ✓ 1309ms | ✓ 978ms | http |
| 47.77.193.180:1080 | ✓ 1082ms | ✓ 1614ms | ✓ 411ms | ✓ 915ms | ✓ 697ms | http |
| 45.136.131.60:8452 | ✓ 1056ms | ✓ 1529ms | ✓ 520ms | ✓ 996ms | ✓ 706ms | http |
| 91.238.105.64:2024 | ✓ 982ms | ✓ 1895ms | ✓ 1555ms | 否 | ✓ 1606ms | http |
| 91.233.223.147:3128 | ✓ 1652ms | 否 | ✓ 1495ms | ✓ 1865ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1081ms | ✓ 1367ms | ✓ 1135ms | ✓ 1390ms | ✓ 1071ms | http |
| 84.247.149.172:3128 | ✓ 1128ms | 否 | ✓ 1526ms | ✓ 1328ms | ✓ 1045ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 1155ms | ✓ 1420ms | ✓ 1574ms | http |
| 142.171.224.229:7890 | ✓ 388ms | ✓ 923ms | ✓ 883ms | ✓ 905ms | ✓ 666ms | http |
| 38.145.218.227:8445 | ✓ 897ms | 否 | ✓ 997ms | ✓ 923ms | ✓ 1084ms | http |
| 38.34.183.225:8449 | 否 | 否 | ✓ 1419ms | ✓ 1429ms | ✓ 1995ms | http |
| 185.114.73.2:1080 | 否 | ✓ 1555ms | ✓ 1314ms | ✓ 1703ms | ✓ 1545ms | http |
| 38.34.179.162:8451 | 否 | 否 | ✓ 1156ms | ✓ 1635ms | ✓ 1848ms | http |
| 88.80.150.82:8080 | ✓ 1900ms | 否 | ✓ 1751ms | ✓ 1901ms | ✓ 1558ms | https |
| 115.231.181.40:8128 | 否 | ✓ 1402ms | 否 | ✓ 1398ms | ✓ 1103ms | http |
| 144.31.25.69:21064 | ✓ 809ms | 否 | ✓ 988ms | 否 | ✓ 1965ms | http |
| 120.92.212.16:7890 | ✓ 1204ms | ✓ 1684ms | ✓ 1152ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 1757ms | ✓ 1628ms | ✓ 1252ms | ✓ 1342ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1236ms | ✓ 1730ms | ✓ 1602ms | ✓ 1750ms | 否 | http |
| 172.212.68.37:3128 | ✓ 347ms | ✓ 1471ms | ✓ 1284ms | ✓ 1211ms | 否 | http |
| 185.115.74.185:8080 | ✓ 1345ms | ✓ 1565ms | ✓ 1509ms | 否 | 否 | http |
| 38.34.179.49:8450 | ✓ 349ms | ✓ 1535ms | ✓ 1179ms | ✓ 1408ms | ✓ 718ms | http |
| 38.145.220.33:8448 | ✓ 496ms | ✓ 1100ms | ✓ 490ms | ✓ 1664ms | ✓ 716ms | http |
| 45.136.131.39:8451 | ✓ 920ms | 否 | ✓ 1933ms | 否 | ✓ 707ms | http |
| 38.34.179.150:8449 | 否 | ✓ 1904ms | ✓ 541ms | ✓ 1488ms | ✓ 1282ms | http |
| 38.34.183.234:8450 | ✓ 981ms | ✓ 1858ms | ✓ 563ms | ✓ 1682ms | ✓ 1671ms | http |
| 38.145.220.198:8448 | ✓ 1168ms | ✓ 1810ms | ✓ 470ms | ✓ 1006ms | ✓ 992ms | http |
| 45.136.131.59:8450 | 否 | 否 | ✓ 1346ms | ✓ 1149ms | ✓ 855ms | http |
| 158.160.215.167:8124 | ✓ 687ms | ✓ 1749ms | ✓ 865ms | ✓ 1866ms | ✓ 1440ms | http |
| 1.231.81.166:3128 | ✓ 1992ms | ✓ 1607ms | ✓ 1707ms | ✓ 1633ms | ✓ 1181ms | http |
| 38.34.183.233:8448 | ✓ 1211ms | ✓ 918ms | ✓ 1280ms | ✓ 1880ms | ✓ 782ms | http |
| 45.136.130.186:8451 | ✓ 952ms | ✓ 1456ms | 否 | 否 | ✓ 1665ms | http |
| 45.136.130.177:8448 | 否 | ✓ 1282ms | ✓ 1951ms | ✓ 1868ms | 否 | http |
| 38.145.203.96:8452 | ✓ 1644ms | 否 | ✓ 473ms | ✓ 936ms | ✓ 753ms | http |
| 38.34.179.61:8445 | ✓ 1106ms | ✓ 1136ms | ✓ 1987ms | ✓ 1663ms | 否 | http |
| 38.34.183.224:8448 | ✓ 1115ms | 否 | ✓ 560ms | ✓ 1521ms | 否 | http |
| 85.208.108.43:2094 | ✓ 1307ms | 否 | ✓ 1321ms | ✓ 1220ms | ✓ 1004ms | http |
| 43.99.54.236:5555 | 否 | 否 | ✓ 1884ms | ✓ 1080ms | ✓ 848ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1520ms | ✓ 1415ms | ✓ 1743ms | ✓ 1317ms | http |
| 45.186.6.104:3128 | ✓ 1137ms | ✓ 1862ms | ✓ 1624ms | 否 | 否 | http |
| 38.34.179.14:8450 | 否 | ✓ 954ms | ✓ 650ms | 否 | ✓ 1233ms | http |
| 38.145.208.244:8448 | ✓ 452ms | ✓ 1491ms | ✓ 1753ms | 否 | 否 | http |
| 38.34.183.225:8450 | ✓ 377ms | ✓ 920ms | ✓ 471ms | ✓ 1108ms | ✓ 1773ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 867ms | ✓ 1053ms | ✓ 1297ms | http |
| 194.67.99.223:1080 | ✓ 1639ms | ✓ 1554ms | 否 | 否 | ✓ 1398ms | http |
| 45.129.141.143:3128 | ✓ 1139ms | ✓ 1656ms | ✓ 1552ms | ✓ 1704ms | 否 | http |
| 210.48.154.94:80 | 否 | 否 | ✓ 1064ms | ✓ 1346ms | ✓ 1054ms | http |
| 45.136.198.40:3128 | ✓ 1472ms | 否 | ✓ 1776ms | 否 | ✓ 1827ms | http |
| 148.153.56.51:80 | ✓ 813ms | ✓ 895ms | ✓ 1033ms | ✓ 1099ms | ✓ 886ms | http |
| 207.254.71.62:8088 | ✓ 1050ms | ✓ 1875ms | ✓ 1210ms | ✓ 1973ms | ✓ 1707ms | http |
| 112.163.160.93:3128 | 否 | 否 | ✓ 1058ms | ✓ 1243ms | ✓ 1061ms | http |
| 38.34.178.154:8445 | ✓ 528ms | ✓ 1593ms | ✓ 1136ms | 否 | 否 | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 944ms | ✓ 1188ms | ✓ 1758ms | http |
| 38.34.179.16:8451 | 否 | ✓ 918ms | ✓ 388ms | ✓ 1370ms | ✓ 1674ms | http |
| 38.34.179.96:8451 | ✓ 481ms | ✓ 1007ms | ✓ 1065ms | ✓ 940ms | ✓ 1104ms | http |
| 34.96.238.40:8080 | ✓ 1428ms | ✓ 1529ms | ✓ 1122ms | ✓ 1288ms | ✓ 1295ms | http |
| 103.179.180.134:8080 | 否 | 否 | ✓ 1808ms | ✓ 1644ms | ✓ 1943ms | http |
| 38.145.203.76:8450 | ✓ 1699ms | ✓ 834ms | ✓ 351ms | ✓ 1119ms | ✓ 775ms | http |
| 45.136.130.186:8452 | ✓ 297ms | ✓ 924ms | ✓ 322ms | ✓ 954ms | ✓ 710ms | http |
| 45.136.130.168:8448 | ✓ 917ms | ✓ 933ms | ✓ 494ms | ✓ 1384ms | ✓ 1434ms | http |
| 45.136.130.173:8448 | ✓ 604ms | ✓ 1469ms | ✓ 569ms | ✓ 1325ms | ✓ 937ms | http |
| 45.136.130.170:8448 | ✓ 828ms | ✓ 1430ms | ✓ 489ms | ✓ 1299ms | ✓ 1188ms | http |
| 106.14.203.63:3333 | ✓ 1012ms | ✓ 1283ms | ✓ 1098ms | ✓ 1282ms | ✓ 1645ms | http |
| 45.136.131.54:8448 | ✓ 264ms | ✓ 984ms | ✓ 1556ms | ✓ 1874ms | ✓ 707ms | http |
| 38.145.220.22:8451 | 否 | 否 | ✓ 897ms | ✓ 1852ms | ✓ 1304ms | http |
| 45.136.130.171:8445 | 否 | 否 | ✓ 838ms | ✓ 922ms | ✓ 737ms | http |
| 45.136.131.49:8444 | ✓ 278ms | ✓ 946ms | ✓ 1372ms | 否 | ✓ 745ms | http |
| 211.217.231.234:8080 | 否 | ✓ 1691ms | ✓ 1346ms | ✓ 1330ms | ✓ 1009ms | http |
| 45.136.130.185:8444 | 否 | 否 | ✓ 1563ms | ✓ 1055ms | ✓ 714ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1491ms | ✓ 1114ms | 否 | ✓ 1135ms | http |
| 192.203.0.250:999 | 否 | ✓ 1284ms | ✓ 1062ms | ✓ 1266ms | 否 | http |
| 190.52.110.31:999 | 否 | ✓ 1655ms | ✓ 1645ms | ✓ 1586ms | 否 | http |
| 125.64.244.100:8889 | ✓ 1855ms | ✓ 1959ms | 否 | 否 | ✓ 1950ms | http |
| 38.145.218.212:8448 | 否 | 否 | ✓ 294ms | ✓ 914ms | ✓ 784ms | http |
| 103.113.70.189:1081 | ✓ 286ms | ✓ 1732ms | ✓ 143ms | ✓ 1116ms | ✓ 726ms | http |
| 103.82.23.118:5216 | ✓ 1740ms | 否 | ✓ 1863ms | ✓ 1992ms | 否 | http |

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
