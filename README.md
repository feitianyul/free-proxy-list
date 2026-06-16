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

最后更新：2026-06-16 16:45:59 UTC（2026-06-17 00:45:59 UTC+8）

**代理总数：48**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.43.46.91:443 | 否 | ✓ 1403ms | ✓ 344ms | 否 | ✓ 1191ms | http |
| 64.188.125.187:3128 | ✓ 1577ms | 否 | ✓ 1620ms | 否 | ✓ 1718ms | http |
| 154.201.77.250:11021 | ✓ 1537ms | 否 | 否 | ✓ 1984ms | ✓ 1480ms | http |
| 34.43.46.91:80 | ✓ 1420ms | 否 | 否 | ✓ 1424ms | ✓ 1190ms | http |
| 82.97.247.37:80 | ✓ 725ms | 否 | ✓ 1304ms | 否 | ✓ 1709ms | http |
| 95.3.69.222:8080 | ✓ 1186ms | 否 | ✓ 1069ms | ✓ 1716ms | ✓ 1555ms | http |
| 113.160.132.26:8080 | ✓ 1536ms | ✓ 1711ms | ✓ 1589ms | ✓ 1218ms | ✓ 998ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1721ms | ✓ 1582ms | ✓ 1721ms | http |
| 8.215.25.3:2080 | 否 | 否 | ✓ 891ms | ✓ 1235ms | ✓ 953ms | http |
| 169.212.15.161:5000 | ✓ 940ms | ✓ 1063ms | ✓ 778ms | ✓ 1272ms | ✓ 1194ms | http |
| 91.107.182.124:82 | ✓ 1005ms | 否 | ✓ 1894ms | 否 | ✓ 1910ms | http |
| 51.250.4.195:2080 | ✓ 1039ms | ✓ 1839ms | ✓ 1883ms | 否 | 否 | http |
| 159.198.35.187:1080 | ✓ 631ms | 否 | ✓ 613ms | 否 | ✓ 846ms | http |
| 81.200.154.236:48503 | ✓ 488ms | 否 | ✓ 796ms | ✓ 1553ms | 否 | http |
| 185.200.188.234:10001 | ✓ 1098ms | 否 | ✓ 1056ms | 否 | ✓ 1572ms | http |
| 185.11.134.227:8443 | ✓ 960ms | 否 | ✓ 1927ms | 否 | ✓ 1640ms | http |
| 117.236.124.166:3128 | ✓ 1510ms | 否 | ✓ 1134ms | ✓ 1539ms | ✓ 1482ms | http |
| 8.154.21.175:3128 | ✓ 1224ms | ✓ 1169ms | ✓ 1689ms | 否 | ✓ 1715ms | http |
| 104.154.186.48:80 | 否 | ✓ 1097ms | ✓ 556ms | ✓ 849ms | ✓ 632ms | http |
| 36.147.78.166:80 | ✓ 1819ms | 否 | 否 | ✓ 1995ms | ✓ 1759ms | http |
| 217.154.155.115:8080 | ✓ 1340ms | 否 | ✓ 1218ms | 否 | ✓ 1132ms | http |
| 85.234.100.149:8080 | ✓ 1937ms | 否 | ✓ 586ms | ✓ 1747ms | ✓ 1302ms | http |
| 3.137.86.220:443 | ✓ 222ms | 否 | ✓ 1687ms | ✓ 1921ms | ✓ 1330ms | http |
| 85.234.100.149:1080 | ✓ 928ms | 否 | ✓ 633ms | ✓ 1792ms | ✓ 1177ms | http |
| 84.47.150.125:1080 | ✓ 1276ms | 否 | ✓ 1287ms | 否 | ✓ 1563ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1589ms | ✓ 1884ms | ✓ 1819ms | http |
| 199.127.62.89:3129 | ✓ 1937ms | 否 | ✓ 1621ms | 否 | ✓ 1682ms | http |
| 103.157.117.226:81 | ✓ 1746ms | 否 | 否 | ✓ 1776ms | ✓ 1569ms | http |
| 43.130.109.66:18088 | ✓ 847ms | 否 | ✓ 919ms | 否 | ✓ 975ms | http |
| 3.137.86.220:1080 | ✓ 860ms | 否 | ✓ 1959ms | ✓ 966ms | 否 | http |
| 18.180.59.181:80 | ✓ 787ms | ✓ 1341ms | ✓ 1246ms | 否 | ✓ 1509ms | http |
| 80.150.246.98:443 | ✓ 720ms | ✓ 1685ms | ✓ 1693ms | 否 | ✓ 1697ms | http |
| 103.157.200.126:3128 | ✓ 1475ms | 否 | ✓ 1162ms | 否 | ✓ 1561ms | http |
| 1.231.81.166:3128 | ✓ 1692ms | ✓ 1184ms | ✓ 1258ms | ✓ 1472ms | ✓ 915ms | http |
| 202.28.194.139:31280 | ✓ 1734ms | 否 | ✓ 1862ms | ✓ 1929ms | ✓ 1941ms | http |
| 8.219.97.248:80 | ✓ 1390ms | 否 | ✓ 1148ms | ✓ 1407ms | ✓ 1363ms | http |
| 144.202.14.153:50000 | ✓ 389ms | ✓ 1139ms | ✓ 251ms | ✓ 1748ms | ✓ 987ms | http |
| 115.231.181.40:8128 | ✓ 959ms | ✓ 1179ms | ✓ 1067ms | 否 | 否 | http |
| 36.147.78.166:443 | 否 | ✓ 1770ms | ✓ 1825ms | 否 | ✓ 1704ms | http |
| 45.67.223.123:2001 | ✓ 1929ms | ✓ 1654ms | 否 | 否 | ✓ 1787ms | http |
| 185.141.26.131:3128 | ✓ 1181ms | 否 | ✓ 1534ms | 否 | ✓ 1835ms | http |
| 167.86.95.198:3128 | ✓ 1405ms | ✓ 1825ms | ✓ 1644ms | ✓ 1833ms | 否 | http |
| 47.85.51.197:1080 | 否 | ✓ 1035ms | ✓ 307ms | ✓ 993ms | ✓ 827ms | http |
| 138.124.114.42:7443 | ✓ 1454ms | 否 | ✓ 1590ms | 否 | ✓ 1369ms | http |
| 89.169.53.40:7443 | ✓ 1747ms | ✓ 1971ms | ✓ 1704ms | 否 | 否 | http |
| 150.241.116.167:443 | ✓ 1343ms | 否 | ✓ 1870ms | ✓ 1987ms | 否 | http |
| 16.52.81.236:31351 | ✓ 1482ms | 否 | ✓ 1970ms | 否 | ✓ 1401ms | http |
| 103.28.37.131:3128 | ✓ 1646ms | 否 | ✓ 1919ms | ✓ 1367ms | 否 | http |

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
