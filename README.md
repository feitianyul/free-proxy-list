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

最后更新：2026-04-03 11:45:42 UTC（2026-04-03 19:45:42 UTC+8）

**代理总数：40**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 40 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 40 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 384ms | ✓ 1226ms | 否 | ✓ 1042ms | ✓ 1972ms | http |
| 147.161.210.140:8800 | ✓ 834ms | ✓ 1323ms | ✓ 1195ms | ✓ 1262ms | ✓ 1372ms | http |
| 159.223.71.162:8080 | ✓ 877ms | 否 | ✓ 970ms | ✓ 1201ms | ✓ 1046ms | http |
| 167.103.115.102:8800 | ✓ 1091ms | 否 | ✓ 1092ms | ✓ 1174ms | ✓ 1343ms | http |
| 95.213.217.168:52004 | ✓ 1420ms | 否 | ✓ 1607ms | 否 | ✓ 1591ms | http |
| 203.80.138.81:50000 | 否 | ✓ 1538ms | ✓ 1290ms | ✓ 1456ms | ✓ 1213ms | http |
| 113.160.132.26:8080 | ✓ 1946ms | 否 | ✓ 1734ms | ✓ 1465ms | ✓ 1142ms | http |
| 167.103.34.108:8800 | ✓ 1553ms | 否 | ✓ 1713ms | 否 | ✓ 1649ms | http |
| 45.167.124.52:8080 | ✓ 1010ms | 否 | ✓ 1352ms | 否 | ✓ 1682ms | http |
| 212.58.132.5:8888 | ✓ 1766ms | 否 | ✓ 1522ms | ✓ 1587ms | ✓ 1273ms | http |
| 167.103.144.127:8800 | 否 | 否 | ✓ 1664ms | ✓ 1726ms | ✓ 1557ms | http |
| 217.76.245.80:999 | 否 | ✓ 1513ms | ✓ 1155ms | ✓ 1336ms | ✓ 1203ms | http |
| 45.167.125.21:999 | 否 | 否 | ✓ 1288ms | ✓ 1872ms | ✓ 1501ms | http |
| 31.207.47.254:3128 | ✓ 1155ms | ✓ 1756ms | 否 | ✓ 1825ms | ✓ 1592ms | http |
| 167.103.31.122:8800 | 否 | 否 | ✓ 1406ms | ✓ 1975ms | ✓ 1612ms | http |
| 147.161.239.240:8800 | ✓ 547ms | 否 | ✓ 845ms | ✓ 1333ms | ✓ 1043ms | http |
| 101.43.127.100:8877 | ✓ 1899ms | 否 | ✓ 1123ms | ✓ 1345ms | ✓ 1352ms | http |
| 177.234.217.88:999 | ✓ 1381ms | 否 | ✓ 1798ms | 否 | ✓ 1680ms | http |
| 62.234.206.73:3128 | ✓ 1123ms | 否 | ✓ 1580ms | 否 | ✓ 1488ms | http |
| 35.225.22.61:80 | ✓ 519ms | ✓ 1010ms | ✓ 561ms | ✓ 988ms | ✓ 823ms | http |
| 46.39.105.157:8080 | ✓ 772ms | ✓ 1877ms | 否 | ✓ 1609ms | ✓ 1363ms | http |
| 164.163.42.25:10000 | ✓ 1693ms | 否 | ✓ 1238ms | 否 | ✓ 1870ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1545ms | ✓ 905ms | ✓ 1098ms | ✓ 854ms | http |
| 150.249.255.91:3128 | ✓ 1759ms | ✓ 1766ms | ✓ 1667ms | 否 | 否 | http |
| 101.32.244.83:8080 | ✓ 1090ms | 否 | ✓ 1092ms | ✓ 1626ms | ✓ 1441ms | http |
| 121.43.196.210:8222 | ✓ 1089ms | ✓ 1234ms | ✓ 1008ms | ✓ 1290ms | ✓ 1054ms | http |
| 121.43.196.213:8222 | ✓ 1150ms | ✓ 1317ms | ✓ 968ms | ✓ 1296ms | ✓ 1026ms | http |
| 114.55.226.123:10086 | ✓ 1232ms | ✓ 1635ms | ✓ 1147ms | ✓ 1503ms | ✓ 1196ms | http |
| 45.136.198.40:3128 | ✓ 739ms | ✓ 1795ms | ✓ 1890ms | 否 | ✓ 1626ms | http |
| 115.231.181.40:8128 | ✓ 1007ms | ✓ 1923ms | 否 | 否 | ✓ 1975ms | http |
| 104.248.243.244:3128 | 否 | 否 | ✓ 1724ms | ✓ 1731ms | ✓ 1401ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1946ms | ✓ 1820ms | 否 | ✓ 1520ms | http |
| 5.253.43.103:3128 | ✓ 1328ms | ✓ 1735ms | ✓ 1676ms | ✓ 1852ms | ✓ 1305ms | http |
| 64.227.76.27:1080 | ✓ 860ms | 否 | 否 | ✓ 1585ms | ✓ 1363ms | http |
| 59.46.216.131:30001 | ✓ 1241ms | ✓ 1502ms | ✓ 1338ms | 否 | 否 | http |
| 217.77.102.18:3128 | ✓ 1150ms | 否 | ✓ 1324ms | ✓ 1873ms | ✓ 1665ms | http |
| 1.231.81.166:3128 | ✓ 1424ms | ✓ 1506ms | ✓ 1681ms | 否 | ✓ 1189ms | http |
| 72.11.151.159:6005 | ✓ 1034ms | ✓ 1810ms | ✓ 783ms | 否 | ✓ 1234ms | http |
| 5.104.87.17:8051 | ✓ 1104ms | 否 | ✓ 1075ms | ✓ 1761ms | 否 | http |
| 168.110.52.228:3128 | ✓ 1782ms | 否 | 否 | ✓ 980ms | ✓ 1814ms | http |

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
