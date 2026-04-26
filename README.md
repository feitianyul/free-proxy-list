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

最后更新：2026-04-26 09:06:34 UTC（2026-04-26 17:06:34 UTC+8）

**代理总数：47**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 703ms | ✓ 1776ms | ✓ 609ms | 否 | 否 | http |
| 218.108.131.186:17890 | ✓ 974ms | ✓ 1164ms | ✓ 939ms | ✓ 1376ms | ✓ 1144ms | http |
| 47.84.73.61:1080 | ✓ 1564ms | ✓ 1819ms | ✓ 872ms | ✓ 1172ms | ✓ 924ms | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 985ms | ✓ 1101ms | ✓ 865ms | http |
| 1.231.81.166:3128 | ✓ 1704ms | ✓ 1464ms | ✓ 1974ms | ✓ 1318ms | ✓ 1351ms | http |
| 194.31.87.77:3128 | ✓ 1984ms | 否 | ✓ 1038ms | ✓ 1853ms | ✓ 1498ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1572ms | ✓ 1538ms | ✓ 1683ms | http |
| 2.27.54.161:1080 | ✓ 996ms | 否 | ✓ 1417ms | 否 | ✓ 1746ms | http |
| 128.199.113.85:9090 | ✓ 1427ms | 否 | ✓ 858ms | ✓ 1271ms | ✓ 984ms | http |
| 120.92.212.16:8890 | ✓ 1071ms | 否 | ✓ 1795ms | ✓ 1259ms | 否 | http |
| 16.163.173.29:1080 | ✓ 1028ms | 否 | ✓ 1688ms | ✓ 1842ms | 否 | http |
| 113.160.132.26:8080 | ✓ 945ms | 否 | ✓ 1761ms | 否 | ✓ 1726ms | http |
| 47.84.59.16:1080 | ✓ 950ms | ✓ 1790ms | ✓ 798ms | ✓ 1161ms | ✓ 907ms | http |
| 128.199.121.61:9090 | ✓ 1292ms | 否 | ✓ 1159ms | ✓ 1141ms | ✓ 1157ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1451ms | ✓ 1777ms | ✓ 1473ms | http |
| 34.96.238.40:8080 | ✓ 1292ms | ✓ 1299ms | ✓ 1611ms | ✓ 1328ms | ✓ 1204ms | http |
| 207.254.71.62:8088 | ✓ 1102ms | 否 | ✓ 1489ms | ✓ 1988ms | ✓ 1895ms | http |
| 36.141.21.200:7890 | ✓ 1158ms | ✓ 1226ms | ✓ 1012ms | 否 | ✓ 1188ms | http |
| 8.219.195.129:1080 | ✓ 1491ms | ✓ 1806ms | ✓ 916ms | ✓ 1172ms | ✓ 912ms | http |
| 54.222.174.194:80 | ✓ 1716ms | ✓ 1958ms | 否 | ✓ 1883ms | 否 | http |
| 43.133.90.161:8888 | ✓ 1326ms | ✓ 1885ms | 否 | 否 | ✓ 711ms | http |
| 210.223.44.230:3128 | ✓ 1890ms | 否 | ✓ 1118ms | ✓ 1095ms | ✓ 826ms | http |
| 117.236.124.166:3128 | ✓ 1721ms | 否 | ✓ 1101ms | ✓ 1911ms | ✓ 1521ms | http |
| 150.107.140.238:3128 | ✓ 1851ms | 否 | ✓ 955ms | ✓ 1283ms | 否 | http |
| 168.144.75.9:3128 | ✓ 1466ms | 否 | ✓ 1606ms | 否 | ✓ 1733ms | http |
| 80.92.204.47:1081 | ✓ 542ms | 否 | ✓ 1354ms | 否 | ✓ 1210ms | http |
| 128.199.116.219:9090 | ✓ 810ms | 否 | ✓ 810ms | ✓ 1185ms | ✓ 950ms | http |
| 152.42.177.32:8888 | ✓ 1034ms | 否 | ✓ 1204ms | ✓ 1377ms | ✓ 1361ms | http |
| 159.89.31.62:8080 | ✓ 538ms | 否 | ✓ 1847ms | 否 | ✓ 1502ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1758ms | ✓ 1014ms | 否 | ✓ 1517ms | http |
| 130.61.174.200:1080 | ✓ 1606ms | ✓ 1541ms | 否 | ✓ 1405ms | 否 | http |
| 183.232.248.73:7890 | ✓ 1028ms | ✓ 1193ms | 否 | ✓ 1219ms | ✓ 928ms | http |
| 86.104.74.110:1081 | ✓ 1604ms | 否 | ✓ 1482ms | 否 | ✓ 1669ms | http |
| 45.76.207.177:40000 | ✓ 1468ms | 否 | ✓ 746ms | ✓ 1144ms | ✓ 1112ms | http |
| 59.46.216.131:30001 | ✓ 1193ms | 否 | ✓ 1184ms | 否 | ✓ 1129ms | http |
| 177.93.132.244:3128 | ✓ 777ms | 否 | ✓ 840ms | 否 | ✓ 1814ms | http |
| 159.223.225.118:8888 | ✓ 1316ms | 否 | ✓ 1460ms | 否 | ✓ 1331ms | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 1123ms | ✓ 1131ms | ✓ 919ms | http |
| 45.186.6.104:3128 | ✓ 1381ms | ✓ 1958ms | ✓ 1757ms | 否 | 否 | http |
| 146.190.80.158:9090 | 否 | 否 | ✓ 1056ms | ✓ 1651ms | ✓ 1587ms | http |
| 8.219.97.248:80 | ✓ 1833ms | 否 | ✓ 1622ms | ✓ 1556ms | 否 | http |
| 199.68.217.2:3128 | ✓ 1424ms | 否 | 否 | ✓ 1692ms | ✓ 1352ms | http |
| 45.140.147.82:1081 | ✓ 918ms | 否 | ✓ 1530ms | 否 | ✓ 1762ms | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1717ms | ✓ 1378ms | ✓ 1156ms | http |
| 103.231.236.236:8182 | ✓ 1820ms | 否 | ✓ 1607ms | ✓ 1534ms | 否 | http |
| 61.52.131.172:8443 | 否 | ✓ 1886ms | ✓ 1038ms | ✓ 1333ms | ✓ 1031ms | http |
| 121.230.8.136:1080 | ✓ 1439ms | 否 | ✓ 1557ms | 否 | ✓ 1505ms | http |

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
