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

最后更新：2026-03-29 12:29:43 UTC（2026-03-29 20:29:43 UTC+8）

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
| 208.87.243.199:7878 | ✓ 612ms | ✓ 658ms | ✓ 895ms | ✓ 680ms | ✓ 523ms | http |
| 43.99.54.236:5555 | ✓ 615ms | ✓ 1961ms | ✓ 613ms | ✓ 788ms | ✓ 614ms | http |
| 147.161.210.140:8800 | ✓ 1197ms | 否 | ✓ 828ms | ✓ 1025ms | ✓ 988ms | http |
| 219.117.204.211:7799 | ✓ 1201ms | 否 | 否 | ✓ 832ms | ✓ 653ms | http |
| 113.160.132.26:8080 | ✓ 1395ms | ✓ 1591ms | ✓ 869ms | ✓ 1154ms | ✓ 968ms | http |
| 167.172.77.49:8080 | ✓ 1340ms | 否 | ✓ 1437ms | ✓ 1022ms | ✓ 815ms | http |
| 167.103.115.102:8800 | ✓ 1341ms | 否 | ✓ 985ms | ✓ 1149ms | ✓ 1148ms | http |
| 167.103.34.108:8800 | ✓ 1394ms | ✓ 1961ms | ✓ 1553ms | ✓ 1386ms | 否 | http |
| 95.213.217.168:52004 | ✓ 1005ms | 否 | ✓ 1837ms | ✓ 1818ms | ✓ 1307ms | http |
| 101.47.73.135:3128 | ✓ 1494ms | 否 | ✓ 1586ms | 否 | ✓ 1848ms | http |
| 45.144.232.5:11741 | ✓ 1180ms | 否 | ✓ 1298ms | 否 | ✓ 1940ms | http |
| 103.113.70.189:1081 | ✓ 445ms | 否 | ✓ 332ms | ✓ 1290ms | ✓ 895ms | http |
| 35.225.22.61:80 | ✓ 578ms | 否 | ✓ 1370ms | ✓ 1142ms | ✓ 1182ms | http |
| 167.103.144.127:8800 | ✓ 1725ms | 否 | ✓ 1329ms | ✓ 1605ms | ✓ 1447ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1421ms | ✓ 1414ms | ✓ 1161ms | http |
| 103.84.95.54:7890 | ✓ 1343ms | 否 | ✓ 684ms | 否 | ✓ 642ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1454ms | ✓ 901ms | ✓ 717ms | http |
| 206.81.27.105:3128 | ✓ 666ms | 否 | ✓ 1734ms | 否 | ✓ 1670ms | http |
| 167.103.31.122:8800 | ✓ 1709ms | 否 | ✓ 1328ms | ✓ 1565ms | ✓ 1489ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1174ms | ✓ 1002ms | 否 | ✓ 939ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1258ms | ✓ 890ms | ✓ 1556ms | 否 | http |
| 180.250.219.58:53281 | ✓ 1751ms | 否 | ✓ 1422ms | ✓ 1872ms | ✓ 1845ms | http |
| 1.231.81.166:3128 | ✓ 1483ms | ✓ 1195ms | ✓ 1191ms | 否 | 否 | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1320ms | ✓ 1507ms | ✓ 931ms | http |
| 128.199.114.189:9090 | ✓ 1186ms | 否 | ✓ 1659ms | ✓ 1836ms | ✓ 1218ms | http |
| 147.161.239.240:8800 | 否 | 否 | ✓ 1280ms | ✓ 1871ms | ✓ 1643ms | http |
| 101.43.127.100:8877 | ✓ 1392ms | 否 | ✓ 1066ms | 否 | ✓ 1892ms | http |
| 177.234.217.88:999 | ✓ 1811ms | 否 | ✓ 1907ms | ✓ 1965ms | ✓ 1740ms | http |
| 194.59.204.87:9080 | ✓ 627ms | ✓ 1655ms | ✓ 1318ms | 否 | 否 | http |
| 45.144.28.81:10808 | ✓ 682ms | ✓ 1581ms | ✓ 1426ms | 否 | 否 | http |
| 160.238.65.9:3128 | 否 | 否 | ✓ 1088ms | ✓ 1540ms | ✓ 1198ms | http |
| 181.41.201.85:3128 | ✓ 1029ms | 否 | ✓ 1537ms | 否 | ✓ 1897ms | http |
| 160.238.65.6:3128 | 否 | ✓ 1668ms | 否 | ✓ 1629ms | ✓ 1769ms | http |
| 143.244.140.119:3128 | ✓ 1561ms | 否 | 否 | ✓ 1742ms | ✓ 1725ms | http |
| 160.238.65.8:3128 | ✓ 696ms | 否 | ✓ 675ms | 否 | ✓ 1285ms | http |
| 115.231.181.40:8128 | ✓ 900ms | ✓ 1167ms | 否 | 否 | ✓ 1212ms | http |
| 113.255.59.226:8080 | ✓ 1410ms | 否 | ✓ 1379ms | 否 | ✓ 1069ms | http |
| 103.39.51.207:8080 | ✓ 1794ms | 否 | 否 | ✓ 1916ms | ✓ 1670ms | http |
| 64.227.76.27:1080 | ✓ 1180ms | 否 | ✓ 1188ms | 否 | ✓ 1756ms | http |
| 222.228.171.92:8080 | ✓ 1321ms | 否 | ✓ 1666ms | 否 | ✓ 1372ms | http |
| 38.145.220.11:8446 | ✓ 1188ms | 否 | ✓ 828ms | 否 | ✓ 1366ms | http |
| 45.12.151.226:2829 | ✓ 1353ms | 否 | ✓ 1500ms | 否 | ✓ 1683ms | http |
| 160.238.65.2:3128 | ✓ 714ms | ✓ 1545ms | ✓ 1632ms | 否 | ✓ 1663ms | http |
| 160.238.65.3:3128 | ✓ 689ms | ✓ 1652ms | ✓ 1724ms | 否 | 否 | http |
| 113.176.92.71:3128 | ✓ 1797ms | ✓ 1604ms | ✓ 1622ms | ✓ 1639ms | ✓ 1389ms | http |
| 38.145.218.161:8444 | ✓ 1692ms | ✓ 682ms | ✓ 307ms | ✓ 827ms | ✓ 814ms | http |
| 106.75.15.167:7890 | ✓ 1145ms | 否 | 否 | ✓ 1318ms | ✓ 890ms | http |
| 128.199.116.219:9090 | ✓ 1534ms | 否 | ✓ 1044ms | ✓ 1045ms | ✓ 915ms | http |
| 106.117.208.101:7890 | ✓ 888ms | ✓ 1207ms | 否 | ✓ 1417ms | 否 | http |
| 91.238.123.230:8000 | 否 | ✓ 1913ms | ✓ 1223ms | 否 | ✓ 1336ms | http |
| 91.238.123.111:8000 | ✓ 618ms | 否 | ✓ 1633ms | 否 | ✓ 1680ms | http |
| 34.101.184.164:3128 | ✓ 1810ms | 否 | ✓ 1764ms | 否 | ✓ 1471ms | http |
| 45.136.198.40:3128 | ✓ 1373ms | ✓ 1834ms | 否 | 否 | ✓ 1913ms | http |
| 103.39.51.190:8080 | ✓ 1873ms | 否 | 否 | ✓ 1619ms | ✓ 1344ms | http |
| 5.104.87.17:8051 | ✓ 1562ms | 否 | ✓ 1727ms | ✓ 1922ms | ✓ 1479ms | http |
| 160.238.65.5:3128 | ✓ 1234ms | 否 | ✓ 1751ms | 否 | ✓ 1665ms | http |
| 61.52.131.172:8443 | ✓ 906ms | ✓ 1093ms | 否 | ✓ 1160ms | ✓ 876ms | http |
| 38.145.208.172:8448 | ✓ 1230ms | 否 | ✓ 1623ms | ✓ 660ms | ✓ 513ms | http |
| 39.185.46.193:5911 | ✓ 642ms | ✓ 1170ms | ✓ 855ms | ✓ 827ms | ✓ 656ms | http |
| 8.219.97.248:80 | ✓ 1758ms | 否 | ✓ 1226ms | ✓ 1663ms | ✓ 1762ms | http |
| 152.69.229.220:3128 | ✓ 1683ms | ✓ 1063ms | ✓ 1657ms | ✓ 1004ms | ✓ 957ms | http |
| 8.137.112.117:3128 | 否 | ✓ 1793ms | 否 | ✓ 1429ms | ✓ 1513ms | http |

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
