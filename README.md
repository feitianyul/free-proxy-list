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

最后更新：2026-05-31 11:52:39 UTC（2026-05-31 19:52:39 UTC+8）

**代理总数：950**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 950 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 950 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 170.106.136.181:31002 | ✓ 295ms | ✓ 1576ms | ✓ 571ms | ✓ 605ms | ✓ 461ms | http |
| 43.161.239.147:8888 | ✓ 646ms | ✓ 1360ms | ✓ 597ms | ✓ 745ms | ✓ 609ms | http |
| 94.158.244.245:1080 | ✓ 307ms | 否 | ✓ 954ms | ✓ 1265ms | ✓ 857ms | http |
| 59.66.24.75:6382 | ✓ 875ms | ✓ 1175ms | ✓ 1055ms | ✓ 1156ms | ✓ 948ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 781ms | ✓ 974ms | ✓ 718ms | http |
| 192.99.8.15:8850 | ✓ 974ms | 否 | ✓ 1188ms | ✓ 1608ms | ✓ 1405ms | http |
| 117.50.248.177:1080 | ✓ 960ms | 否 | ✓ 1002ms | ✓ 1621ms | ✓ 1380ms | http |
| 137.31.47.67:80 | ✓ 1550ms | ✓ 1825ms | ✓ 1163ms | ✓ 1699ms | ✓ 1693ms | http |
| 137.31.47.61:80 | ✓ 1298ms | ✓ 1823ms | ✓ 1143ms | ✓ 1884ms | ✓ 1792ms | http |
| 137.31.45.49:80 | ✓ 1290ms | ✓ 1934ms | ✓ 1724ms | ✓ 1449ms | ✓ 1657ms | http |
| 137.31.45.61:80 | ✓ 1268ms | ✓ 1857ms | ✓ 1909ms | ✓ 1433ms | ✓ 1605ms | http |
| 137.31.47.85:80 | ✓ 1253ms | ✓ 1850ms | ✓ 1904ms | ✓ 1847ms | ✓ 1239ms | http |
| 137.31.47.73:80 | ✓ 1255ms | 否 | ✓ 1172ms | ✓ 1991ms | ✓ 1689ms | http |
| 137.31.47.79:80 | ✓ 1595ms | ✓ 1809ms | ✓ 1918ms | ✓ 1600ms | ✓ 1241ms | http |
| 137.31.45.91:80 | ✓ 1513ms | ✓ 1891ms | 否 | ✓ 1520ms | ✓ 1240ms | http |
| 137.31.45.85:80 | 否 | 否 | ✓ 1221ms | ✓ 1707ms | ✓ 1246ms | http |
| 137.31.45.79:80 | ✓ 1538ms | ✓ 1874ms | ✓ 1787ms | ✓ 1762ms | ✓ 1241ms | http |
| 137.31.45.73:80 | ✓ 1287ms | 否 | ✓ 1922ms | ✓ 1726ms | ✓ 1261ms | http |
| 137.31.47.49:80 | ✓ 1511ms | ✓ 1849ms | ✓ 1900ms | ✓ 1789ms | ✓ 1260ms | http |
| 137.31.45.67:80 | ✓ 1520ms | 否 | ✓ 1805ms | ✓ 1761ms | ✓ 1241ms | http |
| 190.212.131.238:3128 | ✓ 1271ms | 否 | ✓ 1688ms | ✓ 1957ms | ✓ 1703ms | http |
| 137.31.47.55:80 | ✓ 1444ms | ✓ 1820ms | 否 | ✓ 1729ms | ✓ 1457ms | http |
| 137.31.47.91:80 | ✓ 1274ms | ✓ 1790ms | ✓ 1745ms | ✓ 1929ms | ✓ 1812ms | http |
| 137.31.47.79:443 | 否 | 否 | ✓ 1146ms | ✓ 1567ms | ✓ 1236ms | http |
| 194.36.208.108:1080 | ✓ 1435ms | 否 | ✓ 1689ms | ✓ 1878ms | ✓ 1358ms | http |
| 167.86.95.198:3128 | ✓ 1055ms | 否 | 否 | ✓ 1890ms | ✓ 1893ms | http |
| 176.111.37.216:39811 | ✓ 1413ms | ✓ 1640ms | ✓ 1885ms | ✓ 1903ms | 否 | http |
| 137.31.47.73:443 | ✓ 1245ms | 否 | 否 | ✓ 1428ms | ✓ 1773ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1594ms | ✓ 888ms | ✓ 1686ms | 否 | http |
| 2.27.51.203:8080 | ✓ 1441ms | 否 | ✓ 1514ms | 否 | ✓ 1801ms | http |
| 14.143.222.113:57718 | ✓ 1613ms | 否 | ✓ 1144ms | ✓ 1813ms | 否 | http |
| 185.200.188.234:10001 | ✓ 1442ms | 否 | ✓ 1530ms | 否 | ✓ 1769ms | http |
| 14.143.222.113:10158 | ✓ 1620ms | 否 | ✓ 1952ms | ✓ 1859ms | 否 | http |
| 103.174.131.163:3128 | ✓ 1818ms | 否 | ✓ 1734ms | ✓ 1811ms | ✓ 1653ms | http |
| 47.82.151.59:3128 | ✓ 945ms | ✓ 1831ms | ✓ 903ms | ✓ 1328ms | ✓ 1196ms | http |
| 47.82.178.21:3128 | ✓ 955ms | ✓ 1839ms | ✓ 898ms | ✓ 1326ms | ✓ 1188ms | http |
| 47.82.145.66:3128 | ✓ 904ms | ✓ 1846ms | ✓ 1057ms | ✓ 1249ms | ✓ 1198ms | http |
| 47.82.180.80:3128 | ✓ 946ms | ✓ 1851ms | ✓ 890ms | ✓ 1414ms | ✓ 1173ms | http |
| 47.82.145.17:3128 | ✓ 898ms | 否 | ✓ 919ms | ✓ 1307ms | ✓ 1247ms | http |
| 47.82.151.148:3128 | ✓ 1061ms | ✓ 1905ms | ✓ 918ms | ✓ 1244ms | ✓ 1241ms | http |
| 47.82.180.56:3128 | ✓ 975ms | 否 | ✓ 888ms | ✓ 1309ms | ✓ 1216ms | http |
| 47.82.154.105:3128 | ✓ 1099ms | ✓ 1843ms | ✓ 890ms | ✓ 1367ms | ✓ 1237ms | http |
| 47.82.152.191:3128 | ✓ 1290ms | 否 | ✓ 1625ms | ✓ 1365ms | ✓ 1241ms | http |
| 47.82.180.219:3128 | ✓ 948ms | ✓ 1805ms | ✓ 902ms | ✓ 1156ms | ✓ 1189ms | http |
| 47.82.180.243:3128 | ✓ 971ms | ✓ 1802ms | ✓ 893ms | ✓ 1242ms | ✓ 1172ms | http |
| 47.82.151.189:3128 | ✓ 981ms | ✓ 1872ms | ✓ 897ms | ✓ 1136ms | ✓ 1218ms | http |
| 47.82.145.170:3128 | ✓ 895ms | 否 | ✓ 905ms | ✓ 1141ms | ✓ 1183ms | http |
| 47.82.180.167:3128 | ✓ 976ms | ✓ 1835ms | ✓ 893ms | ✓ 1242ms | ✓ 1171ms | http |
| 47.82.145.117:3128 | ✓ 902ms | ✓ 1819ms | ✓ 898ms | ✓ 1311ms | ✓ 1215ms | http |
| 47.82.0.167:3128 | ✓ 895ms | ✓ 1824ms | ✓ 916ms | ✓ 1320ms | ✓ 1214ms | http |
| 47.82.152.109:3128 | ✓ 997ms | ✓ 1817ms | ✓ 900ms | ✓ 1195ms | ✓ 1262ms | http |
| 103.17.140.87:8080 | ✓ 1315ms | ✓ 1779ms | ✓ 1214ms | 否 | ✓ 1368ms | http |
| 47.82.0.83:3128 | ✓ 948ms | ✓ 1992ms | ✓ 896ms | ✓ 1156ms | ✓ 1190ms | http |
| 47.82.152.192:3128 | ✓ 1008ms | ✓ 1786ms | ✓ 899ms | ✓ 1308ms | ✓ 1184ms | http |
| 47.82.154.45:3128 | ✓ 919ms | ✓ 1828ms | ✓ 1089ms | ✓ 1145ms | ✓ 1196ms | http |
| 47.82.151.167:3128 | ✓ 928ms | 否 | ✓ 895ms | ✓ 1138ms | ✓ 1236ms | http |
| 47.82.178.185:3128 | ✓ 933ms | ✓ 1900ms | ✓ 995ms | ✓ 1162ms | ✓ 1184ms | http |
| 47.82.145.99:3128 | ✓ 912ms | 否 | ✓ 910ms | ✓ 1142ms | ✓ 1245ms | http |
| 47.82.178.20:3128 | ✓ 944ms | 否 | ✓ 922ms | ✓ 1163ms | ✓ 1175ms | http |
| 47.82.154.85:3128 | ✓ 958ms | ✓ 1927ms | ✓ 889ms | ✓ 1221ms | ✓ 1219ms | http |
| 47.82.154.214:3128 | ✓ 925ms | ✓ 1849ms | ✓ 901ms | ✓ 1304ms | ✓ 1240ms | http |
| 47.82.151.249:3128 | ✓ 943ms | ✓ 1896ms | ✓ 900ms | ✓ 1304ms | ✓ 1195ms | http |
| 47.82.152.254:3128 | ✓ 938ms | ✓ 1813ms | ✓ 865ms | ✓ 1413ms | ✓ 1216ms | http |
| 47.82.178.175:3128 | ✓ 987ms | ✓ 1781ms | ✓ 1060ms | ✓ 1236ms | ✓ 1171ms | http |
| 47.82.178.204:3128 | ✓ 965ms | 否 | ✓ 902ms | ✓ 1137ms | ✓ 1234ms | http |
| 47.82.180.84:3128 | ✓ 1064ms | ✓ 1844ms | ✓ 932ms | ✓ 1192ms | ✓ 1213ms | http |
| 47.82.152.34:3128 | ✓ 989ms | ✓ 1843ms | ✓ 901ms | ✓ 1305ms | ✓ 1236ms | http |
| 47.82.0.156:3128 | ✓ 1063ms | ✓ 1842ms | ✓ 923ms | ✓ 1245ms | ✓ 1215ms | http |
| 47.82.180.109:3128 | ✓ 974ms | 否 | ✓ 890ms | ✓ 1171ms | ✓ 1235ms | http |
| 47.82.180.74:3128 | ✓ 971ms | ✓ 1852ms | ✓ 896ms | ✓ 1325ms | ✓ 1225ms | http |
| 47.82.154.62:3128 | ✓ 955ms | ✓ 1856ms | ✓ 925ms | ✓ 1332ms | ✓ 1223ms | http |
| 47.82.178.137:3128 | ✓ 945ms | ✓ 1852ms | ✓ 875ms | ✓ 1408ms | ✓ 1213ms | http |
| 47.79.249.204:3128 | ✓ 1017ms | ✓ 1949ms | ✓ 923ms | ✓ 1254ms | ✓ 1182ms | http |
| 47.82.151.70:3128 | ✓ 907ms | 否 | ✓ 1063ms | ✓ 1145ms | ✓ 1199ms | http |
| 47.82.151.186:3128 | ✓ 1004ms | ✓ 1865ms | ✓ 901ms | ✓ 1310ms | ✓ 1241ms | http |
| 47.82.145.150:3128 | ✓ 1001ms | ✓ 1815ms | ✓ 901ms | ✓ 1160ms | ✓ 1441ms | http |
| 47.82.178.30:3128 | ✓ 984ms | ✓ 1905ms | ✓ 887ms | ✓ 1333ms | ✓ 1197ms | http |
| 47.82.178.24:3128 | ✓ 957ms | 否 | ✓ 912ms | ✓ 1245ms | ✓ 1190ms | http |
| 47.82.145.85:3128 | ✓ 1042ms | ✓ 1808ms | ✓ 1069ms | ✓ 1239ms | ✓ 1180ms | http |
| 47.82.151.153:3128 | ✓ 1005ms | ✓ 1856ms | ✓ 905ms | ✓ 1310ms | ✓ 1261ms | http |
| 47.82.180.7:3128 | ✓ 982ms | ✓ 1846ms | ✓ 887ms | ✓ 1411ms | ✓ 1196ms | http |
| 47.82.145.32:3128 | ✓ 906ms | 否 | ✓ 947ms | ✓ 1316ms | ✓ 1179ms | http |
| 47.82.180.238:3128 | ✓ 903ms | 否 | ✓ 891ms | ✓ 1298ms | ✓ 1235ms | http |
| 47.82.151.252:3128 | ✓ 905ms | ✓ 1836ms | ✓ 1083ms | ✓ 1301ms | ✓ 1223ms | http |
| 47.82.178.81:3128 | ✓ 921ms | 否 | ✓ 1057ms | ✓ 1156ms | ✓ 1206ms | http |
| 47.82.151.141:3128 | ✓ 925ms | 否 | ✓ 899ms | ✓ 1316ms | ✓ 1217ms | http |
| 47.82.154.147:3128 | ✓ 1009ms | ✓ 1890ms | ✓ 933ms | ✓ 1317ms | ✓ 1214ms | http |
| 47.82.151.201:3128 | ✓ 932ms | ✓ 1842ms | ✓ 1069ms | ✓ 1310ms | ✓ 1213ms | http |
| 47.82.178.53:3128 | ✓ 983ms | ✓ 1854ms | ✓ 1111ms | ✓ 1166ms | ✓ 1248ms | http |
| 47.82.178.165:3128 | ✓ 920ms | 否 | ✓ 867ms | ✓ 1409ms | ✓ 1174ms | http |
| 47.82.152.64:3128 | ✓ 981ms | 否 | ✓ 906ms | ✓ 1329ms | ✓ 1176ms | http |
| 47.82.154.79:3128 | ✓ 910ms | ✓ 1845ms | ✓ 913ms | ✓ 1482ms | ✓ 1236ms | http |
| 47.82.154.213:3128 | ✓ 930ms | 否 | ✓ 1061ms | ✓ 1171ms | ✓ 1229ms | http |
| 47.82.151.50:3128 | ✓ 987ms | ✓ 1872ms | ✓ 922ms | ✓ 1424ms | ✓ 1199ms | http |
| 47.82.154.159:3128 | ✓ 971ms | 否 | ✓ 906ms | ✓ 1325ms | ✓ 1218ms | http |
| 47.82.178.143:3128 | ✓ 923ms | 否 | ✓ 865ms | ✓ 1430ms | ✓ 1194ms | http |
| 47.82.154.156:3128 | ✓ 986ms | ✓ 1926ms | ✓ 905ms | ✓ 1391ms | ✓ 1217ms | http |
| 47.82.151.76:3128 | ✓ 906ms | 否 | ✓ 899ms | ✓ 1399ms | ✓ 1236ms | http |
| 47.82.145.174:3128 | ✓ 892ms | 否 | ✓ 1018ms | ✓ 1366ms | ✓ 1175ms | http |
| 47.82.178.26:3128 | ✓ 983ms | ✓ 1849ms | ✓ 1069ms | ✓ 1307ms | ✓ 1211ms | http |

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
