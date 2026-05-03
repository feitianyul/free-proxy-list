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

最后更新：2026-05-03 12:39:39 UTC（2026-05-03 20:39:39 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.140.147.82:1082 | 否 | ✓ 1305ms | ✓ 495ms | 否 | ✓ 1336ms | http |
| 218.108.131.186:17890 | ✓ 1005ms | ✓ 1287ms | ✓ 1034ms | ✓ 1601ms | ✓ 1103ms | http |
| 1.231.81.166:3128 | ✓ 1013ms | ✓ 1287ms | ✓ 1945ms | ✓ 1549ms | ✓ 1410ms | http |
| 8.219.97.248:80 | ✓ 1638ms | 否 | ✓ 1295ms | 否 | ✓ 1531ms | http |
| 113.160.132.26:8080 | ✓ 1857ms | ✓ 1716ms | 否 | ✓ 1417ms | ✓ 1215ms | http |
| 45.167.124.71:999 | ✓ 859ms | ✓ 1643ms | ✓ 1345ms | ✓ 1678ms | ✓ 1368ms | http |
| 206.206.126.177:2412 | ✓ 894ms | 否 | ✓ 876ms | ✓ 1228ms | ✓ 958ms | http |
| 38.180.62.47:10808 | ✓ 1914ms | 否 | ✓ 1519ms | ✓ 1466ms | ✓ 1373ms | http |
| 109.120.156.122:8090 | ✓ 1098ms | 否 | ✓ 1530ms | 否 | ✓ 1524ms | http |
| 43.133.44.89:8888 | ✓ 1325ms | 否 | ✓ 882ms | 否 | ✓ 972ms | http |
| 46.105.190.38:3128 | ✓ 1392ms | ✓ 1411ms | ✓ 519ms | 否 | 否 | http |
| 47.85.51.197:1080 | ✓ 1547ms | 否 | ✓ 773ms | ✓ 999ms | 否 | http |
| 130.61.174.200:1080 | 否 | ✓ 1224ms | 否 | ✓ 1476ms | ✓ 1086ms | http |
| 217.182.195.221:30003 | ✓ 944ms | 否 | ✓ 1387ms | 否 | ✓ 1767ms | http |
| 38.188.247.12:999 | ✓ 1847ms | 否 | ✓ 372ms | ✓ 1833ms | ✓ 1409ms | http |
| 61.52.131.172:8443 | ✓ 1293ms | ✓ 1370ms | ✓ 1102ms | ✓ 1369ms | ✓ 1100ms | http |
| 46.105.190.40:3128 | ✓ 905ms | 否 | ✓ 488ms | ✓ 1673ms | ✓ 1293ms | http |
| 220.197.44.36:3128 | ✓ 1824ms | 否 | ✓ 1842ms | 否 | ✓ 1568ms | http |
| 77.110.107.80:1080 | ✓ 614ms | ✓ 1507ms | ✓ 1178ms | 否 | ✓ 1208ms | http |
| 38.180.121.135:10808 | ✓ 462ms | 否 | 否 | ✓ 1809ms | ✓ 1362ms | http |
| 193.123.250.39:1080 | 否 | 否 | ✓ 1888ms | ✓ 1937ms | ✓ 1221ms | http |
| 62.133.60.126:24558 | ✓ 507ms | 否 | ✓ 790ms | ✓ 1466ms | ✓ 1213ms | http |
| 77.110.107.80:8080 | ✓ 1035ms | ✓ 1858ms | ✓ 1043ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1033ms | 否 | ✓ 1980ms | 否 | ✓ 1856ms | http |
| 20.164.75.153:8080 | ✓ 1538ms | 否 | 否 | ✓ 1955ms | ✓ 1784ms | http |
| 86.104.74.110:1081 | ✓ 1059ms | ✓ 1448ms | 否 | ✓ 1555ms | ✓ 1216ms | http |
| 106.10.55.212:1121 | ✓ 1915ms | ✓ 1345ms | ✓ 1264ms | ✓ 1468ms | 否 | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1984ms | ✓ 1848ms | ✓ 1473ms | http |
| 86.104.74.110:1082 | ✓ 597ms | ✓ 1379ms | ✓ 364ms | 否 | 否 | http |
| 34.96.238.40:8080 | 否 | ✓ 1357ms | ✓ 1036ms | ✓ 1250ms | 否 | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1476ms | ✓ 1591ms | ✓ 1247ms | http |
| 49.156.44.114:8080 | ✓ 1664ms | 否 | ✓ 1524ms | ✓ 1711ms | ✓ 1795ms | http |
| 89.208.106.138:10808 | 否 | ✓ 1349ms | 否 | ✓ 1636ms | ✓ 1078ms | http |
| 62.60.231.71:56608 | ✓ 892ms | 否 | ✓ 919ms | 否 | ✓ 1113ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1003ms | ✓ 134ms | ✓ 1102ms | 否 | http |
| 8.154.21.175:3128 | ✓ 1031ms | ✓ 1314ms | ✓ 1120ms | ✓ 1557ms | ✓ 1098ms | http |
| 45.125.67.37:8443 | ✓ 1053ms | 否 | ✓ 1300ms | ✓ 1434ms | ✓ 1311ms | http |
| 121.230.8.220:1080 | 否 | ✓ 1493ms | ✓ 1752ms | ✓ 1612ms | 否 | http |
| 41.196.16.230:1976 | ✓ 1730ms | 否 | ✓ 1637ms | 否 | ✓ 1909ms | http |
| 104.248.195.47:8080 | ✓ 478ms | ✓ 1426ms | ✓ 1449ms | ✓ 1459ms | ✓ 1138ms | http |
| 72.56.87.46:3128 | ✓ 901ms | 否 | 否 | ✓ 1738ms | ✓ 1436ms | http |
| 62.60.149.161:3128 | ✓ 974ms | 否 | ✓ 685ms | 否 | ✓ 1964ms | http |
| 154.90.48.209:9090 | ✓ 1637ms | 否 | ✓ 1288ms | ✓ 1497ms | ✓ 1212ms | http |
| 152.42.177.32:8888 | ✓ 1143ms | 否 | ✓ 1352ms | ✓ 1569ms | ✓ 1531ms | http |
| 101.32.243.189:80 | ✓ 1416ms | 否 | ✓ 1475ms | ✓ 1727ms | ✓ 1874ms | http |
| 80.92.204.47:1081 | ✓ 1029ms | ✓ 1299ms | ✓ 949ms | ✓ 1778ms | ✓ 1304ms | http |
| 150.249.255.91:3128 | ✓ 1292ms | 否 | ✓ 738ms | ✓ 1071ms | ✓ 893ms | http |
| 3.101.133.120:80 | ✓ 968ms | 否 | 否 | ✓ 1250ms | ✓ 1442ms | http |
| 86.104.72.220:1081 | ✓ 511ms | ✓ 1414ms | ✓ 251ms | ✓ 983ms | 否 | http |
| 86.104.72.220:1082 | ✓ 498ms | ✓ 1407ms | ✓ 70ms | ✓ 1168ms | 否 | http |
| 154.64.232.35:8080 | 否 | 否 | ✓ 1982ms | ✓ 1780ms | ✓ 687ms | http |
| 45.129.141.143:3128 | ✓ 1493ms | ✓ 1703ms | ✓ 1910ms | ✓ 1988ms | ✓ 1478ms | http |
| 45.186.6.104:3128 | ✓ 1489ms | ✓ 1887ms | ✓ 1847ms | 否 | 否 | http |
| 139.159.97.82:10900 | ✓ 1424ms | 否 | ✓ 1577ms | ✓ 1750ms | ✓ 1327ms | http |
| 120.92.212.16:7890 | ✓ 1979ms | ✓ 1905ms | 否 | 否 | ✓ 1195ms | http |
| 47.77.216.82:1080 | ✓ 283ms | ✓ 1285ms | ✓ 920ms | ✓ 1721ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1239ms | ✓ 1410ms | 否 | ✓ 1448ms | ✓ 1440ms | http |
| 45.140.147.155:1082 | ✓ 1385ms | 否 | ✓ 1639ms | 否 | ✓ 1758ms | http |
| 20.127.128.70:8080 | ✓ 1461ms | 否 | ✓ 935ms | 否 | ✓ 1534ms | http |
| 103.157.200.126:3128 | ✓ 1542ms | 否 | ✓ 1182ms | ✓ 1666ms | ✓ 1222ms | http |
| 59.46.216.131:30001 | ✓ 1232ms | ✓ 1703ms | ✓ 1242ms | 否 | ✓ 1279ms | http |
| 148.230.4.241:999 | ✓ 1336ms | ✓ 1941ms | ✓ 658ms | 否 | 否 | http |
| 207.254.71.62:8088 | ✓ 1247ms | ✓ 1521ms | ✓ 1438ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 1267ms | ✓ 1593ms | ✓ 881ms | 否 | 否 | http |
| 45.153.231.229:8080 | ✓ 991ms | 否 | ✓ 1691ms | 否 | ✓ 1744ms | http |

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
