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

最后更新：2026-04-16 12:49:51 UTC（2026-04-16 20:49:51 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 632ms | 否 | ✓ 789ms | ✓ 1621ms | ✓ 1817ms | http |
| 147.161.210.140:8800 | ✓ 1688ms | ✓ 1890ms | ✓ 1051ms | ✓ 1162ms | ✓ 1259ms | http |
| 188.246.224.49:7890 | ✓ 636ms | ✓ 1614ms | 否 | 否 | ✓ 1722ms | http |
| 202.141.161.53:10808 | ✓ 1126ms | ✓ 1434ms | ✓ 1266ms | 否 | ✓ 1101ms | http |
| 167.103.115.102:8800 | ✓ 1221ms | 否 | ✓ 1389ms | ✓ 1470ms | ✓ 1412ms | http |
| 113.160.132.26:8080 | ✓ 1752ms | 否 | ✓ 1628ms | 否 | ✓ 1497ms | http |
| 168.144.75.9:3128 | ✓ 1306ms | 否 | 否 | ✓ 1918ms | ✓ 1799ms | http |
| 167.103.34.108:8800 | ✓ 1801ms | 否 | ✓ 1572ms | 否 | ✓ 1769ms | http |
| 1.231.81.166:3128 | ✓ 1759ms | ✓ 1718ms | 否 | ✓ 1713ms | ✓ 1582ms | http |
| 103.113.70.189:1081 | ✓ 287ms | ✓ 1623ms | ✓ 61ms | ✓ 1676ms | ✓ 916ms | http |
| 34.96.238.40:8080 | ✓ 1145ms | 否 | 否 | ✓ 1360ms | ✓ 1498ms | http |
| 84.47.150.125:1080 | ✓ 1233ms | 否 | ✓ 1810ms | 否 | ✓ 1590ms | http |
| 167.103.31.122:8800 | ✓ 1861ms | 否 | ✓ 1905ms | ✓ 1846ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1589ms | ✓ 1595ms | ✓ 834ms | 否 | ✓ 1088ms | http |
| 167.103.144.127:8800 | ✓ 991ms | 否 | ✓ 1517ms | ✓ 1735ms | ✓ 1678ms | http |
| 190.12.150.244:999 | ✓ 1532ms | 否 | 否 | ✓ 1770ms | ✓ 1967ms | http |
| 45.12.151.226:2829 | ✓ 1429ms | 否 | ✓ 1806ms | 否 | ✓ 1669ms | http |
| 78.11.96.22:8888 | ✓ 787ms | 否 | ✓ 1086ms | ✓ 1460ms | ✓ 1363ms | http |
| 35.225.22.61:80 | ✓ 741ms | 否 | ✓ 275ms | ✓ 1199ms | ✓ 943ms | http |
| 212.58.132.5:8888 | ✓ 976ms | 否 | ✓ 972ms | ✓ 1349ms | ✓ 1049ms | http |
| 45.167.125.21:999 | ✓ 810ms | 否 | ✓ 1220ms | ✓ 1872ms | ✓ 1568ms | http |
| 166.249.54.61:7234 | ✓ 953ms | ✓ 1402ms | 否 | ✓ 1725ms | ✓ 1435ms | http |
| 5.104.87.17:8051 | ✓ 1881ms | 否 | ✓ 1590ms | ✓ 1910ms | ✓ 1540ms | http |
| 34.101.184.164:3128 | ✓ 1690ms | 否 | ✓ 1557ms | ✓ 1747ms | 否 | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1368ms | ✓ 1449ms | ✓ 1115ms | http |
| 107.172.102.234:40621 | ✓ 757ms | ✓ 978ms | ✓ 1020ms | ✓ 1963ms | 否 | http |
| 84.47.150.126:1080 | ✓ 765ms | 否 | ✓ 1284ms | 否 | ✓ 1700ms | http |
| 138.124.99.216:8888 | 否 | ✓ 1687ms | 否 | ✓ 1939ms | ✓ 1470ms | http |
| 85.239.59.252:7890 | ✓ 604ms | 否 | ✓ 500ms | 否 | ✓ 1102ms | http |
| 148.135.119.12:3128 | ✓ 755ms | ✓ 1723ms | ✓ 1494ms | 否 | 否 | http |
| 150.249.255.91:3128 | 否 | ✓ 1735ms | 否 | ✓ 1138ms | ✓ 864ms | http |
| 20.127.128.70:8080 | ✓ 1275ms | 否 | ✓ 1208ms | ✓ 1265ms | ✓ 796ms | http |
| 152.32.132.190:7890 | ✓ 872ms | 否 | 否 | ✓ 1103ms | ✓ 928ms | http |
| 158.160.85.248:3128 | ✓ 705ms | 否 | ✓ 762ms | 否 | ✓ 1741ms | http |
| 157.230.178.216:8088 | ✓ 720ms | 否 | 否 | ✓ 1908ms | ✓ 1533ms | http |
| 185.114.73.2:1080 | ✓ 490ms | 否 | ✓ 1183ms | ✓ 1785ms | ✓ 1860ms | http |
| 59.46.216.131:30001 | ✓ 1111ms | 否 | ✓ 1120ms | ✓ 1415ms | 否 | http |
| 47.74.226.8:5001 | 否 | ✓ 1631ms | ✓ 1213ms | ✓ 1808ms | 否 | http |
| 43.132.188.134:443 | ✓ 1627ms | 否 | ✓ 1992ms | ✓ 1566ms | 否 | http |
| 34.71.229.255:3128 | ✓ 957ms | 否 | ✓ 822ms | ✓ 1083ms | ✓ 919ms | http |
| 120.92.212.16:7890 | ✓ 1087ms | ✓ 1296ms | ✓ 1769ms | ✓ 1341ms | ✓ 1118ms | http |
| 12.89.176.82:3128 | ✓ 295ms | 否 | 否 | ✓ 1242ms | ✓ 801ms | http |
| 101.32.243.189:80 | ✓ 1339ms | 否 | ✓ 1564ms | ✓ 1952ms | ✓ 1662ms | http |
| 103.138.70.165:3129 | ✓ 1771ms | 否 | ✓ 1417ms | ✓ 1547ms | ✓ 1538ms | http |
| 36.141.21.200:7890 | ✓ 1252ms | ✓ 1385ms | ✓ 1069ms | ✓ 1438ms | ✓ 1185ms | http |
| 120.92.212.16:8890 | ✓ 1244ms | ✓ 1283ms | ✓ 1124ms | ✓ 1601ms | ✓ 1062ms | http |
| 140.238.242.189:8100 | 否 | 否 | ✓ 1462ms | ✓ 1871ms | ✓ 1239ms | http |
| 147.45.214.210:1080 | ✓ 581ms | 否 | ✓ 1653ms | ✓ 1911ms | 否 | http |
| 57.128.188.167:9174 | ✓ 1839ms | 否 | ✓ 1566ms | ✓ 1962ms | ✓ 1864ms | http |
| 218.108.131.186:17890 | ✓ 822ms | ✓ 1007ms | ✓ 887ms | ✓ 1126ms | ✓ 891ms | http |
| 114.118.82.146:80 | 否 | ✓ 1262ms | ✓ 1206ms | ✓ 1384ms | ✓ 1054ms | http |
| 45.140.147.155:1082 | ✓ 944ms | 否 | ✓ 1371ms | ✓ 1911ms | 否 | http |
| 45.140.147.155:1081 | 否 | 否 | ✓ 390ms | ✓ 1240ms | ✓ 1085ms | http |
| 147.45.186.28:3128 | 否 | 否 | ✓ 484ms | ✓ 1916ms | ✓ 1864ms | http |
| 62.113.119.14:8080 | ✓ 604ms | ✓ 1511ms | ✓ 1190ms | ✓ 1392ms | ✓ 1162ms | http |
| 8.219.195.129:1080 | ✓ 1019ms | ✓ 1984ms | ✓ 894ms | ✓ 1293ms | ✓ 990ms | http |
| 217.77.102.18:3128 | ✓ 1051ms | ✓ 1828ms | ✓ 1897ms | 否 | ✓ 1802ms | http |
| 42.101.8.101:8888 | ✓ 1276ms | ✓ 1538ms | 否 | ✓ 1872ms | 否 | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1296ms | ✓ 1723ms | ✓ 1552ms | http |
| 61.52.131.172:8443 | ✓ 1505ms | ✓ 1172ms | ✓ 990ms | ✓ 1319ms | ✓ 1613ms | http |
| 104.248.243.244:3128 | ✓ 1619ms | 否 | ✓ 1428ms | ✓ 1322ms | 否 | http |
| 18.170.25.193:54009 | ✓ 1652ms | 否 | 否 | ✓ 1807ms | ✓ 1522ms | http |
| 177.93.132.244:3128 | ✓ 1432ms | 否 | ✓ 1815ms | 否 | ✓ 1826ms | http |
| 113.176.92.71:3128 | ✓ 1893ms | ✓ 1561ms | ✓ 1798ms | ✓ 1528ms | ✓ 1396ms | http |

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
