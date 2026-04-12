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

最后更新：2026-04-12 14:04:27 UTC（2026-04-12 22:04:27 UTC+8）

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
| 167.103.115.102:8800 | ✓ 1765ms | 否 | ✓ 931ms | ✓ 1234ms | ✓ 1014ms | http |
| 147.161.210.140:8800 | ✓ 1512ms | 否 | ✓ 1342ms | 否 | ✓ 730ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1347ms | ✓ 1543ms | ✓ 1778ms | ✓ 1018ms | http |
| 167.103.34.108:8800 | ✓ 1823ms | 否 | ✓ 1580ms | ✓ 1643ms | 否 | http |
| 5.196.101.18:3128 | ✓ 1982ms | 否 | ✓ 1212ms | 否 | ✓ 1956ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1928ms | 否 | ✓ 1890ms | ✓ 1794ms | http |
| 167.103.144.127:8800 | ✓ 1424ms | ✓ 1762ms | ✓ 1217ms | 否 | ✓ 1572ms | http |
| 223.84.151.86:30005 | ✓ 1157ms | 否 | ✓ 1852ms | ✓ 1166ms | 否 | http |
| 5.104.87.17:8051 | ✓ 924ms | 否 | ✓ 1371ms | ✓ 1206ms | ✓ 798ms | http |
| 167.103.31.122:8800 | ✓ 1541ms | 否 | ✓ 1384ms | ✓ 1634ms | ✓ 1562ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1065ms | ✓ 1376ms | ✓ 1120ms | http |
| 20.210.39.153:8561 | ✓ 580ms | ✓ 1064ms | ✓ 1328ms | 否 | 否 | http |
| 45.167.125.21:999 | 否 | 否 | ✓ 1584ms | ✓ 1927ms | ✓ 1519ms | http |
| 95.214.9.93:3128 | ✓ 1272ms | 否 | ✓ 972ms | ✓ 1944ms | ✓ 1382ms | http |
| 36.141.21.200:7890 | 否 | ✓ 1084ms | ✓ 1147ms | ✓ 1202ms | 否 | http |
| 147.161.239.240:8800 | ✓ 893ms | 否 | ✓ 1546ms | ✓ 1868ms | ✓ 1739ms | http |
| 120.92.108.86:7890 | ✓ 1225ms | 否 | ✓ 1255ms | 否 | ✓ 1651ms | http |
| 20.78.118.91:8561 | ✓ 1446ms | ✓ 1389ms | ✓ 454ms | ✓ 794ms | ✓ 605ms | http |
| 20.78.26.206:8561 | ✓ 1441ms | 否 | ✓ 501ms | ✓ 824ms | ✓ 691ms | http |
| 45.140.147.82:1082 | ✓ 1315ms | ✓ 1299ms | ✓ 1178ms | 否 | 否 | http |
| 38.145.208.241:8447 | ✓ 503ms | ✓ 1042ms | ✓ 1146ms | 否 | ✓ 562ms | http |
| 103.157.200.126:3128 | ✓ 1941ms | 否 | ✓ 1361ms | ✓ 1910ms | ✓ 1516ms | http |
| 200.125.171.254:999 | ✓ 814ms | 否 | ✓ 1501ms | ✓ 1763ms | ✓ 1328ms | http |
| 103.113.70.189:1082 | ✓ 560ms | ✓ 1253ms | ✓ 305ms | ✓ 1204ms | ✓ 1010ms | http |
| 103.113.70.189:1081 | ✓ 547ms | ✓ 1313ms | ✓ 292ms | ✓ 1189ms | ✓ 984ms | http |
| 20.210.76.178:8561 | ✓ 1820ms | ✓ 946ms | ✓ 706ms | ✓ 804ms | ✓ 642ms | http |
| 20.27.15.49:8561 | ✓ 1820ms | ✓ 1322ms | ✓ 518ms | ✓ 724ms | ✓ 661ms | http |
| 20.210.76.175:8561 | ✓ 1821ms | 否 | ✓ 438ms | ✓ 789ms | ✓ 613ms | http |
| 20.210.76.104:8561 | ✓ 1821ms | 否 | ✓ 433ms | ✓ 794ms | ✓ 614ms | http |
| 45.140.147.82:1081 | ✓ 1655ms | 否 | ✓ 1446ms | 否 | ✓ 1101ms | http |
| 35.225.22.61:80 | ✓ 938ms | 否 | ✓ 793ms | ✓ 1263ms | 否 | http |
| 144.31.25.69:21064 | ✓ 1210ms | 否 | ✓ 1015ms | 否 | ✓ 1981ms | http |
| 217.76.245.80:999 | ✓ 1028ms | 否 | ✓ 1326ms | ✓ 1449ms | ✓ 1198ms | http |
| 114.237.77.245:1080 | ✓ 849ms | ✓ 1128ms | ✓ 872ms | 否 | ✓ 932ms | http |
| 8.219.97.248:80 | ✓ 996ms | ✓ 1747ms | 否 | ✓ 1218ms | 否 | http |
| 210.223.44.230:3128 | ✓ 647ms | ✓ 1170ms | ✓ 1119ms | ✓ 919ms | ✓ 651ms | http |
| 38.145.203.35:8450 | ✓ 293ms | 否 | ✓ 1557ms | ✓ 731ms | 否 | http |
| 101.32.244.83:8080 | ✓ 944ms | 否 | ✓ 878ms | ✓ 1375ms | ✓ 1198ms | http |
| 121.43.196.213:8222 | ✓ 930ms | ✓ 981ms | ✓ 823ms | ✓ 1118ms | ✓ 932ms | http |
| 121.43.196.210:8222 | ✓ 971ms | ✓ 1018ms | ✓ 812ms | ✓ 1106ms | ✓ 896ms | http |
| 114.55.226.123:10086 | ✓ 1099ms | ✓ 1393ms | ✓ 1002ms | ✓ 1259ms | ✓ 1090ms | http |
| 81.169.170.253:3128 | ✓ 752ms | 否 | ✓ 1945ms | 否 | ✓ 1959ms | http |
| 38.145.220.39:8452 | ✓ 1340ms | 否 | 否 | ✓ 1782ms | ✓ 574ms | http |
| 38.145.220.81:8445 | ✓ 1342ms | 否 | 否 | ✓ 1746ms | ✓ 721ms | http |
| 38.145.208.221:8447 | ✓ 583ms | ✓ 864ms | ✓ 126ms | ✓ 892ms | ✓ 1317ms | http |
| 38.34.179.51:8449 | ✓ 265ms | ✓ 805ms | ✓ 1059ms | 否 | ✓ 526ms | http |
| 38.145.208.214:8446 | 否 | ✓ 1540ms | ✓ 837ms | 否 | ✓ 951ms | http |
| 79.132.136.58:3128 | ✓ 656ms | ✓ 1857ms | ✓ 921ms | ✓ 1535ms | ✓ 1195ms | http |
| 171.227.167.109:1009 | ✓ 1654ms | 否 | ✓ 1111ms | 否 | ✓ 1029ms | http |
| 38.34.178.245:8446 | ✓ 683ms | ✓ 985ms | ✓ 1030ms | ✓ 1705ms | ✓ 1085ms | http |
| 34.50.41.219:3128 | ✓ 1592ms | 否 | ✓ 1148ms | ✓ 994ms | ✓ 916ms | http |
| 38.34.179.77:8448 | ✓ 1651ms | 否 | ✓ 200ms | ✓ 1720ms | ✓ 1560ms | http |
| 104.168.93.120:8080 | ✓ 1867ms | 否 | ✓ 1162ms | ✓ 1651ms | 否 | http |
| 38.34.179.13:8451 | ✓ 1681ms | 否 | ✓ 1347ms | 否 | ✓ 583ms | http |
| 38.145.220.39:8449 | ✓ 1421ms | ✓ 1432ms | 否 | ✓ 1596ms | ✓ 1575ms | http |
| 107.172.102.234:40621 | ✓ 1851ms | ✓ 1160ms | ✓ 716ms | ✓ 833ms | ✓ 524ms | http |
| 38.34.179.74:8447 | ✓ 1918ms | ✓ 1617ms | ✓ 1005ms | ✓ 975ms | ✓ 618ms | http |
| 45.136.130.249:8446 | ✓ 1862ms | 否 | ✓ 91ms | ✓ 852ms | ✓ 1736ms | http |
| 38.145.218.13:8446 | ✓ 1858ms | 否 | ✓ 92ms | ✓ 856ms | 否 | http |
| 38.145.220.196:8446 | ✓ 1903ms | 否 | ✓ 1032ms | 否 | ✓ 542ms | http |
| 38.145.218.162:8448 | ✓ 1903ms | ✓ 1405ms | ✓ 1868ms | ✓ 699ms | ✓ 575ms | http |
| 38.145.218.160:8448 | ✓ 1907ms | ✓ 1866ms | ✓ 1415ms | ✓ 928ms | ✓ 683ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1205ms | ✓ 1616ms | ✓ 1563ms | 否 | http |
| 38.145.208.206:8449 | ✓ 1902ms | ✓ 840ms | ✓ 532ms | 否 | ✓ 1439ms | http |
| 38.145.220.198:8446 | ✓ 1900ms | 否 | ✓ 1039ms | 否 | ✓ 816ms | http |
| 38.145.220.193:8446 | ✓ 1903ms | 否 | ✓ 1037ms | 否 | ✓ 547ms | http |
| 38.145.220.56:8446 | ✓ 1903ms | 否 | ✓ 1038ms | 否 | ✓ 1130ms | http |
| 218.108.131.186:17890 | ✓ 1173ms | ✓ 1008ms | ✓ 851ms | ✓ 1082ms | ✓ 892ms | http |
| 1.231.81.166:3128 | ✓ 1457ms | ✓ 1198ms | ✓ 1593ms | 否 | ✓ 1861ms | http |
| 38.34.179.40:8446 | ✓ 399ms | 否 | ✓ 1469ms | ✓ 740ms | ✓ 873ms | http |
| 137.59.47.73:3128 | ✓ 1721ms | 否 | ✓ 1905ms | ✓ 1217ms | 否 | http |
| 45.140.147.155:1082 | ✓ 604ms | ✓ 1336ms | ✓ 1702ms | ✓ 1779ms | 否 | http |
| 103.125.181.135:9999 | ✓ 984ms | 否 | ✓ 1175ms | ✓ 1648ms | ✓ 1182ms | http |
| 61.52.131.172:8443 | ✓ 877ms | ✓ 1177ms | ✓ 1044ms | ✓ 1182ms | ✓ 933ms | http |
| 38.145.208.204:8446 | ✓ 245ms | ✓ 862ms | ✓ 402ms | ✓ 805ms | 否 | http |
| 38.34.183.222:8453 | ✓ 1804ms | ✓ 795ms | ✓ 1014ms | 否 | ✓ 954ms | http |
| 38.34.179.106:8445 | ✓ 1208ms | 否 | 否 | ✓ 689ms | ✓ 713ms | http |
| 20.27.14.220:8561 | ✓ 1449ms | ✓ 965ms | ✓ 797ms | ✓ 1252ms | 否 | http |
| 20.27.11.248:8561 | ✓ 1390ms | ✓ 1134ms | ✓ 803ms | ✓ 1128ms | 否 | http |
| 8.219.195.129:1080 | ✓ 682ms | ✓ 1627ms | ✓ 750ms | ✓ 1027ms | ✓ 808ms | http |
| 103.56.112.120:7890 | ✓ 1209ms | ✓ 1888ms | ✓ 1364ms | ✓ 1417ms | ✓ 998ms | http |
| 38.34.179.98:8451 | ✓ 1485ms | ✓ 1638ms | 否 | ✓ 796ms | ✓ 1403ms | http |
| 36.103.198.235:7890 | ✓ 902ms | ✓ 1300ms | 否 | 否 | ✓ 1760ms | http |
| 115.231.181.40:8128 | ✓ 930ms | ✓ 1105ms | 否 | 否 | ✓ 1464ms | http |
| 139.159.99.242:8080 | ✓ 1171ms | ✓ 1003ms | ✓ 902ms | 否 | 否 | http |
| 38.145.218.161:8445 | ✓ 1340ms | ✓ 1817ms | ✓ 1354ms | 否 | 否 | http |
| 38.34.179.39:8452 | ✓ 1085ms | 否 | ✓ 1135ms | ✓ 1778ms | ✓ 500ms | http |

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
