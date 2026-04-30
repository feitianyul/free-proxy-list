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

最后更新：2026-04-30 00:53:01 UTC（2026-04-30 08:53:01 UTC+8）

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
| 216.180.127.45:1080 | ✓ 1658ms | ✓ 1387ms | ✓ 476ms | ✓ 1344ms | ✓ 990ms | http |
| 34.96.238.40:8080 | ✓ 949ms | ✓ 1175ms | ✓ 956ms | ✓ 1042ms | 否 | http |
| 47.85.51.197:1080 | ✓ 758ms | 否 | ✓ 754ms | ✓ 1179ms | 否 | http |
| 34.71.229.255:3128 | ✓ 931ms | ✓ 1792ms | ✓ 1568ms | ✓ 1428ms | ✓ 1387ms | http |
| 113.160.132.26:8080 | ✓ 1336ms | ✓ 1327ms | ✓ 1415ms | ✓ 1269ms | ✓ 1037ms | http |
| 1.231.81.166:3128 | ✓ 1881ms | ✓ 1125ms | ✓ 1776ms | ✓ 1090ms | ✓ 1126ms | http |
| 154.64.232.35:8080 | ✓ 1201ms | ✓ 1603ms | 否 | ✓ 1693ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1879ms | 否 | ✓ 1507ms | ✓ 1561ms | ✓ 1269ms | http |
| 46.101.95.183:8888 | ✓ 1944ms | 否 | ✓ 1572ms | ✓ 1981ms | 否 | http |
| 45.167.124.71:999 | ✓ 1416ms | 否 | 否 | ✓ 1915ms | ✓ 1790ms | http |
| 103.3.246.71:3128 | ✓ 1449ms | 否 | ✓ 1071ms | ✓ 1125ms | ✓ 958ms | http |
| 218.108.131.186:17890 | ✓ 944ms | 否 | ✓ 1384ms | ✓ 1557ms | ✓ 965ms | http |
| 91.233.223.147:3128 | ✓ 1274ms | 否 | ✓ 1373ms | 否 | ✓ 1913ms | http |
| 77.110.119.136:3128 | ✓ 1711ms | 否 | ✓ 883ms | 否 | ✓ 1147ms | http |
| 64.188.67.154:10808 | 否 | 否 | ✓ 782ms | ✓ 1924ms | ✓ 1171ms | http |
| 159.223.225.118:8888 | ✓ 1164ms | ✓ 1959ms | 否 | 否 | ✓ 1477ms | http |
| 86.104.74.110:1081 | ✓ 1314ms | 否 | ✓ 870ms | ✓ 1869ms | 否 | http |
| 130.61.174.200:1080 | ✓ 1163ms | ✓ 1643ms | 否 | 否 | ✓ 1483ms | http |
| 194.150.220.163:1082 | ✓ 1165ms | 否 | ✓ 718ms | ✓ 1891ms | 否 | http |
| 172.236.145.31:7890 | ✓ 679ms | 否 | ✓ 665ms | ✓ 1129ms | ✓ 756ms | http |
| 59.46.216.131:30001 | ✓ 1023ms | ✓ 1241ms | ✓ 1037ms | 否 | 否 | http |
| 168.110.52.228:3128 | ✓ 1974ms | ✓ 1207ms | ✓ 457ms | 否 | ✓ 1793ms | http |
| 210.223.44.230:3128 | ✓ 543ms | ✓ 891ms | 否 | ✓ 1619ms | ✓ 769ms | http |
| 132.226.235.199:1080 | ✓ 1708ms | 否 | ✓ 1371ms | ✓ 1124ms | ✓ 1373ms | http |
| 54.189.56.88:18083 | ✓ 1136ms | 否 | 否 | ✓ 1887ms | ✓ 1700ms | http |
| 77.110.116.224:3128 | ✓ 1291ms | 否 | ✓ 988ms | 否 | ✓ 1464ms | http |
| 137.59.47.73:3128 | ✓ 923ms | 否 | ✓ 957ms | ✓ 1135ms | ✓ 820ms | http |
| 52.59.51.29:3128 | ✓ 1346ms | 否 | ✓ 1882ms | 否 | ✓ 1434ms | http |
| 120.92.108.86:7890 | ✓ 1237ms | 否 | ✓ 1669ms | 否 | ✓ 1946ms | http |
| 45.129.141.143:3128 | ✓ 1311ms | 否 | ✓ 1920ms | 否 | ✓ 1952ms | http |
| 103.157.200.126:3128 | ✓ 1839ms | 否 | ✓ 1485ms | ✓ 1978ms | ✓ 1594ms | http |
| 152.70.91.193:40000 | ✓ 1527ms | 否 | 否 | ✓ 1406ms | ✓ 1296ms | http |
| 62.113.119.14:8080 | ✓ 1287ms | ✓ 1854ms | ✓ 1050ms | 否 | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1756ms | ✓ 1684ms | ✓ 1910ms | ✓ 1085ms | http |
| 121.230.8.91:1080 | ✓ 842ms | ✓ 1286ms | ✓ 991ms | ✓ 1219ms | ✓ 1012ms | http |
| 52.163.56.148:80 | ✓ 726ms | ✓ 1097ms | ✓ 997ms | ✓ 1123ms | ✓ 1231ms | http |
| 103.144.28.78:3333 | 否 | ✓ 963ms | ✓ 1347ms | ✓ 1347ms | ✓ 997ms | http |
| 38.92.10.98:20058 | ✓ 714ms | ✓ 1574ms | ✓ 642ms | 否 | ✓ 786ms | http |
| 121.230.9.132:1080 | ✓ 983ms | ✓ 1813ms | ✓ 1177ms | ✓ 1544ms | ✓ 1083ms | http |
| 121.230.8.22:1080 | 否 | ✓ 1199ms | ✓ 1074ms | ✓ 1274ms | ✓ 1014ms | http |
| 51.92.173.133:40878 | ✓ 1310ms | 否 | ✓ 1711ms | 否 | ✓ 1803ms | http |
| 13.60.163.108:4145 | ✓ 1276ms | 否 | ✓ 1836ms | 否 | ✓ 1822ms | http |
| 52.56.167.111:25806 | ✓ 1263ms | 否 | ✓ 1906ms | 否 | ✓ 1646ms | http |
| 80.92.204.47:1081 | 否 | 否 | ✓ 692ms | ✓ 1928ms | ✓ 995ms | http |
| 86.104.72.220:1081 | ✓ 344ms | ✓ 1601ms | ✓ 1426ms | ✓ 1981ms | 否 | http |
| 64.188.67.154:1080 | ✓ 1682ms | 否 | ✓ 1640ms | ✓ 1976ms | ✓ 1423ms | http |
| 8.219.97.248:80 | ✓ 1062ms | 否 | ✓ 985ms | 否 | ✓ 1655ms | http |
| 54.189.56.88:31861 | ✓ 1386ms | 否 | ✓ 886ms | 否 | ✓ 1736ms | http |
| 183.232.248.73:7890 | ✓ 1309ms | ✓ 1431ms | ✓ 1397ms | 否 | 否 | http |
| 38.92.10.152:57579 | ✓ 492ms | ✓ 793ms | ✓ 1118ms | 否 | ✓ 1938ms | http |
| 116.254.118.180:80 | ✓ 1691ms | ✓ 1989ms | ✓ 896ms | ✓ 1241ms | ✓ 1990ms | http |
| 113.176.92.71:3128 | ✓ 1153ms | ✓ 1515ms | ✓ 1161ms | ✓ 1148ms | ✓ 911ms | http |
| 4.247.152.147:3128 | ✓ 1658ms | ✓ 1808ms | ✓ 1408ms | 否 | ✓ 1432ms | http |
| 136.244.67.217:1080 | ✓ 1467ms | ✓ 1779ms | ✓ 1493ms | 否 | 否 | http |
| 138.197.68.35:4857 | 否 | ✓ 1806ms | ✓ 312ms | ✓ 1214ms | 否 | http |
| 18.170.25.193:8812 | ✓ 1869ms | 否 | ✓ 1671ms | ✓ 1949ms | ✓ 1855ms | http |
| 86.104.74.110:1082 | ✓ 1852ms | ✓ 1540ms | ✓ 1388ms | ✓ 1882ms | 否 | http |
| 217.77.102.18:3128 | ✓ 1621ms | 否 | ✓ 1696ms | 否 | ✓ 1928ms | http |
| 3.238.34.111:44677 | ✓ 1486ms | 否 | ✓ 1658ms | 否 | ✓ 1582ms | http |
| 40.177.99.164:37358 | ✓ 1481ms | 否 | ✓ 1978ms | 否 | ✓ 1681ms | http |
| 15.160.116.45:31345 | ✓ 1480ms | 否 | ✓ 1682ms | 否 | ✓ 1882ms | http |
| 13.60.181.61:4192 | ✓ 1506ms | 否 | ✓ 1963ms | 否 | ✓ 1636ms | http |
| 8.154.21.175:3128 | ✓ 811ms | ✓ 1005ms | ✓ 815ms | ✓ 1002ms | ✓ 888ms | http |
| 45.140.147.155:1082 | ✓ 1327ms | 否 | ✓ 1471ms | 否 | ✓ 1453ms | http |
| 139.162.153.201:3128 | ✓ 1270ms | ✓ 1688ms | ✓ 724ms | ✓ 1549ms | ✓ 1206ms | http |
| 45.140.147.82:1081 | ✓ 640ms | ✓ 1684ms | ✓ 1306ms | ✓ 1896ms | ✓ 1362ms | http |
| 121.230.9.160:1080 | ✓ 1321ms | ✓ 1645ms | ✓ 1877ms | ✓ 1307ms | ✓ 1113ms | http |
| 94.72.109.214:8888 | ✓ 1258ms | ✓ 1889ms | ✓ 1098ms | ✓ 1648ms | ✓ 1225ms | http |
| 120.92.212.16:7890 | ✓ 980ms | ✓ 1384ms | ✓ 1674ms | ✓ 1069ms | ✓ 931ms | http |
| 120.92.212.16:8890 | ✓ 1821ms | 否 | 否 | ✓ 1680ms | ✓ 960ms | http |
| 103.67.46.225:3125 | ✓ 1887ms | 否 | ✓ 1510ms | ✓ 1661ms | ✓ 1612ms | http |
| 61.52.131.172:8443 | ✓ 830ms | ✓ 1077ms | ✓ 906ms | ✓ 1119ms | ✓ 929ms | http |
| 54.188.236.206:29223 | ✓ 1433ms | 否 | ✓ 1537ms | 否 | ✓ 1910ms | http |
| 13.41.196.179:23714 | ✓ 1262ms | 否 | ✓ 1820ms | 否 | ✓ 1432ms | http |
| 16.62.123.236:9002 | ✓ 1387ms | 否 | ✓ 1824ms | 否 | ✓ 1759ms | http |
| 37.187.109.70:10111 | ✓ 1227ms | ✓ 1614ms | ✓ 1266ms | 否 | 否 | http |
| 183.238.3.150:7897 | ✓ 858ms | ✓ 1060ms | ✓ 998ms | ✓ 1011ms | ✓ 781ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1450ms | ✓ 1958ms | ✓ 1909ms | http |
| 103.39.51.207:8080 | ✓ 1214ms | 否 | ✓ 1699ms | ✓ 1658ms | ✓ 1446ms | http |
| 42.101.8.101:8888 | ✓ 1925ms | ✓ 1914ms | ✓ 1849ms | 否 | 否 | http |

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
