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

最后更新：2026-03-11 19:50:12 UTC（2026-03-12 03:50:12 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 574ms | ✓ 831ms | ✓ 1227ms | ✓ 913ms | ✓ 721ms | http |
| 45.136.131.63:8443 | ✓ 580ms | ✓ 852ms | ✓ 1200ms | ✓ 1080ms | ✓ 803ms | http |
| 152.42.213.210:8080 | ✓ 1520ms | 否 | 否 | ✓ 1604ms | ✓ 1239ms | http |
| 171.251.172.78:5109 | 否 | 否 | ✓ 1534ms | ✓ 1902ms | ✓ 1584ms | http |
| 205.209.118.30:3138 | ✓ 280ms | ✓ 922ms | ✓ 828ms | ✓ 1130ms | ✓ 880ms | http |
| 35.225.22.61:80 | ✓ 468ms | ✓ 1317ms | ✓ 963ms | ✓ 1183ms | 否 | http |
| 103.113.70.189:1081 | ✓ 274ms | ✓ 933ms | 否 | ✓ 1070ms | ✓ 954ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1208ms | ✓ 1191ms | ✓ 893ms | http |
| 38.180.2.107:3128 | ✓ 843ms | ✓ 1629ms | ✓ 1335ms | 否 | 否 | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1874ms | ✓ 1759ms | ✓ 1540ms | http |
| 45.136.130.175:8443 | ✓ 514ms | ✓ 873ms | ✓ 503ms | ✓ 997ms | ✓ 679ms | http |
| 45.136.130.188:8443 | ✓ 516ms | ✓ 883ms | ✓ 496ms | ✓ 992ms | ✓ 772ms | http |
| 45.136.130.191:8443 | ✓ 520ms | ✓ 934ms | ✓ 442ms | ✓ 994ms | ✓ 792ms | http |
| 45.136.130.223:8443 | ✓ 1162ms | 否 | ✓ 749ms | ✓ 1522ms | ✓ 1421ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1156ms | ✓ 1462ms | ✓ 1053ms | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1398ms | ✓ 1430ms | ✓ 1120ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1595ms | ✓ 1471ms | ✓ 1899ms | http |
| 160.238.65.9:3128 | ✓ 1020ms | 否 | ✓ 1385ms | 否 | ✓ 1718ms | http |
| 160.238.65.2:3128 | ✓ 1018ms | 否 | ✓ 1386ms | 否 | ✓ 1716ms | http |
| 190.9.109.198:999 | ✓ 904ms | ✓ 1416ms | ✓ 1320ms | ✓ 1278ms | ✓ 1306ms | http |
| 185.115.74.185:8080 | ✓ 1009ms | ✓ 1640ms | ✓ 1693ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1257ms | ✓ 1634ms | ✓ 1230ms | ✓ 1550ms | ✓ 1285ms | http |
| 91.107.141.42:8081 | 否 | 否 | ✓ 1097ms | ✓ 1356ms | ✓ 1624ms | http |
| 1.231.81.166:3128 | ✓ 754ms | ✓ 1202ms | ✓ 984ms | 否 | 否 | http |
| 111.48.191.1:7890 | ✓ 925ms | ✓ 1124ms | ✓ 1019ms | 否 | 否 | http |
| 152.42.213.210:443 | ✓ 948ms | 否 | ✓ 1851ms | ✓ 1313ms | ✓ 987ms | http |
| 101.43.255.96:80 | ✓ 1197ms | ✓ 1556ms | ✓ 1218ms | ✓ 1549ms | ✓ 1242ms | http |
| 202.155.12.161:443 | ✓ 1886ms | 否 | ✓ 1349ms | ✓ 1382ms | ✓ 1178ms | http |
| 194.213.18.200:443 | ✓ 474ms | ✓ 1697ms | ✓ 1983ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1176ms | ✓ 1473ms | ✓ 1139ms | ✓ 1428ms | ✓ 1213ms | http |
| 101.47.73.135:3128 | ✓ 1114ms | 否 | ✓ 1559ms | ✓ 1528ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1397ms | ✓ 1562ms | ✓ 1759ms | 否 | 否 | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1143ms | ✓ 1143ms | ✓ 923ms | http |
| 45.136.198.40:3128 | ✓ 771ms | ✓ 1615ms | ✓ 1729ms | 否 | 否 | http |
| 113.160.132.26:8080 | ✓ 1738ms | ✓ 1517ms | ✓ 1841ms | ✓ 1464ms | ✓ 1169ms | http |
| 111.79.111.126:3128 | ✓ 1407ms | ✓ 1496ms | ✓ 1660ms | ✓ 1884ms | ✓ 1275ms | http |
| 47.77.193.180:1080 | ✓ 797ms | ✓ 1697ms | ✓ 400ms | ✓ 899ms | ✓ 671ms | http |
| 67.169.98.211:443 | ✓ 1050ms | 否 | 否 | ✓ 1765ms | ✓ 1593ms | http |
| 168.235.110.63:3128 | ✓ 269ms | 否 | ✓ 920ms | 否 | ✓ 1703ms | http |
| 190.212.131.238:3128 | ✓ 1422ms | 否 | ✓ 1419ms | 否 | ✓ 1500ms | http |
| 160.238.65.7:3128 | ✓ 1551ms | 否 | ✓ 999ms | 否 | ✓ 1119ms | http |
| 107.173.52.58:7890 | ✓ 1011ms | 否 | ✓ 1007ms | ✓ 1242ms | ✓ 952ms | http |
| 139.159.99.242:8080 | ✓ 1770ms | 否 | ✓ 988ms | ✓ 1271ms | ✓ 983ms | http |
| 121.230.9.139:1080 | ✓ 1188ms | ✓ 1683ms | ✓ 1376ms | ✓ 1710ms | ✓ 1441ms | http |
| 160.238.65.5:3128 | ✓ 947ms | ✓ 1819ms | ✓ 1390ms | 否 | ✓ 1520ms | http |
| 160.238.65.3:3128 | ✓ 948ms | 否 | ✓ 1208ms | 否 | ✓ 1520ms | http |
| 160.238.65.6:3128 | ✓ 947ms | ✓ 1713ms | ✓ 1497ms | 否 | ✓ 1525ms | http |
| 160.238.65.4:3128 | ✓ 947ms | ✓ 1785ms | ✓ 1424ms | 否 | ✓ 1518ms | http |
| 160.238.65.8:3128 | ✓ 951ms | 否 | ✓ 1204ms | 否 | ✓ 1524ms | http |
| 185.191.236.162:3128 | ✓ 968ms | 否 | ✓ 1693ms | ✓ 1532ms | ✓ 1657ms | http |
| 121.230.9.160:1080 | ✓ 1387ms | ✓ 1596ms | ✓ 1453ms | ✓ 1473ms | ✓ 1322ms | http |
| 103.39.51.190:8080 | ✓ 1540ms | 否 | 否 | ✓ 1582ms | ✓ 1739ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1597ms | ✓ 1531ms | ✓ 1251ms | http |
| 116.80.96.106:3172 | ✓ 1901ms | 否 | ✓ 1704ms | 否 | ✓ 1833ms | http |
| 178.236.245.59:3128 | ✓ 1864ms | 否 | ✓ 1984ms | ✓ 1887ms | ✓ 1728ms | http |
| 106.117.208.101:7890 | 否 | 否 | ✓ 1684ms | ✓ 1586ms | ✓ 1226ms | http |
| 121.138.61.193:8235 | ✓ 955ms | ✓ 1487ms | ✓ 1299ms | ✓ 1359ms | ✓ 1204ms | http |
| 172.212.68.37:3128 | ✓ 309ms | ✓ 1610ms | ✓ 1225ms | ✓ 1392ms | ✓ 899ms | http |
| 144.31.25.69:21064 | ✓ 1036ms | 否 | ✓ 1203ms | 否 | ✓ 1972ms | http |
| 95.3.9.78:3128 | ✓ 1228ms | 否 | ✓ 1868ms | ✓ 1667ms | 否 | http |
| 95.3.9.78:8080 | ✓ 1999ms | 否 | 否 | ✓ 1611ms | ✓ 1934ms | http |
| 61.52.131.172:8443 | ✓ 1064ms | ✓ 1382ms | ✓ 1150ms | ✓ 1321ms | ✓ 1105ms | http |
| 88.80.150.82:8080 | ✓ 1208ms | ✓ 1869ms | 否 | 否 | ✓ 1553ms | https |
| 43.165.195.107:3128 | ✓ 1767ms | ✓ 1916ms | ✓ 1337ms | ✓ 1380ms | ✓ 1119ms | http |

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
