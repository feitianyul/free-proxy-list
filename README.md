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

最后更新：2026-03-14 11:25:48 UTC（2026-03-14 19:25:48 UTC+8）

**代理总数：52**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 202.155.12.161:443 | 否 | 否 | ✓ 1078ms | ✓ 1360ms | ✓ 1170ms | http |
| 205.209.118.30:3138 | 否 | 否 | ✓ 754ms | ✓ 1252ms | ✓ 1847ms | http |
| 86.53.183.16:1080 | ✓ 987ms | 否 | ✓ 1791ms | ✓ 1723ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1860ms | 否 | 否 | ✓ 1474ms | ✓ 1384ms | http |
| 45.167.124.52:8080 | ✓ 1061ms | ✓ 1772ms | ✓ 1998ms | ✓ 1602ms | ✓ 1243ms | http |
| 165.227.5.10:8888 | ✓ 528ms | 否 | 否 | ✓ 1175ms | ✓ 901ms | http |
| 186.148.180.46:999 | ✓ 677ms | 否 | ✓ 1262ms | ✓ 1505ms | ✓ 1185ms | http |
| 216.180.127.45:1080 | ✓ 1610ms | 否 | ✓ 1475ms | ✓ 1928ms | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1853ms | ✓ 1298ms | ✓ 1162ms | ✓ 838ms | http |
| 101.43.255.96:80 | ✓ 1273ms | ✓ 1559ms | ✓ 1276ms | 否 | ✓ 1380ms | http |
| 162.243.149.86:31028 | ✓ 542ms | ✓ 1545ms | 否 | 否 | ✓ 1837ms | http |
| 101.43.127.100:8877 | ✓ 1114ms | 否 | ✓ 1154ms | ✓ 1670ms | ✓ 1976ms | http |
| 138.124.53.25:7443 | ✓ 563ms | 否 | ✓ 1691ms | 否 | ✓ 1531ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1084ms | ✓ 1235ms | ✓ 1200ms | http |
| 45.88.0.115:3128 | ✓ 1811ms | ✓ 1237ms | ✓ 1641ms | 否 | 否 | http |
| 45.88.0.114:3128 | ✓ 1813ms | ✓ 1778ms | ✓ 1585ms | 否 | 否 | http |
| 85.198.96.242:3128 | 否 | ✓ 1663ms | ✓ 1563ms | 否 | ✓ 1249ms | http |
| 81.70.169.194:80 | 否 | ✓ 1472ms | ✓ 1314ms | ✓ 1789ms | ✓ 1608ms | http |
| 157.100.54.4:80 | ✓ 919ms | 否 | ✓ 884ms | ✓ 1709ms | ✓ 1337ms | http |
| 45.186.6.104:3128 | ✓ 1223ms | ✓ 1627ms | ✓ 1556ms | 否 | 否 | http |
| 116.80.65.85:3172 | ✓ 1949ms | 否 | ✓ 1783ms | 否 | ✓ 1882ms | http |
| 62.60.177.204:34094 | ✓ 410ms | 否 | ✓ 816ms | ✓ 1118ms | ✓ 739ms | http |
| 103.113.70.189:1081 | ✓ 1274ms | ✓ 1479ms | 否 | ✓ 1238ms | ✓ 987ms | http |
| 150.249.255.91:3128 | ✓ 736ms | ✓ 1507ms | ✓ 732ms | ✓ 1852ms | 否 | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1146ms | ✓ 1450ms | ✓ 1144ms | http |
| 38.145.218.101:8447 | ✓ 1098ms | ✓ 1472ms | ✓ 846ms | ✓ 964ms | ✓ 747ms | http |
| 120.92.212.16:7890 | ✓ 1137ms | ✓ 1442ms | ✓ 1195ms | ✓ 1524ms | ✓ 1180ms | http |
| 139.159.99.242:8080 | ✓ 1035ms | ✓ 1228ms | ✓ 1297ms | 否 | 否 | http |
| 88.80.150.82:8080 | ✓ 1253ms | ✓ 1947ms | 否 | 否 | ✓ 1841ms | https |
| 45.88.0.113:3128 | ✓ 798ms | 否 | ✓ 1637ms | 否 | ✓ 1492ms | http |
| 45.136.131.39:8443 | 否 | 否 | ✓ 967ms | ✓ 1630ms | ✓ 1490ms | http |
| 38.145.208.201:8447 | ✓ 580ms | ✓ 902ms | ✓ 311ms | ✓ 933ms | ✓ 748ms | http |
| 59.46.216.131:30001 | ✓ 1669ms | 否 | ✓ 1372ms | 否 | ✓ 1342ms | http |
| 89.185.85.138:1080 | ✓ 705ms | 否 | ✓ 818ms | ✓ 1977ms | ✓ 1872ms | http |
| 120.238.159.234:22222 | 否 | 否 | ✓ 1416ms | ✓ 1379ms | ✓ 1097ms | http |
| 45.88.0.111:3128 | ✓ 1079ms | ✓ 1473ms | 否 | ✓ 1588ms | ✓ 1000ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1712ms | ✓ 1251ms | ✓ 1822ms | 否 | http |
| 213.220.62.62:3128 | ✓ 536ms | ✓ 1626ms | 否 | ✓ 1287ms | 否 | http |
| 45.136.198.40:3128 | ✓ 945ms | 否 | ✓ 1816ms | 否 | ✓ 1762ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1948ms | ✓ 1646ms | ✓ 1632ms | http |
| 57.128.188.167:9157 | ✓ 1879ms | 否 | ✓ 1947ms | 否 | ✓ 1949ms | http |
| 150.230.249.50:1080 | 否 | 否 | ✓ 1027ms | ✓ 1203ms | ✓ 994ms | http |
| 45.136.130.245:8447 | ✓ 428ms | ✓ 1017ms | ✓ 501ms | ✓ 977ms | ✓ 851ms | http |
| 119.18.145.49:20326 | ✓ 1768ms | 否 | ✓ 1591ms | 否 | ✓ 1734ms | http |
| 38.210.179.112:999 | ✓ 1350ms | 否 | ✓ 1525ms | ✓ 1841ms | 否 | http |
| 201.71.2.41:999 | ✓ 1296ms | 否 | 否 | ✓ 1827ms | ✓ 1624ms | http |
| 47.77.193.180:1080 | ✓ 1225ms | ✓ 1028ms | ✓ 557ms | ✓ 974ms | ✓ 688ms | http |
| 45.88.0.98:3128 | ✓ 938ms | 否 | ✓ 572ms | ✓ 1270ms | ✓ 958ms | http |
| 45.88.0.117:3128 | ✓ 476ms | ✓ 1173ms | ✓ 1026ms | ✓ 1302ms | ✓ 1032ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1517ms | ✓ 1837ms | ✓ 1849ms | http |
| 45.88.0.116:3128 | ✓ 729ms | 否 | ✓ 467ms | ✓ 1300ms | ✓ 976ms | http |
| 45.88.0.99:3128 | ✓ 959ms | 否 | ✓ 1564ms | ✓ 1864ms | 否 | http |

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
