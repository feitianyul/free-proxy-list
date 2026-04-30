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

最后更新：2026-04-30 20:32:18 UTC（2026-05-01 04:32:18 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 268ms | 否 | ✓ 695ms | ✓ 1250ms | ✓ 937ms | http |
| 43.167.237.94:3128 | ✓ 1326ms | ✓ 1946ms | ✓ 672ms | ✓ 1235ms | ✓ 591ms | http |
| 206.206.126.177:2412 | ✓ 703ms | 否 | ✓ 990ms | ✓ 985ms | ✓ 771ms | http |
| 8.154.21.175:3128 | ✓ 801ms | ✓ 1456ms | ✓ 840ms | ✓ 1104ms | ✓ 1296ms | http |
| 137.59.47.73:3128 | ✓ 1332ms | ✓ 1210ms | ✓ 1324ms | ✓ 1265ms | ✓ 1186ms | http |
| 113.160.132.26:8080 | ✓ 1866ms | ✓ 1324ms | 否 | ✓ 1580ms | ✓ 898ms | http |
| 218.72.124.35:40000 | ✓ 1103ms | ✓ 1276ms | ✓ 893ms | ✓ 1349ms | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1421ms | ✓ 1638ms | ✓ 1502ms | 否 | http |
| 45.167.124.71:999 | ✓ 967ms | 否 | ✓ 1640ms | 否 | ✓ 1735ms | http |
| 107.173.160.222:1080 | ✓ 150ms | ✓ 1854ms | ✓ 802ms | ✓ 988ms | ✓ 659ms | http |
| 154.64.232.35:8080 | ✓ 1375ms | ✓ 1267ms | ✓ 477ms | 否 | ✓ 1429ms | http |
| 1.231.81.166:3128 | ✓ 770ms | ✓ 1162ms | ✓ 1299ms | ✓ 990ms | ✓ 1253ms | http |
| 223.84.151.86:30005 | ✓ 1220ms | ✓ 1562ms | ✓ 1396ms | ✓ 1725ms | ✓ 1500ms | http |
| 62.60.237.68:8080 | ✓ 1332ms | ✓ 1804ms | ✓ 1813ms | 否 | ✓ 1868ms | http |
| 120.92.108.86:7890 | ✓ 1866ms | 否 | ✓ 1522ms | 否 | ✓ 1369ms | http |
| 212.58.132.5:8888 | ✓ 1778ms | 否 | ✓ 1636ms | ✓ 1533ms | ✓ 1477ms | http |
| 37.187.109.70:10111 | ✓ 945ms | 否 | ✓ 1162ms | 否 | ✓ 1573ms | http |
| 103.70.114.149:3128 | ✓ 1433ms | 否 | ✓ 1019ms | ✓ 1719ms | ✓ 1485ms | http |
| 210.223.44.230:3128 | ✓ 1478ms | ✓ 1049ms | ✓ 612ms | ✓ 861ms | ✓ 659ms | http |
| 43.133.44.89:8888 | ✓ 1346ms | ✓ 1637ms | 否 | ✓ 982ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1909ms | 否 | ✓ 1411ms | ✓ 1508ms | 否 | http |
| 190.12.150.244:999 | ✓ 1136ms | 否 | ✓ 1110ms | ✓ 1797ms | 否 | http |
| 159.223.225.118:8888 | ✓ 1217ms | ✓ 1725ms | ✓ 1993ms | ✓ 1880ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1201ms | ✓ 1310ms | ✓ 928ms | ✓ 1444ms | 否 | http |
| 122.224.198.218:808 | ✓ 1997ms | 否 | ✓ 1998ms | 否 | ✓ 1942ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1053ms | ✓ 1269ms | ✓ 1013ms | http |
| 45.153.231.229:8080 | ✓ 1114ms | 否 | ✓ 1988ms | 否 | ✓ 1851ms | http |
| 218.108.131.186:17890 | ✓ 874ms | ✓ 1012ms | 否 | ✓ 1088ms | ✓ 867ms | http |
| 103.35.191.138:1082 | ✓ 340ms | ✓ 1201ms | ✓ 955ms | ✓ 1231ms | ✓ 892ms | http |
| 103.35.191.244:1082 | ✓ 303ms | 否 | ✓ 866ms | ✓ 1199ms | 否 | http |
| 80.92.204.47:1081 | ✓ 1081ms | ✓ 1618ms | ✓ 1962ms | ✓ 1857ms | 否 | http |
| 220.197.44.36:3128 | ✓ 1668ms | ✓ 1879ms | ✓ 1547ms | ✓ 1878ms | ✓ 1935ms | http |
| 152.32.132.190:7890 | ✓ 1519ms | 否 | ✓ 806ms | ✓ 1037ms | ✓ 657ms | http |
| 46.101.95.183:8888 | ✓ 1540ms | 否 | ✓ 1264ms | 否 | ✓ 1741ms | http |
| 45.140.147.82:1081 | ✓ 767ms | ✓ 1747ms | ✓ 812ms | ✓ 1607ms | ✓ 1121ms | http |
| 45.140.147.155:1081 | ✓ 853ms | 否 | ✓ 782ms | 否 | ✓ 1633ms | http |
| 2.83.243.148:7777 | ✓ 1694ms | 否 | ✓ 1465ms | 否 | ✓ 1230ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1505ms | ✓ 1982ms | ✓ 1565ms | http |
| 103.67.80.153:8080 | ✓ 1785ms | 否 | ✓ 1476ms | ✓ 1335ms | ✓ 1319ms | http |
| 8.219.97.248:80 | ✓ 1421ms | 否 | ✓ 1650ms | ✓ 1537ms | 否 | http |
| 103.35.190.182:1082 | ✓ 678ms | ✓ 1089ms | ✓ 321ms | ✓ 1549ms | ✓ 1642ms | http |
| 120.92.212.16:7890 | ✓ 1972ms | 否 | ✓ 880ms | ✓ 1527ms | 否 | http |
| 101.36.105.101:9128 | 否 | ✓ 1547ms | ✓ 1930ms | 否 | ✓ 1173ms | http |
| 103.101.193.170:1111 | 否 | 否 | ✓ 1370ms | ✓ 1330ms | ✓ 1320ms | http |
| 62.113.119.14:8080 | ✓ 851ms | ✓ 1677ms | ✓ 1504ms | 否 | 否 | http |
| 89.208.106.138:10808 | ✓ 701ms | ✓ 1798ms | ✓ 1210ms | 否 | 否 | http |
| 113.176.92.71:3128 | ✓ 1584ms | ✓ 1324ms | ✓ 1120ms | 否 | 否 | http |
| 103.184.99.194:8080 | ✓ 1871ms | 否 | ✓ 1630ms | 否 | ✓ 1405ms | http |
| 77.110.116.224:3128 | ✓ 1529ms | 否 | 否 | ✓ 1750ms | ✓ 1330ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1877ms | 否 | ✓ 1243ms | ✓ 1574ms | http |
| 213.35.113.150:6969 | ✓ 1923ms | 否 | ✓ 1769ms | 否 | ✓ 1169ms | http |
| 109.199.125.66:3128 | ✓ 1420ms | 否 | ✓ 1318ms | 否 | ✓ 1697ms | http |
| 34.101.184.164:3128 | ✓ 1682ms | 否 | ✓ 1402ms | ✓ 1185ms | 否 | http |
| 130.61.174.200:1080 | 否 | ✓ 1816ms | 否 | ✓ 1708ms | ✓ 1541ms | http |
| 94.158.219.111:3128 | ✓ 1406ms | 否 | ✓ 1353ms | 否 | ✓ 1717ms | http |
| 154.90.48.209:9090 | ✓ 1556ms | 否 | ✓ 803ms | ✓ 1268ms | ✓ 1108ms | http |
| 20.120.225.109:3128 | ✓ 574ms | ✓ 1053ms | ✓ 912ms | ✓ 1324ms | ✓ 1040ms | http |
| 185.230.190.195:3128 | ✓ 1536ms | 否 | ✓ 1030ms | 否 | ✓ 1687ms | http |
| 3.101.133.120:80 | ✓ 125ms | ✓ 1151ms | ✓ 1480ms | ✓ 994ms | ✓ 753ms | http |
| 185.234.64.62:1081 | ✓ 848ms | ✓ 1438ms | ✓ 589ms | ✓ 1394ms | 否 | http |
| 158.160.215.167:8124 | ✓ 971ms | 否 | ✓ 1241ms | 否 | ✓ 1894ms | http |
| 157.230.38.173:3128 | ✓ 1443ms | 否 | ✓ 709ms | ✓ 1051ms | ✓ 826ms | http |
| 86.104.74.110:1081 | ✓ 1017ms | ✓ 1320ms | ✓ 1038ms | 否 | ✓ 1244ms | http |
| 86.104.72.219:1081 | ✓ 783ms | ✓ 1271ms | 否 | ✓ 1256ms | 否 | http |
| 86.104.74.110:1082 | ✓ 1785ms | ✓ 1537ms | ✓ 651ms | 否 | 否 | http |
| 103.35.190.69:1082 | ✓ 1048ms | ✓ 1414ms | ✓ 952ms | 否 | 否 | http |
| 124.217.77.49:8082 | ✓ 1328ms | 否 | ✓ 1550ms | ✓ 1321ms | 否 | http |
| 123.112.212.33:9000 | ✓ 1267ms | ✓ 1716ms | ✓ 1974ms | ✓ 1597ms | 否 | http |
| 138.197.68.35:4857 | ✓ 616ms | ✓ 1524ms | ✓ 1411ms | 否 | 否 | http |
| 121.230.8.136:1080 | ✓ 1022ms | ✓ 1128ms | ✓ 944ms | ✓ 1444ms | ✓ 1111ms | http |
| 168.110.52.228:3128 | ✓ 1015ms | 否 | ✓ 1319ms | ✓ 927ms | ✓ 721ms | http |
| 103.138.70.165:3129 | ✓ 1929ms | 否 | 否 | ✓ 1981ms | ✓ 1752ms | http |
| 202.40.186.66:43773 | ✓ 1994ms | 否 | ✓ 1978ms | ✓ 1997ms | 否 | http |
| 194.150.220.163:1082 | ✓ 1792ms | 否 | ✓ 1360ms | 否 | ✓ 1606ms | http |
| 65.109.213.99:1080 | ✓ 1016ms | ✓ 1678ms | ✓ 1070ms | 否 | 否 | http |
| 200.174.198.32:8888 | ✓ 1842ms | 否 | ✓ 950ms | 否 | ✓ 1707ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1166ms | ✓ 1327ms | ✓ 1685ms | ✓ 1314ms | http |
| 64.188.77.26:3128 | ✓ 1180ms | 否 | ✓ 1701ms | ✓ 1931ms | 否 | http |
| 180.125.216.109:8118 | 否 | 否 | ✓ 922ms | ✓ 1287ms | ✓ 960ms | http |
| 8.211.166.184:8081 | ✓ 1554ms | ✓ 928ms | ✓ 759ms | ✓ 759ms | ✓ 617ms | http |
| 109.234.38.35:3128 | ✓ 1695ms | ✓ 1725ms | ✓ 1975ms | 否 | ✓ 1807ms | http |
| 103.67.46.225:3125 | ✓ 1722ms | 否 | 否 | ✓ 1699ms | ✓ 1653ms | http |
| 121.230.8.250:1080 | ✓ 1002ms | ✓ 1246ms | ✓ 989ms | ✓ 1243ms | ✓ 1086ms | http |
| 61.52.131.172:8443 | ✓ 910ms | ✓ 1153ms | ✓ 887ms | ✓ 1221ms | ✓ 913ms | http |
| 137.184.0.30:3128 | ✓ 723ms | ✓ 741ms | 否 | ✓ 633ms | ✓ 639ms | http |
| 47.112.25.109:7890 | ✓ 1688ms | ✓ 1567ms | ✓ 1124ms | 否 | ✓ 1077ms | http |
| 103.229.126.221:7890 | ✓ 1196ms | ✓ 1758ms | ✓ 1661ms | 否 | 否 | http |
| 77.110.107.80:1080 | ✓ 1846ms | ✓ 1925ms | ✓ 1130ms | 否 | ✓ 1508ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1935ms | ✓ 1846ms | ✓ 1921ms | ✓ 1030ms | http |

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
