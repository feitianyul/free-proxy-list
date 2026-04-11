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

最后更新：2026-04-11 15:28:28 UTC（2026-04-11 23:28:28 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1183ms | 否 | ✓ 1071ms | ✓ 1165ms | ✓ 1210ms | http |
| 155.117.18.36:25388 | ✓ 1090ms | ✓ 1268ms | 否 | ✓ 1605ms | ✓ 1894ms | http |
| 167.103.115.102:8800 | ✓ 1835ms | 否 | ✓ 1201ms | ✓ 1507ms | ✓ 1322ms | http |
| 5.104.87.17:8051 | ✓ 1950ms | 否 | ✓ 1625ms | ✓ 1828ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1873ms | 否 | ✓ 1584ms | 否 | ✓ 1824ms | http |
| 137.59.47.73:3128 | ✓ 1881ms | 否 | 否 | ✓ 1553ms | ✓ 1291ms | http |
| 104.129.203.245:10733 | ✓ 578ms | ✓ 1018ms | ✓ 907ms | 否 | 否 | http |
| 104.129.203.245:10026 | ✓ 620ms | ✓ 996ms | ✓ 945ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 581ms | ✓ 1494ms | ✓ 826ms | 否 | ✓ 945ms | http |
| 113.160.132.26:8080 | ✓ 1088ms | ✓ 1606ms | 否 | 否 | ✓ 1166ms | http |
| 167.103.144.127:8800 | ✓ 1636ms | 否 | ✓ 1464ms | ✓ 1753ms | ✓ 1799ms | http |
| 130.61.30.221:8080 | ✓ 893ms | 否 | ✓ 1266ms | ✓ 1747ms | 否 | http |
| 45.167.124.52:8080 | 否 | ✓ 1546ms | ✓ 506ms | ✓ 1540ms | ✓ 1279ms | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1169ms | ✓ 1551ms | ✓ 1069ms | http |
| 45.167.125.21:999 | 否 | ✓ 1934ms | ✓ 1259ms | ✓ 1668ms | ✓ 1524ms | http |
| 91.238.105.43:2023 | ✓ 768ms | ✓ 1772ms | ✓ 1869ms | 否 | ✓ 1775ms | http |
| 167.103.31.122:8800 | ✓ 1497ms | 否 | ✓ 1265ms | ✓ 1506ms | ✓ 1461ms | http |
| 45.12.151.226:2829 | ✓ 1421ms | 否 | ✓ 1585ms | 否 | ✓ 1459ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1522ms | ✓ 1257ms | ✓ 1177ms | 否 | http |
| 147.161.239.240:8800 | ✓ 1240ms | 否 | ✓ 1352ms | ✓ 1491ms | ✓ 1745ms | http |
| 162.240.154.26:3128 | ✓ 1056ms | ✓ 1569ms | ✓ 1642ms | ✓ 1241ms | ✓ 1928ms | http |
| 45.136.130.181:8445 | ✓ 441ms | 否 | ✓ 1766ms | ✓ 1018ms | ✓ 921ms | http |
| 202.141.161.53:10808 | ✓ 1179ms | 否 | 否 | ✓ 1820ms | ✓ 1095ms | http |
| 45.136.131.49:8445 | 否 | 否 | ✓ 734ms | ✓ 1027ms | ✓ 772ms | http |
| 38.34.179.38:8447 | 否 | ✓ 1284ms | ✓ 1232ms | ✓ 1732ms | ✓ 763ms | http |
| 38.145.220.103:8446 | 否 | 否 | ✓ 1551ms | ✓ 1284ms | ✓ 747ms | http |
| 38.34.179.46:8448 | ✓ 498ms | 否 | ✓ 410ms | ✓ 1137ms | ✓ 1397ms | http |
| 45.136.131.28:8449 | ✓ 460ms | ✓ 1324ms | ✓ 1787ms | ✓ 1265ms | ✓ 1465ms | http |
| 45.136.131.36:8450 | ✓ 1249ms | 否 | ✓ 1306ms | 否 | ✓ 1152ms | http |
| 210.223.44.230:3128 | ✓ 1804ms | ✓ 1771ms | ✓ 1571ms | ✓ 1181ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1025ms | 否 | ✓ 1039ms | ✓ 1972ms | 否 | http |
| 51.81.6.158:3128 | ✓ 1616ms | 否 | 否 | ✓ 1500ms | ✓ 1370ms | http |
| 147.161.152.13:10285 | ✓ 481ms | 否 | ✓ 1221ms | 否 | ✓ 877ms | http |
| 165.225.194.14:10610 | ✓ 422ms | ✓ 1461ms | ✓ 1182ms | ✓ 1781ms | ✓ 1962ms | http |
| 147.161.130.13:11101 | ✓ 416ms | 否 | ✓ 886ms | 否 | ✓ 1700ms | http |
| 101.32.244.83:8080 | ✓ 1137ms | 否 | ✓ 1178ms | ✓ 1744ms | ✓ 1506ms | http |
| 121.43.196.210:8222 | ✓ 1063ms | ✓ 1308ms | ✓ 1069ms | ✓ 1316ms | ✓ 1115ms | http |
| 121.43.196.213:8222 | ✓ 1095ms | ✓ 1258ms | ✓ 1091ms | ✓ 1418ms | ✓ 1077ms | http |
| 114.55.226.123:10086 | ✓ 1294ms | ✓ 1620ms | ✓ 1258ms | ✓ 1514ms | ✓ 1262ms | http |
| 101.43.127.100:8877 | ✓ 1095ms | ✓ 1360ms | ✓ 1734ms | ✓ 1293ms | ✓ 1171ms | http |
| 45.136.130.179:8450 | ✓ 1427ms | ✓ 1534ms | ✓ 430ms | ✓ 1195ms | ✓ 1740ms | http |
| 103.157.200.126:3128 | ✓ 1149ms | 否 | ✓ 1130ms | ✓ 1575ms | ✓ 1275ms | http |
| 38.145.220.182:8450 | ✓ 1757ms | 否 | ✓ 485ms | ✓ 1502ms | ✓ 1341ms | http |
| 46.30.46.133:3128 | ✓ 545ms | 否 | ✓ 1773ms | ✓ 1981ms | ✓ 1425ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1956ms | ✓ 1636ms | ✓ 1256ms | http |
| 218.108.131.186:17890 | ✓ 1003ms | ✓ 1299ms | ✓ 1025ms | ✓ 1307ms | ✓ 1076ms | http |
| 120.92.212.16:7890 | ✓ 1938ms | 否 | ✓ 1132ms | ✓ 1489ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1914ms | ✓ 1450ms | ✓ 1224ms | 否 | ✓ 1196ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1707ms | ✓ 1428ms | ✓ 1561ms | http |
| 38.145.203.45:8449 | ✓ 1206ms | 否 | ✓ 364ms | ✓ 1206ms | ✓ 979ms | http |
| 38.145.203.34:8449 | ✓ 1206ms | 否 | ✓ 364ms | ✓ 1229ms | ✓ 968ms | http |
| 38.34.179.23:8451 | ✓ 1506ms | ✓ 1395ms | ✓ 992ms | ✓ 1378ms | ✓ 841ms | http |
| 38.145.208.213:8445 | ✓ 934ms | 否 | 否 | ✓ 1078ms | ✓ 791ms | http |
| 38.145.203.41:8453 | 否 | 否 | ✓ 409ms | ✓ 979ms | ✓ 1757ms | http |
| 45.136.130.183:8447 | 否 | 否 | ✓ 390ms | ✓ 998ms | ✓ 853ms | http |
| 45.136.130.177:8447 | 否 | 否 | ✓ 384ms | ✓ 1019ms | ✓ 829ms | http |
| 38.145.220.79:8450 | ✓ 1768ms | 否 | ✓ 565ms | ✓ 1011ms | ✓ 1528ms | http |
| 38.145.220.43:8450 | ✓ 1767ms | 否 | ✓ 554ms | ✓ 945ms | ✓ 1371ms | http |
| 38.145.208.207:8445 | ✓ 934ms | 否 | 否 | ✓ 1062ms | ✓ 814ms | http |
| 38.145.218.238:8450 | 否 | 否 | ✓ 1453ms | ✓ 1425ms | ✓ 1294ms | http |
| 45.136.130.184:8447 | 否 | 否 | ✓ 394ms | ✓ 970ms | ✓ 1002ms | http |
| 38.145.208.185:8449 | ✓ 980ms | ✓ 1318ms | ✓ 1120ms | ✓ 1376ms | ✓ 1209ms | http |
| 45.136.130.252:8453 | ✓ 1503ms | 否 | ✓ 1840ms | ✓ 1298ms | ✓ 1361ms | http |
| 38.145.208.187:8449 | ✓ 1806ms | ✓ 1459ms | ✓ 495ms | ✓ 1572ms | ✓ 1246ms | http |
| 38.145.208.178:8449 | ✓ 1806ms | 否 | ✓ 863ms | ✓ 1423ms | ✓ 1229ms | http |
| 38.145.220.20:8450 | 否 | 否 | ✓ 1523ms | ✓ 1377ms | ✓ 1124ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1561ms | 否 | ✓ 1633ms | ✓ 1950ms | http |
| 38.34.179.77:8448 | ✓ 1329ms | 否 | ✓ 342ms | ✓ 1074ms | 否 | http |
| 38.145.208.214:8446 | ✓ 1540ms | ✓ 1311ms | ✓ 1912ms | ✓ 1008ms | ✓ 843ms | http |
| 38.145.208.212:8446 | ✓ 1502ms | ✓ 1853ms | ✓ 1446ms | ✓ 1065ms | ✓ 976ms | http |
| 45.136.130.247:8448 | ✓ 1219ms | 否 | ✓ 757ms | 否 | ✓ 1668ms | http |
| 160.119.69.7:8080 | ✓ 1013ms | 否 | ✓ 1638ms | 否 | ✓ 1524ms | http |
| 112.209.22.22:8082 | 否 | 否 | ✓ 1721ms | ✓ 1668ms | ✓ 1638ms | http |
| 212.58.132.5:8888 | ✓ 1453ms | 否 | ✓ 1432ms | ✓ 1499ms | ✓ 1321ms | http |
| 38.34.179.105:8447 | ✓ 1579ms | 否 | ✓ 366ms | ✓ 1606ms | ✓ 1487ms | http |
| 107.172.102.234:40621 | ✓ 1002ms | ✓ 1092ms | ✓ 1125ms | ✓ 1079ms | ✓ 914ms | http |
| 38.145.208.210:8448 | ✓ 1306ms | 否 | ✓ 1088ms | ✓ 1072ms | ✓ 923ms | http |
| 165.225.202.12:9480 | ✓ 732ms | 否 | ✓ 1608ms | 否 | ✓ 1458ms | http |
| 147.161.146.13:11197 | ✓ 1251ms | ✓ 1386ms | ✓ 1422ms | 否 | 否 | http |
| 38.34.179.89:8449 | ✓ 1023ms | 否 | ✓ 1266ms | ✓ 1347ms | ✓ 758ms | http |
| 38.145.208.223:8453 | ✓ 1551ms | ✓ 1105ms | ✓ 992ms | ✓ 1795ms | ✓ 791ms | http |
| 181.78.44.63:999 | ✓ 1515ms | 否 | ✓ 1223ms | ✓ 1646ms | ✓ 1667ms | http |
| 223.16.170.103:80 | ✓ 1718ms | 否 | ✓ 1483ms | ✓ 1327ms | ✓ 1330ms | http |

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
