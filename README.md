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

最后更新：2026-03-11 09:43:34 UTC（2026-03-11 17:43:34 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 500ms | ✓ 1111ms | ✓ 308ms | ✓ 938ms | ✓ 743ms | http |
| 45.136.130.175:8443 | ✓ 306ms | ✓ 875ms | ✓ 887ms | ✓ 970ms | ✓ 729ms | http |
| 45.136.131.47:8443 | ✓ 323ms | ✓ 1921ms | ✓ 819ms | ✓ 1156ms | ✓ 739ms | http |
| 194.213.18.200:443 | ✓ 1418ms | 否 | 否 | ✓ 1538ms | ✓ 1626ms | http |
| 1.231.81.166:3128 | ✓ 1894ms | ✓ 1422ms | ✓ 1516ms | ✓ 1324ms | ✓ 962ms | http |
| 185.115.74.185:8080 | ✓ 1194ms | ✓ 1657ms | ✓ 1307ms | 否 | 否 | http |
| 103.84.95.54:7890 | ✓ 880ms | 否 | ✓ 829ms | ✓ 1104ms | ✓ 924ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1752ms | 否 | ✓ 1407ms | ✓ 1901ms | http |
| 217.76.245.80:999 | ✓ 810ms | ✓ 1462ms | ✓ 1404ms | ✓ 1501ms | ✓ 1196ms | http |
| 24.199.124.151:3128 | ✓ 535ms | ✓ 1255ms | ✓ 1160ms | 否 | 否 | http |
| 160.238.65.6:3128 | ✓ 1242ms | 否 | ✓ 688ms | ✓ 1452ms | ✓ 967ms | http |
| 86.53.183.16:1080 | ✓ 1248ms | 否 | ✓ 1229ms | ✓ 1734ms | ✓ 1433ms | http |
| 160.238.65.2:3128 | ✓ 1247ms | ✓ 1692ms | ✓ 991ms | ✓ 1549ms | ✓ 1013ms | http |
| 115.231.181.40:8128 | ✓ 1078ms | 否 | ✓ 1084ms | ✓ 1522ms | ✓ 1029ms | http |
| 160.238.65.7:3128 | ✓ 1801ms | ✓ 1186ms | ✓ 944ms | ✓ 1546ms | ✓ 1016ms | http |
| 95.3.9.78:3128 | ✓ 1297ms | ✓ 1993ms | 否 | 否 | ✓ 1811ms | http |
| 2.63.162.206:3128 | ✓ 1609ms | 否 | ✓ 1925ms | 否 | ✓ 1938ms | http |
| 160.238.65.9:3128 | ✓ 1800ms | 否 | ✓ 454ms | ✓ 1248ms | ✓ 1009ms | http |
| 45.136.130.188:8443 | ✓ 823ms | ✓ 1926ms | ✓ 295ms | ✓ 924ms | ✓ 1001ms | http |
| 45.136.130.191:8443 | ✓ 810ms | ✓ 914ms | ✓ 980ms | ✓ 1610ms | ✓ 1623ms | http |
| 190.9.109.198:999 | ✓ 778ms | 否 | ✓ 1060ms | ✓ 1298ms | ✓ 1196ms | http |
| 35.225.22.61:80 | ✓ 342ms | ✓ 1324ms | ✓ 341ms | ✓ 1111ms | 否 | http |
| 14.225.212.37:7890 | ✓ 1596ms | 否 | ✓ 1104ms | ✓ 1369ms | ✓ 1095ms | http |
| 91.107.141.42:8081 | ✓ 1664ms | ✓ 1939ms | ✓ 1438ms | 否 | 否 | http |
| 46.183.25.8:443 | ✓ 1600ms | 否 | ✓ 1389ms | ✓ 1505ms | 否 | http |
| 43.167.227.161:1080 | ✓ 644ms | 否 | ✓ 780ms | ✓ 1287ms | ✓ 775ms | http |
| 101.43.255.96:80 | ✓ 1161ms | ✓ 1491ms | ✓ 1137ms | ✓ 1423ms | ✓ 1216ms | http |
| 202.155.12.161:443 | ✓ 1480ms | 否 | ✓ 1337ms | ✓ 1135ms | ✓ 1238ms | http |
| 210.223.44.230:3128 | ✓ 777ms | ✓ 1043ms | ✓ 1105ms | ✓ 1267ms | ✓ 1178ms | http |
| 138.124.53.25:7443 | ✓ 575ms | ✓ 1728ms | ✓ 1252ms | ✓ 1685ms | 否 | http |
| 160.238.65.8:3128 | ✓ 1085ms | 否 | 否 | ✓ 1882ms | ✓ 1389ms | http |
| 160.238.65.5:3128 | ✓ 874ms | ✓ 1534ms | 否 | ✓ 1575ms | 否 | http |
| 95.3.9.78:8080 | 否 | 否 | ✓ 1278ms | ✓ 1622ms | ✓ 1410ms | http |
| 39.104.201.40:7890 | ✓ 1112ms | ✓ 1420ms | ✓ 1841ms | 否 | ✓ 1136ms | http |
| 5.252.33.13:2025 | 否 | 否 | ✓ 1172ms | ✓ 1938ms | ✓ 1691ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1395ms | ✓ 1393ms | 否 | ✓ 1108ms | http |
| 103.35.188.243:3128 | ✓ 588ms | 否 | 否 | ✓ 1128ms | ✓ 858ms | http |
| 152.70.98.46:8888 | ✓ 1357ms | ✓ 1322ms | ✓ 1561ms | ✓ 1234ms | ✓ 945ms | http |
| 94.176.3.43:7443 | ✓ 1085ms | ✓ 1707ms | ✓ 1654ms | 否 | 否 | http |
| 190.212.131.238:3128 | ✓ 1036ms | 否 | ✓ 1628ms | ✓ 1883ms | ✓ 1925ms | http |
| 81.70.169.194:80 | ✓ 1101ms | 否 | ✓ 1155ms | 否 | ✓ 1556ms | http |
| 165.227.5.10:8888 | ✓ 947ms | 否 | ✓ 1246ms | 否 | ✓ 739ms | http |
| 178.236.245.17:3128 | ✓ 606ms | ✓ 1618ms | ✓ 624ms | ✓ 1576ms | ✓ 1211ms | http |
| 178.236.245.59:3128 | ✓ 588ms | ✓ 1466ms | ✓ 652ms | ✓ 1533ms | ✓ 1417ms | http |
| 38.180.2.107:3128 | ✓ 763ms | ✓ 1600ms | 否 | 否 | ✓ 1959ms | http |
| 200.174.198.32:8888 | ✓ 1735ms | 否 | ✓ 1989ms | 否 | ✓ 1631ms | http |
| 103.113.70.189:1081 | ✓ 720ms | ✓ 1716ms | ✓ 423ms | ✓ 1121ms | ✓ 1430ms | http |
| 152.42.213.210:8080 | ✓ 1766ms | 否 | ✓ 1188ms | ✓ 1409ms | 否 | http |
| 14.225.222.213:7890 | ✓ 1560ms | ✓ 1799ms | ✓ 1558ms | 否 | 否 | http |
| 185.191.236.162:3128 | ✓ 972ms | ✓ 1528ms | ✓ 1639ms | 否 | ✓ 1470ms | http |
| 160.238.65.3:3128 | ✓ 1721ms | 否 | ✓ 487ms | ✓ 1960ms | ✓ 1475ms | http |
| 42.96.16.158:1311 | ✓ 1619ms | 否 | ✓ 1126ms | ✓ 1513ms | ✓ 1141ms | http |
| 190.6.54.12:6969 | ✓ 1279ms | 否 | ✓ 1543ms | 否 | ✓ 1909ms | http |
| 160.238.65.4:3128 | 否 | ✓ 1556ms | ✓ 401ms | ✓ 1319ms | ✓ 1020ms | http |
| 14.143.222.113:10155 | ✓ 1861ms | 否 | ✓ 1099ms | ✓ 1432ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1204ms | 否 | ✓ 1184ms | ✓ 1506ms | ✓ 1509ms | http |
| 121.43.196.213:8222 | ✓ 1162ms | ✓ 1303ms | ✓ 995ms | ✓ 1284ms | ✓ 1065ms | http |
| 121.43.196.210:8222 | ✓ 1088ms | ✓ 1247ms | ✓ 1079ms | ✓ 1325ms | ✓ 1062ms | http |
| 114.55.226.123:10086 | ✓ 1169ms | 否 | ✓ 1176ms | ✓ 1499ms | ✓ 1199ms | http |
| 59.46.216.131:30001 | ✓ 1429ms | 否 | ✓ 1855ms | 否 | ✓ 1696ms | http |
| 34.101.184.164:3128 | ✓ 1741ms | 否 | 否 | ✓ 1408ms | ✓ 1129ms | http |
| 14.225.222.164:7890 | ✓ 1597ms | 否 | ✓ 1435ms | 否 | ✓ 1687ms | http |
| 45.136.198.40:3128 | ✓ 1120ms | ✓ 1412ms | ✓ 1799ms | ✓ 1709ms | ✓ 1780ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 823ms | ✓ 1151ms | ✓ 951ms | http |
| 162.248.165.72:1080 | ✓ 1080ms | 否 | ✓ 1938ms | 否 | ✓ 1873ms | http |
| 47.77.193.180:1080 | ✓ 326ms | ✓ 1518ms | ✓ 306ms | ✓ 992ms | ✓ 724ms | http |
| 207.254.71.62:8088 | ✓ 626ms | ✓ 1614ms | ✓ 1375ms | ✓ 1706ms | ✓ 1793ms | http |
| 115.190.91.223:7897 | 否 | ✓ 1487ms | 否 | ✓ 1750ms | ✓ 1094ms | http |
| 88.80.150.82:8080 | ✓ 1481ms | 否 | ✓ 1999ms | 否 | ✓ 1552ms | https |
| 35.181.173.74:9075 | ✓ 1801ms | 否 | ✓ 1631ms | ✓ 1946ms | 否 | http |
| 62.113.119.14:8080 | ✓ 846ms | 否 | ✓ 538ms | ✓ 1364ms | ✓ 1048ms | http |
| 2.83.243.148:7777 | 否 | 否 | ✓ 1551ms | ✓ 1616ms | ✓ 1139ms | http |
| 168.235.110.63:3128 | ✓ 185ms | ✓ 976ms | ✓ 378ms | ✓ 966ms | ✓ 733ms | http |
| 158.69.185.37:3129 | ✓ 207ms | ✓ 1215ms | ✓ 418ms | ✓ 992ms | ✓ 1038ms | http |
| 45.136.130.223:8443 | ✓ 615ms | ✓ 1510ms | ✓ 296ms | ✓ 931ms | ✓ 726ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1939ms | ✓ 1961ms | ✓ 1881ms | http |
| 61.52.131.172:8443 | ✓ 1080ms | ✓ 1439ms | ✓ 1074ms | ✓ 1381ms | ✓ 1118ms | http |
| 8.140.104.98:3128 | ✓ 1083ms | ✓ 1390ms | ✓ 1212ms | ✓ 1491ms | ✓ 1107ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1106ms | 否 | ✓ 929ms | ✓ 832ms | http |

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
