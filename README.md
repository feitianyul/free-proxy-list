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

最后更新：2026-06-13 14:56:23 UTC（2026-06-13 22:56:23 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 83.147.36.155:8080 | ✓ 662ms | ✓ 1462ms | ✓ 1956ms | 否 | ✓ 1699ms | http |
| 91.107.172.30:82 | ✓ 933ms | 否 | 否 | ✓ 1688ms | ✓ 1398ms | http |
| 185.200.188.234:10001 | ✓ 1230ms | ✓ 1726ms | ✓ 1046ms | 否 | ✓ 1538ms | http |
| 138.124.114.42:7443 | ✓ 1695ms | 否 | ✓ 1106ms | 否 | ✓ 1170ms | http |
| 170.106.136.181:31002 | ✓ 580ms | ✓ 1010ms | ✓ 667ms | ✓ 881ms | ✓ 660ms | http |
| 157.230.220.25:4857 | ✓ 881ms | ✓ 1691ms | ✓ 758ms | ✓ 1789ms | 否 | http |
| 34.43.46.91:443 | ✓ 459ms | 否 | ✓ 741ms | ✓ 1219ms | 否 | http |
| 34.43.46.91:80 | ✓ 297ms | ✓ 1533ms | ✓ 1401ms | ✓ 1549ms | 否 | http |
| 104.154.186.48:80 | ✓ 466ms | 否 | ✓ 1832ms | ✓ 1192ms | ✓ 1086ms | http |
| 81.200.154.236:48503 | ✓ 656ms | ✓ 1858ms | ✓ 981ms | ✓ 1388ms | 否 | http |
| 176.111.37.5:39811 | ✓ 975ms | ✓ 1763ms | ✓ 924ms | 否 | 否 | http |
| 159.198.35.187:1080 | 否 | 否 | ✓ 367ms | ✓ 1220ms | ✓ 942ms | http |
| 157.245.100.190:442 | ✓ 1140ms | 否 | ✓ 1500ms | ✓ 1436ms | ✓ 1568ms | http |
| 113.160.132.26:8080 | ✓ 1616ms | ✓ 1907ms | ✓ 1102ms | 否 | ✓ 1318ms | http |
| 185.11.134.227:8443 | ✓ 645ms | ✓ 1732ms | 否 | 否 | ✓ 1196ms | http |
| 202.28.194.139:31280 | ✓ 1857ms | 否 | ✓ 1787ms | 否 | ✓ 1948ms | http |
| 95.3.69.222:8080 | ✓ 1878ms | ✓ 1810ms | ✓ 1676ms | 否 | 否 | http |
| 89.169.53.40:7443 | ✓ 1089ms | 否 | ✓ 1277ms | 否 | ✓ 1891ms | http |
| 91.107.182.124:82 | ✓ 1395ms | 否 | ✓ 1032ms | 否 | ✓ 1932ms | http |
| 89.125.68.33:10000 | ✓ 1077ms | 否 | ✓ 986ms | ✓ 1641ms | 否 | http |
| 94.228.163.232:1080 | ✓ 1169ms | ✓ 1678ms | ✓ 1950ms | 否 | 否 | http |
| 47.79.119.13:8080 | ✓ 1090ms | 否 | ✓ 920ms | ✓ 1319ms | ✓ 1000ms | http |
| 82.97.247.37:80 | ✓ 1124ms | 否 | ✓ 494ms | 否 | ✓ 972ms | http |
| 3.137.86.220:443 | ✓ 598ms | ✓ 1945ms | 否 | 否 | ✓ 1372ms | http |
| 150.241.116.167:443 | ✓ 484ms | 否 | ✓ 1013ms | ✓ 1648ms | ✓ 1411ms | http |
| 176.111.37.216:39811 | ✓ 462ms | 否 | ✓ 1297ms | ✓ 1718ms | ✓ 1387ms | http |
| 77.246.104.106:4433 | ✓ 807ms | ✓ 1897ms | ✓ 1257ms | 否 | 否 | http |
| 138.124.113.102:7443 | 否 | 否 | ✓ 1228ms | ✓ 1505ms | ✓ 885ms | http |
| 152.32.132.190:7890 | ✓ 996ms | 否 | 否 | ✓ 1057ms | ✓ 1605ms | http |
| 46.101.251.110:8080 | ✓ 1086ms | ✓ 1710ms | ✓ 1584ms | ✓ 1922ms | 否 | http |
| 52.188.28.218:3128 | ✓ 333ms | ✓ 1691ms | ✓ 279ms | ✓ 1913ms | ✓ 632ms | http |
| 92.118.112.32:1082 | ✓ 1265ms | ✓ 1038ms | ✓ 1489ms | 否 | 否 | http |
| 47.80.103.120:443 | ✓ 1328ms | 否 | ✓ 1521ms | 否 | ✓ 1445ms | http |
| 180.2.108.38:8080 | ✓ 777ms | 否 | ✓ 1490ms | 否 | ✓ 963ms | http |
| 85.192.60.187:7443 | ✓ 941ms | 否 | ✓ 762ms | ✓ 1779ms | ✓ 1480ms | http |
| 3.137.86.220:1080 | ✓ 722ms | ✓ 895ms | 否 | ✓ 1126ms | 否 | http |
| 79.137.205.130:7443 | ✓ 691ms | 否 | ✓ 1663ms | ✓ 1401ms | 否 | http |
| 91.107.168.255:82 | ✓ 970ms | 否 | ✓ 430ms | ✓ 1349ms | ✓ 1358ms | http |
| 85.192.28.47:7443 | ✓ 1783ms | 否 | ✓ 1712ms | ✓ 1244ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1977ms | 否 | ✓ 1035ms | ✓ 1396ms | ✓ 1984ms | http |
| 31.57.172.220:10808 | ✓ 1114ms | 否 | ✓ 1593ms | ✓ 1382ms | ✓ 778ms | http |
| 77.110.116.93:7443 | ✓ 663ms | 否 | 否 | ✓ 1737ms | ✓ 1162ms | http |
| 213.165.42.185:7443 | ✓ 1202ms | 否 | ✓ 727ms | ✓ 1889ms | 否 | http |
| 144.172.114.214:1080 | ✓ 481ms | 否 | ✓ 682ms | ✓ 1000ms | 否 | http |
| 217.154.155.115:8080 | ✓ 1661ms | ✓ 1783ms | ✓ 1151ms | ✓ 1263ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1719ms | ✓ 1982ms | 否 | ✓ 1760ms | 否 | http |
| 47.80.103.120:8080 | ✓ 1074ms | 否 | 否 | ✓ 1443ms | ✓ 1905ms | http |
| 138.197.68.35:4857 | ✓ 1278ms | 否 | ✓ 180ms | ✓ 1395ms | ✓ 872ms | http |
| 85.192.28.62:7443 | ✓ 998ms | 否 | ✓ 1270ms | 否 | ✓ 1584ms | http |
| 212.58.132.5:8888 | ✓ 1404ms | 否 | ✓ 1263ms | ✓ 1467ms | ✓ 1205ms | http |
| 103.157.117.226:81 | ✓ 1979ms | 否 | ✓ 1435ms | ✓ 1762ms | ✓ 1568ms | http |
| 185.141.26.131:3128 | ✓ 1539ms | 否 | ✓ 1322ms | 否 | ✓ 1632ms | http |
| 185.191.239.248:3128 | ✓ 1691ms | 否 | ✓ 445ms | 否 | ✓ 1041ms | http |
| 18.180.59.181:80 | ✓ 1492ms | 否 | ✓ 1900ms | ✓ 1324ms | ✓ 1081ms | http |
| 103.157.200.126:3128 | ✓ 1094ms | 否 | ✓ 1086ms | ✓ 1858ms | ✓ 1439ms | http |
| 91.149.222.102:22335 | ✓ 1391ms | ✓ 1873ms | 否 | 否 | ✓ 1497ms | http |
| 154.223.77.54:10001 | 否 | 否 | ✓ 1577ms | ✓ 1517ms | ✓ 1117ms | http |
| 174.137.134.182:2999 | ✓ 1136ms | 否 | 否 | ✓ 1668ms | ✓ 1846ms | http |
| 45.84.222.25:1080 | ✓ 1050ms | 否 | ✓ 1889ms | 否 | ✓ 1602ms | http |
| 209.141.35.94:28080 | ✓ 1176ms | 否 | ✓ 1419ms | ✓ 1584ms | ✓ 1440ms | http |
| 209.141.35.94:28017 | ✓ 1318ms | 否 | ✓ 1571ms | ✓ 1700ms | ✓ 1611ms | http |
| 159.223.87.50:443 | ✓ 1483ms | 否 | ✓ 1455ms | ✓ 1700ms | ✓ 1627ms | http |
| 169.212.15.161:5000 | ✓ 955ms | 否 | ✓ 764ms | ✓ 1724ms | ✓ 1941ms | http |
| 85.234.100.149:8080 | ✓ 1435ms | 否 | ✓ 563ms | 否 | ✓ 1487ms | http |
| 82.102.11.164:3460 | ✓ 1537ms | 否 | ✓ 1919ms | ✓ 1576ms | ✓ 1499ms | http |
| 85.192.28.65:7443 | 否 | ✓ 1901ms | ✓ 780ms | 否 | ✓ 1636ms | http |
| 91.186.213.124:1081 | ✓ 1047ms | 否 | ✓ 1181ms | 否 | ✓ 1355ms | http |
| 43.167.192.85:8080 | 否 | 否 | ✓ 1296ms | ✓ 1013ms | ✓ 1847ms | http |
| 34.96.238.40:8080 | ✓ 1044ms | 否 | ✓ 1694ms | 否 | ✓ 1593ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1391ms | ✓ 1566ms | ✓ 1536ms | http |
| 121.230.8.136:1080 | 否 | ✓ 1729ms | ✓ 1478ms | ✓ 1731ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1060ms | 否 | ✓ 1197ms | 否 | ✓ 1559ms | http |
| 151.243.180.211:2080 | ✓ 425ms | 否 | ✓ 1741ms | 否 | ✓ 1593ms | http |
| 168.110.52.228:3128 | ✓ 788ms | 否 | ✓ 1034ms | 否 | ✓ 1850ms | http |
| 144.31.134.103:1080 | 否 | 否 | ✓ 1272ms | ✓ 1581ms | ✓ 1494ms | http |
| 49.12.219.42:8000 | ✓ 1409ms | 否 | 否 | ✓ 1246ms | ✓ 1613ms | http |
| 50.114.102.16:8888 | ✓ 742ms | ✓ 1770ms | ✓ 826ms | ✓ 1698ms | ✓ 1239ms | http |
| 34.87.80.221:30000 | ✓ 1451ms | 否 | ✓ 903ms | ✓ 1240ms | ✓ 1040ms | http |
| 115.199.235.199:10808 | ✓ 1309ms | ✓ 1378ms | ✓ 1130ms | 否 | ✓ 1184ms | http |
| 43.156.228.168:80 | ✓ 1926ms | 否 | 否 | ✓ 1746ms | ✓ 1599ms | http |

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
