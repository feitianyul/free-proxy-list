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

最后更新：2026-03-16 15:56:17 UTC（2026-03-16 23:56:17 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1329ms | 否 | ✓ 1019ms | ✓ 1454ms | ✓ 1245ms | http |
| 1.231.81.166:3128 | ✓ 1363ms | 否 | 否 | ✓ 1351ms | ✓ 1076ms | http |
| 113.160.132.26:8080 | ✓ 1972ms | 否 | ✓ 1404ms | ✓ 1598ms | ✓ 1171ms | http |
| 35.225.22.61:80 | ✓ 1072ms | 否 | ✓ 1086ms | ✓ 1281ms | 否 | http |
| 217.76.245.80:999 | ✓ 792ms | ✓ 1479ms | ✓ 1126ms | ✓ 1750ms | ✓ 1459ms | http |
| 62.60.177.204:34094 | ✓ 1121ms | ✓ 1253ms | ✓ 1451ms | 否 | ✓ 1385ms | http |
| 202.155.12.161:443 | ✓ 1671ms | 否 | ✓ 1308ms | ✓ 1503ms | ✓ 1120ms | http |
| 219.117.204.211:7799 | ✓ 1689ms | ✓ 1753ms | ✓ 1190ms | ✓ 1403ms | ✓ 1191ms | http |
| 149.50.116.240:1080 | 否 | 否 | ✓ 1416ms | ✓ 1572ms | ✓ 1546ms | http |
| 137.220.150.152:6005 | ✓ 1380ms | 否 | ✓ 1704ms | ✓ 1747ms | ✓ 1638ms | http |
| 137.220.151.110:6005 | ✓ 1625ms | 否 | ✓ 998ms | ✓ 1430ms | 否 | http |
| 103.82.23.118:5247 | ✓ 1647ms | 否 | ✓ 1626ms | ✓ 1806ms | ✓ 1796ms | http |
| 186.148.180.46:999 | ✓ 608ms | 否 | ✓ 561ms | ✓ 1564ms | ✓ 1653ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1537ms | 否 | ✓ 1494ms | ✓ 1334ms | http |
| 45.167.124.52:8080 | ✓ 1273ms | 否 | ✓ 1188ms | ✓ 1580ms | ✓ 1296ms | http |
| 45.136.198.40:3128 | ✓ 682ms | ✓ 1481ms | 否 | ✓ 1806ms | 否 | http |
| 117.159.239.56:22222 | ✓ 1099ms | ✓ 1251ms | ✓ 1061ms | ✓ 1297ms | ✓ 1020ms | http |
| 137.220.150.104:6005 | ✓ 1008ms | 否 | ✓ 1031ms | ✓ 1854ms | ✓ 1334ms | http |
| 101.43.127.100:8877 | ✓ 1377ms | 否 | ✓ 1353ms | ✓ 1437ms | ✓ 1006ms | http |
| 190.242.157.215:8080 | ✓ 1507ms | 否 | 否 | ✓ 1761ms | ✓ 1570ms | http |
| 115.231.181.40:8128 | ✓ 1084ms | 否 | ✓ 1736ms | ✓ 1508ms | ✓ 1178ms | http |
| 117.159.239.52:22222 | ✓ 1050ms | 否 | 否 | ✓ 1347ms | ✓ 1044ms | http |
| 183.249.5.111:22222 | ✓ 894ms | ✓ 1535ms | 否 | 否 | ✓ 911ms | http |
| 8.209.239.31:30000 | ✓ 1865ms | ✓ 1848ms | ✓ 717ms | ✓ 1412ms | ✓ 789ms | http |
| 120.92.212.16:8890 | ✓ 1174ms | ✓ 1448ms | ✓ 1182ms | ✓ 1498ms | 否 | http |
| 77.110.113.24:40000 | ✓ 1058ms | 否 | ✓ 1443ms | 否 | ✓ 1781ms | http |
| 8.137.112.117:3128 | ✓ 1092ms | ✓ 1654ms | ✓ 1284ms | ✓ 1707ms | ✓ 1192ms | http |
| 120.92.212.16:7890 | ✓ 1155ms | ✓ 1523ms | 否 | ✓ 1688ms | 否 | http |
| 113.59.32.141:22222 | 否 | ✓ 1538ms | ✓ 1115ms | ✓ 1479ms | ✓ 1209ms | http |
| 120.240.35.178:22222 | ✓ 1752ms | ✓ 1589ms | 否 | ✓ 1424ms | ✓ 1123ms | http |
| 160.250.5.22:1 | ✓ 1944ms | 否 | 否 | ✓ 1453ms | ✓ 1209ms | http |
| 120.240.29.174:22222 | 否 | 否 | ✓ 1637ms | ✓ 1486ms | ✓ 1480ms | http |
| 62.113.119.14:8080 | ✓ 1070ms | 否 | ✓ 903ms | ✓ 1431ms | ✓ 1117ms | http |
| 120.240.35.176:22222 | ✓ 1164ms | 否 | ✓ 1717ms | 否 | ✓ 1821ms | http |
| 160.238.65.6:3128 | ✓ 907ms | 否 | ✓ 1751ms | ✓ 1910ms | 否 | http |
| 160.238.65.3:3128 | ✓ 1861ms | 否 | ✓ 797ms | ✓ 1322ms | ✓ 1273ms | http |
| 160.238.65.4:3128 | ✓ 1944ms | ✓ 1259ms | ✓ 1451ms | ✓ 1603ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1275ms | ✓ 1870ms | ✓ 1132ms | ✓ 1641ms | ✓ 1519ms | http |
| 121.43.196.213:8222 | ✓ 1134ms | ✓ 1321ms | ✓ 1046ms | ✓ 1326ms | ✓ 1089ms | http |
| 121.43.196.210:8222 | ✓ 1147ms | ✓ 1260ms | ✓ 1101ms | ✓ 1370ms | ✓ 1060ms | http |
| 114.55.226.123:10086 | ✓ 1322ms | 否 | ✓ 1173ms | ✓ 1484ms | ✓ 1228ms | http |
| 192.71.213.85:9812 | ✓ 962ms | 否 | ✓ 1855ms | ✓ 1529ms | 否 | http |
| 172.212.68.37:3128 | ✓ 282ms | ✓ 1806ms | ✓ 676ms | ✓ 1494ms | ✓ 1291ms | http |
| 47.77.193.180:1080 | ✓ 963ms | 否 | ✓ 648ms | ✓ 954ms | ✓ 730ms | http |
| 113.59.32.142:22222 | ✓ 1439ms | ✓ 1525ms | ✓ 1307ms | ✓ 1460ms | ✓ 1161ms | http |
| 88.80.150.82:8080 | ✓ 1855ms | ✓ 1859ms | ✓ 1875ms | 否 | ✓ 1822ms | https |
| 14.225.212.37:7890 | ✓ 1664ms | 否 | 否 | ✓ 1718ms | ✓ 1077ms | http |
| 168.235.110.63:3128 | ✓ 198ms | 否 | ✓ 1023ms | ✓ 967ms | ✓ 721ms | http |
| 117.159.239.44:22222 | ✓ 1032ms | ✓ 1292ms | ✓ 1055ms | ✓ 1335ms | ✓ 1119ms | http |
| 160.250.4.13:1 | ✓ 1878ms | 否 | ✓ 1557ms | ✓ 1532ms | ✓ 1283ms | http |
| 222.184.48.235:22222 | 否 | ✓ 1502ms | ✓ 1534ms | 否 | ✓ 1165ms | http |
| 201.144.20.238:3128 | ✓ 762ms | 否 | ✓ 937ms | ✓ 1532ms | ✓ 911ms | http |
| 45.136.131.59:8450 | 否 | 否 | ✓ 574ms | ✓ 947ms | ✓ 748ms | http |
| 133.242.138.34:8100 | 否 | 否 | ✓ 879ms | ✓ 1287ms | ✓ 1056ms | http |
| 192.71.213.85:9091 | ✓ 677ms | 否 | ✓ 578ms | ✓ 1386ms | 否 | http |
| 192.71.213.85:5678 | ✓ 1422ms | 否 | ✓ 1753ms | ✓ 1697ms | 否 | http |
| 117.159.239.50:22222 | ✓ 1039ms | ✓ 1357ms | ✓ 1074ms | ✓ 1368ms | ✓ 1044ms | http |
| 117.159.239.87:22222 | ✓ 1062ms | 否 | ✓ 1301ms | ✓ 1342ms | ✓ 1041ms | http |
| 38.34.179.161:8450 | ✓ 877ms | ✓ 1763ms | ✓ 905ms | ✓ 959ms | ✓ 716ms | http |
| 38.145.218.163:8450 | ✓ 909ms | ✓ 1543ms | ✓ 577ms | ✓ 1006ms | ✓ 984ms | http |
| 38.34.183.130:8450 | ✓ 894ms | 否 | ✓ 362ms | ✓ 1018ms | ✓ 819ms | http |
| 38.34.183.234:8450 | ✓ 894ms | ✓ 993ms | ✓ 1224ms | 否 | ✓ 897ms | http |
| 117.159.239.46:22222 | ✓ 1001ms | ✓ 1292ms | ✓ 1113ms | ✓ 1340ms | ✓ 1008ms | http |
| 117.159.239.51:22222 | ✓ 991ms | ✓ 1300ms | ✓ 1009ms | ✓ 1393ms | ✓ 1114ms | http |
| 159.223.42.219:3128 | ✓ 1312ms | 否 | ✓ 1495ms | ✓ 1243ms | ✓ 987ms | http |
| 38.145.220.15:8450 | ✓ 1749ms | 否 | ✓ 1298ms | ✓ 1917ms | 否 | http |
| 38.145.203.35:8450 | 否 | ✓ 1610ms | 否 | ✓ 1330ms | ✓ 1617ms | http |
| 120.238.159.189:22222 | 否 | ✓ 1603ms | ✓ 1556ms | 否 | ✓ 1148ms | http |
| 38.145.208.182:8448 | ✓ 940ms | ✓ 1860ms | ✓ 1260ms | ✓ 1950ms | ✓ 1671ms | http |
| 38.145.208.186:8448 | ✓ 940ms | 否 | ✓ 1310ms | 否 | ✓ 1895ms | http |
| 106.117.208.101:7890 | ✓ 1635ms | ✓ 1495ms | ✓ 1151ms | ✓ 1623ms | ✓ 1247ms | http |
| 120.232.242.119:22222 | ✓ 1111ms | ✓ 1485ms | ✓ 1193ms | ✓ 1360ms | ✓ 1056ms | http |

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
