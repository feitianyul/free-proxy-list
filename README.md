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

最后更新：2026-04-29 00:54:07 UTC（2026-04-29 08:54:07 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 824ms | 否 | ✓ 1049ms | ✓ 1048ms | ✓ 1101ms | http |
| 15.204.230.28:3128 | ✓ 1388ms | 否 | ✓ 1794ms | ✓ 1922ms | ✓ 1323ms | http |
| 113.160.132.26:8080 | ✓ 1995ms | ✓ 1359ms | ✓ 1327ms | ✓ 1493ms | ✓ 1584ms | http |
| 45.167.124.71:999 | 否 | ✓ 1933ms | ✓ 1588ms | 否 | ✓ 1868ms | http |
| 86.104.72.220:1081 | ✓ 490ms | ✓ 1245ms | ✓ 570ms | ✓ 1388ms | ✓ 1084ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 858ms | ✓ 1407ms | ✓ 1994ms | http |
| 115.231.181.40:8128 | ✓ 1442ms | 否 | ✓ 1099ms | ✓ 1316ms | 否 | http |
| 120.92.212.16:7890 | ✓ 868ms | 否 | ✓ 1094ms | ✓ 1490ms | ✓ 1982ms | http |
| 46.101.95.183:8888 | ✓ 1744ms | ✓ 1762ms | ✓ 1048ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1050ms | ✓ 1561ms | ✓ 1242ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1019ms | ✓ 1360ms | ✓ 1088ms | 否 | 否 | http |
| 8.211.166.184:8081 | ✓ 643ms | ✓ 1044ms | ✓ 818ms | ✓ 837ms | ✓ 670ms | http |
| 154.64.232.35:8080 | ✓ 612ms | ✓ 1999ms | ✓ 1181ms | 否 | 否 | http |
| 47.85.51.197:1080 | 否 | ✓ 1213ms | ✓ 883ms | 否 | ✓ 1835ms | http |
| 77.110.116.224:3128 | ✓ 1573ms | 否 | ✓ 1318ms | ✓ 1986ms | ✓ 1863ms | http |
| 34.71.229.255:3128 | ✓ 1935ms | 否 | ✓ 1706ms | 否 | ✓ 1782ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1184ms | ✓ 1052ms | 否 | ✓ 1372ms | http |
| 172.236.145.31:7890 | ✓ 749ms | 否 | ✓ 749ms | ✓ 1461ms | 否 | http |
| 110.89.51.18:443 | 否 | 否 | ✓ 1927ms | ✓ 1963ms | ✓ 1282ms | http |
| 168.144.75.9:3128 | ✓ 1766ms | 否 | 否 | ✓ 1973ms | ✓ 1723ms | http |
| 120.92.108.86:7890 | ✓ 1296ms | 否 | ✓ 1300ms | ✓ 1908ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1612ms | 否 | ✓ 1420ms | 否 | ✓ 1512ms | http |
| 117.2.224.43:30106 | 否 | 否 | ✓ 1722ms | ✓ 1418ms | ✓ 1564ms | http |
| 121.230.8.41:1080 | ✓ 981ms | ✓ 1566ms | ✓ 1022ms | ✓ 1523ms | ✓ 1032ms | http |
| 152.70.91.193:40000 | ✓ 1273ms | 否 | 否 | ✓ 1401ms | ✓ 1125ms | http |
| 38.180.2.107:3128 | ✓ 984ms | ✓ 1867ms | ✓ 1866ms | 否 | 否 | http |
| 143.198.223.214:1084 | ✓ 848ms | 否 | ✓ 1106ms | ✓ 1024ms | ✓ 856ms | http |
| 45.88.0.115:3128 | ✓ 1017ms | ✓ 1520ms | ✓ 1706ms | 否 | ✓ 1831ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1602ms | ✓ 1958ms | ✓ 1665ms | http |
| 45.88.0.117:3128 | ✓ 1375ms | ✓ 1568ms | ✓ 1772ms | ✓ 1815ms | ✓ 1532ms | http |
| 45.140.147.82:1081 | ✓ 1309ms | ✓ 1510ms | ✓ 1249ms | 否 | ✓ 1451ms | http |
| 121.130.177.28:8888 | ✓ 941ms | 否 | ✓ 1747ms | ✓ 1656ms | ✓ 1165ms | http |
| 183.232.248.73:7890 | ✓ 1355ms | ✓ 1410ms | ✓ 1397ms | ✓ 1902ms | ✓ 1108ms | http |
| 86.104.72.220:1082 | ✓ 412ms | 否 | ✓ 1078ms | ✓ 1235ms | ✓ 945ms | http |
| 77.110.119.136:3128 | ✓ 1148ms | 否 | ✓ 603ms | ✓ 1758ms | ✓ 1161ms | http |
| 80.92.204.47:1081 | ✓ 658ms | ✓ 1697ms | ✓ 995ms | ✓ 1869ms | ✓ 1343ms | http |
| 150.107.140.238:3128 | ✓ 1678ms | 否 | ✓ 1945ms | ✓ 1425ms | 否 | http |
| 213.220.62.62:3128 | ✓ 637ms | ✓ 1427ms | ✓ 640ms | ✓ 1538ms | ✓ 1128ms | http |
| 213.220.62.63:3128 | ✓ 638ms | ✓ 1428ms | ✓ 640ms | ✓ 1481ms | ✓ 1150ms | http |
| 45.88.0.113:3128 | ✓ 638ms | ✓ 1511ms | ✓ 594ms | ✓ 1554ms | ✓ 1178ms | http |
| 45.88.0.111:3128 | ✓ 620ms | ✓ 1529ms | ✓ 636ms | ✓ 1615ms | ✓ 1151ms | http |
| 45.88.0.99:3128 | ✓ 608ms | ✓ 1533ms | ✓ 637ms | ✓ 1593ms | ✓ 1213ms | http |
| 45.88.0.98:3128 | ✓ 604ms | ✓ 1539ms | ✓ 662ms | ✓ 1556ms | ✓ 1214ms | http |
| 45.88.0.116:3128 | ✓ 639ms | ✓ 1551ms | ✓ 600ms | ✓ 1604ms | ✓ 1192ms | http |
| 45.88.0.114:3128 | ✓ 608ms | ✓ 1716ms | ✓ 601ms | ✓ 1575ms | ✓ 1169ms | http |
| 118.113.245.61:1080 | ✓ 1280ms | ✓ 1721ms | ✓ 1314ms | ✓ 1814ms | ✓ 1365ms | http |
| 118.113.246.204:1080 | ✓ 1232ms | ✓ 1731ms | ✓ 1279ms | ✓ 1683ms | ✓ 1279ms | http |
| 217.182.195.221:30000 | ✓ 1407ms | 否 | ✓ 953ms | 否 | ✓ 1871ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1472ms | ✓ 1472ms | ✓ 1412ms | http |
| 112.209.22.22:8083 | 否 | 否 | ✓ 1497ms | ✓ 1236ms | ✓ 1814ms | http |
| 130.61.174.200:1080 | 否 | ✓ 1659ms | ✓ 986ms | ✓ 1971ms | 否 | http |
| 139.162.153.201:3128 | ✓ 704ms | ✓ 1927ms | ✓ 622ms | 否 | ✓ 1597ms | http |
| 89.208.106.138:10808 | ✓ 1138ms | ✓ 1905ms | 否 | ✓ 1537ms | 否 | http |
| 64.188.67.154:1080 | 否 | ✓ 1868ms | ✓ 1763ms | 否 | ✓ 1476ms | http |
| 34.101.184.164:3128 | ✓ 1988ms | 否 | ✓ 1688ms | 否 | ✓ 1858ms | http |
| 129.150.53.35:3128 | ✓ 1006ms | 否 | ✓ 1252ms | ✓ 1076ms | ✓ 865ms | http |
| 43.133.44.89:8888 | 否 | ✓ 1673ms | 否 | ✓ 1907ms | ✓ 1729ms | http |
| 61.52.131.172:8443 | ✓ 835ms | ✓ 1158ms | ✓ 924ms | ✓ 1187ms | ✓ 926ms | http |
| 103.171.255.59:8080 | ✓ 1831ms | 否 | 否 | ✓ 1695ms | ✓ 1562ms | http |
| 91.107.181.137:3128 | ✓ 1364ms | 否 | ✓ 1994ms | 否 | ✓ 1796ms | http |
| 121.230.8.136:1080 | ✓ 828ms | ✓ 1352ms | ✓ 1121ms | 否 | ✓ 900ms | http |
| 103.39.51.207:8080 | ✓ 1287ms | 否 | 否 | ✓ 1488ms | ✓ 1577ms | http |

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
