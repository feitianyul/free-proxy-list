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

最后更新：2026-04-02 09:59:57 UTC（2026-04-02 17:59:57 UTC+8）

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
| 43.99.54.236:5555 | ✓ 838ms | ✓ 1312ms | ✓ 767ms | ✓ 1008ms | ✓ 772ms | http |
| 1.231.81.166:3128 | ✓ 902ms | ✓ 1201ms | ✓ 1301ms | ✓ 1268ms | ✓ 967ms | http |
| 147.161.210.140:8800 | ✓ 786ms | 否 | ✓ 902ms | ✓ 1772ms | ✓ 1165ms | http |
| 147.161.239.240:8800 | ✓ 1372ms | ✓ 1671ms | ✓ 869ms | ✓ 1658ms | ✓ 1470ms | http |
| 159.223.71.162:443 | ✓ 1783ms | 否 | ✓ 1240ms | ✓ 1192ms | ✓ 935ms | http |
| 203.80.138.81:50000 | 否 | 否 | ✓ 1442ms | ✓ 1477ms | ✓ 1440ms | http |
| 167.103.115.102:8800 | ✓ 1786ms | 否 | ✓ 1132ms | ✓ 1647ms | ✓ 1425ms | http |
| 113.160.132.26:8080 | ✓ 1779ms | ✓ 1524ms | 否 | ✓ 1384ms | ✓ 1216ms | http |
| 95.213.217.168:52004 | ✓ 1426ms | ✓ 1930ms | 否 | ✓ 1963ms | ✓ 1739ms | http |
| 167.103.34.108:8800 | ✓ 1840ms | 否 | ✓ 1507ms | ✓ 1464ms | ✓ 1804ms | http |
| 120.92.212.16:8890 | ✓ 1037ms | 否 | ✓ 1091ms | 否 | ✓ 1111ms | http |
| 45.167.124.52:8080 | ✓ 548ms | ✓ 1505ms | ✓ 533ms | ✓ 1595ms | ✓ 1340ms | http |
| 35.225.22.61:80 | ✓ 914ms | ✓ 1669ms | ✓ 214ms | ✓ 972ms | 否 | http |
| 45.12.151.226:2829 | ✓ 1175ms | 否 | ✓ 824ms | ✓ 1687ms | ✓ 1210ms | http |
| 165.232.146.249:3128 | 否 | ✓ 1093ms | ✓ 1771ms | 否 | ✓ 1285ms | http |
| 8.219.97.248:80 | ✓ 1987ms | 否 | 否 | ✓ 1951ms | ✓ 1989ms | http |
| 167.103.144.127:8800 | ✓ 1706ms | 否 | ✓ 1185ms | ✓ 1442ms | ✓ 1313ms | http |
| 167.103.31.122:8800 | ✓ 1710ms | 否 | ✓ 1361ms | 否 | ✓ 1670ms | http |
| 128.199.113.85:9090 | 否 | 否 | ✓ 911ms | ✓ 1353ms | ✓ 1042ms | http |
| 166.0.192.117:8888 | ✓ 757ms | ✓ 1228ms | ✓ 1883ms | ✓ 1137ms | ✓ 1192ms | http |
| 146.190.80.158:9090 | ✓ 1282ms | 否 | 否 | ✓ 1220ms | ✓ 1056ms | http |
| 128.199.116.219:9090 | ✓ 1276ms | 否 | ✓ 918ms | ✓ 1192ms | 否 | http |
| 159.223.71.162:8080 | ✓ 1276ms | 否 | ✓ 851ms | ✓ 1234ms | ✓ 997ms | http |
| 128.199.121.61:9090 | ✓ 1278ms | 否 | ✓ 884ms | ✓ 1200ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1098ms | 否 | ✓ 984ms | ✓ 1550ms | ✓ 1060ms | http |
| 45.167.125.21:999 | ✓ 859ms | 否 | ✓ 1394ms | ✓ 1903ms | ✓ 1426ms | http |
| 177.234.217.88:999 | ✓ 1384ms | 否 | ✓ 1839ms | ✓ 1825ms | ✓ 1556ms | http |
| 128.199.114.189:9090 | ✓ 917ms | 否 | 否 | ✓ 1261ms | ✓ 1115ms | http |
| 120.92.212.16:7890 | ✓ 1253ms | ✓ 1845ms | ✓ 1072ms | ✓ 1363ms | ✓ 1078ms | http |
| 208.87.243.199:7878 | ✓ 605ms | 否 | ✓ 963ms | ✓ 864ms | ✓ 654ms | http |
| 209.126.84.232:8888 | 否 | 否 | ✓ 1983ms | ✓ 1558ms | ✓ 1133ms | http |
| 45.136.198.40:3128 | ✓ 839ms | 否 | ✓ 1757ms | 否 | ✓ 1723ms | http |
| 34.101.184.164:3128 | ✓ 1571ms | 否 | ✓ 1293ms | 否 | ✓ 1505ms | http |
| 49.51.244.112:8888 | 否 | ✓ 1762ms | ✓ 727ms | ✓ 986ms | ✓ 646ms | http |
| 150.241.71.15:1080 | ✓ 1865ms | 否 | ✓ 1363ms | ✓ 1808ms | 否 | http |
| 119.28.156.42:3128 | ✓ 751ms | 否 | ✓ 725ms | ✓ 1010ms | ✓ 798ms | http |
| 34.96.238.40:8080 | ✓ 1014ms | ✓ 1428ms | ✓ 1180ms | ✓ 1215ms | ✓ 1205ms | http |
| 113.176.92.71:3128 | ✓ 1542ms | ✓ 1551ms | ✓ 1474ms | ✓ 1322ms | ✓ 1068ms | http |
| 115.231.181.40:8128 | ✓ 1680ms | ✓ 1513ms | ✓ 1974ms | 否 | 否 | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1210ms | ✓ 1194ms | ✓ 941ms | http |
| 160.250.5.22:1 | ✓ 1456ms | 否 | ✓ 1467ms | ✓ 1348ms | ✓ 1160ms | http |
| 38.34.179.106:8450 | ✓ 868ms | 否 | ✓ 1587ms | ✓ 1049ms | ✓ 1291ms | http |
| 38.34.179.21:8446 | ✓ 870ms | ✓ 1613ms | ✓ 1992ms | ✓ 1013ms | ✓ 1005ms | http |
| 38.34.179.54:8446 | ✓ 871ms | ✓ 1195ms | 否 | ✓ 1157ms | ✓ 1257ms | http |
| 38.145.208.220:8448 | ✓ 1242ms | ✓ 843ms | ✓ 1023ms | 否 | ✓ 828ms | http |
| 38.34.179.100:8449 | ✓ 872ms | ✓ 1927ms | ✓ 1810ms | ✓ 1030ms | ✓ 1848ms | http |
| 38.145.203.45:8452 | ✓ 1594ms | ✓ 1057ms | ✓ 1309ms | 否 | ✓ 738ms | http |
| 38.145.203.46:8448 | ✓ 1596ms | ✓ 1771ms | 否 | ✓ 885ms | ✓ 1174ms | http |
| 38.145.208.172:8448 | ✓ 1688ms | ✓ 1136ms | ✓ 1956ms | ✓ 1509ms | ✓ 812ms | http |
| 45.136.130.171:8445 | ✓ 1892ms | ✓ 848ms | ✓ 1407ms | 否 | ✓ 768ms | http |
| 45.136.130.186:8451 | ✓ 1412ms | 否 | ✓ 772ms | 否 | ✓ 1370ms | http |
| 38.145.220.41:8444 | 否 | ✓ 1846ms | ✓ 1419ms | 否 | ✓ 716ms | http |
| 45.140.147.155:1082 | ✓ 1014ms | 否 | ✓ 425ms | ✓ 1625ms | ✓ 1370ms | http |
| 106.117.208.101:7890 | ✓ 1075ms | ✓ 1385ms | ✓ 1154ms | ✓ 1485ms | ✓ 1087ms | http |
| 116.80.65.77:3172 | ✓ 1658ms | 否 | 否 | ✓ 1996ms | ✓ 1760ms | http |
| 167.160.191.204:6005 | ✓ 577ms | 否 | ✓ 1027ms | ✓ 1102ms | ✓ 1286ms | http |
| 168.222.254.26:8888 | ✓ 759ms | ✓ 1830ms | ✓ 1622ms | 否 | 否 | http |
| 209.97.149.157:80 | ✓ 263ms | ✓ 981ms | ✓ 1176ms | ✓ 972ms | ✓ 709ms | http |
| 104.248.151.93:9090 | ✓ 866ms | 否 | ✓ 940ms | ✓ 1219ms | 否 | http |
| 59.8.203.55:80 | ✓ 1221ms | ✓ 1452ms | ✓ 1718ms | ✓ 1182ms | ✓ 882ms | http |
| 210.223.44.230:3128 | ✓ 784ms | ✓ 1356ms | ✓ 1347ms | 否 | 否 | http |
| 5.102.109.41:999 | ✓ 442ms | ✓ 1228ms | ✓ 682ms | 否 | ✓ 1756ms | http |
| 218.60.0.214:80 | ✓ 1507ms | ✓ 1859ms | ✓ 1466ms | ✓ 1709ms | ✓ 1275ms | http |
| 85.208.108.43:2094 | ✓ 379ms | 否 | ✓ 635ms | ✓ 963ms | ✓ 1565ms | http |
| 38.145.208.210:8448 | ✓ 1238ms | ✓ 1961ms | ✓ 1905ms | ✓ 1033ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1670ms | 否 | 否 | ✓ 1660ms | ✓ 1541ms | http |
| 59.46.216.131:30001 | ✓ 1082ms | 否 | ✓ 1223ms | ✓ 1495ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1296ms | 否 | ✓ 1255ms | ✓ 1633ms | ✓ 1279ms | http |
| 103.113.70.189:1081 | ✓ 517ms | 否 | ✓ 572ms | 否 | ✓ 755ms | http |

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
