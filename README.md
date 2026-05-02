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

最后更新：2026-05-02 23:39:00 UTC（2026-05-03 07:39:00 UTC+8）

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
| 149.51.42.10:3128 | ✓ 1654ms | ✓ 1712ms | 否 | ✓ 1458ms | 否 | http |
| 206.206.126.177:2412 | ✓ 1748ms | 否 | ✓ 1680ms | ✓ 1258ms | ✓ 1008ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1702ms | ✓ 1413ms | 否 | ✓ 1676ms | http |
| 148.230.4.241:999 | ✓ 937ms | 否 | ✓ 649ms | ✓ 1696ms | ✓ 1490ms | http |
| 45.167.124.71:999 | ✓ 1114ms | ✓ 1735ms | ✓ 1324ms | ✓ 1918ms | ✓ 1588ms | http |
| 47.85.51.197:1080 | ✓ 316ms | ✓ 1054ms | ✓ 341ms | 否 | ✓ 1028ms | http |
| 47.77.216.82:1080 | 否 | ✓ 1339ms | ✓ 869ms | ✓ 1092ms | ✓ 839ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1583ms | ✓ 490ms | ✓ 1125ms | 否 | http |
| 92.119.127.208:6005 | ✓ 1326ms | ✓ 1704ms | ✓ 1586ms | ✓ 1882ms | 否 | http |
| 20.164.75.153:8080 | ✓ 1458ms | 否 | ✓ 1162ms | 否 | ✓ 1761ms | http |
| 86.104.72.220:1081 | ✓ 619ms | ✓ 1289ms | ✓ 1067ms | 否 | ✓ 1090ms | http |
| 34.96.238.40:8080 | ✓ 1481ms | ✓ 1348ms | ✓ 1176ms | 否 | ✓ 1218ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1082ms | ✓ 1658ms | ✓ 1580ms | http |
| 147.45.178.211:14658 | ✓ 989ms | 否 | ✓ 639ms | ✓ 1973ms | ✓ 1357ms | http |
| 218.108.131.186:17890 | ✓ 1078ms | ✓ 1205ms | ✓ 1156ms | ✓ 1189ms | ✓ 1068ms | http |
| 223.16.170.103:80 | ✓ 1612ms | ✓ 1940ms | ✓ 1272ms | ✓ 1471ms | ✓ 1351ms | http |
| 103.35.190.69:1082 | ✓ 353ms | ✓ 940ms | ✓ 106ms | 否 | 否 | http |
| 45.59.122.132:80 | ✓ 1414ms | ✓ 1667ms | ✓ 824ms | 否 | ✓ 1122ms | http |
| 1.231.81.166:3128 | ✓ 1853ms | ✓ 1108ms | ✓ 1638ms | ✓ 1534ms | ✓ 1621ms | http |
| 177.93.33.55:999 | ✓ 1413ms | ✓ 1905ms | ✓ 1343ms | 否 | ✓ 1740ms | http |
| 130.61.174.200:1080 | ✓ 964ms | ✓ 1420ms | ✓ 1827ms | ✓ 1849ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1911ms | 否 | ✓ 1466ms | ✓ 1623ms | ✓ 1722ms | http |
| 94.131.118.129:1082 | ✓ 1160ms | ✓ 1074ms | ✓ 650ms | ✓ 1468ms | ✓ 1267ms | http |
| 80.92.204.47:1081 | ✓ 891ms | ✓ 1073ms | ✓ 822ms | ✓ 1516ms | ✓ 1452ms | http |
| 94.131.118.129:1081 | ✓ 1156ms | ✓ 1072ms | ✓ 657ms | ✓ 1488ms | ✓ 1181ms | http |
| 121.230.8.171:1080 | ✓ 1513ms | ✓ 1715ms | ✓ 1262ms | ✓ 1538ms | ✓ 1281ms | http |
| 120.92.212.16:8890 | ✓ 1217ms | ✓ 1630ms | 否 | ✓ 1520ms | ✓ 1072ms | http |
| 77.110.119.136:3128 | 否 | 否 | ✓ 203ms | ✓ 1166ms | ✓ 946ms | http |
| 45.153.231.229:8080 | ✓ 1092ms | ✓ 1839ms | ✓ 1530ms | 否 | ✓ 1881ms | http |
| 109.120.156.122:8090 | ✓ 915ms | ✓ 1989ms | ✓ 1421ms | 否 | 否 | http |
| 94.158.219.111:3128 | ✓ 1406ms | ✓ 1769ms | ✓ 954ms | 否 | ✓ 1666ms | http |
| 103.35.190.69:1081 | 否 | ✓ 1247ms | ✓ 58ms | ✓ 1922ms | ✓ 1726ms | http |
| 212.58.132.5:8888 | ✓ 1844ms | ✓ 1926ms | ✓ 1010ms | ✓ 1447ms | ✓ 1374ms | http |
| 120.92.212.16:7890 | ✓ 1250ms | 否 | ✓ 1632ms | ✓ 1453ms | ✓ 1793ms | http |
| 45.140.147.82:1082 | 否 | ✓ 1768ms | ✓ 1943ms | 否 | ✓ 1974ms | http |
| 59.46.216.131:30001 | ✓ 1151ms | ✓ 1670ms | 否 | ✓ 1596ms | 否 | http |
| 149.51.42.10:8080 | ✓ 486ms | ✓ 1395ms | 否 | ✓ 1795ms | 否 | http |
| 101.32.243.189:80 | ✓ 1653ms | 否 | ✓ 1806ms | ✓ 1883ms | ✓ 1577ms | http |
| 154.90.48.209:9090 | ✓ 1647ms | 否 | ✓ 1804ms | ✓ 1634ms | 否 | http |
| 62.133.60.126:24558 | 否 | 否 | ✓ 1284ms | ✓ 1989ms | ✓ 1743ms | http |
| 45.140.147.82:1081 | 否 | ✓ 1680ms | ✓ 1970ms | 否 | ✓ 1275ms | http |
| 117.236.124.166:3128 | ✓ 927ms | 否 | ✓ 1523ms | 否 | ✓ 1669ms | http |
| 86.104.72.220:1082 | ✓ 498ms | ✓ 1224ms | ✓ 1792ms | ✓ 1180ms | 否 | http |
| 3.101.133.120:80 | ✓ 613ms | ✓ 1567ms | ✓ 1779ms | ✓ 1016ms | ✓ 775ms | http |
| 202.129.206.239:3128 | ✓ 1575ms | 否 | ✓ 1451ms | ✓ 1709ms | ✓ 1802ms | http |
| 157.0.142.246:10057 | ✓ 1577ms | ✓ 1400ms | ✓ 1208ms | 否 | 否 | http |
| 92.119.127.211:6005 | ✓ 841ms | ✓ 1602ms | ✓ 1600ms | ✓ 1832ms | 否 | http |
| 8.211.166.184:8081 | ✓ 1732ms | 否 | ✓ 1042ms | ✓ 1081ms | 否 | http |
| 49.156.44.114:8080 | 否 | ✓ 1914ms | 否 | ✓ 1755ms | ✓ 1764ms | http |
| 42.200.76.16:3888 | ✓ 852ms | 否 | ✓ 1254ms | ✓ 1154ms | ✓ 913ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1469ms | ✓ 1819ms | ✓ 1567ms | http |
| 45.129.141.143:3128 | ✓ 1372ms | ✓ 1653ms | 否 | ✓ 1976ms | ✓ 1647ms | http |
| 121.230.8.136:1080 | ✓ 1621ms | ✓ 1893ms | ✓ 1532ms | ✓ 1984ms | ✓ 1528ms | http |
| 160.238.65.4:3128 | ✓ 1073ms | ✓ 1704ms | ✓ 1293ms | ✓ 1764ms | ✓ 1588ms | http |
| 160.238.65.9:3128 | 否 | ✓ 1337ms | 否 | ✓ 1521ms | ✓ 1547ms | http |
| 185.230.191.240:3128 | ✓ 707ms | ✓ 1496ms | ✓ 989ms | 否 | ✓ 1737ms | http |
| 160.238.65.5:3128 | 否 | ✓ 1265ms | ✓ 1533ms | ✓ 1821ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1043ms | 否 | ✓ 1259ms | 否 | ✓ 1270ms | http |
| 91.217.81.131:1080 | ✓ 732ms | ✓ 1852ms | ✓ 1108ms | 否 | 否 | http |
| 160.238.65.6:3128 | 否 | 否 | ✓ 456ms | ✓ 1899ms | ✓ 1106ms | http |
| 160.238.65.3:3128 | 否 | ✓ 1226ms | ✓ 1144ms | ✓ 1953ms | ✓ 986ms | http |
| 160.238.65.8:3128 | ✓ 1434ms | ✓ 1729ms | ✓ 1213ms | ✓ 1960ms | ✓ 964ms | http |
| 38.180.2.107:3128 | ✓ 1734ms | 否 | ✓ 1641ms | ✓ 1983ms | 否 | http |
| 160.238.65.2:3128 | 否 | ✓ 1342ms | ✓ 1236ms | 否 | ✓ 1941ms | http |
| 61.52.131.172:8443 | ✓ 882ms | ✓ 1313ms | ✓ 989ms | 否 | ✓ 924ms | http |
| 45.140.147.155:1081 | ✓ 1726ms | 否 | ✓ 1078ms | 否 | ✓ 1700ms | http |
| 118.113.247.104:1080 | ✓ 1431ms | ✓ 1873ms | ✓ 1524ms | ✓ 1991ms | ✓ 1531ms | http |
| 160.238.65.7:3128 | ✓ 515ms | ✓ 1241ms | ✓ 719ms | ✓ 1392ms | ✓ 1285ms | http |

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
