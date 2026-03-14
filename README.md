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

最后更新：2026-03-14 06:44:14 UTC（2026-03-14 14:44:14 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 314ms | 否 | ✓ 1040ms | ✓ 1389ms | ✓ 939ms | http |
| 103.113.70.189:1081 | ✓ 339ms | ✓ 1693ms | 否 | 否 | ✓ 1144ms | http |
| 113.160.132.26:8080 | ✓ 1420ms | ✓ 1648ms | ✓ 1303ms | ✓ 1219ms | ✓ 987ms | http |
| 45.167.124.52:8080 | ✓ 1165ms | ✓ 1981ms | ✓ 1205ms | ✓ 1829ms | ✓ 1807ms | http |
| 183.249.5.117:22222 | 否 | ✓ 1245ms | 否 | ✓ 1074ms | ✓ 825ms | http |
| 120.232.242.119:22222 | ✓ 1016ms | ✓ 1362ms | ✓ 1097ms | 否 | ✓ 964ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1218ms | ✓ 1335ms | ✓ 1043ms | http |
| 101.47.73.135:3128 | ✓ 1184ms | 否 | ✓ 1164ms | ✓ 1151ms | ✓ 1203ms | http |
| 222.184.48.251:22222 | ✓ 1047ms | ✓ 1166ms | ✓ 959ms | 否 | ✓ 1030ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1260ms | ✓ 1541ms | ✓ 1274ms | http |
| 186.148.180.46:999 | ✓ 1626ms | 否 | ✓ 1350ms | ✓ 1815ms | 否 | http |
| 216.180.127.45:1080 | ✓ 572ms | 否 | ✓ 929ms | 否 | ✓ 1128ms | http |
| 160.238.65.6:3128 | ✓ 1008ms | ✓ 1679ms | ✓ 1477ms | 否 | 否 | http |
| 160.238.65.7:3128 | ✓ 1003ms | ✓ 1766ms | ✓ 1400ms | 否 | 否 | http |
| 160.238.65.3:3128 | ✓ 1006ms | ✓ 1519ms | ✓ 1642ms | 否 | 否 | http |
| 160.238.65.9:3128 | ✓ 1003ms | ✓ 1679ms | ✓ 1481ms | 否 | 否 | http |
| 160.238.65.4:3128 | ✓ 1006ms | ✓ 1741ms | ✓ 1421ms | 否 | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1240ms | ✓ 1610ms | ✓ 1472ms | http |
| 35.225.22.61:80 | ✓ 660ms | ✓ 1362ms | 否 | ✓ 1050ms | ✓ 923ms | http |
| 81.70.169.194:80 | ✓ 1036ms | 否 | 否 | ✓ 1337ms | ✓ 1407ms | http |
| 8.219.97.248:80 | ✓ 1145ms | 否 | ✓ 1063ms | ✓ 1291ms | 否 | http |
| 38.145.208.93:8443 | ✓ 672ms | 否 | ✓ 569ms | ✓ 1101ms | ✓ 735ms | http |
| 38.145.203.246:8443 | ✓ 672ms | 否 | ✓ 570ms | ✓ 1122ms | ✓ 751ms | http |
| 113.59.32.148:22222 | ✓ 1238ms | 否 | ✓ 1111ms | 否 | ✓ 1036ms | http |
| 120.238.159.250:22222 | ✓ 1070ms | ✓ 1342ms | ✓ 978ms | ✓ 1206ms | ✓ 970ms | http |
| 85.198.96.242:3128 | 否 | ✓ 1850ms | 否 | ✓ 1670ms | ✓ 1347ms | http |
| 101.43.255.96:80 | ✓ 1287ms | ✓ 1334ms | 否 | 否 | ✓ 1064ms | http |
| 62.113.119.14:8080 | ✓ 1125ms | 否 | ✓ 625ms | ✓ 1467ms | ✓ 1121ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1219ms | 否 | ✓ 1307ms | ✓ 1478ms | http |
| 62.60.177.204:34094 | ✓ 492ms | 否 | 否 | ✓ 1002ms | ✓ 762ms | http |
| 117.159.239.51:22222 | ✓ 973ms | ✓ 1135ms | ✓ 998ms | ✓ 1221ms | ✓ 964ms | http |
| 45.140.147.155:1082 | ✓ 590ms | 否 | ✓ 1232ms | ✓ 1599ms | ✓ 922ms | http |
| 45.140.147.155:1081 | ✓ 670ms | 否 | ✓ 1170ms | ✓ 1619ms | ✓ 1349ms | http |
| 160.250.5.22:1 | ✓ 1096ms | 否 | ✓ 1089ms | ✓ 1610ms | ✓ 1064ms | http |
| 160.238.65.2:3128 | ✓ 799ms | ✓ 1832ms | 否 | ✓ 1963ms | ✓ 1472ms | http |
| 160.238.65.5:3128 | ✓ 794ms | ✓ 1797ms | 否 | 否 | ✓ 1473ms | http |
| 222.184.48.252:22222 | ✓ 953ms | ✓ 1309ms | ✓ 1588ms | 否 | ✓ 1781ms | http |
| 117.159.239.50:22222 | ✓ 1179ms | ✓ 1143ms | ✓ 887ms | ✓ 1164ms | ✓ 924ms | http |
| 43.167.227.161:1080 | ✓ 790ms | ✓ 1870ms | ✓ 717ms | 否 | 否 | http |
| 116.80.65.75:3172 | 否 | 否 | ✓ 1586ms | ✓ 1906ms | ✓ 1749ms | http |
| 121.40.231.103:7890 | ✓ 942ms | ✓ 1122ms | ✓ 941ms | ✓ 1176ms | ✓ 912ms | http |
| 45.186.6.104:3128 | ✓ 1208ms | ✓ 1916ms | ✓ 1961ms | 否 | 否 | http |
| 38.145.203.161:8443 | ✓ 1006ms | ✓ 928ms | ✓ 1858ms | ✓ 1159ms | ✓ 782ms | http |
| 165.227.5.10:8888 | ✓ 992ms | ✓ 1544ms | ✓ 1091ms | 否 | ✓ 1242ms | http |
| 113.59.32.162:22222 | ✓ 1199ms | ✓ 1438ms | ✓ 1013ms | ✓ 1374ms | ✓ 1010ms | http |
| 3.137.167.45:34876 | ✓ 1253ms | 否 | ✓ 1562ms | 否 | ✓ 1897ms | http |
| 152.42.213.210:8080 | ✓ 901ms | 否 | 否 | ✓ 1684ms | ✓ 956ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1980ms | ✓ 1966ms | ✓ 1703ms | http |
| 45.136.130.165:8443 | ✓ 1055ms | ✓ 1533ms | ✓ 782ms | ✓ 1221ms | ✓ 902ms | http |
| 150.230.249.50:1080 | 否 | 否 | ✓ 1060ms | ✓ 1174ms | ✓ 861ms | http |
| 47.101.149.27:9010 | ✓ 1427ms | 否 | 否 | ✓ 1548ms | ✓ 1412ms | http |
| 45.136.198.40:3128 | ✓ 692ms | 否 | ✓ 1011ms | ✓ 1556ms | ✓ 1178ms | http |
| 103.84.95.54:7890 | ✓ 727ms | 否 | 否 | ✓ 1042ms | ✓ 744ms | http |
| 120.240.35.173:22222 | 否 | ✓ 1528ms | ✓ 1011ms | 否 | ✓ 990ms | http |
| 120.238.159.230:22222 | ✓ 994ms | ✓ 1355ms | 否 | ✓ 1321ms | ✓ 974ms | http |
| 106.117.208.101:7890 | ✓ 1128ms | ✓ 1450ms | 否 | 否 | ✓ 1858ms | http |
| 45.88.0.113:3128 | ✓ 1563ms | 否 | ✓ 850ms | ✓ 1693ms | 否 | http |
| 45.88.0.115:3128 | ✓ 879ms | ✓ 1587ms | ✓ 1909ms | ✓ 1719ms | 否 | http |
| 213.220.62.62:3128 | ✓ 558ms | 否 | 否 | ✓ 1779ms | ✓ 1054ms | http |
| 45.88.0.98:3128 | ✓ 503ms | 否 | ✓ 539ms | ✓ 1613ms | 否 | http |
| 45.88.0.116:3128 | ✓ 505ms | ✓ 1907ms | ✓ 535ms | ✓ 1345ms | ✓ 1036ms | http |
| 45.88.0.114:3128 | ✓ 944ms | ✓ 1431ms | ✓ 463ms | ✓ 1299ms | ✓ 1028ms | http |
| 183.249.5.109:22222 | ✓ 854ms | ✓ 998ms | ✓ 1031ms | ✓ 1037ms | ✓ 903ms | http |
| 2.56.122.146:10808 | ✓ 942ms | 否 | ✓ 1031ms | ✓ 1731ms | ✓ 1160ms | http |
| 45.88.0.111:3128 | ✓ 947ms | 否 | ✓ 1257ms | 否 | ✓ 1072ms | http |
| 45.88.0.117:3128 | ✓ 948ms | 否 | ✓ 1275ms | ✓ 1989ms | ✓ 1056ms | http |
| 180.127.149.244:1080 | ✓ 1084ms | ✓ 1239ms | 否 | 否 | ✓ 964ms | http |
| 38.145.203.135:8443 | ✓ 493ms | ✓ 961ms | ✓ 525ms | ✓ 1954ms | ✓ 729ms | http |
| 59.46.216.131:30001 | ✓ 1999ms | ✓ 1450ms | 否 | 否 | ✓ 1955ms | http |
| 183.249.5.105:22222 | ✓ 825ms | ✓ 1299ms | 否 | ✓ 1084ms | ✓ 832ms | http |
| 38.145.208.94:8443 | 否 | 否 | ✓ 322ms | ✓ 973ms | ✓ 770ms | http |
| 38.145.208.95:8443 | 否 | 否 | ✓ 337ms | ✓ 1020ms | ✓ 750ms | http |
| 150.249.255.91:3128 | ✓ 1710ms | ✓ 1194ms | ✓ 731ms | ✓ 940ms | 否 | http |
| 45.136.130.245:8447 | ✓ 584ms | 否 | ✓ 335ms | ✓ 943ms | ✓ 776ms | http |
| 38.145.208.97:8443 | 否 | ✓ 1981ms | ✓ 337ms | ✓ 975ms | ✓ 768ms | http |
| 38.145.208.96:8443 | ✓ 1069ms | ✓ 1705ms | ✓ 335ms | ✓ 965ms | ✓ 764ms | http |
| 120.198.141.75:22222 | 否 | ✓ 1783ms | 否 | ✓ 1465ms | ✓ 1241ms | http |
| 91.233.223.147:3128 | ✓ 1259ms | 否 | ✓ 1519ms | 否 | ✓ 1944ms | http |
| 222.184.48.248:22222 | ✓ 1909ms | 否 | ✓ 1306ms | ✓ 1213ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1924ms | 否 | ✓ 1318ms | ✓ 1503ms | ✓ 1652ms | http |
| 222.184.48.235:22222 | ✓ 1832ms | ✓ 1371ms | ✓ 1320ms | ✓ 1686ms | ✓ 1895ms | http |
| 101.43.127.100:8877 | ✓ 990ms | 否 | ✓ 1898ms | 否 | ✓ 1011ms | http |
| 103.39.51.190:8080 | ✓ 1683ms | 否 | 否 | ✓ 1586ms | ✓ 1922ms | http |
| 61.52.131.172:8443 | ✓ 1002ms | ✓ 1318ms | ✓ 1078ms | 否 | ✓ 1020ms | http |
| 45.88.0.99:3128 | ✓ 732ms | ✓ 1655ms | ✓ 1686ms | ✓ 1314ms | ✓ 1004ms | http |
| 120.240.29.51:22222 | ✓ 1262ms | ✓ 1357ms | ✓ 1034ms | ✓ 1403ms | ✓ 1032ms | http |
| 120.238.159.229:22222 | ✓ 956ms | ✓ 1306ms | ✓ 1057ms | ✓ 1394ms | ✓ 1047ms | http |

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
