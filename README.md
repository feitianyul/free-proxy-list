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

最后更新：2026-03-05 19:00:32 UTC（2026-03-06 03:00:32 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 387ms | ✓ 1734ms | ✓ 1012ms | ✓ 1050ms | ✓ 796ms | http |
| 168.235.110.63:3128 | ✓ 762ms | 否 | ✓ 1148ms | ✓ 1250ms | ✓ 892ms | http |
| 45.140.147.155:1081 | ✓ 1570ms | ✓ 1089ms | ✓ 1329ms | 否 | ✓ 1119ms | http |
| 45.136.198.40:3128 | ✓ 680ms | ✓ 1821ms | ✓ 1969ms | ✓ 1832ms | ✓ 1559ms | http |
| 125.128.12.144:3128 | ✓ 1823ms | 否 | ✓ 1141ms | ✓ 1267ms | ✓ 1093ms | http |
| 61.72.110.54:3128 | ✓ 1796ms | 否 | ✓ 1251ms | ✓ 1740ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 805ms | ✓ 1325ms | ✓ 854ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1156ms | ✓ 976ms | ✓ 1151ms | ✓ 876ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1269ms | ✓ 1582ms | ✓ 1972ms | http |
| 150.107.140.238:3128 | ✓ 1022ms | 否 | ✓ 1330ms | 否 | ✓ 1550ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1729ms | ✓ 1997ms | ✓ 993ms | http |
| 120.92.212.16:8890 | ✓ 1052ms | ✓ 1297ms | 否 | ✓ 1554ms | ✓ 1346ms | http |
| 185.191.236.162:3128 | ✓ 682ms | 否 | ✓ 1954ms | 否 | ✓ 1433ms | http |
| 61.72.221.234:3128 | ✓ 1848ms | 否 | 否 | ✓ 1316ms | ✓ 1007ms | http |
| 35.225.22.61:80 | ✓ 757ms | ✓ 1799ms | ✓ 279ms | ✓ 1040ms | 否 | http |
| 81.70.169.194:80 | ✓ 1037ms | ✓ 1362ms | ✓ 1126ms | ✓ 1384ms | ✓ 1126ms | http |
| 101.43.255.96:80 | ✓ 1085ms | ✓ 1296ms | ✓ 1065ms | ✓ 1412ms | ✓ 1163ms | http |
| 138.124.53.25:7443 | ✓ 898ms | ✓ 1610ms | 否 | 否 | ✓ 1034ms | http |
| 91.193.240.157:9877 | ✓ 1190ms | 否 | ✓ 1729ms | 否 | ✓ 1649ms | http |
| 45.186.6.104:3128 | ✓ 1195ms | ✓ 1788ms | ✓ 1700ms | 否 | 否 | http |
| 185.115.74.185:8080 | ✓ 903ms | ✓ 1849ms | ✓ 1947ms | 否 | 否 | http |
| 5.252.33.13:2025 | ✓ 1414ms | 否 | ✓ 1635ms | 否 | ✓ 1581ms | http |
| 107.174.80.186:3128 | ✓ 503ms | ✓ 1770ms | ✓ 975ms | ✓ 1095ms | ✓ 1079ms | http |
| 125.128.12.14:3128 | 否 | 否 | ✓ 1418ms | ✓ 1273ms | ✓ 977ms | http |
| 61.72.221.194:3128 | ✓ 1035ms | ✓ 1893ms | 否 | ✓ 1960ms | 否 | http |
| 223.16.170.103:80 | ✓ 1746ms | ✓ 1970ms | ✓ 1309ms | 否 | 否 | http |
| 121.230.8.111:1080 | ✓ 1714ms | ✓ 1814ms | ✓ 1171ms | ✓ 1567ms | ✓ 1074ms | http |
| 46.249.103.192:443 | ✓ 1531ms | 否 | ✓ 1553ms | ✓ 1619ms | 否 | http |
| 188.132.141.249:443 | ✓ 1535ms | 否 | ✓ 1627ms | 否 | ✓ 1658ms | http |
| 210.223.44.230:3128 | ✓ 1791ms | 否 | ✓ 1594ms | ✓ 1424ms | ✓ 1659ms | http |
| 61.72.221.94:3128 | ✓ 1715ms | 否 | ✓ 1828ms | ✓ 1845ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1914ms | 否 | 否 | ✓ 1205ms | ✓ 1013ms | http |
| 61.72.110.94:3128 | 否 | ✓ 1723ms | ✓ 1446ms | ✓ 1221ms | ✓ 1394ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1301ms | ✓ 1305ms | 否 | ✓ 1058ms | http |
| 14.56.177.44:3128 | ✓ 1561ms | ✓ 1301ms | ✓ 1160ms | ✓ 1492ms | ✓ 1303ms | http |
| 156.225.70.152:39151 | ✓ 1614ms | ✓ 1243ms | ✓ 851ms | 否 | 否 | http |
| 129.159.119.154:1080 | ✓ 387ms | ✓ 1000ms | ✓ 1054ms | ✓ 1034ms | ✓ 664ms | http |
| 209.38.51.97:3128 | ✓ 374ms | 否 | ✓ 1528ms | 否 | ✓ 1150ms | http |
| 154.37.208.132:30000 | ✓ 1191ms | ✓ 1569ms | 否 | ✓ 1471ms | ✓ 1002ms | http |
| 69.48.179.20:3128 | ✓ 216ms | 否 | ✓ 824ms | ✓ 1710ms | 否 | http |
| 91.107.175.112:10801 | ✓ 826ms | ✓ 1786ms | ✓ 1278ms | ✓ 1542ms | ✓ 1419ms | http |
| 47.95.231.180:8084 | ✓ 1070ms | ✓ 1390ms | ✓ 1057ms | ✓ 1388ms | ✓ 1115ms | http |
| 121.128.121.54:3128 | ✓ 1852ms | 否 | ✓ 1140ms | 否 | ✓ 930ms | http |
| 120.55.163.237:10086 | 否 | ✓ 1117ms | ✓ 813ms | ✓ 1790ms | ✓ 920ms | http |
| 47.77.193.180:1080 | ✓ 770ms | ✓ 1063ms | ✓ 289ms | ✓ 886ms | ✓ 689ms | http |
| 180.127.149.244:1080 | ✓ 999ms | ✓ 1168ms | ✓ 1073ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1045ms | 否 | ✓ 1114ms | ✓ 1496ms | 否 | http |
| 103.215.36.88:11929 | ✓ 1123ms | ✓ 1163ms | ✓ 1089ms | ✓ 1158ms | ✓ 1008ms | http |
| 121.230.8.136:1080 | ✓ 1325ms | ✓ 1659ms | ✓ 1318ms | ✓ 1986ms | ✓ 1096ms | http |
| 103.82.23.118:5171 | ✓ 1549ms | 否 | ✓ 1683ms | ✓ 1973ms | ✓ 1735ms | http |
| 154.90.48.209:9090 | ✓ 1033ms | 否 | ✓ 1183ms | ✓ 1531ms | ✓ 1361ms | http |
| 116.58.162.45:3128 | ✓ 1938ms | ✓ 1866ms | ✓ 1747ms | 否 | 否 | http |
| 120.79.99.232:8099 | ✓ 1115ms | ✓ 1404ms | ✓ 1112ms | ✓ 1389ms | ✓ 1130ms | http |
| 185.243.218.43:49153 | ✓ 550ms | 否 | ✓ 1510ms | ✓ 1925ms | ✓ 1629ms | http |
| 103.18.78.250:1111 | ✓ 1929ms | 否 | ✓ 1618ms | ✓ 1650ms | ✓ 1541ms | http |
| 172.212.68.37:3128 | ✓ 275ms | ✓ 1523ms | 否 | ✓ 1157ms | ✓ 833ms | http |
| 175.215.54.252:3040 | ✓ 1753ms | ✓ 1155ms | ✓ 1344ms | ✓ 1923ms | 否 | http |
| 103.215.36.88:19626 | ✓ 1272ms | ✓ 1386ms | ✓ 1103ms | ✓ 1641ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1379ms | ✓ 1923ms | 否 | ✓ 1957ms | 否 | https |
| 183.237.195.130:3128 | ✓ 1308ms | ✓ 1672ms | ✓ 1902ms | ✓ 1704ms | ✓ 1352ms | http |
| 103.39.51.190:8080 | ✓ 1951ms | 否 | 否 | ✓ 1545ms | ✓ 1712ms | http |
| 62.113.119.14:8080 | ✓ 1318ms | 否 | ✓ 926ms | ✓ 1761ms | ✓ 1149ms | http |
| 103.215.36.88:18989 | ✓ 1138ms | ✓ 1495ms | ✓ 977ms | ✓ 1329ms | ✓ 977ms | http |
| 45.129.141.143:3128 | ✓ 958ms | ✓ 1600ms | ✓ 1655ms | ✓ 1638ms | ✓ 1602ms | http |
| 39.104.201.40:7890 | ✓ 1072ms | ✓ 1425ms | ✓ 1050ms | ✓ 1398ms | ✓ 1130ms | http |
| 165.225.112.42:13681 | ✓ 896ms | ✓ 1690ms | ✓ 1657ms | ✓ 1252ms | ✓ 989ms | http |
| 165.225.112.42:11691 | ✓ 888ms | 否 | ✓ 1464ms | ✓ 1246ms | ✓ 984ms | http |
| 165.225.112.42:11975 | ✓ 894ms | 否 | ✓ 1534ms | ✓ 1243ms | ✓ 985ms | http |
| 116.80.82.219:3172 | ✓ 1901ms | 否 | ✓ 1720ms | 否 | ✓ 1832ms | http |

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
