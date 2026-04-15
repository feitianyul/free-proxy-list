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

最后更新：2026-04-15 14:39:50 UTC（2026-04-15 22:39:50 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 624ms | ✓ 1490ms | ✓ 921ms | ✓ 1531ms | ✓ 1351ms | http |
| 147.161.210.140:8800 | ✓ 859ms | 否 | ✓ 1140ms | ✓ 1446ms | ✓ 1125ms | http |
| 113.160.132.26:8080 | ✓ 1956ms | 否 | ✓ 1367ms | ✓ 1408ms | ✓ 1131ms | http |
| 167.103.115.102:8800 | ✓ 1890ms | 否 | ✓ 1371ms | ✓ 1546ms | ✓ 1465ms | http |
| 1.231.81.166:3128 | ✓ 1225ms | 否 | ✓ 1314ms | ✓ 1754ms | ✓ 1379ms | http |
| 167.103.34.108:8800 | ✓ 1950ms | 否 | ✓ 1600ms | 否 | ✓ 1691ms | http |
| 103.113.70.189:1082 | ✓ 309ms | ✓ 1124ms | ✓ 742ms | ✓ 927ms | ✓ 758ms | http |
| 167.103.31.122:8800 | ✓ 1968ms | 否 | ✓ 1886ms | 否 | ✓ 1593ms | http |
| 45.167.125.21:999 | ✓ 1137ms | 否 | ✓ 1525ms | ✓ 1861ms | ✓ 1487ms | http |
| 35.225.22.61:80 | ✓ 240ms | 否 | ✓ 347ms | ✓ 1126ms | ✓ 1022ms | http |
| 167.103.144.127:8800 | ✓ 1213ms | 否 | ✓ 971ms | ✓ 1511ms | ✓ 1339ms | http |
| 137.59.47.73:3128 | ✓ 1649ms | 否 | ✓ 1885ms | 否 | ✓ 1325ms | http |
| 130.61.30.221:8080 | ✓ 816ms | 否 | 否 | ✓ 1881ms | ✓ 1585ms | http |
| 20.127.128.70:8080 | ✓ 522ms | 否 | ✓ 1053ms | ✓ 1270ms | ✓ 1450ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1531ms | ✓ 1429ms | ✓ 1848ms | http |
| 122.248.45.54:8080 | 否 | 否 | ✓ 1859ms | ✓ 1601ms | ✓ 1727ms | http |
| 20.27.13.35:8561 | ✓ 810ms | 否 | ✓ 899ms | ✓ 1109ms | ✓ 822ms | http |
| 138.124.99.216:8888 | ✓ 1640ms | ✓ 1929ms | ✓ 1624ms | ✓ 1966ms | ✓ 1879ms | http |
| 20.78.26.206:8561 | ✓ 1881ms | ✓ 1759ms | ✓ 834ms | ✓ 1124ms | ✓ 856ms | http |
| 20.78.118.91:8561 | ✓ 1884ms | ✓ 1751ms | ✓ 863ms | ✓ 1070ms | ✓ 921ms | http |
| 20.27.11.248:8561 | ✓ 736ms | 否 | ✓ 748ms | ✓ 1072ms | ✓ 927ms | http |
| 20.27.15.111:8561 | ✓ 729ms | 否 | ✓ 764ms | ✓ 1091ms | ✓ 921ms | http |
| 103.113.70.189:1081 | ✓ 476ms | 否 | ✓ 49ms | ✓ 1206ms | ✓ 916ms | http |
| 144.31.27.49:1080 | ✓ 523ms | 否 | ✓ 1321ms | 否 | ✓ 1715ms | http |
| 85.239.59.252:7890 | ✓ 712ms | 否 | ✓ 757ms | 否 | ✓ 1837ms | http |
| 195.26.224.49:3128 | ✓ 651ms | 否 | ✓ 864ms | ✓ 1931ms | ✓ 1432ms | http |
| 150.107.140.238:3128 | ✓ 944ms | 否 | ✓ 1999ms | ✓ 1344ms | 否 | http |
| 45.149.92.147:5001 | ✓ 936ms | 否 | ✓ 1310ms | ✓ 1390ms | ✓ 1065ms | http |
| 94.131.118.129:1081 | ✓ 978ms | ✓ 1731ms | ✓ 551ms | ✓ 1689ms | ✓ 1104ms | http |
| 190.12.150.244:999 | ✓ 876ms | 否 | ✓ 1134ms | ✓ 1593ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1171ms | 否 | ✓ 1149ms | ✓ 1762ms | 否 | http |
| 181.78.240.121:999 | ✓ 1916ms | 否 | ✓ 1586ms | ✓ 1796ms | ✓ 1480ms | http |
| 160.238.65.3:3128 | ✓ 1618ms | 否 | ✓ 1619ms | 否 | ✓ 1964ms | http |
| 160.238.65.9:3128 | 否 | ✓ 1340ms | ✓ 1896ms | 否 | ✓ 998ms | http |
| 36.141.21.200:7890 | 否 | ✓ 1827ms | ✓ 1698ms | ✓ 1882ms | 否 | http |
| 218.153.163.186:8692 | ✓ 1129ms | 否 | ✓ 1996ms | ✓ 1235ms | ✓ 966ms | http |
| 101.32.243.189:80 | ✓ 1531ms | 否 | ✓ 1810ms | 否 | ✓ 1549ms | http |
| 213.131.85.28:1981 | ✓ 1551ms | 否 | ✓ 1905ms | ✓ 1873ms | ✓ 1874ms | http |
| 34.101.184.164:3128 | ✓ 1748ms | 否 | ✓ 1344ms | ✓ 1724ms | ✓ 1210ms | http |
| 116.80.49.174:3172 | ✓ 1765ms | 否 | ✓ 1710ms | 否 | ✓ 1905ms | http |
| 185.132.178.178:1080 | ✓ 516ms | 否 | 否 | ✓ 1653ms | ✓ 1424ms | http |
| 107.173.42.121:7890 | 否 | ✓ 954ms | ✓ 109ms | ✓ 966ms | 否 | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1420ms | ✓ 1463ms | ✓ 1329ms | http |
| 20.210.76.104:8561 | ✓ 1658ms | 否 | ✓ 705ms | ✓ 1198ms | ✓ 951ms | http |
| 5.255.123.43:1080 | ✓ 1041ms | 否 | ✓ 1058ms | ✓ 1985ms | ✓ 1402ms | http |
| 59.46.216.131:30001 | ✓ 1189ms | 否 | ✓ 1607ms | 否 | ✓ 1258ms | http |
| 20.210.39.153:8561 | 否 | 否 | ✓ 1039ms | ✓ 1068ms | ✓ 934ms | http |
| 38.180.2.107:3128 | ✓ 681ms | ✓ 1624ms | ✓ 1701ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1049ms | ✓ 1492ms | ✓ 1221ms | 否 | 否 | http |
| 218.153.163.156:8623 | ✓ 1892ms | 否 | ✓ 1276ms | 否 | ✓ 1531ms | http |
| 147.45.60.34:1082 | ✓ 281ms | ✓ 1353ms | ✓ 1299ms | 否 | ✓ 1248ms | http |
| 152.32.132.190:7890 | ✓ 1513ms | ✓ 1397ms | 否 | ✓ 1138ms | 否 | http |
| 45.12.151.226:2829 | ✓ 1616ms | 否 | ✓ 1184ms | ✓ 1830ms | 否 | http |
| 181.78.44.63:999 | ✓ 996ms | ✓ 1692ms | ✓ 1421ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1162ms | ✓ 1444ms | ✓ 1120ms | ✓ 1469ms | ✓ 1120ms | http |
| 168.222.254.26:8888 | ✓ 675ms | ✓ 1571ms | ✓ 1300ms | ✓ 1967ms | 否 | http |
| 45.140.147.82:1082 | ✓ 413ms | ✓ 1050ms | 否 | ✓ 1578ms | ✓ 1094ms | http |
| 212.58.132.5:8888 | 否 | ✓ 1988ms | ✓ 1725ms | ✓ 1453ms | ✓ 1178ms | http |
| 202.141.161.53:10808 | ✓ 1248ms | ✓ 1855ms | 否 | ✓ 1371ms | 否 | http |

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
