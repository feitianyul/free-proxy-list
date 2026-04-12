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

最后更新：2026-04-12 19:47:24 UTC（2026-04-13 03:47:24 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 1108ms | ✓ 1436ms | 否 | ✓ 1152ms | ✓ 845ms | http |
| 147.161.210.140:8800 | ✓ 797ms | 否 | ✓ 1101ms | ✓ 1275ms | ✓ 1229ms | http |
| 218.108.131.186:17890 | ✓ 1025ms | ✓ 1220ms | ✓ 974ms | ✓ 1302ms | ✓ 1077ms | http |
| 167.103.34.108:8800 | ✓ 1749ms | 否 | ✓ 1599ms | ✓ 1423ms | ✓ 1426ms | http |
| 109.107.179.140:31000 | ✓ 1550ms | 否 | ✓ 1941ms | 否 | ✓ 1356ms | http |
| 34.71.229.255:3128 | ✓ 545ms | ✓ 1766ms | ✓ 1070ms | ✓ 1062ms | ✓ 982ms | http |
| 36.103.198.235:7890 | ✓ 1125ms | 否 | ✓ 1567ms | ✓ 1897ms | ✓ 1601ms | http |
| 80.250.165.242:3128 | ✓ 1389ms | ✓ 1777ms | ✓ 1763ms | 否 | ✓ 1582ms | http |
| 167.103.115.102:8800 | ✓ 1342ms | 否 | 否 | ✓ 1276ms | ✓ 1604ms | http |
| 167.103.144.127:8800 | ✓ 1478ms | 否 | ✓ 1748ms | ✓ 1820ms | ✓ 1801ms | http |
| 120.92.108.86:7890 | ✓ 1783ms | 否 | 否 | ✓ 1972ms | ✓ 1501ms | http |
| 137.59.47.73:3128 | ✓ 1960ms | ✓ 1952ms | ✓ 1333ms | 否 | 否 | http |
| 45.167.124.52:8080 | ✓ 493ms | ✓ 1429ms | ✓ 493ms | 否 | 否 | http |
| 45.167.125.21:999 | ✓ 751ms | ✓ 1819ms | ✓ 1942ms | ✓ 1582ms | ✓ 1339ms | http |
| 91.233.223.147:3128 | ✓ 1259ms | ✓ 1974ms | ✓ 751ms | 否 | ✓ 1412ms | http |
| 162.240.154.26:3128 | ✓ 864ms | ✓ 1296ms | ✓ 962ms | 否 | 否 | http |
| 167.103.31.122:8800 | ✓ 1246ms | 否 | ✓ 1423ms | ✓ 1599ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1542ms | 否 | ✓ 1542ms | ✓ 1824ms | ✓ 1242ms | http |
| 223.16.170.103:80 | ✓ 1393ms | 否 | ✓ 1229ms | ✓ 1305ms | ✓ 1328ms | http |
| 5.255.123.43:1080 | ✓ 619ms | ✓ 1227ms | 否 | 否 | ✓ 1260ms | http |
| 36.141.21.200:7890 | ✓ 1167ms | ✓ 1362ms | ✓ 1155ms | 否 | ✓ 1194ms | http |
| 121.230.8.162:1080 | ✓ 1463ms | ✓ 1686ms | ✓ 1417ms | ✓ 1980ms | ✓ 1289ms | http |
| 130.61.30.221:8080 | ✓ 503ms | 否 | ✓ 755ms | ✓ 1715ms | 否 | http |
| 20.120.225.109:3128 | ✓ 512ms | ✓ 1372ms | ✓ 680ms | ✓ 1476ms | ✓ 973ms | http |
| 147.161.239.240:8800 | ✓ 618ms | ✓ 1456ms | ✓ 680ms | ✓ 1538ms | ✓ 1215ms | http |
| 8.219.195.129:1080 | ✓ 1270ms | ✓ 1878ms | ✓ 934ms | ✓ 1273ms | ✓ 1006ms | http |
| 212.58.132.5:8888 | ✓ 1497ms | 否 | ✓ 1526ms | 否 | ✓ 1399ms | http |
| 103.85.113.66:9999 | ✓ 488ms | 否 | ✓ 963ms | ✓ 1950ms | ✓ 1497ms | http |
| 177.234.217.88:999 | ✓ 1602ms | ✓ 1752ms | ✓ 1661ms | ✓ 1850ms | ✓ 1737ms | http |
| 43.99.54.236:5555 | ✓ 896ms | ✓ 1163ms | ✓ 930ms | ✓ 1022ms | ✓ 825ms | http |
| 113.160.132.26:8080 | ✓ 1713ms | ✓ 1668ms | ✓ 1750ms | 否 | ✓ 1241ms | http |
| 59.46.216.131:30001 | ✓ 1180ms | ✓ 1608ms | 否 | ✓ 1508ms | ✓ 1333ms | http |
| 168.222.254.136:8888 | ✓ 1087ms | ✓ 1567ms | ✓ 1852ms | ✓ 1983ms | 否 | http |
| 192.71.213.85:5678 | ✓ 1464ms | 否 | ✓ 1536ms | ✓ 1423ms | 否 | http |
| 217.217.249.160:8080 | ✓ 1538ms | 否 | ✓ 1301ms | 否 | ✓ 1726ms | http |
| 1.231.81.166:3128 | ✓ 1668ms | ✓ 1551ms | ✓ 1708ms | ✓ 1284ms | ✓ 1028ms | http |
| 118.113.246.123:1080 | 否 | 否 | ✓ 1536ms | ✓ 1728ms | ✓ 1701ms | http |
| 115.231.181.40:8128 | ✓ 1799ms | ✓ 1317ms | ✓ 997ms | 否 | ✓ 1417ms | http |
| 171.227.167.109:1004 | ✓ 1749ms | 否 | ✓ 1142ms | ✓ 1462ms | 否 | http |
| 148.153.56.51:80 | ✓ 645ms | ✓ 1035ms | ✓ 1092ms | ✓ 1100ms | ✓ 727ms | http |
| 58.69.114.117:5050 | 否 | 否 | ✓ 1517ms | ✓ 1551ms | ✓ 1551ms | http |
| 103.235.67.190:80 | ✓ 1465ms | 否 | ✓ 954ms | ✓ 1438ms | ✓ 1122ms | http |
| 121.230.8.111:1080 | ✓ 1296ms | ✓ 1708ms | ✓ 1123ms | 否 | 否 | http |
| 79.132.136.58:3128 | ✓ 974ms | ✓ 1545ms | ✓ 702ms | ✓ 1176ms | ✓ 1658ms | http |
| 178.159.94.76:3128 | ✓ 748ms | ✓ 1754ms | ✓ 1379ms | ✓ 1578ms | ✓ 1570ms | http |
| 46.30.46.133:3128 | ✓ 1656ms | 否 | ✓ 1212ms | ✓ 1952ms | 否 | http |
| 114.237.77.202:1080 | ✓ 1283ms | ✓ 1629ms | ✓ 1160ms | ✓ 1424ms | ✓ 1142ms | http |
| 121.230.9.37:1080 | ✓ 1523ms | ✓ 1829ms | ✓ 1334ms | 否 | ✓ 1422ms | http |
| 103.157.200.126:3128 | ✓ 1173ms | 否 | ✓ 1387ms | ✓ 1505ms | ✓ 1221ms | http |
| 8.209.238.110:47701 | ✓ 1141ms | 否 | ✓ 824ms | ✓ 1054ms | ✓ 854ms | http |
| 95.214.9.93:3128 | ✓ 1044ms | 否 | ✓ 972ms | ✓ 1865ms | ✓ 1438ms | http |
| 5.196.101.18:3128 | ✓ 1848ms | 否 | ✓ 575ms | ✓ 1905ms | ✓ 1365ms | http |
| 168.110.52.228:3128 | ✓ 622ms | 否 | 否 | ✓ 974ms | ✓ 771ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1156ms | ✓ 1320ms | 否 | ✓ 930ms | http |
| 37.187.109.70:10111 | ✓ 1114ms | ✓ 1758ms | ✓ 1951ms | 否 | 否 | http |
| 171.227.167.109:1006 | ✓ 1213ms | 否 | ✓ 1641ms | ✓ 1489ms | ✓ 1221ms | http |
| 171.227.167.109:1009 | ✓ 1672ms | 否 | ✓ 1209ms | ✓ 1449ms | ✓ 1317ms | http |
| 38.180.2.107:3128 | ✓ 662ms | ✓ 1675ms | 否 | 否 | ✓ 1994ms | http |
| 110.42.37.202:20005 | 否 | ✓ 1885ms | ✓ 1760ms | 否 | ✓ 1908ms | http |
| 45.149.92.147:5001 | ✓ 811ms | 否 | ✓ 791ms | ✓ 1023ms | ✓ 795ms | http |
| 171.227.167.109:1005 | ✓ 1360ms | 否 | ✓ 1953ms | 否 | ✓ 1649ms | http |
| 38.191.200.253:999 | 否 | ✓ 1647ms | 否 | ✓ 1812ms | ✓ 1336ms | http |
| 91.132.161.200:3128 | ✓ 665ms | ✓ 1787ms | ✓ 1390ms | ✓ 1767ms | ✓ 1341ms | http |
| 124.243.150.41:3128 | ✓ 911ms | ✓ 1729ms | ✓ 1188ms | ✓ 1267ms | ✓ 1074ms | http |
| 8.219.64.245:3128 | ✓ 899ms | ✓ 1963ms | ✓ 1080ms | ✓ 1271ms | ✓ 1018ms | http |
| 14.29.168.215:1080 | ✓ 1078ms | ✓ 1477ms | ✓ 1248ms | ✓ 1559ms | ✓ 1107ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1301ms | ✓ 1481ms | ✓ 1140ms | http |
| 121.230.9.203:1080 | ✓ 1299ms | ✓ 1662ms | ✓ 1318ms | ✓ 1607ms | ✓ 1213ms | http |
| 106.10.55.212:1121 | ✓ 1396ms | ✓ 1484ms | ✓ 1027ms | 否 | 否 | http |
| 105.159.164.192:4832 | ✓ 1100ms | ✓ 1397ms | ✓ 1832ms | 否 | ✓ 1242ms | http |
| 187.102.195.53:999 | ✓ 1529ms | ✓ 1897ms | ✓ 663ms | ✓ 1826ms | 否 | http |
| 51.79.207.21:8080 | ✓ 1215ms | 否 | ✓ 1254ms | ✓ 1454ms | ✓ 1124ms | http |
| 114.237.77.231:1080 | ✓ 1057ms | ✓ 1375ms | ✓ 1021ms | ✓ 1415ms | ✓ 1179ms | http |
| 103.39.51.207:8080 | ✓ 1560ms | 否 | 否 | ✓ 1863ms | ✓ 1565ms | http |
| 103.217.224.75:3125 | 否 | 否 | ✓ 1965ms | ✓ 1670ms | ✓ 1612ms | http |

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
