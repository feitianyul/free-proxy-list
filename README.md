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

最后更新：2026-04-25 20:42:48 UTC（2026-04-26 04:42:48 UTC+8）

**代理总数：49**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1862ms | ✓ 1312ms | ✓ 1294ms | ✓ 1136ms | ✓ 1013ms | http |
| 47.85.51.197:1080 | ✓ 345ms | ✓ 960ms | ✓ 458ms | 否 | 否 | http |
| 80.92.204.47:1081 | ✓ 846ms | ✓ 1070ms | ✓ 524ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 1665ms | ✓ 1754ms | ✓ 1342ms | ✓ 1254ms | ✓ 1012ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1745ms | ✓ 1909ms | ✓ 1506ms | http |
| 59.46.216.131:30001 | ✓ 1099ms | ✓ 1580ms | ✓ 1280ms | ✓ 1583ms | ✓ 1252ms | http |
| 217.76.245.80:999 | ✓ 862ms | ✓ 1533ms | ✓ 963ms | ✓ 1414ms | ✓ 1308ms | http |
| 82.148.18.242:443 | ✓ 783ms | ✓ 1787ms | ✓ 1547ms | 否 | 否 | http |
| 68.183.199.89:1080 | ✓ 1238ms | 否 | ✓ 1085ms | 否 | ✓ 1744ms | http |
| 84.47.150.125:1080 | ✓ 1537ms | ✓ 1474ms | ✓ 1969ms | ✓ 1727ms | 否 | http |
| 8.209.238.110:47701 | ✓ 652ms | ✓ 1404ms | ✓ 873ms | ✓ 1095ms | ✓ 867ms | http |
| 86.104.74.110:1081 | ✓ 1083ms | ✓ 1290ms | 否 | 否 | ✓ 1943ms | http |
| 62.113.119.14:8080 | ✓ 1429ms | 否 | ✓ 1441ms | ✓ 1453ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1606ms | 否 | ✓ 1511ms | ✓ 1295ms | http |
| 45.129.141.143:3128 | ✓ 1204ms | ✓ 1728ms | ✓ 1731ms | ✓ 1845ms | ✓ 1435ms | http |
| 120.92.212.16:7890 | ✓ 1080ms | ✓ 1583ms | ✓ 1314ms | ✓ 1387ms | ✓ 1109ms | http |
| 47.105.98.23:3128 | ✓ 1450ms | ✓ 1650ms | 否 | 否 | ✓ 1927ms | http |
| 200.125.171.254:999 | ✓ 1121ms | ✓ 1706ms | ✓ 1291ms | ✓ 1554ms | ✓ 1121ms | http |
| 211.95.152.50:45046 | 否 | ✓ 1801ms | ✓ 1612ms | ✓ 1978ms | ✓ 1325ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1366ms | ✓ 1168ms | 否 | ✓ 1716ms | http |
| 147.75.202.36:10063 | ✓ 767ms | ✓ 1123ms | ✓ 1174ms | ✓ 1230ms | ✓ 1200ms | http |
| 168.110.52.228:3128 | ✓ 739ms | ✓ 1422ms | ✓ 1240ms | ✓ 1214ms | ✓ 866ms | http |
| 177.93.132.244:3128 | ✓ 1962ms | 否 | ✓ 1905ms | 否 | ✓ 1969ms | http |
| 2.27.54.161:1080 | ✓ 726ms | 否 | ✓ 1543ms | ✓ 1664ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1027ms | ✓ 1265ms | ✓ 1033ms | ✓ 1260ms | ✓ 1046ms | http |
| 62.60.216.109:3128 | ✓ 858ms | ✓ 1606ms | ✓ 1303ms | ✓ 1651ms | ✓ 1363ms | http |
| 61.171.66.158:3128 | ✓ 988ms | ✓ 1404ms | ✓ 998ms | ✓ 1409ms | ✓ 1063ms | http |
| 8.137.112.117:3128 | ✓ 1493ms | ✓ 1578ms | ✓ 1319ms | ✓ 1519ms | ✓ 1312ms | http |
| 36.141.21.200:7890 | 否 | 否 | ✓ 1201ms | ✓ 1450ms | ✓ 1175ms | http |
| 49.151.187.67:8082 | ✓ 1574ms | 否 | ✓ 1484ms | ✓ 1871ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1210ms | ✓ 1821ms | ✓ 1134ms | ✓ 1419ms | ✓ 1467ms | http |
| 108.181.0.167:8080 | 否 | ✓ 1461ms | ✓ 512ms | ✓ 949ms | ✓ 1477ms | http |
| 121.43.196.213:8222 | ✓ 1132ms | ✓ 1289ms | ✓ 998ms | ✓ 1371ms | ✓ 1080ms | http |
| 121.43.196.210:8222 | ✓ 1125ms | ✓ 1291ms | ✓ 1025ms | ✓ 1304ms | ✓ 1142ms | http |
| 114.55.226.123:10086 | ✓ 1288ms | ✓ 1659ms | ✓ 1393ms | ✓ 1704ms | ✓ 1260ms | http |
| 8.219.195.129:1080 | ✓ 1346ms | ✓ 1844ms | ✓ 1123ms | ✓ 1349ms | ✓ 1026ms | http |
| 159.89.31.62:8080 | ✓ 474ms | ✓ 1559ms | ✓ 926ms | 否 | ✓ 1595ms | http |
| 194.31.87.77:3128 | ✓ 1230ms | 否 | ✓ 1557ms | 否 | ✓ 1897ms | http |
| 183.232.248.73:7890 | ✓ 1476ms | ✓ 1829ms | 否 | 否 | ✓ 1433ms | http |
| 23.143.160.193:999 | 否 | ✓ 1496ms | ✓ 1140ms | 否 | ✓ 1146ms | http |
| 187.102.195.55:999 | 否 | 否 | ✓ 1560ms | ✓ 1847ms | ✓ 1383ms | http |
| 58.63.109.230:10817 | ✓ 968ms | ✓ 1325ms | ✓ 988ms | ✓ 1355ms | ✓ 984ms | http |
| 94.131.106.231:1081 | ✓ 907ms | 否 | ✓ 1061ms | ✓ 1831ms | 否 | http |
| 105.159.136.255:4472 | ✓ 906ms | ✓ 1717ms | ✓ 1982ms | 否 | 否 | http |
| 192.99.44.178:3128 | ✓ 999ms | ✓ 1182ms | ✓ 1195ms | ✓ 1599ms | 否 | http |
| 47.84.73.61:1080 | ✓ 1627ms | ✓ 1961ms | ✓ 1001ms | ✓ 1361ms | ✓ 1103ms | http |
| 45.186.6.104:3128 | ✓ 1513ms | ✓ 1557ms | ✓ 1683ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1097ms | ✓ 1419ms | ✓ 1178ms | ✓ 1411ms | ✓ 1145ms | http |
| 121.230.8.144:1080 | ✓ 1230ms | ✓ 1560ms | ✓ 1248ms | ✓ 1513ms | 否 | http |

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
