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

最后更新：2026-05-04 10:16:45 UTC（2026-05-04 18:16:45 UTC+8）

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
| 34.71.229.255:3128 | ✓ 941ms | 否 | ✓ 931ms | ✓ 1226ms | ✓ 890ms | http |
| 218.108.131.186:17890 | ✓ 1468ms | ✓ 1148ms | ✓ 1012ms | ✓ 1264ms | ✓ 1007ms | http |
| 38.180.62.47:10808 | ✓ 1104ms | 否 | ✓ 1217ms | ✓ 1767ms | ✓ 1453ms | http |
| 62.113.119.14:8080 | ✓ 1203ms | 否 | ✓ 760ms | 否 | ✓ 1298ms | http |
| 113.160.132.26:8080 | ✓ 1576ms | ✓ 1577ms | 否 | ✓ 1604ms | ✓ 1121ms | http |
| 94.131.118.39:1081 | ✓ 1185ms | 否 | ✓ 1148ms | 否 | ✓ 1490ms | http |
| 38.180.121.135:10808 | ✓ 1785ms | ✓ 1572ms | ✓ 1631ms | 否 | 否 | http |
| 101.6.65.112:10080 | ✓ 1509ms | 否 | 否 | ✓ 1510ms | ✓ 1149ms | http |
| 47.77.216.82:1080 | 否 | 否 | ✓ 1036ms | ✓ 976ms | ✓ 872ms | http |
| 148.230.4.241:999 | 否 | ✓ 1780ms | ✓ 851ms | ✓ 1890ms | 否 | http |
| 45.167.124.71:999 | 否 | 否 | ✓ 1473ms | ✓ 1750ms | ✓ 1608ms | http |
| 181.119.97.24:999 | 否 | 否 | ✓ 1606ms | ✓ 1825ms | ✓ 1761ms | http |
| 62.133.60.126:24558 | ✓ 818ms | 否 | ✓ 1764ms | ✓ 1482ms | ✓ 1764ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1689ms | 否 | ✓ 1523ms | ✓ 809ms | http |
| 45.153.231.229:8080 | ✓ 1561ms | ✓ 1859ms | ✓ 1895ms | 否 | ✓ 1978ms | http |
| 154.64.232.35:8080 | ✓ 1589ms | ✓ 965ms | 否 | 否 | ✓ 1155ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1108ms | ✓ 1330ms | ✓ 1280ms | http |
| 45.59.122.132:80 | ✓ 1271ms | 否 | ✓ 804ms | ✓ 1759ms | ✓ 1501ms | http |
| 38.95.77.85:6005 | ✓ 676ms | ✓ 1910ms | ✓ 1476ms | 否 | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1722ms | ✓ 1683ms | ✓ 1376ms | ✓ 1548ms | http |
| 193.123.250.39:1080 | ✓ 977ms | 否 | ✓ 1737ms | 否 | ✓ 1216ms | http |
| 8.219.97.248:80 | ✓ 1805ms | 否 | ✓ 1607ms | ✓ 1354ms | 否 | http |
| 47.85.51.197:1080 | ✓ 273ms | ✓ 1637ms | 否 | 否 | ✓ 1547ms | http |
| 103.157.200.126:3128 | ✓ 1715ms | 否 | ✓ 1596ms | ✓ 1794ms | ✓ 1388ms | http |
| 86.104.72.219:1081 | ✓ 661ms | ✓ 927ms | ✓ 1019ms | ✓ 1281ms | 否 | http |
| 45.140.147.155:1082 | ✓ 423ms | ✓ 1121ms | ✓ 1469ms | ✓ 1736ms | ✓ 1491ms | http |
| 206.206.126.177:2412 | ✓ 1668ms | ✓ 1738ms | ✓ 1105ms | ✓ 1134ms | ✓ 913ms | http |
| 20.127.128.70:8080 | ✓ 1596ms | 否 | ✓ 1766ms | 否 | ✓ 1745ms | http |
| 109.120.156.122:8090 | ✓ 923ms | ✓ 1654ms | ✓ 728ms | 否 | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1206ms | ✓ 1936ms | ✓ 1527ms | http |
| 57.128.188.167:9172 | ✓ 1720ms | 否 | ✓ 1536ms | 否 | ✓ 1708ms | http |
| 130.61.174.200:1080 | ✓ 1127ms | ✓ 1331ms | ✓ 457ms | 否 | 否 | http |
| 46.105.190.40:3128 | ✓ 508ms | ✓ 1509ms | ✓ 771ms | ✓ 1734ms | ✓ 1315ms | http |
| 8.154.21.175:3128 | ✓ 1008ms | ✓ 1252ms | ✓ 963ms | ✓ 1262ms | ✓ 1062ms | http |
| 160.250.134.143:3128 | ✓ 1644ms | 否 | ✓ 1157ms | ✓ 1524ms | ✓ 1082ms | http |
| 120.92.212.16:8890 | ✓ 1284ms | ✓ 1701ms | 否 | ✓ 1650ms | 否 | http |
| 147.45.178.211:14658 | ✓ 778ms | ✓ 1651ms | 否 | 否 | ✓ 1728ms | http |
| 217.77.102.18:3128 | ✓ 1190ms | 否 | ✓ 1938ms | 否 | ✓ 1684ms | http |
| 137.59.47.73:3128 | ✓ 1722ms | ✓ 1417ms | ✓ 1419ms | ✓ 1307ms | ✓ 1060ms | http |
| 45.140.147.82:1082 | ✓ 540ms | ✓ 1426ms | ✓ 1067ms | ✓ 1965ms | ✓ 1120ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 913ms | ✓ 1749ms | ✓ 1097ms | http |
| 37.187.109.70:10111 | ✓ 1386ms | 否 | ✓ 1835ms | 否 | ✓ 1768ms | http |
| 38.188.247.12:999 | ✓ 1912ms | 否 | ✓ 373ms | ✓ 1429ms | 否 | http |
| 101.32.243.189:80 | 否 | 否 | ✓ 1634ms | ✓ 1537ms | ✓ 1586ms | http |
| 45.125.67.37:8443 | ✓ 1303ms | 否 | ✓ 1359ms | ✓ 1268ms | ✓ 1936ms | http |
| 154.90.48.209:9090 | ✓ 1933ms | 否 | ✓ 1755ms | ✓ 1393ms | ✓ 1089ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1769ms | ✓ 1973ms | ✓ 1424ms | http |
| 91.242.229.129:8092 | ✓ 871ms | 否 | ✓ 1769ms | ✓ 1833ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1150ms | 否 | ✓ 1571ms | ✓ 1338ms | ✓ 1211ms | http |
| 59.46.216.131:30001 | ✓ 1055ms | 否 | ✓ 1456ms | 否 | ✓ 1251ms | http |
| 3.101.133.120:80 | ✓ 752ms | ✓ 1270ms | ✓ 884ms | ✓ 1137ms | ✓ 1187ms | http |
| 45.140.147.82:1081 | ✓ 1354ms | 否 | ✓ 1436ms | ✓ 1847ms | 否 | http |
| 45.63.88.46:1080 | 否 | ✓ 1676ms | 否 | ✓ 1283ms | ✓ 1798ms | http |
| 89.208.106.138:10808 | ✓ 1395ms | 否 | ✓ 1763ms | ✓ 1372ms | 否 | http |
| 31.56.48.253:26133 | ✓ 1231ms | ✓ 1565ms | ✓ 848ms | ✓ 1699ms | ✓ 1403ms | http |
| 152.70.91.193:40000 | 否 | 否 | ✓ 1761ms | ✓ 1919ms | ✓ 1817ms | http |
| 45.186.6.104:3128 | ✓ 1794ms | ✓ 1628ms | ✓ 1698ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 748ms | ✓ 1874ms | ✓ 612ms | ✓ 1403ms | ✓ 1530ms | http |
| 38.7.18.188:999 | ✓ 424ms | ✓ 1324ms | ✓ 1627ms | 否 | 否 | http |
| 190.12.150.244:999 | ✓ 1841ms | ✓ 1885ms | ✓ 1275ms | 否 | 否 | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1825ms | ✓ 1869ms | ✓ 1794ms | http |
| 61.52.131.172:8443 | ✓ 1023ms | ✓ 1344ms | ✓ 1204ms | ✓ 1367ms | ✓ 1128ms | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1513ms | ✓ 1530ms | ✓ 1541ms | http |
| 80.92.204.47:1081 | ✓ 544ms | 否 | ✓ 967ms | ✓ 1810ms | 否 | http |
| 139.159.97.82:10900 | ✓ 1866ms | 否 | 否 | ✓ 1618ms | ✓ 1497ms | http |

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
