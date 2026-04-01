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

最后更新：2026-04-01 08:23:29 UTC（2026-04-01 16:23:29 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 473ms | ✓ 1644ms | ✓ 1356ms | ✓ 1399ms | ✓ 1383ms | http |
| 95.213.217.168:52004 | ✓ 803ms | 否 | ✓ 1621ms | 否 | ✓ 1906ms | http |
| 1.231.81.166:3128 | ✓ 1466ms | ✓ 1551ms | ✓ 1744ms | ✓ 1495ms | ✓ 1295ms | http |
| 147.161.210.140:8800 | ✓ 1431ms | 否 | ✓ 1730ms | ✓ 1152ms | 否 | http |
| 167.103.115.102:8800 | ✓ 1192ms | 否 | ✓ 1120ms | 否 | ✓ 1478ms | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1396ms | ✓ 1619ms | ✓ 1103ms | http |
| 133.242.138.34:8100 | ✓ 1452ms | ✓ 1835ms | 否 | ✓ 1720ms | ✓ 1879ms | http |
| 208.87.243.199:7878 | 否 | ✓ 1546ms | ✓ 1542ms | ✓ 1606ms | 否 | http |
| 35.225.22.61:80 | ✓ 900ms | ✓ 1512ms | 否 | ✓ 1343ms | ✓ 902ms | http |
| 167.103.34.108:8800 | ✓ 1242ms | 否 | ✓ 1376ms | ✓ 1480ms | ✓ 1298ms | http |
| 113.160.132.26:8080 | ✓ 1800ms | ✓ 1845ms | ✓ 1449ms | ✓ 1606ms | ✓ 1212ms | http |
| 167.103.31.122:8800 | ✓ 1741ms | 否 | ✓ 1278ms | 否 | ✓ 1702ms | http |
| 45.167.124.52:8080 | ✓ 1189ms | ✓ 1493ms | ✓ 656ms | ✓ 1632ms | ✓ 1301ms | http |
| 178.128.243.121:3128 | ✓ 748ms | ✓ 1328ms | ✓ 847ms | ✓ 1192ms | ✓ 1238ms | http |
| 167.103.144.127:8800 | ✓ 1603ms | 否 | ✓ 1231ms | ✓ 1500ms | ✓ 1439ms | http |
| 31.192.106.135:8010 | ✓ 1026ms | 否 | ✓ 1291ms | ✓ 1995ms | 否 | http |
| 72.11.150.178:6005 | ✓ 1819ms | ✓ 1738ms | ✓ 931ms | ✓ 1196ms | ✓ 922ms | http |
| 128.199.254.13:9090 | ✓ 1649ms | 否 | ✓ 1142ms | ✓ 1300ms | ✓ 1064ms | http |
| 128.199.113.85:9090 | ✓ 1650ms | 否 | ✓ 1471ms | ✓ 1258ms | ✓ 1088ms | http |
| 160.250.5.22:1 | ✓ 1748ms | 否 | ✓ 1485ms | ✓ 1658ms | ✓ 1550ms | http |
| 190.12.150.244:999 | ✓ 902ms | 否 | ✓ 884ms | ✓ 1654ms | ✓ 1405ms | http |
| 177.234.217.88:999 | ✓ 1298ms | 否 | ✓ 1851ms | ✓ 1786ms | ✓ 1437ms | http |
| 103.163.80.172:8085 | 否 | 否 | ✓ 1933ms | ✓ 1726ms | ✓ 1808ms | http |
| 213.131.85.30:1976 | ✓ 1423ms | 否 | ✓ 1528ms | 否 | ✓ 1538ms | http |
| 82.114.228.67:1080 | ✓ 1116ms | ✓ 1903ms | ✓ 1680ms | 否 | 否 | http |
| 128.199.114.189:9090 | ✓ 1237ms | 否 | ✓ 1025ms | ✓ 1616ms | ✓ 1394ms | http |
| 181.78.44.63:999 | ✓ 990ms | 否 | ✓ 421ms | ✓ 1458ms | ✓ 1106ms | http |
| 103.145.34.212:1111 | ✓ 1574ms | 否 | ✓ 1711ms | 否 | ✓ 1654ms | http |
| 209.126.84.232:8888 | ✓ 822ms | 否 | ✓ 833ms | ✓ 1616ms | ✓ 1157ms | http |
| 203.80.138.81:50000 | 否 | ✓ 1315ms | ✓ 1153ms | 否 | ✓ 992ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1903ms | ✓ 1035ms | ✓ 1691ms | ✓ 1821ms | http |
| 45.12.151.226:2829 | ✓ 1123ms | ✓ 1531ms | 否 | ✓ 1492ms | ✓ 1401ms | http |
| 57.128.188.167:9157 | ✓ 1607ms | 否 | ✓ 1488ms | ✓ 1991ms | ✓ 1786ms | http |
| 120.92.212.16:8890 | ✓ 1201ms | 否 | ✓ 1121ms | 否 | ✓ 1421ms | http |
| 103.173.139.220:8080 | ✓ 1718ms | 否 | ✓ 1699ms | 否 | ✓ 1740ms | http |
| 116.171.106.78:3443 | ✓ 1879ms | ✓ 1802ms | ✓ 1700ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1089ms | ✓ 1713ms | ✓ 1121ms | ✓ 1689ms | 否 | http |
| 150.249.255.91:3128 | ✓ 1225ms | ✓ 1090ms | ✓ 700ms | 否 | 否 | http |
| 37.59.112.197:80 | ✓ 958ms | ✓ 1635ms | ✓ 1272ms | 否 | 否 | http |
| 217.217.249.160:8080 | ✓ 1537ms | 否 | ✓ 1679ms | 否 | ✓ 1240ms | http |
| 181.57.142.26:999 | ✓ 1041ms | ✓ 1859ms | ✓ 1669ms | ✓ 1913ms | ✓ 1611ms | http |
| 202.58.66.10:8181 | ✓ 1707ms | 否 | 否 | ✓ 1765ms | ✓ 1718ms | http |
| 57.128.188.167:9174 | ✓ 1692ms | 否 | ✓ 1745ms | ✓ 1879ms | ✓ 1461ms | http |
| 146.190.80.158:9090 | ✓ 904ms | 否 | 否 | ✓ 1273ms | ✓ 1043ms | http |
| 203.150.113.51:8080 | ✓ 1705ms | 否 | 否 | ✓ 1550ms | ✓ 1805ms | http |
| 34.101.184.164:3128 | ✓ 944ms | 否 | ✓ 1445ms | ✓ 1505ms | ✓ 1714ms | http |
| 128.199.116.219:9090 | ✓ 927ms | 否 | 否 | ✓ 1502ms | ✓ 1135ms | http |
| 128.199.121.61:9090 | ✓ 1403ms | 否 | ✓ 955ms | ✓ 1357ms | ✓ 1010ms | http |
| 59.46.216.131:30001 | ✓ 1192ms | 否 | ✓ 1926ms | 否 | ✓ 1204ms | http |
| 212.58.132.5:8888 | ✓ 1098ms | 否 | ✓ 1790ms | ✓ 1504ms | ✓ 1249ms | http |
| 147.45.186.28:3128 | ✓ 1657ms | 否 | ✓ 1905ms | ✓ 1830ms | 否 | http |
| 57.128.188.167:8188 | ✓ 1744ms | 否 | ✓ 1695ms | 否 | ✓ 1766ms | http |
| 168.222.254.26:8888 | 否 | ✓ 1686ms | ✓ 695ms | 否 | ✓ 1329ms | http |
| 217.217.249.160:80 | ✓ 1105ms | 否 | ✓ 928ms | ✓ 1945ms | ✓ 1444ms | http |
| 223.16.170.103:3128 | ✓ 1358ms | 否 | ✓ 1255ms | ✓ 1312ms | ✓ 1716ms | http |
| 5.102.109.41:999 | ✓ 678ms | 否 | ✓ 369ms | 否 | ✓ 1077ms | http |
| 45.129.141.143:3128 | ✓ 667ms | ✓ 1627ms | ✓ 1193ms | ✓ 1752ms | ✓ 1537ms | http |
| 113.176.92.71:3128 | ✓ 1197ms | 否 | ✓ 1511ms | ✓ 1496ms | ✓ 1137ms | http |
| 38.180.2.107:3128 | ✓ 991ms | ✓ 1916ms | ✓ 1238ms | ✓ 1567ms | ✓ 1337ms | http |
| 103.39.51.190:8080 | ✓ 1931ms | 否 | 否 | ✓ 1834ms | ✓ 1942ms | http |
| 167.71.196.28:8080 | ✓ 1153ms | 否 | ✓ 962ms | 否 | ✓ 1010ms | http |
| 115.231.181.40:8128 | ✓ 1084ms | ✓ 1435ms | ✓ 1092ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 632ms | ✓ 1446ms | ✓ 649ms | 否 | ✓ 1546ms | http |
| 61.52.131.172:8443 | ✓ 1151ms | ✓ 1415ms | ✓ 1086ms | 否 | 否 | http |
| 85.235.150.219:3128 | ✓ 457ms | ✓ 1906ms | ✓ 1569ms | 否 | 否 | http |
| 92.119.127.213:6005 | ✓ 529ms | ✓ 1427ms | ✓ 1661ms | ✓ 1857ms | 否 | http |
| 8.222.175.80:6128 | ✓ 912ms | 否 | ✓ 976ms | ✓ 1244ms | ✓ 994ms | http |
| 106.117.208.101:7890 | ✓ 1894ms | ✓ 1619ms | ✓ 1188ms | ✓ 1623ms | ✓ 1187ms | http |

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
