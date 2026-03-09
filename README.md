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

最后更新：2026-03-09 00:19:21 UTC（2026-03-09 08:19:21 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 327ms | ✓ 1936ms | ✓ 927ms | ✓ 1078ms | ✓ 807ms | http |
| 1.231.81.166:3128 | ✓ 1633ms | ✓ 1190ms | ✓ 1159ms | ✓ 1198ms | ✓ 936ms | http |
| 45.140.147.155:1081 | ✓ 431ms | ✓ 1574ms | ✓ 1263ms | 否 | 否 | http |
| 178.236.245.59:3128 | ✓ 967ms | 否 | ✓ 1622ms | ✓ 1523ms | ✓ 1176ms | http |
| 178.236.245.17:3128 | ✓ 1004ms | 否 | ✓ 1621ms | ✓ 1536ms | ✓ 1185ms | http |
| 35.225.22.61:80 | ✓ 608ms | 否 | ✓ 797ms | ✓ 920ms | ✓ 1091ms | http |
| 165.227.5.10:8888 | ✓ 1420ms | ✓ 1274ms | ✓ 1579ms | 否 | ✓ 821ms | http |
| 202.155.12.161:443 | 否 | ✓ 1666ms | ✓ 1405ms | ✓ 1758ms | 否 | http |
| 1.225.116.115:1080 | ✓ 1879ms | 否 | ✓ 1469ms | 否 | ✓ 1491ms | http |
| 14.56.107.244:3128 | ✓ 1876ms | 否 | ✓ 1432ms | 否 | ✓ 1971ms | http |
| 194.213.18.200:443 | ✓ 1716ms | ✓ 1614ms | ✓ 1910ms | 否 | ✓ 1703ms | http |
| 101.43.255.96:80 | ✓ 1214ms | ✓ 1555ms | ✓ 1378ms | ✓ 1899ms | ✓ 1565ms | http |
| 81.70.169.194:80 | ✓ 1149ms | ✓ 1410ms | ✓ 1102ms | ✓ 1748ms | ✓ 1217ms | http |
| 185.115.74.185:8080 | ✓ 1328ms | ✓ 1814ms | ✓ 1463ms | 否 | 否 | http |
| 190.9.109.198:999 | ✓ 729ms | ✓ 1380ms | ✓ 1468ms | ✓ 1457ms | ✓ 1240ms | http |
| 190.9.109.207:999 | ✓ 1686ms | ✓ 1157ms | ✓ 1063ms | ✓ 1186ms | 否 | http |
| 168.235.110.63:3128 | ✓ 488ms | 否 | ✓ 773ms | ✓ 1206ms | ✓ 829ms | http |
| 88.80.150.82:8080 | ✓ 1649ms | 否 | 否 | ✓ 1979ms | ✓ 1726ms | https |
| 103.215.36.88:19290 | ✓ 1298ms | ✓ 1609ms | ✓ 1356ms | 否 | ✓ 1191ms | http |
| 103.165.157.109:3125 | ✓ 1549ms | 否 | ✓ 1460ms | ✓ 1585ms | ✓ 1613ms | http |
| 120.92.212.16:7890 | ✓ 1250ms | ✓ 1444ms | ✓ 1218ms | ✓ 1440ms | ✓ 1145ms | http |
| 120.92.212.16:8890 | ✓ 1191ms | ✓ 1904ms | ✓ 1136ms | ✓ 1657ms | ✓ 1434ms | http |
| 210.223.44.230:3128 | ✓ 1756ms | 否 | ✓ 1362ms | ✓ 1414ms | 否 | http |
| 121.128.121.54:3128 | ✓ 1152ms | ✓ 1888ms | ✓ 1358ms | ✓ 1459ms | ✓ 1307ms | http |
| 138.124.53.25:7443 | ✓ 1749ms | 否 | 否 | ✓ 1559ms | ✓ 1643ms | http |
| 152.42.213.210:8080 | ✓ 943ms | 否 | ✓ 1871ms | ✓ 1373ms | ✓ 1421ms | http |
| 125.128.12.144:3128 | ✓ 1420ms | ✓ 1947ms | ✓ 1434ms | ✓ 1834ms | ✓ 1355ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1449ms | ✓ 1123ms | ✓ 1430ms | ✓ 1134ms | http |
| 113.177.131.2:3128 | 否 | ✓ 1854ms | ✓ 1099ms | ✓ 1759ms | ✓ 1066ms | http |
| 116.80.82.231:3172 | ✓ 1995ms | 否 | ✓ 1941ms | 否 | ✓ 1825ms | http |
| 103.165.157.250:3125 | 否 | 否 | ✓ 1621ms | ✓ 1879ms | ✓ 1658ms | http |
| 150.107.247.195:1080 | 否 | 否 | ✓ 1912ms | ✓ 1671ms | ✓ 1662ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1293ms | ✓ 1142ms | ✓ 1508ms | 否 | http |
| 159.192.133.250:8088 | ✓ 1177ms | 否 | ✓ 1513ms | ✓ 1798ms | ✓ 1339ms | http |
| 115.231.181.40:8128 | ✓ 1040ms | ✓ 1321ms | ✓ 1054ms | ✓ 1416ms | ✓ 1100ms | http |
| 116.80.82.218:3172 | ✓ 1830ms | 否 | ✓ 1738ms | 否 | ✓ 1827ms | http |
| 47.77.193.180:1080 | 否 | ✓ 943ms | ✓ 594ms | ✓ 936ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1176ms | ✓ 1645ms | 否 | ✓ 1804ms | 否 | http |
| 162.248.165.72:1080 | ✓ 1547ms | 否 | ✓ 1612ms | ✓ 1966ms | ✓ 1793ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1758ms | 否 | ✓ 1120ms | ✓ 802ms | http |
| 45.140.147.155:1082 | ✓ 520ms | 否 | ✓ 1219ms | ✓ 1755ms | ✓ 1733ms | http |
| 150.249.255.91:3128 | ✓ 1295ms | ✓ 1126ms | ✓ 719ms | ✓ 1096ms | 否 | http |
| 103.183.10.169:3125 | ✓ 1958ms | 否 | ✓ 1562ms | ✓ 1606ms | ✓ 1633ms | http |
| 159.89.31.62:8080 | ✓ 466ms | ✓ 1561ms | ✓ 1798ms | ✓ 1422ms | ✓ 1358ms | http |
| 91.233.223.147:3128 | ✓ 765ms | ✓ 1863ms | ✓ 1736ms | ✓ 1806ms | ✓ 1424ms | http |
| 223.16.170.103:80 | 否 | ✓ 1849ms | ✓ 1285ms | 否 | ✓ 1359ms | http |
| 45.129.141.143:3128 | ✓ 1193ms | 否 | ✓ 1794ms | ✓ 1939ms | ✓ 1387ms | http |
| 38.180.2.107:3128 | ✓ 1172ms | ✓ 1975ms | 否 | ✓ 1894ms | ✓ 1588ms | http |
| 103.215.36.88:19067 | ✓ 1237ms | ✓ 1907ms | ✓ 1327ms | ✓ 1598ms | ✓ 1254ms | http |
| 34.101.184.164:3128 | ✓ 1161ms | 否 | ✓ 1806ms | ✓ 1393ms | ✓ 1161ms | http |
| 103.215.36.88:17280 | 否 | ✓ 1479ms | ✓ 1360ms | 否 | ✓ 1274ms | http |
| 120.79.99.232:8099 | ✓ 1427ms | ✓ 1728ms | ✓ 1360ms | ✓ 1668ms | 否 | http |
| 62.113.119.14:8080 | ✓ 622ms | 否 | ✓ 744ms | ✓ 1596ms | ✓ 1352ms | http |
| 107.172.125.217:3128 | 否 | 否 | ✓ 1092ms | ✓ 948ms | ✓ 744ms | http |
| 46.249.103.192:443 | ✓ 1640ms | 否 | ✓ 1577ms | ✓ 1901ms | 否 | http |
| 27.254.99.183:8118 | ✓ 1687ms | 否 | ✓ 1472ms | ✓ 1523ms | 否 | http |
| 67.169.98.211:443 | ✓ 1883ms | 否 | ✓ 869ms | ✓ 1818ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1757ms | ✓ 1775ms | 否 | 否 | ✓ 1593ms | http |
| 61.72.110.54:3128 | ✓ 1638ms | ✓ 1246ms | ✓ 1418ms | ✓ 1379ms | ✓ 1880ms | http |
| 45.136.198.40:3128 | ✓ 670ms | ✓ 1396ms | 否 | ✓ 1999ms | ✓ 1677ms | http |
| 103.139.138.194:3128 | ✓ 1994ms | 否 | 否 | ✓ 1559ms | ✓ 1243ms | http |
| 180.103.19.249:1080 | ✓ 1498ms | ✓ 1663ms | ✓ 1402ms | ✓ 1585ms | ✓ 1225ms | http |
| 121.40.231.103:7890 | ✓ 1095ms | ✓ 1256ms | ✓ 1045ms | ✓ 1476ms | ✓ 1142ms | http |
| 172.212.68.37:3128 | 否 | ✓ 1515ms | ✓ 728ms | ✓ 1153ms | ✓ 1054ms | http |
| 20.120.225.109:3128 | ✓ 800ms | ✓ 1521ms | ✓ 1703ms | 否 | 否 | http |
| 47.95.231.180:8084 | ✓ 1143ms | ✓ 1485ms | ✓ 1073ms | ✓ 1439ms | ✓ 1118ms | http |
| 47.110.42.192:9003 | ✓ 1901ms | ✓ 1945ms | ✓ 1888ms | 否 | 否 | http |

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
