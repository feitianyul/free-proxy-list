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

最后更新：2026-03-06 00:48:20 UTC（2026-03-06 08:48:20 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 107.174.80.186:3128 | ✓ 769ms | ✓ 867ms | ✓ 1195ms | ✓ 878ms | ✓ 706ms | http |
| 205.209.118.30:3138 | ✓ 1202ms | 否 | ✓ 909ms | ✓ 1250ms | ✓ 1155ms | http |
| 103.84.95.54:7890 | ✓ 706ms | 否 | ✓ 729ms | ✓ 1056ms | ✓ 726ms | http |
| 125.128.12.144:3128 | ✓ 1863ms | 否 | ✓ 1114ms | ✓ 1494ms | ✓ 1822ms | http |
| 61.72.221.194:3128 | ✓ 1851ms | ✓ 1421ms | ✓ 1038ms | 否 | 否 | http |
| 61.72.221.234:3128 | 否 | ✓ 1362ms | ✓ 1180ms | ✓ 1270ms | 否 | http |
| 61.72.110.94:3128 | 否 | ✓ 1589ms | ✓ 1374ms | 否 | ✓ 1275ms | http |
| 125.128.12.14:3128 | ✓ 1847ms | 否 | ✓ 1890ms | ✓ 1659ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1469ms | ✓ 996ms | ✓ 1295ms | 否 | http |
| 61.72.110.54:3128 | 否 | ✓ 1951ms | ✓ 1341ms | 否 | ✓ 1027ms | http |
| 14.56.177.44:3128 | ✓ 1924ms | ✓ 1937ms | 否 | ✓ 1867ms | ✓ 1479ms | http |
| 217.76.245.80:999 | ✓ 697ms | ✓ 1538ms | ✓ 1309ms | ✓ 1456ms | ✓ 1529ms | http |
| 14.56.107.244:3128 | ✓ 652ms | ✓ 888ms | ✓ 789ms | ✓ 980ms | ✓ 795ms | http |
| 147.161.160.48:10965 | 否 | ✓ 1693ms | ✓ 1069ms | ✓ 1475ms | ✓ 1158ms | http |
| 147.161.160.48:13711 | 否 | ✓ 1677ms | ✓ 1084ms | ✓ 1485ms | ✓ 1165ms | http |
| 147.161.160.48:3128 | 否 | ✓ 1697ms | ✓ 1063ms | ✓ 1506ms | ✓ 1158ms | http |
| 121.128.121.54:3128 | ✓ 1698ms | ✓ 1528ms | 否 | 否 | ✓ 798ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1474ms | ✓ 967ms | ✓ 1485ms | ✓ 996ms | http |
| 5.75.196.26:40000 | ✓ 609ms | 否 | ✓ 1441ms | 否 | ✓ 1805ms | http |
| 61.72.221.94:3128 | ✓ 665ms | 否 | 否 | ✓ 1051ms | ✓ 1051ms | http |
| 101.132.152.80:9530 | ✓ 908ms | ✓ 1041ms | ✓ 903ms | ✓ 1120ms | ✓ 867ms | http |
| 81.70.169.194:80 | ✓ 954ms | ✓ 1241ms | ✓ 1015ms | ✓ 1207ms | ✓ 1076ms | http |
| 101.43.255.96:80 | ✓ 956ms | ✓ 1261ms | ✓ 998ms | ✓ 1326ms | ✓ 997ms | http |
| 168.235.110.63:3128 | ✓ 1018ms | ✓ 1137ms | ✓ 529ms | ✓ 1160ms | ✓ 889ms | http |
| 35.225.22.61:80 | ✓ 1134ms | ✓ 1194ms | ✓ 1110ms | ✓ 1064ms | ✓ 899ms | http |
| 69.48.179.20:3128 | ✓ 943ms | 否 | ✓ 182ms | ✓ 987ms | 否 | http |
| 165.227.5.10:8888 | ✓ 1035ms | 否 | ✓ 706ms | 否 | ✓ 706ms | http |
| 121.230.8.135:1080 | ✓ 1246ms | ✓ 1510ms | ✓ 1205ms | ✓ 1491ms | 否 | http |
| 91.193.240.157:9877 | ✓ 1193ms | 否 | ✓ 1909ms | 否 | ✓ 1807ms | http |
| 147.161.160.48:11716 | ✓ 1947ms | ✓ 1846ms | ✓ 892ms | ✓ 1487ms | ✓ 1166ms | http |
| 38.180.2.107:3128 | ✓ 869ms | 否 | ✓ 1559ms | 否 | ✓ 1879ms | http |
| 103.139.138.194:3128 | ✓ 1631ms | 否 | ✓ 1507ms | ✓ 1920ms | ✓ 1364ms | http |
| 5.252.33.13:2025 | ✓ 1499ms | 否 | ✓ 1298ms | 否 | ✓ 1794ms | http |
| 1.225.116.115:1080 | ✓ 1703ms | 否 | ✓ 1647ms | 否 | ✓ 1323ms | http |
| 59.46.216.131:30001 | ✓ 1161ms | 否 | ✓ 1060ms | 否 | ✓ 1079ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1223ms | 否 | ✓ 1129ms | ✓ 1140ms | http |
| 103.109.96.86:4429 | ✓ 1711ms | 否 | ✓ 1690ms | ✓ 1877ms | 否 | http |
| 116.80.82.232:3172 | ✓ 1879ms | 否 | 否 | ✓ 1832ms | ✓ 1883ms | http |
| 46.249.103.192:443 | ✓ 1218ms | 否 | ✓ 1765ms | ✓ 1831ms | 否 | http |
| 45.140.147.155:1081 | 否 | ✓ 1758ms | ✓ 688ms | 否 | ✓ 939ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1332ms | 否 | ✓ 1732ms | ✓ 1101ms | http |
| 45.140.147.155:1082 | ✓ 1072ms | ✓ 1304ms | ✓ 1030ms | 否 | 否 | http |
| 121.126.185.63:25152 | ✓ 1957ms | ✓ 1948ms | ✓ 1578ms | 否 | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1801ms | ✓ 1787ms | ✓ 1808ms | http |
| 115.231.181.40:8128 | ✓ 1004ms | ✓ 1197ms | ✓ 1020ms | ✓ 1343ms | ✓ 1068ms | http |
| 103.82.23.118:5253 | ✓ 1612ms | 否 | ✓ 1639ms | ✓ 1962ms | ✓ 1258ms | http |
| 154.37.208.132:30000 | ✓ 933ms | ✓ 1911ms | 否 | ✓ 1884ms | 否 | http |
| 210.223.44.230:3128 | ✓ 881ms | ✓ 821ms | 否 | 否 | ✓ 774ms | http |
| 202.73.27.123:8080 | ✓ 1808ms | 否 | ✓ 1253ms | ✓ 1354ms | ✓ 1384ms | http |
| 147.161.160.48:11272 | ✓ 840ms | 否 | ✓ 835ms | ✓ 1476ms | ✓ 1159ms | http |
| 147.161.160.48:10810 | ✓ 853ms | 否 | ✓ 837ms | ✓ 1494ms | ✓ 1161ms | http |
| 94.72.109.169:8080 | ✓ 781ms | ✓ 1995ms | 否 | ✓ 1801ms | ✓ 1638ms | http |
| 45.136.198.40:3128 | ✓ 1165ms | ✓ 1989ms | ✓ 1509ms | 否 | ✓ 1854ms | http |
| 91.107.175.112:10801 | ✓ 1364ms | 否 | ✓ 1572ms | 否 | ✓ 1518ms | http |
| 121.230.8.153:1080 | ✓ 970ms | ✓ 1323ms | ✓ 1048ms | ✓ 1246ms | ✓ 1019ms | http |
| 154.12.231.32:80 | ✓ 1405ms | 否 | 否 | ✓ 978ms | ✓ 1056ms | http |
| 121.204.158.249:3128 | ✓ 1063ms | ✓ 1291ms | ✓ 1054ms | ✓ 1491ms | ✓ 1030ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1226ms | ✓ 926ms | ✓ 1251ms | ✓ 961ms | http |
| 121.230.8.49:1080 | ✓ 1357ms | ✓ 1753ms | ✓ 1129ms | ✓ 1792ms | ✓ 1423ms | http |
| 101.32.244.83:8080 | ✓ 985ms | ✓ 1793ms | ✓ 952ms | ✓ 1443ms | ✓ 1323ms | http |
| 121.43.196.213:8222 | ✓ 941ms | ✓ 1113ms | ✓ 895ms | ✓ 1154ms | ✓ 948ms | http |
| 121.43.196.210:8222 | ✓ 967ms | ✓ 1085ms | ✓ 896ms | ✓ 1234ms | ✓ 891ms | http |
| 114.55.226.123:10086 | ✓ 1160ms | ✓ 1379ms | ✓ 1057ms | ✓ 1241ms | ✓ 1080ms | http |
| 103.215.36.88:17977 | ✓ 1237ms | ✓ 1528ms | ✓ 1004ms | ✓ 1386ms | ✓ 1053ms | http |
| 183.237.195.130:3128 | ✓ 1303ms | ✓ 1295ms | ✓ 1463ms | ✓ 1288ms | ✓ 1171ms | http |
| 74.48.78.224:2080 | ✓ 1935ms | ✓ 1707ms | ✓ 1708ms | 否 | ✓ 789ms | http |
| 138.124.53.25:7443 | ✓ 561ms | 否 | ✓ 1883ms | 否 | ✓ 1724ms | http |
| 45.140.147.82:1082 | ✓ 629ms | ✓ 1494ms | ✓ 1454ms | ✓ 1853ms | ✓ 1198ms | http |
| 45.140.147.82:1081 | ✓ 624ms | 否 | ✓ 957ms | ✓ 1847ms | ✓ 1198ms | http |
| 83.219.250.8:62920 | ✓ 582ms | ✓ 1540ms | ✓ 1133ms | 否 | ✓ 1706ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1823ms | ✓ 1626ms | ✓ 1307ms | ✓ 1224ms | http |
| 172.212.68.37:3128 | ✓ 968ms | 否 | ✓ 1170ms | ✓ 1806ms | ✓ 1109ms | http |
| 103.39.49.98:8080 | 否 | 否 | ✓ 1255ms | ✓ 1576ms | ✓ 1332ms | http |
| 34.96.238.40:8080 | ✓ 1375ms | ✓ 1348ms | ✓ 1124ms | 否 | ✓ 1049ms | http |
| 149.62.191.202:3128 | ✓ 966ms | 否 | ✓ 1964ms | ✓ 1902ms | ✓ 1339ms | http |
| 61.109.216.213:8080 | ✓ 1470ms | 否 | ✓ 771ms | ✓ 1010ms | 否 | http |
| 192.166.82.55:1080 | 否 | 否 | ✓ 1526ms | ✓ 1538ms | ✓ 1488ms | http |
| 103.82.23.118:5247 | 否 | 否 | ✓ 1642ms | ✓ 1905ms | ✓ 1648ms | http |
| 103.39.51.190:8080 | ✓ 1781ms | 否 | 否 | ✓ 1296ms | ✓ 1378ms | http |
| 212.175.29.184:8080 | ✓ 1206ms | 否 | ✓ 1918ms | ✓ 1714ms | ✓ 1989ms | http |
| 154.90.48.209:9090 | ✓ 1839ms | 否 | ✓ 1536ms | ✓ 1471ms | ✓ 1013ms | http |
| 62.113.119.14:8080 | ✓ 1729ms | 否 | 否 | ✓ 1517ms | ✓ 1154ms | http |
| 103.82.23.118:5178 | ✓ 1639ms | 否 | ✓ 1182ms | ✓ 1574ms | ✓ 1231ms | http |
| 116.80.82.224:3172 | ✓ 1844ms | 否 | ✓ 1576ms | 否 | ✓ 1969ms | http |
| 116.80.82.231:3172 | ✓ 1583ms | 否 | ✓ 1943ms | 否 | ✓ 1913ms | http |
| 88.80.150.82:8080 | ✓ 856ms | ✓ 1709ms | ✓ 891ms | ✓ 1763ms | ✓ 1385ms | https |
| 116.80.82.216:3172 | ✓ 1944ms | 否 | 否 | ✓ 1879ms | ✓ 1670ms | http |
| 116.80.82.227:3172 | 否 | 否 | ✓ 1681ms | ✓ 1885ms | ✓ 1674ms | http |
| 116.80.82.223:3172 | ✓ 1954ms | 否 | ✓ 1740ms | 否 | ✓ 1693ms | http |
| 116.80.82.221:3172 | ✓ 1896ms | 否 | 否 | ✓ 1867ms | ✓ 1661ms | http |

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
