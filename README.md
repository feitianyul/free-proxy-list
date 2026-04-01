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

最后更新：2026-04-01 11:57:39 UTC（2026-04-01 19:57:39 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | 否 | ✓ 1315ms | ✓ 976ms | ✓ 1065ms | ✓ 609ms | http |
| 147.161.210.140:8800 | ✓ 1426ms | 否 | ✓ 870ms | ✓ 1115ms | ✓ 983ms | http |
| 1.231.81.166:3128 | ✓ 1459ms | 否 | ✓ 1538ms | ✓ 1121ms | ✓ 1028ms | http |
| 42.96.16.158:1311 | ✓ 1603ms | 否 | ✓ 1234ms | ✓ 1287ms | ✓ 964ms | http |
| 133.242.138.34:8100 | ✓ 1450ms | ✓ 1504ms | 否 | ✓ 1360ms | ✓ 1518ms | http |
| 113.160.132.26:8080 | ✓ 1852ms | 否 | ✓ 934ms | ✓ 1502ms | ✓ 1100ms | http |
| 34.101.184.164:3128 | ✓ 1868ms | 否 | ✓ 1084ms | ✓ 1541ms | ✓ 994ms | http |
| 95.213.217.168:52004 | ✓ 1016ms | ✓ 1724ms | ✓ 1682ms | 否 | ✓ 1746ms | http |
| 167.103.115.102:8800 | ✓ 1616ms | 否 | ✓ 1078ms | ✓ 1952ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1658ms | 否 | ✓ 1585ms | ✓ 1511ms | ✓ 1864ms | http |
| 180.250.219.58:53281 | ✓ 1866ms | 否 | ✓ 1523ms | ✓ 1986ms | ✓ 1945ms | http |
| 35.225.22.61:80 | ✓ 1017ms | 否 | 否 | ✓ 1423ms | ✓ 1297ms | http |
| 167.103.144.127:8800 | ✓ 1185ms | 否 | ✓ 1279ms | ✓ 1632ms | ✓ 1358ms | http |
| 45.12.151.226:2829 | ✓ 933ms | 否 | ✓ 1505ms | 否 | ✓ 1607ms | http |
| 45.167.124.52:8080 | ✓ 640ms | ✓ 1608ms | ✓ 776ms | ✓ 1651ms | ✓ 1396ms | http |
| 91.233.223.147:3128 | ✓ 925ms | 否 | ✓ 970ms | 否 | ✓ 1651ms | http |
| 115.231.181.40:8128 | ✓ 1929ms | ✓ 1782ms | ✓ 974ms | ✓ 1151ms | ✓ 1145ms | http |
| 167.103.31.122:8800 | ✓ 1548ms | 否 | ✓ 1375ms | ✓ 1600ms | ✓ 1586ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1001ms | ✓ 889ms | ✓ 1141ms | ✓ 1127ms | http |
| 203.80.138.81:50000 | 否 | ✓ 1067ms | ✓ 908ms | ✓ 1010ms | ✓ 785ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1461ms | ✓ 994ms | ✓ 1390ms | 否 | http |
| 181.78.44.63:999 | ✓ 923ms | ✓ 1806ms | ✓ 1418ms | ✓ 1521ms | ✓ 1484ms | http |
| 177.234.217.88:999 | ✓ 1325ms | 否 | ✓ 1850ms | ✓ 1870ms | ✓ 1642ms | http |
| 8.219.97.248:80 | ✓ 1285ms | 否 | ✓ 1813ms | ✓ 1650ms | 否 | http |
| 147.161.239.240:8800 | ✓ 706ms | 否 | ✓ 1501ms | ✓ 1818ms | ✓ 1649ms | http |
| 45.129.141.143:3128 | ✓ 1594ms | ✓ 1967ms | ✓ 1810ms | 否 | ✓ 1843ms | http |
| 45.149.92.147:5001 | ✓ 619ms | 否 | ✓ 622ms | ✓ 779ms | ✓ 627ms | http |
| 167.160.191.204:6005 | ✓ 735ms | ✓ 1462ms | ✓ 1932ms | 否 | ✓ 1873ms | http |
| 45.136.198.40:3128 | ✓ 1391ms | ✓ 1859ms | 否 | 否 | ✓ 1885ms | http |
| 194.59.204.87:9080 | ✓ 654ms | ✓ 1719ms | ✓ 1510ms | ✓ 1978ms | ✓ 1353ms | http |
| 116.80.49.165:3172 | ✓ 1676ms | 否 | ✓ 1473ms | ✓ 1797ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1520ms | ✓ 1479ms | 否 | ✓ 1464ms | 否 | http |
| 195.19.217.200:3128 | ✓ 1272ms | 否 | ✓ 1691ms | 否 | ✓ 1683ms | http |
| 31.192.106.135:8010 | ✓ 1404ms | 否 | ✓ 1325ms | 否 | ✓ 1825ms | http |
| 59.46.216.131:30001 | ✓ 993ms | ✓ 1310ms | ✓ 1100ms | ✓ 1334ms | 否 | http |
| 116.80.65.75:3172 | ✓ 1694ms | 否 | ✓ 1480ms | ✓ 1784ms | 否 | http |
| 186.116.148.52:8080 | ✓ 1942ms | 否 | ✓ 1775ms | ✓ 1696ms | 否 | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1470ms | ✓ 1618ms | ✓ 1348ms | http |
| 47.74.226.8:5001 | ✓ 930ms | 否 | 否 | ✓ 1191ms | ✓ 1229ms | http |
| 72.11.150.178:6005 | 否 | 否 | ✓ 1014ms | ✓ 1220ms | ✓ 1318ms | http |
| 43.225.185.4:8000 | ✓ 974ms | 否 | ✓ 1645ms | ✓ 1309ms | ✓ 1317ms | http |
| 38.34.179.27:8452 | ✓ 1143ms | ✓ 659ms | ✓ 391ms | 否 | ✓ 1371ms | http |
| 38.34.179.95:8444 | ✓ 1632ms | ✓ 630ms | 否 | ✓ 1748ms | ✓ 786ms | http |
| 106.117.208.101:7890 | ✓ 961ms | 否 | ✓ 1548ms | 否 | ✓ 1986ms | http |
| 116.80.63.64:7777 | ✓ 1479ms | 否 | 否 | ✓ 1809ms | ✓ 1634ms | http |
| 116.80.96.108:3172 | ✓ 1522ms | 否 | 否 | ✓ 1821ms | ✓ 1631ms | http |
| 150.249.255.91:3128 | ✓ 502ms | ✓ 858ms | 否 | ✓ 1387ms | 否 | http |
| 24.144.86.173:1080 | ✓ 1349ms | 否 | 否 | ✓ 971ms | ✓ 733ms | http |
| 210.223.44.230:3128 | ✓ 1619ms | 否 | 否 | ✓ 1883ms | ✓ 1780ms | http |
| 62.113.119.14:8080 | ✓ 899ms | ✓ 1900ms | ✓ 1061ms | 否 | ✓ 1366ms | http |
| 160.250.5.22:1 | ✓ 1400ms | 否 | ✓ 1236ms | ✓ 1462ms | ✓ 1009ms | http |
| 160.250.4.245:1 | ✓ 1415ms | 否 | 否 | ✓ 1925ms | ✓ 1142ms | http |
| 46.101.190.71:3128 | ✓ 1298ms | 否 | ✓ 1932ms | 否 | ✓ 1462ms | http |
| 217.217.249.160:8080 | ✓ 1820ms | 否 | ✓ 1218ms | 否 | ✓ 1362ms | http |
| 218.60.0.214:80 | ✓ 1001ms | 否 | 否 | ✓ 1305ms | ✓ 1068ms | http |
| 107.174.80.186:3128 | ✓ 1794ms | 否 | ✓ 746ms | 否 | ✓ 786ms | http |
| 103.76.148.162:8181 | ✓ 1767ms | 否 | ✓ 1349ms | ✓ 1642ms | ✓ 1505ms | http |
| 198.59.68.130:3128 | ✓ 1822ms | ✓ 1488ms | ✓ 1663ms | ✓ 1753ms | 否 | http |
| 38.159.63.8:999 | 否 | ✓ 1628ms | ✓ 1760ms | 否 | ✓ 1522ms | http |
| 103.39.51.190:8080 | ✓ 1696ms | 否 | ✓ 1818ms | 否 | ✓ 1868ms | http |

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
