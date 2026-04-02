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

最后更新：2026-04-02 14:19:05 UTC（2026-04-02 22:19:05 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | 否 | ✓ 1361ms | ✓ 845ms | ✓ 986ms | ✓ 790ms | http |
| 159.223.71.162:443 | ✓ 1498ms | 否 | ✓ 1370ms | ✓ 1544ms | 否 | http |
| 167.103.115.102:8800 | ✓ 1498ms | 否 | ✓ 1102ms | ✓ 1222ms | ✓ 1355ms | http |
| 203.80.138.81:50000 | ✓ 1308ms | ✓ 1977ms | ✓ 1678ms | ✓ 1341ms | ✓ 1641ms | http |
| 113.160.132.26:8080 | ✓ 1558ms | 否 | ✓ 1318ms | ✓ 1640ms | ✓ 1168ms | http |
| 45.149.92.147:5001 | ✓ 772ms | 否 | ✓ 753ms | ✓ 956ms | ✓ 755ms | http |
| 147.161.210.140:8800 | 否 | 否 | ✓ 777ms | ✓ 1065ms | ✓ 877ms | http |
| 167.103.34.108:8800 | ✓ 1364ms | 否 | ✓ 1320ms | ✓ 1466ms | ✓ 1423ms | http |
| 167.103.31.122:8800 | ✓ 1579ms | 否 | ✓ 1449ms | ✓ 1763ms | ✓ 1905ms | http |
| 45.167.124.52:8080 | ✓ 551ms | 否 | ✓ 738ms | ✓ 1555ms | ✓ 1299ms | http |
| 45.167.125.21:999 | ✓ 945ms | ✓ 1691ms | ✓ 1263ms | ✓ 1745ms | ✓ 1518ms | http |
| 159.223.71.162:8080 | ✓ 1336ms | 否 | ✓ 910ms | ✓ 1220ms | ✓ 979ms | http |
| 95.213.217.168:52004 | ✓ 648ms | 否 | ✓ 746ms | ✓ 1662ms | ✓ 1269ms | http |
| 208.87.243.199:7878 | ✓ 731ms | ✓ 1383ms | ✓ 1735ms | 否 | 否 | http |
| 167.103.144.127:8800 | 否 | 否 | ✓ 1274ms | ✓ 1445ms | ✓ 1304ms | http |
| 45.140.147.82:1082 | ✓ 1407ms | ✓ 1451ms | ✓ 928ms | 否 | ✓ 1525ms | http |
| 146.190.80.158:9090 | 否 | 否 | ✓ 1137ms | ✓ 1241ms | ✓ 955ms | http |
| 1.231.81.166:3128 | ✓ 1523ms | ✓ 1638ms | ✓ 1374ms | ✓ 1195ms | ✓ 982ms | http |
| 128.199.116.219:9090 | ✓ 1498ms | 否 | ✓ 863ms | ✓ 1229ms | ✓ 1013ms | http |
| 128.199.113.85:9090 | ✓ 1483ms | 否 | ✓ 899ms | ✓ 1212ms | ✓ 973ms | http |
| 147.161.239.240:8800 | 否 | ✓ 1346ms | ✓ 978ms | 否 | ✓ 1701ms | http |
| 38.145.208.172:8448 | ✓ 758ms | 否 | ✓ 1993ms | ✓ 1208ms | 否 | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 945ms | ✓ 1297ms | ✓ 1034ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1784ms | ✓ 1640ms | ✓ 1257ms | http |
| 106.10.55.212:1121 | ✓ 1516ms | 否 | ✓ 1988ms | ✓ 1546ms | ✓ 1025ms | http |
| 45.12.151.226:2829 | ✓ 809ms | ✓ 1743ms | ✓ 823ms | 否 | 否 | http |
| 20.78.118.91:8561 | ✓ 1785ms | ✓ 1405ms | ✓ 740ms | ✓ 1014ms | ✓ 772ms | http |
| 38.207.164.82:6005 | ✓ 1503ms | 否 | ✓ 1392ms | ✓ 1393ms | ✓ 1018ms | http |
| 165.232.146.249:3128 | ✓ 999ms | ✓ 1385ms | ✓ 636ms | 否 | ✓ 720ms | http |
| 45.129.141.143:3128 | ✓ 674ms | 否 | ✓ 1840ms | 否 | ✓ 1666ms | http |
| 209.126.84.232:8888 | ✓ 1186ms | ✓ 1903ms | ✓ 721ms | ✓ 1493ms | ✓ 1182ms | http |
| 42.96.16.158:1311 | ✓ 1436ms | 否 | ✓ 1104ms | ✓ 1307ms | ✓ 1146ms | http |
| 177.234.217.88:999 | ✓ 1283ms | 否 | ✓ 1830ms | ✓ 1813ms | ✓ 1571ms | http |
| 128.199.121.61:9090 | ✓ 1272ms | 否 | ✓ 1002ms | ✓ 1292ms | ✓ 1156ms | http |
| 45.140.147.82:1081 | ✓ 1117ms | 否 | ✓ 1419ms | ✓ 1141ms | 否 | http |
| 121.230.8.17:1080 | ✓ 1337ms | ✓ 1621ms | ✓ 1633ms | 否 | ✓ 1333ms | http |
| 94.159.103.156:3128 | ✓ 774ms | 否 | ✓ 1600ms | ✓ 1893ms | ✓ 1416ms | http |
| 35.225.22.61:80 | ✓ 706ms | 否 | ✓ 1840ms | ✓ 944ms | ✓ 684ms | http |
| 20.27.13.35:8561 | ✓ 1617ms | ✓ 1558ms | ✓ 1018ms | ✓ 1360ms | ✓ 990ms | http |
| 157.230.38.173:3128 | ✓ 838ms | 否 | ✓ 983ms | ✓ 1238ms | ✓ 944ms | http |
| 38.34.179.173:8452 | ✓ 726ms | ✓ 1473ms | ✓ 1409ms | ✓ 1340ms | 否 | http |
| 222.184.48.251:22222 | ✓ 1345ms | ✓ 1242ms | ✓ 1116ms | 否 | 否 | http |
| 212.58.132.5:8888 | ✓ 1577ms | 否 | ✓ 1247ms | ✓ 1474ms | ✓ 1204ms | http |
| 116.171.106.26:3443 | ✓ 1736ms | 否 | ✓ 1636ms | ✓ 1982ms | 否 | http |
| 114.237.77.219:1080 | 否 | ✓ 1284ms | ✓ 1156ms | 否 | ✓ 1416ms | http |
| 192.71.213.85:9812 | ✓ 1309ms | 否 | ✓ 1714ms | ✓ 1886ms | 否 | http |
| 45.136.131.32:8443 | 否 | ✓ 1286ms | ✓ 1800ms | ✓ 1224ms | ✓ 717ms | http |
| 46.101.190.71:3128 | ✓ 501ms | 否 | ✓ 1525ms | ✓ 1790ms | ✓ 1295ms | http |
| 62.113.119.14:8080 | ✓ 598ms | ✓ 1563ms | ✓ 1069ms | ✓ 1512ms | ✓ 1100ms | http |
| 86.53.183.16:1080 | ✓ 615ms | ✓ 1870ms | ✓ 1427ms | 否 | 否 | http |
| 182.53.202.208:8080 | 否 | 否 | ✓ 1370ms | ✓ 1725ms | ✓ 1568ms | http |
| 16.78.119.130:443 | 否 | ✓ 1892ms | ✓ 1684ms | ✓ 1843ms | ✓ 1746ms | http |
| 45.140.147.155:1082 | ✓ 1462ms | ✓ 1682ms | ✓ 884ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 1435ms | ✓ 1407ms | ✓ 1299ms | ✓ 1927ms | 否 | http |
| 128.199.254.13:9090 | ✓ 1455ms | 否 | ✓ 1299ms | ✓ 1198ms | ✓ 1173ms | http |
| 181.78.44.63:999 | 否 | 否 | ✓ 1456ms | ✓ 1363ms | ✓ 1372ms | http |
| 104.248.151.93:9090 | ✓ 965ms | 否 | ✓ 921ms | ✓ 1306ms | ✓ 994ms | http |
| 222.184.48.242:22222 | 否 | ✓ 1230ms | ✓ 1015ms | 否 | ✓ 1930ms | http |
| 213.131.85.30:1976 | 否 | ✓ 1994ms | ✓ 1360ms | 否 | ✓ 1828ms | http |
| 82.114.228.67:1080 | ✓ 1648ms | 否 | ✓ 788ms | ✓ 1564ms | 否 | http |
| 5.102.109.41:999 | 否 | ✓ 1752ms | 否 | ✓ 1518ms | ✓ 1183ms | http |
| 47.74.226.8:5001 | ✓ 1538ms | ✓ 1712ms | ✓ 1102ms | ✓ 1478ms | 否 | http |
| 61.52.131.172:8443 | ✓ 996ms | ✓ 1351ms | ✓ 1004ms | ✓ 1349ms | ✓ 1018ms | http |
| 45.136.198.40:3128 | ✓ 1240ms | 否 | ✓ 1964ms | ✓ 1872ms | ✓ 1617ms | http |
| 120.92.212.16:7890 | ✓ 1118ms | ✓ 1356ms | ✓ 1155ms | ✓ 1359ms | ✓ 1123ms | http |
| 120.92.212.16:8890 | ✓ 1094ms | ✓ 1390ms | ✓ 1144ms | ✓ 1367ms | ✓ 1412ms | http |
| 217.217.249.160:8080 | ✓ 1423ms | 否 | ✓ 1640ms | 否 | ✓ 1922ms | http |
| 185.114.73.2:1080 | ✓ 1632ms | 否 | ✓ 1648ms | ✓ 1904ms | ✓ 1762ms | http |
| 101.43.127.100:8877 | ✓ 1998ms | ✓ 1324ms | ✓ 1050ms | ✓ 1389ms | ✓ 1033ms | http |

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
