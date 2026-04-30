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

最后更新：2026-04-30 13:08:05 UTC（2026-04-30 21:08:05 UTC+8）

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
| 47.85.51.197:1080 | ✓ 658ms | 否 | ✓ 511ms | ✓ 974ms | 否 | http |
| 218.108.131.186:17890 | ✓ 964ms | ✓ 1757ms | ✓ 1007ms | ✓ 1299ms | ✓ 1422ms | http |
| 46.101.95.183:8888 | ✓ 1215ms | 否 | ✓ 910ms | ✓ 1727ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1839ms | ✓ 1687ms | ✓ 1226ms | 否 | ✓ 1180ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1293ms | ✓ 1589ms | ✓ 1257ms | http |
| 45.167.124.71:999 | ✓ 1690ms | 否 | ✓ 1369ms | ✓ 1709ms | ✓ 1523ms | http |
| 62.113.119.14:8080 | ✓ 771ms | 否 | ✓ 723ms | 否 | ✓ 1952ms | http |
| 159.223.225.118:8888 | ✓ 655ms | 否 | ✓ 1321ms | 否 | ✓ 1466ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1701ms | 否 | ✓ 1657ms | ✓ 1781ms | http |
| 115.231.181.40:8128 | ✓ 1771ms | 否 | 否 | ✓ 1287ms | ✓ 1082ms | http |
| 65.109.213.99:1080 | ✓ 1029ms | 否 | ✓ 538ms | 否 | ✓ 1694ms | http |
| 103.3.246.71:3128 | 否 | 否 | ✓ 1479ms | ✓ 1607ms | ✓ 1243ms | http |
| 77.110.107.80:8080 | ✓ 1296ms | 否 | ✓ 619ms | ✓ 1701ms | 否 | http |
| 86.104.72.220:1081 | ✓ 1006ms | 否 | ✓ 962ms | 否 | ✓ 1269ms | http |
| 107.173.160.222:1080 | ✓ 498ms | 否 | ✓ 1448ms | ✓ 1261ms | ✓ 828ms | http |
| 94.72.109.214:8888 | ✓ 1327ms | 否 | ✓ 742ms | ✓ 1228ms | ✓ 890ms | http |
| 8.154.21.175:3128 | ✓ 944ms | ✓ 1727ms | ✓ 974ms | ✓ 1248ms | ✓ 1042ms | http |
| 103.157.200.126:3128 | ✓ 1915ms | 否 | ✓ 1363ms | 否 | ✓ 1899ms | http |
| 185.21.11.140:1080 | ✓ 574ms | 否 | ✓ 810ms | ✓ 1508ms | ✓ 1128ms | http |
| 38.92.10.98:20058 | 否 | ✓ 1146ms | ✓ 411ms | ✓ 1362ms | ✓ 764ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1736ms | ✓ 1899ms | ✓ 1436ms | http |
| 77.110.116.224:3128 | ✓ 1155ms | 否 | 否 | ✓ 1603ms | ✓ 1326ms | http |
| 47.95.231.180:8084 | ✓ 1634ms | ✓ 1444ms | ✓ 955ms | ✓ 1335ms | ✓ 1056ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1966ms | ✓ 1612ms | ✓ 1475ms | ✓ 922ms | http |
| 185.234.66.87:1081 | 否 | 否 | ✓ 1351ms | ✓ 1802ms | ✓ 1404ms | http |
| 77.110.119.136:3128 | ✓ 937ms | 否 | ✓ 233ms | ✓ 1681ms | ✓ 952ms | http |
| 185.234.66.87:1082 | ✓ 1146ms | 否 | 否 | ✓ 1993ms | ✓ 1373ms | http |
| 210.223.44.230:3128 | ✓ 1760ms | 否 | ✓ 1972ms | ✓ 1853ms | ✓ 1193ms | http |
| 154.64.232.35:8080 | ✓ 861ms | ✓ 1049ms | 否 | 否 | ✓ 779ms | http |
| 152.70.91.193:40000 | 否 | 否 | ✓ 1971ms | ✓ 1355ms | ✓ 1556ms | http |
| 118.113.247.40:1080 | 否 | ✓ 1766ms | 否 | ✓ 1787ms | ✓ 1382ms | http |
| 38.92.10.152:57579 | ✓ 1637ms | ✓ 1085ms | ✓ 586ms | ✓ 1021ms | ✓ 696ms | http |
| 206.206.126.177:2412 | ✓ 1610ms | 否 | ✓ 940ms | ✓ 1232ms | ✓ 919ms | http |
| 130.61.174.200:1080 | ✓ 1268ms | 否 | 否 | ✓ 1527ms | ✓ 1278ms | http |
| 77.110.107.80:1080 | ✓ 1266ms | 否 | ✓ 1816ms | ✓ 1715ms | ✓ 1571ms | http |
| 120.92.212.16:7890 | ✓ 1175ms | ✓ 1317ms | 否 | ✓ 1980ms | ✓ 1113ms | http |
| 109.199.125.231:3128 | ✓ 526ms | ✓ 1614ms | 否 | 否 | ✓ 1373ms | http |
| 34.96.238.40:8080 | ✓ 1450ms | ✓ 1379ms | 否 | ✓ 1336ms | 否 | http |
| 203.150.128.136:8080 | ✓ 1908ms | 否 | ✓ 1922ms | ✓ 1800ms | ✓ 1831ms | http |
| 60.249.94.208:3128 | ✓ 1352ms | ✓ 1510ms | ✓ 992ms | ✓ 1359ms | ✓ 1083ms | http |
| 81.26.190.143:1080 | ✓ 753ms | 否 | ✓ 1036ms | 否 | ✓ 1714ms | http |
| 8.219.97.248:80 | ✓ 1905ms | 否 | ✓ 1474ms | ✓ 1948ms | 否 | http |
| 183.232.248.73:7890 | ✓ 963ms | ✓ 1265ms | ✓ 1075ms | ✓ 1226ms | ✓ 1010ms | http |
| 183.238.3.150:7897 | ✓ 963ms | ✓ 1268ms | ✓ 1101ms | ✓ 1278ms | ✓ 989ms | http |
| 59.46.216.131:30001 | ✓ 1623ms | 否 | ✓ 1247ms | 否 | ✓ 1163ms | http |
| 3.101.133.120:80 | ✓ 574ms | ✓ 1225ms | ✓ 1104ms | ✓ 1216ms | ✓ 956ms | http |
| 80.92.204.47:1081 | ✓ 1144ms | ✓ 1995ms | ✓ 460ms | ✓ 1503ms | ✓ 1110ms | http |
| 103.70.114.149:3128 | ✓ 1604ms | 否 | ✓ 1589ms | ✓ 1910ms | ✓ 1708ms | http |
| 5.255.126.157:10001 | ✓ 1038ms | ✓ 1650ms | 否 | 否 | ✓ 1728ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 831ms | ✓ 1388ms | ✓ 1373ms | http |
| 152.32.132.190:7890 | ✓ 1802ms | ✓ 1809ms | 否 | ✓ 1032ms | ✓ 1608ms | http |
| 168.222.254.136:8888 | ✓ 841ms | ✓ 1850ms | ✓ 1374ms | 否 | 否 | http |
| 89.208.106.138:10808 | ✓ 1022ms | 否 | ✓ 1185ms | 否 | ✓ 1085ms | http |
| 91.217.81.131:1080 | ✓ 1869ms | 否 | ✓ 1680ms | 否 | ✓ 1593ms | http |
| 34.101.184.164:3128 | ✓ 990ms | 否 | ✓ 899ms | ✓ 1950ms | ✓ 1100ms | http |
| 185.234.64.63:1081 | 否 | 否 | ✓ 1175ms | ✓ 1223ms | ✓ 895ms | http |
| 103.20.63.232:3128 | ✓ 898ms | ✓ 1901ms | ✓ 881ms | ✓ 1096ms | ✓ 873ms | http |
| 61.52.131.172:8443 | ✓ 989ms | ✓ 1293ms | ✓ 1139ms | ✓ 1358ms | ✓ 1101ms | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1577ms | ✓ 1545ms | ✓ 1678ms | http |

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
