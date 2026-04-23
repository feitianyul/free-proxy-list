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

最后更新：2026-04-23 11:15:45 UTC（2026-04-23 19:15:45 UTC+8）

**代理总数：66**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 74ms | ✓ 959ms | 否 | ✓ 991ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1217ms | ✓ 1202ms | ✓ 1160ms | ✓ 1126ms | ✓ 932ms | http |
| 106.10.55.212:1121 | ✓ 1198ms | ✓ 1760ms | ✓ 1582ms | ✓ 1515ms | 否 | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1510ms | ✓ 1412ms | ✓ 1209ms | http |
| 46.101.95.183:8888 | ✓ 856ms | 否 | ✓ 1240ms | ✓ 1848ms | 否 | http |
| 152.42.208.139:8118 | ✓ 979ms | 否 | ✓ 1370ms | ✓ 1965ms | ✓ 1058ms | http |
| 115.231.181.40:8128 | ✓ 1347ms | ✓ 1245ms | ✓ 1912ms | ✓ 1469ms | 否 | http |
| 45.153.231.229:8080 | ✓ 1239ms | ✓ 1937ms | ✓ 1342ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1953ms | ✓ 1642ms | ✓ 1386ms | ✓ 1558ms | ✓ 1313ms | http |
| 177.93.132.244:3128 | ✓ 645ms | 否 | ✓ 641ms | 否 | ✓ 1677ms | http |
| 113.160.132.26:8080 | ✓ 1078ms | ✓ 1648ms | ✓ 1081ms | ✓ 1524ms | ✓ 1203ms | http |
| 218.108.131.186:17890 | ✓ 873ms | ✓ 1029ms | ✓ 820ms | ✓ 1094ms | ✓ 879ms | http |
| 91.99.15.45:2095 | ✓ 549ms | ✓ 1442ms | ✓ 1531ms | ✓ 1775ms | 否 | http |
| 93.185.156.89:3128 | 否 | 否 | ✓ 729ms | ✓ 1901ms | ✓ 1553ms | http |
| 35.225.22.61:80 | ✓ 378ms | 否 | ✓ 562ms | 否 | ✓ 1041ms | http |
| 168.110.52.228:3128 | ✓ 864ms | 否 | 否 | ✓ 1147ms | ✓ 1010ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1035ms | ✓ 1375ms | ✓ 1122ms | http |
| 212.58.132.5:8888 | ✓ 1095ms | 否 | ✓ 1474ms | ✓ 1496ms | ✓ 1201ms | http |
| 89.208.106.138:10808 | ✓ 441ms | 否 | ✓ 1989ms | 否 | ✓ 1175ms | http |
| 34.71.229.255:3128 | ✓ 346ms | ✓ 1451ms | ✓ 790ms | ✓ 1186ms | 否 | http |
| 130.61.174.200:1080 | 否 | 否 | ✓ 645ms | ✓ 1467ms | ✓ 928ms | http |
| 152.70.91.193:40000 | ✓ 1573ms | 否 | 否 | ✓ 1710ms | ✓ 1215ms | http |
| 135.125.97.184:38833 | ✓ 1134ms | ✓ 1679ms | ✓ 1272ms | 否 | 否 | http |
| 160.250.4.245:1 | ✓ 1944ms | 否 | ✓ 1251ms | ✓ 1451ms | ✓ 1233ms | http |
| 94.131.123.32:3128 | ✓ 771ms | ✓ 1716ms | 否 | ✓ 1730ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1345ms | ✓ 1339ms | ✓ 1348ms | 否 | 否 | http |
| 84.47.150.126:1080 | ✓ 1501ms | 否 | ✓ 1884ms | 否 | ✓ 1609ms | http |
| 168.144.75.9:3128 | ✓ 1599ms | 否 | ✓ 1869ms | ✓ 1818ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1572ms | 否 | ✓ 1491ms | 否 | ✓ 1627ms | http |
| 217.182.195.221:30000 | ✓ 1599ms | 否 | ✓ 1173ms | 否 | ✓ 1738ms | http |
| 14.184.71.31:5107 | ✓ 1950ms | 否 | ✓ 1491ms | ✓ 1810ms | ✓ 1965ms | http |
| 120.92.108.86:7890 | ✓ 1785ms | 否 | ✓ 1873ms | 否 | ✓ 1568ms | http |
| 202.129.206.239:3128 | ✓ 1456ms | 否 | ✓ 1642ms | ✓ 1832ms | 否 | http |
| 168.222.254.136:8888 | ✓ 1166ms | ✓ 1675ms | ✓ 1481ms | 否 | 否 | http |
| 85.190.99.143:443 | ✓ 499ms | 否 | ✓ 1355ms | 否 | ✓ 1947ms | http |
| 8.219.195.129:1080 | ✓ 914ms | 否 | ✓ 1079ms | ✓ 1287ms | ✓ 1258ms | http |
| 103.229.126.221:7890 | ✓ 1038ms | ✓ 1831ms | ✓ 1249ms | ✓ 1130ms | ✓ 864ms | http |
| 1.234.153.14:80 | ✓ 946ms | ✓ 1778ms | ✓ 1031ms | ✓ 1228ms | ✓ 897ms | http |
| 8.137.112.117:3128 | ✓ 1072ms | ✓ 1538ms | ✓ 1036ms | ✓ 1323ms | ✓ 1148ms | http |
| 194.147.90.23:3128 | ✓ 964ms | 否 | ✓ 1629ms | 否 | ✓ 1804ms | http |
| 103.56.112.120:7890 | ✓ 1346ms | 否 | ✓ 1362ms | ✓ 1730ms | ✓ 1324ms | http |
| 145.239.239.82:3128 | ✓ 1073ms | 否 | ✓ 1385ms | 否 | ✓ 1798ms | http |
| 120.92.212.16:7890 | ✓ 1294ms | 否 | ✓ 1863ms | ✓ 1458ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1149ms | 否 | 否 | ✓ 1630ms | ✓ 1420ms | http |
| 210.45.76.58:42992 | ✓ 1205ms | ✓ 1648ms | ✓ 1554ms | ✓ 1697ms | ✓ 1208ms | http |
| 64.188.77.26:3128 | 否 | ✓ 1616ms | 否 | ✓ 1483ms | ✓ 1486ms | http |
| 62.113.119.14:8080 | ✓ 980ms | ✓ 1483ms | ✓ 702ms | ✓ 1435ms | ✓ 1059ms | http |
| 45.140.147.82:1081 | ✓ 1350ms | ✓ 1360ms | ✓ 1738ms | ✓ 1745ms | ✓ 1608ms | http |
| 45.140.147.82:1082 | ✓ 550ms | ✓ 1222ms | ✓ 1661ms | ✓ 1498ms | ✓ 1071ms | http |
| 45.76.207.177:40000 | ✓ 1166ms | 否 | ✓ 1244ms | ✓ 1449ms | ✓ 1299ms | http |
| 45.140.147.155:1082 | ✓ 1673ms | 否 | 否 | ✓ 1312ms | ✓ 1876ms | http |
| 183.232.248.73:7890 | ✓ 1155ms | ✓ 1339ms | ✓ 1434ms | ✓ 1293ms | 否 | http |
| 213.32.85.26:3128 | 否 | ✓ 1631ms | ✓ 1431ms | 否 | ✓ 1649ms | http |
| 45.129.141.143:3128 | ✓ 878ms | 否 | ✓ 1449ms | ✓ 1967ms | ✓ 1838ms | http |
| 37.187.109.70:10111 | ✓ 1243ms | 否 | ✓ 1965ms | 否 | ✓ 1862ms | http |
| 47.84.73.61:1080 | ✓ 1064ms | ✓ 1897ms | ✓ 963ms | ✓ 1285ms | ✓ 1041ms | http |
| 117.236.124.166:3128 | ✓ 1615ms | 否 | ✓ 1189ms | ✓ 1858ms | 否 | http |
| 200.174.198.32:8888 | ✓ 1287ms | 否 | ✓ 1434ms | 否 | ✓ 1828ms | http |
| 20.127.128.70:8080 | ✓ 1937ms | 否 | ✓ 1815ms | 否 | ✓ 1948ms | http |
| 223.84.151.86:30005 | ✓ 1568ms | ✓ 1650ms | ✓ 1579ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 1194ms | 否 | ✓ 1388ms | 否 | ✓ 1770ms | http |
| 218.77.106.10:10150 | ✓ 1670ms | ✓ 1902ms | ✓ 1883ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 864ms | ✓ 1188ms | ✓ 983ms | ✓ 1307ms | ✓ 1069ms | http |
| 58.69.114.117:5050 | ✓ 1545ms | 否 | 否 | ✓ 1804ms | ✓ 1553ms | http |
| 172.208.25.199:3128 | ✓ 1450ms | ✓ 1394ms | 否 | ✓ 1801ms | ✓ 1852ms | http |
| 160.250.5.22:1 | ✓ 1770ms | 否 | 否 | ✓ 1459ms | ✓ 1186ms | http |

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
