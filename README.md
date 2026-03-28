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

最后更新：2026-03-28 09:42:44 UTC（2026-03-28 17:42:44 UTC+8）

**代理总数：163**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 162 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 163 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 520ms | ✓ 1305ms | ✓ 1165ms | ✓ 1049ms | ✓ 843ms | http |
| 168.63.153.150:3128 | ✓ 764ms | ✓ 1379ms | ✓ 939ms | ✓ 1028ms | ✓ 832ms | http |
| 1.231.81.166:3128 | ✓ 1615ms | ✓ 1124ms | ✓ 1523ms | ✓ 1258ms | ✓ 1106ms | http |
| 147.161.210.140:8800 | ✓ 1581ms | 否 | ✓ 1622ms | ✓ 1584ms | ✓ 1042ms | http |
| 113.160.132.26:8080 | ✓ 1992ms | ✓ 1981ms | ✓ 1300ms | ✓ 1492ms | ✓ 1171ms | http |
| 106.75.15.167:7890 | ✓ 1029ms | ✓ 1551ms | 否 | 否 | ✓ 1523ms | http |
| 35.225.22.61:80 | ✓ 131ms | ✓ 1184ms | ✓ 1332ms | ✓ 809ms | ✓ 855ms | http |
| 38.34.183.8:8448 | ✓ 915ms | 否 | ✓ 744ms | ✓ 1032ms | ✓ 732ms | http |
| 193.233.22.29:10808 | ✓ 834ms | 否 | ✓ 1067ms | ✓ 1435ms | ✓ 922ms | http |
| 167.103.115.102:8800 | ✓ 1030ms | 否 | ✓ 1072ms | ✓ 1169ms | ✓ 1122ms | http |
| 167.103.34.108:8800 | ✓ 1249ms | 否 | ✓ 1283ms | ✓ 1481ms | ✓ 1446ms | http |
| 5.104.87.17:8051 | ✓ 1598ms | 否 | 否 | ✓ 1974ms | ✓ 1591ms | http |
| 167.103.144.127:8800 | ✓ 1801ms | 否 | ✓ 1616ms | 否 | ✓ 1568ms | http |
| 101.47.73.135:3128 | ✓ 1315ms | 否 | ✓ 1121ms | ✓ 1948ms | 否 | http |
| 186.148.180.46:999 | ✓ 1060ms | 否 | ✓ 1363ms | ✓ 1734ms | ✓ 1649ms | http |
| 180.250.219.58:53281 | ✓ 1894ms | 否 | ✓ 1858ms | 否 | ✓ 1958ms | http |
| 165.232.146.249:3128 | ✓ 511ms | ✓ 1866ms | ✓ 638ms | 否 | ✓ 1719ms | http |
| 147.161.239.240:8800 | ✓ 1173ms | ✓ 1617ms | ✓ 1041ms | ✓ 1606ms | ✓ 1443ms | http |
| 45.12.151.226:2829 | ✓ 1197ms | ✓ 1549ms | ✓ 649ms | ✓ 1602ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1632ms | 否 | 否 | ✓ 1105ms | ✓ 872ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 993ms | ✓ 1459ms | ✓ 1152ms | http |
| 167.103.31.122:8800 | ✓ 1563ms | 否 | ✓ 1382ms | ✓ 1705ms | ✓ 1493ms | http |
| 158.160.215.167:8124 | ✓ 859ms | 否 | ✓ 1818ms | 否 | ✓ 1783ms | http |
| 166.88.55.83:7890 | ✓ 757ms | ✓ 1237ms | ✓ 752ms | ✓ 954ms | ✓ 746ms | http |
| 185.76.240.150:10001 | ✓ 1196ms | ✓ 1936ms | ✓ 1241ms | 否 | 否 | http |
| 158.160.215.167:8126 | ✓ 1131ms | 否 | ✓ 1831ms | 否 | ✓ 1871ms | http |
| 128.199.113.85:9090 | ✓ 1006ms | 否 | ✓ 1030ms | ✓ 1216ms | 否 | http |
| 38.145.208.172:8448 | ✓ 365ms | 否 | ✓ 1679ms | ✓ 883ms | ✓ 763ms | http |
| 128.199.121.61:9090 | ✓ 1022ms | 否 | ✓ 1612ms | ✓ 1297ms | ✓ 966ms | http |
| 219.117.204.211:7799 | ✓ 1850ms | ✓ 1792ms | ✓ 1112ms | 否 | 否 | http |
| 101.43.127.100:8877 | ✓ 1031ms | ✓ 1241ms | 否 | ✓ 1193ms | ✓ 1762ms | http |
| 45.144.28.81:10808 | ✓ 1618ms | 否 | ✓ 1966ms | ✓ 1730ms | ✓ 1437ms | http |
| 38.145.208.169:8446 | ✓ 312ms | ✓ 1295ms | ✓ 627ms | ✓ 834ms | ✓ 671ms | http |
| 38.145.218.160:8450 | ✓ 1098ms | ✓ 799ms | ✓ 334ms | ✓ 903ms | ✓ 1687ms | http |
| 38.34.179.66:8446 | ✓ 419ms | ✓ 914ms | ✓ 901ms | 否 | ✓ 753ms | http |
| 38.34.179.14:8452 | ✓ 382ms | ✓ 1677ms | ✓ 1105ms | ✓ 1137ms | ✓ 1070ms | http |
| 38.34.179.17:8453 | ✓ 381ms | 否 | ✓ 770ms | ✓ 1344ms | ✓ 879ms | http |
| 38.145.220.27:8444 | ✓ 269ms | ✓ 822ms | ✓ 1143ms | ✓ 1612ms | ✓ 1000ms | http |
| 146.190.80.158:9090 | ✓ 840ms | 否 | 否 | ✓ 1312ms | ✓ 927ms | http |
| 38.145.218.76:8451 | ✓ 736ms | ✓ 804ms | ✓ 1010ms | ✓ 1091ms | ✓ 1144ms | http |
| 38.34.183.224:8451 | ✓ 536ms | ✓ 1591ms | ✓ 410ms | ✓ 988ms | ✓ 1801ms | http |
| 128.199.254.13:9090 | ✓ 865ms | 否 | ✓ 860ms | ✓ 1230ms | 否 | http |
| 128.199.116.219:9090 | ✓ 858ms | 否 | ✓ 894ms | ✓ 1165ms | 否 | http |
| 38.34.179.47:8446 | ✓ 382ms | ✓ 1332ms | 否 | ✓ 1361ms | ✓ 1321ms | http |
| 38.145.208.222:8446 | ✓ 982ms | ✓ 1682ms | ✓ 1459ms | ✓ 1592ms | ✓ 765ms | http |
| 38.145.208.196:8452 | ✓ 566ms | ✓ 1857ms | ✓ 760ms | 否 | ✓ 1334ms | http |
| 38.145.203.34:8445 | ✓ 1272ms | ✓ 864ms | ✓ 1987ms | ✓ 1573ms | 否 | http |
| 38.145.220.79:8453 | ✓ 1216ms | 否 | ✓ 986ms | ✓ 1608ms | ✓ 1641ms | http |
| 116.80.65.80:3172 | ✓ 1625ms | 否 | ✓ 1634ms | ✓ 1962ms | 否 | http |
| 45.136.131.33:8444 | ✓ 671ms | ✓ 1378ms | 否 | ✓ 1383ms | 否 | http |
| 45.136.131.27:8444 | ✓ 678ms | ✓ 1860ms | 否 | ✓ 1177ms | 否 | http |
| 177.234.217.88:999 | ✓ 1239ms | ✓ 1748ms | ✓ 1662ms | 否 | ✓ 1506ms | http |
| 43.99.54.236:5555 | ✓ 753ms | ✓ 1132ms | ✓ 766ms | ✓ 975ms | ✓ 776ms | http |
| 38.145.220.33:8448 | ✓ 377ms | ✓ 849ms | ✓ 406ms | ✓ 1872ms | ✓ 1318ms | http |
| 164.90.206.15:3128 | ✓ 973ms | ✓ 1830ms | ✓ 1849ms | ✓ 1627ms | ✓ 1116ms | http |
| 223.16.170.103:80 | ✓ 1720ms | 否 | ✓ 1772ms | ✓ 1562ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1328ms | 否 | ✓ 1834ms | 否 | ✓ 1819ms | http |
| 64.181.240.152:3128 | 否 | ✓ 1687ms | ✓ 652ms | ✓ 927ms | ✓ 783ms | http |
| 45.88.0.116:3128 | ✓ 1220ms | 否 | 否 | ✓ 1960ms | ✓ 1578ms | http |
| 45.88.0.99:3128 | ✓ 1265ms | ✓ 1483ms | 否 | ✓ 1337ms | ✓ 1032ms | http |
| 45.88.0.115:3128 | ✓ 511ms | ✓ 1432ms | 否 | ✓ 1323ms | ✓ 1074ms | http |
| 38.34.179.106:8450 | ✓ 614ms | ✓ 1021ms | ✓ 1191ms | ✓ 1742ms | ✓ 742ms | http |
| 213.220.62.62:3128 | ✓ 935ms | 否 | ✓ 461ms | ✓ 1463ms | ✓ 1112ms | http |
| 45.136.130.189:8451 | ✓ 999ms | ✓ 830ms | ✓ 652ms | 否 | ✓ 1043ms | http |
| 45.136.130.186:8451 | ✓ 617ms | ✓ 814ms | ✓ 818ms | 否 | ✓ 670ms | http |
| 38.34.179.151:8452 | ✓ 607ms | ✓ 932ms | ✓ 1932ms | ✓ 1763ms | ✓ 1321ms | http |
| 45.77.246.231:80 | ✓ 854ms | 否 | ✓ 1284ms | ✓ 1177ms | ✓ 939ms | http |
| 45.88.0.111:3128 | ✓ 449ms | ✓ 1474ms | ✓ 454ms | ✓ 1550ms | ✓ 1172ms | http |
| 45.88.0.98:3128 | ✓ 461ms | ✓ 1415ms | ✓ 452ms | ✓ 1554ms | ✓ 1196ms | http |
| 45.88.0.113:3128 | ✓ 461ms | 否 | ✓ 463ms | ✓ 1385ms | ✓ 1032ms | http |
| 45.88.0.117:3128 | ✓ 463ms | 否 | ✓ 462ms | ✓ 1646ms | ✓ 1026ms | http |
| 38.145.208.185:8449 | ✓ 1016ms | ✓ 1046ms | 否 | ✓ 1433ms | ✓ 644ms | http |
| 38.145.208.173:8444 | 否 | ✓ 1848ms | ✓ 1060ms | 否 | ✓ 848ms | http |
| 45.88.0.114:3128 | ✓ 536ms | 否 | ✓ 436ms | ✓ 1366ms | ✓ 1072ms | http |
| 59.8.203.55:80 | ✓ 1697ms | ✓ 1242ms | ✓ 1308ms | ✓ 1117ms | ✓ 872ms | http |
| 88.80.150.82:8080 | ✓ 1455ms | ✓ 1975ms | 否 | 否 | ✓ 1728ms | https |
| 74.50.96.247:8888 | 否 | 否 | ✓ 1661ms | ✓ 1099ms | ✓ 1174ms | http |
| 210.223.44.230:3128 | ✓ 1756ms | ✓ 1533ms | ✓ 1080ms | 否 | 否 | http |
| 128.199.114.189:9090 | ✓ 956ms | 否 | ✓ 978ms | ✓ 1275ms | ✓ 1039ms | http |
| 38.145.220.33:8453 | ✓ 343ms | ✓ 1105ms | ✓ 417ms | ✓ 925ms | ✓ 707ms | http |
| 38.145.220.33:8445 | ✓ 342ms | ✓ 1150ms | ✓ 403ms | ✓ 911ms | ✓ 758ms | http |
| 45.136.131.34:8449 | ✓ 309ms | ✓ 1147ms | ✓ 904ms | ✓ 1174ms | ✓ 911ms | http |
| 45.136.130.247:8448 | ✓ 309ms | ✓ 949ms | ✓ 904ms | ✓ 1176ms | ✓ 831ms | http |
| 38.145.220.103:8446 | ✓ 380ms | ✓ 1597ms | ✓ 328ms | ✓ 869ms | ✓ 806ms | http |
| 38.145.218.13:8452 | ✓ 289ms | ✓ 1040ms | ✓ 841ms | ✓ 1255ms | ✓ 773ms | http |
| 38.34.179.35:8448 | ✓ 444ms | ✓ 1327ms | ✓ 386ms | ✓ 1169ms | ✓ 751ms | http |
| 38.145.220.102:8453 | ✓ 528ms | ✓ 1460ms | ✓ 303ms | ✓ 856ms | ✓ 819ms | http |
| 38.34.179.19:8453 | ✓ 1022ms | 否 | ✓ 434ms | ✓ 999ms | ✓ 1002ms | http |
| 38.34.179.91:8445 | ✓ 564ms | ✓ 1472ms | ✓ 678ms | ✓ 1801ms | ✓ 723ms | http |
| 38.145.203.107:8453 | ✓ 809ms | ✓ 1844ms | ✓ 504ms | ✓ 905ms | ✓ 1150ms | http |
| 38.145.203.105:8453 | ✓ 767ms | ✓ 1390ms | ✓ 977ms | ✓ 884ms | ✓ 940ms | http |
| 38.145.203.124:8452 | ✓ 620ms | ✓ 1646ms | ✓ 901ms | ✓ 905ms | ✓ 1272ms | http |
| 38.34.179.62:8453 | ✓ 428ms | ✓ 1180ms | ✓ 617ms | ✓ 1416ms | ✓ 880ms | http |
| 38.145.208.191:8446 | ✓ 475ms | 否 | ✓ 919ms | 否 | ✓ 674ms | http |
| 38.34.179.150:8444 | ✓ 599ms | 否 | ✓ 466ms | ✓ 948ms | ✓ 1102ms | http |
| 38.34.179.69:8452 | ✓ 695ms | 否 | ✓ 567ms | ✓ 984ms | ✓ 1879ms | http |
| 38.145.220.35:8444 | ✓ 452ms | ✓ 1991ms | ✓ 791ms | ✓ 865ms | ✓ 1253ms | http |
| 38.145.218.229:8450 | ✓ 500ms | ✓ 947ms | ✓ 990ms | 否 | ✓ 663ms | http |
| 104.248.151.93:9090 | ✓ 1468ms | 否 | ✓ 1598ms | ✓ 1220ms | ✓ 964ms | http |
| 38.145.208.197:8452 | ✓ 1169ms | ✓ 841ms | ✓ 671ms | 否 | ✓ 983ms | http |

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
