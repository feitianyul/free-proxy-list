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

最后更新：2026-05-01 20:56:36 UTC（2026-05-02 04:56:36 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.211.166.184:8081 | ✓ 1250ms | 否 | ✓ 743ms | ✓ 904ms | ✓ 721ms | http |
| 1.231.81.166:3128 | ✓ 1280ms | ✓ 1342ms | ✓ 1697ms | ✓ 969ms | ✓ 1002ms | http |
| 34.96.238.40:8080 | ✓ 993ms | ✓ 1299ms | 否 | ✓ 1055ms | ✓ 1454ms | http |
| 223.84.151.86:30005 | ✓ 1229ms | ✓ 1262ms | ✓ 1420ms | ✓ 1508ms | ✓ 1413ms | http |
| 113.160.132.26:8080 | ✓ 1697ms | ✓ 1359ms | ✓ 1412ms | ✓ 1226ms | ✓ 993ms | http |
| 218.108.131.186:17890 | ✓ 866ms | ✓ 1092ms | ✓ 893ms | 否 | 否 | http |
| 103.125.181.135:9999 | ✓ 1600ms | 否 | ✓ 1618ms | ✓ 1463ms | ✓ 1158ms | http |
| 45.167.124.71:999 | ✓ 1251ms | ✓ 1895ms | ✓ 1681ms | 否 | ✓ 1694ms | http |
| 47.85.51.197:1080 | ✓ 1269ms | ✓ 1462ms | ✓ 350ms | ✓ 1142ms | ✓ 836ms | http |
| 86.104.72.219:1081 | ✓ 459ms | ✓ 1545ms | ✓ 934ms | ✓ 1168ms | ✓ 1398ms | http |
| 46.101.95.183:8888 | ✓ 843ms | 否 | ✓ 680ms | ✓ 1749ms | ✓ 1344ms | http |
| 45.140.147.155:1082 | ✓ 694ms | ✓ 1854ms | ✓ 1425ms | 否 | ✓ 1527ms | http |
| 212.58.132.5:8888 | ✓ 1291ms | 否 | ✓ 998ms | ✓ 1468ms | ✓ 1158ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1355ms | ✓ 1214ms | ✓ 1375ms | ✓ 1180ms | http |
| 175.194.173.105:3128 | ✓ 729ms | ✓ 1334ms | ✓ 1458ms | 否 | 否 | http |
| 45.153.231.229:8080 | ✓ 844ms | 否 | ✓ 1866ms | ✓ 1935ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1569ms | 否 | ✓ 1100ms | ✓ 1519ms | ✓ 1311ms | http |
| 120.92.108.86:7890 | ✓ 1374ms | 否 | ✓ 1270ms | ✓ 1632ms | ✓ 1339ms | http |
| 91.184.241.12:443 | ✓ 1384ms | 否 | ✓ 1656ms | 否 | ✓ 1930ms | http |
| 20.127.128.70:8080 | ✓ 945ms | ✓ 1536ms | ✓ 1147ms | ✓ 1265ms | ✓ 1035ms | http |
| 206.206.126.177:2412 | ✓ 761ms | ✓ 1628ms | ✓ 1380ms | ✓ 1028ms | ✓ 809ms | http |
| 44.201.32.14:50917 | 否 | 否 | ✓ 544ms | ✓ 1760ms | ✓ 1276ms | http |
| 18.222.132.180:32158 | ✓ 1214ms | 否 | ✓ 731ms | 否 | ✓ 1681ms | http |
| 54.67.110.244:9316 | ✓ 1140ms | 否 | ✓ 1185ms | 否 | ✓ 1767ms | http |
| 99.79.58.74:23397 | ✓ 1141ms | 否 | ✓ 993ms | 否 | ✓ 1979ms | http |
| 51.44.97.6:20383 | ✓ 1280ms | 否 | ✓ 702ms | ✓ 1930ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1669ms | 否 | ✓ 1960ms | ✓ 1813ms | ✓ 1717ms | http |
| 148.230.4.241:999 | ✓ 922ms | ✓ 1550ms | ✓ 757ms | ✓ 1512ms | ✓ 1268ms | http |
| 103.167.170.125:8080 | ✓ 1649ms | 否 | 否 | ✓ 1660ms | ✓ 1428ms | http |
| 86.104.72.220:1081 | ✓ 393ms | ✓ 1155ms | ✓ 880ms | ✓ 1408ms | ✓ 878ms | http |
| 115.231.181.40:8128 | ✓ 1316ms | ✓ 1223ms | 否 | 否 | ✓ 1018ms | http |
| 120.92.212.16:8890 | ✓ 968ms | ✓ 1195ms | ✓ 1010ms | ✓ 1172ms | ✓ 1738ms | http |
| 39.107.58.214:7818 | 否 | 否 | ✓ 1755ms | ✓ 1248ms | ✓ 985ms | http |
| 120.92.212.16:7890 | ✓ 1135ms | 否 | ✓ 1110ms | ✓ 1184ms | ✓ 1053ms | http |
| 103.157.200.126:3128 | ✓ 1585ms | 否 | ✓ 1911ms | 否 | ✓ 1546ms | http |
| 89.208.106.138:10808 | 否 | ✓ 1998ms | ✓ 637ms | 否 | ✓ 1120ms | http |
| 86.104.72.220:1082 | ✓ 446ms | ✓ 1310ms | ✓ 290ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 696ms | ✓ 1114ms | 否 | ✓ 1012ms | ✓ 837ms | http |
| 80.92.204.47:1081 | ✓ 1093ms | ✓ 1505ms | ✓ 1638ms | 否 | ✓ 1366ms | http |
| 130.61.174.200:1080 | ✓ 1102ms | 否 | 否 | ✓ 1764ms | ✓ 1350ms | http |
| 98.153.152.141:7070 | 否 | 否 | ✓ 1362ms | ✓ 1254ms | ✓ 729ms | http |
| 158.160.215.167:8124 | ✓ 1250ms | 否 | ✓ 1569ms | 否 | ✓ 1680ms | http |
| 150.107.140.238:3128 | ✓ 1216ms | 否 | ✓ 1071ms | ✓ 1339ms | ✓ 1972ms | http |
| 77.232.142.164:3128 | ✓ 1295ms | ✓ 1925ms | ✓ 1107ms | 否 | ✓ 1610ms | http |
| 62.113.119.14:8080 | ✓ 1143ms | ✓ 1592ms | ✓ 1118ms | ✓ 1643ms | ✓ 1257ms | http |
| 92.119.127.213:6005 | ✓ 1029ms | ✓ 1737ms | ✓ 1909ms | 否 | ✓ 1742ms | http |
| 77.110.107.80:1080 | ✓ 639ms | ✓ 1840ms | ✓ 1216ms | ✓ 1810ms | 否 | http |
| 172.105.118.164:3128 | ✓ 1463ms | 否 | 否 | ✓ 1773ms | ✓ 1322ms | http |
| 121.230.8.237:1080 | ✓ 1114ms | ✓ 1537ms | ✓ 1112ms | 否 | 否 | http |
| 222.107.27.7:8053 | 否 | ✓ 1298ms | ✓ 905ms | ✓ 1013ms | ✓ 808ms | http |
| 72.11.151.159:6005 | ✓ 1955ms | ✓ 1933ms | ✓ 1069ms | ✓ 1320ms | ✓ 1195ms | http |
| 72.11.150.178:6005 | ✓ 1694ms | 否 | ✓ 1603ms | ✓ 1341ms | ✓ 1011ms | http |
| 120.79.99.232:8099 | ✓ 1197ms | ✓ 1387ms | ✓ 1177ms | ✓ 1312ms | ✓ 1107ms | http |
| 8.154.21.175:3128 | ✓ 1943ms | ✓ 1096ms | ✓ 1907ms | ✓ 1154ms | ✓ 932ms | http |
| 121.230.8.153:1080 | ✓ 1055ms | ✓ 1817ms | ✓ 1311ms | ✓ 1439ms | ✓ 1081ms | http |
| 136.244.96.236:50000 | ✓ 1158ms | 否 | ✓ 1622ms | 否 | ✓ 1874ms | http |
| 92.119.127.211:6005 | ✓ 1802ms | 否 | ✓ 1687ms | ✓ 1752ms | 否 | http |
| 43.133.44.89:8888 | 否 | ✓ 1532ms | 否 | ✓ 1758ms | ✓ 1513ms | http |
| 47.112.25.109:7890 | ✓ 1907ms | 否 | 否 | ✓ 1192ms | ✓ 1893ms | http |
| 35.182.12.78:5000 | ✓ 1299ms | 否 | ✓ 1955ms | 否 | ✓ 1855ms | http |
| 16.18.37.186:41511 | ✓ 1263ms | 否 | ✓ 1825ms | 否 | ✓ 1536ms | http |
| 15.237.108.20:56423 | ✓ 1265ms | 否 | ✓ 1800ms | ✓ 1921ms | ✓ 1916ms | http |
| 16.62.229.137:55285 | ✓ 1259ms | 否 | 否 | ✓ 1990ms | ✓ 1813ms | http |
| 45.140.147.155:1081 | ✓ 1112ms | 否 | ✓ 1912ms | ✓ 1901ms | 否 | http |
| 154.90.48.209:9090 | ✓ 1963ms | 否 | ✓ 1396ms | ✓ 1372ms | ✓ 1082ms | http |
| 185.125.203.192:3128 | 否 | ✓ 1864ms | ✓ 528ms | ✓ 1618ms | ✓ 1339ms | http |
| 165.99.192.23:1111 | 否 | 否 | ✓ 1658ms | ✓ 1425ms | ✓ 1412ms | http |
| 119.18.146.65:10001 | ✓ 1569ms | 否 | ✓ 1516ms | ✓ 1889ms | ✓ 1814ms | http |
| 103.16.72.157:9108 | ✓ 1544ms | 否 | ✓ 1229ms | ✓ 1521ms | ✓ 1754ms | http |
| 3.101.133.120:80 | 否 | ✓ 727ms | ✓ 118ms | ✓ 778ms | ✓ 614ms | http |
| 34.246.223.187:51598 | ✓ 1964ms | 否 | 否 | ✓ 1861ms | ✓ 1480ms | http |
| 13.60.181.61:30921 | 否 | 否 | ✓ 1353ms | ✓ 1947ms | ✓ 1779ms | http |
| 18.170.25.193:41442 | ✓ 1392ms | 否 | ✓ 1351ms | 否 | ✓ 1659ms | http |
| 52.47.115.41:42682 | ✓ 1343ms | 否 | ✓ 1336ms | 否 | ✓ 1864ms | http |
| 94.131.118.39:1081 | ✓ 1204ms | ✓ 1854ms | ✓ 1330ms | 否 | 否 | http |
| 3.99.158.157:25638 | ✓ 1247ms | 否 | ✓ 999ms | 否 | ✓ 1520ms | http |
| 108.131.109.106:4882 | ✓ 1594ms | 否 | ✓ 1173ms | 否 | ✓ 1893ms | http |
| 3.8.3.11:13811 | ✓ 1975ms | 否 | 否 | ✓ 1906ms | ✓ 1300ms | http |
| 150.249.255.91:3128 | ✓ 1268ms | 否 | 否 | ✓ 900ms | ✓ 688ms | http |
| 8.219.97.248:80 | ✓ 1320ms | 否 | ✓ 1394ms | ✓ 1891ms | 否 | http |
| 162.240.154.26:3128 | ✓ 1547ms | ✓ 1768ms | ✓ 1044ms | 否 | 否 | http |
| 43.167.237.94:3128 | 否 | ✓ 1838ms | ✓ 648ms | ✓ 868ms | ✓ 662ms | http |
| 103.67.46.225:3125 | ✓ 1582ms | 否 | ✓ 1919ms | ✓ 1736ms | 否 | http |
| 61.52.131.172:8443 | ✓ 839ms | ✓ 1222ms | ✓ 972ms | ✓ 1200ms | ✓ 970ms | http |
| 103.172.70.173:8080 | ✓ 1973ms | 否 | 否 | ✓ 1486ms | ✓ 1551ms | http |
| 92.119.127.212:6005 | ✓ 1115ms | ✓ 1589ms | ✓ 1453ms | 否 | ✓ 1876ms | http |
| 94.131.118.39:1082 | 否 | ✓ 1572ms | ✓ 873ms | ✓ 1750ms | 否 | http |
| 20.120.225.109:3128 | ✓ 621ms | 否 | ✓ 1042ms | ✓ 1347ms | 否 | http |

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
