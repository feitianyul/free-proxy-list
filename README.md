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

最后更新：2026-04-26 14:06:28 UTC（2026-04-26 22:06:28 UTC+8）

**代理总数：46**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 46 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 46 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 16.163.173.29:1080 | ✓ 847ms | ✓ 1351ms | ✓ 1296ms | ✓ 1110ms | ✓ 853ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1531ms | 否 | ✓ 1496ms | ✓ 1903ms | http |
| 20.78.118.91:8561 | ✓ 1415ms | ✓ 1084ms | ✓ 691ms | ✓ 1152ms | ✓ 831ms | http |
| 20.27.15.111:8561 | ✓ 1395ms | ✓ 1154ms | ✓ 778ms | ✓ 1058ms | ✓ 830ms | http |
| 20.27.14.220:8561 | ✓ 1489ms | ✓ 1069ms | ✓ 848ms | ✓ 1031ms | ✓ 890ms | http |
| 20.27.11.248:8561 | ✓ 1442ms | ✓ 1829ms | ✓ 666ms | ✓ 1055ms | ✓ 854ms | http |
| 124.16.93.233:7890 | ✓ 1019ms | ✓ 1295ms | ✓ 1132ms | ✓ 1357ms | ✓ 1095ms | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1437ms | ✓ 1338ms | ✓ 1192ms | http |
| 211.95.152.50:45046 | 否 | ✓ 1761ms | ✓ 1762ms | ✓ 1790ms | 否 | http |
| 206.206.126.177:2412 | ✓ 986ms | 否 | ✓ 1177ms | ✓ 1201ms | ✓ 950ms | http |
| 177.93.132.244:3128 | ✓ 659ms | 否 | ✓ 1827ms | 否 | ✓ 1667ms | http |
| 20.210.39.153:8561 | ✓ 720ms | ✓ 1335ms | ✓ 690ms | 否 | 否 | http |
| 20.27.13.35:8561 | ✓ 1476ms | ✓ 1620ms | ✓ 699ms | ✓ 1040ms | ✓ 839ms | http |
| 20.78.26.206:8561 | ✓ 1466ms | ✓ 1755ms | ✓ 646ms | ✓ 1078ms | ✓ 784ms | http |
| 103.157.200.126:3128 | ✓ 1339ms | 否 | ✓ 1946ms | ✓ 1589ms | 否 | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1593ms | ✓ 1326ms | ✓ 989ms | http |
| 152.32.132.190:7890 | ✓ 1582ms | ✓ 1180ms | ✓ 1831ms | 否 | ✓ 1138ms | http |
| 62.113.119.14:8080 | ✓ 661ms | ✓ 1611ms | ✓ 1054ms | ✓ 1469ms | ✓ 1050ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 277ms | ✓ 955ms | ✓ 771ms | http |
| 120.92.212.16:8890 | ✓ 1227ms | 否 | ✓ 1870ms | 否 | ✓ 1409ms | http |
| 101.32.244.83:8080 | ✓ 1556ms | 否 | ✓ 1140ms | ✓ 1638ms | ✓ 1539ms | http |
| 121.43.196.210:8222 | ✓ 1092ms | ✓ 1287ms | ✓ 1009ms | ✓ 1416ms | ✓ 1105ms | http |
| 121.43.196.213:8222 | ✓ 1137ms | ✓ 1287ms | ✓ 1004ms | ✓ 1389ms | ✓ 1105ms | http |
| 121.230.8.62:1080 | 否 | 否 | ✓ 1603ms | ✓ 1772ms | ✓ 1344ms | http |
| 158.160.148.221:3128 | ✓ 589ms | 否 | ✓ 994ms | ✓ 1720ms | ✓ 1332ms | http |
| 80.92.204.47:1081 | ✓ 426ms | ✓ 1423ms | ✓ 1549ms | ✓ 1793ms | 否 | http |
| 183.232.248.73:7890 | ✓ 1114ms | 否 | 否 | ✓ 1302ms | ✓ 1065ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1423ms | ✓ 1190ms | 否 | ✓ 1466ms | http |
| 36.141.21.200:7890 | ✓ 1937ms | 否 | 否 | ✓ 1424ms | ✓ 1223ms | http |
| 46.101.95.183:8888 | 否 | 否 | ✓ 857ms | ✓ 1530ms | ✓ 1812ms | http |
| 196.70.71.34:20126 | ✓ 1181ms | ✓ 1744ms | ✓ 1408ms | ✓ 1653ms | ✓ 1364ms | http |
| 45.140.147.82:1081 | ✓ 551ms | ✓ 1890ms | ✓ 1595ms | 否 | ✓ 1487ms | http |
| 20.210.76.178:8561 | ✓ 1632ms | 否 | ✓ 1250ms | ✓ 1606ms | ✓ 1196ms | http |
| 20.210.76.175:8561 | ✓ 1635ms | ✓ 1961ms | ✓ 1287ms | ✓ 1630ms | ✓ 1180ms | http |
| 20.210.76.104:8561 | ✓ 1633ms | 否 | ✓ 1265ms | ✓ 1590ms | ✓ 1203ms | http |
| 20.18.193.135:8561 | ✓ 1639ms | 否 | ✓ 1244ms | ✓ 1608ms | ✓ 1210ms | http |
| 20.27.15.49:8561 | ✓ 1636ms | ✓ 1809ms | ✓ 1420ms | ✓ 1655ms | ✓ 1192ms | http |
| 210.223.44.230:3128 | ✓ 1607ms | 否 | ✓ 1467ms | 否 | ✓ 1647ms | http |
| 86.109.3.28:10080 | ✓ 480ms | ✓ 1164ms | ✓ 211ms | ✓ 911ms | ✓ 710ms | http |
| 200.125.171.254:999 | ✓ 849ms | ✓ 1716ms | ✓ 1316ms | ✓ 1614ms | ✓ 1315ms | http |
| 45.186.6.104:3128 | ✓ 1536ms | ✓ 1800ms | ✓ 1807ms | 否 | 否 | http |
| 103.126.238.13:8081 | ✓ 1990ms | 否 | ✓ 1458ms | ✓ 1850ms | ✓ 1511ms | http |
| 103.158.253.60:1452 | ✓ 1973ms | 否 | ✓ 1507ms | ✓ 1667ms | ✓ 1760ms | http |
| 103.158.253.61:1452 | ✓ 1969ms | 否 | ✓ 1506ms | ✓ 1688ms | ✓ 1816ms | http |
| 61.52.131.172:8443 | ✓ 1032ms | ✓ 1443ms | ✓ 1175ms | ✓ 1450ms | ✓ 1147ms | http |
| 42.200.76.16:3888 | ✓ 1091ms | 否 | ✓ 854ms | ✓ 1141ms | ✓ 842ms | http |

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
