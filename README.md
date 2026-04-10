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

最后更新：2026-04-10 18:45:49 UTC（2026-04-11 02:45:49 UTC+8）

**代理总数：61**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 185.191.236.162:3128 | ✓ 881ms | ✓ 1571ms | ✓ 1923ms | ✓ 1467ms | ✓ 960ms | http |
| 218.108.131.186:17890 | ✓ 1006ms | ✓ 1305ms | ✓ 1090ms | ✓ 1307ms | ✓ 1039ms | http |
| 147.161.210.140:8800 | ✓ 853ms | 否 | ✓ 1278ms | ✓ 1347ms | ✓ 1171ms | http |
| 1.231.81.166:3128 | ✓ 944ms | ✓ 1142ms | 否 | ✓ 1410ms | ✓ 1228ms | http |
| 202.141.161.53:10808 | ✓ 1117ms | ✓ 1491ms | ✓ 1306ms | ✓ 1373ms | ✓ 1287ms | http |
| 159.223.225.118:8888 | ✓ 785ms | 否 | 否 | ✓ 1495ms | ✓ 1197ms | http |
| 113.160.132.26:8080 | ✓ 1591ms | ✓ 1888ms | ✓ 1467ms | ✓ 1508ms | ✓ 1179ms | http |
| 152.32.132.190:7890 | ✓ 1091ms | 否 | ✓ 1392ms | ✓ 1652ms | 否 | http |
| 167.103.115.102:8800 | ✓ 1961ms | 否 | ✓ 1487ms | ✓ 1258ms | ✓ 1457ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1852ms | ✓ 1706ms | ✓ 1391ms | http |
| 35.225.22.61:80 | 否 | ✓ 1367ms | ✓ 850ms | 否 | ✓ 894ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1452ms | ✓ 1623ms | ✓ 1268ms | http |
| 167.103.144.127:8800 | ✓ 1953ms | 否 | ✓ 1374ms | ✓ 1698ms | ✓ 1810ms | http |
| 167.103.31.122:8800 | ✓ 1906ms | 否 | ✓ 1606ms | ✓ 1882ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1127ms | 否 | ✓ 1296ms | 否 | ✓ 1230ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1517ms | ✓ 487ms | ✓ 1494ms | ✓ 1276ms | http |
| 91.238.105.64:2024 | ✓ 1284ms | ✓ 1482ms | ✓ 1537ms | 否 | 否 | http |
| 91.238.105.43:2023 | ✓ 1285ms | ✓ 1720ms | 否 | ✓ 1896ms | ✓ 1852ms | http |
| 85.239.59.252:7890 | 否 | ✓ 1498ms | ✓ 690ms | 否 | ✓ 1236ms | http |
| 45.167.125.21:999 | ✓ 687ms | ✓ 1507ms | 否 | ✓ 1843ms | ✓ 1583ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1454ms | ✓ 1115ms | 否 | ✓ 1680ms | http |
| 130.61.30.221:8080 | ✓ 503ms | 否 | ✓ 1267ms | ✓ 1846ms | 否 | http |
| 155.117.18.36:25388 | ✓ 961ms | ✓ 1037ms | ✓ 1491ms | ✓ 1527ms | ✓ 872ms | http |
| 77.93.89.128:47146 | ✓ 1710ms | 否 | ✓ 1898ms | ✓ 1604ms | ✓ 1293ms | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 1818ms | ✓ 1889ms | ✓ 1492ms | http |
| 147.161.239.240:8800 | ✓ 450ms | ✓ 1529ms | ✓ 1307ms | ✓ 1839ms | ✓ 1343ms | http |
| 101.43.127.100:8877 | ✓ 1051ms | ✓ 1253ms | ✓ 977ms | ✓ 1379ms | ✓ 1128ms | http |
| 47.105.98.23:3128 | ✓ 1043ms | ✓ 1379ms | ✓ 1137ms | 否 | ✓ 1132ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1407ms | ✓ 1440ms | 否 | ✓ 1175ms | http |
| 115.204.166.236:40000 | ✓ 1683ms | ✓ 1829ms | ✓ 1189ms | 否 | ✓ 1861ms | http |
| 35.194.4.51:3128 | ✓ 180ms | ✓ 1442ms | ✓ 1056ms | ✓ 985ms | ✓ 969ms | http |
| 201.144.25.226:3128 | 否 | 否 | ✓ 1405ms | ✓ 1683ms | ✓ 940ms | http |
| 177.234.217.88:999 | ✓ 1915ms | 否 | ✓ 1831ms | ✓ 1767ms | ✓ 1689ms | http |
| 130.162.53.123:50960 | ✓ 426ms | ✓ 1269ms | ✓ 361ms | 否 | ✓ 1238ms | http |
| 121.230.9.125:1080 | ✓ 1265ms | ✓ 1620ms | ✓ 1189ms | ✓ 1534ms | ✓ 1252ms | http |
| 212.58.132.5:8888 | ✓ 1098ms | 否 | ✓ 1460ms | ✓ 1480ms | ✓ 1225ms | http |
| 45.236.129.64:3128 | ✓ 1827ms | 否 | ✓ 1705ms | 否 | ✓ 1818ms | http |
| 8.219.97.248:80 | ✓ 1551ms | 否 | ✓ 1297ms | ✓ 1461ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1597ms | ✓ 1220ms | ✓ 375ms | 否 | 否 | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1350ms | ✓ 1479ms | ✓ 1178ms | http |
| 113.176.92.71:3128 | ✓ 1857ms | ✓ 1666ms | ✓ 1481ms | ✓ 1390ms | ✓ 1544ms | http |
| 103.235.67.190:80 | ✓ 1104ms | 否 | ✓ 1467ms | ✓ 1434ms | ✓ 1200ms | http |
| 121.230.8.111:1080 | ✓ 1288ms | 否 | ✓ 1302ms | 否 | ✓ 1405ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1442ms | ✓ 1188ms | ✓ 1402ms | 否 | http |
| 66.231.71.76:999 | ✓ 1620ms | ✓ 1613ms | 否 | ✓ 1879ms | ✓ 1612ms | http |
| 210.223.44.230:3128 | ✓ 889ms | ✓ 1283ms | ✓ 1122ms | ✓ 1309ms | ✓ 932ms | http |
| 170.106.137.214:7890 | ✓ 1604ms | ✓ 1790ms | ✓ 655ms | ✓ 895ms | ✓ 916ms | http |
| 222.228.171.92:8080 | ✓ 974ms | ✓ 1957ms | ✓ 1426ms | ✓ 1169ms | ✓ 852ms | http |
| 104.168.93.120:8080 | ✓ 1276ms | 否 | ✓ 1046ms | ✓ 1987ms | ✓ 1161ms | http |
| 168.110.52.228:3128 | ✓ 1792ms | 否 | 否 | ✓ 1323ms | ✓ 1039ms | http |
| 157.66.2.100:1111 | 否 | 否 | ✓ 1853ms | ✓ 1632ms | ✓ 1549ms | http |
| 65.108.203.35:18080 | ✓ 705ms | ✓ 1746ms | ✓ 935ms | 否 | 否 | http |
| 43.99.86.50:30219 | ✓ 1845ms | ✓ 1168ms | ✓ 832ms | ✓ 1058ms | ✓ 846ms | http |
| 217.77.102.18:3128 | ✓ 889ms | ✓ 1941ms | 否 | ✓ 1816ms | ✓ 1247ms | http |
| 150.241.106.173:8080 | ✓ 1490ms | 否 | ✓ 745ms | ✓ 1977ms | 否 | http |
| 61.76.95.217:40088 | 否 | ✓ 1889ms | ✓ 1511ms | ✓ 1979ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1725ms | ✓ 1757ms | ✓ 1798ms | 否 | 否 | http |
| 104.129.202.127:12354 | ✓ 969ms | ✓ 1015ms | ✓ 1064ms | ✓ 1408ms | ✓ 856ms | http |
| 104.129.202.127:10810 | ✓ 1536ms | ✓ 1067ms | ✓ 969ms | ✓ 1133ms | ✓ 824ms | http |
| 103.39.51.207:8080 | ✓ 1678ms | 否 | 否 | ✓ 1890ms | ✓ 1932ms | http |
| 46.101.126.84:8888 | ✓ 415ms | ✓ 1754ms | ✓ 1842ms | ✓ 1558ms | ✓ 1372ms | http |

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
