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

最后更新：2026-03-04 12:32:52 UTC（2026-03-04 20:32:52 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | 否 | 否 | ✓ 550ms | ✓ 1348ms | ✓ 914ms | http |
| 103.84.95.54:7890 | ✓ 980ms | 否 | ✓ 619ms | ✓ 1175ms | ✓ 766ms | http |
| 120.232.242.119:22222 | ✓ 1038ms | ✓ 1188ms | ✓ 1010ms | ✓ 1104ms | ✓ 962ms | http |
| 120.238.159.228:22222 | ✓ 981ms | ✓ 1213ms | ✓ 976ms | ✓ 1159ms | ✓ 954ms | http |
| 211.171.114.154:3128 | ✓ 1054ms | 否 | ✓ 1392ms | ✓ 1278ms | ✓ 1914ms | http |
| 59.46.216.131:30001 | ✓ 1048ms | ✓ 1284ms | ✓ 1207ms | ✓ 1295ms | ✓ 1030ms | http |
| 113.59.32.141:22222 | ✓ 1553ms | ✓ 1929ms | ✓ 1644ms | ✓ 1879ms | ✓ 1619ms | http |
| 115.231.181.40:8128 | ✓ 901ms | 否 | ✓ 1572ms | 否 | ✓ 1433ms | http |
| 113.59.32.162:22222 | 否 | ✓ 1402ms | ✓ 1160ms | ✓ 1617ms | ✓ 1230ms | http |
| 5.75.196.26:40000 | ✓ 1200ms | ✓ 1856ms | ✓ 924ms | 否 | 否 | http |
| 183.249.5.105:22222 | ✓ 1049ms | ✓ 1225ms | ✓ 911ms | ✓ 1767ms | 否 | http |
| 113.59.32.161:22222 | ✓ 1536ms | ✓ 1655ms | ✓ 1785ms | 否 | 否 | http |
| 121.141.161.55:1080 | 否 | ✓ 1359ms | ✓ 1866ms | 否 | ✓ 972ms | http |
| 183.249.5.214:22222 | ✓ 720ms | ✓ 842ms | ✓ 690ms | ✓ 1151ms | ✓ 680ms | http |
| 4.216.195.194:3128 | ✓ 1838ms | 否 | ✓ 1373ms | ✓ 1209ms | ✓ 738ms | http |
| 101.43.255.96:80 | ✓ 1065ms | ✓ 1337ms | 否 | ✓ 1191ms | ✓ 1038ms | http |
| 81.70.169.194:80 | 否 | ✓ 1319ms | ✓ 1144ms | 否 | ✓ 1026ms | http |
| 91.193.240.157:9877 | ✓ 1482ms | 否 | ✓ 1999ms | 否 | ✓ 1645ms | http |
| 120.92.212.16:8890 | ✓ 959ms | ✓ 1412ms | ✓ 961ms | 否 | 否 | http |
| 121.40.231.103:7890 | 否 | 否 | ✓ 819ms | ✓ 1092ms | ✓ 805ms | http |
| 117.159.239.51:22222 | ✓ 799ms | ✓ 1012ms | ✓ 780ms | ✓ 1061ms | ✓ 836ms | http |
| 120.240.35.176:22222 | 否 | ✓ 1199ms | 否 | ✓ 1186ms | ✓ 962ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1210ms | 否 | ✓ 1372ms | ✓ 1202ms | http |
| 120.198.141.79:22222 | 否 | ✓ 1500ms | ✓ 1162ms | 否 | ✓ 1054ms | http |
| 120.240.35.160:22222 | ✓ 958ms | ✓ 1175ms | ✓ 996ms | ✓ 1168ms | 否 | http |
| 120.240.110.112:22222 | ✓ 959ms | ✓ 1192ms | ✓ 907ms | ✓ 1140ms | ✓ 1089ms | http |
| 120.240.35.178:22222 | ✓ 904ms | ✓ 1169ms | ✓ 1038ms | ✓ 1153ms | ✓ 1064ms | http |
| 160.238.65.9:3128 | ✓ 862ms | ✓ 1503ms | ✓ 699ms | ✓ 1748ms | ✓ 1266ms | http |
| 160.238.65.2:3128 | ✓ 701ms | 否 | ✓ 637ms | ✓ 1880ms | ✓ 1191ms | http |
| 160.238.65.4:3128 | ✓ 691ms | ✓ 1669ms | ✓ 659ms | ✓ 1734ms | ✓ 1778ms | http |
| 160.238.65.3:3128 | ✓ 621ms | ✓ 1643ms | ✓ 786ms | ✓ 1760ms | ✓ 1782ms | http |
| 160.238.65.7:3128 | ✓ 654ms | 否 | ✓ 611ms | ✓ 1967ms | ✓ 1255ms | http |
| 160.238.65.5:3128 | ✓ 674ms | 否 | ✓ 616ms | ✓ 1956ms | ✓ 1259ms | http |
| 160.238.65.6:3128 | ✓ 729ms | 否 | ✓ 1191ms | ✓ 1714ms | ✓ 1247ms | http |
| 160.238.65.8:3128 | ✓ 748ms | 否 | ✓ 1258ms | ✓ 1703ms | ✓ 1241ms | http |
| 120.240.29.174:22222 | ✓ 947ms | ✓ 1195ms | ✓ 891ms | ✓ 1180ms | ✓ 984ms | http |
| 58.220.95.12:12120 | ✓ 1205ms | 否 | ✓ 920ms | 否 | ✓ 888ms | http |
| 154.90.48.209:9090 | ✓ 755ms | 否 | ✓ 1392ms | ✓ 1511ms | 否 | http |
| 49.146.62.183:8082 | 否 | 否 | ✓ 1707ms | ✓ 1561ms | ✓ 1616ms | http |
| 120.198.141.84:22222 | ✓ 1374ms | ✓ 1284ms | ✓ 1045ms | ✓ 1232ms | ✓ 985ms | http |
| 120.240.29.51:22222 | 否 | ✓ 1221ms | ✓ 1021ms | ✓ 1149ms | ✓ 937ms | http |
| 35.234.17.221:8080 | 否 | 否 | ✓ 905ms | ✓ 985ms | ✓ 835ms | http |
| 83.219.250.8:62920 | ✓ 1326ms | 否 | ✓ 1415ms | 否 | ✓ 1641ms | http |
| 91.107.148.58:53967 | ✓ 658ms | ✓ 1788ms | ✓ 1577ms | 否 | 否 | http |
| 157.0.142.246:10061 | ✓ 1015ms | ✓ 1277ms | ✓ 1031ms | ✓ 1323ms | ✓ 974ms | http |
| 106.14.205.114:483 | ✓ 1001ms | ✓ 1062ms | ✓ 1119ms | ✓ 1095ms | ✓ 825ms | http |
| 113.59.32.148:22222 | 否 | ✓ 1314ms | ✓ 1041ms | ✓ 1341ms | ✓ 1140ms | http |
| 113.59.32.163:22222 | ✓ 1074ms | ✓ 1313ms | ✓ 975ms | ✓ 1486ms | ✓ 1124ms | http |
| 222.184.48.248:22222 | ✓ 866ms | ✓ 1177ms | ✓ 902ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1083ms | 否 | ✓ 1744ms | 否 | ✓ 1327ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1360ms | ✓ 1520ms | ✓ 1888ms | 否 | http |
| 61.72.221.94:3128 | ✓ 1751ms | 否 | ✓ 1574ms | ✓ 1501ms | ✓ 1886ms | http |
| 120.92.212.16:7890 | ✓ 1557ms | 否 | ✓ 1234ms | 否 | ✓ 1592ms | http |
| 205.209.118.30:3138 | ✓ 425ms | 否 | 否 | ✓ 1584ms | ✓ 1051ms | http |
| 103.215.36.88:19475 | ✓ 1765ms | 否 | ✓ 1026ms | ✓ 1856ms | 否 | http |
| 61.72.221.194:3128 | ✓ 608ms | 否 | ✓ 1179ms | ✓ 943ms | ✓ 730ms | http |
| 125.128.12.14:3128 | ✓ 1635ms | ✓ 1032ms | ✓ 1119ms | ✓ 946ms | ✓ 750ms | http |
| 8.217.147.173:8080 | 否 | ✓ 1267ms | ✓ 1633ms | ✓ 1556ms | ✓ 1425ms | http |
| 61.72.221.234:3128 | ✓ 1634ms | 否 | ✓ 1717ms | ✓ 1643ms | ✓ 777ms | http |
| 8.219.97.248:80 | ✓ 884ms | 否 | ✓ 1051ms | ✓ 1346ms | 否 | http |
| 61.72.110.54:3128 | ✓ 828ms | ✓ 1311ms | ✓ 540ms | ✓ 1188ms | ✓ 754ms | http |
| 14.56.177.44:3128 | ✓ 818ms | ✓ 1538ms | ✓ 538ms | ✓ 997ms | ✓ 718ms | http |
| 121.128.121.54:3128 | ✓ 840ms | ✓ 1180ms | ✓ 980ms | ✓ 948ms | ✓ 758ms | http |
| 61.72.110.94:3128 | ✓ 841ms | ✓ 1614ms | ✓ 539ms | ✓ 990ms | ✓ 731ms | http |
| 14.56.107.244:3128 | ✓ 826ms | ✓ 1133ms | ✓ 1308ms | ✓ 975ms | ✓ 781ms | http |
| 125.128.12.144:3128 | ✓ 839ms | ✓ 1162ms | ✓ 1002ms | 否 | ✓ 756ms | http |
| 188.166.208.168:9876 | ✓ 1330ms | 否 | ✓ 755ms | ✓ 1018ms | ✓ 771ms | http |
| 120.79.99.232:8099 | ✓ 1237ms | ✓ 1441ms | ✓ 1217ms | ✓ 1379ms | ✓ 1187ms | http |
| 210.223.44.230:3128 | ✓ 895ms | ✓ 1006ms | ✓ 1238ms | 否 | ✓ 1091ms | http |
| 103.82.93.219:3128 | ✓ 1524ms | 否 | 否 | ✓ 1202ms | ✓ 904ms | http |
| 103.166.185.54:3128 | ✓ 1697ms | ✓ 1633ms | ✓ 1393ms | ✓ 1214ms | ✓ 929ms | http |
| 45.136.198.40:3128 | ✓ 871ms | 否 | ✓ 1226ms | ✓ 1742ms | ✓ 1623ms | http |
| 90.84.188.97:8000 | ✓ 1115ms | 否 | 否 | ✓ 1732ms | ✓ 1335ms | http |
| 150.107.140.238:3128 | ✓ 1126ms | 否 | ✓ 1006ms | ✓ 1484ms | ✓ 1853ms | http |
| 117.159.239.42:22222 | ✓ 832ms | ✓ 972ms | ✓ 790ms | ✓ 1045ms | ✓ 820ms | http |
| 120.240.29.53:22222 | 否 | ✓ 1209ms | 否 | ✓ 1208ms | ✓ 1008ms | http |
| 120.240.29.168:22222 | ✓ 1093ms | ✓ 1206ms | ✓ 1261ms | ✓ 1263ms | ✓ 924ms | http |
| 47.77.180.205:1080 | ✓ 545ms | ✓ 1359ms | ✓ 699ms | ✓ 648ms | ✓ 479ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1142ms | ✓ 1162ms | ✓ 906ms | http |
| 159.89.31.62:8080 | ✓ 1152ms | 否 | ✓ 1434ms | ✓ 1799ms | ✓ 1331ms | http |
| 43.134.166.79:8888 | 否 | ✓ 1314ms | ✓ 1448ms | ✓ 1389ms | ✓ 1214ms | http |
| 27.254.99.183:8118 | ✓ 1916ms | 否 | ✓ 1257ms | ✓ 1261ms | 否 | http |
| 117.159.239.44:22222 | ✓ 785ms | ✓ 1025ms | ✓ 868ms | ✓ 1045ms | ✓ 809ms | http |
| 117.159.239.54:22222 | ✓ 1017ms | ✓ 1420ms | ✓ 1089ms | ✓ 1332ms | 否 | http |

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
